// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataLakeTablebaseInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListDataLakeTablebaseInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDataLakeTablebaseInfoResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListDataLakeTablebaseInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataLakeTablebaseInfoResponseBody
	GetSuccess() *bool
	SetTablebaseInfoList(v []*DLTablebaseInfo) *ListDataLakeTablebaseInfoResponseBody
	GetTablebaseInfoList() []*DLTablebaseInfo
	SetTotalCount(v string) *ListDataLakeTablebaseInfoResponseBody
	GetTotalCount() *string
}

type ListDataLakeTablebaseInfoResponseBody struct {
	// example:
	//
	// 400
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Rows is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// B43AD641-49C2-5299-9E06-1B37EC1B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success           *bool              `json:"Success,omitempty" xml:"Success,omitempty"`
	TablebaseInfoList []*DLTablebaseInfo `json:"TablebaseInfoList,omitempty" xml:"TablebaseInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDataLakeTablebaseInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataLakeTablebaseInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetTablebaseInfoList() []*DLTablebaseInfo {
	return s.TablebaseInfoList
}

func (s *ListDataLakeTablebaseInfoResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetErrorCode(v string) *ListDataLakeTablebaseInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetErrorMessage(v string) *ListDataLakeTablebaseInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetRequestId(v string) *ListDataLakeTablebaseInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetSuccess(v bool) *ListDataLakeTablebaseInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetTablebaseInfoList(v []*DLTablebaseInfo) *ListDataLakeTablebaseInfoResponseBody {
	s.TablebaseInfoList = v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) SetTotalCount(v string) *ListDataLakeTablebaseInfoResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataLakeTablebaseInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
