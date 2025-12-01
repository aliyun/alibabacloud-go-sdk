// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVulItemPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVulItemPageResponseBody
	GetCode() *string
	SetData(v []*GetVulItemPageResponseBodyData) *GetVulItemPageResponseBody
	GetData() []*GetVulItemPageResponseBodyData
	SetHttpStatusCode(v int32) *GetVulItemPageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetVulItemPageResponseBody
	GetMessage() *string
	SetPageInfo(v *GetVulItemPageResponseBodyPageInfo) *GetVulItemPageResponseBody
	GetPageInfo() *GetVulItemPageResponseBodyPageInfo
	SetRequestId(v string) *GetVulItemPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetVulItemPageResponseBody
	GetSuccess() *bool
}

type GetVulItemPageResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data []*GetVulItemPageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Pagination information.
	PageInfo *GetVulItemPageResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Request response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 02F8BBF3-2D61-5982-8911-EEB387BE3AF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetVulItemPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVulItemPageResponseBody) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVulItemPageResponseBody) GetData() []*GetVulItemPageResponseBodyData {
	return s.Data
}

func (s *GetVulItemPageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetVulItemPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVulItemPageResponseBody) GetPageInfo() *GetVulItemPageResponseBodyPageInfo {
	return s.PageInfo
}

