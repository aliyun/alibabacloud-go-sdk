// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAnnotationMissionTagInfoListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetAnnotationMissionTagInfoListRequest
	GetInstanceId() *string
	SetPageIndex(v int32) *GetAnnotationMissionTagInfoListRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetAnnotationMissionTagInfoListRequest
	GetPageSize() *int32
}

type GetAnnotationMissionTagInfoListRequest struct {
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s GetAnnotationMissionTagInfoListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAnnotationMissionTagInfoListRequest) GoString() string {
	return s.String()
}

func (s *GetAnnotationMissionTagInfoListRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAnnotationMissionTagInfoListRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetAnnotationMissionTagInfoListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetAnnotationMissionTagInfoListRequest) SetInstanceId(v string) *GetAnnotationMissionTagInfoListRequest {
	s.InstanceId = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListRequest) SetPageIndex(v int32) *GetAnnotationMissionTagInfoListRequest {
	s.PageIndex = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListRequest) SetPageSize(v int32) *GetAnnotationMissionTagInfoListRequest {
	s.PageSize = &v
	return s
}

func (s *GetAnnotationMissionTagInfoListRequest) Validate() error {
	return dara.Validate(s)
}
