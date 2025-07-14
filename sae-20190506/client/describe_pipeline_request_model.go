// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPipelineId(v string) *DescribePipelineRequest
	GetPipelineId() *string
}

type DescribePipelineRequest struct {
	// The ID of the batch. You can call the [DescribeChangeOrder](https://help.aliyun.com/document_detail/126617.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 917660ba-5092-44ca-b8e0-80012c96****
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s DescribePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineRequest) GoString() string {
	return s.String()
}

func (s *DescribePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *DescribePipelineRequest) SetPipelineId(v string) *DescribePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *DescribePipelineRequest) Validate() error {
	return dara.Validate(s)
}
