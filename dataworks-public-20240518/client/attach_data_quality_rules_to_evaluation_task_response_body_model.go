// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDataQualityRulesToEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachDataQualityRulesToEvaluationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AttachDataQualityRulesToEvaluationTaskResponseBody
	GetSuccess() *bool
}

type AttachDataQualityRulesToEvaluationTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E6F0DBDD-5AD8-4870-A6A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The value of the association is as follows:
	//
	// - true: The call is successful.
	//
	// - false: the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AttachDataQualityRulesToEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachDataQualityRulesToEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *AttachDataQualityRulesToEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachDataQualityRulesToEvaluationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AttachDataQualityRulesToEvaluationTaskResponseBody) SetRequestId(v string) *AttachDataQualityRulesToEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskResponseBody) SetSuccess(v bool) *AttachDataQualityRulesToEvaluationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *AttachDataQualityRulesToEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
