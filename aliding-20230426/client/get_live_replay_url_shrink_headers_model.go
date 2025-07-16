// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLiveReplayUrlShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetLiveReplayUrlShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetLiveReplayUrlShrinkHeaders
	GetAccountContextShrink() *string
}

type GetLiveReplayUrlShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetLiveReplayUrlShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetLiveReplayUrlShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetLiveReplayUrlShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetLiveReplayUrlShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetLiveReplayUrlShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetLiveReplayUrlShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetLiveReplayUrlShrinkHeaders) SetAccountContextShrink(v string) *GetLiveReplayUrlShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetLiveReplayUrlShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
