// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSkillsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSkillsHeadersAccountContext) *GetSkillsHeaders
	GetAccountContext() *GetSkillsHeadersAccountContext
}

type GetSkillsHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSkillsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSkillsHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsHeaders) GoString() string {
	return s.String()
}

func (s *GetSkillsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSkillsHeaders) GetAccountContext() *GetSkillsHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSkillsHeaders) SetCommonHeaders(v map[string]*string) *GetSkillsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSkillsHeaders) SetAccountContext(v *GetSkillsHeadersAccountContext) *GetSkillsHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSkillsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillsHeadersAccountContext struct {
	// Buc SsoTicket
	//
	// example:
	//
	// bucxxx
	SsoTicket *string `json:"SsoTicket,omitempty" xml:"SsoTicket,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetSkillsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSkillsHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *GetSkillsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSkillsHeadersAccountContext) SetSsoTicket(v string) *GetSkillsHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *GetSkillsHeadersAccountContext) SetAccountId(v string) *GetSkillsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSkillsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
