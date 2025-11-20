// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomGroupHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomGroupHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteMeetingRoomGroupHeadersAccountContext) *DeleteMeetingRoomGroupHeaders
	GetAccountContext() *DeleteMeetingRoomGroupHeadersAccountContext
}

type DeleteMeetingRoomGroupHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteMeetingRoomGroupHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteMeetingRoomGroupHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMeetingRoomGroupHeaders) GetAccountContext() *DeleteMeetingRoomGroupHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteMeetingRoomGroupHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomGroupHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomGroupHeaders) SetAccountContext(v *DeleteMeetingRoomGroupHeadersAccountContext) *DeleteMeetingRoomGroupHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteMeetingRoomGroupHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMeetingRoomGroupHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteMeetingRoomGroupHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomGroupHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomGroupHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMeetingRoomGroupHeadersAccountContext) SetAccountId(v string) *DeleteMeetingRoomGroupHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteMeetingRoomGroupHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
