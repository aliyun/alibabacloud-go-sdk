// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVulAutoRepairConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListVulAutoRepairConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListVulAutoRepairConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListVulAutoRepairConfigResponseBody
	GetMessage() *string
	SetPageInfo(v *ListVulAutoRepairConfigResponseBodyPageInfo) *ListVulAutoRepairConfigResponseBody
	GetPageInfo() *ListVulAutoRepairConfigResponseBodyPageInfo
	SetRequestId(v string) *ListVulAutoRepairConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListVulAutoRepairConfigResponseBody
	GetSuccess() *bool
	SetVulAutoRepairConfigList(v []*ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) *ListVulAutoRepairConfigResponseBody
	GetVulAutoRepairConfigList() []*ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList
}

type ListVulAutoRepairConfigResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *ListVulAutoRepairConfigResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3B3F3A90-46A5-4023-A2D8-D68B14262F96
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// An array consisting of the existing configurations of vulnerabilities that can be automatically fixed.
	VulAutoRepairConfigList []*ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList `json:"VulAutoRepairConfigList,omitempty" xml:"VulAutoRepairConfigList,omitempty" type:"Repeated"`
}

func (s ListVulAutoRepairConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVulAutoRepairConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ListVulAutoRepairConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListVulAutoRepairConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListVulAutoRepairConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListVulAutoRepairConfigResponseBody) GetPageInfo() *ListVulAutoRepairConfigResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListVulAutoRepairConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVulAutoRepairConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListVulAutoRepairConfigResponseBody) GetVulAutoRepairConfigList() []*ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	return s.VulAutoRepairConfigList
}

func (s *ListVulAutoRepairConfigResponseBody) SetCode(v string) *ListVulAutoRepairConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetHttpStatusCode(v int32) *ListVulAutoRepairConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetMessage(v string) *ListVulAutoRepairConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetPageInfo(v *ListVulAutoRepairConfigResponseBodyPageInfo) *ListVulAutoRepairConfigResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetRequestId(v string) *ListVulAutoRepairConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetSuccess(v bool) *ListVulAutoRepairConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) SetVulAutoRepairConfigList(v []*ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) *ListVulAutoRepairConfigResponseBody {
	s.VulAutoRepairConfigList = v
	return s
}

func (s *ListVulAutoRepairConfigResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VulAutoRepairConfigList != nil {
		for _, item := range s.VulAutoRepairConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVulAutoRepairConfigResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **10**.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 11
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVulAutoRepairConfigResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVulAutoRepairConfigResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) SetCount(v int32) *ListVulAutoRepairConfigResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) SetCurrentPage(v int32) *ListVulAutoRepairConfigResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) SetPageSize(v int32) *ListVulAutoRepairConfigResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) SetTotalCount(v int32) *ListVulAutoRepairConfigResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList struct {
	// The alias of the vulnerability.
	//
	// example:
	//
	// RHSA-2022:0274-Important: polkit pkexec Local Privilege Escalation Vulnerability(CVE-2021-4034)
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The configuration ID of the vulnerability.
	//
	// example:
	//
	// 37338
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// anolisos:8.4:ANSA-2022:0001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The reason why the vulnerability can be automatically fixed.
	//
	// example:
	//
	// The vulnerability fix is risk-free and can be configured to automate the fix.
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) String() string {
	return dara.Prettify(s)
}

func (s ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GoString() string {
	return s.String()
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GetAliasName() *string {
	return s.AliasName
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GetId() *int64 {
	return s.Id
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GetName() *string {
	return s.Name
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GetReason() *string {
	return s.Reason
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) GetType() *string {
	return s.Type
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) SetAliasName(v string) *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	s.AliasName = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) SetId(v int64) *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	s.Id = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) SetName(v string) *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	s.Name = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) SetReason(v string) *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	s.Reason = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) SetType(v string) *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList {
	s.Type = &v
	return s
}

func (s *ListVulAutoRepairConfigResponseBodyVulAutoRepairConfigList) Validate() error {
	return dara.Validate(s)
}
