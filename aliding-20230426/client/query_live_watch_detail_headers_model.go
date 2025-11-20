// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveWatchDetailHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryLiveWatchDetailHeadersAccountContext) *QueryLiveWatchDetailHeaders
	GetAccountContext() *QueryLiveWatchDetailHeadersAccountContext
}

type QueryLiveWatchDetailHeaders struct {
	CommonHeaders  map[string]*string                         `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryLiveWatchDetailHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryLiveWatchDetailHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveWatchDetailHeaders) GetAccountContext() *QueryLiveWatchDetailHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryLiveWatchDetailHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchDetailHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchDetailHeaders) SetAccountContext(v *QueryLiveWatchDetailHeadersAccountContext) *QueryLiveWatchDetailHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryLiveWatchDetailHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLiveWatchDetailHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryLiveWatchDetailHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryLiveWatchDetailHeadersAccountContext) SetAccountId(v string) *QueryLiveWatchDetailHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryLiveWatchDetailHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
