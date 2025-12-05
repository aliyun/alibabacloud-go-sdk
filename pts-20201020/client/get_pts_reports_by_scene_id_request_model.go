// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsReportsBySceneIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetPtsReportsBySceneIdRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *GetPtsReportsBySceneIdRequest
	GetPageSize() *int32
	SetSceneId(v string) *GetPtsReportsBySceneIdRequest
	GetSceneId() *string
}

type GetPtsReportsBySceneIdRequest struct {
	// The number of the page to display in the paging operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of reports to display per page. Valid values: 5 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The scenario ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// NGBCD4K
	SceneId *string `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s GetPtsReportsBySceneIdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPtsReportsBySceneIdRequest) GoString() string {
	return s.String()
}

func (s *GetPtsReportsBySceneIdRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetPtsReportsBySceneIdRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetPtsReportsBySceneIdRequest) GetSceneId() *string {
	return s.SceneId
}

func (s *GetPtsReportsBySceneIdRequest) SetPageNumber(v int32) *GetPtsReportsBySceneIdRequest {
	s.PageNumber = &v
	return s
}

func (s *GetPtsReportsBySceneIdRequest) SetPageSize(v int32) *GetPtsReportsBySceneIdRequest {
	s.PageSize = &v
	return s
}

func (s *GetPtsReportsBySceneIdRequest) SetSceneId(v string) *GetPtsReportsBySceneIdRequest {
	s.SceneId = &v
	return s
}

func (s *GetPtsReportsBySceneIdRequest) Validate() error {
	return dara.Validate(s)
}
