// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreCheckItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckAndRiskTypeList(v []*IgnoreCheckItemsRequestCheckAndRiskTypeList) *IgnoreCheckItemsRequest
	GetCheckAndRiskTypeList() []*IgnoreCheckItemsRequestCheckAndRiskTypeList
	SetCheckIds(v []*int64) *IgnoreCheckItemsRequest
	GetCheckIds() []*int64
	SetContainerItems(v []*IgnoreCheckItemsRequestContainerItems) *IgnoreCheckItemsRequest
	GetContainerItems() []*IgnoreCheckItemsRequestContainerItems
	SetLang(v string) *IgnoreCheckItemsRequest
	GetLang() *string
	SetReason(v string) *IgnoreCheckItemsRequest
	GetReason() *string
	SetSource(v string) *IgnoreCheckItemsRequest
	GetSource() *string
	SetType(v int32) *IgnoreCheckItemsRequest
	GetType() *int32
	SetUuidList(v []*string) *IgnoreCheckItemsRequest
	GetUuidList() []*string
}

type IgnoreCheckItemsRequest struct {
	// The information about check items.
	CheckAndRiskTypeList []*IgnoreCheckItemsRequestCheckAndRiskTypeList `json:"CheckAndRiskTypeList,omitempty" xml:"CheckAndRiskTypeList,omitempty" type:"Repeated"`
	// The IDs of check items.
	CheckIds []*int64 `json:"CheckIds,omitempty" xml:"CheckIds,omitempty" type:"Repeated"`
	// List of container names that need to be whitelisted.
	ContainerItems []*IgnoreCheckItemsRequestContainerItems `json:"ContainerItems,omitempty" xml:"ContainerItems,omitempty" type:"Repeated"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The reason why you add the risk item to the whitelist.
	//
	// example:
	//
	// already config in another way
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The data source. Valid values:
	//
	// 	- **default**: host baseline
	//
	// 	- **agentless**: agentless baseline
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The operation that you want to perform on the risk item.Valid values:
	//
	// 	- **1**: adds the risk item to the whitelist
	//
	// 	- **2**: removes the risk item from the whitelist
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUIDs of the servers.
	//
	// > You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to query the UUIDs of servers.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s IgnoreCheckItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreCheckItemsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreCheckItemsRequest) GetCheckAndRiskTypeList() []*IgnoreCheckItemsRequestCheckAndRiskTypeList {
	return s.CheckAndRiskTypeList
}

func (s *IgnoreCheckItemsRequest) GetCheckIds() []*int64 {
	return s.CheckIds
}

func (s *IgnoreCheckItemsRequest) GetContainerItems() []*IgnoreCheckItemsRequestContainerItems {
	return s.ContainerItems
}

func (s *IgnoreCheckItemsRequest) GetLang() *string {
	return s.Lang
}

func (s *IgnoreCheckItemsRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreCheckItemsRequest) GetSource() *string {
	return s.Source
}

func (s *IgnoreCheckItemsRequest) GetType() *int32 {
	return s.Type
}

func (s *IgnoreCheckItemsRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *IgnoreCheckItemsRequest) SetCheckAndRiskTypeList(v []*IgnoreCheckItemsRequestCheckAndRiskTypeList) *IgnoreCheckItemsRequest {
	s.CheckAndRiskTypeList = v
	return s
}

func (s *IgnoreCheckItemsRequest) SetCheckIds(v []*int64) *IgnoreCheckItemsRequest {
	s.CheckIds = v
	return s
}

func (s *IgnoreCheckItemsRequest) SetContainerItems(v []*IgnoreCheckItemsRequestContainerItems) *IgnoreCheckItemsRequest {
	s.ContainerItems = v
	return s
}

func (s *IgnoreCheckItemsRequest) SetLang(v string) *IgnoreCheckItemsRequest {
	s.Lang = &v
	return s
}

func (s *IgnoreCheckItemsRequest) SetReason(v string) *IgnoreCheckItemsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreCheckItemsRequest) SetSource(v string) *IgnoreCheckItemsRequest {
	s.Source = &v
	return s
}

func (s *IgnoreCheckItemsRequest) SetType(v int32) *IgnoreCheckItemsRequest {
	s.Type = &v
	return s
}

func (s *IgnoreCheckItemsRequest) SetUuidList(v []*string) *IgnoreCheckItemsRequest {
	s.UuidList = v
	return s
}

func (s *IgnoreCheckItemsRequest) Validate() error {
	return dara.Validate(s)
}

type IgnoreCheckItemsRequestCheckAndRiskTypeList struct {
	// The ID of the check item.
	//
	// example:
	//
	// 52
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The baseline type of the check item.
	//
	// example:
	//
	// weak_password
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s IgnoreCheckItemsRequestCheckAndRiskTypeList) String() string {
	return dara.Prettify(s)
}

func (s IgnoreCheckItemsRequestCheckAndRiskTypeList) GoString() string {
	return s.String()
}

func (s *IgnoreCheckItemsRequestCheckAndRiskTypeList) GetCheckId() *int64 {
	return s.CheckId
}

func (s *IgnoreCheckItemsRequestCheckAndRiskTypeList) GetRiskType() *string {
	return s.RiskType
}

func (s *IgnoreCheckItemsRequestCheckAndRiskTypeList) SetCheckId(v int64) *IgnoreCheckItemsRequestCheckAndRiskTypeList {
	s.CheckId = &v
	return s
}

func (s *IgnoreCheckItemsRequestCheckAndRiskTypeList) SetRiskType(v string) *IgnoreCheckItemsRequestCheckAndRiskTypeList {
	s.RiskType = &v
	return s
}

func (s *IgnoreCheckItemsRequestCheckAndRiskTypeList) Validate() error {
	return dara.Validate(s)
}

type IgnoreCheckItemsRequestContainerItems struct {
	// The names of the containers that need to be whitelisted for the current asset, separated by English commas.
	//
	// example:
	//
	// "anythingllm,ChuanhuChat"
	ContainerNames *string `json:"ContainerNames,omitempty" xml:"ContainerNames,omitempty"`
	// The UUID of the server.
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// example:
	//
	// 14eb2fb6-ab02-4869-a1e1-2cdb0f7*****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s IgnoreCheckItemsRequestContainerItems) String() string {
	return dara.Prettify(s)
}

func (s IgnoreCheckItemsRequestContainerItems) GoString() string {
	return s.String()
}

func (s *IgnoreCheckItemsRequestContainerItems) GetContainerNames() *string {
	return s.ContainerNames
}

func (s *IgnoreCheckItemsRequestContainerItems) GetUuid() *string {
	return s.Uuid
}

func (s *IgnoreCheckItemsRequestContainerItems) SetContainerNames(v string) *IgnoreCheckItemsRequestContainerItems {
	s.ContainerNames = &v
	return s
}

func (s *IgnoreCheckItemsRequestContainerItems) SetUuid(v string) *IgnoreCheckItemsRequestContainerItems {
	s.Uuid = &v
	return s
}

func (s *IgnoreCheckItemsRequestContainerItems) Validate() error {
	return dara.Validate(s)
}
