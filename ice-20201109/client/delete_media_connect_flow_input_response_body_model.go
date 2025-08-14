// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaConnectFlowInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *DeleteMediaConnectFlowInputResponseBody
	GetContent() *string
	SetDescription(v string) *DeleteMediaConnectFlowInputResponseBody
	GetDescription() *string
	SetRequestId(v string) *DeleteMediaConnectFlowInputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *DeleteMediaConnectFlowInputResponseBody
	GetRetCode() *int32
}

type DeleteMediaConnectFlowInputResponseBody struct {
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
	// C0C02296-113C-5838-8FE9-8F3A32998DDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s DeleteMediaConnectFlowInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaConnectFlowInputResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaConnectFlowInputResponseBody) GetContent() *string {
	return s.Content
}

func (s *DeleteMediaConnectFlowInputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DeleteMediaConnectFlowInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaConnectFlowInputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *DeleteMediaConnectFlowInputResponseBody) SetContent(v string) *DeleteMediaConnectFlowInputResponseBody {
	s.Content = &v
	return s
}

func (s *DeleteMediaConnectFlowInputResponseBody) SetDescription(v string) *DeleteMediaConnectFlowInputResponseBody {
	s.Description = &v
	return s
}

func (s *DeleteMediaConnectFlowInputResponseBody) SetRequestId(v string) *DeleteMediaConnectFlowInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaConnectFlowInputResponseBody) SetRetCode(v int32) *DeleteMediaConnectFlowInputResponseBody {
	s.RetCode = &v
	return s
}

func (s *DeleteMediaConnectFlowInputResponseBody) Validate() error {
	return dara.Validate(s)
}
