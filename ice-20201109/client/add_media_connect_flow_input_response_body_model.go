// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowInputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *AddMediaConnectFlowInputResponseBodyContent) *AddMediaConnectFlowInputResponseBody
	GetContent() *AddMediaConnectFlowInputResponseBodyContent
	SetDescription(v string) *AddMediaConnectFlowInputResponseBody
	GetDescription() *string
	SetRequestId(v string) *AddMediaConnectFlowInputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *AddMediaConnectFlowInputResponseBody
	GetRetCode() *int32
}

type AddMediaConnectFlowInputResponseBody struct {
	// The response body.
	Content *AddMediaConnectFlowInputResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 11357BE8-4C54-58EA-890A-5AB646EDE4B2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s AddMediaConnectFlowInputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowInputResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowInputResponseBody) GetContent() *AddMediaConnectFlowInputResponseBodyContent {
	return s.Content
}

func (s *AddMediaConnectFlowInputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *AddMediaConnectFlowInputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaConnectFlowInputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *AddMediaConnectFlowInputResponseBody) SetContent(v *AddMediaConnectFlowInputResponseBodyContent) *AddMediaConnectFlowInputResponseBody {
	s.Content = v
	return s
}

func (s *AddMediaConnectFlowInputResponseBody) SetDescription(v string) *AddMediaConnectFlowInputResponseBody {
	s.Description = &v
	return s
}

func (s *AddMediaConnectFlowInputResponseBody) SetRequestId(v string) *AddMediaConnectFlowInputResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaConnectFlowInputResponseBody) SetRetCode(v int32) *AddMediaConnectFlowInputResponseBody {
	s.RetCode = &v
	return s
}

func (s *AddMediaConnectFlowInputResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddMediaConnectFlowInputResponseBodyContent struct {
	// The source URL.
	//
	// example:
	//
	// rtmp://1.2.3.4:1935/live/AliTestInput_8666ec062190f00e263012666319a5be
	InputUrl *string `json:"InputUrl,omitempty" xml:"InputUrl,omitempty"`
}

func (s AddMediaConnectFlowInputResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowInputResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowInputResponseBodyContent) GetInputUrl() *string {
	return s.InputUrl
}

func (s *AddMediaConnectFlowInputResponseBodyContent) SetInputUrl(v string) *AddMediaConnectFlowInputResponseBodyContent {
	s.InputUrl = &v
	return s
}

func (s *AddMediaConnectFlowInputResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
