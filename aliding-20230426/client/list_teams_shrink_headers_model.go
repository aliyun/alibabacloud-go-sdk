// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTeamsShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTeamsShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListTeamsShrinkHeaders
	GetAccountContextShrink() *string
}

type ListTeamsShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListTeamsShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTeamsShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListTeamsShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTeamsShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListTeamsShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListTeamsShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTeamsShrinkHeaders) SetAccountContextShrink(v string) *ListTeamsShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListTeamsShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
