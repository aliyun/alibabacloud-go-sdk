// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommitFileShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommitFileShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CommitFileShrinkHeaders
	GetAccountContextShrink() *string
}

type CommitFileShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CommitFileShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommitFileShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CommitFileShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommitFileShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CommitFileShrinkHeaders) SetCommonHeaders(v map[string]*string) *CommitFileShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommitFileShrinkHeaders) SetAccountContextShrink(v string) *CommitFileShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CommitFileShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
