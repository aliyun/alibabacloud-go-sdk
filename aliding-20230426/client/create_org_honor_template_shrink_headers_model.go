// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateOrgHonorTemplateShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateOrgHonorTemplateShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateOrgHonorTemplateShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateOrgHonorTemplateShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateOrgHonorTemplateShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateOrgHonorTemplateShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateOrgHonorTemplateShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrgHonorTemplateShrinkHeaders) SetAccountContextShrink(v string) *CreateOrgHonorTemplateShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateOrgHonorTemplateShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
