// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetScheduleHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetScheduleHeadersAccountContext) *GetScheduleHeaders
	GetAccountContext() *GetScheduleHeadersAccountContext
}

type GetScheduleHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetScheduleHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetScheduleHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleHeaders) GoString() string {
	return s.String()
}

func (s *GetScheduleHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetScheduleHeaders) GetAccountContext() *GetScheduleHeadersAccountContext {
	return s.AccountContext
}

func (s *GetScheduleHeaders) SetCommonHeaders(v map[string]*string) *GetScheduleHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetScheduleHeaders) SetAccountContext(v *GetScheduleHeadersAccountContext) *GetScheduleHeaders {
	s.AccountContext = v
	return s
}

func (s *GetScheduleHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetScheduleHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetScheduleHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetScheduleHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetScheduleHeadersAccountContext) SetAccountId(v string) *GetScheduleHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetScheduleHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
