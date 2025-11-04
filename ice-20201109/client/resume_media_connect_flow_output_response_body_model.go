// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeMediaConnectFlowOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *ResumeMediaConnectFlowOutputResponseBody
	GetContent() *string
	SetDescription(v string) *ResumeMediaConnectFlowOutputResponseBody
	GetDescription() *string
	SetRequestId(v string) *ResumeMediaConnectFlowOutputResponseBody
	GetRequestId() *string
	SetRetCode(v int32) *ResumeMediaConnectFlowOutputResponseBody
	GetRetCode() *int32
}

type ResumeMediaConnectFlowOutputResponseBody struct {
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
	// 31E30781-9495-5E2D-A84D-759B0A01E262
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned error code. A value of 0 indicates the call is successful.
	//
	// example:
	//
	// 0
	RetCode *int32 `json:"RetCode,omitempty" xml:"RetCode,omitempty"`
}

func (s ResumeMediaConnectFlowOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResumeMediaConnectFlowOutputResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeMediaConnectFlowOutputResponseBody) GetContent() *string {
	return s.Content
}

func (s *ResumeMediaConnectFlowOutputResponseBody) GetDescription() *string {
	return s.Description
}

func (s *ResumeMediaConnectFlowOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResumeMediaConnectFlowOutputResponseBody) GetRetCode() *int32 {
	return s.RetCode
}

func (s *ResumeMediaConnectFlowOutputResponseBody) SetContent(v string) *ResumeMediaConnectFlowOutputResponseBody {
	s.Content = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponseBody) SetDescription(v string) *ResumeMediaConnectFlowOutputResponseBody {
	s.Description = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponseBody) SetRequestId(v string) *ResumeMediaConnectFlowOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponseBody) SetRetCode(v int32) *ResumeMediaConnectFlowOutputResponseBody {
	s.RetCode = &v
	return s
}

func (s *ResumeMediaConnectFlowOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
