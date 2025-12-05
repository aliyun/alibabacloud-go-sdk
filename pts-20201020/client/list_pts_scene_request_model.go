// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPtsSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyWord(v string) *ListPtsSceneRequest
	GetKeyWord() *string
	SetPageNumber(v int32) *ListPtsSceneRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPtsSceneRequest
	GetPageSize() *int32
}

type ListPtsSceneRequest struct {
	// The keyword based on which you can search for the PTS scenario. You can perform a fuzzy search on the scenario name (**SceneName**) or an exact search on the scenario ID (**SceneId**).
	KeyWord *string `json:"KeyWord,omitempty" xml:"KeyWord,omitempty"`
	// The number of the page to return. Valid values: 1 to 1073741824.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of scenario entries to return per page. Valid values: 10 to 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListPtsSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPtsSceneRequest) GoString() string {
	return s.String()
}

func (s *ListPtsSceneRequest) GetKeyWord() *string {
	return s.KeyWord
}

func (s *ListPtsSceneRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPtsSceneRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPtsSceneRequest) SetKeyWord(v string) *ListPtsSceneRequest {
	s.KeyWord = &v
	return s
}

func (s *ListPtsSceneRequest) SetPageNumber(v int32) *ListPtsSceneRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPtsSceneRequest) SetPageSize(v int32) *ListPtsSceneRequest {
	s.PageSize = &v
	return s
}

func (s *ListPtsSceneRequest) Validate() error {
	return dara.Validate(s)
}
