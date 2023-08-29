// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccessRequest struct {
	CatalogId          *string              `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	Principal          *Principal           `json:"Principal,omitempty" xml:"Principal,omitempty"`
	PrivilegeResources []*PrivilegeResource `json:"PrivilegeResources,omitempty" xml:"PrivilegeResources,omitempty" type:"Repeated"`
}

func (s AccessRequest) String() string {
	return tea.Prettify(s)
}

func (s AccessRequest) GoString() string {
	return s.String()
}

func (s *AccessRequest) SetCatalogId(v string) *AccessRequest {
	s.CatalogId = &v
	return s
}

func (s *AccessRequest) SetPrincipal(v *Principal) *AccessRequest {
	s.Principal = v
	return s
}

func (s *AccessRequest) SetPrivilegeResources(v []*PrivilegeResource) *AccessRequest {
	s.PrivilegeResources = v
	return s
}

type Catalog struct {
	CatalogId   *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy   *string `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocationUri *string `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Catalog) String() string {
	return tea.Prettify(s)
}

func (s Catalog) GoString() string {
	return s.String()
}

func (s *Catalog) SetCatalogId(v string) *Catalog {
	s.CatalogId = &v
	return s
}

func (s *Catalog) SetCreateTime(v int64) *Catalog {
	s.CreateTime = &v
	return s
}

func (s *Catalog) SetCreatedBy(v string) *Catalog {
	s.CreatedBy = &v
	return s
}

func (s *Catalog) SetDescription(v string) *Catalog {
	s.Description = &v
	return s
}

func (s *Catalog) SetLocationUri(v string) *Catalog {
	s.LocationUri = &v
	return s
}

func (s *Catalog) SetOwner(v string) *Catalog {
	s.Owner = &v
	return s
}

func (s *Catalog) SetStatus(v string) *Catalog {
	s.Status = &v
	return s
}

func (s *Catalog) SetUpdateTime(v int64) *Catalog {
	s.UpdateTime = &v
	return s
}

type CatalogInput struct {
	CatalogId   *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocationUri *string `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	Owner       *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s CatalogInput) String() string {
	return tea.Prettify(s)
}

func (s CatalogInput) GoString() string {
	return s.String()
}

func (s *CatalogInput) SetCatalogId(v string) *CatalogInput {
	s.CatalogId = &v
	return s
}

func (s *CatalogInput) SetDescription(v string) *CatalogInput {
	s.Description = &v
	return s
}

func (s *CatalogInput) SetLocationUri(v string) *CatalogInput {
	s.LocationUri = &v
	return s
}

func (s *CatalogInput) SetOwner(v string) *CatalogInput {
	s.Owner = &v
	return s
}

type CatalogResource struct {
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
}

func (s CatalogResource) String() string {
	return tea.Prettify(s)
}

func (s CatalogResource) GoString() string {
	return s.String()
}

func (s *CatalogResource) SetCatalogId(v string) *CatalogResource {
	s.CatalogId = &v
	return s
}

type CatalogSettings struct {
	Config map[string]*string `json:"Config,omitempty" xml:"Config,omitempty"`
}

func (s CatalogSettings) String() string {
	return tea.Prettify(s)
}

func (s CatalogSettings) GoString() string {
	return s.String()
}

func (s *CatalogSettings) SetConfig(v map[string]*string) *CatalogSettings {
	s.Config = v
	return s
}

type ColumnResource struct {
	ColumnNames  []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ColumnResource) String() string {
	return tea.Prettify(s)
}

func (s ColumnResource) GoString() string {
	return s.String()
}

func (s *ColumnResource) SetColumnNames(v []*string) *ColumnResource {
	s.ColumnNames = v
	return s
}

func (s *ColumnResource) SetDatabaseName(v string) *ColumnResource {
	s.DatabaseName = &v
	return s
}

func (s *ColumnResource) SetTableName(v string) *ColumnResource {
	s.TableName = &v
	return s
}

type ColumnStatistics struct {
	ColumnStatisticsDesc    *ColumnStatisticsDesc  `json:"ColumnStatisticsDesc,omitempty" xml:"ColumnStatisticsDesc,omitempty"`
	ColumnStatisticsObjList []*ColumnStatisticsObj `json:"ColumnStatisticsObjList,omitempty" xml:"ColumnStatisticsObjList,omitempty" type:"Repeated"`
	Engine                  *string                `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsStatsCompliant        *bool                  `json:"IsStatsCompliant,omitempty" xml:"IsStatsCompliant,omitempty"`
}

func (s ColumnStatistics) String() string {
	return tea.Prettify(s)
}

func (s ColumnStatistics) GoString() string {
	return s.String()
}

func (s *ColumnStatistics) SetColumnStatisticsDesc(v *ColumnStatisticsDesc) *ColumnStatistics {
	s.ColumnStatisticsDesc = v
	return s
}

func (s *ColumnStatistics) SetColumnStatisticsObjList(v []*ColumnStatisticsObj) *ColumnStatistics {
	s.ColumnStatisticsObjList = v
	return s
}

func (s *ColumnStatistics) SetEngine(v string) *ColumnStatistics {
	s.Engine = &v
	return s
}

func (s *ColumnStatistics) SetIsStatsCompliant(v bool) *ColumnStatistics {
	s.IsStatsCompliant = &v
	return s
}

type ColumnStatisticsDesc struct {
	LastAnalyzedTime *int64  `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	PartitionName    *string `json:"PartitionName,omitempty" xml:"PartitionName,omitempty"`
}

func (s ColumnStatisticsDesc) String() string {
	return tea.Prettify(s)
}

func (s ColumnStatisticsDesc) GoString() string {
	return s.String()
}

func (s *ColumnStatisticsDesc) SetLastAnalyzedTime(v int64) *ColumnStatisticsDesc {
	s.LastAnalyzedTime = &v
	return s
}

func (s *ColumnStatisticsDesc) SetPartitionName(v string) *ColumnStatisticsDesc {
	s.PartitionName = &v
	return s
}

type ColumnStatisticsObj struct {
	ColumnName           *string                                  `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	ColumnStatisticsData *ColumnStatisticsObjColumnStatisticsData `json:"ColumnStatisticsData,omitempty" xml:"ColumnStatisticsData,omitempty" type:"Struct"`
	ColumnType           *string                                  `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
}

func (s ColumnStatisticsObj) String() string {
	return tea.Prettify(s)
}

func (s ColumnStatisticsObj) GoString() string {
	return s.String()
}

func (s *ColumnStatisticsObj) SetColumnName(v string) *ColumnStatisticsObj {
	s.ColumnName = &v
	return s
}

func (s *ColumnStatisticsObj) SetColumnStatisticsData(v *ColumnStatisticsObjColumnStatisticsData) *ColumnStatisticsObj {
	s.ColumnStatisticsData = v
	return s
}

func (s *ColumnStatisticsObj) SetColumnType(v string) *ColumnStatisticsObj {
	s.ColumnType = &v
	return s
}

type ColumnStatisticsObjColumnStatisticsData struct {
	StatisticsData *string `json:"StatisticsData,omitempty" xml:"StatisticsData,omitempty"`
	StatisticsType *string `json:"StatisticsType,omitempty" xml:"StatisticsType,omitempty"`
}

func (s ColumnStatisticsObjColumnStatisticsData) String() string {
	return tea.Prettify(s)
}

func (s ColumnStatisticsObjColumnStatisticsData) GoString() string {
	return s.String()
}

func (s *ColumnStatisticsObjColumnStatisticsData) SetStatisticsData(v string) *ColumnStatisticsObjColumnStatisticsData {
	s.StatisticsData = &v
	return s
}

func (s *ColumnStatisticsObjColumnStatisticsData) SetStatisticsType(v string) *ColumnStatisticsObjColumnStatisticsData {
	s.StatisticsType = &v
	return s
}

type Database struct {
	CreateTime  *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy   *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	LocationUri *string                `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	Name        *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName   *string                `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerType   *string                `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters  map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Privileges  *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	UpdateTime  *int32                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Database) String() string {
	return tea.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) SetCreateTime(v int32) *Database {
	s.CreateTime = &v
	return s
}

func (s *Database) SetCreatedBy(v string) *Database {
	s.CreatedBy = &v
	return s
}

func (s *Database) SetDescription(v string) *Database {
	s.Description = &v
	return s
}

func (s *Database) SetLocationUri(v string) *Database {
	s.LocationUri = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetOwnerName(v string) *Database {
	s.OwnerName = &v
	return s
}

func (s *Database) SetOwnerType(v string) *Database {
	s.OwnerType = &v
	return s
}

func (s *Database) SetParameters(v map[string]*string) *Database {
	s.Parameters = v
	return s
}

func (s *Database) SetPrivileges(v *PrincipalPrivilegeSet) *Database {
	s.Privileges = v
	return s
}

func (s *Database) SetUpdateTime(v int32) *Database {
	s.UpdateTime = &v
	return s
}

type DatabaseInput struct {
	CreateTime  *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string                `json:"Description,omitempty" xml:"Description,omitempty"`
	LocationUri *string                `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	Name        *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerName   *string                `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerType   *string                `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters  map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Privileges  *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
}

func (s DatabaseInput) String() string {
	return tea.Prettify(s)
}

func (s DatabaseInput) GoString() string {
	return s.String()
}

func (s *DatabaseInput) SetCreateTime(v int32) *DatabaseInput {
	s.CreateTime = &v
	return s
}

func (s *DatabaseInput) SetDescription(v string) *DatabaseInput {
	s.Description = &v
	return s
}

func (s *DatabaseInput) SetLocationUri(v string) *DatabaseInput {
	s.LocationUri = &v
	return s
}

func (s *DatabaseInput) SetName(v string) *DatabaseInput {
	s.Name = &v
	return s
}

func (s *DatabaseInput) SetOwnerName(v string) *DatabaseInput {
	s.OwnerName = &v
	return s
}

func (s *DatabaseInput) SetOwnerType(v string) *DatabaseInput {
	s.OwnerType = &v
	return s
}

func (s *DatabaseInput) SetParameters(v map[string]*string) *DatabaseInput {
	s.Parameters = v
	return s
}

func (s *DatabaseInput) SetPrivileges(v *PrincipalPrivilegeSet) *DatabaseInput {
	s.Privileges = v
	return s
}

type DatabaseProfile struct {
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FileCnt    *int64  `json:"FileCnt,omitempty" xml:"FileCnt,omitempty"`
	FileSize   *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	LatestDate *string `json:"LatestDate,omitempty" xml:"LatestDate,omitempty"`
	Location   *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ObjectCnt  *int64  `json:"ObjectCnt,omitempty" xml:"ObjectCnt,omitempty"`
	ObjectSize *int64  `json:"ObjectSize,omitempty" xml:"ObjectSize,omitempty"`
}

func (s DatabaseProfile) String() string {
	return tea.Prettify(s)
}

func (s DatabaseProfile) GoString() string {
	return s.String()
}

func (s *DatabaseProfile) SetCreateTime(v string) *DatabaseProfile {
	s.CreateTime = &v
	return s
}

func (s *DatabaseProfile) SetFileCnt(v int64) *DatabaseProfile {
	s.FileCnt = &v
	return s
}

func (s *DatabaseProfile) SetFileSize(v int64) *DatabaseProfile {
	s.FileSize = &v
	return s
}

func (s *DatabaseProfile) SetLatestDate(v string) *DatabaseProfile {
	s.LatestDate = &v
	return s
}

func (s *DatabaseProfile) SetLocation(v string) *DatabaseProfile {
	s.Location = &v
	return s
}

func (s *DatabaseProfile) SetName(v string) *DatabaseProfile {
	s.Name = &v
	return s
}

func (s *DatabaseProfile) SetObjectCnt(v int64) *DatabaseProfile {
	s.ObjectCnt = &v
	return s
}

func (s *DatabaseProfile) SetObjectSize(v int64) *DatabaseProfile {
	s.ObjectSize = &v
	return s
}

type DatabaseResource struct {
	DatabaseName     *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DatabaseWildcard *string `json:"DatabaseWildcard,omitempty" xml:"DatabaseWildcard,omitempty"`
}

func (s DatabaseResource) String() string {
	return tea.Prettify(s)
}

func (s DatabaseResource) GoString() string {
	return s.String()
}

func (s *DatabaseResource) SetDatabaseName(v string) *DatabaseResource {
	s.DatabaseName = &v
	return s
}

func (s *DatabaseResource) SetDatabaseWildcard(v string) *DatabaseResource {
	s.DatabaseWildcard = &v
	return s
}

type DbStorageRank struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Quantity *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s DbStorageRank) String() string {
	return tea.Prettify(s)
}

func (s DbStorageRank) GoString() string {
	return s.String()
}

func (s *DbStorageRank) SetDbName(v string) *DbStorageRank {
	s.DbName = &v
	return s
}

func (s *DbStorageRank) SetQuantity(v int64) *DbStorageRank {
	s.Quantity = &v
	return s
}

type ErrorDetail struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s ErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s ErrorDetail) GoString() string {
	return s.String()
}

func (s *ErrorDetail) SetCode(v string) *ErrorDetail {
	s.Code = &v
	return s
}

func (s *ErrorDetail) SetMessage(v string) *ErrorDetail {
	s.Message = &v
	return s
}

type FieldSchema struct {
	Comment    *string            `json:"Comment,omitempty" xml:"Comment,omitempty"`
	Name       *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Type       *string            `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FieldSchema) String() string {
	return tea.Prettify(s)
}

func (s FieldSchema) GoString() string {
	return s.String()
}

func (s *FieldSchema) SetComment(v string) *FieldSchema {
	s.Comment = &v
	return s
}

func (s *FieldSchema) SetName(v string) *FieldSchema {
	s.Name = &v
	return s
}

func (s *FieldSchema) SetParameters(v map[string]*string) *FieldSchema {
	s.Parameters = v
	return s
}

func (s *FieldSchema) SetType(v string) *FieldSchema {
	s.Type = &v
	return s
}

type FileCnt struct {
	Large  *int64 `json:"Large,omitempty" xml:"Large,omitempty"`
	Middle *int64 `json:"Middle,omitempty" xml:"Middle,omitempty"`
	Small  *int64 `json:"Small,omitempty" xml:"Small,omitempty"`
	Tiny   *int64 `json:"Tiny,omitempty" xml:"Tiny,omitempty"`
}

func (s FileCnt) String() string {
	return tea.Prettify(s)
}

func (s FileCnt) GoString() string {
	return s.String()
}

func (s *FileCnt) SetLarge(v int64) *FileCnt {
	s.Large = &v
	return s
}

func (s *FileCnt) SetMiddle(v int64) *FileCnt {
	s.Middle = &v
	return s
}

func (s *FileCnt) SetSmall(v int64) *FileCnt {
	s.Small = &v
	return s
}

func (s *FileCnt) SetTiny(v int64) *FileCnt {
	s.Tiny = &v
	return s
}

type Function struct {
	CatalogId    *string        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ClassName    *string        `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	CreateTime   *int32         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy    *string        `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	DatabaseName *string        `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionName *string        `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType *string        `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OwnerName    *string        `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerType    *string        `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	ResourceUri  []*ResourceUri `json:"ResourceUri,omitempty" xml:"ResourceUri,omitempty" type:"Repeated"`
	UpdateTime   *int32         `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Function) String() string {
	return tea.Prettify(s)
}

func (s Function) GoString() string {
	return s.String()
}

func (s *Function) SetCatalogId(v string) *Function {
	s.CatalogId = &v
	return s
}

func (s *Function) SetClassName(v string) *Function {
	s.ClassName = &v
	return s
}

func (s *Function) SetCreateTime(v int32) *Function {
	s.CreateTime = &v
	return s
}

func (s *Function) SetCreatedBy(v string) *Function {
	s.CreatedBy = &v
	return s
}

func (s *Function) SetDatabaseName(v string) *Function {
	s.DatabaseName = &v
	return s
}

func (s *Function) SetFunctionName(v string) *Function {
	s.FunctionName = &v
	return s
}

func (s *Function) SetFunctionType(v string) *Function {
	s.FunctionType = &v
	return s
}

func (s *Function) SetOwnerName(v string) *Function {
	s.OwnerName = &v
	return s
}

func (s *Function) SetOwnerType(v string) *Function {
	s.OwnerType = &v
	return s
}

func (s *Function) SetResourceUri(v []*ResourceUri) *Function {
	s.ResourceUri = v
	return s
}

func (s *Function) SetUpdateTime(v int32) *Function {
	s.UpdateTime = &v
	return s
}

type FunctionInput struct {
	ClassName    *string        `json:"ClassName,omitempty" xml:"ClassName,omitempty"`
	CreateTime   *int32         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	FunctionName *string        `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	FunctionType *string        `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	OwnerName    *string        `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	OwnerType    *string        `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	ResourceUri  []*ResourceUri `json:"ResourceUri,omitempty" xml:"ResourceUri,omitempty" type:"Repeated"`
}

func (s FunctionInput) String() string {
	return tea.Prettify(s)
}

func (s FunctionInput) GoString() string {
	return s.String()
}

func (s *FunctionInput) SetClassName(v string) *FunctionInput {
	s.ClassName = &v
	return s
}

func (s *FunctionInput) SetCreateTime(v int32) *FunctionInput {
	s.CreateTime = &v
	return s
}

func (s *FunctionInput) SetFunctionName(v string) *FunctionInput {
	s.FunctionName = &v
	return s
}

func (s *FunctionInput) SetFunctionType(v string) *FunctionInput {
	s.FunctionType = &v
	return s
}

func (s *FunctionInput) SetOwnerName(v string) *FunctionInput {
	s.OwnerName = &v
	return s
}

func (s *FunctionInput) SetOwnerType(v string) *FunctionInput {
	s.OwnerType = &v
	return s
}

func (s *FunctionInput) SetResourceUri(v []*ResourceUri) *FunctionInput {
	s.ResourceUri = v
	return s
}

type FunctionResource struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s FunctionResource) String() string {
	return tea.Prettify(s)
}

func (s FunctionResource) GoString() string {
	return s.String()
}

func (s *FunctionResource) SetDatabaseName(v string) *FunctionResource {
	s.DatabaseName = &v
	return s
}

func (s *FunctionResource) SetFunctionName(v string) *FunctionResource {
	s.FunctionName = &v
	return s
}

type GrantRevokeEntry struct {
	Accesses         []*string     `json:"Accesses,omitempty" xml:"Accesses,omitempty" type:"Repeated"`
	DelegateAccesses []*string     `json:"DelegateAccesses,omitempty" xml:"DelegateAccesses,omitempty" type:"Repeated"`
	Id               *string       `json:"Id,omitempty" xml:"Id,omitempty"`
	MetaResource     *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	Principal        *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
}

func (s GrantRevokeEntry) String() string {
	return tea.Prettify(s)
}

func (s GrantRevokeEntry) GoString() string {
	return s.String()
}

func (s *GrantRevokeEntry) SetAccesses(v []*string) *GrantRevokeEntry {
	s.Accesses = v
	return s
}

func (s *GrantRevokeEntry) SetDelegateAccesses(v []*string) *GrantRevokeEntry {
	s.DelegateAccesses = v
	return s
}

func (s *GrantRevokeEntry) SetId(v string) *GrantRevokeEntry {
	s.Id = &v
	return s
}

func (s *GrantRevokeEntry) SetMetaResource(v *MetaResource) *GrantRevokeEntry {
	s.MetaResource = v
	return s
}

func (s *GrantRevokeEntry) SetPrincipal(v *Principal) *GrantRevokeEntry {
	s.Principal = v
	return s
}

type GrantRevokeFailureEntry struct {
	ErrorDetail      *ErrorDetail      `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	GrantRevokeEntry *GrantRevokeEntry `json:"GrantRevokeEntry,omitempty" xml:"GrantRevokeEntry,omitempty"`
}

func (s GrantRevokeFailureEntry) String() string {
	return tea.Prettify(s)
}

func (s GrantRevokeFailureEntry) GoString() string {
	return s.String()
}

func (s *GrantRevokeFailureEntry) SetErrorDetail(v *ErrorDetail) *GrantRevokeFailureEntry {
	s.ErrorDetail = v
	return s
}

func (s *GrantRevokeFailureEntry) SetGrantRevokeEntry(v *GrantRevokeEntry) *GrantRevokeFailureEntry {
	s.GrantRevokeEntry = v
	return s
}

type HighLight struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s HighLight) String() string {
	return tea.Prettify(s)
}

func (s HighLight) GoString() string {
	return s.String()
}

func (s *HighLight) SetKey(v string) *HighLight {
	s.Key = &v
	return s
}

func (s *HighLight) SetValue(v string) *HighLight {
	s.Value = &v
	return s
}

type IndicatorStatistic struct {
	Data *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
}

func (s IndicatorStatistic) String() string {
	return tea.Prettify(s)
}

func (s IndicatorStatistic) GoString() string {
	return s.String()
}

func (s *IndicatorStatistic) SetData(v int64) *IndicatorStatistic {
	s.Data = &v
	return s
}

func (s *IndicatorStatistic) SetDate(v string) *IndicatorStatistic {
	s.Date = &v
	return s
}

type LifecycleResource struct {
	BizId              *string                    `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CatalogId          *string                    `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	Database           *LifecycleResourceDatabase `json:"Database,omitempty" xml:"Database,omitempty" type:"Struct"`
	DatabaseName       *string                    `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DatabaseProfile    *DatabaseProfile           `json:"DatabaseProfile,omitempty" xml:"DatabaseProfile,omitempty"`
	GmtCreate          *string                    `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	LifecycleRuleBizId *string                    `json:"LifecycleRuleBizId,omitempty" xml:"LifecycleRuleBizId,omitempty"`
	Owner              *int64                     `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Table              *LifecycleResourceTable    `json:"Table,omitempty" xml:"Table,omitempty" type:"Struct"`
	TableName          *string                    `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableProfile       *TableProfile              `json:"TableProfile,omitempty" xml:"TableProfile,omitempty"`
}

func (s LifecycleResource) String() string {
	return tea.Prettify(s)
}

func (s LifecycleResource) GoString() string {
	return s.String()
}

func (s *LifecycleResource) SetBizId(v string) *LifecycleResource {
	s.BizId = &v
	return s
}

func (s *LifecycleResource) SetCatalogId(v string) *LifecycleResource {
	s.CatalogId = &v
	return s
}

func (s *LifecycleResource) SetDatabase(v *LifecycleResourceDatabase) *LifecycleResource {
	s.Database = v
	return s
}

func (s *LifecycleResource) SetDatabaseName(v string) *LifecycleResource {
	s.DatabaseName = &v
	return s
}

func (s *LifecycleResource) SetDatabaseProfile(v *DatabaseProfile) *LifecycleResource {
	s.DatabaseProfile = v
	return s
}

func (s *LifecycleResource) SetGmtCreate(v string) *LifecycleResource {
	s.GmtCreate = &v
	return s
}

func (s *LifecycleResource) SetLifecycleRuleBizId(v string) *LifecycleResource {
	s.LifecycleRuleBizId = &v
	return s
}

func (s *LifecycleResource) SetOwner(v int64) *LifecycleResource {
	s.Owner = &v
	return s
}

func (s *LifecycleResource) SetTable(v *LifecycleResourceTable) *LifecycleResource {
	s.Table = v
	return s
}

func (s *LifecycleResource) SetTableName(v string) *LifecycleResource {
	s.TableName = &v
	return s
}

func (s *LifecycleResource) SetTableProfile(v *TableProfile) *LifecycleResource {
	s.TableProfile = v
	return s
}

type LifecycleResourceDatabase struct {
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocationUri *string `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	UpdateTime  *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s LifecycleResourceDatabase) String() string {
	return tea.Prettify(s)
}

func (s LifecycleResourceDatabase) GoString() string {
	return s.String()
}

func (s *LifecycleResourceDatabase) SetCreateTime(v int64) *LifecycleResourceDatabase {
	s.CreateTime = &v
	return s
}

func (s *LifecycleResourceDatabase) SetDescription(v string) *LifecycleResourceDatabase {
	s.Description = &v
	return s
}

func (s *LifecycleResourceDatabase) SetLocationUri(v string) *LifecycleResourceDatabase {
	s.LocationUri = &v
	return s
}

func (s *LifecycleResourceDatabase) SetName(v string) *LifecycleResourceDatabase {
	s.Name = &v
	return s
}

func (s *LifecycleResourceDatabase) SetUpdateTime(v int64) *LifecycleResourceDatabase {
	s.UpdateTime = &v
	return s
}

type LifecycleResourceTable struct {
	CreateTime   *int64                    `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName *string                   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Parameters   map[string]*string        `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Sd           *LifecycleResourceTableSd `json:"Sd,omitempty" xml:"Sd,omitempty" type:"Struct"`
	TableName    *string                   `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType    *string                   `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s LifecycleResourceTable) String() string {
	return tea.Prettify(s)
}

func (s LifecycleResourceTable) GoString() string {
	return s.String()
}

func (s *LifecycleResourceTable) SetCreateTime(v int64) *LifecycleResourceTable {
	s.CreateTime = &v
	return s
}

func (s *LifecycleResourceTable) SetDatabaseName(v string) *LifecycleResourceTable {
	s.DatabaseName = &v
	return s
}

func (s *LifecycleResourceTable) SetParameters(v map[string]*string) *LifecycleResourceTable {
	s.Parameters = v
	return s
}

func (s *LifecycleResourceTable) SetSd(v *LifecycleResourceTableSd) *LifecycleResourceTable {
	s.Sd = v
	return s
}

func (s *LifecycleResourceTable) SetTableName(v string) *LifecycleResourceTable {
	s.TableName = &v
	return s
}

func (s *LifecycleResourceTable) SetTableType(v string) *LifecycleResourceTable {
	s.TableType = &v
	return s
}

type LifecycleResourceTableSd struct {
	BucketCols   []*string                          `json:"BucketCols,omitempty" xml:"BucketCols,omitempty" type:"Repeated"`
	InputFormat  *string                            `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	Location     *string                            `json:"Location,omitempty" xml:"Location,omitempty"`
	OutputFormat *string                            `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters   map[string]*string                 `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerDeInfo    *LifecycleResourceTableSdSerDeInfo `json:"SerDeInfo,omitempty" xml:"SerDeInfo,omitempty" type:"Struct"`
}

func (s LifecycleResourceTableSd) String() string {
	return tea.Prettify(s)
}

func (s LifecycleResourceTableSd) GoString() string {
	return s.String()
}

func (s *LifecycleResourceTableSd) SetBucketCols(v []*string) *LifecycleResourceTableSd {
	s.BucketCols = v
	return s
}

func (s *LifecycleResourceTableSd) SetInputFormat(v string) *LifecycleResourceTableSd {
	s.InputFormat = &v
	return s
}

func (s *LifecycleResourceTableSd) SetLocation(v string) *LifecycleResourceTableSd {
	s.Location = &v
	return s
}

func (s *LifecycleResourceTableSd) SetOutputFormat(v string) *LifecycleResourceTableSd {
	s.OutputFormat = &v
	return s
}

func (s *LifecycleResourceTableSd) SetParameters(v map[string]*string) *LifecycleResourceTableSd {
	s.Parameters = v
	return s
}

func (s *LifecycleResourceTableSd) SetSerDeInfo(v *LifecycleResourceTableSdSerDeInfo) *LifecycleResourceTableSd {
	s.SerDeInfo = v
	return s
}

type LifecycleResourceTableSdSerDeInfo struct {
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerializationLib *string            `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
}

func (s LifecycleResourceTableSdSerDeInfo) String() string {
	return tea.Prettify(s)
}

func (s LifecycleResourceTableSdSerDeInfo) GoString() string {
	return s.String()
}

func (s *LifecycleResourceTableSdSerDeInfo) SetName(v string) *LifecycleResourceTableSdSerDeInfo {
	s.Name = &v
	return s
}

func (s *LifecycleResourceTableSdSerDeInfo) SetParameters(v map[string]*string) *LifecycleResourceTableSdSerDeInfo {
	s.Parameters = v
	return s
}

func (s *LifecycleResourceTableSdSerDeInfo) SetSerializationLib(v string) *LifecycleResourceTableSdSerDeInfo {
	s.SerializationLib = &v
	return s
}

type LifecycleRule struct {
	ArchiveDays      *int32            `json:"ArchiveDays,omitempty" xml:"ArchiveDays,omitempty"`
	BindCount        *int32            `json:"BindCount,omitempty" xml:"BindCount,omitempty"`
	BizId            *string           `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CatalogId        *string           `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColdArchiveDays  *int32            `json:"ColdArchiveDays,omitempty" xml:"ColdArchiveDays,omitempty"`
	Config           *string           `json:"Config,omitempty" xml:"Config,omitempty"`
	Description      *string           `json:"Description,omitempty" xml:"Description,omitempty"`
	GmtCreate        *string           `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified      *string           `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	IaDays           *int32            `json:"IaDays,omitempty" xml:"IaDays,omitempty"`
	Name             *string           `json:"Name,omitempty" xml:"Name,omitempty"`
	ResourceType     *string           `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RuleType         *string           `json:"RuleType,omitempty" xml:"RuleType,omitempty"`
	ScheduleStatus   *string           `json:"ScheduleStatus,omitempty" xml:"ScheduleStatus,omitempty"`
	Workflow         *Workflow         `json:"Workflow,omitempty" xml:"Workflow,omitempty"`
	WorkflowId       *string           `json:"WorkflowId,omitempty" xml:"WorkflowId,omitempty"`
	WorkflowInstance *WorkflowInstance `json:"WorkflowInstance,omitempty" xml:"WorkflowInstance,omitempty"`
}

func (s LifecycleRule) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRule) GoString() string {
	return s.String()
}

func (s *LifecycleRule) SetArchiveDays(v int32) *LifecycleRule {
	s.ArchiveDays = &v
	return s
}

func (s *LifecycleRule) SetBindCount(v int32) *LifecycleRule {
	s.BindCount = &v
	return s
}

func (s *LifecycleRule) SetBizId(v string) *LifecycleRule {
	s.BizId = &v
	return s
}

func (s *LifecycleRule) SetCatalogId(v string) *LifecycleRule {
	s.CatalogId = &v
	return s
}

func (s *LifecycleRule) SetColdArchiveDays(v int32) *LifecycleRule {
	s.ColdArchiveDays = &v
	return s
}

func (s *LifecycleRule) SetConfig(v string) *LifecycleRule {
	s.Config = &v
	return s
}

func (s *LifecycleRule) SetDescription(v string) *LifecycleRule {
	s.Description = &v
	return s
}

func (s *LifecycleRule) SetGmtCreate(v string) *LifecycleRule {
	s.GmtCreate = &v
	return s
}

func (s *LifecycleRule) SetGmtModified(v string) *LifecycleRule {
	s.GmtModified = &v
	return s
}

func (s *LifecycleRule) SetIaDays(v int32) *LifecycleRule {
	s.IaDays = &v
	return s
}

func (s *LifecycleRule) SetName(v string) *LifecycleRule {
	s.Name = &v
	return s
}

func (s *LifecycleRule) SetResourceType(v string) *LifecycleRule {
	s.ResourceType = &v
	return s
}

func (s *LifecycleRule) SetRuleType(v string) *LifecycleRule {
	s.RuleType = &v
	return s
}

func (s *LifecycleRule) SetScheduleStatus(v string) *LifecycleRule {
	s.ScheduleStatus = &v
	return s
}

func (s *LifecycleRule) SetWorkflow(v *Workflow) *LifecycleRule {
	s.Workflow = v
	return s
}

func (s *LifecycleRule) SetWorkflowId(v string) *LifecycleRule {
	s.WorkflowId = &v
	return s
}

func (s *LifecycleRule) SetWorkflowInstance(v *WorkflowInstance) *LifecycleRule {
	s.WorkflowInstance = v
	return s
}

type LifecycleTask struct {
	BizId            *string           `json:"BizId,omitempty" xml:"BizId,omitempty"`
	LifecycleRule    *LifecycleRule    `json:"LifecycleRule,omitempty" xml:"LifecycleRule,omitempty"`
	Name             *string           `json:"Name,omitempty" xml:"Name,omitempty"`
	WorkflowInstance *WorkflowInstance `json:"WorkflowInstance,omitempty" xml:"WorkflowInstance,omitempty"`
}

func (s LifecycleTask) String() string {
	return tea.Prettify(s)
}

func (s LifecycleTask) GoString() string {
	return s.String()
}

func (s *LifecycleTask) SetBizId(v string) *LifecycleTask {
	s.BizId = &v
	return s
}

func (s *LifecycleTask) SetLifecycleRule(v *LifecycleRule) *LifecycleTask {
	s.LifecycleRule = v
	return s
}

func (s *LifecycleTask) SetName(v string) *LifecycleTask {
	s.Name = &v
	return s
}

func (s *LifecycleTask) SetWorkflowInstance(v *WorkflowInstance) *LifecycleTask {
	s.WorkflowInstance = v
	return s
}

type LocationStorageRankDTO struct {
	FileCnt  *int64  `json:"FileCnt,omitempty" xml:"FileCnt,omitempty"`
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Storage  *int64  `json:"Storage,omitempty" xml:"Storage,omitempty"`
}

func (s LocationStorageRankDTO) String() string {
	return tea.Prettify(s)
}

func (s LocationStorageRankDTO) GoString() string {
	return s.String()
}

func (s *LocationStorageRankDTO) SetFileCnt(v int64) *LocationStorageRankDTO {
	s.FileCnt = &v
	return s
}

func (s *LocationStorageRankDTO) SetLocation(v string) *LocationStorageRankDTO {
	s.Location = &v
	return s
}

func (s *LocationStorageRankDTO) SetStorage(v int64) *LocationStorageRankDTO {
	s.Storage = &v
	return s
}

type LockObj struct {
	CatalogId     *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName  *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionName *string `json:"PartitionName,omitempty" xml:"PartitionName,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s LockObj) String() string {
	return tea.Prettify(s)
}

func (s LockObj) GoString() string {
	return s.String()
}

func (s *LockObj) SetCatalogId(v string) *LockObj {
	s.CatalogId = &v
	return s
}

func (s *LockObj) SetDatabaseName(v string) *LockObj {
	s.DatabaseName = &v
	return s
}

func (s *LockObj) SetPartitionName(v string) *LockObj {
	s.PartitionName = &v
	return s
}

func (s *LockObj) SetTableName(v string) *LockObj {
	s.TableName = &v
	return s
}

type LockStatus struct {
	LockId    *int64  `json:"LockId,omitempty" xml:"LockId,omitempty"`
	LockState *string `json:"LockState,omitempty" xml:"LockState,omitempty"`
}

func (s LockStatus) String() string {
	return tea.Prettify(s)
}

func (s LockStatus) GoString() string {
	return s.String()
}

func (s *LockStatus) SetLockId(v int64) *LockStatus {
	s.LockId = &v
	return s
}

func (s *LockStatus) SetLockState(v string) *LockStatus {
	s.LockState = &v
	return s
}

type LogInfo struct {
	BizTime    *string `json:"BizTime,omitempty" xml:"BizTime,omitempty"`
	GmtCreate  *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
	LogId      *string `json:"LogId,omitempty" xml:"LogId,omitempty"`
	LogSummary *string `json:"LogSummary,omitempty" xml:"LogSummary,omitempty"`
	LogType    *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
}

func (s LogInfo) String() string {
	return tea.Prettify(s)
}

func (s LogInfo) GoString() string {
	return s.String()
}

func (s *LogInfo) SetBizTime(v string) *LogInfo {
	s.BizTime = &v
	return s
}

func (s *LogInfo) SetGmtCreate(v string) *LogInfo {
	s.GmtCreate = &v
	return s
}

func (s *LogInfo) SetInstanceId(v string) *LogInfo {
	s.InstanceId = &v
	return s
}

func (s *LogInfo) SetLogContent(v string) *LogInfo {
	s.LogContent = &v
	return s
}

func (s *LogInfo) SetLogId(v string) *LogInfo {
	s.LogId = &v
	return s
}

func (s *LogInfo) SetLogSummary(v string) *LogInfo {
	s.LogSummary = &v
	return s
}

func (s *LogInfo) SetLogType(v string) *LogInfo {
	s.LogType = &v
	return s
}

type MetaResource struct {
	CatalogResource  *CatalogResource  `json:"CatalogResource,omitempty" xml:"CatalogResource,omitempty"`
	ColumnResource   *ColumnResource   `json:"ColumnResource,omitempty" xml:"ColumnResource,omitempty"`
	DatabaseResource *DatabaseResource `json:"DatabaseResource,omitempty" xml:"DatabaseResource,omitempty"`
	FunctionResource *FunctionResource `json:"FunctionResource,omitempty" xml:"FunctionResource,omitempty"`
	ResourceType     *string           `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TableResource    *TableResource    `json:"TableResource,omitempty" xml:"TableResource,omitempty"`
}

func (s MetaResource) String() string {
	return tea.Prettify(s)
}

func (s MetaResource) GoString() string {
	return s.String()
}

func (s *MetaResource) SetCatalogResource(v *CatalogResource) *MetaResource {
	s.CatalogResource = v
	return s
}

func (s *MetaResource) SetColumnResource(v *ColumnResource) *MetaResource {
	s.ColumnResource = v
	return s
}

func (s *MetaResource) SetDatabaseResource(v *DatabaseResource) *MetaResource {
	s.DatabaseResource = v
	return s
}

func (s *MetaResource) SetFunctionResource(v *FunctionResource) *MetaResource {
	s.FunctionResource = v
	return s
}

func (s *MetaResource) SetResourceType(v string) *MetaResource {
	s.ResourceType = &v
	return s
}

func (s *MetaResource) SetTableResource(v *TableResource) *MetaResource {
	s.TableResource = v
	return s
}

type Order struct {
	Col   *string `json:"Col,omitempty" xml:"Col,omitempty"`
	Order *int32  `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s Order) String() string {
	return tea.Prettify(s)
}

func (s Order) GoString() string {
	return s.String()
}

func (s *Order) SetCol(v string) *Order {
	s.Col = &v
	return s
}

func (s *Order) SetOrder(v int32) *Order {
	s.Order = &v
	return s
}

type Partition struct {
	CreateTime       *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName     *string                `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LastAccessTime   *int32                 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastAnalyzedTime *int32                 `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	Parameters       map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Privileges       *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	Sd               *StorageDescriptor     `json:"Sd,omitempty" xml:"Sd,omitempty"`
	TableName        *string                `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Values           []*string              `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s Partition) String() string {
	return tea.Prettify(s)
}

func (s Partition) GoString() string {
	return s.String()
}

func (s *Partition) SetCreateTime(v int32) *Partition {
	s.CreateTime = &v
	return s
}

func (s *Partition) SetDatabaseName(v string) *Partition {
	s.DatabaseName = &v
	return s
}

func (s *Partition) SetLastAccessTime(v int32) *Partition {
	s.LastAccessTime = &v
	return s
}

func (s *Partition) SetLastAnalyzedTime(v int32) *Partition {
	s.LastAnalyzedTime = &v
	return s
}

func (s *Partition) SetParameters(v map[string]*string) *Partition {
	s.Parameters = v
	return s
}

func (s *Partition) SetPrivileges(v *PrincipalPrivilegeSet) *Partition {
	s.Privileges = v
	return s
}

func (s *Partition) SetSd(v *StorageDescriptor) *Partition {
	s.Sd = v
	return s
}

func (s *Partition) SetTableName(v string) *Partition {
	s.TableName = &v
	return s
}

func (s *Partition) SetValues(v []*string) *Partition {
	s.Values = v
	return s
}

type PartitionError struct {
	ErrorDetail     *ErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	PartitionValues []*string    `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
}

