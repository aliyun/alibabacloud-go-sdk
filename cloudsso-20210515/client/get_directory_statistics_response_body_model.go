// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryStatisticsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryStatistics(v *GetDirectoryStatisticsResponseBodyDirectoryStatistics) *GetDirectoryStatisticsResponseBody
	GetDirectoryStatistics() *GetDirectoryStatisticsResponseBodyDirectoryStatistics
	SetRequestId(v string) *GetDirectoryStatisticsResponseBody
	GetRequestId() *string
}

type GetDirectoryStatisticsResponseBody struct {
	// The statistics of the directory.
	DirectoryStatistics *GetDirectoryStatisticsResponseBodyDirectoryStatistics `json:"DirectoryStatistics,omitempty" xml:"DirectoryStatistics,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B7228B0-A435-5D27-A6B2-ED3571F0654B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryStatisticsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponseBody) GetDirectoryStatistics() *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	return s.DirectoryStatistics
}

func (s *GetDirectoryStatisticsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectoryStatisticsResponseBody) SetDirectoryStatistics(v *GetDirectoryStatisticsResponseBodyDirectoryStatistics) *GetDirectoryStatisticsResponseBody {
	s.DirectoryStatistics = v
	return s
}

func (s *GetDirectoryStatisticsResponseBody) SetRequestId(v string) *GetDirectoryStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBody) Validate() error {
	if s.DirectoryStatistics != nil {
		if err := s.DirectoryStatistics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectoryStatisticsResponseBodyDirectoryStatistics struct {
	// The number of access permissions that are assigned.
	//
	// example:
	//
	// 5
	AccessAssignmentCount *int32 `json:"AccessAssignmentCount,omitempty" xml:"AccessAssignmentCount,omitempty"`
	// The number of access configurations.
	//
	// example:
	//
	// 6
	AccessConfigurationCount *int32 `json:"AccessConfigurationCount,omitempty" xml:"AccessConfigurationCount,omitempty"`
	// The quota for access configurations.
	//
	// example:
	//
	// 1000
	AccessConfigurationQuota *int32 `json:"AccessConfigurationQuota,omitempty" xml:"AccessConfigurationQuota,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	//
	// example:
	//
	// new-example
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The number of user groups.
	//
	// example:
	//
	// 4
	GroupCount *int32 `json:"GroupCount,omitempty" xml:"GroupCount,omitempty"`
	// The quota for user groups.
	//
	// example:
	//
	// 500
	GroupQuota *int32 `json:"GroupQuota,omitempty" xml:"GroupQuota,omitempty"`
	// The number of tasks that are being performed.
	//
	// example:
	//
	// 0
	InProgressTaskCount *int32 `json:"InProgressTaskCount,omitempty" xml:"InProgressTaskCount,omitempty"`
	// The number of inline policies that can be configured for an access configuration.
	//
	// example:
	//
	// 1
	InlinePolicyPerAccessConfigurationQuota *int32 `json:"InlinePolicyPerAccessConfigurationQuota,omitempty" xml:"InlinePolicyPerAccessConfigurationQuota,omitempty"`
	// The region ID of the directory.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of SCIM credentials.
	//
	// example:
	//
	// 2
	SCIMServerCredentialCount *int32 `json:"SCIMServerCredentialCount,omitempty" xml:"SCIMServerCredentialCount,omitempty"`
	// Indicates whether SCIM synchronization is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	SCIMSyncEnabled *bool `json:"SCIMSyncEnabled,omitempty" xml:"SCIMSyncEnabled,omitempty"`
	// Indicates whether SSO is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	SSOEnabled *bool `json:"SSOEnabled,omitempty" xml:"SSOEnabled,omitempty"`
	// The quota for system policies that can be configured for an access configuration.
	//
	// example:
	//
	// 20
	SystemPolicyPerAccessConfigurationQuota *int32 `json:"SystemPolicyPerAccessConfigurationQuota,omitempty" xml:"SystemPolicyPerAccessConfigurationQuota,omitempty"`
	// The number of users.
	//
	// example:
	//
	// 16
	UserCount *int32 `json:"UserCount,omitempty" xml:"UserCount,omitempty"`
	// The quota for users.
	//
	// example:
	//
	// 1000
	UserQuota *int32 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetDirectoryStatisticsResponseBodyDirectoryStatistics) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryStatisticsResponseBodyDirectoryStatistics) GoString() string {
	return s.String()
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetAccessAssignmentCount() *int32 {
	return s.AccessAssignmentCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetAccessConfigurationCount() *int32 {
	return s.AccessConfigurationCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetAccessConfigurationQuota() *int32 {
	return s.AccessConfigurationQuota
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetGroupCount() *int32 {
	return s.GroupCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetGroupQuota() *int32 {
	return s.GroupQuota
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetInProgressTaskCount() *int32 {
	return s.InProgressTaskCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetInlinePolicyPerAccessConfigurationQuota() *int32 {
	return s.InlinePolicyPerAccessConfigurationQuota
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetRegion() *string {
	return s.Region
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetSCIMServerCredentialCount() *int32 {
	return s.SCIMServerCredentialCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetSCIMSyncEnabled() *bool {
	return s.SCIMSyncEnabled
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetSSOEnabled() *bool {
	return s.SSOEnabled
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetSystemPolicyPerAccessConfigurationQuota() *int32 {
	return s.SystemPolicyPerAccessConfigurationQuota
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetUserCount() *int32 {
	return s.UserCount
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) GetUserQuota() *int32 {
	return s.UserQuota
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessAssignmentCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessAssignmentCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessConfigurationCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessConfigurationCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetAccessConfigurationQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.AccessConfigurationQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetDirectoryId(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetDirectoryName(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.DirectoryName = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetGroupCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.GroupCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetGroupQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.GroupQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetInProgressTaskCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.InProgressTaskCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetInlinePolicyPerAccessConfigurationQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.InlinePolicyPerAccessConfigurationQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetRegion(v string) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.Region = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSCIMServerCredentialCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SCIMServerCredentialCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSCIMSyncEnabled(v bool) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SCIMSyncEnabled = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSSOEnabled(v bool) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SSOEnabled = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetSystemPolicyPerAccessConfigurationQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.SystemPolicyPerAccessConfigurationQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserCount(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserCount = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) SetUserQuota(v int32) *GetDirectoryStatisticsResponseBodyDirectoryStatistics {
	s.UserQuota = &v
	return s
}

func (s *GetDirectoryStatisticsResponseBodyDirectoryStatistics) Validate() error {
	return dara.Validate(s)
}
