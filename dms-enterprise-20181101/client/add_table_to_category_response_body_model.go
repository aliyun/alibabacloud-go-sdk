// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTableToCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *AddTableToCategoryResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *AddTableToCategoryResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *AddTableToCategoryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTableToCategoryResponseBody
	GetSuccess() *bool
}

type AddTableToCategoryResponseBody struct {
	// The error code that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
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

func (s AddTableToCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTableToCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *AddTableToCategoryResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *AddTableToCategoryResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *AddTableToCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTableToCategoryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTableToCategoryResponseBody) SetErrorCode(v string) *AddTableToCategoryResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *AddTableToCategoryResponseBody) SetErrorMessage(v string) *AddTableToCategoryResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *AddTableToCategoryResponseBody) SetRequestId(v string) *AddTableToCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTableToCategoryResponseBody) SetSuccess(v bool) *AddTableToCategoryResponseBody {
	s.Success = &v
	return s
}

func (s *AddTableToCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