func (s PartitionError) String() string {
	return tea.Prettify(s)
}

func (s PartitionError) GoString() string {
	return s.String()
}

func (s *PartitionError) SetErrorDetail(v *ErrorDetail) *PartitionError {
	s.ErrorDetail = v
	return s
}

func (s *PartitionError) SetPartitionValues(v []*string) *PartitionError {
	s.PartitionValues = v
	return s
}

type PartitionInput struct {
	CreateTime       *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DatabaseName     *string                `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LastAccessTime   *int32                 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastAnalyzedTime *int32                 `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	Parameters       map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	Privileges       *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	Sd               *StorageDescriptor     `json:"Sd,omitempty" xml:"Sd,omitempty"`
	TableName        *string                `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Values           []*string              `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s PartitionInput) String() string {
	return tea.Prettify(s)
}

func (s PartitionInput) GoString() string {
	return s.String()
}

func (s *PartitionInput) SetCreateTime(v int32) *PartitionInput {
	s.CreateTime = &v
	return s
}

func (s *PartitionInput) SetDatabaseName(v string) *PartitionInput {
	s.DatabaseName = &v
	return s
}

func (s *PartitionInput) SetLastAccessTime(v int32) *PartitionInput {
	s.LastAccessTime = &v
	return s
}

func (s *PartitionInput) SetLastAnalyzedTime(v int32) *PartitionInput {
	s.LastAnalyzedTime = &v
	return s
}

func (s *PartitionInput) SetParameters(v map[string]*string) *PartitionInput {
	s.Parameters = v
	return s
}

func (s *PartitionInput) SetPrivileges(v *PrincipalPrivilegeSet) *PartitionInput {
	s.Privileges = v
	return s
}

func (s *PartitionInput) SetSd(v *StorageDescriptor) *PartitionInput {
	s.Sd = v
	return s
}

func (s *PartitionInput) SetTableName(v string) *PartitionInput {
	s.TableName = &v
	return s
}

func (s *PartitionInput) SetValues(v []*string) *PartitionInput {
	s.Values = v
	return s
}

type PartitionProfile struct {
	AccessNum              *int64  `json:"AccessNum,omitempty" xml:"AccessNum,omitempty"`
	AccessNumMonthly       *int64  `json:"AccessNumMonthly,omitempty" xml:"AccessNumMonthly,omitempty"`
	AccessNumWeekly        *int64  `json:"AccessNumWeekly,omitempty" xml:"AccessNumWeekly,omitempty"`
	ArchiveStatus          *string `json:"ArchiveStatus,omitempty" xml:"ArchiveStatus,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSourceType         *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DatabaseName           *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FileCnt                *int64  `json:"FileCnt,omitempty" xml:"FileCnt,omitempty"`
	FileSize               *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	LastAccessNumTime      *string `json:"LastAccessNumTime,omitempty" xml:"LastAccessNumTime,omitempty"`
	LastAccessTime         *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastModifyTime         *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	Location               *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ObjectAccessNum        *int64  `json:"ObjectAccessNum,omitempty" xml:"ObjectAccessNum,omitempty"`
	ObjectAccessNumMonthly *int64  `json:"ObjectAccessNumMonthly,omitempty" xml:"ObjectAccessNumMonthly,omitempty"`
	ObjectAccessNumWeekly  *int64  `json:"ObjectAccessNumWeekly,omitempty" xml:"ObjectAccessNumWeekly,omitempty"`
	ObjectCnt              *int64  `json:"ObjectCnt,omitempty" xml:"ObjectCnt,omitempty"`
	ObjectSize             *int64  `json:"ObjectSize,omitempty" xml:"ObjectSize,omitempty"`
	PartitionName          *string `json:"PartitionName,omitempty" xml:"PartitionName,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s PartitionProfile) String() string {
	return tea.Prettify(s)
}

func (s PartitionProfile) GoString() string {
	return s.String()
}

func (s *PartitionProfile) SetAccessNum(v int64) *PartitionProfile {
	s.AccessNum = &v
	return s
}

func (s *PartitionProfile) SetAccessNumMonthly(v int64) *PartitionProfile {
	s.AccessNumMonthly = &v
	return s
}

func (s *PartitionProfile) SetAccessNumWeekly(v int64) *PartitionProfile {
	s.AccessNumWeekly = &v
	return s
}

func (s *PartitionProfile) SetArchiveStatus(v string) *PartitionProfile {
	s.ArchiveStatus = &v
	return s
}

func (s *PartitionProfile) SetCreateTime(v string) *PartitionProfile {
	s.CreateTime = &v
	return s
}

func (s *PartitionProfile) SetDataSourceType(v string) *PartitionProfile {
	s.DataSourceType = &v
	return s
}

func (s *PartitionProfile) SetDatabaseName(v string) *PartitionProfile {
	s.DatabaseName = &v
	return s
}

func (s *PartitionProfile) SetFileCnt(v int64) *PartitionProfile {
	s.FileCnt = &v
	return s
}

func (s *PartitionProfile) SetFileSize(v int64) *PartitionProfile {
	s.FileSize = &v
	return s
}

func (s *PartitionProfile) SetLastAccessNumTime(v string) *PartitionProfile {
	s.LastAccessNumTime = &v
	return s
}

func (s *PartitionProfile) SetLastAccessTime(v string) *PartitionProfile {
	s.LastAccessTime = &v
	return s
}

func (s *PartitionProfile) SetLastModifyTime(v string) *PartitionProfile {
	s.LastModifyTime = &v
	return s
}

func (s *PartitionProfile) SetLocation(v string) *PartitionProfile {
	s.Location = &v
	return s
}

func (s *PartitionProfile) SetObjectAccessNum(v int64) *PartitionProfile {
	s.ObjectAccessNum = &v
	return s
}

func (s *PartitionProfile) SetObjectAccessNumMonthly(v int64) *PartitionProfile {
	s.ObjectAccessNumMonthly = &v
	return s
}

func (s *PartitionProfile) SetObjectAccessNumWeekly(v int64) *PartitionProfile {
	s.ObjectAccessNumWeekly = &v
	return s
}

func (s *PartitionProfile) SetObjectCnt(v int64) *PartitionProfile {
	s.ObjectCnt = &v
	return s
}

func (s *PartitionProfile) SetObjectSize(v int64) *PartitionProfile {
	s.ObjectSize = &v
	return s
}

func (s *PartitionProfile) SetPartitionName(v string) *PartitionProfile {
	s.PartitionName = &v
	return s
}

func (s *PartitionProfile) SetTableName(v string) *PartitionProfile {
	s.TableName = &v
	return s
}

type PartitionSpec struct {
	SharedSDPartitions      []*Partition                          `json:"SharedSDPartitions,omitempty" xml:"SharedSDPartitions,omitempty" type:"Repeated"`
	SharedStorageDescriptor *PartitionSpecSharedStorageDescriptor `json:"SharedStorageDescriptor,omitempty" xml:"SharedStorageDescriptor,omitempty" type:"Struct"`
}

func (s PartitionSpec) String() string {
	return tea.Prettify(s)
}

func (s PartitionSpec) GoString() string {
	return s.String()
}

func (s *PartitionSpec) SetSharedSDPartitions(v []*Partition) *PartitionSpec {
	s.SharedSDPartitions = v
	return s
}

func (s *PartitionSpec) SetSharedStorageDescriptor(v *PartitionSpecSharedStorageDescriptor) *PartitionSpec {
	s.SharedStorageDescriptor = v
	return s
}

type PartitionSpecSharedStorageDescriptor struct {
	Cols     []*FieldSchema `json:"Cols,omitempty" xml:"Cols,omitempty" type:"Repeated"`
	Location *string        `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s PartitionSpecSharedStorageDescriptor) String() string {
	return tea.Prettify(s)
}

func (s PartitionSpecSharedStorageDescriptor) GoString() string {
	return s.String()
}

func (s *PartitionSpecSharedStorageDescriptor) SetCols(v []*FieldSchema) *PartitionSpecSharedStorageDescriptor {
	s.Cols = v
	return s
}

func (s *PartitionSpecSharedStorageDescriptor) SetLocation(v string) *PartitionSpecSharedStorageDescriptor {
	s.Location = &v
	return s
}

type Principal struct {
	PrincipalArn *string `json:"PrincipalArn,omitempty" xml:"PrincipalArn,omitempty"`
}

func (s Principal) String() string {
	return tea.Prettify(s)
}

func (s Principal) GoString() string {
	return s.String()
}

func (s *Principal) SetPrincipalArn(v string) *Principal {
	s.PrincipalArn = &v
	return s
}

type PrincipalPrivilegeSet struct {
	GroupPrivileges map[string][]*PrivilegeGrantInfo `json:"GroupPrivileges,omitempty" xml:"GroupPrivileges,omitempty"`
	RolePrivileges  map[string][]*PrivilegeGrantInfo `json:"RolePrivileges,omitempty" xml:"RolePrivileges,omitempty"`
	UserPrivileges  map[string][]*PrivilegeGrantInfo `json:"UserPrivileges,omitempty" xml:"UserPrivileges,omitempty"`
}

func (s PrincipalPrivilegeSet) String() string {
	return tea.Prettify(s)
}

func (s PrincipalPrivilegeSet) GoString() string {
	return s.String()
}

func (s *PrincipalPrivilegeSet) SetGroupPrivileges(v map[string][]*PrivilegeGrantInfo) *PrincipalPrivilegeSet {
	s.GroupPrivileges = v
	return s
}

func (s *PrincipalPrivilegeSet) SetRolePrivileges(v map[string][]*PrivilegeGrantInfo) *PrincipalPrivilegeSet {
	s.RolePrivileges = v
	return s
}

func (s *PrincipalPrivilegeSet) SetUserPrivileges(v map[string][]*PrivilegeGrantInfo) *PrincipalPrivilegeSet {
	s.UserPrivileges = v
	return s
}

type PrincipalResourcePermissions struct {
	Accesses         []*string     `json:"Accesses,omitempty" xml:"Accesses,omitempty" type:"Repeated"`
	DelegateAccesses []*string     `json:"DelegateAccesses,omitempty" xml:"DelegateAccesses,omitempty" type:"Repeated"`
	MetaResource     *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	Principal        *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
}

func (s PrincipalResourcePermissions) String() string {
	return tea.Prettify(s)
}

func (s PrincipalResourcePermissions) GoString() string {
	return s.String()
}

func (s *PrincipalResourcePermissions) SetAccesses(v []*string) *PrincipalResourcePermissions {
	s.Accesses = v
	return s
}

func (s *PrincipalResourcePermissions) SetDelegateAccesses(v []*string) *PrincipalResourcePermissions {
	s.DelegateAccesses = v
	return s
}

func (s *PrincipalResourcePermissions) SetMetaResource(v *MetaResource) *PrincipalResourcePermissions {
	s.MetaResource = v
	return s
}

func (s *PrincipalResourcePermissions) SetPrincipal(v *Principal) *PrincipalResourcePermissions {
	s.Principal = v
	return s
}

type PrivilegeGrantInfo struct {
	CreateTime  *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantOption *bool   `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	Grantor     *string `json:"Grantor,omitempty" xml:"Grantor,omitempty"`
	GrantorType *string `json:"GrantorType,omitempty" xml:"GrantorType,omitempty"`
	Privilege   *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s PrivilegeGrantInfo) String() string {
	return tea.Prettify(s)
}

func (s PrivilegeGrantInfo) GoString() string {
	return s.String()
}

func (s *PrivilegeGrantInfo) SetCreateTime(v int32) *PrivilegeGrantInfo {
	s.CreateTime = &v
	return s
}

func (s *PrivilegeGrantInfo) SetGrantOption(v bool) *PrivilegeGrantInfo {
	s.GrantOption = &v
	return s
}

func (s *PrivilegeGrantInfo) SetGrantor(v string) *PrivilegeGrantInfo {
	s.Grantor = &v
	return s
}

func (s *PrivilegeGrantInfo) SetGrantorType(v string) *PrivilegeGrantInfo {
	s.GrantorType = &v
	return s
}

func (s *PrivilegeGrantInfo) SetPrivilege(v string) *PrivilegeGrantInfo {
	s.Privilege = &v
	return s
}

type PrivilegeResource struct {
	Access       *string       `json:"Access,omitempty" xml:"Access,omitempty"`
	MetaResource *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
}

func (s PrivilegeResource) String() string {
	return tea.Prettify(s)
}

func (s PrivilegeResource) GoString() string {
	return s.String()
}

func (s *PrivilegeResource) SetAccess(v string) *PrivilegeResource {
	s.Access = &v
	return s
}

func (s *PrivilegeResource) SetMetaResource(v *MetaResource) *PrivilegeResource {
	s.MetaResource = v
	return s
}

type ResourceUri struct {
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Uri          *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ResourceUri) String() string {
	return tea.Prettify(s)
}

func (s ResourceUri) GoString() string {
	return s.String()
}

func (s *ResourceUri) SetResourceType(v string) *ResourceUri {
	s.ResourceType = &v
	return s
}

func (s *ResourceUri) SetUri(v string) *ResourceUri {
	s.Uri = &v
	return s
}

type Role struct {
	CreateTime   *int64       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description  *string      `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName  *string      `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	IsPredefined *int32       `json:"IsPredefined,omitempty" xml:"IsPredefined,omitempty"`
	Name         *string      `json:"Name,omitempty" xml:"Name,omitempty"`
	PrincipalArn *string      `json:"PrincipalArn,omitempty" xml:"PrincipalArn,omitempty"`
	UpdateTime   *int64       `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Users        []*Principal `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s Role) String() string {
	return tea.Prettify(s)
}

func (s Role) GoString() string {
	return s.String()
}

func (s *Role) SetCreateTime(v int64) *Role {
	s.CreateTime = &v
	return s
}

func (s *Role) SetDescription(v string) *Role {
	s.Description = &v
	return s
}

func (s *Role) SetDisplayName(v string) *Role {
	s.DisplayName = &v
	return s
}

func (s *Role) SetIsPredefined(v int32) *Role {
	s.IsPredefined = &v
	return s
}

func (s *Role) SetName(v string) *Role {
	s.Name = &v
	return s
}

func (s *Role) SetPrincipalArn(v string) *Role {
	s.PrincipalArn = &v
	return s
}

func (s *Role) SetUpdateTime(v int64) *Role {
	s.UpdateTime = &v
	return s
}

func (s *Role) SetUsers(v []*Principal) *Role {
	s.Users = v
	return s
}

type RoleInput struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RoleInput) String() string {
	return tea.Prettify(s)
}

func (s RoleInput) GoString() string {
	return s.String()
}

func (s *RoleInput) SetDescription(v string) *RoleInput {
	s.Description = &v
	return s
}

func (s *RoleInput) SetDisplayName(v string) *RoleInput {
	s.DisplayName = &v
	return s
}

func (s *RoleInput) SetName(v string) *RoleInput {
	s.Name = &v
	return s
}

type SerDeInfo struct {
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerializationLib *string            `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
}

func (s SerDeInfo) String() string {
	return tea.Prettify(s)
}

func (s SerDeInfo) GoString() string {
	return s.String()
}

func (s *SerDeInfo) SetName(v string) *SerDeInfo {
	s.Name = &v
	return s
}

func (s *SerDeInfo) SetParameters(v map[string]*string) *SerDeInfo {
	s.Parameters = v
	return s
}

func (s *SerDeInfo) SetSerializationLib(v string) *SerDeInfo {
	s.SerializationLib = &v
	return s
}

type SingleIndicatorDTO struct {
	DayIncrement   *int64   `json:"DayIncrement,omitempty" xml:"DayIncrement,omitempty"`
	DayOnDay       *float64 `json:"DayOnDay,omitempty" xml:"DayOnDay,omitempty"`
	MonthIncrement *int64   `json:"MonthIncrement,omitempty" xml:"MonthIncrement,omitempty"`
	MonthOnMonth   *float64 `json:"MonthOnMonth,omitempty" xml:"MonthOnMonth,omitempty"`
	Total          *int64   `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s SingleIndicatorDTO) String() string {
	return tea.Prettify(s)
}

func (s SingleIndicatorDTO) GoString() string {
	return s.String()
}

func (s *SingleIndicatorDTO) SetDayIncrement(v int64) *SingleIndicatorDTO {
	s.DayIncrement = &v
	return s
}

func (s *SingleIndicatorDTO) SetDayOnDay(v float64) *SingleIndicatorDTO {
	s.DayOnDay = &v
	return s
}

func (s *SingleIndicatorDTO) SetMonthIncrement(v int64) *SingleIndicatorDTO {
	s.MonthIncrement = &v
	return s
}

func (s *SingleIndicatorDTO) SetMonthOnMonth(v float64) *SingleIndicatorDTO {
	s.MonthOnMonth = &v
	return s
}

func (s *SingleIndicatorDTO) SetTotal(v int64) *SingleIndicatorDTO {
	s.Total = &v
	return s
}

type SkewedInfo struct {
	SkewedColNames             []*string          `json:"SkewedColNames,omitempty" xml:"SkewedColNames,omitempty" type:"Repeated"`
	SkewedColValueLocationMaps map[string]*string `json:"SkewedColValueLocationMaps,omitempty" xml:"SkewedColValueLocationMaps,omitempty"`
	SkewedColValues            [][]*string        `json:"SkewedColValues,omitempty" xml:"SkewedColValues,omitempty" type:"Repeated"`
}

func (s SkewedInfo) String() string {
	return tea.Prettify(s)
}

func (s SkewedInfo) GoString() string {
	return s.String()
}

func (s *SkewedInfo) SetSkewedColNames(v []*string) *SkewedInfo {
	s.SkewedColNames = v
	return s
}

func (s *SkewedInfo) SetSkewedColValueLocationMaps(v map[string]*string) *SkewedInfo {
	s.SkewedColValueLocationMaps = v
	return s
}

func (s *SkewedInfo) SetSkewedColValues(v [][]*string) *SkewedInfo {
	s.SkewedColValues = v
	return s
}

type SmallFileCntRank struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Location  *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Quantity  *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s SmallFileCntRank) String() string {
	return tea.Prettify(s)
}

func (s SmallFileCntRank) GoString() string {
	return s.String()
}

func (s *SmallFileCntRank) SetDbName(v string) *SmallFileCntRank {
	s.DbName = &v
	return s
}

func (s *SmallFileCntRank) SetLocation(v string) *SmallFileCntRank {
	s.Location = &v
	return s
}

func (s *SmallFileCntRank) SetQuantity(v int64) *SmallFileCntRank {
	s.Quantity = &v
	return s
}

func (s *SmallFileCntRank) SetTableName(v string) *SmallFileCntRank {
	s.TableName = &v
	return s
}

type SortCriterion struct {
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Sort      *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s SortCriterion) String() string {
	return tea.Prettify(s)
}

func (s SortCriterion) GoString() string {
	return s.String()
}

func (s *SortCriterion) SetFieldName(v string) *SortCriterion {
	s.FieldName = &v
	return s
}

func (s *SortCriterion) SetSort(v string) *SortCriterion {
	s.Sort = &v
	return s
}

type StorageCollectTaskOperationResult struct {
	DlfCreated *bool   `json:"DlfCreated,omitempty" xml:"DlfCreated,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskType   *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s StorageCollectTaskOperationResult) String() string {
	return tea.Prettify(s)
}

func (s StorageCollectTaskOperationResult) GoString() string {
	return s.String()
}

func (s *StorageCollectTaskOperationResult) SetDlfCreated(v bool) *StorageCollectTaskOperationResult {
	s.DlfCreated = &v
	return s
}

func (s *StorageCollectTaskOperationResult) SetErrCode(v string) *StorageCollectTaskOperationResult {
	s.ErrCode = &v
	return s
}

func (s *StorageCollectTaskOperationResult) SetErrMessage(v string) *StorageCollectTaskOperationResult {
	s.ErrMessage = &v
	return s
}

func (s *StorageCollectTaskOperationResult) SetSuccess(v bool) *StorageCollectTaskOperationResult {
	s.Success = &v
	return s
}

func (s *StorageCollectTaskOperationResult) SetTaskId(v string) *StorageCollectTaskOperationResult {
	s.TaskId = &v
	return s
}

func (s *StorageCollectTaskOperationResult) SetTaskType(v string) *StorageCollectTaskOperationResult {
	s.TaskType = &v
	return s
}

type StorageDescriptor struct {
	BucketCols             []*string          `json:"BucketCols,omitempty" xml:"BucketCols,omitempty" type:"Repeated"`
	Cols                   []*FieldSchema     `json:"Cols,omitempty" xml:"Cols,omitempty" type:"Repeated"`
	Compressed             *bool              `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	InputFormat            *string            `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	Location               *string            `json:"Location,omitempty" xml:"Location,omitempty"`
	NumBuckets             *int32             `json:"NumBuckets,omitempty" xml:"NumBuckets,omitempty"`
	OutputFormat           *string            `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters             map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerDeInfo              *SerDeInfo         `json:"SerDeInfo,omitempty" xml:"SerDeInfo,omitempty"`
	SkewedInfo             *SkewedInfo        `json:"SkewedInfo,omitempty" xml:"SkewedInfo,omitempty"`
	SortCols               []*Order           `json:"SortCols,omitempty" xml:"SortCols,omitempty" type:"Repeated"`
	StoredAsSubDirectories *bool              `json:"StoredAsSubDirectories,omitempty" xml:"StoredAsSubDirectories,omitempty"`
}

func (s StorageDescriptor) String() string {
	return tea.Prettify(s)
}

func (s StorageDescriptor) GoString() string {
	return s.String()
}

func (s *StorageDescriptor) SetBucketCols(v []*string) *StorageDescriptor {
	s.BucketCols = v
	return s
}

func (s *StorageDescriptor) SetCols(v []*FieldSchema) *StorageDescriptor {
	s.Cols = v
	return s
}

func (s *StorageDescriptor) SetCompressed(v bool) *StorageDescriptor {
	s.Compressed = &v
	return s
}

func (s *StorageDescriptor) SetInputFormat(v string) *StorageDescriptor {
	s.InputFormat = &v
	return s
}

func (s *StorageDescriptor) SetLocation(v string) *StorageDescriptor {
	s.Location = &v
	return s
}

func (s *StorageDescriptor) SetNumBuckets(v int32) *StorageDescriptor {
	s.NumBuckets = &v
	return s
}

func (s *StorageDescriptor) SetOutputFormat(v string) *StorageDescriptor {
	s.OutputFormat = &v
	return s
}

func (s *StorageDescriptor) SetParameters(v map[string]*string) *StorageDescriptor {
	s.Parameters = v
	return s
}

func (s *StorageDescriptor) SetSerDeInfo(v *SerDeInfo) *StorageDescriptor {
	s.SerDeInfo = v
	return s
}

func (s *StorageDescriptor) SetSkewedInfo(v *SkewedInfo) *StorageDescriptor {
	s.SkewedInfo = v
	return s
}

func (s *StorageDescriptor) SetSortCols(v []*Order) *StorageDescriptor {
	s.SortCols = v
	return s
}

func (s *StorageDescriptor) SetStoredAsSubDirectories(v bool) *StorageDescriptor {
	s.StoredAsSubDirectories = &v
	return s
}

type StorageFormat struct {
	Avro          *int64 `json:"Avro,omitempty" xml:"Avro,omitempty"`
	Csv           *int64 `json:"Csv,omitempty" xml:"Csv,omitempty"`
	Delta         *int64 `json:"Delta,omitempty" xml:"Delta,omitempty"`
	Hudi          *int64 `json:"Hudi,omitempty" xml:"Hudi,omitempty"`
	Iceberg       *int64 `json:"Iceberg,omitempty" xml:"Iceberg,omitempty"`
	Json          *int64 `json:"Json,omitempty" xml:"Json,omitempty"`
	Orc           *int64 `json:"Orc,omitempty" xml:"Orc,omitempty"`
	Parquet       *int64 `json:"Parquet,omitempty" xml:"Parquet,omitempty"`
	Uncategorized *int64 `json:"Uncategorized,omitempty" xml:"Uncategorized,omitempty"`
}

func (s StorageFormat) String() string {
	return tea.Prettify(s)
}

func (s StorageFormat) GoString() string {
	return s.String()
}

func (s *StorageFormat) SetAvro(v int64) *StorageFormat {
	s.Avro = &v
	return s
}

func (s *StorageFormat) SetCsv(v int64) *StorageFormat {
	s.Csv = &v
	return s
}

func (s *StorageFormat) SetDelta(v int64) *StorageFormat {
	s.Delta = &v
	return s
}

func (s *StorageFormat) SetHudi(v int64) *StorageFormat {
	s.Hudi = &v
	return s
}

func (s *StorageFormat) SetIceberg(v int64) *StorageFormat {
	s.Iceberg = &v
	return s
}

func (s *StorageFormat) SetJson(v int64) *StorageFormat {
	s.Json = &v
	return s
}

func (s *StorageFormat) SetOrc(v int64) *StorageFormat {
	s.Orc = &v
	return s
}

func (s *StorageFormat) SetParquet(v int64) *StorageFormat {
	s.Parquet = &v
	return s
}

func (s *StorageFormat) SetUncategorized(v int64) *StorageFormat {
	s.Uncategorized = &v
	return s
}

type StorageLayer struct {
	Archive     *int64 `json:"Archive,omitempty" xml:"Archive,omitempty"`
	ColdArchive *int64 `json:"ColdArchive,omitempty" xml:"ColdArchive,omitempty"`
	Infrequent  *int64 `json:"Infrequent,omitempty" xml:"Infrequent,omitempty"`
	Standard    *int64 `json:"Standard,omitempty" xml:"Standard,omitempty"`
	Unknown     *int64 `json:"Unknown,omitempty" xml:"Unknown,omitempty"`
}

func (s StorageLayer) String() string {
	return tea.Prettify(s)
}

func (s StorageLayer) GoString() string {
	return s.String()
}

func (s *StorageLayer) SetArchive(v int64) *StorageLayer {
	s.Archive = &v
	return s
}

func (s *StorageLayer) SetColdArchive(v int64) *StorageLayer {
	s.ColdArchive = &v
	return s
}

func (s *StorageLayer) SetInfrequent(v int64) *StorageLayer {
	s.Infrequent = &v
	return s
}

func (s *StorageLayer) SetStandard(v int64) *StorageLayer {
	s.Standard = &v
	return s
}

func (s *StorageLayer) SetUnknown(v int64) *StorageLayer {
	s.Unknown = &v
	return s
}

type StorageRankDTO struct {
	DbStorageRank    []*DbStorageRank    `json:"dbStorageRank,omitempty" xml:"dbStorageRank,omitempty" type:"Repeated"`
	SmallFileCntRank []*SmallFileCntRank `json:"smallFileCntRank,omitempty" xml:"smallFileCntRank,omitempty" type:"Repeated"`
	TableStorageRank []*TableStorageRank `json:"tableStorageRank,omitempty" xml:"tableStorageRank,omitempty" type:"Repeated"`
}

func (s StorageRankDTO) String() string {
	return tea.Prettify(s)
}

func (s StorageRankDTO) GoString() string {
	return s.String()
}

func (s *StorageRankDTO) SetDbStorageRank(v []*DbStorageRank) *StorageRankDTO {
	s.DbStorageRank = v
	return s
}

func (s *StorageRankDTO) SetSmallFileCntRank(v []*SmallFileCntRank) *StorageRankDTO {
	s.SmallFileCntRank = v
	return s
}

func (s *StorageRankDTO) SetTableStorageRank(v []*TableStorageRank) *StorageRankDTO {
	s.TableStorageRank = v
	return s
}

type StorageSummary struct {
	DatabaseNum  *int32 `json:"DatabaseNum,omitempty" xml:"DatabaseNum,omitempty"`
	PartitionNum *int32 `json:"PartitionNum,omitempty" xml:"PartitionNum,omitempty"`
	TableNum     *int32 `json:"TableNum,omitempty" xml:"TableNum,omitempty"`
}

func (s StorageSummary) String() string {
	return tea.Prettify(s)
}

func (s StorageSummary) GoString() string {
	return s.String()
}

func (s *StorageSummary) SetDatabaseNum(v int32) *StorageSummary {
	s.DatabaseNum = &v
	return s
}

func (s *StorageSummary) SetPartitionNum(v int32) *StorageSummary {
	s.PartitionNum = &v
	return s
}

func (s *StorageSummary) SetTableNum(v int32) *StorageSummary {
	s.TableNum = &v
	return s
}

type StrogeCollectTask struct {
	DestinationBucketName *string `json:"DestinationBucketName,omitempty" xml:"DestinationBucketName,omitempty"`
	DestinationPrefix     *string `json:"DestinationPrefix,omitempty" xml:"DestinationPrefix,omitempty"`
	DlfCreated            *bool   `json:"DlfCreated,omitempty" xml:"DlfCreated,omitempty"`
	GmtCreate             *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified           *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                    *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InventoryId           *string `json:"InventoryId,omitempty" xml:"InventoryId,omitempty"`
	Location              *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskType              *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s StrogeCollectTask) String() string {
	return tea.Prettify(s)
}

func (s StrogeCollectTask) GoString() string {
	return s.String()
}

func (s *StrogeCollectTask) SetDestinationBucketName(v string) *StrogeCollectTask {
	s.DestinationBucketName = &v
	return s
}

func (s *StrogeCollectTask) SetDestinationPrefix(v string) *StrogeCollectTask {
	s.DestinationPrefix = &v
	return s
}

func (s *StrogeCollectTask) SetDlfCreated(v bool) *StrogeCollectTask {
	s.DlfCreated = &v
	return s
}

func (s *StrogeCollectTask) SetGmtCreate(v string) *StrogeCollectTask {
	s.GmtCreate = &v
	return s
}

func (s *StrogeCollectTask) SetGmtModified(v string) *StrogeCollectTask {
	s.GmtModified = &v
	return s
}

func (s *StrogeCollectTask) SetId(v string) *StrogeCollectTask {
	s.Id = &v
	return s
}

func (s *StrogeCollectTask) SetInventoryId(v string) *StrogeCollectTask {
	s.InventoryId = &v
	return s
}

func (s *StrogeCollectTask) SetLocation(v string) *StrogeCollectTask {
	s.Location = &v
	return s
}

func (s *StrogeCollectTask) SetStatus(v string) *StrogeCollectTask {
	s.Status = &v
	return s
}

func (s *StrogeCollectTask) SetTaskType(v string) *StrogeCollectTask {
	s.TaskType = &v
	return s
}

type Table struct {
	Cascade          *bool                  `json:"Cascade,omitempty" xml:"Cascade,omitempty"`
	CreateTime       *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	DatabaseName     *string                `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LastAccessTime   *int32                 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastAnalyzedTime *int32                 `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	Owner            *string                `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerType        *string                `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters       map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionKeys    []*FieldSchema         `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	Privileges       *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	Retention        *int32                 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RewriteEnabled   *bool                  `json:"RewriteEnabled,omitempty" xml:"RewriteEnabled,omitempty"`
	Sd               *StorageDescriptor     `json:"Sd,omitempty" xml:"Sd,omitempty"`
	TableId          *string                `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName        *string                `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType        *string                `json:"TableType,omitempty" xml:"TableType,omitempty"`
	TableVersion     *TableVersion          `json:"TableVersion,omitempty" xml:"TableVersion,omitempty"`
	Temporary        *bool                  `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	UpdateTime       *int32                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ViewExpandedText *string                `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText *string                `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s Table) String() string {
	return tea.Prettify(s)
}

func (s Table) GoString() string {
	return s.String()
}

func (s *Table) SetCascade(v bool) *Table {
	s.Cascade = &v
	return s
}

func (s *Table) SetCreateTime(v int32) *Table {
	s.CreateTime = &v
	return s
}

func (s *Table) SetCreatedBy(v string) *Table {
	s.CreatedBy = &v
	return s
}

func (s *Table) SetDatabaseName(v string) *Table {
	s.DatabaseName = &v
	return s
}

func (s *Table) SetLastAccessTime(v int32) *Table {
	s.LastAccessTime = &v
	return s
}

func (s *Table) SetLastAnalyzedTime(v int32) *Table {
	s.LastAnalyzedTime = &v
	return s
}

func (s *Table) SetOwner(v string) *Table {
	s.Owner = &v
	return s
}

func (s *Table) SetOwnerType(v string) *Table {
	s.OwnerType = &v
	return s
}

func (s *Table) SetParameters(v map[string]*string) *Table {
	s.Parameters = v
	return s
}

func (s *Table) SetPartitionKeys(v []*FieldSchema) *Table {
	s.PartitionKeys = v
	return s
}

func (s *Table) SetPrivileges(v *PrincipalPrivilegeSet) *Table {
	s.Privileges = v
	return s
}

func (s *Table) SetRetention(v int32) *Table {
	s.Retention = &v
	return s
}

func (s *Table) SetRewriteEnabled(v bool) *Table {
	s.RewriteEnabled = &v
	return s
}

func (s *Table) SetSd(v *StorageDescriptor) *Table {
	s.Sd = v
	return s
}

func (s *Table) SetTableId(v string) *Table {
	s.TableId = &v
	return s
}

func (s *Table) SetTableName(v string) *Table {
	s.TableName = &v
	return s
}

func (s *Table) SetTableType(v string) *Table {
	s.TableType = &v
	return s
}

func (s *Table) SetTableVersion(v *TableVersion) *Table {
	s.TableVersion = v
	return s
}

func (s *Table) SetTemporary(v bool) *Table {
	s.Temporary = &v
	return s
}

func (s *Table) SetUpdateTime(v int32) *Table {
	s.UpdateTime = &v
	return s
}

func (s *Table) SetViewExpandedText(v string) *Table {
	s.ViewExpandedText = &v
	return s
}

func (s *Table) SetViewOriginalText(v string) *Table {
	s.ViewOriginalText = &v
	return s
}

type TableError struct {
	ErrorDetail *ErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty"`
	TableName   *string      `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableError) String() string {
	return tea.Prettify(s)
}

func (s TableError) GoString() string {
	return s.String()
}

func (s *TableError) SetErrorDetail(v *ErrorDetail) *TableError {
	s.ErrorDetail = v
	return s
}

func (s *TableError) SetTableName(v string) *TableError {
	s.TableName = &v
	return s
}

type TableExtended struct {
	Cascade          *bool                    `json:"Cascade,omitempty" xml:"Cascade,omitempty"`
	CreateTime       *int32                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy        *string                  `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	DatabaseName     *string                  `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LastAccessTime   *int32                   `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastAnalyzedTime *int32                   `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	Owner            *string                  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerType        *string                  `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters       map[string]*string       `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionKeys    []*FieldSchema           `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	Privileges       *TableExtendedPrivileges `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Struct"`
	Retention        *int32                   `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RewriteEnabled   *bool                    `json:"RewriteEnabled,omitempty" xml:"RewriteEnabled,omitempty"`
	Sd               *TableExtendedSd         `json:"Sd,omitempty" xml:"Sd,omitempty" type:"Struct"`
	TableFormat      *string                  `json:"TableFormat,omitempty" xml:"TableFormat,omitempty"`
	TableName        *string                  `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType        *string                  `json:"TableType,omitempty" xml:"TableType,omitempty"`
	Temporary        *bool                    `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	UpdateTime       *int32                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	ViewExpandedText *string                  `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText *string                  `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s TableExtended) String() string {
	return tea.Prettify(s)
}

func (s TableExtended) GoString() string {
	return s.String()
}

func (s *TableExtended) SetCascade(v bool) *TableExtended {
	s.Cascade = &v
	return s
}

func (s *TableExtended) SetCreateTime(v int32) *TableExtended {
	s.CreateTime = &v
	return s
}

func (s *TableExtended) SetCreatedBy(v string) *TableExtended {
	s.CreatedBy = &v
	return s
}

func (s *TableExtended) SetDatabaseName(v string) *TableExtended {
	s.DatabaseName = &v
	return s
}

func (s *TableExtended) SetLastAccessTime(v int32) *TableExtended {
	s.LastAccessTime = &v
	return s
}

func (s *TableExtended) SetLastAnalyzedTime(v int32) *TableExtended {
	s.LastAnalyzedTime = &v
	return s
}

func (s *TableExtended) SetOwner(v string) *TableExtended {
	s.Owner = &v
	return s
}

func (s *TableExtended) SetOwnerType(v string) *TableExtended {
	s.OwnerType = &v
	return s
}

func (s *TableExtended) SetParameters(v map[string]*string) *TableExtended {
	s.Parameters = v
	return s
}

func (s *TableExtended) SetPartitionKeys(v []*FieldSchema) *TableExtended {
	s.PartitionKeys = v
	return s
}

func (s *TableExtended) SetPrivileges(v *TableExtendedPrivileges) *TableExtended {
	s.Privileges = v
	return s
}

func (s *TableExtended) SetRetention(v int32) *TableExtended {
	s.Retention = &v
	return s
}

func (s *TableExtended) SetRewriteEnabled(v bool) *TableExtended {
	s.RewriteEnabled = &v
	return s
}

func (s *TableExtended) SetSd(v *TableExtendedSd) *TableExtended {
	s.Sd = v
	return s
}

func (s *TableExtended) SetTableFormat(v string) *TableExtended {
	s.TableFormat = &v
	return s
}

func (s *TableExtended) SetTableName(v string) *TableExtended {
	s.TableName = &v
	return s
}

func (s *TableExtended) SetTableType(v string) *TableExtended {
	s.TableType = &v
	return s
}

func (s *TableExtended) SetTemporary(v bool) *TableExtended {
	s.Temporary = &v
	return s
}

func (s *TableExtended) SetUpdateTime(v int32) *TableExtended {
	s.UpdateTime = &v
	return s
}

func (s *TableExtended) SetViewExpandedText(v string) *TableExtended {
	s.ViewExpandedText = &v
	return s
}

func (s *TableExtended) SetViewOriginalText(v string) *TableExtended {
	s.ViewOriginalText = &v
	return s
}

type TableExtendedPrivileges struct {
	RolePrivileges  map[string][]*TableExtendedPrivilegesRolePrivilegesValue  `json:"RolePrivileges,omitempty" xml:"RolePrivileges,omitempty"`
	UserPrivileges  map[string][]*TableExtendedPrivilegesUserPrivilegesValue  `json:"UserPrivileges,omitempty" xml:"UserPrivileges,omitempty"`
	GroupPrivileges map[string][]*TableExtendedPrivilegesGroupPrivilegesValue `json:"groupPrivileges,omitempty" xml:"groupPrivileges,omitempty"`
}

func (s TableExtendedPrivileges) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedPrivileges) GoString() string {
	return s.String()
}

func (s *TableExtendedPrivileges) SetRolePrivileges(v map[string][]*TableExtendedPrivilegesRolePrivilegesValue) *TableExtendedPrivileges {
	s.RolePrivileges = v
	return s
}

func (s *TableExtendedPrivileges) SetUserPrivileges(v map[string][]*TableExtendedPrivilegesUserPrivilegesValue) *TableExtendedPrivileges {
	s.UserPrivileges = v
	return s
}

func (s *TableExtendedPrivileges) SetGroupPrivileges(v map[string][]*TableExtendedPrivilegesGroupPrivilegesValue) *TableExtendedPrivileges {
	s.GroupPrivileges = v
	return s
}

type TableExtendedSd struct {
	BucketCols             []*string                  `json:"BucketCols,omitempty" xml:"BucketCols,omitempty" type:"Repeated"`
	Cols                   []*FieldSchema             `json:"Cols,omitempty" xml:"Cols,omitempty" type:"Repeated"`
	Compressed             *bool                      `json:"Compressed,omitempty" xml:"Compressed,omitempty"`
	InputFormat            *string                    `json:"InputFormat,omitempty" xml:"InputFormat,omitempty"`
	Location               *string                    `json:"Location,omitempty" xml:"Location,omitempty"`
	NumBuckets             *int32                     `json:"NumBuckets,omitempty" xml:"NumBuckets,omitempty"`
	OutputFormat           *string                    `json:"OutputFormat,omitempty" xml:"OutputFormat,omitempty"`
	Parameters             map[string]*string         `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerDeInfo              *TableExtendedSdSerDeInfo  `json:"SerDeInfo,omitempty" xml:"SerDeInfo,omitempty" type:"Struct"`
	SkewedInfo             *TableExtendedSdSkewedInfo `json:"SkewedInfo,omitempty" xml:"SkewedInfo,omitempty" type:"Struct"`
	SortCols               []*Order                   `json:"SortCols,omitempty" xml:"SortCols,omitempty" type:"Repeated"`
	StoredAsSubDirectories *bool                      `json:"StoredAsSubDirectories,omitempty" xml:"StoredAsSubDirectories,omitempty"`
}

