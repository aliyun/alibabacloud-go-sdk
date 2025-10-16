// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeInterpreterSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCodeInterpreterSessionResult
	GetCode() *string
	SetData(v *CodeInterpreterSessionOut) *GetCodeInterpreterSessionResult
	GetData() *CodeInterpreterSessionOut
	SetRequestId(v string) *GetCodeInterpreterSessionResult
	GetRequestId() *string
}

type GetCodeInterpreterSessionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code      *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CodeInterpreterSessionOut `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetCodeInterpreterSessionResult) String() string {
	return dara.Prettify(s)
}

func (s GetCodeInterpreterSessionResult) GoString() string {
	return s.String()
}

func (s *GetCodeInterpreterSessionResult) GetCode() *string {
	return s.Code
}

func (s *GetCodeInterpreterSessionResult) GetData() *CodeInterpreterSessionOut {
	return s.Data
}

func (s *GetCodeInterpreterSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCodeInterpreterSessionResult) SetCode(v string) *GetCodeInterpreterSessionResult {
	s.Code = &v
	return s
}

func (s *GetCodeInterpreterSessionResult) SetData(v *CodeInterpreterSessionOut) *GetCodeInterpreterSessionResult {
	s.Data = v
	return s
}

func (s *GetCodeInterpreterSessionResult) SetRequestId(v string) *GetCodeInterpreterSessionResult {
	s.RequestId = &v
	return s
}

func (s *GetCodeInterpreterSessionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
