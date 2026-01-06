// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingInfoHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingInfoHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkMeetingInfoHeadersAccountContext) *GetDingtalkMeetingInfoHeaders
	GetAccountContext() *GetDingtalkMeetingInfoHeadersAccountContext
}

type GetDingtalkMeetingInfoHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkMeetingInfoHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkMeetingInfoHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingInfoHeaders) GetAccountContext() *GetDingtalkMeetingInfoHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkMeetingInfoHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingInfoHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingInfoHeaders) SetAccountContext(v *GetDingtalkMeetingInfoHeadersAccountContext) *GetDingtalkMeetingInfoHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkMeetingInfoHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingInfoHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkMeetingInfoHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingInfoHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingInfoHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkMeetingInfoHeadersAccountContext) SetAccountId(v string) *GetDingtalkMeetingInfoHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkMeetingInfoHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
