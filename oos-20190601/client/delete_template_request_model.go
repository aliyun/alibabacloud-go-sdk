// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeleteExecutions(v bool) *DeleteTemplateRequest
	GetAutoDeleteExecutions() *bool
	SetRegionId(v string) *DeleteTemplateRequest
	GetRegionId() *string
	SetTemplateName(v string) *DeleteTemplateRequest
	GetTemplateName() *string
}

type DeleteTemplateRequest struct {
	// Specifies whether to delete the related executions when a template is deleted.
	//
	// example:
	//
	// false
	AutoDeleteExecutions *bool `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template. The name can be 1 to 200 characters in length and can contain letters, digits, hyphens (-), and underscores (_). It cannot start with ALIYUN, ACS, ALIBABA, or ALICLOUD.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeleteTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplateRequest) GetAutoDeleteExecutions() *bool {
	return s.AutoDeleteExecutions
}

func (s *DeleteTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DeleteTemplateRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplateRequest {
	s.AutoDeleteExecutions = &v
	return s
}

func (s *DeleteTemplateRequest) SetRegionId(v string) *DeleteTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplateRequest) SetTemplateName(v string) *DeleteTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteTemplateRequest) Validate() error {
	return dara.Validate(s)
}
