// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetColumnsVisibilityHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SetColumnsVisibilityHeadersAccountContext) *SetColumnsVisibilityHeaders
	GetAccountContext() *SetColumnsVisibilityHeadersAccountContext
}

type SetColumnsVisibilityHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SetColumnsVisibilityHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SetColumnsVisibilityHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityHeaders) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetColumnsVisibilityHeaders) GetAccountContext() *SetColumnsVisibilityHeadersAccountContext {
	return s.AccountContext
}

func (s *SetColumnsVisibilityHeaders) SetCommonHeaders(v map[string]*string) *SetColumnsVisibilityHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetColumnsVisibilityHeaders) SetAccountContext(v *SetColumnsVisibilityHeadersAccountContext) *SetColumnsVisibilityHeaders {
	s.AccountContext = v
	return s
}

func (s *SetColumnsVisibilityHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetColumnsVisibilityHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SetColumnsVisibilityHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SetColumnsVisibilityHeadersAccountContext) SetAccountId(v string) *SetColumnsVisibilityHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SetColumnsVisibilityHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
