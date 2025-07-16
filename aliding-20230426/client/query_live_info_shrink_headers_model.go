// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryLiveInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryLiveInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryLiveInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryLiveInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryLiveInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryLiveInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryLiveInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryLiveInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryLiveInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryLiveInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
