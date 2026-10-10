// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFunctionMetasResponseBody
	GetCode() *string
	SetData(v *ListFunctionMetasResponseBodyData) *ListFunctionMetasResponseBody
	GetData() *ListFunctionMetasResponseBodyData
	SetHttpStatusCode(v int32) *ListFunctionMetasResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFunctionMetasResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListFunctionMetasResponseBody
	GetRequestId() *string
}

type ListFunctionMetasResponseBody struct {
	// example:
	//
	// OK
	Code *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFunctionMetasResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 26D1277F-55BF-453E-A6B2-02B7E6F97699
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFunctionMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionMetasResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionMetasResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFunctionMetasResponseBody) GetData() *ListFunctionMetasResponseBodyData {
	return s.Data
}

func (s *ListFunctionMetasResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFunctionMetasResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFunctionMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFunctionMetasResponseBody) SetCode(v string) *ListFunctionMetasResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionMetasResponseBody) SetData(v *ListFunctionMetasResponseBodyData) *ListFunctionMetasResponseBody {
	s.Data = v
	return s
}

func (s *ListFunctionMetasResponseBody) SetHttpStatusCode(v int32) *ListFunctionMetasResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFunctionMetasResponseBody) SetMessage(v string) *ListFunctionMetasResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionMetasResponseBody) SetRequestId(v string) *ListFunctionMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionMetasResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFunctionMetasResponseBodyData struct {
	List []*ListFunctionMetasResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionMetasResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionMetasResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFunctionMetasResponseBodyData) GetList() []*ListFunctionMetasResponseBodyDataList {
	return s.List
}

func (s *ListFunctionMetasResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFunctionMetasResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFunctionMetasResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFunctionMetasResponseBodyData) SetList(v []*ListFunctionMetasResponseBodyDataList) *ListFunctionMetasResponseBodyData {
	s.List = v
	return s
}

func (s *ListFunctionMetasResponseBodyData) SetPageNumber(v int32) *ListFunctionMetasResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionMetasResponseBodyData) SetPageSize(v int32) *ListFunctionMetasResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFunctionMetasResponseBodyData) SetTotalCount(v int32) *ListFunctionMetasResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFunctionMetasResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFunctionMetasResponseBodyDataList struct {
	// example:
	//
	// 15772400000****
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// example:
	//
	// 王先生
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// cn-shanghai
	FailoverRegion *string `json:"FailoverRegion,omitempty" xml:"FailoverRegion,omitempty"`
	// example:
	//
	// 5
	FailoverRegionWeight *float64 `json:"FailoverRegionWeight,omitempty" xml:"FailoverRegionWeight,omitempty"`
	// example:
	//
	// 4bbcd898-xxxxx-47e5-85bb-718166
	FunctionMetaId *string `json:"FunctionMetaId,omitempty" xml:"FunctionMetaId,omitempty"`
	// example:
	//
	// sql_hra
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// example:
	//
	// http://xxxx
	HttpTriggerUrl *string `json:"HttpTriggerUrl,omitempty" xml:"HttpTriggerUrl,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	Region *int32 `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// logical
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// url_detection_pro
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ListFunctionMetasResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionMetasResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFunctionMetasResponseBodyDataList) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *ListFunctionMetasResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListFunctionMetasResponseBodyDataList) GetFailoverRegion() *string {
	return s.FailoverRegion
}

func (s *ListFunctionMetasResponseBodyDataList) GetFailoverRegionWeight() *float64 {
	return s.FailoverRegionWeight
}

func (s *ListFunctionMetasResponseBodyDataList) GetFunctionMetaId() *string {
	return s.FunctionMetaId
}

func (s *ListFunctionMetasResponseBodyDataList) GetFunctionName() *string {
	return s.FunctionName
}

func (s *ListFunctionMetasResponseBodyDataList) GetHttpTriggerUrl() *string {
	return s.HttpTriggerUrl
}

func (s *ListFunctionMetasResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFunctionMetasResponseBodyDataList) GetRegion() *int32 {
	return s.Region
}

func (s *ListFunctionMetasResponseBodyDataList) GetRole() *string {
	return s.Role
}

func (s *ListFunctionMetasResponseBodyDataList) GetService() *string {
	return s.Service
}

func (s *ListFunctionMetasResponseBodyDataList) SetAliyunUid(v string) *ListFunctionMetasResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetDescription(v string) *ListFunctionMetasResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetFailoverRegion(v string) *ListFunctionMetasResponseBodyDataList {
	s.FailoverRegion = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetFailoverRegionWeight(v float64) *ListFunctionMetasResponseBodyDataList {
	s.FailoverRegionWeight = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetFunctionMetaId(v string) *ListFunctionMetasResponseBodyDataList {
	s.FunctionMetaId = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetFunctionName(v string) *ListFunctionMetasResponseBodyDataList {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetHttpTriggerUrl(v string) *ListFunctionMetasResponseBodyDataList {
	s.HttpTriggerUrl = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetInstanceId(v string) *ListFunctionMetasResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetRegion(v int32) *ListFunctionMetasResponseBodyDataList {
	s.Region = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetRole(v string) *ListFunctionMetasResponseBodyDataList {
	s.Role = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) SetService(v string) *ListFunctionMetasResponseBodyDataList {
	s.Service = &v
	return s
}

func (s *ListFunctionMetasResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
