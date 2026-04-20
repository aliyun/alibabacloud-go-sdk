// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallCenterProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListCallCenterProvidersResponseBody
	GetCode() *string
	SetData(v *ListCallCenterProvidersResponseBodyData) *ListCallCenterProvidersResponseBody
	GetData() *ListCallCenterProvidersResponseBodyData
	SetHttpStatusCode(v int32) *ListCallCenterProvidersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListCallCenterProvidersResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListCallCenterProvidersResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListCallCenterProvidersResponseBody
	GetRequestId() *string
}

type ListCallCenterProvidersResponseBody struct {
	// example:
	//
	// OK
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListCallCenterProvidersResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 75BAAB9B-40B2-5FF5-A59A-7BCF8154C6EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCallCenterProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCallCenterProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ListCallCenterProvidersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListCallCenterProvidersResponseBody) GetData() *ListCallCenterProvidersResponseBodyData {
	return s.Data
}

func (s *ListCallCenterProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListCallCenterProvidersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCallCenterProvidersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListCallCenterProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCallCenterProvidersResponseBody) SetCode(v string) *ListCallCenterProvidersResponseBody {
	s.Code = &v
	return s
}

func (s *ListCallCenterProvidersResponseBody) SetData(v *ListCallCenterProvidersResponseBodyData) *ListCallCenterProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ListCallCenterProvidersResponseBody) SetHttpStatusCode(v int32) *ListCallCenterProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListCallCenterProvidersResponseBody) SetMessage(v string) *ListCallCenterProvidersResponseBody {
	s.Message = &v
	return s
}

func (s *ListCallCenterProvidersResponseBody) SetParams(v []*string) *ListCallCenterProvidersResponseBody {
	s.Params = v
	return s
}

func (s *ListCallCenterProvidersResponseBody) SetRequestId(v string) *ListCallCenterProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCallCenterProvidersResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCallCenterProvidersResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize  *int32                                              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Providers []*ListCallCenterProvidersResponseBodyDataProviders `json:"Providers,omitempty" xml:"Providers,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCallCenterProvidersResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCallCenterProvidersResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCallCenterProvidersResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCallCenterProvidersResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCallCenterProvidersResponseBodyData) GetProviders() []*ListCallCenterProvidersResponseBodyDataProviders {
	return s.Providers
}

func (s *ListCallCenterProvidersResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCallCenterProvidersResponseBodyData) SetPageNumber(v int32) *ListCallCenterProvidersResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyData) SetPageSize(v int32) *ListCallCenterProvidersResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyData) SetProviders(v []*ListCallCenterProvidersResponseBodyDataProviders) *ListCallCenterProvidersResponseBodyData {
	s.Providers = v
	return s
}

func (s *ListCallCenterProvidersResponseBodyData) SetTotalCount(v int32) *ListCallCenterProvidersResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyData) Validate() error {
	if s.Providers != nil {
		for _, item := range s.Providers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCallCenterProvidersResponseBodyDataProviders struct {
	// example:
	//
	// 153*********
	Destination *string `json:"Destination,omitempty" xml:"Destination,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// key1=value1
	Extras *string `json:"Extras,omitempty" xml:"Extras,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 010****
	Originator *string `json:"Originator,omitempty" xml:"Originator,omitempty"`
	// example:
	//
	// xxxxxxxxx
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// CCC
	ProviderType *string `json:"ProviderType,omitempty" xml:"ProviderType,omitempty"`
	// example:
	//
	// -
	ReferTo *string `json:"ReferTo,omitempty" xml:"ReferTo,omitempty"`
	// example:
	//
	// trunk-xxx
	TrunkId *string `json:"TrunkId,omitempty" xml:"TrunkId,omitempty"`
}

func (s ListCallCenterProvidersResponseBodyDataProviders) String() string {
	return dara.Prettify(s)
}

func (s ListCallCenterProvidersResponseBodyDataProviders) GoString() string {
	return s.String()
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetDestination() *string {
	return s.Destination
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetExtras() *string {
	return s.Extras
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetName() *string {
	return s.Name
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetOriginator() *string {
	return s.Originator
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetReferTo() *string {
	return s.ReferTo
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) GetTrunkId() *string {
	return s.TrunkId
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetDestination(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.Destination = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetDisplayName(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.DisplayName = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetExtras(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.Extras = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetInstanceId(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.InstanceId = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetName(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.Name = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetOriginator(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.Originator = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetProviderId(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.ProviderId = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetProviderType(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.ProviderType = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetReferTo(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.ReferTo = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) SetTrunkId(v string) *ListCallCenterProvidersResponseBodyDataProviders {
	s.TrunkId = &v
	return s
}

func (s *ListCallCenterProvidersResponseBodyDataProviders) Validate() error {
	return dara.Validate(s)
}
