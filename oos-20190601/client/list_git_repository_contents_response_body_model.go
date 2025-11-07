// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGitRepositoryContentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContents(v []*ListGitRepositoryContentsResponseBodyContents) *ListGitRepositoryContentsResponseBody
	GetContents() []*ListGitRepositoryContentsResponseBodyContents
	SetCount(v int32) *ListGitRepositoryContentsResponseBody
	GetCount() *int32
	SetRequestId(v string) *ListGitRepositoryContentsResponseBody
	GetRequestId() *string
}

type ListGitRepositoryContentsResponseBody struct {
	Contents []*ListGitRepositoryContentsResponseBodyContents `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 2075899A-585D-4A41-A9B2-28DA8534
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListGitRepositoryContentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoryContentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGitRepositoryContentsResponseBody) GetContents() []*ListGitRepositoryContentsResponseBodyContents {
	return s.Contents
}

func (s *ListGitRepositoryContentsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *ListGitRepositoryContentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGitRepositoryContentsResponseBody) SetContents(v []*ListGitRepositoryContentsResponseBodyContents) *ListGitRepositoryContentsResponseBody {
	s.Contents = v
	return s
}

func (s *ListGitRepositoryContentsResponseBody) SetCount(v int32) *ListGitRepositoryContentsResponseBody {
	s.Count = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBody) SetRequestId(v string) *ListGitRepositoryContentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBody) Validate() error {
	if s.Contents != nil {
		for _, item := range s.Contents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGitRepositoryContentsResponseBodyContents struct {
	// example:
	//
	// mydir
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// mydir
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// dir1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// dir
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListGitRepositoryContentsResponseBodyContents) String() string {
	return dara.Prettify(s)
}

func (s ListGitRepositoryContentsResponseBodyContents) GoString() string {
	return s.String()
}

func (s *ListGitRepositoryContentsResponseBodyContents) GetContent() *string {
	return s.Content
}

func (s *ListGitRepositoryContentsResponseBodyContents) GetName() *string {
	return s.Name
}

func (s *ListGitRepositoryContentsResponseBodyContents) GetPath() *string {
	return s.Path
}

func (s *ListGitRepositoryContentsResponseBodyContents) GetType() *string {
	return s.Type
}

func (s *ListGitRepositoryContentsResponseBodyContents) SetContent(v string) *ListGitRepositoryContentsResponseBodyContents {
	s.Content = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBodyContents) SetName(v string) *ListGitRepositoryContentsResponseBodyContents {
	s.Name = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBodyContents) SetPath(v string) *ListGitRepositoryContentsResponseBodyContents {
	s.Path = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBodyContents) SetType(v string) *ListGitRepositoryContentsResponseBodyContents {
	s.Type = &v
	return s
}

func (s *ListGitRepositoryContentsResponseBodyContents) Validate() error {
	return dara.Validate(s)
}
