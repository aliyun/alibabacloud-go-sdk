// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAging(v int32) *CreateFlowLogRequest
	GetActiveAging() *int32
	SetDescription(v string) *CreateFlowLogRequest
	GetDescription() *string
	SetInactiveAging(v int32) *CreateFlowLogRequest
	GetInactiveAging() *int32
	SetLogstoreName(v string) *CreateFlowLogRequest
	GetLogstoreName() *string
	SetName(v string) *CreateFlowLogRequest
	GetName() *string
	SetNetflowServerIp(v string) *CreateFlowLogRequest
	GetNetflowServerIp() *string
	SetNetflowServerPort(v int32) *CreateFlowLogRequest
	GetNetflowServerPort() *int32
	SetNetflowVersion(v string) *CreateFlowLogRequest
	GetNetflowVersion() *string
	SetOutputType(v string) *CreateFlowLogRequest
	GetOutputType() *string
	SetOwnerAccount(v string) *CreateFlowLogRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFlowLogRequest
	GetOwnerId() *int64
	SetProjectName(v string) *CreateFlowLogRequest
	GetProjectName() *string
	SetRegionId(v string) *CreateFlowLogRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateFlowLogRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFlowLogRequest
	GetResourceOwnerId() *int64
	SetSlsRegionId(v string) *CreateFlowLogRequest
	GetSlsRegionId() *string
}

type CreateFlowLogRequest struct {
	// The output interval under active connections. Valid values: **60 to 6000**. Unit: seconds. Default value: **300**.
	//
	// example:
	//
	// 300
	ActiveAging *int32 `json:"ActiveAging,omitempty" xml:"ActiveAging,omitempty"`
	// The description of the flow log.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The output interval under inactive connections. Valid values: **10 to 600**. Unit: seconds. Default value: **15**.
	//
	// example:
	//
	// 15
	InactiveAging *int32 `json:"InactiveAging,omitempty" xml:"InactiveAging,omitempty"`
	// The Logstore in Log Service.
	//
	// If OutputType is set to **sls*	- or **all**, this parameter is required.
	//
	// example:
	//
	// config-operation-log
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the flow log.
	//
	// example:
	//
	// sag-flowlog-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IP address of the NetFlow collector where the flow log is stored.
	//
	// If OutputType is set to **netflow*	- or **all**, this parameter is required.
	//
	// example:
	//
	// 192.168.0.2
	NetflowServerIp *string `json:"NetflowServerIp,omitempty" xml:"NetflowServerIp,omitempty"`
	// The port number of the NetFlow collector where the flow log is stored. Default value: **9995**.
	//
	// If OutputType is set to **netflow*	- or **all**, this parameter is required.
	//
	// example:
	//
	// 9995
	NetflowServerPort *int32 `json:"NetflowServerPort,omitempty" xml:"NetflowServerPort,omitempty"`
	// The version of the NetFlow collector where the flow log is stored. Valid values: **V5**, **V9**, and **V10**. Default value: **V9**.
	//
	// If OutputType is set to **netflow*	- or **all**, this parameter is required.
	//
	// example:
	//
	// V9
	NetflowVersion *string `json:"NetflowVersion,omitempty" xml:"NetflowVersion,omitempty"`
	// The type of the flow log. Valid values:
	//
	// 	- **sls**: The flow log is stored in Log Service.
	//
	// 	- **netflow**: The flow log is stored on a NetFlow collector.
	//
	// 	- **all**: The flow log is stored both in Log Service and on a NetFlow collector.
	//
	// example:
	//
	// all
	OutputType   *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project in Log Service.
	//
	// If OutputType is set to **sls*	- or **all**, this parameter is required.
	//
	// example:
	//
	// sag-flowlog-shanghai
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region to which the flow log belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the region where Log Service is deployed.
	//
	// If OutputType is set to **sls*	- or **all**, this parameter is required.
	//
	// example:
	//
	// cn-shanghai
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s CreateFlowLogRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowLogRequest) GoString() string {
	return s.String()
}

func (s *CreateFlowLogRequest) GetActiveAging() *int32 {
	return s.ActiveAging
}

func (s *CreateFlowLogRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFlowLogRequest) GetInactiveAging() *int32 {
	return s.InactiveAging
}

func (s *CreateFlowLogRequest) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *CreateFlowLogRequest) GetName() *string {
	return s.Name
}

func (s *CreateFlowLogRequest) GetNetflowServerIp() *string {
	return s.NetflowServerIp
}

func (s *CreateFlowLogRequest) GetNetflowServerPort() *int32 {
	return s.NetflowServerPort
}

func (s *CreateFlowLogRequest) GetNetflowVersion() *string {
	return s.NetflowVersion
}

func (s *CreateFlowLogRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *CreateFlowLogRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFlowLogRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFlowLogRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateFlowLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateFlowLogRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFlowLogRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFlowLogRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *CreateFlowLogRequest) SetActiveAging(v int32) *CreateFlowLogRequest {
	s.ActiveAging = &v
	return s
}

func (s *CreateFlowLogRequest) SetDescription(v string) *CreateFlowLogRequest {
	s.Description = &v
	return s
}

func (s *CreateFlowLogRequest) SetInactiveAging(v int32) *CreateFlowLogRequest {
	s.InactiveAging = &v
	return s
}

func (s *CreateFlowLogRequest) SetLogstoreName(v string) *CreateFlowLogRequest {
	s.LogstoreName = &v
	return s
}

func (s *CreateFlowLogRequest) SetName(v string) *CreateFlowLogRequest {
	s.Name = &v
	return s
}

func (s *CreateFlowLogRequest) SetNetflowServerIp(v string) *CreateFlowLogRequest {
	s.NetflowServerIp = &v
	return s
}

func (s *CreateFlowLogRequest) SetNetflowServerPort(v int32) *CreateFlowLogRequest {
	s.NetflowServerPort = &v
	return s
}

func (s *CreateFlowLogRequest) SetNetflowVersion(v string) *CreateFlowLogRequest {
	s.NetflowVersion = &v
	return s
}

func (s *CreateFlowLogRequest) SetOutputType(v string) *CreateFlowLogRequest {
	s.OutputType = &v
	return s
}

func (s *CreateFlowLogRequest) SetOwnerAccount(v string) *CreateFlowLogRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFlowLogRequest) SetOwnerId(v int64) *CreateFlowLogRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFlowLogRequest) SetProjectName(v string) *CreateFlowLogRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateFlowLogRequest) SetRegionId(v string) *CreateFlowLogRequest {
	s.RegionId = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceOwnerAccount(v string) *CreateFlowLogRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFlowLogRequest) SetResourceOwnerId(v int64) *CreateFlowLogRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFlowLogRequest) SetSlsRegionId(v string) *CreateFlowLogRequest {
	s.SlsRegionId = &v
	return s
}

func (s *CreateFlowLogRequest) Validate() error {
	return dara.Validate(s)
}
