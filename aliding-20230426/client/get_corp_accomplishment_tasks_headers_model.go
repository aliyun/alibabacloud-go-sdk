// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCorpAccomplishmentTasksHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetCorpAccomplishmentTasksHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetCorpAccomplishmentTasksHeadersAccountContext) *GetCorpAccomplishmentTasksHeaders
	GetAccountContext() *GetCorpAccomplishmentTasksHeadersAccountContext
}

type GetCorpAccomplishmentTasksHeaders struct {
	CommonHeaders  map[string]*string                               `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetCorpAccomplishmentTasksHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetCorpAccomplishmentTasksHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksHeaders) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetCorpAccomplishmentTasksHeaders) GetAccountContext() *GetCorpAccomplishmentTasksHeadersAccountContext {
	return s.AccountContext
}

func (s *GetCorpAccomplishmentTasksHeaders) SetCommonHeaders(v map[string]*string) *GetCorpAccomplishmentTasksHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetCorpAccomplishmentTasksHeaders) SetAccountContext(v *GetCorpAccomplishmentTasksHeadersAccountContext) *GetCorpAccomplishmentTasksHeaders {
	s.AccountContext = v
	return s
}

func (s *GetCorpAccomplishmentTasksHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCorpAccomplishmentTasksHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetCorpAccomplishmentTasksHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetCorpAccomplishmentTasksHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetCorpAccomplishmentTasksHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetCorpAccomplishmentTasksHeadersAccountContext) SetAccountId(v string) *GetCorpAccomplishmentTasksHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetCorpAccomplishmentTasksHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
