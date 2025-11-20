// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCommentListReportHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *CommentListReportHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *CommentListReportHeadersAccountContext) *CommentListReportHeaders
	GetAccountContext() *CommentListReportHeadersAccountContext
}

type CommentListReportHeaders struct {
	CommonHeaders  map[string]*string                      `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *CommentListReportHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s CommentListReportHeaders) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportHeaders) GoString() string {
	return s.String()
}

func (s *CommentListReportHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *CommentListReportHeaders) GetAccountContext() *CommentListReportHeadersAccountContext {
	return s.AccountContext
}

func (s *CommentListReportHeaders) SetCommonHeaders(v map[string]*string) *CommentListReportHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CommentListReportHeaders) SetAccountContext(v *CommentListReportHeadersAccountContext) *CommentListReportHeaders {
	s.AccountContext = v
	return s
}

func (s *CommentListReportHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CommentListReportHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s CommentListReportHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s CommentListReportHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *CommentListReportHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *CommentListReportHeadersAccountContext) SetAccountId(v string) *CommentListReportHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *CommentListReportHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
