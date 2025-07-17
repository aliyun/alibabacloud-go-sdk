// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataLakeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteDataLakeDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteDataLakeDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteDataLakeDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataLakeDatabaseResponseBody
	GetSuccess() *bool
}

type DeleteDataLakeDatabaseResponseBody struct {
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
	// C51420E3-144A-4A94-B473-8662FCF4AD10
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataLakeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataLakeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataLakeDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteDataLakeDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteDataLakeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataLakeDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataLakeDatabaseResponseBody) SetErrorCode(v string) *DeleteDataLakeDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteDataLakeDatabaseResponseBody) SetErrorMessage(v string) *DeleteDataLakeDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteDataLakeDatabaseResponseBody) SetRequestId(v string) *DeleteDataLakeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataLakeDatabaseResponseBody) SetSuccess(v bool) *DeleteDataLakeDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataLakeDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
