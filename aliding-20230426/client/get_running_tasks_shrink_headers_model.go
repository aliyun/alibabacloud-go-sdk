// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunningTasksShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetRunningTasksShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *GetRunningTasksShrinkHeaders
	GetAccountContextShrink() *string
}

type GetRunningTasksShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s GetRunningTasksShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetRunningTasksShrinkHeaders) GoString() string {
	return s.String()
}

func (s *GetRunningTasksShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetRunningTasksShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *GetRunningTasksShrinkHeaders) SetCommonHeaders(v map[string]*string) *GetRunningTasksShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetRunningTasksShrinkHeaders) SetAccountContextShrink(v string) *GetRunningTasksShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *GetRunningTasksShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
