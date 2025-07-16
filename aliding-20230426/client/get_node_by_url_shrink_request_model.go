// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeByUrlShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOptionShrink(v string) *GetNodeByUrlShrinkRequest
	GetOptionShrink() *string
	SetTenantContextShrink(v string) *GetNodeByUrlShrinkRequest
	GetTenantContextShrink() *string
	SetUrl(v string) *GetNodeByUrlShrinkRequest
	GetUrl() *string
}

type GetNodeByUrlShrinkRequest struct {
	OptionShrink        *string `json:"Option,omitempty" xml:"Option,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://alidocs.dingtalk.com/i/nodes/EpGBa2L*********gN7R35y
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetNodeByUrlShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNodeByUrlShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetNodeByUrlShrinkRequest) GetOptionShrink() *string {
	return s.OptionShrink
}

func (s *GetNodeByUrlShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetNodeByUrlShrinkRequest) GetUrl() *string {
	return s.Url
}

func (s *GetNodeByUrlShrinkRequest) SetOptionShrink(v string) *GetNodeByUrlShrinkRequest {
	s.OptionShrink = &v
	return s
}

func (s *GetNodeByUrlShrinkRequest) SetTenantContextShrink(v string) *GetNodeByUrlShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetNodeByUrlShrinkRequest) SetUrl(v string) *GetNodeByUrlShrinkRequest {
	s.Url = &v
	return s
}

func (s *GetNodeByUrlShrinkRequest) Validate() error {
	return dara.Validate(s)
}
