// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActiveAnnouncementsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListActiveAnnouncementsResponseBody
	GetCode() *string
	SetItems(v []*ListActiveAnnouncementsResponseBodyItems) *ListActiveAnnouncementsResponseBody
	GetItems() []*ListActiveAnnouncementsResponseBodyItems
	SetMessage(v string) *ListActiveAnnouncementsResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *ListActiveAnnouncementsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListActiveAnnouncementsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListActiveAnnouncementsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListActiveAnnouncementsResponseBody
	GetTotal() *int64
}

type ListActiveAnnouncementsResponseBody struct {
	// The business status code. A value of 200 indicates success. A failure returns a backend error code (ERR.	- / InvalidParameter.*).
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of MCP cards.
	Items []*ListActiveAnnouncementsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number, starting from 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListActiveAnnouncementsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListActiveAnnouncementsResponseBody) GoString() string {
	return s.String()
}

func (s *ListActiveAnnouncementsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListActiveAnnouncementsResponseBody) GetItems() []*ListActiveAnnouncementsResponseBodyItems {
	return s.Items
}

func (s *ListActiveAnnouncementsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListActiveAnnouncementsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListActiveAnnouncementsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListActiveAnnouncementsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListActiveAnnouncementsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListActiveAnnouncementsResponseBody) SetCode(v string) *ListActiveAnnouncementsResponseBody {
	s.Code = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetItems(v []*ListActiveAnnouncementsResponseBodyItems) *ListActiveAnnouncementsResponseBody {
	s.Items = v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetMessage(v string) *ListActiveAnnouncementsResponseBody {
	s.Message = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetPageNumber(v int64) *ListActiveAnnouncementsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetPageSize(v int64) *ListActiveAnnouncementsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetRequestId(v string) *ListActiveAnnouncementsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) SetTotal(v int64) *ListActiveAnnouncementsResponseBody {
	s.Total = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListActiveAnnouncementsResponseBodyItems struct {
	// The business ID of the announcement.
	//
	// example:
	//
	// 1001
	AnnouncementId *int64 `json:"announcementId,omitempty" xml:"announcementId,omitempty"`
	// The returned content.
	//
	// example:
	//
	// The system will undergo maintenance tonight
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The user ID of the project creator.
	//
	// example:
	//
	// 10001
	CreatedBy *int64 `json:"createdBy,omitempty" xml:"createdBy,omitempty"`
	// The display page. Valid values: ALL, FRONTEND, BACKEND.
	//
	// example:
	//
	// ALL
	DisplayPage *string `json:"displayPage,omitempty" xml:"displayPage,omitempty"`
	// The display type and group label.
	//
	// example:
	//
	// LIST
	DisplayType *string `json:"displayType,omitempty" xml:"displayType,omitempty"`
	// The custom effective end time.
	//
	// example:
	//
	// 2026-08-21T14:00:00+08:00
	EffectiveEnd *string `json:"effectiveEnd,omitempty" xml:"effectiveEnd,omitempty"`
	// The effective start time in ISO 8601 format with time zone. If not specified, the announcement takes effect immediately.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	EffectiveStart *string `json:"effectiveStart,omitempty" xml:"effectiveStart,omitempty"`
	// The priority of the free task.
	//
	// - Uses the default priority of the project, as shown in the following figure. The default priorities are as follows:
	//
	//      - **-10**: Low. This is the default value.
	//
	//      - **0**: Normal.
	//
	//     - **1**: Urgent.
	//
	//      - **2**: Very urgent.
	//
	// ![](https://img.alicdn.com/imgextra/i1/O1CN01hNuSPz25juCzgxhmW_!!6000000007563-2-tps-2682-1304.png)
	//
	// - Custom priority, as shown in the following figure, with an additional "Moderately urgent" level.
	//
	// The value of this parameter is subject to the actual response of the API. A higher priority corresponds to a larger value.
	//
	// ![](https://img.alicdn.com/imgextra/i1/O1CN01V67b3i1mkNvJiW8D1_!!6000000004992-2-tps-2128-1126.png)
	//
	// example:
	//
	// GENERAL
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The publish time in ISO 8601 format.
	//
	// example:
	//
	// 2026-08-20T14:00:00+08:00
	PublishedAt *string `json:"publishedAt,omitempty" xml:"publishedAt,omitempty"`
	// The task status. Running is returned upon submission.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The title of the scheduled meeting.
	//
	// example:
	//
	// System Maintenance Notice
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s ListActiveAnnouncementsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListActiveAnnouncementsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetAnnouncementId() *int64 {
	return s.AnnouncementId
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetContent() *string {
	return s.Content
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetCreatedBy() *int64 {
	return s.CreatedBy
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetDisplayPage() *string {
	return s.DisplayPage
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetDisplayType() *string {
	return s.DisplayType
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetEffectiveEnd() *string {
	return s.EffectiveEnd
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetEffectiveStart() *string {
	return s.EffectiveStart
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetPriority() *string {
	return s.Priority
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetPublishedAt() *string {
	return s.PublishedAt
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListActiveAnnouncementsResponseBodyItems) GetTitle() *string {
	return s.Title
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetAnnouncementId(v int64) *ListActiveAnnouncementsResponseBodyItems {
	s.AnnouncementId = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetContent(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.Content = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetCreatedBy(v int64) *ListActiveAnnouncementsResponseBodyItems {
	s.CreatedBy = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetDisplayPage(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.DisplayPage = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetDisplayType(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.DisplayType = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetEffectiveEnd(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.EffectiveEnd = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetEffectiveStart(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.EffectiveStart = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetPriority(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.Priority = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetPublishedAt(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.PublishedAt = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetStatus(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) SetTitle(v string) *ListActiveAnnouncementsResponseBodyItems {
	s.Title = &v
	return s
}

func (s *ListActiveAnnouncementsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
