// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeColumnSecLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ChangeColumnSecLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ChangeColumnSecLevelResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ChangeColumnSecLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ChangeColumnSecLevelResponseBody
	GetSuccess() *bool
}

type ChangeColumnSecLevelResponseBody struct {
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
	// 	- true: The request was successful.
	//
	// 	- false: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ChangeColumnSecLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeColumnSecLevelResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeColumnSecLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ChangeColumnSecLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ChangeColumnSecLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeColumnSecLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ChangeColumnSecLevelResponseBody) SetErrorCode(v string) *ChangeColumnSecLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ChangeColumnSecLevelResponseBody) SetErrorMessage(v string) *ChangeColumnSecLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ChangeColumnSecLevelResponseBody) SetRequestId(v string) *ChangeColumnSecLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeColumnSecLevelResponseBody) SetSuccess(v bool) *ChangeColumnSecLevelResponseBody {
	s.Success = &v
	return s
}

func (s *ChangeColumnSecLevelResponseBody) Validate() error {
	return dara.Validate(s)
}
