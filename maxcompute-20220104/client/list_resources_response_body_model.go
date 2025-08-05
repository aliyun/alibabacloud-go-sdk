// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListResourcesResponseBodyData) *ListResourcesResponseBody
	GetData() *ListResourcesResponseBodyData
	SetRequestId(v string) *ListResourcesResponseBody
	GetRequestId() *string
}

type ListResourcesResponseBody struct {
	// The returned data.
	Data *ListResourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0bc3b4ae16685836687916212e7850
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) GetData() *ListResourcesResponseBodyData {
	return s.Data
}

func (s *ListResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourcesResponseBody) SetData(v *ListResourcesResponseBodyData) *ListResourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyData struct {
	// Indicates the marker after which the returned list begins.
	//
	// example:
	//
	// ZmN0X21vbnRoX3Rhb2Jhb19pbmRleCE=
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The maximum number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxItem *int32 `json:"maxItem,omitempty" xml:"maxItem,omitempty"`
	// The list of resources.
	Resources []*ListResourcesResponseBodyDataResources `json:"resources,omitempty" xml:"resources,omitempty" type:"Repeated"`
}

func (s ListResourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *ListResourcesResponseBodyData) GetMaxItem() *int32 {
	return s.MaxItem
}

func (s *ListResourcesResponseBodyData) GetResources() []*ListResourcesResponseBodyDataResources {
	return s.Resources
}

func (s *ListResourcesResponseBodyData) SetMarker(v string) *ListResourcesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetMaxItem(v int32) *ListResourcesResponseBodyData {
	s.MaxItem = &v
	return s
}

func (s *ListResourcesResponseBodyData) SetResources(v []*ListResourcesResponseBodyDataResources) *ListResourcesResponseBodyData {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListResourcesResponseBodyDataResources struct {
	// The remarks.
	//
	// example:
	//
	// file
	Comment *string `json:"comment,omitempty" xml:"comment,omitempty"`
	// The Base64-encoded 128-bit MD5 hash value of the HTTP request body.
	//
	// example:
	//
	// MACiECZtnLiNkNS1v5****=1
	ContentMD5 *string `json:"contentMD5,omitempty" xml:"contentMD5,omitempty"`
	// The time when the resource was created.
	//
	// example:
	//
	// 2022-01-29T03:34:09Z
	CreationTime *int64 `json:"creationTime,omitempty" xml:"creationTime,omitempty"`
	// The display name.
	//
	// example:
	//
	// res_1
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The time when the resource was modified.
	//
	// example:
	//
	// 2023-04-18T06:15:05Z
	LastModifiedTime *int64 `json:"lastModifiedTime,omitempty" xml:"lastModifiedTime,omitempty"`
	// The user who updated the resource.
	//
	// example:
	//
	// ALIYUN$xxx@test.aliyunid.com
	LastUpdator *string `json:"lastUpdator,omitempty" xml:"lastUpdator,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// res_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The owner of the resource.
	//
	// example:
	//
	// 1265860483008101
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The schema to which the resource belongs.
	//
	// example:
	//
	// schemaA
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
	// The size of the resource.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- py
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- jar
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- volumefile
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- table
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListResourcesResponseBodyDataResources) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesResponseBodyDataResources) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBodyDataResources) GetComment() *string {
	return s.Comment
}

func (s *ListResourcesResponseBodyDataResources) GetContentMD5() *string {
	return s.ContentMD5
}

func (s *ListResourcesResponseBodyDataResources) GetCreationTime() *int64 {
	return s.CreationTime
}

func (s *ListResourcesResponseBodyDataResources) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListResourcesResponseBodyDataResources) GetLastModifiedTime() *int64 {
	return s.LastModifiedTime
}

func (s *ListResourcesResponseBodyDataResources) GetLastUpdator() *string {
	return s.LastUpdator
}

func (s *ListResourcesResponseBodyDataResources) GetName() *string {
	return s.Name
}

func (s *ListResourcesResponseBodyDataResources) GetOwner() *string {
	return s.Owner
}

func (s *ListResourcesResponseBodyDataResources) GetSchema() *string {
	return s.Schema
}

func (s *ListResourcesResponseBodyDataResources) GetSize() *int64 {
	return s.Size
}

func (s *ListResourcesResponseBodyDataResources) GetType() *string {
	return s.Type
}

func (s *ListResourcesResponseBodyDataResources) SetComment(v string) *ListResourcesResponseBodyDataResources {
	s.Comment = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetContentMD5(v string) *ListResourcesResponseBodyDataResources {
	s.ContentMD5 = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetCreationTime(v int64) *ListResourcesResponseBodyDataResources {
	s.CreationTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetDisplayName(v string) *ListResourcesResponseBodyDataResources {
	s.DisplayName = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastModifiedTime(v int64) *ListResourcesResponseBodyDataResources {
	s.LastModifiedTime = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetLastUpdator(v string) *ListResourcesResponseBodyDataResources {
	s.LastUpdator = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetName(v string) *ListResourcesResponseBodyDataResources {
	s.Name = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetOwner(v string) *ListResourcesResponseBodyDataResources {
	s.Owner = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSchema(v string) *ListResourcesResponseBodyDataResources {
	s.Schema = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetSize(v int64) *ListResourcesResponseBodyDataResources {
	s.Size = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) SetType(v string) *ListResourcesResponseBodyDataResources {
	s.Type = &v
	return s
}

func (s *ListResourcesResponseBodyDataResources) Validate() error {
	return dara.Validate(s)
}
