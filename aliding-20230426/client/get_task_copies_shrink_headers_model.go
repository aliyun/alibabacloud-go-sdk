// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskCopiesShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetTaskCopiesShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetTaskCopiesShrinkHeaders
	GetAccountContextShrink() *string
}

type GetTaskCopiesShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetTaskCopiesShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetTaskCopiesShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetTaskCopiesShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetTaskCopiesShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetTaskCopiesShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetTaskCopiesShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetTaskCopiesShrinkHeaders) SetAccountContextShrink(v string) *GetTaskCopiesShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetTaskCopiesShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
