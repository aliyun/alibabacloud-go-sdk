// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogRunMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *LogRunMetricsResponseBody
	GetRequestId() *string
}

type LogRunMetricsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s LogRunMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s LogRunMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *LogRunMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *LogRunMetricsResponseBody) SetRequestId(v string) *LogRunMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *LogRunMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
