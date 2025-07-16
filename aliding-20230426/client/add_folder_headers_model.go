// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFolderHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddFolderHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddFolderHeadersAccountContext) *AddFolderHeaders
	GetAccountContext() *AddFolderHeadersAccountContext
}

type AddFolderHeaders struct {
	CommonHeaders  map[string]*string              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *AddFolderHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddFolderHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddFolderHeaders) GoString() string {
	return s.String()
}

func (s *AddFolderHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddFolderHeaders) GetAccountContext() *AddFolderHeadersAccountContext {
	return s.AccountContext
}

func (s *AddFolderHeaders) SetCommonHeaders(v map[string]*string) *AddFolderHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddFolderHeaders) SetAccountContext(v *AddFolderHeadersAccountContext) *AddFolderHeaders {
	s.AccountContext = v
	return s
}

func (s *AddFolderHeaders) Validate() error {
	return dara.Validate(s)
}

type AddFolderHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddFolderHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddFolderHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddFolderHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddFolderHeadersAccountContext) SetAccountId(v string) *AddFolderHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddFolderHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
