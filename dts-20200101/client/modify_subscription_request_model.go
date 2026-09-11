// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbList(v string) *ModifySubscriptionRequest
	GetDbList() *string
	SetDtsInstanceId(v string) *ModifySubscriptionRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *ModifySubscriptionRequest
	GetDtsJobId() *string
	SetModifyType(v string) *ModifySubscriptionRequest
	GetModifyType() *string
	SetRegionId(v string) *ModifySubscriptionRequest
	GetRegionId() *string
	SetReserved(v string) *ModifySubscriptionRequest
	GetReserved() *string
	SetResourceGroupId(v string) *ModifySubscriptionRequest
	GetResourceGroupId() *string
	SetSubscriptionDataTypeDDL(v bool) *ModifySubscriptionRequest
	GetSubscriptionDataTypeDDL() *bool
	SetSubscriptionDataTypeDML(v bool) *ModifySubscriptionRequest
	GetSubscriptionDataTypeDML() *bool
}

type ModifySubscriptionRequest struct {
	// The modified subscription objects, in JSON format. For more information about the definition, see [Objects of DTS tasks](https://help.aliyun.com/document_detail/209545.html).
	//
	// > - The new DbList value overwrites the original DbList value. Make sure that the new DbList value contains all the objects that you want to track. Otherwise, objects may be lost. Modify this parameter with caution.
	//
	// - You can call [DescribeDtsJobDetail](https://help.aliyun.com/document_detail/208925.html) to query the original subscription objects.
	//
	// example:
	//
	// {"dtstest":{"name":"dtstest","all":true}}
	DbList *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	// The ID of the change tracking instance. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the instance ID.
	//
	// example:
	//
	// dtsboss6pn1w******
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	// The ID of the change tracking task. You can call [DescribeDtsJobs](https://help.aliyun.com/document_detail/209702.html) to query the task ID.
	//
	// example:
	//
	// boss6pn1w******
	DtsJobId   *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	// The region in which the DTS instance resides. For more information, see [Supported regions](https://help.aliyun.com/document_detail/141033.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Reserved *string `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmzawhxxc****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Specifies whether the modified task subscribes to DDL data. Valid values:
	//
	// - **true**: The task subscribes to DDL data.
	//
	// - **false**: The task does not subscribe to DDL data.
	//
	// example:
	//
	// true
	SubscriptionDataTypeDDL *bool `json:"SubscriptionDataTypeDDL,omitempty" xml:"SubscriptionDataTypeDDL,omitempty"`
	// Specifies whether the modified task subscribes to DML data. Valid values:
	//
	// - **true**: The task subscribes to DML data.
	//
	// - **false**: The task does not subscribe to DML data.
	//
	// example:
	//
	// true
	SubscriptionDataTypeDML *bool `json:"SubscriptionDataTypeDML,omitempty" xml:"SubscriptionDataTypeDML,omitempty"`
}

func (s ModifySubscriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionRequest) GetDbList() *string {
	return s.DbList
}

func (s *ModifySubscriptionRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *ModifySubscriptionRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *ModifySubscriptionRequest) GetModifyType() *string {
	return s.ModifyType
}

func (s *ModifySubscriptionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySubscriptionRequest) GetReserved() *string {
	return s.Reserved
}

func (s *ModifySubscriptionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySubscriptionRequest) GetSubscriptionDataTypeDDL() *bool {
	return s.SubscriptionDataTypeDDL
}

func (s *ModifySubscriptionRequest) GetSubscriptionDataTypeDML() *bool {
	return s.SubscriptionDataTypeDML
}

func (s *ModifySubscriptionRequest) SetDbList(v string) *ModifySubscriptionRequest {
	s.DbList = &v
	return s
}

func (s *ModifySubscriptionRequest) SetDtsInstanceId(v string) *ModifySubscriptionRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetDtsJobId(v string) *ModifySubscriptionRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetModifyType(v string) *ModifySubscriptionRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifySubscriptionRequest) SetRegionId(v string) *ModifySubscriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetReserved(v string) *ModifySubscriptionRequest {
	s.Reserved = &v
	return s
}

func (s *ModifySubscriptionRequest) SetResourceGroupId(v string) *ModifySubscriptionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetSubscriptionDataTypeDDL(v bool) *ModifySubscriptionRequest {
	s.SubscriptionDataTypeDDL = &v
	return s
}

func (s *ModifySubscriptionRequest) SetSubscriptionDataTypeDML(v bool) *ModifySubscriptionRequest {
	s.SubscriptionDataTypeDML = &v
	return s
}

func (s *ModifySubscriptionRequest) Validate() error {
	return dara.Validate(s)
}
