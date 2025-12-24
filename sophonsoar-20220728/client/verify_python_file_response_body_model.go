// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyPythonFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyPythonFileResponseBody
	GetRequestId() *string
	SetSyntax(v []*VerifyPythonFileResponseBodySyntax) *VerifyPythonFileResponseBody
	GetSyntax() []*VerifyPythonFileResponseBodySyntax
}

type VerifyPythonFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F72685FB-A6E6-5A9A-97F7-6DC1056E63CE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result. If the parameter is left empty, the syntax of the code snippet is correct.
	Syntax []*VerifyPythonFileResponseBodySyntax `json:"Syntax,omitempty" xml:"Syntax,omitempty" type:"Repeated"`
}

func (s VerifyPythonFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyPythonFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyPythonFileResponseBody) GetSyntax() []*VerifyPythonFileResponseBodySyntax {
	return s.Syntax
}

func (s *VerifyPythonFileResponseBody) SetRequestId(v string) *VerifyPythonFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyPythonFileResponseBody) SetSyntax(v []*VerifyPythonFileResponseBodySyntax) *VerifyPythonFileResponseBody {
	s.Syntax = v
	return s
}

func (s *VerifyPythonFileResponseBody) Validate() error {
	if s.Syntax != nil {
		for _, item := range s.Syntax {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type VerifyPythonFileResponseBodySyntax struct {
	// The number that indicates the end column of the error code.
	//
	// example:
	//
	// 5
	EndColumn *int32 `json:"EndColumn,omitempty" xml:"EndColumn,omitempty"`
	// The number that indicates the end line of the error code.
	//
	// example:
	//
	// 5
	EndLineNumber *int32 `json:"EndLineNumber,omitempty" xml:"EndLineNumber,omitempty"`
	// The error message for the error code.
	//
	// example:
	//
	// undefined name \\"ab\\"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The severity level of the error code. Valid values:
	//
	// 	- 4: moderate
	//
	// 	- 8: serious
	//
	// example:
	//
	// 4
	Severity *int32 `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// The number that indicates the start column of the error code.
	//
	// example:
	//
	// 2
	StartColumn *int32 `json:"StartColumn,omitempty" xml:"StartColumn,omitempty"`
	// The number that indicates the start line of the error code.
	//
	// example:
	//
	// 2
	StartLineNumber *int32 `json:"StartLineNumber,omitempty" xml:"StartLineNumber,omitempty"`
}

func (s VerifyPythonFileResponseBodySyntax) String() string {
	return dara.Prettify(s)
}

func (s VerifyPythonFileResponseBodySyntax) GoString() string {
	return s.String()
}

func (s *VerifyPythonFileResponseBodySyntax) GetEndColumn() *int32 {
	return s.EndColumn
}

func (s *VerifyPythonFileResponseBodySyntax) GetEndLineNumber() *int32 {
	return s.EndLineNumber
}

func (s *VerifyPythonFileResponseBodySyntax) GetMessage() *string {
	return s.Message
}

func (s *VerifyPythonFileResponseBodySyntax) GetSeverity() *int32 {
	return s.Severity
}

func (s *VerifyPythonFileResponseBodySyntax) GetStartColumn() *int32 {
	return s.StartColumn
}

func (s *VerifyPythonFileResponseBodySyntax) GetStartLineNumber() *int32 {
	return s.StartLineNumber
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetEndLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.EndLineNumber = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetMessage(v string) *VerifyPythonFileResponseBodySyntax {
	s.Message = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetSeverity(v int32) *VerifyPythonFileResponseBodySyntax {
	s.Severity = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartColumn(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartColumn = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) SetStartLineNumber(v int32) *VerifyPythonFileResponseBodySyntax {
	s.StartLineNumber = &v
	return s
}

func (s *VerifyPythonFileResponseBodySyntax) Validate() error {
	return dara.Validate(s)
}
