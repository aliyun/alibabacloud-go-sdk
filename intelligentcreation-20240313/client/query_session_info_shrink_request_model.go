// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *QuerySessionInfoShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QuerySessionInfoShrinkRequest
	GetPageSize() *int32
	SetProjectId(v string) *QuerySessionInfoShrinkRequest
	GetProjectId() *string
	SetStatusListShrink(v string) *QuerySessionInfoShrinkRequest
	GetStatusListShrink() *string
}

type QuerySessionInfoShrinkRequest struct {
	// example:
	//
	// 1
	PageNo *int32 `json:"pageNo,omitempty" xml:"pageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 805800890535673856
	ProjectId        *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StatusListShrink *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s QuerySessionInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QuerySessionInfoShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySessionInfoShrinkRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QuerySessionInfoShrinkRequest) GetStatusListShrink() *string {
	return s.StatusListShrink
}

func (s *QuerySessionInfoShrinkRequest) SetPageNo(v int32) *QuerySessionInfoShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetPageSize(v int32) *QuerySessionInfoShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetProjectId(v string) *QuerySessionInfoShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) SetStatusListShrink(v string) *QuerySessionInfoShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *QuerySessionInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
