// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFormDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteFormDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *DeleteFormDataShrinkHeaders
	GetAccountContextShrink() *string
}

type DeleteFormDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s DeleteFormDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteFormDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *DeleteFormDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteFormDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *DeleteFormDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *DeleteFormDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteFormDataShrinkHeaders) SetAccountContextShrink(v string) *DeleteFormDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *DeleteFormDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
