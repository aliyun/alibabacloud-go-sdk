// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeVisibilityModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataPortalId(v string) *ChangeVisibilityModelRequest
	GetDataPortalId() *string
	SetMenuIds(v string) *ChangeVisibilityModelRequest
	GetMenuIds() *string
	SetShowOnlyWithAccess(v bool) *ChangeVisibilityModelRequest
	GetShowOnlyWithAccess() *bool
}

type ChangeVisibilityModelRequest struct {
	// The ID of the BI portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The menu ID of the BI portal leaf node.
	//
	// 	- The directory menu cannot be authorized.
	//
	// 	- You can upload multiple parameters at a time. Separate multiple IDs with commas (,). The maximum number of parameters that can be modified at a time is 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// Whether only authorization is visible. Valid values:
	//
	// 	- true: Only the authorization is visible.
	//
	// 	- false: Both are visible.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	ShowOnlyWithAccess *bool `json:"ShowOnlyWithAccess,omitempty" xml:"ShowOnlyWithAccess,omitempty"`
}

func (s ChangeVisibilityModelRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeVisibilityModelRequest) GoString() string {
	return s.String()
}

func (s *ChangeVisibilityModelRequest) GetDataPortalId() *string {
	return s.DataPortalId
}

func (s *ChangeVisibilityModelRequest) GetMenuIds() *string {
	return s.MenuIds
}

func (s *ChangeVisibilityModelRequest) GetShowOnlyWithAccess() *bool {
	return s.ShowOnlyWithAccess
}

func (s *ChangeVisibilityModelRequest) SetDataPortalId(v string) *ChangeVisibilityModelRequest {
	s.DataPortalId = &v
	return s
}

func (s *ChangeVisibilityModelRequest) SetMenuIds(v string) *ChangeVisibilityModelRequest {
	s.MenuIds = &v
	return s
}

func (s *ChangeVisibilityModelRequest) SetShowOnlyWithAccess(v bool) *ChangeVisibilityModelRequest {
	s.ShowOnlyWithAccess = &v
	return s
}

func (s *ChangeVisibilityModelRequest) Validate() error {
	return dara.Validate(s)
}
