// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSourcesResponseBody
	GetCode() *string
	SetData(v *ListSourcesResponseBodyData) *ListSourcesResponseBody
	GetData() *ListSourcesResponseBodyData
	SetMessage(v string) *ListSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSourcesResponseBody
	GetRequestId() *string
}

type ListSourcesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned data.
	Data *ListSourcesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FA9BB94-915B-5299-A694-49FCC7F5DD00
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSourcesResponseBody) GetData() *ListSourcesResponseBodyData {
	return s.Data
}

func (s *ListSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSourcesResponseBody) SetCode(v string) *ListSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListSourcesResponseBody) SetData(v *ListSourcesResponseBodyData) *ListSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListSourcesResponseBody) SetMessage(v string) *ListSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListSourcesResponseBody) SetRequestId(v string) *ListSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSourcesResponseBodyData struct {
	// The list of sources.
	Items []*ListSourcesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBodyData) GetItems() []*ListSourcesResponseBodyDataItems {
	return s.Items
}

func (s *ListSourcesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSourcesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSourcesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListSourcesResponseBodyData) SetItems(v []*ListSourcesResponseBodyDataItems) *ListSourcesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListSourcesResponseBodyData) SetPageNumber(v int32) *ListSourcesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSourcesResponseBodyData) SetPageSize(v int32) *ListSourcesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSourcesResponseBodyData) SetTotalSize(v int32) *ListSourcesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListSourcesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSourcesResponseBodyDataItems struct {
	// example:
	//
	// Association completed
	AssociationReason *string `json:"associationReason,omitempty" xml:"associationReason,omitempty"`
	// example:
	//
	// ASSOCIATED
	AssociationStatus *string `json:"associationStatus,omitempty" xml:"associationStatus,omitempty"`
	// The creation timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// The source information when the source type is K8S.
	K8sSourceInfo *ListSourcesResponseBodyDataItemsK8sSourceInfo `json:"k8sSourceInfo,omitempty" xml:"k8sSourceInfo,omitempty" type:"Struct"`
	// The source information when the source type is MSE_NACOS.
	NacosSourceInfo *ListSourcesResponseBodyDataItemsNacosSourceInfo `json:"nacosSourceInfo,omitempty" xml:"nacosSourceInfo,omitempty" type:"Struct"`
	// The source name. If the source type is K8S, the name is the container cluster name. If the source type is MSE_NACOS, the name is the Nacos instance name.
	//
	// example:
	//
	// itemcenter-dev-cluster
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// rg-xxxx
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The source ID.
	//
	// example:
	//
	// src-crdddallhtgtria***
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The update timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1719386834548
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s ListSourcesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBodyDataItems) GetAssociationReason() *string {
	return s.AssociationReason
}

func (s *ListSourcesResponseBodyDataItems) GetAssociationStatus() *string {
	return s.AssociationStatus
}

func (s *ListSourcesResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListSourcesResponseBodyDataItems) GetK8sSourceInfo() *ListSourcesResponseBodyDataItemsK8sSourceInfo {
	return s.K8sSourceInfo
}

func (s *ListSourcesResponseBodyDataItems) GetNacosSourceInfo() *ListSourcesResponseBodyDataItemsNacosSourceInfo {
	return s.NacosSourceInfo
}

func (s *ListSourcesResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListSourcesResponseBodyDataItems) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListSourcesResponseBodyDataItems) GetSourceId() *string {
	return s.SourceId
}

func (s *ListSourcesResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListSourcesResponseBodyDataItems) SetAssociationReason(v string) *ListSourcesResponseBodyDataItems {
	s.AssociationReason = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetAssociationStatus(v string) *ListSourcesResponseBodyDataItems {
	s.AssociationStatus = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetCreateTimestamp(v int64) *ListSourcesResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetK8sSourceInfo(v *ListSourcesResponseBodyDataItemsK8sSourceInfo) *ListSourcesResponseBodyDataItems {
	s.K8sSourceInfo = v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetNacosSourceInfo(v *ListSourcesResponseBodyDataItemsNacosSourceInfo) *ListSourcesResponseBodyDataItems {
	s.NacosSourceInfo = v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetName(v string) *ListSourcesResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetResourceGroupId(v string) *ListSourcesResponseBodyDataItems {
	s.ResourceGroupId = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetSourceId(v string) *ListSourcesResponseBodyDataItems {
	s.SourceId = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListSourcesResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListSourcesResponseBodyDataItems) Validate() error {
	if s.K8sSourceInfo != nil {
		if err := s.K8sSourceInfo.Validate(); err != nil {
			return err
		}
	}
	if s.NacosSourceInfo != nil {
		if err := s.NacosSourceInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSourcesResponseBodyDataItemsK8sSourceInfo struct {
	// The cluster ID.
	//
	// example:
	//
	// c09212180612a42adbed6a940d01d***
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
}

func (s ListSourcesResponseBodyDataItemsK8sSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBodyDataItemsK8sSourceInfo) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBodyDataItemsK8sSourceInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSourcesResponseBodyDataItemsK8sSourceInfo) SetClusterId(v string) *ListSourcesResponseBodyDataItemsK8sSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *ListSourcesResponseBodyDataItemsK8sSourceInfo) Validate() error {
	return dara.Validate(s)
}

type ListSourcesResponseBodyDataItemsNacosSourceInfo struct {
	// The endpoint of the Nacos instance.
	//
	// example:
	//
	// mse-3353***-nacos-ans.mse.aliyuncs.com:8848
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// The registry ID.
	//
	// example:
	//
	// mse-3353***
	ClusterId *string `json:"clusterId,omitempty" xml:"clusterId,omitempty"`
	// The Nacos instance ID.
	//
	// example:
	//
	// mse_prepaid_public_cn-wuf***
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListSourcesResponseBodyDataItemsNacosSourceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListSourcesResponseBodyDataItemsNacosSourceInfo) GoString() string {
	return s.String()
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) GetAddress() *string {
	return s.Address
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) SetAddress(v string) *ListSourcesResponseBodyDataItemsNacosSourceInfo {
	s.Address = &v
	return s
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) SetClusterId(v string) *ListSourcesResponseBodyDataItemsNacosSourceInfo {
	s.ClusterId = &v
	return s
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) SetInstanceId(v string) *ListSourcesResponseBodyDataItemsNacosSourceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListSourcesResponseBodyDataItemsNacosSourceInfo) Validate() error {
	return dara.Validate(s)
}
