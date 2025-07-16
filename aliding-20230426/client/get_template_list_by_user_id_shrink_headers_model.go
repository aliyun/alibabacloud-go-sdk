// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateListByUserIdShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTemplateListByUserIdShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetTemplateListByUserIdShrinkHeaders
	GetAccountContextShrink() *string
}

type GetTemplateListByUserIdShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetTemplateListByUserIdShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateListByUserIdShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetTemplateListByUserIdShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTemplateListByUserIdShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetTemplateListByUserIdShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetTemplateListByUserIdShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTemplateListByUserIdShrinkHeaders) SetAccountContextShrink(v string) *GetTemplateListByUserIdShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetTemplateListByUserIdShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
