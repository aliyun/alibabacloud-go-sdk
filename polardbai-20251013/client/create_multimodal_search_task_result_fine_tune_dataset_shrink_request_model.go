// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetDatasetName() *string
	SetResultIndexShrink(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetResultIndexShrink() *string
	SetResultMode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetResultMode() *string
	SetTaskId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetTaskId() *string
	SetTopN(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest
	GetTopN() *int32
}

type CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// 多模态数据集
	DatasetDescription *string `json:"DatasetDescription,omitempty" xml:"DatasetDescription,omitempty"`
	// example:
	//
	// dataset-1
	DatasetName       *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ResultIndexShrink *string `json:"ResultIndex,omitempty" xml:"ResultIndex,omitempty"`
	// example:
	//
	// ["raw", "rerank"]
	ResultMode *string `json:"ResultMode,omitempty" xml:"ResultMode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 476751c5-*****-39f6aff1df96
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 100
	TopN *int32 `json:"TopN,omitempty" xml:"TopN,omitempty"`
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetResultIndexShrink() *string {
	return s.ResultIndexShrink
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetResultMode() *string {
	return s.ResultMode
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetDBClusterId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetDatasetDescription(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetDatasetName(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetResultIndexShrink(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.ResultIndexShrink = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetResultMode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.ResultMode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetTaskId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) SetTopN(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest {
	s.TopN = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetShrinkRequest) Validate() error {
	return dara.Validate(s)
}
