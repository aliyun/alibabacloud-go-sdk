// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveWatchUserListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryLiveWatchUserListShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryLiveWatchUserListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryLiveWatchUserListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveWatchUserListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryLiveWatchUserListShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchUserListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchUserListShrinkHeaders) SetAccountContextShrink(v string) *QueryLiveWatchUserListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryLiveWatchUserListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
