// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMeetingRoomHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteMeetingRoomHeadersAccountContext) *DeleteMeetingRoomHeaders
	GetAccountContext() *DeleteMeetingRoomHeadersAccountContext
}

type DeleteMeetingRoomHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DeleteMeetingRoomHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteMeetingRoomHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomHeaders) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteMeetingRoomHeaders) GetAccountContext() *DeleteMeetingRoomHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteMeetingRoomHeaders) SetCommonHeaders(v map[string]*string) *DeleteMeetingRoomHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteMeetingRoomHeaders) SetAccountContext(v *DeleteMeetingRoomHeadersAccountContext) *DeleteMeetingRoomHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteMeetingRoomHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteMeetingRoomHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteMeetingRoomHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteMeetingRoomHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteMeetingRoomHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteMeetingRoomHeadersAccountContext) SetAccountId(v string) *DeleteMeetingRoomHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteMeetingRoomHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
