// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityEvaluationTaskResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityEvaluationTaskResponseBody
	GetRequestId() *string
}

type CreateDataQualityEvaluationTaskResponseBody struct {
	// The ID of the new monitor.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2d9ce-38ef-4923-baf6-391a7e656
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityEvaluationTaskResponseBody) SetId(v int64) *CreateDataQualityEvaluationTaskResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskResponseBody) SetRequestId(v string) *CreateDataQualityEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
