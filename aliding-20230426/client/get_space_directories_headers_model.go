// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpaceDirectoriesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetSpaceDirectoriesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetSpaceDirectoriesHeadersAccountContext) *GetSpaceDirectoriesHeaders
	GetAccountContext() *GetSpaceDirectoriesHeadersAccountContext
}

type GetSpaceDirectoriesHeaders struct {
	CommonHeaders  map[string]*string                        `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetSpaceDirectoriesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetSpaceDirectoriesHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesHeaders) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetSpaceDirectoriesHeaders) GetAccountContext() *GetSpaceDirectoriesHeadersAccountContext {
	return s.AccountContext
}

func (s *GetSpaceDirectoriesHeaders) SetCommonHeaders(v map[string]*string) *GetSpaceDirectoriesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetSpaceDirectoriesHeaders) SetAccountContext(v *GetSpaceDirectoriesHeadersAccountContext) *GetSpaceDirectoriesHeaders {
	s.AccountContext = v
	return s
}

func (s *GetSpaceDirectoriesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSpaceDirectoriesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetSpaceDirectoriesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetSpaceDirectoriesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetSpaceDirectoriesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetSpaceDirectoriesHeadersAccountContext) SetAccountId(v string) *GetSpaceDirectoriesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetSpaceDirectoriesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
