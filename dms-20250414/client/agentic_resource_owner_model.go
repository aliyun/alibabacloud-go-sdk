// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgenticResourceOwner interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogAlias(v string) *AgenticResourceOwner
	GetCatalogAlias() *string
	SetCatalogDataSourceType(v string) *AgenticResourceOwner
	GetCatalogDataSourceType() *string
	SetCatalogEnvType(v string) *AgenticResourceOwner
	GetCatalogEnvType() *string
	SetCatalogHost(v string) *AgenticResourceOwner
	GetCatalogHost() *string
	SetCatalogPort(v string) *AgenticResourceOwner
	GetCatalogPort() *string
	SetCatalogUuid(v string) *AgenticResourceOwner
	GetCatalogUuid() *string
	SetDataSourceUuid(v string) *AgenticResourceOwner
	GetDataSourceUuid() *string
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
	// The instance alias. This display property comes from the catalog to which the resource belongs and is used only for frontend display to align with the data permission list. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// Production database
	CatalogAlias *string `json:"CatalogAlias,omitempty" xml:"CatalogAlias,omitempty"`
	// The instance data source type, used only for frontend display. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// MYSQL
	CatalogDataSourceType *string `json:"CatalogDataSourceType,omitempty" xml:"CatalogDataSourceType,omitempty"`
	// The instance environment type, used only for frontend display. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// prod
	CatalogEnvType *string `json:"CatalogEnvType,omitempty" xml:"CatalogEnvType,omitempty"`
	// The instance host, used only for frontend display. Built-in sample data sources are masked as ******. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// rm-xxxxxxxx.mysql.rds.aliyuncs.com
	CatalogHost *string `json:"CatalogHost,omitempty" xml:"CatalogHost,omitempty"`
	// The instance port, used only for frontend display. Built-in sample data sources are masked as ******. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// 3306
	CatalogPort *string `json:"CatalogPort,omitempty" xml:"CatalogPort,omitempty"`
	// The UUID of the catalog to which the resource belongs.
	//
	// example:
	//
	// mc-HZ-5d9fbt8wW4AnGZNddXg4f
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the instance data source to which the resource belongs. It is used only for frontend display and instance locating. Non-sensitive connection information is not masked. It is populated only in the My Assets or Managed Resources query scenarios and is null in other scenarios.
	//
	// example:
	//
	// dsc-HZ-3f2b9a7c1d
	DataSourceUuid *string `json:"DataSourceUuid,omitempty" xml:"DataSourceUuid,omitempty"`
	// The fully qualified name of the database. It has a value only when ResourceType is DATABASE and is populated by the downstream service, which facilitates direct frontend display. For MySQL, it is the database name itself. For PostgreSQL or SQL Server, it is in the format of DatabaseName.SchemaName.
	//
	// example:
	//
	// finance.public
	DatabaseQualifiedName *string `json:"DatabaseQualifiedName,omitempty" xml:"DatabaseQualifiedName,omitempty"`
	// The UUID of the database. It has a value only when ResourceType is DATABASE.
	//
	// example:
	//
	// md-HZ-vXR1ezNGjiDjiV13Gos1N
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The principal ID of the operator who registered the ownership relationship. The downstream service does not return this field in the My Assets query scenario, and the value is null.
	//
	// example:
	//
	// usr_6ieggks7zuy6gpfdmgzsjdso
	GrantBy *string `json:"GrantBy,omitempty" xml:"GrantBy,omitempty"`
	// The source channel of the ownership. CONSOLE indicates manual registration in the console. Other values indicate that the ownership is written by the built-in system. The downstream service does not return this field in the My Assets query scenario, and the value is null.
	//
	// example:
	//
	// CONSOLE
	GrantFrom *string `json:"GrantFrom,omitempty" xml:"GrantFrom,omitempty"`
	// The principal ID of the owner. This is the internal principal ID of the gateway with a usr_ or agt_ prefix, not the Alibaba Cloud UID.
	//
	// example:
	//
	// usr_wlwp5a7uruanebg5bbdqqf5n
	OwnerPrincipalId *string `json:"OwnerPrincipalId,omitempty" xml:"OwnerPrincipalId,omitempty"`
	// The principal type of the owner. Valid values:
	//
	// - USER: Human user.
	//
	// - AGENT: Managed agent.
	//
	// example:
	//
	// USER
	OwnerPrincipalType *string `json:"OwnerPrincipalType,omitempty" xml:"OwnerPrincipalType,omitempty"`
	// The ownership level. Valid values:
	//
	// - INSTANCE: Instance-level ownership. The coordinate includes only CatalogUuid.
	//
	// - DATABASE: Database-level ownership. The coordinate includes CatalogUuid and DatabaseUuid.
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

func (s *AgenticResourceOwner) GetCatalogAlias() *string {
	return s.CatalogAlias
}

func (s *AgenticResourceOwner) GetCatalogDataSourceType() *string {
	return s.CatalogDataSourceType
}

func (s *AgenticResourceOwner) GetCatalogEnvType() *string {
	return s.CatalogEnvType
}

func (s *AgenticResourceOwner) GetCatalogHost() *string {
	return s.CatalogHost
}

func (s *AgenticResourceOwner) GetCatalogPort() *string {
	return s.CatalogPort
}

func (s *AgenticResourceOwner) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *AgenticResourceOwner) GetDataSourceUuid() *string {
	return s.DataSourceUuid
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

func (s *AgenticResourceOwner) SetCatalogAlias(v string) *AgenticResourceOwner {
	s.CatalogAlias = &v
	return s
}

func (s *AgenticResourceOwner) SetCatalogDataSourceType(v string) *AgenticResourceOwner {
	s.CatalogDataSourceType = &v
	return s
}

func (s *AgenticResourceOwner) SetCatalogEnvType(v string) *AgenticResourceOwner {
	s.CatalogEnvType = &v
	return s
}

func (s *AgenticResourceOwner) SetCatalogHost(v string) *AgenticResourceOwner {
	s.CatalogHost = &v
	return s
}

func (s *AgenticResourceOwner) SetCatalogPort(v string) *AgenticResourceOwner {
	s.CatalogPort = &v
	return s
}

func (s *AgenticResourceOwner) SetCatalogUuid(v string) *AgenticResourceOwner {
	s.CatalogUuid = &v
	return s
}

func (s *AgenticResourceOwner) SetDataSourceUuid(v string) *AgenticResourceOwner {
	s.DataSourceUuid = &v
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
