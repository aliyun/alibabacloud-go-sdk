// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExposedInstanceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *DescribeExposedInstanceListRequest
	GetAssetType() *string
	SetCspmStatus(v bool) *DescribeExposedInstanceListRequest
	GetCspmStatus() *bool
	SetCurrentPage(v int32) *DescribeExposedInstanceListRequest
	GetCurrentPage() *int32
	SetExposureComponent(v string) *DescribeExposedInstanceListRequest
	GetExposureComponent() *string
	SetExposureComponentBizType(v string) *DescribeExposedInstanceListRequest
	GetExposureComponentBizType() *string
	SetExposureIp(v string) *DescribeExposedInstanceListRequest
	GetExposureIp() *string
	SetExposurePort(v string) *DescribeExposedInstanceListRequest
	GetExposurePort() *string
	SetGroupId(v int64) *DescribeExposedInstanceListRequest
	GetGroupId() *int64
	SetHealthStatus(v bool) *DescribeExposedInstanceListRequest
	GetHealthStatus() *bool
	SetInstanceId(v string) *DescribeExposedInstanceListRequest
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeExposedInstanceListRequest
	GetInstanceName() *string
	SetPageSize(v int32) *DescribeExposedInstanceListRequest
	GetPageSize() *int32
	SetResourceDirectoryAccountId(v int64) *DescribeExposedInstanceListRequest
	GetResourceDirectoryAccountId() *int64
	SetVulStatus(v bool) *DescribeExposedInstanceListRequest
	GetVulStatus() *bool
}

type DescribeExposedInstanceListRequest struct {
	// The type of the asset. Valid values:
	//
	// 	- **0**: an Elastic Compute Service (ECS) instance.
	//
	// 	- **3**: an ApsaraDB RDS instance.
	//
	// 	- **4**: an ApsaraDB for MongoDB instance.
	//
	// 	- **5**: an ApsaraDB for Redis instance.
	//
	// example:
	//
	// 0
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Specifies whether the asset has Cloud Security Posture Management (CSPM) risks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CspmStatus *bool `json:"CspmStatus,omitempty" xml:"CspmStatus,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The server component that is exposed on the Internet.
	//
	// example:
	//
	// openssl
	ExposureComponent *string `json:"ExposureComponent,omitempty" xml:"ExposureComponent,omitempty"`
	// Expose component type.
	//
	// example:
	//
	// system_service
	ExposureComponentBizType *string `json:"ExposureComponentBizType,omitempty" xml:"ExposureComponentBizType,omitempty"`
	// The public IP address of the server or the public endpoint of the database.
	//
	// example:
	//
	// 116.12.XX.XX
	ExposureIp *string `json:"ExposureIp,omitempty" xml:"ExposureIp,omitempty"`
	// The port that is exposed on the Internet.
	//
	// example:
	//
	// 22
	ExposurePort *string `json:"ExposurePort,omitempty" xml:"ExposurePort,omitempty"`
	// The ID of the server group.
	//
	// > You can call the [DescribeAllGroups](~~DescribeAllGroups~~) operation to query the IDs of server groups.
	//
	// example:
	//
	// 9535356
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Specifies whether the asset has weak password risks. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// Unhealthy
	HealthStatus *bool `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The instance ID of the asset.
	//
	// example:
	//
	// i-bp1g6wxdwps7s9dz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the asset.
	//
	// example:
	//
	// abc_centos7.2_005
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of entries to return on each page. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The Alibaba Cloud account ID of the member in the resource directory.
	//
	// >  You can call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain the ID.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether the asset has vulnerabilities. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	VulStatus *bool `json:"VulStatus,omitempty" xml:"VulStatus,omitempty"`
}

func (s DescribeExposedInstanceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExposedInstanceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeExposedInstanceListRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeExposedInstanceListRequest) GetCspmStatus() *bool {
	return s.CspmStatus
}

func (s *DescribeExposedInstanceListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeExposedInstanceListRequest) GetExposureComponent() *string {
	return s.ExposureComponent
}

func (s *DescribeExposedInstanceListRequest) GetExposureComponentBizType() *string {
	return s.ExposureComponentBizType
}

func (s *DescribeExposedInstanceListRequest) GetExposureIp() *string {
	return s.ExposureIp
}

func (s *DescribeExposedInstanceListRequest) GetExposurePort() *string {
	return s.ExposurePort
}

func (s *DescribeExposedInstanceListRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *DescribeExposedInstanceListRequest) GetHealthStatus() *bool {
	return s.HealthStatus
}

func (s *DescribeExposedInstanceListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeExposedInstanceListRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeExposedInstanceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeExposedInstanceListRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *DescribeExposedInstanceListRequest) GetVulStatus() *bool {
	return s.VulStatus
}

func (s *DescribeExposedInstanceListRequest) SetAssetType(v string) *DescribeExposedInstanceListRequest {
	s.AssetType = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetCspmStatus(v bool) *DescribeExposedInstanceListRequest {
	s.CspmStatus = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetCurrentPage(v int32) *DescribeExposedInstanceListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureComponent(v string) *DescribeExposedInstanceListRequest {
	s.ExposureComponent = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureComponentBizType(v string) *DescribeExposedInstanceListRequest {
	s.ExposureComponentBizType = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposureIp(v string) *DescribeExposedInstanceListRequest {
	s.ExposureIp = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetExposurePort(v string) *DescribeExposedInstanceListRequest {
	s.ExposurePort = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetGroupId(v int64) *DescribeExposedInstanceListRequest {
	s.GroupId = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetHealthStatus(v bool) *DescribeExposedInstanceListRequest {
	s.HealthStatus = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetInstanceId(v string) *DescribeExposedInstanceListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetInstanceName(v string) *DescribeExposedInstanceListRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetPageSize(v int32) *DescribeExposedInstanceListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetResourceDirectoryAccountId(v int64) *DescribeExposedInstanceListRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) SetVulStatus(v bool) *DescribeExposedInstanceListRequest {
	s.VulStatus = &v
	return s
}

func (s *DescribeExposedInstanceListRequest) Validate() error {
	return dara.Validate(s)
}
