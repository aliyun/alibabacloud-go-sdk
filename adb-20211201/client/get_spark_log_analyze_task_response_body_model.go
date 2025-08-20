// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSparkLogAnalyzeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkAnalyzeLogTask) *GetSparkLogAnalyzeTaskResponseBody
	GetData() *SparkAnalyzeLogTask
	SetRequestId(v string) *GetSparkLogAnalyzeTaskResponseBody
	GetRequestId() *string
}

type GetSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1DF5AF5B-C803-1861-A0FF-63666A557709
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSparkLogAnalyzeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetSparkLogAnalyzeTaskResponseBody) GetData() *SparkAnalyzeLogTask {
	return s.Data
}

func (s *GetSparkLogAnalyzeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *GetSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *GetSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSparkLogAnalyzeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
