// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDingTypeShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *SyncDingTypeShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *SyncDingTypeShrinkHeaders
	GetAccountContextShrink() *string
}

type SyncDingTypeShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s SyncDingTypeShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s SyncDingTypeShrinkHeaders) GoString() string {
	return s.String()
}

func (s *SyncDingTypeShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *SyncDingTypeShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *SyncDingTypeShrinkHeaders) SetCommonHeaders(v map[string]*string) *SyncDingTypeShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *SyncDingTypeShrinkHeaders) SetAccountContextShrink(v string) *SyncDingTypeShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *SyncDingTypeShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
