// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAndroidInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAndroidInstanceIds(v []*string) *DescribeAndroidInstancesRequest
	GetAndroidInstanceIds() []*string
	SetAndroidInstanceName(v string) *DescribeAndroidInstancesRequest
	GetAndroidInstanceName() *string
	SetAppManagePolicyId(v string) *DescribeAndroidInstancesRequest
	GetAppManagePolicyId() *string
	SetAuthorizedUserId(v string) *DescribeAndroidInstancesRequest
	GetAuthorizedUserId() *string
	SetBizRegionId(v string) *DescribeAndroidInstancesRequest
	GetBizRegionId() *string
	SetChargeType(v string) *DescribeAndroidInstancesRequest
	GetChargeType() *string
	SetInstanceGroupId(v string) *DescribeAndroidInstancesRequest
	GetInstanceGroupId() *string
	SetInstanceGroupIds(v []*string) *DescribeAndroidInstancesRequest
	GetInstanceGroupIds() []*string
	SetInstanceGroupName(v string) *DescribeAndroidInstancesRequest
	GetInstanceGroupName() *string
	SetKeyPairId(v string) *DescribeAndroidInstancesRequest
	GetKeyPairId() *string
	SetMaxResults(v int32) *DescribeAndroidInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeAndroidInstancesRequest
	GetNextToken() *string
	SetNodeId(v string) *DescribeAndroidInstancesRequest
	GetNodeId() *string
	SetNodeName(v string) *DescribeAndroidInstancesRequest
	GetNodeName() *string
	SetOfficeSiteIds(v []*string) *DescribeAndroidInstancesRequest
	GetOfficeSiteIds() []*string
	SetPrivateIpAddress(v string) *DescribeAndroidInstancesRequest
	GetPrivateIpAddress() *string
	SetQosRuleIds(v []*string) *DescribeAndroidInstancesRequest
	GetQosRuleIds() []*string
	SetSaleMode(v string) *DescribeAndroidInstancesRequest
	GetSaleMode() *string
	SetStatus(v string) *DescribeAndroidInstancesRequest
	GetStatus() *string
	SetTag(v []*DescribeAndroidInstancesRequestTag) *DescribeAndroidInstancesRequest
	GetTag() []*DescribeAndroidInstancesRequestTag
}

