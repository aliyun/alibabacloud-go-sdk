// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIdsShrink(v string) *DescribeBackupClientsShrinkRequest
	GetClientIdsShrink() *string
	SetClientType(v string) *DescribeBackupClientsShrinkRequest
	GetClientType() *string
	SetClusterId(v string) *DescribeBackupClientsShrinkRequest
	GetClusterId() *string
	SetCrossAccountRoleName(v string) *DescribeBackupClientsShrinkRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DescribeBackupClientsShrinkRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DescribeBackupClientsShrinkRequest
	GetCrossAccountUserId() *int64
	SetFilters(v []*DescribeBackupClientsShrinkRequestFilters) *DescribeBackupClientsShrinkRequest
	GetFilters() []*DescribeBackupClientsShrinkRequestFilters
	SetInstanceIdsShrink(v string) *DescribeBackupClientsShrinkRequest
	GetInstanceIdsShrink() *string
	SetPageNumber(v int32) *DescribeBackupClientsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupClientsShrinkRequest
	GetPageSize() *int32
	SetTag(v []*DescribeBackupClientsShrinkRequestTag) *DescribeBackupClientsShrinkRequest
	GetTag() []*DescribeBackupClientsShrinkRequestTag
}

type DescribeBackupClientsShrinkRequest struct {
	// The list of backup client IDs.
	//
	// example:
	//
	// ["c-*********************"]
	ClientIdsShrink *string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty"`
	// The type of the backup client. Valid values:
	//
	// - **ECS_CLIENT**: ECS File Backup client.
	//
	// - **CONTAINER_CLIENT**: container backup client.
	//
	// - **LOCAL_CLIENT**: local NAS backup, CPFS backup, archive, or data synchronization client.
	//
	// This parameter is required.
	//
	// example:
	//
	// ECS_CLIENT
	ClientType *string `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	// The ID of the backup cluster.
	//
	// example:
	//
	// cl-000ge4wa61b4d337xblq
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the RAM role created in the source account for cross-account backup.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// The type of cross-account backup. Valid values:
	//
	// - SELF_ACCOUNT: backup within the current account.
	//
	// - CROSS_ACCOUNT: cross-account backup.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source account used for cross-account backup.
	//
	// example:
	//
	// 129374672382xxxx
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
	// The query filters.
	Filters []*DescribeBackupClientsShrinkRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The list of ECS instance IDs.
	//
	// example:
	//
	// ["i-*********************"]
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Minimum value: 1. Maximum value: 99. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The tag information to return.
	//
	// example:
	//
	// 33738719#
	Tag []*DescribeBackupClientsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsShrinkRequest) GetClientIdsShrink() *string {
	return s.ClientIdsShrink
}

func (s *DescribeBackupClientsShrinkRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeBackupClientsShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupClientsShrinkRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeBackupClientsShrinkRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeBackupClientsShrinkRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeBackupClientsShrinkRequest) GetFilters() []*DescribeBackupClientsShrinkRequestFilters {
	return s.Filters
}

func (s *DescribeBackupClientsShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeBackupClientsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupClientsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupClientsShrinkRequest) GetTag() []*DescribeBackupClientsShrinkRequestTag {
	return s.Tag
}

func (s *DescribeBackupClientsShrinkRequest) SetClientIdsShrink(v string) *DescribeBackupClientsShrinkRequest {
	s.ClientIdsShrink = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetClientType(v string) *DescribeBackupClientsShrinkRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetClusterId(v string) *DescribeBackupClientsShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountRoleName(v string) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountType(v string) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetCrossAccountUserId(v int64) *DescribeBackupClientsShrinkRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetFilters(v []*DescribeBackupClientsShrinkRequestFilters) *DescribeBackupClientsShrinkRequest {
	s.Filters = v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetInstanceIdsShrink(v string) *DescribeBackupClientsShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetPageNumber(v int32) *DescribeBackupClientsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetPageSize(v int32) *DescribeBackupClientsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) SetTag(v []*DescribeBackupClientsShrinkRequestTag) *DescribeBackupClientsShrinkRequest {
	s.Tag = v
	return s
}

func (s *DescribeBackupClientsShrinkRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeBackupClientsShrinkRequestFilters struct {
	// The key of the query filter.
	//
	// example:
	//
	// InstanceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values to match in the query filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsShrinkRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsShrinkRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsShrinkRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupClientsShrinkRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribeBackupClientsShrinkRequestFilters) SetKey(v string) *DescribeBackupClientsShrinkRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequestFilters) SetValues(v []*string) *DescribeBackupClientsShrinkRequestFilters {
	s.Values = v
	return s
}

func (s *DescribeBackupClientsShrinkRequestFilters) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupClientsShrinkRequestTag struct {
	// The tag key of the backup vault. Valid values of N: 1 to 20.
	//
	// - The tag key cannot start with `aliyun` or `acs:`.
	//
	// - The tag key cannot contain `http://` or `https://`.
	//
	// - The tag key cannot be an empty string.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value of the backup vault. Valid values of N: 1 to 20.
	//
	// - The tag value cannot start with `aliyun` or `acs:`.
	//
	// - The tag value cannot contain `http://` or `https://`.
	//
	// - The tag value cannot be an empty string.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeBackupClientsShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupClientsShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeBackupClientsShrinkRequestTag) SetKey(v string) *DescribeBackupClientsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequestTag) SetValue(v string) *DescribeBackupClientsShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeBackupClientsShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
