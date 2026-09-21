// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticResourceOwner interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogUuid(v string) *AgenticResourceOwner
	GetCatalogUuid() *string
	SetDatabaseQualifiedName(v string) *AgenticResourceOwner
	GetDatabaseQualifiedName() *string
	SetDatabaseUuid(v string) *AgenticResourceOwner
	GetDatabaseUuid() *string
	SetGrantBy(v string) *AgenticResourceOwner
	GetGrantBy() *string
	SetGrantFrom(v string) *AgenticResourceOwner
	GetGrantFrom() *string
	SetOwnerPrincipalId(v string) *AgenticResourceOwner
	GetOwnerPrincipalId() *string
	SetOwnerPrincipalType(v string) *AgenticResourceOwner
	GetOwnerPrincipalType() *string
	SetResourceType(v string) *AgenticResourceOwner
	GetResourceType() *string
}

type AgenticResourceOwner struct {
	// The UUID of the Catalog to which the resource belongs.
	//
	// example:
	//
	// mc-HZ-5d9fbt8wW4AnGZNddXg4f
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The qualified name of the database. This field has a value only when ResourceType is DATABASE and the downstream backfills the value. This field is provided for direct display on the frontend. For MySQL, this is the database name itself. For PostgreSQL or SQL Server, this is in the format of DatabaseName.SchemaName.
	//
	// example:
	//
	// finance.public
	DatabaseQualifiedName *string `json:"DatabaseQualifiedName,omitempty" xml:"DatabaseQualifiedName,omitempty"`
	// The UUID of the database. This field has a value only when ResourceType is DATABASE.
	//
	// example:
	//
	// md-HZ-vXR1ezNGjiDjiV13Gos1N
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The principal ID of the operator who registered this ownership relationship. In the "My Assets" scenario, the downstream does not return this field, and the value is null.
	//
	// example:
	//
	// usr_6ieggks7zuy6gpfdmgzsjdso
	GrantBy *string `json:"GrantBy,omitempty" xml:"GrantBy,omitempty"`
	// The source channel of the ownership. Valid values:
	//
	// - CONSOLE: Manually registered in the console.
	//
	// - Other values: Written by the system built-in ownership mechanism.
	//
	// In the "My Assets" scenario, the downstream does not return this field, and the value is null.
	//
	// example:
	//
	// CONSOLE
	GrantFrom *string `json:"GrantFrom,omitempty" xml:"GrantFrom,omitempty"`
	// The Owner principal ID. This is a gateway internal principal ID with the usr_ or agt_ prefix, not an Alibaba Cloud UID.
	//
	// example:
	//
	// usr_wlwp5a7uruanebg5bbdqqf5n
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The Owner principal type. Valid values:
	//
	// - USER: Human user.
	//
	// - AGENT: Managed Agent.
	//
	// example:
	//
	// USER
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// The ownership level. Valid values:
	//
	// - INSTANCE: Instance-level ownership. The coordinate contains only CatalogUuid.
	//
	// - DATABASE: Database-level ownership. The coordinate contains CatalogUuid + DatabaseUuid.
	//
	// example:
	//
	// DATABASE
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s AgenticResourceOwner) String() string {
	return dara.Prettify(s)
}

func (s AgenticResourceOwner) GoString() string {
	return s.String()
}

func (s *AgenticResourceOwner) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticResourceOwner) GetDatabaseQualifiedName() *string {
	return s.DatabaseQualifiedName
}

func (s *AgenticResourceOwner) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *AgenticResourceOwner) GetGrantBy() *string {
	return s.GrantBy
}

func (s *AgenticResourceOwner) GetGrantFrom() *string {
	return s.GrantFrom
}

func (s *AgenticResourceOwner) GetOwnerPrincipalId() *string {
	return s.OwnerPrincipalId
}

func (s *AgenticResourceOwner) GetOwnerPrincipalType() *string {
	return s.OwnerPrincipalType
}

func (s *AgenticResourceOwner) GetResourceType() *string {
	return s.ResourceType
}

func (s *AgenticResourceOwner) SetCatalogUuid(v string) *AgenticResourceOwner {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticResourceOwner) SetDatabaseQualifiedName(v string) *AgenticResourceOwner {
	s.DatabaseQualifiedName = &v
	return s
}

func (s *AgenticResourceOwner) SetDatabaseUuid(v string) *AgenticResourceOwner {
	s.DatabaseUuid = &v
	return s
}

func (s *AgenticResourceOwner) SetGrantBy(v string) *AgenticResourceOwner {
	s.GrantBy = &v
	return s
}

func (s *AgenticResourceOwner) SetGrantFrom(v string) *AgenticResourceOwner {
	s.GrantFrom = &v
	return s
}

func (s *AgenticResourceOwner) SetOwnerPrincipalId(v string) *AgenticResourceOwner {
	s.OwnerPrincipalId = &v
	return s
}

func (s *AgenticResourceOwner) SetOwnerPrincipalType(v string) *AgenticResourceOwner {
	s.OwnerPrincipalType = &v
	return s
}

func (s *AgenticResourceOwner) SetResourceType(v string) *AgenticResourceOwner {
	s.ResourceType = &v
	return s
}

func (s *AgenticResourceOwner) Validate() error {
	return dara.Validate(s)
}
