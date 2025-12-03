// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrganizationId(v string) *GetUserInfoRequest
	GetOrganizationId() *string
}

type GetUserInfoRequest struct {
	// example:
	//
	// 61e54b0e0bb300d827e1ae27
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
}

func (s GetUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoRequest) GoString() string {
	return s.String()
}

func (s *GetUserInfoRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetUserInfoRequest) SetOrganizationId(v string) *GetUserInfoRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
