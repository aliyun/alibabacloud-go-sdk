// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSimilarSecurityEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageInfo(v *DescribeSimilarSecurityEventsResponseBodyPageInfo) *DescribeSimilarSecurityEventsResponseBody
	GetPageInfo() *DescribeSimilarSecurityEventsResponseBodyPageInfo
	SetRequestId(v string) *DescribeSimilarSecurityEventsResponseBody
	GetRequestId() *string
	SetSecurityEventsResponse(v []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) *DescribeSimilarSecurityEventsResponseBody
	GetSecurityEventsResponse() []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse
}

type DescribeSimilarSecurityEventsResponseBody struct {
	// The pagination information.
	PageInfo *DescribeSimilarSecurityEventsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9F4217C8-763F-51EF-84D4-5535E072B2D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the alert events that are triggered by the same rule or of the same alert type.
	SecurityEventsResponse []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse `json:"SecurityEventsResponse,omitempty" xml:"SecurityEventsResponse,omitempty" type:"Repeated"`
}

func (s DescribeSimilarSecurityEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBody) GetPageInfo() *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *DescribeSimilarSecurityEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSimilarSecurityEventsResponseBody) GetSecurityEventsResponse() []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	return s.SecurityEventsResponse
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetPageInfo(v *DescribeSimilarSecurityEventsResponseBodyPageInfo) *DescribeSimilarSecurityEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetRequestId(v string) *DescribeSimilarSecurityEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBody) SetSecurityEventsResponse(v []*DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) *DescribeSimilarSecurityEventsResponseBody {
	s.SecurityEventsResponse = v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBody) Validate() error {
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityEventsResponse != nil {
		for _, item := range s.SecurityEventsResponse {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSimilarSecurityEventsResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 5
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
	// 5
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSimilarSecurityEventsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetCount(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetCurrentPage(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetPageSize(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) SetTotalCount(v int32) *DescribeSimilarSecurityEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse struct {
	// The name of the alert event.
	//
	// example:
	//
	// Trojan
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- Suspicious Process
	//
	// 	- Webshell
	//
	// 	- Unusual Logon
	//
	// 	- Malicious Software
	//
	// 	- Sensitive File Tampering
	//
	// 	- Unusual Network Connection
	//
	// 	- Other
	//
	// 	- Suspicious Account
	//
	// 	- Cloud threat detection
	//
	// 	- Precision defense
	//
	// 	- Application Whitelist
	//
	// 	- Persistence
	//
	// 	- Web Application Threat Detection
	//
	// 	- Malicious scripts
	//
	// 	- Malicious Network Activity
	//
	// 	- K8s Abnormal Behavior
	//
	// 	- Website backdoor (local engine)
	//
	// 	- Exploit
	//
	// 	- Image Scan
	//
	// 	- Trusted exception
	//
	// For more information about alert types, see [Overview](https://help.aliyun.com/document_detail/68388.html).
	//
	// example:
	//
	// Malicious Software
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The timestamp generated when the alert event was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1648544361480
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The timestamp generated when the alert event was first detected. Unit: milliseconds.
	//
	// example:
	//
	// 1648457961000
	OccurrenceTime *int64 `json:"OccurrenceTime,omitempty" xml:"OccurrenceTime,omitempty"`
	// The ID of the alert event.
	//
	// example:
	//
	// 158661
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The UUID of the server that was affected by the alert event.
	//
	// example:
	//
	// qweeqq-13232-daweq-w****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetEventName() *string {
	return s.EventName
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetEventType() *string {
	return s.EventType
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetOccurrenceTime() *int64 {
	return s.OccurrenceTime
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetEventName(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.EventName = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetEventType(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.EventType = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetLastTime(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.LastTime = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetOccurrenceTime(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.OccurrenceTime = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetSecurityEventId(v int64) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.SecurityEventId = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) SetUuid(v string) *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse {
	s.Uuid = &v
	return s
}

func (s *DescribeSimilarSecurityEventsResponseBodySecurityEventsResponse) Validate() error {
	return dara.Validate(s)
}
