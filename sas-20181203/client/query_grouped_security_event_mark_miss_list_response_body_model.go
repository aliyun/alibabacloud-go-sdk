// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupedSecurityEventMarkMissListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetCode() *string
	SetList(v []*QueryGroupedSecurityEventMarkMissListResponseBodyList) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetList() []*QueryGroupedSecurityEventMarkMissListResponseBodyList
	SetMessage(v string) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetMessage() *string
	SetPageInfo(v *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetPageInfo() *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo
	SetRequestId(v string) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryGroupedSecurityEventMarkMissListResponseBody
	GetSuccess() *bool
}

type QueryGroupedSecurityEventMarkMissListResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request is successful. Other status codes indicate that the request fails. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// An array that consists of the whitelist rules.
	List []*QueryGroupedSecurityEventMarkMissListResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The error message returned.
	//
	// example:
	//
	// There was an error with your request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination information.
	PageInfo *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 965F9282-D403-4FA2-B1B9-10F62DC719BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryGroupedSecurityEventMarkMissListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupedSecurityEventMarkMissListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetList() []*QueryGroupedSecurityEventMarkMissListResponseBodyList {
	return s.List
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetPageInfo() *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo {
	return s.PageInfo
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetCode(v string) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetList(v []*QueryGroupedSecurityEventMarkMissListResponseBodyList) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.List = v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetMessage(v string) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetPageInfo(v *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.PageInfo = v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetRequestId(v string) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) SetSuccess(v bool) *QueryGroupedSecurityEventMarkMissListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
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

type QueryGroupedSecurityEventMarkMissListResponseBodyList struct {
	// The ID of the user.
	//
	// example:
	//
	// 31412647
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The handling method. Valid values:
	//
	// 	- **auto_add_white**: Automatically Added to Whitelist
	//
	// 	- **defense_not_notification**: Defense Without Notification
	//
	// example:
	//
	// auto_add_white
	DisposalWay *string `json:"DisposalWay,omitempty" xml:"DisposalWay,omitempty"`
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
	// Unusual logon
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
	// The operator. Valid values:
	//
	// 	- **contains**: contains
	//
	// 	- **notContains**: does not contain
	//
	// 	- **strEqual**: equals
	//
	// 	- **strNotEqual**: does not equal
	//
	// 	- **regex**: regular expression
	//
	// example:
	//
	// contains
	Operate *string `json:"Operate,omitempty" xml:"Operate,omitempty"`
	// The UUIDs of assets. Multiple UUIDs are separated by commas (,).
	//
	// example:
	//
	// 6985b88c-eb19-4d27-98ad-e4a42312****,5721d503-9b04-4243-89ca-1fb8ca5e****,db2678c3-10e3-4a20-92f1-265f6****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s QueryGroupedSecurityEventMarkMissListResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupedSecurityEventMarkMissListResponseBodyList) GoString() string {
	return s.String()
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetDisposalWay() *string {
	return s.DisposalWay
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetEventName() *string {
	return s.EventName
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetEventNameOriginal() *string {
	return s.EventNameOriginal
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetEventType() *string {
	return s.EventType
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetEventTypeOriginal() *string {
	return s.EventTypeOriginal
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetField() *string {
	return s.Field
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetFiledAliasName() *string {
	return s.FiledAliasName
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetOperate() *string {
	return s.Operate
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) GetUuids() *string {
	return s.Uuids
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetAliUid(v int64) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.AliUid = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetDisposalWay(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.DisposalWay = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetEventName(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.EventName = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetEventNameOriginal(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.EventNameOriginal = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetEventType(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.EventType = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetEventTypeOriginal(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.EventTypeOriginal = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetField(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.Field = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetFieldValue(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.FieldValue = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetFiledAliasName(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.FiledAliasName = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetOperate(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.Operate = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) SetUuids(v string) *QueryGroupedSecurityEventMarkMissListResponseBodyList {
	s.Uuids = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 9
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page. Default value: **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) SetCount(v int32) *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) SetCurrentPage(v int32) *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) SetPageSize(v int32) *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) SetTotalCount(v int32) *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *QueryGroupedSecurityEventMarkMissListResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
