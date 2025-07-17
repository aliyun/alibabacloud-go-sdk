// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallManagedPrometheusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UninstallManagedPrometheusResponseBody
	GetCode() *int32
	SetData(v string) *UninstallManagedPrometheusResponseBody
	GetData() *string
	SetMessage(v string) *UninstallManagedPrometheusResponseBody
	GetMessage() *string
	SetRequestId(v string) *UninstallManagedPrometheusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UninstallManagedPrometheusResponseBody
	GetSuccess() *bool
}

type UninstallManagedPrometheusResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response content. The status of the Prometheus instance is returned.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message that is returned if the request fails.
	//
	// example:
	//
	// vpcId is blank
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the Prometheus instance was removed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UninstallManagedPrometheusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UninstallManagedPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallManagedPrometheusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UninstallManagedPrometheusResponseBody) GetData() *string {
	return s.Data
}

func (s *UninstallManagedPrometheusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UninstallManagedPrometheusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UninstallManagedPrometheusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UninstallManagedPrometheusResponseBody) SetCode(v int32) *UninstallManagedPrometheusResponseBody {
	s.Code = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetData(v string) *UninstallManagedPrometheusResponseBody {
	s.Data = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetMessage(v string) *UninstallManagedPrometheusResponseBody {
	s.Message = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetRequestId(v string) *UninstallManagedPrometheusResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) SetSuccess(v bool) *UninstallManagedPrometheusResponseBody {
	s.Success = &v
	return s
}

func (s *UninstallManagedPrometheusResponseBody) Validate() error {
	return dara.Validate(s)
}
