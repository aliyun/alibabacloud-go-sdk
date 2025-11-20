// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMinutesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopMinutesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StopMinutesHeadersAccountContext) *StopMinutesHeaders
	GetAccountContext() *StopMinutesHeadersAccountContext
}

type StopMinutesHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StopMinutesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StopMinutesHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesHeaders) GoString() string {
	return s.String()
}

func (s *StopMinutesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopMinutesHeaders) GetAccountContext() *StopMinutesHeadersAccountContext {
	return s.AccountContext
}

func (s *StopMinutesHeaders) SetCommonHeaders(v map[string]*string) *StopMinutesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopMinutesHeaders) SetAccountContext(v *StopMinutesHeadersAccountContext) *StopMinutesHeaders {
	s.AccountContext = v
	return s
}

func (s *StopMinutesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopMinutesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StopMinutesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StopMinutesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StopMinutesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StopMinutesHeadersAccountContext) SetAccountId(v string) *StopMinutesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StopMinutesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
