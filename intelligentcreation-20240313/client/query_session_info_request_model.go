// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySessionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *QuerySessionInfoRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QuerySessionInfoRequest
	GetPageSize() *int32
	SetProjectId(v string) *QuerySessionInfoRequest
	GetProjectId() *string
	SetStatusList(v []*string) *QuerySessionInfoRequest
	GetStatusList() []*string
}

type QuerySessionInfoRequest struct {
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
	ProjectId  *string   `json:"projectId,omitempty" xml:"projectId,omitempty"`
	StatusList []*string `json:"statusList,omitempty" xml:"statusList,omitempty" type:"Repeated"`
}

func (s QuerySessionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySessionInfoRequest) GoString() string {
	return s.String()
}

func (s *QuerySessionInfoRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QuerySessionInfoRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QuerySessionInfoRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *QuerySessionInfoRequest) GetStatusList() []*string {
	return s.StatusList
}

func (s *QuerySessionInfoRequest) SetPageNo(v int32) *QuerySessionInfoRequest {
	s.PageNo = &v
	return s
}

func (s *QuerySessionInfoRequest) SetPageSize(v int32) *QuerySessionInfoRequest {
	s.PageSize = &v
	return s
}

func (s *QuerySessionInfoRequest) SetProjectId(v string) *QuerySessionInfoRequest {
	s.ProjectId = &v
	return s
}

func (s *QuerySessionInfoRequest) SetStatusList(v []*string) *QuerySessionInfoRequest {
	s.StatusList = v
	return s
}

func (s *QuerySessionInfoRequest) Validate() error {
	return dara.Validate(s)
}
