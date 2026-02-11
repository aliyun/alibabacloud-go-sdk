// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMetricResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *QueryMetricResponseBody
	GetData() *string
	SetRequestId(v string) *QueryMetricResponseBody
	GetRequestId() *string
}

type QueryMetricResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryMetricResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMetricResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMetricResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryMetricResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMetricResponseBody) SetData(v string) *QueryMetricResponseBody {
	s.Data = &v
	return s
}

func (s *QueryMetricResponseBody) SetRequestId(v string) *QueryMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMetricResponseBody) Validate() error {
	return dara.Validate(s)
}
