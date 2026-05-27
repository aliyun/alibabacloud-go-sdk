// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetUserIdByOpenDingtalkIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) *BatchGetUserIdByOpenDingtalkIdHeaders
	GetAccountContext() *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext
}

type BatchGetUserIdByOpenDingtalkIdHeaders struct {
	CommonHeaders  map[string]*string                                   `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s BatchGetUserIdByOpenDingtalkIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdHeaders) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *BatchGetUserIdByOpenDingtalkIdHeaders) GetAccountContext() *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext {
	return s.AccountContext
}

func (s *BatchGetUserIdByOpenDingtalkIdHeaders) SetCommonHeaders(v map[string]*string) *BatchGetUserIdByOpenDingtalkIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdHeaders) SetAccountContext(v *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) *BatchGetUserIdByOpenDingtalkIdHeaders {
	s.AccountContext = v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchGetUserIdByOpenDingtalkIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) SetAccountId(v string) *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *BatchGetUserIdByOpenDingtalkIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
