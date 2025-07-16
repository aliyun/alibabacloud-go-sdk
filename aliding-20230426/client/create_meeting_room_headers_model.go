// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMeetingRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateMeetingRoomHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CreateMeetingRoomHeadersAccountContext) *CreateMeetingRoomHeaders
	GetAccountContext() *CreateMeetingRoomHeadersAccountContext
}

type CreateMeetingRoomHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CreateMeetingRoomHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CreateMeetingRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateMeetingRoomHeaders) GetAccountContext() *CreateMeetingRoomHeadersAccountContext {
	return s.AccountContext
}

func (s *CreateMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *CreateMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateMeetingRoomHeaders) SetAccountContext(v *CreateMeetingRoomHeadersAccountContext) *CreateMeetingRoomHeaders {
	s.AccountContext = v
	return s
}

func (s *CreateMeetingRoomHeaders) Validate() error {
	return dara.Validate(s)
}

type CreateMeetingRoomHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CreateMeetingRoomHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CreateMeetingRoomHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CreateMeetingRoomHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateMeetingRoomHeadersAccountContext) SetAccountId(v string) *CreateMeetingRoomHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CreateMeetingRoomHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
