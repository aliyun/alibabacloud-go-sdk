// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseMediaConnectFlowFailoverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *CloseMediaConnectFlowFailoverResponseBody
	GetContent() *string
	SetDescription(v string) *CloseMediaConnectFlowFailoverResponseBody
	GetDescription() *string
	SetRequestId(v string) *CloseMediaConnectFlowFailoverResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *CloseMediaConnectFlowFailoverResponseBody
	GetRetCode() *int32
}

type CloseMediaConnectFlowFailoverResponseBody struct {
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
	// 391DDF25-705C-5B38-9DB9-7A6B00D6065A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s CloseMediaConnectFlowFailoverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseMediaConnectFlowFailoverResponseBody) GoString() string {
	return s.String()
}

func (s *CloseMediaConnectFlowFailoverResponseBody) GetContent() *string {
	return s.Content
}

func (s *CloseMediaConnectFlowFailoverResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CloseMediaConnectFlowFailoverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseMediaConnectFlowFailoverResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *CloseMediaConnectFlowFailoverResponseBody) SetContent(v string) *CloseMediaConnectFlowFailoverResponseBody {
	s.Content = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponseBody) SetDescription(v string) *CloseMediaConnectFlowFailoverResponseBody {
	s.Description = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponseBody) SetRequestId(v string) *CloseMediaConnectFlowFailoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponseBody) SetRetCode(v int32) *CloseMediaConnectFlowFailoverResponseBody {
	s.RetCode = &v
	return s
}

func (s *CloseMediaConnectFlowFailoverResponseBody) Validate() error {
	return dara.Validate(s)
}
