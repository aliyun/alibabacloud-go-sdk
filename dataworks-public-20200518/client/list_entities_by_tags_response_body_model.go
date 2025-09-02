// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEntitiesByTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListEntitiesByTagsResponseBodyData) *ListEntitiesByTagsResponseBody
	GetData() *ListEntitiesByTagsResponseBodyData
	SetErrorCode(v string) *ListEntitiesByTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListEntitiesByTagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListEntitiesByTagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListEntitiesByTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListEntitiesByTagsResponseBody
	GetSuccess() *bool
}

type ListEntitiesByTagsResponseBody struct {
	// The data returned.
	Data *ListEntitiesByTagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 101011005
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Invalid.Entity.EntityTypeNotSupported
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
	// 0000-ABCD-E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// true\\
	//
	// false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListEntitiesByTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesByTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEntitiesByTagsResponseBody) GetData() *ListEntitiesByTagsResponseBodyData {
	return s.Data
}

func (s *ListEntitiesByTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListEntitiesByTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListEntitiesByTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListEntitiesByTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEntitiesByTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListEntitiesByTagsResponseBody) SetData(v *ListEntitiesByTagsResponseBodyData) *ListEntitiesByTagsResponseBody {
	s.Data = v
	return s
}

func (s *ListEntitiesByTagsResponseBody) SetErrorCode(v string) *ListEntitiesByTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListEntitiesByTagsResponseBody) SetErrorMessage(v string) *ListEntitiesByTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListEntitiesByTagsResponseBody) SetHttpStatusCode(v int32) *ListEntitiesByTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListEntitiesByTagsResponseBody) SetRequestId(v string) *ListEntitiesByTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEntitiesByTagsResponseBody) SetSuccess(v bool) *ListEntitiesByTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListEntitiesByTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEntitiesByTagsResponseBodyData struct {
	// The entities.
	EntityList []*Entity `json:"EntityList,omitempty" xml:"EntityList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 12345
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListEntitiesByTagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEntitiesByTagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEntitiesByTagsResponseBodyData) GetEntityList() []*Entity {
	return s.EntityList
}

func (s *ListEntitiesByTagsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEntitiesByTagsResponseBodyData) SetEntityList(v []*Entity) *ListEntitiesByTagsResponseBodyData {
	s.EntityList = v
	return s
}

func (s *ListEntitiesByTagsResponseBodyData) SetNextToken(v string) *ListEntitiesByTagsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListEntitiesByTagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
