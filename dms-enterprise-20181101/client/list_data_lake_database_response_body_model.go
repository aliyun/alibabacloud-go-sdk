// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseList(v []*DLDatabase) *ListDataLakeDatabaseResponseBody
	GetDatabaseList() []*DLDatabase
	SetErrorCode(v string) *ListDataLakeDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeDatabaseResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakeDatabaseResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeDatabaseResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataLakeDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeDatabaseResponseBody
	GetSuccess() *bool
}

type ListDataLakeDatabaseResponseBody struct {
	DatabaseList []*DLDatabase `json:"DatabaseList,omitempty" xml:"DatabaseList,omitempty" type:"Repeated"`
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	MaxResults   *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataLakeDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeDatabaseResponseBody) GetDatabaseList() []*DLDatabase {
	return s.DatabaseList
}

func (s *ListDataLakeDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeDatabaseResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeDatabaseResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeDatabaseResponseBody) SetDatabaseList(v []*DLDatabase) *ListDataLakeDatabaseResponseBody {
	s.DatabaseList = v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetErrorCode(v string) *ListDataLakeDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetErrorMessage(v string) *ListDataLakeDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetMaxResults(v int32) *ListDataLakeDatabaseResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetNextToken(v string) *ListDataLakeDatabaseResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetRequestId(v string) *ListDataLakeDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) SetSuccess(v bool) *ListDataLakeDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
