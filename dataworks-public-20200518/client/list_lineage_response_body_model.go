// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLineageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListLineageResponseBodyData) *ListLineageResponseBody
	GetData() *ListLineageResponseBodyData
	SetErrorCode(v string) *ListLineageResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListLineageResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListLineageResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListLineageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLineageResponseBody
	GetSuccess() *bool
}

type ListLineageResponseBody struct {
	// The structure returned.
	Data *ListLineageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 1010040007
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// qualifiedName should be in format as entity-table.entity-guid
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
	// 64B-587A-8CED-969E1973887FXXX-TT
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

func (s ListLineageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLineageResponseBody) GoString() string {
	return s.String()
}

func (s *ListLineageResponseBody) GetData() *ListLineageResponseBodyData {
	return s.Data
}

func (s *ListLineageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListLineageResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListLineageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLineageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLineageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLineageResponseBody) SetData(v *ListLineageResponseBodyData) *ListLineageResponseBody {
	s.Data = v
	return s
}

func (s *ListLineageResponseBody) SetErrorCode(v string) *ListLineageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListLineageResponseBody) SetErrorMessage(v string) *ListLineageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListLineageResponseBody) SetHttpStatusCode(v int32) *ListLineageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLineageResponseBody) SetRequestId(v string) *ListLineageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLineageResponseBody) SetSuccess(v bool) *ListLineageResponseBody {
	s.Success = &v
	return s
}

func (s *ListLineageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLineageResponseBodyData struct {
	// The array of the entity structure.
	DataEntityList []*ListLineageResponseBodyDataDataEntityList `json:"DataEntityList,omitempty" xml:"DataEntityList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// nextTokenFromRequest-xxxsd-ff
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListLineageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListLineageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListLineageResponseBodyData) GetDataEntityList() []*ListLineageResponseBodyDataDataEntityList {
	return s.DataEntityList
}

func (s *ListLineageResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLineageResponseBodyData) SetDataEntityList(v []*ListLineageResponseBodyDataDataEntityList) *ListLineageResponseBodyData {
	s.DataEntityList = v
	return s
}

func (s *ListLineageResponseBodyData) SetNextToken(v string) *ListLineageResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListLineageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListLineageResponseBodyDataDataEntityList struct {
	// The time when the lineage was generated.
	//
	// example:
	//
	// 1686215809269
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The information about the entity.
	Entity *Entity `json:"Entity,omitempty" xml:"Entity,omitempty"`
	// The array of the relationship structure.
	RelationList []*ListLineageResponseBodyDataDataEntityListRelationList `json:"RelationList,omitempty" xml:"RelationList,omitempty" type:"Repeated"`
}

func (s ListLineageResponseBodyDataDataEntityList) String() string {
	return dara.Prettify(s)
}

func (s ListLineageResponseBodyDataDataEntityList) GoString() string {
	return s.String()
}

func (s *ListLineageResponseBodyDataDataEntityList) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListLineageResponseBodyDataDataEntityList) GetEntity() *Entity {
	return s.Entity
}

func (s *ListLineageResponseBodyDataDataEntityList) GetRelationList() []*ListLineageResponseBodyDataDataEntityListRelationList {
	return s.RelationList
}

func (s *ListLineageResponseBodyDataDataEntityList) SetCreateTimestamp(v int64) *ListLineageResponseBodyDataDataEntityList {
	s.CreateTimestamp = &v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityList) SetEntity(v *Entity) *ListLineageResponseBodyDataDataEntityList {
	s.Entity = v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityList) SetRelationList(v []*ListLineageResponseBodyDataDataEntityListRelationList) *ListLineageResponseBodyDataDataEntityList {
	s.RelationList = v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityList) Validate() error {
	return dara.Validate(s)
}

type ListLineageResponseBodyDataDataEntityListRelationList struct {
	// The data channel. Valid values:
	//
	// 	- **FIRST_PARTY: DataWorks platform**
	//
	// 	- **THIRD_PARTY: user registration**
	//
	// example:
	//
	// THIRD_PARTY
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// The data source.
	//
	// example:
	//
	// mysql
	Datasource *string `json:"Datasource,omitempty" xml:"Datasource,omitempty"`
	// The unique relationship ID.
	//
	// example:
	//
	// aaabbccddguid
	Guid *string `json:"Guid,omitempty" xml:"Guid,omitempty"`
	// The task type, which is used to describe the relationship between entities, such as SQL-based calculation, mapping based on report fields, or API operation definition.
	//
	// example:
	//
	// sql
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListLineageResponseBodyDataDataEntityListRelationList) String() string {
	return dara.Prettify(s)
}

func (s ListLineageResponseBodyDataDataEntityListRelationList) GoString() string {
	return s.String()
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) GetChannel() *string {
	return s.Channel
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) GetDatasource() *string {
	return s.Datasource
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) GetGuid() *string {
	return s.Guid
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) GetType() *string {
	return s.Type
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) SetChannel(v string) *ListLineageResponseBodyDataDataEntityListRelationList {
	s.Channel = &v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) SetDatasource(v string) *ListLineageResponseBodyDataDataEntityListRelationList {
	s.Datasource = &v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) SetGuid(v string) *ListLineageResponseBodyDataDataEntityListRelationList {
	s.Guid = &v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) SetType(v string) *ListLineageResponseBodyDataDataEntityListRelationList {
	s.Type = &v
	return s
}

func (s *ListLineageResponseBodyDataDataEntityListRelationList) Validate() error {
	return dara.Validate(s)
}
