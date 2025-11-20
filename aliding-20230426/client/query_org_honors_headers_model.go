// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryOrgHonorsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryOrgHonorsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryOrgHonorsHeadersAccountContext) *QueryOrgHonorsHeaders
	GetAccountContext() *QueryOrgHonorsHeadersAccountContext
}

type QueryOrgHonorsHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryOrgHonorsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryOrgHonorsHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryOrgHonorsHeaders) GetAccountContext() *QueryOrgHonorsHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryOrgHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryOrgHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryOrgHonorsHeaders) SetAccountContext(v *QueryOrgHonorsHeadersAccountContext) *QueryOrgHonorsHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryOrgHonorsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryOrgHonorsHeadersAccountContext struct {
	// example:
	//
	// 243331014234180628
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryOrgHonorsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryOrgHonorsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryOrgHonorsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryOrgHonorsHeadersAccountContext) SetAccountId(v string) *QueryOrgHonorsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryOrgHonorsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
