// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSupportDBTableRecoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SupportDBTableRecoveryResponseBody
	GetCode() *string
	SetData(v string) *SupportDBTableRecoveryResponseBody
	GetData() *string
	SetErrCode(v string) *SupportDBTableRecoveryResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SupportDBTableRecoveryResponseBody
	GetErrMessage() *string
	SetMessage(v string) *SupportDBTableRecoveryResponseBody
	GetMessage() *string
	SetRequestId(v string) *SupportDBTableRecoveryResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SupportDBTableRecoveryResponseBody
	GetSuccess() *string
}

type SupportDBTableRecoveryResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SupportDBTableRecoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SupportDBTableRecoveryResponseBody) GoString() string {
	return s.String()
}

func (s *SupportDBTableRecoveryResponseBody) GetCode() *string {
	return s.Code
}

func (s *SupportDBTableRecoveryResponseBody) GetData() *string {
	return s.Data
}

func (s *SupportDBTableRecoveryResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SupportDBTableRecoveryResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SupportDBTableRecoveryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SupportDBTableRecoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SupportDBTableRecoveryResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SupportDBTableRecoveryResponseBody) SetCode(v string) *SupportDBTableRecoveryResponseBody {
	s.Code = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetData(v string) *SupportDBTableRecoveryResponseBody {
	s.Data = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetErrCode(v string) *SupportDBTableRecoveryResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetErrMessage(v string) *SupportDBTableRecoveryResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetMessage(v string) *SupportDBTableRecoveryResponseBody {
	s.Message = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetRequestId(v string) *SupportDBTableRecoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) SetSuccess(v string) *SupportDBTableRecoveryResponseBody {
	s.Success = &v
	return s
}

func (s *SupportDBTableRecoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
