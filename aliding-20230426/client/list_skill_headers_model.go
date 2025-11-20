// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListSkillHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListSkillHeadersAccountContext) *ListSkillHeaders
	GetAccountContext() *ListSkillHeadersAccountContext
}

type ListSkillHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListSkillHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListSkillHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListSkillHeaders) GoString() string {
	return s.String()
}

func (s *ListSkillHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListSkillHeaders) GetAccountContext() *ListSkillHeadersAccountContext {
	return s.AccountContext
}

func (s *ListSkillHeaders) SetCommonHeaders(v map[string]*string) *ListSkillHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListSkillHeaders) SetAccountContext(v *ListSkillHeadersAccountContext) *ListSkillHeaders {
	s.AccountContext = v
	return s
}

func (s *ListSkillHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListSkillHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListSkillHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListSkillHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListSkillHeadersAccountContext) SetAccountId(v string) *ListSkillHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListSkillHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
