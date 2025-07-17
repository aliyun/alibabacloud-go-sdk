// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveAliClusterIdsFromPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody
	GetData() *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData
	SetMessage(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned struct.
	Data *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
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

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GetData() *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetCode(v int32) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetData(v *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetMessage(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) SetRequestId(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBody) Validate() error {
	return dara.Validate(s)
}

type RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData struct {
	// The Info-level information.
	//
	// example:
	//
	// {regionId: the region where the global aggregation instance resides. globalViewClusterId: the ID of the global aggregation instance. failedClusterIds: the IDs of the clusters that failed to be added. A cluster may fail to be added if the specified cluster ID is invalid.}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The additional information.
	//
	// example:
	//
	// OK
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

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) GetInfo() *string {
	return s.Info
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetInfo(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Info = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetMsg(v string) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *RemoveAliClusterIdsFromPrometheusGlobalViewResponseBodyData) Validate() error {
	return dara.Validate(s)
}
