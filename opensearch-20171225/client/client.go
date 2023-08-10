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

type ABTestExperiment struct {
	Name         *string            `json:"name,omitempty" xml:"name,omitempty"`
	Online       *bool              `json:"online,omitempty" xml:"online,omitempty"`
	Params       map[string]*string `json:"params,omitempty" xml:"params,omitempty"`
	SerialNumber *int32             `json:"serialNumber,omitempty" xml:"serialNumber,omitempty"`
	Traffic      *int32             `json:"traffic,omitempty" xml:"traffic,omitempty"`
}

func (s ABTestExperiment) String() string {
	return tea.Prettify(s)
}

func (s ABTestExperiment) GoString() string {
	return s.String()
}

func (s *ABTestExperiment) SetName(v string) *ABTestExperiment {
	s.Name = &v
	return s
}

func (s *ABTestExperiment) SetOnline(v bool) *ABTestExperiment {
	s.Online = &v
	return s
}

func (s *ABTestExperiment) SetParams(v map[string]*string) *ABTestExperiment {
	s.Params = v
	return s
}

func (s *ABTestExperiment) SetSerialNumber(v int32) *ABTestExperiment {
	s.SerialNumber = &v
	return s
}

func (s *ABTestExperiment) SetTraffic(v int32) *ABTestExperiment {
	s.Traffic = &v
	return s
}

type ABTestGroup struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *int32  `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ABTestGroup) String() string {
	return tea.Prettify(s)
}

func (s ABTestGroup) GoString() string {
	return s.String()
}

func (s *ABTestGroup) SetName(v string) *ABTestGroup {
	s.Name = &v
	return s
}

func (s *ABTestGroup) SetStatus(v int32) *ABTestGroup {
	s.Status = &v
	return s
}

type ABTestScene struct {
	Name   *string   `json:"name,omitempty" xml:"name,omitempty"`
	Status *int32    `json:"status,omitempty" xml:"status,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ABTestScene) String() string {
	return tea.Prettify(s)
}

func (s ABTestScene) GoString() string {
	return s.String()
}

func (s *ABTestScene) SetName(v string) *ABTestScene {
	s.Name = &v
	return s
}

func (s *ABTestScene) SetStatus(v int32) *ABTestScene {
	s.Status = &v
	return s
}

func (s *ABTestScene) SetValues(v []*string) *ABTestScene {
	s.Values = v
	return s
}

type App struct {
	AutoSwitch      *bool             `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster         *AppCluster       `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	DataSources     []*DataSource     `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	Description     *string           `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *Domain           `json:"domain,omitempty" xml:"domain,omitempty"`
	FetchFields     []*string         `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks      []*FirstRank      `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	NetworkType     *string           `json:"networkType,omitempty" xml:"networkType,omitempty"`
	QueryProcessors []*QueryProcessor `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	Quota           *Quota            `json:"quota,omitempty" xml:"quota,omitempty"`
	RealtimeShared  *bool             `json:"realtimeShared,omitempty" xml:"realtimeShared,omitempty"`
	Schema          *Schema           `json:"schema,omitempty" xml:"schema,omitempty"`
	Schemas         []*Schema         `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks     []*SecondRank     `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	Summaries       []*Summary        `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	Type            *string           `json:"type,omitempty" xml:"type,omitempty"`
}

func (s App) String() string {
	return tea.Prettify(s)
}

func (s App) GoString() string {
	return s.String()
}

func (s *App) SetAutoSwitch(v bool) *App {
	s.AutoSwitch = &v
	return s
}

func (s *App) SetCluster(v *AppCluster) *App {
	s.Cluster = v
	return s
}

func (s *App) SetDataSources(v []*DataSource) *App {
	s.DataSources = v
	return s
}

func (s *App) SetDescription(v string) *App {
	s.Description = &v
	return s
}

func (s *App) SetDomain(v *Domain) *App {
	s.Domain = v
	return s
}

func (s *App) SetFetchFields(v []*string) *App {
	s.FetchFields = v
	return s
}

func (s *App) SetFirstRanks(v []*FirstRank) *App {
	s.FirstRanks = v
	return s
}

func (s *App) SetNetworkType(v string) *App {
	s.NetworkType = &v
	return s
}

func (s *App) SetQueryProcessors(v []*QueryProcessor) *App {
	s.QueryProcessors = v
	return s
}

func (s *App) SetQuota(v *Quota) *App {
	s.Quota = v
	return s
}

func (s *App) SetRealtimeShared(v bool) *App {
	s.RealtimeShared = &v
	return s
}

func (s *App) SetSchema(v *Schema) *App {
	s.Schema = v
	return s
}

func (s *App) SetSchemas(v []*Schema) *App {
	s.Schemas = v
	return s
}

func (s *App) SetSecondRanks(v []*SecondRank) *App {
	s.SecondRanks = v
	return s
}

func (s *App) SetSummaries(v []*Summary) *App {
	s.Summaries = v
	return s
}

func (s *App) SetType(v string) *App {
	s.Type = &v
	return s
}

type AppCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s AppCluster) String() string {
	return tea.Prettify(s)
}

func (s AppCluster) GoString() string {
	return s.String()
}

func (s *AppCluster) SetMaxQueryClauseLength(v int32) *AppCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *AppCluster) SetMaxTimeoutMS(v int32) *AppCluster {
	s.MaxTimeoutMS = &v
	return s
}

type AppGroup struct {
	ChargeType  *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	ChargingWay *string `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	Quota       *Quota  `json:"quota,omitempty" xml:"quota,omitempty"`
	Type        *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s AppGroup) String() string {
	return tea.Prettify(s)
}

func (s AppGroup) GoString() string {
	return s.String()
}

func (s *AppGroup) SetChargeType(v string) *AppGroup {
	s.ChargeType = &v
	return s
}

func (s *AppGroup) SetChargingWay(v string) *AppGroup {
	s.ChargingWay = &v
	return s
}

func (s *AppGroup) SetDescription(v string) *AppGroup {
	s.Description = &v
	return s
}

func (s *AppGroup) SetDomain(v string) *AppGroup {
	s.Domain = &v
	return s
}

func (s *AppGroup) SetName(v string) *AppGroup {
	s.Name = &v
	return s
}

func (s *AppGroup) SetQuota(v *Quota) *AppGroup {
	s.Quota = v
	return s
}

func (s *AppGroup) SetType(v string) *AppGroup {
	s.Type = &v
	return s
}

type DataSource struct {
	Fields     []map[string]*string               `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                            `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}             `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]*DataSourcePluginsValue `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                            `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                            `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                            `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DataSource) String() string {
	return tea.Prettify(s)
}

func (s DataSource) GoString() string {
	return s.String()
}

func (s *DataSource) SetFields(v []map[string]*string) *DataSource {
	s.Fields = v
	return s
}

func (s *DataSource) SetKeyField(v string) *DataSource {
	s.KeyField = &v
	return s
}

func (s *DataSource) SetParameters(v map[string]interface{}) *DataSource {
	s.Parameters = v
	return s
}

func (s *DataSource) SetPlugins(v map[string]*DataSourcePluginsValue) *DataSource {
	s.Plugins = v
	return s
}

func (s *DataSource) SetSchemaName(v string) *DataSource {
	s.SchemaName = &v
	return s
}

func (s *DataSource) SetTableName(v string) *DataSource {
	s.TableName = &v
	return s
}

func (s *DataSource) SetType(v string) *DataSource {
	s.Type = &v
	return s
}

type Domain struct {
	Category  *string              `json:"category,omitempty" xml:"category,omitempty"`
	Functions map[string][]*string `json:"functions,omitempty" xml:"functions,omitempty"`
	Name      *string              `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Domain) String() string {
	return tea.Prettify(s)
}

func (s Domain) GoString() string {
	return s.String()
}

func (s *Domain) SetCategory(v string) *Domain {
	s.Category = &v
	return s
}

func (s *Domain) SetFunctions(v map[string][]*string) *Domain {
	s.Functions = v
	return s
}

func (s *Domain) SetName(v string) *Domain {
	s.Name = &v
	return s
}

type FirstRank struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FirstRank) String() string {
	return tea.Prettify(s)
}

func (s FirstRank) GoString() string {
	return s.String()
}

func (s *FirstRank) SetActive(v bool) *FirstRank {
	s.Active = &v
	return s
}

func (s *FirstRank) SetDescription(v string) *FirstRank {
	s.Description = &v
	return s
}

func (s *FirstRank) SetMeta(v interface{}) *FirstRank {
	s.Meta = v
	return s
}

func (s *FirstRank) SetName(v string) *FirstRank {
	s.Name = &v
	return s
}

func (s *FirstRank) SetType(v string) *FirstRank {
	s.Type = &v
	return s
}

type PrepayOrderInfo struct {
	AutoRenew    *bool   `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	Duration     *int32  `json:"duration,omitempty" xml:"duration,omitempty"`
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s PrepayOrderInfo) String() string {
	return tea.Prettify(s)
}

func (s PrepayOrderInfo) GoString() string {
	return s.String()
}

func (s *PrepayOrderInfo) SetAutoRenew(v bool) *PrepayOrderInfo {
	s.AutoRenew = &v
	return s
}

func (s *PrepayOrderInfo) SetDuration(v int32) *PrepayOrderInfo {
	s.Duration = &v
	return s
}

func (s *PrepayOrderInfo) SetPricingCycle(v string) *PrepayOrderInfo {
	s.PricingCycle = &v
	return s
}

type QueryProcessor struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s QueryProcessor) String() string {
	return tea.Prettify(s)
}

func (s QueryProcessor) GoString() string {
	return s.String()
}

func (s *QueryProcessor) SetActive(v bool) *QueryProcessor {
	s.Active = &v
	return s
}

func (s *QueryProcessor) SetCategory(v string) *QueryProcessor {
	s.Category = &v
	return s
}

func (s *QueryProcessor) SetDomain(v string) *QueryProcessor {
	s.Domain = &v
	return s
}

func (s *QueryProcessor) SetIndexes(v []*string) *QueryProcessor {
	s.Indexes = v
	return s
}

func (s *QueryProcessor) SetName(v string) *QueryProcessor {
	s.Name = &v
	return s
}

func (s *QueryProcessor) SetProcessors(v []map[string]interface{}) *QueryProcessor {
	s.Processors = v
	return s
}

type Quota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	OrderType       *string `json:"orderType,omitempty" xml:"orderType,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s Quota) String() string {
	return tea.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) SetComputeResource(v int32) *Quota {
	s.ComputeResource = &v
	return s
}

func (s *Quota) SetDocSize(v int32) *Quota {
	s.DocSize = &v
	return s
}

func (s *Quota) SetOrderType(v string) *Quota {
	s.OrderType = &v
	return s
}

func (s *Quota) SetSpec(v string) *Quota {
	s.Spec = &v
	return s
}

type ScheduledTask struct {
	AutoSwitch  *bool                `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cron        *string              `json:"cron,omitempty" xml:"cron,omitempty"`
	Enabled     *bool                `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Filter      *ScheduledTaskFilter `json:"filter,omitempty" xml:"filter,omitempty" type:"Struct"`
	ForkedAppId *string              `json:"forkedAppId,omitempty" xml:"forkedAppId,omitempty"`
	Permanent   *bool                `json:"permanent,omitempty" xml:"permanent,omitempty"`
	RunNow      *bool                `json:"runNow,omitempty" xml:"runNow,omitempty"`
	Type        *string              `json:"type,omitempty" xml:"type,omitempty"`
	Version     *string              `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ScheduledTask) String() string {
	return tea.Prettify(s)
}

func (s ScheduledTask) GoString() string {
	return s.String()
}

func (s *ScheduledTask) SetAutoSwitch(v bool) *ScheduledTask {
	s.AutoSwitch = &v
	return s
}

func (s *ScheduledTask) SetCron(v string) *ScheduledTask {
	s.Cron = &v
	return s
}

func (s *ScheduledTask) SetEnabled(v bool) *ScheduledTask {
	s.Enabled = &v
	return s
}

func (s *ScheduledTask) SetFilter(v *ScheduledTaskFilter) *ScheduledTask {
	s.Filter = v
	return s
}

func (s *ScheduledTask) SetForkedAppId(v string) *ScheduledTask {
	s.ForkedAppId = &v
	return s
}

func (s *ScheduledTask) SetPermanent(v bool) *ScheduledTask {
	s.Permanent = &v
	return s
}

func (s *ScheduledTask) SetRunNow(v bool) *ScheduledTask {
	s.RunNow = &v
	return s
}

func (s *ScheduledTask) SetType(v string) *ScheduledTask {
	s.Type = &v
	return s
}

func (s *ScheduledTask) SetVersion(v string) *ScheduledTask {
	s.Version = &v
	return s
}

type ScheduledTaskFilter struct {
	Days       *int32  `json:"days,omitempty" xml:"days,omitempty"`
	Expression *string `json:"expression,omitempty" xml:"expression,omitempty"`
	Field      *string `json:"field,omitempty" xml:"field,omitempty"`
	Unit       *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s ScheduledTaskFilter) String() string {
	return tea.Prettify(s)
}

func (s ScheduledTaskFilter) GoString() string {
	return s.String()
}

func (s *ScheduledTaskFilter) SetDays(v int32) *ScheduledTaskFilter {
	s.Days = &v
	return s
}

func (s *ScheduledTaskFilter) SetExpression(v string) *ScheduledTaskFilter {
	s.Expression = &v
	return s
}

func (s *ScheduledTaskFilter) SetField(v string) *ScheduledTaskFilter {
	s.Field = &v
	return s
}

func (s *ScheduledTaskFilter) SetUnit(v string) *ScheduledTaskFilter {
	s.Unit = &v
	return s
}

type Schema struct {
	IndexSortConfig  []*SchemaIndexSortConfig      `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *SchemaIndexes                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                       `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                       `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                     `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                       `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]*SchemaTablesValue `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *SchemaTtlField               `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s Schema) String() string {
	return tea.Prettify(s)
}

func (s Schema) GoString() string {
	return s.String()
}

func (s *Schema) SetIndexSortConfig(v []*SchemaIndexSortConfig) *Schema {
	s.IndexSortConfig = v
	return s
}

func (s *Schema) SetIndexes(v *SchemaIndexes) *Schema {
	s.Indexes = v
	return s
}

func (s *Schema) SetName(v string) *Schema {
	s.Name = &v
	return s
}

func (s *Schema) SetRouteField(v string) *Schema {
	s.RouteField = &v
	return s
}

func (s *Schema) SetRouteFieldValues(v []*string) *Schema {
	s.RouteFieldValues = v
	return s
}

func (s *Schema) SetSecondRouteField(v string) *Schema {
	s.SecondRouteField = &v
	return s
}

func (s *Schema) SetTables(v map[string]*SchemaTablesValue) *Schema {
	s.Tables = v
	return s
}

func (s *Schema) SetTtlField(v *SchemaTtlField) *Schema {
	s.TtlField = v
	return s
}

type SchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s SchemaIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s SchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *SchemaIndexSortConfig) SetDirection(v string) *SchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *SchemaIndexSortConfig) SetField(v string) *SchemaIndexSortConfig {
	s.Field = &v
	return s
}

type SchemaIndexes struct {
	FilterFields []*string                                  `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]*SchemaIndexesSearchFieldsValue `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s SchemaIndexes) String() string {
	return tea.Prettify(s)
}

func (s SchemaIndexes) GoString() string {
	return s.String()
}

func (s *SchemaIndexes) SetFilterFields(v []*string) *SchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *SchemaIndexes) SetSearchFields(v map[string]*SchemaIndexesSearchFieldsValue) *SchemaIndexes {
	s.SearchFields = v
	return s
}

type SchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s SchemaTtlField) String() string {
	return tea.Prettify(s)
}

func (s SchemaTtlField) GoString() string {
	return s.String()
}

func (s *SchemaTtlField) SetName(v string) *SchemaTtlField {
	s.Name = &v
	return s
}

func (s *SchemaTtlField) SetTtl(v int64) *SchemaTtlField {
	s.Ttl = &v
	return s
}

type SearchStrategy struct {
	Description   *string                        `json:"description,omitempty" xml:"description,omitempty"`
	IsDefault     *bool                          `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	MergeConfig   *SearchStrategyMergeConfig     `json:"mergeConfig,omitempty" xml:"mergeConfig,omitempty" type:"Struct"`
	Name          *string                        `json:"name,omitempty" xml:"name,omitempty"`
	SearchConfigs []*SearchStrategySearchConfigs `json:"searchConfigs,omitempty" xml:"searchConfigs,omitempty" type:"Repeated"`
}

func (s SearchStrategy) String() string {
	return tea.Prettify(s)
}

func (s SearchStrategy) GoString() string {
	return s.String()
}

func (s *SearchStrategy) SetDescription(v string) *SearchStrategy {
	s.Description = &v
	return s
}

func (s *SearchStrategy) SetIsDefault(v bool) *SearchStrategy {
	s.IsDefault = &v
	return s
}

func (s *SearchStrategy) SetMergeConfig(v *SearchStrategyMergeConfig) *SearchStrategy {
	s.MergeConfig = v
	return s
}

func (s *SearchStrategy) SetName(v string) *SearchStrategy {
	s.Name = &v
	return s
}

func (s *SearchStrategy) SetSearchConfigs(v []*SearchStrategySearchConfigs) *SearchStrategy {
	s.SearchConfigs = v
	return s
}

type SearchStrategyMergeConfig struct {
	DocCount *int32  `json:"docCount,omitempty" xml:"docCount,omitempty"`
	RankName *string `json:"rankName,omitempty" xml:"rankName,omitempty"`
}

func (s SearchStrategyMergeConfig) String() string {
	return tea.Prettify(s)
}

func (s SearchStrategyMergeConfig) GoString() string {
	return s.String()
}

func (s *SearchStrategyMergeConfig) SetDocCount(v int32) *SearchStrategyMergeConfig {
	s.DocCount = &v
	return s
}

func (s *SearchStrategyMergeConfig) SetRankName(v string) *SearchStrategyMergeConfig {
	s.RankName = &v
	return s
}

type SearchStrategySearchConfigs struct {
	FirstRankName   *string `json:"firstRankName,omitempty" xml:"firstRankName,omitempty"`
	MergeProportion *int32  `json:"mergeProportion,omitempty" xml:"mergeProportion,omitempty"`
	QueryType       *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	SecondRankName  *string `json:"secondRankName,omitempty" xml:"secondRankName,omitempty"`
}

func (s SearchStrategySearchConfigs) String() string {
	return tea.Prettify(s)
}

func (s SearchStrategySearchConfigs) GoString() string {
	return s.String()
}

func (s *SearchStrategySearchConfigs) SetFirstRankName(v string) *SearchStrategySearchConfigs {
	s.FirstRankName = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetMergeProportion(v int32) *SearchStrategySearchConfigs {
	s.MergeProportion = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetQueryType(v string) *SearchStrategySearchConfigs {
	s.QueryType = &v
	return s
}

func (s *SearchStrategySearchConfigs) SetSecondRankName(v string) *SearchStrategySearchConfigs {
	s.SecondRankName = &v
	return s
}

type SecondRank struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s SecondRank) String() string {
	return tea.Prettify(s)
}

func (s SecondRank) GoString() string {
	return s.String()
}

func (s *SecondRank) SetActive(v bool) *SecondRank {
	s.Active = &v
	return s
}

func (s *SecondRank) SetDescription(v string) *SecondRank {
	s.Description = &v
	return s
}

func (s *SecondRank) SetMeta(v interface{}) *SecondRank {
	s.Meta = v
	return s
}

func (s *SecondRank) SetName(v string) *SecondRank {
	s.Name = &v
	return s
}

type Summary struct {
	Active *bool        `json:"active,omitempty" xml:"active,omitempty"`
	Meta   *SummaryMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	Name   *string      `json:"name,omitempty" xml:"name,omitempty"`
}

func (s Summary) String() string {
	return tea.Prettify(s)
}

func (s Summary) GoString() string {
	return s.String()
}

func (s *Summary) SetActive(v bool) *Summary {
	s.Active = &v
	return s
}

func (s *Summary) SetMeta(v *SummaryMeta) *Summary {
	s.Meta = v
	return s
}

func (s *Summary) SetName(v string) *Summary {
	s.Name = &v
	return s
}

type SummaryMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s SummaryMeta) String() string {
	return tea.Prettify(s)
}

func (s SummaryMeta) GoString() string {
	return s.String()
}

func (s *SummaryMeta) SetElement(v string) *SummaryMeta {
	s.Element = &v
	return s
}

func (s *SummaryMeta) SetEllipsis(v string) *SummaryMeta {
	s.Ellipsis = &v
	return s
}

func (s *SummaryMeta) SetField(v string) *SummaryMeta {
	s.Field = &v
	return s
}

func (s *SummaryMeta) SetLen(v int32) *SummaryMeta {
	s.Len = &v
	return s
}

func (s *SummaryMeta) SetSnippet(v string) *SummaryMeta {
	s.Snippet = &v
	return s
}

type DataSourcePluginsValue struct {
	Name       *string            `json:"name,omitempty" xml:"name,omitempty"`
	FromFields *string            `json:"fromFields,omitempty" xml:"fromFields,omitempty"`
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s DataSourcePluginsValue) String() string {
	return tea.Prettify(s)
}

func (s DataSourcePluginsValue) GoString() string {
	return s.String()
}

func (s *DataSourcePluginsValue) SetName(v string) *DataSourcePluginsValue {
	s.Name = &v
	return s
}

func (s *DataSourcePluginsValue) SetFromFields(v string) *DataSourcePluginsValue {
	s.FromFields = &v
	return s
}

func (s *DataSourcePluginsValue) SetParameters(v map[string]*string) *DataSourcePluginsValue {
	s.Parameters = v
	return s
}

type SchemaIndexesSearchFieldsValue struct {
	Analyzer           *string   `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	AnalyzerType       *string   `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	AnalyzerGeneration *string   `json:"analyzerGeneration,omitempty" xml:"analyzerGeneration,omitempty"`
	Fields             []*string `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	Label              *string   `json:"label,omitempty" xml:"label,omitempty"`
}

func (s SchemaIndexesSearchFieldsValue) String() string {
	return tea.Prettify(s)
}

func (s SchemaIndexesSearchFieldsValue) GoString() string {
	return s.String()
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzer(v string) *SchemaIndexesSearchFieldsValue {
	s.Analyzer = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzerType(v string) *SchemaIndexesSearchFieldsValue {
	s.AnalyzerType = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetAnalyzerGeneration(v string) *SchemaIndexesSearchFieldsValue {
	s.AnalyzerGeneration = &v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetFields(v []*string) *SchemaIndexesSearchFieldsValue {
	s.Fields = v
	return s
}

func (s *SchemaIndexesSearchFieldsValue) SetLabel(v string) *SchemaIndexesSearchFieldsValue {
	s.Label = &v
	return s
}

type SchemaTablesValue struct {
	Name         *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	PrimaryTable *bool                                    `json:"primaryTable,omitempty" xml:"primaryTable,omitempty"`
	Fields       map[string]*SchemaTablesValueFieldsValue `json:"fields,omitempty" xml:"fields,omitempty"`
}

func (s SchemaTablesValue) String() string {
	return tea.Prettify(s)
}

func (s SchemaTablesValue) GoString() string {
	return s.String()
}

func (s *SchemaTablesValue) SetName(v string) *SchemaTablesValue {
	s.Name = &v
	return s
}

func (s *SchemaTablesValue) SetPrimaryTable(v bool) *SchemaTablesValue {
	s.PrimaryTable = &v
	return s
}

func (s *SchemaTablesValue) SetFields(v map[string]*SchemaTablesValueFieldsValue) *SchemaTablesValue {
	s.Fields = v
	return s
}

type SchemaTablesValueFieldsValue struct {
	Name       *string   `json:"name,omitempty" xml:"name,omitempty"`
	PrimaryKey *bool     `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
	Type       *string   `json:"type,omitempty" xml:"type,omitempty"`
	JoinWith   []*string `json:"joinWith,omitempty" xml:"joinWith,omitempty" type:"Repeated"`
	Label      *string   `json:"label,omitempty" xml:"label,omitempty"`
}

func (s SchemaTablesValueFieldsValue) String() string {
	return tea.Prettify(s)
}

func (s SchemaTablesValueFieldsValue) GoString() string {
	return s.String()
}

func (s *SchemaTablesValueFieldsValue) SetName(v string) *SchemaTablesValueFieldsValue {
	s.Name = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetPrimaryKey(v bool) *SchemaTablesValueFieldsValue {
	s.PrimaryKey = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetType(v string) *SchemaTablesValueFieldsValue {
	s.Type = &v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetJoinWith(v []*string) *SchemaTablesValueFieldsValue {
	s.JoinWith = v
	return s
}

func (s *SchemaTablesValueFieldsValue) SetLabel(v string) *SchemaTablesValueFieldsValue {
	s.Label = &v
	return s
}

type BindESUserAnalyzerRequest struct {
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindESUserAnalyzerRequest) String() string {
	return tea.Prettify(s)
}

func (s BindESUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerRequest) SetBody(v interface{}) *BindESUserAnalyzerRequest {
	s.Body = v
	return s
}

type BindESUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindESUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponseBody) SetRequestId(v string) *BindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *BindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type BindESUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindESUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s BindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *BindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *BindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *BindESUserAnalyzerResponse) SetStatusCode(v int32) *BindESUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *BindESUserAnalyzerResponse) SetBody(v *BindESUserAnalyzerResponseBody) *BindESUserAnalyzerResponse {
	s.Body = v
	return s
}

type BindEsInstanceRequest struct {
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindEsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s BindEsInstanceRequest) GoString() string {
	return s.String()
}

func (s *BindEsInstanceRequest) SetBody(v map[string]interface{}) *BindEsInstanceRequest {
	s.Body = v
	return s
}

type BindEsInstanceResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s BindEsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponseBody) SetRequestId(v string) *BindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindEsInstanceResponseBody) SetResult(v map[string]interface{}) *BindEsInstanceResponseBody {
	s.Result = v
	return s
}

type BindEsInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BindEsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s BindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *BindEsInstanceResponse) SetHeaders(v map[string]*string) *BindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *BindEsInstanceResponse) SetStatusCode(v int32) *BindEsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindEsInstanceResponse) SetBody(v *BindEsInstanceResponseBody) *BindEsInstanceResponse {
	s.Body = v
	return s
}

type CompileSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CompileSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompileSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponseBody) SetRequestId(v string) *CompileSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type CompileSortScriptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompileSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompileSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s CompileSortScriptResponse) GoString() string {
	return s.String()
}

func (s *CompileSortScriptResponse) SetHeaders(v map[string]*string) *CompileSortScriptResponse {
	s.Headers = v
	return s
}

func (s *CompileSortScriptResponse) SetStatusCode(v int32) *CompileSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *CompileSortScriptResponse) SetBody(v *CompileSortScriptResponseBody) *CompileSortScriptResponse {
	s.Body = v
	return s
}

type CreateABTestExperimentRequest struct {
	Body   *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool             `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentRequest) SetBody(v *ABTestExperiment) *CreateABTestExperimentRequest {
	s.Body = v
	return s
}

func (s *CreateABTestExperimentRequest) SetDryRun(v bool) *CreateABTestExperimentRequest {
	s.DryRun = &v
	return s
}

type CreateABTestExperimentResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test.
	Result *CreateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBody) SetRequestId(v string) *CreateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestExperimentResponseBody) SetResult(v *CreateABTestExperimentResponseBodyResult) *CreateABTestExperimentResponseBody {
	s.Result = v
	return s
}

type CreateABTestExperimentResponseBodyResult struct {
	// The time when the test was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// *   true: in effect
	// *   false: not in effect
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the test.
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponseBodyResult) SetCreated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetId(v string) *CreateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetName(v string) *CreateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetOnline(v bool) *CreateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *CreateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetTraffic(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *CreateABTestExperimentResponseBodyResult) SetUpdated(v int32) *CreateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponse) SetHeaders(v map[string]*string) *CreateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestExperimentResponse) SetStatusCode(v int32) *CreateABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestExperimentResponse) SetBody(v *CreateABTestExperimentResponseBody) *CreateABTestExperimentResponse {
	s.Body = v
	return s
}

type CreateABTestGroupRequest struct {
	Body   *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool        `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupRequest) SetBody(v *ABTestGroup) *CreateABTestGroupRequest {
	s.Body = v
	return s
}

func (s *CreateABTestGroupRequest) SetDryRun(v bool) *CreateABTestGroupRequest {
	s.DryRun = &v
	return s
}

type CreateABTestGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *CreateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBody) SetRequestId(v string) *CreateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestGroupResponseBody) SetResult(v *CreateABTestGroupResponseBodyResult) *CreateABTestGroupResponseBody {
	s.Result = v
	return s
}

type CreateABTestGroupResponseBodyResult struct {
	// The time when the test group was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponseBodyResult) SetCreated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetId(v string) *CreateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetName(v string) *CreateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetStatus(v int32) *CreateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestGroupResponseBodyResult) SetUpdated(v int32) *CreateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestGroupResponse) SetHeaders(v map[string]*string) *CreateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestGroupResponse) SetStatusCode(v int32) *CreateABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestGroupResponse) SetBody(v *CreateABTestGroupResponseBody) *CreateABTestGroupResponse {
	s.Body = v
	return s
}

type CreateABTestSceneRequest struct {
	Body   *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool        `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneRequest) SetBody(v *ABTestScene) *CreateABTestSceneRequest {
	s.Body = v
	return s
}

func (s *CreateABTestSceneRequest) SetDryRun(v bool) *CreateABTestSceneRequest {
	s.DryRun = &v
	return s
}

type CreateABTestSceneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *CreateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBody) SetRequestId(v string) *CreateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateABTestSceneResponseBody) SetResult(v *CreateABTestSceneResponseBodyResult) *CreateABTestSceneResponseBody {
	s.Result = v
	return s
}

type CreateABTestSceneResponseBodyResult struct {
	// The time when the test scenario was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The tag of the test scenario.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s CreateABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponseBodyResult) SetCreated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetId(v string) *CreateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetName(v string) *CreateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetStatus(v int32) *CreateABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetUpdated(v int32) *CreateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *CreateABTestSceneResponseBodyResult) SetValues(v []*string) *CreateABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

type CreateABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneResponse) SetHeaders(v map[string]*string) *CreateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestSceneResponse) SetStatusCode(v int32) *CreateABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestSceneResponse) SetBody(v *CreateABTestSceneResponseBody) *CreateABTestSceneResponse {
	s.Body = v
	return s
}

type CreateAppRequest struct {
	Body *App `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetBody(v *App) *CreateAppRequest {
	s.Body = v
	return s
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

type CreateAppResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBody) SetRequestId(v string) *CreateAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppResponseBody) SetResult(v map[string]interface{}) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponse) GoString() string {
	return s.String()
}

func (s *CreateAppResponse) SetHeaders(v map[string]*string) *CreateAppResponse {
	s.Headers = v
	return s
}

func (s *CreateAppResponse) SetStatusCode(v int32) *CreateAppResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppResponse) SetBody(v *CreateAppResponseBody) *CreateAppResponse {
	s.Body = v
	return s
}

