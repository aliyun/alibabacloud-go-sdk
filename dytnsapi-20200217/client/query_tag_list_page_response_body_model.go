// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagListPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTagListPageResponseBody
	GetCode() *string
	SetData(v *QueryTagListPageResponseBodyData) *QueryTagListPageResponseBody
	GetData() *QueryTagListPageResponseBodyData
	SetMessage(v string) *QueryTagListPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTagListPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTagListPageResponseBody
	GetSuccess() *bool
}

type QueryTagListPageResponseBody struct {
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryTagListPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4475657-BB7E-585D-9E09-37934F096103
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagListPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagListPageResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTagListPageResponseBody) GetData() *QueryTagListPageResponseBodyData {
	return s.Data
}

func (s *QueryTagListPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTagListPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagListPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTagListPageResponseBody) SetCode(v string) *QueryTagListPageResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetData(v *QueryTagListPageResponseBodyData) *QueryTagListPageResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagListPageResponseBody) SetMessage(v string) *QueryTagListPageResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetRequestId(v string) *QueryTagListPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagListPageResponseBody) SetSuccess(v bool) *QueryTagListPageResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTagListPageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTagListPageResponseBodyData struct {
	// The page number.
	//
	// example:
	//
	// 11
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 24
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The retruned data.
	Records []*QueryTagListPageResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// The total number of returned entries.
	//
	// example:
	//
	// 32
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 91
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s QueryTagListPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTagListPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *QueryTagListPageResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryTagListPageResponseBodyData) GetRecords() []*QueryTagListPageResponseBodyDataRecords {
	return s.Records
}

func (s *QueryTagListPageResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *QueryTagListPageResponseBodyData) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *QueryTagListPageResponseBodyData) SetPageNo(v int64) *QueryTagListPageResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetPageSize(v int64) *QueryTagListPageResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetRecords(v []*QueryTagListPageResponseBodyDataRecords) *QueryTagListPageResponseBodyData {
	s.Records = v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetTotalCount(v int64) *QueryTagListPageResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) SetTotalPage(v int64) *QueryTagListPageResponseBodyData {
	s.TotalPage = &v
	return s
}

func (s *QueryTagListPageResponseBodyData) Validate() error {
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

type QueryTagListPageResponseBodyDataRecords struct {
	// The API operation that is called by the frontend.
	//
	// example:
	//
	// TwoElementsVerification
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The URL for the API documentation.
	//
	// example:
	//
	// https://help.aliyun.com/document_detail/388997.html?spm=a2c4g.388997.0.0.cf804cc7DX4vlP
	DocAddress *string `json:"DocAddress,omitempty" xml:"DocAddress,omitempty"`
	// The tag ID.
	//
	// example:
	//
	// 75
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The industry ID.
	//
	// example:
	//
	// 2
	IndustryId *int64 `json:"IndustryId,omitempty" xml:"IndustryId,omitempty"`
	// The industry name.
	//
	// example:
	//
	// Test
	IndustryName *string `json:"IndustryName,omitempty" xml:"IndustryName,omitempty"`
	// The tag description.
	//
	// example:
	//
	// for autotest new
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// Indicates whether the number is activated.
	//
	// example:
	//
	// 45
	IsOpen *int64 `json:"IsOpen,omitempty" xml:"IsOpen,omitempty"`
	// The tag name.
	//
	// example:
	//
	// Aliyun
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 	- 0: The number is hidden.
	//
	// 	- 1: The number is public.
	//
	// example:
	//
	// 1
	SaleStatusStr *string `json:"SaleStatusStr,omitempty" xml:"SaleStatusStr,omitempty"`
	// The scene ID.
	//
	// example:
	//
	// 13
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
	// The scene name.
	//
	// example:
	//
	// check
	SceneName *string `json:"SceneName,omitempty" xml:"SceneName,omitempty"`
}

func (s QueryTagListPageResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s QueryTagListPageResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *QueryTagListPageResponseBodyDataRecords) GetApiName() *string {
	return s.ApiName
}

func (s *QueryTagListPageResponseBodyDataRecords) GetCode() *string {
	return s.Code
}

func (s *QueryTagListPageResponseBodyDataRecords) GetDocAddress() *string {
	return s.DocAddress
}

func (s *QueryTagListPageResponseBodyDataRecords) GetId() *int64 {
	return s.Id
}

func (s *QueryTagListPageResponseBodyDataRecords) GetIndustryId() *int64 {
	return s.IndustryId
}

func (s *QueryTagListPageResponseBodyDataRecords) GetIndustryName() *string {
	return s.IndustryName
}

func (s *QueryTagListPageResponseBodyDataRecords) GetIntroduction() *string {
	return s.Introduction
}

func (s *QueryTagListPageResponseBodyDataRecords) GetIsOpen() *int64 {
	return s.IsOpen
}

func (s *QueryTagListPageResponseBodyDataRecords) GetName() *string {
	return s.Name
}

func (s *QueryTagListPageResponseBodyDataRecords) GetSaleStatusStr() *string {
	return s.SaleStatusStr
}

func (s *QueryTagListPageResponseBodyDataRecords) GetSceneId() *int64 {
	return s.SceneId
}

func (s *QueryTagListPageResponseBodyDataRecords) GetSceneName() *string {
	return s.SceneName
}

func (s *QueryTagListPageResponseBodyDataRecords) SetApiName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.ApiName = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetCode(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Code = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetDocAddress(v string) *QueryTagListPageResponseBodyDataRecords {
	s.DocAddress = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.Id = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIndustryId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.IndustryId = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIndustryName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.IndustryName = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIntroduction(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Introduction = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetIsOpen(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.IsOpen = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.Name = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSaleStatusStr(v string) *QueryTagListPageResponseBodyDataRecords {
	s.SaleStatusStr = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSceneId(v int64) *QueryTagListPageResponseBodyDataRecords {
	s.SceneId = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) SetSceneName(v string) *QueryTagListPageResponseBodyDataRecords {
	s.SceneName = &v
	return s
}

func (s *QueryTagListPageResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
