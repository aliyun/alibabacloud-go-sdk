// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTeamsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListTeamsHeadersAccountContext) *ListTeamsHeaders
	GetAccountContext() *ListTeamsHeadersAccountContext
}

type ListTeamsHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListTeamsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListTeamsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsHeaders) GoString() string {
	return s.String()
}

func (s *ListTeamsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTeamsHeaders) GetAccountContext() *ListTeamsHeadersAccountContext {
	return s.AccountContext
}

func (s *ListTeamsHeaders) SetCommonHeaders(v map[string]*string) *ListTeamsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTeamsHeaders) SetAccountContext(v *ListTeamsHeadersAccountContext) *ListTeamsHeaders {
	s.AccountContext = v
	return s
}

func (s *ListTeamsHeaders) Validate() error {
	return dara.Validate(s)
}

type ListTeamsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListTeamsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListTeamsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListTeamsHeadersAccountContext) SetAccountId(v string) *ListTeamsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListTeamsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
