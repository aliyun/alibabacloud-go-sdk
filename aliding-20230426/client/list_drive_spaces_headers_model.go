// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDriveSpacesHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListDriveSpacesHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListDriveSpacesHeadersAccountContext) *ListDriveSpacesHeaders
	GetAccountContext() *ListDriveSpacesHeadersAccountContext
}

type ListDriveSpacesHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContext *ListDriveSpacesHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListDriveSpacesHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesHeaders) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListDriveSpacesHeaders) GetAccountContext() *ListDriveSpacesHeadersAccountContext {
	return s.AccountContext
}

func (s *ListDriveSpacesHeaders) SetCommonHeaders(v map[string]*string) *ListDriveSpacesHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListDriveSpacesHeaders) SetAccountContext(v *ListDriveSpacesHeadersAccountContext) *ListDriveSpacesHeaders {
	s.AccountContext = v
	return s
}

func (s *ListDriveSpacesHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDriveSpacesHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListDriveSpacesHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListDriveSpacesHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListDriveSpacesHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListDriveSpacesHeadersAccountContext) SetAccountId(v string) *ListDriveSpacesHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListDriveSpacesHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
