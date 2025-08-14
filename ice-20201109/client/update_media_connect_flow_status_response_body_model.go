// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateMediaConnectFlowStatusResponseBody
	GetContent() *string
	SetDescription(v string) *UpdateMediaConnectFlowStatusResponseBody
	GetDescription() *string
	SetRequestId(v string) *UpdateMediaConnectFlowStatusResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *UpdateMediaConnectFlowStatusResponseBody
	GetRetCode() *int32
}

type UpdateMediaConnectFlowStatusResponseBody struct {
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
	// ok
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 20B3A1B6-4BD2-5DE6-BCBC-098C9B4F4E91
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates that the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s UpdateMediaConnectFlowStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowStatusResponseBody) GetContent() *string {
	return s.Content
}

func (s *UpdateMediaConnectFlowStatusResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaConnectFlowStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaConnectFlowStatusResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *UpdateMediaConnectFlowStatusResponseBody) SetContent(v string) *UpdateMediaConnectFlowStatusResponseBody {
	s.Content = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponseBody) SetDescription(v string) *UpdateMediaConnectFlowStatusResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponseBody) SetRequestId(v string) *UpdateMediaConnectFlowStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponseBody) SetRetCode(v int32) *UpdateMediaConnectFlowStatusResponseBody {
	s.RetCode = &v
	return s
}

func (s *UpdateMediaConnectFlowStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
