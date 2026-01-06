// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMemberListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkMeetingMemberListHeadersAccountContext) *GetDingtalkMeetingMemberListHeaders
	GetAccountContext() *GetDingtalkMeetingMemberListHeadersAccountContext
}

type GetDingtalkMeetingMemberListHeaders struct {
	CommonHeaders  map[string]*string                                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkMeetingMemberListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkMeetingMemberListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingMemberListHeaders) GetAccountContext() *GetDingtalkMeetingMemberListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkMeetingMemberListHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMemberListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingMemberListHeaders) SetAccountContext(v *GetDingtalkMeetingMemberListHeadersAccountContext) *GetDingtalkMeetingMemberListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkMeetingMemberListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingMemberListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkMeetingMemberListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMemberListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMemberListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkMeetingMemberListHeadersAccountContext) SetAccountId(v string) *GetDingtalkMeetingMemberListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkMeetingMemberListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
