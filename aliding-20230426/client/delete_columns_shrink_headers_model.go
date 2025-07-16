// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteColumnsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteColumnsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteColumnsShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteColumnsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteColumnsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteColumnsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteColumnsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteColumnsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteColumnsShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteColumnsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteColumnsShrinkHeaders) SetAccountContextShrink(v string) *DeleteColumnsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteColumnsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
