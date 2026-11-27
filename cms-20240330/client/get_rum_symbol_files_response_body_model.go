// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*GetRumSymbolFilesResponseBodyFiles) *GetRumSymbolFilesResponseBody
	GetFiles() []*GetRumSymbolFilesResponseBodyFiles
	SetHasMore(v bool) *GetRumSymbolFilesResponseBody
	GetHasMore() *bool
	SetNextToken(v string) *GetRumSymbolFilesResponseBody
	GetNextToken() *string
	SetPageSize(v int64) *GetRumSymbolFilesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *GetRumSymbolFilesResponseBody
	GetRequestId() *string
}

type GetRumSymbolFilesResponseBody struct {
	// The list of symbol files.
	Files []*GetRumSymbolFilesResponseBodyFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// Indicates whether more data is available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The pagination token for the next page. The example is for illustration purposes only and cannot be used for actual queries. To query the next page, use the actual value returned in the response.
	//
	// example:
	//
	// next
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// request-1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetRumSymbolFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFilesResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFilesResponseBody) GetFiles() []*GetRumSymbolFilesResponseBodyFiles {
	return s.Files
}

func (s *GetRumSymbolFilesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *GetRumSymbolFilesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetRumSymbolFilesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetRumSymbolFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumSymbolFilesResponseBody) SetFiles(v []*GetRumSymbolFilesResponseBodyFiles) *GetRumSymbolFilesResponseBody {
	s.Files = v
	return s
}

func (s *GetRumSymbolFilesResponseBody) SetHasMore(v bool) *GetRumSymbolFilesResponseBody {
	s.HasMore = &v
	return s
}

func (s *GetRumSymbolFilesResponseBody) SetNextToken(v string) *GetRumSymbolFilesResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetRumSymbolFilesResponseBody) SetPageSize(v int64) *GetRumSymbolFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetRumSymbolFilesResponseBody) SetRequestId(v string) *GetRumSymbolFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumSymbolFilesResponseBody) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRumSymbolFilesResponseBodyFiles struct {
	// The object path of the symbol file.
	//
	// example:
	//
	// user/workspace/service-id/3.0.0/build-1/new-app.map
	ArtifactPath *string `json:"artifactPath,omitempty" xml:"artifactPath,omitempty"`
	// The type of the symbol file.
	//
	// example:
	//
	// map
	ArtifactType *string `json:"artifactType,omitempty" xml:"artifactType,omitempty"`
	// The name of the symbol file.
	//
	// example:
	//
	// new-app.map
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The list of matched search keywords.
	//
	// example:
	//
	// ["new-app"]
	MatchedKeywords []*string `json:"matchedKeywords,omitempty" xml:"matchedKeywords,omitempty" type:"Repeated"`
	// The platform type to which the file belongs.
	//
	// example:
	//
	// web
	Platform *string `json:"platform,omitempty" xml:"platform,omitempty"`
	// The size of the file in bytes.
	//
	// example:
	//
	// 512
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The source of the symbol file. This query API returns only symbol files uploaded by users, which is fixed to user. The value system is reserved for internal system symbols of the platform.
	//
	// example:
	//
	// user
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The timestamp of the last time the file was updated, in milliseconds.
	//
	// example:
	//
	// 1778806923000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The UUID or Debug ID of the symbol file.
	//
	// example:
	//
	// build-1
	Uuid *string `json:"uuid,omitempty" xml:"uuid,omitempty"`
	// The associated application version number.
	//
	// example:
	//
	// 3.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetRumSymbolFilesResponseBodyFiles) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolFilesResponseBodyFiles) GoString() string {
	return s.String()
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetArtifactPath() *string {
	return s.ArtifactPath
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetArtifactType() *string {
	return s.ArtifactType
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetFileName() *string {
	return s.FileName
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetMatchedKeywords() []*string {
	return s.MatchedKeywords
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetPlatform() *string {
	return s.Platform
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetSize() *int64 {
	return s.Size
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetSource() *string {
	return s.Source
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetUuid() *string {
	return s.Uuid
}

func (s *GetRumSymbolFilesResponseBodyFiles) GetVersion() *string {
	return s.Version
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetArtifactPath(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.ArtifactPath = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetArtifactType(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.ArtifactType = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetFileName(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.FileName = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetMatchedKeywords(v []*string) *GetRumSymbolFilesResponseBodyFiles {
	s.MatchedKeywords = v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetPlatform(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.Platform = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetSize(v int64) *GetRumSymbolFilesResponseBodyFiles {
	s.Size = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetSource(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.Source = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetUpdateTime(v int64) *GetRumSymbolFilesResponseBodyFiles {
	s.UpdateTime = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetUuid(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.Uuid = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) SetVersion(v string) *GetRumSymbolFilesResponseBodyFiles {
	s.Version = &v
	return s
}

func (s *GetRumSymbolFilesResponseBodyFiles) Validate() error {
	return dara.Validate(s)
}
