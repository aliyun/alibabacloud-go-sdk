// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityEventMarkMissListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*DescribeSecurityEventMarkMissListResponseBodyList) *DescribeSecurityEventMarkMissListResponseBody
	GetList() []*DescribeSecurityEventMarkMissListResponseBodyList
	SetPageInfo(v *DescribeSecurityEventMarkMissListResponseBodyPageInfo) *DescribeSecurityEventMarkMissListResponseBody
	GetPageInfo() *DescribeSecurityEventMarkMissListResponseBodyPageInfo
	SetRequestId(v string) *DescribeSecurityEventMarkMissListResponseBody
	GetRequestId() *string
}

type DescribeSecurityEventMarkMissListResponseBody struct {
	// The ID of the rule.
	List []*DescribeSecurityEventMarkMissListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *DescribeSecurityEventMarkMissListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSecurityEventMarkMissListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListResponseBody) GetList() []*DescribeSecurityEventMarkMissListResponseBodyList {
	return s.List
}

func (s *DescribeSecurityEventMarkMissListResponseBody) GetPageInfo() *DescribeSecurityEventMarkMissListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeSecurityEventMarkMissListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityEventMarkMissListResponseBody) SetList(v []*DescribeSecurityEventMarkMissListResponseBodyList) *DescribeSecurityEventMarkMissListResponseBody {
	s.List = v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBody) SetPageInfo(v *DescribeSecurityEventMarkMissListResponseBodyPageInfo) *DescribeSecurityEventMarkMissListResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBody) SetRequestId(v string) *DescribeSecurityEventMarkMissListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventMarkMissListResponseBodyList struct {
	// The user ID.
	//
	// example:
	//
	// 176618589410****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The name of the alert event. The value indicates a subtype.
	//
	// example:
	//
	// Login with unusual location
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The name of the alert event. The value indicates a type.
	//
	// example:
	//
	// login_common_location
	EventNameOriginal *string `json:"EventNameOriginal,omitempty" xml:"EventNameOriginal,omitempty"`
	// The subtype of the alert event.
	//
	// example:
	//
	// Unusual Logon
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The type of the alert event.
	//
	// example:
	//
	// login_common_location
	EventTypeOriginal *string `json:"EventTypeOriginal,omitempty" xml:"EventTypeOriginal,omitempty"`
	// The field that is used in the whitelist rule.
	//
	// example:
	//
	// type
	Field *string `json:"Field,omitempty" xml:"Field,omitempty"`
	// The value of the field.
	//
	// example:
	//
	// root
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
	// The alias of the field.
	//
	// example:
	//
	// Logon Time
	FiledAliasName *string `json:"FiledAliasName,omitempty" xml:"FiledAliasName,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 104037
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// rm-bp1e8t4q15sr3****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the asset.
	//
	// example:
	//
	// sql-test-001
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the server.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The operator. Valid values:
	//
	// - **contains**: contains
	//
	// - **notContains**: does not contain
	//
	// - **strEqual**: equals
	//
	// - **strNotEqual**: does not equal
	//
	// - **regex**: regular expression
	//
	// example:
	//
	// contains
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The UUID of the asset.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSecurityEventMarkMissListResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListResponseBodyList) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetEventNameOriginal() *string {
	return s.EventNameOriginal
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetEventTypeOriginal() *string {
	return s.EventTypeOriginal
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetField() *string {
	return s.Field
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetId() *int64 {
	return s.Id
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetOperate() *string {
	return s.Operate
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetAliUid(v int64) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.AliUid = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetEventName(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.EventName = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetEventNameOriginal(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.EventNameOriginal = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetEventType(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.EventType = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetEventTypeOriginal(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.EventTypeOriginal = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetField(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.Field = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetFieldValue(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.FieldValue = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetFiledAliasName(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.FiledAliasName = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetId(v int64) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.Id = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetInstanceId(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetInstanceName(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetInternetIp(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetIntranetIp(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetOperate(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.Operate = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) SetUuid(v string) *DescribeSecurityEventMarkMissListResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityEventMarkMissListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityEventMarkMissListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityEventMarkMissListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) SetCount(v int32) *DescribeSecurityEventMarkMissListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSecurityEventMarkMissListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) SetPageSize(v int32) *DescribeSecurityEventMarkMissListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSecurityEventMarkMissListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityEventMarkMissListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
