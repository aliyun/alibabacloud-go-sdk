// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListClustersResponseBody
	GetCode() *int32
	SetData(v *ListClustersResponseBodyData) *ListClustersResponseBody
	GetData() *ListClustersResponseBodyData
	SetMessage(v string) *ListClustersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListClustersResponseBody
	GetSuccess() *bool
}

type ListClustersResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListClustersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39938688-0BAB-5AD8-BF02-F4910FAC7589
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListClustersResponseBody) GetData() *ListClustersResponseBodyData {
	return s.Data
}

func (s *ListClustersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClustersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListClustersResponseBody) SetCode(v int32) *ListClustersResponseBody {
	s.Code = &v
	return s
}

func (s *ListClustersResponseBody) SetData(v *ListClustersResponseBodyData) *ListClustersResponseBody {
	s.Data = v
	return s
}

func (s *ListClustersResponseBody) SetMessage(v string) *ListClustersResponseBody {
	s.Message = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClustersResponseBody) SetSuccess(v bool) *ListClustersResponseBody {
	s.Success = &v
	return s
}

func (s *ListClustersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListClustersResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// -
	Records []*ListClustersResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 30
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListClustersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListClustersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListClustersResponseBodyData) GetRecords() []*ListClustersResponseBodyDataRecords {
	return s.Records
}

func (s *ListClustersResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListClustersResponseBodyData) SetPageNumber(v int32) *ListClustersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListClustersResponseBodyData) SetPageSize(v int32) *ListClustersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListClustersResponseBodyData) SetRecords(v []*ListClustersResponseBodyDataRecords) *ListClustersResponseBodyData {
	s.Records = v
	return s
}

func (s *ListClustersResponseBodyData) SetTotal(v int32) *ListClustersResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListClustersResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyDataRecords struct {
	// example:
	//
	// PREPAY
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// xxljob-c20f7ec9a78
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// xxl-job-test-1730427510169
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// example:
	//
	// scx.small.x2
	ClusterSpec *string `json:"ClusterSpec,omitempty" xml:"ClusterSpec,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2024-10-29 15:56:36
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// xxljob
	EngineType *string `json:"EngineType,omitempty" xml:"EngineType,omitempty"`
	// example:
	//
	// 2.0.0
	EngineVersion  *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty"`
	InternetDomain *string `json:"InternetDomain,omitempty" xml:"InternetDomain,omitempty"`
	// example:
	//
	// http://xxljob-b9e19e46c4e.schedulerx.mse.aliyuncs.com
	IntranetDomain *string `json:"IntranetDomain,omitempty" xml:"IntranetDomain,omitempty"`
	// example:
	//
	// 1
	ProductType  *int32  `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	SpInstanceId *string `json:"SpInstanceId,omitempty" xml:"SpInstanceId,omitempty"`
	// example:
	//
	// 1
	Status           *int32                                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags             map[string]interface{}                          `json:"Tags,omitempty" xml:"Tags,omitempty"`
	VSwitches        []*ListClustersResponseBodyDataRecordsVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" type:"Repeated"`
	VersionLifecycle *string                                         `json:"VersionLifecycle,omitempty" xml:"VersionLifecycle,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-bp1fxort6ag5h9752i305
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyDataRecords) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListClustersResponseBodyDataRecords) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyDataRecords) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersResponseBodyDataRecords) GetClusterSpec() *string {
	return s.ClusterSpec
}

func (s *ListClustersResponseBodyDataRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClustersResponseBodyDataRecords) GetEndTime() *string {
	return s.EndTime
}

func (s *ListClustersResponseBodyDataRecords) GetEngineType() *string {
	return s.EngineType
}

func (s *ListClustersResponseBodyDataRecords) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ListClustersResponseBodyDataRecords) GetInternetDomain() *string {
	return s.InternetDomain
}

func (s *ListClustersResponseBodyDataRecords) GetIntranetDomain() *string {
	return s.IntranetDomain
}

func (s *ListClustersResponseBodyDataRecords) GetProductType() *int32 {
	return s.ProductType
}

func (s *ListClustersResponseBodyDataRecords) GetSpInstanceId() *string {
	return s.SpInstanceId
}

func (s *ListClustersResponseBodyDataRecords) GetStatus() *int32 {
	return s.Status
}

func (s *ListClustersResponseBodyDataRecords) GetTags() map[string]interface{} {
	return s.Tags
}

func (s *ListClustersResponseBodyDataRecords) GetVSwitches() []*ListClustersResponseBodyDataRecordsVSwitches {
	return s.VSwitches
}

func (s *ListClustersResponseBodyDataRecords) GetVersionLifecycle() *string {
	return s.VersionLifecycle
}

func (s *ListClustersResponseBodyDataRecords) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClustersResponseBodyDataRecords) SetChargeType(v string) *ListClustersResponseBodyDataRecords {
	s.ChargeType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterId(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterName(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetClusterSpec(v string) *ListClustersResponseBodyDataRecords {
	s.ClusterSpec = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetCreateTime(v string) *ListClustersResponseBodyDataRecords {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEndTime(v string) *ListClustersResponseBodyDataRecords {
	s.EndTime = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEngineType(v string) *ListClustersResponseBodyDataRecords {
	s.EngineType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetEngineVersion(v string) *ListClustersResponseBodyDataRecords {
	s.EngineVersion = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetInternetDomain(v string) *ListClustersResponseBodyDataRecords {
	s.InternetDomain = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetIntranetDomain(v string) *ListClustersResponseBodyDataRecords {
	s.IntranetDomain = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetProductType(v int32) *ListClustersResponseBodyDataRecords {
	s.ProductType = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetSpInstanceId(v string) *ListClustersResponseBodyDataRecords {
	s.SpInstanceId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetStatus(v int32) *ListClustersResponseBodyDataRecords {
	s.Status = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetTags(v map[string]interface{}) *ListClustersResponseBodyDataRecords {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetVSwitches(v []*ListClustersResponseBodyDataRecordsVSwitches) *ListClustersResponseBodyDataRecords {
	s.VSwitches = v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetVersionLifecycle(v string) *ListClustersResponseBodyDataRecords {
	s.VersionLifecycle = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) SetVpcId(v string) *ListClustersResponseBodyDataRecords {
	s.VpcId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecords) Validate() error {
	if s.VSwitches != nil {
		for _, item := range s.VSwitches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClustersResponseBodyDataRecordsVSwitches struct {
	// example:
	//
	// vsw-8vbl54xzux86usy61r5zm
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClustersResponseBodyDataRecordsVSwitches) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyDataRecordsVSwitches) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) SetVSwitchId(v string) *ListClustersResponseBodyDataRecordsVSwitches {
	s.VSwitchId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) SetZoneId(v string) *ListClustersResponseBodyDataRecordsVSwitches {
	s.ZoneId = &v
	return s
}

func (s *ListClustersResponseBodyDataRecordsVSwitches) Validate() error {
	return dara.Validate(s)
}
