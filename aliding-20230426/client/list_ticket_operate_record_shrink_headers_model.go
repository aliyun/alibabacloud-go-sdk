// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTicketOperateRecordShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *ListTicketOperateRecordShrinkHeaders
	GetAccountContextShrink() *string
}

type ListTicketOperateRecordShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s ListTicketOperateRecordShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListTicketOperateRecordShrinkHeaders) GoString() string {
	return s.String()
}

func (s *ListTicketOperateRecordShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListTicketOperateRecordShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *ListTicketOperateRecordShrinkHeaders) SetCommonHeaders(v map[string]*string) *ListTicketOperateRecordShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListTicketOperateRecordShrinkHeaders) SetAccountContextShrink(v string) *ListTicketOperateRecordShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *ListTicketOperateRecordShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
