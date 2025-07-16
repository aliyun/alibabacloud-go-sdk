// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSpaceDirectoriesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetSpaceDirectoriesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetSpaceDirectoriesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetSpaceDirectoriesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSpaceDirectoriesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetSpaceDirectoriesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceDirectoriesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceDirectoriesShrinkHeaders) SetAccountContextShrink(v string) *GetSpaceDirectoriesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetSpaceDirectoriesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
