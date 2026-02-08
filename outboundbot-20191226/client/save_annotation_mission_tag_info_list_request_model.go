// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnnotationMissionTagInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnnotationMissionTagInfoList(v []*SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) *SaveAnnotationMissionTagInfoListRequest
	GetAnnotationMissionTagInfoList() []*SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList
	SetAnnotationMissionTagInfoListJsonString(v string) *SaveAnnotationMissionTagInfoListRequest
	GetAnnotationMissionTagInfoListJsonString() *string
	SetInstanceId(v string) *SaveAnnotationMissionTagInfoListRequest
	GetInstanceId() *string
	SetReset(v bool) *SaveAnnotationMissionTagInfoListRequest
	GetReset() *bool
}

type SaveAnnotationMissionTagInfoListRequest struct {
	// List of tags
	//
	// > This parameter is equivalent to AnnotationMissionTagInfoListJsonString. You only need to specify one of them.
	AnnotationMissionTagInfoList []*SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList `json:"AnnotationMissionTagInfoList,omitempty" xml:"AnnotationMissionTagInfoList,omitempty" type:"Repeated"`
	// Tag JSON data
	//
	// > The specific parameters in the JSON are the same as those in AnnotationMissionTagInfoList. You can specify either one.
	//
	// example:
	//
	// [{"delete":false,				"InstanceId":"00b37342-e759-4fe5-b296-aef775933af0",					"AnnotationMissionTagInfoName":"测试数据1",					"AnnotationMissionTagInfoDescription":"测试"}]
	AnnotationMissionTagInfoListJsonString *string `json:"AnnotationMissionTagInfoListJsonString,omitempty" xml:"AnnotationMissionTagInfoListJsonString,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// af81a389-91f0-4157-8d82-720edd02b66a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Whether to recover the default tag value
	//
	// example:
	//
	// false
	Reset *bool `json:"Reset,omitempty" xml:"Reset,omitempty"`
}

func (s SaveAnnotationMissionTagInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionTagInfoListRequest) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionTagInfoListRequest) GetAnnotationMissionTagInfoList() []*SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	return s.AnnotationMissionTagInfoList
}

func (s *SaveAnnotationMissionTagInfoListRequest) GetAnnotationMissionTagInfoListJsonString() *string {
	return s.AnnotationMissionTagInfoListJsonString
}

func (s *SaveAnnotationMissionTagInfoListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionTagInfoListRequest) GetReset() *bool {
	return s.Reset
}

func (s *SaveAnnotationMissionTagInfoListRequest) SetAnnotationMissionTagInfoList(v []*SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) *SaveAnnotationMissionTagInfoListRequest {
	s.AnnotationMissionTagInfoList = v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequest) SetAnnotationMissionTagInfoListJsonString(v string) *SaveAnnotationMissionTagInfoListRequest {
	s.AnnotationMissionTagInfoListJsonString = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequest) SetInstanceId(v string) *SaveAnnotationMissionTagInfoListRequest {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequest) SetReset(v bool) *SaveAnnotationMissionTagInfoListRequest {
	s.Reset = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequest) Validate() error {
	if s.AnnotationMissionTagInfoList != nil {
		for _, item := range s.AnnotationMissionTagInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList struct {
	// Tag Description
	//
	// example:
	//
	// 标签描述
	AnnotationMissionTagInfoDescription *string `json:"AnnotationMissionTagInfoDescription,omitempty" xml:"AnnotationMissionTagInfoDescription,omitempty"`
	// tag id
	//
	// example:
	//
	// bdbff6a5-1f68-4b41-8d37-6ff805ce165a
	AnnotationMissionTagInfoId *string `json:"AnnotationMissionTagInfoId,omitempty" xml:"AnnotationMissionTagInfoId,omitempty"`
	// Tag Name
	//
	// example:
	//
	// 标签
	AnnotationMissionTagInfoName *string `json:"AnnotationMissionTagInfoName,omitempty" xml:"AnnotationMissionTagInfoName,omitempty"`
	// Indicates whether to delete the tag.
	//
	// > Set this parameter to true to delete the tag or to false to add the tag.
	//
	// example:
	//
	// true
	Delete *bool `json:"Delete,omitempty" xml:"Delete,omitempty"`
	// Instance ID
	//
	// example:
	//
	// bdbff6a5-1f68-4b41-8d37-6ff805ce165a
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Tenant ID.
	//
	// example:
	//
	// bdbff6a5-1f68-4b41-8d37-6ff805ce165a
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) String() string {
	return dara.Prettify(s)
}

func (s SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GoString() string {
	return s.String()
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoDescription() *string {
	return s.AnnotationMissionTagInfoDescription
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoId() *string {
	return s.AnnotationMissionTagInfoId
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetAnnotationMissionTagInfoName() *string {
	return s.AnnotationMissionTagInfoName
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetDelete() *bool {
	return s.Delete
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoDescription(v string) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoDescription = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoId(v string) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoId = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetAnnotationMissionTagInfoName(v string) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.AnnotationMissionTagInfoName = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetDelete(v bool) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.Delete = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetInstanceId(v string) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.InstanceId = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) SetTenantId(v string) *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList {
	s.TenantId = &v
	return s
}

func (s *SaveAnnotationMissionTagInfoListRequestAnnotationMissionTagInfoList) Validate() error {
	return dara.Validate(s)
}
