// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWhitelistTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateWhitelistTemplateRequest
	GetRegionId() *string
	SetSecurityIPList(v string) *CreateWhitelistTemplateRequest
	GetSecurityIPList() *string
	SetTemplateName(v string) *CreateWhitelistTemplateRequest
	GetTemplateName() *string
}

type CreateWhitelistTemplateRequest struct {
	// RegionId
	//
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
	// InvoiceTemplate
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s CreateWhitelistTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWhitelistTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateWhitelistTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateWhitelistTemplateRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *CreateWhitelistTemplateRequest) GetTemplateName() *string {
	return s.TemplateName
}

func (s *CreateWhitelistTemplateRequest) SetRegionId(v string) *CreateWhitelistTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateWhitelistTemplateRequest) SetSecurityIPList(v string) *CreateWhitelistTemplateRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateWhitelistTemplateRequest) SetTemplateName(v string) *CreateWhitelistTemplateRequest {
	s.TemplateName = &v
	return s
}

func (s *CreateWhitelistTemplateRequest) Validate() error {
	return dara.Validate(s)
}
