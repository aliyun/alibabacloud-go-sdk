// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListPipelinesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListPipelinesResponseBody
	GetErrorMessage() *string
	SetNextToken(v string) *ListPipelinesResponseBody
	GetNextToken() *string
	SetPipelines(v []*ListPipelinesResponseBodyPipelines) *ListPipelinesResponseBody
	GetPipelines() []*ListPipelinesResponseBodyPipelines
	SetRequestId(v string) *ListPipelinesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListPipelinesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListPipelinesResponseBody
	GetTotalCount() *int64
}

type ListPipelinesResponseBody struct {
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
	// ssaassasass
	NextToken *string                               `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	Pipelines []*ListPipelinesResponseBodyPipelines `json:"pipelines,omitempty" xml:"pipelines,omitempty" type:"Repeated"`
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
	// 50
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPipelinesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListPipelinesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
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

func (s *ListPipelinesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListPipelinesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPipelinesResponseBody) SetErrorCode(v string) *ListPipelinesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListPipelinesResponseBody) SetErrorMessage(v string) *ListPipelinesResponseBody {
	s.ErrorMessage = &v
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

func (s *ListPipelinesResponseBody) SetSuccess(v bool) *ListPipelinesResponseBody {
	s.Success = &v
	return s
}

func (s *ListPipelinesResponseBody) SetTotalCount(v int64) *ListPipelinesResponseBody {
	s.TotalCount = &v
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
	// example:
	//
	// 1586863220000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 22121222
	CreatorAccountId *string `json:"creatorAccountId,omitempty" xml:"creatorAccountId,omitempty"`
	GroupId          *int64  `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// 124
	PipelineId *int64 `json:"pipelineId,omitempty" xml:"pipelineId,omitempty"`
	// example:
	//
	// 流水线
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
}

func (s ListPipelinesResponseBodyPipelines) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponseBodyPipelines) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponseBodyPipelines) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListPipelinesResponseBodyPipelines) GetCreatorAccountId() *string {
	return s.CreatorAccountId
}

func (s *ListPipelinesResponseBodyPipelines) GetGroupId() *int64 {
	return s.GroupId
}

func (s *ListPipelinesResponseBodyPipelines) GetPipelineId() *int64 {
	return s.PipelineId
}

func (s *ListPipelinesResponseBodyPipelines) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesResponseBodyPipelines) SetCreateTime(v int64) *ListPipelinesResponseBodyPipelines {
	s.CreateTime = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetCreatorAccountId(v string) *ListPipelinesResponseBodyPipelines {
	s.CreatorAccountId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetGroupId(v int64) *ListPipelinesResponseBodyPipelines {
	s.GroupId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineId(v int64) *ListPipelinesResponseBodyPipelines {
	s.PipelineId = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) SetPipelineName(v string) *ListPipelinesResponseBodyPipelines {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesResponseBodyPipelines) Validate() error {
	return dara.Validate(s)
}