type DescribeAndroidInstancesRequest struct {
	// The IDs of the instances.
	AndroidInstanceIds []*string `json:"AndroidInstanceIds,omitempty" xml:"AndroidInstanceIds,omitempty" type:"Repeated"`
	// The name of the instance.
	//
	// example:
	//
	// name
	AndroidInstanceName *string `json:"AndroidInstanceName,omitempty" xml:"AndroidInstanceName,omitempty"`
	AppManagePolicyId   *string `json:"AppManagePolicyId,omitempty" xml:"AppManagePolicyId,omitempty"`
	AuthorizedUserId    *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2807298.html) operation to query the regions where Cloud Phone is supported.
	//
	// example:
	//
	// cn-shanghai
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The billing method.
	//
	// Valid values:
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// 	- PrePaid: subscription.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the instance group.
	//
	// example:
	//
	// ag-25nt4kk9whjh****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The IDs of the instance groups.
	InstanceGroupIds []*string `json:"InstanceGroupIds,omitempty" xml:"InstanceGroupIds,omitempty" type:"Repeated"`
	// The name of the instance group.
	//
	// example:
	//
	// test
	InstanceGroupName *string `json:"InstanceGroupName,omitempty" xml:"InstanceGroupName,omitempty"`
	// The ID of the bound key pair.
	//
	// example:
	//
	// kp-5hh431emkpuoi****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The maximum number of entries per page. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. If the parameter is left empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kw9dGL5jves2FS9RLq****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// node_id
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The name of the node.
	//
	// example:
	//
	// node_name
	NodeName         *string   `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	OfficeSiteIds    []*string `json:"OfficeSiteIds,omitempty" xml:"OfficeSiteIds,omitempty" type:"Repeated"`
	PrivateIpAddress *string   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	QosRuleIds       []*string `json:"QosRuleIds,omitempty" xml:"QosRuleIds,omitempty" type:"Repeated"`
	// The sales mode.
	//
	// Valid values:
	//
	// 	- Instance: the standard mode.
	//
	// 	- Node: the node mode.
	//
	// example:
	//
	// Instance
	SaleMode *string `json:"SaleMode,omitempty" xml:"SaleMode,omitempty"`
	// The state of the instance.
	//
	// Valid values:
	//
	// 	- BACKUPING: The instance is being backed up.
	//
	// 	- STARTING: The instance is being started.
	//
	// 	- RUNNING: The instance group is available.
	//
	// 	- DELETING: The instance is being deleted.
	//
	// 	- BACKUP_FAILED: The backup operation failed.
	//
	// 	- DELETED: The instance is deleted.
	//
	// 	- FAILED: The instance failed to be created.
	//
	// 	- STOPPED: The instance is stopped.
	//
	// 	- RECOVERING: The instance has an ongoing file recovery task.
	//
	// 	- UNAVAILABLE: The instance has an exception.
	//
	// 	- REBOOTING: The instance is being restarted.
	//
	// 	- RESETTING: The instance is being reset.
	//
	// 	- STOPPING: The instance is being stopped.
	//
	// 	- RECOVER_FAILED: The file recovery task failed.
	//
	// 	- CREATING: The instance is being created.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resources.
	Tag []*DescribeAndroidInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAndroidInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesRequest) GetAndroidInstanceIds() []*string {
	return s.AndroidInstanceIds
}

func (s *DescribeAndroidInstancesRequest) GetAndroidInstanceName() *string {
	return s.AndroidInstanceName
}

func (s *DescribeAndroidInstancesRequest) GetAppManagePolicyId() *string {
	return s.AppManagePolicyId
}

func (s *DescribeAndroidInstancesRequest) GetAuthorizedUserId() *string {
	return s.AuthorizedUserId
}

func (s *DescribeAndroidInstancesRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *DescribeAndroidInstancesRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeAndroidInstancesRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *DescribeAndroidInstancesRequest) GetInstanceGroupIds() []*string {
	return s.InstanceGroupIds
}

func (s *DescribeAndroidInstancesRequest) GetInstanceGroupName() *string {
	return s.InstanceGroupName
}

func (s *DescribeAndroidInstancesRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DescribeAndroidInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeAndroidInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAndroidInstancesRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeAndroidInstancesRequest) GetNodeName() *string {
	return s.NodeName
}

func (s *DescribeAndroidInstancesRequest) GetOfficeSiteIds() []*string {
	return s.OfficeSiteIds
}

func (s *DescribeAndroidInstancesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *DescribeAndroidInstancesRequest) GetQosRuleIds() []*string {
	return s.QosRuleIds
}

func (s *DescribeAndroidInstancesRequest) GetSaleMode() *string {
	return s.SaleMode
}

func (s *DescribeAndroidInstancesRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeAndroidInstancesRequest) GetTag() []*DescribeAndroidInstancesRequestTag {
	return s.Tag
}

func (s *DescribeAndroidInstancesRequest) SetAndroidInstanceIds(v []*string) *DescribeAndroidInstancesRequest {
	s.AndroidInstanceIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetAndroidInstanceName(v string) *DescribeAndroidInstancesRequest {
	s.AndroidInstanceName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetAppManagePolicyId(v string) *DescribeAndroidInstancesRequest {
	s.AppManagePolicyId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetAuthorizedUserId(v string) *DescribeAndroidInstancesRequest {
	s.AuthorizedUserId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetBizRegionId(v string) *DescribeAndroidInstancesRequest {
	s.BizRegionId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetChargeType(v string) *DescribeAndroidInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupId(v string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupIds(v []*string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetInstanceGroupName(v string) *DescribeAndroidInstancesRequest {
	s.InstanceGroupName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetKeyPairId(v string) *DescribeAndroidInstancesRequest {
	s.KeyPairId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetMaxResults(v int32) *DescribeAndroidInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNextToken(v string) *DescribeAndroidInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNodeId(v string) *DescribeAndroidInstancesRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetNodeName(v string) *DescribeAndroidInstancesRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetOfficeSiteIds(v []*string) *DescribeAndroidInstancesRequest {
	s.OfficeSiteIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetPrivateIpAddress(v string) *DescribeAndroidInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetQosRuleIds(v []*string) *DescribeAndroidInstancesRequest {
	s.QosRuleIds = v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetSaleMode(v string) *DescribeAndroidInstancesRequest {
	s.SaleMode = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetStatus(v string) *DescribeAndroidInstancesRequest {
	s.Status = &v
	return s
}

func (s *DescribeAndroidInstancesRequest) SetTag(v []*DescribeAndroidInstancesRequestTag) *DescribeAndroidInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeAndroidInstancesRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeAndroidInstancesRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAndroidInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAndroidInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAndroidInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAndroidInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAndroidInstancesRequestTag) SetKey(v string) *DescribeAndroidInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAndroidInstancesRequestTag) SetValue(v string) *DescribeAndroidInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAndroidInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
