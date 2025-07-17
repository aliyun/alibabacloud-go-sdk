// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityEvaluationTaskInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityEvaluationTaskInstanceResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityEvaluationTaskInstanceResponseBody
	GetRequestId() *string
}

type CreateDataQualityEvaluationTaskInstanceResponseBody struct {
	// The ID of the data quality monitoring instance.
	//
	// example:
	//
	// 22130
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ecb967ec-c137-48****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityEvaluationTaskInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityEvaluationTaskInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityEvaluationTaskInstanceResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityEvaluationTaskInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityEvaluationTaskInstanceResponseBody) SetId(v int64) *CreateDataQualityEvaluationTaskInstanceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceResponseBody) SetRequestId(v string) *CreateDataQualityEvaluationTaskInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityEvaluationTaskInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
