// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefundPayAsYouGoOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *RefundPayAsYouGoOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *RefundPayAsYouGoOrderResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *RefundPayAsYouGoOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefundPayAsYouGoOrderResponseBody
	GetSuccess() *bool
}

type RefundPayAsYouGoOrderResponseBody struct {
	// The error code returned if the request failed.
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
	// The ID of the request.
	//
	// example:
	//
	// C51420E3-144A-4A94-B473-8662FCF4AD10
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

func (s RefundPayAsYouGoOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefundPayAsYouGoOrderResponseBody) GoString() string {
	return s.String()
}

func (s *RefundPayAsYouGoOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RefundPayAsYouGoOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *RefundPayAsYouGoOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefundPayAsYouGoOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefundPayAsYouGoOrderResponseBody) SetErrorCode(v string) *RefundPayAsYouGoOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RefundPayAsYouGoOrderResponseBody) SetErrorMessage(v string) *RefundPayAsYouGoOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RefundPayAsYouGoOrderResponseBody) SetRequestId(v string) *RefundPayAsYouGoOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefundPayAsYouGoOrderResponseBody) SetSuccess(v bool) *RefundPayAsYouGoOrderResponseBody {
	s.Success = &v
	return s
}

func (s *RefundPayAsYouGoOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
