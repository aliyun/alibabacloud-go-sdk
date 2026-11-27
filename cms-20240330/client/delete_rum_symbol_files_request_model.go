// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRumSymbolFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePaths(v []*string) *DeleteRumSymbolFilesRequest
	GetFilePaths() []*string
	SetServiceId(v string) *DeleteRumSymbolFilesRequest
	GetServiceId() *string
	SetWorkspace(v string) *DeleteRumSymbolFilesRequest
	GetWorkspace() *string
}

type DeleteRumSymbolFilesRequest struct {
	// The list of object paths of the symbol files to delete. This parameter is required and cannot be an empty array. Set the value to the artifactPath returned by GetRumSymbolFiles. Up to 1,000 paths are supported per call. Only user-uploaded symbol files under the current application can be deleted. System symbol files cannot be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["user/workspace/service-id/3.0.0/build-1/new-app.map"]
	FilePaths []*string `json:"filePaths,omitempty" xml:"filePaths,omitempty" type:"Repeated"`
	// The ID of the RUM application service.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4e5@f67890
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-1234567890-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s DeleteRumSymbolFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRumSymbolFilesRequest) GoString() string {
	return s.String()
}

func (s *DeleteRumSymbolFilesRequest) GetFilePaths() []*string {
	return s.FilePaths
}

func (s *DeleteRumSymbolFilesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *DeleteRumSymbolFilesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *DeleteRumSymbolFilesRequest) SetFilePaths(v []*string) *DeleteRumSymbolFilesRequest {
	s.FilePaths = v
	return s
}

func (s *DeleteRumSymbolFilesRequest) SetServiceId(v string) *DeleteRumSymbolFilesRequest {
	s.ServiceId = &v
	return s
}

func (s *DeleteRumSymbolFilesRequest) SetWorkspace(v string) *DeleteRumSymbolFilesRequest {
	s.Workspace = &v
	return s
}

func (s *DeleteRumSymbolFilesRequest) Validate() error {
	return dara.Validate(s)
}
