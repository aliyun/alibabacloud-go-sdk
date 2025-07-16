// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantHonorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GrantHonorHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GrantHonorHeadersAccountContext) *GrantHonorHeaders
	GetAccountContext() *GrantHonorHeadersAccountContext
}

type GrantHonorHeaders struct {
	CommonHeaders  map[string]*string               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GrantHonorHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GrantHonorHeaders) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorHeaders) GoString() string {
	return s.String()
}

func (s *GrantHonorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GrantHonorHeaders) GetAccountContext() *GrantHonorHeadersAccountContext {
	return s.AccountContext
}

func (s *GrantHonorHeaders) SetCommonHeaders(v map[string]*string) *GrantHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GrantHonorHeaders) SetAccountContext(v *GrantHonorHeadersAccountContext) *GrantHonorHeaders {
	s.AccountContext = v
	return s
}

func (s *GrantHonorHeaders) Validate() error {
	return dara.Validate(s)
}

type GrantHonorHeadersAccountContext struct {
	// example:
	//
	// null
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GrantHonorHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GrantHonorHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GrantHonorHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GrantHonorHeadersAccountContext) SetAccountId(v string) *GrantHonorHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GrantHonorHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
