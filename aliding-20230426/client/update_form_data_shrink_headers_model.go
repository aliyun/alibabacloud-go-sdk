// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFormDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *UpdateFormDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *UpdateFormDataShrinkHeaders
	GetAccountContextShrink() *string
}

type UpdateFormDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s UpdateFormDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s UpdateFormDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *UpdateFormDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *UpdateFormDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *UpdateFormDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *UpdateFormDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UpdateFormDataShrinkHeaders) SetAccountContextShrink(v string) *UpdateFormDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *UpdateFormDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
