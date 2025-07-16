// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryUserHonorsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *QueryUserHonorsShrinkHeaders
	GetAccountContextShrink() *string
}

type QueryUserHonorsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s QueryUserHonorsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryUserHonorsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *QueryUserHonorsShrinkHeaders) SetCommonHeaders(v map[string]*string) *QueryUserHonorsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserHonorsShrinkHeaders) SetAccountContextShrink(v string) *QueryUserHonorsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *QueryUserHonorsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
