// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWearOrgHonorShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *WearOrgHonorShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *WearOrgHonorShrinkHeaders
	GetAccountContextShrink() *string
}

type WearOrgHonorShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s WearOrgHonorShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s WearOrgHonorShrinkHeaders) GoString() string {
	return s.String()
}

func (s *WearOrgHonorShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *WearOrgHonorShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *WearOrgHonorShrinkHeaders) SetCommonHeaders(v map[string]*string) *WearOrgHonorShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *WearOrgHonorShrinkHeaders) SetAccountContextShrink(v string) *WearOrgHonorShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *WearOrgHonorShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
