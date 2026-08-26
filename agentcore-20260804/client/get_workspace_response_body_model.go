// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkspaceResponseBody
	GetCode() *string
	SetData(v *GetWorkspaceResponseBodyData) *GetWorkspaceResponseBody
	GetData() *GetWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *GetWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWorkspaceResponseBody
	GetSuccess() *bool
}

type GetWorkspaceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The workspace details.
	Data *GetWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkspaceResponseBody) GetData() *GetWorkspaceResponseBodyData {
	return s.Data
}

func (s *GetWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWorkspaceResponseBody) SetCode(v string) *GetWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetData(v *GetWorkspaceResponseBodyData) *GetWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *GetWorkspaceResponseBody) SetHttpStatusCode(v int32) *GetWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetMessage(v string) *GetWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetSuccess(v bool) *GetWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *GetWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyData struct {
	// The creation time.
	//
	// example:
	//
	// 2026-08-06T03:56:56Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The workspace name.
	//
	// example:
	//
	// production-agents
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The workspace network configuration.
	NetworkConfiguration *GetWorkspaceResponseBodyDataNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The workspace status. Valid values: Initializing, InitializationFailed, Initialized, Deleting, Deleted.
	//
	// example:
	//
	// Initialized
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The reason for the workspace status.
	//
	// example:
	//
	// InitializationFailed: VPC not found
	StatusReason *string `json:"statusReason,omitempty" xml:"statusReason,omitempty"`
	// The ID of the tenant to which the workspace belongs.
	//
	// example:
	//
	// tenant-123456
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetWorkspaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetWorkspaceResponseBodyData) GetNetworkConfiguration() *GetWorkspaceResponseBodyDataNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *GetWorkspaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWorkspaceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetWorkspaceResponseBodyData) GetStatusReason() *string {
	return s.StatusReason
}

func (s *GetWorkspaceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *GetWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetWorkspaceResponseBodyData) SetCreateTime(v string) *GetWorkspaceResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetName(v string) *GetWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetNetworkConfiguration(v *GetWorkspaceResponseBodyDataNetworkConfiguration) *GetWorkspaceResponseBodyData {
	s.NetworkConfiguration = v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetRegionId(v string) *GetWorkspaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetStatus(v string) *GetWorkspaceResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetStatusReason(v string) *GetWorkspaceResponseBodyData {
	s.StatusReason = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetTenantId(v string) *GetWorkspaceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) SetWorkspaceId(v string) *GetWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetWorkspaceResponseBodyData) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyDataNetworkConfiguration struct {
	// The user VPC network configuration.
	Vpc *GetWorkspaceResponseBodyDataNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s GetWorkspaceResponseBodyDataNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyDataNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyDataNetworkConfiguration) GetVpc() *GetWorkspaceResponseBodyDataNetworkConfigurationVpc {
	return s.Vpc
}

func (s *GetWorkspaceResponseBodyDataNetworkConfiguration) SetVpc(v *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) *GetWorkspaceResponseBodyDataNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *GetWorkspaceResponseBodyDataNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetWorkspaceResponseBodyDataNetworkConfigurationVpc struct {
	// Indicates whether the VPC network is enabled.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the user VPC.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetWorkspaceResponseBodyDataNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s GetWorkspaceResponseBodyDataNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) SetEnabled(v bool) *GetWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVSwitchIds(v []*string) *GetWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVpcId(v string) *GetWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *GetWorkspaceResponseBodyDataNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}
