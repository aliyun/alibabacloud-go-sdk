// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsRealNumberCallConnectionRateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetCode() *string
	SetMessage(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetMessage() *string
	SetModel(v *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetModel() *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel
	SetRequestId(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryVmsRealNumberCallConnectionRateInfoResponseBody
	GetSuccess() *bool
}

type QueryVmsRealNumberCallConnectionRateInfoResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 95B86652-B234-5387-BAC6-E441FR49399F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetModel() *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	return s.Model
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetAccessDeniedDetail(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetCode(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetMessage(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetModel(v *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.Model = v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetRequestId(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) SetSuccess(v bool) *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel struct {
	// 接通率
	//
	// example:
	//
	// 6.71672
	CallConnectionRate *float64 `json:"CallConnectionRate,omitempty" xml:"CallConnectionRate,omitempty"`
	// 接通数
	//
	// example:
	//
	// 52
	CompleteCount *int64 `json:"CompleteCount,omitempty" xml:"CompleteCount,omitempty"`
	// example:
	//
	// 示例值示例值
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// 请求通话数
	//
	// example:
	//
	// 49
	RequestCount *int64 `json:"RequestCount,omitempty" xml:"RequestCount,omitempty"`
	// example:
	//
	// 示例值
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// EcsInstance, EcsDisk, EcsImage, EcsSnapshot, EcsSecurityGroup, EcsEip, EcsVpc, EcsVRouter, EcsVSwitch, EcsVRouteTable, EcsAuthImage, EcsAll, SlbLoadbalancer, SlbVm, RdsInstance, RdsAllInstance, KvsInstance, OcsInstance, OdpsInstance
	//
	// example:
	//
	// 示例值
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// 响铃数
	//
	// example:
	//
	// 3
	RingingCount *int64 `json:"RingingCount,omitempty" xml:"RingingCount,omitempty"`
	// 响铃率
	//
	// example:
	//
	// 25.4222
	RingingRate *float64 `json:"RingingRate,omitempty" xml:"RingingRate,omitempty"`
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GoString() string {
	return s.String()
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetCallConnectionRate() *float64 {
	return s.CallConnectionRate
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetCompleteCount() *int64 {
	return s.CompleteCount
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetRequestCount() *int64 {
	return s.RequestCount
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetResourceType() *string {
	return s.ResourceType
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetRingingCount() *int64 {
	return s.RingingCount
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) GetRingingRate() *float64 {
	return s.RingingRate
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetCallConnectionRate(v float64) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.CallConnectionRate = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetCompleteCount(v int64) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.CompleteCount = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetRegionId(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.RegionId = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetRequestCount(v int64) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.RequestCount = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetResourceId(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.ResourceId = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetResourceType(v string) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.ResourceType = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetRingingCount(v int64) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.RingingCount = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) SetRingingRate(v float64) *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel {
	s.RingingRate = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
