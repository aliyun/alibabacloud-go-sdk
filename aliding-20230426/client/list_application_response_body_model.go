// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody
	GetData() []*ListApplicationResponseBodyData
	SetPageNumber(v int64) *ListApplicationResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *ListApplicationResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *ListApplicationResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListApplicationResponseBody
	GetVendorType() *string
}

type ListApplicationResponseBody struct {
	// example:
	//
	// [{}]
	Data []*ListApplicationResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBody) GetData() []*ListApplicationResponseBodyData {
	return s.Data
}

func (s *ListApplicationResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListApplicationResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListApplicationResponseBody) SetData(v []*ListApplicationResponseBodyData) *ListApplicationResponseBody {
	s.Data = v
	return s
}

func (s *ListApplicationResponseBody) SetPageNumber(v int64) *ListApplicationResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListApplicationResponseBody) SetRequestId(v string) *ListApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationResponseBody) SetTotalCount(v int64) *ListApplicationResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationResponseBody) SetVendorRequestId(v string) *ListApplicationResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListApplicationResponseBody) SetVendorType(v string) *ListApplicationResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListApplicationResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationResponseBodyData struct {
	// example:
	//
	// {\\"ODIN_TOPIC_ID\\":\\"256\\"}
	AppConfig *string `json:"AppConfig,omitempty" xml:"AppConfig,omitempty"`
	// example:
	//
	// APP_XCxxx
	AppType *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	// example:
	//
	// ONLINE
	ApplicationStatus *string `json:"ApplicationStatus,omitempty" xml:"ApplicationStatus,omitempty"`
	// example:
	//
	// ding5xxx
	CorpId *string `json:"CorpId,omitempty" xml:"CorpId,omitempty"`
	// example:
	//
	// 123456
	CreatorUserId *string `json:"CreatorUserId,omitempty" xml:"CreatorUserId,omitempty"`
	// example:
	//
	// 小明创建的宜搭应用
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// appdiqiu%%#0089FF
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// y
	Inexistence *string `json:"Inexistence,omitempty" xml:"Inexistence,omitempty"`
	// example:
	//
	// app
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ding5xxx
	SubCorpId *string `json:"SubCorpId,omitempty" xml:"SubCorpId,omitempty"`
}

func (s ListApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListApplicationResponseBodyData) GetAppConfig() *string {
	return s.AppConfig
}

func (s *ListApplicationResponseBodyData) GetAppType() *string {
	return s.AppType
}

func (s *ListApplicationResponseBodyData) GetApplicationStatus() *string {
	return s.ApplicationStatus
}

func (s *ListApplicationResponseBodyData) GetCorpId() *string {
	return s.CorpId
}

func (s *ListApplicationResponseBodyData) GetCreatorUserId() *string {
	return s.CreatorUserId
}

func (s *ListApplicationResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListApplicationResponseBodyData) GetIcon() *string {
	return s.Icon
}

func (s *ListApplicationResponseBodyData) GetInexistence() *string {
	return s.Inexistence
}

func (s *ListApplicationResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListApplicationResponseBodyData) GetSubCorpId() *string {
	return s.SubCorpId
}

func (s *ListApplicationResponseBodyData) SetAppConfig(v string) *ListApplicationResponseBodyData {
	s.AppConfig = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetAppType(v string) *ListApplicationResponseBodyData {
	s.AppType = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetApplicationStatus(v string) *ListApplicationResponseBodyData {
	s.ApplicationStatus = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCorpId(v string) *ListApplicationResponseBodyData {
	s.CorpId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetCreatorUserId(v string) *ListApplicationResponseBodyData {
	s.CreatorUserId = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetDescription(v string) *ListApplicationResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetIcon(v string) *ListApplicationResponseBodyData {
	s.Icon = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetInexistence(v string) *ListApplicationResponseBodyData {
	s.Inexistence = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetName(v string) *ListApplicationResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListApplicationResponseBodyData) SetSubCorpId(v string) *ListApplicationResponseBodyData {
	s.SubCorpId = &v
	return s
}

func (s *ListApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
