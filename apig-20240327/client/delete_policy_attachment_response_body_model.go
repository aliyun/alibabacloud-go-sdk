// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolicyAttachmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePolicyAttachmentResponseBody
	GetCode() *string
	SetMessage(v string) *DeletePolicyAttachmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePolicyAttachmentResponseBody
	GetRequestId() *string
}

type DeletePolicyAttachmentResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// ID of the request
	//
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePolicyAttachmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePolicyAttachmentResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePolicyAttachmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePolicyAttachmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePolicyAttachmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePolicyAttachmentResponseBody) SetCode(v string) *DeletePolicyAttachmentResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePolicyAttachmentResponseBody) SetMessage(v string) *DeletePolicyAttachmentResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePolicyAttachmentResponseBody) SetRequestId(v string) *DeletePolicyAttachmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePolicyAttachmentResponseBody) Validate() error {
	return dara.Validate(s)
}
