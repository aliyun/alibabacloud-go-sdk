// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateMediaConnectFlowOutputResponseBody
	GetContent() *string
	SetDescription(v string) *UpdateMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *UpdateMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *UpdateMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type UpdateMediaConnectFlowOutputResponseBody struct {
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
	// D737D0BC-4CB5-55AA-8119-B540C95DFE6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s UpdateMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMediaConnectFlowOutputResponseBody) GetContent() *string {
	return s.Content
}

func (s *UpdateMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *UpdateMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *UpdateMediaConnectFlowOutputResponseBody) SetContent(v string) *UpdateMediaConnectFlowOutputResponseBody {
	s.Content = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponseBody) SetDescription(v string) *UpdateMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponseBody) SetRequestId(v string) *UpdateMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *UpdateMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *UpdateMediaConnectFlowOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
