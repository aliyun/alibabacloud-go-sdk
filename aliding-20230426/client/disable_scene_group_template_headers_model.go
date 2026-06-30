// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSceneGroupTemplateHeaders interface {
	dara.Model
	String() string
	GoString() string
	SetCommonHeaders(v map[string]*string) *DisableSceneGroupTemplateHeaders
	GetCommonHeaders() map[string]*string
	SetAccountContext(v *DisableSceneGroupTemplateHeadersAccountContext) *DisableSceneGroupTemplateHeaders
	GetAccountContext() *DisableSceneGroupTemplateHeadersAccountContext
}

type DisableSceneGroupTemplateHeaders struct {
	CommonHeaders  map[string]*string                              `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccountContext *DisableSceneGroupTemplateHeadersAccountContext `json:"AccountContext,omitempty" xml:"AccountContext,omitempty" type:"Struct"`
}

func (s DisableSceneGroupTemplateHeaders) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateHeaders) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateHeaders) GetCommonHeaders() map[string]*string {
	return s.CommonHeaders
}

func (s *DisableSceneGroupTemplateHeaders) GetAccountContext() *DisableSceneGroupTemplateHeadersAccountContext {
	return s.AccountContext
}

func (s *DisableSceneGroupTemplateHeaders) SetCommonHeaders(v map[string]*string) *DisableSceneGroupTemplateHeaders {
	s.CommonHeaders = v
	return s
}

func (s *DisableSceneGroupTemplateHeaders) SetAccountContext(v *DisableSceneGroupTemplateHeadersAccountContext) *DisableSceneGroupTemplateHeaders {
	s.AccountContext = v
	return s
}

func (s *DisableSceneGroupTemplateHeaders) Validate() error {
	if s.AccountContext != nil {
		if err := s.AccountContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DisableSceneGroupTemplateHeadersAccountContext struct {
	// This parameter is required.
	//
	// example:
	//
	// 012345
	AccountId *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
}

func (s DisableSceneGroupTemplateHeadersAccountContext) String() string {
	return dara.Prettify(s)
}

func (s DisableSceneGroupTemplateHeadersAccountContext) GoString() string {
	return s.String()
}

func (s *DisableSceneGroupTemplateHeadersAccountContext) GetAccountId() *string {
	return s.AccountId
}

func (s *DisableSceneGroupTemplateHeadersAccountContext) SetAccountId(v string) *DisableSceneGroupTemplateHeadersAccountContext {
	s.AccountId = &v
	return s
}

func (s *DisableSceneGroupTemplateHeadersAccountContext) Validate() error {
	return dara.Validate(s)
}
