// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePrometheusMonitoringResponseBody
	GetCode() *int32
	SetData(v string) *UpdatePrometheusMonitoringResponseBody
	GetData() *string
	SetMessage(v string) *UpdatePrometheusMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrometheusMonitoringResponseBody
	GetRequestId() *string
}

type UpdatePrometheusMonitoringResponseBody struct {
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
	// 626037F5-FDEB-45B0-804C-B3C92797****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePrometheusMonitoringResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdatePrometheusMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusMonitoringResponseBody) SetCode(v int32) *UpdatePrometheusMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusMonitoringResponseBody) SetData(v string) *UpdatePrometheusMonitoringResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrometheusMonitoringResponseBody) SetMessage(v string) *UpdatePrometheusMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusMonitoringResponseBody) SetRequestId(v string) *UpdatePrometheusMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}
