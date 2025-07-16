// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchMainOrgShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SwitchMainOrgShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SwitchMainOrgShrinkHeaders
	GetAccountContextShrink() *string
}

type SwitchMainOrgShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SwitchMainOrgShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SwitchMainOrgShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SwitchMainOrgShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SwitchMainOrgShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SwitchMainOrgShrinkHeaders) SetCommonHeaders(v map[string]*string) *SwitchMainOrgShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SwitchMainOrgShrinkHeaders) SetAccountContextShrink(v string) *SwitchMainOrgShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SwitchMainOrgShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
