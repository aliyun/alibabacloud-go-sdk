// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddPrometheusInstanceResponseBody
	GetCode() *int32
	SetData(v string) *AddPrometheusInstanceResponseBody
	GetData() *string
	SetMessage(v string) *AddPrometheusInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPrometheusInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddPrometheusInstanceResponseBody
	GetSuccess() *bool
}

type AddPrometheusInstanceResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned. { "RequestId": the request ID, "Data": "{ "clusterType": the cluster type, "remoteWriteUrl": the public URL for remote write, "internetGrafanaUrl": the internal URL for Grafana, "authToken": indicates whether authentication is enabled, "internetPushGatewayUrl": the internal URL for Pushgateway, "clusterId": the cluster ID, "internetRemoteReadUrl": the internal URL for remote read, "remoteReadUrl": the public URL for remote read, "grafanaUrl": the public URL for Grafana, "pushGatewayUrl": the public URL for Pushgateway, "internetRemoteWriteUrl": the internal URL for remote write}" }
	//
	// example:
	//
	// {
	//
	//   "RequestId": "1293091C-54AD-50FE-B787-E314B94B35AB",
	//
	//   "Data": "{
	//
	//   "clusterType":"remote-write-prometheus",
	//
	//   "remoteWriteUrl":"http://cn-hu/api/v3/write",
	//
	//   "internetGrafanaUrl":"https://cn-hanga/cn-hangzhou",
	//
	//   "authToken":false,
	//
	//   "internetPushGatewayUrl":"https://cangzhou/api/v2",
	//
	//   "clusterId":"vrju1lj3sa|123456",
	//
	//   "internetRemoteReadUrl":"https://cn-hangzh67cn-hangzhou/api/v1/read",
	//
	//   "remoteReadUrl":"http://cn-hanou/api/v1/read",
	//
	//   "grafanaUrl":"http://cn-angzhou",
	//
	//   "pushGatewayUrl":"htt1lj3sa/cn-hangzhou/api/v2",
	//
	//   "internetRemoteWriteUrl":"httpsngzhou/api/v3/write"}"
	//
	// }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9319A57D-2D9E-472A-B69B-CF3CD16D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddPrometheusInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddPrometheusInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *AddPrometheusInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPrometheusInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrometheusInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddPrometheusInstanceResponseBody) SetCode(v int32) *AddPrometheusInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) SetData(v string) *AddPrometheusInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) SetMessage(v string) *AddPrometheusInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) SetRequestId(v string) *AddPrometheusInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) SetSuccess(v bool) *AddPrometheusInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *AddPrometheusInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
