// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePluginAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceIds(v []*string) *CreatePluginAttachmentRequest
	GetAttachResourceIds() []*string
	SetAttachResourceType(v string) *CreatePluginAttachmentRequest
	GetAttachResourceType() *string
	SetEnable(v bool) *CreatePluginAttachmentRequest
	GetEnable() *bool
	SetEnvironmentId(v string) *CreatePluginAttachmentRequest
	GetEnvironmentId() *string
	SetGatewayId(v string) *CreatePluginAttachmentRequest
	GetGatewayId() *string
	SetPluginConfig(v string) *CreatePluginAttachmentRequest
	GetPluginConfig() *string
	SetPluginId(v string) *CreatePluginAttachmentRequest
	GetPluginId() *string
}

type CreatePluginAttachmentRequest struct {
	// The attachment IDs.
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// The type of the resource to which the plug-in is attached. Valid values: GatewayRoute, Gateway, GatewayDomain, HttpApi, and Operation.
	//
	// example:
	//
	// HttpApi
	AttachResourceType *string `json:"attachResourceType,omitempty" xml:"attachResourceType,omitempty"`
	// Specifies whether to enable the plug-in. Default value: false.
	//
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The environment ID.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// gw-cq7l5s5lhtg***
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The Base64-encoded configurations of the plug-in.
	//
	// example:
	//
	// cHJlcGVuZDoKLSByb2xlOiBzeXN0ZW0KICBjb250ZW50OiDor7fkvb/nlKjoi7Hor63lm57nrZTpl67popgKYXBwZW5kOgotIHJvbGU6IHVzZXIKICBjb250ZW50OiDmr4/mrKHlm57nrZTlrozpl67popjvvIzlsJ3or5Xov5vooYzlj43pl64K
	PluginConfig *string `json:"pluginConfig,omitempty" xml:"pluginConfig,omitempty"`
	// The plug-in ID.
	//
	// example:
	//
	// pl-cvo8udem1hkob1qd67i0
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
}

func (s CreatePluginAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePluginAttachmentRequest) GoString() string {
	return s.String()
}

func (s *CreatePluginAttachmentRequest) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
}

func (s *CreatePluginAttachmentRequest) GetAttachResourceType() *string {
	return s.AttachResourceType
}

func (s *CreatePluginAttachmentRequest) GetEnable() *bool {
	return s.Enable
}

func (s *CreatePluginAttachmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreatePluginAttachmentRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreatePluginAttachmentRequest) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *CreatePluginAttachmentRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *CreatePluginAttachmentRequest) SetAttachResourceIds(v []*string) *CreatePluginAttachmentRequest {
	s.AttachResourceIds = v
	return s
}

func (s *CreatePluginAttachmentRequest) SetAttachResourceType(v string) *CreatePluginAttachmentRequest {
	s.AttachResourceType = &v
	return s
}

func (s *CreatePluginAttachmentRequest) SetEnable(v bool) *CreatePluginAttachmentRequest {
	s.Enable = &v
	return s
}

func (s *CreatePluginAttachmentRequest) SetEnvironmentId(v string) *CreatePluginAttachmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *CreatePluginAttachmentRequest) SetGatewayId(v string) *CreatePluginAttachmentRequest {
	s.GatewayId = &v
	return s
}

func (s *CreatePluginAttachmentRequest) SetPluginConfig(v string) *CreatePluginAttachmentRequest {
	s.PluginConfig = &v
	return s
}

func (s *CreatePluginAttachmentRequest) SetPluginId(v string) *CreatePluginAttachmentRequest {
	s.PluginId = &v
	return s
}

func (s *CreatePluginAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
