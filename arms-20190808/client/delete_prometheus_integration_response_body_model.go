// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeletePrometheusIntegrationResponseBody
	GetCode() *int32
	SetData(v string) *DeletePrometheusIntegrationResponseBody
	GetData() *string
	SetMessage(v string) *DeletePrometheusIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrometheusIntegrationResponseBody
	GetRequestId() *string
}

type DeletePrometheusIntegrationResponseBody struct {
	// The status code. The status code 200 indicates that the request was successful. If another status code is returned, the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// success or an error message.
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
	// Id of the request
	//
	// example:
	//
	// 19F54318-CC92-5567-BF66-CB029EC44C84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrometheusIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePrometheusIntegrationResponseBody) GetData() *string {
	return s.Data
}

func (s *DeletePrometheusIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrometheusIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusIntegrationResponseBody) SetCode(v int32) *DeletePrometheusIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrometheusIntegrationResponseBody) SetData(v string) *DeletePrometheusIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePrometheusIntegrationResponseBody) SetMessage(v string) *DeletePrometheusIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrometheusIntegrationResponseBody) SetRequestId(v string) *DeletePrometheusIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
