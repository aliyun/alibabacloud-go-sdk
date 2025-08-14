// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateMediaConnectFlowInputResponseBody
	GetContent() *string
	SetDescription(v string) *UpdateMediaConnectFlowInputResponseBody
	GetDescription() *string
	SetRequestId(v string) *UpdateMediaConnectFlowInputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *UpdateMediaConnectFlowInputResponseBody
	GetRetCode() *int32
}

type UpdateMediaConnectFlowInputResponseBody struct {
	// The response body.
	//
	// example:
	//
	// ""
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The call description.
	//
	// example:
	//
	// OK
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 52451256-FFEA-5D2E-AA60-EE7053000F22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s UpdateMediaConnectFlowInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowInputResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowInputResponseBody) GetContent() *string {
	return s.Content
}

func (s *UpdateMediaConnectFlowInputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaConnectFlowInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaConnectFlowInputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *UpdateMediaConnectFlowInputResponseBody) SetContent(v string) *UpdateMediaConnectFlowInputResponseBody {
	s.Content = &v
	return s
}

func (s *UpdateMediaConnectFlowInputResponseBody) SetDescription(v string) *UpdateMediaConnectFlowInputResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateMediaConnectFlowInputResponseBody) SetRequestId(v string) *UpdateMediaConnectFlowInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaConnectFlowInputResponseBody) SetRetCode(v int32) *UpdateMediaConnectFlowInputResponseBody {
	s.RetCode = &v
	return s
}

func (s *UpdateMediaConnectFlowInputResponseBody) Validate() error {
	return dara.Validate(s)
}