func (s TableExtendedSd) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedSd) GoString() string {
	return s.String()
}

func (s *TableExtendedSd) SetBucketCols(v []*string) *TableExtendedSd {
	s.BucketCols = v
	return s
}

func (s *TableExtendedSd) SetCols(v []*FieldSchema) *TableExtendedSd {
	s.Cols = v
	return s
}

func (s *TableExtendedSd) SetCompressed(v bool) *TableExtendedSd {
	s.Compressed = &v
	return s
}

func (s *TableExtendedSd) SetInputFormat(v string) *TableExtendedSd {
	s.InputFormat = &v
	return s
}

func (s *TableExtendedSd) SetLocation(v string) *TableExtendedSd {
	s.Location = &v
	return s
}

func (s *TableExtendedSd) SetNumBuckets(v int32) *TableExtendedSd {
	s.NumBuckets = &v
	return s
}

func (s *TableExtendedSd) SetOutputFormat(v string) *TableExtendedSd {
	s.OutputFormat = &v
	return s
}

func (s *TableExtendedSd) SetParameters(v map[string]*string) *TableExtendedSd {
	s.Parameters = v
	return s
}

func (s *TableExtendedSd) SetSerDeInfo(v *TableExtendedSdSerDeInfo) *TableExtendedSd {
	s.SerDeInfo = v
	return s
}

func (s *TableExtendedSd) SetSkewedInfo(v *TableExtendedSdSkewedInfo) *TableExtendedSd {
	s.SkewedInfo = v
	return s
}

func (s *TableExtendedSd) SetSortCols(v []*Order) *TableExtendedSd {
	s.SortCols = v
	return s
}

func (s *TableExtendedSd) SetStoredAsSubDirectories(v bool) *TableExtendedSd {
	s.StoredAsSubDirectories = &v
	return s
}

type TableExtendedSdSerDeInfo struct {
	Name             *string            `json:"Name,omitempty" xml:"Name,omitempty"`
	Parameters       map[string]*string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	SerializationLib *string            `json:"SerializationLib,omitempty" xml:"SerializationLib,omitempty"`
}

func (s TableExtendedSdSerDeInfo) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedSdSerDeInfo) GoString() string {
	return s.String()
}

func (s *TableExtendedSdSerDeInfo) SetName(v string) *TableExtendedSdSerDeInfo {
	s.Name = &v
	return s
}

func (s *TableExtendedSdSerDeInfo) SetParameters(v map[string]*string) *TableExtendedSdSerDeInfo {
	s.Parameters = v
	return s
}

func (s *TableExtendedSdSerDeInfo) SetSerializationLib(v string) *TableExtendedSdSerDeInfo {
	s.SerializationLib = &v
	return s
}

type TableExtendedSdSkewedInfo struct {
	SkewedColNames             []*string          `json:"SkewedColNames,omitempty" xml:"SkewedColNames,omitempty" type:"Repeated"`
	SkewedColValueLocationMaps map[string]*string `json:"SkewedColValueLocationMaps,omitempty" xml:"SkewedColValueLocationMaps,omitempty"`
	SkewedColValues            [][]*string        `json:"SkewedColValues,omitempty" xml:"SkewedColValues,omitempty" type:"Repeated"`
}

func (s TableExtendedSdSkewedInfo) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedSdSkewedInfo) GoString() string {
	return s.String()
}

func (s *TableExtendedSdSkewedInfo) SetSkewedColNames(v []*string) *TableExtendedSdSkewedInfo {
	s.SkewedColNames = v
	return s
}

func (s *TableExtendedSdSkewedInfo) SetSkewedColValueLocationMaps(v map[string]*string) *TableExtendedSdSkewedInfo {
	s.SkewedColValueLocationMaps = v
	return s
}

func (s *TableExtendedSdSkewedInfo) SetSkewedColValues(v [][]*string) *TableExtendedSdSkewedInfo {
	s.SkewedColValues = v
	return s
}

type TableInput struct {
	Cascade          *bool                  `json:"Cascade,omitempty" xml:"Cascade,omitempty"`
	CreateTime       *int32                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreatedBy        *string                `json:"CreatedBy,omitempty" xml:"CreatedBy,omitempty"`
	DatabaseName     *string                `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	LastAccessTime   *int32                 `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastAnalyzedTime *int32                 `json:"LastAnalyzedTime,omitempty" xml:"LastAnalyzedTime,omitempty"`
	Owner            *string                `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerType        *string                `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	Parameters       map[string]*string     `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	PartitionKeys    []*FieldSchema         `json:"PartitionKeys,omitempty" xml:"PartitionKeys,omitempty" type:"Repeated"`
	Privileges       *PrincipalPrivilegeSet `json:"Privileges,omitempty" xml:"Privileges,omitempty"`
	Retention        *int32                 `json:"Retention,omitempty" xml:"Retention,omitempty"`
	RewriteEnabled   *bool                  `json:"RewriteEnabled,omitempty" xml:"RewriteEnabled,omitempty"`
	Sd               *StorageDescriptor     `json:"Sd,omitempty" xml:"Sd,omitempty"`
	TableName        *string                `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TableType        *string                `json:"TableType,omitempty" xml:"TableType,omitempty"`
	Temporary        *bool                  `json:"Temporary,omitempty" xml:"Temporary,omitempty"`
	ViewExpandedText *string                `json:"ViewExpandedText,omitempty" xml:"ViewExpandedText,omitempty"`
	ViewOriginalText *string                `json:"ViewOriginalText,omitempty" xml:"ViewOriginalText,omitempty"`
}

func (s TableInput) String() string {
	return tea.Prettify(s)
}

func (s TableInput) GoString() string {
	return s.String()
}

func (s *TableInput) SetCascade(v bool) *TableInput {
	s.Cascade = &v
	return s
}

func (s *TableInput) SetCreateTime(v int32) *TableInput {
	s.CreateTime = &v
	return s
}

func (s *TableInput) SetCreatedBy(v string) *TableInput {
	s.CreatedBy = &v
	return s
}

func (s *TableInput) SetDatabaseName(v string) *TableInput {
	s.DatabaseName = &v
	return s
}

func (s *TableInput) SetLastAccessTime(v int32) *TableInput {
	s.LastAccessTime = &v
	return s
}

func (s *TableInput) SetLastAnalyzedTime(v int32) *TableInput {
	s.LastAnalyzedTime = &v
	return s
}

func (s *TableInput) SetOwner(v string) *TableInput {
	s.Owner = &v
	return s
}

func (s *TableInput) SetOwnerType(v string) *TableInput {
	s.OwnerType = &v
	return s
}

func (s *TableInput) SetParameters(v map[string]*string) *TableInput {
	s.Parameters = v
	return s
}

func (s *TableInput) SetPartitionKeys(v []*FieldSchema) *TableInput {
	s.PartitionKeys = v
	return s
}

func (s *TableInput) SetPrivileges(v *PrincipalPrivilegeSet) *TableInput {
	s.Privileges = v
	return s
}

func (s *TableInput) SetRetention(v int32) *TableInput {
	s.Retention = &v
	return s
}

func (s *TableInput) SetRewriteEnabled(v bool) *TableInput {
	s.RewriteEnabled = &v
	return s
}

func (s *TableInput) SetSd(v *StorageDescriptor) *TableInput {
	s.Sd = v
	return s
}

func (s *TableInput) SetTableName(v string) *TableInput {
	s.TableName = &v
	return s
}

func (s *TableInput) SetTableType(v string) *TableInput {
	s.TableType = &v
	return s
}

func (s *TableInput) SetTemporary(v bool) *TableInput {
	s.Temporary = &v
	return s
}

func (s *TableInput) SetViewExpandedText(v string) *TableInput {
	s.ViewExpandedText = &v
	return s
}

func (s *TableInput) SetViewOriginalText(v string) *TableInput {
	s.ViewOriginalText = &v
	return s
}

type TableProfile struct {
	AccessNum              *int64  `json:"AccessNum,omitempty" xml:"AccessNum,omitempty"`
	AccessNumMonthly       *int64  `json:"AccessNumMonthly,omitempty" xml:"AccessNumMonthly,omitempty"`
	AccessNumWeekly        *int64  `json:"AccessNumWeekly,omitempty" xml:"AccessNumWeekly,omitempty"`
	CreateTime             *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSourceType         *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DatabaseName           *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FileCnt                *int64  `json:"FileCnt,omitempty" xml:"FileCnt,omitempty"`
	FileSize               *int64  `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	IsPartitioned          *bool   `json:"IsPartitioned,omitempty" xml:"IsPartitioned,omitempty"`
	LastAccessNumTime      *string `json:"LastAccessNumTime,omitempty" xml:"LastAccessNumTime,omitempty"`
	LastAccessTime         *string `json:"LastAccessTime,omitempty" xml:"LastAccessTime,omitempty"`
	LastDdlTime            *string `json:"LastDdlTime,omitempty" xml:"LastDdlTime,omitempty"`
	LastModifyTime         *string `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	LatestAccessNumDate    *string `json:"LatestAccessNumDate,omitempty" xml:"LatestAccessNumDate,omitempty"`
	LatestDate             *string `json:"LatestDate,omitempty" xml:"LatestDate,omitempty"`
	Location               *string `json:"Location,omitempty" xml:"Location,omitempty"`
	ObjectAccessNum        *int64  `json:"ObjectAccessNum,omitempty" xml:"ObjectAccessNum,omitempty"`
	ObjectAccessNumMonthly *int64  `json:"ObjectAccessNumMonthly,omitempty" xml:"ObjectAccessNumMonthly,omitempty"`
	ObjectAccessNumWeekly  *int64  `json:"ObjectAccessNumWeekly,omitempty" xml:"ObjectAccessNumWeekly,omitempty"`
	ObjectCnt              *int64  `json:"ObjectCnt,omitempty" xml:"ObjectCnt,omitempty"`
	ObjectSize             *int64  `json:"ObjectSize,omitempty" xml:"ObjectSize,omitempty"`
	PartitionCnt           *int64  `json:"PartitionCnt,omitempty" xml:"PartitionCnt,omitempty"`
	RecordCnt              *int64  `json:"RecordCnt,omitempty" xml:"RecordCnt,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableProfile) String() string {
	return tea.Prettify(s)
}

func (s TableProfile) GoString() string {
	return s.String()
}

func (s *TableProfile) SetAccessNum(v int64) *TableProfile {
	s.AccessNum = &v
	return s
}

func (s *TableProfile) SetAccessNumMonthly(v int64) *TableProfile {
	s.AccessNumMonthly = &v
	return s
}

func (s *TableProfile) SetAccessNumWeekly(v int64) *TableProfile {
	s.AccessNumWeekly = &v
	return s
}

func (s *TableProfile) SetCreateTime(v string) *TableProfile {
	s.CreateTime = &v
	return s
}

func (s *TableProfile) SetDataSourceType(v string) *TableProfile {
	s.DataSourceType = &v
	return s
}

func (s *TableProfile) SetDatabaseName(v string) *TableProfile {
	s.DatabaseName = &v
	return s
}

func (s *TableProfile) SetFileCnt(v int64) *TableProfile {
	s.FileCnt = &v
	return s
}

func (s *TableProfile) SetFileSize(v int64) *TableProfile {
	s.FileSize = &v
	return s
}

func (s *TableProfile) SetIsPartitioned(v bool) *TableProfile {
	s.IsPartitioned = &v
	return s
}

func (s *TableProfile) SetLastAccessNumTime(v string) *TableProfile {
	s.LastAccessNumTime = &v
	return s
}

func (s *TableProfile) SetLastAccessTime(v string) *TableProfile {
	s.LastAccessTime = &v
	return s
}

func (s *TableProfile) SetLastDdlTime(v string) *TableProfile {
	s.LastDdlTime = &v
	return s
}

func (s *TableProfile) SetLastModifyTime(v string) *TableProfile {
	s.LastModifyTime = &v
	return s
}

func (s *TableProfile) SetLatestAccessNumDate(v string) *TableProfile {
	s.LatestAccessNumDate = &v
	return s
}

func (s *TableProfile) SetLatestDate(v string) *TableProfile {
	s.LatestDate = &v
	return s
}

func (s *TableProfile) SetLocation(v string) *TableProfile {
	s.Location = &v
	return s
}

func (s *TableProfile) SetObjectAccessNum(v int64) *TableProfile {
	s.ObjectAccessNum = &v
	return s
}

func (s *TableProfile) SetObjectAccessNumMonthly(v int64) *TableProfile {
	s.ObjectAccessNumMonthly = &v
	return s
}

func (s *TableProfile) SetObjectAccessNumWeekly(v int64) *TableProfile {
	s.ObjectAccessNumWeekly = &v
	return s
}

func (s *TableProfile) SetObjectCnt(v int64) *TableProfile {
	s.ObjectCnt = &v
	return s
}

func (s *TableProfile) SetObjectSize(v int64) *TableProfile {
	s.ObjectSize = &v
	return s
}

func (s *TableProfile) SetPartitionCnt(v int64) *TableProfile {
	s.PartitionCnt = &v
	return s
}

func (s *TableProfile) SetRecordCnt(v int64) *TableProfile {
	s.RecordCnt = &v
	return s
}

func (s *TableProfile) SetTableName(v string) *TableProfile {
	s.TableName = &v
	return s
}

type TableResource struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableResource) String() string {
	return tea.Prettify(s)
}

func (s TableResource) GoString() string {
	return s.String()
}

func (s *TableResource) SetDatabaseName(v string) *TableResource {
	s.DatabaseName = &v
	return s
}

func (s *TableResource) SetTableName(v string) *TableResource {
	s.TableName = &v
	return s
}

type TableStorageRank struct {
	DbName    *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Quantity  *int64  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s TableStorageRank) String() string {
	return tea.Prettify(s)
}

func (s TableStorageRank) GoString() string {
	return s.String()
}

func (s *TableStorageRank) SetDbName(v string) *TableStorageRank {
	s.DbName = &v
	return s
}

func (s *TableStorageRank) SetQuantity(v int64) *TableStorageRank {
	s.Quantity = &v
	return s
}

func (s *TableStorageRank) SetTableName(v string) *TableStorageRank {
	s.TableName = &v
	return s
}

type TableVersion struct {
	Table     *Table `json:"Table,omitempty" xml:"Table,omitempty"`
	VersionId *int32 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s TableVersion) String() string {
	return tea.Prettify(s)
}

func (s TableVersion) GoString() string {
	return s.String()
}

func (s *TableVersion) SetTable(v *Table) *TableVersion {
	s.Table = v
	return s
}

func (s *TableVersion) SetVersionId(v int32) *TableVersion {
	s.VersionId = &v
	return s
}

type TaskStatus struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s TaskStatus) String() string {
	return tea.Prettify(s)
}

func (s TaskStatus) GoString() string {
	return s.String()
}

func (s *TaskStatus) SetMessage(v string) *TaskStatus {
	s.Message = &v
	return s
}

func (s *TaskStatus) SetStatus(v string) *TaskStatus {
	s.Status = &v
	return s
}

type UnarchiveDetail struct {
	ApiCallTimes        *int64  `json:"ApiCallTimes,omitempty" xml:"ApiCallTimes,omitempty"`
	Cost                *int64  `json:"Cost,omitempty" xml:"Cost,omitempty"`
	StorageSize         *int64  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty"`
	StorageType         *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	UnarchiveTaskStatus *string `json:"UnarchiveTaskStatus,omitempty" xml:"UnarchiveTaskStatus,omitempty"`
}

func (s UnarchiveDetail) String() string {
	return tea.Prettify(s)
}

func (s UnarchiveDetail) GoString() string {
	return s.String()
}

func (s *UnarchiveDetail) SetApiCallTimes(v int64) *UnarchiveDetail {
	s.ApiCallTimes = &v
	return s
}

func (s *UnarchiveDetail) SetCost(v int64) *UnarchiveDetail {
	s.Cost = &v
	return s
}

func (s *UnarchiveDetail) SetStorageSize(v int64) *UnarchiveDetail {
	s.StorageSize = &v
	return s
}

func (s *UnarchiveDetail) SetStorageType(v string) *UnarchiveDetail {
	s.StorageType = &v
	return s
}

func (s *UnarchiveDetail) SetUnarchiveTaskStatus(v string) *UnarchiveDetail {
	s.UnarchiveTaskStatus = &v
	return s
}

type UpdateTablePartitionColumnStatisticsRequest struct {
	CatalogId            *string             `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnStatisticsList []*ColumnStatistics `json:"ColumnStatisticsList,omitempty" xml:"ColumnStatisticsList,omitempty" type:"Repeated"`
	DatabaseName         *string             `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Engine               *string             `json:"Engine,omitempty" xml:"Engine,omitempty"`
	IsStatsCompliant     *bool               `json:"IsStatsCompliant,omitempty" xml:"IsStatsCompliant,omitempty"`
	TableName            *string             `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ValidWriteIdList     *string             `json:"ValidWriteIdList,omitempty" xml:"ValidWriteIdList,omitempty"`
	WriteId              *string             `json:"WriteId,omitempty" xml:"WriteId,omitempty"`
}

func (s UpdateTablePartitionColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTablePartitionColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetCatalogId(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetColumnStatisticsList(v []*ColumnStatistics) *UpdateTablePartitionColumnStatisticsRequest {
	s.ColumnStatisticsList = v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetDatabaseName(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetEngine(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.Engine = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetIsStatsCompliant(v bool) *UpdateTablePartitionColumnStatisticsRequest {
	s.IsStatsCompliant = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetTableName(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.TableName = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetValidWriteIdList(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.ValidWriteIdList = &v
	return s
}

func (s *UpdateTablePartitionColumnStatisticsRequest) SetWriteId(v string) *UpdateTablePartitionColumnStatisticsRequest {
	s.WriteId = &v
	return s
}

type UserRole struct {
	GrantTime *int64     `json:"GrantTime,omitempty" xml:"GrantTime,omitempty"`
	Role      *Role      `json:"Role,omitempty" xml:"Role,omitempty"`
	User      *Principal `json:"User,omitempty" xml:"User,omitempty"`
}

func (s UserRole) String() string {
	return tea.Prettify(s)
}

func (s UserRole) GoString() string {
	return s.String()
}

func (s *UserRole) SetGrantTime(v int64) *UserRole {
	s.GrantTime = &v
	return s
}

func (s *UserRole) SetRole(v *Role) *UserRole {
	s.Role = v
	return s
}

func (s *UserRole) SetUser(v *Principal) *UserRole {
	s.User = v
	return s
}

type Workflow struct {
	LatestEndTime        *string `json:"LatestEndTime,omitempty" xml:"LatestEndTime,omitempty"`
	LatestInstanceId     *string `json:"LatestInstanceId,omitempty" xml:"LatestInstanceId,omitempty"`
	LatestInstanceStatus *string `json:"LatestInstanceStatus,omitempty" xml:"LatestInstanceStatus,omitempty"`
	LatestStartTime      *string `json:"LatestStartTime,omitempty" xml:"LatestStartTime,omitempty"`
}

func (s Workflow) String() string {
	return tea.Prettify(s)
}

func (s Workflow) GoString() string {
	return s.String()
}

func (s *Workflow) SetLatestEndTime(v string) *Workflow {
	s.LatestEndTime = &v
	return s
}

func (s *Workflow) SetLatestInstanceId(v string) *Workflow {
	s.LatestInstanceId = &v
	return s
}

func (s *Workflow) SetLatestInstanceStatus(v string) *Workflow {
	s.LatestInstanceStatus = &v
	return s
}

func (s *Workflow) SetLatestStartTime(v string) *Workflow {
	s.LatestStartTime = &v
	return s
}

type WorkflowInstance struct {
	BatchProgress      *int32     `json:"BatchProgress,omitempty" xml:"BatchProgress,omitempty"`
	DlfWorkflowId      *string    `json:"DlfWorkflowId,omitempty" xml:"DlfWorkflowId,omitempty"`
	EndTime            *int64     `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ExternalInstanceId *string    `json:"ExternalInstanceId,omitempty" xml:"ExternalInstanceId,omitempty"`
	RuntimeLogs        []*LogInfo `json:"RuntimeLogs,omitempty" xml:"RuntimeLogs,omitempty" type:"Repeated"`
	StartTime          *int64     `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status             *string    `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s WorkflowInstance) String() string {
	return tea.Prettify(s)
}

func (s WorkflowInstance) GoString() string {
	return s.String()
}

func (s *WorkflowInstance) SetBatchProgress(v int32) *WorkflowInstance {
	s.BatchProgress = &v
	return s
}

func (s *WorkflowInstance) SetDlfWorkflowId(v string) *WorkflowInstance {
	s.DlfWorkflowId = &v
	return s
}

func (s *WorkflowInstance) SetEndTime(v int64) *WorkflowInstance {
	s.EndTime = &v
	return s
}

func (s *WorkflowInstance) SetExternalInstanceId(v string) *WorkflowInstance {
	s.ExternalInstanceId = &v
	return s
}

func (s *WorkflowInstance) SetRuntimeLogs(v []*LogInfo) *WorkflowInstance {
	s.RuntimeLogs = v
	return s
}

func (s *WorkflowInstance) SetStartTime(v int64) *WorkflowInstance {
	s.StartTime = &v
	return s
}

func (s *WorkflowInstance) SetStatus(v string) *WorkflowInstance {
	s.Status = &v
	return s
}

type TableExtendedPrivilegesRolePrivilegesValue struct {
	CreateTime  *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantOption *bool   `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	Grantor     *string `json:"Grantor,omitempty" xml:"Grantor,omitempty"`
	GrantorType *string `json:"GrantorType,omitempty" xml:"GrantorType,omitempty"`
	Privilege   *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s TableExtendedPrivilegesRolePrivilegesValue) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedPrivilegesRolePrivilegesValue) GoString() string {
	return s.String()
}

func (s *TableExtendedPrivilegesRolePrivilegesValue) SetCreateTime(v int32) *TableExtendedPrivilegesRolePrivilegesValue {
	s.CreateTime = &v
	return s
}

func (s *TableExtendedPrivilegesRolePrivilegesValue) SetGrantOption(v bool) *TableExtendedPrivilegesRolePrivilegesValue {
	s.GrantOption = &v
	return s
}

func (s *TableExtendedPrivilegesRolePrivilegesValue) SetGrantor(v string) *TableExtendedPrivilegesRolePrivilegesValue {
	s.Grantor = &v
	return s
}

func (s *TableExtendedPrivilegesRolePrivilegesValue) SetGrantorType(v string) *TableExtendedPrivilegesRolePrivilegesValue {
	s.GrantorType = &v
	return s
}

func (s *TableExtendedPrivilegesRolePrivilegesValue) SetPrivilege(v string) *TableExtendedPrivilegesRolePrivilegesValue {
	s.Privilege = &v
	return s
}

type TableExtendedPrivilegesUserPrivilegesValue struct {
	CreateTime  *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantOption *bool   `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	Grantor     *string `json:"Grantor,omitempty" xml:"Grantor,omitempty"`
	GrantorType *string `json:"GrantorType,omitempty" xml:"GrantorType,omitempty"`
	Privilege   *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s TableExtendedPrivilegesUserPrivilegesValue) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedPrivilegesUserPrivilegesValue) GoString() string {
	return s.String()
}

func (s *TableExtendedPrivilegesUserPrivilegesValue) SetCreateTime(v int32) *TableExtendedPrivilegesUserPrivilegesValue {
	s.CreateTime = &v
	return s
}

func (s *TableExtendedPrivilegesUserPrivilegesValue) SetGrantOption(v bool) *TableExtendedPrivilegesUserPrivilegesValue {
	s.GrantOption = &v
	return s
}

func (s *TableExtendedPrivilegesUserPrivilegesValue) SetGrantor(v string) *TableExtendedPrivilegesUserPrivilegesValue {
	s.Grantor = &v
	return s
}

func (s *TableExtendedPrivilegesUserPrivilegesValue) SetGrantorType(v string) *TableExtendedPrivilegesUserPrivilegesValue {
	s.GrantorType = &v
	return s
}

func (s *TableExtendedPrivilegesUserPrivilegesValue) SetPrivilege(v string) *TableExtendedPrivilegesUserPrivilegesValue {
	s.Privilege = &v
	return s
}

type TableExtendedPrivilegesGroupPrivilegesValue struct {
	CreateTime  *int32  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	GrantOption *bool   `json:"GrantOption,omitempty" xml:"GrantOption,omitempty"`
	Grantor     *string `json:"Grantor,omitempty" xml:"Grantor,omitempty"`
	GrantorType *string `json:"GrantorType,omitempty" xml:"GrantorType,omitempty"`
	Privilege   *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
}

func (s TableExtendedPrivilegesGroupPrivilegesValue) String() string {
	return tea.Prettify(s)
}

func (s TableExtendedPrivilegesGroupPrivilegesValue) GoString() string {
	return s.String()
}

func (s *TableExtendedPrivilegesGroupPrivilegesValue) SetCreateTime(v int32) *TableExtendedPrivilegesGroupPrivilegesValue {
	s.CreateTime = &v
	return s
}

func (s *TableExtendedPrivilegesGroupPrivilegesValue) SetGrantOption(v bool) *TableExtendedPrivilegesGroupPrivilegesValue {
	s.GrantOption = &v
	return s
}

func (s *TableExtendedPrivilegesGroupPrivilegesValue) SetGrantor(v string) *TableExtendedPrivilegesGroupPrivilegesValue {
	s.Grantor = &v
	return s
}

func (s *TableExtendedPrivilegesGroupPrivilegesValue) SetGrantorType(v string) *TableExtendedPrivilegesGroupPrivilegesValue {
	s.GrantorType = &v
	return s
}

func (s *TableExtendedPrivilegesGroupPrivilegesValue) SetPrivilege(v string) *TableExtendedPrivilegesGroupPrivilegesValue {
	s.Privilege = &v
	return s
}

type AbortLockRequest struct {
	// LockId
	LockId *int64 `json:"LockId,omitempty" xml:"LockId,omitempty"`
}

func (s AbortLockRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortLockRequest) GoString() string {
	return s.String()
}

func (s *AbortLockRequest) SetLockId(v int64) *AbortLockRequest {
	s.LockId = &v
	return s
}

type AbortLockResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AbortLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AbortLockResponseBody) GoString() string {
	return s.String()
}

func (s *AbortLockResponseBody) SetCode(v string) *AbortLockResponseBody {
	s.Code = &v
	return s
}

func (s *AbortLockResponseBody) SetMessage(v string) *AbortLockResponseBody {
	s.Message = &v
	return s
}

func (s *AbortLockResponseBody) SetRequestId(v string) *AbortLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *AbortLockResponseBody) SetSuccess(v bool) *AbortLockResponseBody {
	s.Success = &v
	return s
}

type AbortLockResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AbortLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AbortLockResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortLockResponse) GoString() string {
	return s.String()
}

func (s *AbortLockResponse) SetHeaders(v map[string]*string) *AbortLockResponse {
	s.Headers = v
	return s
}

func (s *AbortLockResponse) SetStatusCode(v int32) *AbortLockResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortLockResponse) SetBody(v *AbortLockResponseBody) *AbortLockResponse {
	s.Body = v
	return s
}

type BatchCreatePartitionsRequest struct {
	CatalogId       *string           `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName    *string           `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IfNotExists     *bool             `json:"IfNotExists,omitempty" xml:"IfNotExists,omitempty"`
	NeedResult      *bool             `json:"NeedResult,omitempty" xml:"NeedResult,omitempty"`
	PartitionInputs []*PartitionInput `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty" type:"Repeated"`
	TableName       *string           `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s BatchCreatePartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreatePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchCreatePartitionsRequest) SetCatalogId(v string) *BatchCreatePartitionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchCreatePartitionsRequest) SetDatabaseName(v string) *BatchCreatePartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchCreatePartitionsRequest) SetIfNotExists(v bool) *BatchCreatePartitionsRequest {
	s.IfNotExists = &v
	return s
}

func (s *BatchCreatePartitionsRequest) SetNeedResult(v bool) *BatchCreatePartitionsRequest {
	s.NeedResult = &v
	return s
}

func (s *BatchCreatePartitionsRequest) SetPartitionInputs(v []*PartitionInput) *BatchCreatePartitionsRequest {
	s.PartitionInputs = v
	return s
}

func (s *BatchCreatePartitionsRequest) SetTableName(v string) *BatchCreatePartitionsRequest {
	s.TableName = &v
	return s
}

type BatchCreatePartitionsResponseBody struct {
	Code            *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionErrors []*PartitionError `json:"PartitionErrors,omitempty" xml:"PartitionErrors,omitempty" type:"Repeated"`
	Partitions      []*Partition      `json:"Partitions,omitempty" xml:"Partitions,omitempty" type:"Repeated"`
	RequestId       *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchCreatePartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreatePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreatePartitionsResponseBody) SetCode(v string) *BatchCreatePartitionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreatePartitionsResponseBody) SetMessage(v string) *BatchCreatePartitionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreatePartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchCreatePartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchCreatePartitionsResponseBody) SetPartitions(v []*Partition) *BatchCreatePartitionsResponseBody {
	s.Partitions = v
	return s
}

func (s *BatchCreatePartitionsResponseBody) SetRequestId(v string) *BatchCreatePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreatePartitionsResponseBody) SetSuccess(v bool) *BatchCreatePartitionsResponseBody {
	s.Success = &v
	return s
}

type BatchCreatePartitionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCreatePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreatePartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreatePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchCreatePartitionsResponse) SetHeaders(v map[string]*string) *BatchCreatePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchCreatePartitionsResponse) SetStatusCode(v int32) *BatchCreatePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreatePartitionsResponse) SetBody(v *BatchCreatePartitionsResponseBody) *BatchCreatePartitionsResponse {
	s.Body = v
	return s
}

type BatchCreateTablesRequest struct {
	CatalogId    *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string       `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IfNotExists  *bool         `json:"IfNotExists,omitempty" xml:"IfNotExists,omitempty"`
	TableInputs  []*TableInput `json:"TableInputs,omitempty" xml:"TableInputs,omitempty" type:"Repeated"`
}

func (s BatchCreateTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTablesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateTablesRequest) SetCatalogId(v string) *BatchCreateTablesRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchCreateTablesRequest) SetDatabaseName(v string) *BatchCreateTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchCreateTablesRequest) SetIfNotExists(v bool) *BatchCreateTablesRequest {
	s.IfNotExists = &v
	return s
}

func (s *BatchCreateTablesRequest) SetTableInputs(v []*TableInput) *BatchCreateTablesRequest {
	s.TableInputs = v
	return s
}

type BatchCreateTablesResponseBody struct {
	Code        *string       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableErrors []*TableError `json:"TableErrors,omitempty" xml:"TableErrors,omitempty" type:"Repeated"`
}

func (s BatchCreateTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTablesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchCreateTablesResponseBody) SetCode(v string) *BatchCreateTablesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchCreateTablesResponseBody) SetMessage(v string) *BatchCreateTablesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchCreateTablesResponseBody) SetRequestId(v string) *BatchCreateTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchCreateTablesResponseBody) SetSuccess(v bool) *BatchCreateTablesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchCreateTablesResponseBody) SetTableErrors(v []*TableError) *BatchCreateTablesResponseBody {
	s.TableErrors = v
	return s
}

type BatchCreateTablesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchCreateTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchCreateTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchCreateTablesResponse) GoString() string {
	return s.String()
}

func (s *BatchCreateTablesResponse) SetHeaders(v map[string]*string) *BatchCreateTablesResponse {
	s.Headers = v
	return s
}

func (s *BatchCreateTablesResponse) SetStatusCode(v int32) *BatchCreateTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchCreateTablesResponse) SetBody(v *BatchCreateTablesResponseBody) *BatchCreateTablesResponse {
	s.Body = v
	return s
}

type BatchDeletePartitionsRequest struct {
	CatalogId          *string                                           `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName       *string                                           `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IfExists           *bool                                             `json:"IfExists,omitempty" xml:"IfExists,omitempty"`
	PartitionValueList []*BatchDeletePartitionsRequestPartitionValueList `json:"PartitionValueList,omitempty" xml:"PartitionValueList,omitempty" type:"Repeated"`
	TableName          *string                                           `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s BatchDeletePartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeletePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeletePartitionsRequest) SetCatalogId(v string) *BatchDeletePartitionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchDeletePartitionsRequest) SetDatabaseName(v string) *BatchDeletePartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchDeletePartitionsRequest) SetIfExists(v bool) *BatchDeletePartitionsRequest {
	s.IfExists = &v
	return s
}

func (s *BatchDeletePartitionsRequest) SetPartitionValueList(v []*BatchDeletePartitionsRequestPartitionValueList) *BatchDeletePartitionsRequest {
	s.PartitionValueList = v
	return s
}

func (s *BatchDeletePartitionsRequest) SetTableName(v string) *BatchDeletePartitionsRequest {
	s.TableName = &v
	return s
}

type BatchDeletePartitionsRequestPartitionValueList struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s BatchDeletePartitionsRequestPartitionValueList) String() string {
	return tea.Prettify(s)
}

func (s BatchDeletePartitionsRequestPartitionValueList) GoString() string {
	return s.String()
}

func (s *BatchDeletePartitionsRequestPartitionValueList) SetValues(v []*string) *BatchDeletePartitionsRequestPartitionValueList {
	s.Values = v
	return s
}

type BatchDeletePartitionsResponseBody struct {
	Code            *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionErrors []*PartitionError `json:"PartitionErrors,omitempty" xml:"PartitionErrors,omitempty" type:"Repeated"`
	RequestId       *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeletePartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeletePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeletePartitionsResponseBody) SetCode(v string) *BatchDeletePartitionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeletePartitionsResponseBody) SetMessage(v string) *BatchDeletePartitionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeletePartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchDeletePartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchDeletePartitionsResponseBody) SetRequestId(v string) *BatchDeletePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeletePartitionsResponseBody) SetSuccess(v bool) *BatchDeletePartitionsResponseBody {
	s.Success = &v
	return s
}

type BatchDeletePartitionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeletePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeletePartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeletePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeletePartitionsResponse) SetHeaders(v map[string]*string) *BatchDeletePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeletePartitionsResponse) SetStatusCode(v int32) *BatchDeletePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeletePartitionsResponse) SetBody(v *BatchDeletePartitionsResponseBody) *BatchDeletePartitionsResponse {
	s.Body = v
	return s
}

type BatchDeleteTableVersionsRequest struct {
	CatalogId    *string  `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string  `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string  `json:"TableName,omitempty" xml:"TableName,omitempty"`
	VersionIds   []*int32 `json:"VersionIds,omitempty" xml:"VersionIds,omitempty" type:"Repeated"`
}

func (s BatchDeleteTableVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTableVersionsRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteTableVersionsRequest) SetCatalogId(v string) *BatchDeleteTableVersionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchDeleteTableVersionsRequest) SetDatabaseName(v string) *BatchDeleteTableVersionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchDeleteTableVersionsRequest) SetTableName(v string) *BatchDeleteTableVersionsRequest {
	s.TableName = &v
	return s
}

func (s *BatchDeleteTableVersionsRequest) SetVersionIds(v []*int32) *BatchDeleteTableVersionsRequest {
	s.VersionIds = v
	return s
}

type BatchDeleteTableVersionsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchDeleteTableVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTableVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteTableVersionsResponseBody) SetCode(v string) *BatchDeleteTableVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteTableVersionsResponseBody) SetMessage(v string) *BatchDeleteTableVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteTableVersionsResponseBody) SetRequestId(v string) *BatchDeleteTableVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteTableVersionsResponseBody) SetSuccess(v bool) *BatchDeleteTableVersionsResponseBody {
	s.Success = &v
	return s
}

type BatchDeleteTableVersionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteTableVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteTableVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTableVersionsResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteTableVersionsResponse) SetHeaders(v map[string]*string) *BatchDeleteTableVersionsResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteTableVersionsResponse) SetStatusCode(v int32) *BatchDeleteTableVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteTableVersionsResponse) SetBody(v *BatchDeleteTableVersionsResponseBody) *BatchDeleteTableVersionsResponse {
	s.Body = v
	return s
}

type BatchDeleteTablesRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// IfExists
	IfExists   *bool     `json:"IfExists,omitempty" xml:"IfExists,omitempty"`
	TableNames []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s BatchDeleteTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTablesRequest) GoString() string {
	return s.String()
}

func (s *BatchDeleteTablesRequest) SetCatalogId(v string) *BatchDeleteTablesRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchDeleteTablesRequest) SetDatabaseName(v string) *BatchDeleteTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchDeleteTablesRequest) SetIfExists(v bool) *BatchDeleteTablesRequest {
	s.IfExists = &v
	return s
}

func (s *BatchDeleteTablesRequest) SetTableNames(v []*string) *BatchDeleteTablesRequest {
	s.TableNames = v
	return s
}

type BatchDeleteTablesResponseBody struct {
	Code        *string       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableErrors []*TableError `json:"TableErrors,omitempty" xml:"TableErrors,omitempty" type:"Repeated"`
}

func (s BatchDeleteTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTablesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchDeleteTablesResponseBody) SetCode(v string) *BatchDeleteTablesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchDeleteTablesResponseBody) SetMessage(v string) *BatchDeleteTablesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchDeleteTablesResponseBody) SetRequestId(v string) *BatchDeleteTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchDeleteTablesResponseBody) SetSuccess(v bool) *BatchDeleteTablesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchDeleteTablesResponseBody) SetTableErrors(v []*TableError) *BatchDeleteTablesResponseBody {
	s.TableErrors = v
	return s
}

type BatchDeleteTablesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchDeleteTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchDeleteTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchDeleteTablesResponse) GoString() string {
	return s.String()
}

func (s *BatchDeleteTablesResponse) SetHeaders(v map[string]*string) *BatchDeleteTablesResponse {
	s.Headers = v
	return s
}

func (s *BatchDeleteTablesResponse) SetStatusCode(v int32) *BatchDeleteTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchDeleteTablesResponse) SetBody(v *BatchDeleteTablesResponseBody) *BatchDeleteTablesResponse {
	s.Body = v
	return s
}

type BatchGetPartitionColumnStatisticsRequest struct {
	CatalogId      *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNames    []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName   *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionNames []*string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty" type:"Repeated"`
	TableName      *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s BatchGetPartitionColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionColumnStatisticsRequest) SetCatalogId(v string) *BatchGetPartitionColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsRequest) SetColumnNames(v []*string) *BatchGetPartitionColumnStatisticsRequest {
	s.ColumnNames = v
	return s
}

