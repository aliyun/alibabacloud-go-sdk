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
	SetOpUserId(v string) *ResetDataServiceAppSecretRequest
	GetOpUserId() *string
	SetUpdateCommand(v *ResetDataServiceAppSecretRequestUpdateCommand) *ResetDataServiceAppSecretRequest
	GetUpdateCommand() *ResetDataServiceAppSecretRequestUpdateCommand
}

type ResetDataServiceAppSecretRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The update command to reset the application key of a data service.
	//
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

func (s *ResetDataServiceAppSecretRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *ResetDataServiceAppSecretRequest) GetUpdateCommand() *ResetDataServiceAppSecretRequestUpdateCommand {
	return s.UpdateCommand
}

func (s *ResetDataServiceAppSecretRequest) SetOpTenantId(v int64) *ResetDataServiceAppSecretRequest {
	s.OpTenantId = &v
	return s
}

func (s *ResetDataServiceAppSecretRequest) SetOpUserId(v string) *ResetDataServiceAppSecretRequest {
	s.OpUserId = &v
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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The new application key. The key must be globally unique and is used when you call API operations. The key must be 8 to 128 characters in length, and can contain digits, underscores (_), and hyphens (-). You can customize this parameter only when you use Alibaba Cloud API Gateway or the built-in gateway. If you use Apsara Stack Gateway, this parameter is ignored.
	//
	// example:
	//
	// 200000001
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The new application secret. If you leave this parameter empty, the system automatically generates a new AppSecret value. The secret must be 8 to 127 characters in length, and can contain digits, underscores (_), and hyphens (-). You can customize this parameter only when you use Alibaba Cloud API Gateway or the built-in gateway. If you use Apsara Stack Gateway, this parameter is ignored.
	//
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
