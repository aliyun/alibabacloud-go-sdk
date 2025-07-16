// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveWatchDetailShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryLiveWatchDetailShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryLiveWatchDetailShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryLiveWatchDetailShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveWatchDetailShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryLiveWatchDetailShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveWatchDetailShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveWatchDetailShrinkHeaders) SetAccountContextShrink(v string) *QueryLiveWatchDetailShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryLiveWatchDetailShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
