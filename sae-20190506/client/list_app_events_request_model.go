// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListAppEventsRequest
	GetAppId() *string
	SetCurrentPage(v int32) *ListAppEventsRequest
	GetCurrentPage() *int32
	SetEventType(v string) *ListAppEventsRequest
	GetEventType() *string
	SetNamespace(v string) *ListAppEventsRequest
	GetNamespace() *string
	SetObjectKind(v string) *ListAppEventsRequest
	GetObjectKind() *string
	SetObjectName(v string) *ListAppEventsRequest
	GetObjectName() *string
	SetPageSize(v int32) *ListAppEventsRequest
	GetPageSize() *int32
	SetReason(v string) *ListAppEventsRequest
	GetReason() *string
}

type ListAppEventsRequest struct {
	// The application ID.
	//
	// example:
	//
	// f7730764-d88f-4b9a-8d8e-cd8efbfe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The page number of the page to return.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The type of the event. Valid values:
	//
	// 	- **Warning**: an alert.
	//
	// 	- **Normal**: a normal event.
	//
	// example:
	//
	// Warning
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The namespace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The type of the object. Valid values:
	//
	// 	- **Deployment**: an application.
	//
	// 	- **Pod**: an application instance.
	//
	// 	- **Service**: a Server Load Balancer (SLB) instance.
	//
	// 	- **HorizontalPodAutoscaler**: an auto scaling policy.
	//
	// 	- **CloneSet**: an application.
	//
	// example:
	//
	// Pod
	ObjectKind *string `json:"ObjectKind,omitempty" xml:"ObjectKind,omitempty"`
	// The name of the object. Fuzzy search by prefix is supported.
	//
	// example:
	//
	// errew-b86bf540-b4dc-47d8-a42f-b4997c14bd8f-5595cbddd6-x****
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// The number of entries to return on each page. Valid values: 0 to 10000.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The cause of the event. Fuzzy search by prefix is supported.
	//
	// example:
	//
	// Started
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ListAppEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppEventsRequest) GoString() string {
	return s.String()
}

func (s *ListAppEventsRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListAppEventsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListAppEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *ListAppEventsRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListAppEventsRequest) GetObjectKind() *string {
	return s.ObjectKind
}

func (s *ListAppEventsRequest) GetObjectName() *string {
	return s.ObjectName
}

func (s *ListAppEventsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppEventsRequest) GetReason() *string {
	return s.Reason
}

func (s *ListAppEventsRequest) SetAppId(v string) *ListAppEventsRequest {
	s.AppId = &v
	return s
}

func (s *ListAppEventsRequest) SetCurrentPage(v int32) *ListAppEventsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListAppEventsRequest) SetEventType(v string) *ListAppEventsRequest {
	s.EventType = &v
	return s
}

func (s *ListAppEventsRequest) SetNamespace(v string) *ListAppEventsRequest {
	s.Namespace = &v
	return s
}

func (s *ListAppEventsRequest) SetObjectKind(v string) *ListAppEventsRequest {
	s.ObjectKind = &v
	return s
}

func (s *ListAppEventsRequest) SetObjectName(v string) *ListAppEventsRequest {
	s.ObjectName = &v
	return s
}

func (s *ListAppEventsRequest) SetPageSize(v int32) *ListAppEventsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppEventsRequest) SetReason(v string) *ListAppEventsRequest {
	s.Reason = &v
	return s
}

func (s *ListAppEventsRequest) Validate() error {
	return dara.Validate(s)
}
