// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTraceInfoDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeTraceInfoDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTraceInfoDetailResponseBody
	GetSuccess() *bool
	SetTraceInfoDetail(v *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) *DescribeTraceInfoDetailResponseBody
	GetTraceInfoDetail() *DescribeTraceInfoDetailResponseBodyTraceInfoDetail
}

type DescribeTraceInfoDetailResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The details of the tracing diagram.
	TraceInfoDetail *DescribeTraceInfoDetailResponseBodyTraceInfoDetail `json:"TraceInfoDetail,omitempty" xml:"TraceInfoDetail,omitempty" type:"Struct"`
}

func (s DescribeTraceInfoDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTraceInfoDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTraceInfoDetailResponseBody) GetTraceInfoDetail() *DescribeTraceInfoDetailResponseBodyTraceInfoDetail {
	return s.TraceInfoDetail
}

func (s *DescribeTraceInfoDetailResponseBody) SetRequestId(v string) *DescribeTraceInfoDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBody) SetSuccess(v bool) *DescribeTraceInfoDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBody) SetTraceInfoDetail(v *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) *DescribeTraceInfoDetailResponseBody {
	s.TraceInfoDetail = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetail struct {
	// An array that consists of the edges of the tracing diagram.
	EdgeList []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList `json:"EdgeList,omitempty" xml:"EdgeList,omitempty" type:"Repeated"`
	// An array that consists of the metadata configurations of the vertex type.
	EntityTypeList []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList `json:"EntityTypeList,omitempty" xml:"EntityTypeList,omitempty" type:"Repeated"`
	// An array that consists of the metadata configurations of the edge type.
	RelationTypeList []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList `json:"RelationTypeList,omitempty" xml:"RelationTypeList,omitempty" type:"Repeated"`
	// An array that consists of all vertexes of the tracing diagram.
	VertexList []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList `json:"VertexList,omitempty" xml:"VertexList,omitempty" type:"Repeated"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetail) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) GetEdgeList() []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	return s.EdgeList
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) GetEntityTypeList() []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	return s.EntityTypeList
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) GetRelationTypeList() []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	return s.RelationTypeList
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) GetVertexList() []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	return s.VertexList
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) SetEdgeList(v []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) *DescribeTraceInfoDetailResponseBodyTraceInfoDetail {
	s.EdgeList = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) SetEntityTypeList(v []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) *DescribeTraceInfoDetailResponseBodyTraceInfoDetail {
	s.EntityTypeList = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) SetRelationTypeList(v []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) *DescribeTraceInfoDetailResponseBodyTraceInfoDetail {
	s.RelationTypeList = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) SetVertexList(v []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) *DescribeTraceInfoDetailResponseBodyTraceInfoDetail {
	s.VertexList = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList struct {
	// The number of times.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ending vertex ID of the edge of the tracing diagram.
	//
	// example:
	//
	// a1d1fa39e5345dcef3f9712172cxxxxx
	EndId *string `json:"EndId,omitempty" xml:"EndId,omitempty"`
	// The starting vertex ID of the edge of the tracing diagram.
	//
	// example:
	//
	// 02b4bf933c8e3bb8b9465eee502xxxxx
	StartId *string `json:"StartId,omitempty" xml:"StartId,omitempty"`
	// The point in time.
	//
	// example:
	//
	// 2022-12-21 10:24:42
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the edge of the tracing diagram.
	//
	// example:
	//
	// trigger_file_alert
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GetEndId() *string {
	return s.EndId
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GetStartId() *string {
	return s.StartId
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GetTime() *string {
	return s.Time
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) SetCount(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	s.Count = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) SetEndId(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	s.EndId = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) SetStartId(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	s.StartId = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) SetTime(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	s.Time = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) SetType(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEdgeList) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList struct {
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	DbId *int32 `json:"DbId,omitempty" xml:"DbId,omitempty"`
	// The rendering color of the vertex.
	//
	// example:
	//
	// #fff
	DisplayColor *string `json:"DisplayColor,omitempty" xml:"DisplayColor,omitempty"`
	// The icon style of the vertex.
	//
	// example:
	//
	// https://img.alicdn.com/tfs/TB176P5OgDqK1RjSZSyXXaxEVXa-49-48.png
	DisplayIcon *string `json:"DisplayIcon,omitempty" xml:"DisplayIcon,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// [{"name":"${logtime}","value":"$!{time}"}]
	DisplayTemplate *string `json:"DisplayTemplate,omitempty" xml:"DisplayTemplate,omitempty"`
	// The timestamp when the vertex was created.
	//
	// example:
	//
	// 2022-10-09T11:47Z
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the vertex was last modified.
	//
	// example:
	//
	// 2022-10-09T11:47Z
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the vertex type.
	//
	// example:
	//
	// Alert
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the vertex type.
	//
	// example:
	//
	// Alert
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// *
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetDbId() *int32 {
	return s.DbId
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetDisplayColor() *string {
	return s.DisplayColor
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetDisplayIcon() *string {
	return s.DisplayIcon
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetDisplayTemplate() *string {
	return s.DisplayTemplate
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetId() *string {
	return s.Id
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetLimit() *int32 {
	return s.Limit
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetName() *string {
	return s.Name
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetNamespace() *string {
	return s.Namespace
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) GetOffset() *int32 {
	return s.Offset
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetDbId(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.DbId = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetDisplayColor(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.DisplayColor = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetDisplayIcon(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.DisplayIcon = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetDisplayTemplate(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.DisplayTemplate = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetGmtCreate(v int64) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.GmtCreate = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetGmtModified(v int64) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.GmtModified = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetId(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.Id = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetLimit(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.Limit = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetName(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.Name = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetNamespace(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.Namespace = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) SetOffset(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList {
	s.Offset = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailEntityTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList struct {
	// Indicates whether the edge is a directional edge. Valid values:
	//
	// 	- **1**: yes
	//
	// 	- **0**: no
	//
	// example:
	//
	// 1
	Directed *int32 `json:"Directed,omitempty" xml:"Directed,omitempty"`
	// The rendering color of the edge.
	//
	// example:
	//
	// #fff
	DisplayColor *string `json:"DisplayColor,omitempty" xml:"DisplayColor,omitempty"`
	// The name of the edge type.
	//
	// example:
	//
	// file
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the edge type.
	//
	// example:
	//
	// netflow_to_process
	RelationTypeId *string `json:"RelationTypeId,omitempty" xml:"RelationTypeId,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// Deprecated
	ShowType *string `json:"ShowType,omitempty" xml:"ShowType,omitempty"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GetDirected() *int32 {
	return s.Directed
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GetDisplayColor() *string {
	return s.DisplayColor
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GetName() *string {
	return s.Name
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GetRelationTypeId() *string {
	return s.RelationTypeId
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) GetShowType() *string {
	return s.ShowType
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) SetDirected(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	s.Directed = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) SetDisplayColor(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	s.DisplayColor = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) SetName(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	s.Name = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) SetRelationTypeId(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	s.RelationTypeId = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) SetShowType(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList {
	s.ShowType = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailRelationTypeList) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList struct {
	// The number of times.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the vertex.
	//
	// example:
	//
	// a1d1fa39e5345dcef3f9712172xxxxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the entity represented by the vertex.
	//
	// example:
	//
	// /usr/local/tomcat
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// An array that consists of the neighbor nodes.
	NeighborList []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList `json:"NeighborList,omitempty" xml:"NeighborList,omitempty" type:"Repeated"`
	// The point in time.
	//
	// example:
	//
	// 2022-12-21 10:24:42
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// The type of the entity represented by the vertex.
	//
	// example:
	//
	// file_path
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetId() *string {
	return s.Id
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetName() *string {
	return s.Name
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetNeighborList() []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList {
	return s.NeighborList
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetTime() *string {
	return s.Time
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetCount(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.Count = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetId(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.Id = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetName(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.Name = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetNeighborList(v []*DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.NeighborList = v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetTime(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.Time = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) SetType(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexList) Validate() error {
	return dara.Validate(s)
}

type DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList struct {
	// The number of neighbor nodes.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Indicates whether one more page is returned.
	//
	// example:
	//
	// False
	HasMore *bool `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// The type of the neighbor node. The value is fixed as **alert**.
	//
	// example:
	//
	// alert
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) GoString() string {
	return s.String()
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) GetCount() *int32 {
	return s.Count
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) GetHasMore() *bool {
	return s.HasMore
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) GetType() *string {
	return s.Type
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) SetCount(v int32) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList {
	s.Count = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) SetHasMore(v bool) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList {
	s.HasMore = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) SetType(v string) *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList {
	s.Type = &v
	return s
}

func (s *DescribeTraceInfoDetailResponseBodyTraceInfoDetailVertexListNeighborList) Validate() error {
	return dara.Validate(s)
}
