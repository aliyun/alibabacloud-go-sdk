// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecordMinutesUrlHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryRecordMinutesUrlHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryRecordMinutesUrlHeadersAccountContext) *QueryRecordMinutesUrlHeaders
	GetAccountContext() *QueryRecordMinutesUrlHeadersAccountContext
}

type QueryRecordMinutesUrlHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryRecordMinutesUrlHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryRecordMinutesUrlHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlHeaders) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryRecordMinutesUrlHeaders) GetAccountContext() *QueryRecordMinutesUrlHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryRecordMinutesUrlHeaders) SetCommonHeaders(v map[string]*string) *QueryRecordMinutesUrlHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryRecordMinutesUrlHeaders) SetAccountContext(v *QueryRecordMinutesUrlHeadersAccountContext) *QueryRecordMinutesUrlHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryRecordMinutesUrlHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryRecordMinutesUrlHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryRecordMinutesUrlHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryRecordMinutesUrlHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryRecordMinutesUrlHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryRecordMinutesUrlHeadersAccountContext) SetAccountId(v string) *QueryRecordMinutesUrlHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryRecordMinutesUrlHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
