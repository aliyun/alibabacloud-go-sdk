// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *RevokeDataServiceApiRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *RevokeDataServiceApiRequest
	GetProjectId() *int32
	SetRevokeCommand(v *RevokeDataServiceApiRequestRevokeCommand) *RevokeDataServiceApiRequest
	GetRevokeCommand() *RevokeDataServiceApiRequestRevokeCommand
}

type RevokeDataServiceApiRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	RevokeCommand *RevokeDataServiceApiRequestRevokeCommand `json:"RevokeCommand,omitempty" xml:"RevokeCommand,omitempty" type:"Struct"`
}

func (s RevokeDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *RevokeDataServiceApiRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *RevokeDataServiceApiRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *RevokeDataServiceApiRequest) GetRevokeCommand() *RevokeDataServiceApiRequestRevokeCommand {
	return s.RevokeCommand
}

func (s *RevokeDataServiceApiRequest) SetOpTenantId(v int64) *RevokeDataServiceApiRequest {
	s.OpTenantId = &v
	return s
}

func (s *RevokeDataServiceApiRequest) SetProjectId(v int32) *RevokeDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *RevokeDataServiceApiRequest) SetRevokeCommand(v *RevokeDataServiceApiRequestRevokeCommand) *RevokeDataServiceApiRequest {
	s.RevokeCommand = v
	return s
}

func (s *RevokeDataServiceApiRequest) Validate() error {
	if s.RevokeCommand != nil {
		if err := s.RevokeCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RevokeDataServiceApiRequestRevokeCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1021
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// example:
	//
	// 1203
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// USE
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// APP
	GranteeType *string `json:"GranteeType,omitempty" xml:"GranteeType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s RevokeDataServiceApiRequestRevokeCommand) String() string {
	return dara.Prettify(s)
}

func (s RevokeDataServiceApiRequestRevokeCommand) GoString() string {
	return s.String()
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetApiId() *int64 {
	return s.ApiId
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetAuthType() *string {
	return s.AuthType
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetEnv() *string {
	return s.Env
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetGranteeType() *string {
	return s.GranteeType
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetReason() *string {
	return s.Reason
}

func (s *RevokeDataServiceApiRequestRevokeCommand) GetUserId() *string {
	return s.UserId
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetApiId(v int64) *RevokeDataServiceApiRequestRevokeCommand {
	s.ApiId = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetAppId(v int32) *RevokeDataServiceApiRequestRevokeCommand {
	s.AppId = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetAuthType(v string) *RevokeDataServiceApiRequestRevokeCommand {
	s.AuthType = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetEnv(v string) *RevokeDataServiceApiRequestRevokeCommand {
	s.Env = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetGranteeType(v string) *RevokeDataServiceApiRequestRevokeCommand {
	s.GranteeType = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetReason(v string) *RevokeDataServiceApiRequestRevokeCommand {
	s.Reason = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) SetUserId(v string) *RevokeDataServiceApiRequestRevokeCommand {
	s.UserId = &v
	return s
}

func (s *RevokeDataServiceApiRequestRevokeCommand) Validate() error {
	return dara.Validate(s)
}
