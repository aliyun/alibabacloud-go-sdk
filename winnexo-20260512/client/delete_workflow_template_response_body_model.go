// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkflowTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkflowTemplateResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteWorkflowTemplateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkflowTemplateResponseBody
	GetRequestId() *string
}

type DeleteWorkflowTemplateResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request tracking ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkflowTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkflowTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkflowTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkflowTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkflowTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkflowTemplateResponseBody) SetCode(v string) *DeleteWorkflowTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkflowTemplateResponseBody) SetMessage(v string) *DeleteWorkflowTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkflowTemplateResponseBody) SetRequestId(v string) *DeleteWorkflowTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkflowTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
