// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserHonorsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *QueryUserHonorsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *QueryUserHonorsHeadersAccountContext) *QueryUserHonorsHeaders
	GetAccountContext() *QueryUserHonorsHeadersAccountContext
}

type QueryUserHonorsHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *QueryUserHonorsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s QueryUserHonorsHeaders) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsHeaders) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *QueryUserHonorsHeaders) GetAccountContext() *QueryUserHonorsHeadersAccountContext {
	return s.AccountContext
}

func (s *QueryUserHonorsHeaders) SetCommonHeaders(v map[string]*string) *QueryUserHonorsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *QueryUserHonorsHeaders) SetAccountContext(v *QueryUserHonorsHeadersAccountContext) *QueryUserHonorsHeaders {
	s.AccountContext = v
	return s
}

func (s *QueryUserHonorsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryUserHonorsHeadersAccountContext struct {
	// example:
	//
	// 243331014234180628
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s QueryUserHonorsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s QueryUserHonorsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *QueryUserHonorsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *QueryUserHonorsHeadersAccountContext) SetAccountId(v string) *QueryUserHonorsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *QueryUserHonorsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
