// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationProvisioningScopeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetApplicationProvisioningScopeHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *GetApplicationProvisioningScopeHeaders
	GetAuthorization() *string
}

type GetApplicationProvisioningScopeHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. Format: Bearer ${access_token}. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s GetApplicationProvisioningScopeHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationProvisioningScopeHeaders) GoString() string {
	return s.String()
}

func (s *GetApplicationProvisioningScopeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetApplicationProvisioningScopeHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *GetApplicationProvisioningScopeHeaders) SetCommonHeaders(v map[string]*string) *GetApplicationProvisioningScopeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetApplicationProvisioningScopeHeaders) SetAuthorization(v string) *GetApplicationProvisioningScopeHeaders {
	s.Authorization = &v
	return s
}

func (s *GetApplicationProvisioningScopeHeaders) Validate() error {
	return dara.Validate(s)
}
