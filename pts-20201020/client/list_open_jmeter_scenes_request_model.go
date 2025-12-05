// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpenJMeterScenesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListOpenJMeterScenesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListOpenJMeterScenesRequest
	GetPageSize() *int32
	SetSceneId(v string) *ListOpenJMeterScenesRequest
	GetSceneId() *string
	SetSceneName(v string) *ListOpenJMeterScenesRequest
	GetSceneName() *string
}

type ListOpenJMeterScenesRequest struct {
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of scenarios to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scenario ID.
	//
	// example:
	//
	// DYYPZIH
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scenario name.
	//
	// example:
	//
	// test
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s ListOpenJMeterScenesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpenJMeterScenesRequest) GoString() string {
	return s.String()
}

func (s *ListOpenJMeterScenesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListOpenJMeterScenesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListOpenJMeterScenesRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *ListOpenJMeterScenesRequest) GetSceneName() *string {
	return s.SceneName
}

func (s *ListOpenJMeterScenesRequest) SetPageNumber(v int32) *ListOpenJMeterScenesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetPageSize(v int32) *ListOpenJMeterScenesRequest {
	s.PageSize = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetSceneId(v string) *ListOpenJMeterScenesRequest {
	s.SceneId = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) SetSceneName(v string) *ListOpenJMeterScenesRequest {
	s.SceneName = &v
	return s
}

func (s *ListOpenJMeterScenesRequest) Validate() error {
	return dara.Validate(s)
}
