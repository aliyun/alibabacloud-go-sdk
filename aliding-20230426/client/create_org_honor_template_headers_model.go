// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrgHonorTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateOrgHonorTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateOrgHonorTemplateHeadersAccountContext) *CreateOrgHonorTemplateHeaders
	GetAccountContext() *CreateOrgHonorTemplateHeadersAccountContext
}

type CreateOrgHonorTemplateHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateOrgHonorTemplateHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateOrgHonorTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateOrgHonorTemplateHeaders) GetAccountContext() *CreateOrgHonorTemplateHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateOrgHonorTemplateHeaders) SetCommonHeaders(v map[string]*string) *CreateOrgHonorTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrgHonorTemplateHeaders) SetAccountContext(v *CreateOrgHonorTemplateHeadersAccountContext) *CreateOrgHonorTemplateHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateOrgHonorTemplateHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrgHonorTemplateHeadersAccountContext struct {
	// example:
	//
	// 208579
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateOrgHonorTemplateHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateOrgHonorTemplateHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateOrgHonorTemplateHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateOrgHonorTemplateHeadersAccountContext) SetAccountId(v string) *CreateOrgHonorTemplateHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateOrgHonorTemplateHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
