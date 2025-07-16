// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryScheduleConferenceHeadersAccountContext) *QueryScheduleConferenceHeaders
	GetAccountContext() *QueryScheduleConferenceHeadersAccountContext
}

type QueryScheduleConferenceHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryScheduleConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryScheduleConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryScheduleConferenceHeaders) GetAccountContext() *QueryScheduleConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *QueryScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryScheduleConferenceHeaders) SetAccountContext(v *QueryScheduleConferenceHeadersAccountContext) *QueryScheduleConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryScheduleConferenceHeaders) Validate() error {
	return dara.Validate(s)
}

type QueryScheduleConferenceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryScheduleConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryScheduleConferenceHeadersAccountContext) SetAccountId(v string) *QueryScheduleConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryScheduleConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
