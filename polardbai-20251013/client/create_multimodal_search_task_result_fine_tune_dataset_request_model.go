// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMultimodalSearchTaskResultFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetDBClusterId() *string
	SetDatasetDescription(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetDatasetDescription() *string
	SetDatasetName(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetDatasetName() *string
	SetResultIndex(v []*int32) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetResultIndex() []*int32
	SetResultMode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetResultMode() *string
	SetTaskId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetTaskId() *string
	SetTopN(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest
	GetTopN() *int32
}

type CreateMultimodalSearchTaskResultFineTuneDatasetRequest struct {
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
	DatasetName *string  `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	ResultIndex []*int32 `json:"ResultIndex,omitempty" xml:"ResultIndex,omitempty" type:"Repeated"`
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

func (s CreateMultimodalSearchTaskResultFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetDatasetDescription() *string {
	return s.DatasetDescription
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetResultIndex() []*int32 {
	return s.ResultIndex
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetResultMode() *string {
	return s.ResultMode
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) GetTopN() *int32 {
	return s.TopN
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetDBClusterId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetDatasetDescription(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.DatasetDescription = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetDatasetName(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.DatasetName = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetResultIndex(v []*int32) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.ResultIndex = v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetResultMode(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.ResultMode = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetTaskId(v string) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.TaskId = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) SetTopN(v int32) *CreateMultimodalSearchTaskResultFineTuneDatasetRequest {
	s.TopN = &v
	return s
}

func (s *CreateMultimodalSearchTaskResultFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
