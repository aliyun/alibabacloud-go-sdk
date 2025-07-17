// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrderAttachmentFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetOrderAttachmentFileResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetOrderAttachmentFileResponseBody
	GetErrorMessage() *string
	SetFileUrl(v string) *GetOrderAttachmentFileResponseBody
	GetFileUrl() *string
	SetRequestId(v string) *GetOrderAttachmentFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetOrderAttachmentFileResponseBody
	GetSuccess() *bool
}

type GetOrderAttachmentFileResponseBody struct {
	// The error code returned. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request is successful.
	//
	// 	- The **ErrorCode*	- parameter is returned if the request fails. For more information, see the **Error codes*	- section of this topic.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The download URL of the attachment.
	//
	// example:
	//
	// https://dmsxxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// FE8EE2F1-4880-46BC-A704-5CF63EAF9A04
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetOrderAttachmentFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOrderAttachmentFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetOrderAttachmentFileResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetOrderAttachmentFileResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetOrderAttachmentFileResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetOrderAttachmentFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOrderAttachmentFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetOrderAttachmentFileResponseBody) SetErrorCode(v string) *GetOrderAttachmentFileResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetOrderAttachmentFileResponseBody) SetErrorMessage(v string) *GetOrderAttachmentFileResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetOrderAttachmentFileResponseBody) SetFileUrl(v string) *GetOrderAttachmentFileResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetOrderAttachmentFileResponseBody) SetRequestId(v string) *GetOrderAttachmentFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOrderAttachmentFileResponseBody) SetSuccess(v bool) *GetOrderAttachmentFileResponseBody {
	s.Success = &v
	return s
}

func (s *GetOrderAttachmentFileResponseBody) Validate() error {
	return dara.Validate(s)
}
