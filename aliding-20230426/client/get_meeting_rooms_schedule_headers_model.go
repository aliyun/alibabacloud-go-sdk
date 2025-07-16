// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMeetingRoomsScheduleHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetMeetingRoomsScheduleHeadersAccountContext) *GetMeetingRoomsScheduleHeaders
	GetAccountContext() *GetMeetingRoomsScheduleHeadersAccountContext
}

type GetMeetingRoomsScheduleHeaders struct {
	CommonHeaders  map[string]*string                            `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetMeetingRoomsScheduleHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetMeetingRoomsScheduleHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleHeaders) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMeetingRoomsScheduleHeaders) GetAccountContext() *GetMeetingRoomsScheduleHeadersAccountContext {
	return s.AccountContext
}

func (s *GetMeetingRoomsScheduleHeaders) SetCommonHeaders(v map[string]*string) *GetMeetingRoomsScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMeetingRoomsScheduleHeaders) SetAccountContext(v *GetMeetingRoomsScheduleHeadersAccountContext) *GetMeetingRoomsScheduleHeaders {
	s.AccountContext = v
	return s
}

func (s *GetMeetingRoomsScheduleHeaders) Validate() error {
	return dara.Validate(s)
}

type GetMeetingRoomsScheduleHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetMeetingRoomsScheduleHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetMeetingRoomsScheduleHeadersAccountContext) SetAccountId(v string) *GetMeetingRoomsScheduleHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetMeetingRoomsScheduleHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
