// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDataSourceFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *UploadDataSourceFileRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *UploadDataSourceFileRequest
	GetOpUserId() *string
	SetUploadCommand(v *UploadDataSourceFileRequestUploadCommand) *UploadDataSourceFileRequest
	GetUploadCommand() *UploadDataSourceFileRequestUploadCommand
}

type UploadDataSourceFileRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The request object for uploading a datasource authentication file.
	//
	// This parameter is required.
	UploadCommand *UploadDataSourceFileRequestUploadCommand `json:"UploadCommand,omitempty" xml:"UploadCommand,omitempty" type:"Struct"`
}

func (s UploadDataSourceFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSourceFileRequest) GoString() string {
	return s.String()
}

func (s *UploadDataSourceFileRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *UploadDataSourceFileRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *UploadDataSourceFileRequest) GetUploadCommand() *UploadDataSourceFileRequestUploadCommand {
	return s.UploadCommand
}

func (s *UploadDataSourceFileRequest) SetOpTenantId(v int64) *UploadDataSourceFileRequest {
	s.OpTenantId = &v
	return s
}

func (s *UploadDataSourceFileRequest) SetOpUserId(v string) *UploadDataSourceFileRequest {
	s.OpUserId = &v
	return s
}

func (s *UploadDataSourceFileRequest) SetUploadCommand(v *UploadDataSourceFileRequestUploadCommand) *UploadDataSourceFileRequest {
	s.UploadCommand = v
	return s
}

func (s *UploadDataSourceFileRequest) Validate() error {
	if s.UploadCommand != nil {
		if err := s.UploadCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadDataSourceFileRequestUploadCommand struct {
	// The Base64-encoded file content. The decoded file size must be between 0 and 5 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2V5dGFiIGNvbnRlbnQ=
	FileContentBase64 *string `json:"FileContentBase64,omitempty" xml:"FileContentBase64,omitempty"`
	// The file name, including the extension. The extension is validated against a whitelist. Supported extensions: jar, xml, conf, keytab, jks, rsa, pem, yaml, keystore, properties, and key.
	//
	// This parameter is required.
	//
	// example:
	//
	// user.keytab
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s UploadDataSourceFileRequestUploadCommand) String() string {
	return dara.Prettify(s)
}

func (s UploadDataSourceFileRequestUploadCommand) GoString() string {
	return s.String()
}

func (s *UploadDataSourceFileRequestUploadCommand) GetFileContentBase64() *string {
	return s.FileContentBase64
}

func (s *UploadDataSourceFileRequestUploadCommand) GetFileName() *string {
	return s.FileName
}

func (s *UploadDataSourceFileRequestUploadCommand) SetFileContentBase64(v string) *UploadDataSourceFileRequestUploadCommand {
	s.FileContentBase64 = &v
	return s
}

func (s *UploadDataSourceFileRequestUploadCommand) SetFileName(v string) *UploadDataSourceFileRequestUploadCommand {
	s.FileName = &v
	return s
}

func (s *UploadDataSourceFileRequestUploadCommand) Validate() error {
	return dara.Validate(s)
}
