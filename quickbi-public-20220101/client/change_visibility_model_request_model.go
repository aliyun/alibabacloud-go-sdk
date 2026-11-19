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
	// The ID of the data portal.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0d173abb53e84c8ca7495429163b****
	DataPortalId *string `json:"DataPortalId,omitempty" xml:"DataPortalId,omitempty"`
	// The IDs of the leaf node menus in the data portal.
	//
	// - You can specify multiple IDs. Separate multiple IDs with commas (,). A maximum of 100 IDs can be modified at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 54kqgoa****,pg1n135****
	MenuIds *string `json:"MenuIds,omitempty" xml:"MenuIds,omitempty"`
	// Specifies whether the menus are visible only to authorized users. Valid values:
	//
	// - true: The menus are visible only to authorized users.
	//
	// - false: The menus are visible to all users.
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
