// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupLiveInfoShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryGroupLiveInfoShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryGroupLiveInfoShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryGroupLiveInfoShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryGroupLiveInfoShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupLiveInfoShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryGroupLiveInfoShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryGroupLiveInfoShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryGroupLiveInfoShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryGroupLiveInfoShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryGroupLiveInfoShrinkHeaders) SetAccountContextShrink(v string) *QueryGroupLiveInfoShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryGroupLiveInfoShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
