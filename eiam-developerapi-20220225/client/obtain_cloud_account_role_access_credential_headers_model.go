// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iObtainCloudAccountRoleAccessCredentialHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ObtainCloudAccountRoleAccessCredentialHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *ObtainCloudAccountRoleAccessCredentialHeaders
	GetAuthorization() *string
}

type ObtainCloudAccountRoleAccessCredentialHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s ObtainCloudAccountRoleAccessCredentialHeaders) String() string {
	return dara.Prettify(s)
}

func (s ObtainCloudAccountRoleAccessCredentialHeaders) GoString() string {
	return s.String()
}

func (s *ObtainCloudAccountRoleAccessCredentialHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ObtainCloudAccountRoleAccessCredentialHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *ObtainCloudAccountRoleAccessCredentialHeaders) SetCommonHeaders(v map[string]*string) *ObtainCloudAccountRoleAccessCredentialHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialHeaders) SetAuthorization(v string) *ObtainCloudAccountRoleAccessCredentialHeaders {
	s.Authorization = &v
	return s
}

func (s *ObtainCloudAccountRoleAccessCredentialHeaders) Validate() error {
	return dara.Validate(s)
}
