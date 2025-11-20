// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelScheduleConferenceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CancelScheduleConferenceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CancelScheduleConferenceHeadersAccountContext) *CancelScheduleConferenceHeaders
	GetAccountContext() *CancelScheduleConferenceHeadersAccountContext
}

type CancelScheduleConferenceHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CancelScheduleConferenceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CancelScheduleConferenceHeaders) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceHeaders) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CancelScheduleConferenceHeaders) GetAccountContext() *CancelScheduleConferenceHeadersAccountContext {
	return s.AccountContext
}

func (s *CancelScheduleConferenceHeaders) SetCommonHeaders(v map[string]*string) *CancelScheduleConferenceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CancelScheduleConferenceHeaders) SetAccountContext(v *CancelScheduleConferenceHeadersAccountContext) *CancelScheduleConferenceHeaders {
	s.AccountContext = v
	return s
}

func (s *CancelScheduleConferenceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelScheduleConferenceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CancelScheduleConferenceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CancelScheduleConferenceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CancelScheduleConferenceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CancelScheduleConferenceHeadersAccountContext) SetAccountId(v string) *CancelScheduleConferenceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CancelScheduleConferenceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
