// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallManagedPrometheusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InstallManagedPrometheusResponseBody
	GetCode() *int32
	SetData(v string) *InstallManagedPrometheusResponseBody
	GetData() *string
	SetMessage(v string) *InstallManagedPrometheusResponseBody
	GetMessage() *string
	SetRequestId(v string) *InstallManagedPrometheusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InstallManagedPrometheusResponseBody
	GetSuccess() *bool
}

type InstallManagedPrometheusResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response content. In most cases, the installation status of the Prometheus agent is returned.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned if the Prometheus agent failed to be installed.
	//
	// example:
	//
	// vpcId is blank
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// String	2A0CEDF1-06FE-44AC-8E21-21A5BE65****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the Prometheus agent was installed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InstallManagedPrometheusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallManagedPrometheusResponseBody) GoString() string {
	return s.String()
}

func (s *InstallManagedPrometheusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallManagedPrometheusResponseBody) GetData() *string {
	return s.Data
}

func (s *InstallManagedPrometheusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallManagedPrometheusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallManagedPrometheusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InstallManagedPrometheusResponseBody) SetCode(v int32) *InstallManagedPrometheusResponseBody {
	s.Code = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetData(v string) *InstallManagedPrometheusResponseBody {
	s.Data = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetMessage(v string) *InstallManagedPrometheusResponseBody {
	s.Message = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetRequestId(v string) *InstallManagedPrometheusResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) SetSuccess(v bool) *InstallManagedPrometheusResponseBody {
	s.Success = &v
	return s
}

func (s *InstallManagedPrometheusResponseBody) Validate() error {
	return dara.Validate(s)
}
