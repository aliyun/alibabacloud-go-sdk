// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNotificationPoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectedMode(v bool) *ListNotificationPoliciesRequest
	GetDirectedMode() *bool
	SetIds(v string) *ListNotificationPoliciesRequest
	GetIds() *string
	SetIsDetail(v bool) *ListNotificationPoliciesRequest
	GetIsDetail() *bool
	SetName(v string) *ListNotificationPoliciesRequest
	GetName() *string
	SetPage(v int64) *ListNotificationPoliciesRequest
	GetPage() *int64
	SetRegionId(v string) *ListNotificationPoliciesRequest
	GetRegionId() *string
	SetSize(v int64) *ListNotificationPoliciesRequest
	GetSize() *int64
}

type ListNotificationPoliciesRequest struct {
	// Specifies whether to enable simple mode.
	//
	// example:
	//
	// true
	DirectedMode *bool `json:"DirectedMode,omitempty" xml:"DirectedMode,omitempty"`
	// The ID of the notification policy.
	//
	// example:
	//
	// 12345
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// Specifies whether to query the details about notification policies. Valid values:
	//
	// 	- `true`: Details about notification policies are queried.
	//
	// 	- `false`: Details about notification policies are not queried.
	//
	// example:
	//
	// false
	IsDetail *bool `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// The name of the notification policy.
	//
	// example:
	//
	// notificationpolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the region. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListNotificationPoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNotificationPoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListNotificationPoliciesRequest) GetDirectedMode() *bool {
	return s.DirectedMode
}

func (s *ListNotificationPoliciesRequest) GetIds() *string {
	return s.Ids
}

func (s *ListNotificationPoliciesRequest) GetIsDetail() *bool {
	return s.IsDetail
}

func (s *ListNotificationPoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListNotificationPoliciesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListNotificationPoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListNotificationPoliciesRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListNotificationPoliciesRequest) SetDirectedMode(v bool) *ListNotificationPoliciesRequest {
	s.DirectedMode = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetIds(v string) *ListNotificationPoliciesRequest {
	s.Ids = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetIsDetail(v bool) *ListNotificationPoliciesRequest {
	s.IsDetail = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetName(v string) *ListNotificationPoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetPage(v int64) *ListNotificationPoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetRegionId(v string) *ListNotificationPoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListNotificationPoliciesRequest) SetSize(v int64) *ListNotificationPoliciesRequest {
	s.Size = &v
	return s
}

func (s *ListNotificationPoliciesRequest) Validate() error {
	return dara.Validate(s)
}
