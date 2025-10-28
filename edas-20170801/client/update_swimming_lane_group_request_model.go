// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *UpdateSwimmingLaneGroupRequest
	GetAppIds() *string
	SetEntryApp(v string) *UpdateSwimmingLaneGroupRequest
	GetEntryApp() *string
	SetGroupId(v int64) *UpdateSwimmingLaneGroupRequest
	GetGroupId() *int64
	SetName(v string) *UpdateSwimmingLaneGroupRequest
	GetName() *string
}

type UpdateSwimmingLaneGroupRequest struct {
	// The list of application IDs related to the lane group.
	//
	// example:
	//
	// 8e7689af-6ddd-4676-8ee6-5fbecdf2****,f72deaac-26ba-429a-948d-5fa47c4a****,5049d2c8-f997-4fc9-92a2-153506a6****,99a2d4b5-99a5-4e25-a964-1bd03a17****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The ingress application. The application is in the EDAS:{application ID} format.
	//
	// example:
	//
	// EDAS:dd2690a7-3fe4-4975-9a4c-5a60ffd6****
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	// The ID of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 98
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the lane group.
	//
	// example:
	//
	// test-swimlanegroup
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneGroupRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *UpdateSwimmingLaneGroupRequest) GetEntryApp() *string {
	return s.EntryApp
}

func (s *UpdateSwimmingLaneGroupRequest) GetGroupId() *int64 {
	return s.GroupId
}

func (s *UpdateSwimmingLaneGroupRequest) GetName() *string {
	return s.Name
}

func (s *UpdateSwimmingLaneGroupRequest) SetAppIds(v string) *UpdateSwimmingLaneGroupRequest {
	s.AppIds = &v
	return s
}

func (s *UpdateSwimmingLaneGroupRequest) SetEntryApp(v string) *UpdateSwimmingLaneGroupRequest {
	s.EntryApp = &v
	return s
}

func (s *UpdateSwimmingLaneGroupRequest) SetGroupId(v int64) *UpdateSwimmingLaneGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateSwimmingLaneGroupRequest) SetName(v string) *UpdateSwimmingLaneGroupRequest {
	s.Name = &v
	return s
}

func (s *UpdateSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
