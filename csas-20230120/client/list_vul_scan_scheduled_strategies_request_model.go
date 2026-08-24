// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulScanScheduledStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVulScanScheduledStrategiesRequest
	GetCurrentPage() *int64
	SetMatchMode(v string) *ListVulScanScheduledStrategiesRequest
	GetMatchMode() *string
	SetPageSize(v int64) *ListVulScanScheduledStrategiesRequest
	GetPageSize() *int64
	SetStatus(v string) *ListVulScanScheduledStrategiesRequest
	GetStatus() *string
	SetStrategyIds(v []*string) *ListVulScanScheduledStrategiesRequest
	GetStrategyIds() []*string
	SetStrategyName(v string) *ListVulScanScheduledStrategiesRequest
	GetStrategyName() *string
	SetUserGroupId(v string) *ListVulScanScheduledStrategiesRequest
	GetUserGroupId() *string
}

type ListVulScanScheduledStrategiesRequest struct {
	// The page number of the current page in a paging query. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Filters by the matching mode of the effective scope. Valid values:
	//
	// - **UserGroupAll**: Takes effect for all users under the current Alibaba Cloud account.
	//
	// - **UserGroupNormal**: Takes effect only for users in specified user groups.
	//
	// example:
	//
	// UserGroupNormal
	MatchMode *string `json:"MatchMode,omitempty" xml:"MatchMode,omitempty"`
	// The number of entries per page in a paging query. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by enabled status. Valid values:
	//
	// - **Enabled**: Enabled.
	//
	// - **Disabled**: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IDs of scheduled vulnerability scan policies used for filtering. A maximum of 100 IDs can be specified. Duplicate IDs are not allowed.
	StrategyIds []*string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty" type:"Repeated"`
	// The policy name. Fuzzy match is supported. The name can be up to 128 characters in length.
	//
	// example:
	//
	// Weekly vulnerability scanning for R&D department
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The user group ID. Used to filter records whose effective scope includes the specified user group. You can obtain the value from the following operation:
	//
	// - [ListUserGroups](~~ListUserGroups~~): lists user groups.
	//
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListVulScanScheduledStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVulScanScheduledStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListVulScanScheduledStrategiesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVulScanScheduledStrategiesRequest) GetMatchMode() *string {
	return s.MatchMode
}

func (s *ListVulScanScheduledStrategiesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVulScanScheduledStrategiesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVulScanScheduledStrategiesRequest) GetStrategyIds() []*string {
	return s.StrategyIds
}

func (s *ListVulScanScheduledStrategiesRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListVulScanScheduledStrategiesRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListVulScanScheduledStrategiesRequest) SetCurrentPage(v int64) *ListVulScanScheduledStrategiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetMatchMode(v string) *ListVulScanScheduledStrategiesRequest {
	s.MatchMode = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetPageSize(v int64) *ListVulScanScheduledStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetStatus(v string) *ListVulScanScheduledStrategiesRequest {
	s.Status = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetStrategyIds(v []*string) *ListVulScanScheduledStrategiesRequest {
	s.StrategyIds = v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetStrategyName(v string) *ListVulScanScheduledStrategiesRequest {
	s.StrategyName = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) SetUserGroupId(v string) *ListVulScanScheduledStrategiesRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListVulScanScheduledStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
