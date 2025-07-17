// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityEvaluationTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityEvaluationTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityEvaluationTaskResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityEvaluationTaskResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 0bc1ec92159376****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the deletion is successful.
	//
	// - true: Successful
	//
	// - false: Failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataQualityEvaluationTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityEvaluationTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityEvaluationTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityEvaluationTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityEvaluationTaskResponseBody) SetRequestId(v string) *DeleteDataQualityEvaluationTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityEvaluationTaskResponseBody) SetSuccess(v bool) *DeleteDataQualityEvaluationTaskResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityEvaluationTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
