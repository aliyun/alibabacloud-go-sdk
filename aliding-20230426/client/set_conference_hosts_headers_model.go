// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetConferenceHostsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SetConferenceHostsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SetConferenceHostsHeadersAccountContext) *SetConferenceHostsHeaders
	GetAccountContext() *SetConferenceHostsHeadersAccountContext
}

type SetConferenceHostsHeaders struct {
	CommonHeaders  map[string]*string                       `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SetConferenceHostsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SetConferenceHostsHeaders) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsHeaders) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SetConferenceHostsHeaders) GetAccountContext() *SetConferenceHostsHeadersAccountContext {
	return s.AccountContext
}

func (s *SetConferenceHostsHeaders) SetCommonHeaders(v map[string]*string) *SetConferenceHostsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SetConferenceHostsHeaders) SetAccountContext(v *SetConferenceHostsHeadersAccountContext) *SetConferenceHostsHeaders {
	s.AccountContext = v
	return s
}

func (s *SetConferenceHostsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetConferenceHostsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SetConferenceHostsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SetConferenceHostsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SetConferenceHostsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SetConferenceHostsHeadersAccountContext) SetAccountId(v string) *SetConferenceHostsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SetConferenceHostsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