type CreateAppGroupRequest struct {
	Body *AppGroup `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) SetBody(v *AppGroup) *CreateAppGroupRequest {
	s.Body = v
	return s
}

type CreateAppGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result *CreateAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBody) SetRequestId(v string) *CreateAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupResponseBody) SetResult(v *CreateAppGroupResponseBodyResult) *CreateAppGroupResponseBody {
	s.Result = v
	return s
}

type CreateAppGroupResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *CreateAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResult) SetChargeType(v string) *CreateAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetChargingWay(v int32) *CreateAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCommodityCode(v string) *CreateAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCreated(v int32) *CreateAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetCurrentVersion(v string) *CreateAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDescription(v string) *CreateAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetDomain(v string) *CreateAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetExpireOn(v string) *CreateAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *CreateAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetId(v string) *CreateAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetInstanceId(v string) *CreateAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetLockMode(v string) *CreateAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *CreateAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetName(v string) *CreateAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProcessingOrderId(v string) *CreateAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProduced(v int32) *CreateAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetProjectId(v string) *CreateAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetQuota(v *CreateAppGroupResponseBodyResultQuota) *CreateAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *CreateAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetStatus(v string) *CreateAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetSwitchedTime(v int32) *CreateAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetType(v string) *CreateAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetUpdated(v int32) *CreateAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateAppGroupResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetDocSize(v int32) *CreateAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppGroupResponseBodyResultQuota) SetSpec(v string) *CreateAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type CreateAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupResponse) SetHeaders(v map[string]*string) *CreateAppGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupResponse) SetStatusCode(v int32) *CreateAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupResponse) SetBody(v *CreateAppGroupResponseBody) *CreateAppGroupResponse {
	s.Body = v
	return s
}

type CreateFirstRankRequest struct {
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateFirstRankRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankRequest) GoString() string {
	return s.String()
}

func (s *CreateFirstRankRequest) SetBody(v *FirstRank) *CreateFirstRankRequest {
	s.Body = v
	return s
}

func (s *CreateFirstRankRequest) SetDryRun(v bool) *CreateFirstRankRequest {
	s.DryRun = &v
	return s
}

type CreateFirstRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *CreateFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBody) SetRequestId(v string) *CreateFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFirstRankResponseBody) SetResult(v *CreateFirstRankResponseBodyResult) *CreateFirstRankResponseBody {
	s.Result = v
	return s
}

type CreateFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The content of the expression.
	Meta []*CreateFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResult) SetActive(v bool) *CreateFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetMeta(v []*CreateFirstRankResponseBodyResultMeta) *CreateFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *CreateFirstRankResponseBodyResult) SetName(v string) *CreateFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type CreateFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: \[-100000,100000] (excluding 0).
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s CreateFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponseBodyResultMeta) SetArg(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetAttribute(v string) *CreateFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *CreateFirstRankResponseBodyResultMeta) SetWeight(v float32) *CreateFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type CreateFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFirstRankResponse) GoString() string {
	return s.String()
}

func (s *CreateFirstRankResponse) SetHeaders(v map[string]*string) *CreateFirstRankResponse {
	s.Headers = v
	return s
}

func (s *CreateFirstRankResponse) SetStatusCode(v int32) *CreateFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFirstRankResponse) SetBody(v *CreateFirstRankResponseBody) *CreateFirstRankResponse {
	s.Body = v
	return s
}

type CreateFunctionInstanceRequest struct {
	// The parameters that are used to create the instance.
	CreateParameters []*CreateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The cron expression used to schedule periodic training, in the format of (Minutes Hours DayofMonth Month DayofWeek). The default value is empty, which indicates that no periodic training is performed. DayofWeek 0 indicates Sunday.
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the feature. Valid values:
	//
	// *   PAAS: This is the default value. Training is required before you can use the feature.
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The name of the instance. The name must be 1 to 30 characters in length and can contain letters, digits, and underscores (\_). The name is case-sensitive and must start with a letter.
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The type of the model. The following features correspond to different model types:
	//
	// *   click-through rate (CTR) model: tf_checkpoint
	// *   Popularity model: pop
	// *   Category model: offline_inference
	// *   Hotword model: offline_inference
	// *   Shading model: offline_inference
	// *   Drop-down suggestion model: offline_inference
	// *   Word segmentation model: text
	// *   Term weight model: tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The parameters that are used to use the instance.
	UsageParameters []*CreateFunctionInstanceRequestUsageParameters `json:"usageParameters,omitempty" xml:"usageParameters,omitempty" type:"Repeated"`
}

func (s CreateFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequest) SetCreateParameters(v []*CreateFunctionInstanceRequestCreateParameters) *CreateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *CreateFunctionInstanceRequest) SetCron(v string) *CreateFunctionInstanceRequest {
	s.Cron = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetDescription(v string) *CreateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetFunctionType(v string) *CreateFunctionInstanceRequest {
	s.FunctionType = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetInstanceName(v string) *CreateFunctionInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetModelType(v string) *CreateFunctionInstanceRequest {
	s.ModelType = &v
	return s
}

func (s *CreateFunctionInstanceRequest) SetUsageParameters(v []*CreateFunctionInstanceRequestUsageParameters) *CreateFunctionInstanceRequest {
	s.UsageParameters = v
	return s
}

type CreateFunctionInstanceRequestCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateFunctionInstanceRequestCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetName(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *CreateFunctionInstanceRequestCreateParameters) SetValue(v string) *CreateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

type CreateFunctionInstanceRequestUsageParameters struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreateFunctionInstanceRequestUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceRequestUsageParameters) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceRequestUsageParameters) SetName(v string) *CreateFunctionInstanceRequestUsageParameters {
	s.Name = &v
	return s
}

func (s *CreateFunctionInstanceRequestUsageParameters) SetValue(v string) *CreateFunctionInstanceRequestUsageParameters {
	s.Value = &v
	return s
}

type CreateFunctionInstanceResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponseBody) SetCode(v string) *CreateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetHttpCode(v int64) *CreateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetLatency(v int64) *CreateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetMessage(v string) *CreateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetRequestId(v string) *CreateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionInstanceResponseBody) SetStatus(v string) *CreateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type CreateFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionInstanceResponse) SetHeaders(v map[string]*string) *CreateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionInstanceResponse) SetStatusCode(v int32) *CreateFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionInstanceResponse) SetBody(v *CreateFunctionInstanceResponseBody) *CreateFunctionInstanceResponse {
	s.Body = v
	return s
}

type CreateFunctionTaskResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponseBody) SetCode(v string) *CreateFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetHttpCode(v int64) *CreateFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetLatency(v int64) *CreateFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetMessage(v string) *CreateFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetRequestId(v string) *CreateFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionTaskResponseBody) SetStatus(v string) *CreateFunctionTaskResponseBody {
	s.Status = &v
	return s
}

type CreateFunctionTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateFunctionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionTaskResponse) SetHeaders(v map[string]*string) *CreateFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionTaskResponse) SetStatusCode(v int32) *CreateFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionTaskResponse) SetBody(v *CreateFunctionTaskResponseBody) *CreateFunctionTaskResponse {
	s.Body = v
	return s
}

type CreateInterventionDictionaryRequest struct {
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	Name         *string `json:"name,omitempty" xml:"name,omitempty"`
	Type         *string `json:"type,omitempty" xml:"type,omitempty"`
	DryRun       *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateInterventionDictionaryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryRequest) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryRequest) SetAnalyzerType(v string) *CreateInterventionDictionaryRequest {
	s.AnalyzerType = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetName(v string) *CreateInterventionDictionaryRequest {
	s.Name = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetType(v string) *CreateInterventionDictionaryRequest {
	s.Type = &v
	return s
}

func (s *CreateInterventionDictionaryRequest) SetDryRun(v bool) *CreateInterventionDictionaryRequest {
	s.DryRun = &v
	return s
}

type CreateInterventionDictionaryResponseBody struct {
	// The type of the intervention dictionary. Valid values:
	//
	// *   stopword: an intervention dictionary for stop word filtering
	// *   synonym: an intervention dictionary for synonym configuration
	// *   correction: an intervention dictionary for spelling correction
	// *   category_prediction: an intervention dictionary for category prediction
	// *   ner: an intervention dictionary for named entity recognition (NER)
	// *   term_weighting: an intervention dictionary for term weight analysis
	// *   suggest_allowlist: a drop-down suggestion whitelist
	// *   suggest_denylist: a drop-down suggestion blacklist
	// *   hot_allowlist: a top search whitelist
	// *   hot_denylist: a top search blacklist
	// *   hint_allowlist: a shading whitelist
	// *   hint_denylist: a shading blacklist
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The ID of the request.
	Result *CreateInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBody) SetRequestId(v string) *CreateInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBody) SetResult(v *CreateInterventionDictionaryResponseBodyResult) *CreateInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type CreateInterventionDictionaryResponseBodyResult struct {
	// Creates an intervention dictionary.
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The name of the intervention dictionary.
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Type    *string `json:"type,omitempty" xml:"type,omitempty"`
	// CreateInterventionDictionary
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetCreated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetName(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetType(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *CreateInterventionDictionaryResponseBodyResult) SetUpdated(v string) *CreateInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateInterventionDictionaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *CreateInterventionDictionaryResponse) SetHeaders(v map[string]*string) *CreateInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *CreateInterventionDictionaryResponse) SetStatusCode(v int32) *CreateInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInterventionDictionaryResponse) SetBody(v *CreateInterventionDictionaryResponseBody) *CreateInterventionDictionaryResponse {
	s.Body = v
	return s
}

type CreateQueryProcessorRequest struct {
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateQueryProcessorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorRequest) SetBody(v interface{}) *CreateQueryProcessorRequest {
	s.Body = v
	return s
}

func (s *CreateQueryProcessorRequest) SetDryRun(v bool) *CreateQueryProcessorRequest {
	s.DryRun = &v
	return s
}

type CreateQueryProcessorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	Result *CreateQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBody) SetRequestId(v string) *CreateQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQueryProcessorResponseBody) SetResult(v *CreateQueryProcessorResponseBodyResult) *CreateQueryProcessorResponseBody {
	s.Result = v
	return s
}

type CreateQueryProcessorResponseBodyResult struct {
	// Indicates whether the query analysis rule is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s CreateQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponseBodyResult) SetActive(v bool) *CreateQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetCreated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetDomain(v string) *CreateQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetIndexes(v []*string) *CreateQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetName(v string) *CreateQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *CreateQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *CreateQueryProcessorResponseBodyResult) SetUpdated(v int32) *CreateQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type CreateQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorResponse) SetHeaders(v map[string]*string) *CreateQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *CreateQueryProcessorResponse) SetStatusCode(v int32) *CreateQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateQueryProcessorResponse) SetBody(v *CreateQueryProcessorResponseBody) *CreateQueryProcessorResponse {
	s.Body = v
	return s
}

type CreateScheduledTaskRequest struct {
	Body *ScheduledTask `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskRequest) SetBody(v *ScheduledTask) *CreateScheduledTaskRequest {
	s.Body = v
	return s
}

type CreateScheduledTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponseBody) SetRequestId(v string) *CreateScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledTaskResponseBody) SetResult(v map[string]interface{}) *CreateScheduledTaskResponseBody {
	s.Result = v
	return s
}

type CreateScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledTaskResponse) SetHeaders(v map[string]*string) *CreateScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledTaskResponse) SetStatusCode(v int32) *CreateScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledTaskResponse) SetBody(v *CreateScheduledTaskResponseBody) *CreateScheduledTaskResponse {
	s.Body = v
	return s
}

