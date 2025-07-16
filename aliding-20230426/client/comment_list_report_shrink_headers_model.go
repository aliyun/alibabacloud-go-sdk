// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportShrinkHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommentListReportShrinkHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContextShrink(v string) *CommentListReportShrinkHeaders
	GetAccountContextShrink() *string
}

type CommentListReportShrinkHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContextShrink *string            `json:"AccountContext,omitempty" xml:"AccountContext,omitempty"`
}

func (s CommentListReportShrinkHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportShrinkHeaders) GoString() string {
	return s.String()
}

func (s *CommentListReportShrinkHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommentListReportShrinkHeaders) GetAccountContextShrink() *string {
	return s.AccountContextShrink
}

func (s *CommentListReportShrinkHeaders) SetCommonHeaders(v map[string]*string) *CommentListReportShrinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommentListReportShrinkHeaders) SetAccountContextShrink(v string) *CommentListReportShrinkHeaders {
	s.AccountContextShrink = &v
	return s
}

func (s *CommentListReportShrinkHeaders) Validate() error {
	return dara.Validate(s)
}
