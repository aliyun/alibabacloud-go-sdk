// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v string) *InsertSwimmingLaneGroupRequest
	GetAppIds() *string
	SetEntryApp(v string) *InsertSwimmingLaneGroupRequest
	GetEntryApp() *string
	SetLogicalRegionId(v string) *InsertSwimmingLaneGroupRequest
	GetLogicalRegionId() *string
	SetName(v string) *InsertSwimmingLaneGroupRequest
	GetName() *string
}

type InsertSwimmingLaneGroupRequest struct {
	// IDs of all applications that are related to the lane group. Separate multiple applications with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// bdb251cc-02a6-48dd-891b-2ab21b25****,ee33ed0c-fddc-47b5-9f63-e1ccc4be****
	AppIds *string `json:"AppIds,omitempty" xml:"AppIds,omitempty"`
	// The ingress application. The application is in the EDAS:{application ID} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// EDAS:5cc89013-9232-4b36-b3eb-ff89b3d2****
	EntryApp *string `json:"EntryApp,omitempty" xml:"EntryApp,omitempty"`
	// The ID of the custom namespace. The ID is in the physical region ID:custom namespace identifier format. Example: cn-hangzhou:test.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou:test
	LogicalRegionId *string `json:"LogicalRegionId,omitempty" xml:"LogicalRegionId,omitempty"`
	// The name of the lane group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s InsertSwimmingLaneGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupRequest) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupRequest) GetAppIds() *string {
	return s.AppIds
}

func (s *InsertSwimmingLaneGroupRequest) GetEntryApp() *string {
	return s.EntryApp
}

func (s *InsertSwimmingLaneGroupRequest) GetLogicalRegionId() *string {
	return s.LogicalRegionId
}

func (s *InsertSwimmingLaneGroupRequest) GetName() *string {
	return s.Name
}

func (s *InsertSwimmingLaneGroupRequest) SetAppIds(v string) *InsertSwimmingLaneGroupRequest {
	s.AppIds = &v
	return s
}

func (s *InsertSwimmingLaneGroupRequest) SetEntryApp(v string) *InsertSwimmingLaneGroupRequest {
	s.EntryApp = &v
	return s
}

func (s *InsertSwimmingLaneGroupRequest) SetLogicalRegionId(v string) *InsertSwimmingLaneGroupRequest {
	s.LogicalRegionId = &v
	return s
}

func (s *InsertSwimmingLaneGroupRequest) SetName(v string) *InsertSwimmingLaneGroupRequest {
	s.Name = &v
	return s
}

func (s *InsertSwimmingLaneGroupRequest) Validate() error {
	return dara.Validate(s)
}
