// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdatePrometheusInstanceResponseBody
	GetCode() *int32
	SetData(v string) *UpdatePrometheusInstanceResponseBody
	GetData() *string
	SetMessage(v string) *UpdatePrometheusInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdatePrometheusInstanceResponseBody
	GetRequestId() *string
}

type UpdatePrometheusInstanceResponseBody struct {
	// The returned status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
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
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdatePrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdatePrometheusInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdatePrometheusInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdatePrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdatePrometheusInstanceResponseBody) SetCode(v int32) *UpdatePrometheusInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) SetData(v string) *UpdatePrometheusInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) SetMessage(v string) *UpdatePrometheusInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) SetRequestId(v string) *UpdatePrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
