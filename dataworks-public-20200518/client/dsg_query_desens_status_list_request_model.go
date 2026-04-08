// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryDesensStatusListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *DsgQueryDesensStatusListRequest
	GetKeyword() *string
	SetPageNumber(v int32) *DsgQueryDesensStatusListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgQueryDesensStatusListRequest
	GetPageSize() *int32
	SetSceneCode(v string) *DsgQueryDesensStatusListRequest
	GetSceneCode() *string
	SetSceneId(v string) *DsgQueryDesensStatusListRequest
	GetSceneId() *string
}

type DsgQueryDesensStatusListRequest struct {
	// example:
	//
	// my
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dataworks_display_desense_code
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
	// example:
	//
	// 124
	SceneId *string `json:"sceneId,omitempty" xml:"sceneId,omitempty"`
}

func (s DsgQueryDesensStatusListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryDesensStatusListRequest) GoString() string {
	return s.String()
}

func (s *DsgQueryDesensStatusListRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DsgQueryDesensStatusListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgQueryDesensStatusListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgQueryDesensStatusListRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *DsgQueryDesensStatusListRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *DsgQueryDesensStatusListRequest) SetKeyword(v string) *DsgQueryDesensStatusListRequest {
	s.Keyword = &v
	return s
}

func (s *DsgQueryDesensStatusListRequest) SetPageNumber(v int32) *DsgQueryDesensStatusListRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgQueryDesensStatusListRequest) SetPageSize(v int32) *DsgQueryDesensStatusListRequest {
	s.PageSize = &v
	return s
}

func (s *DsgQueryDesensStatusListRequest) SetSceneCode(v string) *DsgQueryDesensStatusListRequest {
	s.SceneCode = &v
	return s
}

func (s *DsgQueryDesensStatusListRequest) SetSceneId(v string) *DsgQueryDesensStatusListRequest {
	s.SceneId = &v
	return s
}

func (s *DsgQueryDesensStatusListRequest) Validate() error {
	return dara.Validate(s)
}
