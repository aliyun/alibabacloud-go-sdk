// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDataLakeDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataLakeDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataLakeDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataLakeDatabaseResponseBody
	GetSuccess() *bool
}

type CreateDataLakeDatabaseResponseBody struct {
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// D911009F-3E95-5AFD-8CF1-73F7B4F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataLakeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLakeDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataLakeDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataLakeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataLakeDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataLakeDatabaseResponseBody) SetErrorCode(v string) *CreateDataLakeDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataLakeDatabaseResponseBody) SetErrorMessage(v string) *CreateDataLakeDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataLakeDatabaseResponseBody) SetRequestId(v string) *CreateDataLakeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataLakeDatabaseResponseBody) SetSuccess(v bool) *CreateDataLakeDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataLakeDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
