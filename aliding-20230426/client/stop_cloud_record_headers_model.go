// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCloudRecordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StopCloudRecordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StopCloudRecordHeadersAccountContext) *StopCloudRecordHeaders
	GetAccountContext() *StopCloudRecordHeadersAccountContext
}

type StopCloudRecordHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StopCloudRecordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StopCloudRecordHeaders) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StopCloudRecordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StopCloudRecordHeaders) GetAccountContext() *StopCloudRecordHeadersAccountContext {
	return s.AccountContext
}

func (s *StopCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StopCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StopCloudRecordHeaders) SetAccountContext(v *StopCloudRecordHeadersAccountContext) *StopCloudRecordHeaders {
	s.AccountContext = v
	return s
}

func (s *StopCloudRecordHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StopCloudRecordHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StopCloudRecordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StopCloudRecordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StopCloudRecordHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StopCloudRecordHeadersAccountContext) SetAccountId(v string) *StopCloudRecordHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StopCloudRecordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
