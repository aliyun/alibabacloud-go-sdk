// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RestartOutboundTaskResponseBody
	GetCode() *string
	SetData(v string) *RestartOutboundTaskResponseBody
	GetData() *string
	SetMessage(v string) *RestartOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartOutboundTaskResponseBody
	GetSuccess() *bool
}

type RestartOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RestartOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *RestartOutboundTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *RestartOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartOutboundTaskResponseBody) SetCode(v string) *RestartOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetData(v string) *RestartOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetMessage(v string) *RestartOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetRequestId(v string) *RestartOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) SetSuccess(v bool) *RestartOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RestartOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
