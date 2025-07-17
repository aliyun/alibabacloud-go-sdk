// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataImportSQLTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataImportSQLTypeResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataImportSQLTypeResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataImportSQLTypeResponseBody
	GetRequestId() *string
	SetSqlTypeResult(v []*string) *ListDataImportSQLTypeResponseBody
	GetSqlTypeResult() []*string
	SetSuccess(v bool) *ListDataImportSQLTypeResponseBody
	GetSuccess() *bool
}

type ListDataImportSQLTypeResponseBody struct {
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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// E5EE2B9E-2F95-57FA-B284-CB441CEE49D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The types of SQL statements.
	SqlTypeResult []*string `json:"SqlTypeResult,omitempty" xml:"SqlTypeResult,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataImportSQLTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataImportSQLTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataImportSQLTypeResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataImportSQLTypeResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataImportSQLTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataImportSQLTypeResponseBody) GetSqlTypeResult() []*string {
	return s.SqlTypeResult
}

func (s *ListDataImportSQLTypeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataImportSQLTypeResponseBody) SetErrorCode(v string) *ListDataImportSQLTypeResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataImportSQLTypeResponseBody) SetErrorMessage(v string) *ListDataImportSQLTypeResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataImportSQLTypeResponseBody) SetRequestId(v string) *ListDataImportSQLTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataImportSQLTypeResponseBody) SetSqlTypeResult(v []*string) *ListDataImportSQLTypeResponseBody {
	s.SqlTypeResult = v
	return s
}

func (s *ListDataImportSQLTypeResponseBody) SetSuccess(v bool) *ListDataImportSQLTypeResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataImportSQLTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
