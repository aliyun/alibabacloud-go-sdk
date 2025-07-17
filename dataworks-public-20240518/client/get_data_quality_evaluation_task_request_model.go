// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityEvaluationTaskRequest
	GetId() *int64
}

type GetDataQualityEvaluationTaskRequest struct {
	// The ID of the data quality monitor.
	//
	// example:
	//
	// 1006455182
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityEvaluationTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskRequest) SetId(v int64) *GetDataQualityEvaluationTaskRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskRequest) Validate() error {
	return dara.Validate(s)
}
