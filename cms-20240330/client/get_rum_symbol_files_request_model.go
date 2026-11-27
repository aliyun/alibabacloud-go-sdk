// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArtifactType(v string) *GetRumSymbolFilesRequest
	GetArtifactType() *string
	SetKeywords(v []*string) *GetRumSymbolFilesRequest
	GetKeywords() []*string
	SetNextToken(v string) *GetRumSymbolFilesRequest
	GetNextToken() *string
	SetPageSize(v int64) *GetRumSymbolFilesRequest
	GetPageSize() *int64
	SetPlatform(v string) *GetRumSymbolFilesRequest
	GetPlatform() *string
	SetServiceId(v string) *GetRumSymbolFilesRequest
	GetServiceId() *string
	SetVersion(v string) *GetRumSymbolFilesRequest
	GetVersion() *string
	SetWorkspace(v string) *GetRumSymbolFilesRequest
	GetWorkspace() *string
}

type GetRumSymbolFilesRequest struct {
	// The filter for the symbol file type. Valid values: map, dsym, txt, so, sym, pdb, and json. This parameter is used for query filtering and differs from the fileKind dimension in the upload API GetRumSymbolFileParams: uploading dsym (.dSYM.zip) corresponds to artifactType=dsym; standard symbol-zip (.zip) is not currently supported for filtering by artifactType.
	//
	// example:
	//
	// txt
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The list of search keywords. You can specify up to 50 keywords to match file names, UUIDs, version numbers, and more.
	//
	// example:
	//
	// ["new-app"]
	Keywords []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	// The paging token. To query the next page, specify the nextToken value returned in the previous response. The example is for illustration purposes only and cannot be used for actual queries.
	//
	// example:
	//
	// next
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The filter for the platform type. Valid values: web, miniapp, minigame, uniapp, android, ios, harmonyos, windows, macos, and linux.
	//
	// example:
	//
	// android
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The service ID of the RUM application. This parameter is required, cannot be an empty string, and must be used together with workspace to locate the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4e5@f67890
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The filter for the application version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The name of the workspace to which the RUM application belongs. This parameter is required, cannot be an empty string, and must be used together with serviceId to locate the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-1234567890-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetRumSymbolFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFilesRequest) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFilesRequest) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetRumSymbolFilesRequest) GetKeywords() []*string {
	return s.Keywords
}

func (s *GetRumSymbolFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetRumSymbolFilesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRumSymbolFilesRequest) GetPlatform() *string {
	return s.Platform
}

func (s *GetRumSymbolFilesRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetRumSymbolFilesRequest) GetVersion() *string {
	return s.Version
}

func (s *GetRumSymbolFilesRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetRumSymbolFilesRequest) SetArtifactType(v string) *GetRumSymbolFilesRequest {
	s.ArtifactType = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetKeywords(v []*string) *GetRumSymbolFilesRequest {
	s.Keywords = v
	return s
}

func (s *GetRumSymbolFilesRequest) SetNextToken(v string) *GetRumSymbolFilesRequest {
	s.NextToken = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetPageSize(v int64) *GetRumSymbolFilesRequest {
	s.PageSize = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetPlatform(v string) *GetRumSymbolFilesRequest {
	s.Platform = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetServiceId(v string) *GetRumSymbolFilesRequest {
	s.ServiceId = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetVersion(v string) *GetRumSymbolFilesRequest {
	s.Version = &v
	return s
}

func (s *GetRumSymbolFilesRequest) SetWorkspace(v string) *GetRumSymbolFilesRequest {
	s.Workspace = &v
	return s
}

func (s *GetRumSymbolFilesRequest) Validate() error {
	return dara.Validate(s)
}
