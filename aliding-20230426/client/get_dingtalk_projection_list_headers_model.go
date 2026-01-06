// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionListHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionListHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetDingtalkProjectionListHeadersAccountContext) *GetDingtalkProjectionListHeaders
	GetAccountContext() *GetDingtalkProjectionListHeadersAccountContext
}

type GetDingtalkProjectionListHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetDingtalkProjectionListHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetDingtalkProjectionListHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListHeaders) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetDingtalkProjectionListHeaders) GetAccountContext() *GetDingtalkProjectionListHeadersAccountContext {
	return s.AccountContext
}

func (s *GetDingtalkProjectionListHeaders) SetCommonHeaders(v map[string]*string) *GetDingtalkProjectionListHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetDingtalkProjectionListHeaders) SetAccountContext(v *GetDingtalkProjectionListHeadersAccountContext) *GetDingtalkProjectionListHeaders {
	s.AccountContext = v
	return s
}

func (s *GetDingtalkProjectionListHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkProjectionListHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetDingtalkProjectionListHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionListHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionListHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetDingtalkProjectionListHeadersAccountContext) SetAccountId(v string) *GetDingtalkProjectionListHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetDingtalkProjectionListHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
