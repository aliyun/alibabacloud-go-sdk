// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteLiveShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteLiveShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteLiveShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteLiveShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteLiveShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteLiveShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteLiveShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteLiveShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteLiveShrinkHeaders) SetAccountContextShrink(v string) *DeleteLiveShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteLiveShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
