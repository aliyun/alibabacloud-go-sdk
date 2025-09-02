// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMetaCollectionsResponseBodyData) *ListMetaCollectionsResponseBody
	GetData() *ListMetaCollectionsResponseBodyData
	SetErrorCode(v string) *ListMetaCollectionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMetaCollectionsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListMetaCollectionsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListMetaCollectionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaCollectionsResponseBody
	GetSuccess() *bool
}

type ListMetaCollectionsResponseBody struct {
	// The returned result.
	Data *ListMetaCollectionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Collection.NotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A0DE367DDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMetaCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponseBody) GetData() *ListMetaCollectionsResponseBodyData {
	return s.Data
}

func (s *ListMetaCollectionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMetaCollectionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMetaCollectionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMetaCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaCollectionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaCollectionsResponseBody) SetData(v *ListMetaCollectionsResponseBodyData) *ListMetaCollectionsResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetErrorCode(v string) *ListMetaCollectionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetErrorMessage(v string) *ListMetaCollectionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetHttpStatusCode(v int32) *ListMetaCollectionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetRequestId(v string) *ListMetaCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) SetSuccess(v bool) *ListMetaCollectionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaCollectionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMetaCollectionsResponseBodyData struct {
	// The collections.
	CollectionList []*Collection `json:"CollectionList,omitempty" xml:"CollectionList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 37ae2053d87d380f28ce0dc0853ca51e
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMetaCollectionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionsResponseBodyData) GetCollectionList() []*Collection {
	return s.CollectionList
}

func (s *ListMetaCollectionsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaCollectionsResponseBodyData) SetCollectionList(v []*Collection) *ListMetaCollectionsResponseBodyData {
	s.CollectionList = v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) SetNextToken(v string) *ListMetaCollectionsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMetaCollectionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
