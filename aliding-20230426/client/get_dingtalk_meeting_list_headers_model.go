// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkMeetingListHeadersAccountContext) *GetDingtalkMeetingListHeaders
	GetAccountContext() *GetDingtalkMeetingListHeadersAccountContext
}

type GetDingtalkMeetingListHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkMeetingListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkMeetingListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingListHeaders) GetAccountContext() *GetDingtalkMeetingListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkMeetingListHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingListHeaders) SetAccountContext(v *GetDingtalkMeetingListHeadersAccountContext) *GetDingtalkMeetingListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkMeetingListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkMeetingListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkMeetingListHeadersAccountContext) SetAccountId(v string) *GetDingtalkMeetingListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkMeetingListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
