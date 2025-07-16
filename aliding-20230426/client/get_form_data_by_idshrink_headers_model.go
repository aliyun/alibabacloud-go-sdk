// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormDataByIDShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormDataByIDShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetFormDataByIDShrinkHeaders
	GetAccountContextShrink() *string
}

type GetFormDataByIDShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetFormDataByIDShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormDataByIDShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetFormDataByIDShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormDataByIDShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetFormDataByIDShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetFormDataByIDShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormDataByIDShrinkHeaders) SetAccountContextShrink(v string) *GetFormDataByIDShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetFormDataByIDShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
