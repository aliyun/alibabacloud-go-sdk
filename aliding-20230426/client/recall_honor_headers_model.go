// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecallHonorHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *RecallHonorHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *RecallHonorHeadersAccountContext) *RecallHonorHeaders
	GetAccountContext() *RecallHonorHeadersAccountContext
}

type RecallHonorHeaders struct {
	CommonHeaders  map[string]*string                `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *RecallHonorHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s RecallHonorHeaders) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorHeaders) GoString() string {
	return s.String()
}

func (s *RecallHonorHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *RecallHonorHeaders) GetAccountContext() *RecallHonorHeadersAccountContext {
	return s.AccountContext
}

func (s *RecallHonorHeaders) SetCommonHeaders(v map[string]*string) *RecallHonorHeaders {
	s.CommonHeaders = v
	return s
}

func (s *RecallHonorHeaders) SetAccountContext(v *RecallHonorHeadersAccountContext) *RecallHonorHeaders {
	s.AccountContext = v
	return s
}

func (s *RecallHonorHeaders) Validate() error {
	return dara.Validate(s)
}

type RecallHonorHeadersAccountContext struct {
	// example:
	//
	// 1323845308033417
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s RecallHonorHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s RecallHonorHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *RecallHonorHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *RecallHonorHeadersAccountContext) SetAccountId(v string) *RecallHonorHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *RecallHonorHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
