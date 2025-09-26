// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCodeInterpreterSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCodeInterpreterSessionResult
	GetCode() *string
	SetData(v *CodeInterpreterSessionListOut) *ListCodeInterpreterSessionResult
	GetData() *CodeInterpreterSessionListOut
	SetRequestId(v string) *ListCodeInterpreterSessionResult
	GetRequestId() *string
}

type ListCodeInterpreterSessionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code      *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CodeInterpreterSessionListOut `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListCodeInterpreterSessionResult) String() string {
	return dara.Prettify(s)
}

func (s ListCodeInterpreterSessionResult) GoString() string {
	return s.String()
}

func (s *ListCodeInterpreterSessionResult) GetCode() *string {
	return s.Code
}

func (s *ListCodeInterpreterSessionResult) GetData() *CodeInterpreterSessionListOut {
	return s.Data
}

func (s *ListCodeInterpreterSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCodeInterpreterSessionResult) SetCode(v string) *ListCodeInterpreterSessionResult {
	s.Code = &v
	return s
}

func (s *ListCodeInterpreterSessionResult) SetData(v *CodeInterpreterSessionListOut) *ListCodeInterpreterSessionResult {
	s.Data = v
	return s
}

func (s *ListCodeInterpreterSessionResult) SetRequestId(v string) *ListCodeInterpreterSessionResult {
	s.RequestId = &v
	return s
}

func (s *ListCodeInterpreterSessionResult) Validate() error {
	return dara.Validate(s)
}
