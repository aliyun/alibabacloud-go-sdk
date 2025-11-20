// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCloudRecordHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *StartCloudRecordHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *StartCloudRecordHeadersAccountContext) *StartCloudRecordHeaders
	GetAccountContext() *StartCloudRecordHeadersAccountContext
}

type StartCloudRecordHeaders struct {
	CommonHeaders  map[string]*string                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *StartCloudRecordHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s StartCloudRecordHeaders) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordHeaders) GoString() string {
	return s.String()
}

func (s *StartCloudRecordHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *StartCloudRecordHeaders) GetAccountContext() *StartCloudRecordHeadersAccountContext {
	return s.AccountContext
}

func (s *StartCloudRecordHeaders) SetCommonHeaders(v map[string]*string) *StartCloudRecordHeaders {
	s.CommonHeaders = v
	return s
}

func (s *StartCloudRecordHeaders) SetAccountContext(v *StartCloudRecordHeadersAccountContext) *StartCloudRecordHeaders {
	s.AccountContext = v
	return s
}

func (s *StartCloudRecordHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartCloudRecordHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s StartCloudRecordHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s StartCloudRecordHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *StartCloudRecordHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *StartCloudRecordHeadersAccountContext) SetAccountId(v string) *StartCloudRecordHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *StartCloudRecordHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
