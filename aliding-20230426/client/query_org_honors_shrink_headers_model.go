// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryOrgHonorsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryOrgHonorsShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryOrgHonorsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryOrgHonorsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryOrgHonorsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryOrgHonorsShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgHonorsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgHonorsShrinkHeaders) SetAccountContextShrink(v string) *QueryOrgHonorsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryOrgHonorsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
