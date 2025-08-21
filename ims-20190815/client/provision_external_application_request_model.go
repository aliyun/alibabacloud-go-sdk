// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProvisionExternalApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ProvisionExternalApplicationRequest
	GetAppId() *string
	SetScopes(v string) *ProvisionExternalApplicationRequest
	GetScopes() *string
}

type ProvisionExternalApplicationRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 403550611646604****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The permissions that are granted to the application. Separate multiple permissions with a semicolon (;).
	//
	// >  For more information about the supported permissions, see [Overview](https://help.aliyun.com/document_detail/93693.html).
	//
	// example:
	//
	// openid;aliuid
	Scopes *string `json:"Scopes,omitempty" xml:"Scopes,omitempty"`
}

func (s ProvisionExternalApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ProvisionExternalApplicationRequest) GoString() string {
	return s.String()
}

func (s *ProvisionExternalApplicationRequest) GetAppId() *string {
	return s.AppId
}

func (s *ProvisionExternalApplicationRequest) GetScopes() *string {
	return s.Scopes
}

func (s *ProvisionExternalApplicationRequest) SetAppId(v string) *ProvisionExternalApplicationRequest {
	s.AppId = &v
	return s
}

func (s *ProvisionExternalApplicationRequest) SetScopes(v string) *ProvisionExternalApplicationRequest {
	s.Scopes = &v
	return s
}

func (s *ProvisionExternalApplicationRequest) Validate() error {
	return dara.Validate(s)
}
