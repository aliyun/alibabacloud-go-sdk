// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduleConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateScheduleConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateScheduleConferenceHeadersAccountContext) *UpdateScheduleConferenceHeaders
	GetAccountContext() *UpdateScheduleConferenceHeadersAccountContext
}

type UpdateScheduleConferenceHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateScheduleConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateScheduleConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateScheduleConferenceHeaders) GetAccountContext() *UpdateScheduleConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *UpdateScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateScheduleConferenceHeaders) SetAccountContext(v *UpdateScheduleConferenceHeadersAccountContext) *UpdateScheduleConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateScheduleConferenceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateScheduleConferenceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateScheduleConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduleConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateScheduleConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateScheduleConferenceHeadersAccountContext) SetAccountId(v string) *UpdateScheduleConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateScheduleConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
