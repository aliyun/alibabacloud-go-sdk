// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SwitchMainOrgHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SwitchMainOrgHeadersAccountContext) *SwitchMainOrgHeaders
	GetAccountContext() *SwitchMainOrgHeadersAccountContext
}

type SwitchMainOrgHeaders struct {
	CommonHeaders  map[string]*string                  `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SwitchMainOrgHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SwitchMainOrgHeaders) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgHeaders) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SwitchMainOrgHeaders) GetAccountContext() *SwitchMainOrgHeadersAccountContext {
	return s.AccountContext
}

func (s *SwitchMainOrgHeaders) SetCommonHeaders(v map[string]*string) *SwitchMainOrgHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SwitchMainOrgHeaders) SetAccountContext(v *SwitchMainOrgHeadersAccountContext) *SwitchMainOrgHeaders {
	s.AccountContext = v
	return s
}

func (s *SwitchMainOrgHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SwitchMainOrgHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SwitchMainOrgHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SwitchMainOrgHeadersAccountContext) SetAccountId(v string) *SwitchMainOrgHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SwitchMainOrgHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
