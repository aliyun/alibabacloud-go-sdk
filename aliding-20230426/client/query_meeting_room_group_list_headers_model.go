// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMeetingRoomGroupListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryMeetingRoomGroupListHeadersAccountContext) *QueryMeetingRoomGroupListHeaders
	GetAccountContext() *QueryMeetingRoomGroupListHeadersAccountContext
}

type QueryMeetingRoomGroupListHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryMeetingRoomGroupListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryMeetingRoomGroupListHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListHeaders) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryMeetingRoomGroupListHeaders) GetAccountContext() *QueryMeetingRoomGroupListHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryMeetingRoomGroupListHeaders) SetCommonHeaders(v map[string]*string) *QueryMeetingRoomGroupListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryMeetingRoomGroupListHeaders) SetAccountContext(v *QueryMeetingRoomGroupListHeadersAccountContext) *QueryMeetingRoomGroupListHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryMeetingRoomGroupListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMeetingRoomGroupListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryMeetingRoomGroupListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryMeetingRoomGroupListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryMeetingRoomGroupListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryMeetingRoomGroupListHeadersAccountContext) SetAccountId(v string) *QueryMeetingRoomGroupListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryMeetingRoomGroupListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
