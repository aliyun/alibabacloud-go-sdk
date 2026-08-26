// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkspaceResponseBody
	GetCode() *string
	SetData(v *CreateWorkspaceResponseBodyData) *CreateWorkspaceResponseBody
	GetData() *CreateWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkspaceResponseBody
	GetSuccess() *bool
}

type CreateWorkspaceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The workspace details.
	Data *CreateWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkspaceResponseBody) GetData() *CreateWorkspaceResponseBodyData {
	return s.Data
}

func (s *CreateWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkspaceResponseBody) SetCode(v string) *CreateWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetData(v *CreateWorkspaceResponseBodyData) *CreateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *CreateWorkspaceResponseBody) SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetMessage(v string) *CreateWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyData struct {
	// The workspace name.
	//
	// example:
	//
	// production-agents
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The network configuration of the workspace.
	NetworkConfiguration *CreateWorkspaceResponseBodyDataNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The workspace status.
	//
	// example:
	//
	// Initialized
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s CreateWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceResponseBodyData) GetNetworkConfiguration() *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *CreateWorkspaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWorkspaceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateWorkspaceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBodyData) SetName(v string) *CreateWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetNetworkConfiguration(v *CreateWorkspaceResponseBodyDataNetworkConfiguration) *CreateWorkspaceResponseBodyData {
	s.NetworkConfiguration = v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetRegionId(v string) *CreateWorkspaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetStatus(v string) *CreateWorkspaceResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetTenantId(v string) *CreateWorkspaceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) SetWorkspaceId(v string) *CreateWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyData) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyDataNetworkConfiguration struct {
	// The VPC network configuration of the user.
	Vpc *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s CreateWorkspaceResponseBodyDataNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyDataNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) GetVpc() *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	return s.Vpc
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) SetVpc(v *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) *CreateWorkspaceResponseBodyDataNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyDataNetworkConfigurationVpc struct {
	// Specifies whether to enable VPC networking.
	//
	// example:
	//
	// true
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The list of vSwitch IDs.
	VSwitchIds []*string `json:"vSwitchIds,omitempty" xml:"vSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC.
	//
	// example:
	//
	// vpc-bp1234567890
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetEnabled(v bool) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVSwitchIds(v []*string) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVpcId(v string) *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyDataNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}
