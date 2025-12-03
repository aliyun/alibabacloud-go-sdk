// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchUserHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchUserHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *PatchUserHeaders
	GetAuthorization() *string
}

type PatchUserHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. The value is in the Bearer ${access_token} format. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer AT8csE2seYxxxxxij
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchUserHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchUserHeaders) GoString() string {
	return s.String()
}

func (s *PatchUserHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchUserHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PatchUserHeaders) SetCommonHeaders(v map[string]*string) *PatchUserHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchUserHeaders) SetAuthorization(v string) *PatchUserHeaders {
	s.Authorization = &v
	return s
}

func (s *PatchUserHeaders) Validate() error {
	return dara.Validate(s)
}
