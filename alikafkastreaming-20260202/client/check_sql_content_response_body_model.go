// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CheckSqlContentResponseBody
	GetCode() *int64
	SetData(v *CheckSqlContentResponseBodyData) *CheckSqlContentResponseBody
	GetData() *CheckSqlContentResponseBodyData
	SetRequestId(v string) *CheckSqlContentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckSqlContentResponseBody
	GetSuccess() *bool
}

type CheckSqlContentResponseBody struct {
	Code      *int64                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CheckSqlContentResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckSqlContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlContentResponseBody) GoString() string {
	return s.String()
}

func (s *CheckSqlContentResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CheckSqlContentResponseBody) GetData() *CheckSqlContentResponseBodyData {
	return s.Data
}

func (s *CheckSqlContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckSqlContentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckSqlContentResponseBody) SetCode(v int64) *CheckSqlContentResponseBody {
	s.Code = &v
	return s
}

func (s *CheckSqlContentResponseBody) SetData(v *CheckSqlContentResponseBodyData) *CheckSqlContentResponseBody {
	s.Data = v
	return s
}

func (s *CheckSqlContentResponseBody) SetRequestId(v string) *CheckSqlContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckSqlContentResponseBody) SetSuccess(v bool) *CheckSqlContentResponseBody {
	s.Success = &v
	return s
}

func (s *CheckSqlContentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckSqlContentResponseBodyData struct {
	ErrorList []*CheckSqlContentResponseBodyDataErrorList `json:"ErrorList,omitempty" xml:"ErrorList,omitempty" type:"Repeated"`
	Valid     *bool                                       `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s CheckSqlContentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlContentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckSqlContentResponseBodyData) GetErrorList() []*CheckSqlContentResponseBodyDataErrorList {
	return s.ErrorList
}

func (s *CheckSqlContentResponseBodyData) GetValid() *bool {
	return s.Valid
}

func (s *CheckSqlContentResponseBodyData) SetErrorList(v []*CheckSqlContentResponseBodyDataErrorList) *CheckSqlContentResponseBodyData {
	s.ErrorList = v
	return s
}

func (s *CheckSqlContentResponseBodyData) SetValid(v bool) *CheckSqlContentResponseBodyData {
	s.Valid = &v
	return s
}

func (s *CheckSqlContentResponseBodyData) Validate() error {
	if s.ErrorList != nil {
		for _, item := range s.ErrorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckSqlContentResponseBodyDataErrorList struct {
	CodeSnippet     *string `json:"CodeSnippet,omitempty" xml:"CodeSnippet,omitempty"`
	ColumnNumber    *int32  `json:"ColumnNumber,omitempty" xml:"ColumnNumber,omitempty"`
	EndColumnNumber *int32  `json:"EndColumnNumber,omitempty" xml:"EndColumnNumber,omitempty"`
	EndLineNumber   *int32  `json:"EndLineNumber,omitempty" xml:"EndLineNumber,omitempty"`
	ErrorType       *string `json:"ErrorType,omitempty" xml:"ErrorType,omitempty"`
	LineNumber      *int32  `json:"LineNumber,omitempty" xml:"LineNumber,omitempty"`
	Message         *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s CheckSqlContentResponseBodyDataErrorList) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlContentResponseBodyDataErrorList) GoString() string {
	return s.String()
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetCodeSnippet() *string {
	return s.CodeSnippet
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetColumnNumber() *int32 {
	return s.ColumnNumber
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetEndColumnNumber() *int32 {
	return s.EndColumnNumber
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetEndLineNumber() *int32 {
	return s.EndLineNumber
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetErrorType() *string {
	return s.ErrorType
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetLineNumber() *int32 {
	return s.LineNumber
}

func (s *CheckSqlContentResponseBodyDataErrorList) GetMessage() *string {
	return s.Message
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetCodeSnippet(v string) *CheckSqlContentResponseBodyDataErrorList {
	s.CodeSnippet = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetColumnNumber(v int32) *CheckSqlContentResponseBodyDataErrorList {
	s.ColumnNumber = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetEndColumnNumber(v int32) *CheckSqlContentResponseBodyDataErrorList {
	s.EndColumnNumber = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetEndLineNumber(v int32) *CheckSqlContentResponseBodyDataErrorList {
	s.EndLineNumber = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetErrorType(v string) *CheckSqlContentResponseBodyDataErrorList {
	s.ErrorType = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetLineNumber(v int32) *CheckSqlContentResponseBodyDataErrorList {
	s.LineNumber = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) SetMessage(v string) *CheckSqlContentResponseBodyDataErrorList {
	s.Message = &v
	return s
}

func (s *CheckSqlContentResponseBodyDataErrorList) Validate() error {
	return dara.Validate(s)
}
