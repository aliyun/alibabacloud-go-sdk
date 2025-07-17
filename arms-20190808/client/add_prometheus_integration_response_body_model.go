// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddPrometheusIntegrationResponseBody
	GetCode() *int32
	SetData(v *AddPrometheusIntegrationResponseBodyData) *AddPrometheusIntegrationResponseBody
	GetData() *AddPrometheusIntegrationResponseBodyData
	SetMessage(v string) *AddPrometheusIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPrometheusIntegrationResponseBody
	GetRequestId() *string
}

type AddPrometheusIntegrationResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *AddPrometheusIntegrationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3703B98C-335E-5BA7-972E-F90E9E768A85
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddPrometheusIntegrationResponseBody) GetData() *AddPrometheusIntegrationResponseBodyData {
	return s.Data
}

func (s *AddPrometheusIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPrometheusIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrometheusIntegrationResponseBody) SetCode(v int32) *AddPrometheusIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *AddPrometheusIntegrationResponseBody) SetData(v *AddPrometheusIntegrationResponseBodyData) *AddPrometheusIntegrationResponseBody {
	s.Data = v
	return s
}

func (s *AddPrometheusIntegrationResponseBody) SetMessage(v string) *AddPrometheusIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *AddPrometheusIntegrationResponseBody) SetRequestId(v string) *AddPrometheusIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrometheusIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddPrometheusIntegrationResponseBodyData struct {
	// The ID of the exporter.
	//
	// example:
	//
	// 2829
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the exporter.
	//
	// example:
	//
	// hw-cloud02
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s AddPrometheusIntegrationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusIntegrationResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrometheusIntegrationResponseBodyData) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *AddPrometheusIntegrationResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *AddPrometheusIntegrationResponseBodyData) SetInstanceId(v int64) *AddPrometheusIntegrationResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *AddPrometheusIntegrationResponseBodyData) SetInstanceName(v string) *AddPrometheusIntegrationResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *AddPrometheusIntegrationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
