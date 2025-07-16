// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateFormDataShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CreateOrUpdateFormDataShrinkHeaders
	GetAccountContextShrink() *string
}

type CreateOrUpdateFormDataShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CreateOrUpdateFormDataShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CreateOrUpdateFormDataShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CreateOrUpdateFormDataShrinkHeaders) SetCommonHeaders(v map[string]*string) *CreateOrUpdateFormDataShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CreateOrUpdateFormDataShrinkHeaders) SetAccountContextShrink(v string) *CreateOrUpdateFormDataShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CreateOrUpdateFormDataShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
