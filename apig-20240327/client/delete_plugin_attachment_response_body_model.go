// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePluginAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePluginAttachmentResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePluginAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePluginAttachmentResponseBody
	GetRequestId() *string
}

type DeletePluginAttachmentResponseBody struct {
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 76BDFFC7-0764-5168-B047-92EE0BC7FDDE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePluginAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePluginAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePluginAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePluginAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePluginAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePluginAttachmentResponseBody) SetCode(v string) *DeletePluginAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePluginAttachmentResponseBody) SetMessage(v string) *DeletePluginAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePluginAttachmentResponseBody) SetRequestId(v string) *DeletePluginAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePluginAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
