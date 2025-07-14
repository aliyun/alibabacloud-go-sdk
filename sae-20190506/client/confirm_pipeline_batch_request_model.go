// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPipelineBatchRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirm(v bool) *ConfirmPipelineBatchRequest
	GetConfirm() *bool
	SetPipelineId(v string) *ConfirmPipelineBatchRequest
	GetPipelineId() *string
}

type ConfirmPipelineBatchRequest struct {
	// true
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	Confirm *bool `json:"Confirm,omitempty" xml:"Confirm,omitempty"`
	// The ID of the batch. You can call the [DescribeChangeOrder](https://www.alibabacloud.com/help/zh/sae/serverless-app-engine-classic/developer-reference/api-sae-2019-05-06-describechangeorder-old?spm=a2c63.p38356.help-menu-search-118957.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2e-vds-feh-***
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ConfirmPipelineBatchRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPipelineBatchRequest) GoString() string {
	return s.String()
}

func (s *ConfirmPipelineBatchRequest) GetConfirm() *bool {
	return s.Confirm
}

func (s *ConfirmPipelineBatchRequest) GetPipelineId() *string {
	return s.PipelineId
}

func (s *ConfirmPipelineBatchRequest) SetConfirm(v bool) *ConfirmPipelineBatchRequest {
	s.Confirm = &v
	return s
}

func (s *ConfirmPipelineBatchRequest) SetPipelineId(v string) *ConfirmPipelineBatchRequest {
	s.PipelineId = &v
	return s
}

func (s *ConfirmPipelineBatchRequest) Validate() error {
	return dara.Validate(s)
}
