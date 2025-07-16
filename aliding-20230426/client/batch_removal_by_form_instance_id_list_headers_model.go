// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchRemovalByFormInstanceIdListHeadersAccountContext) *BatchRemovalByFormInstanceIdListHeaders
	GetAccountContext() *BatchRemovalByFormInstanceIdListHeadersAccountContext
}

type BatchRemovalByFormInstanceIdListHeaders struct {
	CommonHeaders  map[string]*string                                     `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchRemovalByFormInstanceIdListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchRemovalByFormInstanceIdListHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListHeaders) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchRemovalByFormInstanceIdListHeaders) GetAccountContext() *BatchRemovalByFormInstanceIdListHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchRemovalByFormInstanceIdListHeaders) SetCommonHeaders(v map[string]*string) *BatchRemovalByFormInstanceIdListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListHeaders) SetAccountContext(v *BatchRemovalByFormInstanceIdListHeadersAccountContext) *BatchRemovalByFormInstanceIdListHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchRemovalByFormInstanceIdListHeaders) Validate() error {
	return dara.Validate(s)
}

type BatchRemovalByFormInstanceIdListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchRemovalByFormInstanceIdListHeadersAccountContext) SetAccountId(v string) *BatchRemovalByFormInstanceIdListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
