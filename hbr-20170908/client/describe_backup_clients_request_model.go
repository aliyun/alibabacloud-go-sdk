// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIds(v []*string) *DescribeBackupClientsRequest
	GetClientIds() []*string
	SetClientType(v string) *DescribeBackupClientsRequest
	GetClientType() *string
	SetClusterId(v string) *DescribeBackupClientsRequest
	GetClusterId() *string
	SetCrossAccountRoleName(v string) *DescribeBackupClientsRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DescribeBackupClientsRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DescribeBackupClientsRequest
	GetCrossAccountUserId() *int64
	SetFilters(v []*DescribeBackupClientsRequestFilters) *DescribeBackupClientsRequest
	GetFilters() []*DescribeBackupClientsRequestFilters
	SetInstanceIds(v []*string) *DescribeBackupClientsRequest
	GetInstanceIds() []*string
	SetPageNumber(v int32) *DescribeBackupClientsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeBackupClientsRequest
	GetPageSize() *int32
	SetTag(v []*DescribeBackupClientsRequestTag) *DescribeBackupClientsRequest
	GetTag() []*DescribeBackupClientsRequestTag
}

type DescribeBackupClientsRequest struct {
	// The list of backup client IDs.
	//
	// example:
	//
	// ["c-*********************"]
	ClientIds []*string `json:"ClientIds,omitempty" xml:"ClientIds,omitempty" type:"Repeated"`
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
	Filters []*DescribeBackupClientsRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// The list of ECS instance IDs.
	//
	// example:
	//
	// ["i-*********************"]
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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
	Tag []*DescribeBackupClientsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequest) GetClientIds() []*string {
	return s.ClientIds
}

func (s *DescribeBackupClientsRequest) GetClientType() *string {
	return s.ClientType
}

func (s *DescribeBackupClientsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeBackupClientsRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeBackupClientsRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeBackupClientsRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeBackupClientsRequest) GetFilters() []*DescribeBackupClientsRequestFilters {
	return s.Filters
}

func (s *DescribeBackupClientsRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeBackupClientsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeBackupClientsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeBackupClientsRequest) GetTag() []*DescribeBackupClientsRequestTag {
	return s.Tag
}

func (s *DescribeBackupClientsRequest) SetClientIds(v []*string) *DescribeBackupClientsRequest {
	s.ClientIds = v
	return s
}

func (s *DescribeBackupClientsRequest) SetClientType(v string) *DescribeBackupClientsRequest {
	s.ClientType = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetClusterId(v string) *DescribeBackupClientsRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountRoleName(v string) *DescribeBackupClientsRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountType(v string) *DescribeBackupClientsRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetCrossAccountUserId(v int64) *DescribeBackupClientsRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetFilters(v []*DescribeBackupClientsRequestFilters) *DescribeBackupClientsRequest {
	s.Filters = v
	return s
}

func (s *DescribeBackupClientsRequest) SetInstanceIds(v []*string) *DescribeBackupClientsRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeBackupClientsRequest) SetPageNumber(v int32) *DescribeBackupClientsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetPageSize(v int32) *DescribeBackupClientsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBackupClientsRequest) SetTag(v []*DescribeBackupClientsRequestTag) *DescribeBackupClientsRequest {
	s.Tag = v
	return s
}

func (s *DescribeBackupClientsRequest) Validate() error {
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

type DescribeBackupClientsRequestFilters struct {
	// The key of the query filter.
	//
	// example:
	//
	// InstanceId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The values to match in the query filter.
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s DescribeBackupClientsRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupClientsRequestFilters) GetValues() []*string {
	return s.Values
}

func (s *DescribeBackupClientsRequestFilters) SetKey(v string) *DescribeBackupClientsRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsRequestFilters) SetValues(v []*string) *DescribeBackupClientsRequestFilters {
	s.Values = v
	return s
}

func (s *DescribeBackupClientsRequestFilters) Validate() error {
	return dara.Validate(s)
}

type DescribeBackupClientsRequestTag struct {
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

func (s DescribeBackupClientsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeBackupClientsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeBackupClientsRequestTag) SetKey(v string) *DescribeBackupClientsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeBackupClientsRequestTag) SetValue(v string) *DescribeBackupClientsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeBackupClientsRequestTag) Validate() error {
	return dara.Validate(s)
}
