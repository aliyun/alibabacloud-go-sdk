// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWorkitemAttachmentCreateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *WorkitemAttachmentCreateResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *WorkitemAttachmentCreateResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *WorkitemAttachmentCreateResponseBody
	GetRequestId() *string
	SetSuccess(v string) *WorkitemAttachmentCreateResponseBody
	GetSuccess() *string
}

type WorkitemAttachmentCreateResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// InvalidTagGroup.IdNotFoundntraceId: 2137844496.4337.16624448853053831
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// F590C9D8-E908-5B6C-95AC-56B7E8011FFA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s WorkitemAttachmentCreateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WorkitemAttachmentCreateResponseBody) GoString() string {
	return s.String()
}

func (s *WorkitemAttachmentCreateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *WorkitemAttachmentCreateResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *WorkitemAttachmentCreateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WorkitemAttachmentCreateResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *WorkitemAttachmentCreateResponseBody) SetErrorCode(v string) *WorkitemAttachmentCreateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *WorkitemAttachmentCreateResponseBody) SetErrorMessage(v string) *WorkitemAttachmentCreateResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *WorkitemAttachmentCreateResponseBody) SetRequestId(v string) *WorkitemAttachmentCreateResponseBody {
	s.RequestId = &v
	return s
}

func (s *WorkitemAttachmentCreateResponseBody) SetSuccess(v string) *WorkitemAttachmentCreateResponseBody {
	s.Success = &v
	return s
}

func (s *WorkitemAttachmentCreateResponseBody) Validate() error {
	return dara.Validate(s)
}
