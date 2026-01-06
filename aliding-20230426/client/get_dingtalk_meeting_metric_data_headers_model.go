// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkMeetingMetricDataHeadersAccountContext) *GetDingtalkMeetingMetricDataHeaders
	GetAccountContext() *GetDingtalkMeetingMetricDataHeadersAccountContext
}

type GetDingtalkMeetingMetricDataHeaders struct {
	CommonHeaders  map[string]*string                                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkMeetingMetricDataHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkMeetingMetricDataHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkMeetingMetricDataHeaders) GetAccountContext() *GetDingtalkMeetingMetricDataHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkMeetingMetricDataHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkMeetingMetricDataHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkMeetingMetricDataHeaders) SetAccountContext(v *GetDingtalkMeetingMetricDataHeadersAccountContext) *GetDingtalkMeetingMetricDataHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkMeetingMetricDataHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkMeetingMetricDataHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkMeetingMetricDataHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkMeetingMetricDataHeadersAccountContext) SetAccountId(v string) *GetDingtalkMeetingMetricDataHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
