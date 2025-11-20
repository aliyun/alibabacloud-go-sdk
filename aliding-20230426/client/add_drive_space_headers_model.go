// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDriveSpaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *AddDriveSpaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *AddDriveSpaceHeadersAccountContext) *AddDriveSpaceHeaders
	GetAccountContext() *AddDriveSpaceHeadersAccountContext
}

type AddDriveSpaceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContext *AddDriveSpaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s AddDriveSpaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceHeaders) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *AddDriveSpaceHeaders) GetAccountContext() *AddDriveSpaceHeadersAccountContext {
	return s.AccountContext
}

func (s *AddDriveSpaceHeaders) SetCommonHeaders(v map[string]*string) *AddDriveSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AddDriveSpaceHeaders) SetAccountContext(v *AddDriveSpaceHeadersAccountContext) *AddDriveSpaceHeaders {
	s.AccountContext = v
	return s
}

func (s *AddDriveSpaceHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDriveSpaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s AddDriveSpaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s AddDriveSpaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *AddDriveSpaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *AddDriveSpaceHeadersAccountContext) SetAccountId(v string) *AddDriveSpaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *AddDriveSpaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
