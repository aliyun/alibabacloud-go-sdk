// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataLakeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *DLDatabase) *UpdateDataLakeDatabaseResponseBody
	GetDatabase() *DLDatabase
	SetErrorCode(v string) *UpdateDataLakeDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateDataLakeDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *UpdateDataLakeDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataLakeDatabaseResponseBody
	GetSuccess() *bool
}

type UpdateDataLakeDatabaseResponseBody struct {
	Database *DLDatabase `json:"Database,omitempty" xml:"Database,omitempty"`
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
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataLakeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataLakeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataLakeDatabaseResponseBody) GetDatabase() *DLDatabase {
	return s.Database
}

func (s *UpdateDataLakeDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateDataLakeDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateDataLakeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataLakeDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataLakeDatabaseResponseBody) SetDatabase(v *DLDatabase) *UpdateDataLakeDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *UpdateDataLakeDatabaseResponseBody) SetErrorCode(v string) *UpdateDataLakeDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateDataLakeDatabaseResponseBody) SetErrorMessage(v string) *UpdateDataLakeDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateDataLakeDatabaseResponseBody) SetRequestId(v string) *UpdateDataLakeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataLakeDatabaseResponseBody) SetSuccess(v bool) *UpdateDataLakeDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataLakeDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}
