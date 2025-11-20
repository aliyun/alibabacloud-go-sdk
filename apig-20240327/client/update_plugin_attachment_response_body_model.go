// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdatePluginAttachmentResponseBody
	GetCode() *string
	SetMessage(v string) *UpdatePluginAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePluginAttachmentResponseBody
	GetRequestId() *string
}

type UpdatePluginAttachmentResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The status message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F330090D-80F8-557B-8610-7EC7E386B4A4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdatePluginAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePluginAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdatePluginAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePluginAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePluginAttachmentResponseBody) SetCode(v string) *UpdatePluginAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePluginAttachmentResponseBody) SetMessage(v string) *UpdatePluginAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePluginAttachmentResponseBody) SetRequestId(v string) *UpdatePluginAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePluginAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
