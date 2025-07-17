// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAliClusterIdsToPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddAliClusterIdsToPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) *AddAliClusterIdsToPrometheusGlobalViewResponseBody
	GetData() *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData
	SetMessage(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type AddAliClusterIdsToPrometheusGlobalViewResponseBody struct {
	// The status code. The HTTP 200 status code indicates a successful request, while others indicate error conditions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the array object.
	Data *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// F7781D4A-2818-41E7-B7BB-79D809E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) GetData() *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetCode(v int32) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetData(v *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetMessage(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) SetRequestId(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddAliClusterIdsToPrometheusGlobalViewResponseBodyData struct {
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
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) GetInfo() *string {
	return s.Info
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetInfo(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddAliClusterIdsToPrometheusGlobalViewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
