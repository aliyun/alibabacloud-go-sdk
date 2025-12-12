// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteWhitelistTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *DeleteWhitelistTemplateRequest
	GetTemplateId() *string
	SetTemplateName(v string) *DeleteWhitelistTemplateRequest
	GetTemplateName() *string
}

type DeleteWhitelistTemplateRequest struct {
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
	// 98a6d3db05984dca
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// 98a6d3db05984dca
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s DeleteWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteWhitelistTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteWhitelistTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DeleteWhitelistTemplateRequest) SetRegionId(v string) *DeleteWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteWhitelistTemplateRequest) SetTemplateId(v string) *DeleteWhitelistTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteWhitelistTemplateRequest) SetTemplateName(v string) *DeleteWhitelistTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *DeleteWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
