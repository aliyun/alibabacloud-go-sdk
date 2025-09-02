// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListMetaDBRequest
	GetClusterId() *string
	SetDataSourceType(v string) *ListMetaDBRequest
	GetDataSourceType() *string
	SetPageNum(v int32) *ListMetaDBRequest
	GetPageNum() *int32
	SetPageSize(v int32) *ListMetaDBRequest
	GetPageSize() *int32
	SetProjectId(v int64) *ListMetaDBRequest
	GetProjectId() *int64
}

type ListMetaDBRequest struct {
	// The ID of the E-MapReduce (EMR) cluster. You can log on to the [EMR console](https://emr.console.aliyun.com/?spm=a2c4g.11186623.0.0.965cc5c2GeiHet#/cn-hangzhou) to query the ID.
	//
	// example:
	//
	// abc
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The type of the data source. Set the value to emr.
	//
	// This parameter is required.
	//
	// example:
	//
	// emr
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListMetaDBRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDBRequest) GoString() string {
	return s.String()
}

func (s *ListMetaDBRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListMetaDBRequest) GetDataSourceType() *string {
	return s.DataSourceType
}

func (s *ListMetaDBRequest) GetPageNum() *int32 {
	return s.PageNum
}

func (s *ListMetaDBRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMetaDBRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListMetaDBRequest) SetClusterId(v string) *ListMetaDBRequest {
	s.ClusterId = &v
	return s
}

func (s *ListMetaDBRequest) SetDataSourceType(v string) *ListMetaDBRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListMetaDBRequest) SetPageNum(v int32) *ListMetaDBRequest {
	s.PageNum = &v
	return s
}

func (s *ListMetaDBRequest) SetPageSize(v int32) *ListMetaDBRequest {
	s.PageSize = &v
	return s
}

func (s *ListMetaDBRequest) SetProjectId(v int64) *ListMetaDBRequest {
	s.ProjectId = &v
	return s
}

func (s *ListMetaDBRequest) Validate() error {
	return dara.Validate(s)
}
