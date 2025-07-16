// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsViewHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListEventsViewHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListEventsViewHeadersAccountContext) *ListEventsViewHeaders
	GetAccountContext() *ListEventsViewHeadersAccountContext
}

type ListEventsViewHeaders struct {
	CommonHeaders  map[string]*string                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListEventsViewHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListEventsViewHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewHeaders) GoString() string {
	return s.String()
}

func (s *ListEventsViewHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListEventsViewHeaders) GetAccountContext() *ListEventsViewHeadersAccountContext {
	return s.AccountContext
}

func (s *ListEventsViewHeaders) SetCommonHeaders(v map[string]*string) *ListEventsViewHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListEventsViewHeaders) SetAccountContext(v *ListEventsViewHeadersAccountContext) *ListEventsViewHeaders {
	s.AccountContext = v
	return s
}

func (s *ListEventsViewHeaders) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListEventsViewHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListEventsViewHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListEventsViewHeadersAccountContext) SetAccountId(v string) *ListEventsViewHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListEventsViewHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
