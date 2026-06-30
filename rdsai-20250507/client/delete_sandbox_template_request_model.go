// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSandboxTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceName(v string) *DeleteSandboxTemplateRequest
	GetInstanceName() *string
	SetRegionId(v string) *DeleteSandboxTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *DeleteSandboxTemplateRequest
	GetTemplateId() *string
}

type DeleteSandboxTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// desktop-xxxx
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteSandboxTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSandboxTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSandboxTemplateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteSandboxTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSandboxTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteSandboxTemplateRequest) SetInstanceName(v string) *DeleteSandboxTemplateRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteSandboxTemplateRequest) SetRegionId(v string) *DeleteSandboxTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSandboxTemplateRequest) SetTemplateId(v string) *DeleteSandboxTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteSandboxTemplateRequest) Validate() error {
	return dara.Validate(s)
}
