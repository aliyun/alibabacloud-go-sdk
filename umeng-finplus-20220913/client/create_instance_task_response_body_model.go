// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateInstanceTaskResponseBody
	GetData() *string
	SetErrCode(v int32) *CreateInstanceTaskResponseBody
	GetErrCode() *int32
	SetErrMessage(v string) *CreateInstanceTaskResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateInstanceTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateInstanceTaskResponseBody
	GetSuccess() *bool
}

type CreateInstanceTaskResponseBody struct {
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *int32  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateInstanceTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateInstanceTaskResponseBody) GetErrCode() *int32 {
	return s.ErrCode
}

func (s *CreateInstanceTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateInstanceTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateInstanceTaskResponseBody) SetData(v string) *CreateInstanceTaskResponseBody {
	s.Data = &v
	return s
}

func (s *CreateInstanceTaskResponseBody) SetErrCode(v int32) *CreateInstanceTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateInstanceTaskResponseBody) SetErrMessage(v string) *CreateInstanceTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateInstanceTaskResponseBody) SetRequestId(v string) *CreateInstanceTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceTaskResponseBody) SetSuccess(v bool) *CreateInstanceTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateInstanceTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
