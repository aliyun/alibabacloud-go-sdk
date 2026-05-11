// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListArtifactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody
	GetArtifacts() []*ListArtifactsResponseBodyArtifacts
	SetMaxResults(v int32) *ListArtifactsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListArtifactsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListArtifactsResponseBody
	GetRequestId() *string
}

type ListArtifactsResponseBody struct {
	Artifacts []*ListArtifactsResponseBodyArtifacts `json:"artifacts,omitempty" xml:"artifacts,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// xxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListArtifactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBody) GetArtifacts() []*ListArtifactsResponseBodyArtifacts {
	return s.Artifacts
}

func (s *ListArtifactsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListArtifactsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListArtifactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListArtifactsResponseBody) SetArtifacts(v []*ListArtifactsResponseBodyArtifacts) *ListArtifactsResponseBody {
	s.Artifacts = v
	return s
}

func (s *ListArtifactsResponseBody) SetMaxResults(v int32) *ListArtifactsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListArtifactsResponseBody) SetNextToken(v string) *ListArtifactsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListArtifactsResponseBody) SetRequestId(v string) *ListArtifactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListArtifactsResponseBody) Validate() error {
	if s.Artifacts != nil {
		for _, item := range s.Artifacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListArtifactsResponseBodyArtifacts struct {
	// example:
	//
	// false
	IsDirectory *bool `json:"isDirectory,omitempty" xml:"isDirectory,omitempty"`
	// example:
	//
	// 2026-04-30T16:03:54Z
	LastModified *string `json:"lastModified,omitempty" xml:"lastModified,omitempty"`
	// example:
	//
	// missions/mission-xxx/artifacts/2026-05/05-01/xxxx.md
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// example:
	//
	// 21950
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListArtifactsResponseBodyArtifacts) String() string {
	return dara.Prettify(s)
}

func (s ListArtifactsResponseBodyArtifacts) GoString() string {
	return s.String()
}

func (s *ListArtifactsResponseBodyArtifacts) GetIsDirectory() *bool {
	return s.IsDirectory
}

func (s *ListArtifactsResponseBodyArtifacts) GetLastModified() *string {
	return s.LastModified
}

func (s *ListArtifactsResponseBodyArtifacts) GetPath() *string {
	return s.Path
}

func (s *ListArtifactsResponseBodyArtifacts) GetSize() *int64 {
	return s.Size
}

func (s *ListArtifactsResponseBodyArtifacts) SetIsDirectory(v bool) *ListArtifactsResponseBodyArtifacts {
	s.IsDirectory = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetLastModified(v string) *ListArtifactsResponseBodyArtifacts {
	s.LastModified = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetPath(v string) *ListArtifactsResponseBodyArtifacts {
	s.Path = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) SetSize(v int64) *ListArtifactsResponseBodyArtifacts {
	s.Size = &v
	return s
}

func (s *ListArtifactsResponseBodyArtifacts) Validate() error {
	return dara.Validate(s)
}
