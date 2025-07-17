// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeletePrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v string) *DeletePrometheusGlobalViewResponseBody
	GetData() *string
	SetMessage(v string) *DeletePrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type DeletePrometheusGlobalViewResponseBody struct {
	// The status code. The status code 200 indicates a successful request, whereas others indicate a failed request.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters in the JSON format.
	//
	// example:
	//
	// {"Success":true,"Msg":"OK"}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The supplemental message providing additional context about the response.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 337B8F7E-0A64-5768-9225-E9B3CF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeletePrometheusGlobalViewResponseBody) GetData() *string {
	return s.Data
}

func (s *DeletePrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrometheusGlobalViewResponseBody) SetCode(v int32) *DeletePrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponseBody) SetData(v string) *DeletePrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponseBody) SetMessage(v string) *DeletePrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponseBody) SetRequestId(v string) *DeletePrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}
