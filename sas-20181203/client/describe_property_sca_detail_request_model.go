// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePropertyScaDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBiz(v string) *DescribePropertyScaDetailRequest
	GetBiz() *string
	SetBizType(v string) *DescribePropertyScaDetailRequest
	GetBizType() *string
	SetCurrentPage(v int32) *DescribePropertyScaDetailRequest
	GetCurrentPage() *int32
	SetLang(v string) *DescribePropertyScaDetailRequest
	GetLang() *string
	SetName(v int64) *DescribePropertyScaDetailRequest
	GetName() *int64
	SetNextToken(v string) *DescribePropertyScaDetailRequest
	GetNextToken() *string
	SetPageSize(v int32) *DescribePropertyScaDetailRequest
	GetPageSize() *int32
	SetPid(v string) *DescribePropertyScaDetailRequest
	GetPid() *string
	SetPort(v string) *DescribePropertyScaDetailRequest
	GetPort() *string
	SetProcessStartedEnd(v int64) *DescribePropertyScaDetailRequest
	GetProcessStartedEnd() *int64
	SetProcessStartedStart(v int64) *DescribePropertyScaDetailRequest
	GetProcessStartedStart() *int64
	SetRemark(v string) *DescribePropertyScaDetailRequest
	GetRemark() *string
	SetScaName(v string) *DescribePropertyScaDetailRequest
	GetScaName() *string
	SetScaNamePattern(v string) *DescribePropertyScaDetailRequest
	GetScaNamePattern() *string
	SetScaVersion(v string) *DescribePropertyScaDetailRequest
	GetScaVersion() *string
	SetSearchCriteriaList(v []*DescribePropertyScaDetailRequestSearchCriteriaList) *DescribePropertyScaDetailRequest
	GetSearchCriteriaList() []*DescribePropertyScaDetailRequestSearchCriteriaList
	SetSearchInfo(v string) *DescribePropertyScaDetailRequest
	GetSearchInfo() *string
	SetSearchInfoSub(v string) *DescribePropertyScaDetailRequest
	GetSearchInfoSub() *string
	SetSearchItem(v string) *DescribePropertyScaDetailRequest
	GetSearchItem() *string
	SetSearchItemSub(v string) *DescribePropertyScaDetailRequest
	GetSearchItemSub() *string
	SetUseNextToken(v bool) *DescribePropertyScaDetailRequest
	GetUseNextToken() *bool
	SetUser(v string) *DescribePropertyScaDetailRequest
	GetUser() *string
	SetUuid(v string) *DescribePropertyScaDetailRequest
	GetUuid() *string
}

