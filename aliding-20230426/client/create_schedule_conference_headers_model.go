// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateScheduleConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateScheduleConferenceHeadersAccountContext) *CreateScheduleConferenceHeaders
	GetAccountContext() *CreateScheduleConferenceHeadersAccountContext
}

type CreateScheduleConferenceHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateScheduleConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateScheduleConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateScheduleConferenceHeaders) GetAccountContext() *CreateScheduleConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *CreateScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateScheduleConferenceHeaders) SetAccountContext(v *CreateScheduleConferenceHeadersAccountContext) *CreateScheduleConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateScheduleConferenceHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateScheduleConferenceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateScheduleConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateScheduleConferenceHeadersAccountContext) SetAccountId(v string) *CreateScheduleConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateScheduleConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
