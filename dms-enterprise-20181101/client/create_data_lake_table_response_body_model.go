// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataLakeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDataLakeTableResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDataLakeTableResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDataLakeTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataLakeTableResponseBody
	GetSuccess() *bool
	SetTable(v *DLTable) *CreateDataLakeTableResponseBody
	GetTable() *DLTable
}

type CreateDataLakeTableResponseBody struct {
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
	// 7FAD400F-7A5C-4193-8F9A-39D86C4F0231
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	Table   *DLTable `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s CreateDataLakeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataLakeTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataLakeTableResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDataLakeTableResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDataLakeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataLakeTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataLakeTableResponseBody) GetTable() *DLTable {
	return s.Table
}

func (s *CreateDataLakeTableResponseBody) SetErrorCode(v string) *CreateDataLakeTableResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDataLakeTableResponseBody) SetErrorMessage(v string) *CreateDataLakeTableResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDataLakeTableResponseBody) SetRequestId(v string) *CreateDataLakeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataLakeTableResponseBody) SetSuccess(v bool) *CreateDataLakeTableResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataLakeTableResponseBody) SetTable(v *DLTable) *CreateDataLakeTableResponseBody {
	s.Table = v
	return s
}

func (s *CreateDataLakeTableResponseBody) Validate() error {
	return dara.Validate(s)
}
