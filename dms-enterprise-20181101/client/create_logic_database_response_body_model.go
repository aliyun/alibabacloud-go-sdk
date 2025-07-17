// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogicDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateLogicDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateLogicDatabaseResponseBody
	GetErrorMessage() *string
	SetLogicDbId(v int64) *CreateLogicDatabaseResponseBody
	GetLogicDbId() *int64
	SetRequestId(v string) *CreateLogicDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateLogicDatabaseResponseBody
	GetSuccess() *bool
}

type CreateLogicDatabaseResponseBody struct {
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
	// The ID of the logical database.
	//
	// example:
	//
	// 1***
	LogicDbId *int64 `json:"LogicDbId,omitempty" xml:"LogicDbId,omitempty"`
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

func (s CreateLogicDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLogicDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLogicDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateLogicDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateLogicDatabaseResponseBody) GetLogicDbId() *int64 {
	return s.LogicDbId
}

func (s *CreateLogicDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLogicDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateLogicDatabaseResponseBody) SetErrorCode(v string) *CreateLogicDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateLogicDatabaseResponseBody) SetErrorMessage(v string) *CreateLogicDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateLogicDatabaseResponseBody) SetLogicDbId(v int64) *CreateLogicDatabaseResponseBody {
	s.LogicDbId = &v
	return s
}

func (s *CreateLogicDatabaseResponseBody) SetRequestId(v string) *CreateLogicDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLogicDatabaseResponseBody) SetSuccess(v bool) *CreateLogicDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateLogicDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
