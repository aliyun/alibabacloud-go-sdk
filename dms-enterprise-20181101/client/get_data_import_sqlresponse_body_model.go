// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataImportSQLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetDataImportSQLResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataImportSQLResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataImportSQLResponseBody
	GetRequestId() *string
	SetSQLDetail(v *GetDataImportSQLResponseBodySQLDetail) *GetDataImportSQLResponseBody
	GetSQLDetail() *GetDataImportSQLResponseBodySQLDetail
	SetSuccess(v bool) *GetDataImportSQLResponseBody
	GetSuccess() *bool
}

type GetDataImportSQLResponseBody struct {
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
	// B43AD641-49C2-5299-9E06-1B37EC1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of SQL statements.
	SQLDetail *GetDataImportSQLResponseBodySQLDetail `json:"SQLDetail,omitempty" xml:"SQLDetail,omitempty" type:"Struct"`
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

func (s GetDataImportSQLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataImportSQLResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataImportSQLResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataImportSQLResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataImportSQLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataImportSQLResponseBody) GetSQLDetail() *GetDataImportSQLResponseBodySQLDetail {
	return s.SQLDetail
}

func (s *GetDataImportSQLResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataImportSQLResponseBody) SetErrorCode(v string) *GetDataImportSQLResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataImportSQLResponseBody) SetErrorMessage(v string) *GetDataImportSQLResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataImportSQLResponseBody) SetRequestId(v string) *GetDataImportSQLResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataImportSQLResponseBody) SetSQLDetail(v *GetDataImportSQLResponseBodySQLDetail) *GetDataImportSQLResponseBody {
	s.SQLDetail = v
	return s
}

func (s *GetDataImportSQLResponseBody) SetSuccess(v bool) *GetDataImportSQLResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataImportSQLResponseBody) Validate() error {
	if s.SQLDetail != nil {
		if err := s.SQLDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataImportSQLResponseBodySQLDetail struct {
	// The SQL script.
	//
	// example:
	//
	// insert into t1 values (1);
	ExecSql *string `json:"ExecSql,omitempty" xml:"ExecSql,omitempty"`
}

func (s GetDataImportSQLResponseBodySQLDetail) String() string {
	return dara.Prettify(s)
}

func (s GetDataImportSQLResponseBodySQLDetail) GoString() string {
	return s.String()
}

func (s *GetDataImportSQLResponseBodySQLDetail) GetExecSql() *string {
	return s.ExecSql
}

func (s *GetDataImportSQLResponseBodySQLDetail) SetExecSql(v string) *GetDataImportSQLResponseBodySQLDetail {
	s.ExecSql = &v
	return s
}

func (s *GetDataImportSQLResponseBodySQLDetail) Validate() error {
	return dara.Validate(s)
}
