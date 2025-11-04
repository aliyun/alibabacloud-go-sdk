// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *AddMediaConnectFlowOutputResponseBodyContent) *AddMediaConnectFlowOutputResponseBody
	GetContent() *AddMediaConnectFlowOutputResponseBodyContent
	SetDescription(v string) *AddMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *AddMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *AddMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type AddMediaConnectFlowOutputResponseBody struct {
	// The response body.
	Content *AddMediaConnectFlowOutputResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
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
	// 11AA9E73-FBA0-58DC-97BA-D606D847BCB6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates that the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s AddMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowOutputResponseBody) GetContent() *AddMediaConnectFlowOutputResponseBodyContent {
	return s.Content
}

func (s *AddMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *AddMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *AddMediaConnectFlowOutputResponseBody) SetContent(v *AddMediaConnectFlowOutputResponseBodyContent) *AddMediaConnectFlowOutputResponseBody {
	s.Content = v
	return s
}

func (s *AddMediaConnectFlowOutputResponseBody) SetDescription(v string) *AddMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *AddMediaConnectFlowOutputResponseBody) SetRequestId(v string) *AddMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *AddMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *AddMediaConnectFlowOutputResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddMediaConnectFlowOutputResponseBodyContent struct {
	// The output URL.
	//
	// example:
	//
	// srt://1.2.3.4:1025
	OutputUrl *string `json:"OutputUrl,omitempty" xml:"OutputUrl,omitempty"`
}

func (s AddMediaConnectFlowOutputResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s AddMediaConnectFlowOutputResponseBodyContent) GoString() string {
	return s.String()
}

func (s *AddMediaConnectFlowOutputResponseBodyContent) GetOutputUrl() *string {
	return s.OutputUrl
}

func (s *AddMediaConnectFlowOutputResponseBodyContent) SetOutputUrl(v string) *AddMediaConnectFlowOutputResponseBodyContent {
	s.OutputUrl = &v
	return s
}

func (s *AddMediaConnectFlowOutputResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
