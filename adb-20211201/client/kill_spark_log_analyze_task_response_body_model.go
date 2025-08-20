// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSparkLogAnalyzeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SparkAnalyzeLogTask) *KillSparkLogAnalyzeTaskResponseBody
	GetData() *SparkAnalyzeLogTask
	SetRequestId(v string) *KillSparkLogAnalyzeTaskResponseBody
	GetRequestId() *string
}

type KillSparkLogAnalyzeTaskResponseBody struct {
	// The information about the Spark log analysis task.
	Data *SparkAnalyzeLogTask `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1DF5AF5B-C803-1861-A0FF-63666A557709
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s KillSparkLogAnalyzeTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillSparkLogAnalyzeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *KillSparkLogAnalyzeTaskResponseBody) GetData() *SparkAnalyzeLogTask {
	return s.Data
}

func (s *KillSparkLogAnalyzeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillSparkLogAnalyzeTaskResponseBody) SetData(v *SparkAnalyzeLogTask) *KillSparkLogAnalyzeTaskResponseBody {
	s.Data = v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponseBody) SetRequestId(v string) *KillSparkLogAnalyzeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillSparkLogAnalyzeTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
