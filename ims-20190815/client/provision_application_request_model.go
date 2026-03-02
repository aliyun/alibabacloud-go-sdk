// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ProvisionApplicationRequest
	GetAppId() *string
	SetScopes(v string) *ProvisionApplicationRequest
	GetScopes() *string
}

type ProvisionApplicationRequest struct {
	// This parameter is required.
	AppId  *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Scopes *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
}

func (s ProvisionApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ProvisionApplicationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ProvisionApplicationRequest) GetScopes() *string {
	return s.Scopes
}

func (s *ProvisionApplicationRequest) SetAppId(v string) *ProvisionApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ProvisionApplicationRequest) SetScopes(v string) *ProvisionApplicationRequest {
	s.Scopes = &v
	return s
}

func (s *ProvisionApplicationRequest) Validate() error {
	return dara.Validate(s)
}
