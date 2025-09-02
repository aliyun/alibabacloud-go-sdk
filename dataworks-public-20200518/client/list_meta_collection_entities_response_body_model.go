// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaCollectionEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListMetaCollectionEntitiesResponseBodyData) *ListMetaCollectionEntitiesResponseBody
	GetData() *ListMetaCollectionEntitiesResponseBodyData
	SetErrorCode(v string) *ListMetaCollectionEntitiesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListMetaCollectionEntitiesResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListMetaCollectionEntitiesResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListMetaCollectionEntitiesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMetaCollectionEntitiesResponseBody
	GetSuccess() *bool
}

type ListMetaCollectionEntitiesResponseBody struct {
	// The response parameters.
	Data *ListMetaCollectionEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID. You can locate logs and troubleshoot issues based on the ID.
	//
	// example:
	//
	// E25887B7-579C-54A5-9C4F-83A0DE367DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true: The request was successful.
	//
	// false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListMetaCollectionEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionEntitiesResponseBody) GetData() *ListMetaCollectionEntitiesResponseBodyData {
	return s.Data
}

func (s *ListMetaCollectionEntitiesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListMetaCollectionEntitiesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListMetaCollectionEntitiesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListMetaCollectionEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaCollectionEntitiesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMetaCollectionEntitiesResponseBody) SetData(v *ListMetaCollectionEntitiesResponseBodyData) *ListMetaCollectionEntitiesResponseBody {
	s.Data = v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) SetErrorCode(v string) *ListMetaCollectionEntitiesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) SetErrorMessage(v string) *ListMetaCollectionEntitiesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) SetHttpStatusCode(v int32) *ListMetaCollectionEntitiesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) SetRequestId(v string) *ListMetaCollectionEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) SetSuccess(v bool) *ListMetaCollectionEntitiesResponseBody {
	s.Success = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListMetaCollectionEntitiesResponseBodyData struct {
	// The entities.
	EntityList []*Entity `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 123344
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListMetaCollectionEntitiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMetaCollectionEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMetaCollectionEntitiesResponseBodyData) GetEntityList() []*Entity {
	return s.EntityList
}

func (s *ListMetaCollectionEntitiesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMetaCollectionEntitiesResponseBodyData) SetEntityList(v []*Entity) *ListMetaCollectionEntitiesResponseBodyData {
	s.EntityList = v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBodyData) SetNextToken(v string) *ListMetaCollectionEntitiesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListMetaCollectionEntitiesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
