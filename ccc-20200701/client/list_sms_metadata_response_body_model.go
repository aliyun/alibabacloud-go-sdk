// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSmsMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSmsMetadataResponseBody
	GetCode() *string
	SetData(v *ListSmsMetadataResponseBodyData) *ListSmsMetadataResponseBody
	GetData() *ListSmsMetadataResponseBodyData
	SetHttpStatusCode(v int32) *ListSmsMetadataResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListSmsMetadataResponseBody
	GetMessage() *string
	SetParams(v []*string) *ListSmsMetadataResponseBody
	GetParams() []*string
	SetRequestId(v string) *ListSmsMetadataResponseBody
	GetRequestId() *string
}

type ListSmsMetadataResponseBody struct {
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListSmsMetadataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSmsMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSmsMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *ListSmsMetadataResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSmsMetadataResponseBody) GetData() *ListSmsMetadataResponseBodyData {
	return s.Data
}

func (s *ListSmsMetadataResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSmsMetadataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSmsMetadataResponseBody) GetParams() []*string {
	return s.Params
}

func (s *ListSmsMetadataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSmsMetadataResponseBody) SetCode(v string) *ListSmsMetadataResponseBody {
	s.Code = &v
	return s
}

func (s *ListSmsMetadataResponseBody) SetData(v *ListSmsMetadataResponseBodyData) *ListSmsMetadataResponseBody {
	s.Data = v
	return s
}

func (s *ListSmsMetadataResponseBody) SetHttpStatusCode(v int32) *ListSmsMetadataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSmsMetadataResponseBody) SetMessage(v string) *ListSmsMetadataResponseBody {
	s.Message = &v
	return s
}

func (s *ListSmsMetadataResponseBody) SetParams(v []*string) *ListSmsMetadataResponseBody {
	s.Params = v
	return s
}

func (s *ListSmsMetadataResponseBody) SetRequestId(v string) *ListSmsMetadataResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSmsMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSmsMetadataResponseBodyData struct {
	List []*ListSmsMetadataResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
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
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListSmsMetadataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSmsMetadataResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSmsMetadataResponseBodyData) GetList() []*ListSmsMetadataResponseBodyDataList {
	return s.List
}

func (s *ListSmsMetadataResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSmsMetadataResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSmsMetadataResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSmsMetadataResponseBodyData) SetList(v []*ListSmsMetadataResponseBodyDataList) *ListSmsMetadataResponseBodyData {
	s.List = v
	return s
}

func (s *ListSmsMetadataResponseBodyData) SetPageNumber(v int32) *ListSmsMetadataResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSmsMetadataResponseBodyData) SetPageSize(v int32) *ListSmsMetadataResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListSmsMetadataResponseBodyData) SetTotalCount(v int32) *ListSmsMetadataResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSmsMetadataResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListSmsMetadataResponseBodyDataList struct {
	// example:
	//
	// 15772400000****
	AliyunUid   *int64  `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// MISSED_CALL_NOTIFICATION
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// example:
	//
	// 5ffc1c9a-4d3d-4019-*****-73255fb01d1c
	SmsMetadataId *string `json:"SmsMetadataId,omitempty" xml:"SmsMetadataId,omitempty"`
	// example:
	//
	// SMS_468xxxx298
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
}

func (s ListSmsMetadataResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s ListSmsMetadataResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *ListSmsMetadataResponseBodyDataList) GetAliyunUid() *int64 {
	return s.AliyunUid
}

func (s *ListSmsMetadataResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *ListSmsMetadataResponseBodyDataList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSmsMetadataResponseBodyDataList) GetName() *string {
	return s.Name
}

func (s *ListSmsMetadataResponseBodyDataList) GetScenario() *string {
	return s.Scenario
}

func (s *ListSmsMetadataResponseBodyDataList) GetSignName() *string {
	return s.SignName
}

func (s *ListSmsMetadataResponseBodyDataList) GetSmsMetadataId() *string {
	return s.SmsMetadataId
}

func (s *ListSmsMetadataResponseBodyDataList) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *ListSmsMetadataResponseBodyDataList) SetAliyunUid(v int64) *ListSmsMetadataResponseBodyDataList {
	s.AliyunUid = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetDescription(v string) *ListSmsMetadataResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetInstanceId(v string) *ListSmsMetadataResponseBodyDataList {
	s.InstanceId = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetName(v string) *ListSmsMetadataResponseBodyDataList {
	s.Name = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetScenario(v string) *ListSmsMetadataResponseBodyDataList {
	s.Scenario = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetSignName(v string) *ListSmsMetadataResponseBodyDataList {
	s.SignName = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetSmsMetadataId(v string) *ListSmsMetadataResponseBodyDataList {
	s.SmsMetadataId = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) SetTemplateCode(v string) *ListSmsMetadataResponseBodyDataList {
	s.TemplateCode = &v
	return s
}

func (s *ListSmsMetadataResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