func (s *BatchGetPartitionColumnStatisticsRequest) SetDatabaseName(v string) *BatchGetPartitionColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsRequest) SetPartitionNames(v []*string) *BatchGetPartitionColumnStatisticsRequest {
	s.PartitionNames = v
	return s
}

func (s *BatchGetPartitionColumnStatisticsRequest) SetTableName(v string) *BatchGetPartitionColumnStatisticsRequest {
	s.TableName = &v
	return s
}

type BatchGetPartitionColumnStatisticsResponseBody struct {
	Code                   *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionStatisticsMap map[string][]*ColumnStatisticsObj `json:"PartitionStatisticsMap,omitempty" xml:"PartitionStatisticsMap,omitempty"`
	RequestId              *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchGetPartitionColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionColumnStatisticsResponseBody) SetCode(v string) *BatchGetPartitionColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponseBody) SetMessage(v string) *BatchGetPartitionColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponseBody) SetPartitionStatisticsMap(v map[string][]*ColumnStatisticsObj) *BatchGetPartitionColumnStatisticsResponseBody {
	s.PartitionStatisticsMap = v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponseBody) SetRequestId(v string) *BatchGetPartitionColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponseBody) SetSuccess(v bool) *BatchGetPartitionColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type BatchGetPartitionColumnStatisticsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetPartitionColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetPartitionColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionColumnStatisticsResponse) SetHeaders(v map[string]*string) *BatchGetPartitionColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponse) SetStatusCode(v int32) *BatchGetPartitionColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetPartitionColumnStatisticsResponse) SetBody(v *BatchGetPartitionColumnStatisticsResponseBody) *BatchGetPartitionColumnStatisticsResponse {
	s.Body = v
	return s
}

type BatchGetPartitionsRequest struct {
	CatalogId          *string                                        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName       *string                                        `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IsShareSd          *bool                                          `json:"IsShareSd,omitempty" xml:"IsShareSd,omitempty"`
	PartitionValueList []*BatchGetPartitionsRequestPartitionValueList `json:"PartitionValueList,omitempty" xml:"PartitionValueList,omitempty" type:"Repeated"`
	TableName          *string                                        `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s BatchGetPartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionsRequest) SetCatalogId(v string) *BatchGetPartitionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchGetPartitionsRequest) SetDatabaseName(v string) *BatchGetPartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchGetPartitionsRequest) SetIsShareSd(v bool) *BatchGetPartitionsRequest {
	s.IsShareSd = &v
	return s
}

func (s *BatchGetPartitionsRequest) SetPartitionValueList(v []*BatchGetPartitionsRequestPartitionValueList) *BatchGetPartitionsRequest {
	s.PartitionValueList = v
	return s
}

func (s *BatchGetPartitionsRequest) SetTableName(v string) *BatchGetPartitionsRequest {
	s.TableName = &v
	return s
}

type BatchGetPartitionsRequestPartitionValueList struct {
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s BatchGetPartitionsRequestPartitionValueList) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionsRequestPartitionValueList) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionsRequestPartitionValueList) SetValues(v []*string) *BatchGetPartitionsRequestPartitionValueList {
	s.Values = v
	return s
}

type BatchGetPartitionsResponseBody struct {
	Code            *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionErrors []*PartitionError `json:"PartitionErrors,omitempty" xml:"PartitionErrors,omitempty" type:"Repeated"`
	PartitionSpecs  []*PartitionSpec  `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
	Partitions      []*Partition      `json:"Partitions,omitempty" xml:"Partitions,omitempty" type:"Repeated"`
	RequestId       *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchGetPartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionsResponseBody) SetCode(v string) *BatchGetPartitionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetMessage(v string) *BatchGetPartitionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchGetPartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetPartitionSpecs(v []*PartitionSpec) *BatchGetPartitionsResponseBody {
	s.PartitionSpecs = v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetPartitions(v []*Partition) *BatchGetPartitionsResponseBody {
	s.Partitions = v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetRequestId(v string) *BatchGetPartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetPartitionsResponseBody) SetSuccess(v bool) *BatchGetPartitionsResponseBody {
	s.Success = &v
	return s
}

type BatchGetPartitionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetPartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetPartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetPartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetPartitionsResponse) SetHeaders(v map[string]*string) *BatchGetPartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetPartitionsResponse) SetStatusCode(v int32) *BatchGetPartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetPartitionsResponse) SetBody(v *BatchGetPartitionsResponseBody) *BatchGetPartitionsResponse {
	s.Body = v
	return s
}

type BatchGetTablesRequest struct {
	CatalogId    *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableNames   []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s BatchGetTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTablesRequest) GoString() string {
	return s.String()
}

func (s *BatchGetTablesRequest) SetCatalogId(v string) *BatchGetTablesRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchGetTablesRequest) SetDatabaseName(v string) *BatchGetTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchGetTablesRequest) SetTableNames(v []*string) *BatchGetTablesRequest {
	s.TableNames = v
	return s
}

type BatchGetTablesResponseBody struct {
	Code        *string       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableErrors []*TableError `json:"TableErrors,omitempty" xml:"TableErrors,omitempty" type:"Repeated"`
	Tables      []*Table      `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s BatchGetTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTablesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetTablesResponseBody) SetCode(v string) *BatchGetTablesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchGetTablesResponseBody) SetMessage(v string) *BatchGetTablesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchGetTablesResponseBody) SetRequestId(v string) *BatchGetTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGetTablesResponseBody) SetSuccess(v bool) *BatchGetTablesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchGetTablesResponseBody) SetTableErrors(v []*TableError) *BatchGetTablesResponseBody {
	s.TableErrors = v
	return s
}

func (s *BatchGetTablesResponseBody) SetTables(v []*Table) *BatchGetTablesResponseBody {
	s.Tables = v
	return s
}

type BatchGetTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGetTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetTablesResponse) GoString() string {
	return s.String()
}

func (s *BatchGetTablesResponse) SetHeaders(v map[string]*string) *BatchGetTablesResponse {
	s.Headers = v
	return s
}

func (s *BatchGetTablesResponse) SetStatusCode(v int32) *BatchGetTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGetTablesResponse) SetBody(v *BatchGetTablesResponseBody) *BatchGetTablesResponse {
	s.Body = v
	return s
}

type BatchGrantPermissionsRequest struct {
	// catalogId
	CatalogId          *string             `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	GrantRevokeEntries []*GrantRevokeEntry `json:"GrantRevokeEntries,omitempty" xml:"GrantRevokeEntries,omitempty" type:"Repeated"`
	Type               *string             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchGrantPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsRequest) SetCatalogId(v string) *BatchGrantPermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchGrantPermissionsRequest) SetGrantRevokeEntries(v []*GrantRevokeEntry) *BatchGrantPermissionsRequest {
	s.GrantRevokeEntries = v
	return s
}

func (s *BatchGrantPermissionsRequest) SetType(v string) *BatchGrantPermissionsRequest {
	s.Type = &v
	return s
}

type BatchGrantPermissionsResponseBody struct {
	// result
	BatchGrantRevokeFailureResult []*GrantRevokeFailureEntry `json:"BatchGrantRevokeFailureResult,omitempty" xml:"BatchGrantRevokeFailureResult,omitempty" type:"Repeated"`
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchGrantPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponseBody) SetBatchGrantRevokeFailureResult(v []*GrantRevokeFailureEntry) *BatchGrantPermissionsResponseBody {
	s.BatchGrantRevokeFailureResult = v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetCode(v string) *BatchGrantPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetMessage(v string) *BatchGrantPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetRequestId(v string) *BatchGrantPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchGrantPermissionsResponseBody) SetSuccess(v bool) *BatchGrantPermissionsResponseBody {
	s.Success = &v
	return s
}

type BatchGrantPermissionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchGrantPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGrantPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchGrantPermissionsResponse) SetHeaders(v map[string]*string) *BatchGrantPermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchGrantPermissionsResponse) SetStatusCode(v int32) *BatchGrantPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchGrantPermissionsResponse) SetBody(v *BatchGrantPermissionsResponseBody) *BatchGrantPermissionsResponse {
	s.Body = v
	return s
}

type BatchRevokePermissionsRequest struct {
	// catalogId
	CatalogId          *string             `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	GrantRevokeEntries []*GrantRevokeEntry `json:"GrantRevokeEntries,omitempty" xml:"GrantRevokeEntries,omitempty" type:"Repeated"`
	Type               *string             `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s BatchRevokePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsRequest) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsRequest) SetCatalogId(v string) *BatchRevokePermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchRevokePermissionsRequest) SetGrantRevokeEntries(v []*GrantRevokeEntry) *BatchRevokePermissionsRequest {
	s.GrantRevokeEntries = v
	return s
}

func (s *BatchRevokePermissionsRequest) SetType(v string) *BatchRevokePermissionsRequest {
	s.Type = &v
	return s
}

type BatchRevokePermissionsResponseBody struct {
	// result
	BatchGrantRevokeFailureResult []*GrantRevokeFailureEntry `json:"BatchGrantRevokeFailureResult,omitempty" xml:"BatchGrantRevokeFailureResult,omitempty" type:"Repeated"`
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchRevokePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponseBody) SetBatchGrantRevokeFailureResult(v []*GrantRevokeFailureEntry) *BatchRevokePermissionsResponseBody {
	s.BatchGrantRevokeFailureResult = v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetCode(v string) *BatchRevokePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetMessage(v string) *BatchRevokePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetRequestId(v string) *BatchRevokePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRevokePermissionsResponseBody) SetSuccess(v bool) *BatchRevokePermissionsResponseBody {
	s.Success = &v
	return s
}

type BatchRevokePermissionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchRevokePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchRevokePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchRevokePermissionsResponse) GoString() string {
	return s.String()
}

func (s *BatchRevokePermissionsResponse) SetHeaders(v map[string]*string) *BatchRevokePermissionsResponse {
	s.Headers = v
	return s
}

func (s *BatchRevokePermissionsResponse) SetStatusCode(v int32) *BatchRevokePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRevokePermissionsResponse) SetBody(v *BatchRevokePermissionsResponseBody) *BatchRevokePermissionsResponse {
	s.Body = v
	return s
}

type BatchUpdatePartitionsRequest struct {
	CatalogId       *string           `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName    *string           `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionInputs []*PartitionInput `json:"PartitionInputs,omitempty" xml:"PartitionInputs,omitempty" type:"Repeated"`
	TableName       *string           `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s BatchUpdatePartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdatePartitionsRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdatePartitionsRequest) SetCatalogId(v string) *BatchUpdatePartitionsRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchUpdatePartitionsRequest) SetDatabaseName(v string) *BatchUpdatePartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchUpdatePartitionsRequest) SetPartitionInputs(v []*PartitionInput) *BatchUpdatePartitionsRequest {
	s.PartitionInputs = v
	return s
}

func (s *BatchUpdatePartitionsRequest) SetTableName(v string) *BatchUpdatePartitionsRequest {
	s.TableName = &v
	return s
}

type BatchUpdatePartitionsResponseBody struct {
	Code            *string           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message         *string           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionErrors []*PartitionError `json:"PartitionErrors,omitempty" xml:"PartitionErrors,omitempty" type:"Repeated"`
	RequestId       *string           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success         *bool             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchUpdatePartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdatePartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdatePartitionsResponseBody) SetCode(v string) *BatchUpdatePartitionsResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUpdatePartitionsResponseBody) SetMessage(v string) *BatchUpdatePartitionsResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUpdatePartitionsResponseBody) SetPartitionErrors(v []*PartitionError) *BatchUpdatePartitionsResponseBody {
	s.PartitionErrors = v
	return s
}

func (s *BatchUpdatePartitionsResponseBody) SetRequestId(v string) *BatchUpdatePartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdatePartitionsResponseBody) SetSuccess(v bool) *BatchUpdatePartitionsResponseBody {
	s.Success = &v
	return s
}

type BatchUpdatePartitionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdatePartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdatePartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdatePartitionsResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdatePartitionsResponse) SetHeaders(v map[string]*string) *BatchUpdatePartitionsResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdatePartitionsResponse) SetStatusCode(v int32) *BatchUpdatePartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdatePartitionsResponse) SetBody(v *BatchUpdatePartitionsResponseBody) *BatchUpdatePartitionsResponse {
	s.Body = v
	return s
}

type BatchUpdateTablesRequest struct {
	CatalogId    *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string       `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IsAsync      *bool         `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
	TableInputs  []*TableInput `json:"TableInputs,omitempty" xml:"TableInputs,omitempty" type:"Repeated"`
}

func (s BatchUpdateTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTablesRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateTablesRequest) SetCatalogId(v string) *BatchUpdateTablesRequest {
	s.CatalogId = &v
	return s
}

func (s *BatchUpdateTablesRequest) SetDatabaseName(v string) *BatchUpdateTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *BatchUpdateTablesRequest) SetIsAsync(v bool) *BatchUpdateTablesRequest {
	s.IsAsync = &v
	return s
}

func (s *BatchUpdateTablesRequest) SetTableInputs(v []*TableInput) *BatchUpdateTablesRequest {
	s.TableInputs = v
	return s
}

type BatchUpdateTablesResponseBody struct {
	Code        *string       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableErrors []*TableError `json:"TableErrors,omitempty" xml:"TableErrors,omitempty" type:"Repeated"`
	TaskId      *string       `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s BatchUpdateTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTablesResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateTablesResponseBody) SetCode(v string) *BatchUpdateTablesResponseBody {
	s.Code = &v
	return s
}

func (s *BatchUpdateTablesResponseBody) SetMessage(v string) *BatchUpdateTablesResponseBody {
	s.Message = &v
	return s
}

func (s *BatchUpdateTablesResponseBody) SetRequestId(v string) *BatchUpdateTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateTablesResponseBody) SetSuccess(v bool) *BatchUpdateTablesResponseBody {
	s.Success = &v
	return s
}

func (s *BatchUpdateTablesResponseBody) SetTableErrors(v []*TableError) *BatchUpdateTablesResponseBody {
	s.TableErrors = v
	return s
}

func (s *BatchUpdateTablesResponseBody) SetTaskId(v string) *BatchUpdateTablesResponseBody {
	s.TaskId = &v
	return s
}

type BatchUpdateTablesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchUpdateTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchUpdateTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchUpdateTablesResponse) GoString() string {
	return s.String()
}

func (s *BatchUpdateTablesResponse) SetHeaders(v map[string]*string) *BatchUpdateTablesResponse {
	s.Headers = v
	return s
}

func (s *BatchUpdateTablesResponse) SetStatusCode(v int32) *BatchUpdateTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchUpdateTablesResponse) SetBody(v *BatchUpdateTablesResponseBody) *BatchUpdateTablesResponse {
	s.Body = v
	return s
}

type CancelQueryRequest struct {
	QueryId *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s CancelQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelQueryRequest) GoString() string {
	return s.String()
}

func (s *CancelQueryRequest) SetQueryId(v string) *CancelQueryRequest {
	s.QueryId = &v
	return s
}

type CancelQueryResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CancelQueryResponseBody) SetData(v string) *CancelQueryResponseBody {
	s.Data = &v
	return s
}

func (s *CancelQueryResponseBody) SetRequestId(v string) *CancelQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelQueryResponseBody) SetSuccess(v bool) *CancelQueryResponseBody {
	s.Success = &v
	return s
}

type CancelQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelQueryResponse) GoString() string {
	return s.String()
}

func (s *CancelQueryResponse) SetHeaders(v map[string]*string) *CancelQueryResponse {
	s.Headers = v
	return s
}

func (s *CancelQueryResponse) SetStatusCode(v int32) *CancelQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelQueryResponse) SetBody(v *CancelQueryResponseBody) *CancelQueryResponse {
	s.Body = v
	return s
}

type CheckPermissionsRequest struct {
	Body *AccessRequest `json:"Body,omitempty" xml:"Body,omitempty"`
}

func (s CheckPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CheckPermissionsRequest) SetBody(v *AccessRequest) *CheckPermissionsRequest {
	s.Body = v
	return s
}

type CheckPermissionsResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CheckPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckPermissionsResponseBody) SetCode(v string) *CheckPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *CheckPermissionsResponseBody) SetMessage(v string) *CheckPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *CheckPermissionsResponseBody) SetRequestId(v string) *CheckPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckPermissionsResponseBody) SetSuccess(v bool) *CheckPermissionsResponseBody {
	s.Success = &v
	return s
}

type CheckPermissionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CheckPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CheckPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckPermissionsResponse) GoString() string {
	return s.String()
}

func (s *CheckPermissionsResponse) SetHeaders(v map[string]*string) *CheckPermissionsResponse {
	s.Headers = v
	return s
}

func (s *CheckPermissionsResponse) SetStatusCode(v int32) *CheckPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckPermissionsResponse) SetBody(v *CheckPermissionsResponseBody) *CheckPermissionsResponse {
	s.Body = v
	return s
}

type CreateCatalogRequest struct {
	// cataloginput
	CatalogInput *CatalogInput `json:"CatalogInput,omitempty" xml:"CatalogInput,omitempty"`
}

func (s CreateCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCatalogRequest) GoString() string {
	return s.String()
}

func (s *CreateCatalogRequest) SetCatalogInput(v *CatalogInput) *CreateCatalogRequest {
	s.CatalogInput = v
	return s
}

type CreateCatalogResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCatalogResponseBody) SetCode(v string) *CreateCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCatalogResponseBody) SetMessage(v string) *CreateCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCatalogResponseBody) SetRequestId(v string) *CreateCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCatalogResponseBody) SetSuccess(v bool) *CreateCatalogResponseBody {
	s.Success = &v
	return s
}

type CreateCatalogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCatalogResponse) GoString() string {
	return s.String()
}

func (s *CreateCatalogResponse) SetHeaders(v map[string]*string) *CreateCatalogResponse {
	s.Headers = v
	return s
}

func (s *CreateCatalogResponse) SetStatusCode(v int32) *CreateCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCatalogResponse) SetBody(v *CreateCatalogResponseBody) *CreateCatalogResponse {
	s.Body = v
	return s
}

type CreateDatabaseRequest struct {
	CatalogId     *string        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseInput *DatabaseInput `json:"DatabaseInput,omitempty" xml:"DatabaseInput,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) SetCatalogId(v string) *CreateDatabaseRequest {
	s.CatalogId = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseInput(v *DatabaseInput) *CreateDatabaseRequest {
	s.DatabaseInput = v
	return s
}

type CreateDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) SetCode(v string) *CreateDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetMessage(v string) *CreateDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetSuccess(v bool) *CreateDatabaseResponseBody {
	s.Success = &v
	return s
}

type CreateDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponse) SetHeaders(v map[string]*string) *CreateDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseResponse) SetStatusCode(v int32) *CreateDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseResponse) SetBody(v *CreateDatabaseResponseBody) *CreateDatabaseResponse {
	s.Body = v
	return s
}

type CreateFunctionRequest struct {
	CatalogId     *string        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName  *string        `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionInput *FunctionInput `json:"FunctionInput,omitempty" xml:"FunctionInput,omitempty"`
}

func (s CreateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionRequest) SetCatalogId(v string) *CreateFunctionRequest {
	s.CatalogId = &v
	return s
}

func (s *CreateFunctionRequest) SetDatabaseName(v string) *CreateFunctionRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateFunctionRequest) SetFunctionInput(v *FunctionInput) *CreateFunctionRequest {
	s.FunctionInput = v
	return s
}

type CreateFunctionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponseBody) SetCode(v string) *CreateFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionResponseBody) SetMessage(v string) *CreateFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionResponseBody) SetRequestId(v string) *CreateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResponseBody) SetSuccess(v bool) *CreateFunctionResponseBody {
	s.Success = &v
	return s
}

type CreateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResponse) SetHeaders(v map[string]*string) *CreateFunctionResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResponse) SetStatusCode(v int32) *CreateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResponse) SetBody(v *CreateFunctionResponseBody) *CreateFunctionResponse {
	s.Body = v
	return s
}

type CreateLockRequest struct {
	// LockObjList
	LockObjList []*LockObj `json:"LockObjList,omitempty" xml:"LockObjList,omitempty" type:"Repeated"`
}

func (s CreateLockRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateLockRequest) GoString() string {
	return s.String()
}

func (s *CreateLockRequest) SetLockObjList(v []*LockObj) *CreateLockRequest {
	s.LockObjList = v
	return s
}

type CreateLockResponseBody struct {
	// Code
	Code       *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	LockStatus *LockStatus `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateLockResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLockResponseBody) SetCode(v string) *CreateLockResponseBody {
	s.Code = &v
	return s
}

func (s *CreateLockResponseBody) SetLockStatus(v *LockStatus) *CreateLockResponseBody {
	s.LockStatus = v
	return s
}

func (s *CreateLockResponseBody) SetMessage(v string) *CreateLockResponseBody {
	s.Message = &v
	return s
}

func (s *CreateLockResponseBody) SetRequestId(v string) *CreateLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLockResponseBody) SetSuccess(v bool) *CreateLockResponseBody {
	s.Success = &v
	return s
}

type CreateLockResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateLockResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateLockResponse) GoString() string {
	return s.String()
}

func (s *CreateLockResponse) SetHeaders(v map[string]*string) *CreateLockResponse {
	s.Headers = v
	return s
}

func (s *CreateLockResponse) SetStatusCode(v int32) *CreateLockResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLockResponse) SetBody(v *CreateLockResponseBody) *CreateLockResponse {
	s.Body = v
	return s
}

type CreatePartitionRequest struct {
	CatalogId      *string         `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName   *string         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IfNotExists    *bool           `json:"IfNotExists,omitempty" xml:"IfNotExists,omitempty"`
	NeedResult     *bool           `json:"NeedResult,omitempty" xml:"NeedResult,omitempty"`
	PartitionInput *PartitionInput `json:"PartitionInput,omitempty" xml:"PartitionInput,omitempty"`
	TableName      *string         `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CreatePartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePartitionRequest) GoString() string {
	return s.String()
}

func (s *CreatePartitionRequest) SetCatalogId(v string) *CreatePartitionRequest {
	s.CatalogId = &v
	return s
}

func (s *CreatePartitionRequest) SetDatabaseName(v string) *CreatePartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreatePartitionRequest) SetIfNotExists(v bool) *CreatePartitionRequest {
	s.IfNotExists = &v
	return s
}

func (s *CreatePartitionRequest) SetNeedResult(v bool) *CreatePartitionRequest {
	s.NeedResult = &v
	return s
}

func (s *CreatePartitionRequest) SetPartitionInput(v *PartitionInput) *CreatePartitionRequest {
	s.PartitionInput = v
	return s
}

func (s *CreatePartitionRequest) SetTableName(v string) *CreatePartitionRequest {
	s.TableName = &v
	return s
}

type CreatePartitionResponseBody struct {
	Code      *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string    `json:"Message,omitempty" xml:"Message,omitempty"`
	Partition *Partition `json:"Partition,omitempty" xml:"Partition,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreatePartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePartitionResponseBody) SetCode(v string) *CreatePartitionResponseBody {
	s.Code = &v
	return s
}

func (s *CreatePartitionResponseBody) SetMessage(v string) *CreatePartitionResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePartitionResponseBody) SetPartition(v *Partition) *CreatePartitionResponseBody {
	s.Partition = v
	return s
}

func (s *CreatePartitionResponseBody) SetRequestId(v string) *CreatePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePartitionResponseBody) SetSuccess(v bool) *CreatePartitionResponseBody {
	s.Success = &v
	return s
}

type CreatePartitionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreatePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePartitionResponse) GoString() string {
	return s.String()
}

func (s *CreatePartitionResponse) SetHeaders(v map[string]*string) *CreatePartitionResponse {
	s.Headers = v
	return s
}

func (s *CreatePartitionResponse) SetStatusCode(v int32) *CreatePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePartitionResponse) SetBody(v *CreatePartitionResponseBody) *CreatePartitionResponse {
	s.Body = v
	return s
}

type CreateRoleRequest struct {
	Body *RoleInput `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateRoleRequest) SetBody(v *RoleInput) *CreateRoleRequest {
	s.Body = v
	return s
}

type CreateRoleResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoleResponseBody) SetCode(v string) *CreateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRoleResponseBody) SetMessage(v string) *CreateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRoleResponseBody) SetRequestId(v string) *CreateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoleResponseBody) SetSuccess(v bool) *CreateRoleResponseBody {
	s.Success = &v
	return s
}

type CreateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateRoleResponse) SetHeaders(v map[string]*string) *CreateRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateRoleResponse) SetStatusCode(v int32) *CreateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoleResponse) SetBody(v *CreateRoleResponseBody) *CreateRoleResponse {
	s.Body = v
	return s
}

type CreateTableRequest struct {
	CatalogId    *string     `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string     `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableInput   *TableInput `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
}

func (s CreateTableRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTableRequest) GoString() string {
	return s.String()
}

func (s *CreateTableRequest) SetCatalogId(v string) *CreateTableRequest {
	s.CatalogId = &v
	return s
}

func (s *CreateTableRequest) SetDatabaseName(v string) *CreateTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateTableRequest) SetTableInput(v *TableInput) *CreateTableRequest {
	s.TableInput = v
	return s
}

type CreateTableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTableResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTableResponseBody) SetCode(v string) *CreateTableResponseBody {
	s.Code = &v
	return s
}

func (s *CreateTableResponseBody) SetMessage(v string) *CreateTableResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTableResponseBody) SetRequestId(v string) *CreateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTableResponseBody) SetSuccess(v bool) *CreateTableResponseBody {
	s.Success = &v
	return s
}

type CreateTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTableResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTableResponse) SetHeaders(v map[string]*string) *CreateTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTableResponse) SetStatusCode(v int32) *CreateTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableResponse) SetBody(v *CreateTableResponseBody) *CreateTableResponse {
	s.Body = v
	return s
}

type DeleteCatalogRequest struct {
	// CatalogId
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IsAsync   *bool   `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
}

func (s DeleteCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteCatalogRequest) GoString() string {
	return s.String()
}

func (s *DeleteCatalogRequest) SetCatalogId(v string) *DeleteCatalogRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteCatalogRequest) SetIsAsync(v bool) *DeleteCatalogRequest {
	s.IsAsync = &v
	return s
}

type DeleteCatalogResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Request is success or not
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId  *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCatalogResponseBody) SetCode(v string) *DeleteCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteCatalogResponseBody) SetMessage(v string) *DeleteCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteCatalogResponseBody) SetRequestId(v string) *DeleteCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCatalogResponseBody) SetSuccess(v bool) *DeleteCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCatalogResponseBody) SetTaskId(v string) *DeleteCatalogResponseBody {
	s.TaskId = &v
	return s
}

type DeleteCatalogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCatalogResponse) GoString() string {
	return s.String()
}

func (s *DeleteCatalogResponse) SetHeaders(v map[string]*string) *DeleteCatalogResponse {
	s.Headers = v
	return s
}

func (s *DeleteCatalogResponse) SetStatusCode(v int32) *DeleteCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCatalogResponse) SetBody(v *DeleteCatalogResponseBody) *DeleteCatalogResponse {
	s.Body = v
	return s
}

type DeleteDatabaseRequest struct {
	Async     *bool   `json:"Async,omitempty" xml:"Async,omitempty"`
	Cascade   *bool   `json:"Cascade,omitempty" xml:"Cascade,omitempty"`
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetAsync(v bool) *DeleteDatabaseRequest {
	s.Async = &v
	return s
}

func (s *DeleteDatabaseRequest) SetCascade(v bool) *DeleteDatabaseRequest {
	s.Cascade = &v
	return s
}

func (s *DeleteDatabaseRequest) SetCatalogId(v string) *DeleteDatabaseRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetName(v string) *DeleteDatabaseRequest {
	s.Name = &v
	return s
}

type DeleteDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DeleteDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponseBody) SetCode(v string) *DeleteDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetMessage(v string) *DeleteDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetRequestId(v string) *DeleteDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetSuccess(v bool) *DeleteDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatabaseResponseBody) SetTaskId(v string) *DeleteDatabaseResponseBody {
	s.TaskId = &v
	return s
}

type DeleteDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetHeaders(v map[string]*string) *DeleteDatabaseResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabaseResponse) SetStatusCode(v int32) *DeleteDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabaseResponse) SetBody(v *DeleteDatabaseResponseBody) *DeleteDatabaseResponse {
	s.Body = v
	return s
}

type DeleteFunctionRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DeleteFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionRequest) GoString() string {
	return s.String()
}

func (s *DeleteFunctionRequest) SetCatalogId(v string) *DeleteFunctionRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteFunctionRequest) SetDatabaseName(v string) *DeleteFunctionRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteFunctionRequest) SetFunctionName(v string) *DeleteFunctionRequest {
	s.FunctionName = &v
	return s
}

type DeleteFunctionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponseBody) SetCode(v string) *DeleteFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetMessage(v string) *DeleteFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetRequestId(v string) *DeleteFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResponseBody) SetSuccess(v bool) *DeleteFunctionResponseBody {
	s.Success = &v
	return s
}

type DeleteFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResponse) SetHeaders(v map[string]*string) *DeleteFunctionResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResponse) SetStatusCode(v int32) *DeleteFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResponse) SetBody(v *DeleteFunctionResponseBody) *DeleteFunctionResponse {
	s.Body = v
	return s
}

type DeletePartitionRequest struct {
	CatalogId       *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName    *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IfExists        *bool     `json:"IfExists,omitempty" xml:"IfExists,omitempty"`
	PartitionValues []*string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
	TableName       *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeletePartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionRequest) GoString() string {
	return s.String()
}

func (s *DeletePartitionRequest) SetCatalogId(v string) *DeletePartitionRequest {
	s.CatalogId = &v
	return s
}

func (s *DeletePartitionRequest) SetDatabaseName(v string) *DeletePartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeletePartitionRequest) SetIfExists(v bool) *DeletePartitionRequest {
	s.IfExists = &v
	return s
}

func (s *DeletePartitionRequest) SetPartitionValues(v []*string) *DeletePartitionRequest {
	s.PartitionValues = v
	return s
}

func (s *DeletePartitionRequest) SetTableName(v string) *DeletePartitionRequest {
	s.TableName = &v
	return s
}

type DeletePartitionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePartitionResponseBody) SetCode(v string) *DeletePartitionResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePartitionResponseBody) SetMessage(v string) *DeletePartitionResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePartitionResponseBody) SetRequestId(v string) *DeletePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePartitionResponseBody) SetSuccess(v bool) *DeletePartitionResponseBody {
	s.Success = &v
	return s
}

type DeletePartitionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionResponse) GoString() string {
	return s.String()
}

func (s *DeletePartitionResponse) SetHeaders(v map[string]*string) *DeletePartitionResponse {
	s.Headers = v
	return s
}

func (s *DeletePartitionResponse) SetStatusCode(v int32) *DeletePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePartitionResponse) SetBody(v *DeletePartitionResponseBody) *DeletePartitionResponse {
	s.Body = v
	return s
}

type DeletePartitionColumnStatisticsRequest struct {
	CatalogId      *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNames    []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName   *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionNames []*string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty" type:"Repeated"`
	TableName      *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeletePartitionColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DeletePartitionColumnStatisticsRequest) SetCatalogId(v string) *DeletePartitionColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *DeletePartitionColumnStatisticsRequest) SetColumnNames(v []*string) *DeletePartitionColumnStatisticsRequest {
	s.ColumnNames = v
	return s
}

func (s *DeletePartitionColumnStatisticsRequest) SetDatabaseName(v string) *DeletePartitionColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeletePartitionColumnStatisticsRequest) SetPartitionNames(v []*string) *DeletePartitionColumnStatisticsRequest {
	s.PartitionNames = v
	return s
}

func (s *DeletePartitionColumnStatisticsRequest) SetTableName(v string) *DeletePartitionColumnStatisticsRequest {
	s.TableName = &v
	return s
}

type DeletePartitionColumnStatisticsShrinkRequest struct {
	CatalogId            *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNamesShrink    *string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty"`
	DatabaseName         *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionNamesShrink *string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeletePartitionColumnStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionColumnStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeletePartitionColumnStatisticsShrinkRequest) SetCatalogId(v string) *DeletePartitionColumnStatisticsShrinkRequest {
	s.CatalogId = &v
	return s
}

func (s *DeletePartitionColumnStatisticsShrinkRequest) SetColumnNamesShrink(v string) *DeletePartitionColumnStatisticsShrinkRequest {
	s.ColumnNamesShrink = &v
	return s
}

func (s *DeletePartitionColumnStatisticsShrinkRequest) SetDatabaseName(v string) *DeletePartitionColumnStatisticsShrinkRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeletePartitionColumnStatisticsShrinkRequest) SetPartitionNamesShrink(v string) *DeletePartitionColumnStatisticsShrinkRequest {
	s.PartitionNamesShrink = &v
	return s
}

func (s *DeletePartitionColumnStatisticsShrinkRequest) SetTableName(v string) *DeletePartitionColumnStatisticsShrinkRequest {
	s.TableName = &v
	return s
}

type DeletePartitionColumnStatisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePartitionColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePartitionColumnStatisticsResponseBody) SetCode(v string) *DeletePartitionColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePartitionColumnStatisticsResponseBody) SetMessage(v string) *DeletePartitionColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePartitionColumnStatisticsResponseBody) SetRequestId(v string) *DeletePartitionColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePartitionColumnStatisticsResponseBody) SetSuccess(v bool) *DeletePartitionColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type DeletePartitionColumnStatisticsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeletePartitionColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePartitionColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePartitionColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DeletePartitionColumnStatisticsResponse) SetHeaders(v map[string]*string) *DeletePartitionColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DeletePartitionColumnStatisticsResponse) SetStatusCode(v int32) *DeletePartitionColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePartitionColumnStatisticsResponse) SetBody(v *DeletePartitionColumnStatisticsResponseBody) *DeletePartitionColumnStatisticsResponse {
	s.Body = v
	return s
}

type DeleteRoleRequest struct {
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s DeleteRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRoleRequest) SetRoleName(v string) *DeleteRoleRequest {
	s.RoleName = &v
	return s
}

type DeleteRoleResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponseBody) SetCode(v string) *DeleteRoleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRoleResponseBody) SetMessage(v string) *DeleteRoleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRoleResponseBody) SetRequestId(v string) *DeleteRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoleResponseBody) SetSuccess(v bool) *DeleteRoleResponseBody {
	s.Success = &v
	return s
}

type DeleteRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRoleResponse) GoString() string {
	return s.String()
}

func (s *DeleteRoleResponse) SetHeaders(v map[string]*string) *DeleteRoleResponse {
	s.Headers = v
	return s
}

func (s *DeleteRoleResponse) SetStatusCode(v int32) *DeleteRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRoleResponse) SetBody(v *DeleteRoleResponseBody) *DeleteRoleResponse {
	s.Body = v
	return s
}

type DeleteTableRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeleteTableRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableRequest) SetCatalogId(v string) *DeleteTableRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteTableRequest) SetDatabaseName(v string) *DeleteTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteTableRequest) SetTableName(v string) *DeleteTableRequest {
	s.TableName = &v
	return s
}

type DeleteTableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableResponseBody) SetCode(v string) *DeleteTableResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTableResponseBody) SetMessage(v string) *DeleteTableResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTableResponseBody) SetRequestId(v string) *DeleteTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableResponseBody) SetSuccess(v bool) *DeleteTableResponseBody {
	s.Success = &v
	return s
}

type DeleteTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTableResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableResponse) SetHeaders(v map[string]*string) *DeleteTableResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableResponse) SetStatusCode(v int32) *DeleteTableResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableResponse) SetBody(v *DeleteTableResponseBody) *DeleteTableResponse {
	s.Body = v
	return s
}

type DeleteTableColumnStatisticsRequest struct {
	CatalogId    *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNames  []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeleteTableColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableColumnStatisticsRequest) SetCatalogId(v string) *DeleteTableColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteTableColumnStatisticsRequest) SetColumnNames(v []*string) *DeleteTableColumnStatisticsRequest {
	s.ColumnNames = v
	return s
}

func (s *DeleteTableColumnStatisticsRequest) SetDatabaseName(v string) *DeleteTableColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteTableColumnStatisticsRequest) SetTableName(v string) *DeleteTableColumnStatisticsRequest {
	s.TableName = &v
	return s
}

type DeleteTableColumnStatisticsShrinkRequest struct {
	CatalogId         *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNamesShrink *string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty"`
	DatabaseName      *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DeleteTableColumnStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableColumnStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableColumnStatisticsShrinkRequest) SetCatalogId(v string) *DeleteTableColumnStatisticsShrinkRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteTableColumnStatisticsShrinkRequest) SetColumnNamesShrink(v string) *DeleteTableColumnStatisticsShrinkRequest {
	s.ColumnNamesShrink = &v
	return s
}

func (s *DeleteTableColumnStatisticsShrinkRequest) SetDatabaseName(v string) *DeleteTableColumnStatisticsShrinkRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteTableColumnStatisticsShrinkRequest) SetTableName(v string) *DeleteTableColumnStatisticsShrinkRequest {
	s.TableName = &v
	return s
}

type DeleteTableColumnStatisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTableColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableColumnStatisticsResponseBody) SetCode(v string) *DeleteTableColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTableColumnStatisticsResponseBody) SetMessage(v string) *DeleteTableColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTableColumnStatisticsResponseBody) SetRequestId(v string) *DeleteTableColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableColumnStatisticsResponseBody) SetSuccess(v bool) *DeleteTableColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type DeleteTableColumnStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTableColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTableColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableColumnStatisticsResponse) SetHeaders(v map[string]*string) *DeleteTableColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableColumnStatisticsResponse) SetStatusCode(v int32) *DeleteTableColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableColumnStatisticsResponse) SetBody(v *DeleteTableColumnStatisticsResponseBody) *DeleteTableColumnStatisticsResponse {
	s.Body = v
	return s
}

type DeleteTableVersionRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	VersionId    *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteTableVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteTableVersionRequest) SetCatalogId(v string) *DeleteTableVersionRequest {
	s.CatalogId = &v
	return s
}

func (s *DeleteTableVersionRequest) SetDatabaseName(v string) *DeleteTableVersionRequest {
	s.DatabaseName = &v
	return s
}

func (s *DeleteTableVersionRequest) SetTableName(v string) *DeleteTableVersionRequest {
	s.TableName = &v
	return s
}

func (s *DeleteTableVersionRequest) SetVersionId(v int32) *DeleteTableVersionRequest {
	s.VersionId = &v
	return s
}

type DeleteTableVersionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteTableVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTableVersionResponseBody) SetCode(v string) *DeleteTableVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteTableVersionResponseBody) SetMessage(v string) *DeleteTableVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteTableVersionResponseBody) SetRequestId(v string) *DeleteTableVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTableVersionResponseBody) SetSuccess(v bool) *DeleteTableVersionResponseBody {
	s.Success = &v
	return s
}

type DeleteTableVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTableVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTableVersionResponse) GoString() string {
	return s.String()
}

func (s *DeleteTableVersionResponse) SetHeaders(v map[string]*string) *DeleteTableVersionResponse {
	s.Headers = v
	return s
}

func (s *DeleteTableVersionResponse) SetStatusCode(v int32) *DeleteTableVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTableVersionResponse) SetBody(v *DeleteTableVersionResponseBody) *DeleteTableVersionResponse {
	s.Body = v
	return s
}

type DeregisterLocationRequest struct {
	LocationId *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
}

func (s DeregisterLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeregisterLocationRequest) GoString() string {
	return s.String()
}

func (s *DeregisterLocationRequest) SetLocationId(v string) *DeregisterLocationRequest {
	s.LocationId = &v
	return s
}

type DeregisterLocationResponseBody struct {
	Data      *DeregisterLocationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                               `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeregisterLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeregisterLocationResponseBody) GoString() string {
	return s.String()
}

func (s *DeregisterLocationResponseBody) SetData(v *DeregisterLocationResponseBodyData) *DeregisterLocationResponseBody {
	s.Data = v
	return s
}

func (s *DeregisterLocationResponseBody) SetRequestId(v string) *DeregisterLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeregisterLocationResponseBody) SetSuccess(v bool) *DeregisterLocationResponseBody {
	s.Success = &v
	return s
}

type DeregisterLocationResponseBodyData struct {
	// Location ID
	LocationId                            *string                              `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	StorageCollectTaskOperationResultList []*StorageCollectTaskOperationResult `json:"StorageCollectTaskOperationResultList,omitempty" xml:"StorageCollectTaskOperationResultList,omitempty" type:"Repeated"`
}

func (s DeregisterLocationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DeregisterLocationResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeregisterLocationResponseBodyData) SetLocationId(v string) *DeregisterLocationResponseBodyData {
	s.LocationId = &v
	return s
}

func (s *DeregisterLocationResponseBodyData) SetStorageCollectTaskOperationResultList(v []*StorageCollectTaskOperationResult) *DeregisterLocationResponseBodyData {
	s.StorageCollectTaskOperationResultList = v
	return s
}

type DeregisterLocationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeregisterLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeregisterLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeregisterLocationResponse) GoString() string {
	return s.String()
}

func (s *DeregisterLocationResponse) SetHeaders(v map[string]*string) *DeregisterLocationResponse {
	s.Headers = v
	return s
}

func (s *DeregisterLocationResponse) SetStatusCode(v int32) *DeregisterLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterLocationResponse) SetBody(v *DeregisterLocationResponseBody) *DeregisterLocationResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	Regions   []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetSuccess(v bool) *DescribeRegionsResponseBody {
	s.Success = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ShowName    *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetDescription(v string) *DescribeRegionsResponseBodyRegions {
	s.Description = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetName(v string) *DescribeRegionsResponseBodyRegions {
	s.Name = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetShowName(v string) *DescribeRegionsResponseBodyRegions {
	s.ShowName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetType(v string) *DescribeRegionsResponseBodyRegions {
	s.Type = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type GetAsyncTaskStatusRequest struct {
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s GetAsyncTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskStatusRequest) SetCatalogId(v string) *GetAsyncTaskStatusRequest {
	s.CatalogId = &v
	return s
}

func (s *GetAsyncTaskStatusRequest) SetTaskId(v string) *GetAsyncTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetAsyncTaskStatusResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success    *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskStatus *TaskStatus `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s GetAsyncTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskStatusResponseBody) SetCode(v string) *GetAsyncTaskStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetAsyncTaskStatusResponseBody) SetMessage(v string) *GetAsyncTaskStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncTaskStatusResponseBody) SetRequestId(v string) *GetAsyncTaskStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAsyncTaskStatusResponseBody) SetSuccess(v bool) *GetAsyncTaskStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetAsyncTaskStatusResponseBody) SetTaskStatus(v *TaskStatus) *GetAsyncTaskStatusResponseBody {
	s.TaskStatus = v
	return s
}

type GetAsyncTaskStatusResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAsyncTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskStatusResponse) SetHeaders(v map[string]*string) *GetAsyncTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskStatusResponse) SetStatusCode(v int32) *GetAsyncTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTaskStatusResponse) SetBody(v *GetAsyncTaskStatusResponseBody) *GetAsyncTaskStatusResponse {
	s.Body = v
	return s
}

type GetCatalogRequest struct {
	// catalogId
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
}

func (s GetCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogRequest) SetCatalogId(v string) *GetCatalogRequest {
	s.CatalogId = &v
	return s
}

type GetCatalogResponseBody struct {
	Catalog *Catalog `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogResponseBody) SetCatalog(v *Catalog) *GetCatalogResponseBody {
	s.Catalog = v
	return s
}

func (s *GetCatalogResponseBody) SetCode(v string) *GetCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *GetCatalogResponseBody) SetMessage(v string) *GetCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *GetCatalogResponseBody) SetRequestId(v string) *GetCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogResponseBody) SetSuccess(v bool) *GetCatalogResponseBody {
	s.Success = &v
	return s
}

type GetCatalogResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogResponse) SetHeaders(v map[string]*string) *GetCatalogResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogResponse) SetStatusCode(v int32) *GetCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogResponse) SetBody(v *GetCatalogResponseBody) *GetCatalogResponse {
	s.Body = v
	return s
}

type GetCatalogSettingsRequest struct {
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
}

func (s GetCatalogSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSettingsRequest) GoString() string {
	return s.String()
}

func (s *GetCatalogSettingsRequest) SetCatalogId(v string) *GetCatalogSettingsRequest {
	s.CatalogId = &v
	return s
}

type GetCatalogSettingsResponseBody struct {
	CatalogSettings *CatalogSettings `json:"CatalogSettings,omitempty" xml:"CatalogSettings,omitempty"`
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCatalogSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCatalogSettingsResponseBody) SetCatalogSettings(v *CatalogSettings) *GetCatalogSettingsResponseBody {
	s.CatalogSettings = v
	return s
}

func (s *GetCatalogSettingsResponseBody) SetCode(v string) *GetCatalogSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCatalogSettingsResponseBody) SetMessage(v string) *GetCatalogSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *GetCatalogSettingsResponseBody) SetRequestId(v string) *GetCatalogSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCatalogSettingsResponseBody) SetSuccess(v bool) *GetCatalogSettingsResponseBody {
	s.Success = &v
	return s
}

type GetCatalogSettingsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCatalogSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCatalogSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCatalogSettingsResponse) GoString() string {
	return s.String()
}

func (s *GetCatalogSettingsResponse) SetHeaders(v map[string]*string) *GetCatalogSettingsResponse {
	s.Headers = v
	return s
}

func (s *GetCatalogSettingsResponse) SetStatusCode(v int32) *GetCatalogSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCatalogSettingsResponse) SetBody(v *GetCatalogSettingsResponseBody) *GetCatalogSettingsResponse {
	s.Body = v
	return s
}

type GetDatabaseRequest struct {
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDatabaseRequest) SetCatalogId(v string) *GetDatabaseRequest {
	s.CatalogId = &v
	return s
}

func (s *GetDatabaseRequest) SetName(v string) *GetDatabaseRequest {
	s.Name = &v
	return s
}

type GetDatabaseResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Database  *Database `json:"Database,omitempty" xml:"Database,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponseBody) SetCode(v string) *GetDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *GetDatabaseResponseBody) SetDatabase(v *Database) *GetDatabaseResponseBody {
	s.Database = v
	return s
}

func (s *GetDatabaseResponseBody) SetMessage(v string) *GetDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *GetDatabaseResponseBody) SetRequestId(v string) *GetDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatabaseResponseBody) SetSuccess(v bool) *GetDatabaseResponseBody {
	s.Success = &v
	return s
}

type GetDatabaseResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetDatabaseResponse) SetHeaders(v map[string]*string) *GetDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetDatabaseResponse) SetStatusCode(v int32) *GetDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDatabaseResponse) SetBody(v *GetDatabaseResponseBody) *GetDatabaseResponse {
	s.Body = v
	return s
}

type GetFunctionRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s GetFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionRequest) SetCatalogId(v string) *GetFunctionRequest {
	s.CatalogId = &v
	return s
}

func (s *GetFunctionRequest) SetDatabaseName(v string) *GetFunctionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetFunctionRequest) SetFunctionName(v string) *GetFunctionRequest {
	s.FunctionName = &v
	return s
}

type GetFunctionResponseBody struct {
	Code      *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Function  *Function `json:"Function,omitempty" xml:"Function,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResponseBody) SetCode(v string) *GetFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionResponseBody) SetFunction(v *Function) *GetFunctionResponseBody {
	s.Function = v
	return s
}

func (s *GetFunctionResponseBody) SetMessage(v string) *GetFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionResponseBody) SetRequestId(v string) *GetFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionResponseBody) SetSuccess(v bool) *GetFunctionResponseBody {
	s.Success = &v
	return s
}

type GetFunctionResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResponse) SetHeaders(v map[string]*string) *GetFunctionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResponse) SetStatusCode(v int32) *GetFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResponse) SetBody(v *GetFunctionResponseBody) *GetFunctionResponse {
	s.Body = v
	return s
}

type GetLifecycleRuleRequest struct {
	BizId        *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s GetLifecycleRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleRuleRequest) GoString() string {
	return s.String()
}

func (s *GetLifecycleRuleRequest) SetBizId(v string) *GetLifecycleRuleRequest {
	s.BizId = &v
	return s
}

func (s *GetLifecycleRuleRequest) SetResourceName(v string) *GetLifecycleRuleRequest {
	s.ResourceName = &v
	return s
}

type GetLifecycleRuleResponseBody struct {
	Data *LifecycleRule `json:"Data,omitempty" xml:"Data,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLifecycleRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetLifecycleRuleResponseBody) SetData(v *LifecycleRule) *GetLifecycleRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetLifecycleRuleResponseBody) SetRequestId(v string) *GetLifecycleRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLifecycleRuleResponseBody) SetSuccess(v bool) *GetLifecycleRuleResponseBody {
	s.Success = &v
	return s
}

type GetLifecycleRuleResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLifecycleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLifecycleRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLifecycleRuleResponse) GoString() string {
	return s.String()
}

func (s *GetLifecycleRuleResponse) SetHeaders(v map[string]*string) *GetLifecycleRuleResponse {
	s.Headers = v
	return s
}

func (s *GetLifecycleRuleResponse) SetStatusCode(v int32) *GetLifecycleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLifecycleRuleResponse) SetBody(v *GetLifecycleRuleResponseBody) *GetLifecycleRuleResponse {
	s.Body = v
	return s
}

type GetLockRequest struct {
	// LockId
	LockId *int64 `json:"LockId,omitempty" xml:"LockId,omitempty"`
}

func (s GetLockRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLockRequest) GoString() string {
	return s.String()
}

func (s *GetLockRequest) SetLockId(v int64) *GetLockRequest {
	s.LockId = &v
	return s
}

type GetLockResponseBody struct {
	// Code
	Code       *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	LockStatus *LockStatus `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLockResponseBody) GoString() string {
	return s.String()
}

func (s *GetLockResponseBody) SetCode(v string) *GetLockResponseBody {
	s.Code = &v
	return s
}

func (s *GetLockResponseBody) SetLockStatus(v *LockStatus) *GetLockResponseBody {
	s.LockStatus = v
	return s
}

func (s *GetLockResponseBody) SetMessage(v string) *GetLockResponseBody {
	s.Message = &v
	return s
}

func (s *GetLockResponseBody) SetRequestId(v string) *GetLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLockResponseBody) SetSuccess(v bool) *GetLockResponseBody {
	s.Success = &v
	return s
}

type GetLockResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLockResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLockResponse) GoString() string {
	return s.String()
}

func (s *GetLockResponse) SetHeaders(v map[string]*string) *GetLockResponse {
	s.Headers = v
	return s
}

func (s *GetLockResponse) SetStatusCode(v int32) *GetLockResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLockResponse) SetBody(v *GetLockResponseBody) *GetLockResponse {
	s.Body = v
	return s
}

type GetPartitionRequest struct {
	CatalogId       *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName    *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionValues []*string `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
	TableName       *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionRequest) GoString() string {
	return s.String()
}

func (s *GetPartitionRequest) SetCatalogId(v string) *GetPartitionRequest {
	s.CatalogId = &v
	return s
}

func (s *GetPartitionRequest) SetDatabaseName(v string) *GetPartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetPartitionRequest) SetPartitionValues(v []*string) *GetPartitionRequest {
	s.PartitionValues = v
	return s
}

func (s *GetPartitionRequest) SetTableName(v string) *GetPartitionRequest {
	s.TableName = &v
	return s
}

type GetPartitionResponseBody struct {
	Code      *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string    `json:"Message,omitempty" xml:"Message,omitempty"`
	Partition *Partition `json:"Partition,omitempty" xml:"Partition,omitempty"`
	RequestId *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartitionResponseBody) SetCode(v string) *GetPartitionResponseBody {
	s.Code = &v
	return s
}

func (s *GetPartitionResponseBody) SetMessage(v string) *GetPartitionResponseBody {
	s.Message = &v
	return s
}

func (s *GetPartitionResponseBody) SetPartition(v *Partition) *GetPartitionResponseBody {
	s.Partition = v
	return s
}

func (s *GetPartitionResponseBody) SetRequestId(v string) *GetPartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPartitionResponseBody) SetSuccess(v bool) *GetPartitionResponseBody {
	s.Success = &v
	return s
}

type GetPartitionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPartitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionResponse) GoString() string {
	return s.String()
}

func (s *GetPartitionResponse) SetHeaders(v map[string]*string) *GetPartitionResponse {
	s.Headers = v
	return s
}

func (s *GetPartitionResponse) SetStatusCode(v int32) *GetPartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartitionResponse) SetBody(v *GetPartitionResponseBody) *GetPartitionResponse {
	s.Body = v
	return s
}

type GetPartitionColumnStatisticsRequest struct {
	CatalogId      *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNames    []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName   *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionNames []*string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty" type:"Repeated"`
	TableName      *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPartitionColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetPartitionColumnStatisticsRequest) SetCatalogId(v string) *GetPartitionColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *GetPartitionColumnStatisticsRequest) SetColumnNames(v []*string) *GetPartitionColumnStatisticsRequest {
	s.ColumnNames = v
	return s
}

func (s *GetPartitionColumnStatisticsRequest) SetDatabaseName(v string) *GetPartitionColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetPartitionColumnStatisticsRequest) SetPartitionNames(v []*string) *GetPartitionColumnStatisticsRequest {
	s.PartitionNames = v
	return s
}

func (s *GetPartitionColumnStatisticsRequest) SetTableName(v string) *GetPartitionColumnStatisticsRequest {
	s.TableName = &v
	return s
}

type GetPartitionColumnStatisticsShrinkRequest struct {
	CatalogId            *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNamesShrink    *string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty"`
	DatabaseName         *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionNamesShrink *string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty"`
	TableName            *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetPartitionColumnStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionColumnStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetPartitionColumnStatisticsShrinkRequest) SetCatalogId(v string) *GetPartitionColumnStatisticsShrinkRequest {
	s.CatalogId = &v
	return s
}

func (s *GetPartitionColumnStatisticsShrinkRequest) SetColumnNamesShrink(v string) *GetPartitionColumnStatisticsShrinkRequest {
	s.ColumnNamesShrink = &v
	return s
}

func (s *GetPartitionColumnStatisticsShrinkRequest) SetDatabaseName(v string) *GetPartitionColumnStatisticsShrinkRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetPartitionColumnStatisticsShrinkRequest) SetPartitionNamesShrink(v string) *GetPartitionColumnStatisticsShrinkRequest {
	s.PartitionNamesShrink = &v
	return s
}

func (s *GetPartitionColumnStatisticsShrinkRequest) SetTableName(v string) *GetPartitionColumnStatisticsShrinkRequest {
	s.TableName = &v
	return s
}

type GetPartitionColumnStatisticsResponseBody struct {
	Code                   *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Message                *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PartitionStatisticsMap map[string][]*ColumnStatisticsObj `json:"PartitionStatisticsMap,omitempty" xml:"PartitionStatisticsMap,omitempty"`
	RequestId              *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPartitionColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPartitionColumnStatisticsResponseBody) SetCode(v string) *GetPartitionColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPartitionColumnStatisticsResponseBody) SetMessage(v string) *GetPartitionColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPartitionColumnStatisticsResponseBody) SetPartitionStatisticsMap(v map[string][]*ColumnStatisticsObj) *GetPartitionColumnStatisticsResponseBody {
	s.PartitionStatisticsMap = v
	return s
}

func (s *GetPartitionColumnStatisticsResponseBody) SetRequestId(v string) *GetPartitionColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPartitionColumnStatisticsResponseBody) SetSuccess(v bool) *GetPartitionColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetPartitionColumnStatisticsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPartitionColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPartitionColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPartitionColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetPartitionColumnStatisticsResponse) SetHeaders(v map[string]*string) *GetPartitionColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetPartitionColumnStatisticsResponse) SetStatusCode(v int32) *GetPartitionColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPartitionColumnStatisticsResponse) SetBody(v *GetPartitionColumnStatisticsResponseBody) *GetPartitionColumnStatisticsResponse {
	s.Body = v
	return s
}

type GetQueryResultRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryId    *string `json:"QueryId,omitempty" xml:"QueryId,omitempty"`
}

func (s GetQueryResultRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQueryResultRequest) GoString() string {
	return s.String()
}

func (s *GetQueryResultRequest) SetPageNumber(v int32) *GetQueryResultRequest {
	s.PageNumber = &v
	return s
}

func (s *GetQueryResultRequest) SetPageSize(v int32) *GetQueryResultRequest {
	s.PageSize = &v
	return s
}

func (s *GetQueryResultRequest) SetQueryId(v string) *GetQueryResultRequest {
	s.QueryId = &v
	return s
}

type GetQueryResultResponseBody struct {
	Duration            *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EndTime             *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GmtCreate           *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified         *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Id                  *string `json:"Id,omitempty" xml:"Id,omitempty"`
	JobCompleted        *bool   `json:"JobCompleted,omitempty" xml:"JobCompleted,omitempty"`
	Logs                *string `json:"Logs,omitempty" xml:"Logs,omitempty"`
	Owner               *int64  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Progress            *int32  `json:"Progress,omitempty" xml:"Progress,omitempty"`
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId           *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultTmpDb         *string `json:"ResultTmpDb,omitempty" xml:"ResultTmpDb,omitempty"`
	ResultTmpTable      *string `json:"ResultTmpTable,omitempty" xml:"ResultTmpTable,omitempty"`
	RowCount            *int32  `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	RowCountOverLimit   *bool   `json:"RowCountOverLimit,omitempty" xml:"RowCountOverLimit,omitempty"`
	Rows                *string `json:"Rows,omitempty" xml:"Rows,omitempty"`
	Schema              *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Sql                 *string `json:"Sql,omitempty" xml:"Sql,omitempty"`
	StartTime           *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Success             *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalBytesProcessed *int64  `json:"TotalBytesProcessed,omitempty" xml:"TotalBytesProcessed,omitempty"`
}

func (s GetQueryResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQueryResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetQueryResultResponseBody) SetDuration(v int64) *GetQueryResultResponseBody {
	s.Duration = &v
	return s
}

func (s *GetQueryResultResponseBody) SetEndTime(v string) *GetQueryResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *GetQueryResultResponseBody) SetErrorMessage(v string) *GetQueryResultResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetQueryResultResponseBody) SetGmtCreate(v string) *GetQueryResultResponseBody {
	s.GmtCreate = &v
	return s
}

func (s *GetQueryResultResponseBody) SetGmtModified(v string) *GetQueryResultResponseBody {
	s.GmtModified = &v
	return s
}

func (s *GetQueryResultResponseBody) SetId(v string) *GetQueryResultResponseBody {
	s.Id = &v
	return s
}

func (s *GetQueryResultResponseBody) SetJobCompleted(v bool) *GetQueryResultResponseBody {
	s.JobCompleted = &v
	return s
}

func (s *GetQueryResultResponseBody) SetLogs(v string) *GetQueryResultResponseBody {
	s.Logs = &v
	return s
}

func (s *GetQueryResultResponseBody) SetOwner(v int64) *GetQueryResultResponseBody {
	s.Owner = &v
	return s
}

func (s *GetQueryResultResponseBody) SetProgress(v int32) *GetQueryResultResponseBody {
	s.Progress = &v
	return s
}

func (s *GetQueryResultResponseBody) SetRegionId(v string) *GetQueryResultResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetQueryResultResponseBody) SetRequestId(v string) *GetQueryResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQueryResultResponseBody) SetResultTmpDb(v string) *GetQueryResultResponseBody {
	s.ResultTmpDb = &v
	return s
}

func (s *GetQueryResultResponseBody) SetResultTmpTable(v string) *GetQueryResultResponseBody {
	s.ResultTmpTable = &v
	return s
}

func (s *GetQueryResultResponseBody) SetRowCount(v int32) *GetQueryResultResponseBody {
	s.RowCount = &v
	return s
}

func (s *GetQueryResultResponseBody) SetRowCountOverLimit(v bool) *GetQueryResultResponseBody {
	s.RowCountOverLimit = &v
	return s
}

func (s *GetQueryResultResponseBody) SetRows(v string) *GetQueryResultResponseBody {
	s.Rows = &v
	return s
}

func (s *GetQueryResultResponseBody) SetSchema(v string) *GetQueryResultResponseBody {
	s.Schema = &v
	return s
}

func (s *GetQueryResultResponseBody) SetSql(v string) *GetQueryResultResponseBody {
	s.Sql = &v
	return s
}

func (s *GetQueryResultResponseBody) SetStartTime(v string) *GetQueryResultResponseBody {
	s.StartTime = &v
	return s
}

func (s *GetQueryResultResponseBody) SetStatus(v string) *GetQueryResultResponseBody {
	s.Status = &v
	return s
}

func (s *GetQueryResultResponseBody) SetSuccess(v bool) *GetQueryResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetQueryResultResponseBody) SetTotalBytesProcessed(v int64) *GetQueryResultResponseBody {
	s.TotalBytesProcessed = &v
	return s
}

type GetQueryResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetQueryResultResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQueryResultResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQueryResultResponse) GoString() string {
	return s.String()
}

func (s *GetQueryResultResponse) SetHeaders(v map[string]*string) *GetQueryResultResponse {
	s.Headers = v
	return s
}

func (s *GetQueryResultResponse) SetStatusCode(v int32) *GetQueryResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQueryResultResponse) SetBody(v *GetQueryResultResponseBody) *GetQueryResultResponse {
	s.Body = v
	return s
}

type GetRegionStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetRegionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRegionStatusRequest) SetRegionId(v string) *GetRegionStatusRequest {
	s.RegionId = &v
	return s
}

type GetRegionStatusResponseBody struct {
	Data      *GetRegionStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                          `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRegionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponseBody) SetData(v *GetRegionStatusResponseBodyData) *GetRegionStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetRegionStatusResponseBody) SetRequestId(v string) *GetRegionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionStatusResponseBody) SetSuccess(v string) *GetRegionStatusResponseBody {
	s.Success = &v
	return s
}

type GetRegionStatusResponseBodyData struct {
	AccountStatus     *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty"`
	IsDependencyReady *bool   `json:"IsDependencyReady,omitempty" xml:"IsDependencyReady,omitempty"`
	IsDlfServiceOpen  *bool   `json:"IsDlfServiceOpen,omitempty" xml:"IsDlfServiceOpen,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RegionStatus      *string `json:"RegionStatus,omitempty" xml:"RegionStatus,omitempty"`
}

func (s GetRegionStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponseBodyData) SetAccountStatus(v string) *GetRegionStatusResponseBodyData {
	s.AccountStatus = &v
	return s
}

func (s *GetRegionStatusResponseBodyData) SetIsDependencyReady(v bool) *GetRegionStatusResponseBodyData {
	s.IsDependencyReady = &v
	return s
}

func (s *GetRegionStatusResponseBodyData) SetIsDlfServiceOpen(v bool) *GetRegionStatusResponseBodyData {
	s.IsDlfServiceOpen = &v
	return s
}

func (s *GetRegionStatusResponseBodyData) SetRegionId(v string) *GetRegionStatusResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetRegionStatusResponseBodyData) SetRegionStatus(v string) *GetRegionStatusResponseBodyData {
	s.RegionStatus = &v
	return s
}

type GetRegionStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRegionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRegionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRegionStatusResponse) GoString() string {
	return s.String()
}

func (s *GetRegionStatusResponse) SetHeaders(v map[string]*string) *GetRegionStatusResponse {
	s.Headers = v
	return s
}

func (s *GetRegionStatusResponse) SetStatusCode(v int32) *GetRegionStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRegionStatusResponse) SetBody(v *GetRegionStatusResponseBody) *GetRegionStatusResponse {
	s.Body = v
	return s
}

type GetRoleRequest struct {
	// roleName
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRoleRequest) GoString() string {
	return s.String()
}

func (s *GetRoleRequest) SetRoleName(v string) *GetRoleRequest {
	s.RoleName = &v
	return s
}

type GetRoleResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// role
	Role *Role `json:"Role,omitempty" xml:"Role,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponseBody) GoString() string {
	return s.String()
}

func (s *GetRoleResponseBody) SetCode(v string) *GetRoleResponseBody {
	s.Code = &v
	return s
}

func (s *GetRoleResponseBody) SetMessage(v string) *GetRoleResponseBody {
	s.Message = &v
	return s
}

func (s *GetRoleResponseBody) SetRequestId(v string) *GetRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRoleResponseBody) SetRole(v *Role) *GetRoleResponseBody {
	s.Role = v
	return s
}

func (s *GetRoleResponseBody) SetSuccess(v bool) *GetRoleResponseBody {
	s.Success = &v
	return s
}

type GetRoleResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRoleResponse) GoString() string {
	return s.String()
}

func (s *GetRoleResponse) SetHeaders(v map[string]*string) *GetRoleResponse {
	s.Headers = v
	return s
}

func (s *GetRoleResponse) SetStatusCode(v int32) *GetRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoleResponse) SetBody(v *GetRoleResponseBody) *GetRoleResponse {
	s.Body = v
	return s
}

type GetServiceStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetServiceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusRequest) GoString() string {
	return s.String()
}

func (s *GetServiceStatusRequest) SetRegionId(v string) *GetServiceStatusRequest {
	s.RegionId = &v
	return s
}

type GetServiceStatusResponseBody struct {
	Data      *GetServiceStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBody) SetData(v *GetServiceStatusResponseBodyData) *GetServiceStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceStatusResponseBody) SetRequestId(v string) *GetServiceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceStatusResponseBody) SetSuccess(v bool) *GetServiceStatusResponseBody {
	s.Success = &v
	return s
}

type GetServiceStatusResponseBodyData struct {
	HasRamPermissions *bool `json:"HasRamPermissions,omitempty" xml:"HasRamPermissions,omitempty"`
	IsDlfServiceOpen  *bool `json:"IsDlfServiceOpen,omitempty" xml:"IsDlfServiceOpen,omitempty"`
	IsOssOpen         *bool `json:"IsOssOpen,omitempty" xml:"IsOssOpen,omitempty"`
}

func (s GetServiceStatusResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponseBodyData) SetHasRamPermissions(v bool) *GetServiceStatusResponseBodyData {
	s.HasRamPermissions = &v
	return s
}

func (s *GetServiceStatusResponseBodyData) SetIsDlfServiceOpen(v bool) *GetServiceStatusResponseBodyData {
	s.IsDlfServiceOpen = &v
	return s
}

func (s *GetServiceStatusResponseBodyData) SetIsOssOpen(v bool) *GetServiceStatusResponseBodyData {
	s.IsOssOpen = &v
	return s
}

type GetServiceStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetServiceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetServiceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetServiceStatusResponse) GoString() string {
	return s.String()
}

func (s *GetServiceStatusResponse) SetHeaders(v map[string]*string) *GetServiceStatusResponse {
	s.Headers = v
	return s
}

func (s *GetServiceStatusResponse) SetStatusCode(v int32) *GetServiceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceStatusResponse) SetBody(v *GetServiceStatusResponseBody) *GetServiceStatusResponse {
	s.Body = v
	return s
}

type GetTableRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableRequest) GoString() string {
	return s.String()
}

func (s *GetTableRequest) SetCatalogId(v string) *GetTableRequest {
	s.CatalogId = &v
	return s
}

func (s *GetTableRequest) SetDatabaseName(v string) *GetTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetTableRequest) SetTableName(v string) *GetTableRequest {
	s.TableName = &v
	return s
}

type GetTableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	Table     *Table  `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s GetTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableResponseBody) SetCode(v string) *GetTableResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableResponseBody) SetMessage(v string) *GetTableResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableResponseBody) SetRequestId(v string) *GetTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableResponseBody) SetSuccess(v bool) *GetTableResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableResponseBody) SetTable(v *Table) *GetTableResponseBody {
	s.Table = v
	return s
}

type GetTableResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableResponse) GoString() string {
	return s.String()
}

func (s *GetTableResponse) SetHeaders(v map[string]*string) *GetTableResponse {
	s.Headers = v
	return s
}

func (s *GetTableResponse) SetStatusCode(v int32) *GetTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableResponse) SetBody(v *GetTableResponseBody) *GetTableResponse {
	s.Body = v
	return s
}

type GetTableColumnStatisticsRequest struct {
	CatalogId    *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNames  []*string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty" type:"Repeated"`
	DatabaseName *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnStatisticsRequest) SetCatalogId(v string) *GetTableColumnStatisticsRequest {
	s.CatalogId = &v
	return s
}

func (s *GetTableColumnStatisticsRequest) SetColumnNames(v []*string) *GetTableColumnStatisticsRequest {
	s.ColumnNames = v
	return s
}

func (s *GetTableColumnStatisticsRequest) SetDatabaseName(v string) *GetTableColumnStatisticsRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetTableColumnStatisticsRequest) SetTableName(v string) *GetTableColumnStatisticsRequest {
	s.TableName = &v
	return s
}

type GetTableColumnStatisticsShrinkRequest struct {
	CatalogId         *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	ColumnNamesShrink *string `json:"ColumnNames,omitempty" xml:"ColumnNames,omitempty"`
	DatabaseName      *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableColumnStatisticsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnStatisticsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetTableColumnStatisticsShrinkRequest) SetCatalogId(v string) *GetTableColumnStatisticsShrinkRequest {
	s.CatalogId = &v
	return s
}

func (s *GetTableColumnStatisticsShrinkRequest) SetColumnNamesShrink(v string) *GetTableColumnStatisticsShrinkRequest {
	s.ColumnNamesShrink = &v
	return s
}

func (s *GetTableColumnStatisticsShrinkRequest) SetDatabaseName(v string) *GetTableColumnStatisticsShrinkRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetTableColumnStatisticsShrinkRequest) SetTableName(v string) *GetTableColumnStatisticsShrinkRequest {
	s.TableName = &v
	return s
}

type GetTableColumnStatisticsResponseBody struct {
	Code                    *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	ColumnStatisticsObjList []*ColumnStatisticsObj `json:"ColumnStatisticsObjList,omitempty" xml:"ColumnStatisticsObjList,omitempty" type:"Repeated"`
	Message                 *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTableColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableColumnStatisticsResponseBody) SetCode(v string) *GetTableColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableColumnStatisticsResponseBody) SetColumnStatisticsObjList(v []*ColumnStatisticsObj) *GetTableColumnStatisticsResponseBody {
	s.ColumnStatisticsObjList = v
	return s
}

func (s *GetTableColumnStatisticsResponseBody) SetMessage(v string) *GetTableColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableColumnStatisticsResponseBody) SetRequestId(v string) *GetTableColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableColumnStatisticsResponseBody) SetSuccess(v bool) *GetTableColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type GetTableColumnStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnStatisticsResponse) SetHeaders(v map[string]*string) *GetTableColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnStatisticsResponse) SetStatusCode(v int32) *GetTableColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnStatisticsResponse) SetBody(v *GetTableColumnStatisticsResponseBody) *GetTableColumnStatisticsResponse {
	s.Body = v
	return s
}

type GetTableProfileRequest struct {
	// CatalogId
	CatalogId *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	// DatabaseName
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// TableName
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetTableProfileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableProfileRequest) GoString() string {
	return s.String()
}

func (s *GetTableProfileRequest) SetCatalogId(v string) *GetTableProfileRequest {
	s.CatalogId = &v
	return s
}

func (s *GetTableProfileRequest) SetDatabaseName(v string) *GetTableProfileRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetTableProfileRequest) SetTableName(v string) *GetTableProfileRequest {
	s.TableName = &v
	return s
}

type GetTableProfileResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message      *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableProfile *TableProfile `json:"TableProfile,omitempty" xml:"TableProfile,omitempty"`
}

func (s GetTableProfileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableProfileResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableProfileResponseBody) SetCode(v string) *GetTableProfileResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableProfileResponseBody) SetMessage(v string) *GetTableProfileResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableProfileResponseBody) SetRequestId(v string) *GetTableProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableProfileResponseBody) SetSuccess(v bool) *GetTableProfileResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableProfileResponseBody) SetTableProfile(v *TableProfile) *GetTableProfileResponseBody {
	s.TableProfile = v
	return s
}

type GetTableProfileResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableProfileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableProfileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableProfileResponse) GoString() string {
	return s.String()
}

func (s *GetTableProfileResponse) SetHeaders(v map[string]*string) *GetTableProfileResponse {
	s.Headers = v
	return s
}

func (s *GetTableProfileResponse) SetStatusCode(v int32) *GetTableProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableProfileResponse) SetBody(v *GetTableProfileResponseBody) *GetTableProfileResponse {
	s.Body = v
	return s
}

type GetTableVersionRequest struct {
	CatalogId    *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableName    *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	VersionId    *int32  `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetTableVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTableVersionRequest) GoString() string {
	return s.String()
}

func (s *GetTableVersionRequest) SetCatalogId(v string) *GetTableVersionRequest {
	s.CatalogId = &v
	return s
}

func (s *GetTableVersionRequest) SetDatabaseName(v string) *GetTableVersionRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetTableVersionRequest) SetTableName(v string) *GetTableVersionRequest {
	s.TableName = &v
	return s
}

func (s *GetTableVersionRequest) SetVersionId(v int32) *GetTableVersionRequest {
	s.VersionId = &v
	return s
}

type GetTableVersionResponseBody struct {
	Code         *string       `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success      *bool         `json:"Success,omitempty" xml:"Success,omitempty"`
	TableVersion *TableVersion `json:"TableVersion,omitempty" xml:"TableVersion,omitempty"`
}

func (s GetTableVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTableVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetTableVersionResponseBody) SetCode(v string) *GetTableVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetTableVersionResponseBody) SetMessage(v string) *GetTableVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetTableVersionResponseBody) SetRequestId(v string) *GetTableVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTableVersionResponseBody) SetSuccess(v bool) *GetTableVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetTableVersionResponseBody) SetTableVersion(v *TableVersion) *GetTableVersionResponseBody {
	s.TableVersion = v
	return s
}

type GetTableVersionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetTableVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTableVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTableVersionResponse) GoString() string {
	return s.String()
}

func (s *GetTableVersionResponse) SetHeaders(v map[string]*string) *GetTableVersionResponse {
	s.Headers = v
	return s
}

func (s *GetTableVersionResponse) SetStatusCode(v int32) *GetTableVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableVersionResponse) SetBody(v *GetTableVersionResponseBody) *GetTableVersionResponse {
	s.Body = v
	return s
}

type GrantPermissionsRequest struct {
	Accesses []*string `json:"Accesses,omitempty" xml:"Accesses,omitempty" type:"Repeated"`
	// CatalogId
	CatalogId        *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DelegateAccesses []*string     `json:"DelegateAccesses,omitempty" xml:"DelegateAccesses,omitempty" type:"Repeated"`
	MetaResource     *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	Principal        *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
	Type             *string       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GrantPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsRequest) GoString() string {
	return s.String()
}

func (s *GrantPermissionsRequest) SetAccesses(v []*string) *GrantPermissionsRequest {
	s.Accesses = v
	return s
}

func (s *GrantPermissionsRequest) SetCatalogId(v string) *GrantPermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *GrantPermissionsRequest) SetDelegateAccesses(v []*string) *GrantPermissionsRequest {
	s.DelegateAccesses = v
	return s
}

func (s *GrantPermissionsRequest) SetMetaResource(v *MetaResource) *GrantPermissionsRequest {
	s.MetaResource = v
	return s
}

func (s *GrantPermissionsRequest) SetPrincipal(v *Principal) *GrantPermissionsRequest {
	s.Principal = v
	return s
}

func (s *GrantPermissionsRequest) SetType(v string) *GrantPermissionsRequest {
	s.Type = &v
	return s
}

type GrantPermissionsResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *GrantPermissionsResponseBody) SetCode(v string) *GrantPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *GrantPermissionsResponseBody) SetMessage(v string) *GrantPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *GrantPermissionsResponseBody) SetRequestId(v string) *GrantPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantPermissionsResponseBody) SetSuccess(v bool) *GrantPermissionsResponseBody {
	s.Success = &v
	return s
}

type GrantPermissionsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantPermissionsResponse) GoString() string {
	return s.String()
}

func (s *GrantPermissionsResponse) SetHeaders(v map[string]*string) *GrantPermissionsResponse {
	s.Headers = v
	return s
}

func (s *GrantPermissionsResponse) SetStatusCode(v int32) *GrantPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantPermissionsResponse) SetBody(v *GrantPermissionsResponseBody) *GrantPermissionsResponse {
	s.Body = v
	return s
}

type GrantRoleToUsersRequest struct {
	// RoleName
	RoleName *string      `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Users    []*Principal `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s GrantRoleToUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersRequest) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersRequest) SetRoleName(v string) *GrantRoleToUsersRequest {
	s.RoleName = &v
	return s
}

