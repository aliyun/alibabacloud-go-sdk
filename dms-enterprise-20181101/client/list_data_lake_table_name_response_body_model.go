// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTableNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakeTableNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeTableNameResponseBody
	GetErrorMessage() *string
	SetMaxResults(v int32) *ListDataLakeTableNameResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataLakeTableNameResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDataLakeTableNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeTableNameResponseBody
	GetSuccess() *bool
	SetTableNameList(v []*string) *ListDataLakeTableNameResponseBody
	GetTableNameList() []*string
}

type ListDataLakeTableNameResponseBody struct {
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
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// C5B8E84B-42B6-4374-AD5A-6264E1753378
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success       *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	TableNameList []*string `json:"TableNameList,omitempty" xml:"TableNameList,omitempty" type:"Repeated"`
}

func (s ListDataLakeTableNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTableNameResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeTableNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeTableNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeTableNameResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataLakeTableNameResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataLakeTableNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeTableNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeTableNameResponseBody) GetTableNameList() []*string {
	return s.TableNameList
}

func (s *ListDataLakeTableNameResponseBody) SetErrorCode(v string) *ListDataLakeTableNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetErrorMessage(v string) *ListDataLakeTableNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetMaxResults(v int32) *ListDataLakeTableNameResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetNextToken(v string) *ListDataLakeTableNameResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetRequestId(v string) *ListDataLakeTableNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetSuccess(v bool) *ListDataLakeTableNameResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeTableNameResponseBody) SetTableNameList(v []*string) *ListDataLakeTableNameResponseBody {
	s.TableNameList = v
	return s
}

func (s *ListDataLakeTableNameResponseBody) Validate() error {
	return dara.Validate(s)
}
