// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryScheduleConferenceInfoHeadersAccountContext) *QueryScheduleConferenceInfoHeaders
	GetAccountContext() *QueryScheduleConferenceInfoHeadersAccountContext
}

type QueryScheduleConferenceInfoHeaders struct {
	CommonHeaders  map[string]*string                                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryScheduleConferenceInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryScheduleConferenceInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryScheduleConferenceInfoHeaders) GetAccountContext() *QueryScheduleConferenceInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryScheduleConferenceInfoHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceInfoHeaders) SetAccountContext(v *QueryScheduleConferenceInfoHeadersAccountContext) *QueryScheduleConferenceInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryScheduleConferenceInfoHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryScheduleConferenceInfoHeadersAccountContext struct {
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryScheduleConferenceInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryScheduleConferenceInfoHeadersAccountContext) SetAccountId(v string) *QueryScheduleConferenceInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryScheduleConferenceInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
