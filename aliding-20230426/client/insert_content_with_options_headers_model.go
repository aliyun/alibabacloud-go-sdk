// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertContentWithOptionsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *InsertContentWithOptionsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *InsertContentWithOptionsHeadersAccountContext) *InsertContentWithOptionsHeaders
	GetAccountContext() *InsertContentWithOptionsHeadersAccountContext
}

type InsertContentWithOptionsHeaders struct {
	CommonHeaders  map[string]*string                             `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *InsertContentWithOptionsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s InsertContentWithOptionsHeaders) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsHeaders) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *InsertContentWithOptionsHeaders) GetAccountContext() *InsertContentWithOptionsHeadersAccountContext {
	return s.AccountContext
}

func (s *InsertContentWithOptionsHeaders) SetCommonHeaders(v map[string]*string) *InsertContentWithOptionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InsertContentWithOptionsHeaders) SetAccountContext(v *InsertContentWithOptionsHeadersAccountContext) *InsertContentWithOptionsHeaders {
	s.AccountContext = v
	return s
}

func (s *InsertContentWithOptionsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertContentWithOptionsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s InsertContentWithOptionsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s InsertContentWithOptionsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *InsertContentWithOptionsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *InsertContentWithOptionsHeadersAccountContext) SetAccountId(v string) *InsertContentWithOptionsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *InsertContentWithOptionsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
