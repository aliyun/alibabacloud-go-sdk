// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDocUpdateContentShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DocUpdateContentShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DocUpdateContentShrinkHeaders
	GetAccountContextShrink() *string
}

type DocUpdateContentShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DocUpdateContentShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DocUpdateContentShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DocUpdateContentShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DocUpdateContentShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DocUpdateContentShrinkHeaders) SetCommonHeaders(v map[string]*string) *DocUpdateContentShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DocUpdateContentShrinkHeaders) SetAccountContextShrink(v string) *DocUpdateContentShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DocUpdateContentShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
