// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExportTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLanguage(v string) *ListExportTasksRequest
	GetLanguage() *string
	SetPageNo(v int32) *ListExportTasksRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListExportTasksRequest
	GetPageSize() *int32
	SetSceneCode(v string) *ListExportTasksRequest
	GetSceneCode() *string
}

type ListExportTasksRequest struct {
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// intlExportUsageDeductHistory
	SceneCode *string `json:"SceneCode,omitempty" xml:"SceneCode,omitempty"`
}

func (s ListExportTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExportTasksRequest) GoString() string {
	return s.String()
}

func (s *ListExportTasksRequest) GetLanguage() *string {
	return s.Language
}

func (s *ListExportTasksRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListExportTasksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExportTasksRequest) GetSceneCode() *string {
	return s.SceneCode
}

func (s *ListExportTasksRequest) SetLanguage(v string) *ListExportTasksRequest {
	s.Language = &v
	return s
}

func (s *ListExportTasksRequest) SetPageNo(v int32) *ListExportTasksRequest {
	s.PageNo = &v
	return s
}

func (s *ListExportTasksRequest) SetPageSize(v int32) *ListExportTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListExportTasksRequest) SetSceneCode(v string) *ListExportTasksRequest {
	s.SceneCode = &v
	return s
}

func (s *ListExportTasksRequest) Validate() error {
	return dara.Validate(s)
}
