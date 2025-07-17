// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v string) *ListPrometheusGlobalViewResponseBody
	GetData() *string
	SetMessage(v string) *ListPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type ListPrometheusGlobalViewResponseBody struct {
	// Status code. Description 200 means success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of global aggregation instances. The value of this parameter is a string in the JSON format.
	//
	// example:
	//
	// [ {groupName: "the name of the global aggregation instance", clusterId: "global-v2-clusterid", endpoint: "cn-hangzhou"}, // ..... more items ]
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// More information.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// DBDCE95A-A0DD-5FC5-97CC-EEFC3D814385
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *ListPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListPrometheusGlobalViewResponseBody) GetData() *string {
	return s.Data
}

func (s *ListPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPrometheusGlobalViewResponseBody) SetCode(v int32) *ListPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *ListPrometheusGlobalViewResponseBody) SetData(v string) *ListPrometheusGlobalViewResponseBody {
	s.Data = &v
	return s
}

func (s *ListPrometheusGlobalViewResponseBody) SetMessage(v string) *ListPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *ListPrometheusGlobalViewResponseBody) SetRequestId(v string) *ListPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}