func (s *GrantRoleToUsersRequest) SetUsers(v []*Principal) *GrantRoleToUsersRequest {
	s.Users = v
	return s
}

type GrantRoleToUsersResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantRoleToUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponseBody) SetCode(v string) *GrantRoleToUsersResponseBody {
	s.Code = &v
	return s
}

func (s *GrantRoleToUsersResponseBody) SetMessage(v string) *GrantRoleToUsersResponseBody {
	s.Message = &v
	return s
}

func (s *GrantRoleToUsersResponseBody) SetRequestId(v string) *GrantRoleToUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantRoleToUsersResponseBody) SetSuccess(v bool) *GrantRoleToUsersResponseBody {
	s.Success = &v
	return s
}

type GrantRoleToUsersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantRoleToUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantRoleToUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRoleToUsersResponse) GoString() string {
	return s.String()
}

func (s *GrantRoleToUsersResponse) SetHeaders(v map[string]*string) *GrantRoleToUsersResponse {
	s.Headers = v
	return s
}

func (s *GrantRoleToUsersResponse) SetStatusCode(v int32) *GrantRoleToUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRoleToUsersResponse) SetBody(v *GrantRoleToUsersResponseBody) *GrantRoleToUsersResponse {
	s.Body = v
	return s
}

type GrantRolesToUserRequest struct {
	RoleNames []*string  `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
	User      *Principal `json:"User,omitempty" xml:"User,omitempty"`
}

func (s GrantRolesToUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesToUserRequest) GoString() string {
	return s.String()
}

func (s *GrantRolesToUserRequest) SetRoleNames(v []*string) *GrantRolesToUserRequest {
	s.RoleNames = v
	return s
}

func (s *GrantRolesToUserRequest) SetUser(v *Principal) *GrantRolesToUserRequest {
	s.User = v
	return s
}

type GrantRolesToUserResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantRolesToUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesToUserResponseBody) GoString() string {
	return s.String()
}

func (s *GrantRolesToUserResponseBody) SetCode(v string) *GrantRolesToUserResponseBody {
	s.Code = &v
	return s
}

func (s *GrantRolesToUserResponseBody) SetMessage(v string) *GrantRolesToUserResponseBody {
	s.Message = &v
	return s
}

func (s *GrantRolesToUserResponseBody) SetRequestId(v string) *GrantRolesToUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantRolesToUserResponseBody) SetSuccess(v bool) *GrantRolesToUserResponseBody {
	s.Success = &v
	return s
}

type GrantRolesToUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GrantRolesToUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GrantRolesToUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GrantRolesToUserResponse) GoString() string {
	return s.String()
}

func (s *GrantRolesToUserResponse) SetHeaders(v map[string]*string) *GrantRolesToUserResponse {
	s.Headers = v
	return s
}

func (s *GrantRolesToUserResponse) SetStatusCode(v int32) *GrantRolesToUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantRolesToUserResponse) SetBody(v *GrantRolesToUserResponseBody) *GrantRolesToUserResponse {
	s.Body = v
	return s
}

type ListCatalogsRequest struct {
	IdPattern     *string `json:"IdPattern,omitempty" xml:"IdPattern,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCatalogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) SetIdPattern(v string) *ListCatalogsRequest {
	s.IdPattern = &v
	return s
}

func (s *ListCatalogsRequest) SetNextPageToken(v string) *ListCatalogsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListCatalogsRequest) SetPageSize(v int32) *ListCatalogsRequest {
	s.PageSize = &v
	return s
}

type ListCatalogsResponseBody struct {
	Catalogs      []*Catalog `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	Code          *string    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string    `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string    `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool      `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCatalogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponseBody) SetCatalogs(v []*Catalog) *ListCatalogsResponseBody {
	s.Catalogs = v
	return s
}

func (s *ListCatalogsResponseBody) SetCode(v string) *ListCatalogsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCatalogsResponseBody) SetMessage(v string) *ListCatalogsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCatalogsResponseBody) SetNextPageToken(v string) *ListCatalogsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListCatalogsResponseBody) SetRequestId(v string) *ListCatalogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCatalogsResponseBody) SetSuccess(v bool) *ListCatalogsResponseBody {
	s.Success = &v
	return s
}

type ListCatalogsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCatalogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCatalogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCatalogsResponse) GoString() string {
	return s.String()
}

func (s *ListCatalogsResponse) SetHeaders(v map[string]*string) *ListCatalogsResponse {
	s.Headers = v
	return s
}

func (s *ListCatalogsResponse) SetStatusCode(v int32) *ListCatalogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCatalogsResponse) SetBody(v *ListCatalogsResponseBody) *ListCatalogsResponse {
	s.Body = v
	return s
}

type ListDatabasesRequest struct {
	CatalogId     *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	NamePattern   *string `json:"NamePattern,omitempty" xml:"NamePattern,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesRequest) GoString() string {
	return s.String()
}

func (s *ListDatabasesRequest) SetCatalogId(v string) *ListDatabasesRequest {
	s.CatalogId = &v
	return s
}

func (s *ListDatabasesRequest) SetNamePattern(v string) *ListDatabasesRequest {
	s.NamePattern = &v
	return s
}

func (s *ListDatabasesRequest) SetNextPageToken(v string) *ListDatabasesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListDatabasesRequest) SetPageSize(v int32) *ListDatabasesRequest {
	s.PageSize = &v
	return s
}

type ListDatabasesResponseBody struct {
	Code          *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Databases     []*Database `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Message       *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string     `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponseBody) SetCode(v string) *ListDatabasesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDatabasesResponseBody) SetDatabases(v []*Database) *ListDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *ListDatabasesResponseBody) SetMessage(v string) *ListDatabasesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDatabasesResponseBody) SetNextPageToken(v string) *ListDatabasesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListDatabasesResponseBody) SetRequestId(v string) *ListDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabasesResponseBody) SetSuccess(v bool) *ListDatabasesResponseBody {
	s.Success = &v
	return s
}

type ListDatabasesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListDatabasesResponse) SetHeaders(v map[string]*string) *ListDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListDatabasesResponse) SetStatusCode(v int32) *ListDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabasesResponse) SetBody(v *ListDatabasesResponseBody) *ListDatabasesResponse {
	s.Body = v
	return s
}

type ListFunctionNamesRequest struct {
	CatalogId           *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionNamePattern *string `json:"FunctionNamePattern,omitempty" xml:"FunctionNamePattern,omitempty"`
	NextPageToken       *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionNamesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionNamesRequest) SetCatalogId(v string) *ListFunctionNamesRequest {
	s.CatalogId = &v
	return s
}

func (s *ListFunctionNamesRequest) SetDatabaseName(v string) *ListFunctionNamesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionNamesRequest) SetFunctionNamePattern(v string) *ListFunctionNamesRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListFunctionNamesRequest) SetNextPageToken(v string) *ListFunctionNamesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListFunctionNamesRequest) SetPageSize(v int32) *ListFunctionNamesRequest {
	s.PageSize = &v
	return s
}

type ListFunctionNamesResponseBody struct {
	Code          *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	FunctionNames []*string `json:"FunctionNames,omitempty" xml:"FunctionNames,omitempty" type:"Repeated"`
	Message       *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string   `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFunctionNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionNamesResponseBody) SetCode(v string) *ListFunctionNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionNamesResponseBody) SetFunctionNames(v []*string) *ListFunctionNamesResponseBody {
	s.FunctionNames = v
	return s
}

func (s *ListFunctionNamesResponseBody) SetMessage(v string) *ListFunctionNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionNamesResponseBody) SetNextPageToken(v string) *ListFunctionNamesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListFunctionNamesResponseBody) SetRequestId(v string) *ListFunctionNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionNamesResponseBody) SetSuccess(v bool) *ListFunctionNamesResponseBody {
	s.Success = &v
	return s
}

type ListFunctionNamesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionNamesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionNamesResponse) SetHeaders(v map[string]*string) *ListFunctionNamesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionNamesResponse) SetStatusCode(v int32) *ListFunctionNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionNamesResponse) SetBody(v *ListFunctionNamesResponseBody) *ListFunctionNamesResponse {
	s.Body = v
	return s
}

type ListFunctionsRequest struct {
	CatalogId           *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName        *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionNamePattern *string `json:"FunctionNamePattern,omitempty" xml:"FunctionNamePattern,omitempty"`
	NextPageToken       *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListFunctionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionsRequest) SetCatalogId(v string) *ListFunctionsRequest {
	s.CatalogId = &v
	return s
}

func (s *ListFunctionsRequest) SetDatabaseName(v string) *ListFunctionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListFunctionsRequest) SetFunctionNamePattern(v string) *ListFunctionsRequest {
	s.FunctionNamePattern = &v
	return s
}

func (s *ListFunctionsRequest) SetNextPageToken(v string) *ListFunctionsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListFunctionsRequest) SetPageSize(v int32) *ListFunctionsRequest {
	s.PageSize = &v
	return s
}

type ListFunctionsResponseBody struct {
	Code          *string     `json:"Code,omitempty" xml:"Code,omitempty"`
	Functions     []*Function `json:"Functions,omitempty" xml:"Functions,omitempty" type:"Repeated"`
	Message       *string     `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string     `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) SetCode(v string) *ListFunctionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionsResponseBody) SetFunctions(v []*Function) *ListFunctionsResponseBody {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBody) SetMessage(v string) *ListFunctionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionsResponseBody) SetNextPageToken(v string) *ListFunctionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListFunctionsResponseBody) SetRequestId(v string) *ListFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionsResponseBody) SetSuccess(v bool) *ListFunctionsResponseBody {
	s.Success = &v
	return s
}

type ListFunctionsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionsResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponse) SetHeaders(v map[string]*string) *ListFunctionsResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionsResponse) SetStatusCode(v int32) *ListFunctionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionsResponse) SetBody(v *ListFunctionsResponseBody) *ListFunctionsResponse {
	s.Body = v
	return s
}

type ListPartitionNamesRequest struct {
	CatalogId         *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName      *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	NextPageToken     *string   `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PartialPartValues []*string `json:"PartialPartValues,omitempty" xml:"PartialPartValues,omitempty" type:"Repeated"`
	TableName         *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListPartitionNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionNamesRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionNamesRequest) SetCatalogId(v string) *ListPartitionNamesRequest {
	s.CatalogId = &v
	return s
}

func (s *ListPartitionNamesRequest) SetDatabaseName(v string) *ListPartitionNamesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListPartitionNamesRequest) SetNextPageToken(v string) *ListPartitionNamesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionNamesRequest) SetPageSize(v int32) *ListPartitionNamesRequest {
	s.PageSize = &v
	return s
}

func (s *ListPartitionNamesRequest) SetPartialPartValues(v []*string) *ListPartitionNamesRequest {
	s.PartialPartValues = v
	return s
}

func (s *ListPartitionNamesRequest) SetTableName(v string) *ListPartitionNamesRequest {
	s.TableName = &v
	return s
}

type ListPartitionNamesResponseBody struct {
	Code           *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken  *string   `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PartitionNames []*string `json:"PartitionNames,omitempty" xml:"PartitionNames,omitempty" type:"Repeated"`
	RequestId      *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPartitionNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionNamesResponseBody) SetCode(v string) *ListPartitionNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPartitionNamesResponseBody) SetMessage(v string) *ListPartitionNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPartitionNamesResponseBody) SetNextPageToken(v string) *ListPartitionNamesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionNamesResponseBody) SetPartitionNames(v []*string) *ListPartitionNamesResponseBody {
	s.PartitionNames = v
	return s
}

func (s *ListPartitionNamesResponseBody) SetRequestId(v string) *ListPartitionNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPartitionNamesResponseBody) SetSuccess(v bool) *ListPartitionNamesResponseBody {
	s.Success = &v
	return s
}

type ListPartitionNamesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPartitionNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPartitionNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionNamesResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionNamesResponse) SetHeaders(v map[string]*string) *ListPartitionNamesResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionNamesResponse) SetStatusCode(v int32) *ListPartitionNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionNamesResponse) SetBody(v *ListPartitionNamesResponseBody) *ListPartitionNamesResponse {
	s.Body = v
	return s
}

type ListPartitionsRequest struct {
	CatalogId         *string   `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName      *string   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IsShareSd         *bool     `json:"IsShareSd,omitempty" xml:"IsShareSd,omitempty"`
	NextPageToken     *string   `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize          *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PartialPartValues []*string `json:"PartialPartValues,omitempty" xml:"PartialPartValues,omitempty" type:"Repeated"`
	TableName         *string   `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListPartitionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionsRequest) SetCatalogId(v string) *ListPartitionsRequest {
	s.CatalogId = &v
	return s
}

func (s *ListPartitionsRequest) SetDatabaseName(v string) *ListPartitionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListPartitionsRequest) SetIsShareSd(v bool) *ListPartitionsRequest {
	s.IsShareSd = &v
	return s
}

func (s *ListPartitionsRequest) SetNextPageToken(v string) *ListPartitionsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionsRequest) SetPageSize(v int32) *ListPartitionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPartitionsRequest) SetPartialPartValues(v []*string) *ListPartitionsRequest {
	s.PartialPartValues = v
	return s
}

func (s *ListPartitionsRequest) SetTableName(v string) *ListPartitionsRequest {
	s.TableName = &v
	return s
}

type ListPartitionsResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// PartitionSpecs
	PartitionSpecs []*PartitionSpec `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
	Partitions     []*Partition     `json:"Partitions,omitempty" xml:"Partitions,omitempty" type:"Repeated"`
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPartitionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponseBody) SetCode(v string) *ListPartitionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPartitionsResponseBody) SetMessage(v string) *ListPartitionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPartitionsResponseBody) SetNextPageToken(v string) *ListPartitionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionsResponseBody) SetPartitionSpecs(v []*PartitionSpec) *ListPartitionsResponseBody {
	s.PartitionSpecs = v
	return s
}

func (s *ListPartitionsResponseBody) SetPartitions(v []*Partition) *ListPartitionsResponseBody {
	s.Partitions = v
	return s
}

func (s *ListPartitionsResponseBody) SetRequestId(v string) *ListPartitionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPartitionsResponseBody) SetSuccess(v bool) *ListPartitionsResponseBody {
	s.Success = &v
	return s
}

type ListPartitionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPartitionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPartitionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionsResponse) SetHeaders(v map[string]*string) *ListPartitionsResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionsResponse) SetStatusCode(v int32) *ListPartitionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionsResponse) SetBody(v *ListPartitionsResponseBody) *ListPartitionsResponse {
	s.Body = v
	return s
}

type ListPartitionsByExprResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ListPartitionsByExprResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsByExprResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionsByExprResponse) SetHeaders(v map[string]*string) *ListPartitionsByExprResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionsByExprResponse) SetStatusCode(v int32) *ListPartitionsByExprResponse {
	s.StatusCode = &v
	return s
}

type ListPartitionsByFilterRequest struct {
	CatalogId     *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName  *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Filter        *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	IsShareSd     *bool   `json:"IsShareSd,omitempty" xml:"IsShareSd,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListPartitionsByFilterRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsByFilterRequest) GoString() string {
	return s.String()
}

func (s *ListPartitionsByFilterRequest) SetCatalogId(v string) *ListPartitionsByFilterRequest {
	s.CatalogId = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetDatabaseName(v string) *ListPartitionsByFilterRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetFilter(v string) *ListPartitionsByFilterRequest {
	s.Filter = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetIsShareSd(v bool) *ListPartitionsByFilterRequest {
	s.IsShareSd = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetNextPageToken(v string) *ListPartitionsByFilterRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetPageSize(v int32) *ListPartitionsByFilterRequest {
	s.PageSize = &v
	return s
}

func (s *ListPartitionsByFilterRequest) SetTableName(v string) *ListPartitionsByFilterRequest {
	s.TableName = &v
	return s
}

type ListPartitionsByFilterResponseBody struct {
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// PartitionSpecs
	PartitionSpecs []*PartitionSpec `json:"PartitionSpecs,omitempty" xml:"PartitionSpecs,omitempty" type:"Repeated"`
	Partitions     []*Partition     `json:"Partitions,omitempty" xml:"Partitions,omitempty" type:"Repeated"`
	RequestId      *string          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListPartitionsByFilterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsByFilterResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartitionsByFilterResponseBody) SetCode(v string) *ListPartitionsByFilterResponseBody {
	s.Code = &v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetMessage(v string) *ListPartitionsByFilterResponseBody {
	s.Message = &v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetNextPageToken(v string) *ListPartitionsByFilterResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetPartitionSpecs(v []*PartitionSpec) *ListPartitionsByFilterResponseBody {
	s.PartitionSpecs = v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetPartitions(v []*Partition) *ListPartitionsByFilterResponseBody {
	s.Partitions = v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetRequestId(v string) *ListPartitionsByFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPartitionsByFilterResponseBody) SetSuccess(v bool) *ListPartitionsByFilterResponseBody {
	s.Success = &v
	return s
}

type ListPartitionsByFilterResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPartitionsByFilterResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPartitionsByFilterResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartitionsByFilterResponse) GoString() string {
	return s.String()
}

func (s *ListPartitionsByFilterResponse) SetHeaders(v map[string]*string) *ListPartitionsByFilterResponse {
	s.Headers = v
	return s
}

func (s *ListPartitionsByFilterResponse) SetStatusCode(v int32) *ListPartitionsByFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartitionsByFilterResponse) SetBody(v *ListPartitionsByFilterResponseBody) *ListPartitionsByFilterResponse {
	s.Body = v
	return s
}

type ListPermissionsRequest struct {
	// CatalogId
	CatalogId                 *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	IsListUserRolePermissions *bool         `json:"IsListUserRolePermissions,omitempty" xml:"IsListUserRolePermissions,omitempty"`
	MetaResource              *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	MetaResourceType          *string       `json:"MetaResourceType,omitempty" xml:"MetaResourceType,omitempty"`
	NextPageToken             *string       `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize                  *int32        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Principal                 *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
	Type                      *string       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListPermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsRequest) GoString() string {
	return s.String()
}

func (s *ListPermissionsRequest) SetCatalogId(v string) *ListPermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *ListPermissionsRequest) SetIsListUserRolePermissions(v bool) *ListPermissionsRequest {
	s.IsListUserRolePermissions = &v
	return s
}

func (s *ListPermissionsRequest) SetMetaResource(v *MetaResource) *ListPermissionsRequest {
	s.MetaResource = v
	return s
}

func (s *ListPermissionsRequest) SetMetaResourceType(v string) *ListPermissionsRequest {
	s.MetaResourceType = &v
	return s
}

func (s *ListPermissionsRequest) SetNextPageToken(v string) *ListPermissionsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListPermissionsRequest) SetPageSize(v int32) *ListPermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListPermissionsRequest) SetPrincipal(v *Principal) *ListPermissionsRequest {
	s.Principal = v
	return s
}

func (s *ListPermissionsRequest) SetType(v string) *ListPermissionsRequest {
	s.Type = &v
	return s
}

type ListPermissionsResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// NextPageToken
	NextPageToken                    *string                         `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PrincipalResourcePermissionsList []*PrincipalResourcePermissions `json:"PrincipalResourcePermissionsList,omitempty" xml:"PrincipalResourcePermissionsList,omitempty" type:"Repeated"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// TotalCount
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponseBody) SetCode(v string) *ListPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPermissionsResponseBody) SetMessage(v string) *ListPermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPermissionsResponseBody) SetNextPageToken(v string) *ListPermissionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListPermissionsResponseBody) SetPrincipalResourcePermissionsList(v []*PrincipalResourcePermissions) *ListPermissionsResponseBody {
	s.PrincipalResourcePermissionsList = v
	return s
}

func (s *ListPermissionsResponseBody) SetRequestId(v string) *ListPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPermissionsResponseBody) SetSuccess(v bool) *ListPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListPermissionsResponseBody) SetTotalCount(v int64) *ListPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListPermissionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPermissionsResponse) GoString() string {
	return s.String()
}

func (s *ListPermissionsResponse) SetHeaders(v map[string]*string) *ListPermissionsResponse {
	s.Headers = v
	return s
}

func (s *ListPermissionsResponse) SetStatusCode(v int32) *ListPermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPermissionsResponse) SetBody(v *ListPermissionsResponseBody) *ListPermissionsResponse {
	s.Body = v
	return s
}

type ListRoleUsersRequest struct {
	// NextPageToken
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// PageSize
	PageSize *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	// use name pattern filter
	UserNamePattern *string `json:"UserNamePattern,omitempty" xml:"UserNamePattern,omitempty"`
}

func (s ListRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRoleUsersRequest) SetNextPageToken(v string) *ListRoleUsersRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListRoleUsersRequest) SetPageSize(v int32) *ListRoleUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListRoleUsersRequest) SetRoleName(v string) *ListRoleUsersRequest {
	s.RoleName = &v
	return s
}

func (s *ListRoleUsersRequest) SetUserNamePattern(v string) *ListRoleUsersRequest {
	s.UserNamePattern = &v
	return s
}

type ListRoleUsersResponseBody struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// NextPageToken
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// user roles
	UserRoles []*UserRole `json:"UserRoles,omitempty" xml:"UserRoles,omitempty" type:"Repeated"`
}

func (s ListRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponseBody) SetCode(v string) *ListRoleUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetMessage(v string) *ListRoleUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetNextPageToken(v string) *ListRoleUsersResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetRequestId(v string) *ListRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetSuccess(v bool) *ListRoleUsersResponseBody {
	s.Success = &v
	return s
}

func (s *ListRoleUsersResponseBody) SetUserRoles(v []*UserRole) *ListRoleUsersResponseBody {
	s.UserRoles = v
	return s
}

type ListRoleUsersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *ListRoleUsersResponse) SetHeaders(v map[string]*string) *ListRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *ListRoleUsersResponse) SetStatusCode(v int32) *ListRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoleUsersResponse) SetBody(v *ListRoleUsersResponseBody) *ListRoleUsersResponse {
	s.Body = v
	return s
}

type ListRolesRequest struct {
	// Next PageToken
	NextPageToken   *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RoleNamePattern *string `json:"RoleNamePattern,omitempty" xml:"RoleNamePattern,omitempty"`
}

func (s ListRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRolesRequest) GoString() string {
	return s.String()
}

func (s *ListRolesRequest) SetNextPageToken(v string) *ListRolesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListRolesRequest) SetPageSize(v int32) *ListRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListRolesRequest) SetRoleNamePattern(v string) *ListRolesRequest {
	s.RoleNamePattern = &v
	return s
}

type ListRolesResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// data
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// role list data
	Roles []*Role `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRolesResponseBody) SetCode(v string) *ListRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListRolesResponseBody) SetMessage(v string) *ListRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListRolesResponseBody) SetNextPageToken(v string) *ListRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRolesResponseBody) SetRequestId(v string) *ListRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRolesResponseBody) SetRoles(v []*Role) *ListRolesResponseBody {
	s.Roles = v
	return s
}

func (s *ListRolesResponseBody) SetSuccess(v bool) *ListRolesResponseBody {
	s.Success = &v
	return s
}

type ListRolesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRolesResponse) GoString() string {
	return s.String()
}

func (s *ListRolesResponse) SetHeaders(v map[string]*string) *ListRolesResponse {
	s.Headers = v
	return s
}

func (s *ListRolesResponse) SetStatusCode(v int32) *ListRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRolesResponse) SetBody(v *ListRolesResponseBody) *ListRolesResponse {
	s.Body = v
	return s
}

type ListTableNamesRequest struct {
	CatalogId        *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName     *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	NextPageToken    *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TableNamePattern *string `json:"TableNamePattern,omitempty" xml:"TableNamePattern,omitempty"`
	TableType        *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListTableNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTableNamesRequest) GoString() string {
	return s.String()
}

func (s *ListTableNamesRequest) SetCatalogId(v string) *ListTableNamesRequest {
	s.CatalogId = &v
	return s
}

func (s *ListTableNamesRequest) SetDatabaseName(v string) *ListTableNamesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListTableNamesRequest) SetNextPageToken(v string) *ListTableNamesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListTableNamesRequest) SetPageSize(v int32) *ListTableNamesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableNamesRequest) SetTableNamePattern(v string) *ListTableNamesRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListTableNamesRequest) SetTableType(v string) *ListTableNamesRequest {
	s.TableType = &v
	return s
}

type ListTableNamesResponseBody struct {
	Code          *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string   `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
	TableNames    []*string `json:"TableNames,omitempty" xml:"TableNames,omitempty" type:"Repeated"`
}

func (s ListTableNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTableNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableNamesResponseBody) SetCode(v string) *ListTableNamesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTableNamesResponseBody) SetMessage(v string) *ListTableNamesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTableNamesResponseBody) SetNextPageToken(v string) *ListTableNamesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTableNamesResponseBody) SetRequestId(v string) *ListTableNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableNamesResponseBody) SetSuccess(v bool) *ListTableNamesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTableNamesResponseBody) SetTableNames(v []*string) *ListTableNamesResponseBody {
	s.TableNames = v
	return s
}

type ListTableNamesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTableNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTableNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTableNamesResponse) GoString() string {
	return s.String()
}

func (s *ListTableNamesResponse) SetHeaders(v map[string]*string) *ListTableNamesResponse {
	s.Headers = v
	return s
}

func (s *ListTableNamesResponse) SetStatusCode(v int32) *ListTableNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableNamesResponse) SetBody(v *ListTableNamesResponseBody) *ListTableNamesResponse {
	s.Body = v
	return s
}

type ListTableVersionsRequest struct {
	CatalogId     *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName  *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TableName     *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s ListTableVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTableVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListTableVersionsRequest) SetCatalogId(v string) *ListTableVersionsRequest {
	s.CatalogId = &v
	return s
}

func (s *ListTableVersionsRequest) SetDatabaseName(v string) *ListTableVersionsRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListTableVersionsRequest) SetNextPageToken(v string) *ListTableVersionsRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListTableVersionsRequest) SetPageSize(v int32) *ListTableVersionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTableVersionsRequest) SetTableName(v string) *ListTableVersionsRequest {
	s.TableName = &v
	return s
}

type ListTableVersionsResponseBody struct {
	Code          *string         `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string         `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string         `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool           `json:"Success,omitempty" xml:"Success,omitempty"`
	TableVersions []*TableVersion `json:"TableVersions,omitempty" xml:"TableVersions,omitempty" type:"Repeated"`
}

func (s ListTableVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTableVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableVersionsResponseBody) SetCode(v string) *ListTableVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTableVersionsResponseBody) SetMessage(v string) *ListTableVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTableVersionsResponseBody) SetNextPageToken(v string) *ListTableVersionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTableVersionsResponseBody) SetRequestId(v string) *ListTableVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableVersionsResponseBody) SetSuccess(v bool) *ListTableVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListTableVersionsResponseBody) SetTableVersions(v []*TableVersion) *ListTableVersionsResponseBody {
	s.TableVersions = v
	return s
}

type ListTableVersionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTableVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTableVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTableVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListTableVersionsResponse) SetHeaders(v map[string]*string) *ListTableVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListTableVersionsResponse) SetStatusCode(v int32) *ListTableVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTableVersionsResponse) SetBody(v *ListTableVersionsResponseBody) *ListTableVersionsResponse {
	s.Body = v
	return s
}

type ListTablesRequest struct {
	CatalogId        *string `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName     *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	NextPageToken    *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TableNamePattern *string `json:"TableNamePattern,omitempty" xml:"TableNamePattern,omitempty"`
	TableType        *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ListTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTablesRequest) GoString() string {
	return s.String()
}

func (s *ListTablesRequest) SetCatalogId(v string) *ListTablesRequest {
	s.CatalogId = &v
	return s
}

func (s *ListTablesRequest) SetDatabaseName(v string) *ListTablesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ListTablesRequest) SetNextPageToken(v string) *ListTablesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListTablesRequest) SetPageSize(v int32) *ListTablesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTablesRequest) SetTableNamePattern(v string) *ListTablesRequest {
	s.TableNamePattern = &v
	return s
}

func (s *ListTablesRequest) SetTableType(v string) *ListTablesRequest {
	s.TableType = &v
	return s
}

type ListTablesResponseBody struct {
	Code          *string  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message       *string  `json:"Message,omitempty" xml:"Message,omitempty"`
	NextPageToken *string  `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	RequestId     *string  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success       *bool    `json:"Success,omitempty" xml:"Success,omitempty"`
	Tables        []*Table `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s ListTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTablesResponseBody) SetCode(v string) *ListTablesResponseBody {
	s.Code = &v
	return s
}

func (s *ListTablesResponseBody) SetMessage(v string) *ListTablesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTablesResponseBody) SetNextPageToken(v string) *ListTablesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListTablesResponseBody) SetRequestId(v string) *ListTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTablesResponseBody) SetSuccess(v bool) *ListTablesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTablesResponseBody) SetTables(v []*Table) *ListTablesResponseBody {
	s.Tables = v
	return s
}

type ListTablesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTablesResponse) GoString() string {
	return s.String()
}

func (s *ListTablesResponse) SetHeaders(v map[string]*string) *ListTablesResponse {
	s.Headers = v
	return s
}

func (s *ListTablesResponse) SetStatusCode(v int32) *ListTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTablesResponse) SetBody(v *ListTablesResponseBody) *ListTablesResponse {
	s.Body = v
	return s
}

type ListUserRolesRequest struct {
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// PageSize
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PrincipalArn *string `json:"PrincipalArn,omitempty" xml:"PrincipalArn,omitempty"`
	// role name pattern filter
	RoleNamePattern *string `json:"RoleNamePattern,omitempty" xml:"RoleNamePattern,omitempty"`
}

func (s ListUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ListUserRolesRequest) SetNextPageToken(v string) *ListUserRolesRequest {
	s.NextPageToken = &v
	return s
}

func (s *ListUserRolesRequest) SetPageSize(v int32) *ListUserRolesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserRolesRequest) SetPrincipalArn(v string) *ListUserRolesRequest {
	s.PrincipalArn = &v
	return s
}

func (s *ListUserRolesRequest) SetRoleNamePattern(v string) *ListUserRolesRequest {
	s.RoleNamePattern = &v
	return s
}

type ListUserRolesResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// NextPageToken
	NextPageToken *string `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// UserRoles
	UserRoles []*UserRole `json:"UserRoles,omitempty" xml:"UserRoles,omitempty" type:"Repeated"`
}

func (s ListUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponseBody) SetCode(v string) *ListUserRolesResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserRolesResponseBody) SetMessage(v string) *ListUserRolesResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserRolesResponseBody) SetNextPageToken(v string) *ListUserRolesResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListUserRolesResponseBody) SetRequestId(v string) *ListUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserRolesResponseBody) SetSuccess(v bool) *ListUserRolesResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserRolesResponseBody) SetUserRoles(v []*UserRole) *ListUserRolesResponseBody {
	s.UserRoles = v
	return s
}

type ListUserRolesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ListUserRolesResponse) SetHeaders(v map[string]*string) *ListUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ListUserRolesResponse) SetStatusCode(v int32) *ListUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserRolesResponse) SetBody(v *ListUserRolesResponseBody) *ListUserRolesResponse {
	s.Body = v
	return s
}

type RefreshLockRequest struct {
	// LockId
	LockId *int64 `json:"LockId,omitempty" xml:"LockId,omitempty"`
}

func (s RefreshLockRequest) String() string {
	return tea.Prettify(s)
}

func (s RefreshLockRequest) GoString() string {
	return s.String()
}

func (s *RefreshLockRequest) SetLockId(v int64) *RefreshLockRequest {
	s.LockId = &v
	return s
}

type RefreshLockResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RefreshLockResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshLockResponseBody) SetCode(v string) *RefreshLockResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshLockResponseBody) SetMessage(v string) *RefreshLockResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshLockResponseBody) SetRequestId(v string) *RefreshLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshLockResponseBody) SetSuccess(v bool) *RefreshLockResponseBody {
	s.Success = &v
	return s
}

type RefreshLockResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RefreshLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RefreshLockResponse) String() string {
	return tea.Prettify(s)
}

func (s RefreshLockResponse) GoString() string {
	return s.String()
}

func (s *RefreshLockResponse) SetHeaders(v map[string]*string) *RefreshLockResponse {
	s.Headers = v
	return s
}

func (s *RefreshLockResponse) SetStatusCode(v int32) *RefreshLockResponse {
	s.StatusCode = &v
	return s
}

func (s *RefreshLockResponse) SetBody(v *RefreshLockResponseBody) *RefreshLockResponse {
	s.Body = v
	return s
}

type RegisterLocationRequest struct {
	InventoryCollectEnabled *bool   `json:"InventoryCollectEnabled,omitempty" xml:"InventoryCollectEnabled,omitempty"`
	Location                *string `json:"Location,omitempty" xml:"Location,omitempty"`
	OssLogCollectEnabled    *bool   `json:"OssLogCollectEnabled,omitempty" xml:"OssLogCollectEnabled,omitempty"`
	RoleName                *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s RegisterLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterLocationRequest) GoString() string {
	return s.String()
}

func (s *RegisterLocationRequest) SetInventoryCollectEnabled(v bool) *RegisterLocationRequest {
	s.InventoryCollectEnabled = &v
	return s
}

func (s *RegisterLocationRequest) SetLocation(v string) *RegisterLocationRequest {
	s.Location = &v
	return s
}

func (s *RegisterLocationRequest) SetOssLogCollectEnabled(v bool) *RegisterLocationRequest {
	s.OssLogCollectEnabled = &v
	return s
}

func (s *RegisterLocationRequest) SetRoleName(v string) *RegisterLocationRequest {
	s.RoleName = &v
	return s
}

type RegisterLocationResponseBody struct {
	Data      *RegisterLocationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RegisterLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterLocationResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterLocationResponseBody) SetData(v *RegisterLocationResponseBodyData) *RegisterLocationResponseBody {
	s.Data = v
	return s
}

func (s *RegisterLocationResponseBody) SetRequestId(v string) *RegisterLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterLocationResponseBody) SetSuccess(v bool) *RegisterLocationResponseBody {
	s.Success = &v
	return s
}

type RegisterLocationResponseBodyData struct {
	// Location ID
	LocationId                            *string                              `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	StorageCollectTaskOperationResultList []*StorageCollectTaskOperationResult `json:"StorageCollectTaskOperationResultList,omitempty" xml:"StorageCollectTaskOperationResultList,omitempty" type:"Repeated"`
}

func (s RegisterLocationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s RegisterLocationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RegisterLocationResponseBodyData) SetLocationId(v string) *RegisterLocationResponseBodyData {
	s.LocationId = &v
	return s
}

func (s *RegisterLocationResponseBodyData) SetStorageCollectTaskOperationResultList(v []*StorageCollectTaskOperationResult) *RegisterLocationResponseBodyData {
	s.StorageCollectTaskOperationResultList = v
	return s
}

type RegisterLocationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RegisterLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterLocationResponse) GoString() string {
	return s.String()
}

func (s *RegisterLocationResponse) SetHeaders(v map[string]*string) *RegisterLocationResponse {
	s.Headers = v
	return s
}

func (s *RegisterLocationResponse) SetStatusCode(v int32) *RegisterLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterLocationResponse) SetBody(v *RegisterLocationResponseBody) *RegisterLocationResponse {
	s.Body = v
	return s
}

type RenamePartitionRequest struct {
	CatalogId       *string         `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName    *string         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PartitionInput  *PartitionInput `json:"PartitionInput,omitempty" xml:"PartitionInput,omitempty"`
	PartitionValues []*string       `json:"PartitionValues,omitempty" xml:"PartitionValues,omitempty" type:"Repeated"`
	TableName       *string         `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RenamePartitionRequest) String() string {
	return tea.Prettify(s)
}

func (s RenamePartitionRequest) GoString() string {
	return s.String()
}

func (s *RenamePartitionRequest) SetCatalogId(v string) *RenamePartitionRequest {
	s.CatalogId = &v
	return s
}

func (s *RenamePartitionRequest) SetDatabaseName(v string) *RenamePartitionRequest {
	s.DatabaseName = &v
	return s
}

func (s *RenamePartitionRequest) SetPartitionInput(v *PartitionInput) *RenamePartitionRequest {
	s.PartitionInput = v
	return s
}

func (s *RenamePartitionRequest) SetPartitionValues(v []*string) *RenamePartitionRequest {
	s.PartitionValues = v
	return s
}

func (s *RenamePartitionRequest) SetTableName(v string) *RenamePartitionRequest {
	s.TableName = &v
	return s
}

type RenamePartitionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenamePartitionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenamePartitionResponseBody) GoString() string {
	return s.String()
}

func (s *RenamePartitionResponseBody) SetCode(v string) *RenamePartitionResponseBody {
	s.Code = &v
	return s
}

func (s *RenamePartitionResponseBody) SetMessage(v string) *RenamePartitionResponseBody {
	s.Message = &v
	return s
}

func (s *RenamePartitionResponseBody) SetRequestId(v string) *RenamePartitionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenamePartitionResponseBody) SetSuccess(v bool) *RenamePartitionResponseBody {
	s.Success = &v
	return s
}

type RenamePartitionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenamePartitionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenamePartitionResponse) String() string {
	return tea.Prettify(s)
}

func (s RenamePartitionResponse) GoString() string {
	return s.String()
}

func (s *RenamePartitionResponse) SetHeaders(v map[string]*string) *RenamePartitionResponse {
	s.Headers = v
	return s
}

func (s *RenamePartitionResponse) SetStatusCode(v int32) *RenamePartitionResponse {
	s.StatusCode = &v
	return s
}

func (s *RenamePartitionResponse) SetBody(v *RenamePartitionResponseBody) *RenamePartitionResponse {
	s.Body = v
	return s
}

