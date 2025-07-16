// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RecallHonorShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *RecallHonorShrinkHeaders
	GetAccountContextShrink() *string
}

type RecallHonorShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s RecallHonorShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorShrinkHeaders) GoString() string {
	return s.String()
}

func (s *RecallHonorShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RecallHonorShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *RecallHonorShrinkHeaders) SetCommonHeaders(v map[string]*string) *RecallHonorShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallHonorShrinkHeaders) SetAccountContextShrink(v string) *RecallHonorShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *RecallHonorShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
