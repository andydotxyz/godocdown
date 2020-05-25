package main

import (
	"fmt"
	"go/doc"
	"io"
)

func renderConstantSectionTo(writer io.Writer, list []*doc.Value) {
	for _, entry := range list {
		fmt.Fprintf(writer, "%s\n%s\n", indentCode(sourceOfNode(entry.Decl)), formatIndent(filterText(entry.Doc)))
	}
}

func renderVariableSectionTo(writer io.Writer, list []*doc.Value) {
	for _, entry := range list {
		fmt.Fprintf(writer, "%s\n%s\n", indentCode(sourceOfNode(entry.Decl)), formatIndent(filterText(entry.Doc)))
	}
}

func renderFunctionSectionTo(writer io.Writer, list []*doc.Func, inTypeSection bool) {

	header := RenderStyle.FunctionHeader
	if inTypeSection {
		header = RenderStyle.TypeFunctionHeader
	}

	for _, entry := range list {
		receiver := " "
		if entry.Recv != "" {
			receiver = fmt.Sprintf("(%s) ", entry.Recv)
		}
		fmt.Fprintf(writer, "%s func %s%s\n\n%s\n%s\n", header, receiver, entry.Name, indentCode(sourceOfNode(entry.Decl)), formatIndent(filterText(entry.Doc)))
	}
}

func renderTypeTo(writer io.Writer, entry *doc.Type) {
	header := RenderStyle.TypeHeader

	fmt.Fprintf(writer, "%s type %s\n\n%s\n\n%s\n", header, entry.Name, indentCode(sourceOfNode(entry.Decl)), formatIndent(filterText(entry.Doc)))
	renderConstantSectionTo(writer, entry.Consts)
	renderVariableSectionTo(writer, entry.Vars)
	renderFunctionSectionTo(writer, entry.Funcs, true)
	renderFunctionSectionTo(writer, entry.Methods, true)
}

func renderTypeListTo(writer io.Writer, list []*doc.Type) {
	fmt.Fprintf(writer, "%s types\n\n",  RenderStyle.TypeHeader) // TODO
	for _, t := range list {
		fmt.Fprintf(writer, " * [%s](%s)\n", t.Name, typeFileName(t)+".html")
	}
}

func renderHeaderTo(writer io.Writer, document *_document) {
	fmt.Fprintf(writer, "# %s\n--\n", document.Name)

	if !document.IsCommand {
		// Import
		if RenderStyle.IncludeImport {
			if document.ImportPath != "" {
				fmt.Fprintf(writer, spacer(4)+"import \"%s\"\n\n", document.ImportPath)
			}
		}
	}
}

func renderSynopsisTo(writer io.Writer, document *_document) {
	fmt.Fprintf(writer, "%s\n", headifySynopsis(formatIndent(filterText(document.pkg.Doc))))
}

func renderUsageTo(writer io.Writer, document *_document) {
	// Usage
	fmt.Fprintf(writer, "%s\n", RenderStyle.UsageHeader)

	if document.Type != nil {
		return
	}

	// Constant Section
	renderConstantSectionTo(writer, document.pkg.Consts)

	// Variable Section
	renderVariableSectionTo(writer, document.pkg.Vars)

	// Function Section
	renderFunctionSectionTo(writer, document.pkg.Funcs, false)
}

func renderSignatureTo(writer io.Writer) {
	if RenderStyle.IncludeSignature {
		fmt.Fprintf(writer, "\n\n--\n**godocdown** http://github.com/robertkrimen/godocdown\n")
	}
}
