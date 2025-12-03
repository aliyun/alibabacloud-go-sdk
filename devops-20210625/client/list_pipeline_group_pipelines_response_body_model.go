// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineGroupPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelineGroupPipelinesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelineGroupPipelinesResponseBody
	GetErrorMessage() *string
	SetNextToken(v string) *ListPipelineGroupPipelinesResponseBody
	GetNextToken() *string
	SetPipelines(v []*ListPipelineGroupPipelinesResponseBodyPipelines) *ListPipelineGroupPipelinesResponseBody
	GetPipelines() []*ListPipelineGroupPipelinesResponseBodyPipelines
	SetRequestId(v string) *ListPipelineGroupPipelinesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelineGroupPipelinesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListPipelineGroupPipelinesResponseBody
	GetTotalCount() *int32
}

type ListPipelineGroupPipelinesResponseBody struct {
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// aaaa
	NextToken *string                                            `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Pipelines []*ListPipelineGroupPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelineGroupPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelineGroupPipelinesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListPipelineGroupPipelinesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelineGroupPipelinesResponseBody) GetPipelines() []*ListPipelineGroupPipelinesResponseBodyPipelines {
	return s.Pipelines
}

func (s *ListPipelineGroupPipelinesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPipelineGroupPipelinesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelineGroupPipelinesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListPipelineGroupPipelinesResponseBody) SetErrorCode(v string) *ListPipelineGroupPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetErrorMessage(v string) *ListPipelineGroupPipelinesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetNextToken(v string) *ListPipelineGroupPipelinesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetPipelines(v []*ListPipelineGroupPipelinesResponseBodyPipelines) *ListPipelineGroupPipelinesResponseBody {
	s.Pipelines = v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetRequestId(v string) *ListPipelineGroupPipelinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetSuccess(v bool) *ListPipelineGroupPipelinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) SetTotalCount(v int32) *ListPipelineGroupPipelinesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBody) Validate() error {
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

type ListPipelineGroupPipelinesResponseBodyPipelines struct {
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 1111
	PipelineId   *int64  `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
}

func (s ListPipelineGroupPipelinesResponseBodyPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineGroupPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetCreateTime(v int64) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetPipelineId(v int64) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelineGroupPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

func (s *ListPipelineGroupPipelinesResponseBodyPipelines) Validate() error {
	return dara.Validate(s)
}
