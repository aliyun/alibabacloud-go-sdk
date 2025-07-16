// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteInstanceShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteInstanceShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteInstanceShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteInstanceShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteInstanceShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteInstanceShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteInstanceShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteInstanceShrinkHeaders) SetAccountContextShrink(v string) *DeleteInstanceShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteInstanceShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
