// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *MeetingFlashMinutesHeadersAccountContext) *MeetingFlashMinutesHeaders
	GetAccountContext() *MeetingFlashMinutesHeadersAccountContext
}

type MeetingFlashMinutesHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *MeetingFlashMinutesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s MeetingFlashMinutesHeaders) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesHeaders) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MeetingFlashMinutesHeaders) GetAccountContext() *MeetingFlashMinutesHeadersAccountContext {
	return s.AccountContext
}

func (s *MeetingFlashMinutesHeaders) SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MeetingFlashMinutesHeaders) SetAccountContext(v *MeetingFlashMinutesHeadersAccountContext) *MeetingFlashMinutesHeaders {
	s.AccountContext = v
	return s
}

func (s *MeetingFlashMinutesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s MeetingFlashMinutesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *MeetingFlashMinutesHeadersAccountContext) SetAccountId(v string) *MeetingFlashMinutesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *MeetingFlashMinutesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
