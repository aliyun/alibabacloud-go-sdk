// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DeleteMediaConnectFlowOutputResponseBody
	GetContent() *string
	SetDescription(v string) *DeleteMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *DeleteMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type DeleteMediaConnectFlowOutputResponseBody struct {
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
	// DF73E08E-F807-50F5-A2BD-B76391EAE8FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DeleteMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowOutputResponseBody) GetContent() *string {
	return s.Content
}

func (s *DeleteMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *DeleteMediaConnectFlowOutputResponseBody) SetContent(v string) *DeleteMediaConnectFlowOutputResponseBody {
	s.Content = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponseBody) SetDescription(v string) *DeleteMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponseBody) SetRequestId(v string) *DeleteMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *DeleteMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *DeleteMediaConnectFlowOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
