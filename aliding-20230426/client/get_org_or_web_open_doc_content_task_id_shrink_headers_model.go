// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrgOrWebOpenDocContentTaskIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetOrgOrWebOpenDocContentTaskIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) SetAccountContextShrink(v string) *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetOrgOrWebOpenDocContentTaskIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
