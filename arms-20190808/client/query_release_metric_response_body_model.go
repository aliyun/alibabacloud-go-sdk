// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReleaseMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryReleaseMetricResponseBody
	GetData() *string
	SetRequestId(v string) *QueryReleaseMetricResponseBody
	GetRequestId() *string
}

type QueryReleaseMetricResponseBody struct {
	// The returned metric data.
	//
	// example:
	//
	// {"data":{"SystemCpuUser":{"all":[{"date":1632798718000,"val":4.3277,"dim":"SystemCpuUser"},{"date":1632798733000,"val":8.1091,"dim":"SystemCpuUser"}]},"SystemMemUtil":{"all":[{"date":1632798718000,"val":73.4227,"dim":"SystemMemUtil"},{"date":1632798733000,"val":93.0977,"dim":"SystemMemUtil"}]}},"success":true}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 46355DD8-FC56-40C5-BFC6-269DE4F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryReleaseMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReleaseMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReleaseMetricResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryReleaseMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReleaseMetricResponseBody) SetData(v string) *QueryReleaseMetricResponseBody {
	s.Data = &v
	return s
}

func (s *QueryReleaseMetricResponseBody) SetRequestId(v string) *QueryReleaseMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReleaseMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
