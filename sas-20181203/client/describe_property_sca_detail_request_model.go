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
	// The type of the asset fingerprint that you want to query. Default value: **sca**. Valid values:
	//
	// 	- **sca**: middleware
	//
	// 	- **sca_database**: database
	//
	// 	- **sca_web**: web service
	//
	// >  If you do not specify this parameter, the default value **sca*	- is used, which indicates that middleware fingerprints are queried.
	//
	// example:
	//
	// sca
	Biz *string `json:"Biz,omitempty" xml:"Biz,omitempty"`
	// The type of the middleware, database, or web service that you want to query. Valid values:
	//
	// 	- **system_service**: system service
	//
	// 	- **software_library**: software library
	//
	// 	- **docker_component**: container component
	//
	// 	- **database**: database
	//
	// 	- **web_container**: web container
	//
	// 	- **jar**: JAR package
	//
	// 	- **web_framework**: web framework
	//
	// example:
	//
	// system_service
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The name of the middleware, database, or web service.
	//
	// >  This parameter is deprecated. You can ignore it.
	//
	// example:
	//
	// 1
	Name      *int64  `json:"Name,omitempty" xml:"Name,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of entries to return on each page. Default value: **10**.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The PID.
	//
	// example:
	//
	// 756
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The port that the process monitors.
	//
	// example:
	//
	// 68
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The timestamp when the process ends. Unit: milliseconds.
	//
	// example:
	//
	// 1641110965
	ProcessStartedEnd *int64 `json:"ProcessStartedEnd,omitempty" xml:"ProcessStartedEnd,omitempty"`
	// The timestamp when the process starts. Unit: milliseconds.
	//
	// example:
	//
	// 1641024565
	ProcessStartedStart *int64 `json:"ProcessStartedStart,omitempty" xml:"ProcessStartedStart,omitempty"`
	// The search condition, such as a server name or a server IP address.
	//
	// >  Fuzzy match is supported.
	//
	// example:
	//
	// 192.168
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The name of the asset fingerprint that you want to query.
	//
	// example:
	//
	// openssl
	ScaName *string `json:"ScaName,omitempty" xml:"ScaName,omitempty"`
	// The name of the process.
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
	// The search keyword. You must specify this parameter based on the value of the **SearchItem*	- parameter.
	//
	// 	- If the **SearchItem*	- parameter is set to **name**, you must enter the name of an asset fingerprint.
	//
	// 	- If the **SearchItem*	- parameter is set to **type**, you must enter the type of an asset fingerprint. Valid values:
	//
	//     	- **system_service**: system service
	//
	//     	- **software_library**: software library
	//
	//     	- **docker_component**: container component
	//
	//     	- **database**: database
	//
	//     	- **web_container**: web container
	//
	//     	- **jar**: JAR package
	//
	//     	- **web_framework**: web framework
	//
	// >  You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the asset fingerprints based on the specified name or type.
	//
	// example:
	//
	// openssl
	SearchInfo *string `json:"SearchInfo,omitempty" xml:"SearchInfo,omitempty"`
	// The keyword of the subquery. You must specify this parameter based on the value of the **SearchItemSub*	- parameter.
	//
	// 	- If the **SearchItemSub*	- parameter is set to **port**, you must enter a port number.
	//
	// 	- If the **SearchItemSub*	- parameter is set to **pid**, you must enter a process ID (PID).
	//
	// 	- If the **SearchItemSub*	- parameter is set to **version**, you must enter the version of a database, middleware, or web service.
	//
	// 	- If the **SearchItemSub*	- parameter is set to **user**, you must enter a username.
	//
	// >  The subquery is used to search for data of a specified database, middleware, or web service.
	//
	// example:
	//
	// 1.0.2k
	SearchInfoSub *string `json:"SearchInfoSub,omitempty" xml:"SearchInfoSub,omitempty"`
	// The type of the search condition. Valid values:
	//
	// 	- **name**: the name of a database, middleware, or web service
	//
	// 	- **type**: the type of a database, middleware, or web service
	//
	// >  You must specify both the **SearchItem*	- and **SearchInfo*	- parameters before you can query the asset fingerprints based on the specified name or type.
	//
	// example:
	//
	// name
	SearchItem *string `json:"SearchItem,omitempty" xml:"SearchItem,omitempty"`
	// The type of the subquery. Valid values:
	//
	// 	- **port**
	//
	// 	- **pid**
	//
	// 	- **version**
	//
	// 	- **user**
	//
	// example:
	//
	// version
	SearchItemSub *string `json:"SearchItemSub,omitempty" xml:"SearchItemSub,omitempty"`
	UseNextToken  *bool   `json:"UseNextToken,omitempty" xml:"UseNextToken,omitempty"`
	// The user who runs the process.
	//
	// example:
	//
	// root
	User *string `json:"User,omitempty" xml:"User,omitempty"`
	// The UUID of the server on which the middleware, database, or web service is run.
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
	return dara.Validate(s)
}
