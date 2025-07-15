// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsApplicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListFlashSmsApplicationsResponseBody
	GetCode() *string
	SetData(v *ListFlashSmsApplicationsResponseBodyData) *ListFlashSmsApplicationsResponseBody
	GetData() *ListFlashSmsApplicationsResponseBodyData
	SetHttpStatusCode(v int32) *ListFlashSmsApplicationsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListFlashSmsApplicationsResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListFlashSmsApplicationsResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListFlashSmsApplicationsResponseBody
	GetRequestId() *string
}

type ListFlashSmsApplicationsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListFlashSmsApplicationsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFlashSmsApplicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsApplicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFlashSmsApplicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListFlashSmsApplicationsResponseBody) GetData() *ListFlashSmsApplicationsResponseBodyData {
	return s.Data
}

func (s *ListFlashSmsApplicationsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListFlashSmsApplicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListFlashSmsApplicationsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListFlashSmsApplicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFlashSmsApplicationsResponseBody) SetCode(v string) *ListFlashSmsApplicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) SetData(v *ListFlashSmsApplicationsResponseBodyData) *ListFlashSmsApplicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) SetHttpStatusCode(v int32) *ListFlashSmsApplicationsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) SetMessage(v string) *ListFlashSmsApplicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) SetParams(v []*string) *ListFlashSmsApplicationsResponseBody {
	s.Params = v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) SetRequestId(v string) *ListFlashSmsApplicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFlashSmsApplicationsResponseBodyData struct {
	List []*ListFlashSmsApplicationsResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFlashSmsApplicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsApplicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFlashSmsApplicationsResponseBodyData) GetList() []*ListFlashSmsApplicationsResponseBodyDataList {
	return s.List
}

func (s *ListFlashSmsApplicationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsApplicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsApplicationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFlashSmsApplicationsResponseBodyData) SetList(v []*ListFlashSmsApplicationsResponseBodyDataList) *ListFlashSmsApplicationsResponseBodyData {
	s.List = v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyData) SetPageNumber(v int32) *ListFlashSmsApplicationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyData) SetPageSize(v int32) *ListFlashSmsApplicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyData) SetTotalCount(v int32) *ListFlashSmsApplicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListFlashSmsApplicationsResponseBodyDataList struct {
	// example:
	//
	// 71b396fa-***********-bd80e070b7c0
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// Uincall
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
	// example:
	//
	// {"user":"600******_dev","pwd":"85abf3**********f494e","account":"6004******"}
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFlashSmsApplicationsResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsApplicationsResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) GetProviderId() *string {
	return s.ProviderId
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) GetValue() *string {
	return s.Value
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) SetApplicationId(v string) *ListFlashSmsApplicationsResponseBodyDataList {
	s.ApplicationId = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) SetInstanceId(v string) *ListFlashSmsApplicationsResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) SetName(v string) *ListFlashSmsApplicationsResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) SetProviderId(v string) *ListFlashSmsApplicationsResponseBodyDataList {
	s.ProviderId = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) SetValue(v string) *ListFlashSmsApplicationsResponseBodyDataList {
	s.Value = &v
	return s
}

func (s *ListFlashSmsApplicationsResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
