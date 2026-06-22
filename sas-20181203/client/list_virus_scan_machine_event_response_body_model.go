// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVirusScanMachineEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListVirusScanMachineEventResponseBodyData) *ListVirusScanMachineEventResponseBody
	GetData() []*ListVirusScanMachineEventResponseBodyData
	SetPageInfo(v *ListVirusScanMachineEventResponseBodyPageInfo) *ListVirusScanMachineEventResponseBody
	GetPageInfo() *ListVirusScanMachineEventResponseBodyPageInfo
	SetRequestId(v string) *ListVirusScanMachineEventResponseBody
	GetRequestId() *string
}

type ListVirusScanMachineEventResponseBody struct {
	// The details of the alert events.
	Data []*ListVirusScanMachineEventResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListVirusScanMachineEventResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// 2DAEF40F-8E1A-550D-8793-99C61C401DD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVirusScanMachineEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBody) GetData() []*ListVirusScanMachineEventResponseBodyData {
	return s.Data
}

func (s *ListVirusScanMachineEventResponseBody) GetPageInfo() *ListVirusScanMachineEventResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListVirusScanMachineEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVirusScanMachineEventResponseBody) SetData(v []*ListVirusScanMachineEventResponseBodyData) *ListVirusScanMachineEventResponseBody {
	s.Data = v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetPageInfo(v *ListVirusScanMachineEventResponseBodyPageInfo) *ListVirusScanMachineEventResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) SetRequestId(v string) *ListVirusScanMachineEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBody) Validate() error {
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

type ListVirusScanMachineEventResponseBodyData struct {
	// The details of the exception event.
	Details []*ListVirusScanMachineEventResponseBodyDataDetails `json:"Details,omitempty" xml:"Details,omitempty" type:"Repeated"`
	// The ID of the alert event.
	//
	// example:
	//
	// 911273
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The name (subtype) of the alert event.
	//
	// example:
	//
	// 恶意脚本代码执行
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The instance name.
	//
	// example:
	//
	// i-wz92q7m5hsbgfhdss***
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address.
	//
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address.
	//
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp of the last occurrence, in milliseconds.
	//
	// example:
	//
	// 1682046733628
	LastTimeStamp *int64 `json:"LastTimeStamp,omitempty" xml:"LastTimeStamp,omitempty"`
	// The risk level of the alert event. Valid values:
	//
	// - **serious**: Critical.
	//
	// - **suspicious**: Suspicious.
	//
	// - **remind**: Reminder.
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyData) GetDetails() []*ListVirusScanMachineEventResponseBodyDataDetails {
	return s.Details
}

func (s *ListVirusScanMachineEventResponseBodyData) GetEventId() *int64 {
	return s.EventId
}

func (s *ListVirusScanMachineEventResponseBodyData) GetEventName() *string {
	return s.EventName
}

func (s *ListVirusScanMachineEventResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListVirusScanMachineEventResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListVirusScanMachineEventResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListVirusScanMachineEventResponseBodyData) GetLastTimeStamp() *int64 {
	return s.LastTimeStamp
}

func (s *ListVirusScanMachineEventResponseBodyData) GetLevel() *string {
	return s.Level
}

func (s *ListVirusScanMachineEventResponseBodyData) SetDetails(v []*ListVirusScanMachineEventResponseBodyDataDetails) *ListVirusScanMachineEventResponseBodyData {
	s.Details = v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetEventId(v int64) *ListVirusScanMachineEventResponseBodyData {
	s.EventId = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetEventName(v string) *ListVirusScanMachineEventResponseBodyData {
	s.EventName = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetInstanceName(v string) *ListVirusScanMachineEventResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetInternetIp(v string) *ListVirusScanMachineEventResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetIntranetIp(v string) *ListVirusScanMachineEventResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetLastTimeStamp(v int64) *ListVirusScanMachineEventResponseBodyData {
	s.LastTimeStamp = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) SetLevel(v string) *ListVirusScanMachineEventResponseBodyData {
	s.Level = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyData) Validate() error {
	if s.Details != nil {
		for _, item := range s.Details {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListVirusScanMachineEventResponseBodyDataDetails struct {
	// The display type of valueDisplay.
	//
	// Valid values:
	//
	// - **download_url**: download link.
	//
	// example:
	//
	// download_url
	InfoType *string `json:"InfoType,omitempty" xml:"InfoType,omitempty"`
	// The display name of the alert event.
	//
	// example:
	//
	// Trojan Path
	NameDisplay *string `json:"NameDisplay,omitempty" xml:"NameDisplay,omitempty"`
	// The display format of the exception event details.
	//
	// Valid values:
	//
	// - **text**: plain text
	//
	// - **html**: rich text.
	//
	// example:
	//
	// html
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The additional attribute information of the exception event.
	//
	// For example, the logon time or logon location of an unusual logon alert, or the trojan file path or trojan type of an alert.
	//
	// example:
	//
	// getopt
	ValueDisplay *string `json:"ValueDisplay,omitempty" xml:"ValueDisplay,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) GetInfoType() *string {
	return s.InfoType
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) GetNameDisplay() *string {
	return s.NameDisplay
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) GetType() *string {
	return s.Type
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) GetValueDisplay() *string {
	return s.ValueDisplay
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) SetInfoType(v string) *ListVirusScanMachineEventResponseBodyDataDetails {
	s.InfoType = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) SetNameDisplay(v string) *ListVirusScanMachineEventResponseBodyDataDetails {
	s.NameDisplay = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) SetType(v string) *ListVirusScanMachineEventResponseBodyDataDetails {
	s.Type = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) SetValueDisplay(v string) *ListVirusScanMachineEventResponseBodyDataDetails {
	s.ValueDisplay = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyDataDetails) Validate() error {
	return dara.Validate(s)
}

type ListVirusScanMachineEventResponseBodyPageInfo struct {
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The maximum number of entries per page in a paged query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of alert events returned.
	//
	// example:
	//
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListVirusScanMachineEventResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListVirusScanMachineEventResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) SetCurrentPage(v int32) *ListVirusScanMachineEventResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) SetPageSize(v int32) *ListVirusScanMachineEventResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) SetTotalCount(v int32) *ListVirusScanMachineEventResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListVirusScanMachineEventResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
