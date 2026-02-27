// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataServiceAppSecretRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *ResetDataServiceAppSecretRequest
	GetOpTenantId() *int64
	SetUpdateCommand(v *ResetDataServiceAppSecretRequestUpdateCommand) *ResetDataServiceAppSecretRequest
	GetUpdateCommand() *ResetDataServiceAppSecretRequestUpdateCommand
}

type ResetDataServiceAppSecretRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	UpdateCommand *ResetDataServiceAppSecretRequestUpdateCommand `json:"UpdateCommand,omitempty" xml:"UpdateCommand,omitempty" type:"Struct"`
}

func (s ResetDataServiceAppSecretRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretRequest) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ResetDataServiceAppSecretRequest) GetUpdateCommand() *ResetDataServiceAppSecretRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *ResetDataServiceAppSecretRequest) SetOpTenantId(v int64) *ResetDataServiceAppSecretRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResetDataServiceAppSecretRequest) SetUpdateCommand(v *ResetDataServiceAppSecretRequestUpdateCommand) *ResetDataServiceAppSecretRequest {
	s.UpdateCommand = v
	return s
}

func (s *ResetDataServiceAppSecretRequest) Validate() error {
	if s.UpdateCommand != nil {
		if err := s.UpdateCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetDataServiceAppSecretRequestUpdateCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 200000001
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// abc123456789
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
}

func (s ResetDataServiceAppSecretRequestUpdateCommand) String() string {
	return dara.Prettify(s)
}

func (s ResetDataServiceAppSecretRequestUpdateCommand) GoString() string {
	return s.String()
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) GetAppKey() *string {
	return s.AppKey
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) GetAppSecret() *string {
	return s.AppSecret
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) SetAppId(v int32) *ResetDataServiceAppSecretRequestUpdateCommand {
	s.AppId = &v
	return s
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) SetAppKey(v string) *ResetDataServiceAppSecretRequestUpdateCommand {
	s.AppKey = &v
	return s
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) SetAppSecret(v string) *ResetDataServiceAppSecretRequestUpdateCommand {
	s.AppSecret = &v
	return s
}

func (s *ResetDataServiceAppSecretRequestUpdateCommand) Validate() error {
	return dara.Validate(s)
}
