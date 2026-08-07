// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDocumentPermissionHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *GetUserDocumentPermissionHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *GetUserDocumentPermissionHeadersAccountContext) *GetUserDocumentPermissionHeaders
	GetAccountContext() *GetUserDocumentPermissionHeadersAccountContext
}

type GetUserDocumentPermissionHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *GetUserDocumentPermissionHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s GetUserDocumentPermissionHeaders) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionHeaders) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *GetUserDocumentPermissionHeaders) GetAccountContext() *GetUserDocumentPermissionHeadersAccountContext {
	return s.AccountContext
}

func (s *GetUserDocumentPermissionHeaders) SetCommonHeaders(v map[string]*string) *GetUserDocumentPermissionHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetUserDocumentPermissionHeaders) SetAccountContext(v *GetUserDocumentPermissionHeadersAccountContext) *GetUserDocumentPermissionHeaders {
	s.AccountContext = v
	return s
}

func (s *GetUserDocumentPermissionHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserDocumentPermissionHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s GetUserDocumentPermissionHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s GetUserDocumentPermissionHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *GetUserDocumentPermissionHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *GetUserDocumentPermissionHeadersAccountContext) SetAccountId(v string) *GetUserDocumentPermissionHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *GetUserDocumentPermissionHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