type DescribePropertyScaDetailRequest struct {
	// The type of Asset Fingerprints to query. Default value: **sca**. Valid values:
	//
	// - **sca**: middleware
	//
	// - **sca_database**: database
	//
	// - **sca_web**: web service
	//
	// > If this parameter is not set, the default value **sca*	- is used, which queries Asset Fingerprints information of the middleware type.
	//
	// example:
	//
	// sca
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The type of middleware, database, or web service to query. Valid values:
	//
	// - **system_service**: system service
	//
	// - **software_library**: software library
	//
	// - **docker_component**: container component
	//
	// - **database**: database
	//
	// - **web_container**: web container
	//
	// - **jar**: JAR package
	//
	// - **web_framework**: web framework.
	//
	// example:
	//
	// system_service
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The page number of the page to return in the query results. Default value: **1**, which indicates that the results are displayed starting from page 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the middleware, database, or web service.
	//
	// > This parameter is deprecated. You do not need to specify this parameter.
	//
	// example:
	//
	// 1
	Name *int64 `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that marks the current position from which to start reading. Leave this parameter empty to start reading from the beginning.
	//
	// > Do not specify this parameter for the first call. The response includes the NextToken value for the second call. Each subsequent response contains the NextToken value for the next call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6k+AtdhNE3kgQEK36GujZ5on+tWdc+4WoaoMP/kUNxxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Sets the number of Asset Fingerprints entries per page in a paged query. Default value: **10**, which indicates that 10 Asset Fingerprints entries are displayed per page.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process ID.
	//
	// example:
	//
	// 756
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The port on which the process listens.
	//
	// example:
	//
	// 68
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The end of the time range for querying the process start timestamp. Unit: seconds.
	//
	// example:
	//
	// 1641110965
	ProcessStartedEnd *int64 `json:"ProcessStartedEnd,omitempty" xml:"ProcessStartedEnd,omitempty"`
	// The start of the time range for querying the process start timestamp. Unit: seconds.
	//
	// example:
	//
	// 1641024565
	ProcessStartedStart *int64 `json:"ProcessStartedStart,omitempty" xml:"ProcessStartedStart,omitempty"`
	// The search condition (server name or IP address).
	//
	// > Fuzzy match is supported.
	//
	// example:
	//
	// 192.168
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the Asset Fingerprints entry to query.
	//
	// example:
	//
	// openssl
	ScaName *string `json:"ScaName,omitempty" xml:"ScaName,omitempty"`
	// The process name.
	//
	// example:
	//
	// open
	ScaNamePattern *string `json:"ScaNamePattern,omitempty" xml:"ScaNamePattern,omitempty"`
	// The version of the middleware, database, or web service.
	//
	// example:
	//
	// 1.0.2k
	ScaVersion *string `json:"ScaVersion,omitempty" xml:"ScaVersion,omitempty"`
	// The list of search criteria.
	SearchCriteriaList []*DescribePropertyScaDetailRequestSearchCriteriaList `json:"SearchCriteriaList,omitempty" xml:"SearchCriteriaList,omitempty" type:"Repeated"`
	// The content to query. Depending on the value of **SearchItem**, you need to enter different query content:
	//
	// - If **SearchItem*	- is set to **name**, enter the name of the asset fingerprint as the query condition.
	//
	// - If **SearchItem*	- is set to **type**, select the type of asset fingerprint to query. Valid values:
	//
	//     - **system_service**: system service
	//
	//     - **software_library**: software library
	//
	//     - **docker_component**: container component
	//
	//     - **database**: database
	//
	//     - **web_container**: web container
	//
	//     - **jar**: JAR package
	//
	//     - **web_framework**: web framework
	//
	// > The **SearchItem*	- and **SearchInfo*	- parameters must be used together. Setting only one of them has no effect. By setting both parameters, you can view all data for asset fingerprints of a specified name or type.
	//
	// example:
	//
	// openssl
	SearchInfo *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	// The content of the sub-query condition. Depending on the value of **SearchItemSub**, you need to enter different query content:
	//
	// - If **SearchItemSub*	- is set to **port**, enter the port as the sub-query condition.
	//
	// - If **SearchItemSub*	- is set to **pid**, enter the process ID as the sub-query condition.
	//
	// - If **SearchItemSub*	- is set to **version**, enter the middleware, database, or web service version as the sub-query condition.
	//
	// - If **SearchItemSub*	- is set to **user**, enter the username as the sub-query condition.
	//
	// > Sub-query conditions help you search for the data list of a specific middleware, database, or web service.
	//
	// example:
	//
	// 1.0.2k
	SearchInfoSub *string `json:"SearchInfoSub,omitempty" xml:"SearchInfoSub,omitempty"`
	// The type of query condition. Valid values:
	//
	// - **name**: the name of the middleware, database, or web service.
	//
	// - **type**: the type of the middleware, database, or web service.
	//
	// > The **SearchItem*	- and **SearchInfo*	- parameters must be used together. Setting only one of them has no effect. By setting both parameters, you can view all data for asset fingerprints of a specified name or type.
	//
	// example:
	//
	// name
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The type of sub-query condition. Valid values:
	//
	// - **port**: port
	//
	// - **pid**: process ID
	//
	// - **version**: version
	//
	// - **user**: user.
	//
	// example:
	//
	// version
	SearchItemSub *string `json:"SearchItemSub,omitempty" xml:"SearchItemSub,omitempty"`
	// Specifies whether to use the NextToken method to retrieve asset list data. If this parameter is used, TotalCount is no longer returned. Valid values:
	//
	// - **true**: Use the NextToken method.
	//
	// - **false**: Do not use the NextToken method.
	//
	// example:
	//
	// true
	UseNextToken *bool `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The user that runs the process.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server on which the middleware, database, or web service is deployed.
	//
	// example:
	//
	// uuid-02ebabe7-1c19-ab****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribePropertyScaDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailRequest) GetBiz() *string {
	return s.Biz
}

func (s *DescribePropertyScaDetailRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribePropertyScaDetailRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribePropertyScaDetailRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePropertyScaDetailRequest) GetName() *int64 {
	return s.Name
}

func (s *DescribePropertyScaDetailRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePropertyScaDetailRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePropertyScaDetailRequest) GetPid() *string {
	return s.Pid
}

func (s *DescribePropertyScaDetailRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePropertyScaDetailRequest) GetProcessStartedEnd() *int64 {
	return s.ProcessStartedEnd
}

func (s *DescribePropertyScaDetailRequest) GetProcessStartedStart() *int64 {
	return s.ProcessStartedStart
}

func (s *DescribePropertyScaDetailRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribePropertyScaDetailRequest) GetScaName() *string {
	return s.ScaName
}

