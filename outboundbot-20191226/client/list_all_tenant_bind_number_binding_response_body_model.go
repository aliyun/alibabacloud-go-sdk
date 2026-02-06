// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllTenantBindNumberBindingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllTenantBindNumberBindingResponseBody
	GetCode() *string
	SetData(v *ListAllTenantBindNumberBindingResponseBodyData) *ListAllTenantBindNumberBindingResponseBody
	GetData() *ListAllTenantBindNumberBindingResponseBodyData
	SetHttpStatusCode(v int32) *ListAllTenantBindNumberBindingResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllTenantBindNumberBindingResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllTenantBindNumberBindingResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllTenantBindNumberBindingResponseBody
	GetSuccess() *bool
}

type ListAllTenantBindNumberBindingResponseBody struct {
	// example:
	//
	// OK
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAllTenantBindNumberBindingResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A4BEAB4B-C810-5386-B72A-1A35FF1E6B15
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllTenantBindNumberBindingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllTenantBindNumberBindingResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetData() *ListAllTenantBindNumberBindingResponseBodyData {
	return s.Data
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllTenantBindNumberBindingResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetCode(v string) *ListAllTenantBindNumberBindingResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetData(v *ListAllTenantBindNumberBindingResponseBodyData) *ListAllTenantBindNumberBindingResponseBody {
	s.Data = v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetHttpStatusCode(v int32) *ListAllTenantBindNumberBindingResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetMessage(v string) *ListAllTenantBindNumberBindingResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetRequestId(v string) *ListAllTenantBindNumberBindingResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) SetSuccess(v bool) *ListAllTenantBindNumberBindingResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAllTenantBindNumberBindingResponseBodyData struct {
	List []*ListAllTenantBindNumberBindingResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s ListAllTenantBindNumberBindingResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllTenantBindNumberBindingResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllTenantBindNumberBindingResponseBodyData) GetList() []*ListAllTenantBindNumberBindingResponseBodyDataList {
	return s.List
}

func (s *ListAllTenantBindNumberBindingResponseBodyData) SetList(v []*ListAllTenantBindNumberBindingResponseBodyDataList) *ListAllTenantBindNumberBindingResponseBodyData {
	s.List = v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyData) Validate() error {
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

type ListAllTenantBindNumberBindingResponseBodyDataList struct {
	BillingType *string `json:"BillingType,omitempty" xml:"BillingType,omitempty"`
	// example:
	//
	// b5bfb844-ce85-4779-bc8f-161fba46aa07
	BindingId        *string   `json:"BindingId,omitempty" xml:"BindingId,omitempty"`
	InstanceNameList []*string `json:"InstanceNameList,omitempty" xml:"InstanceNameList,omitempty" type:"Repeated"`
	// example:
	//
	// 15005059355
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// {\\"rateLimitPeriod\\":\\"1\\",\\"rateLimitCount\\":\\"1\\"}
	SerializedParams *string `json:"SerializedParams,omitempty" xml:"SerializedParams,omitempty"`
	TrunkName        *string `json:"TrunkName,omitempty" xml:"TrunkName,omitempty"`
}

func (s ListAllTenantBindNumberBindingResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListAllTenantBindNumberBindingResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetBillingType() *string {
	return s.BillingType
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetBindingId() *string {
	return s.BindingId
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetInstanceNameList() []*string {
	return s.InstanceNameList
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetNumber() *string {
	return s.Number
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetSerializedParams() *string {
	return s.SerializedParams
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) GetTrunkName() *string {
	return s.TrunkName
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetBillingType(v string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.BillingType = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetBindingId(v string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.BindingId = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetInstanceNameList(v []*string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.InstanceNameList = v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetNumber(v string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.Number = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetSerializedParams(v string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.SerializedParams = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) SetTrunkName(v string) *ListAllTenantBindNumberBindingResponseBodyDataList {
	s.TrunkName = &v
	return s
}

func (s *ListAllTenantBindNumberBindingResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
