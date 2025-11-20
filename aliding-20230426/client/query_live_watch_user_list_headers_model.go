// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveWatchUserListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryLiveWatchUserListHeadersAccountContext) *QueryLiveWatchUserListHeaders
	GetAccountContext() *QueryLiveWatchUserListHeadersAccountContext
}

type QueryLiveWatchUserListHeaders struct {
	CommonHeaders  map[string]*string                           `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryLiveWatchUserListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryLiveWatchUserListHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveWatchUserListHeaders) GetAccountContext() *QueryLiveWatchUserListHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryLiveWatchUserListHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchUserListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchUserListHeaders) SetAccountContext(v *QueryLiveWatchUserListHeadersAccountContext) *QueryLiveWatchUserListHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryLiveWatchUserListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryLiveWatchUserListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryLiveWatchUserListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryLiveWatchUserListHeadersAccountContext) SetAccountId(v string) *QueryLiveWatchUserListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryLiveWatchUserListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
