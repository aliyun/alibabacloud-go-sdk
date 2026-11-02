// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGatewayApiGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateGatewayApiGroupShrinkRequest
	GetAppId() *string
	SetCentersShrink(v string) *CreateGatewayApiGroupShrinkRequest
	GetCentersShrink() *string
	SetDescription(v string) *CreateGatewayApiGroupShrinkRequest
	GetDescription() *string
	SetDomainName(v string) *CreateGatewayApiGroupShrinkRequest
	GetDomainName() *string
	SetEnableMTLS(v bool) *CreateGatewayApiGroupShrinkRequest
	GetEnableMTLS() *bool
	SetGroupName(v string) *CreateGatewayApiGroupShrinkRequest
	GetGroupName() *string
	SetHost(v string) *CreateGatewayApiGroupShrinkRequest
	GetHost() *string
	SetHttpsModel(v bool) *CreateGatewayApiGroupShrinkRequest
	GetHttpsModel() *bool
	SetInstanceId(v string) *CreateGatewayApiGroupShrinkRequest
	GetInstanceId() *string
	SetInstancePort(v int32) *CreateGatewayApiGroupShrinkRequest
	GetInstancePort() *int32
	SetRegionMode(v bool) *CreateGatewayApiGroupShrinkRequest
	GetRegionMode() *bool
	SetTenantId(v string) *CreateGatewayApiGroupShrinkRequest
	GetTenantId() *string
	SetTimeout(v int32) *CreateGatewayApiGroupShrinkRequest
	GetTimeout() *int32
	SetUpstreamType(v string) *CreateGatewayApiGroupShrinkRequest
	GetUpstreamType() *string
	SetVpcId(v string) *CreateGatewayApiGroupShrinkRequest
	GetVpcId() *string
	SetVpcModel(v bool) *CreateGatewayApiGroupShrinkRequest
	GetVpcModel() *bool
	SetWorkspaceId(v string) *CreateGatewayApiGroupShrinkRequest
	GetWorkspaceId() *string
}

type CreateGatewayApiGroupShrinkRequest struct {
	// This parameter is required.
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	CentersShrink *string `json:"Centers,omitempty" xml:"Centers,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DomainName    *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	EnableMTLS    *bool   `json:"EnableMTLS,omitempty" xml:"EnableMTLS,omitempty"`
	// This parameter is required.
	GroupName    *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	Host         *string `json:"Host,omitempty" xml:"Host,omitempty"`
	HttpsModel   *bool   `json:"HttpsModel,omitempty" xml:"HttpsModel,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstancePort *int32  `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	RegionMode   *bool   `json:"RegionMode,omitempty" xml:"RegionMode,omitempty"`
	// This parameter is required.
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// This parameter is required.
	UpstreamType *string `json:"UpstreamType,omitempty" xml:"UpstreamType,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcModel     *bool   `json:"VpcModel,omitempty" xml:"VpcModel,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateGatewayApiGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateGatewayApiGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateGatewayApiGroupShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateGatewayApiGroupShrinkRequest) GetCentersShrink() *string {
	return s.CentersShrink
}

func (s *CreateGatewayApiGroupShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateGatewayApiGroupShrinkRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateGatewayApiGroupShrinkRequest) GetEnableMTLS() *bool {
	return s.EnableMTLS
}

func (s *CreateGatewayApiGroupShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateGatewayApiGroupShrinkRequest) GetHost() *string {
	return s.Host
}

func (s *CreateGatewayApiGroupShrinkRequest) GetHttpsModel() *bool {
	return s.HttpsModel
}

func (s *CreateGatewayApiGroupShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateGatewayApiGroupShrinkRequest) GetInstancePort() *int32 {
	return s.InstancePort
}

func (s *CreateGatewayApiGroupShrinkRequest) GetRegionMode() *bool {
	return s.RegionMode
}

func (s *CreateGatewayApiGroupShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateGatewayApiGroupShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateGatewayApiGroupShrinkRequest) GetUpstreamType() *string {
	return s.UpstreamType
}

func (s *CreateGatewayApiGroupShrinkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateGatewayApiGroupShrinkRequest) GetVpcModel() *bool {
	return s.VpcModel
}

func (s *CreateGatewayApiGroupShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateGatewayApiGroupShrinkRequest) SetAppId(v string) *CreateGatewayApiGroupShrinkRequest {
	s.AppId = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetCentersShrink(v string) *CreateGatewayApiGroupShrinkRequest {
	s.CentersShrink = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetDescription(v string) *CreateGatewayApiGroupShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetDomainName(v string) *CreateGatewayApiGroupShrinkRequest {
	s.DomainName = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetEnableMTLS(v bool) *CreateGatewayApiGroupShrinkRequest {
	s.EnableMTLS = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetGroupName(v string) *CreateGatewayApiGroupShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetHost(v string) *CreateGatewayApiGroupShrinkRequest {
	s.Host = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetHttpsModel(v bool) *CreateGatewayApiGroupShrinkRequest {
	s.HttpsModel = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetInstanceId(v string) *CreateGatewayApiGroupShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetInstancePort(v int32) *CreateGatewayApiGroupShrinkRequest {
	s.InstancePort = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetRegionMode(v bool) *CreateGatewayApiGroupShrinkRequest {
	s.RegionMode = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetTenantId(v string) *CreateGatewayApiGroupShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetTimeout(v int32) *CreateGatewayApiGroupShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetUpstreamType(v string) *CreateGatewayApiGroupShrinkRequest {
	s.UpstreamType = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetVpcId(v string) *CreateGatewayApiGroupShrinkRequest {
	s.VpcId = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetVpcModel(v bool) *CreateGatewayApiGroupShrinkRequest {
	s.VpcModel = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) SetWorkspaceId(v string) *CreateGatewayApiGroupShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateGatewayApiGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}
