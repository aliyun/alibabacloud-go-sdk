// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPermissionsHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *ListPermissionsHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *ListPermissionsHeadersAccountContext) *ListPermissionsHeaders
	GetAccountContext() *ListPermissionsHeadersAccountContext
}

type ListPermissionsHeaders struct {
	CommonHeaders  map[string]*string                    `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *ListPermissionsHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s ListPermissionsHeaders) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsHeaders) GoString() string {
	return s.String()
}

func (s *ListPermissionsHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *ListPermissionsHeaders) GetAccountContext() *ListPermissionsHeadersAccountContext {
	return s.AccountContext
}

func (s *ListPermissionsHeaders) SetCommonHeaders(v map[string]*string) *ListPermissionsHeaders {
	s.CommonHeaders = v
	return s
}

func (s *ListPermissionsHeaders) SetAccountContext(v *ListPermissionsHeadersAccountContext) *ListPermissionsHeaders {
	s.AccountContext = v
	return s
}

func (s *ListPermissionsHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPermissionsHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s ListPermissionsHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s ListPermissionsHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *ListPermissionsHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *ListPermissionsHeadersAccountContext) SetAccountId(v string) *ListPermissionsHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *ListPermissionsHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
