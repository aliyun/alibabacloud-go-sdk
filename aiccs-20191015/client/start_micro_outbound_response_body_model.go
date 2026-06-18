// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMicroOutboundResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartMicroOutboundResponseBody
	GetCode() *string
	SetCustomerInfo(v string) *StartMicroOutboundResponseBody
	GetCustomerInfo() *string
	SetInvokeCmdId(v string) *StartMicroOutboundResponseBody
	GetInvokeCmdId() *string
	SetInvokeCreateTime(v string) *StartMicroOutboundResponseBody
	GetInvokeCreateTime() *string
	SetMessage(v string) *StartMicroOutboundResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartMicroOutboundResponseBody
	GetRequestId() *string
}

type StartMicroOutboundResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Customer information.
	//
	// example:
	//
	// {\\"caseId\\":2323****}
	CustomerInfo *string `json:"CustomerInfo,omitempty" xml:"CustomerInfo,omitempty"`
	// Command ID.
	//
	// example:
	//
	// 8883f165-4a0d-4da2-a2d2
	InvokeCmdId *string `json:"InvokeCmdId,omitempty" xml:"InvokeCmdId,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2019-05-23 17:30:32.525
	InvokeCreateTime *string `json:"InvokeCreateTime,omitempty" xml:"InvokeCreateTime,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartMicroOutboundResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMicroOutboundResponseBody) GoString() string {
	return s.String()
}

func (s *StartMicroOutboundResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartMicroOutboundResponseBody) GetCustomerInfo() *string {
	return s.CustomerInfo
}

func (s *StartMicroOutboundResponseBody) GetInvokeCmdId() *string {
	return s.InvokeCmdId
}

func (s *StartMicroOutboundResponseBody) GetInvokeCreateTime() *string {
	return s.InvokeCreateTime
}

func (s *StartMicroOutboundResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartMicroOutboundResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMicroOutboundResponseBody) SetCode(v string) *StartMicroOutboundResponseBody {
	s.Code = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetCustomerInfo(v string) *StartMicroOutboundResponseBody {
	s.CustomerInfo = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetInvokeCmdId(v string) *StartMicroOutboundResponseBody {
	s.InvokeCmdId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetInvokeCreateTime(v string) *StartMicroOutboundResponseBody {
	s.InvokeCreateTime = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetMessage(v string) *StartMicroOutboundResponseBody {
	s.Message = &v
	return s
}

func (s *StartMicroOutboundResponseBody) SetRequestId(v string) *StartMicroOutboundResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMicroOutboundResponseBody) Validate() error {
	return dara.Validate(s)
}
