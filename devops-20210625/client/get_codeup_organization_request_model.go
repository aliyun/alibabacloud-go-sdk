// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeupOrganizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessToken(v string) *GetCodeupOrganizationRequest
	GetAccessToken() *string
}

type GetCodeupOrganizationRequest struct {
	// example:
	//
	// f0b1e61db5961df5975a93f9129d2513
	AccessToken *string `json:"accessToken,omitempty" xml:"accessToken,omitempty"`
}

func (s GetCodeupOrganizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCodeupOrganizationRequest) GoString() string {
	return s.String()
}

func (s *GetCodeupOrganizationRequest) GetAccessToken() *string {
	return s.AccessToken
}

func (s *GetCodeupOrganizationRequest) SetAccessToken(v string) *GetCodeupOrganizationRequest {
	s.AccessToken = &v
	return s
}

func (s *GetCodeupOrganizationRequest) Validate() error {
	return dara.Validate(s)
}
