// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetProcessDefinitionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetProcessDefinitionShrinkHeaders
	GetAccountContextShrink() *string
}

type GetProcessDefinitionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetProcessDefinitionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetProcessDefinitionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetProcessDefinitionShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetProcessDefinitionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetProcessDefinitionShrinkHeaders) SetAccountContextShrink(v string) *GetProcessDefinitionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetProcessDefinitionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
