// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConnectFlowAllOutputNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*string) *GetMediaConnectFlowAllOutputNameResponseBody
	GetContent() []*string
	SetDescription(v string) *GetMediaConnectFlowAllOutputNameResponseBody
	GetDescription() *string
	SetRequestId(v string) *GetMediaConnectFlowAllOutputNameResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *GetMediaConnectFlowAllOutputNameResponseBody
	GetRetCode() *int32
}

type GetMediaConnectFlowAllOutputNameResponseBody struct {
	// The response body, as an array of strings.
	Content []*string `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
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
	// 559E9828-245D-5CBA-9C7A-4E01453F091F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s GetMediaConnectFlowAllOutputNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConnectFlowAllOutputNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) GetContent() []*string {
	return s.Content
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) SetContent(v []*string) *GetMediaConnectFlowAllOutputNameResponseBody {
	s.Content = v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) SetDescription(v string) *GetMediaConnectFlowAllOutputNameResponseBody {
	s.Description = &v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) SetRequestId(v string) *GetMediaConnectFlowAllOutputNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) SetRetCode(v int32) *GetMediaConnectFlowAllOutputNameResponseBody {
	s.RetCode = &v
	return s
}

func (s *GetMediaConnectFlowAllOutputNameResponseBody) Validate() error {
	return dara.Validate(s)
}
