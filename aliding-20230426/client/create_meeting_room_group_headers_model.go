// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMeetingRoomGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateMeetingRoomGroupHeadersAccountContext) *CreateMeetingRoomGroupHeaders
	GetAccountContext() *CreateMeetingRoomGroupHeadersAccountContext
}

type CreateMeetingRoomGroupHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateMeetingRoomGroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateMeetingRoomGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMeetingRoomGroupHeaders) GetAccountContext() *CreateMeetingRoomGroupHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomGroupHeaders) SetAccountContext(v *CreateMeetingRoomGroupHeadersAccountContext) *CreateMeetingRoomGroupHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateMeetingRoomGroupHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomGroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateMeetingRoomGroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomGroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomGroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateMeetingRoomGroupHeadersAccountContext) SetAccountId(v string) *CreateMeetingRoomGroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateMeetingRoomGroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
