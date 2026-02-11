// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCloudSiemAssetsResponseBody
	GetCode() *int32
	SetData(v *DescribeCloudSiemAssetsResponseBodyData) *DescribeCloudSiemAssetsResponseBody
	GetData() *DescribeCloudSiemAssetsResponseBodyData
	SetMessage(v string) *DescribeCloudSiemAssetsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudSiemAssetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudSiemAssetsResponseBody
	GetSuccess() *bool
}

type DescribeCloudSiemAssetsResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data *DescribeCloudSiemAssetsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
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

func (s DescribeCloudSiemAssetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCloudSiemAssetsResponseBody) GetData() *DescribeCloudSiemAssetsResponseBodyData {
	return s.Data
}

func (s *DescribeCloudSiemAssetsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudSiemAssetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudSiemAssetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudSiemAssetsResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetData(v *DescribeCloudSiemAssetsResponseBodyData) *DescribeCloudSiemAssetsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudSiemAssetsResponseBodyData struct {
	// The pagination information.
	PageInfo *DescribeCloudSiemAssetsResponseBodyDataPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The detailed data.
	ResponseData []*DescribeCloudSiemAssetsResponseBodyDataResponseData `json:"ResponseData,omitempty" xml:"ResponseData,omitempty" type:"Repeated"`
}

func (s DescribeCloudSiemAssetsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyData) GetPageInfo() *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	return s.PageInfo
}

func (s *DescribeCloudSiemAssetsResponseBodyData) GetResponseData() []*DescribeCloudSiemAssetsResponseBodyDataResponseData {
	return s.ResponseData
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetPageInfo(v *DescribeCloudSiemAssetsResponseBodyDataPageInfo) *DescribeCloudSiemAssetsResponseBodyData {
	s.PageInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyData) SetResponseData(v []*DescribeCloudSiemAssetsResponseBodyDataResponseData) *DescribeCloudSiemAssetsResponseBodyData {
	s.ResponseData = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyData) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResponseData != nil {
		for _, item := range s.ResponseData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudSiemAssetsResponseBodyDataPageInfo struct {
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetCurrentPage(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetPageSize(v int32) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) SetTotalCount(v int64) *DescribeCloudSiemAssetsResponseBodyDataPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemAssetsResponseBodyDataResponseData struct {
	// The UUID of the alert associated with the event.
	//
	// example:
	//
	// sas_71e24437d2797ce8fc59692905a4****
	AlertUuid *string `json:"AlertUuid,omitempty" xml:"AlertUuid,omitempty"`
	// The ID of the Alibaba Cloud account in SIEM.
	//
	// example:
	//
	// 1276085894174392
	Aliuid *int64 `json:"Aliuid,omitempty" xml:"Aliuid,omitempty"`
	// The logical ID of the asset.
	//
	// example:
	//
	// 0616caeb-acb8-45e0-8520-4ee5fbe251f0
	AssetId *string `json:"AssetId,omitempty" xml:"AssetId,omitempty"`
	// The display information of the asset is in the JSON format.
	//
	// example:
	//
	// [{"KeyName": "${aliyun.siem.asset.asset_name}","Values": "zsw-agentless-ubuntu20","Key": "asset_name"}]
	AssetInfo []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo `json:"AssetInfo,omitempty" xml:"AssetInfo,omitempty" type:"Repeated"`
	// The name of the asset.
	//
	// example:
	//
	// zsw-agentless-centos****
	AssetName *string `json:"AssetName,omitempty" xml:"AssetName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// domain
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The cloud code of the entity. Valid values:
	//
	// 	- aliyun: Alibaba Cloud
	//
	// 	- qcloud: Tencent Cloud
	//
	// 	- hcloud: Huawei Cloud
	//
	// example:
	//
	// aliyun
	CloudCode *string `json:"CloudCode,omitempty" xml:"CloudCode,omitempty"`
	// The time when the asset was synchronized.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the asset was last updated.
	//
	// example:
	//
	// 2021-01-06 16:37:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the asset.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The UUID of the event.
	//
	// example:
	//
	// 85ea4241-798f-4684-a876-65d4f0c3****
	IncidentUuid *string `json:"IncidentUuid,omitempty" xml:"IncidentUuid,omitempty"`
	// The ID of the associated account to which the asset belongs.
	//
	// example:
	//
	// 176555323***
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAlertUuid() *string {
	return s.AlertUuid
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAliuid() *int64 {
	return s.Aliuid
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAssetId() *string {
	return s.AssetId
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAssetInfo() []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	return s.AssetInfo
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAssetName() *string {
	return s.AssetName
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetCloudCode() *string {
	return s.CloudCode
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetId() *int64 {
	return s.Id
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetIncidentUuid() *string {
	return s.IncidentUuid
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAlertUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AlertUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAliuid(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Aliuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetId(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetInfo(v []*DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetInfo = v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetAssetType(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetCloudCode(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.CloudCode = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtCreate(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtCreate = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetGmtModified(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.GmtModified = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.Id = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetIncidentUuid(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.IncidentUuid = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) SetSubUserId(v int64) *DescribeCloudSiemAssetsResponseBodyDataResponseData {
	s.SubUserId = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseData) Validate() error {
	if s.AssetInfo != nil {
		for _, item := range s.AssetInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo struct {
	// The attribute key.
	//
	// example:
	//
	// suspicious.wbd.wb.trojanpath
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// Trojan Path
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The value of the key.
	//
	// example:
	//
	// /root/test33.php
	Values *string `json:"Values,omitempty" xml:"Values,omitempty"`
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GetKey() *string {
	return s.Key
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GetKeyName() *string {
	return s.KeyName
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) GetValues() *string {
	return s.Values
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKey(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Key = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetKeyName(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.KeyName = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) SetValues(v string) *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo {
	s.Values = &v
	return s
}

func (s *DescribeCloudSiemAssetsResponseBodyDataResponseDataAssetInfo) Validate() error {
	return dara.Validate(s)
}