func (s *DescribePropertyScaDetailRequest) GetScaNamePattern() *string {
	return s.ScaNamePattern
}

func (s *DescribePropertyScaDetailRequest) GetScaVersion() *string {
	return s.ScaVersion
}

func (s *DescribePropertyScaDetailRequest) GetSearchCriteriaList() []*DescribePropertyScaDetailRequestSearchCriteriaList {
	return s.SearchCriteriaList
}

func (s *DescribePropertyScaDetailRequest) GetSearchInfo() *string {
	return s.SearchInfo
}

func (s *DescribePropertyScaDetailRequest) GetSearchInfoSub() *string {
	return s.SearchInfoSub
}

func (s *DescribePropertyScaDetailRequest) GetSearchItem() *string {
	return s.SearchItem
}

func (s *DescribePropertyScaDetailRequest) GetSearchItemSub() *string {
	return s.SearchItemSub
}

func (s *DescribePropertyScaDetailRequest) GetUseNextToken() *bool {
	return s.UseNextToken
}

func (s *DescribePropertyScaDetailRequest) GetUser() *string {
	return s.User
}

func (s *DescribePropertyScaDetailRequest) GetUuid() *string {
	return s.Uuid
}

func (s *DescribePropertyScaDetailRequest) SetBiz(v string) *DescribePropertyScaDetailRequest {
	s.Biz = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetBizType(v string) *DescribePropertyScaDetailRequest {
	s.BizType = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetCurrentPage(v int32) *DescribePropertyScaDetailRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetLang(v string) *DescribePropertyScaDetailRequest {
	s.Lang = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetName(v int64) *DescribePropertyScaDetailRequest {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetNextToken(v string) *DescribePropertyScaDetailRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetPageSize(v int32) *DescribePropertyScaDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetPid(v string) *DescribePropertyScaDetailRequest {
	s.Pid = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetPort(v string) *DescribePropertyScaDetailRequest {
	s.Port = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetProcessStartedEnd(v int64) *DescribePropertyScaDetailRequest {
	s.ProcessStartedEnd = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetProcessStartedStart(v int64) *DescribePropertyScaDetailRequest {
	s.ProcessStartedStart = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetRemark(v string) *DescribePropertyScaDetailRequest {
	s.Remark = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetScaName(v string) *DescribePropertyScaDetailRequest {
	s.ScaName = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetScaNamePattern(v string) *DescribePropertyScaDetailRequest {
	s.ScaNamePattern = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetScaVersion(v string) *DescribePropertyScaDetailRequest {
	s.ScaVersion = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchCriteriaList(v []*DescribePropertyScaDetailRequestSearchCriteriaList) *DescribePropertyScaDetailRequest {
	s.SearchCriteriaList = v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchInfo(v string) *DescribePropertyScaDetailRequest {
	s.SearchInfo = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchInfoSub(v string) *DescribePropertyScaDetailRequest {
	s.SearchInfoSub = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchItem(v string) *DescribePropertyScaDetailRequest {
	s.SearchItem = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetSearchItemSub(v string) *DescribePropertyScaDetailRequest {
	s.SearchItemSub = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetUseNextToken(v bool) *DescribePropertyScaDetailRequest {
	s.UseNextToken = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetUser(v string) *DescribePropertyScaDetailRequest {
	s.User = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) SetUuid(v string) *DescribePropertyScaDetailRequest {
	s.Uuid = &v
	return s
}

func (s *DescribePropertyScaDetailRequest) Validate() error {
	if s.SearchCriteriaList != nil {
		for _, item := range s.SearchCriteriaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePropertyScaDetailRequestSearchCriteriaList struct {
	// The name of the search criterion.
	//
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The filter value of the search criterion.
	//
	// example:
	//
	// test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePropertyScaDetailRequestSearchCriteriaList) String() string {
	return dara.Prettify(s)
}

func (s DescribePropertyScaDetailRequestSearchCriteriaList) GoString() string {
	return s.String()
}

func (s *DescribePropertyScaDetailRequestSearchCriteriaList) GetName() *string {
	return s.Name
}

func (s *DescribePropertyScaDetailRequestSearchCriteriaList) GetValue() *string {
	return s.Value
}

func (s *DescribePropertyScaDetailRequestSearchCriteriaList) SetName(v string) *DescribePropertyScaDetailRequestSearchCriteriaList {
	s.Name = &v
	return s
}

func (s *DescribePropertyScaDetailRequestSearchCriteriaList) SetValue(v string) *DescribePropertyScaDetailRequestSearchCriteriaList {
	s.Value = &v
	return s
}

func (s *DescribePropertyScaDetailRequestSearchCriteriaList) Validate() error {
	return dara.Validate(s)
}
