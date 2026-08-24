// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanScheduledStrategiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int64) *ListVirusScanScheduledStrategiesRequest
	GetCurrentPage() *int64
	SetPageSize(v int64) *ListVirusScanScheduledStrategiesRequest
	GetPageSize() *int64
	SetPerformanceModes(v []*string) *ListVirusScanScheduledStrategiesRequest
	GetPerformanceModes() []*string
	SetScanModes(v []*string) *ListVirusScanScheduledStrategiesRequest
	GetScanModes() []*string
	SetStatus(v string) *ListVirusScanScheduledStrategiesRequest
	GetStatus() *string
	SetStrategyIds(v []*string) *ListVirusScanScheduledStrategiesRequest
	GetStrategyIds() []*string
	SetStrategyName(v string) *ListVirusScanScheduledStrategiesRequest
	GetStrategyName() *string
	SetUserGroupId(v string) *ListVirusScanScheduledStrategiesRequest
	GetUserGroupId() *string
}

type ListVirusScanScheduledStrategiesRequest struct {
	// The page number of the current page in paging. Valid values: 1 to 10000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page in paging. Valid values: 1 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The collection of scan performance modes. Duplicate values are not allowed.
	PerformanceModes []*string `json:"PerformanceModes,omitempty" xml:"PerformanceModes,omitempty" type:"Repeated"`
	// The collection of scan path scopes. Duplicate values are not allowed.
	ScanModes []*string `json:"ScanModes,omitempty" xml:"ScanModes,omitempty" type:"Repeated"`
	// Filters policies by enabled status. Valid values:
	//
	// - **Enabled**: enabled.
	//
	// - **Disabled**: disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The collection of virus scheduled scan policy IDs. Duplicate values are not allowed.
	StrategyIds []*string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty" type:"Repeated"`
	// The policy name. Fuzzy match is supported. The name can be up to 128 characters in length and can contain Chinese characters, uppercase and lowercase letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Weekly_Scan_DevTeam
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
	// The user group ID. This parameter is used to filter policies whose effective scope includes the specified user group. You can obtain the value from:
	//
	// - [ListUserGroups](~~ListUserGroups~~): lists user groups.
	//
	// example:
	//
	// usergroup-9d4f2a7b3c1e****
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
}

func (s ListVirusScanScheduledStrategiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanScheduledStrategiesRequest) GoString() string {
	return s.String()
}

func (s *ListVirusScanScheduledStrategiesRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *ListVirusScanScheduledStrategiesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListVirusScanScheduledStrategiesRequest) GetPerformanceModes() []*string {
	return s.PerformanceModes
}

func (s *ListVirusScanScheduledStrategiesRequest) GetScanModes() []*string {
	return s.ScanModes
}

func (s *ListVirusScanScheduledStrategiesRequest) GetStatus() *string {
	return s.Status
}

func (s *ListVirusScanScheduledStrategiesRequest) GetStrategyIds() []*string {
	return s.StrategyIds
}

func (s *ListVirusScanScheduledStrategiesRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *ListVirusScanScheduledStrategiesRequest) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *ListVirusScanScheduledStrategiesRequest) SetCurrentPage(v int64) *ListVirusScanScheduledStrategiesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetPageSize(v int64) *ListVirusScanScheduledStrategiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetPerformanceModes(v []*string) *ListVirusScanScheduledStrategiesRequest {
	s.PerformanceModes = v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetScanModes(v []*string) *ListVirusScanScheduledStrategiesRequest {
	s.ScanModes = v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetStatus(v string) *ListVirusScanScheduledStrategiesRequest {
	s.Status = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetStrategyIds(v []*string) *ListVirusScanScheduledStrategiesRequest {
	s.StrategyIds = v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetStrategyName(v string) *ListVirusScanScheduledStrategiesRequest {
	s.StrategyName = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) SetUserGroupId(v string) *ListVirusScanScheduledStrategiesRequest {
	s.UserGroupId = &v
	return s
}

func (s *ListVirusScanScheduledStrategiesRequest) Validate() error {
	return dara.Validate(s)
}
