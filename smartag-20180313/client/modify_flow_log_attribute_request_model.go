// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFlowLogAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAging(v int32) *ModifyFlowLogAttributeRequest
	GetActiveAging() *int32
	SetDescription(v string) *ModifyFlowLogAttributeRequest
	GetDescription() *string
	SetFlowLogId(v string) *ModifyFlowLogAttributeRequest
	GetFlowLogId() *string
	SetInactiveAging(v int32) *ModifyFlowLogAttributeRequest
	GetInactiveAging() *int32
	SetLogstoreName(v string) *ModifyFlowLogAttributeRequest
	GetLogstoreName() *string
	SetName(v string) *ModifyFlowLogAttributeRequest
	GetName() *string
	SetNetflowServerIp(v string) *ModifyFlowLogAttributeRequest
	GetNetflowServerIp() *string
	SetNetflowServerPort(v int32) *ModifyFlowLogAttributeRequest
	GetNetflowServerPort() *int32
	SetNetflowVersion(v string) *ModifyFlowLogAttributeRequest
	GetNetflowVersion() *string
	SetOutputType(v string) *ModifyFlowLogAttributeRequest
	GetOutputType() *string
	SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyFlowLogAttributeRequest
	GetOwnerId() *int64
	SetProjectName(v string) *ModifyFlowLogAttributeRequest
	GetProjectName() *string
	SetRegionId(v string) *ModifyFlowLogAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyFlowLogAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyFlowLogAttributeRequest
	GetResourceOwnerId() *int64
	SetSlsRegionId(v string) *ModifyFlowLogAttributeRequest
	GetSlsRegionId() *string
}

type ModifyFlowLogAttributeRequest struct {
	// The interval at which log data of active network connections is collected. Default value: **300**. Unit: seconds.
	//
	// example:
	//
	// 300
	ActiveAging *int32 `json:"ActiveAging,omitempty" xml:"ActiveAging,omitempty"`
	// The description of the flow log.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the flow log.
	//
	// This parameter is required.
	//
	// example:
	//
	// fl-7a56mar1kfw9vj****
	FlowLogId *string `json:"FlowLogId,omitempty" xml:"FlowLogId,omitempty"`
	// The interval at which log data of inactive network connections is collected. Default value: **15**. Unit: seconds.
	//
	// example:
	//
	// 15
	InactiveAging *int32 `json:"InactiveAging,omitempty" xml:"InactiveAging,omitempty"`
	// The Logstore of Log Service. This parameter is required when OutputType is set to **sls**.
	//
	// example:
	//
	// ssfghgh
	LogstoreName *string `json:"LogstoreName,omitempty" xml:"LogstoreName,omitempty"`
	// The name of the flow log.
	//
	// example:
	//
	// DDE
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The IP address of the NetFlow collector where the flow log is stored. This parameter is required when OutputType is set to **netflow**.
	//
	// example:
	//
	// 192.168.0.2
	NetflowServerIp *string `json:"NetflowServerIp,omitempty" xml:"NetflowServerIp,omitempty"`
	// The port of the NetFlow collector. Default value: **9995**. This parameter is required when OutputType is set to **netflow**.
	//
	// example:
	//
	// 9995
	NetflowServerPort *int32 `json:"NetflowServerPort,omitempty" xml:"NetflowServerPort,omitempty"`
	// The NetFlow version. Valid values: **V5**, **V9**, and **V10**. Default value: **V9**. This parameter is required when OutputType is set to **netflow**.
	//
	// example:
	//
	// V9
	NetflowVersion *string `json:"NetflowVersion,omitempty" xml:"NetflowVersion,omitempty"`
	// The location where the flow log is stored. Valid values:
	//
	// 	- **sls**: The flow log is stored in Log Service.
	//
	// 	- **netflow**: The flow log is stored on a NetFlow collector.
	//
	// 	- **all**: The flow log is stored both in Log Service and on a NetFlow collector.
	//
	// example:
	//
	// sls
	OutputType   *string `json:"OutputType,omitempty" xml:"OutputType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project to which the Logstore of Log Service belongs. This parameter is required when OutputType is set to **sls**.
	//
	// example:
	//
	// ddrrgt
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where the flow log is deployed.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the region where Log Service is deployed. This parameter is required when OutputType is set to **sls**.
	//
	// example:
	//
	// cn-shanghai
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s ModifyFlowLogAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFlowLogAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyFlowLogAttributeRequest) GetActiveAging() *int32 {
	return s.ActiveAging
}

func (s *ModifyFlowLogAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFlowLogAttributeRequest) GetFlowLogId() *string {
	return s.FlowLogId
}

func (s *ModifyFlowLogAttributeRequest) GetInactiveAging() *int32 {
	return s.InactiveAging
}

func (s *ModifyFlowLogAttributeRequest) GetLogstoreName() *string {
	return s.LogstoreName
}

func (s *ModifyFlowLogAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyFlowLogAttributeRequest) GetNetflowServerIp() *string {
	return s.NetflowServerIp
}

func (s *ModifyFlowLogAttributeRequest) GetNetflowServerPort() *int32 {
	return s.NetflowServerPort
}

func (s *ModifyFlowLogAttributeRequest) GetNetflowVersion() *string {
	return s.NetflowVersion
}

func (s *ModifyFlowLogAttributeRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *ModifyFlowLogAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyFlowLogAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyFlowLogAttributeRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *ModifyFlowLogAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyFlowLogAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyFlowLogAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyFlowLogAttributeRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *ModifyFlowLogAttributeRequest) SetActiveAging(v int32) *ModifyFlowLogAttributeRequest {
	s.ActiveAging = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetDescription(v string) *ModifyFlowLogAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetFlowLogId(v string) *ModifyFlowLogAttributeRequest {
	s.FlowLogId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetInactiveAging(v int32) *ModifyFlowLogAttributeRequest {
	s.InactiveAging = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetLogstoreName(v string) *ModifyFlowLogAttributeRequest {
	s.LogstoreName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetName(v string) *ModifyFlowLogAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetNetflowServerIp(v string) *ModifyFlowLogAttributeRequest {
	s.NetflowServerIp = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetNetflowServerPort(v int32) *ModifyFlowLogAttributeRequest {
	s.NetflowServerPort = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetNetflowVersion(v string) *ModifyFlowLogAttributeRequest {
	s.NetflowVersion = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOutputType(v string) *ModifyFlowLogAttributeRequest {
	s.OutputType = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetProjectName(v string) *ModifyFlowLogAttributeRequest {
	s.ProjectName = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetRegionId(v string) *ModifyFlowLogAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerAccount(v string) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetResourceOwnerId(v int64) *ModifyFlowLogAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) SetSlsRegionId(v string) *ModifyFlowLogAttributeRequest {
	s.SlsRegionId = &v
	return s
}

func (s *ModifyFlowLogAttributeRequest) Validate() error {
	return dara.Validate(s)
}
