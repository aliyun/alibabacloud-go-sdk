// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityEvaluationTaskInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityEvaluationTaskInstanceRequest
	GetId() *int64
}

type GetDataQualityEvaluationTaskInstanceRequest struct {
	// The ID of the data quality monitoring instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7227550902
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetDataQualityEvaluationTaskInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityEvaluationTaskInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityEvaluationTaskInstanceRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityEvaluationTaskInstanceRequest) SetId(v int64) *GetDataQualityEvaluationTaskInstanceRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityEvaluationTaskInstanceRequest) Validate() error {
	return dara.Validate(s)
}
