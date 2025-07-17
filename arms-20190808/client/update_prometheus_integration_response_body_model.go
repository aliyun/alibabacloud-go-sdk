// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePrometheusIntegrationResponseBody
	GetCode() *int32
	SetData(v *UpdatePrometheusIntegrationResponseBodyData) *UpdatePrometheusIntegrationResponseBody
	GetData() *UpdatePrometheusIntegrationResponseBodyData
	SetMessage(v string) *UpdatePrometheusIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrometheusIntegrationResponseBody
	GetRequestId() *string
}

type UpdatePrometheusIntegrationResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *UpdatePrometheusIntegrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2DB771C3-D1BB-5363-8A5F-ADB2AF2948DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePrometheusIntegrationResponseBody) GetData() *UpdatePrometheusIntegrationResponseBodyData {
	return s.Data
}

func (s *UpdatePrometheusIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusIntegrationResponseBody) SetCode(v int32) *UpdatePrometheusIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBody) SetData(v *UpdatePrometheusIntegrationResponseBodyData) *UpdatePrometheusIntegrationResponseBody {
	s.Data = v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBody) SetMessage(v string) *UpdatePrometheusIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBody) SetRequestId(v string) *UpdatePrometheusIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdatePrometheusIntegrationResponseBodyData struct {
	// The exporter ID.
	//
	// example:
	//
	// 2866
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The exporter name.
	//
	// example:
	//
	// inet
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdatePrometheusIntegrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusIntegrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusIntegrationResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *UpdatePrometheusIntegrationResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdatePrometheusIntegrationResponseBodyData) SetInstanceId(v int64) *UpdatePrometheusIntegrationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBodyData) SetInstanceName(v string) *UpdatePrometheusIntegrationResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *UpdatePrometheusIntegrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
