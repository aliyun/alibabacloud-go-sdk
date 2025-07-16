// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDriveSpaceHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DeleteDriveSpaceHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DeleteDriveSpaceHeadersAccountContext) *DeleteDriveSpaceHeaders
	GetAccountContext() *DeleteDriveSpaceHeadersAccountContext
}

type DeleteDriveSpaceHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	// This parameter is required.
	AccountContext *DeleteDriveSpaceHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DeleteDriveSpaceHeaders) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceHeaders) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DeleteDriveSpaceHeaders) GetAccountContext() *DeleteDriveSpaceHeadersAccountContext {
	return s.AccountContext
}

func (s *DeleteDriveSpaceHeaders) SetCommonHeaders(v map[string]*string) *DeleteDriveSpaceHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DeleteDriveSpaceHeaders) SetAccountContext(v *DeleteDriveSpaceHeadersAccountContext) *DeleteDriveSpaceHeaders {
	s.AccountContext = v
	return s
}

func (s *DeleteDriveSpaceHeaders) Validate() error {
	return dara.Validate(s)
}

type DeleteDriveSpaceHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DeleteDriveSpaceHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DeleteDriveSpaceHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DeleteDriveSpaceHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteDriveSpaceHeadersAccountContext) SetAccountId(v string) *DeleteDriveSpaceHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DeleteDriveSpaceHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
