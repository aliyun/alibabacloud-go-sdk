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
	// The ID of the external application that was created by another Alibaba Cloud account and can be installed by the current account. The application IDs returned by `ListApplications` for the current account are not applicable. `ListExternalApplications` and `ListApplicationProvisionInfos` only query installed records and cannot discover external application IDs that have not been installed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 403550611646604****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The permission scopes granted to the application. You can specify multiple scopes separated by semicolons (;).
	//
	//
	// > For supported permission scopes, refer to "OAuth Scopes" in [OAuth application overview](https://help.aliyun.com/document_detail/93693.html).
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
