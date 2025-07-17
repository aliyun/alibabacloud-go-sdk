// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogicDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteLogicDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteLogicDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteLogicDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteLogicDatabaseResponseBody
	GetSuccess() *bool
}

type DeleteLogicDatabaseResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
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
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteLogicDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogicDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogicDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteLogicDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteLogicDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogicDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteLogicDatabaseResponseBody) SetErrorCode(v string) *DeleteLogicDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteLogicDatabaseResponseBody) SetErrorMessage(v string) *DeleteLogicDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteLogicDatabaseResponseBody) SetRequestId(v string) *DeleteLogicDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogicDatabaseResponseBody) SetSuccess(v bool) *DeleteLogicDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteLogicDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
