// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAppLlmApiKeyForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateAppLlmApiKeyForPartnerRequest
	GetBizId() *string
	SetClientToken(v string) *CreateAppLlmApiKeyForPartnerRequest
	GetClientToken() *string
	SetDescription(v string) *CreateAppLlmApiKeyForPartnerRequest
	GetDescription() *string
	SetIpWhiteList(v []*string) *CreateAppLlmApiKeyForPartnerRequest
	GetIpWhiteList() []*string
}

type CreateAppLlmApiKeyForPartnerRequest struct {
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// xxxxx-xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// success
	Description *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	IpWhiteList []*string `json:"IpWhiteList,omitempty" xml:"IpWhiteList,omitempty" type:"Repeated"`
}

func (s CreateAppLlmApiKeyForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAppLlmApiKeyForPartnerRequest) GoString() string {
	return s.String()
}

func (s *CreateAppLlmApiKeyForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateAppLlmApiKeyForPartnerRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateAppLlmApiKeyForPartnerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateAppLlmApiKeyForPartnerRequest) GetIpWhiteList() []*string {
	return s.IpWhiteList
}

func (s *CreateAppLlmApiKeyForPartnerRequest) SetBizId(v string) *CreateAppLlmApiKeyForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerRequest) SetClientToken(v string) *CreateAppLlmApiKeyForPartnerRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerRequest) SetDescription(v string) *CreateAppLlmApiKeyForPartnerRequest {
	s.Description = &v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerRequest) SetIpWhiteList(v []*string) *CreateAppLlmApiKeyForPartnerRequest {
	s.IpWhiteList = v
	return s
}

func (s *CreateAppLlmApiKeyForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
