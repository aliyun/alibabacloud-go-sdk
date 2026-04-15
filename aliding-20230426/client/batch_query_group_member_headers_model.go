// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryGroupMemberHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchQueryGroupMemberHeadersAccountContext) *BatchQueryGroupMemberHeaders
	GetAccountContext() *BatchQueryGroupMemberHeadersAccountContext
}

type BatchQueryGroupMemberHeaders struct {
	CommonHeaders  map[string]*string                          `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchQueryGroupMemberHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchQueryGroupMemberHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberHeaders) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchQueryGroupMemberHeaders) GetAccountContext() *BatchQueryGroupMemberHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchQueryGroupMemberHeaders) SetCommonHeaders(v map[string]*string) *BatchQueryGroupMemberHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchQueryGroupMemberHeaders) SetAccountContext(v *BatchQueryGroupMemberHeadersAccountContext) *BatchQueryGroupMemberHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchQueryGroupMemberHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchQueryGroupMemberHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchQueryGroupMemberHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryGroupMemberHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchQueryGroupMemberHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchQueryGroupMemberHeadersAccountContext) SetAccountId(v string) *BatchQueryGroupMemberHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchQueryGroupMemberHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
