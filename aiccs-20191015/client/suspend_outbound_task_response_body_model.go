// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendOutboundTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SuspendOutboundTaskResponseBody
	GetCode() *string
	SetData(v string) *SuspendOutboundTaskResponseBody
	GetData() *string
	SetMessage(v string) *SuspendOutboundTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuspendOutboundTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SuspendOutboundTaskResponseBody
	GetSuccess() *bool
}

type SuspendOutboundTaskResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendOutboundTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendOutboundTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendOutboundTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SuspendOutboundTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *SuspendOutboundTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuspendOutboundTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendOutboundTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SuspendOutboundTaskResponseBody) SetCode(v string) *SuspendOutboundTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetData(v string) *SuspendOutboundTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetMessage(v string) *SuspendOutboundTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetRequestId(v string) *SuspendOutboundTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) SetSuccess(v bool) *SuspendOutboundTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendOutboundTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
