// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWorkspaceResponseBody
	GetCode() *string
	SetData(v *UpdateWorkspaceResponseBodyData) *UpdateWorkspaceResponseBody
	GetData() *UpdateWorkspaceResponseBodyData
	SetHttpStatusCode(v int32) *UpdateWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWorkspaceResponseBody
	GetSuccess() *bool
}

type UpdateWorkspaceResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The details of the updated workspace.
	Data *UpdateWorkspaceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWorkspaceResponseBody) GetData() *UpdateWorkspaceResponseBodyData {
	return s.Data
}

func (s *UpdateWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWorkspaceResponseBody) SetCode(v string) *UpdateWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetData(v *UpdateWorkspaceResponseBodyData) *UpdateWorkspaceResponseBody {
	s.Data = v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetHttpStatusCode(v int32) *UpdateWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetMessage(v string) *UpdateWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetRequestId(v string) *UpdateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) SetSuccess(v bool) *UpdateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWorkspaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceResponseBodyData struct {
	// The workspace name.
	//
	// example:
	//
	// production-agents-v2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The workspace network configuration.
	NetworkConfiguration *UpdateWorkspaceResponseBodyDataNetworkConfiguration `json:"networkConfiguration,omitempty" xml:"networkConfiguration,omitempty" type:"Struct"`
	// The region ID of the workspace.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The workspace status. Valid values: Initializing, Initialized, Deleting, Deleted.
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

func (s UpdateWorkspaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateWorkspaceResponseBodyData) GetNetworkConfiguration() *UpdateWorkspaceResponseBodyDataNetworkConfiguration {
	return s.NetworkConfiguration
}

func (s *UpdateWorkspaceResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWorkspaceResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateWorkspaceResponseBodyData) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkspaceResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateWorkspaceResponseBodyData) SetName(v string) *UpdateWorkspaceResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetNetworkConfiguration(v *UpdateWorkspaceResponseBodyDataNetworkConfiguration) *UpdateWorkspaceResponseBodyData {
	s.NetworkConfiguration = v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetRegionId(v string) *UpdateWorkspaceResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetStatus(v string) *UpdateWorkspaceResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetTenantId(v string) *UpdateWorkspaceResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) SetWorkspaceId(v string) *UpdateWorkspaceResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyData) Validate() error {
	if s.NetworkConfiguration != nil {
		if err := s.NetworkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceResponseBodyDataNetworkConfiguration struct {
	// The user VPC network configuration.
	Vpc *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc `json:"vpc,omitempty" xml:"vpc,omitempty" type:"Struct"`
}

func (s UpdateWorkspaceResponseBodyDataNetworkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyDataNetworkConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfiguration) GetVpc() *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	return s.Vpc
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfiguration) SetVpc(v *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) *UpdateWorkspaceResponseBodyDataNetworkConfiguration {
	s.Vpc = v
	return s
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfiguration) Validate() error {
	if s.Vpc != nil {
		if err := s.Vpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc struct {
	// Specifies whether to enable VPC networking.
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

func (s UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) GoString() string {
	return s.String()
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetEnabled() *bool {
	return s.Enabled
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetEnabled(v bool) *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.Enabled = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVSwitchIds(v []*string) *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VSwitchIds = v
	return s
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) SetVpcId(v string) *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc {
	s.VpcId = &v
	return s
}

func (s *UpdateWorkspaceResponseBodyDataNetworkConfigurationVpc) Validate() error {
	return dara.Validate(s)
}
