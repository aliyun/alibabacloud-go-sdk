// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSourceServersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *DescribeSourceServersRequest
	GetJobId() *string
	SetName(v string) *DescribeSourceServersRequest
	GetName() *string
	SetOwnerId(v int64) *DescribeSourceServersRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSourceServersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSourceServersRequest
	GetPageSize() *int32
	SetRelatedJobType(v []*string) *DescribeSourceServersRequest
	GetRelatedJobType() []*string
	SetResourceGroupId(v string) *DescribeSourceServersRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSourceServersRequest
	GetResourceOwnerAccount() *string
	SetSourceId(v []*string) *DescribeSourceServersRequest
	GetSourceId() []*string
	SetState(v string) *DescribeSourceServersRequest
	GetState() *string
	SetTag(v []*DescribeSourceServersRequestTag) *DescribeSourceServersRequest
	GetTag() []*DescribeSourceServersRequestTag
	SetWorkgroupId(v string) *DescribeSourceServersRequest
	GetWorkgroupId() *string
}

type DescribeSourceServersRequest struct {
	// The ID of the migration job.
	//
	// example:
	//
	// j-bp19vlwm0tyigbmj****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The name of the migration source. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// testSourceServerName
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Minimum value: 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of migration job that is associated with the migration source.
	RelatedJobType []*string `json:"RelatedJobType,omitempty" xml:"RelatedJobType,omitempty" type:"Repeated"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmw3ty5y7****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The migration source ID. You can specify multiple IDs.
	//
	// example:
	//
	// s-bp1e2fsl57knvuug****
	SourceId []*string `json:"SourceId,omitempty" xml:"SourceId,omitempty" type:"Repeated"`
	// The state of the migration source. Valid values:
	//
	// 	- Unavailable: The migration source is inactive, or an error occurs in the migration source.
	//
	// 	- Available: The migration source is active.
	//
	// 	- InUse: The migration source is being migrated.
	//
	// 	- Deleting: The migration source is being deleted from Server Migration Center (SMC).
	//
	// example:
	//
	// Available
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The tag.
	Tag []*DescribeSourceServersRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The workgroup ID.
	//
	// example:
	//
	// w-bp1ja22kdqphehlj****
	WorkgroupId *string `json:"WorkgroupId,omitempty" xml:"WorkgroupId,omitempty"`
}

func (s DescribeSourceServersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersRequest) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSourceServersRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSourceServersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSourceServersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSourceServersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSourceServersRequest) GetRelatedJobType() []*string {
	return s.RelatedJobType
}

func (s *DescribeSourceServersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSourceServersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSourceServersRequest) GetSourceId() []*string {
	return s.SourceId
}

func (s *DescribeSourceServersRequest) GetState() *string {
	return s.State
}

func (s *DescribeSourceServersRequest) GetTag() []*DescribeSourceServersRequestTag {
	return s.Tag
}

func (s *DescribeSourceServersRequest) GetWorkgroupId() *string {
	return s.WorkgroupId
}

func (s *DescribeSourceServersRequest) SetJobId(v string) *DescribeSourceServersRequest {
	s.JobId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetName(v string) *DescribeSourceServersRequest {
	s.Name = &v
	return s
}

func (s *DescribeSourceServersRequest) SetOwnerId(v int64) *DescribeSourceServersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetPageNumber(v int32) *DescribeSourceServersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSourceServersRequest) SetPageSize(v int32) *DescribeSourceServersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSourceServersRequest) SetRelatedJobType(v []*string) *DescribeSourceServersRequest {
	s.RelatedJobType = v
	return s
}

func (s *DescribeSourceServersRequest) SetResourceGroupId(v string) *DescribeSourceServersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSourceServersRequest) SetResourceOwnerAccount(v string) *DescribeSourceServersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSourceServersRequest) SetSourceId(v []*string) *DescribeSourceServersRequest {
	s.SourceId = v
	return s
}

func (s *DescribeSourceServersRequest) SetState(v string) *DescribeSourceServersRequest {
	s.State = &v
	return s
}

func (s *DescribeSourceServersRequest) SetTag(v []*DescribeSourceServersRequestTag) *DescribeSourceServersRequest {
	s.Tag = v
	return s
}

func (s *DescribeSourceServersRequest) SetWorkgroupId(v string) *DescribeSourceServersRequest {
	s.WorkgroupId = &v
	return s
}

func (s *DescribeSourceServersRequest) Validate() error {
	return dara.Validate(s)
}

type DescribeSourceServersRequestTag struct {
	// The key of tag N that is attached to the SMC resource. Valid values of N: 1 to 20.
	//
	// You cannot specify an empty string as a tag key. The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is attached to the SMC resource. Valid values of N: 1 to 20.
	//
	// You can specify an empty string as a tag key. The tag value can be up to 64 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSourceServersRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSourceServersRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSourceServersRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSourceServersRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSourceServersRequestTag) SetKey(v string) *DescribeSourceServersRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSourceServersRequestTag) SetValue(v string) *DescribeSourceServersRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSourceServersRequestTag) Validate() error {
	return dara.Validate(s)
}
