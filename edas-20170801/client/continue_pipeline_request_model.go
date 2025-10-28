// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuePipelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirm(v bool) *ContinuePipelineRequest
	GetConfirm() *bool
	SetPipelineId(v string) *ContinuePipelineRequest
	GetPipelineId() *string
}

type ContinuePipelineRequest struct {
	// Specifies whether to release the next batch. Valid values:
	//
	// 	- true: releases the next batch.
	//
	// 	- false: does not release the next batch.
	//
	// example:
	//
	// true
	Confirm *bool `json:"Confirm,omitempty" xml:"Confirm,omitempty"`
	// The ID of the change process. You can call the GetChangeOrderInfo operation to query the ID of the change process that corresponds to a specific batch. For more information, see [GetChangeOrderInfo](https://help.aliyun.com/document_detail/62072.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5648dbd7-df13********************
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ContinuePipelineRequest) String() string {
	return dara.Prettify(s)
}

func (s ContinuePipelineRequest) GoString() string {
	return s.String()
}

func (s *ContinuePipelineRequest) GetConfirm() *bool {
	return s.Confirm
}

func (s *ContinuePipelineRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ContinuePipelineRequest) SetConfirm(v bool) *ContinuePipelineRequest {
	s.Confirm = &v
	return s
}

func (s *ContinuePipelineRequest) SetPipelineId(v string) *ContinuePipelineRequest {
	s.PipelineId = &v
	return s
}

func (s *ContinuePipelineRequest) Validate() error {
	return dara.Validate(s)
}
