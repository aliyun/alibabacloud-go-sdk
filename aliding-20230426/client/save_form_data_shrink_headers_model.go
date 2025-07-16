// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveFormDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SaveFormDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SaveFormDataShrinkHeaders
	GetAccountContextShrink() *string
}

type SaveFormDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SaveFormDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SaveFormDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SaveFormDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SaveFormDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SaveFormDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *SaveFormDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SaveFormDataShrinkHeaders) SetAccountContextShrink(v string) *SaveFormDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SaveFormDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
