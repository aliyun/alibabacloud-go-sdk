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
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 407426893752729****
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
