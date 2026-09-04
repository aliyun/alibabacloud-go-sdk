// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetGraphName() *string
	SetObjectIdsShrink(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetObjectIdsShrink() *string
	SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetObjectType() *string
	SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetOperatingObjectName() *string
	SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest
	GetTenantId() *string
}

type BatchRemoveOperatingObjectFavoritesShrinkRequest struct {
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// The list of primary object business IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract-001
	ObjectIdsShrink *string `json:"objectIds,omitempty" xml:"objectIds,omitempty"`
	// The object type, such as customer. This parameter has a value when type is set to mention.
	//
	// This parameter is required.
	//
	// example:
	//
	// contract
	ObjectType *string `json:"objectType,omitempty" xml:"objectType,omitempty"`
	// The operating object name.
	//
	// This parameter is required.
	//
	// example:
	//
	// customer_assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The tenant ID. This is a common parameter. Pass it explicitly in winnexo-cli by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s BatchRemoveOperatingObjectFavoritesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesShrinkRequest) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetObjectIdsShrink() *string {
	return s.ObjectIdsShrink
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetGraphName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.GraphName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetObjectIdsShrink(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.ObjectIdsShrink = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetObjectType(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.ObjectType = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetOperatingObjectName(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) SetTenantId(v string) *BatchRemoveOperatingObjectFavoritesShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
