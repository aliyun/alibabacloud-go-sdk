// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSparkLogAnalyzeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkAnalyzeLogTask) *SubmitSparkLogAnalyzeTaskResponseBody
	GetData() *SparkAnalyzeLogTask
	SetRequestId(v string) *SubmitSparkLogAnalyzeTaskResponseBody
	GetRequestId() *string
}

type SubmitSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1DF5AF5B-C803-1861-A0FF-63666A557709
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSparkLogAnalyzeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) GetData() *SparkAnalyzeLogTask {
	return s.Data
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *SubmitSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *SubmitSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSparkLogAnalyzeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