type CreateSearchStrategyRequest struct {
	Body *SearchStrategy `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSearchStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchStrategyRequest) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyRequest) SetBody(v *SearchStrategy) *CreateSearchStrategyRequest {
	s.Body = v
	return s
}

type CreateSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponseBody) SetRequestId(v string) *CreateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type CreateSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *CreateSearchStrategyResponse) SetHeaders(v map[string]*string) *CreateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *CreateSearchStrategyResponse) SetStatusCode(v int32) *CreateSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSearchStrategyResponse) SetBody(v *CreateSearchStrategyResponseBody) *CreateSearchStrategyResponse {
	s.Body = v
	return s
}

type CreateSecondRankRequest struct {
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateSecondRankRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankRequest) GoString() string {
	return s.String()
}

func (s *CreateSecondRankRequest) SetBody(v *SecondRank) *CreateSecondRankRequest {
	s.Body = v
	return s
}

func (s *CreateSecondRankRequest) SetDryRun(v bool) *CreateSecondRankRequest {
	s.DryRun = &v
	return s
}

type CreateSecondRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the fine sort expression.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponseBody) SetRequestId(v string) *CreateSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecondRankResponseBody) SetResult(v map[string]interface{}) *CreateSecondRankResponseBody {
	s.Result = v
	return s
}

type CreateSecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecondRankResponse) GoString() string {
	return s.String()
}

func (s *CreateSecondRankResponse) SetHeaders(v map[string]*string) *CreateSecondRankResponse {
	s.Headers = v
	return s
}

func (s *CreateSecondRankResponse) SetStatusCode(v int32) *CreateSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecondRankResponse) SetBody(v *CreateSecondRankResponseBody) *CreateSecondRankResponse {
	s.Body = v
	return s
}

type CreateSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponseBody) SetRequestId(v string) *CreateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type CreateSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptResponse) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponse) SetHeaders(v map[string]*string) *CreateSortScriptResponse {
	s.Headers = v
	return s
}

func (s *CreateSortScriptResponse) SetStatusCode(v int32) *CreateSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSortScriptResponse) SetBody(v *CreateSortScriptResponseBody) *CreateSortScriptResponse {
	s.Body = v
	return s
}

type CreateUserAnalyzerRequest struct {
	// 基础分词器
	Business           *string `json:"business,omitempty" xml:"business,omitempty"`
	BusinessAppGroupId *string `json:"businessAppGroupId,omitempty" xml:"businessAppGroupId,omitempty"`
	// 基础分词器类型 (AUTO, MODEL, SYSTEM, USER)
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 引擎类型 (HA3, ES)
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	DryRun *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateUserAnalyzerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerRequest) SetBusiness(v string) *CreateUserAnalyzerRequest {
	s.Business = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetBusinessAppGroupId(v string) *CreateUserAnalyzerRequest {
	s.BusinessAppGroupId = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetBusinessType(v string) *CreateUserAnalyzerRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetName(v string) *CreateUserAnalyzerRequest {
	s.Name = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetType(v string) *CreateUserAnalyzerRequest {
	s.Type = &v
	return s
}

func (s *CreateUserAnalyzerRequest) SetDryRun(v bool) *CreateUserAnalyzerRequest {
	s.DryRun = &v
	return s
}

type CreateUserAnalyzerResponseBody struct {
	RequestId *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s CreateUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponseBody) SetRequestId(v string) *CreateUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *CreateUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type CreateUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *CreateUserAnalyzerResponse) SetHeaders(v map[string]*string) *CreateUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *CreateUserAnalyzerResponse) SetStatusCode(v int32) *CreateUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserAnalyzerResponse) SetBody(v *CreateUserAnalyzerResponseBody) *CreateUserAnalyzerResponse {
	s.Body = v
	return s
}

type DeleteABTestExperimentResponseBody struct {
	// The ID of the request.
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponseBody) SetRequestId(v string) *DeleteABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestExperimentResponseBody) SetResult(v map[string]interface{}) *DeleteABTestExperimentResponseBody {
	s.Result = v
	return s
}

type DeleteABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponse) SetHeaders(v map[string]*string) *DeleteABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestExperimentResponse) SetStatusCode(v int32) *DeleteABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestExperimentResponse) SetBody(v *DeleteABTestExperimentResponseBody) *DeleteABTestExperimentResponse {
	s.Body = v
	return s
}

type DeleteABTestGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponseBody) SetRequestId(v string) *DeleteABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestGroupResponseBody) SetResult(v map[string]interface{}) *DeleteABTestGroupResponseBody {
	s.Result = v
	return s
}

type DeleteABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestGroupResponse) SetHeaders(v map[string]*string) *DeleteABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestGroupResponse) SetStatusCode(v int32) *DeleteABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestGroupResponse) SetBody(v *DeleteABTestGroupResponseBody) *DeleteABTestGroupResponse {
	s.Body = v
	return s
}

type DeleteABTestSceneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DeleteABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponseBody) SetRequestId(v string) *DeleteABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteABTestSceneResponseBody) SetResult(v map[string]interface{}) *DeleteABTestSceneResponseBody {
	s.Result = v
	return s
}

type DeleteABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponse) SetHeaders(v map[string]*string) *DeleteABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestSceneResponse) SetStatusCode(v int32) *DeleteABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestSceneResponse) SetBody(v *DeleteABTestSceneResponseBody) *DeleteABTestSceneResponse {
	s.Body = v
	return s
}

type DeleteFunctionInstanceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// *   OK: The request is successful.
	// *   FAIL: The request fails.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponseBody) SetCode(v string) *DeleteFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetHttpCode(v int64) *DeleteFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetLatency(v int64) *DeleteFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetMessage(v string) *DeleteFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetRequestId(v string) *DeleteFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionInstanceResponseBody) SetStatus(v string) *DeleteFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type DeleteFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionInstanceResponse) SetHeaders(v map[string]*string) *DeleteFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionInstanceResponse) SetStatusCode(v int32) *DeleteFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionInstanceResponse) SetBody(v *DeleteFunctionInstanceResponseBody) *DeleteFunctionInstanceResponse {
	s.Body = v
	return s
}

type DeleteFunctionTaskResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionTaskResponseBody) SetCode(v string) *DeleteFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetHttpCode(v int64) *DeleteFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetLatency(v int64) *DeleteFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetMessage(v string) *DeleteFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetRequestId(v string) *DeleteFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionTaskResponseBody) SetStatus(v string) *DeleteFunctionTaskResponseBody {
	s.Status = &v
	return s
}

type DeleteFunctionTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFunctionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionTaskResponse) SetHeaders(v map[string]*string) *DeleteFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionTaskResponse) SetStatusCode(v int32) *DeleteFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionTaskResponse) SetBody(v *DeleteFunctionTaskResponseBody) *DeleteFunctionTaskResponse {
	s.Body = v
	return s
}

type DeleteSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponseBody) SetRequestId(v string) *DeleteSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptResponse) SetHeaders(v map[string]*string) *DeleteSortScriptResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptResponse) SetStatusCode(v int32) *DeleteSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSortScriptResponse) SetBody(v *DeleteSortScriptResponseBody) *DeleteSortScriptResponse {
	s.Body = v
	return s
}

type DeleteSortScriptFileResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponseBody) SetRequestId(v string) *DeleteSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSortScriptFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteSortScriptFileResponse) SetHeaders(v map[string]*string) *DeleteSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteSortScriptFileResponse) SetStatusCode(v int32) *DeleteSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSortScriptFileResponse) SetBody(v *DeleteSortScriptFileResponseBody) *DeleteSortScriptFileResponse {
	s.Body = v
	return s
}

type DescribeABTestExperimentResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test.
	Result *DescribeABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBody) SetRequestId(v string) *DescribeABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestExperimentResponseBody) SetResult(v *DescribeABTestExperimentResponseBodyResult) *DescribeABTestExperimentResponseBody {
	s.Result = v
	return s
}

type DescribeABTestExperimentResponseBodyResult struct {
	// The time when the test was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// *   true: in effect
	// *   false: not in effect
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params *DescribeABTestExperimentResponseBodyResultParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The percentage of traffic that is routed to the test.
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResult) SetCreated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetId(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetName(v string) *DescribeABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetOnline(v bool) *DescribeABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetParams(v *DescribeABTestExperimentResponseBodyResultParams) *DescribeABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetTraffic(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *DescribeABTestExperimentResponseBodyResult) SetUpdated(v int32) *DescribeABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeABTestExperimentResponseBodyResultParams struct {
	// The name of the rough sort policy.
	FirstFormulaName *string `json:"first_formula_name,omitempty" xml:"first_formula_name,omitempty"`
}

func (s DescribeABTestExperimentResponseBodyResultParams) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponseBodyResultParams) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponseBodyResultParams) SetFirstFormulaName(v string) *DescribeABTestExperimentResponseBodyResultParams {
	s.FirstFormulaName = &v
	return s
}

type DescribeABTestExperimentResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestExperimentResponse) SetHeaders(v map[string]*string) *DescribeABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestExperimentResponse) SetStatusCode(v int32) *DescribeABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestExperimentResponse) SetBody(v *DescribeABTestExperimentResponseBody) *DescribeABTestExperimentResponse {
	s.Body = v
	return s
}

type DescribeABTestGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test group.
	Result *DescribeABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBody) SetRequestId(v string) *DescribeABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestGroupResponseBody) SetResult(v *DescribeABTestGroupResponseBodyResult) *DescribeABTestGroupResponseBody {
	s.Result = v
	return s
}

type DescribeABTestGroupResponseBodyResult struct {
	// The time when the test group was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponseBodyResult) SetCreated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetId(v string) *DescribeABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetName(v string) *DescribeABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetStatus(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestGroupResponseBodyResult) SetUpdated(v int32) *DescribeABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeABTestGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestGroupResponse) SetHeaders(v map[string]*string) *DescribeABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestGroupResponse) SetStatusCode(v int32) *DescribeABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestGroupResponse) SetBody(v *DescribeABTestGroupResponseBody) *DescribeABTestGroupResponse {
	s.Body = v
	return s
}

type DescribeABTestSceneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario.
	Result *DescribeABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBody) SetRequestId(v string) *DescribeABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeABTestSceneResponseBody) SetResult(v *DescribeABTestSceneResponseBodyResult) *DescribeABTestSceneResponseBody {
	s.Result = v
	return s
}

type DescribeABTestSceneResponseBodyResult struct {
	// The time when the test scenario was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test scenario.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test scenario.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The tag of the test scenario.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s DescribeABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponseBodyResult) SetCreated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetId(v string) *DescribeABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetName(v string) *DescribeABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetStatus(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetUpdated(v int32) *DescribeABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeABTestSceneResponseBodyResult) SetValues(v []*string) *DescribeABTestSceneResponseBodyResult {
	s.Values = v
	return s
}

type DescribeABTestSceneResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DescribeABTestSceneResponse) SetHeaders(v map[string]*string) *DescribeABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DescribeABTestSceneResponse) SetStatusCode(v int32) *DescribeABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeABTestSceneResponse) SetBody(v *DescribeABTestSceneResponseBody) *DescribeABTestSceneResponse {
	s.Body = v
	return s
}

type DescribeAppResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the version.
	Result *DescribeAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBody) SetRequestId(v string) *DescribeAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppResponseBody) SetResult(v *DescribeAppResponseBodyResult) *DescribeAppResponseBody {
	s.Result = v
	return s
}

type DescribeAppResponseBodyResult struct {
	// The ID of the created rough sort expression.
	AlgoDeploymentId *int32 `json:"algoDeploymentId,omitempty" xml:"algoDeploymentId,omitempty"`
	// Indicates whether the version is automatically published to the online environment.
	AutoSwitch  *bool   `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	ClusterName *string `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	// The timestamp when the version was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the version.
	Description *string                              `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *DescribeAppResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	FetchFields []*string `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	// The ID of the version.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The progress of data import, in percentage. For example, a value of 83 indicates 83%.
	ProgressPercent *int32 `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	// The quota information about the version.
	Quota *DescribeAppResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The application schema.
	Schema map[string]interface{} `json:"schema,omitempty" xml:"schema,omitempty"`
	// The status of the version. Valid values:
	//
	// *   ok
	// *   stopped
	// *   frozen
	// *   initializing
	// *   unavailable
	// *   data_waiting
	// *   data_preparing
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResult) SetAlgoDeploymentId(v int32) *DescribeAppResponseBodyResult {
	s.AlgoDeploymentId = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetClusterName(v string) *DescribeAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetCreated(v int32) *DescribeAppResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDescription(v string) *DescribeAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDomain(v *DescribeAppResponseBodyResultDomain) *DescribeAppResponseBodyResult {
	s.Domain = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetFetchFields(v []*string) *DescribeAppResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetId(v string) *DescribeAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetProgressPercent(v int32) *DescribeAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQuota(v *DescribeAppResponseBodyResultQuota) *DescribeAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchema(v map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetStatus(v string) *DescribeAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetType(v string) *DescribeAppResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeAppResponseBodyResultDomain struct {
	Category  *string                                       `json:"category,omitempty" xml:"category,omitempty"`
	Functions *DescribeAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	Name      *string                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomain) SetCategory(v string) *DescribeAppResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetFunctions(v *DescribeAppResponseBodyResultDomainFunctions) *DescribeAppResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *DescribeAppResponseBodyResultDomain) SetName(v string) *DescribeAppResponseBodyResultDomain {
	s.Name = &v
	return s
}

type DescribeAppResponseBodyResultDomainFunctions struct {
	Algo    []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	Qp      []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s DescribeAppResponseBodyResultDomainFunctions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetAlgo(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetQp(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *DescribeAppResponseBodyResultDomainFunctions) SetService(v []*string) *DescribeAppResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

type DescribeAppResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The number of search requests per second. You are charged based on the number of search requests per second in the earlier billing model.
	Qps *int32 `json:"qps,omitempty" xml:"qps,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeAppResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetQps(v int32) *DescribeAppResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *DescribeAppResponseBodyResultQuota) SetSpec(v string) *DescribeAppResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type DescribeAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppResponse) SetHeaders(v map[string]*string) *DescribeAppResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppResponse) SetStatusCode(v int32) *DescribeAppResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppResponse) SetBody(v *DescribeAppResponseBody) *DescribeAppResponse {
	s.Body = v
	return s
}

type DescribeAppGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	Result *DescribeAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBody) SetRequestId(v string) *DescribeAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppGroupResponseBody) SetResult(v *DescribeAppGroupResponseBodyResult) *DescribeAppGroupResponseBody {
	s.Result = v
	return s
}

type DescribeAppGroupResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	Domain      *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota           *DescribeAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId *string                                  `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32                                    `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Tags         []*DescribeAppGroupResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResult) SetChargeType(v string) *DescribeAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetChargingWay(v int32) *DescribeAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCommodityCode(v string) *DescribeAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCreated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetCurrentVersion(v string) *DescribeAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDescription(v string) *DescribeAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetDomain(v string) *DescribeAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetExpireOn(v string) *DescribeAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *DescribeAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetId(v string) *DescribeAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetInstanceId(v string) *DescribeAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockMode(v string) *DescribeAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *DescribeAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetName(v string) *DescribeAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProcessingOrderId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProduced(v int32) *DescribeAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetProjectId(v string) *DescribeAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetQuota(v *DescribeAppGroupResponseBodyResultQuota) *DescribeAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetResourceGroupId(v string) *DescribeAppGroupResponseBodyResult {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *DescribeAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetStatus(v string) *DescribeAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetSwitchedTime(v int32) *DescribeAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetTags(v []*DescribeAppGroupResponseBodyResultTags) *DescribeAppGroupResponseBodyResult {
	s.Tags = v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetType(v string) *DescribeAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResult) SetUpdated(v int32) *DescribeAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeAppGroupResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultQuota) SetSpec(v string) *DescribeAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type DescribeAppGroupResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeAppGroupResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponseBodyResultTags) SetKey(v string) *DescribeAppGroupResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *DescribeAppGroupResponseBodyResultTags) SetValue(v string) *DescribeAppGroupResponseBodyResultTags {
	s.Value = &v
	return s
}

type DescribeAppGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppGroupResponse) SetHeaders(v map[string]*string) *DescribeAppGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppGroupResponse) SetStatusCode(v int32) *DescribeAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppGroupResponse) SetBody(v *DescribeAppGroupResponseBody) *DescribeAppGroupResponse {
	s.Body = v
	return s
}

type DescribeAppStatisticsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The statistics.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeAppStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponseBody) SetRequestId(v string) *DescribeAppStatisticsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppStatisticsResponseBody) SetResult(v map[string]interface{}) *DescribeAppStatisticsResponseBody {
	s.Result = v
	return s
}

type DescribeAppStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppStatisticsResponse) SetHeaders(v map[string]*string) *DescribeAppStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppStatisticsResponse) SetStatusCode(v int32) *DescribeAppStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppStatisticsResponse) SetBody(v *DescribeAppStatisticsResponseBody) *DescribeAppStatisticsResponse {
	s.Body = v
	return s
}

type DescribeAppsResponseBody struct {
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBody) SetRequestId(v string) *DescribeAppsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppsResponseBody) SetResult(v []map[string]interface{}) *DescribeAppsResponseBody {
	s.Result = v
	return s
}

type DescribeAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponse) SetHeaders(v map[string]*string) *DescribeAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsResponse) SetStatusCode(v int32) *DescribeAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsResponse) SetBody(v *DescribeAppsResponseBody) *DescribeAppsResponse {
	s.Body = v
	return s
}

type DescribeDataCollctionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the data collection task.
	Result *DescribeDataCollctionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeDataCollctionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBody) SetRequestId(v string) *DescribeDataCollctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCollctionResponseBody) SetResult(v *DescribeDataCollctionResponseBodyResult) *DescribeDataCollctionResponseBody {
	s.Result = v
	return s
}

type DescribeDataCollctionResponseBodyResult struct {
	// The time when the data collection task was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the data that is collected by the task. Valid values:
	//
	// *   behavior: behavioral data
	// *   item_info: project data
	// *   industry_specific: industry-specific data
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	// The ID of the data collection task.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry to which the data collection task applies. Valid values:
	//
	// *   general
	// *   ecommerce
	IndustryName *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	// The name of the data collection task.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data collection task. Valid values:
	//
	// *   0: disabled
	// *   1: being enabled
	// *   2: enabled
	// *   3: failed to be enabled
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the sundial.
	SundialId *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	// The type of the data source. Valid values:
	//
	// *   server
	//
	// *   web
	//
	// *   app
	//
	//     Note: Only server is supported.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the data collection task was updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeDataCollctionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBodyResult) SetCreated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetDataCollectionType(v string) *DescribeDataCollctionResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetId(v string) *DescribeDataCollctionResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetIndustryName(v string) *DescribeDataCollctionResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetName(v string) *DescribeDataCollctionResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetStatus(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetSundialId(v string) *DescribeDataCollctionResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetType(v string) *DescribeDataCollctionResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetUpdated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeDataCollctionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDataCollctionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDataCollctionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDataCollctionResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponse) SetHeaders(v map[string]*string) *DescribeDataCollctionResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataCollctionResponse) SetStatusCode(v int32) *DescribeDataCollctionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataCollctionResponse) SetBody(v *DescribeDataCollctionResponseBody) *DescribeDataCollctionResponse {
	s.Body = v
	return s
}

type DescribeFirstRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *DescribeFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBody) SetRequestId(v string) *DescribeFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFirstRankResponseBody) SetResult(v *DescribeFirstRankResponseBodyResult) *DescribeFirstRankResponseBody {
	s.Result = v
	return s
}

type DescribeFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*DescribeFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResult) SetActive(v bool) *DescribeFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetDescription(v string) *DescribeFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetMeta(v []*DescribeFirstRankResponseBodyResultMeta) *DescribeFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *DescribeFirstRankResponseBodyResult) SetName(v string) *DescribeFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type DescribeFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: \[-100000,100000] (excluding 0).
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s DescribeFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetArg(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetAttribute(v string) *DescribeFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *DescribeFirstRankResponseBodyResultMeta) SetWeight(v float32) *DescribeFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type DescribeFirstRankResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeFirstRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeFirstRankResponse) SetHeaders(v map[string]*string) *DescribeFirstRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeFirstRankResponse) SetStatusCode(v int32) *DescribeFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFirstRankResponse) SetBody(v *DescribeFirstRankResponseBody) *DescribeFirstRankResponse {
	s.Body = v
	return s
}

type DescribeInterventionDictionaryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information the intervention dictionary.
	Result *DescribeInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBody) SetRequestId(v string) *DescribeInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBody) SetResult(v *DescribeInterventionDictionaryResponseBodyResult) *DescribeInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type DescribeInterventionDictionaryResponseBodyResult struct {
	// The custom analyzer.
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the intervention dictionary.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// *   stopword: an intervention dictionary for stop word filtering
	// *   synonym: an intervention dictionary for synonym configuration
	// *   correction: an intervention dictionary for spelling correction
	// *   category_prediction: an intervention dictionary for category prediction
	// *   ner: an intervention dictionary for named entity recognition (NER)
	// *   term_weighting: an intervention dictionary for term weight analysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetCreated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetName(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetType(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeInterventionDictionaryResponseBodyResult) SetUpdated(v string) *DescribeInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeInterventionDictionaryResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeInterventionDictionaryResponse) SetHeaders(v map[string]*string) *DescribeInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeInterventionDictionaryResponse) SetStatusCode(v int32) *DescribeInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInterventionDictionaryResponse) SetBody(v *DescribeInterventionDictionaryResponseBody) *DescribeInterventionDictionaryResponse {
	s.Body = v
	return s
}

type DescribeQueryProcessorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	Result *DescribeQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBody) SetRequestId(v string) *DescribeQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeQueryProcessorResponseBody) SetResult(v *DescribeQueryProcessorResponseBodyResult) *DescribeQueryProcessorResponseBody {
	s.Result = v
	return s
}

type DescribeQueryProcessorResponseBodyResult struct {
	// Indicates whether the query analysis rule is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponseBodyResult) SetActive(v bool) *DescribeQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetCreated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetDomain(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetIndexes(v []*string) *DescribeQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetName(v string) *DescribeQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *DescribeQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *DescribeQueryProcessorResponseBodyResult) SetUpdated(v int32) *DescribeQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeQueryProcessorResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *DescribeQueryProcessorResponse) SetHeaders(v map[string]*string) *DescribeQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *DescribeQueryProcessorResponse) SetStatusCode(v int32) *DescribeQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeQueryProcessorResponse) SetBody(v *DescribeQueryProcessorResponseBody) *DescribeQueryProcessorResponse {
	s.Body = v
	return s
}

type DescribeRegionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result *DescribeRegionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeRegionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponseBody) SetRequestId(v string) *DescribeRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionResponseBody) SetResult(v *DescribeRegionResponseBodyResult) *DescribeRegionResponseBody {
	s.Result = v
	return s
}

type DescribeRegionResponseBodyResult struct {
	// The configurations.
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	// The ID of the region. Valid values:
	//
	// cn-hangzhou: China (Hangzhou)
	//
	// cn-shanghai: China (Shanghai)
	//
	// cn-qingdao: China (Qingdao)
	//
	// cn-beijing: China (Beijing)
	//
	// cn-zhangjiakou: China (Zhangjiakou)
	//
	// cn-shenzhen: China (Shenzhen)
	//
	// ap-southeast-1: Singapore (Singapore)
	//
	// cn-internal: Internal Center
	//
	// cn-zhangbei-in: Internal Center (Zhangjiakou)
	//
	// us-west-1-in: Internal Center (US)
	//
	// rus-west-1-in: Internal Center (Russia)
	//
	// cn-daily: Daily Environment
	//
	// cn-test: Joint Debugging
	//
	// pre-hangzhou: China (Hangzhou)-Staging
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponseBodyResult) SetConfig(v map[string]interface{}) *DescribeRegionResponseBodyResult {
	s.Config = v
	return s
}

func (s *DescribeRegionResponseBodyResult) SetRegionId(v string) *DescribeRegionResponseBodyResult {
	s.RegionId = &v
	return s
}

type DescribeRegionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionResponse) SetHeaders(v map[string]*string) *DescribeRegionResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionResponse) SetStatusCode(v int32) *DescribeRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionResponse) SetBody(v *DescribeRegionResponseBody) *DescribeRegionResponse {
	s.Body = v
	return s
}

type DescribeRegionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result []*DescribeRegionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) SetResult(v []*DescribeRegionsResponseBodyResult) *DescribeRegionsResponseBody {
	s.Result = v
	return s
}

type DescribeRegionsResponseBodyResult struct {
	// The URL of the OpenSearch console.
	ConsoleUrl *string `json:"consoleUrl,omitempty" xml:"consoleUrl,omitempty"`
	// The endpoint of the region.
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The name of the region.
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// The ID of the region.
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DescribeRegionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyResult) SetConsoleUrl(v string) *DescribeRegionsResponseBodyResult {
	s.ConsoleUrl = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetEndpoint(v string) *DescribeRegionsResponseBodyResult {
	s.Endpoint = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetLocalName(v string) *DescribeRegionsResponseBodyResult {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyResult) SetRegionId(v string) *DescribeRegionsResponseBodyResult {
	s.RegionId = &v
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

type DescribeScheduledTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponseBody) SetRequestId(v string) *DescribeScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTaskResponseBody) SetResult(v map[string]interface{}) *DescribeScheduledTaskResponseBody {
	s.Result = v
	return s
}

type DescribeScheduledTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskResponse) SetHeaders(v map[string]*string) *DescribeScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeScheduledTaskResponse) SetStatusCode(v int32) *DescribeScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScheduledTaskResponse) SetBody(v *DescribeScheduledTaskResponseBody) *DescribeScheduledTaskResponse {
	s.Body = v
	return s
}

type DescribeSecondRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	Result *DescribeSecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBody) SetRequestId(v string) *DescribeSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecondRankResponseBody) SetResult(v *DescribeSecondRankResponseBodyResult) *DescribeSecondRankResponseBody {
	s.Result = v
	return s
}

type DescribeSecondRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression.
	//
	// You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s DescribeSecondRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponseBodyResult) SetActive(v bool) *DescribeSecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetCreated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetDescription(v string) *DescribeSecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetId(v string) *DescribeSecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsDefault(v string) *DescribeSecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetIsSys(v string) *DescribeSecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetMeta(v string) *DescribeSecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetName(v string) *DescribeSecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeSecondRankResponseBodyResult) SetUpdated(v int32) *DescribeSecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

type DescribeSecondRankResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecondRankResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecondRankResponse) SetHeaders(v map[string]*string) *DescribeSecondRankResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecondRankResponse) SetStatusCode(v int32) *DescribeSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecondRankResponse) SetBody(v *DescribeSecondRankResponseBody) *DescribeSecondRankResponse {
	s.Body = v
	return s
}

type DescribeSlowQueryStatusResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *DescribeSlowQueryStatusResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeSlowQueryStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBody) SetRequestId(v string) *DescribeSlowQueryStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBody) SetResult(v *DescribeSlowQueryStatusResponseBodyResult) *DescribeSlowQueryStatusResponseBody {
	s.Result = v
	return s
}

type DescribeSlowQueryStatusResponseBodyResult struct {
	// The ID of the application.
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The network type of the slow query optimization service. Valid values:
	//
	// *   outer: the Internet
	// *   internal: the internal network
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The status of the slow query optimization service. Valid values:
	//
	// *   enabled
	// *   disabled
	// *   n/a
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeSlowQueryStatusResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetAppGroupId(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetRegion(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Region = &v
	return s
}

func (s *DescribeSlowQueryStatusResponseBodyResult) SetStatus(v string) *DescribeSlowQueryStatusResponseBodyResult {
	s.Status = &v
	return s
}

type DescribeSlowQueryStatusResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowQueryStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowQueryStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowQueryStatusResponse) SetHeaders(v map[string]*string) *DescribeSlowQueryStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowQueryStatusResponse) SetStatusCode(v int32) *DescribeSlowQueryStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowQueryStatusResponse) SetBody(v *DescribeSlowQueryStatusResponseBody) *DescribeSlowQueryStatusResponse {
	s.Body = v
	return s
}

type DescribeUserAnalyzerRequest struct {
	// all
	With *string `json:"with,omitempty" xml:"with,omitempty"`
}

func (s DescribeUserAnalyzerRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerRequest) SetWith(v string) *DescribeUserAnalyzerRequest {
	s.With = &v
	return s
}

type DescribeUserAnalyzerResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the custom analyzer.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DescribeUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponseBody) SetRequestId(v string) *DescribeUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *DescribeUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type DescribeUserAnalyzerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserAnalyzerResponse) SetHeaders(v map[string]*string) *DescribeUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserAnalyzerResponse) SetStatusCode(v int32) *DescribeUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserAnalyzerResponse) SetBody(v *DescribeUserAnalyzerResponseBody) *DescribeUserAnalyzerResponse {
	s.Body = v
	return s
}

type DisableSlowQueryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s DisableSlowQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DisableSlowQueryResponseBody) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponseBody) SetRequestId(v string) *DisableSlowQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableSlowQueryResponseBody) SetResult(v map[string]interface{}) *DisableSlowQueryResponseBody {
	s.Result = v
	return s
}

type DisableSlowQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DisableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DisableSlowQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s DisableSlowQueryResponse) GoString() string {
	return s.String()
}

func (s *DisableSlowQueryResponse) SetHeaders(v map[string]*string) *DisableSlowQueryResponse {
	s.Headers = v
	return s
}

func (s *DisableSlowQueryResponse) SetStatusCode(v int32) *DisableSlowQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSlowQueryResponse) SetBody(v *DisableSlowQueryResponseBody) *DisableSlowQueryResponse {
	s.Body = v
	return s
}

type EnableSlowQueryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s EnableSlowQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EnableSlowQueryResponseBody) GoString() string {
	return s.String()
}

func (s *EnableSlowQueryResponseBody) SetRequestId(v string) *EnableSlowQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *EnableSlowQueryResponseBody) SetResult(v map[string]interface{}) *EnableSlowQueryResponseBody {
	s.Result = v
	return s
}

type EnableSlowQueryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *EnableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EnableSlowQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s EnableSlowQueryResponse) GoString() string {
	return s.String()
}

func (s *EnableSlowQueryResponse) SetHeaders(v map[string]*string) *EnableSlowQueryResponse {
	s.Headers = v
	return s
}

func (s *EnableSlowQueryResponse) SetStatusCode(v int32) *EnableSlowQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *EnableSlowQueryResponse) SetBody(v *EnableSlowQueryResponseBody) *EnableSlowQueryResponse {
	s.Body = v
	return s
}

type GenerateMergedTableRequest struct {
	Body *Schema `json:"body,omitempty" xml:"body,omitempty"`
	// \-
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s GenerateMergedTableRequest) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableRequest) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableRequest) SetBody(v *Schema) *GenerateMergedTableRequest {
	s.Body = v
	return s
}

func (s *GenerateMergedTableRequest) SetSpec(v string) *GenerateMergedTableRequest {
	s.Spec = &v
	return s
}

type GenerateMergedTableResponseBody struct {
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *GenerateMergedTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GenerateMergedTableResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBody) SetRequestId(v string) *GenerateMergedTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateMergedTableResponseBody) SetResult(v *GenerateMergedTableResponseBodyResult) *GenerateMergedTableResponseBody {
	s.Result = v
	return s
}

type GenerateMergedTableResponseBodyResult struct {
	FromTable  map[string]interface{} `json:"fromTable,omitempty" xml:"fromTable,omitempty"`
	MergeTable map[string]interface{} `json:"mergeTable,omitempty" xml:"mergeTable,omitempty"`
	PrimaryKey *string                `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
}

func (s GenerateMergedTableResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponseBodyResult) SetFromTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.FromTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetMergeTable(v map[string]interface{}) *GenerateMergedTableResponseBodyResult {
	s.MergeTable = v
	return s
}

func (s *GenerateMergedTableResponseBodyResult) SetPrimaryKey(v string) *GenerateMergedTableResponseBodyResult {
	s.PrimaryKey = &v
	return s
}

type GenerateMergedTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GenerateMergedTableResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GenerateMergedTableResponse) String() string {
	return tea.Prettify(s)
}

func (s GenerateMergedTableResponse) GoString() string {
	return s.String()
}

func (s *GenerateMergedTableResponse) SetHeaders(v map[string]*string) *GenerateMergedTableResponse {
	s.Headers = v
	return s
}

func (s *GenerateMergedTableResponse) SetStatusCode(v int32) *GenerateMergedTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateMergedTableResponse) SetBody(v *GenerateMergedTableResponseBody) *GenerateMergedTableResponse {
	s.Body = v
	return s
}

type GetDomainRequest struct {
	// my_app_group_name
	AppGroupIdentity *string `json:"appGroupIdentity,omitempty" xml:"appGroupIdentity,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetAppGroupIdentity(v string) *GetDomainRequest {
	s.AppGroupIdentity = &v
	return s
}

type GetDomainResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return results.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetDomainResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainResponseBody) SetRequestId(v string) *GetDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDomainResponseBody) SetResult(v map[string]interface{}) *GetDomainResponseBody {
	s.Result = v
	return s
}

type GetDomainResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *GetDomainResponseBody) *GetDomainResponse {
	s.Body = v
	return s
}

type GetFunctionCurrentVersionRequest struct {
	// The category. By default, this parameter is left empty.
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry. By default, this parameter is left empty, which indicates General-purpose Edition.
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The type of the feature. Valid values:
	//
	// *   PAAS. This is the default value.
	// *   SAAS.
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model. The following features correspond to different model types:
	//
	// *   CTR model: tf_checkpoint
	// *   Popularity model: pop
	// *   Category model: offline_inference
	// *   Hotword model: offline_inference
	// *   Shading model: offline_inference
	// *   Drop-down suggestion model: offline_inference
	// *   Word segmentation model: text
	// *   Word weight model: tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
}

func (s GetFunctionCurrentVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionRequest) SetCategory(v string) *GetFunctionCurrentVersionRequest {
	s.Category = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetDomain(v string) *GetFunctionCurrentVersionRequest {
	s.Domain = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetFunctionType(v string) *GetFunctionCurrentVersionRequest {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionRequest) SetModelType(v string) *GetFunctionCurrentVersionRequest {
	s.ModelType = &v
	return s
}

type GetFunctionCurrentVersionResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionCurrentVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBody) SetCode(v string) *GetFunctionCurrentVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetHttpCode(v int64) *GetFunctionCurrentVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetLatency(v int64) *GetFunctionCurrentVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetMessage(v string) *GetFunctionCurrentVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetRequestId(v string) *GetFunctionCurrentVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetResult(v *GetFunctionCurrentVersionResponseBodyResult) *GetFunctionCurrentVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBody) SetStatus(v string) *GetFunctionCurrentVersionResponseBody {
	s.Status = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResult struct {
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature. Valid values:
	//
	// *   PAAS
	// *   SAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The type of the model.
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The configuration information about the instance.
	VersionConfig *GetFunctionCurrentVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	// The ID of the version.
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The name of the version.
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetModelType(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionConfig(v *GetFunctionCurrentVersionResponseBodyResultVersionConfig) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResult) SetVersionName(v string) *GetFunctionCurrentVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfig struct {
	// The parameters that are used to create the instance.
	CreateParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used to use the instance online.
	UsageParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionCurrentVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends struct {
	// The condition.
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The dependency.
	Dependency *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	// The description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

type GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionCurrentVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

type GetFunctionCurrentVersionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionCurrentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionCurrentVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionCurrentVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionCurrentVersionResponse) SetHeaders(v map[string]*string) *GetFunctionCurrentVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionCurrentVersionResponse) SetStatusCode(v int32) *GetFunctionCurrentVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionCurrentVersionResponse) SetBody(v *GetFunctionCurrentVersionResponseBody) *GetFunctionCurrentVersionResponse {
	s.Body = v
	return s
}

type GetFunctionDefaultInstanceResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The default running time.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionDefaultInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBody) SetCode(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetFunctionName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBody {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetLatency(v int64) *GetFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetMessage(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetRequestId(v string) *GetFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetResult(v *GetFunctionDefaultInstanceResponseBodyResult) *GetFunctionDefaultInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionDefaultInstanceResponseBody) SetStatus(v string) *GetFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

type GetFunctionDefaultInstanceResponseBodyResult struct {
	// The default instance name.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s GetFunctionDefaultInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionDefaultInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

type GetFunctionDefaultInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionDefaultInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) SetStatusCode(v int32) *GetFunctionDefaultInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionDefaultInstanceResponse) SetBody(v *GetFunctionDefaultInstanceResponseBody) *GetFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

type GetFunctionInstanceRequest struct {
	// Specifies the richness of returned information. Valid values:
	//
	// *   simple: displays only the basic information.
	// *   normal: displays information such as createParameters and cron. This is the default value.
	// *   detail: returns the details of the training task.
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
}

func (s GetFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceRequest) SetOutput(v string) *GetFunctionInstanceRequest {
	s.Output = &v
	return s
}

type GetFunctionInstanceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the instance.
	Result *GetFunctionInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBody) SetCode(v string) *GetFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetHttpCode(v int64) *GetFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetLatency(v int64) *GetFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetMessage(v string) *GetFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetRequestId(v string) *GetFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetResult(v *GetFunctionInstanceResponseBodyResult) *GetFunctionInstanceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionInstanceResponseBody) SetStatus(v string) *GetFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type GetFunctionInstanceResponseBodyResult struct {
	// The information about the instance.
	Belongs *GetFunctionInstanceResponseBodyResultBelongs `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	// The parameters that are used to create the instance.
	CreateParameters []*GetFunctionInstanceResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The time when the task was created. Unit: milliseconds.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression used to schedule training, in the format of (Minutes Hours DayofMonth Month DayofWeek). If the value is empty, it indicates that no periodic training is performed.
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The description of the instance.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string.
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// *   user: The instance is created by user.
	// *   builtin: The instance is created by the system.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	// 2.  available: Models can be used.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the training task. This parameter is not displayed if no task is available.
	Task *GetFunctionInstanceResponseBodyResultTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The parameters that are used.
	UsageParameters []*GetFunctionInstanceResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResult) SetBelongs(v *GetFunctionInstanceResponseBodyResultBelongs) *GetFunctionInstanceResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateParameters(v []*GetFunctionInstanceResponseBodyResultCreateParameters) *GetFunctionInstanceResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCreateTime(v int64) *GetFunctionInstanceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetCron(v string) *GetFunctionInstanceResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetDescription(v string) *GetFunctionInstanceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetExtendInfo(v string) *GetFunctionInstanceResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionName(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetFunctionType(v string) *GetFunctionInstanceResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetInstanceName(v string) *GetFunctionInstanceResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetModelType(v string) *GetFunctionInstanceResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetSource(v string) *GetFunctionInstanceResponseBodyResult {
	s.Source = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetStatus(v string) *GetFunctionInstanceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetTask(v *GetFunctionInstanceResponseBodyResultTask) *GetFunctionInstanceResponseBodyResult {
	s.Task = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetUsageParameters(v []*GetFunctionInstanceResponseBodyResultUsageParameters) *GetFunctionInstanceResponseBodyResult {
	s.UsageParameters = v
	return s
}

func (s *GetFunctionInstanceResponseBodyResult) SetVersionId(v int64) *GetFunctionInstanceResponseBodyResult {
	s.VersionId = &v
	return s
}

type GetFunctionInstanceResponseBodyResultBelongs struct {
	// The category.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultBelongs) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetCategory(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetDomain(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultBelongs) SetLanguage(v string) *GetFunctionInstanceResponseBodyResultBelongs {
	s.Language = &v
	return s
}

type GetFunctionInstanceResponseBodyResultCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultCreateParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

type GetFunctionInstanceResponseBodyResultTask struct {
	// The status of the task. Valid values:
	//
	// *   success: succeeded
	// *   failed: failed
	// *   untrained: to be trained
	// *   pending: being scheduled
	// *   running: being trained
	DagStatus *string `json:"DagStatus,omitempty" xml:"DagStatus,omitempty"`
	// The time consumed for the most recent run, in milliseconds.
	LastRunTime *int64 `json:"LastRunTime,omitempty" xml:"LastRunTime,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultTask) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultTask) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetDagStatus(v string) *GetFunctionInstanceResponseBodyResultTask {
	s.DagStatus = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultTask) SetLastRunTime(v int64) *GetFunctionInstanceResponseBodyResultTask {
	s.LastRunTime = &v
	return s
}

type GetFunctionInstanceResponseBodyResultUsageParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetFunctionInstanceResponseBodyResultUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponseBodyResultUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) SetName(v string) *GetFunctionInstanceResponseBodyResultUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionInstanceResponseBodyResultUsageParameters) SetValue(v string) *GetFunctionInstanceResponseBodyResultUsageParameters {
	s.Value = &v
	return s
}

type GetFunctionInstanceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionInstanceResponse) SetHeaders(v map[string]*string) *GetFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionInstanceResponse) SetStatusCode(v int32) *GetFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionInstanceResponse) SetBody(v *GetFunctionInstanceResponseBody) *GetFunctionInstanceResponse {
	s.Body = v
	return s
}

type GetFunctionTaskResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponseBody) SetCode(v string) *GetFunctionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetHttpCode(v int64) *GetFunctionTaskResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetLatency(v int64) *GetFunctionTaskResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetMessage(v string) *GetFunctionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetRequestId(v string) *GetFunctionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionTaskResponseBody) SetResult(v *GetFunctionTaskResponseBodyResult) *GetFunctionTaskResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionTaskResponseBody) SetStatus(v string) *GetFunctionTaskResponseBody {
	s.Status = &v
	return s
}

type GetFunctionTaskResponseBodyResult struct {
	// The timestamp that indicates the end time of the task. Unit: milliseconds.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information, which is a JSON string.
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The number of iterations.
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The progress. 90 indicates 90%.
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the task.
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// *   success
	// *   failed
	// *   running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionTaskResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionTaskResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponseBodyResult) SetEndTime(v int64) *GetFunctionTaskResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetExtendInfo(v string) *GetFunctionTaskResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetFunctionName(v string) *GetFunctionTaskResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetGeneration(v string) *GetFunctionTaskResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetProgress(v int64) *GetFunctionTaskResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetRunId(v string) *GetFunctionTaskResponseBodyResult {
	s.RunId = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetStartTime(v int64) *GetFunctionTaskResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *GetFunctionTaskResponseBodyResult) SetStatus(v string) *GetFunctionTaskResponseBodyResult {
	s.Status = &v
	return s
}

type GetFunctionTaskResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionTaskResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionTaskResponse) SetHeaders(v map[string]*string) *GetFunctionTaskResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionTaskResponse) SetStatusCode(v int32) *GetFunctionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionTaskResponse) SetBody(v *GetFunctionTaskResponseBody) *GetFunctionTaskResponse {
	s.Body = v
	return s
}

type GetFunctionVersionResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The maximum duration for which a task can be executed.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result body.
	Result *GetFunctionVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBody) SetCode(v string) *GetFunctionVersionResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetHttpCode(v int64) *GetFunctionVersionResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetLatency(v int64) *GetFunctionVersionResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetMessage(v string) *GetFunctionVersionResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetRequestId(v string) *GetFunctionVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionVersionResponseBody) SetResult(v *GetFunctionVersionResponseBodyResult) *GetFunctionVersionResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionVersionResponseBody) SetStatus(v string) *GetFunctionVersionResponseBody {
	s.Status = &v
	return s
}

type GetFunctionVersionResponseBodyResult struct {
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature. Valid values:
	//
	// *   PAAS
	// *   SAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The type of the model.
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The configuration information.
	VersionConfig *GetFunctionVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	// The ID of the version.
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The name of the version.
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s GetFunctionVersionResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionName(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetFunctionType(v string) *GetFunctionVersionResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetModelType(v string) *GetFunctionVersionResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionConfig(v *GetFunctionVersionResponseBodyResultVersionConfig) *GetFunctionVersionResponseBodyResult {
	s.VersionConfig = v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionId(v int64) *GetFunctionVersionResponseBodyResult {
	s.VersionId = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResult) SetVersionName(v string) *GetFunctionVersionResponseBodyResult {
	s.VersionName = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfig struct {
	// The parameters that are used to create the instance.
	CreateParameters []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used during online use of the instance.
	UsageParameters []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfig) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetCreateParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.CreateParameters = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetDepends(v []*GetFunctionVersionResponseBodyResultVersionConfigDepends) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.Depends = v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfig) SetUsageParameters(v []*GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) *GetFunctionVersionResponseBodyResultVersionConfig {
	s.UsageParameters = v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigCreateParameters {
	s.Required = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigDepends struct {
	// The condition.
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The dependency.
	Dependency *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	// The description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigDepends) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetCondition(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Condition = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDependency(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Dependency = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigDepends) SetDescription(v string) *GetFunctionVersionResponseBodyResultVersionConfigDepends {
	s.Description = &v
	return s
}

type GetFunctionVersionResponseBodyResultVersionConfigUsageParameters struct {
	// The name of the instance.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	Required *string `json:"Required,omitempty" xml:"Required,omitempty"`
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetName(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Name = &v
	return s
}

func (s *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters) SetRequired(v string) *GetFunctionVersionResponseBodyResultVersionConfigUsageParameters {
	s.Required = &v
	return s
}

type GetFunctionVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetFunctionVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFunctionVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionVersionResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionVersionResponse) SetHeaders(v map[string]*string) *GetFunctionVersionResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionVersionResponse) SetStatusCode(v int32) *GetFunctionVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionVersionResponse) SetBody(v *GetFunctionVersionResponseBody) *GetFunctionVersionResponse {
	s.Body = v
	return s
}

type GetModelReportResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s GetModelReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetModelReportResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelReportResponseBody) SetRequestId(v string) *GetModelReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelReportResponseBody) SetResult(v map[string]interface{}) *GetModelReportResponseBody {
	s.Result = v
	return s
}

type GetModelReportResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetModelReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetModelReportResponse) String() string {
	return tea.Prettify(s)
}

func (s GetModelReportResponse) GoString() string {
	return s.String()
}

func (s *GetModelReportResponse) SetHeaders(v map[string]*string) *GetModelReportResponse {
	s.Headers = v
	return s
}

func (s *GetModelReportResponse) SetStatusCode(v int32) *GetModelReportResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModelReportResponse) SetBody(v *GetModelReportResponseBody) *GetModelReportResponse {
	s.Body = v
	return s
}

type GetScriptFileNamesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The files of the script.
	Result []*GetScriptFileNamesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s GetScriptFileNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBody) SetRequestId(v string) *GetScriptFileNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScriptFileNamesResponseBody) SetResult(v []*GetScriptFileNamesResponseBodyResult) *GetScriptFileNamesResponseBody {
	s.Result = v
	return s
}

type GetScriptFileNamesResponseBodyResult struct {
	// The time when the script file was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the script file.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The time when the script file was last modified.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	PathName   *string `json:"pathName,omitempty" xml:"pathName,omitempty"`
}

func (s GetScriptFileNamesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponseBodyResult) SetCreateTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetFileName(v string) *GetScriptFileNamesResponseBodyResult {
	s.FileName = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetModifyTime(v string) *GetScriptFileNamesResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetScriptFileNamesResponseBodyResult) SetPathName(v string) *GetScriptFileNamesResponseBodyResult {
	s.PathName = &v
	return s
}

type GetScriptFileNamesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetScriptFileNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetScriptFileNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetScriptFileNamesResponse) GoString() string {
	return s.String()
}

func (s *GetScriptFileNamesResponse) SetHeaders(v map[string]*string) *GetScriptFileNamesResponse {
	s.Headers = v
	return s
}

func (s *GetScriptFileNamesResponse) SetStatusCode(v int32) *GetScriptFileNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScriptFileNamesResponse) SetBody(v *GetScriptFileNamesResponseBody) *GetScriptFileNamesResponse {
	s.Body = v
	return s
}

type GetSearchStrategyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponseBody) SetRequestId(v string) *GetSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type GetSearchStrategyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *GetSearchStrategyResponse) SetHeaders(v map[string]*string) *GetSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *GetSearchStrategyResponse) SetStatusCode(v int32) *GetSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSearchStrategyResponse) SetBody(v *GetSearchStrategyResponseBody) *GetSearchStrategyResponse {
	s.Body = v
	return s
}

type GetSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the script
	Result *GetSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBody) SetRequestId(v string) *GetSortScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptResponseBody) SetResult(v *GetSortScriptResponseBodyResult) *GetSortScriptResponseBody {
	s.Result = v
	return s
}

type GetSortScriptResponseBodyResult struct {
	// The time when the script was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The sort phase to which the script applies.
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The status of the script. For more information, see the Script status table.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the script.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetSortScriptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponseBodyResult) SetCreateTime(v string) *GetSortScriptResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetModifyTime(v string) *GetSortScriptResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetScope(v string) *GetSortScriptResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetStatus(v string) *GetSortScriptResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetSortScriptResponseBodyResult) SetType(v string) *GetSortScriptResponseBodyResult {
	s.Type = &v
	return s
}

type GetSortScriptResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptResponse) SetHeaders(v map[string]*string) *GetSortScriptResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptResponse) SetStatusCode(v int32) *GetSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSortScriptResponse) SetBody(v *GetSortScriptResponseBody) *GetSortScriptResponse {
	s.Body = v
	return s
}

type GetSortScriptFileResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The content of the sort script.
	Result *GetSortScriptFileResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s GetSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBody) SetRequestId(v string) *GetSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSortScriptFileResponseBody) SetResult(v *GetSortScriptFileResponseBodyResult) *GetSortScriptFileResponseBody {
	s.Result = v
	return s
}

type GetSortScriptFileResponseBodyResult struct {
	// The script content that is encoded in the Base64 format.
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The time when the script was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The last time when the script was last modified.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The version of the script.
	Version *int64 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetSortScriptFileResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponseBodyResult) SetContent(v string) *GetSortScriptFileResponseBodyResult {
	s.Content = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetCreateTime(v string) *GetSortScriptFileResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetModifyTime(v string) *GetSortScriptFileResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetSortScriptFileResponseBodyResult) SetVersion(v int64) *GetSortScriptFileResponseBodyResult {
	s.Version = &v
	return s
}

type GetSortScriptFileResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *GetSortScriptFileResponse) SetHeaders(v map[string]*string) *GetSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *GetSortScriptFileResponse) SetStatusCode(v int32) *GetSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSortScriptFileResponse) SetBody(v *GetSortScriptFileResponseBody) *GetSortScriptFileResponse {
	s.Body = v
	return s
}

type ListABTestExperimentsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the tests.
	//
	// For more information, see [ABTestExperiment](~~173617~~).
	Result []*ListABTestExperimentsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestExperimentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBody) SetRequestId(v string) *ListABTestExperimentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestExperimentsResponseBody) SetResult(v []*ListABTestExperimentsResponseBodyResult) *ListABTestExperimentsResponseBody {
	s.Result = v
	return s
}

type ListABTestExperimentsResponseBodyResult struct {
	// The time when the test was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// *   true: in effect
	// *   false: not in effect
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the test.
	//
	// Valid values: \[0,100].
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListABTestExperimentsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponseBodyResult) SetCreated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetId(v string) *ListABTestExperimentsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetName(v string) *ListABTestExperimentsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetOnline(v bool) *ListABTestExperimentsResponseBodyResult {
	s.Online = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetParams(v map[string]interface{}) *ListABTestExperimentsResponseBodyResult {
	s.Params = v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetTraffic(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *ListABTestExperimentsResponseBodyResult) SetUpdated(v int32) *ListABTestExperimentsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListABTestExperimentsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListABTestExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestExperimentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestExperimentsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestExperimentsResponse) SetHeaders(v map[string]*string) *ListABTestExperimentsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestExperimentsResponse) SetStatusCode(v int32) *ListABTestExperimentsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestExperimentsResponse) SetBody(v *ListABTestExperimentsResponseBody) *ListABTestExperimentsResponse {
	s.Body = v
	return s
}

type ListABTestFixedFlowDividersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The queried whitelists.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestFixedFlowDividersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponseBody) SetRequestId(v string) *ListABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestFixedFlowDividersResponseBody) SetResult(v []*string) *ListABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

type ListABTestFixedFlowDividersResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestFixedFlowDividersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *ListABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *ListABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) SetStatusCode(v int32) *ListABTestFixedFlowDividersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestFixedFlowDividersResponse) SetBody(v *ListABTestFixedFlowDividersResponseBody) *ListABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

type ListABTestGroupsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The test groups.
	//
	// For more information, see [ABTestGroup](~~178935~~).
	Result []*ListABTestGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBody) SetRequestId(v string) *ListABTestGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestGroupsResponseBody) SetResult(v []*ListABTestGroupsResponseBodyResult) *ListABTestGroupsResponseBody {
	s.Result = v
	return s
}

type ListABTestGroupsResponseBodyResult struct {
	// The time when the test group was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListABTestGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponseBodyResult) SetCreated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetId(v string) *ListABTestGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetName(v string) *ListABTestGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetStatus(v int32) *ListABTestGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestGroupsResponseBodyResult) SetUpdated(v int32) *ListABTestGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListABTestGroupsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListABTestGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListABTestGroupsResponse) SetHeaders(v map[string]*string) *ListABTestGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListABTestGroupsResponse) SetStatusCode(v int32) *ListABTestGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestGroupsResponse) SetBody(v *ListABTestGroupsResponseBody) *ListABTestGroupsResponse {
	s.Body = v
	return s
}

type ListABTestScenesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenarios.
	//
	// For more information, see [ABTestScene](~~173618~~).
	Result []*ListABTestScenesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponseBody) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBody) SetRequestId(v string) *ListABTestScenesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListABTestScenesResponseBody) SetResult(v []*ListABTestScenesResponseBodyResult) *ListABTestScenesResponseBody {
	s.Result = v
	return s
}

type ListABTestScenesResponseBodyResult struct {
	// The time when the test scenario was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The name of the test scenario.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListABTestScenesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponseBodyResult) SetCreated(v int32) *ListABTestScenesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetId(v string) *ListABTestScenesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetName(v string) *ListABTestScenesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetStatus(v int32) *ListABTestScenesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetUpdated(v int32) *ListABTestScenesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListABTestScenesResponseBodyResult) SetValues(v []*string) *ListABTestScenesResponseBodyResult {
	s.Values = v
	return s
}

type ListABTestScenesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListABTestScenesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListABTestScenesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListABTestScenesResponse) GoString() string {
	return s.String()
}

func (s *ListABTestScenesResponse) SetHeaders(v map[string]*string) *ListABTestScenesResponse {
	s.Headers = v
	return s
}

func (s *ListABTestScenesResponse) SetStatusCode(v int32) *ListABTestScenesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListABTestScenesResponse) SetBody(v *ListABTestScenesResponseBody) *ListABTestScenesResponse {
	s.Body = v
	return s
}

type ListAppGroupsRequest struct {
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// ""
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 0
	SortBy *int32                      `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	Tags   []*ListAppGroupsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAppGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupsRequest) SetInstanceId(v string) *ListAppGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsRequest) SetName(v string) *ListAppGroupsRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageNumber(v int32) *ListAppGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupsRequest) SetPageSize(v int32) *ListAppGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupsRequest) SetResourceGroupId(v string) *ListAppGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsRequest) SetSortBy(v int32) *ListAppGroupsRequest {
	s.SortBy = &v
	return s
}

