// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetricQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *MetricQueryResponseBody
	GetData() *string
	SetRequestId(v string) *MetricQueryResponseBody
	GetRequestId() *string
}

type MetricQueryResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MetricQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MetricQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MetricQueryResponseBody) GetData() *string {
	return s.Data
}

func (s *MetricQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MetricQueryResponseBody) SetData(v string) *MetricQueryResponseBody {
	s.Data = &v
	return s
}

func (s *MetricQueryResponseBody) SetRequestId(v string) *MetricQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MetricQueryResponseBody) Validate() error {
	return dara.Validate(s)
}
