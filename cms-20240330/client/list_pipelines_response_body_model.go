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
	// The number of results returned on the current page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A pagination token. If this parameter is not empty, use it in a subsequent request to get the next page of results. If this parameter is empty, all results have been returned.
	//
	// example:
	//
	// 3f0d6785770d5fb308f0605d718d422a227c38f96117633678f029842acd19039329e8281583b3da7bc598dfc4c1973e
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// A list of pipelines.
	Pipelines []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0B9377D9-C56B-5C2E-A8A4-************
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
	// The creation time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-02-28T07:14:17Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the pipeline.
	//
	// example:
	//
	// test pipeline
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The pipeline name.
	//
	// example:
	//
	// pipeline-name-1
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The update time.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	//
	// example:
	//
	// 2026-03-24T06:58:22Z
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// workspace-test
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
