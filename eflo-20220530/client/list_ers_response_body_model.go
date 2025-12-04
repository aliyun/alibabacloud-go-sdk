// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListErsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListErsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *ListErsResponseBody
	GetCode() *int32
	SetContent(v *ListErsResponseBodyContent) *ListErsResponseBody
	GetContent() *ListErsResponseBodyContent
	SetMessage(v string) *ListErsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListErsResponseBody
	GetRequestId() *string
}

type ListErsResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Content *ListErsResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The error message. (If the instance is in the Exception state, the exception cause is displayed.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListErsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListErsResponseBody) GoString() string {
	return s.String()
}

func (s *ListErsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListErsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListErsResponseBody) GetContent() *ListErsResponseBodyContent {
	return s.Content
}

func (s *ListErsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListErsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListErsResponseBody) SetAccessDeniedDetail(v string) *ListErsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListErsResponseBody) SetCode(v int32) *ListErsResponseBody {
	s.Code = &v
	return s
}

func (s *ListErsResponseBody) SetContent(v *ListErsResponseBodyContent) *ListErsResponseBody {
	s.Content = v
	return s
}

func (s *ListErsResponseBody) SetMessage(v string) *ListErsResponseBody {
	s.Message = &v
	return s
}

func (s *ListErsResponseBody) SetRequestId(v string) *ListErsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListErsResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListErsResponseBodyContent struct {
	// lingjun hub information list.
	Data []*ListErsResponseBodyContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListErsResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListErsResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListErsResponseBodyContent) GetData() []*ListErsResponseBodyContentData {
	return s.Data
}

func (s *ListErsResponseBodyContent) GetTotal() *int64 {
	return s.Total
}

func (s *ListErsResponseBodyContent) SetData(v []*ListErsResponseBodyContentData) *ListErsResponseBodyContent {
	s.Data = v
	return s
}

func (s *ListErsResponseBodyContent) SetTotal(v int64) *ListErsResponseBodyContent {
	s.Total = &v
	return s
}

func (s *ListErsResponseBodyContent) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErsResponseBodyContentData struct {
	// The number of connections to the Lingjun HUB network instance.
	//
	// example:
	//
	// 2
	Connections *int64 `json:"Connections,omitempty" xml:"Connections,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 1640930671000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the synchronization task.
	//
	// example:
	//
	// No description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the Lingjun HUB instance.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The name of the Lingjun HUB instance.
	//
	// example:
	//
	// er-wulanchabu-main
	ErName *string `json:"ErName,omitempty" xml:"ErName,omitempty"`
	// The time when the O\\&M task was modified.
	//
	// example:
	//
	// 1640930671000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The primary zone.
	//
	// example:
	//
	// cn-wulanchabu-b
	MasterZoneId *string `json:"MasterZoneId,omitempty" xml:"MasterZoneId,omitempty"`
	// The returned message.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Resource group instance ID
	//
	// example:
	//
	// rg-acfmv2m2w43japa
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Number of Lingjun HUB routing policy.
	//
	// example:
	//
	// 2
	RouteMaps *int64 `json:"RouteMaps,omitempty" xml:"RouteMaps,omitempty"`
	// The task status.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*ListErsResponseBodyContentDataTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The tenant ID.
	//
	// example:
	//
	// 1655449505171
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListErsResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s ListErsResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *ListErsResponseBodyContentData) GetConnections() *int64 {
	return s.Connections
}

func (s *ListErsResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListErsResponseBodyContentData) GetDescription() *string {
	return s.Description
}

func (s *ListErsResponseBodyContentData) GetErId() *string {
	return s.ErId
}

func (s *ListErsResponseBodyContentData) GetErName() *string {
	return s.ErName
}

func (s *ListErsResponseBodyContentData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListErsResponseBodyContentData) GetMasterZoneId() *string {
	return s.MasterZoneId
}

func (s *ListErsResponseBodyContentData) GetMessage() *string {
	return s.Message
}

func (s *ListErsResponseBodyContentData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListErsResponseBodyContentData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListErsResponseBodyContentData) GetRouteMaps() *int64 {
	return s.RouteMaps
}

func (s *ListErsResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *ListErsResponseBodyContentData) GetTags() []*ListErsResponseBodyContentDataTags {
	return s.Tags
}

func (s *ListErsResponseBodyContentData) GetTenantId() *string {
	return s.TenantId
}

func (s *ListErsResponseBodyContentData) SetConnections(v int64) *ListErsResponseBodyContentData {
	s.Connections = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetCreateTime(v string) *ListErsResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetDescription(v string) *ListErsResponseBodyContentData {
	s.Description = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetErId(v string) *ListErsResponseBodyContentData {
	s.ErId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetErName(v string) *ListErsResponseBodyContentData {
	s.ErName = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetGmtModified(v string) *ListErsResponseBodyContentData {
	s.GmtModified = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetMasterZoneId(v string) *ListErsResponseBodyContentData {
	s.MasterZoneId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetMessage(v string) *ListErsResponseBodyContentData {
	s.Message = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetRegionId(v string) *ListErsResponseBodyContentData {
	s.RegionId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetResourceGroupId(v string) *ListErsResponseBodyContentData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetRouteMaps(v int64) *ListErsResponseBodyContentData {
	s.RouteMaps = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetStatus(v string) *ListErsResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *ListErsResponseBodyContentData) SetTags(v []*ListErsResponseBodyContentDataTags) *ListErsResponseBodyContentData {
	s.Tags = v
	return s
}

func (s *ListErsResponseBodyContentData) SetTenantId(v string) *ListErsResponseBodyContentData {
	s.TenantId = &v
	return s
}

func (s *ListErsResponseBodyContentData) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListErsResponseBodyContentDataTags struct {
	// The tag key.
	//
	// example:
	//
	// myTest
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListErsResponseBodyContentDataTags) String() string {
	return dara.Prettify(s)
}

func (s ListErsResponseBodyContentDataTags) GoString() string {
	return s.String()
}

func (s *ListErsResponseBodyContentDataTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListErsResponseBodyContentDataTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListErsResponseBodyContentDataTags) SetTagKey(v string) *ListErsResponseBodyContentDataTags {
	s.TagKey = &v
	return s
}

func (s *ListErsResponseBodyContentDataTags) SetTagValue(v string) *ListErsResponseBodyContentDataTags {
	s.TagValue = &v
	return s
}

func (s *ListErsResponseBodyContentDataTags) Validate() error {
	return dara.Validate(s)
}
