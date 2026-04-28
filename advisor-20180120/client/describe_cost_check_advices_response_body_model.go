// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostCheckAdvicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCostCheckAdvicesResponseBody
	GetCode() *string
	SetData(v *DescribeCostCheckAdvicesResponseBodyData) *DescribeCostCheckAdvicesResponseBody
	GetData() *DescribeCostCheckAdvicesResponseBodyData
	SetMessage(v string) *DescribeCostCheckAdvicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCostCheckAdvicesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCostCheckAdvicesResponseBody
	GetSuccess() *string
}

type DescribeCostCheckAdvicesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeCostCheckAdvicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCostCheckAdvicesResponseBody) GetData() *DescribeCostCheckAdvicesResponseBodyData {
	return s.Data
}

func (s *DescribeCostCheckAdvicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCostCheckAdvicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostCheckAdvicesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCostCheckAdvicesResponseBody) SetCode(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetData(v *DescribeCostCheckAdvicesResponseBodyData) *DescribeCostCheckAdvicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetMessage(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetRequestId(v string) *DescribeCostCheckAdvicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) SetSuccess(v string) *DescribeCostCheckAdvicesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCostCheckAdvicesResponseBodyData struct {
	AdviceList []*DescribeCostCheckAdvicesResponseBodyDataAdviceList `json:"AdviceList,omitempty" xml:"AdviceList,omitempty" type:"Repeated"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// example:
	//
	// 4
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyData) GetAdviceList() []*DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	return s.AdviceList
}

func (s *DescribeCostCheckAdvicesResponseBodyData) GetCheckId() *string {
	return s.CheckId
}

func (s *DescribeCostCheckAdvicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCostCheckAdvicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCostCheckAdvicesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetAdviceList(v []*DescribeCostCheckAdvicesResponseBodyDataAdviceList) *DescribeCostCheckAdvicesResponseBodyData {
	s.AdviceList = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetCheckId(v string) *DescribeCostCheckAdvicesResponseBodyData {
	s.CheckId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetPageNumber(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetPageSize(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) SetTotalCount(v int32) *DescribeCostCheckAdvicesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyData) Validate() error {
	if s.AdviceList != nil {
		for _, item := range s.AdviceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostCheckAdvicesResponseBodyDataAdviceList struct {
	// example:
	//
	// 1
	AccountFolderId *string `json:"AccountFolderId,omitempty" xml:"AccountFolderId,omitempty"`
	// example:
	//
	// 1
	AccountFolderName *string `json:"AccountFolderName,omitempty" xml:"AccountFolderName,omitempty"`
	// example:
	//
	// 111******767
	AliyunId *int64 `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	// example:
	//
	// {\\"Domains\\": [{\\"Status\\": \\"success\\", \\"\\": \\"cn\\", \\"DomainName\\": \\"www.****.com\\", Region\\"Desc\\": \\"ok\\"}]}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// Email
	//
	// example:
	//
	// xxx
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 2025-03-05T02:02:00Z
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 2025-03-05T02:02:00Z
	GmtDeleted *int64 `json:"GmtDeleted,omitempty" xml:"GmtDeleted,omitempty"`
	// example:
	//
	// ecs
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// i-2ze5*****ef7d2lk63in
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// 1200***bles_df
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// 1
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 2025-02-04T16:00:00Z
	StartTime *int64                                                    `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Tags      []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	Url       *string                                                   `json:"Url,omitempty" xml:"Url,omitempty"`
	// example:
	//
	// test
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// example:
	//
	// cn-beijing-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceList) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetAccountFolderId() *string {
	return s.AccountFolderId
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetAccountFolderName() *string {
	return s.AccountFolderName
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetAliyunId() *int64 {
	return s.AliyunId
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetContent() *string {
	return s.Content
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetEmail() *string {
	return s.Email
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetGmtDeleted() *int64 {
	return s.GmtDeleted
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetProduct() *string {
	return s.Product
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetResourceName() *string {
	return s.ResourceName
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetSeverity() *string {
	return s.Severity
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetTags() []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags {
	return s.Tags
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetUrl() *string {
	return s.Url
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetUserName() *string {
	return s.UserName
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAccountFolderId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AccountFolderId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAccountFolderName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AccountFolderName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetAliyunId(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.AliyunId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetContent(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Content = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetEmail(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Email = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetEndTime(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.EndTime = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetGmtDeleted(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.GmtDeleted = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetProduct(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Product = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetRegionId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.RegionId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetResourceId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ResourceId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetResourceName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ResourceName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetSeverity(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Severity = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetStartTime(v int64) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.StartTime = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetTags(v []*DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Tags = v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetUrl(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.Url = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetUserName(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.UserName = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) SetZoneId(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceList {
	s.ZoneId = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceList) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCostCheckAdvicesResponseBodyDataAdviceListTags struct {
	// example:
	//
	// autoTest-7
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// basic
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) GoString() string {
	return s.String()
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) SetTagKey(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) SetTagValue(v string) *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags {
	s.TagValue = &v
	return s
}

func (s *DescribeCostCheckAdvicesResponseBodyDataAdviceListTags) Validate() error {
	return dara.Validate(s)
}