func (s *ListAppGroupsRequest) SetTags(v []*ListAppGroupsRequestTags) *ListAppGroupsRequest {
	s.Tags = v
	return s
}

func (s *ListAppGroupsRequest) SetType(v string) *ListAppGroupsRequest {
	s.Type = &v
	return s
}

type ListAppGroupsRequestTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppGroupsRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsRequestTags) GoString() string {
	return s.String()
}

func (s *ListAppGroupsRequestTags) SetKey(v string) *ListAppGroupsRequestTags {
	s.Key = &v
	return s
}

func (s *ListAppGroupsRequestTags) SetValue(v string) *ListAppGroupsRequestTags {
	s.Value = &v
	return s
}

type ListAppGroupsShrinkRequest struct {
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// ""
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// 0
	SortBy     *int32  `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListAppGroupsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAppGroupsShrinkRequest) SetInstanceId(v string) *ListAppGroupsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetName(v string) *ListAppGroupsShrinkRequest {
	s.Name = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetPageNumber(v int32) *ListAppGroupsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetPageSize(v int32) *ListAppGroupsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetResourceGroupId(v string) *ListAppGroupsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetSortBy(v int32) *ListAppGroupsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetTagsShrink(v string) *ListAppGroupsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListAppGroupsShrinkRequest) SetType(v string) *ListAppGroupsShrinkRequest {
	s.Type = &v
	return s
}

type ListAppGroupsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each application.
	//
	// For more information, see [AppGroup](~~170000~~).
	Result []*ListAppGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListAppGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBody) SetRequestId(v string) *ListAppGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAppGroupsResponseBody) SetResult(v []*ListAppGroupsResponseBodyResult) *ListAppGroupsResponseBody {
	s.Result = v
	return s
}

func (s *ListAppGroupsResponseBody) SetTotalCount(v int32) *ListAppGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListAppGroupsResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// domain
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	//
	// For more information, see [Quota](https://www.alibabacloud.com/help/doc-detail/170001.htm).
	Quota *ListAppGroupsResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32                                 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	Tags         []*ListAppGroupsResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListAppGroupsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResult) SetChargeType(v string) *ListAppGroupsResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetChargingWay(v int32) *ListAppGroupsResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCommodityCode(v string) *ListAppGroupsResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCreated(v int32) *ListAppGroupsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetCurrentVersion(v string) *ListAppGroupsResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDescription(v string) *ListAppGroupsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetDomain(v string) *ListAppGroupsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetExpireOn(v string) *ListAppGroupsResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ListAppGroupsResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetId(v string) *ListAppGroupsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetInstanceId(v string) *ListAppGroupsResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockMode(v string) *ListAppGroupsResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetLockedByExpiration(v int32) *ListAppGroupsResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetName(v string) *ListAppGroupsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProcessingOrderId(v string) *ListAppGroupsResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProduced(v int32) *ListAppGroupsResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetProjectId(v string) *ListAppGroupsResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetQuota(v *ListAppGroupsResponseBodyResultQuota) *ListAppGroupsResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ListAppGroupsResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetStatus(v string) *ListAppGroupsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetSwitchedTime(v int32) *ListAppGroupsResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetTags(v []*ListAppGroupsResponseBodyResultTags) *ListAppGroupsResponseBodyResult {
	s.Tags = v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetType(v string) *ListAppGroupsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetUpdated(v int32) *ListAppGroupsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListAppGroupsResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ListAppGroupsResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResultQuota) SetComputeResource(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetDocSize(v int32) *ListAppGroupsResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultQuota) SetSpec(v string) *ListAppGroupsResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ListAppGroupsResponseBodyResultTags struct {
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListAppGroupsResponseBodyResultTags) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponseBodyResultTags) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponseBodyResultTags) SetKey(v string) *ListAppGroupsResponseBodyResultTags {
	s.Key = &v
	return s
}

func (s *ListAppGroupsResponseBodyResultTags) SetValue(v string) *ListAppGroupsResponseBodyResultTags {
	s.Value = &v
	return s
}

type ListAppGroupsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAppGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAppGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAppGroupsResponse) SetHeaders(v map[string]*string) *ListAppGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAppGroupsResponse) SetStatusCode(v int32) *ListAppGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppGroupsResponse) SetBody(v *ListAppGroupsResponseBody) *ListAppGroupsResponse {
	s.Body = v
	return s
}

type ListAppsRequest struct {
	// true
	Group *bool `json:"group,omitempty" xml:"group,omitempty"`
	// 0
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// 0
	Size *int32 `json:"size,omitempty" xml:"size,omitempty"`
}

func (s ListAppsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAppsRequest) GoString() string {
	return s.String()
}

func (s *ListAppsRequest) SetGroup(v bool) *ListAppsRequest {
	s.Group = &v
	return s
}

func (s *ListAppsRequest) SetPage(v int32) *ListAppsRequest {
	s.Page = &v
	return s
}

func (s *ListAppsRequest) SetSize(v int32) *ListAppsRequest {
	s.Size = &v
	return s
}

type ListAppsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ListAppsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAppsResponse) GoString() string {
	return s.String()
}

func (s *ListAppsResponse) SetHeaders(v map[string]*string) *ListAppsResponse {
	s.Headers = v
	return s
}

func (s *ListAppsResponse) SetStatusCode(v int32) *ListAppsResponse {
	s.StatusCode = &v
	return s
}

type ListDataCollectionsRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListDataCollectionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsRequest) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsRequest) SetPageNumber(v int32) *ListDataCollectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataCollectionsRequest) SetPageSize(v int32) *ListDataCollectionsRequest {
	s.PageSize = &v
	return s
}

type ListDataCollectionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the data collection tasks.
	//
	// For more information, see [DataCollection](~~173605~~).
	Result []*ListDataCollectionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the returned data collection tasks.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCollectionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBody) SetRequestId(v string) *ListDataCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCollectionsResponseBody) SetResult(v []*ListDataCollectionsResponseBodyResult) *ListDataCollectionsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataCollectionsResponseBody) SetTotalCount(v int32) *ListDataCollectionsResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataCollectionsResponseBodyResult struct {
	// The time when the data collection task was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the data that is collected by the task. Valid values:
	//
	// *   behavior: behavioral data
	// *   item_info: project data
	// *   industry_specific: industry-specific data
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	// The ID of the data collection task.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry to which the data collection task applies. Valid values:
	//
	// *   general
	// *   ecommerce
	IndustryName *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	// The name of the data collection task.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data collection task. Valid values:
	//
	// *   0: disabled
	// *   1: being enabled
	// *   2: enabled
	// *   3: failed to be enabled
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the sundial.
	SundialId *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	// The type of the data source. Valid values:
	//
	// *   server
	// *   web
	// *   app
	//
	// Note: Only server is supported.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the data collection task was updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListDataCollectionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBodyResult) SetCreated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetDataCollectionType(v string) *ListDataCollectionsResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetId(v string) *ListDataCollectionsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetIndustryName(v string) *ListDataCollectionsResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetName(v string) *ListDataCollectionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetStatus(v int32) *ListDataCollectionsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetSundialId(v string) *ListDataCollectionsResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetType(v string) *ListDataCollectionsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetUpdated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListDataCollectionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataCollectionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataCollectionsResponse) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponse) SetHeaders(v map[string]*string) *ListDataCollectionsResponse {
	s.Headers = v
	return s
}

func (s *ListDataCollectionsResponse) SetStatusCode(v int32) *ListDataCollectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCollectionsResponse) SetBody(v *ListDataCollectionsResponseBody) *ListDataCollectionsResponse {
	s.Body = v
	return s
}

type ListDataSourceTableFieldsRequest struct {
	// {}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// Whether to return the original field type of the data source
	RawType *bool `json:"rawType,omitempty" xml:"rawType,omitempty"`
}

func (s ListDataSourceTableFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsRequest) SetParams(v string) *ListDataSourceTableFieldsRequest {
	s.Params = &v
	return s
}

func (s *ListDataSourceTableFieldsRequest) SetRawType(v bool) *ListDataSourceTableFieldsRequest {
	s.RawType = &v
	return s
}

type ListDataSourceTableFieldsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListDataSourceTableFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponseBody) SetRequestId(v string) *ListDataSourceTableFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTableFieldsResponseBody) SetResult(v map[string]interface{}) *ListDataSourceTableFieldsResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTableFieldsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourceTableFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceTableFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTableFieldsResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTableFieldsResponse) SetHeaders(v map[string]*string) *ListDataSourceTableFieldsResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTableFieldsResponse) SetStatusCode(v int32) *ListDataSourceTableFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTableFieldsResponse) SetBody(v *ListDataSourceTableFieldsResponseBody) *ListDataSourceTableFieldsResponse {
	s.Body = v
	return s
}

type ListDataSourceTablesRequest struct {
	// N/A
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
}

func (s ListDataSourceTablesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesRequest) SetParams(v string) *ListDataSourceTablesRequest {
	s.Params = &v
	return s
}

type ListDataSourceTablesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data tables.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListDataSourceTablesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponseBody) SetRequestId(v string) *ListDataSourceTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourceTablesResponseBody) SetResult(v []*string) *ListDataSourceTablesResponseBody {
	s.Result = v
	return s
}

type ListDataSourceTablesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourceTablesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourceTablesResponse) SetHeaders(v map[string]*string) *ListDataSourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourceTablesResponse) SetStatusCode(v int32) *ListDataSourceTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourceTablesResponse) SetBody(v *ListDataSourceTablesResponseBody) *ListDataSourceTablesResponse {
	s.Body = v
	return s
}

type ListFirstRanksResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each rough sort expression.
	//
	// For more information, see [FirstRank](~~170007~~).
	Result []*ListFirstRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListFirstRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBody) SetRequestId(v string) *ListFirstRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFirstRanksResponseBody) SetResult(v []*ListFirstRanksResponseBodyResult) *ListFirstRanksResponseBody {
	s.Result = v
	return s
}

type ListFirstRanksResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active  *bool  `json:"active,omitempty" xml:"active,omitempty"`
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*ListFirstRanksResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListFirstRanksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResult) SetActive(v bool) *ListFirstRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetCreated(v int32) *ListFirstRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetDescription(v string) *ListFirstRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetMeta(v []*ListFirstRanksResponseBodyResultMeta) *ListFirstRanksResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetName(v string) *ListFirstRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListFirstRanksResponseBodyResult) SetUpdated(v int32) *ListFirstRanksResponseBodyResult {
	s.Updated = &v
	return s
}

type ListFirstRanksResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// For more information, see [Rough sort functions](~~180765~~).
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see [Rough sort functions](~~180765~~).
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: \[-100000,100000] (excluding 0).
	Weight *int32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ListFirstRanksResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponseBodyResultMeta) SetArg(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetAttribute(v string) *ListFirstRanksResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ListFirstRanksResponseBodyResultMeta) SetWeight(v int32) *ListFirstRanksResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type ListFirstRanksResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFirstRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFirstRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFirstRanksResponse) GoString() string {
	return s.String()
}

func (s *ListFirstRanksResponse) SetHeaders(v map[string]*string) *ListFirstRanksResponse {
	s.Headers = v
	return s
}

func (s *ListFirstRanksResponse) SetStatusCode(v int32) *ListFirstRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFirstRanksResponse) SetBody(v *ListFirstRanksResponseBody) *ListFirstRanksResponse {
	s.Body = v
	return s
}

type ListFunctionInstancesRequest struct {
	// The type of the feature.
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model.
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The richness of the returned information. Valid values:
	//
	// *   normal: displays information such as createParameters and cron. This is the default value.
	// *   simple: displays only the basic information.
	// *   detail: returns the details of the training task.
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The number of the page to return. Default value: 1.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// How the instance is created. Valid values:
	//
	// *   builtin: The instance is created by system.
	// *   user: The instance is created by user. This is the default value.
	// *   all: all instances
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListFunctionInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesRequest) SetFunctionType(v string) *ListFunctionInstancesRequest {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetModelType(v string) *ListFunctionInstancesRequest {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetOutput(v string) *ListFunctionInstancesRequest {
	s.Output = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageNumber(v int32) *ListFunctionInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetPageSize(v int32) *ListFunctionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionInstancesRequest) SetSource(v string) *ListFunctionInstancesRequest {
	s.Source = &v
	return s
}

type ListFunctionInstancesResponseBody struct {
	// The error code. If no error occurs, the parameter is left empty.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, the parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the instances.
	Result []*ListFunctionInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries returned.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBody) SetCode(v string) *ListFunctionInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetHttpCode(v int64) *ListFunctionInstancesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetLatency(v int64) *ListFunctionInstancesResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetMessage(v string) *ListFunctionInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetRequestId(v string) *ListFunctionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetResult(v []*ListFunctionInstancesResponseBodyResult) *ListFunctionInstancesResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetStatus(v string) *ListFunctionInstancesResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBody) SetTotalCount(v int64) *ListFunctionInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFunctionInstancesResponseBodyResult struct {
	// The information about the instance.
	Belongs *ListFunctionInstancesResponseBodyResultBelongs `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	// The parameters of the instance.
	CreateParameters []*ListFunctionInstancesResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The time when the instance was created.
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression used to schedule training, in the format of (Minutes Hours DayofMonth Month DayofWeek). If the value is empty, it indicates that no periodic training is performed.
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string. It includes model evaluation information and error information.
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// *   user: The instance is created by user.
	// *   builtin: The instance is created by system.
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	// 2.  available: Models can be used.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters that are used.
	UsageParameters []*ListFunctionInstancesResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResult) SetBelongs(v *ListFunctionInstancesResponseBodyResultBelongs) *ListFunctionInstancesResponseBodyResult {
	s.Belongs = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateParameters(v []*ListFunctionInstancesResponseBodyResultCreateParameters) *ListFunctionInstancesResponseBodyResult {
	s.CreateParameters = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCreateTime(v int64) *ListFunctionInstancesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetCron(v string) *ListFunctionInstancesResponseBodyResult {
	s.Cron = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetDescription(v string) *ListFunctionInstancesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetExtendInfo(v string) *ListFunctionInstancesResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionName(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetFunctionType(v string) *ListFunctionInstancesResponseBodyResult {
	s.FunctionType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetInstanceName(v string) *ListFunctionInstancesResponseBodyResult {
	s.InstanceName = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetModelType(v string) *ListFunctionInstancesResponseBodyResult {
	s.ModelType = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetSource(v string) *ListFunctionInstancesResponseBodyResult {
	s.Source = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetStatus(v string) *ListFunctionInstancesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetUsageParameters(v []*ListFunctionInstancesResponseBodyResultUsageParameters) *ListFunctionInstancesResponseBodyResult {
	s.UsageParameters = v
	return s
}

func (s *ListFunctionInstancesResponseBodyResult) SetVersionId(v int64) *ListFunctionInstancesResponseBodyResult {
	s.VersionId = &v
	return s
}

type ListFunctionInstancesResponseBodyResultBelongs struct {
	// The category.
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultBelongs) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultBelongs) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetCategory(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Category = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetDomain(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Domain = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultBelongs) SetLanguage(v string) *ListFunctionInstancesResponseBodyResultBelongs {
	s.Language = &v
	return s
}

type ListFunctionInstancesResponseBodyResultCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultCreateParameters) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetName(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Name = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultCreateParameters) SetValue(v string) *ListFunctionInstancesResponseBodyResultCreateParameters {
	s.Value = &v
	return s
}

type ListFunctionInstancesResponseBodyResultUsageParameters struct {
	// The name of the parameter.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFunctionInstancesResponseBodyResultUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponseBodyResultUsageParameters) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) SetName(v string) *ListFunctionInstancesResponseBodyResultUsageParameters {
	s.Name = &v
	return s
}

func (s *ListFunctionInstancesResponseBodyResultUsageParameters) SetValue(v string) *ListFunctionInstancesResponseBodyResultUsageParameters {
	s.Value = &v
	return s
}

type ListFunctionInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionInstancesResponse) SetHeaders(v map[string]*string) *ListFunctionInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionInstancesResponse) SetStatusCode(v int32) *ListFunctionInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionInstancesResponse) SetBody(v *ListFunctionInstancesResponseBody) *ListFunctionInstancesResponse {
	s.Body = v
	return s
}

type ListFunctionTasksRequest struct {
	// The end time is less than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of the page to return. Default value: 1.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 1.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The start time is greater than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the task. Valid values:
	//
	// *   success
	// *   failed
	// *   running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListFunctionTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksRequest) SetEndTime(v int64) *ListFunctionTasksRequest {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageNumber(v int32) *ListFunctionTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionTasksRequest) SetPageSize(v int32) *ListFunctionTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStartTime(v int64) *ListFunctionTasksRequest {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksRequest) SetStatus(v string) *ListFunctionTasksRequest {
	s.Status = &v
	return s
}

type ListFunctionTasksResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	Result []*ListFunctionTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of records that meet the requirements.
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBody) SetCode(v string) *ListFunctionTasksResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetHttpCode(v int64) *ListFunctionTasksResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetLatency(v int64) *ListFunctionTasksResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetMessage(v string) *ListFunctionTasksResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetRequestId(v string) *ListFunctionTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetResult(v []*ListFunctionTasksResponseBodyResult) *ListFunctionTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionTasksResponseBody) SetStatus(v string) *ListFunctionTasksResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionTasksResponseBody) SetTotalCount(v int64) *ListFunctionTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListFunctionTasksResponseBodyResult struct {
	// The timestamp that indicates the end time. Unit: milliseconds. 0 indicates that the task has not ended.
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The value is a JSON string. It includes model evaluation information and training error information.
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The number of iterations.
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The progress. 90 indicates 90%.
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the task.
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The timestamp that indicates the start time. Unit: milliseconds.
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// *   success
	// *   failed
	// *   running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFunctionTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponseBodyResult) SetEndTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.EndTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetExtendInfo(v string) *ListFunctionTasksResponseBodyResult {
	s.ExtendInfo = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetFunctionName(v string) *ListFunctionTasksResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetGeneration(v string) *ListFunctionTasksResponseBodyResult {
	s.Generation = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetProgress(v int64) *ListFunctionTasksResponseBodyResult {
	s.Progress = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetRunId(v string) *ListFunctionTasksResponseBodyResult {
	s.RunId = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStartTime(v int64) *ListFunctionTasksResponseBodyResult {
	s.StartTime = &v
	return s
}

func (s *ListFunctionTasksResponseBodyResult) SetStatus(v string) *ListFunctionTasksResponseBodyResult {
	s.Status = &v
	return s
}

type ListFunctionTasksResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFunctionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFunctionTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionTasksResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionTasksResponse) SetHeaders(v map[string]*string) *ListFunctionTasksResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionTasksResponse) SetStatusCode(v int32) *ListFunctionTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionTasksResponse) SetBody(v *ListFunctionTasksResponseBody) *ListFunctionTasksResponse {
	s.Body = v
	return s
}

type ListInterventionDictionariesRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// \-
	Types *string `json:"types,omitempty" xml:"types,omitempty"`
}

func (s ListInterventionDictionariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesRequest) SetPageNumber(v int32) *ListInterventionDictionariesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetPageSize(v int32) *ListInterventionDictionariesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionariesRequest) SetTypes(v string) *ListInterventionDictionariesRequest {
	s.Types = &v
	return s
}

type ListInterventionDictionariesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each intervention dictionary.
	//
	// For more information, see [InterventionDictionary](~~173608~~).
	Result []*ListInterventionDictionariesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBody) SetRequestId(v string) *ListInterventionDictionariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetResult(v []*ListInterventionDictionariesResponseBodyResult) *ListInterventionDictionariesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionariesResponseBody) SetTotalCount(v int32) *ListInterventionDictionariesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInterventionDictionariesResponseBodyResult struct {
	// The custom analyzer.
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the intervention dictionary.
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the intervention dictionary.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// *   stopword: an intervention dictionary for stop word filtering
	// *   synonym: an intervention dictionary for synonym configuration
	// *   correction: an intervention dictionary for spelling correction
	// *   category_prediction: an intervention dictionary for category prediction
	// *   ner: an intervention dictionary for named entity recognition (NER)
	// *   term_weighting: an intervention dictionary for term weight analysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListInterventionDictionariesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponseBodyResult) SetAnalyzer(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetCreated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetId(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetName(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetType(v string) *ListInterventionDictionariesResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListInterventionDictionariesResponseBodyResult) SetUpdated(v int32) *ListInterventionDictionariesResponseBodyResult {
	s.Updated = &v
	return s
}

type ListInterventionDictionariesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInterventionDictionariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionariesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionariesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionariesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionariesResponse) SetStatusCode(v int32) *ListInterventionDictionariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionariesResponse) SetBody(v *ListInterventionDictionariesResponseBody) *ListInterventionDictionariesResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryEntriesRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Test
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageNumber(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetPageSize(v int32) *ListInterventionDictionaryEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListInterventionDictionaryEntriesRequest) SetWord(v string) *ListInterventionDictionaryEntriesRequest {
	s.Word = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each intervention entry.
	//
	// For more information, see [InterventionDictionaryEntry](~~173606~~).
	Result []*ListInterventionDictionaryEntriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *ListInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetResult(v []*ListInterventionDictionaryEntriesResponseBodyResult) *ListInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBody) SetTotalCount(v int32) *ListInterventionDictionaryEntriesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBodyResult struct {
	// The action. Valid values:
	//
	// *   add
	// *   delete
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// The timestamp when the intervention entry was created.
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The content of an intervention entry for category prediction.
	//
	// The parameter returns key-value pairs. The key in a key-value pair indicates the ID of the category. The value in a key-value pair indicates the relevance value of the category. A value of 0 indicates irrelevant. A value of 1 indicates slightly relevant. A value of 2 indicates relevant.
	//
	// Example: {"2":1, "100":0}
	Relevance map[string]interface{} `json:"relevance,omitempty" xml:"relevance,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// *   ACTIVE: The intervention entry takes effect.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The content of an intervention entry for term weight analysis.
	Tokens []*ListInterventionDictionaryEntriesResponseBodyResultTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
	// The timestamp when the intervention entry was last updated.
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The intervention query in the intervention entry.
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCmd(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Cmd = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetCreated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetRelevance(v map[string]interface{}) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Relevance = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetStatus(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetTokens(v []*ListInterventionDictionaryEntriesResponseBodyResultTokens) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Tokens = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetUpdated(v int64) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResult) SetWord(v string) *ListInterventionDictionaryEntriesResponseBodyResult {
	s.Word = &v
	return s
}

type ListInterventionDictionaryEntriesResponseBodyResultTokens struct {
	// The sequence number.
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The internal name of the identified entity type. Valid values:
	//
	// *   brand
	// *   category
	// *   material
	// *   element
	// *   style
	// *   color
	// *   function
	// *   scenario
	// *   people
	// *   season
	// *   model
	// *   region
	// *   name
	// *   adjective
	// *   category-modifier
	// *   size
	// *   quality
	// *   suit
	// *   new-release
	// *   series
	// *   marketing
	// *   entertainment
	// *   organization
	// *   movie
	// *   game
	// *   number
	// *   unit
	// *   common
	// *   new-word
	// *   proper-noun
	// *   symbol
	// *   prefix
	// *   suffix
	// *   gift
	// *   negative
	// *   agent
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The description of the internal name of the identified entity type.
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The entity.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponseBodyResultTokens) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetOrder(v int32) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTag(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetTagLabel(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponseBodyResultTokens) SetToken(v string) *ListInterventionDictionaryEntriesResponseBodyResultTokens {
	s.Token = &v
	return s
}

type ListInterventionDictionaryEntriesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) SetStatusCode(v int32) *ListInterventionDictionaryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryEntriesResponse) SetBody(v *ListInterventionDictionaryEntriesResponseBody) *ListInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryNerResultsRequest struct {
	// Soymilk
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s ListInterventionDictionaryNerResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsRequest) SetQuery(v string) *ListInterventionDictionaryNerResultsRequest {
	s.Query = &v
	return s
}

type ListInterventionDictionaryNerResultsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The NER result.
	//
	// For more information, see [InterventionDictionaryEntry](~~173606~~).
	Result []*ListInterventionDictionaryNerResultsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryNerResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetRequestId(v string) *ListInterventionDictionaryNerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBody) SetResult(v []*ListInterventionDictionaryNerResultsResponseBodyResult) *ListInterventionDictionaryNerResultsResponseBody {
	s.Result = v
	return s
}

type ListInterventionDictionaryNerResultsResponseBodyResult struct {
	// The sequence number.
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The internal name of the identified entity type. Valid values:
	//
	// *   brand
	// *   category
	// *   material
	// *   element
	// *   style
	// *   color
	// *   function
	// *   scenario
	// *   people
	// *   season
	// *   model
	// *   region
	// *   name
	// *   adjective
	// *   category-modifier
	// *   size
	// *   quality
	// *   suit
	// *   new-release
	// *   series
	// *   marketing
	// *   entertainment
	// *   organization
	// *   movie
	// *   game
	// *   number
	// *   unit
	// *   common
	// *   new-word
	// *   proper-noun
	// *   symbol
	// *   prefix
	// *   suffix
	// *   gift
	// *   negative
	// *   agent
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The description of the internal name of the identified entity type.
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The entity.
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetOrder(v int32) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTag(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Tag = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetTagLabel(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.TagLabel = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponseBodyResult) SetToken(v string) *ListInterventionDictionaryNerResultsResponseBodyResult {
	s.Token = &v
	return s
}

type ListInterventionDictionaryNerResultsResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInterventionDictionaryNerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryNerResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryNerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryNerResultsResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryNerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) SetStatusCode(v int32) *ListInterventionDictionaryNerResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryNerResultsResponse) SetBody(v *ListInterventionDictionaryNerResultsResponseBody) *ListInterventionDictionaryNerResultsResponse {
	s.Body = v
	return s
}

type ListInterventionDictionaryRelatedEntitiesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each application and each query analysis rule. If no query analysis rule references the intervention dictionary, the value of the result parameter is an empty list.
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetRequestId(v string) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponseBody) SetResult(v []map[string]interface{}) *ListInterventionDictionaryRelatedEntitiesResponseBody {
	s.Result = v
	return s
}

type ListInterventionDictionaryRelatedEntitiesResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInterventionDictionaryRelatedEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInterventionDictionaryRelatedEntitiesResponse) GoString() string {
	return s.String()
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetHeaders(v map[string]*string) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Headers = v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetStatusCode(v int32) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterventionDictionaryRelatedEntitiesResponse) SetBody(v *ListInterventionDictionaryRelatedEntitiesResponseBody) *ListInterventionDictionaryRelatedEntitiesResponse {
	s.Body = v
	return s
}

type ListModelsRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// pop
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListModelsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListModelsRequest) GoString() string {
	return s.String()
}

func (s *ListModelsRequest) SetPageNumber(v int32) *ListModelsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelsRequest) SetPageSize(v int32) *ListModelsRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelsRequest) SetType(v string) *ListModelsRequest {
	s.Type = &v
	return s
}

type ListModelsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the models.
	//
	// For more information, see [Model](~~180898~~).
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListModelsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetResult(v []map[string]interface{}) *ListModelsResponseBody {
	s.Result = v
	return s
}

type ListModelsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListModelsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListModelsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListModelsResponse) GoString() string {
	return s.String()
}

func (s *ListModelsResponse) SetHeaders(v map[string]*string) *ListModelsResponse {
	s.Headers = v
	return s
}

func (s *ListModelsResponse) SetStatusCode(v int32) *ListModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListModelsResponse) SetBody(v *ListModelsResponseBody) *ListModelsResponse {
	s.Body = v
	return s
}

type ListProceedingsRequest struct {
	// Specifies whether the filtering is complete.
	FilterFinished *bool `json:"filterFinished,omitempty" xml:"filterFinished,omitempty"`
}

func (s ListProceedingsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListProceedingsRequest) GoString() string {
	return s.String()
}

func (s *ListProceedingsRequest) SetFilterFinished(v bool) *ListProceedingsRequest {
	s.FilterFinished = &v
	return s
}

type ListProceedingsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProceedingsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListProceedingsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponseBody) SetRequestId(v string) *ListProceedingsResponseBody {
	s.RequestId = &v
	return s
}

type ListProceedingsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListProceedingsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListProceedingsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListProceedingsResponse) GoString() string {
	return s.String()
}

func (s *ListProceedingsResponse) SetHeaders(v map[string]*string) *ListProceedingsResponse {
	s.Headers = v
	return s
}

func (s *ListProceedingsResponse) SetStatusCode(v int32) *ListProceedingsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProceedingsResponse) SetBody(v *ListProceedingsResponseBody) *ListProceedingsResponse {
	s.Body = v
	return s
}

type ListQueryProcessorAnalyzerResultsRequest struct {
	// The text to be tested.
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsRequest) SetText(v string) *ListQueryProcessorAnalyzerResultsRequest {
	s.Text = &v
	return s
}

type ListQueryProcessorAnalyzerResultsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetRequestId(v string) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponseBody) SetResult(v map[string]interface{}) *ListQueryProcessorAnalyzerResultsResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorAnalyzerResultsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueryProcessorAnalyzerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorAnalyzerResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorAnalyzerResultsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorAnalyzerResultsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetStatusCode(v int32) *ListQueryProcessorAnalyzerResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorAnalyzerResultsResponse) SetBody(v *ListQueryProcessorAnalyzerResultsResponseBody) *ListQueryProcessorAnalyzerResultsResponse {
	s.Body = v
	return s
}

type ListQueryProcessorNersRequest struct {
	// ECOMMERCE
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
}

func (s ListQueryProcessorNersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersRequest) SetDomain(v string) *ListQueryProcessorNersRequest {
	s.Domain = &v
	return s
}

type ListQueryProcessorNersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The priority settings of entity types.
	//
	// For more information, see [Priority settings of entity types](~~173616~~).
	Result []*ListQueryProcessorNersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorNersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBody) SetRequestId(v string) *ListQueryProcessorNersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorNersResponseBody) SetResult(v []*ListQueryProcessorNersResponseBodyResult) *ListQueryProcessorNersResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorNersResponseBodyResult struct {
	// The name of the entity type.
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The priority of an entity type among entity types that have the same priority level.
	//
	// A smaller value indicates a higher priority. Default value: 0.
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The priority level of the entity type.
	//
	// *   HIGH
	// *   MIDDLE
	// *   LOW
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The internal name of the entity type.
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListQueryProcessorNersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponseBodyResult) SetLabel(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Label = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetOrder(v int32) *ListQueryProcessorNersResponseBodyResult {
	s.Order = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetPriority(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Priority = &v
	return s
}

func (s *ListQueryProcessorNersResponseBodyResult) SetTag(v string) *ListQueryProcessorNersResponseBodyResult {
	s.Tag = &v
	return s
}

type ListQueryProcessorNersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueryProcessorNersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorNersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorNersResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponse) SetHeaders(v map[string]*string) *ListQueryProcessorNersResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorNersResponse) SetStatusCode(v int32) *ListQueryProcessorNersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorNersResponse) SetBody(v *ListQueryProcessorNersResponseBody) *ListQueryProcessorNersResponse {
	s.Body = v
	return s
}

type ListQueryProcessorsRequest struct {
	// 0
	IsActive *int32 `json:"isActive,omitempty" xml:"isActive,omitempty"`
}

func (s ListQueryProcessorsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsRequest) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsRequest) SetIsActive(v int32) *ListQueryProcessorsRequest {
	s.IsActive = &v
	return s
}

type ListQueryProcessorsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each query analysis rule.
	//
	// For more information, see [QueryProcessor](~~170014~~).
	Result []*ListQueryProcessorsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListQueryProcessorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBody) SetRequestId(v string) *ListQueryProcessorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQueryProcessorsResponseBody) SetResult(v []*ListQueryProcessorsResponseBodyResult) *ListQueryProcessorsResponseBody {
	s.Result = v
	return s
}

type ListQueryProcessorsResponseBodyResult struct {
	// Indicates whether the query analysis rule is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	//
	// For more information, see the ["Processor"](~~170014~~) section of the QueryProcessor topic.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListQueryProcessorsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponseBodyResult) SetActive(v bool) *ListQueryProcessorsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetCreated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetDomain(v string) *ListQueryProcessorsResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetIndexes(v []*string) *ListQueryProcessorsResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetName(v string) *ListQueryProcessorsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetProcessors(v []map[string]interface{}) *ListQueryProcessorsResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ListQueryProcessorsResponseBodyResult) SetUpdated(v int32) *ListQueryProcessorsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListQueryProcessorsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQueryProcessorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQueryProcessorsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQueryProcessorsResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorsResponse) SetHeaders(v map[string]*string) *ListQueryProcessorsResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorsResponse) SetStatusCode(v int32) *ListQueryProcessorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorsResponse) SetBody(v *ListQueryProcessorsResponseBody) *ListQueryProcessorsResponse {
	s.Body = v
	return s
}

type ListQuotaReviewTasksRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListQuotaReviewTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksRequest) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksRequest) SetPageNumber(v int32) *ListQuotaReviewTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotaReviewTasksRequest) SetPageSize(v int32) *ListQuotaReviewTasksRequest {
	s.PageSize = &v
	return s
}

type ListQuotaReviewTasksResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the ticket for application quota approval.
	//
	// For more information, see [QuotaReviewTask](~~173609~~).
	Result []*ListQuotaReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the returned tickets.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListQuotaReviewTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBody) SetRequestId(v string) *ListQuotaReviewTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetResult(v []*ListQuotaReviewTasksResponseBodyResult) *ListQuotaReviewTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListQuotaReviewTasksResponseBody) SetTotalCount(v int32) *ListQuotaReviewTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListQuotaReviewTasksResponseBodyResult struct {
	// The ID of the application.
	AppGroupId *int32 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The name of the application.
	AppGroupName *string `json:"appGroupName,omitempty" xml:"appGroupName,omitempty"`
	// The type of the application.
	AppGroupType *string `json:"appGroupType,omitempty" xml:"appGroupType,omitempty"`
	// Indicates whether the ticket is approved.
	Approved *bool `json:"approved,omitempty" xml:"approved,omitempty"`
	// Indicates whether the model is available.
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The time when the ticket was created.
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the ticket was last updated.
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ID of the ticket.
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The remarks of the ticket.
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// The computing resource quota that is applied for.
	NewComputeResource *int32 `json:"newComputeResource,omitempty" xml:"newComputeResource,omitempty"`
	// The storage capacity quota that is applied for.
	NewSocSize *int32 `json:"newSocSize,omitempty" xml:"newSocSize,omitempty"`
	// The application specifications that are applied for.
	NewSpec *string `json:"newSpec,omitempty" xml:"newSpec,omitempty"`
	// The original quota of computing resources.
	OldComputeResource *int32 `json:"oldComputeResource,omitempty" xml:"oldComputeResource,omitempty"`
	// The original quota of storage capacity.
	OldDocSize *int32 `json:"oldDocSize,omitempty" xml:"oldDocSize,omitempty"`
	// The original application specifications.
	OldSpec *string `json:"oldSpec,omitempty" xml:"oldSpec,omitempty"`
	// Indicates whether the ticket is pending.
	Pending *bool `json:"pending,omitempty" xml:"pending,omitempty"`
}

func (s ListQuotaReviewTasksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupName(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupName = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAppGroupType(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.AppGroupType = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetApproved(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Approved = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetAvailable(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtCreate(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetGmtModified(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetId(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetMemo(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.Memo = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetNewSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.NewSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldComputeResource(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldComputeResource = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldDocSize(v int32) *ListQuotaReviewTasksResponseBodyResult {
	s.OldDocSize = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetOldSpec(v string) *ListQuotaReviewTasksResponseBodyResult {
	s.OldSpec = &v
	return s
}

func (s *ListQuotaReviewTasksResponseBodyResult) SetPending(v bool) *ListQuotaReviewTasksResponseBodyResult {
	s.Pending = &v
	return s
}

type ListQuotaReviewTasksResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListQuotaReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuotaReviewTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotaReviewTasksResponse) GoString() string {
	return s.String()
}

func (s *ListQuotaReviewTasksResponse) SetHeaders(v map[string]*string) *ListQuotaReviewTasksResponse {
	s.Headers = v
	return s
}

func (s *ListQuotaReviewTasksResponse) SetStatusCode(v int32) *ListQuotaReviewTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQuotaReviewTasksResponse) SetBody(v *ListQuotaReviewTasksResponseBody) *ListQuotaReviewTasksResponse {
	s.Body = v
	return s
}

type ListScheduledTasksRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// wipe
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListScheduledTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksRequest) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksRequest) SetPageNumber(v int32) *ListScheduledTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListScheduledTasksRequest) SetPageSize(v int32) *ListScheduledTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListScheduledTasksRequest) SetType(v string) *ListScheduledTasksRequest {
	s.Type = &v
	return s
}

type ListScheduledTasksResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled tasks.
	//
	// For more information, see [ScheduledTask](~~173610~~).
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the returned scheduled tasks.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListScheduledTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponseBody) SetRequestId(v string) *ListScheduledTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledTasksResponseBody) SetResult(v []map[string]interface{}) *ListScheduledTasksResponseBody {
	s.Result = v
	return s
}

func (s *ListScheduledTasksResponseBody) SetTotalCount(v int64) *ListScheduledTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListScheduledTasksResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListScheduledTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListScheduledTasksResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledTasksResponse) SetHeaders(v map[string]*string) *ListScheduledTasksResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledTasksResponse) SetStatusCode(v int32) *ListScheduledTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledTasksResponse) SetBody(v *ListScheduledTasksResponseBody) *ListScheduledTasksResponse {
	s.Body = v
	return s
}

type ListSearchStrategiesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSearchStrategiesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSearchStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponseBody) SetRequestId(v string) *ListSearchStrategiesResponseBody {
	s.RequestId = &v
	return s
}

type ListSearchStrategiesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSearchStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSearchStrategiesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSearchStrategiesResponse) GoString() string {
	return s.String()
}

func (s *ListSearchStrategiesResponse) SetHeaders(v map[string]*string) *ListSearchStrategiesResponse {
	s.Headers = v
	return s
}

func (s *ListSearchStrategiesResponse) SetStatusCode(v int32) *ListSearchStrategiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSearchStrategiesResponse) SetBody(v *ListSearchStrategiesResponseBody) *ListSearchStrategiesResponse {
	s.Body = v
	return s
}

type ListSecondRanksResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each fine sort expression.
	//
	// For more information, see [SecondRank](~~170008~~).
	Result []*ListSecondRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSecondRanksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBody) SetRequestId(v string) *ListSecondRanksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSecondRanksResponseBody) SetResult(v []*ListSecondRanksResponseBodyResult) *ListSecondRanksResponseBody {
	s.Result = v
	return s
}

func (s *ListSecondRanksResponseBody) SetTotalCount(v int32) *ListSecondRanksResponseBody {
	s.TotalCount = &v
	return s
}

type ListSecondRanksResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression.
	//
	// You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSecondRanksResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponseBodyResult) SetActive(v bool) *ListSecondRanksResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetCreated(v int32) *ListSecondRanksResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetDescription(v string) *ListSecondRanksResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetId(v string) *ListSecondRanksResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsDefault(v string) *ListSecondRanksResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetIsSys(v string) *ListSecondRanksResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetMeta(v string) *ListSecondRanksResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetName(v string) *ListSecondRanksResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSecondRanksResponseBodyResult) SetUpdated(v int32) *ListSecondRanksResponseBodyResult {
	s.Updated = &v
	return s
}

type ListSecondRanksResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSecondRanksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecondRanksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecondRanksResponse) GoString() string {
	return s.String()
}

func (s *ListSecondRanksResponse) SetHeaders(v map[string]*string) *ListSecondRanksResponse {
	s.Headers = v
	return s
}

func (s *ListSecondRanksResponse) SetStatusCode(v int32) *ListSecondRanksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecondRanksResponse) SetBody(v *ListSecondRanksResponseBody) *ListSecondRanksResponse {
	s.Body = v
	return s
}

type ListSlowQueryCategoriesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data returned.
	Result *ListSlowQueryCategoriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryCategoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBody) SetRequestId(v string) *ListSlowQueryCategoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBody) SetResult(v *ListSlowQueryCategoriesResponseBodyResult) *ListSlowQueryCategoriesResponseBody {
	s.Result = v
	return s
}

type ListSlowQueryCategoriesResponseBodyResult struct {
	// The status of the analysis. Valid values:
	//
	// *   PENDING: preparing
	// *   SUCCESS: succeeded
	// *   RUNNING: running
	// *   FAILED: failed
	// *   N/A: unknown
	AnalyzeStatus *string `json:"analyzeStatus,omitempty" xml:"analyzeStatus,omitempty"`
	// The timestamp that indicates the end of the time range to query.
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The timestamp that indicates the beginning of the time range to query.
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryCategoriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetAnalyzeStatus(v string) *ListSlowQueryCategoriesResponseBodyResult {
	s.AnalyzeStatus = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryCategoriesResponseBodyResult) SetStart(v int32) *ListSlowQueryCategoriesResponseBodyResult {
	s.Start = &v
	return s
}

type ListSlowQueryCategoriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSlowQueryCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlowQueryCategoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryCategoriesResponse) GoString() string {
	return s.String()
}

func (s *ListSlowQueryCategoriesResponse) SetHeaders(v map[string]*string) *ListSlowQueryCategoriesResponse {
	s.Headers = v
	return s
}

func (s *ListSlowQueryCategoriesResponse) SetStatusCode(v int32) *ListSlowQueryCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlowQueryCategoriesResponse) SetBody(v *ListSlowQueryCategoriesResponseBody) *ListSlowQueryCategoriesResponse {
	s.Body = v
	return s
}

type ListSlowQueryQueriesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *ListSlowQueryQueriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ListSlowQueryQueriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBody) SetRequestId(v string) *ListSlowQueryQueriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBody) SetResult(v *ListSlowQueryQueriesResponseBodyResult) *ListSlowQueryQueriesResponseBody {
	s.Result = v
	return s
}

type ListSlowQueryQueriesResponseBodyResult struct {
	// The content of the optimization suggestion for the query.
	AppQuery *string `json:"appQuery,omitempty" xml:"appQuery,omitempty"`
	// The end of the time range that was queried.
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The ID of the optimization suggestion.
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The beginning of the time range that was queried.
	Start *int32 `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ListSlowQueryQueriesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetAppQuery(v string) *ListSlowQueryQueriesResponseBodyResult {
	s.AppQuery = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetEnd(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.End = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetIndex(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Index = &v
	return s
}

func (s *ListSlowQueryQueriesResponseBodyResult) SetStart(v int32) *ListSlowQueryQueriesResponseBodyResult {
	s.Start = &v
	return s
}

type ListSlowQueryQueriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSlowQueryQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSlowQueryQueriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSlowQueryQueriesResponse) GoString() string {
	return s.String()
}

func (s *ListSlowQueryQueriesResponse) SetHeaders(v map[string]*string) *ListSlowQueryQueriesResponse {
	s.Headers = v
	return s
}

func (s *ListSlowQueryQueriesResponse) SetStatusCode(v int32) *ListSlowQueryQueriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSlowQueryQueriesResponse) SetBody(v *ListSlowQueryQueriesResponseBody) *ListSlowQueryQueriesResponse {
	s.Body = v
	return s
}

type ListSortExpressionsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort or fine sort expressions that were returned.
	//
	// For more information, see [FirstRank](~~170007~~) and [SecondRank](~~170008~~).
	Result []*ListSortExpressionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortExpressionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBody) SetRequestId(v string) *ListSortExpressionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortExpressionsResponseBody) SetResult(v []*ListSortExpressionsResponseBodyResult) *ListSortExpressionsResponseBody {
	s.Result = v
	return s
}

type ListSortExpressionsResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListSortExpressionsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponseBodyResult) SetActive(v bool) *ListSortExpressionsResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetCreated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetDescription(v string) *ListSortExpressionsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetName(v string) *ListSortExpressionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListSortExpressionsResponseBodyResult) SetUpdated(v int32) *ListSortExpressionsResponseBodyResult {
	s.Updated = &v
	return s
}

type ListSortExpressionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSortExpressionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSortExpressionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSortExpressionsResponse) GoString() string {
	return s.String()
}

func (s *ListSortExpressionsResponse) SetHeaders(v map[string]*string) *ListSortExpressionsResponse {
	s.Headers = v
	return s
}

func (s *ListSortExpressionsResponse) SetStatusCode(v int32) *ListSortExpressionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSortExpressionsResponse) SetBody(v *ListSortExpressionsResponseBody) *ListSortExpressionsResponse {
	s.Body = v
	return s
}

type ListSortScriptsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scripts of the application version.
	Result []*ListSortScriptsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListSortScriptsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBody) SetRequestId(v string) *ListSortScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSortScriptsResponseBody) SetResult(v []*ListSortScriptsResponseBodyResult) *ListSortScriptsResponseBody {
	s.Result = v
	return s
}

type ListSortScriptsResponseBodyResult struct {
	// The time when the script was created.
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The sort phase to which the script applies.
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The name of the script.
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The status of the script. Valid values:
	//
	// *   configurable: The script is created, but no script files are uploaded.
	// *   not compiled: The script is not compiled.
	// *   compile failed: The compilation of the script failed.
	// *   compile successful: The script is compiled.
	// *   released: The script is published.
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the script.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListSortScriptsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponseBodyResult) SetCreateTime(v string) *ListSortScriptsResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetModifyTime(v string) *ListSortScriptsResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScope(v string) *ListSortScriptsResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetScriptName(v string) *ListSortScriptsResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetStatus(v string) *ListSortScriptsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListSortScriptsResponseBodyResult) SetType(v string) *ListSortScriptsResponseBodyResult {
	s.Type = &v
	return s
}

type ListSortScriptsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListSortScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSortScriptsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSortScriptsResponse) GoString() string {
	return s.String()
}

func (s *ListSortScriptsResponse) SetHeaders(v map[string]*string) *ListSortScriptsResponse {
	s.Headers = v
	return s
}

func (s *ListSortScriptsResponse) SetStatusCode(v int32) *ListSortScriptsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSortScriptsResponse) SetBody(v *ListSortScriptsResponseBody) *ListSortScriptsResponse {
	s.Body = v
	return s
}

type ListStatisticLogsRequest struct {
	// The fields to query. Example: columns=wordsTopPv.
	//
	// For more information, see [Metrics in statistical reports](https://www.alibabacloud.com/help/en/opensearch/latest/statistical-report).
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// The content of the query clause.
	Distinct *bool `json:"distinct,omitempty" xml:"distinct,omitempty"`
	// The number of the page to return. Default value: 1.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The content of the query clause.
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The content of the sort clause.
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The beginning of the time range to query. The default value is the timestamp of 00:00:00 on the current day.
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The end of the time range to query. The default value is the timestamp of 24:00:00 on the current day.
	StopTime *int32 `json:"stopTime,omitempty" xml:"stopTime,omitempty"`
}

func (s ListStatisticLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsRequest) SetColumns(v string) *ListStatisticLogsRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticLogsRequest) SetDistinct(v bool) *ListStatisticLogsRequest {
	s.Distinct = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageNumber(v int32) *ListStatisticLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticLogsRequest) SetPageSize(v int32) *ListStatisticLogsRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticLogsRequest) SetQuery(v string) *ListStatisticLogsRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticLogsRequest) SetSortBy(v string) *ListStatisticLogsRequest {
	s.SortBy = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStartTime(v int32) *ListStatisticLogsRequest {
	s.StartTime = &v
	return s
}

func (s *ListStatisticLogsRequest) SetStopTime(v int32) *ListStatisticLogsRequest {
	s.StopTime = &v
	return s
}

type ListStatisticLogsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result. For more information, see:
	//
	// *   [Parameters of hotwords rankings](~~421248~~)
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the queried logs.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponseBody) SetRequestId(v string) *ListStatisticLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticLogsResponseBody) SetResult(v []map[string]interface{}) *ListStatisticLogsResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticLogsResponseBody) SetTotalCount(v int64) *ListStatisticLogsResponseBody {
	s.TotalCount = &v
	return s
}

type ListStatisticLogsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStatisticLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatisticLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticLogsResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticLogsResponse) SetHeaders(v map[string]*string) *ListStatisticLogsResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticLogsResponse) SetStatusCode(v int32) *ListStatisticLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStatisticLogsResponse) SetBody(v *ListStatisticLogsResponseBody) *ListStatisticLogsResponse {
	s.Body = v
	return s
}

type ListStatisticReportRequest struct {
	// pv,uv
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// 1582646399
	EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// bizType:test,sceneTag:myTag
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 1582214400
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListStatisticReportRequest) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportRequest) GoString() string {
	return s.String()
}

func (s *ListStatisticReportRequest) SetColumns(v string) *ListStatisticReportRequest {
	s.Columns = &v
	return s
}

func (s *ListStatisticReportRequest) SetEndTime(v int32) *ListStatisticReportRequest {
	s.EndTime = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageNumber(v int32) *ListStatisticReportRequest {
	s.PageNumber = &v
	return s
}

func (s *ListStatisticReportRequest) SetPageSize(v int32) *ListStatisticReportRequest {
	s.PageSize = &v
	return s
}

func (s *ListStatisticReportRequest) SetQuery(v string) *ListStatisticReportRequest {
	s.Query = &v
	return s
}

func (s *ListStatisticReportRequest) SetStartTime(v int32) *ListStatisticReportRequest {
	s.StartTime = &v
	return s
}

type ListStatisticReportResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The queried reports. Valid values:
	//
	// For more information about the metrics in data quality reports, see the Upload behavioral data section of [Data collection 2.0](~~131547~~).
	//
	// For more information about the metrics in application and A/B test reports, see the Core metrics section of [Metrics of statistical reports](~~187654~~).
	//
	// For more information about the metrics in query analysis reports, see the Query analysis metrics section of [Metrics of statistical reports](~~187654~~).
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the queried reports.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListStatisticReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportResponseBody) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponseBody) SetRequestId(v string) *ListStatisticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStatisticReportResponseBody) SetResult(v []map[string]interface{}) *ListStatisticReportResponseBody {
	s.Result = v
	return s
}

func (s *ListStatisticReportResponseBody) SetTotalCount(v int64) *ListStatisticReportResponseBody {
	s.TotalCount = &v
	return s
}

type ListStatisticReportResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStatisticReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStatisticReportResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStatisticReportResponse) GoString() string {
	return s.String()
}

func (s *ListStatisticReportResponse) SetHeaders(v map[string]*string) *ListStatisticReportResponse {
	s.Headers = v
	return s
}

func (s *ListStatisticReportResponse) SetStatusCode(v int32) *ListStatisticReportResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStatisticReportResponse) SetBody(v *ListStatisticReportResponseBody) *ListStatisticReportResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// The token that is used to retrieve the next page.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags. You can specify a maximum number of 20 tags.
	Tag []*ListTagResourcesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// The key of the tag.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesShrinkRequest struct {
	// The token that is used to retrieve the next page.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags. You can specify a maximum number of 20 tags.
	TagShrink *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s ListTagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesShrinkRequest) SetNextToken(v string) *ListTagResourcesShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceIdShrink(v string) *ListTagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetResourceType(v string) *ListTagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesShrinkRequest) SetTagShrink(v string) *ListTagResourcesShrinkRequest {
	s.TagShrink = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// The token that is used to retrieve the next page.
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resources.
	Result []*ListTagResourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetResult(v []*ListTagResourcesResponseBodyResult) *ListTagResourcesResponseBody {
	s.Result = v
	return s
}

type ListTagResourcesResponseBodyResult struct {
	// The ID of the resource.
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The key of the tag.
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The value of the tag.
	TagValue *string `json:"tagValue,omitempty" xml:"tagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyResult) SetResourceId(v string) *ListTagResourcesResponseBodyResult {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetResourceType(v string) *ListTagResourcesResponseBodyResult {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetTagKey(v string) *ListTagResourcesResponseBodyResult {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyResult) SetTagValue(v string) *ListTagResourcesResponseBodyResult {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUserAnalyzerEntriesRequest struct {
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// kevintest
	Word *string `json:"word,omitempty" xml:"word,omitempty"`
}

func (s ListUserAnalyzerEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesRequest) SetPageNumber(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetPageSize(v int32) *ListUserAnalyzerEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *ListUserAnalyzerEntriesRequest) SetWord(v string) *ListUserAnalyzerEntriesRequest {
	s.Word = &v
	return s
}

type ListUserAnalyzerEntriesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entries of the custom analyzer.
	//
	// For more information, see [UserAnalyzerEntry](~~178932~~).
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ListUserAnalyzerEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponseBody) SetRequestId(v string) *ListUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *ListUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

type ListUserAnalyzerEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserAnalyzerEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *ListUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) SetStatusCode(v int32) *ListUserAnalyzerEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) SetBody(v *ListUserAnalyzerEntriesResponseBody) *ListUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

type ListUserAnalyzersRequest struct {
	// The number of the page to return. Default value: 1.
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListUserAnalyzersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersRequest) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersRequest) SetPageNumber(v int32) *ListUserAnalyzersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListUserAnalyzersRequest) SetPageSize(v int32) *ListUserAnalyzersRequest {
	s.PageSize = &v
	return s
}

type ListUserAnalyzersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// For more information, see [UserAnalyzer](~~178934~~).
	Result []*ListUserAnalyzersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number.
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListUserAnalyzersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBody) SetRequestId(v string) *ListUserAnalyzersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetResult(v []*ListUserAnalyzersResponseBodyResult) *ListUserAnalyzersResponseBody {
	s.Result = v
	return s
}

func (s *ListUserAnalyzersResponseBody) SetTotalCount(v int32) *ListUserAnalyzersResponseBody {
	s.TotalCount = &v
	return s
}

type ListUserAnalyzersResponseBodyResult struct {
	// Indicates whether the application is available.
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The basic analyzer. Valid values:
	//
	// *   chn_standard: [a common analyzer in Chinese](~~179424~~)
	// *   chn_scene_name: an analyzer for person names in Chinese
	// *   chn_ecommerce: [an analyzer for E-commerce in Chinese](~~179424~~)
	// *   chn_it_content: [an analyzer for IT content in Chinese](~~179424~~)
	// *   en_min: a small-granularity analyzer in English
	// *   th_standard: a common analyzer in Thai
	// *   th_ecommerce: an analyzer for E-commerce in Thai
	// *   vn_standard: a common analyzer in Vietnamese
	// *   chn_community_it: an analyzer for IT community content in Chinese
	// *   chn_ecommerce_general: a common analyzer for the E-commerce industry in Chinese
	// *   chn_esports_general: a common analyzer for the gaming industry in Chinese
	// *   chn_edu_question: an analyzer for question search of the education industry in Chinese
	Business *string `json:"business,omitempty" xml:"business,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The dictionaries that are used by the custom analyzer.
	//
	// For more information, see [UserDict](~~178933~~).
	Dicts []*ListUserAnalyzersResponseBodyResultDicts `json:"dicts,omitempty" xml:"dicts,omitempty" type:"Repeated"`
	// The ID of the custom analyzer.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the custom analyzer.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResult) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResult {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetBusiness(v string) *ListUserAnalyzersResponseBodyResult {
	s.Business = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetCreated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetDicts(v []*ListUserAnalyzersResponseBodyResultDicts) *ListUserAnalyzersResponseBodyResult {
	s.Dicts = v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetId(v string) *ListUserAnalyzersResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetName(v string) *ListUserAnalyzersResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResult) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResult {
	s.Updated = &v
	return s
}

type ListUserAnalyzersResponseBodyResultDicts struct {
	// Indicates whether the application is available.
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The number of intervention entries.
	EntriesCount *int32 `json:"entriesCount,omitempty" xml:"entriesCount,omitempty"`
	// The maximum number of intervention entries that can be created in the dictionary.
	EntriesLimit *int32 `json:"entriesLimit,omitempty" xml:"entriesLimit,omitempty"`
	// The ID of the dictionary.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type. Valid value:
	//
	// *   segment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListUserAnalyzersResponseBodyResultDicts) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponseBodyResultDicts) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetAvailable(v bool) *ListUserAnalyzersResponseBodyResultDicts {
	s.Available = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetCreated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Created = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesCount(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesCount = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetEntriesLimit(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.EntriesLimit = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetId(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Id = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetType(v string) *ListUserAnalyzersResponseBodyResultDicts {
	s.Type = &v
	return s
}

func (s *ListUserAnalyzersResponseBodyResultDicts) SetUpdated(v int32) *ListUserAnalyzersResponseBodyResultDicts {
	s.Updated = &v
	return s
}

type ListUserAnalyzersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserAnalyzersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListUserAnalyzersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserAnalyzersResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzersResponse) SetHeaders(v map[string]*string) *ListUserAnalyzersResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzersResponse) SetStatusCode(v int32) *ListUserAnalyzersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAnalyzersResponse) SetBody(v *ListUserAnalyzersResponseBody) *ListUserAnalyzersResponse {
	s.Body = v
	return s
}

type ModifyAppGroupRequest struct {
	// currentVersion
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the instance.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupRequest) SetCurrentVersion(v string) *ModifyAppGroupRequest {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDescription(v string) *ModifyAppGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDomain(v string) *ModifyAppGroupRequest {
	s.Domain = &v
	return s
}

func (s *ModifyAppGroupRequest) SetResourceGroupId(v string) *ModifyAppGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyAppGroupRequest) SetDryRun(v bool) *ModifyAppGroupRequest {
	s.DryRun = &v
	return s
}

type ModifyAppGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Response parameters
	Result *ModifyAppGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBody) SetRequestId(v string) *ModifyAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupResponseBody) SetResult(v *ModifyAppGroupResponseBodyResult) *ModifyAppGroupResponseBody {
	s.Result = v
	return s
}

type ModifyAppGroupResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *ModifyAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResult) SetChargeType(v string) *ModifyAppGroupResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCreated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDescription(v string) *ModifyAppGroupResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetDomain(v string) *ModifyAppGroupResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetId(v string) *ModifyAppGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetLockMode(v string) *ModifyAppGroupResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetLockedByExpiration(v int32) *ModifyAppGroupResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetName(v string) *ModifyAppGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProcessingOrderId(v string) *ModifyAppGroupResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProduced(v int32) *ModifyAppGroupResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetProjectId(v string) *ModifyAppGroupResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetQuota(v *ModifyAppGroupResponseBodyResultQuota) *ModifyAppGroupResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetStatus(v string) *ModifyAppGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetType(v string) *ModifyAppGroupResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyAppGroupResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity.
	//
	// Unit: GB )
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ModifyAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupResponse) SetHeaders(v map[string]*string) *ModifyAppGroupResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupResponse) SetStatusCode(v int32) *ModifyAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppGroupResponse) SetBody(v *ModifyAppGroupResponseBody) *ModifyAppGroupResponse {
	s.Body = v
	return s
}

type ModifyAppGroupQuotaRequest struct {
	Body   *Quota `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool  `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyAppGroupQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaRequest) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaRequest) SetBody(v *Quota) *ModifyAppGroupQuotaRequest {
	s.Body = v
	return s
}

func (s *ModifyAppGroupQuotaRequest) SetDryRun(v bool) *ModifyAppGroupQuotaRequest {
	s.DryRun = &v
	return s
}

type ModifyAppGroupQuotaResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	Result *ModifyAppGroupQuotaResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyAppGroupQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBody) SetRequestId(v string) *ModifyAppGroupQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBody) SetResult(v *ModifyAppGroupQuotaResponseBodyResult) *ModifyAppGroupQuotaResponseBody {
	s.Result = v
	return s
}

type ModifyAppGroupQuotaResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode           *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	LockedByExpiration *int32  `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *ModifyAppGroupQuotaResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargeType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetChargingWay(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCommodityCode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCreated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetCurrentVersion(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetDescription(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetInstanceId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetLockMode(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetLockedByExpiration(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetName(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProcessingOrderId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProduced(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetProjectId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetQuota(v *ModifyAppGroupQuotaResponseBodyResultQuota) *ModifyAppGroupQuotaResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetStatus(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetSwitchedTime(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetUpdated(v int32) *ModifyAppGroupQuotaResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyAppGroupQuotaResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetComputeResource(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetDocSize(v int32) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResultQuota) SetSpec(v string) *ModifyAppGroupQuotaResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ModifyAppGroupQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyAppGroupQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyAppGroupQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppGroupQuotaResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppGroupQuotaResponse) SetHeaders(v map[string]*string) *ModifyAppGroupQuotaResponse {
	s.Headers = v
	return s
}

func (s *ModifyAppGroupQuotaResponse) SetStatusCode(v int32) *ModifyAppGroupQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAppGroupQuotaResponse) SetBody(v *ModifyAppGroupQuotaResponseBody) *ModifyAppGroupQuotaResponse {
	s.Body = v
	return s
}

type ModifyFirstRankRequest struct {
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyFirstRankRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankRequest) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankRequest) SetBody(v *FirstRank) *ModifyFirstRankRequest {
	s.Body = v
	return s
}

func (s *ModifyFirstRankRequest) SetDryRun(v bool) *ModifyFirstRankRequest {
	s.DryRun = &v
	return s
}

type ModifyFirstRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *ModifyFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBody) SetRequestId(v string) *ModifyFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyFirstRankResponseBody) SetResult(v *ModifyFirstRankResponseBodyResult) *ModifyFirstRankResponseBody {
	s.Result = v
	return s
}

type ModifyFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*ModifyFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ModifyFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResult) SetActive(v bool) *ModifyFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetDescription(v string) *ModifyFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetMeta(v []*ModifyFirstRankResponseBodyResultMeta) *ModifyFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *ModifyFirstRankResponseBodyResult) SetName(v string) *ModifyFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type ModifyFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see Rough sort functions.
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: \[-100000,100000] (excluding 0).
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ModifyFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetArg(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetAttribute(v string) *ModifyFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *ModifyFirstRankResponseBodyResultMeta) SetWeight(v float32) *ModifyFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type ModifyFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyFirstRankResponse) GoString() string {
	return s.String()
}

func (s *ModifyFirstRankResponse) SetHeaders(v map[string]*string) *ModifyFirstRankResponse {
	s.Headers = v
	return s
}

func (s *ModifyFirstRankResponse) SetStatusCode(v int32) *ModifyFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyFirstRankResponse) SetBody(v *ModifyFirstRankResponseBody) *ModifyFirstRankResponse {
	s.Body = v
	return s
}

type ModifyQueryProcessorRequest struct {
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifyQueryProcessorRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorRequest) SetBody(v interface{}) *ModifyQueryProcessorRequest {
	s.Body = v
	return s
}

func (s *ModifyQueryProcessorRequest) SetDryRun(v bool) *ModifyQueryProcessorRequest {
	s.DryRun = &v
	return s
}

type ModifyQueryProcessorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	Result *ModifyQueryProcessorResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifyQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBody) SetRequestId(v string) *ModifyQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyQueryProcessorResponseBody) SetResult(v *ModifyQueryProcessorResponseBodyResult) *ModifyQueryProcessorResponseBody {
	s.Result = v
	return s
}

type ModifyQueryProcessorResponseBodyResult struct {
	// Indicates whether the query analysis rule is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the rule was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry. Valid values:
	//
	// *   GENERAL
	// *   ECOMMERCE
	// *   IT_CONTENT
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The most recent update time.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifyQueryProcessorResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponseBodyResult) SetActive(v bool) *ModifyQueryProcessorResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetCreated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetDomain(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Domain = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetIndexes(v []*string) *ModifyQueryProcessorResponseBodyResult {
	s.Indexes = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetName(v string) *ModifyQueryProcessorResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetProcessors(v []map[string]interface{}) *ModifyQueryProcessorResponseBodyResult {
	s.Processors = v
	return s
}

func (s *ModifyQueryProcessorResponseBodyResult) SetUpdated(v int32) *ModifyQueryProcessorResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifyQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *ModifyQueryProcessorResponse) SetHeaders(v map[string]*string) *ModifyQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *ModifyQueryProcessorResponse) SetStatusCode(v int32) *ModifyQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyQueryProcessorResponse) SetBody(v *ModifyQueryProcessorResponseBody) *ModifyQueryProcessorResponse {
	s.Body = v
	return s
}

type ModifyScheduledTaskRequest struct {
	// 请求参数。
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyScheduledTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskRequest) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskRequest) SetBody(v interface{}) *ModifyScheduledTaskRequest {
	s.Body = v
	return s
}

type ModifyScheduledTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s ModifyScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponseBody) SetRequestId(v string) *ModifyScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScheduledTaskResponseBody) SetResult(v map[string]interface{}) *ModifyScheduledTaskResponseBody {
	s.Result = v
	return s
}

type ModifyScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *ModifyScheduledTaskResponse) SetHeaders(v map[string]*string) *ModifyScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *ModifyScheduledTaskResponse) SetStatusCode(v int32) *ModifyScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyScheduledTaskResponse) SetBody(v *ModifyScheduledTaskResponseBody) *ModifyScheduledTaskResponse {
	s.Body = v
	return s
}

type ModifySecondRankRequest struct {
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s ModifySecondRankRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankRequest) GoString() string {
	return s.String()
}

func (s *ModifySecondRankRequest) SetBody(v *SecondRank) *ModifySecondRankRequest {
	s.Body = v
	return s
}

func (s *ModifySecondRankRequest) SetDryRun(v bool) *ModifySecondRankRequest {
	s.DryRun = &v
	return s
}

type ModifySecondRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	Result *ModifySecondRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ModifySecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBody) SetRequestId(v string) *ModifySecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecondRankResponseBody) SetResult(v *ModifySecondRankResponseBodyResult) *ModifySecondRankResponseBody {
	s.Result = v
	return s
}

type ModifySecondRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active  *bool  `json:"active,omitempty" xml:"active,omitempty"`
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// *   true
	// *   false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression.
	//
	// You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the expression.
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Updated *int32  `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ModifySecondRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponseBodyResult) SetActive(v bool) *ModifySecondRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetCreated(v int32) *ModifySecondRankResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetDescription(v string) *ModifySecondRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetId(v string) *ModifySecondRankResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsDefault(v string) *ModifySecondRankResponseBodyResult {
	s.IsDefault = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetIsSys(v string) *ModifySecondRankResponseBodyResult {
	s.IsSys = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetMeta(v string) *ModifySecondRankResponseBodyResult {
	s.Meta = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetName(v string) *ModifySecondRankResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ModifySecondRankResponseBodyResult) SetUpdated(v int32) *ModifySecondRankResponseBodyResult {
	s.Updated = &v
	return s
}

type ModifySecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecondRankResponse) GoString() string {
	return s.String()
}

func (s *ModifySecondRankResponse) SetHeaders(v map[string]*string) *ModifySecondRankResponse {
	s.Headers = v
	return s
}

func (s *ModifySecondRankResponse) SetStatusCode(v int32) *ModifySecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecondRankResponse) SetBody(v *ModifySecondRankResponseBody) *ModifySecondRankResponse {
	s.Body = v
	return s
}

type PreviewModelRequest struct {
	// query
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
}

func (s PreviewModelRequest) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelRequest) GoString() string {
	return s.String()
}

func (s *PreviewModelRequest) SetQuery(v string) *PreviewModelRequest {
	s.Query = &v
	return s
}

type PreviewModelResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s PreviewModelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewModelResponseBody) SetRequestId(v string) *PreviewModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewModelResponseBody) SetResult(v []map[string]interface{}) *PreviewModelResponseBody {
	s.Result = v
	return s
}

func (s *PreviewModelResponseBody) SetTotalCount(v int64) *PreviewModelResponseBody {
	s.TotalCount = &v
	return s
}

type PreviewModelResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PreviewModelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PreviewModelResponse) String() string {
	return tea.Prettify(s)
}

func (s PreviewModelResponse) GoString() string {
	return s.String()
}

func (s *PreviewModelResponse) SetHeaders(v map[string]*string) *PreviewModelResponse {
	s.Headers = v
	return s
}

func (s *PreviewModelResponse) SetStatusCode(v int32) *PreviewModelResponse {
	s.StatusCode = &v
	return s
}

func (s *PreviewModelResponse) SetBody(v *PreviewModelResponseBody) *PreviewModelResponse {
	s.Body = v
	return s
}

type PushInterventionDictionaryEntriesRequest struct {
	Body   []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	DryRun *bool                    `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s PushInterventionDictionaryEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionDictionaryEntriesRequest) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesRequest) SetBody(v []map[string]interface{}) *PushInterventionDictionaryEntriesRequest {
	s.Body = v
	return s
}

func (s *PushInterventionDictionaryEntriesRequest) SetDryRun(v bool) *PushInterventionDictionaryEntriesRequest {
	s.DryRun = &v
	return s
}

type PushInterventionDictionaryEntriesResponseBody struct {
	RequestId *string   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s PushInterventionDictionaryEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetRequestId(v string) *PushInterventionDictionaryEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushInterventionDictionaryEntriesResponseBody) SetResult(v []*string) *PushInterventionDictionaryEntriesResponseBody {
	s.Result = v
	return s
}

type PushInterventionDictionaryEntriesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushInterventionDictionaryEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInterventionDictionaryEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushInterventionDictionaryEntriesResponse) SetHeaders(v map[string]*string) *PushInterventionDictionaryEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) SetStatusCode(v int32) *PushInterventionDictionaryEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PushInterventionDictionaryEntriesResponse) SetBody(v *PushInterventionDictionaryEntriesResponseBody) *PushInterventionDictionaryEntriesResponse {
	s.Body = v
	return s
}

type PushUserAnalyzerEntriesRequest struct {
	Entries []*PushUserAnalyzerEntriesRequestEntries `json:"entries,omitempty" xml:"entries,omitempty" type:"Repeated"`
	DryRun  *bool                                    `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s PushUserAnalyzerEntriesRequest) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesRequest) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesRequest) SetEntries(v []*PushUserAnalyzerEntriesRequestEntries) *PushUserAnalyzerEntriesRequest {
	s.Entries = v
	return s
}

func (s *PushUserAnalyzerEntriesRequest) SetDryRun(v bool) *PushUserAnalyzerEntriesRequest {
	s.DryRun = &v
	return s
}

type PushUserAnalyzerEntriesRequestEntries struct {
	Cmd          *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	Key          *string `json:"key,omitempty" xml:"key,omitempty"`
	SplitEnabled *bool   `json:"splitEnabled,omitempty" xml:"splitEnabled,omitempty"`
	Value        *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PushUserAnalyzerEntriesRequestEntries) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesRequestEntries) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetCmd(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Cmd = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetKey(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Key = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetSplitEnabled(v bool) *PushUserAnalyzerEntriesRequestEntries {
	s.SplitEnabled = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetValue(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Value = &v
	return s
}

type PushUserAnalyzerEntriesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// N/A
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s PushUserAnalyzerEntriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponseBody) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponseBody) SetRequestId(v string) *PushUserAnalyzerEntriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushUserAnalyzerEntriesResponseBody) SetResult(v map[string]interface{}) *PushUserAnalyzerEntriesResponseBody {
	s.Result = v
	return s
}

type PushUserAnalyzerEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PushUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PushUserAnalyzerEntriesResponse) String() string {
	return tea.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *PushUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) SetStatusCode(v int32) *PushUserAnalyzerEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) SetBody(v *PushUserAnalyzerEntriesResponseBody) *PushUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

type RankPreviewQueryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RankPreviewQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RankPreviewQueryResponseBody) GoString() string {
	return s.String()
}

func (s *RankPreviewQueryResponseBody) SetRequestId(v string) *RankPreviewQueryResponseBody {
	s.RequestId = &v
	return s
}

type RankPreviewQueryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RankPreviewQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RankPreviewQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s RankPreviewQueryResponse) GoString() string {
	return s.String()
}

func (s *RankPreviewQueryResponse) SetHeaders(v map[string]*string) *RankPreviewQueryResponse {
	s.Headers = v
	return s
}

func (s *RankPreviewQueryResponse) SetStatusCode(v int32) *RankPreviewQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *RankPreviewQueryResponse) SetBody(v *RankPreviewQueryResponseBody) *RankPreviewQueryResponse {
	s.Body = v
	return s
}

type ReleaseSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ReleaseSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponseBody) SetRequestId(v string) *ReleaseSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseSortScriptResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseSortScriptResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSortScriptResponse) SetHeaders(v map[string]*string) *ReleaseSortScriptResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSortScriptResponse) SetStatusCode(v int32) *ReleaseSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSortScriptResponse) SetBody(v *ReleaseSortScriptResponseBody) *ReleaseSortScriptResponse {
	s.Body = v
	return s
}

type RemoveAppResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppResponseBody) SetRequestId(v string) *RemoveAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppResponseBody) SetResult(v []*int32) *RemoveAppResponseBody {
	s.Result = v
	return s
}

type RemoveAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAppResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAppResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppResponse) SetHeaders(v map[string]*string) *RemoveAppResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppResponse) SetStatusCode(v int32) *RemoveAppResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppResponse) SetBody(v *RemoveAppResponseBody) *RemoveAppResponse {
	s.Body = v
	return s
}

type RemoveAppGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponseBody) SetRequestId(v string) *RemoveAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAppGroupResponseBody) SetResult(v []*int32) *RemoveAppGroupResponseBody {
	s.Result = v
	return s
}

type RemoveAppGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveAppGroupResponse) SetHeaders(v map[string]*string) *RemoveAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveAppGroupResponse) SetStatusCode(v int32) *RemoveAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveAppGroupResponse) SetBody(v *RemoveAppGroupResponseBody) *RemoveAppGroupResponse {
	s.Body = v
	return s
}

type RemoveDataCollectionResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveDataCollectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveDataCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponseBody) SetRequestId(v string) *RemoveDataCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveDataCollectionResponseBody) SetResult(v string) *RemoveDataCollectionResponseBody {
	s.Result = &v
	return s
}

type RemoveDataCollectionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveDataCollectionResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveDataCollectionResponse) GoString() string {
	return s.String()
}

func (s *RemoveDataCollectionResponse) SetHeaders(v map[string]*string) *RemoveDataCollectionResponse {
	s.Headers = v
	return s
}

func (s *RemoveDataCollectionResponse) SetStatusCode(v int32) *RemoveDataCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveDataCollectionResponse) SetBody(v *RemoveDataCollectionResponseBody) *RemoveDataCollectionResponse {
	s.Body = v
	return s
}

type RemoveFirstRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort expression.
	Result *RemoveFirstRankResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveFirstRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBody) SetRequestId(v string) *RemoveFirstRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFirstRankResponseBody) SetResult(v *RemoveFirstRankResponseBodyResult) *RemoveFirstRankResponseBody {
	s.Result = v
	return s
}

type RemoveFirstRankResponseBodyResult struct {
	// Indicates whether the expression is the default one.
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the expression.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*RemoveFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s RemoveFirstRankResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResult) SetActive(v bool) *RemoveFirstRankResponseBodyResult {
	s.Active = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetDescription(v string) *RemoveFirstRankResponseBodyResult {
	s.Description = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetMeta(v []*RemoveFirstRankResponseBodyResultMeta) *RemoveFirstRankResponseBodyResult {
	s.Meta = v
	return s
}

func (s *RemoveFirstRankResponseBodyResult) SetName(v string) *RemoveFirstRankResponseBodyResult {
	s.Name = &v
	return s
}

type RemoveFirstRankResponseBodyResultMeta struct {
	// The parameters that are used by a function in the expression.
	//
	// For more information, see Rough sort functions.
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see Rough sort functions.
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: \[-100000,100000] (excluding 0).
	Weight *float32 `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s RemoveFirstRankResponseBodyResultMeta) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponseBodyResultMeta) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetArg(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Arg = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetAttribute(v string) *RemoveFirstRankResponseBodyResultMeta {
	s.Attribute = &v
	return s
}

func (s *RemoveFirstRankResponseBodyResultMeta) SetWeight(v float32) *RemoveFirstRankResponseBodyResultMeta {
	s.Weight = &v
	return s
}

type RemoveFirstRankResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveFirstRankResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFirstRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveFirstRankResponse) SetHeaders(v map[string]*string) *RemoveFirstRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveFirstRankResponse) SetStatusCode(v int32) *RemoveFirstRankResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveFirstRankResponse) SetBody(v *RemoveFirstRankResponseBody) *RemoveFirstRankResponse {
	s.Body = v
	return s
}

type RemoveInterventionDictionaryResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the intervention dictionary.
	Result *RemoveInterventionDictionaryResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s RemoveInterventionDictionaryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBody) SetRequestId(v string) *RemoveInterventionDictionaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBody) SetResult(v *RemoveInterventionDictionaryResponseBodyResult) *RemoveInterventionDictionaryResponseBody {
	s.Result = v
	return s
}

type RemoveInterventionDictionaryResponseBodyResult struct {
	// The custom analyzer.
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the intervention dictionary.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// *   stopword: an intervention dictionary for stop word filtering
	// *   synonym: an intervention dictionary for synonym configuration
	// *   correction: an intervention dictionary for spelling correction
	// *   category_prediction: an intervention dictionary for category prediction
	// *   ner: an intervention dictionary for named entity recognition (NER)
	// *   term_weighting: an intervention dictionary for term weight analysis
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	Updated *string `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s RemoveInterventionDictionaryResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponseBodyResult) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetAnalyzer(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Analyzer = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetCreated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Created = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetName(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Name = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetType(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Type = &v
	return s
}

func (s *RemoveInterventionDictionaryResponseBodyResult) SetUpdated(v string) *RemoveInterventionDictionaryResponseBodyResult {
	s.Updated = &v
	return s
}

type RemoveInterventionDictionaryResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveInterventionDictionaryResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveInterventionDictionaryResponse) GoString() string {
	return s.String()
}

func (s *RemoveInterventionDictionaryResponse) SetHeaders(v map[string]*string) *RemoveInterventionDictionaryResponse {
	s.Headers = v
	return s
}

func (s *RemoveInterventionDictionaryResponse) SetStatusCode(v int32) *RemoveInterventionDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveInterventionDictionaryResponse) SetBody(v *RemoveInterventionDictionaryResponseBody) *RemoveInterventionDictionaryResponse {
	s.Body = v
	return s
}

type RemoveQueryProcessorResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveQueryProcessorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveQueryProcessorResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponseBody) SetRequestId(v string) *RemoveQueryProcessorResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveQueryProcessorResponseBody) SetResult(v string) *RemoveQueryProcessorResponseBody {
	s.Result = &v
	return s
}

type RemoveQueryProcessorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveQueryProcessorResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveQueryProcessorResponse) GoString() string {
	return s.String()
}

func (s *RemoveQueryProcessorResponse) SetHeaders(v map[string]*string) *RemoveQueryProcessorResponse {
	s.Headers = v
	return s
}

func (s *RemoveQueryProcessorResponse) SetStatusCode(v int32) *RemoveQueryProcessorResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveQueryProcessorResponse) SetBody(v *RemoveQueryProcessorResponseBody) *RemoveQueryProcessorResponse {
	s.Body = v
	return s
}

type RemoveScheduledTaskResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result []*int32 `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s RemoveScheduledTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveScheduledTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponseBody) SetRequestId(v string) *RemoveScheduledTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveScheduledTaskResponseBody) SetResult(v []*int32) *RemoveScheduledTaskResponseBody {
	s.Result = v
	return s
}

type RemoveScheduledTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveScheduledTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveScheduledTaskResponse) GoString() string {
	return s.String()
}

func (s *RemoveScheduledTaskResponse) SetHeaders(v map[string]*string) *RemoveScheduledTaskResponse {
	s.Headers = v
	return s
}

func (s *RemoveScheduledTaskResponse) SetStatusCode(v int32) *RemoveScheduledTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveScheduledTaskResponse) SetBody(v *RemoveScheduledTaskResponseBody) *RemoveScheduledTaskResponse {
	s.Body = v
	return s
}

type RemoveSearchStrategyResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s RemoveSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponseBody) SetRequestId(v string) *RemoveSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type RemoveSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *RemoveSearchStrategyResponse) SetHeaders(v map[string]*string) *RemoveSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *RemoveSearchStrategyResponse) SetStatusCode(v int32) *RemoveSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSearchStrategyResponse) SetBody(v *RemoveSearchStrategyResponseBody) *RemoveSearchStrategyResponse {
	s.Body = v
	return s
}

type RemoveSecondRankResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveSecondRankResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveSecondRankResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponseBody) SetRequestId(v string) *RemoveSecondRankResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveSecondRankResponseBody) SetResult(v map[string]interface{}) *RemoveSecondRankResponseBody {
	s.Result = v
	return s
}

type RemoveSecondRankResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveSecondRankResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveSecondRankResponse) GoString() string {
	return s.String()
}

func (s *RemoveSecondRankResponse) SetHeaders(v map[string]*string) *RemoveSecondRankResponse {
	s.Headers = v
	return s
}

func (s *RemoveSecondRankResponse) SetStatusCode(v int32) *RemoveSecondRankResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveSecondRankResponse) SetBody(v *RemoveSecondRankResponseBody) *RemoveSecondRankResponse {
	s.Body = v
	return s
}

type RemoveUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RemoveUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponseBody) SetRequestId(v string) *RemoveUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *RemoveUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type RemoveUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RemoveUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponse) SetHeaders(v map[string]*string) *RemoveUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserAnalyzerResponse) SetStatusCode(v int32) *RemoveUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserAnalyzerResponse) SetBody(v *RemoveUserAnalyzerResponseBody) *RemoveUserAnalyzerResponse {
	s.Body = v
	return s
}

type RenewAppGroupRequest struct {
	Body *PrepayOrderInfo `json:"body,omitempty" xml:"body,omitempty"`
	// Guaranteed request idempotence
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s RenewAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewAppGroupRequest) GoString() string {
	return s.String()
}

func (s *RenewAppGroupRequest) SetBody(v *PrepayOrderInfo) *RenewAppGroupRequest {
	s.Body = v
	return s
}

func (s *RenewAppGroupRequest) SetClientToken(v string) *RenewAppGroupRequest {
	s.ClientToken = &v
	return s
}

type RenewAppGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s RenewAppGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponseBody) SetRequestId(v string) *RenewAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppGroupResponseBody) SetResult(v bool) *RenewAppGroupResponseBody {
	s.Result = &v
	return s
}

type RenewAppGroupResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewAppGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewAppGroupResponse) GoString() string {
	return s.String()
}

func (s *RenewAppGroupResponse) SetHeaders(v map[string]*string) *RenewAppGroupResponse {
	s.Headers = v
	return s
}

func (s *RenewAppGroupResponse) SetStatusCode(v int32) *RenewAppGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewAppGroupResponse) SetBody(v *RenewAppGroupResponseBody) *RenewAppGroupResponse {
	s.Body = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	Result *ReplaceAppGroupCommodityCodeResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s ReplaceAppGroupCommodityCodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetRequestId(v string) *ReplaceAppGroupCommodityCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBody) SetResult(v *ReplaceAppGroupCommodityCodeResponseBodyResult) *ReplaceAppGroupCommodityCodeResponseBody {
	s.Result = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBodyResult struct {
	// The billing method of the application. Valid values:
	//
	// *   POSTPAY: pay-as-you-go
	// *   PREPAY: subscription
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// *   1: computing resources
	// *   2: queries per second (QPS)
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time.
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// *   0: The quotas are approved.
	// *   1: The quotas are being approved.
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// *   Unlock: The instance is not locked.
	// *   LockByExpiration: The instance is automatically locked after it expires.
	// *   ManualLock: The instance is manually locked.
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the application.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete for the instance. For example, an order is one that is initiated to create the instance or change the quotas or billing method.
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// *   0: The order is in progress.
	// *   1: The order is complete.
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The quota information.
	Quota *ReplaceAppGroupCommodityCodeResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the created fine sort expression.
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application. Valid values:
	//
	// *   producing
	// *   review_pending
	// *   config_pending
	// *   normal
	// *   frozen
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// *   standard: a standard application.
	// *   advance: an advanced application which is of an old application type. New applications cannot be of this type.
	// *   enhanced: an advanced application which is of a new application type.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the test group was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The status information.
	Versions []*string `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargeType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetChargingWay(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ChargingWay = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCommodityCode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCreated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetCurrentVersion(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.CurrentVersion = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetDescription(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetExpireOn(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ExpireOn = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetFirstRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.FirstRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetHasPendingQuotaReviewTask(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.HasPendingQuotaReviewTask = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetInstanceId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockMode(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockMode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetLockedByExpiration(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.LockedByExpiration = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetName(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetPendingSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.PendingSecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProcessingOrderId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProcessingOrderId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProduced(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Produced = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetProjectId(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.ProjectId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetQuota(v *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Quota = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSecondRankAlgoDeploymentId(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SecondRankAlgoDeploymentId = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetStatus(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetSwitchedTime(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.SwitchedTime = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetType(v string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetUpdated(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResult) SetVersions(v []*string) *ReplaceAppGroupCommodityCodeResponseBodyResult {
	s.Versions = v
	return s
}

type ReplaceAppGroupCommodityCodeResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// *   opensearch.share.junior: basic
	// *   opensearch.share.common: shared general-purpose
	// *   opensearch.share.compute: shared computing
	// *   opensearch.share.storage: shared storage
	// *   opensearch.private.common: exclusive general-purpose
	// *   opensearch.private.compute: exclusive computing
	// *   opensearch.private.storage: exclusive storage
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetComputeResource(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetDocSize(v int32) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponseBodyResultQuota) SetSpec(v string) *ReplaceAppGroupCommodityCodeResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type ReplaceAppGroupCommodityCodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReplaceAppGroupCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReplaceAppGroupCommodityCodeResponse) String() string {
	return tea.Prettify(s)
}

func (s ReplaceAppGroupCommodityCodeResponse) GoString() string {
	return s.String()
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetHeaders(v map[string]*string) *ReplaceAppGroupCommodityCodeResponse {
	s.Headers = v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetStatusCode(v int32) *ReplaceAppGroupCommodityCodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceAppGroupCommodityCodeResponse) SetBody(v *ReplaceAppGroupCommodityCodeResponseBody) *ReplaceAppGroupCommodityCodeResponse {
	s.Body = v
	return s
}

type SaveSortScriptFileRequest struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Version *int32  `json:"version,omitempty" xml:"version,omitempty"`
}

func (s SaveSortScriptFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveSortScriptFileRequest) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileRequest) SetContent(v string) *SaveSortScriptFileRequest {
	s.Content = &v
	return s
}

func (s *SaveSortScriptFileRequest) SetVersion(v int32) *SaveSortScriptFileRequest {
	s.Version = &v
	return s
}

type SaveSortScriptFileResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SaveSortScriptFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SaveSortScriptFileResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponseBody) SetRequestId(v string) *SaveSortScriptFileResponseBody {
	s.RequestId = &v
	return s
}

type SaveSortScriptFileResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SaveSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SaveSortScriptFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveSortScriptFileResponse) GoString() string {
	return s.String()
}

func (s *SaveSortScriptFileResponse) SetHeaders(v map[string]*string) *SaveSortScriptFileResponse {
	s.Headers = v
	return s
}

func (s *SaveSortScriptFileResponse) SetStatusCode(v int32) *SaveSortScriptFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveSortScriptFileResponse) SetBody(v *SaveSortScriptFileResponseBody) *SaveSortScriptFileResponse {
	s.Body = v
	return s
}

type StartSlowQueryAnalyzerResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s StartSlowQueryAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponseBody) SetRequestId(v string) *StartSlowQueryAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSlowQueryAnalyzerResponseBody) SetResult(v map[string]interface{}) *StartSlowQueryAnalyzerResponseBody {
	s.Result = v
	return s
}

type StartSlowQueryAnalyzerResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartSlowQueryAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartSlowQueryAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSlowQueryAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *StartSlowQueryAnalyzerResponse) SetHeaders(v map[string]*string) *StartSlowQueryAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) SetStatusCode(v int32) *StartSlowQueryAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSlowQueryAnalyzerResponse) SetBody(v *StartSlowQueryAnalyzerResponseBody) *StartSlowQueryAnalyzerResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags. You can specify a maximum number of 20 tags.
	Tag []*TagResourcesRequestTag `json:"tag,omitempty" xml:"tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// The key of the tag.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UnbindESUserAnalyzerRequest struct {
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindESUserAnalyzerRequest) String() string {
	return tea.Prettify(s)
}

func (s UnbindESUserAnalyzerRequest) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerRequest) SetBody(v interface{}) *UnbindESUserAnalyzerRequest {
	s.Body = v
	return s
}

type UnbindESUserAnalyzerResponseBody struct {
	RequestId *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindESUserAnalyzerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindESUserAnalyzerResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponseBody) SetRequestId(v string) *UnbindESUserAnalyzerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindESUserAnalyzerResponseBody) SetResult(v map[string]interface{}) *UnbindESUserAnalyzerResponseBody {
	s.Result = v
	return s
}

type UnbindESUserAnalyzerResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindESUserAnalyzerResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindESUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *UnbindESUserAnalyzerResponse) SetHeaders(v map[string]*string) *UnbindESUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *UnbindESUserAnalyzerResponse) SetStatusCode(v int32) *UnbindESUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindESUserAnalyzerResponse) SetBody(v *UnbindESUserAnalyzerResponseBody) *UnbindESUserAnalyzerResponse {
	s.Body = v
	return s
}

type UnbindEsInstanceResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data returned.
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UnbindEsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UnbindEsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponseBody) SetRequestId(v string) *UnbindEsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindEsInstanceResponseBody) SetResult(v map[string]interface{}) *UnbindEsInstanceResponseBody {
	s.Result = v
	return s
}

type UnbindEsInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UnbindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UnbindEsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UnbindEsInstanceResponse) GoString() string {
	return s.String()
}

func (s *UnbindEsInstanceResponse) SetHeaders(v map[string]*string) *UnbindEsInstanceResponse {
	s.Headers = v
	return s
}

func (s *UnbindEsInstanceResponse) SetStatusCode(v int32) *UnbindEsInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindEsInstanceResponse) SetBody(v *UnbindEsInstanceResponseBody) *UnbindEsInstanceResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Specifies whether to remove all tags from the specified one or more resources. This parameter takes effect only if the tagKey parameter is not specified. Valid values: true and false. Default value: false.
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The keys of tags. You can specify a maximum number of 20 keys.
	TagKey []*string `json:"tagKey,omitempty" xml:"tagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesShrinkRequest struct {
	// Specifies whether to remove all tags from the specified one or more resources. This parameter takes effect only if the tagKey parameter is not specified. Valid values: true and false. Default value: false.
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The keys of tags. You can specify a maximum number of 20 keys.
	TagKeyShrink *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
}

func (s UntagResourcesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesShrinkRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesShrinkRequest) SetAll(v bool) *UntagResourcesShrinkRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceIdShrink(v string) *UntagResourcesShrinkRequest {
	s.ResourceIdShrink = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetResourceType(v string) *UntagResourcesShrinkRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesShrinkRequest) SetTagKeyShrink(v string) *UntagResourcesShrinkRequest {
	s.TagKeyShrink = &v
	return s
}

type UntagResourcesResponseBody struct {
	// The ID of the request.
	TequestId *string `json:"tequestId,omitempty" xml:"tequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetTequestId(v string) *UntagResourcesResponseBody {
	s.TequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateABTestExperimentRequest struct {
	Body   *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool             `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestExperimentRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentRequest) SetBody(v *ABTestExperiment) *UpdateABTestExperimentRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestExperimentRequest) SetDryRun(v bool) *UpdateABTestExperimentRequest {
	s.DryRun = &v
	return s
}

type UpdateABTestExperimentResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test.
	Result *UpdateABTestExperimentResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestExperimentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBody) SetRequestId(v string) *UpdateABTestExperimentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestExperimentResponseBody) SetResult(v *UpdateABTestExperimentResponseBodyResult) *UpdateABTestExperimentResponseBody {
	s.Result = v
	return s
}

type UpdateABTestExperimentResponseBodyResult struct {
	// The time when the test was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// *   true: in effect
	// *   false: not in effect
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the test.
	//
	// Value values: 0 to 100.
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestExperimentResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponseBodyResult) SetCreated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetId(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetName(v string) *UpdateABTestExperimentResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetOnline(v bool) *UpdateABTestExperimentResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestExperimentResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetTraffic(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestExperimentResponseBodyResult) SetUpdated(v int32) *UpdateABTestExperimentResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestExperimentResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponse) SetHeaders(v map[string]*string) *UpdateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestExperimentResponse) SetStatusCode(v int32) *UpdateABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestExperimentResponse) SetBody(v *UpdateABTestExperimentResponseBody) *UpdateABTestExperimentResponse {
	s.Body = v
	return s
}

type UpdateABTestFixedFlowDividersRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s UpdateABTestFixedFlowDividersRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersRequest) SetBody(v []*string) *UpdateABTestFixedFlowDividersRequest {
	s.Body = v
	return s
}

type UpdateABTestFixedFlowDividersResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The whitelists after the update.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s UpdateABTestFixedFlowDividersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetRequestId(v string) *UpdateABTestFixedFlowDividersResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponseBody) SetResult(v []*string) *UpdateABTestFixedFlowDividersResponseBody {
	s.Result = v
	return s
}

type UpdateABTestFixedFlowDividersResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestFixedFlowDividersResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestFixedFlowDividersResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestFixedFlowDividersResponse) SetHeaders(v map[string]*string) *UpdateABTestFixedFlowDividersResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) SetStatusCode(v int32) *UpdateABTestFixedFlowDividersResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestFixedFlowDividersResponse) SetBody(v *UpdateABTestFixedFlowDividersResponseBody) *UpdateABTestFixedFlowDividersResponse {
	s.Body = v
	return s
}

type UpdateABTestGroupRequest struct {
	Body   *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool        `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupRequest) SetBody(v *ABTestGroup) *UpdateABTestGroupRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestGroupRequest) SetDryRun(v bool) *UpdateABTestGroupRequest {
	s.DryRun = &v
	return s
}

type UpdateABTestGroupResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test group.
	Result *UpdateABTestGroupResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBody) SetRequestId(v string) *UpdateABTestGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestGroupResponseBody) SetResult(v *UpdateABTestGroupResponseBodyResult) *UpdateABTestGroupResponseBody {
	s.Result = v
	return s
}

type UpdateABTestGroupResponseBodyResult struct {
	// The time when the test group was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// *   0: not in effect
	// *   1: in effect
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestGroupResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponseBodyResult) SetCreated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetId(v string) *UpdateABTestGroupResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetName(v string) *UpdateABTestGroupResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetStatus(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpdateABTestGroupResponseBodyResult) SetUpdated(v int32) *UpdateABTestGroupResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestGroupResponse) SetHeaders(v map[string]*string) *UpdateABTestGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestGroupResponse) SetStatusCode(v int32) *UpdateABTestGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestGroupResponse) SetBody(v *UpdateABTestGroupResponseBody) *UpdateABTestGroupResponse {
	s.Body = v
	return s
}

type UpdateABTestSceneRequest struct {
	Body   *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	DryRun *bool        `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestSceneRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneRequest) SetBody(v *ABTestScene) *UpdateABTestSceneRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestSceneRequest) SetDryRun(v bool) *UpdateABTestSceneRequest {
	s.DryRun = &v
	return s
}

type UpdateABTestSceneResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario. For more information, see [ABTestScene](https://www.alibabacloud.com/help/en/opensearch/latest/abtestscene).
	Result *UpdateABTestSceneResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s UpdateABTestSceneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBody) SetRequestId(v string) *UpdateABTestSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateABTestSceneResponseBody) SetResult(v *UpdateABTestSceneResponseBodyResult) *UpdateABTestSceneResponseBody {
	s.Result = v
	return s
}

type UpdateABTestSceneResponseBodyResult struct {
	// The time when the test scenario was created.
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test scenario.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test scenario.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	// - true: enabled
	// - false: stopped
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the A/B test.
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is allocated to the A/B test. Valid values: [0,100].
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test scenario was last modified.
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s UpdateABTestSceneResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponseBodyResult) SetCreated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Created = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetId(v string) *UpdateABTestSceneResponseBodyResult {
	s.Id = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetName(v string) *UpdateABTestSceneResponseBodyResult {
	s.Name = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetOnline(v bool) *UpdateABTestSceneResponseBodyResult {
	s.Online = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetParams(v map[string]interface{}) *UpdateABTestSceneResponseBodyResult {
	s.Params = v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetTraffic(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Traffic = &v
	return s
}

func (s *UpdateABTestSceneResponseBodyResult) SetUpdated(v int32) *UpdateABTestSceneResponseBodyResult {
	s.Updated = &v
	return s
}

type UpdateABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateABTestSceneResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponse) SetHeaders(v map[string]*string) *UpdateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestSceneResponse) SetStatusCode(v int32) *UpdateABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestSceneResponse) SetBody(v *UpdateABTestSceneResponseBody) *UpdateABTestSceneResponse {
	s.Body = v
	return s
}

type UpdateFetchFieldsRequest struct {
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateFetchFieldsRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsRequest) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsRequest) SetBody(v []*string) *UpdateFetchFieldsRequest {
	s.Body = v
	return s
}

func (s *UpdateFetchFieldsRequest) SetDryRun(v bool) *UpdateFetchFieldsRequest {
	s.DryRun = &v
	return s
}

type UpdateFetchFieldsResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateFetchFieldsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponseBody) SetRequestId(v string) *UpdateFetchFieldsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFetchFieldsResponseBody) SetResult(v bool) *UpdateFetchFieldsResponseBody {
	s.Result = &v
	return s
}

type UpdateFetchFieldsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFetchFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFetchFieldsResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFetchFieldsResponse) GoString() string {
	return s.String()
}

func (s *UpdateFetchFieldsResponse) SetHeaders(v map[string]*string) *UpdateFetchFieldsResponse {
	s.Headers = v
	return s
}

func (s *UpdateFetchFieldsResponse) SetStatusCode(v int32) *UpdateFetchFieldsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFetchFieldsResponse) SetBody(v *UpdateFetchFieldsResponseBody) *UpdateFetchFieldsResponse {
	s.Body = v
	return s
}

type UpdateFunctionDefaultInstanceRequest struct {
	// The name of the instance.
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
}

func (s UpdateFunctionDefaultInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceRequest) SetInstanceName(v string) *UpdateFunctionDefaultInstanceRequest {
	s.InstanceName = &v
	return s
}

type UpdateFunctionDefaultInstanceResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionDefaultInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetCode(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetLatency(v int64) *UpdateFunctionDefaultInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetMessage(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetRequestId(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponseBody) SetStatus(v string) *UpdateFunctionDefaultInstanceResponseBody {
	s.Status = &v
	return s
}

type UpdateFunctionDefaultInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionDefaultInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionDefaultInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionDefaultInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionDefaultInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) SetStatusCode(v int32) *UpdateFunctionDefaultInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionDefaultInstanceResponse) SetBody(v *UpdateFunctionDefaultInstanceResponseBody) *UpdateFunctionDefaultInstanceResponse {
	s.Body = v
	return s
}

type UpdateFunctionInstanceRequest struct {
	// The parameters that are used to create the instance.
	CreateParameters []*UpdateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The cron expression used to schedule periodic training, in the format of (Minutes Hours DayofMonth Month DayofWeek). The default value is empty, which indicates that no periodic training is performed. DayofWeek 0 indicates Sunday.
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description of the instance.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The parameters that are used.
	UsageParameters []*UpdateFunctionInstanceRequestUsageParameters `json:"usageParameters,omitempty" xml:"usageParameters,omitempty" type:"Repeated"`
}

func (s UpdateFunctionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequest) SetCreateParameters(v []*UpdateFunctionInstanceRequestCreateParameters) *UpdateFunctionInstanceRequest {
	s.CreateParameters = v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetCron(v string) *UpdateFunctionInstanceRequest {
	s.Cron = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetDescription(v string) *UpdateFunctionInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateFunctionInstanceRequest) SetUsageParameters(v []*UpdateFunctionInstanceRequestUsageParameters) *UpdateFunctionInstanceRequest {
	s.UsageParameters = v
	return s
}

type UpdateFunctionInstanceRequestCreateParameters struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestCreateParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceRequestCreateParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetName(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestCreateParameters) SetValue(v string) *UpdateFunctionInstanceRequestCreateParameters {
	s.Value = &v
	return s
}

type UpdateFunctionInstanceRequestUsageParameters struct {
	// The name of the parameter.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UpdateFunctionInstanceRequestUsageParameters) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceRequestUsageParameters) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceRequestUsageParameters) SetName(v string) *UpdateFunctionInstanceRequestUsageParameters {
	s.Name = &v
	return s
}

func (s *UpdateFunctionInstanceRequestUsageParameters) SetValue(v string) *UpdateFunctionInstanceRequestUsageParameters {
	s.Value = &v
	return s
}

type UpdateFunctionInstanceResponseBody struct {
	// The error code.
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// *       OK: The request was successful.
	// *       FAIL: The request failed.
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponseBody) SetCode(v string) *UpdateFunctionInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetHttpCode(v int64) *UpdateFunctionInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetLatency(v int64) *UpdateFunctionInstanceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetMessage(v string) *UpdateFunctionInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetRequestId(v string) *UpdateFunctionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionInstanceResponseBody) SetStatus(v string) *UpdateFunctionInstanceResponseBody {
	s.Status = &v
	return s
}

type UpdateFunctionInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateFunctionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionInstanceResponse) SetHeaders(v map[string]*string) *UpdateFunctionInstanceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionInstanceResponse) SetStatusCode(v int32) *UpdateFunctionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionInstanceResponse) SetBody(v *UpdateFunctionInstanceResponseBody) *UpdateFunctionInstanceResponse {
	s.Body = v
	return s
}

type UpdateSearchStrategyRequest struct {
	Body *SearchStrategy `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSearchStrategyRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyRequest) SetBody(v *SearchStrategy) *UpdateSearchStrategyRequest {
	s.Body = v
	return s
}

type UpdateSearchStrategyResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSearchStrategyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponseBody) SetRequestId(v string) *UpdateSearchStrategyResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSearchStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSearchStrategyResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSearchStrategyResponse) GoString() string {
	return s.String()
}

func (s *UpdateSearchStrategyResponse) SetHeaders(v map[string]*string) *UpdateSearchStrategyResponse {
	s.Headers = v
	return s
}

func (s *UpdateSearchStrategyResponse) SetStatusCode(v int32) *UpdateSearchStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSearchStrategyResponse) SetBody(v *UpdateSearchStrategyResponseBody) *UpdateSearchStrategyResponse {
	s.Body = v
	return s
}

type UpdateSortScriptResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateSortScriptResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSortScriptResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponseBody) SetRequestId(v string) *UpdateSortScriptResponseBody {
	s.RequestId = &v
	return s
}

type UpdateSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSortScriptResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSortScriptResponse) GoString() string {
	return s.String()
}

func (s *UpdateSortScriptResponse) SetHeaders(v map[string]*string) *UpdateSortScriptResponse {
	s.Headers = v
	return s
}

func (s *UpdateSortScriptResponse) SetStatusCode(v int32) *UpdateSortScriptResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSortScriptResponse) SetBody(v *UpdateSortScriptResponseBody) *UpdateSortScriptResponse {
	s.Body = v
	return s
}

type UpdateSummariesRequest struct {
	Body []*UpdateSummariesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateSummariesRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesRequest) GoString() string {
	return s.String()
}

func (s *UpdateSummariesRequest) SetBody(v []*UpdateSummariesRequestBody) *UpdateSummariesRequest {
	s.Body = v
	return s
}

func (s *UpdateSummariesRequest) SetDryRun(v bool) *UpdateSummariesRequest {
	s.DryRun = &v
	return s
}

type UpdateSummariesRequestBody struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *int32  `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s UpdateSummariesRequestBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesRequestBody) GoString() string {
	return s.String()
}

func (s *UpdateSummariesRequestBody) SetElement(v string) *UpdateSummariesRequestBody {
	s.Element = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetEllipsis(v string) *UpdateSummariesRequestBody {
	s.Ellipsis = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetField(v string) *UpdateSummariesRequestBody {
	s.Field = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetLen(v int32) *UpdateSummariesRequestBody {
	s.Len = &v
	return s
}

func (s *UpdateSummariesRequestBody) SetSnippet(v int32) *UpdateSummariesRequestBody {
	s.Snippet = &v
	return s
}

type UpdateSummariesResponseBody struct {
	// The ID of the request.
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s UpdateSummariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponseBody) SetRequestId(v string) *UpdateSummariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSummariesResponseBody) SetResult(v bool) *UpdateSummariesResponseBody {
	s.Result = &v
	return s
}

type UpdateSummariesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateSummariesResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateSummariesResponse) GoString() string {
	return s.String()
}

func (s *UpdateSummariesResponse) SetHeaders(v map[string]*string) *UpdateSummariesResponse {
	s.Headers = v
	return s
}

func (s *UpdateSummariesResponse) SetStatusCode(v int32) *UpdateSummariesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSummariesResponse) SetBody(v *UpdateSummariesResponseBody) *UpdateSummariesResponse {
	s.Body = v
	return s
}

