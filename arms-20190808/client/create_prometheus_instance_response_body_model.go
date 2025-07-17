// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreatePrometheusInstanceResponseBody
	GetCode() *int32
	SetData(v string) *CreatePrometheusInstanceResponseBody
	GetData() *string
	SetMessage(v string) *CreatePrometheusInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreatePrometheusInstanceResponseBody
	GetRequestId() *string
}

type CreatePrometheusInstanceResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created Prometheus instance.
	//
	// example:
	//
	// qduukd****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 70675725-8F11-4817-8106-CFE0AD71****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePrometheusInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreatePrometheusInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *CreatePrometheusInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreatePrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePrometheusInstanceResponseBody) SetCode(v int32) *CreatePrometheusInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) SetData(v string) *CreatePrometheusInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) SetMessage(v string) *CreatePrometheusInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) SetRequestId(v string) *CreatePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
