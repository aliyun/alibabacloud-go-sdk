// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityEvaluationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityEvaluationTaskResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityEvaluationTaskResponseBody struct {
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 8abcb91f-d266-4073-b907-2ed670378ed1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataQualityEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityEvaluationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityEvaluationTaskResponseBody) SetRequestId(v string) *UpdateDataQualityEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskResponseBody) SetSuccess(v bool) *UpdateDataQualityEvaluationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
