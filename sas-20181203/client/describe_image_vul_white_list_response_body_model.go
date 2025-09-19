// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageVulWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeImageVulWhiteListResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeImageVulWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetImageVulWhitelist(v []*DescribeImageVulWhiteListResponseBodyImageVulWhitelist) *DescribeImageVulWhiteListResponseBody
	GetImageVulWhitelist() []*DescribeImageVulWhiteListResponseBodyImageVulWhitelist
	SetMessage(v string) *DescribeImageVulWhiteListResponseBody
	GetMessage() *string
	SetPageInfo(v *DescribeImageVulWhiteListResponseBodyPageInfo) *DescribeImageVulWhiteListResponseBody
	GetPageInfo() *DescribeImageVulWhiteListResponseBodyPageInfo
	SetRequestId(v string) *DescribeImageVulWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeImageVulWhiteListResponseBody
	GetSuccess() *bool
	SetTimeCost(v int64) *DescribeImageVulWhiteListResponseBody
	GetTimeCost() *int64
}

type DescribeImageVulWhiteListResponseBody struct {
	// The status code returned. A value of **200*	- indicates that the request was successful. Other values indicate that the request failed. You can identify the cause of the failure based on the value of this parameter.
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
	// The information about the whitelist of image vulnerabilities.
	ImageVulWhitelist []*DescribeImageVulWhiteListResponseBodyImageVulWhitelist `json:"ImageVulWhitelist,omitempty" xml:"ImageVulWhitelist,omitempty" type:"Repeated"`
	// The message returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *DescribeImageVulWhiteListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 911025D0-3D1E-5213-A18A-37EA0C92****
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
	// The amount of time that was consumed to process the request. Unit: milliseconds.
	//
	// example:
	//
	// 1
	TimeCost *int64 `json:"TimeCost,omitempty" xml:"TimeCost,omitempty"`
}

func (s DescribeImageVulWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageVulWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeImageVulWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeImageVulWhiteListResponseBody) GetImageVulWhitelist() []*DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	return s.ImageVulWhitelist
}

func (s *DescribeImageVulWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeImageVulWhiteListResponseBody) GetPageInfo() *DescribeImageVulWhiteListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeImageVulWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageVulWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeImageVulWhiteListResponseBody) GetTimeCost() *int64 {
	return s.TimeCost
}

func (s *DescribeImageVulWhiteListResponseBody) SetCode(v string) *DescribeImageVulWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetHttpStatusCode(v int32) *DescribeImageVulWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetImageVulWhitelist(v []*DescribeImageVulWhiteListResponseBodyImageVulWhitelist) *DescribeImageVulWhiteListResponseBody {
	s.ImageVulWhitelist = v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetMessage(v string) *DescribeImageVulWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetPageInfo(v *DescribeImageVulWhiteListResponseBodyPageInfo) *DescribeImageVulWhiteListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetRequestId(v string) *DescribeImageVulWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetSuccess(v bool) *DescribeImageVulWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) SetTimeCost(v int64) *DescribeImageVulWhiteListResponseBody {
	s.TimeCost = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageVulWhiteListResponseBodyImageVulWhitelist struct {
	// The alias of the vulnerability that is specified in Common Vulnerabilities and Exposures (CVE).
	//
	// example:
	//
	// CVE-2019-19906:in_sasl_add_string
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The primary key ID of the vulnerability.
	//
	// example:
	//
	// 34032043
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the vulnerability.
	//
	// example:
	//
	// scan:AVD-2022-953356
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The reason why the vulnerability is added to the whitelist.
	//
	// example:
	//
	// already config in another way
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The object on which the query is performed. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **type**: the object type. The value is fixed to repo.
	//
	// 	- **target**: the object content. The value is in the Namespace/Image repository format.
	//
	// example:
	//
	// {\\"type\\":\\"repo\\",\\"target\\":[\\"sas_test/script_0209\\",\\"sas_test/script\\"]}
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: system vulnerability
	//
	// 	- **sca**: application vulnerability
	//
	// example:
	//
	// sca
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageVulWhiteListResponseBodyImageVulWhitelist) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GoString() string {
	return s.String()
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetId() *int64 {
	return s.Id
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetName() *string {
	return s.Name
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetReason() *string {
	return s.Reason
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetTarget() *string {
	return s.Target
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) GetType() *string {
	return s.Type
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetAliasName(v string) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.AliasName = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetId(v int64) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.Id = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetName(v string) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.Name = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetReason(v string) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.Reason = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetTarget(v string) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.Target = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) SetType(v string) *DescribeImageVulWhiteListResponseBodyImageVulWhitelist {
	s.Type = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyImageVulWhitelist) Validate() error {
	return dara.Validate(s)
}

type DescribeImageVulWhiteListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageVulWhiteListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageVulWhiteListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) SetCount(v int32) *DescribeImageVulWhiteListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeImageVulWhiteListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) SetPageSize(v int32) *DescribeImageVulWhiteListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeImageVulWhiteListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageVulWhiteListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
