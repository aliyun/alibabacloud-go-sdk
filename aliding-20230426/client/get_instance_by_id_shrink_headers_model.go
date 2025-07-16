// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceByIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetInstanceByIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetInstanceByIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetInstanceByIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetInstanceByIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceByIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetInstanceByIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetInstanceByIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetInstanceByIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetInstanceByIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetInstanceByIdShrinkHeaders) SetAccountContextShrink(v string) *GetInstanceByIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetInstanceByIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
