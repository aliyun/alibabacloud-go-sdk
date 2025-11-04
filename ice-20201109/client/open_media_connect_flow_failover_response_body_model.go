// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenMediaConnectFlowFailoverResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *OpenMediaConnectFlowFailoverResponseBody
	GetContent() *string
	SetDescription(v string) *OpenMediaConnectFlowFailoverResponseBody
	GetDescription() *string
	SetRequestId(v string) *OpenMediaConnectFlowFailoverResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *OpenMediaConnectFlowFailoverResponseBody
	GetRetCode() *int32
}

type OpenMediaConnectFlowFailoverResponseBody struct {
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
	// The request ID.
	//
	// example:
	//
	// 11357BE8-4C54-58EA-890A-5AB646EDE4B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s OpenMediaConnectFlowFailoverResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OpenMediaConnectFlowFailoverResponseBody) GoString() string {
	return s.String()
}

func (s *OpenMediaConnectFlowFailoverResponseBody) GetContent() *string {
	return s.Content
}

func (s *OpenMediaConnectFlowFailoverResponseBody) GetDescription() *string {
	return s.Description
}

func (s *OpenMediaConnectFlowFailoverResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OpenMediaConnectFlowFailoverResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *OpenMediaConnectFlowFailoverResponseBody) SetContent(v string) *OpenMediaConnectFlowFailoverResponseBody {
	s.Content = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponseBody) SetDescription(v string) *OpenMediaConnectFlowFailoverResponseBody {
	s.Description = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponseBody) SetRequestId(v string) *OpenMediaConnectFlowFailoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponseBody) SetRetCode(v int32) *OpenMediaConnectFlowFailoverResponseBody {
	s.RetCode = &v
	return s
}

func (s *OpenMediaConnectFlowFailoverResponseBody) Validate() error {
	return dara.Validate(s)
}