type ValidateDataSourcesRequest struct {
	Body *DataSource `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesRequest) SetBody(v *DataSource) *ValidateDataSourcesRequest {
	s.Body = v
	return s
}

type ValidateDataSourcesResponseBody struct {
	RequestId *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []*ValidateDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
}

func (s ValidateDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBody) SetRequestId(v string) *ValidateDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateDataSourcesResponseBody) SetResult(v []*ValidateDataSourcesResponseBodyResult) *ValidateDataSourcesResponseBody {
	s.Result = v
	return s
}

type ValidateDataSourcesResponseBodyResult struct {
	Code       *string                                          `json:"code,omitempty" xml:"code,omitempty"`
	DataSource *ValidateDataSourcesResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	Message    *string                                          `json:"message,omitempty" xml:"message,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResult) SetCode(v string) *ValidateDataSourcesResponseBodyResult {
	s.Code = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetDataSource(v *ValidateDataSourcesResponseBodyResultDataSource) *ValidateDataSourcesResponseBodyResult {
	s.DataSource = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResult) SetMessage(v string) *ValidateDataSourcesResponseBodyResult {
	s.Message = &v
	return s
}

type ValidateDataSourcesResponseBodyResultDataSource struct {
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	TableName  *string                `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ValidateDataSourcesResponseBodyResultDataSource) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponseBodyResultDataSource) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetParameters(v map[string]interface{}) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Parameters = v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetTableName(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.TableName = &v
	return s
}

func (s *ValidateDataSourcesResponseBodyResultDataSource) SetType(v string) *ValidateDataSourcesResponseBodyResultDataSource {
	s.Type = &v
	return s
}

type ValidateDataSourcesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ValidateDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ValidateDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ValidateDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ValidateDataSourcesResponse) SetHeaders(v map[string]*string) *ValidateDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ValidateDataSourcesResponse) SetStatusCode(v int32) *ValidateDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateDataSourcesResponse) SetBody(v *ValidateDataSourcesResponseBody) *ValidateDataSourcesResponse {
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
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("opensearch"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) BindESUserAnalyzerWithOptions(appGroupIdentity *string, esInstanceId *string, request *BindESUserAnalyzerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindESUserAnalyzerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("BindESUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/es/" + tea.StringValue(openapiutil.GetEncodeParam(esInstanceId)) + "/actions/bind-analyzer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindESUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindESUserAnalyzer(appGroupIdentity *string, esInstanceId *string, request *BindESUserAnalyzerRequest) (_result *BindESUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindESUserAnalyzerResponse{}
	_body, _err := client.BindESUserAnalyzerWithOptions(appGroupIdentity, esInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BindEsInstanceWithOptions(appGroupIdentity *string, request *BindEsInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BindEsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Body)) {
		body["body"] = request.Body
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("BindEsInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/actions/bind-es-instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BindEsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BindEsInstance(appGroupIdentity *string, request *BindEsInstanceRequest) (_result *BindEsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BindEsInstanceResponse{}
	_body, _err := client.BindEsInstanceWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompileSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CompileSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CompileSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/actions/compiling"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CompileSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompileSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *CompileSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompileSortScriptResponse{}
	_body, _err := client.CompileSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, request *CreateABTestExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, request *CreateABTestExperimentRequest) (_result *CreateABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestExperimentResponse{}
	_body, _err := client.CreateABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, request *CreateABTestGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestGroup(appGroupIdentity *string, sceneId *string, request *CreateABTestGroupRequest) (_result *CreateABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestGroupResponse{}
	_body, _err := client.CreateABTestGroupWithOptions(appGroupIdentity, sceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateABTestSceneWithOptions(appGroupIdentity *string, request *CreateABTestSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateABTestSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateABTestScene(appGroupIdentity *string, request *CreateABTestSceneRequest) (_result *CreateABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateABTestSceneResponse{}
	_body, _err := client.CreateABTestSceneWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   When you create a standard application, a new version of the application is created if the specified application name already exists.
 * *   When you create a version of an existing application, you must set the autoSwitch and realtimeShared parameters.
 * *   When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
 * *   When you create a version of an existing application, the modification of the quota parameter does not take effect.
 *
 * @param request CreateAppRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateAppResponse
 */
func (client *Client) CreateAppWithOptions(appGroupIdentity *string, request *CreateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   When you create a standard application, a new version of the application is created if the specified application name already exists.
 * *   When you create a version of an existing application, you must set the autoSwitch and realtimeShared parameters.
 * *   When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
 * *   When you create a version of an existing application, the modification of the quota parameter does not take effect.
 *
 * @param request CreateAppRequest
 * @return CreateAppResponse
 */
func (client *Client) CreateApp(appGroupIdentity *string, request *CreateAppRequest) (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppGroupWithOptions(request *CreateAppGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAppGroup(request *CreateAppGroupRequest) (_result *CreateAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppGroupResponse{}
	_body, _err := client.CreateAppGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFirstRankWithOptions(appGroupIdentity *string, appId *string, request *CreateFirstRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFirstRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/first-ranks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFirstRank(appGroupIdentity *string, appId *string, request *CreateFirstRankRequest) (_result *CreateFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFirstRankResponse{}
	_body, _err := client.CreateFirstRankWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call the [GetFunctionCurrentVersion](~~421377~~) operation to query the latest version of the current feature. The response of the operation includes the createParameters parameter that is used to create an algorithm instance, the usageParameters parameter, and the requirements for setting these parameters.
 *
 * @param request CreateFunctionInstanceRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateFunctionInstanceResponse
 */
func (client *Client) CreateFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, request *CreateFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateParameters)) {
		body["createParameters"] = request.CreateParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Cron)) {
		body["cron"] = request.Cron
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		body["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		body["modelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.UsageParameters)) {
		body["usageParameters"] = request.UsageParameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call the [GetFunctionCurrentVersion](~~421377~~) operation to query the latest version of the current feature. The response of the operation includes the createParameters parameter that is used to create an algorithm instance, the usageParameters parameter, and the requirements for setting these parameters.
 *
 * @param request CreateFunctionInstanceRequest
 * @return CreateFunctionInstanceResponse
 */
func (client *Client) CreateFunctionInstance(appGroupIdentity *string, functionName *string, request *CreateFunctionInstanceRequest) (_result *CreateFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionInstanceResponse{}
	_body, _err := client.CreateFunctionInstanceWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateFunctionTaskWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName)) + "/tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateFunctionTask(appGroupIdentity *string, functionName *string, instanceName *string) (_result *CreateFunctionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionTaskResponse{}
	_body, _err := client.CreateFunctionTaskWithOptions(appGroupIdentity, functionName, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInterventionDictionaryWithOptions(request *CreateInterventionDictionaryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateInterventionDictionaryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AnalyzerType)) {
		body["analyzerType"] = request.AnalyzerType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInterventionDictionary(request *CreateInterventionDictionaryRequest) (_result *CreateInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateInterventionDictionaryResponse{}
	_body, _err := client.CreateInterventionDictionaryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQueryProcessorWithOptions(appGroupIdentity *string, appId *string, request *CreateQueryProcessorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQueryProcessorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQueryProcessor(appGroupIdentity *string, appId *string, request *CreateQueryProcessorRequest) (_result *CreateQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQueryProcessorResponse{}
	_body, _err := client.CreateQueryProcessorWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ****
 *
 * @param request CreateScheduledTaskRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return CreateScheduledTaskResponse
 */
func (client *Client) CreateScheduledTaskWithOptions(appGroupIdentity *string, request *CreateScheduledTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scheduled-tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ****
 *
 * @param request CreateScheduledTaskRequest
 * @return CreateScheduledTaskResponse
 */
func (client *Client) CreateScheduledTask(appGroupIdentity *string, request *CreateScheduledTaskRequest) (_result *CreateScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateScheduledTaskResponse{}
	_body, _err := client.CreateScheduledTaskWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSearchStrategyWithOptions(appGroupIdentity *string, appId *string, request *CreateSearchStrategyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSearchStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/search-strategies"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSearchStrategy(appGroupIdentity *string, appId *string, request *CreateSearchStrategyRequest) (_result *CreateSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSearchStrategyResponse{}
	_body, _err := client.CreateSearchStrategyWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecondRankWithOptions(appGroupIdentity *string, appId *string, request *CreateSecondRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSecondRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/second-ranks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecondRank(appGroupIdentity *string, appId *string, request *CreateSecondRankRequest) (_result *CreateSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSecondRankResponse{}
	_body, _err := client.CreateSecondRankWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSortScriptWithOptions(appGroupIdentity *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSortScript(appGroupIdentity *string, appVersionId *string) (_result *CreateSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CreateSortScriptWithOptions(appGroupIdentity, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateUserAnalyzerWithOptions(request *CreateUserAnalyzerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserAnalyzerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Business)) {
		body["business"] = request.Business
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessAppGroupId)) {
		body["businessAppGroupId"] = request.BusinessAppGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessType)) {
		body["businessType"] = request.BusinessType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateUserAnalyzer(request *CreateUserAnalyzerRequest) (_result *CreateUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserAnalyzerResponse{}
	_body, _err := client.CreateUserAnalyzerWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(experimentId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *DeleteABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestExperimentResponse{}
	_body, _err := client.DeleteABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string) (_result *DeleteABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestGroupResponse{}
	_body, _err := client.DeleteABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteABTestSceneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteABTestScene(appGroupIdentity *string, sceneId *string) (_result *DeleteABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteABTestSceneResponse{}
	_body, _err := client.DeleteABTestSceneWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string) (_result *DeleteFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionInstanceResponse{}
	_body, _err := client.DeleteFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFunctionTaskWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, generation *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(generation))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFunctionTask(appGroupIdentity *string, functionName *string, instanceName *string, generation *string) (_result *DeleteFunctionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionTaskResponse{}
	_body, _err := client.DeleteFunctionTaskWithOptions(appGroupIdentity, functionName, instanceName, generation, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *DeleteSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSortScriptResponse{}
	_body, _err := client.DeleteSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSortScriptFileWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, fileName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSortScriptFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/files/src/" + tea.StringValue(openapiutil.GetEncodeParam(fileName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSortScriptFile(appGroupIdentity *string, appVersionId *string, scriptName *string, fileName *string) (_result *DeleteSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSortScriptFileResponse{}
	_body, _err := client.DeleteSortScriptFileWithOptions(appGroupIdentity, appVersionId, scriptName, fileName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestExperimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(experimentId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *DescribeABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestExperimentResponse{}
	_body, _err := client.DescribeABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string) (_result *DescribeABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestGroupResponse{}
	_body, _err := client.DescribeABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeABTestSceneResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeABTestScene(appGroupIdentity *string, sceneId *string) (_result *DescribeABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeABTestSceneResponse{}
	_body, _err := client.DescribeABTestSceneWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApp(appGroupIdentity *string, appId *string) (_result *DescribeAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppResponse{}
	_body, _err := client.DescribeAppWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppGroup(appGroupIdentity *string) (_result *DescribeAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppGroupResponse{}
	_body, _err := client.DescribeAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppStatisticsWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppStatisticsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAppStatistics"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/statistics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppStatisticsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAppStatistics(appGroupIdentity *string, appId *string) (_result *DescribeAppStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppStatisticsResponse{}
	_body, _err := client.DescribeAppStatisticsWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAppsWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeAppsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeApps"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApps(appGroupIdentity *string) (_result *DescribeAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeAppsResponse{}
	_body, _err := client.DescribeAppsWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDataCollctionWithOptions(appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeDataCollctionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDataCollction"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/data-collections/" + tea.StringValue(openapiutil.GetEncodeParam(dataCollectionIdentity))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDataCollctionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDataCollction(appGroupIdentity *string, dataCollectionIdentity *string) (_result *DescribeDataCollctionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeDataCollctionResponse{}
	_body, _err := client.DescribeDataCollctionWithOptions(appGroupIdentity, dataCollectionIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeFirstRankResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/first-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeFirstRank(appGroupIdentity *string, appId *string, name *string) (_result *DescribeFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeFirstRankResponse{}
	_body, _err := client.DescribeFirstRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInterventionDictionaryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeInterventionDictionaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInterventionDictionary(name *string) (_result *DescribeInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeInterventionDictionaryResponse{}
	_body, _err := client.DescribeInterventionDictionaryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeQueryProcessorResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeQueryProcessor(appGroupIdentity *string, appId *string, name *string) (_result *DescribeQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeQueryProcessorResponse{}
	_body, _err := client.DescribeQueryProcessorWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/region"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegion() (_result *DescribeRegionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionResponse{}
	_body, _err := client.DescribeRegionWithOptions(headers, runtime)
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
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/regions"),
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

func (client *Client) DescribeScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeScheduledTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scheduled-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeScheduledTask(appGroupIdentity *string, taskId *string) (_result *DescribeScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeScheduledTaskResponse{}
	_body, _err := client.DescribeScheduledTaskWithOptions(appGroupIdentity, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSecondRankResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/second-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecondRank(appGroupIdentity *string, appId *string, name *string) (_result *DescribeSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSecondRankResponse{}
	_body, _err := client.DescribeSecondRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowQueryStatusWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeSlowQueryStatusResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowQueryStatus"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowQueryStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowQueryStatus(appGroupIdentity *string) (_result *DescribeSlowQueryStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeSlowQueryStatusResponse{}
	_body, _err := client.DescribeSlowQueryStatusWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserAnalyzerWithOptions(name *string, request *DescribeUserAnalyzerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeUserAnalyzerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.With)) {
		query["with"] = request.With
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserAnalyzer(name *string, request *DescribeUserAnalyzerRequest) (_result *DescribeUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeUserAnalyzerResponse{}
	_body, _err := client.DescribeUserAnalyzerWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DisableSlowQueryWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DisableSlowQueryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DisableSlowQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query/actions/disable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DisableSlowQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DisableSlowQuery(appGroupIdentity *string) (_result *DisableSlowQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DisableSlowQueryResponse{}
	_body, _err := client.DisableSlowQueryWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EnableSlowQueryWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EnableSlowQueryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("EnableSlowQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query/actions/enable"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &EnableSlowQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EnableSlowQuery(appGroupIdentity *string) (_result *EnableSlowQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EnableSlowQueryResponse{}
	_body, _err := client.EnableSlowQueryWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GenerateMergedTableWithOptions(request *GenerateMergedTableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GenerateMergedTableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Spec)) {
		query["spec"] = request.Spec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("GenerateMergedTable"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/schema/actions/merge"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GenerateMergedTableResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GenerateMergedTable(request *GenerateMergedTableRequest) (_result *GenerateMergedTableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GenerateMergedTableResponse{}
	_body, _err := client.GenerateMergedTableWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDomainWithOptions(domainName *string, request *GetDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AppGroupIdentity)) {
		query["appGroupIdentity"] = request.AppGroupIdentity
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/domains/" + tea.StringValue(openapiutil.GetEncodeParam(domainName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDomain(domainName *string, request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(domainName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionCurrentVersionWithOptions(functionName *string, request *GetFunctionCurrentVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionCurrentVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		query["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		query["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["modelType"] = request.ModelType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionCurrentVersion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/current-version"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionCurrentVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionCurrentVersion(functionName *string, request *GetFunctionCurrentVersionRequest) (_result *GetFunctionCurrentVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionCurrentVersionResponse{}
	_body, _err := client.GetFunctionCurrentVersionWithOptions(functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionDefaultInstanceWithOptions(appGroupIdentity *string, functionName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionDefaultInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionDefaultInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/default-instance"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionDefaultInstance(appGroupIdentity *string, functionName *string) (_result *GetFunctionDefaultInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionDefaultInstanceResponse{}
	_body, _err := client.GetFunctionDefaultInstanceWithOptions(appGroupIdentity, functionName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string, request *GetFunctionInstanceRequest) (_result *GetFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionInstanceResponse{}
	_body, _err := client.GetFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionTaskWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, generation *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName)) + "/tasks/" + tea.StringValue(openapiutil.GetEncodeParam(generation))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionTask(appGroupIdentity *string, functionName *string, instanceName *string, generation *string) (_result *GetFunctionTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionTaskResponse{}
	_body, _err := client.GetFunctionTaskWithOptions(appGroupIdentity, functionName, instanceName, generation, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFunctionVersionWithOptions(functionName *string, versionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionVersionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetFunctionVersion"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/versions/" + tea.StringValue(openapiutil.GetEncodeParam(versionId))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFunctionVersion(functionName *string, versionId *string) (_result *GetFunctionVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionVersionResponse{}
	_body, _err := client.GetFunctionVersionWithOptions(functionName, versionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetModelReportWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetModelReportResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetModelReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/algorithm/models/" + tea.StringValue(openapiutil.GetEncodeParam(modelName)) + "/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetModelReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetModelReport(appGroupIdentity *string, modelName *string) (_result *GetModelReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetModelReportResponse{}
	_body, _err := client.GetModelReportWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetScriptFileNamesWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetScriptFileNamesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetScriptFileNames"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/file-names"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetScriptFileNamesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetScriptFileNames(appGroupIdentity *string, appVersionId *string, scriptName *string) (_result *GetScriptFileNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetScriptFileNamesResponse{}
	_body, _err := client.GetScriptFileNamesWithOptions(appGroupIdentity, appVersionId, scriptName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSearchStrategyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/search-strategies/" + tea.StringValue(openapiutil.GetEncodeParam(strategyName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string) (_result *GetSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSearchStrategyResponse{}
	_body, _err := client.GetSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *GetSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSortScriptResponse{}
	_body, _err := client.GetSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSortScriptFileWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSortScriptFileResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/files/src/" + tea.StringValue(openapiutil.GetEncodeParam(fileName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSortScriptFile(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string) (_result *GetSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSortScriptFileResponse{}
	_body, _err := client.GetSortScriptFileWithOptions(appGroupIdentity, scriptName, appVersionId, fileName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestExperimentsWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestExperimentsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestExperiments"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestExperimentsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestExperiments(appGroupIdentity *string, sceneId *string, groupId *string) (_result *ListABTestExperimentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestExperimentsResponse{}
	_body, _err := client.ListABTestExperimentsWithOptions(appGroupIdentity, sceneId, groupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestFixedFlowDividersWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestFixedFlowDividersResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestFixedFlowDividers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(experimentId)) + "/fixed-flow-dividers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestFixedFlowDividers(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string) (_result *ListABTestFixedFlowDividersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestFixedFlowDividersResponse{}
	_body, _err := client.ListABTestFixedFlowDividersWithOptions(appGroupIdentity, sceneId, groupId, experimentId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestGroupsWithOptions(appGroupIdentity *string, sceneId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestGroupsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestGroups"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestGroups(appGroupIdentity *string, sceneId *string) (_result *ListABTestGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestGroupsResponse{}
	_body, _err := client.ListABTestGroupsWithOptions(appGroupIdentity, sceneId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListABTestScenesWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListABTestScenesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListABTestScenes"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListABTestScenesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListABTestScenes(appGroupIdentity *string) (_result *ListABTestScenesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListABTestScenesResponse{}
	_body, _err := client.ListABTestScenesWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * *   This operation allows you to query applications by application name, instance ID, and application type.
 * *   This operation can sort the applications based on their creation time.
 * *   This operation supports the parameters for paging.
 *
 * @param tmpReq ListAppGroupsRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ListAppGroupsResponse
 */
func (client *Client) ListAppGroupsWithOptions(tmpReq *ListAppGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppGroupsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListAppGroupsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["instanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		query["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAppGroups"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAppGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * *   This operation allows you to query applications by application name, instance ID, and application type.
 * *   This operation can sort the applications based on their creation time.
 * *   This operation supports the parameters for paging.
 *
 * @param request ListAppGroupsRequest
 * @return ListAppGroupsResponse
 */
func (client *Client) ListAppGroups(request *ListAppGroupsRequest) (_result *ListAppGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppGroupsResponse{}
	_body, _err := client.ListAppGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(request *ListAppsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Group)) {
		query["group"] = request.Group
	}

	if !tea.BoolValue(util.IsUnset(request.Page)) {
		query["page"] = request.Page
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		query["size"] = request.Size
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListApps"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/apps"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("none"),
	}
	_result = &ListAppsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps(request *ListAppsRequest) (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataCollectionsWithOptions(appGroupIdentity *string, request *ListDataCollectionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataCollectionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataCollections"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/data-collections"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataCollectionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataCollections(appGroupIdentity *string, request *ListDataCollectionsRequest) (_result *ListDataCollectionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataCollectionsResponse{}
	_body, _err := client.ListDataCollectionsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceTableFieldsWithOptions(dataSourceType *string, request *ListDataSourceTableFieldsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTableFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["params"] = request.Params
	}

	if !tea.BoolValue(util.IsUnset(request.RawType)) {
		query["rawType"] = request.RawType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTableFields"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceType)) + "/fields"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTableFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceTableFields(dataSourceType *string, request *ListDataSourceTableFieldsRequest) (_result *ListDataSourceTableFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTableFieldsResponse{}
	_body, _err := client.ListDataSourceTableFieldsWithOptions(dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourceTablesWithOptions(dataSourceType *string, request *ListDataSourceTablesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourceTablesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Params)) {
		query["params"] = request.Params
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSourceTables"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/" + tea.StringValue(openapiutil.GetEncodeParam(dataSourceType)) + "/tables"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourceTablesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListDataSourceTables(dataSourceType *string, request *ListDataSourceTablesRequest) (_result *ListDataSourceTablesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourceTablesResponse{}
	_body, _err := client.ListDataSourceTablesWithOptions(dataSourceType, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFirstRanksWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFirstRanksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListFirstRanks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/first-ranks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFirstRanksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFirstRanks(appGroupIdentity *string, appId *string) (_result *ListFirstRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFirstRanksResponse{}
	_body, _err := client.ListFirstRanksWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionInstancesWithOptions(appGroupIdentity *string, functionName *string, request *ListFunctionInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FunctionType)) {
		query["functionType"] = request.FunctionType
	}

	if !tea.BoolValue(util.IsUnset(request.ModelType)) {
		query["modelType"] = request.ModelType
	}

	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Source)) {
		query["source"] = request.Source
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionInstances"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionInstances(appGroupIdentity *string, functionName *string, request *ListFunctionInstancesRequest) (_result *ListFunctionInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionInstancesResponse{}
	_body, _err := client.ListFunctionInstancesWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFunctionTasksWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *ListFunctionTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName)) + "/tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFunctionTasks(appGroupIdentity *string, functionName *string, instanceName *string, request *ListFunctionTasksRequest) (_result *ListFunctionTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionTasksResponse{}
	_body, _err := client.ListFunctionTasksWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionariesWithOptions(request *ListInterventionDictionariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Types)) {
		query["types"] = request.Types
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaries(request *ListInterventionDictionariesRequest) (_result *ListInterventionDictionariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionariesResponse{}
	_body, _err := client.ListInterventionDictionariesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryEntriesWithOptions(name *string, request *ListInterventionDictionaryEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		query["word"] = request.Word
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/entries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryEntries(name *string, request *ListInterventionDictionaryEntriesRequest) (_result *ListInterventionDictionaryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryEntriesResponse{}
	_body, _err := client.ListInterventionDictionaryEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryNerResultsWithOptions(name *string, request *ListInterventionDictionaryNerResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryNerResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryNerResults"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/ner-results"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryNerResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryNerResults(name *string, request *ListInterventionDictionaryNerResultsRequest) (_result *ListInterventionDictionaryNerResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryNerResultsResponse{}
	_body, _err := client.ListInterventionDictionaryNerResultsWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInterventionDictionaryRelatedEntitiesWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInterventionDictionaryRelatedEntitiesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListInterventionDictionaryRelatedEntities"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/related"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInterventionDictionaryRelatedEntitiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInterventionDictionaryRelatedEntities(name *string) (_result *ListInterventionDictionaryRelatedEntitiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInterventionDictionaryRelatedEntitiesResponse{}
	_body, _err := client.ListInterventionDictionaryRelatedEntitiesWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListModelsWithOptions(appGroupIdentity *string, request *ListModelsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListModelsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListModels"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/algorithm/models"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListModelsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListModels(appGroupIdentity *string, request *ListModelsRequest) (_result *ListModelsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListModelsResponse{}
	_body, _err := client.ListModelsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListProceedingsWithOptions(appGroupIdentity *string, request *ListProceedingsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListProceedingsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterFinished)) {
		query["filterFinished"] = request.FilterFinished
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListProceedings"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/proceedings"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListProceedingsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListProceedings(appGroupIdentity *string, request *ListProceedingsRequest) (_result *ListProceedingsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListProceedingsResponse{}
	_body, _err := client.ListProceedingsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorAnalyzerResultsWithOptions(appGroupIdentity *string, appId *string, name *string, request *ListQueryProcessorAnalyzerResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorAnalyzerResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Text)) {
		query["text"] = request.Text
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessorAnalyzerResults"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/analyze"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorAnalyzerResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessorAnalyzerResults(appGroupIdentity *string, appId *string, name *string, request *ListQueryProcessorAnalyzerResultsRequest) (_result *ListQueryProcessorAnalyzerResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorAnalyzerResultsResponse{}
	_body, _err := client.ListQueryProcessorAnalyzerResultsWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorNersWithOptions(request *ListQueryProcessorNersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorNersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		query["domain"] = request.Domain
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessorNers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/query-processor/ner/default-priorities"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorNersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessorNers(request *ListQueryProcessorNersRequest) (_result *ListQueryProcessorNersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorNersResponse{}
	_body, _err := client.ListQueryProcessorNersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQueryProcessorsWithOptions(appGroupIdentity *string, appId *string, request *ListQueryProcessorsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQueryProcessorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IsActive)) {
		query["isActive"] = request.IsActive
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQueryProcessors"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQueryProcessorsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQueryProcessors(appGroupIdentity *string, appId *string, request *ListQueryProcessorsRequest) (_result *ListQueryProcessorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQueryProcessorsResponse{}
	_body, _err := client.ListQueryProcessorsWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotaReviewTasksWithOptions(appGroupIdentity *string, request *ListQuotaReviewTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotaReviewTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListQuotaReviewTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/quota-review-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListQuotaReviewTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotaReviewTasks(appGroupIdentity *string, request *ListQuotaReviewTasksRequest) (_result *ListQuotaReviewTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotaReviewTasksResponse{}
	_body, _err := client.ListQuotaReviewTasksWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListScheduledTasksWithOptions(appGroupIdentity *string, request *ListScheduledTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListScheduledTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		query["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListScheduledTasks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scheduled-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListScheduledTasks(appGroupIdentity *string, request *ListScheduledTasksRequest) (_result *ListScheduledTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListScheduledTasksResponse{}
	_body, _err := client.ListScheduledTasksWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSearchStrategiesWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSearchStrategiesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSearchStrategies"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/search-strategies"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSearchStrategiesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSearchStrategies(appGroupIdentity *string, appId *string) (_result *ListSearchStrategiesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSearchStrategiesResponse{}
	_body, _err := client.ListSearchStrategiesWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecondRanksWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSecondRanksResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSecondRanks"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/second-ranks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSecondRanksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecondRanks(appGroupIdentity *string, appId *string) (_result *ListSecondRanksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSecondRanksResponse{}
	_body, _err := client.ListSecondRanksWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlowQueryCategoriesWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSlowQueryCategoriesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlowQueryCategories"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query/categories"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlowQueryCategoriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlowQueryCategories(appGroupIdentity *string) (_result *ListSlowQueryCategoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlowQueryCategoriesResponse{}
	_body, _err := client.ListSlowQueryCategoriesWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSlowQueryQueriesWithOptions(appGroupIdentity *string, categoryIndex *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSlowQueryQueriesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSlowQueryQueries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query/categories/" + tea.StringValue(openapiutil.GetEncodeParam(categoryIndex)) + "/queries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSlowQueryQueriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSlowQueryQueries(appGroupIdentity *string, categoryIndex *string) (_result *ListSlowQueryQueriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSlowQueryQueriesResponse{}
	_body, _err := client.ListSlowQueryQueriesWithOptions(appGroupIdentity, categoryIndex, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSortExpressionsWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSortExpressionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSortExpressions"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/sort-expressions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSortExpressionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSortExpressions(appGroupIdentity *string, appId *string) (_result *ListSortExpressionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSortExpressionsResponse{}
	_body, _err := client.ListSortExpressionsWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSortScriptsWithOptions(appGroupIdentity *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSortScriptsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListSortScripts"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSortScriptsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSortScripts(appGroupIdentity *string, appVersionId *string) (_result *ListSortScriptsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSortScriptsResponse{}
	_body, _err := client.ListSortScriptsWithOptions(appGroupIdentity, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatisticLogsWithOptions(appGroupIdentity *string, moduleName *string, request *ListStatisticLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStatisticLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		query["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.Distinct)) {
		query["distinct"] = request.Distinct
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["sortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StopTime)) {
		query["stopTime"] = request.StopTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStatisticLogs"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/statistic-logs/" + tea.StringValue(openapiutil.GetEncodeParam(moduleName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStatisticLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatisticLogs(appGroupIdentity *string, moduleName *string, request *ListStatisticLogsRequest) (_result *ListStatisticLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStatisticLogsResponse{}
	_body, _err := client.ListStatisticLogsWithOptions(appGroupIdentity, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStatisticReportWithOptions(appGroupIdentity *string, moduleName *string, request *ListStatisticReportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStatisticReportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Columns)) {
		query["columns"] = request.Columns
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListStatisticReport"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/statistic-report/" + tea.StringValue(openapiutil.GetEncodeParam(moduleName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStatisticReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStatisticReport(appGroupIdentity *string, moduleName *string, request *ListStatisticReportRequest) (_result *ListStatisticReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStatisticReportResponse{}
	_body, _err := client.ListStatisticReportWithOptions(appGroupIdentity, moduleName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(tmpReq *ListTagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListTagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("resourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Tag)) {
		request.TagShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tag, tea.String("tag"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["nextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagShrink)) {
		query["tag"] = request.TagShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/resource-tags"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserAnalyzerEntriesWithOptions(name *string, request *ListUserAnalyzerEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserAnalyzerEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Word)) {
		query["word"] = request.Word
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserAnalyzerEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/entries"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserAnalyzerEntries(name *string, request *ListUserAnalyzerEntriesRequest) (_result *ListUserAnalyzerEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserAnalyzerEntriesResponse{}
	_body, _err := client.ListUserAnalyzerEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListUserAnalyzersWithOptions(request *ListUserAnalyzersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserAnalyzersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUserAnalyzers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserAnalyzersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListUserAnalyzers(request *ListUserAnalyzersRequest) (_result *ListUserAnalyzersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserAnalyzersResponse{}
	_body, _err := client.ListUserAnalyzersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Debugging
 * [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=OpenSearch\\&api=ModifyAppGroup\\&type=ROA\\&version=2017-12-25)
 *
 * @param request ModifyAppGroupRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return ModifyAppGroupResponse
 */
func (client *Client) ModifyAppGroupWithOptions(appGroupIdentity *string, request *ModifyAppGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CurrentVersion)) {
		body["currentVersion"] = request.CurrentVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Debugging
 * [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=OpenSearch\\&api=ModifyAppGroup\\&type=ROA\\&version=2017-12-25)
 *
 * @param request ModifyAppGroupRequest
 * @return ModifyAppGroupResponse
 */
func (client *Client) ModifyAppGroup(appGroupIdentity *string, request *ModifyAppGroupRequest) (_result *ModifyAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAppGroupResponse{}
	_body, _err := client.ModifyAppGroupWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppGroupQuotaWithOptions(appGroupIdentity *string, request *ModifyAppGroupQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAppGroupQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyAppGroupQuota"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/quota"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyAppGroupQuotaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAppGroupQuota(appGroupIdentity *string, request *ModifyAppGroupQuotaRequest) (_result *ModifyAppGroupQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAppGroupQuotaResponse{}
	_body, _err := client.ModifyAppGroupQuotaWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifyFirstRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyFirstRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/first-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyFirstRank(appGroupIdentity *string, appId *string, name *string, request *ModifyFirstRankRequest) (_result *ModifyFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyFirstRankResponse{}
	_body, _err := client.ModifyFirstRankWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifyQueryProcessorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyQueryProcessorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyQueryProcessor(appGroupIdentity *string, appId *string, name *string, request *ModifyQueryProcessorRequest) (_result *ModifyQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyQueryProcessorResponse{}
	_body, _err := client.ModifyQueryProcessorWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, request *ModifyScheduledTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyScheduledTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scheduled-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyScheduledTask(appGroupIdentity *string, taskId *string, request *ModifyScheduledTaskRequest) (_result *ModifyScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyScheduledTaskResponse{}
	_body, _err := client.ModifyScheduledTaskWithOptions(appGroupIdentity, taskId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, request *ModifySecondRankRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifySecondRankResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/second-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecondRank(appGroupIdentity *string, appId *string, name *string, request *ModifySecondRankRequest) (_result *ModifySecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifySecondRankResponse{}
	_body, _err := client.ModifySecondRankWithOptions(appGroupIdentity, appId, name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PreviewModelWithOptions(appGroupIdentity *string, modelName *string, request *PreviewModelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PreviewModelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Query)) {
		query["query"] = request.Query
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PreviewModel"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/algorithm/models/" + tea.StringValue(openapiutil.GetEncodeParam(modelName)) + "/actions/preview"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PreviewModelResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PreviewModel(appGroupIdentity *string, modelName *string, request *PreviewModelRequest) (_result *PreviewModelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PreviewModelResponse{}
	_body, _err := client.PreviewModelWithOptions(appGroupIdentity, modelName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushInterventionDictionaryEntriesWithOptions(name *string, request *PushInterventionDictionaryEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInterventionDictionaryEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushInterventionDictionaryEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/entries/actions/bulk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushInterventionDictionaryEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushInterventionDictionaryEntries(name *string, request *PushInterventionDictionaryEntriesRequest) (_result *PushInterventionDictionaryEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInterventionDictionaryEntriesResponse{}
	_body, _err := client.PushInterventionDictionaryEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PushUserAnalyzerEntriesWithOptions(name *string, request *PushUserAnalyzerEntriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushUserAnalyzerEntriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Entries)) {
		body["entries"] = request.Entries
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("PushUserAnalyzerEntries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(openapiutil.GetEncodeParam(name)) + "/entries/actions/bulk"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PushUserAnalyzerEntriesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PushUserAnalyzerEntries(name *string, request *PushUserAnalyzerEntriesRequest) (_result *PushUserAnalyzerEntriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushUserAnalyzerEntriesResponse{}
	_body, _err := client.PushUserAnalyzerEntriesWithOptions(name, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RankPreviewQueryWithOptions(appGroupIdentity *string, modelName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RankPreviewQueryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RankPreviewQuery"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/algorithm/models/" + tea.StringValue(openapiutil.GetEncodeParam(modelName)) + "/actions/query-rank"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RankPreviewQueryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RankPreviewQuery(appGroupIdentity *string, modelName *string) (_result *RankPreviewQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RankPreviewQueryResponse{}
	_body, _err := client.RankPreviewQueryWithOptions(appGroupIdentity, modelName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseSortScriptWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/actions/release"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseSortScript(appGroupIdentity *string, scriptName *string, appVersionId *string) (_result *ReleaseSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseSortScriptResponse{}
	_body, _err := client.ReleaseSortScriptWithOptions(appGroupIdentity, scriptName, appVersionId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * > If an application has two versions, you can delete only the offline version.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveAppResponse
 */
func (client *Client) RemoveAppWithOptions(appGroupIdentity *string, appId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveApp"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * > If an application has two versions, you can delete only the offline version.
 *
 * @return RemoveAppResponse
 */
func (client *Client) RemoveApp(appGroupIdentity *string, appId *string) (_result *RemoveAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveAppResponse{}
	_body, _err := client.RemoveAppWithOptions(appGroupIdentity, appId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * ## Debugging
 * [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=OpenSearch\\&api=RemoveAppGroup\\&type=ROA\\&version=2017-12-25)
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return RemoveAppGroupResponse
 */
func (client *Client) RemoveAppGroupWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveAppGroupResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * ## Debugging
 * [OpenAPI Explorer automatically calculates the signature value. For your convenience, we recommend that you call this operation in OpenAPI Explorer. OpenAPI Explorer dynamically generates the sample code of the operation for different SDKs.](https://api.aliyun.com/#product=OpenSearch\\&api=RemoveAppGroup\\&type=ROA\\&version=2017-12-25)
 *
 * @return RemoveAppGroupResponse
 */
func (client *Client) RemoveAppGroup(appGroupIdentity *string) (_result *RemoveAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveAppGroupResponse{}
	_body, _err := client.RemoveAppGroupWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveDataCollectionWithOptions(appGroupIdentity *string, dataCollectionIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveDataCollectionResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveDataCollection"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/data-collections/" + tea.StringValue(openapiutil.GetEncodeParam(dataCollectionIdentity))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveDataCollectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveDataCollection(appGroupIdentity *string, dataCollectionIdentity *string) (_result *RemoveDataCollectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveDataCollectionResponse{}
	_body, _err := client.RemoveDataCollectionWithOptions(appGroupIdentity, dataCollectionIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveFirstRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveFirstRankResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFirstRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/first-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFirstRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveFirstRank(appGroupIdentity *string, appId *string, name *string) (_result *RemoveFirstRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveFirstRankResponse{}
	_body, _err := client.RemoveFirstRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveInterventionDictionaryWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveInterventionDictionaryResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveInterventionDictionary"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/intervention-dictionaries/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveInterventionDictionaryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveInterventionDictionary(name *string) (_result *RemoveInterventionDictionaryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveInterventionDictionaryResponse{}
	_body, _err := client.RemoveInterventionDictionaryWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveQueryProcessorWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveQueryProcessorResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveQueryProcessor"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/query-processors/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveQueryProcessorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveQueryProcessor(appGroupIdentity *string, appId *string, name *string) (_result *RemoveQueryProcessorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveQueryProcessorResponse{}
	_body, _err := client.RemoveQueryProcessorWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveScheduledTaskWithOptions(appGroupIdentity *string, taskId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveScheduledTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveScheduledTask"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scheduled-tasks/" + tea.StringValue(openapiutil.GetEncodeParam(taskId))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveScheduledTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveScheduledTask(appGroupIdentity *string, taskId *string) (_result *RemoveScheduledTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveScheduledTaskResponse{}
	_body, _err := client.RemoveScheduledTaskWithOptions(appGroupIdentity, taskId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveSearchStrategyResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/search-strategies/" + tea.StringValue(openapiutil.GetEncodeParam(strategyName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string) (_result *RemoveSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveSearchStrategyResponse{}
	_body, _err := client.RemoveSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveSecondRankWithOptions(appGroupIdentity *string, appId *string, name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveSecondRankResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveSecondRank"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/second-ranks/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveSecondRankResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveSecondRank(appGroupIdentity *string, appId *string, name *string) (_result *RemoveSecondRankResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveSecondRankResponse{}
	_body, _err := client.RemoveSecondRankWithOptions(appGroupIdentity, appId, name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveUserAnalyzerWithOptions(name *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveUserAnalyzerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/user-analyzers/" + tea.StringValue(openapiutil.GetEncodeParam(name))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveUserAnalyzer(name *string) (_result *RemoveUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveUserAnalyzerResponse{}
	_body, _err := client.RemoveUserAnalyzerWithOptions(name, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewAppGroupWithOptions(appGroupIdentity *string, request *RenewAppGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["clientToken"] = request.ClientToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewAppGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/actions/renew"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewAppGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewAppGroup(appGroupIdentity *string, request *RenewAppGroupRequest) (_result *RenewAppGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewAppGroupResponse{}
	_body, _err := client.RenewAppGroupWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReplaceAppGroupCommodityCodeWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReplaceAppGroupCommodityCodeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReplaceAppGroupCommodityCode"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/actions/to-instance-typed"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReplaceAppGroupCommodityCodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReplaceAppGroupCommodityCode(appGroupIdentity *string) (_result *ReplaceAppGroupCommodityCodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReplaceAppGroupCommodityCodeResponse{}
	_body, _err := client.ReplaceAppGroupCommodityCodeWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SaveSortScriptFileWithOptions(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, request *SaveSortScriptFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveSortScriptFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SaveSortScriptFile"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName)) + "/files/src/" + tea.StringValue(openapiutil.GetEncodeParam(fileName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SaveSortScriptFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SaveSortScriptFile(appGroupIdentity *string, scriptName *string, appVersionId *string, fileName *string, request *SaveSortScriptFileRequest) (_result *SaveSortScriptFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveSortScriptFileResponse{}
	_body, _err := client.SaveSortScriptFileWithOptions(appGroupIdentity, scriptName, appVersionId, fileName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSlowQueryAnalyzerWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartSlowQueryAnalyzerResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartSlowQueryAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/optimizers/slow-query/actions/run"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSlowQueryAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSlowQueryAnalyzer(appGroupIdentity *string) (_result *StartSlowQueryAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartSlowQueryAnalyzerResponse{}
	_body, _err := client.StartSlowQueryAnalyzerWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["resourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/resource-tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request UnbindESUserAnalyzerRequest
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UnbindESUserAnalyzerResponse
 */
func (client *Client) UnbindESUserAnalyzerWithOptions(appGroupIdentity *string, esInstanceId *string, request *UnbindESUserAnalyzerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindESUserAnalyzerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindESUserAnalyzer"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/es/" + tea.StringValue(openapiutil.GetEncodeParam(esInstanceId)) + "/actions/unbind-analyzer"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindESUserAnalyzerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * The ID of the request.
 *
 * @param request UnbindESUserAnalyzerRequest
 * @return UnbindESUserAnalyzerResponse
 */
func (client *Client) UnbindESUserAnalyzer(appGroupIdentity *string, esInstanceId *string, request *UnbindESUserAnalyzerRequest) (_result *UnbindESUserAnalyzerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindESUserAnalyzerResponse{}
	_body, _err := client.UnbindESUserAnalyzerWithOptions(appGroupIdentity, esInstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UnbindEsInstanceWithOptions(appGroupIdentity *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnbindEsInstanceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UnbindEsInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/actions/unbind-es-instance"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnbindEsInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UnbindEsInstance(appGroupIdentity *string) (_result *UnbindEsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnbindEsInstanceResponse{}
	_body, _err := client.UnbindEsInstanceWithOptions(appGroupIdentity, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(tmpReq *UntagResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &UntagResourcesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.ResourceId)) {
		request.ResourceIdShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.ResourceId, tea.String("resourceId"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TagKey)) {
		request.TagKeyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TagKey, tea.String("tagKey"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["all"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceIdShrink)) {
		query["resourceId"] = request.ResourceIdShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKeyShrink)) {
		query["tagKey"] = request.TagKeyShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/resource-tags"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestExperimentWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestExperimentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestExperimentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestExperiment"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(experimentId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestExperimentResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestExperiment(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestExperimentRequest) (_result *UpdateABTestExperimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestExperimentResponse{}
	_body, _err := client.UpdateABTestExperimentWithOptions(appGroupIdentity, sceneId, groupId, experimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestFixedFlowDividersWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestFixedFlowDividersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestFixedFlowDividersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestFixedFlowDividers"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId)) + "/experiments/" + tea.StringValue(openapiutil.GetEncodeParam(experimentId)) + "/fixed-flow-dividers"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestFixedFlowDividersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestFixedFlowDividers(appGroupIdentity *string, sceneId *string, groupId *string, experimentId *string, request *UpdateABTestFixedFlowDividersRequest) (_result *UpdateABTestFixedFlowDividersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestFixedFlowDividersResponse{}
	_body, _err := client.UpdateABTestFixedFlowDividersWithOptions(appGroupIdentity, sceneId, groupId, experimentId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestGroupWithOptions(appGroupIdentity *string, sceneId *string, groupId *string, request *UpdateABTestGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestGroup"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId)) + "/groups/" + tea.StringValue(openapiutil.GetEncodeParam(groupId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestGroup(appGroupIdentity *string, sceneId *string, groupId *string, request *UpdateABTestGroupRequest) (_result *UpdateABTestGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestGroupResponse{}
	_body, _err := client.UpdateABTestGroupWithOptions(appGroupIdentity, sceneId, groupId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateABTestSceneWithOptions(appGroupIdentity *string, sceneId *string, request *UpdateABTestSceneRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateABTestSceneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateABTestScene"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/scenes/" + tea.StringValue(openapiutil.GetEncodeParam(sceneId))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateABTestSceneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateABTestScene(appGroupIdentity *string, sceneId *string, request *UpdateABTestSceneRequest) (_result *UpdateABTestSceneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateABTestSceneResponse{}
	_body, _err := client.UpdateABTestSceneWithOptions(appGroupIdentity, sceneId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFetchFieldsWithOptions(appGroupIdentity *string, appId *string, request *UpdateFetchFieldsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFetchFieldsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFetchFields"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/fetch-fields"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFetchFieldsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFetchFields(appGroupIdentity *string, appId *string, request *UpdateFetchFieldsRequest) (_result *UpdateFetchFieldsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFetchFieldsResponse{}
	_body, _err := client.UpdateFetchFieldsWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionDefaultInstanceWithOptions(appGroupIdentity *string, functionName *string, request *UpdateFunctionDefaultInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionDefaultInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["instanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunctionDefaultInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/default-instance"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionDefaultInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunctionDefaultInstance(appGroupIdentity *string, functionName *string, request *UpdateFunctionDefaultInstanceRequest) (_result *UpdateFunctionDefaultInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionDefaultInstanceResponse{}
	_body, _err := client.UpdateFunctionDefaultInstanceWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateFunctionInstanceWithOptions(appGroupIdentity *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CreateParameters)) {
		body["createParameters"] = request.CreateParameters
	}

	if !tea.BoolValue(util.IsUnset(request.Cron)) {
		body["cron"] = request.Cron
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.UsageParameters)) {
		body["usageParameters"] = request.UsageParameters
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunctionInstance"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/instances/" + tea.StringValue(openapiutil.GetEncodeParam(instanceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateFunctionInstance(appGroupIdentity *string, functionName *string, instanceName *string, request *UpdateFunctionInstanceRequest) (_result *UpdateFunctionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionInstanceResponse{}
	_body, _err := client.UpdateFunctionInstanceWithOptions(appGroupIdentity, functionName, instanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSearchStrategyWithOptions(appGroupIdentity *string, appId *string, strategyName *string, request *UpdateSearchStrategyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSearchStrategyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSearchStrategy"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/search-strategies/" + tea.StringValue(openapiutil.GetEncodeParam(strategyName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSearchStrategyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSearchStrategy(appGroupIdentity *string, appId *string, strategyName *string, request *UpdateSearchStrategyRequest) (_result *UpdateSearchStrategyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSearchStrategyResponse{}
	_body, _err := client.UpdateSearchStrategyWithOptions(appGroupIdentity, appId, strategyName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * You can call this operation to modify the description of a sort script.
 *
 * @param headers map
 * @param runtime runtime options for this request RuntimeOptions
 * @return UpdateSortScriptResponse
 */
func (client *Client) UpdateSortScriptWithOptions(appGroupIdentity *string, appVersionId *string, scriptName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSortScriptResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSortScript"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appVersionId)) + "/sort-scripts/" + tea.StringValue(openapiutil.GetEncodeParam(scriptName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSortScriptResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * You can call this operation to modify the description of a sort script.
 *
 * @return UpdateSortScriptResponse
 */
func (client *Client) UpdateSortScript(appGroupIdentity *string, appVersionId *string, scriptName *string) (_result *UpdateSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSortScriptResponse{}
	_body, _err := client.UpdateSortScriptWithOptions(appGroupIdentity, appVersionId, scriptName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateSummariesWithOptions(appGroupIdentity *string, appId *string, request *UpdateSummariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateSummariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    util.ToArray(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateSummaries"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/apps/" + tea.StringValue(openapiutil.GetEncodeParam(appId)) + "/summaries"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateSummariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateSummaries(appGroupIdentity *string, appId *string, request *UpdateSummariesRequest) (_result *UpdateSummariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateSummariesResponse{}
	_body, _err := client.UpdateSummariesWithOptions(appGroupIdentity, appId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ValidateDataSourcesWithOptions(request *ValidateDataSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ValidateDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(request.Body),
	}
	params := &openapi.Params{
		Action:      tea.String("ValidateDataSources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/assist/data-sources/validations"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ValidateDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ValidateDataSources(request *ValidateDataSourcesRequest) (_result *ValidateDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ValidateDataSourcesResponse{}
	_body, _err := client.ValidateDataSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
