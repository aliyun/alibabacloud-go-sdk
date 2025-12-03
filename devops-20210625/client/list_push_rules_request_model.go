// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPushRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *ListPushRulesRequest
	GetAccessToken() *string
	SetOrganizationId(v string) *ListPushRulesRequest
	GetOrganizationId() *string
}

type ListPushRulesRequest struct {
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

func (s ListPushRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPushRulesRequest) GoString() string {
	return s.String()
}

func (s *ListPushRulesRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *ListPushRulesRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListPushRulesRequest) SetAccessToken(v string) *ListPushRulesRequest {
	s.AccessToken = &v
	return s
}

func (s *ListPushRulesRequest) SetOrganizationId(v string) *ListPushRulesRequest {
	s.OrganizationId = &v
	return s
}

func (s *ListPushRulesRequest) Validate() error {
	return dara.Validate(s)
}
