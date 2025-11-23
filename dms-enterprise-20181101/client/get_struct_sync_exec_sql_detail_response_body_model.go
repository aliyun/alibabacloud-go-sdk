// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncExecSqlDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetStructSyncExecSqlDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetStructSyncExecSqlDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetStructSyncExecSqlDetailResponseBody
	GetRequestId() *string
	SetStructSyncExecSqlDetail(v *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) *GetStructSyncExecSqlDetailResponseBody
	GetStructSyncExecSqlDetail() *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail
	SetSuccess(v bool) *GetStructSyncExecSqlDetailResponseBody
	GetSuccess() *bool
}

type GetStructSyncExecSqlDetailResponseBody struct {
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
	// 1810E635-A2D7-428B-BAA9-85DAEB9B1A77
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the SQL statements.
	StructSyncExecSqlDetail *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail `json:"StructSyncExecSqlDetail,omitempty" xml:"StructSyncExecSqlDetail,omitempty" type:"Struct"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStructSyncExecSqlDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStructSyncExecSqlDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStructSyncExecSqlDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStructSyncExecSqlDetailResponseBody) GetStructSyncExecSqlDetail() *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail {
	return s.StructSyncExecSqlDetail
}

func (s *GetStructSyncExecSqlDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetErrorCode(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetErrorMessage(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetRequestId(v string) *GetStructSyncExecSqlDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetStructSyncExecSqlDetail(v *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) *GetStructSyncExecSqlDetailResponseBody {
	s.StructSyncExecSqlDetail = v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) SetSuccess(v bool) *GetStructSyncExecSqlDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBody) Validate() error {
	if s.StructSyncExecSqlDetail != nil {
		if err := s.StructSyncExecSqlDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail struct {
	// The SQL statements that are executed.
	ExecSql *string `json:"ExecSql,omitempty" xml:"ExecSql,omitempty"`
	// The total number of SQL statements.
	//
	// example:
	//
	// 1
	TotalSqlCount *int64 `json:"TotalSqlCount,omitempty" xml:"TotalSqlCount,omitempty"`
}

func (s GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) GetExecSql() *string {
	return s.ExecSql
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) GetTotalSqlCount() *int64 {
	return s.TotalSqlCount
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) SetExecSql(v string) *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail {
	s.ExecSql = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) SetTotalSqlCount(v int64) *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail {
	s.TotalSqlCount = &v
	return s
}

func (s *GetStructSyncExecSqlDetailResponseBodyStructSyncExecSqlDetail) Validate() error {
	return dara.Validate(s)
}
