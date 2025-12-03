// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *PatchGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAuthorization(v string) *PatchGroupHeaders
	GetAuthorization() *string
}

type PatchGroupHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// The authentication information. The value is in the Bearer ${access_token} format. Example: Bearer ATxxxx.
	//
	// This parameter is required.
	//
	// example:
	//
	// Bearer xxxx
	Authorization *string `json:"Authorization,omitempty" xml:"Authorization,omitempty"`
}

func (s PatchGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s PatchGroupHeaders) GoString() string {
	return s.String()
}

func (s *PatchGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *PatchGroupHeaders) GetAuthorization() *string {
	return s.Authorization
}

func (s *PatchGroupHeaders) SetCommonHeaders(v map[string]*string) *PatchGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PatchGroupHeaders) SetAuthorization(v string) *PatchGroupHeaders {
	s.Authorization = &v
	return s
}

func (s *PatchGroupHeaders) Validate() error {
	return dara.Validate(s)
}
