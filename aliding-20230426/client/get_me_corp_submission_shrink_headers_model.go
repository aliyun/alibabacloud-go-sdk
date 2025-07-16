// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeCorpSubmissionShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetMeCorpSubmissionShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetMeCorpSubmissionShrinkHeaders
	GetAccountContextShrink() *string
}

type GetMeCorpSubmissionShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetMeCorpSubmissionShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetMeCorpSubmissionShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetMeCorpSubmissionShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetMeCorpSubmissionShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetMeCorpSubmissionShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetMeCorpSubmissionShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetMeCorpSubmissionShrinkHeaders) SetAccountContextShrink(v string) *GetMeCorpSubmissionShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetMeCorpSubmissionShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
