// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListIndexesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListIndexesResponseBody
	GetErrorMessage() *string
	SetIndexList(v *ListIndexesResponseBodyIndexList) *ListIndexesResponseBody
	GetIndexList() *ListIndexesResponseBodyIndexList
	SetRequestId(v string) *ListIndexesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIndexesResponseBody
	GetSuccess() *bool
}

type ListIndexesResponseBody struct {
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
	// The details of indexes.
	IndexList *ListIndexesResponseBodyIndexList `json:"IndexList,omitempty" xml:"IndexList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 1F4DE2F1-5B47-462A-A973-E02EB7AF386B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIndexesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListIndexesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListIndexesResponseBody) GetIndexList() *ListIndexesResponseBodyIndexList {
	return s.IndexList
}

func (s *ListIndexesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIndexesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIndexesResponseBody) SetErrorCode(v string) *ListIndexesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListIndexesResponseBody) SetErrorMessage(v string) *ListIndexesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListIndexesResponseBody) SetIndexList(v *ListIndexesResponseBodyIndexList) *ListIndexesResponseBody {
	s.IndexList = v
	return s
}

func (s *ListIndexesResponseBody) SetRequestId(v string) *ListIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIndexesResponseBody) SetSuccess(v bool) *ListIndexesResponseBody {
	s.Success = &v
	return s
}

func (s *ListIndexesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyIndexList struct {
	Index []*ListIndexesResponseBodyIndexListIndex `json:"Index,omitempty" xml:"Index,omitempty" type:"Repeated"`
}

func (s ListIndexesResponseBodyIndexList) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyIndexList) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyIndexList) GetIndex() []*ListIndexesResponseBodyIndexListIndex {
	return s.Index
}

func (s *ListIndexesResponseBodyIndexList) SetIndex(v []*ListIndexesResponseBodyIndexListIndex) *ListIndexesResponseBodyIndexList {
	s.Index = v
	return s
}

func (s *ListIndexesResponseBodyIndexList) Validate() error {
	return dara.Validate(s)
}

type ListIndexesResponseBodyIndexListIndex struct {
	// The description of the index.
	//
	// example:
	//
	// test
	IndexComment *string `json:"IndexComment,omitempty" xml:"IndexComment,omitempty"`
	// The ID of the index.
	//
	// example:
	//
	// 1
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// idx_test
	IndexName *string `json:"IndexName,omitempty" xml:"IndexName,omitempty"`
	// The type of the index. Valid values:
	//
	// 	- Primary
	//
	// 	- Unique
	//
	// 	- Normal
	//
	// 	- FullText
	//
	// 	- Spatial
	//
	// example:
	//
	// Primary
	IndexType *string `json:"IndexType,omitempty" xml:"IndexType,omitempty"`
	// The ID of the table.
	//
	// example:
	//
	// 1
	TableId *string `json:"TableId,omitempty" xml:"TableId,omitempty"`
}

func (s ListIndexesResponseBodyIndexListIndex) String() string {
	return dara.Prettify(s)
}

func (s ListIndexesResponseBodyIndexListIndex) GoString() string {
	return s.String()
}

func (s *ListIndexesResponseBodyIndexListIndex) GetIndexComment() *string {
	return s.IndexComment
}

func (s *ListIndexesResponseBodyIndexListIndex) GetIndexId() *string {
	return s.IndexId
}

func (s *ListIndexesResponseBodyIndexListIndex) GetIndexName() *string {
	return s.IndexName
}

func (s *ListIndexesResponseBodyIndexListIndex) GetIndexType() *string {
	return s.IndexType
}

func (s *ListIndexesResponseBodyIndexListIndex) GetTableId() *string {
	return s.TableId
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexComment(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexComment = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexId(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexId = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexName(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexName = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetIndexType(v string) *ListIndexesResponseBodyIndexListIndex {
	s.IndexType = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) SetTableId(v string) *ListIndexesResponseBodyIndexListIndex {
	s.TableId = &v
	return s
}

func (s *ListIndexesResponseBodyIndexListIndex) Validate() error {
	return dara.Validate(s)
}
