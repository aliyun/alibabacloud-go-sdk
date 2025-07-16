// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchRemovalByFormInstanceIdListShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchRemovalByFormInstanceIdListShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchRemovalByFormInstanceIdListShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchRemovalByFormInstanceIdListShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkHeaders) SetAccountContextShrink(v string) *BatchRemovalByFormInstanceIdListShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
