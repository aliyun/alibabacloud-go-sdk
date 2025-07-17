// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeletePrometheusMonitoringResponseBody
	GetCode() *int32
	SetData(v string) *DeletePrometheusMonitoringResponseBody
	GetData() *string
	SetMessage(v string) *DeletePrometheusMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrometheusMonitoringResponseBody
	GetRequestId() *string
}

type DeletePrometheusMonitoringResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4D6C358A-A58B-4F4B-94CE-F5AAF023****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrometheusMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePrometheusMonitoringResponseBody) GetData() *string {
	return s.Data
}

func (s *DeletePrometheusMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrometheusMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusMonitoringResponseBody) SetCode(v int32) *DeletePrometheusMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrometheusMonitoringResponseBody) SetData(v string) *DeletePrometheusMonitoringResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePrometheusMonitoringResponseBody) SetMessage(v string) *DeletePrometheusMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrometheusMonitoringResponseBody) SetRequestId(v string) *DeletePrometheusMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}
