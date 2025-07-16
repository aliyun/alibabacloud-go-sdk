// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateMeetingRoomGroupHeadersAccountContext) *UpdateMeetingRoomGroupHeaders
	GetAccountContext() *UpdateMeetingRoomGroupHeadersAccountContext
}

type UpdateMeetingRoomGroupHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateMeetingRoomGroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateMeetingRoomGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMeetingRoomGroupHeaders) GetAccountContext() *UpdateMeetingRoomGroupHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomGroupHeaders) SetAccountContext(v *UpdateMeetingRoomGroupHeadersAccountContext) *UpdateMeetingRoomGroupHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateMeetingRoomGroupHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateMeetingRoomGroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateMeetingRoomGroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomGroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomGroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateMeetingRoomGroupHeadersAccountContext) SetAccountId(v string) *UpdateMeetingRoomGroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateMeetingRoomGroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
