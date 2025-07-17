// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedPrometheusStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetManagedPrometheusStatusResponseBody
	GetCode() *int32
	SetData(v string) *GetManagedPrometheusStatusResponseBody
	GetData() *string
	SetMessage(v string) *GetManagedPrometheusStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetManagedPrometheusStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetManagedPrometheusStatusResponseBody
	GetSuccess() *bool
}

type GetManagedPrometheusStatusResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The installation status of the Prometheus agent.
	//
	// 	- Installing: The Prometheus agent is installed and no registration information is available.
	//
	// 	- Succeed: The Prometheus agent is installed and registered.
	//
	// 	- Failure: The Prometheus agent failed to be installed or registered.
	//
	// 	- Unknown: The installation status of the Prometheus agent is unknown.
	//
	// example:
	//
	// Installing
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetManagedPrometheusStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetManagedPrometheusStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetManagedPrometheusStatusResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetManagedPrometheusStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetManagedPrometheusStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetManagedPrometheusStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetManagedPrometheusStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetManagedPrometheusStatusResponseBody) SetCode(v int32) *GetManagedPrometheusStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetManagedPrometheusStatusResponseBody) SetData(v string) *GetManagedPrometheusStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetManagedPrometheusStatusResponseBody) SetMessage(v string) *GetManagedPrometheusStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetManagedPrometheusStatusResponseBody) SetRequestId(v string) *GetManagedPrometheusStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetManagedPrometheusStatusResponseBody) SetSuccess(v bool) *GetManagedPrometheusStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetManagedPrometheusStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
