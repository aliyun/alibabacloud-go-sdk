// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataLakeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabase(v *DLDatabase) *GetDataLakeDatabaseResponseBody
	GetDatabase() *DLDatabase
	SetErrorCode(v string) *GetDataLakeDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataLakeDatabaseResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataLakeDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataLakeDatabaseResponseBody
	GetSuccess() *bool
}

type GetDataLakeDatabaseResponseBody struct {
	// The database information.
	Database *DLDatabase `json:"Database,omitempty" xml:"Database,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that is returned if the request failed.
	//
	// example:
	//
	// code: 404, can not find catalog, name : hive1 request id: FF737753-9641-1F51-AFDA-7DF541114B29
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E5EE2B9E-2F95-57FA-B284-CB441CEE49D6
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

func (s GetDataLakeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataLakeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataLakeDatabaseResponseBody) GetDatabase() *DLDatabase {
	return s.Database
}

func (s *GetDataLakeDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataLakeDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataLakeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataLakeDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataLakeDatabaseResponseBody) SetDatabase(v *DLDatabase) *GetDataLakeDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDataLakeDatabaseResponseBody) SetErrorCode(v string) *GetDataLakeDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataLakeDatabaseResponseBody) SetErrorMessage(v string) *GetDataLakeDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataLakeDatabaseResponseBody) SetRequestId(v string) *GetDataLakeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataLakeDatabaseResponseBody) SetSuccess(v bool) *GetDataLakeDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataLakeDatabaseResponseBody) Validate() error {
	if s.Database != nil {
		if err := s.Database.Validate(); err != nil {
			return err
		}
	}
	return nil
}
