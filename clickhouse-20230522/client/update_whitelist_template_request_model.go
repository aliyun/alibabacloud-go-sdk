// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateWhitelistTemplateRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *UpdateWhitelistTemplateRequest
	GetSecurityIPList() *string
	SetTemplateId(v string) *UpdateWhitelistTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *UpdateWhitelistTemplateRequest
	GetTemplateName() *string
}

type UpdateWhitelistTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.1,10.0.0.1
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 98a6d3db05984dca
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 98a6d3db05984dca
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s UpdateWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateWhitelistTemplateRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *UpdateWhitelistTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWhitelistTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *UpdateWhitelistTemplateRequest) SetRegionId(v string) *UpdateWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateWhitelistTemplateRequest) SetSecurityIPList(v string) *UpdateWhitelistTemplateRequest {
	s.SecurityIPList = &v
	return s
}

func (s *UpdateWhitelistTemplateRequest) SetTemplateId(v string) *UpdateWhitelistTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWhitelistTemplateRequest) SetTemplateName(v string) *UpdateWhitelistTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *UpdateWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