type RenameTableRequest struct {
	CatalogId    *string     `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName *string     `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IsAsync      *bool       `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
	TableInput   *TableInput `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	TableName    *string     `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s RenameTableRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameTableRequest) GoString() string {
	return s.String()
}

func (s *RenameTableRequest) SetCatalogId(v string) *RenameTableRequest {
	s.CatalogId = &v
	return s
}

func (s *RenameTableRequest) SetDatabaseName(v string) *RenameTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *RenameTableRequest) SetIsAsync(v bool) *RenameTableRequest {
	s.IsAsync = &v
	return s
}

func (s *RenameTableRequest) SetTableInput(v *TableInput) *RenameTableRequest {
	s.TableInput = v
	return s
}

func (s *RenameTableRequest) SetTableName(v string) *RenameTableRequest {
	s.TableName = &v
	return s
}

type RenameTableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// Async task Id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RenameTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameTableResponseBody) GoString() string {
	return s.String()
}

func (s *RenameTableResponseBody) SetCode(v string) *RenameTableResponseBody {
	s.Code = &v
	return s
}

func (s *RenameTableResponseBody) SetMessage(v string) *RenameTableResponseBody {
	s.Message = &v
	return s
}

func (s *RenameTableResponseBody) SetRequestId(v string) *RenameTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameTableResponseBody) SetSuccess(v bool) *RenameTableResponseBody {
	s.Success = &v
	return s
}

func (s *RenameTableResponseBody) SetTaskId(v string) *RenameTableResponseBody {
	s.TaskId = &v
	return s
}

type RenameTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenameTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameTableResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameTableResponse) GoString() string {
	return s.String()
}

func (s *RenameTableResponse) SetHeaders(v map[string]*string) *RenameTableResponse {
	s.Headers = v
	return s
}

func (s *RenameTableResponse) SetStatusCode(v int32) *RenameTableResponse {
	s.StatusCode = &v
	return s
}

func (s *RenameTableResponse) SetBody(v *RenameTableResponseBody) *RenameTableResponse {
	s.Body = v
	return s
}

type RevokePermissionsRequest struct {
	Accesses []*string `json:"Accesses,omitempty" xml:"Accesses,omitempty" type:"Repeated"`
	// CatalogId
	CatalogId        *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DelegateAccesses []*string     `json:"DelegateAccesses,omitempty" xml:"DelegateAccesses,omitempty" type:"Repeated"`
	MetaResource     *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	Principal        *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
	Type             *string       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RevokePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionsRequest) GoString() string {
	return s.String()
}

func (s *RevokePermissionsRequest) SetAccesses(v []*string) *RevokePermissionsRequest {
	s.Accesses = v
	return s
}

func (s *RevokePermissionsRequest) SetCatalogId(v string) *RevokePermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *RevokePermissionsRequest) SetDelegateAccesses(v []*string) *RevokePermissionsRequest {
	s.DelegateAccesses = v
	return s
}

func (s *RevokePermissionsRequest) SetMetaResource(v *MetaResource) *RevokePermissionsRequest {
	s.MetaResource = v
	return s
}

func (s *RevokePermissionsRequest) SetPrincipal(v *Principal) *RevokePermissionsRequest {
	s.Principal = v
	return s
}

func (s *RevokePermissionsRequest) SetType(v string) *RevokePermissionsRequest {
	s.Type = &v
	return s
}

type RevokePermissionsResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *RevokePermissionsResponseBody) SetCode(v string) *RevokePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *RevokePermissionsResponseBody) SetMessage(v string) *RevokePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *RevokePermissionsResponseBody) SetRequestId(v string) *RevokePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokePermissionsResponseBody) SetSuccess(v bool) *RevokePermissionsResponseBody {
	s.Success = &v
	return s
}

type RevokePermissionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokePermissionsResponse) GoString() string {
	return s.String()
}

func (s *RevokePermissionsResponse) SetHeaders(v map[string]*string) *RevokePermissionsResponse {
	s.Headers = v
	return s
}

func (s *RevokePermissionsResponse) SetStatusCode(v int32) *RevokePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePermissionsResponse) SetBody(v *RevokePermissionsResponseBody) *RevokePermissionsResponse {
	s.Body = v
	return s
}

type RevokeRoleFromUsersRequest struct {
	RoleName *string      `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Users    []*Principal `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s RevokeRoleFromUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeRoleFromUsersRequest) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersRequest) SetRoleName(v string) *RevokeRoleFromUsersRequest {
	s.RoleName = &v
	return s
}

func (s *RevokeRoleFromUsersRequest) SetUsers(v []*Principal) *RevokeRoleFromUsersRequest {
	s.Users = v
	return s
}

type RevokeRoleFromUsersResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeRoleFromUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeRoleFromUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersResponseBody) SetCode(v string) *RevokeRoleFromUsersResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeRoleFromUsersResponseBody) SetMessage(v string) *RevokeRoleFromUsersResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeRoleFromUsersResponseBody) SetRequestId(v string) *RevokeRoleFromUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeRoleFromUsersResponseBody) SetSuccess(v bool) *RevokeRoleFromUsersResponseBody {
	s.Success = &v
	return s
}

type RevokeRoleFromUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeRoleFromUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeRoleFromUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeRoleFromUsersResponse) GoString() string {
	return s.String()
}

func (s *RevokeRoleFromUsersResponse) SetHeaders(v map[string]*string) *RevokeRoleFromUsersResponse {
	s.Headers = v
	return s
}

func (s *RevokeRoleFromUsersResponse) SetStatusCode(v int32) *RevokeRoleFromUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRoleFromUsersResponse) SetBody(v *RevokeRoleFromUsersResponseBody) *RevokeRoleFromUsersResponse {
	s.Body = v
	return s
}

type RevokeRolesFromUserRequest struct {
	RoleNames []*string  `json:"RoleNames,omitempty" xml:"RoleNames,omitempty" type:"Repeated"`
	User      *Principal `json:"User,omitempty" xml:"User,omitempty"`
}

func (s RevokeRolesFromUserRequest) String() string {
	return tea.Prettify(s)
}

func (s RevokeRolesFromUserRequest) GoString() string {
	return s.String()
}

func (s *RevokeRolesFromUserRequest) SetRoleNames(v []*string) *RevokeRolesFromUserRequest {
	s.RoleNames = v
	return s
}

func (s *RevokeRolesFromUserRequest) SetUser(v *Principal) *RevokeRolesFromUserRequest {
	s.User = v
	return s
}

type RevokeRolesFromUserResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RevokeRolesFromUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RevokeRolesFromUserResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeRolesFromUserResponseBody) SetCode(v string) *RevokeRolesFromUserResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeRolesFromUserResponseBody) SetMessage(v string) *RevokeRolesFromUserResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeRolesFromUserResponseBody) SetRequestId(v string) *RevokeRolesFromUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeRolesFromUserResponseBody) SetSuccess(v bool) *RevokeRolesFromUserResponseBody {
	s.Success = &v
	return s
}

type RevokeRolesFromUserResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RevokeRolesFromUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RevokeRolesFromUserResponse) String() string {
	return tea.Prettify(s)
}

func (s RevokeRolesFromUserResponse) GoString() string {
	return s.String()
}

func (s *RevokeRolesFromUserResponse) SetHeaders(v map[string]*string) *RevokeRolesFromUserResponse {
	s.Headers = v
	return s
}

func (s *RevokeRolesFromUserResponse) SetStatusCode(v int32) *RevokeRolesFromUserResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRolesFromUserResponse) SetBody(v *RevokeRolesFromUserResponseBody) *RevokeRolesFromUserResponse {
	s.Body = v
	return s
}

type RunMigrationWorkflowRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RunMigrationWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s RunMigrationWorkflowRequest) GoString() string {
	return s.String()
}

func (s *RunMigrationWorkflowRequest) SetInstanceId(v string) *RunMigrationWorkflowRequest {
	s.InstanceId = &v
	return s
}

type RunMigrationWorkflowResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunMigrationWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunMigrationWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *RunMigrationWorkflowResponseBody) SetData(v string) *RunMigrationWorkflowResponseBody {
	s.Data = &v
	return s
}

func (s *RunMigrationWorkflowResponseBody) SetRequestId(v string) *RunMigrationWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunMigrationWorkflowResponseBody) SetSuccess(v bool) *RunMigrationWorkflowResponseBody {
	s.Success = &v
	return s
}

type RunMigrationWorkflowResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunMigrationWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunMigrationWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s RunMigrationWorkflowResponse) GoString() string {
	return s.String()
}

func (s *RunMigrationWorkflowResponse) SetHeaders(v map[string]*string) *RunMigrationWorkflowResponse {
	s.Headers = v
	return s
}

func (s *RunMigrationWorkflowResponse) SetStatusCode(v int32) *RunMigrationWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMigrationWorkflowResponse) SetBody(v *RunMigrationWorkflowResponseBody) *RunMigrationWorkflowResponse {
	s.Body = v
	return s
}

type SearchRequest struct {
	// catalogid
	CatalogId    *string          `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	PageNumber   *int64           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchText   *string          `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	SearchType   *string          `json:"SearchType,omitempty" xml:"SearchType,omitempty"`
	SortCriteria []*SortCriterion `json:"SortCriteria,omitempty" xml:"SortCriteria,omitempty" type:"Repeated"`
}

func (s SearchRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchRequest) GoString() string {
	return s.String()
}

func (s *SearchRequest) SetCatalogId(v string) *SearchRequest {
	s.CatalogId = &v
	return s
}

func (s *SearchRequest) SetPageNumber(v int64) *SearchRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRequest) SetPageSize(v int64) *SearchRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRequest) SetSearchText(v string) *SearchRequest {
	s.SearchText = &v
	return s
}

func (s *SearchRequest) SetSearchType(v string) *SearchRequest {
	s.SearchType = &v
	return s
}

func (s *SearchRequest) SetSortCriteria(v []*SortCriterion) *SearchRequest {
	s.SortCriteria = v
	return s
}

type SearchResponseBody struct {
	Code           *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DatabaseResult *SearchResponseBodyDatabaseResult `json:"DatabaseResult,omitempty" xml:"DatabaseResult,omitempty" type:"Struct"`
	Message        *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TableResult    *SearchResponseBodyTableResult    `json:"TableResult,omitempty" xml:"TableResult,omitempty" type:"Struct"`
}

func (s SearchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBody) GoString() string {
	return s.String()
}

func (s *SearchResponseBody) SetCode(v string) *SearchResponseBody {
	s.Code = &v
	return s
}

func (s *SearchResponseBody) SetDatabaseResult(v *SearchResponseBodyDatabaseResult) *SearchResponseBody {
	s.DatabaseResult = v
	return s
}

func (s *SearchResponseBody) SetMessage(v string) *SearchResponseBody {
	s.Message = &v
	return s
}

func (s *SearchResponseBody) SetRequestId(v string) *SearchResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchResponseBody) SetSuccess(v bool) *SearchResponseBody {
	s.Success = &v
	return s
}

func (s *SearchResponseBody) SetTableResult(v *SearchResponseBodyTableResult) *SearchResponseBody {
	s.TableResult = v
	return s
}

type SearchResponseBodyDatabaseResult struct {
	Databases  []*SearchResponseBodyDatabaseResultDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	TotalCount *int64                                       `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchResponseBodyDatabaseResult) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDatabaseResult) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDatabaseResult) SetDatabases(v []*SearchResponseBodyDatabaseResultDatabases) *SearchResponseBodyDatabaseResult {
	s.Databases = v
	return s
}

func (s *SearchResponseBodyDatabaseResult) SetTotalCount(v int64) *SearchResponseBodyDatabaseResult {
	s.TotalCount = &v
	return s
}

type SearchResponseBodyDatabaseResultDatabases struct {
	Database      *Database    `json:"Database,omitempty" xml:"Database,omitempty"`
	HighLightList []*HighLight `json:"HighLightList,omitempty" xml:"HighLightList,omitempty" type:"Repeated"`
}

func (s SearchResponseBodyDatabaseResultDatabases) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyDatabaseResultDatabases) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyDatabaseResultDatabases) SetDatabase(v *Database) *SearchResponseBodyDatabaseResultDatabases {
	s.Database = v
	return s
}

func (s *SearchResponseBodyDatabaseResultDatabases) SetHighLightList(v []*HighLight) *SearchResponseBodyDatabaseResultDatabases {
	s.HighLightList = v
	return s
}

type SearchResponseBodyTableResult struct {
	Tables     []*SearchResponseBodyTableResultTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchResponseBodyTableResult) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyTableResult) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyTableResult) SetTables(v []*SearchResponseBodyTableResultTables) *SearchResponseBodyTableResult {
	s.Tables = v
	return s
}

func (s *SearchResponseBodyTableResult) SetTotalCount(v int64) *SearchResponseBodyTableResult {
	s.TotalCount = &v
	return s
}

type SearchResponseBodyTableResultTables struct {
	HighLightList []*HighLight `json:"HighLightList,omitempty" xml:"HighLightList,omitempty" type:"Repeated"`
	Table         *Table       `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s SearchResponseBodyTableResultTables) String() string {
	return tea.Prettify(s)
}

func (s SearchResponseBodyTableResultTables) GoString() string {
	return s.String()
}

func (s *SearchResponseBodyTableResultTables) SetHighLightList(v []*HighLight) *SearchResponseBodyTableResultTables {
	s.HighLightList = v
	return s
}

func (s *SearchResponseBodyTableResultTables) SetTable(v *Table) *SearchResponseBodyTableResultTables {
	s.Table = v
	return s
}

type SearchResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchResponse) GoString() string {
	return s.String()
}

func (s *SearchResponse) SetHeaders(v map[string]*string) *SearchResponse {
	s.Headers = v
	return s
}

func (s *SearchResponse) SetStatusCode(v int32) *SearchResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchResponse) SetBody(v *SearchResponseBody) *SearchResponse {
	s.Body = v
	return s
}

type SearchAcrossCatalogRequest struct {
	CatalogIds   []*string        `json:"CatalogIds,omitempty" xml:"CatalogIds,omitempty" type:"Repeated"`
	PageNumber   *int64           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int64           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchText   *string          `json:"SearchText,omitempty" xml:"SearchText,omitempty"`
	SearchTypes  []*string        `json:"SearchTypes,omitempty" xml:"SearchTypes,omitempty" type:"Repeated"`
	SortCriteria []*SortCriterion `json:"SortCriteria,omitempty" xml:"SortCriteria,omitempty" type:"Repeated"`
}

func (s SearchAcrossCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogRequest) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogRequest) SetCatalogIds(v []*string) *SearchAcrossCatalogRequest {
	s.CatalogIds = v
	return s
}

func (s *SearchAcrossCatalogRequest) SetPageNumber(v int64) *SearchAcrossCatalogRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchAcrossCatalogRequest) SetPageSize(v int64) *SearchAcrossCatalogRequest {
	s.PageSize = &v
	return s
}

func (s *SearchAcrossCatalogRequest) SetSearchText(v string) *SearchAcrossCatalogRequest {
	s.SearchText = &v
	return s
}

func (s *SearchAcrossCatalogRequest) SetSearchTypes(v []*string) *SearchAcrossCatalogRequest {
	s.SearchTypes = v
	return s
}

func (s *SearchAcrossCatalogRequest) SetSortCriteria(v []*SortCriterion) *SearchAcrossCatalogRequest {
	s.SortCriteria = v
	return s
}

type SearchAcrossCatalogResponseBody struct {
	CatalogResult  *SearchAcrossCatalogResponseBodyCatalogResult  `json:"CatalogResult,omitempty" xml:"CatalogResult,omitempty" type:"Struct"`
	Code           *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	DatabaseResult *SearchAcrossCatalogResponseBodyDatabaseResult `json:"DatabaseResult,omitempty" xml:"DatabaseResult,omitempty" type:"Struct"`
	Message        *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TableResult    *SearchAcrossCatalogResponseBodyTableResult    `json:"TableResult,omitempty" xml:"TableResult,omitempty" type:"Struct"`
}

func (s SearchAcrossCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBody) SetCatalogResult(v *SearchAcrossCatalogResponseBodyCatalogResult) *SearchAcrossCatalogResponseBody {
	s.CatalogResult = v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetCode(v string) *SearchAcrossCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetDatabaseResult(v *SearchAcrossCatalogResponseBodyDatabaseResult) *SearchAcrossCatalogResponseBody {
	s.DatabaseResult = v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetMessage(v string) *SearchAcrossCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetRequestId(v string) *SearchAcrossCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetSuccess(v bool) *SearchAcrossCatalogResponseBody {
	s.Success = &v
	return s
}

func (s *SearchAcrossCatalogResponseBody) SetTableResult(v *SearchAcrossCatalogResponseBodyTableResult) *SearchAcrossCatalogResponseBody {
	s.TableResult = v
	return s
}

type SearchAcrossCatalogResponseBodyCatalogResult struct {
	Catalogs   []*SearchAcrossCatalogResponseBodyCatalogResultCatalogs `json:"Catalogs,omitempty" xml:"Catalogs,omitempty" type:"Repeated"`
	TotalCount *int64                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAcrossCatalogResponseBodyCatalogResult) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyCatalogResult) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyCatalogResult) SetCatalogs(v []*SearchAcrossCatalogResponseBodyCatalogResultCatalogs) *SearchAcrossCatalogResponseBodyCatalogResult {
	s.Catalogs = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyCatalogResult) SetTotalCount(v int64) *SearchAcrossCatalogResponseBodyCatalogResult {
	s.TotalCount = &v
	return s
}

type SearchAcrossCatalogResponseBodyCatalogResultCatalogs struct {
	Catalog       *Catalog     `json:"Catalog,omitempty" xml:"Catalog,omitempty"`
	HighLightList []*HighLight `json:"HighLightList,omitempty" xml:"HighLightList,omitempty" type:"Repeated"`
}

func (s SearchAcrossCatalogResponseBodyCatalogResultCatalogs) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyCatalogResultCatalogs) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyCatalogResultCatalogs) SetCatalog(v *Catalog) *SearchAcrossCatalogResponseBodyCatalogResultCatalogs {
	s.Catalog = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyCatalogResultCatalogs) SetHighLightList(v []*HighLight) *SearchAcrossCatalogResponseBodyCatalogResultCatalogs {
	s.HighLightList = v
	return s
}

type SearchAcrossCatalogResponseBodyDatabaseResult struct {
	Databases  []*SearchAcrossCatalogResponseBodyDatabaseResultDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	TotalCount *int64                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAcrossCatalogResponseBodyDatabaseResult) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyDatabaseResult) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyDatabaseResult) SetDatabases(v []*SearchAcrossCatalogResponseBodyDatabaseResultDatabases) *SearchAcrossCatalogResponseBodyDatabaseResult {
	s.Databases = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyDatabaseResult) SetTotalCount(v int64) *SearchAcrossCatalogResponseBodyDatabaseResult {
	s.TotalCount = &v
	return s
}

type SearchAcrossCatalogResponseBodyDatabaseResultDatabases struct {
	Database      *Database    `json:"Database,omitempty" xml:"Database,omitempty"`
	HighLightList []*HighLight `json:"HighLightList,omitempty" xml:"HighLightList,omitempty" type:"Repeated"`
}

func (s SearchAcrossCatalogResponseBodyDatabaseResultDatabases) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyDatabaseResultDatabases) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyDatabaseResultDatabases) SetDatabase(v *Database) *SearchAcrossCatalogResponseBodyDatabaseResultDatabases {
	s.Database = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyDatabaseResultDatabases) SetHighLightList(v []*HighLight) *SearchAcrossCatalogResponseBodyDatabaseResultDatabases {
	s.HighLightList = v
	return s
}

type SearchAcrossCatalogResponseBodyTableResult struct {
	Tables     []*SearchAcrossCatalogResponseBodyTableResultTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TotalCount *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchAcrossCatalogResponseBodyTableResult) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyTableResult) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyTableResult) SetTables(v []*SearchAcrossCatalogResponseBodyTableResultTables) *SearchAcrossCatalogResponseBodyTableResult {
	s.Tables = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyTableResult) SetTotalCount(v int64) *SearchAcrossCatalogResponseBodyTableResult {
	s.TotalCount = &v
	return s
}

type SearchAcrossCatalogResponseBodyTableResultTables struct {
	HighLightList []*HighLight `json:"HighLightList,omitempty" xml:"HighLightList,omitempty" type:"Repeated"`
	Table         *Table       `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s SearchAcrossCatalogResponseBodyTableResultTables) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponseBodyTableResultTables) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponseBodyTableResultTables) SetHighLightList(v []*HighLight) *SearchAcrossCatalogResponseBodyTableResultTables {
	s.HighLightList = v
	return s
}

func (s *SearchAcrossCatalogResponseBodyTableResultTables) SetTable(v *Table) *SearchAcrossCatalogResponseBodyTableResultTables {
	s.Table = v
	return s
}

type SearchAcrossCatalogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAcrossCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchAcrossCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAcrossCatalogResponse) GoString() string {
	return s.String()
}

func (s *SearchAcrossCatalogResponse) SetHeaders(v map[string]*string) *SearchAcrossCatalogResponse {
	s.Headers = v
	return s
}

func (s *SearchAcrossCatalogResponse) SetStatusCode(v int32) *SearchAcrossCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAcrossCatalogResponse) SetBody(v *SearchAcrossCatalogResponseBody) *SearchAcrossCatalogResponse {
	s.Body = v
	return s
}

type StopMigrationWorkflowRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopMigrationWorkflowRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationWorkflowRequest) GoString() string {
	return s.String()
}

func (s *StopMigrationWorkflowRequest) SetInstanceId(v string) *StopMigrationWorkflowRequest {
	s.InstanceId = &v
	return s
}

type StopMigrationWorkflowResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopMigrationWorkflowResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationWorkflowResponseBody) GoString() string {
	return s.String()
}

func (s *StopMigrationWorkflowResponseBody) SetRequestId(v string) *StopMigrationWorkflowResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMigrationWorkflowResponseBody) SetSuccess(v bool) *StopMigrationWorkflowResponseBody {
	s.Success = &v
	return s
}

type StopMigrationWorkflowResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopMigrationWorkflowResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMigrationWorkflowResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationWorkflowResponse) GoString() string {
	return s.String()
}

func (s *StopMigrationWorkflowResponse) SetHeaders(v map[string]*string) *StopMigrationWorkflowResponse {
	s.Headers = v
	return s
}

func (s *StopMigrationWorkflowResponse) SetStatusCode(v int32) *StopMigrationWorkflowResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMigrationWorkflowResponse) SetBody(v *StopMigrationWorkflowResponseBody) *StopMigrationWorkflowResponse {
	s.Body = v
	return s
}

type SubmitQueryRequest struct {
	CatalogId   *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	Sql         *string `json:"sql,omitempty" xml:"sql,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s SubmitQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s SubmitQueryRequest) GoString() string {
	return s.String()
}

func (s *SubmitQueryRequest) SetCatalogId(v string) *SubmitQueryRequest {
	s.CatalogId = &v
	return s
}

func (s *SubmitQueryRequest) SetSql(v string) *SubmitQueryRequest {
	s.Sql = &v
	return s
}

func (s *SubmitQueryRequest) SetWorkspaceId(v string) *SubmitQueryRequest {
	s.WorkspaceId = &v
	return s
}

type SubmitQueryResponseBody struct {
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SubmitQueryResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitQueryResponseBody) SetData(v string) *SubmitQueryResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitQueryResponseBody) SetRequestId(v string) *SubmitQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitQueryResponseBody) SetSuccess(v bool) *SubmitQueryResponseBody {
	s.Success = &v
	return s
}

type SubmitQueryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SubmitQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SubmitQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s SubmitQueryResponse) GoString() string {
	return s.String()
}

func (s *SubmitQueryResponse) SetHeaders(v map[string]*string) *SubmitQueryResponse {
	s.Headers = v
	return s
}

func (s *SubmitQueryResponse) SetStatusCode(v int32) *SubmitQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitQueryResponse) SetBody(v *SubmitQueryResponseBody) *SubmitQueryResponse {
	s.Body = v
	return s
}

type UnLockRequest struct {
	// LockId
	LockId *int64 `json:"LockId,omitempty" xml:"LockId,omitempty"`
}

func (s UnLockRequest) String() string {
	return tea.Prettify(s)
}

func (s UnLockRequest) GoString() string {
	return s.String()
}

func (s *UnLockRequest) SetLockId(v int64) *UnLockRequest {
	s.LockId = &v
	return s
}

type UnLockResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UnLockResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnLockResponseBody) GoString() string {
	return s.String()
}

func (s *UnLockResponseBody) SetCode(v string) *UnLockResponseBody {
	s.Code = &v
	return s
}

func (s *UnLockResponseBody) SetMessage(v string) *UnLockResponseBody {
	s.Message = &v
	return s
}

func (s *UnLockResponseBody) SetRequestId(v string) *UnLockResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnLockResponseBody) SetSuccess(v bool) *UnLockResponseBody {
	s.Success = &v
	return s
}

type UnLockResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnLockResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnLockResponse) String() string {
	return tea.Prettify(s)
}

func (s UnLockResponse) GoString() string {
	return s.String()
}

func (s *UnLockResponse) SetHeaders(v map[string]*string) *UnLockResponse {
	s.Headers = v
	return s
}

func (s *UnLockResponse) SetStatusCode(v int32) *UnLockResponse {
	s.StatusCode = &v
	return s
}

func (s *UnLockResponse) SetBody(v *UnLockResponseBody) *UnLockResponse {
	s.Body = v
	return s
}

type UpdateCatalogRequest struct {
	// cataloginput
	CatalogInput *CatalogInput `json:"CatalogInput,omitempty" xml:"CatalogInput,omitempty"`
}

func (s UpdateCatalogRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogRequest) GoString() string {
	return s.String()
}

func (s *UpdateCatalogRequest) SetCatalogInput(v *CatalogInput) *UpdateCatalogRequest {
	s.CatalogInput = v
	return s
}

type UpdateCatalogResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCatalogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCatalogResponseBody) SetCode(v string) *UpdateCatalogResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCatalogResponseBody) SetMessage(v string) *UpdateCatalogResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCatalogResponseBody) SetRequestId(v string) *UpdateCatalogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCatalogResponseBody) SetSuccess(v bool) *UpdateCatalogResponseBody {
	s.Success = &v
	return s
}

type UpdateCatalogResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCatalogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCatalogResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogResponse) GoString() string {
	return s.String()
}

func (s *UpdateCatalogResponse) SetHeaders(v map[string]*string) *UpdateCatalogResponse {
	s.Headers = v
	return s
}

func (s *UpdateCatalogResponse) SetStatusCode(v int32) *UpdateCatalogResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCatalogResponse) SetBody(v *UpdateCatalogResponseBody) *UpdateCatalogResponse {
	s.Body = v
	return s
}

type UpdateCatalogSettingsRequest struct {
	// CatalogId
	CatalogId       *string          `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	CatalogSettings *CatalogSettings `json:"CatalogSettings,omitempty" xml:"CatalogSettings,omitempty"`
}

func (s UpdateCatalogSettingsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogSettingsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCatalogSettingsRequest) SetCatalogId(v string) *UpdateCatalogSettingsRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateCatalogSettingsRequest) SetCatalogSettings(v *CatalogSettings) *UpdateCatalogSettingsRequest {
	s.CatalogSettings = v
	return s
}

type UpdateCatalogSettingsResponseBody struct {
	// Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateCatalogSettingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCatalogSettingsResponseBody) SetCode(v string) *UpdateCatalogSettingsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateCatalogSettingsResponseBody) SetMessage(v string) *UpdateCatalogSettingsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateCatalogSettingsResponseBody) SetRequestId(v string) *UpdateCatalogSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCatalogSettingsResponseBody) SetSuccess(v bool) *UpdateCatalogSettingsResponseBody {
	s.Success = &v
	return s
}

type UpdateCatalogSettingsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateCatalogSettingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateCatalogSettingsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCatalogSettingsResponse) GoString() string {
	return s.String()
}

func (s *UpdateCatalogSettingsResponse) SetHeaders(v map[string]*string) *UpdateCatalogSettingsResponse {
	s.Headers = v
	return s
}

func (s *UpdateCatalogSettingsResponse) SetStatusCode(v int32) *UpdateCatalogSettingsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCatalogSettingsResponse) SetBody(v *UpdateCatalogSettingsResponseBody) *UpdateCatalogSettingsResponse {
	s.Body = v
	return s
}

type UpdateDatabaseRequest struct {
	CatalogId     *string        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseInput *DatabaseInput `json:"DatabaseInput,omitempty" xml:"DatabaseInput,omitempty"`
	Name          *string        `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatabaseRequest) SetCatalogId(v string) *UpdateDatabaseRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateDatabaseRequest) SetDatabaseInput(v *DatabaseInput) *UpdateDatabaseRequest {
	s.DatabaseInput = v
	return s
}

func (s *UpdateDatabaseRequest) SetName(v string) *UpdateDatabaseRequest {
	s.Name = &v
	return s
}

type UpdateDatabaseResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatabaseResponseBody) SetCode(v string) *UpdateDatabaseResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateDatabaseResponseBody) SetMessage(v string) *UpdateDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateDatabaseResponseBody) SetRequestId(v string) *UpdateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatabaseResponseBody) SetSuccess(v bool) *UpdateDatabaseResponseBody {
	s.Success = &v
	return s
}

type UpdateDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDatabaseResponse) GoString() string {
	return s.String()
}

func (s *UpdateDatabaseResponse) SetHeaders(v map[string]*string) *UpdateDatabaseResponse {
	s.Headers = v
	return s
}

func (s *UpdateDatabaseResponse) SetStatusCode(v int32) *UpdateDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDatabaseResponse) SetBody(v *UpdateDatabaseResponseBody) *UpdateDatabaseResponse {
	s.Body = v
	return s
}

type UpdateFunctionRequest struct {
	CatalogId     *string        `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName  *string        `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	FunctionInput *FunctionInput `json:"FunctionInput,omitempty" xml:"FunctionInput,omitempty"`
	FunctionName  *string        `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s UpdateFunctionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionRequest) SetCatalogId(v string) *UpdateFunctionRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateFunctionRequest) SetDatabaseName(v string) *UpdateFunctionRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateFunctionRequest) SetFunctionInput(v *FunctionInput) *UpdateFunctionRequest {
	s.FunctionInput = v
	return s
}

func (s *UpdateFunctionRequest) SetFunctionName(v string) *UpdateFunctionRequest {
	s.FunctionName = &v
	return s
}

type UpdateFunctionResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateFunctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponseBody) SetCode(v string) *UpdateFunctionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetMessage(v string) *UpdateFunctionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetRequestId(v string) *UpdateFunctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResponseBody) SetSuccess(v bool) *UpdateFunctionResponseBody {
	s.Success = &v
	return s
}

type UpdateFunctionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFunctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResponse) SetHeaders(v map[string]*string) *UpdateFunctionResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResponse) SetStatusCode(v int32) *UpdateFunctionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResponse) SetBody(v *UpdateFunctionResponseBody) *UpdateFunctionResponse {
	s.Body = v
	return s
}

type UpdatePartitionColumnStatisticsRequest struct {
	UpdateTablePartitionColumnStatisticsRequest *UpdateTablePartitionColumnStatisticsRequest `json:"UpdateTablePartitionColumnStatisticsRequest,omitempty" xml:"UpdateTablePartitionColumnStatisticsRequest,omitempty"`
}

func (s UpdatePartitionColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartitionColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePartitionColumnStatisticsRequest) SetUpdateTablePartitionColumnStatisticsRequest(v *UpdateTablePartitionColumnStatisticsRequest) *UpdatePartitionColumnStatisticsRequest {
	s.UpdateTablePartitionColumnStatisticsRequest = v
	return s
}

type UpdatePartitionColumnStatisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePartitionColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartitionColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePartitionColumnStatisticsResponseBody) SetCode(v string) *UpdatePartitionColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePartitionColumnStatisticsResponseBody) SetMessage(v string) *UpdatePartitionColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePartitionColumnStatisticsResponseBody) SetRequestId(v string) *UpdatePartitionColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePartitionColumnStatisticsResponseBody) SetSuccess(v bool) *UpdatePartitionColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type UpdatePartitionColumnStatisticsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePartitionColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePartitionColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePartitionColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePartitionColumnStatisticsResponse) SetHeaders(v map[string]*string) *UpdatePartitionColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePartitionColumnStatisticsResponse) SetStatusCode(v int32) *UpdatePartitionColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePartitionColumnStatisticsResponse) SetBody(v *UpdatePartitionColumnStatisticsResponseBody) *UpdatePartitionColumnStatisticsResponse {
	s.Body = v
	return s
}

type UpdatePermissionsRequest struct {
	Accesses []*string `json:"Accesses,omitempty" xml:"Accesses,omitempty" type:"Repeated"`
	// CatalogId
	CatalogId        *string       `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DelegateAccesses []*string     `json:"DelegateAccesses,omitempty" xml:"DelegateAccesses,omitempty" type:"Repeated"`
	MetaResource     *MetaResource `json:"MetaResource,omitempty" xml:"MetaResource,omitempty"`
	Principal        *Principal    `json:"Principal,omitempty" xml:"Principal,omitempty"`
	Type             *string       `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdatePermissionsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionsRequest) GoString() string {
	return s.String()
}

func (s *UpdatePermissionsRequest) SetAccesses(v []*string) *UpdatePermissionsRequest {
	s.Accesses = v
	return s
}

func (s *UpdatePermissionsRequest) SetCatalogId(v string) *UpdatePermissionsRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdatePermissionsRequest) SetDelegateAccesses(v []*string) *UpdatePermissionsRequest {
	s.DelegateAccesses = v
	return s
}

func (s *UpdatePermissionsRequest) SetMetaResource(v *MetaResource) *UpdatePermissionsRequest {
	s.MetaResource = v
	return s
}

func (s *UpdatePermissionsRequest) SetPrincipal(v *Principal) *UpdatePermissionsRequest {
	s.Principal = v
	return s
}

func (s *UpdatePermissionsRequest) SetType(v string) *UpdatePermissionsRequest {
	s.Type = &v
	return s
}

type UpdatePermissionsResponseBody struct {
	// Response Code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Message Code
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdatePermissionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdatePermissionsResponseBody) SetCode(v string) *UpdatePermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdatePermissionsResponseBody) SetMessage(v string) *UpdatePermissionsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdatePermissionsResponseBody) SetRequestId(v string) *UpdatePermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdatePermissionsResponseBody) SetSuccess(v bool) *UpdatePermissionsResponseBody {
	s.Success = &v
	return s
}

type UpdatePermissionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdatePermissionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdatePermissionsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdatePermissionsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePermissionsResponse) SetHeaders(v map[string]*string) *UpdatePermissionsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePermissionsResponse) SetStatusCode(v int32) *UpdatePermissionsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePermissionsResponse) SetBody(v *UpdatePermissionsResponseBody) *UpdatePermissionsResponse {
	s.Body = v
	return s
}

