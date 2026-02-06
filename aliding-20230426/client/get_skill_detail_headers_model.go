// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSkillDetailHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSkillDetailHeadersAccountContext) *GetSkillDetailHeaders
	GetAccountContext() *GetSkillDetailHeadersAccountContext
}

type GetSkillDetailHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSkillDetailHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSkillDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailHeaders) GoString() string {
	return s.String()
}

func (s *GetSkillDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSkillDetailHeaders) GetAccountContext() *GetSkillDetailHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSkillDetailHeaders) SetCommonHeaders(v map[string]*string) *GetSkillDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSkillDetailHeaders) SetAccountContext(v *GetSkillDetailHeadersAccountContext) *GetSkillDetailHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSkillDetailHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillDetailHeadersAccountContext struct {
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

func (s GetSkillDetailHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSkillDetailHeadersAccountContext) GetSsoTicket() *string {
	return s.SsoTicket
}

func (s *GetSkillDetailHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSkillDetailHeadersAccountContext) SetSsoTicket(v string) *GetSkillDetailHeadersAccountContext {
	s.SsoTicket = &v
	return s
}

func (s *GetSkillDetailHeadersAccountContext) SetAccountId(v string) *GetSkillDetailHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSkillDetailHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
