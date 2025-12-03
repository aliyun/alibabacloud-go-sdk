// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePipelineBaseInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvId(v int64) *UpdatePipelineBaseInfoRequest
	GetEnvId() *int64
	SetPipelineName(v string) *UpdatePipelineBaseInfoRequest
	GetPipelineName() *string
	SetTagList(v string) *UpdatePipelineBaseInfoRequest
	GetTagList() *string
}

type UpdatePipelineBaseInfoRequest struct {
	// example:
	//
	// 0
	EnvId *int64 `json:"envId,omitempty" xml:"envId,omitempty"`
	// This parameter is required.
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
	// example:
	//
	// "11,222,33"
	TagList *string `json:"tagList,omitempty" xml:"tagList,omitempty"`
}

func (s UpdatePipelineBaseInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePipelineBaseInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdatePipelineBaseInfoRequest) GetEnvId() *int64 {
	return s.EnvId
}

func (s *UpdatePipelineBaseInfoRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *UpdatePipelineBaseInfoRequest) GetTagList() *string {
	return s.TagList
}

func (s *UpdatePipelineBaseInfoRequest) SetEnvId(v int64) *UpdatePipelineBaseInfoRequest {
	s.EnvId = &v
	return s
}

func (s *UpdatePipelineBaseInfoRequest) SetPipelineName(v string) *UpdatePipelineBaseInfoRequest {
	s.PipelineName = &v
	return s
}

func (s *UpdatePipelineBaseInfoRequest) SetTagList(v string) *UpdatePipelineBaseInfoRequest {
	s.TagList = &v
	return s
}

func (s *UpdatePipelineBaseInfoRequest) Validate() error {
	return dara.Validate(s)
}