func (s *GetVulItemPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVulItemPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetVulItemPageResponseBody) SetCode(v string) *GetVulItemPageResponseBody {
	s.Code = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetData(v []*GetVulItemPageResponseBodyData) *GetVulItemPageResponseBody {
	s.Data = v
	return s
}

func (s *GetVulItemPageResponseBody) SetHttpStatusCode(v int32) *GetVulItemPageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetMessage(v string) *GetVulItemPageResponseBody {
	s.Message = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetPageInfo(v *GetVulItemPageResponseBodyPageInfo) *GetVulItemPageResponseBody {
	s.PageInfo = v
	return s
}

func (s *GetVulItemPageResponseBody) SetRequestId(v string) *GetVulItemPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVulItemPageResponseBody) SetSuccess(v bool) *GetVulItemPageResponseBody {
	s.Success = &v
	return s
}

func (s *GetVulItemPageResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetVulItemPageResponseBodyData struct {
	// Vulnerability alias.
	//
	// example:
	//
	// RHSA-2024:4620: libndp 安全更新
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// Number of high-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 74
	AsapCount *int32 `json:"AsapCount,omitempty" xml:"AsapCount,omitempty"`
	// User ID.
	//
	// example:
	//
	// 1940494487193744
	CustomerId *string `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// Prefix for the CVE remediation advice URL.
	//
	// example:
	//
	// https://avd.aliyun.com/detail/
	CveUrlPrefix *string `json:"CveUrlPrefix,omitempty" xml:"CveUrlPrefix,omitempty"`
	// Processing status.
	//
	// example:
	//
	// y
	Dealed *string `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// Timestamp of the last discovery of the vulnerability.
	//
	// example:
	//
	// 2023-04-23 14:47:34
	FindTime *string `json:"FindTime,omitempty" xml:"FindTime,omitempty"`
	// Number of processed vulnerabilities.
	//
	// example:
	//
	// 20
	HandledCount *int32 `json:"HandledCount,omitempty" xml:"HandledCount,omitempty"`
	// Primary key ID.
	//
	// example:
	//
	// 353845
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Number of medium-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 10
	LaterCount *int32 `json:"LaterCount,omitempty" xml:"LaterCount,omitempty"`
	// Risk level
	//
	// example:
	//
	// later
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// Vulnerability name.
	//
	// example:
	//
	// oval:com.redhat.rhsa:def:20205002
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of low-priority vulnerabilities to be fixed.
	//
	// example:
	//
	// 8
	NntfCount *int32 `json:"NntfCount,omitempty" xml:"NntfCount,omitempty"`
	// CVE number.
	//
	// example:
	//
	// CVE-2019-20907
	Related *string `json:"Related,omitempty" xml:"Related,omitempty"`
	// Number of related CVE numbers.
	//
	// example:
	//
	// 20
	RelatedCveCount *int32 `json:"RelatedCveCount,omitempty" xml:"RelatedCveCount,omitempty"`
	// Vulnerability type.
	//
	// example:
	//
	// sca
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
	// Tags.
	//
	// example:
	//
	// Elevation of Privilege
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// Total number of fixed vulnerabilities.
	//
	// example:
	//
	// 50
	TotalFixCount *int64 `json:"TotalFixCount,omitempty" xml:"TotalFixCount,omitempty"`
}

func (s GetVulItemPageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVulItemPageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyData) GetAliasName() *string {
	return s.AliasName
}

func (s *GetVulItemPageResponseBodyData) GetAsapCount() *int32 {
	return s.AsapCount
}

func (s *GetVulItemPageResponseBodyData) GetCustomerId() *string {
	return s.CustomerId
}

func (s *GetVulItemPageResponseBodyData) GetCveUrlPrefix() *string {
	return s.CveUrlPrefix
}

func (s *GetVulItemPageResponseBodyData) GetDealed() *string {
	return s.Dealed
}

func (s *GetVulItemPageResponseBodyData) GetFindTime() *string {
	return s.FindTime
}

func (s *GetVulItemPageResponseBodyData) GetHandledCount() *int32 {
	return s.HandledCount
}

func (s *GetVulItemPageResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetVulItemPageResponseBodyData) GetLaterCount() *int32 {
	return s.LaterCount
}

func (s *GetVulItemPageResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *GetVulItemPageResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetVulItemPageResponseBodyData) GetNntfCount() *int32 {
	return s.NntfCount
}

func (s *GetVulItemPageResponseBodyData) GetRelated() *string {
	return s.Related
}

func (s *GetVulItemPageResponseBodyData) GetRelatedCveCount() *int32 {
	return s.RelatedCveCount
}

func (s *GetVulItemPageResponseBodyData) GetScanType() *string {
	return s.ScanType
}

func (s *GetVulItemPageResponseBodyData) GetTags() *string {
	return s.Tags
}

func (s *GetVulItemPageResponseBodyData) GetTotalFixCount() *int64 {
	return s.TotalFixCount
}

func (s *GetVulItemPageResponseBodyData) SetAliasName(v string) *GetVulItemPageResponseBodyData {
	s.AliasName = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetAsapCount(v int32) *GetVulItemPageResponseBodyData {
	s.AsapCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCustomerId(v string) *GetVulItemPageResponseBodyData {
	s.CustomerId = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetCveUrlPrefix(v string) *GetVulItemPageResponseBodyData {
	s.CveUrlPrefix = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetDealed(v string) *GetVulItemPageResponseBodyData {
	s.Dealed = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetFindTime(v string) *GetVulItemPageResponseBodyData {
	s.FindTime = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetHandledCount(v int32) *GetVulItemPageResponseBodyData {
	s.HandledCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetId(v int64) *GetVulItemPageResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLaterCount(v int32) *GetVulItemPageResponseBodyData {
	s.LaterCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetLevel(v string) *GetVulItemPageResponseBodyData {
	s.Level = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetName(v string) *GetVulItemPageResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetNntfCount(v int32) *GetVulItemPageResponseBodyData {
	s.NntfCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelated(v string) *GetVulItemPageResponseBodyData {
	s.Related = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetRelatedCveCount(v int32) *GetVulItemPageResponseBodyData {
	s.RelatedCveCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetScanType(v string) *GetVulItemPageResponseBodyData {
	s.ScanType = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTags(v string) *GetVulItemPageResponseBodyData {
	s.Tags = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) SetTotalFixCount(v int64) *GetVulItemPageResponseBodyData {
	s.TotalFixCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetVulItemPageResponseBodyPageInfo struct {
	// The current page number for pagination queries.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Number of items to display per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Total number of records in the query result.
	//
	// example:
	//
	// 163
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetVulItemPageResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s GetVulItemPageResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *GetVulItemPageResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *GetVulItemPageResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetVulItemPageResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetVulItemPageResponseBodyPageInfo) SetCurrentPage(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetPageSize(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) SetTotalCount(v int32) *GetVulItemPageResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *GetVulItemPageResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
