// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMeetingRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *UpdateMeetingRoomHeadersAccountContext) *UpdateMeetingRoomHeaders
	GetAccountContext() *UpdateMeetingRoomHeadersAccountContext
}

type UpdateMeetingRoomHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *UpdateMeetingRoomHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s UpdateMeetingRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateMeetingRoomHeaders) GetAccountContext() *UpdateMeetingRoomHeadersAccountContext {
	return s.AccountContext
}

func (s *UpdateMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *UpdateMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateMeetingRoomHeaders) SetAccountContext(v *UpdateMeetingRoomHeadersAccountContext) *UpdateMeetingRoomHeaders {
	s.AccountContext = v
	return s
}

func (s *UpdateMeetingRoomHeaders) Validate() error {
	return dara.Validate(s)
}

type UpdateMeetingRoomHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s UpdateMeetingRoomHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s UpdateMeetingRoomHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *UpdateMeetingRoomHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *UpdateMeetingRoomHeadersAccountContext) SetAccountId(v string) *UpdateMeetingRoomHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *UpdateMeetingRoomHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
