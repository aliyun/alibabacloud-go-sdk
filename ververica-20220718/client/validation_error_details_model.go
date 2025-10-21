// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidationErrorDetails interface {
	dara.Model
	String() string
	GoString() string
	SetColumnNumber(v string) *ValidationErrorDetails
	GetColumnNumber() *string
	SetEndColumnNumber(v string) *ValidationErrorDetails
	GetEndColumnNumber() *string
	SetEndLineNumber(v string) *ValidationErrorDetails
	GetEndLineNumber() *string
	SetLineNumber(v string) *ValidationErrorDetails
	GetLineNumber() *string
	SetMessage(v string) *ValidationErrorDetails
	GetMessage() *string
}

type ValidationErrorDetails struct {
	ColumnNumber    *string `json:"columnNumber,omitempty" xml:"columnNumber,omitempty"`
	EndColumnNumber *string `json:"endColumnNumber,omitempty" xml:"endColumnNumber,omitempty"`
	EndLineNumber   *string `json:"endLineNumber,omitempty" xml:"endLineNumber,omitempty"`
	LineNumber      *string `json:"lineNumber,omitempty" xml:"lineNumber,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidationErrorDetails) String() string {
	return dara.Prettify(s)
}

func (s ValidationErrorDetails) GoString() string {
	return s.String()
}

func (s *ValidationErrorDetails) GetColumnNumber() *string {
	return s.ColumnNumber
}

func (s *ValidationErrorDetails) GetEndColumnNumber() *string {
	return s.EndColumnNumber
}

func (s *ValidationErrorDetails) GetEndLineNumber() *string {
	return s.EndLineNumber
}

func (s *ValidationErrorDetails) GetLineNumber() *string {
	return s.LineNumber
}

func (s *ValidationErrorDetails) GetMessage() *string {
	return s.Message
}

func (s *ValidationErrorDetails) SetColumnNumber(v string) *ValidationErrorDetails {
	s.ColumnNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetEndColumnNumber(v string) *ValidationErrorDetails {
	s.EndColumnNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetEndLineNumber(v string) *ValidationErrorDetails {
	s.EndLineNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetLineNumber(v string) *ValidationErrorDetails {
	s.LineNumber = &v
	return s
}

func (s *ValidationErrorDetails) SetMessage(v string) *ValidationErrorDetails {
	s.Message = &v
	return s
}

func (s *ValidationErrorDetails) Validate() error {
	return dara.Validate(s)
}
