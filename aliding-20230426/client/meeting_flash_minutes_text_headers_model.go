// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMeetingFlashMinutesTextHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesTextHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *MeetingFlashMinutesTextHeadersAccountContext) *MeetingFlashMinutesTextHeaders
	GetAccountContext() *MeetingFlashMinutesTextHeadersAccountContext
}

type MeetingFlashMinutesTextHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *MeetingFlashMinutesTextHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s MeetingFlashMinutesTextHeaders) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextHeaders) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *MeetingFlashMinutesTextHeaders) GetAccountContext() *MeetingFlashMinutesTextHeadersAccountContext {
	return s.AccountContext
}

func (s *MeetingFlashMinutesTextHeaders) SetCommonHeaders(v map[string]*string) *MeetingFlashMinutesTextHeaders {
	s.CommonHeaders = v
	return s
}

func (s *MeetingFlashMinutesTextHeaders) SetAccountContext(v *MeetingFlashMinutesTextHeadersAccountContext) *MeetingFlashMinutesTextHeaders {
	s.AccountContext = v
	return s
}

func (s *MeetingFlashMinutesTextHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MeetingFlashMinutesTextHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s MeetingFlashMinutesTextHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s MeetingFlashMinutesTextHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *MeetingFlashMinutesTextHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *MeetingFlashMinutesTextHeadersAccountContext) SetAccountId(v string) *MeetingFlashMinutesTextHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *MeetingFlashMinutesTextHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
