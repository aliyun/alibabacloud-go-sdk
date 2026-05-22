// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserIdByOpenDingtalkIdHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserIdByOpenDingtalkIdHeadersAccountContext) *GetUserIdByOpenDingtalkIdHeaders
	GetAccountContext() *GetUserIdByOpenDingtalkIdHeadersAccountContext
}

type GetUserIdByOpenDingtalkIdHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserIdByOpenDingtalkIdHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserIdByOpenDingtalkIdHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdHeaders) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserIdByOpenDingtalkIdHeaders) GetAccountContext() *GetUserIdByOpenDingtalkIdHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserIdByOpenDingtalkIdHeaders) SetCommonHeaders(v map[string]*string) *GetUserIdByOpenDingtalkIdHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdHeaders) SetAccountContext(v *GetUserIdByOpenDingtalkIdHeadersAccountContext) *GetUserIdByOpenDingtalkIdHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserIdByOpenDingtalkIdHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserIdByOpenDingtalkIdHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserIdByOpenDingtalkIdHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserIdByOpenDingtalkIdHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserIdByOpenDingtalkIdHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserIdByOpenDingtalkIdHeadersAccountContext) SetAccountId(v string) *GetUserIdByOpenDingtalkIdHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserIdByOpenDingtalkIdHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
