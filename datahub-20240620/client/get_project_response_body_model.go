// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *GetProjectResponseBody
	GetComment() *string
	SetCreateTime(v string) *GetProjectResponseBody
	GetCreateTime() *string
	SetCreator(v string) *GetProjectResponseBody
	GetCreator() *string
	SetProjectName(v string) *GetProjectResponseBody
	GetProjectName() *string
	SetRequestId(v string) *GetProjectResponseBody
	GetRequestId() *string
	SetStorage(v int64) *GetProjectResponseBody
	GetStorage() *int64
	SetSuccess(v bool) *GetProjectResponseBody
	GetSuccess() *bool
	SetUpdateTime(v string) *GetProjectResponseBody
	GetUpdateTime() *string
	SetVpcWhitelist(v string) *GetProjectResponseBody
	GetVpcWhitelist() *string
}

type GetProjectResponseBody struct {
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1724041098000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1559031978056215
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// test_project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// A20A7093-8FE0-058C-BE0C-3C8057D5F1A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 12252454
	Storage *int64 `json:"Storage,omitempty" xml:"Storage,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1724041098000
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// [11.22.33.44]
	VpcWhitelist *string `json:"VpcWhitelist,omitempty" xml:"VpcWhitelist,omitempty"`
}

func (s GetProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectResponseBody) GetComment() *string {
	return s.Comment
}

func (s *GetProjectResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProjectResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *GetProjectResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectResponseBody) GetStorage() *int64 {
	return s.Storage
}

func (s *GetProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetProjectResponseBody) GetVpcWhitelist() *string {
	return s.VpcWhitelist
}

func (s *GetProjectResponseBody) SetComment(v string) *GetProjectResponseBody {
	s.Comment = &v
	return s
}

func (s *GetProjectResponseBody) SetCreateTime(v string) *GetProjectResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetCreator(v string) *GetProjectResponseBody {
	s.Creator = &v
	return s
}

func (s *GetProjectResponseBody) SetProjectName(v string) *GetProjectResponseBody {
	s.ProjectName = &v
	return s
}

func (s *GetProjectResponseBody) SetRequestId(v string) *GetProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectResponseBody) SetStorage(v int64) *GetProjectResponseBody {
	s.Storage = &v
	return s
}

func (s *GetProjectResponseBody) SetSuccess(v bool) *GetProjectResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectResponseBody) SetUpdateTime(v string) *GetProjectResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetProjectResponseBody) SetVpcWhitelist(v string) *GetProjectResponseBody {
	s.VpcWhitelist = &v
	return s
}

func (s *GetProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
