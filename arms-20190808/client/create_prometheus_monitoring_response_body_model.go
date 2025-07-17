// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusMonitoringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePrometheusMonitoringResponseBody
	GetCode() *int32
	SetData(v string) *CreatePrometheusMonitoringResponseBody
	GetData() *string
	SetMessage(v string) *CreatePrometheusMonitoringResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePrometheusMonitoringResponseBody
	GetRequestId() *string
}

type CreatePrometheusMonitoringResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the monitoring configuration that was added, or the exception information.
	//
	// example:
	//
	// name1
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

func (s CreatePrometheusMonitoringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusMonitoringResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusMonitoringResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePrometheusMonitoringResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePrometheusMonitoringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusMonitoringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusMonitoringResponseBody) SetCode(v int32) *CreatePrometheusMonitoringResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrometheusMonitoringResponseBody) SetData(v string) *CreatePrometheusMonitoringResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePrometheusMonitoringResponseBody) SetMessage(v string) *CreatePrometheusMonitoringResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrometheusMonitoringResponseBody) SetRequestId(v string) *CreatePrometheusMonitoringResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusMonitoringResponseBody) Validate() error {
	return dara.Validate(s)
}