type UpdateRegisteredLocationRequest struct {
	InventoryCollectEnabled *bool   `json:"InventoryCollectEnabled,omitempty" xml:"InventoryCollectEnabled,omitempty"`
	LocationId              *string `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	OssLogCollectEnabled    *bool   `json:"OssLogCollectEnabled,omitempty" xml:"OssLogCollectEnabled,omitempty"`
}

func (s UpdateRegisteredLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegisteredLocationRequest) GoString() string {
	return s.String()
}

func (s *UpdateRegisteredLocationRequest) SetInventoryCollectEnabled(v bool) *UpdateRegisteredLocationRequest {
	s.InventoryCollectEnabled = &v
	return s
}

func (s *UpdateRegisteredLocationRequest) SetLocationId(v string) *UpdateRegisteredLocationRequest {
	s.LocationId = &v
	return s
}

func (s *UpdateRegisteredLocationRequest) SetOssLogCollectEnabled(v bool) *UpdateRegisteredLocationRequest {
	s.OssLogCollectEnabled = &v
	return s
}

type UpdateRegisteredLocationResponseBody struct {
	Data      *UpdateRegisteredLocationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRegisteredLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegisteredLocationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRegisteredLocationResponseBody) SetData(v *UpdateRegisteredLocationResponseBodyData) *UpdateRegisteredLocationResponseBody {
	s.Data = v
	return s
}

func (s *UpdateRegisteredLocationResponseBody) SetRequestId(v string) *UpdateRegisteredLocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRegisteredLocationResponseBody) SetSuccess(v bool) *UpdateRegisteredLocationResponseBody {
	s.Success = &v
	return s
}

type UpdateRegisteredLocationResponseBodyData struct {
	// Location ID
	LocationId                            *string                              `json:"LocationId,omitempty" xml:"LocationId,omitempty"`
	StorageCollectTaskOperationResultList []*StorageCollectTaskOperationResult `json:"StorageCollectTaskOperationResultList,omitempty" xml:"StorageCollectTaskOperationResultList,omitempty" type:"Repeated"`
}

func (s UpdateRegisteredLocationResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegisteredLocationResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateRegisteredLocationResponseBodyData) SetLocationId(v string) *UpdateRegisteredLocationResponseBodyData {
	s.LocationId = &v
	return s
}

func (s *UpdateRegisteredLocationResponseBodyData) SetStorageCollectTaskOperationResultList(v []*StorageCollectTaskOperationResult) *UpdateRegisteredLocationResponseBodyData {
	s.StorageCollectTaskOperationResultList = v
	return s
}

type UpdateRegisteredLocationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRegisteredLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRegisteredLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRegisteredLocationResponse) GoString() string {
	return s.String()
}

func (s *UpdateRegisteredLocationResponse) SetHeaders(v map[string]*string) *UpdateRegisteredLocationResponse {
	s.Headers = v
	return s
}

func (s *UpdateRegisteredLocationResponse) SetStatusCode(v int32) *UpdateRegisteredLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRegisteredLocationResponse) SetBody(v *UpdateRegisteredLocationResponseBody) *UpdateRegisteredLocationResponse {
	s.Body = v
	return s
}

type UpdateRoleRequest struct {
	RoleInput *RoleInput `json:"RoleInput,omitempty" xml:"RoleInput,omitempty"`
	// RoleName
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s UpdateRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleRequest) SetRoleInput(v *RoleInput) *UpdateRoleRequest {
	s.RoleInput = v
	return s
}

func (s *UpdateRoleRequest) SetRoleName(v string) *UpdateRoleRequest {
	s.RoleName = &v
	return s
}

type UpdateRoleResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRoleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponseBody) SetCode(v string) *UpdateRoleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleResponseBody) SetMessage(v string) *UpdateRoleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleResponseBody) SetRequestId(v string) *UpdateRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleResponseBody) SetSuccess(v bool) *UpdateRoleResponseBody {
	s.Success = &v
	return s
}

type UpdateRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleResponse) SetHeaders(v map[string]*string) *UpdateRoleResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleResponse) SetStatusCode(v int32) *UpdateRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleResponse) SetBody(v *UpdateRoleResponseBody) *UpdateRoleResponse {
	s.Body = v
	return s
}

type UpdateRoleUsersRequest struct {
	RoleName *string      `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	Users    []*Principal `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s UpdateRoleUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleUsersRequest) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersRequest) SetRoleName(v string) *UpdateRoleUsersRequest {
	s.RoleName = &v
	return s
}

func (s *UpdateRoleUsersRequest) SetUsers(v []*Principal) *UpdateRoleUsersRequest {
	s.Users = v
	return s
}

type UpdateRoleUsersResponseBody struct {
	// code
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// requestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// success
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRoleUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleUsersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersResponseBody) SetCode(v string) *UpdateRoleUsersResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRoleUsersResponseBody) SetMessage(v string) *UpdateRoleUsersResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRoleUsersResponseBody) SetRequestId(v string) *UpdateRoleUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoleUsersResponseBody) SetSuccess(v bool) *UpdateRoleUsersResponseBody {
	s.Success = &v
	return s
}

type UpdateRoleUsersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateRoleUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateRoleUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRoleUsersResponse) GoString() string {
	return s.String()
}

func (s *UpdateRoleUsersResponse) SetHeaders(v map[string]*string) *UpdateRoleUsersResponse {
	s.Headers = v
	return s
}

func (s *UpdateRoleUsersResponse) SetStatusCode(v int32) *UpdateRoleUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRoleUsersResponse) SetBody(v *UpdateRoleUsersResponseBody) *UpdateRoleUsersResponse {
	s.Body = v
	return s
}

type UpdateTableRequest struct {
	AllowPartitionKeyChange *bool       `json:"AllowPartitionKeyChange,omitempty" xml:"AllowPartitionKeyChange,omitempty"`
	CatalogId               *string     `json:"CatalogId,omitempty" xml:"CatalogId,omitempty"`
	DatabaseName            *string     `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	IsAsync                 *bool       `json:"IsAsync,omitempty" xml:"IsAsync,omitempty"`
	SkipArchive             *bool       `json:"SkipArchive,omitempty" xml:"SkipArchive,omitempty"`
	TableInput              *TableInput `json:"TableInput,omitempty" xml:"TableInput,omitempty"`
	TableName               *string     `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s UpdateTableRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableRequest) SetAllowPartitionKeyChange(v bool) *UpdateTableRequest {
	s.AllowPartitionKeyChange = &v
	return s
}

func (s *UpdateTableRequest) SetCatalogId(v string) *UpdateTableRequest {
	s.CatalogId = &v
	return s
}

func (s *UpdateTableRequest) SetDatabaseName(v string) *UpdateTableRequest {
	s.DatabaseName = &v
	return s
}

func (s *UpdateTableRequest) SetIsAsync(v bool) *UpdateTableRequest {
	s.IsAsync = &v
	return s
}

func (s *UpdateTableRequest) SetSkipArchive(v bool) *UpdateTableRequest {
	s.SkipArchive = &v
	return s
}

func (s *UpdateTableRequest) SetTableInput(v *TableInput) *UpdateTableRequest {
	s.TableInput = v
	return s
}

func (s *UpdateTableRequest) SetTableName(v string) *UpdateTableRequest {
	s.TableName = &v
	return s
}

type UpdateTableResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	// Async task Id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableResponseBody) SetCode(v string) *UpdateTableResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTableResponseBody) SetMessage(v string) *UpdateTableResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTableResponseBody) SetRequestId(v string) *UpdateTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableResponseBody) SetSuccess(v bool) *UpdateTableResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateTableResponseBody) SetTaskId(v string) *UpdateTableResponseBody {
	s.TaskId = &v
	return s
}

type UpdateTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTableResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableResponse) SetHeaders(v map[string]*string) *UpdateTableResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableResponse) SetStatusCode(v int32) *UpdateTableResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableResponse) SetBody(v *UpdateTableResponseBody) *UpdateTableResponse {
	s.Body = v
	return s
}

type UpdateTableColumnStatisticsRequest struct {
	UpdateTablePartitionColumnStatisticsRequest *UpdateTablePartitionColumnStatisticsRequest `json:"UpdateTablePartitionColumnStatisticsRequest,omitempty" xml:"UpdateTablePartitionColumnStatisticsRequest,omitempty"`
}

func (s UpdateTableColumnStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableColumnStatisticsRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableColumnStatisticsRequest) SetUpdateTablePartitionColumnStatisticsRequest(v *UpdateTablePartitionColumnStatisticsRequest) *UpdateTableColumnStatisticsRequest {
	s.UpdateTablePartitionColumnStatisticsRequest = v
	return s
}

type UpdateTableColumnStatisticsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateTableColumnStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableColumnStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTableColumnStatisticsResponseBody) SetCode(v string) *UpdateTableColumnStatisticsResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateTableColumnStatisticsResponseBody) SetMessage(v string) *UpdateTableColumnStatisticsResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateTableColumnStatisticsResponseBody) SetRequestId(v string) *UpdateTableColumnStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTableColumnStatisticsResponseBody) SetSuccess(v bool) *UpdateTableColumnStatisticsResponseBody {
	s.Success = &v
	return s
}

type UpdateTableColumnStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTableColumnStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTableColumnStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTableColumnStatisticsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTableColumnStatisticsResponse) SetHeaders(v map[string]*string) *UpdateTableColumnStatisticsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTableColumnStatisticsResponse) SetStatusCode(v int32) *UpdateTableColumnStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTableColumnStatisticsResponse) SetBody(v *UpdateTableColumnStatisticsResponseBody) *UpdateTableColumnStatisticsResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":              tea.String("datalake-daily.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("datalake-daily.aliyuncs.com"),
		"ap-south-1":                  tea.String("datalake-daily.aliyuncs.com"),
		"ap-southeast-1":              tea.String("datalake-daily.aliyuncs.com"),
		"ap-southeast-2":              tea.String("datalake-daily.aliyuncs.com"),
		"ap-southeast-3":              tea.String("datalake-daily.aliyuncs.com"),
		"ap-southeast-5":              tea.String("datalake-daily.aliyuncs.com"),
		"cn-beijing":                  tea.String("dlf.cn-beijing.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("datalake-daily.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("datalake-daily.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("datalake-daily.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("datalake-daily.aliyuncs.com"),
		"cn-chengdu":                  tea.String("datalake-daily.aliyuncs.com"),
		"cn-edge-1":                   tea.String("datalake-daily.aliyuncs.com"),
		"cn-fujian":                   tea.String("datalake-daily.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("dlf.cn-hangzhou.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("datalake-daily.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("datalake-daily.aliyuncs.com"),
		"cn-hongkong":                 tea.String("datalake-daily.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("datalake-daily.aliyuncs.com"),
		"cn-huhehaote":                tea.String("datalake-daily.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("datalake-daily.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("datalake-daily.aliyuncs.com"),
		"cn-qingdao":                  tea.String("datalake-daily.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("datalake-daily.aliyuncs.com"),
		"cn-shanghai":                 tea.String("dlf.cn-shanghai.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("datalake-daily.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("datalake-daily.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("datalake-daily.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("datalake-daily.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("datalake-daily.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("dlf.cn-shenzhen.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("datalake-daily.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("datalake-daily.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("datalake-daily.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("datalake-daily.aliyuncs.com"),
		"cn-wuhan":                    tea.String("datalake-daily.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("datalake-daily.aliyuncs.com"),
		"cn-yushanfang":               tea.String("datalake-daily.aliyuncs.com"),
		"cn-zhangbei":                 tea.String("datalake-daily.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("datalake-daily.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("datalake-daily.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("datalake-daily.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("datalake-daily.aliyuncs.com"),
		"eu-central-1":                tea.String("datalake-daily.aliyuncs.com"),
		"eu-west-1":                   tea.String("datalake-daily.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("datalake-daily.aliyuncs.com"),
		"me-east-1":                   tea.String("datalake-daily.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("datalake-daily.aliyuncs.com"),
		"us-east-1":                   tea.String("datalake-daily.aliyuncs.com"),
		"us-west-1":                   tea.String("datalake-daily.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("datalake"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortLockWithOptions(request *AbortLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LockId)) {
		query["LockId"] = request.LockId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortLock"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/locks/abort"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AbortLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortLock(request *AbortLockRequest) (_result *AbortLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortLockResponse{}
	_body, _err := client.AbortLockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreatePartitionsWithOptions(request *BatchCreatePartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchCreatePartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfNotExists)) {
		body["IfNotExists"] = request.IfNotExists
	}

	if !tea.BoolValue(util.IsUnset(request.NeedResult)) {
		body["NeedResult"] = request.NeedResult
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionInputs)) {
		body["PartitionInputs"] = request.PartitionInputs
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreatePartitions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/batchcreate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreatePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreatePartitions(request *BatchCreatePartitionsRequest) (_result *BatchCreatePartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreatePartitionsResponse{}
	_body, _err := client.BatchCreatePartitionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchCreateTablesWithOptions(request *BatchCreateTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchCreateTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfNotExists)) {
		body["IfNotExists"] = request.IfNotExists
	}

	if !tea.BoolValue(util.IsUnset(request.TableInputs)) {
		body["TableInputs"] = request.TableInputs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchCreateTables"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/batchcreate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchCreateTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchCreateTables(request *BatchCreateTablesRequest) (_result *BatchCreateTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchCreateTablesResponse{}
	_body, _err := client.BatchCreateTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeletePartitionsWithOptions(request *BatchDeletePartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchDeletePartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfExists)) {
		body["IfExists"] = request.IfExists
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionValueList)) {
		body["PartitionValueList"] = request.PartitionValueList
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeletePartitions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/batchdelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeletePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeletePartitions(request *BatchDeletePartitionsRequest) (_result *BatchDeletePartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeletePartitionsResponse{}
	_body, _err := client.BatchDeletePartitionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteTableVersionsWithOptions(request *BatchDeleteTableVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchDeleteTableVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionIds)) {
		body["VersionIds"] = request.VersionIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteTableVersions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/versions/batchdelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteTableVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteTableVersions(request *BatchDeleteTableVersionsRequest) (_result *BatchDeleteTableVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteTableVersionsResponse{}
	_body, _err := client.BatchDeleteTableVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchDeleteTablesWithOptions(request *BatchDeleteTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchDeleteTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfExists)) {
		body["IfExists"] = request.IfExists
	}

	if !tea.BoolValue(util.IsUnset(request.TableNames)) {
		body["TableNames"] = request.TableNames
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchDeleteTables"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/batchdelete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchDeleteTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchDeleteTables(request *BatchDeleteTablesRequest) (_result *BatchDeleteTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchDeleteTablesResponse{}
	_body, _err := client.BatchDeleteTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetPartitionColumnStatisticsWithOptions(request *BatchGetPartitionColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetPartitionColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnNames)) {
		body["ColumnNames"] = request.ColumnNames
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNames)) {
		body["PartitionNames"] = request.PartitionNames
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetPartitionColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/columnstatistics/batchget"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetPartitionColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetPartitionColumnStatistics(request *BatchGetPartitionColumnStatisticsRequest) (_result *BatchGetPartitionColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetPartitionColumnStatisticsResponse{}
	_body, _err := client.BatchGetPartitionColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetPartitionsWithOptions(request *BatchGetPartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetPartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IsShareSd)) {
		body["IsShareSd"] = request.IsShareSd
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionValueList)) {
		body["PartitionValueList"] = request.PartitionValueList
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetPartitions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/batchget"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetPartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetPartitions(request *BatchGetPartitionsRequest) (_result *BatchGetPartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetPartitionsResponse{}
	_body, _err := client.BatchGetPartitionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetTablesWithOptions(request *BatchGetTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableNames)) {
		body["TableNames"] = request.TableNames
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGetTables"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/batchget"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGetTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetTables(request *BatchGetTablesRequest) (_result *BatchGetTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetTablesResponse{}
	_body, _err := client.BatchGetTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGrantPermissionsWithOptions(request *BatchGrantPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGrantPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRevokeEntries)) {
		body["GrantRevokeEntries"] = request.GrantRevokeEntries
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchGrantPermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/batchgrant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGrantPermissions(request *BatchGrantPermissionsRequest) (_result *BatchGrantPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGrantPermissionsResponse{}
	_body, _err := client.BatchGrantPermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchRevokePermissionsWithOptions(request *BatchRevokePermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchRevokePermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.GrantRevokeEntries)) {
		body["GrantRevokeEntries"] = request.GrantRevokeEntries
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchRevokePermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/batchrevoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchRevokePermissions(request *BatchRevokePermissionsRequest) (_result *BatchRevokePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchRevokePermissionsResponse{}
	_body, _err := client.BatchRevokePermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdatePartitionsWithOptions(request *BatchUpdatePartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchUpdatePartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionInputs)) {
		body["PartitionInputs"] = request.PartitionInputs
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdatePartitions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/batchupdate"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdatePartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdatePartitions(request *BatchUpdatePartitionsRequest) (_result *BatchUpdatePartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdatePartitionsResponse{}
	_body, _err := client.BatchUpdatePartitionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchUpdateTablesWithOptions(request *BatchUpdateTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchUpdateTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		body["IsAsync"] = request.IsAsync
	}

	if !tea.BoolValue(util.IsUnset(request.TableInputs)) {
		body["TableInputs"] = request.TableInputs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BatchUpdateTables"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/batchupdate"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchUpdateTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchUpdateTables(request *BatchUpdateTablesRequest) (_result *BatchUpdateTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchUpdateTablesResponse{}
	_body, _err := client.BatchUpdateTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelQueryWithOptions(request *CancelQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelQuery"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/query/cancelQuery"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelQuery(request *CancelQueryRequest) (_result *CancelQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelQueryResponse{}
	_body, _err := client.CancelQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckPermissionsWithOptions(request *CheckPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CheckPermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CheckPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckPermissions(request *CheckPermissionsRequest) (_result *CheckPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckPermissionsResponse{}
	_body, _err := client.CheckPermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCatalogWithOptions(request *CreateCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogInput)) {
		body["CatalogInput"] = request.CatalogInput
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCatalog"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCatalog(request *CreateCatalogRequest) (_result *CreateCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCatalogResponse{}
	_body, _err := client.CreateCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInput)) {
		body["DatabaseInput"] = request.DatabaseInput
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatabase"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDatabase(request *CreateDatabaseRequest) (_result *CreateDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionWithOptions(request *CreateFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionInput)) {
		body["FunctionInput"] = request.FunctionInput
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunction"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunction(request *CreateFunctionRequest) (_result *CreateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionResponse{}
	_body, _err := client.CreateFunctionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateLockWithOptions(request *CreateLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LockObjList)) {
		body["LockObjList"] = request.LockObjList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateLock"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/locks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateLock(request *CreateLockRequest) (_result *CreateLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateLockResponse{}
	_body, _err := client.CreateLockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePartitionWithOptions(request *CreatePartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfNotExists)) {
		body["IfNotExists"] = request.IfNotExists
	}

	if !tea.BoolValue(util.IsUnset(request.NeedResult)) {
		body["NeedResult"] = request.NeedResult
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionInput)) {
		body["PartitionInput"] = request.PartitionInput
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreatePartition"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreatePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePartition(request *CreatePartitionRequest) (_result *CreatePartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePartitionResponse{}
	_body, _err := client.CreatePartitionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateRoleWithOptions(request *CreateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateRole"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateRole(request *CreateRoleRequest) (_result *CreateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateRoleResponse{}
	_body, _err := client.CreateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTableWithOptions(request *CreateTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableInput)) {
		body["TableInput"] = request.TableInput
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTable"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTable(request *CreateTableRequest) (_result *CreateTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTableResponse{}
	_body, _err := client.CreateTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCatalogWithOptions(request *DeleteCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		query["IsAsync"] = request.IsAsync
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCatalog"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCatalog(request *DeleteCatalogRequest) (_result *DeleteCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCatalogResponse{}
	_body, _err := client.DeleteCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Async)) {
		query["Async"] = request.Async
	}

	if !tea.BoolValue(util.IsUnset(request.Cascade)) {
		query["Cascade"] = request.Cascade
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabase"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionWithOptions(request *DeleteFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunction"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunction(request *DeleteFunctionRequest) (_result *DeleteFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionResponse{}
	_body, _err := client.DeleteFunctionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePartitionWithOptions(request *DeletePartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IfExists)) {
		body["IfExists"] = request.IfExists
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionValues)) {
		body["PartitionValues"] = request.PartitionValues
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePartition"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePartition(request *DeletePartitionRequest) (_result *DeletePartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePartitionResponse{}
	_body, _err := client.DeletePartitionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePartitionColumnStatisticsWithOptions(tmpReq *DeletePartitionColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePartitionColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeletePartitionColumnStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ColumnNames)) {
		request.ColumnNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnNames, tea.String("ColumnNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PartitionNames)) {
		request.PartitionNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionNames, tea.String("PartitionNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnNamesShrink)) {
		query["ColumnNames"] = request.ColumnNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNamesShrink)) {
		query["PartitionNames"] = request.PartitionNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePartitionColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/columnstatistics"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePartitionColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePartitionColumnStatistics(request *DeletePartitionColumnStatisticsRequest) (_result *DeletePartitionColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePartitionColumnStatisticsResponse{}
	_body, _err := client.DeletePartitionColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteRoleWithOptions(request *DeleteRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRole"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteRole(request *DeleteRoleRequest) (_result *DeleteRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRoleResponse{}
	_body, _err := client.DeleteRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTableWithOptions(request *DeleteTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTable"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTable(request *DeleteTableRequest) (_result *DeleteTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableResponse{}
	_body, _err := client.DeleteTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTableColumnStatisticsWithOptions(tmpReq *DeleteTableColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DeleteTableColumnStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ColumnNames)) {
		request.ColumnNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnNames, tea.String("ColumnNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnNamesShrink)) {
		query["ColumnNames"] = request.ColumnNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTableColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/columnstatistics"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTableColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTableColumnStatistics(request *DeleteTableColumnStatisticsRequest) (_result *DeleteTableColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableColumnStatisticsResponse{}
	_body, _err := client.DeleteTableColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTableVersionWithOptions(request *DeleteTableVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTableVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTableVersion"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/versions"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTableVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTableVersion(request *DeleteTableVersionRequest) (_result *DeleteTableVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTableVersionResponse{}
	_body, _err := client.DeleteTableVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeregisterLocationWithOptions(request *DeregisterLocationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeregisterLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LocationId)) {
		query["LocationId"] = request.LocationId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeregisterLocation"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/locations"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeregisterLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeregisterLocation(request *DeregisterLocationRequest) (_result *DeregisterLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeregisterLocationResponse{}
	_body, _err := client.DeregisterLocationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/service/describeRegions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions() (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAsyncTaskStatusWithOptions(request *GetAsyncTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAsyncTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncTaskStatus"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncTaskStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAsyncTaskStatus(request *GetAsyncTaskStatusRequest) (_result *GetAsyncTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAsyncTaskStatusResponse{}
	_body, _err := client.GetAsyncTaskStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCatalogWithOptions(request *GetCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalog"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCatalog(request *GetCatalogRequest) (_result *GetCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogResponse{}
	_body, _err := client.GetCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCatalogSettingsWithOptions(request *GetCatalogSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCatalogSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetCatalogSettings"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/settings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCatalogSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCatalogSettings(request *GetCatalogSettingsRequest) (_result *GetCatalogSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCatalogSettingsResponse{}
	_body, _err := client.GetCatalogSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDatabaseWithOptions(request *GetDatabaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDatabase"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDatabase(request *GetDatabaseRequest) (_result *GetDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDatabaseResponse{}
	_body, _err := client.GetDatabaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionWithOptions(request *GetFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		query["FunctionName"] = request.FunctionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunction"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunction(request *GetFunctionRequest) (_result *GetFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionResponse{}
	_body, _err := client.GetFunctionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLifecycleRuleWithOptions(request *GetLifecycleRuleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLifecycleRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BizId)) {
		query["BizId"] = request.BizId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		query["ResourceName"] = request.ResourceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLifecycleRule"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/metastorehouse/lifecycle/rule/getLifecycleRule"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLifecycleRuleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLifecycleRule(request *GetLifecycleRuleRequest) (_result *GetLifecycleRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLifecycleRuleResponse{}
	_body, _err := client.GetLifecycleRuleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLockWithOptions(request *GetLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LockId)) {
		query["LockId"] = request.LockId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLock"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/locks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLock(request *GetLockRequest) (_result *GetLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLockResponse{}
	_body, _err := client.GetLockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPartitionWithOptions(request *GetPartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionValues)) {
		body["PartitionValues"] = request.PartitionValues
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPartition"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPartition(request *GetPartitionRequest) (_result *GetPartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPartitionResponse{}
	_body, _err := client.GetPartitionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPartitionColumnStatisticsWithOptions(tmpReq *GetPartitionColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPartitionColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetPartitionColumnStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ColumnNames)) {
		request.ColumnNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnNames, tea.String("ColumnNames"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.PartitionNames)) {
		request.PartitionNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.PartitionNames, tea.String("PartitionNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnNamesShrink)) {
		query["ColumnNames"] = request.ColumnNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionNamesShrink)) {
		query["PartitionNames"] = request.PartitionNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPartitionColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/columnstatistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPartitionColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPartitionColumnStatistics(request *GetPartitionColumnStatisticsRequest) (_result *GetPartitionColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPartitionColumnStatisticsResponse{}
	_body, _err := client.GetPartitionColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQueryResultWithOptions(request *GetQueryResultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQueryResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.QueryId)) {
		query["QueryId"] = request.QueryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetQueryResult"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/query/getQueryResult"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetQueryResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQueryResult(request *GetQueryResultRequest) (_result *GetQueryResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQueryResultResponse{}
	_body, _err := client.GetQueryResultWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRegionStatusWithOptions(request *GetRegionStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRegionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRegionStatus"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/service/getRegionStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRegionStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRegionStatus(request *GetRegionStatusRequest) (_result *GetRegionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRegionStatusResponse{}
	_body, _err := client.GetRegionStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetRoleWithOptions(request *GetRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRole"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetRole(request *GetRoleRequest) (_result *GetRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRoleResponse{}
	_body, _err := client.GetRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetServiceStatusWithOptions(request *GetServiceStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetServiceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetServiceStatus"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/service/getServiceStatus"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetServiceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetServiceStatus(request *GetServiceStatusRequest) (_result *GetServiceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetServiceStatusResponse{}
	_body, _err := client.GetServiceStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableWithOptions(request *GetTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTable"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTable(request *GetTableRequest) (_result *GetTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableResponse{}
	_body, _err := client.GetTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableColumnStatisticsWithOptions(tmpReq *GetTableColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &GetTableColumnStatisticsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ColumnNames)) {
		request.ColumnNamesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ColumnNames, tea.String("ColumnNames"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.ColumnNamesShrink)) {
		query["ColumnNames"] = request.ColumnNamesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/columnstatistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableColumnStatistics(request *GetTableColumnStatisticsRequest) (_result *GetTableColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableColumnStatisticsResponse{}
	_body, _err := client.GetTableColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableProfileWithOptions(request *GetTableProfileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableProfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableProfile"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/metastorehouse/catalog/database/tableprofile"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableProfileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableProfile(request *GetTableProfileRequest) (_result *GetTableProfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableProfileResponse{}
	_body, _err := client.GetTableProfileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTableVersionWithOptions(request *GetTableVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTableVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["VersionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTableVersion"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTableVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTableVersion(request *GetTableVersionRequest) (_result *GetTableVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTableVersionResponse{}
	_body, _err := client.GetTableVersionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantPermissionsWithOptions(request *GrantPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GrantPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accesses)) {
		body["Accesses"] = request.Accesses
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DelegateAccesses)) {
		body["DelegateAccesses"] = request.DelegateAccesses
	}

	if !tea.BoolValue(util.IsUnset(request.MetaResource)) {
		body["MetaResource"] = request.MetaResource
	}

	if !tea.BoolValue(util.IsUnset(request.Principal)) {
		body["Principal"] = request.Principal
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantPermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/grant"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantPermissions(request *GrantPermissionsRequest) (_result *GrantPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantPermissionsResponse{}
	_body, _err := client.GrantPermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantRoleToUsersWithOptions(request *GrantRoleToUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GrantRoleToUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantRoleToUsers"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/grantusers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantRoleToUsers(request *GrantRoleToUsersRequest) (_result *GrantRoleToUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantRoleToUsersResponse{}
	_body, _err := client.GrantRoleToUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GrantRolesToUserWithOptions(request *GrantRolesToUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GrantRolesToUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleNames)) {
		body["RoleNames"] = request.RoleNames
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		body["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GrantRolesToUser"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/grantroles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GrantRolesToUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GrantRolesToUser(request *GrantRolesToUserRequest) (_result *GrantRolesToUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GrantRolesToUserResponse{}
	_body, _err := client.GrantRolesToUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCatalogsWithOptions(request *ListCatalogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCatalogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdPattern)) {
		query["IdPattern"] = request.IdPattern
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCatalogs"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCatalogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCatalogs(request *ListCatalogsRequest) (_result *ListCatalogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCatalogsResponse{}
	_body, _err := client.ListCatalogsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDatabasesWithOptions(request *ListDatabasesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.NamePattern)) {
		query["NamePattern"] = request.NamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDatabases"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDatabases(request *ListDatabasesRequest) (_result *ListDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDatabasesResponse{}
	_body, _err := client.ListDatabasesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionNamesWithOptions(request *ListFunctionNamesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNamePattern)) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionNames"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions/names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionNames(request *ListFunctionNamesRequest) (_result *ListFunctionNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionNamesResponse{}
	_body, _err := client.ListFunctionNamesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionsWithOptions(request *ListFunctionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionNamePattern)) {
		query["FunctionNamePattern"] = request.FunctionNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctions(request *ListFunctionsRequest) (_result *ListFunctionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionsResponse{}
	_body, _err := client.ListFunctionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartitionNamesWithOptions(request *ListPartitionNamesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartitionNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		body["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PartialPartValues)) {
		body["PartialPartValues"] = request.PartialPartValues
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPartitionNames"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/names"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPartitionNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPartitionNames(request *ListPartitionNamesRequest) (_result *ListPartitionNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionNamesResponse{}
	_body, _err := client.ListPartitionNamesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartitionsWithOptions(request *ListPartitionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartitionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IsShareSd)) {
		body["IsShareSd"] = request.IsShareSd
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		body["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PartialPartValues)) {
		body["PartialPartValues"] = request.PartialPartValues
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPartitions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPartitionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPartitions(request *ListPartitionsRequest) (_result *ListPartitionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionsResponse{}
	_body, _err := client.ListPartitionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartitionsByExprWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartitionsByExprResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListPartitionsByExpr"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/listbyexpr"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ListPartitionsByExprResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPartitionsByExpr() (_result *ListPartitionsByExprResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionsByExprResponse{}
	_body, _err := client.ListPartitionsByExprWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartitionsByFilterWithOptions(request *ListPartitionsByFilterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartitionsByFilterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		body["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.IsShareSd)) {
		body["IsShareSd"] = request.IsShareSd
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		body["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPartitionsByFilter"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/listbyfilter"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPartitionsByFilterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPartitionsByFilter(request *ListPartitionsByFilterRequest) (_result *ListPartitionsByFilterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartitionsByFilterResponse{}
	_body, _err := client.ListPartitionsByFilterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPermissionsWithOptions(request *ListPermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.IsListUserRolePermissions)) {
		body["IsListUserRolePermissions"] = request.IsListUserRolePermissions
	}

	if !tea.BoolValue(util.IsUnset(request.MetaResource)) {
		body["MetaResource"] = request.MetaResource
	}

	if !tea.BoolValue(util.IsUnset(request.MetaResourceType)) {
		body["MetaResourceType"] = request.MetaResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		body["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Principal)) {
		body["Principal"] = request.Principal
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListPermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListPermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPermissions(request *ListPermissionsRequest) (_result *ListPermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPermissionsResponse{}
	_body, _err := client.ListPermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRoleUsersWithOptions(request *ListRoleUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		query["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.UserNamePattern)) {
		query["UserNamePattern"] = request.UserNamePattern
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoleUsers"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/roleusers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoleUsers(request *ListRoleUsersRequest) (_result *ListRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRoleUsersResponse{}
	_body, _err := client.ListRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRolesWithOptions(request *ListRolesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.RoleNamePattern)) {
		query["RoleNamePattern"] = request.RoleNamePattern
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRoles"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRoles(request *ListRolesRequest) (_result *ListRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRolesResponse{}
	_body, _err := client.ListRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTableNamesWithOptions(request *ListTableNamesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTableNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TableNamePattern)) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTableNames"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTableNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTableNames(request *ListTableNamesRequest) (_result *ListTableNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableNamesResponse{}
	_body, _err := client.ListTableNamesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTableVersionsWithOptions(request *ListTableVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTableVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		query["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTableVersions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/versions/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTableVersionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTableVersions(request *ListTableVersionsRequest) (_result *ListTableVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTableVersionsResponse{}
	_body, _err := client.ListTableVersionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTablesWithOptions(request *ListTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		query["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		query["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.TableNamePattern)) {
		query["TableNamePattern"] = request.TableNamePattern
	}

	if !tea.BoolValue(util.IsUnset(request.TableType)) {
		query["TableType"] = request.TableType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTables"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/databases/tables/list"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTables(request *ListTablesRequest) (_result *ListTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTablesResponse{}
	_body, _err := client.ListTablesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserRolesWithOptions(request *ListUserRolesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextPageToken)) {
		query["NextPageToken"] = request.NextPageToken
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.PrincipalArn)) {
		query["PrincipalArn"] = request.PrincipalArn
	}

	if !tea.BoolValue(util.IsUnset(request.RoleNamePattern)) {
		query["RoleNamePattern"] = request.RoleNamePattern
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserRoles"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/userroles"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserRoles(request *ListUserRolesRequest) (_result *ListUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserRolesResponse{}
	_body, _err := client.ListUserRolesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RefreshLockWithOptions(request *RefreshLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RefreshLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LockId)) {
		query["LockId"] = request.LockId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RefreshLock"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/locks"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RefreshLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RefreshLock(request *RefreshLockRequest) (_result *RefreshLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RefreshLockResponse{}
	_body, _err := client.RefreshLockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterLocationWithOptions(request *RegisterLocationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RegisterLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InventoryCollectEnabled)) {
		body["InventoryCollectEnabled"] = request.InventoryCollectEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.Location)) {
		body["Location"] = request.Location
	}

	if !tea.BoolValue(util.IsUnset(request.OssLogCollectEnabled)) {
		body["OssLogCollectEnabled"] = request.OssLogCollectEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RegisterLocation"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/locations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RegisterLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterLocation(request *RegisterLocationRequest) (_result *RegisterLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RegisterLocationResponse{}
	_body, _err := client.RegisterLocationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenamePartitionWithOptions(request *RenamePartitionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenamePartitionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionInput)) {
		body["PartitionInput"] = request.PartitionInput
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionValues)) {
		body["PartitionValues"] = request.PartitionValues
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenamePartition"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/rename"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenamePartitionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenamePartition(request *RenamePartitionRequest) (_result *RenamePartitionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenamePartitionResponse{}
	_body, _err := client.RenamePartitionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameTableWithOptions(request *RenameTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenameTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		body["IsAsync"] = request.IsAsync
	}

	if !tea.BoolValue(util.IsUnset(request.TableInput)) {
		body["TableInput"] = request.TableInput
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenameTable"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/rename"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenameTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameTable(request *RenameTableRequest) (_result *RenameTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenameTableResponse{}
	_body, _err := client.RenameTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokePermissionsWithOptions(request *RevokePermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokePermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accesses)) {
		body["Accesses"] = request.Accesses
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DelegateAccesses)) {
		body["DelegateAccesses"] = request.DelegateAccesses
	}

	if !tea.BoolValue(util.IsUnset(request.MetaResource)) {
		body["MetaResource"] = request.MetaResource
	}

	if !tea.BoolValue(util.IsUnset(request.Principal)) {
		body["Principal"] = request.Principal
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokePermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/revoke"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokePermissions(request *RevokePermissionsRequest) (_result *RevokePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokePermissionsResponse{}
	_body, _err := client.RevokePermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeRoleFromUsersWithOptions(request *RevokeRoleFromUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeRoleFromUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeRoleFromUsers"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/revokeusers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeRoleFromUsers(request *RevokeRoleFromUsersRequest) (_result *RevokeRoleFromUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeRoleFromUsersResponse{}
	_body, _err := client.RevokeRoleFromUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RevokeRolesFromUserWithOptions(request *RevokeRolesFromUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RevokeRolesFromUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleNames)) {
		body["RoleNames"] = request.RoleNames
	}

	if !tea.BoolValue(util.IsUnset(request.User)) {
		body["User"] = request.User
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RevokeRolesFromUser"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles/revokeroles"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RevokeRolesFromUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RevokeRolesFromUser(request *RevokeRolesFromUserRequest) (_result *RevokeRolesFromUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RevokeRolesFromUserResponse{}
	_body, _err := client.RevokeRolesFromUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunMigrationWorkflowWithOptions(request *RunMigrationWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RunMigrationWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunMigrationWorkflow"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/migration/workflow/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RunMigrationWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunMigrationWorkflow(request *RunMigrationWorkflowRequest) (_result *RunMigrationWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RunMigrationWorkflowResponse{}
	_body, _err := client.RunMigrationWorkflowWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchWithOptions(request *SearchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchText)) {
		body["SearchText"] = request.SearchText
	}

	if !tea.BoolValue(util.IsUnset(request.SearchType)) {
		body["SearchType"] = request.SearchType
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriteria)) {
		body["SortCriteria"] = request.SortCriteria
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Search"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Search(request *SearchRequest) (_result *SearchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchResponse{}
	_body, _err := client.SearchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchAcrossCatalogWithOptions(request *SearchAcrossCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchAcrossCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogIds)) {
		body["CatalogIds"] = request.CatalogIds
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchText)) {
		body["SearchText"] = request.SearchText
	}

	if !tea.BoolValue(util.IsUnset(request.SearchTypes)) {
		body["SearchTypes"] = request.SearchTypes
	}

	if !tea.BoolValue(util.IsUnset(request.SortCriteria)) {
		body["SortCriteria"] = request.SortCriteria
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAcrossCatalog"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/search/search-across-catalog"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAcrossCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchAcrossCatalog(request *SearchAcrossCatalogRequest) (_result *SearchAcrossCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchAcrossCatalogResponse{}
	_body, _err := client.SearchAcrossCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMigrationWorkflowWithOptions(request *StopMigrationWorkflowRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopMigrationWorkflowResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMigrationWorkflow"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/migration/workflow/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMigrationWorkflowResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMigrationWorkflow(request *StopMigrationWorkflowRequest) (_result *StopMigrationWorkflowResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopMigrationWorkflowResponse{}
	_body, _err := client.StopMigrationWorkflowWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SubmitQueryWithOptions(request *SubmitQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SubmitQueryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["catalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.Sql)) {
		body["sql"] = request.Sql
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["workspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SubmitQuery"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/query/submitQueryRequestBody"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SubmitQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SubmitQuery(request *SubmitQueryRequest) (_result *SubmitQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SubmitQueryResponse{}
	_body, _err := client.SubmitQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnLockWithOptions(request *UnLockRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnLockResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LockId)) {
		query["LockId"] = request.LockId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UnLock"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/locks"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnLockResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnLock(request *UnLockRequest) (_result *UnLockResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnLockResponse{}
	_body, _err := client.UnLockWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCatalogWithOptions(request *UpdateCatalogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCatalogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogInput)) {
		body["CatalogInput"] = request.CatalogInput
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCatalog"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCatalogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCatalog(request *UpdateCatalogRequest) (_result *UpdateCatalogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCatalogResponse{}
	_body, _err := client.UpdateCatalogWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateCatalogSettingsWithOptions(request *UpdateCatalogSettingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCatalogSettingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogSettings)) {
		body["CatalogSettings"] = request.CatalogSettings
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateCatalogSettings"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/settings"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateCatalogSettingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateCatalogSettings(request *UpdateCatalogSettingsRequest) (_result *UpdateCatalogSettingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCatalogSettingsResponse{}
	_body, _err := client.UpdateCatalogSettingsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateDatabaseWithOptions(request *UpdateDatabaseRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseInput)) {
		body["DatabaseInput"] = request.DatabaseInput
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDatabase"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDatabaseResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateDatabase(request *UpdateDatabaseRequest) (_result *UpdateDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDatabaseResponse{}
	_body, _err := client.UpdateDatabaseWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionWithOptions(request *UpdateFunctionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionInput)) {
		body["FunctionInput"] = request.FunctionInput
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionName)) {
		body["FunctionName"] = request.FunctionName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunction"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/functions"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunction(request *UpdateFunctionRequest) (_result *UpdateFunctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionResponse{}
	_body, _err := client.UpdateFunctionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePartitionColumnStatisticsWithOptions(request *UpdatePartitionColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePartitionColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.UpdateTablePartitionColumnStatisticsRequest),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePartitionColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/partitions/columnstatistics"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePartitionColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePartitionColumnStatistics(request *UpdatePartitionColumnStatisticsRequest) (_result *UpdatePartitionColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePartitionColumnStatisticsResponse{}
	_body, _err := client.UpdatePartitionColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdatePermissionsWithOptions(request *UpdatePermissionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdatePermissionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Accesses)) {
		body["Accesses"] = request.Accesses
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DelegateAccesses)) {
		body["DelegateAccesses"] = request.DelegateAccesses
	}

	if !tea.BoolValue(util.IsUnset(request.MetaResource)) {
		body["MetaResource"] = request.MetaResource
	}

	if !tea.BoolValue(util.IsUnset(request.Principal)) {
		body["Principal"] = request.Principal
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdatePermissions"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/permissions/"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdatePermissionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdatePermissions(request *UpdatePermissionsRequest) (_result *UpdatePermissionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdatePermissionsResponse{}
	_body, _err := client.UpdatePermissionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRegisteredLocationWithOptions(request *UpdateRegisteredLocationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRegisteredLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InventoryCollectEnabled)) {
		body["InventoryCollectEnabled"] = request.InventoryCollectEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.LocationId)) {
		body["LocationId"] = request.LocationId
	}

	if !tea.BoolValue(util.IsUnset(request.OssLogCollectEnabled)) {
		body["OssLogCollectEnabled"] = request.OssLogCollectEnabled
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRegisteredLocation"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/webapi/locations"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRegisteredLocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRegisteredLocation(request *UpdateRegisteredLocationRequest) (_result *UpdateRegisteredLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRegisteredLocationResponse{}
	_body, _err := client.UpdateRegisteredLocationWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleWithOptions(request *UpdateRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleInput)) {
		body["RoleInput"] = request.RoleInput
	}

	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRole"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/roles"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoleResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRole(request *UpdateRoleRequest) (_result *UpdateRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleResponse{}
	_body, _err := client.UpdateRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateRoleUsersWithOptions(request *UpdateRoleUsersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRoleUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RoleName)) {
		body["RoleName"] = request.RoleName
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRoleUsers"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/auth/updateroleusers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateRoleUsers(request *UpdateRoleUsersRequest) (_result *UpdateRoleUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRoleUsersResponse{}
	_body, _err := client.UpdateRoleUsersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTableWithOptions(request *UpdateTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllowPartitionKeyChange)) {
		body["AllowPartitionKeyChange"] = request.AllowPartitionKeyChange
	}

	if !tea.BoolValue(util.IsUnset(request.CatalogId)) {
		body["CatalogId"] = request.CatalogId
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.IsAsync)) {
		body["IsAsync"] = request.IsAsync
	}

	if !tea.BoolValue(util.IsUnset(request.SkipArchive)) {
		body["SkipArchive"] = request.SkipArchive
	}

	if !tea.BoolValue(util.IsUnset(request.TableInput)) {
		body["TableInput"] = request.TableInput
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTable"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTable(request *UpdateTableRequest) (_result *UpdateTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableResponse{}
	_body, _err := client.UpdateTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTableColumnStatisticsWithOptions(request *UpdateTableColumnStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTableColumnStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.UpdateTablePartitionColumnStatisticsRequest),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTableColumnStatistics"),
		Version:     tea.String("2020-07-10"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/metastore/catalogs/databases/tables/columnstatistics"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTableColumnStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTableColumnStatistics(request *UpdateTableColumnStatisticsRequest) (_result *UpdateTableColumnStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTableColumnStatisticsResponse{}
	_body, _err := client.UpdateTableColumnStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
