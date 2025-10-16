// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCodeInterpreterSessionResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartCodeInterpreterSessionResult
	GetCode() *string
	SetData(v *CodeInterpreterSessionOut) *StartCodeInterpreterSessionResult
	GetData() *CodeInterpreterSessionOut
	SetRequestId(v string) *StartCodeInterpreterSessionResult
	GetRequestId() *string
}

type StartCodeInterpreterSessionResult struct {
	// SUCCESS 为成功，失败情况返回对应错误类型，比如 ERR_BAD_REQUEST ERR_VALIDATION_FAILED ERR_INTERNAL_SERVER_ERROR
	Code      *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data      *CodeInterpreterSessionOut `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string                    `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartCodeInterpreterSessionResult) String() string {
	return dara.Prettify(s)
}

func (s StartCodeInterpreterSessionResult) GoString() string {
	return s.String()
}

func (s *StartCodeInterpreterSessionResult) GetCode() *string {
	return s.Code
}

func (s *StartCodeInterpreterSessionResult) GetData() *CodeInterpreterSessionOut {
	return s.Data
}

func (s *StartCodeInterpreterSessionResult) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCodeInterpreterSessionResult) SetCode(v string) *StartCodeInterpreterSessionResult {
	s.Code = &v
	return s
}

func (s *StartCodeInterpreterSessionResult) SetData(v *CodeInterpreterSessionOut) *StartCodeInterpreterSessionResult {
	s.Data = v
	return s
}

func (s *StartCodeInterpreterSessionResult) SetRequestId(v string) *StartCodeInterpreterSessionResult {
	s.RequestId = &v
	return s
}

func (s *StartCodeInterpreterSessionResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
