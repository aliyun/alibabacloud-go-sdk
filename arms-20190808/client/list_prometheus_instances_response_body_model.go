// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPrometheusInstancesResponseBody
	GetCode() *int32
	SetData(v string) *ListPrometheusInstancesResponseBody
	GetData() *string
	SetMessage(v string) *ListPrometheusInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrometheusInstancesResponseBody
	GetRequestId() *string
}

type ListPrometheusInstancesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The Prometheus instances in the region in the JSON format.
	//
	// example:
	//
	// [{"agentStatus":"0","clusterId":"global-v2-cn-1672753017899-dmjnwtzz","clusterName":"test-GlobalView","clusterType":"GlobalViewV2","commercialConfig":{},"createTime":1656579981000,"id":13785300,"isAdvancedClusterInstalled":false,"isClusterRunning":true,"isControllerInstalled":true,"isIntegrationCenter":false,"regionId":"cn-hongkong","updateTime":1657616273000,"userId":"1672753017899"},{"agentStatus":"0","clusterId":"51a123a61a8f31f0","clusterName":"cloud-product-prometheus_cn-qingdao","clusterType":"cloud-product-prometheus","commercialConfig":{},"controllerId":"51a123a61a8f31f0","createTime":1653532488000,"id":13746658,"isAdvancedClusterInstalled":false,"isClusterRunning":true,"isControllerInstalled":true,"isIntegrationCenter":false,"regionId":"cn-qingdao","updateTime":1653532518000,"userId":"1672753017899"}]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// E9C9DA3D-10FE-472E-9EEF-2D0A3E41****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusInstancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPrometheusInstancesResponseBody) GetData() *string {
	return s.Data
}

func (s *ListPrometheusInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusInstancesResponseBody) SetCode(v int32) *ListPrometheusInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetData(v string) *ListPrometheusInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetMessage(v string) *ListPrometheusInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) SetRequestId(v string) *ListPrometheusInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
