// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePrimaryObjectFavoriteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *TogglePrimaryObjectFavoriteRequest
	GetAction() *string
	SetObjectIds(v []*string) *TogglePrimaryObjectFavoriteRequest
	GetObjectIds() []*string
	SetObjectType(v string) *TogglePrimaryObjectFavoriteRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *TogglePrimaryObjectFavoriteRequest
	GetTenantId() *string
}

type TogglePrimaryObjectFavoriteRequest struct {
	// The operation type. Valid values:
	//
	// - **1**: Add to whitelist.
	//
	// - **2**: Remove from whitelist.
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The list of primary object business IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectIds []*string `json:"objectIds,omitempty" xml:"objectIds,omitempty" type:"Repeated"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The ID of the effective tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s TogglePrimaryObjectFavoriteRequest) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteRequest) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteRequest) GetAction() *string {
	return s.Action
}

func (s *TogglePrimaryObjectFavoriteRequest) GetObjectIds() []*string {
	return s.ObjectIds
}

func (s *TogglePrimaryObjectFavoriteRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *TogglePrimaryObjectFavoriteRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *TogglePrimaryObjectFavoriteRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *TogglePrimaryObjectFavoriteRequest) SetAction(v string) *TogglePrimaryObjectFavoriteRequest {
	s.Action = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteRequest) SetObjectIds(v []*string) *TogglePrimaryObjectFavoriteRequest {
	s.ObjectIds = v
	return s
}

func (s *TogglePrimaryObjectFavoriteRequest) SetObjectType(v string) *TogglePrimaryObjectFavoriteRequest {
	s.ObjectType = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteRequest) SetOperatingObjectName(v string) *TogglePrimaryObjectFavoriteRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteRequest) SetTenantId(v string) *TogglePrimaryObjectFavoriteRequest {
	s.TenantId = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteRequest) Validate() error {
	return dara.Validate(s)
}
