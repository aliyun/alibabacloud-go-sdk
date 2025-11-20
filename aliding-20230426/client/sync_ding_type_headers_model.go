// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SyncDingTypeHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *SyncDingTypeHeadersAccountContext) *SyncDingTypeHeaders
	GetAccountContext() *SyncDingTypeHeadersAccountContext
}

type SyncDingTypeHeaders struct {
	CommonHeaders  map[string]*string                 `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *SyncDingTypeHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s SyncDingTypeHeaders) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeHeaders) GoString() string {
	return s.String()
}

func (s *SyncDingTypeHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SyncDingTypeHeaders) GetAccountContext() *SyncDingTypeHeadersAccountContext {
	return s.AccountContext
}

func (s *SyncDingTypeHeaders) SetCommonHeaders(v map[string]*string) *SyncDingTypeHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDingTypeHeaders) SetAccountContext(v *SyncDingTypeHeadersAccountContext) *SyncDingTypeHeaders {
	s.AccountContext = v
	return s
}

func (s *SyncDingTypeHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SyncDingTypeHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s SyncDingTypeHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *SyncDingTypeHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *SyncDingTypeHeadersAccountContext) SetAccountId(v string) *SyncDingTypeHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *SyncDingTypeHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
