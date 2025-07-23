// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegion(v string) *GetTemplateRequest
	GetRegion() *string
	SetResourceGroupId(v string) *GetTemplateRequest
	GetResourceGroupId() *string
	SetTemplateId(v string) *GetTemplateRequest
	GetTemplateId() *string
}

type GetTemplateRequest struct {
	// Template Area
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// ResourceGroup ID
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Template ID
	//
	// This parameter is required.
	//
	// example:
	//
	// XFKR6WYRVS24S07R
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetRegion() *string {
	return s.Region
}

func (s *GetTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateRequest) SetRegion(v string) *GetTemplateRequest {
	s.Region = &v
	return s
}

func (s *GetTemplateRequest) SetResourceGroupId(v string) *GetTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateId(v string) *GetTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
