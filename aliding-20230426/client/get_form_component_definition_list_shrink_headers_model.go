// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFormComponentDefinitionListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetFormComponentDefinitionListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetFormComponentDefinitionListShrinkHeaders
	GetAccountContextShrink() *string
}

type GetFormComponentDefinitionListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetFormComponentDefinitionListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetFormComponentDefinitionListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetFormComponentDefinitionListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetFormComponentDefinitionListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetFormComponentDefinitionListShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetFormComponentDefinitionListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetFormComponentDefinitionListShrinkHeaders) SetAccountContextShrink(v string) *GetFormComponentDefinitionListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetFormComponentDefinitionListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
