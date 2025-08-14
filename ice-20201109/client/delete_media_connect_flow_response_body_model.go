// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DeleteMediaConnectFlowResponseBody
	GetContent() *string
	SetDescription(v string) *DeleteMediaConnectFlowResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteMediaConnectFlowResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *DeleteMediaConnectFlowResponseBody
	GetRetCode() *int32
}

type DeleteMediaConnectFlowResponseBody struct {
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
	// 5AEC17BD-D80B-5F78-BE1B-F07DFA0C8622
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of `0` indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DeleteMediaConnectFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowResponseBody) GetContent() *string {
	return s.Content
}

func (s *DeleteMediaConnectFlowResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteMediaConnectFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaConnectFlowResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *DeleteMediaConnectFlowResponseBody) SetContent(v string) *DeleteMediaConnectFlowResponseBody {
	s.Content = &v
	return s
}

func (s *DeleteMediaConnectFlowResponseBody) SetDescription(v string) *DeleteMediaConnectFlowResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteMediaConnectFlowResponseBody) SetRequestId(v string) *DeleteMediaConnectFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaConnectFlowResponseBody) SetRetCode(v int32) *DeleteMediaConnectFlowResponseBody {
	s.RetCode = &v
	return s
}

func (s *DeleteMediaConnectFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
