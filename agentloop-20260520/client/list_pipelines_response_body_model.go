// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPipelinesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelinesResponseBody
	GetNextToken() *string
	SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody
	GetPipelines() []*ListPipelinesResponseBodyPipelines
	SetRequestId(v string) *ListPipelinesResponseBody
	GetRequestId() *string
}

type ListPipelinesResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// MTIzNDU2Nzg5MA==
	NextToken *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Pipelines []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// example:
	//
	// 9ACFB10A-1B2C-3D4E-5F6G-7H8I9J0K1L2M
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelinesResponseBody) GetPipelines() []*ListPipelinesResponseBodyPipelines {
	return s.Pipelines
}

func (s *ListPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelinesResponseBody) SetMaxResults(v int32) *ListPipelinesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListPipelinesResponseBody) SetNextToken(v string) *ListPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesResponseBody) SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelinesResponseBody) SetRequestId(v string) *ListPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelinesResponseBody) Validate() error {
	if s.Pipelines != nil {
		for _, item := range s.Pipelines {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPipelinesResponseBodyPipelines struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-01T00:00:00Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 我的流水线
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// my-pipeline
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-01-02T00:00:00Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// example:
	//
	// my-workspace
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListPipelinesResponseBodyPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelines) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListPipelinesResponseBodyPipelines) GetDescription() *string {
	return s.Description
}

func (s *ListPipelinesResponseBodyPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesResponseBodyPipelines) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPipelinesResponseBodyPipelines) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListPipelinesResponseBodyPipelines) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListPipelinesResponseBodyPipelines) SetCreateTime(v string) *ListPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetDescription(v string) *ListPipelinesResponseBodyPipelines {
	s.Description = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetRegionId(v string) *ListPipelinesResponseBodyPipelines {
	s.RegionId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetUpdateTime(v string) *ListPipelinesResponseBodyPipelines {
	s.UpdateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetWorkspace(v string) *ListPipelinesResponseBodyPipelines {
	s.Workspace = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) Validate() error {
	return dara.Validate(s)
}
