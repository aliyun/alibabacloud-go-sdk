// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetWhitelistTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *GetWhitelistTemplateRequest
	GetTemplateId() *string
}

type GetWhitelistTemplateRequest struct {
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
}

func (s GetWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetWhitelistTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetWhitelistTemplateRequest) SetRegionId(v string) *GetWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetWhitelistTemplateRequest) SetTemplateId(v string) *GetWhitelistTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
