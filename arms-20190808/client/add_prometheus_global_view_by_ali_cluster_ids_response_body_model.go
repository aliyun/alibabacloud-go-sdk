// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewByAliClusterIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddPrometheusGlobalViewByAliClusterIdsResponseBody
	GetCode() *int32
	SetData(v *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) *AddPrometheusGlobalViewByAliClusterIdsResponseBody
	GetData() *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData
	SetMessage(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBody
	GetRequestId() *string
}

type AddPrometheusGlobalViewByAliClusterIdsResponseBody struct {
	// The status code. The HTTP 200 status code indicates a successful request, while others indicate error conditions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The struct returned.
	Data *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 3A0EA2AF-C9B3-555C-B9D5-5DD8F5EF98A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) GetData() *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	return s.Data
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetCode(v int32) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.Code = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetData(v *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.Data = v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetMessage(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.Message = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) SetRequestId(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddPrometheusGlobalViewByAliClusterIdsResponseBodyData struct {
	// The Info-level information.
	//
	// example:
	//
	// {regionId: the region where the aggregation instance resides. globalViewClusterId: the ID of the aggregation instance. failedClusterIds: the ID of the cluster that failed to be added. A cluster may fail to be added because the specified cluster ID is invalid or the cluster is added across continents.}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The additional information.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
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

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) GetInfo() *string {
	return s.Info
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetInfo(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Info = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetMsg(v string) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) SetSuccess(v bool) *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddPrometheusGlobalViewByAliClusterIdsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
