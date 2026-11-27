// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFileParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugId(v string) *GetRumSymbolFileParamsRequest
	GetDebugId() *string
	SetFileKind(v string) *GetRumSymbolFileParamsRequest
	GetFileKind() *string
	SetFileName(v string) *GetRumSymbolFileParamsRequest
	GetFileName() *string
	SetServiceId(v string) *GetRumSymbolFileParamsRequest
	GetServiceId() *string
	SetVersion(v string) *GetRumSymbolFileParamsRequest
	GetVersion() *string
	SetWorkspace(v string) *GetRumSymbolFileParamsRequest
	GetWorkspace() *string
}

type GetRumSymbolFileParamsRequest struct {
	// The debug identifier of the symbol file, such as a UUID or Debug ID. This parameter is required for all file types except dsym.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426614174000
	DebugId *string `json:"debugId,omitempty" xml:"debugId,omitempty"`
	// The type of the file to upload. Valid values: sourcemap (.map), dsym (.dSYM.zip), proguard (.txt), native-symbol (.so), name-cache (.json/.txt), pdb (.pdb), sym (.sym), and symbol-zip (.zip). This parameter is used to apply for upload signature parameters. The value dimension differs from the artifactType filter parameter in the GetRumSymbolFiles query operation. Specifically, dsym (.dSYM.zip) is classified as artifactType=dsym during queries. symbol-zip (.zip) files are directly uploaded and archived with their original file names, and the current query operation does not support filtering by zip for artifactType.
	//
	// This parameter is required.
	//
	// example:
	//
	// proguard
	FileKind *string `json:"fileKind,omitempty" xml:"fileKind,omitempty"`
	// The name of the symbol file to upload. The file extension must match the fileKind.
	//
	// This parameter is required.
	//
	// example:
	//
	// mapping.txt
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The ID of the RUM application service.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4e5@f67890
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The version number of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-1234567890-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetRumSymbolFileParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFileParamsRequest) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFileParamsRequest) GetDebugId() *string {
	return s.DebugId
}

func (s *GetRumSymbolFileParamsRequest) GetFileKind() *string {
	return s.FileKind
}

func (s *GetRumSymbolFileParamsRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetRumSymbolFileParamsRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetRumSymbolFileParamsRequest) GetVersion() *string {
	return s.Version
}

func (s *GetRumSymbolFileParamsRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetRumSymbolFileParamsRequest) SetDebugId(v string) *GetRumSymbolFileParamsRequest {
	s.DebugId = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) SetFileKind(v string) *GetRumSymbolFileParamsRequest {
	s.FileKind = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) SetFileName(v string) *GetRumSymbolFileParamsRequest {
	s.FileName = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) SetServiceId(v string) *GetRumSymbolFileParamsRequest {
	s.ServiceId = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) SetVersion(v string) *GetRumSymbolFileParamsRequest {
	s.Version = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) SetWorkspace(v string) *GetRumSymbolFileParamsRequest {
	s.Workspace = &v
	return s
}

func (s *GetRumSymbolFileParamsRequest) Validate() error {
	return dara.Validate(s)
}
