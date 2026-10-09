// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmgVulSubmitRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyEmgVulSubmitRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyEmgVulSubmitRequest
	GetDryRun() *bool
	SetLang(v string) *ModifyEmgVulSubmitRequest
	GetLang() *string
	SetName(v string) *ModifyEmgVulSubmitRequest
	GetName() *string
	SetResourceDirectoryAccountId(v int64) *ModifyEmgVulSubmitRequest
	GetResourceDirectoryAccountId() *int64
	SetUserAgreement(v string) *ModifyEmgVulSubmitRequest
	GetUserAgreement() *string
}

type ModifyEmgVulSubmitRequest struct {
	// The client token used to ensure the idempotence of the request. Use a different token for different requests. Only ASCII characters are supported. The token can be up to 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-0016e04115b
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run for this request. Valid values: true: performs only a dry run without executing the actual operation. false: executes the request normally. Default value: false.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The language of the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the vulnerability to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// scan:ASCV-2019-032401
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the member accounts in the resource directory (Alibaba Cloud account).
	//
	// >Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 16670360956*****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// Specifies whether to run vulnerability detection. Valid values:
	//
	// - **yes**: Run.
	//
	// - **no**: Do not run.
	//
	// This parameter is required.
	//
	// example:
	//
	// yes
	UserAgreement *string `json:"UserAgreement,omitempty" xml:"UserAgreement,omitempty"`
}

func (s ModifyEmgVulSubmitRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmgVulSubmitRequest) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyEmgVulSubmitRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyEmgVulSubmitRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifyEmgVulSubmitRequest) GetName() *string {
	return s.Name
}

func (s *ModifyEmgVulSubmitRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyEmgVulSubmitRequest) GetUserAgreement() *string {
	return s.UserAgreement
}

func (s *ModifyEmgVulSubmitRequest) SetClientToken(v string) *ModifyEmgVulSubmitRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetDryRun(v bool) *ModifyEmgVulSubmitRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetLang(v string) *ModifyEmgVulSubmitRequest {
	s.Lang = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetName(v string) *ModifyEmgVulSubmitRequest {
	s.Name = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetResourceDirectoryAccountId(v int64) *ModifyEmgVulSubmitRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) SetUserAgreement(v string) *ModifyEmgVulSubmitRequest {
	s.UserAgreement = &v
	return s
}

func (s *ModifyEmgVulSubmitRequest) Validate() error {
	return dara.Validate(s)
}
