// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusMonitoringStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePrometheusMonitoringStatusResponseBody
	GetCode() *int32
	SetData(v string) *UpdatePrometheusMonitoringStatusResponseBody
	GetData() *string
	SetMessage(v string) *UpdatePrometheusMonitoringStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrometheusMonitoringStatusResponseBody
	GetRequestId() *string
}

type UpdatePrometheusMonitoringStatusResponseBody struct {
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
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusMonitoringStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusMonitoringStatusResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) SetCode(v int32) *UpdatePrometheusMonitoringStatusResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) SetData(v string) *UpdatePrometheusMonitoringStatusResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) SetMessage(v string) *UpdatePrometheusMonitoringStatusResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) SetRequestId(v string) *UpdatePrometheusMonitoringStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusMonitoringStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
