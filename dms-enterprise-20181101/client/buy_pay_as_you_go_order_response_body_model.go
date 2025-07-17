// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuyPayAsYouGoOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *BuyPayAsYouGoOrderResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *BuyPayAsYouGoOrderResponseBody
	GetErrorMessage() *string
	SetInstanceId(v string) *BuyPayAsYouGoOrderResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *BuyPayAsYouGoOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuyPayAsYouGoOrderResponseBody
	GetSuccess() *bool
}

type BuyPayAsYouGoOrderResponseBody struct {
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
	// If the purchase is successful, the ID of the purchased instance is returned.
	//
	// example:
	//
	// dms_pre_public_cn-nif23l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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

func (s BuyPayAsYouGoOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuyPayAsYouGoOrderResponseBody) GoString() string {
	return s.String()
}

func (s *BuyPayAsYouGoOrderResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *BuyPayAsYouGoOrderResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *BuyPayAsYouGoOrderResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BuyPayAsYouGoOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuyPayAsYouGoOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuyPayAsYouGoOrderResponseBody) SetErrorCode(v string) *BuyPayAsYouGoOrderResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponseBody) SetErrorMessage(v string) *BuyPayAsYouGoOrderResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponseBody) SetInstanceId(v string) *BuyPayAsYouGoOrderResponseBody {
	s.InstanceId = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponseBody) SetRequestId(v string) *BuyPayAsYouGoOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponseBody) SetSuccess(v bool) *BuyPayAsYouGoOrderResponseBody {
	s.Success = &v
	return s
}

func (s *BuyPayAsYouGoOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
