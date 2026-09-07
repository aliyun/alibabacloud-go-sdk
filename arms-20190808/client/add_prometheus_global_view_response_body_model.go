// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrometheusGlobalViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddPrometheusGlobalViewResponseBody
	GetCode() *int32
	SetData(v *AddPrometheusGlobalViewResponseBodyData) *AddPrometheusGlobalViewResponseBody
	GetData() *AddPrometheusGlobalViewResponseBodyData
	SetMessage(v string) *AddPrometheusGlobalViewResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddPrometheusGlobalViewResponseBody
	GetRequestId() *string
}

type AddPrometheusGlobalViewResponseBody struct {
	// The status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response struct.
	Data *AddPrometheusGlobalViewResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned for the result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 34ED024E-9E31-434A-9E4E-D9D15C3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPrometheusGlobalViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewResponseBody) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddPrometheusGlobalViewResponseBody) GetData() *AddPrometheusGlobalViewResponseBodyData {
	return s.Data
}

func (s *AddPrometheusGlobalViewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddPrometheusGlobalViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddPrometheusGlobalViewResponseBody) SetCode(v int32) *AddPrometheusGlobalViewResponseBody {
	s.Code = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBody) SetData(v *AddPrometheusGlobalViewResponseBodyData) *AddPrometheusGlobalViewResponseBody {
	s.Data = v
	return s
}

func (s *AddPrometheusGlobalViewResponseBody) SetMessage(v string) *AddPrometheusGlobalViewResponseBody {
	s.Message = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBody) SetRequestId(v string) *AddPrometheusGlobalViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddPrometheusGlobalViewResponseBodyData struct {
	// The Info-level information.
	//
	// example:
	//
	// {regionId: 实例所属region, globalViewClusterId: 实例Id, failedInstances: 数据源JSON list中，添加失败的单个JSON的list}
	Info *AddPrometheusGlobalViewResponseBodyDataInfo `json:"Info,omitempty" xml:"Info,omitempty" type:"Struct"`
	// The additional description.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// Indicates whether the query was successful.
	//
	// - `true`: Successful.
	//
	// - `false`: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddPrometheusGlobalViewResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponseBodyData) GetInfo() *AddPrometheusGlobalViewResponseBodyDataInfo {
	return s.Info
}

func (s *AddPrometheusGlobalViewResponseBodyData) GetMsg() *string {
	return s.Msg
}

func (s *AddPrometheusGlobalViewResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetInfo(v *AddPrometheusGlobalViewResponseBodyDataInfo) *AddPrometheusGlobalViewResponseBodyData {
	s.Info = v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetMsg(v string) *AddPrometheusGlobalViewResponseBodyData {
	s.Msg = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyData) SetSuccess(v bool) *AddPrometheusGlobalViewResponseBodyData {
	s.Success = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyData) Validate() error {
	if s.Info != nil {
		if err := s.Info.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddPrometheusGlobalViewResponseBodyDataInfo struct {
	// The list of instances that failed to be added.
	//
	// example:
	//
	// [{"sourceName": "数据源名称- ArmsPrometheus","sourceType":"AlibabaPrometheus","userId":"UserID","clusterId":"ClusterId",}]
	FailedInstances *string `json:"FailedInstances,omitempty" xml:"FailedInstances,omitempty"`
	// The aggregation instance ID.
	//
	// example:
	//
	// global-v2-cn-1483223059272121-jmjjfznz
	GlobalViewClusterId *string `json:"GlobalViewClusterId,omitempty" xml:"GlobalViewClusterId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPrometheusGlobalViewResponseBodyDataInfo) String() string {
	return dara.Prettify(s)
}

func (s AddPrometheusGlobalViewResponseBodyDataInfo) GoString() string {
	return s.String()
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) GetFailedInstances() *string {
	return s.FailedInstances
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) GetGlobalViewClusterId() *string {
	return s.GlobalViewClusterId
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) SetFailedInstances(v string) *AddPrometheusGlobalViewResponseBodyDataInfo {
	s.FailedInstances = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) SetGlobalViewClusterId(v string) *AddPrometheusGlobalViewResponseBodyDataInfo {
	s.GlobalViewClusterId = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) SetRegionId(v string) *AddPrometheusGlobalViewResponseBodyDataInfo {
	s.RegionId = &v
	return s
}

func (s *AddPrometheusGlobalViewResponseBodyDataInfo) Validate() error {
	return dara.Validate(s)
}
