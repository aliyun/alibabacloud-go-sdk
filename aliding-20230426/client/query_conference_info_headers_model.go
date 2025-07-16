// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryConferenceInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryConferenceInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryConferenceInfoHeadersAccountContext) *QueryConferenceInfoHeaders
	GetAccountContext() *QueryConferenceInfoHeadersAccountContext
}

type QueryConferenceInfoHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryConferenceInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryConferenceInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryConferenceInfoHeaders) GetAccountContext() *QueryConferenceInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryConferenceInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryConferenceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryConferenceInfoHeaders) SetAccountContext(v *QueryConferenceInfoHeadersAccountContext) *QueryConferenceInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryConferenceInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryConferenceInfoHeadersAccountContext struct {
	// example:
	//
	// 208579
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryConferenceInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryConferenceInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryConferenceInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryConferenceInfoHeadersAccountContext) SetAccountId(v string) *QueryConferenceInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryConferenceInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
