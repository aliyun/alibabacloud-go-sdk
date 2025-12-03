// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePushRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *DeletePushRuleRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *DeletePushRuleRequest
	GetOrganizationId() *string
}

type DeletePushRuleRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 60d54f3daccf2bbd6659f3ad
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s DeletePushRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePushRuleRequest) GoString() string {
	return s.String()
}

func (s *DeletePushRuleRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *DeletePushRuleRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *DeletePushRuleRequest) SetAccessToken(v string) *DeletePushRuleRequest {
	s.AccessToken = &v
	return s
}

func (s *DeletePushRuleRequest) SetOrganizationId(v string) *DeletePushRuleRequest {
	s.OrganizationId = &v
	return s
}

func (s *DeletePushRuleRequest) Validate() error {
	return dara.Validate(s)
}
