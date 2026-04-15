// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *BatchQueryGroupMemberShrinkHeaders
	GetAccountContextShrink() *string
}

type BatchQueryGroupMemberShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s BatchQueryGroupMemberShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberShrinkHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchQueryGroupMemberShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *BatchQueryGroupMemberShrinkHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryGroupMemberShrinkHeaders) SetAccountContextShrink(v string) *BatchQueryGroupMemberShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *BatchQueryGroupMemberShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
