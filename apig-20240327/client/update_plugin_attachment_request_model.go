// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePluginAttachmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttachResourceIds(v []*string) *UpdatePluginAttachmentRequest
	GetAttachResourceIds() []*string
	SetEnable(v bool) *UpdatePluginAttachmentRequest
	GetEnable() *bool
	SetPluginConfig(v string) *UpdatePluginAttachmentRequest
	GetPluginConfig() *string
}

type UpdatePluginAttachmentRequest struct {
	AttachResourceIds []*string `json:"attachResourceIds,omitempty" xml:"attachResourceIds,omitempty" type:"Repeated"`
	// example:
	//
	// false
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// example:
	//
	// cHJlcGVuZDoKLSByb2xlOiBzeXN0ZW0KICBjb250ZW50OiDor7fkvb/nlKjoi7Hor63lm57nrZTpl67popgKYXBwZW5kOgotIHJvbGU6IHVzZXIKICBjb250ZW50OiDmr4/mrKHlm57nrZTlrozpl67popjvvIzlsJ3or5Xov5vooYzlj43pl64K
	PluginConfig *string `json:"pluginConfig,omitempty" xml:"pluginConfig,omitempty"`
}

func (s UpdatePluginAttachmentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePluginAttachmentRequest) GoString() string {
	return s.String()
}

func (s *UpdatePluginAttachmentRequest) GetAttachResourceIds() []*string {
	return s.AttachResourceIds
}

func (s *UpdatePluginAttachmentRequest) GetEnable() *bool {
	return s.Enable
}

func (s *UpdatePluginAttachmentRequest) GetPluginConfig() *string {
	return s.PluginConfig
}

func (s *UpdatePluginAttachmentRequest) SetAttachResourceIds(v []*string) *UpdatePluginAttachmentRequest {
	s.AttachResourceIds = v
	return s
}

func (s *UpdatePluginAttachmentRequest) SetEnable(v bool) *UpdatePluginAttachmentRequest {
	s.Enable = &v
	return s
}

func (s *UpdatePluginAttachmentRequest) SetPluginConfig(v string) *UpdatePluginAttachmentRequest {
	s.PluginConfig = &v
	return s
}

func (s *UpdatePluginAttachmentRequest) Validate() error {
	return dara.Validate(s)
}
