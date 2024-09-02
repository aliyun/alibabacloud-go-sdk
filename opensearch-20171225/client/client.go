// This file is auto-generated, don't edit it. Thanks.
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
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0 停止实验 1 开通实验
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
	ChargeType      *string        `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	Description     *string        `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *string        `json:"domain,omitempty" xml:"domain,omitempty"`
	Name            *string        `json:"name,omitempty" xml:"name,omitempty"`
	Order           *AppGroupOrder `json:"order,omitempty" xml:"order,omitempty" type:"Struct"`
	Quota           *Quota         `json:"quota,omitempty" xml:"quota,omitempty"`
	ResourceGroupId *string        `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Type            *string        `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *AppGroup) SetOrder(v *AppGroupOrder) *AppGroup {
	s.Order = v
	return s
}

func (s *AppGroup) SetQuota(v *Quota) *AppGroup {
	s.Quota = v
	return s
}

func (s *AppGroup) SetResourceGroupId(v string) *AppGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *AppGroup) SetType(v string) *AppGroup {
	s.Type = &v
	return s
}

type AppGroupOrder struct {
	// example:
	//
	// false
	AutoRenew *bool `json:"autoRenew,omitempty" xml:"autoRenew,omitempty"`
	// example:
	//
	// 1
	Duration *int64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// Month
	PricingCycle *string `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
}

func (s AppGroupOrder) String() string {
	return tea.Prettify(s)
}

func (s AppGroupOrder) GoString() string {
	return s.String()
}

func (s *AppGroupOrder) SetAutoRenew(v bool) *AppGroupOrder {
	s.AutoRenew = &v
	return s
}

func (s *AppGroupOrder) SetDuration(v int64) *AppGroupOrder {
	s.Duration = &v
	return s
}

func (s *AppGroupOrder) SetPricingCycle(v string) *AppGroupOrder {
	s.PricingCycle = &v
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
	// example:
	//
	// keyword: 关键字查询 vector: 向量查询
	QueryType      *string `json:"queryType,omitempty" xml:"queryType,omitempty"`
	SecondRankName *string `json:"secondRankName,omitempty" xml:"secondRankName,omitempty"`
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
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//   "name": "kevintest-analyzer"
	//
	// }
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
	// The ID of the request.
	//
	// example:
	//
	// 3AD34CAD-9603-5251-AFF5-3916C848A1D3
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The body of the request.
	//
	// example:
	//
	// {
	//
	//   "esInstanceId": "es-cn-abcde"
	//
	// }
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
	// The ID of the request.
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompileSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The experiment details.
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
	// The time when the experiment was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The experiment ID.
	//
	// example:
	//
	// 12889
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The experiment alias.
	//
	// example:
	//
	// test3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the experiment is in effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The experiment parameters.
	//
	// example:
	//
	// {"firstFormulaName": "default"}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the experiment.
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the experiment was last modified.
	//
	// example:
	//
	// 1589017861
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body. For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
	Body *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to check the validity of input parameters. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**: checks only the validity of input parameters.
	//
	// 	- **false**: checks the validity of input parameters and creates an attribution configuration.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
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
	//
	// example:
	//
	// 1588839490
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 13466
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The alias of the test group.
	//
	// example:
	//
	// Group_2020-5-7_15:23:3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group.
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last updated.
	//
	// example:
	//
	// 1588839490
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 20405
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// kevintest_2020-5-7_15:21:48
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last modified.
	//
	// example:
	//
	// 1589012351
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	AutoSwitch      *bool                              `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster         *CreateAppRequestCluster           `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	DataSources     []*CreateAppRequestDataSources     `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	Description     *string                            `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *CreateAppRequestDomain            `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	FetchFields     []*string                          `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks      []*CreateAppRequestFirstRanks      `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	NetworkType     *string                            `json:"networkType,omitempty" xml:"networkType,omitempty"`
	QueryProcessors []*CreateAppRequestQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	Schema          *CreateAppRequestSchema            `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	Schemas         []*CreateAppRequestSchemas         `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks     []*CreateAppRequestSecondRanks     `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	Summaries       []*CreateAppRequestSummaries       `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequest) GoString() string {
	return s.String()
}

func (s *CreateAppRequest) SetAutoSwitch(v bool) *CreateAppRequest {
	s.AutoSwitch = &v
	return s
}

func (s *CreateAppRequest) SetCluster(v *CreateAppRequestCluster) *CreateAppRequest {
	s.Cluster = v
	return s
}

func (s *CreateAppRequest) SetDataSources(v []*CreateAppRequestDataSources) *CreateAppRequest {
	s.DataSources = v
	return s
}

func (s *CreateAppRequest) SetDescription(v string) *CreateAppRequest {
	s.Description = &v
	return s
}

func (s *CreateAppRequest) SetDomain(v *CreateAppRequestDomain) *CreateAppRequest {
	s.Domain = v
	return s
}

func (s *CreateAppRequest) SetFetchFields(v []*string) *CreateAppRequest {
	s.FetchFields = v
	return s
}

func (s *CreateAppRequest) SetFirstRanks(v []*CreateAppRequestFirstRanks) *CreateAppRequest {
	s.FirstRanks = v
	return s
}

func (s *CreateAppRequest) SetNetworkType(v string) *CreateAppRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateAppRequest) SetQueryProcessors(v []*CreateAppRequestQueryProcessors) *CreateAppRequest {
	s.QueryProcessors = v
	return s
}

func (s *CreateAppRequest) SetSchema(v *CreateAppRequestSchema) *CreateAppRequest {
	s.Schema = v
	return s
}

func (s *CreateAppRequest) SetSchemas(v []*CreateAppRequestSchemas) *CreateAppRequest {
	s.Schemas = v
	return s
}

func (s *CreateAppRequest) SetSecondRanks(v []*CreateAppRequestSecondRanks) *CreateAppRequest {
	s.SecondRanks = v
	return s
}

func (s *CreateAppRequest) SetSummaries(v []*CreateAppRequestSummaries) *CreateAppRequest {
	s.Summaries = v
	return s
}

func (s *CreateAppRequest) SetDryRun(v bool) *CreateAppRequest {
	s.DryRun = &v
	return s
}

type CreateAppRequestCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s CreateAppRequestCluster) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestCluster) GoString() string {
	return s.String()
}

func (s *CreateAppRequestCluster) SetMaxQueryClauseLength(v int32) *CreateAppRequestCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *CreateAppRequestCluster) SetMaxTimeoutMS(v int32) *CreateAppRequestCluster {
	s.MaxTimeoutMS = &v
	return s
}

type CreateAppRequestDataSources struct {
	Fields     []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                  `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}   `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]interface{}   `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppRequestDataSources) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDataSources) SetFields(v []map[string]interface{}) *CreateAppRequestDataSources {
	s.Fields = v
	return s
}

func (s *CreateAppRequestDataSources) SetKeyField(v string) *CreateAppRequestDataSources {
	s.KeyField = &v
	return s
}

func (s *CreateAppRequestDataSources) SetParameters(v map[string]interface{}) *CreateAppRequestDataSources {
	s.Parameters = v
	return s
}

func (s *CreateAppRequestDataSources) SetPlugins(v map[string]interface{}) *CreateAppRequestDataSources {
	s.Plugins = v
	return s
}

func (s *CreateAppRequestDataSources) SetSchemaName(v string) *CreateAppRequestDataSources {
	s.SchemaName = &v
	return s
}

func (s *CreateAppRequestDataSources) SetTableName(v string) *CreateAppRequestDataSources {
	s.TableName = &v
	return s
}

func (s *CreateAppRequestDataSources) SetType(v string) *CreateAppRequestDataSources {
	s.Type = &v
	return s
}

type CreateAppRequestDomain struct {
	Category  *string                `json:"category,omitempty" xml:"category,omitempty"`
	Functions map[string]interface{} `json:"functions,omitempty" xml:"functions,omitempty"`
	Name      *string                `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestDomain) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestDomain) GoString() string {
	return s.String()
}

func (s *CreateAppRequestDomain) SetCategory(v string) *CreateAppRequestDomain {
	s.Category = &v
	return s
}

func (s *CreateAppRequestDomain) SetFunctions(v map[string]interface{}) *CreateAppRequestDomain {
	s.Functions = v
	return s
}

func (s *CreateAppRequestDomain) SetName(v string) *CreateAppRequestDomain {
	s.Name = &v
	return s
}

type CreateAppRequestFirstRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppRequestFirstRanks) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestFirstRanks) GoString() string {
	return s.String()
}

func (s *CreateAppRequestFirstRanks) SetActive(v bool) *CreateAppRequestFirstRanks {
	s.Active = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetDescription(v string) *CreateAppRequestFirstRanks {
	s.Description = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetMeta(v interface{}) *CreateAppRequestFirstRanks {
	s.Meta = v
	return s
}

func (s *CreateAppRequestFirstRanks) SetName(v string) *CreateAppRequestFirstRanks {
	s.Name = &v
	return s
}

func (s *CreateAppRequestFirstRanks) SetType(v string) *CreateAppRequestFirstRanks {
	s.Type = &v
	return s
}

type CreateAppRequestQueryProcessors struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateAppRequestQueryProcessors) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestQueryProcessors) GoString() string {
	return s.String()
}

func (s *CreateAppRequestQueryProcessors) SetActive(v bool) *CreateAppRequestQueryProcessors {
	s.Active = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetCategory(v string) *CreateAppRequestQueryProcessors {
	s.Category = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetDomain(v string) *CreateAppRequestQueryProcessors {
	s.Domain = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetIndexes(v []*string) *CreateAppRequestQueryProcessors {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetName(v string) *CreateAppRequestQueryProcessors {
	s.Name = &v
	return s
}

func (s *CreateAppRequestQueryProcessors) SetProcessors(v []map[string]interface{}) *CreateAppRequestQueryProcessors {
	s.Processors = v
	return s
}

type CreateAppRequestSchema struct {
	IndexSortConfig  []*CreateAppRequestSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *CreateAppRequestSchemaIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                  `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                  `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                  `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                   `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *CreateAppRequestSchemaTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppRequestSchema) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchema) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchema) SetIndexSortConfig(v []*CreateAppRequestSchemaIndexSortConfig) *CreateAppRequestSchema {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppRequestSchema) SetIndexes(v *CreateAppRequestSchemaIndexes) *CreateAppRequestSchema {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestSchema) SetName(v string) *CreateAppRequestSchema {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchema) SetRouteField(v string) *CreateAppRequestSchema {
	s.RouteField = &v
	return s
}

func (s *CreateAppRequestSchema) SetRouteFieldValues(v []*string) *CreateAppRequestSchema {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppRequestSchema) SetSecondRouteField(v string) *CreateAppRequestSchema {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppRequestSchema) SetTables(v map[string]interface{}) *CreateAppRequestSchema {
	s.Tables = v
	return s
}

func (s *CreateAppRequestSchema) SetTtlField(v *CreateAppRequestSchemaTtlField) *CreateAppRequestSchema {
	s.TtlField = v
	return s
}

type CreateAppRequestSchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppRequestSchemaIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaIndexSortConfig) SetDirection(v string) *CreateAppRequestSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppRequestSchemaIndexSortConfig) SetField(v string) *CreateAppRequestSchemaIndexSortConfig {
	s.Field = &v
	return s
}

type CreateAppRequestSchemaIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppRequestSchemaIndexes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemaIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaIndexes) SetFilterFields(v []*string) *CreateAppRequestSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppRequestSchemaIndexes) SetSearchFields(v map[string]interface{}) *CreateAppRequestSchemaIndexes {
	s.SearchFields = v
	return s
}

type CreateAppRequestSchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppRequestSchemaTtlField) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemaTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemaTtlField) SetName(v string) *CreateAppRequestSchemaTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemaTtlField) SetTtl(v int64) *CreateAppRequestSchemaTtlField {
	s.Ttl = &v
	return s
}

type CreateAppRequestSchemas struct {
	IndexSortConfig  []*CreateAppRequestSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *CreateAppRequestSchemasIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                   `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                   `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                 `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                   `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                    `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *CreateAppRequestSchemasTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppRequestSchemas) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemas) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemas) SetIndexSortConfig(v []*CreateAppRequestSchemasIndexSortConfig) *CreateAppRequestSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppRequestSchemas) SetIndexes(v *CreateAppRequestSchemasIndexes) *CreateAppRequestSchemas {
	s.Indexes = v
	return s
}

func (s *CreateAppRequestSchemas) SetName(v string) *CreateAppRequestSchemas {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemas) SetRouteField(v string) *CreateAppRequestSchemas {
	s.RouteField = &v
	return s
}

func (s *CreateAppRequestSchemas) SetRouteFieldValues(v []*string) *CreateAppRequestSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppRequestSchemas) SetSecondRouteField(v string) *CreateAppRequestSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppRequestSchemas) SetTables(v map[string]interface{}) *CreateAppRequestSchemas {
	s.Tables = v
	return s
}

func (s *CreateAppRequestSchemas) SetTtlField(v *CreateAppRequestSchemasTtlField) *CreateAppRequestSchemas {
	s.TtlField = v
	return s
}

type CreateAppRequestSchemasIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppRequestSchemasIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasIndexSortConfig) SetDirection(v string) *CreateAppRequestSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppRequestSchemasIndexSortConfig) SetField(v string) *CreateAppRequestSchemasIndexSortConfig {
	s.Field = &v
	return s
}

type CreateAppRequestSchemasIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppRequestSchemasIndexes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemasIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasIndexes) SetFilterFields(v []*string) *CreateAppRequestSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppRequestSchemasIndexes) SetSearchFields(v map[string]interface{}) *CreateAppRequestSchemasIndexes {
	s.SearchFields = v
	return s
}

type CreateAppRequestSchemasTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppRequestSchemasTtlField) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSchemasTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSchemasTtlField) SetName(v string) *CreateAppRequestSchemasTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppRequestSchemasTtlField) SetTtl(v int64) *CreateAppRequestSchemasTtlField {
	s.Ttl = &v
	return s
}

type CreateAppRequestSecondRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestSecondRanks) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSecondRanks) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSecondRanks) SetActive(v bool) *CreateAppRequestSecondRanks {
	s.Active = &v
	return s
}

func (s *CreateAppRequestSecondRanks) SetDescription(v string) *CreateAppRequestSecondRanks {
	s.Description = &v
	return s
}

func (s *CreateAppRequestSecondRanks) SetMeta(v interface{}) *CreateAppRequestSecondRanks {
	s.Meta = v
	return s
}

func (s *CreateAppRequestSecondRanks) SetName(v string) *CreateAppRequestSecondRanks {
	s.Name = &v
	return s
}

type CreateAppRequestSummaries struct {
	Meta []*CreateAppRequestSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name *string                          `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppRequestSummaries) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSummaries) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSummaries) SetMeta(v []*CreateAppRequestSummariesMeta) *CreateAppRequestSummaries {
	s.Meta = v
	return s
}

func (s *CreateAppRequestSummaries) SetName(v string) *CreateAppRequestSummaries {
	s.Name = &v
	return s
}

type CreateAppRequestSummariesMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s CreateAppRequestSummariesMeta) String() string {
	return tea.Prettify(s)
}

func (s CreateAppRequestSummariesMeta) GoString() string {
	return s.String()
}

func (s *CreateAppRequestSummariesMeta) SetElement(v string) *CreateAppRequestSummariesMeta {
	s.Element = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetEllipsis(v string) *CreateAppRequestSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetField(v string) *CreateAppRequestSummariesMeta {
	s.Field = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetLen(v int32) *CreateAppRequestSummariesMeta {
	s.Len = &v
	return s
}

func (s *CreateAppRequestSummariesMeta) SetSnippet(v string) *CreateAppRequestSummariesMeta {
	s.Snippet = &v
	return s
}

type CreateAppResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFG
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
	Result *CreateAppResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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

func (s *CreateAppResponseBody) SetResult(v *CreateAppResponseBodyResult) *CreateAppResponseBody {
	s.Result = v
	return s
}

type CreateAppResponseBodyResult struct {
	AutoSwitch      *bool                                         `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster         *CreateAppResponseBodyResultCluster           `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	ClusterName     *string                                       `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	DataSources     []*CreateAppResponseBodyResultDataSources     `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	Description     *string                                       `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *CreateAppResponseBodyResultDomain            `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	FetchFields     []*string                                     `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks      []*CreateAppResponseBodyResultFirstRanks      `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	Id              *string                                       `json:"id,omitempty" xml:"id,omitempty"`
	Interpretations map[string]interface{}                        `json:"interpretations,omitempty" xml:"interpretations,omitempty"`
	IsCurrent       *bool                                         `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	ProgressPercent *int32                                        `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Prompts         []map[string]interface{}                      `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	QueryProcessors []*CreateAppResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	Quota           *CreateAppResponseBodyResultQuota             `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	Schema          *CreateAppResponseBodyResultSchema            `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	Schemas         []*CreateAppResponseBodyResultSchemas         `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks     []*CreateAppResponseBodyResultSecondRanks     `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	Status          *string                                       `json:"status,omitempty" xml:"status,omitempty"`
	Summaries       []*CreateAppResponseBodyResultSummaries       `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	Type            *string                                       `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResult) SetAutoSwitch(v bool) *CreateAppResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetCluster(v *CreateAppResponseBodyResultCluster) *CreateAppResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *CreateAppResponseBodyResult) SetClusterName(v string) *CreateAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetDataSources(v []*CreateAppResponseBodyResultDataSources) *CreateAppResponseBodyResult {
	s.DataSources = v
	return s
}

func (s *CreateAppResponseBodyResult) SetDescription(v string) *CreateAppResponseBodyResult {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetDomain(v *CreateAppResponseBodyResultDomain) *CreateAppResponseBodyResult {
	s.Domain = v
	return s
}

func (s *CreateAppResponseBodyResult) SetFetchFields(v []*string) *CreateAppResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *CreateAppResponseBodyResult) SetFirstRanks(v []*CreateAppResponseBodyResultFirstRanks) *CreateAppResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *CreateAppResponseBodyResult) SetId(v string) *CreateAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetInterpretations(v map[string]interface{}) *CreateAppResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *CreateAppResponseBodyResult) SetIsCurrent(v bool) *CreateAppResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetProgressPercent(v int32) *CreateAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetPrompts(v []map[string]interface{}) *CreateAppResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *CreateAppResponseBodyResult) SetQueryProcessors(v []*CreateAppResponseBodyResultQueryProcessors) *CreateAppResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *CreateAppResponseBodyResult) SetQuota(v *CreateAppResponseBodyResultQuota) *CreateAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSchema(v *CreateAppResponseBodyResultSchema) *CreateAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSchemas(v []*CreateAppResponseBodyResultSchemas) *CreateAppResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *CreateAppResponseBodyResult) SetSecondRanks(v []*CreateAppResponseBodyResultSecondRanks) *CreateAppResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *CreateAppResponseBodyResult) SetStatus(v string) *CreateAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *CreateAppResponseBodyResult) SetSummaries(v []*CreateAppResponseBodyResultSummaries) *CreateAppResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *CreateAppResponseBodyResult) SetType(v string) *CreateAppResponseBodyResult {
	s.Type = &v
	return s
}

type CreateAppResponseBodyResultCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s CreateAppResponseBodyResultCluster) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *CreateAppResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *CreateAppResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *CreateAppResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

type CreateAppResponseBodyResultDataSources struct {
	Fields     []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                  `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}   `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]interface{}   `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppResponseBodyResultDataSources) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetKeyField(v string) *CreateAppResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *CreateAppResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetSchemaName(v string) *CreateAppResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetTableName(v string) *CreateAppResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *CreateAppResponseBodyResultDataSources) SetType(v string) *CreateAppResponseBodyResultDataSources {
	s.Type = &v
	return s
}

type CreateAppResponseBodyResultDomain struct {
	Category  *string                                     `json:"category,omitempty" xml:"category,omitempty"`
	Functions *CreateAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	Name      *string                                     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppResponseBodyResultDomain) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDomain) SetCategory(v string) *CreateAppResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *CreateAppResponseBodyResultDomain) SetFunctions(v *CreateAppResponseBodyResultDomainFunctions) *CreateAppResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *CreateAppResponseBodyResultDomain) SetName(v string) *CreateAppResponseBodyResultDomain {
	s.Name = &v
	return s
}

type CreateAppResponseBodyResultDomainFunctions struct {
	Algo    []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	Qp      []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s CreateAppResponseBodyResultDomainFunctions) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetAlgo(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetQp(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *CreateAppResponseBodyResultDomainFunctions) SetService(v []*string) *CreateAppResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

type CreateAppResponseBodyResultFirstRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppResponseBodyResultFirstRanks) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultFirstRanks) SetActive(v bool) *CreateAppResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetDescription(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetMeta(v interface{}) *CreateAppResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetName(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultFirstRanks) SetType(v string) *CreateAppResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

type CreateAppResponseBodyResultQueryProcessors struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s CreateAppResponseBodyResultQueryProcessors) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetActive(v bool) *CreateAppResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetCategory(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetDomain(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetIndexes(v []*string) *CreateAppResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetName(v string) *CreateAppResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *CreateAppResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

type CreateAppResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Qps             *int32  `json:"qps,omitempty" xml:"qps,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateAppResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultQuota) SetComputeResource(v int32) *CreateAppResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetDocSize(v int32) *CreateAppResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetQps(v int32) *CreateAppResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *CreateAppResponseBodyResultQuota) SetSpec(v string) *CreateAppResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type CreateAppResponseBodyResultSchema struct {
	IndexSortConfig  []*CreateAppResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *CreateAppResponseBodyResultSchemaIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                             `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                             `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                           `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                             `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                              `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *CreateAppResponseBodyResultSchemaTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppResponseBodyResultSchema) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchema) SetIndexSortConfig(v []*CreateAppResponseBodyResultSchemaIndexSortConfig) *CreateAppResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetIndexes(v *CreateAppResponseBodyResultSchemaIndexes) *CreateAppResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetName(v string) *CreateAppResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetRouteField(v string) *CreateAppResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetRouteFieldValues(v []*string) *CreateAppResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetSecondRouteField(v string) *CreateAppResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetTables(v map[string]interface{}) *CreateAppResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *CreateAppResponseBodyResultSchema) SetTtlField(v *CreateAppResponseBodyResultSchemaTtlField) *CreateAppResponseBodyResultSchema {
	s.TtlField = v
	return s
}

type CreateAppResponseBodyResultSchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppResponseBodyResultSchemaIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *CreateAppResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexSortConfig) SetField(v string) *CreateAppResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

type CreateAppResponseBodyResultSchemaIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppResponseBodyResultSchemaIndexes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *CreateAppResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *CreateAppResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

type CreateAppResponseBodyResultSchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppResponseBodyResultSchemaTtlField) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemaTtlField) SetName(v string) *CreateAppResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemaTtlField) SetTtl(v int64) *CreateAppResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

type CreateAppResponseBodyResultSchemas struct {
	IndexSortConfig  []*CreateAppResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *CreateAppResponseBodyResultSchemasIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                              `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                              `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                            `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                              `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                               `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *CreateAppResponseBodyResultSchemasTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s CreateAppResponseBodyResultSchemas) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemas) SetIndexSortConfig(v []*CreateAppResponseBodyResultSchemasIndexSortConfig) *CreateAppResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetIndexes(v *CreateAppResponseBodyResultSchemasIndexes) *CreateAppResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetName(v string) *CreateAppResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetRouteField(v string) *CreateAppResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *CreateAppResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetSecondRouteField(v string) *CreateAppResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetTables(v map[string]interface{}) *CreateAppResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *CreateAppResponseBodyResultSchemas) SetTtlField(v *CreateAppResponseBodyResultSchemasTtlField) *CreateAppResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

type CreateAppResponseBodyResultSchemasIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s CreateAppResponseBodyResultSchemasIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *CreateAppResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexSortConfig) SetField(v string) *CreateAppResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

type CreateAppResponseBodyResultSchemasIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s CreateAppResponseBodyResultSchemasIndexes) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *CreateAppResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *CreateAppResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *CreateAppResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

type CreateAppResponseBodyResultSchemasTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s CreateAppResponseBodyResultSchemasTtlField) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSchemasTtlField) SetName(v string) *CreateAppResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *CreateAppResponseBodyResultSchemasTtlField) SetTtl(v int64) *CreateAppResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

type CreateAppResponseBodyResultSecondRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppResponseBodyResultSecondRanks) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSecondRanks) SetActive(v bool) *CreateAppResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetDescription(v string) *CreateAppResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetMeta(v interface{}) *CreateAppResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultSecondRanks) SetName(v string) *CreateAppResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

type CreateAppResponseBodyResultSummaries struct {
	Meta []*CreateAppResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name *string                                     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateAppResponseBodyResultSummaries) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSummaries) SetMeta(v []*CreateAppResponseBodyResultSummariesMeta) *CreateAppResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *CreateAppResponseBodyResultSummaries) SetName(v string) *CreateAppResponseBodyResultSummaries {
	s.Name = &v
	return s
}

type CreateAppResponseBodyResultSummariesMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s CreateAppResponseBodyResultSummariesMeta) String() string {
	return tea.Prettify(s)
}

func (s CreateAppResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetElement(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetEllipsis(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetField(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetLen(v int32) *CreateAppResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *CreateAppResponseBodyResultSummariesMeta) SetSnippet(v string) *CreateAppResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

type CreateAppResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	ChargeType      *string                     `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	Name            *string                     `json:"name,omitempty" xml:"name,omitempty"`
	Quota           *CreateAppGroupRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId *string                     `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	Type            *string                     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequest) SetChargeType(v string) *CreateAppGroupRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateAppGroupRequest) SetName(v string) *CreateAppGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateAppGroupRequest) SetQuota(v *CreateAppGroupRequestQuota) *CreateAppGroupRequest {
	s.Quota = v
	return s
}

func (s *CreateAppGroupRequest) SetResourceGroupId(v string) *CreateAppGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateAppGroupRequest) SetType(v string) *CreateAppGroupRequest {
	s.Type = &v
	return s
}

type CreateAppGroupRequestQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateAppGroupRequestQuota) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateAppGroupRequestQuota) SetComputeResource(v int32) *CreateAppGroupRequestQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateAppGroupRequestQuota) SetDocSize(v int32) *CreateAppGroupRequestQuota {
	s.DocSize = &v
	return s
}

func (s *CreateAppGroupRequestQuota) SetSpec(v string) *CreateAppGroupRequestQuota {
	s.Spec = &v
	return s
}

type CreateAppGroupResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 766CF6DB-CA02-3E12-7CBA-6AC21FC978FD
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
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// 	- 1: computing resources
	//
	// 	- 2: queries per second (QPS)
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1590139542
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL: general.
	//
	// 	- ECOMMERCE: e-commerce.
	//
	// 	- IT_CONTENT: IT content.
	//
	// example:
	//
	// GENERAL
	Domain     *string `json:"domain,omitempty" xml:"domain,omitempty"`
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// -
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// 	- 0: The quotas are approved.
	//
	// 	- 1: The quotas are being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// -
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// 	- Unlock: The instance is not locked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// 	- 0: The order is in progress.
	//
	// 	- 1: The order is complete.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// -
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *CreateAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The status of the application. Valid values:
	//
	// 	- producing
	//
	// 	- review_pending
	//
	// 	- config_pending
	//
	// 	- normal
	//
	// 	- frozen
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard: a standard application.
	//
	// 	- advance: an advanced application which is of an old application type. New applications cannot be of this type.
	//
	// 	- enhanced: an advanced application which is of a new application type.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1590978265
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

func (s *CreateAppGroupResponseBodyResult) SetEngineType(v string) *CreateAppGroupResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *CreateAppGroupResponseBodyResult) SetExpireOn(v string) *CreateAppGroupResponseBodyResult {
	s.ExpireOn = &v
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

func (s *CreateAppGroupResponseBodyResult) SetName(v string) *CreateAppGroupResponseBodyResult {
	s.Name = &v
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
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// 	- opensearch.share.junior: basic
	//
	// 	- opensearch.share.common: shared general-purpose
	//
	// 	- opensearch.share.compute: shared computing
	//
	// 	- opensearch.share.storage: shared storage
	//
	// 	- opensearch.private.common: exclusive general-purpose
	//
	// 	- opensearch.private.compute: exclusive computing
	//
	// 	- opensearch.private.storage: exclusive storage
	//
	// example:
	//
	// opensearch.share.common
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateAppGroupCredentialsRequest struct {
	// example:
	//
	// api-token
	Type   *string `json:"type,omitempty" xml:"type,omitempty"`
	DryRun *bool   `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateAppGroupCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupCredentialsRequest) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsRequest) SetType(v string) *CreateAppGroupCredentialsRequest {
	s.Type = &v
	return s
}

func (s *CreateAppGroupCredentialsRequest) SetDryRun(v bool) *CreateAppGroupCredentialsRequest {
	s.DryRun = &v
	return s
}

type CreateAppGroupCredentialsResponseBody struct {
	// example:
	//
	// 1-2-3-4
	RequestId *string                                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateAppGroupCredentialsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateAppGroupCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponseBody) SetRequestId(v string) *CreateAppGroupCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBody) SetResult(v *CreateAppGroupCredentialsResponseBodyResult) *CreateAppGroupCredentialsResponseBody {
	s.Result = v
	return s
}

type CreateAppGroupCredentialsResponseBodyResult struct {
	// example:
	//
	// app_group_123
	AppGroupId *int64 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	Enabled    *bool  `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// example:
	//
	// generated_token_string
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// example:
	//
	// api-token
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateAppGroupCredentialsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupCredentialsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetAppGroupId(v int64) *CreateAppGroupCredentialsResponseBodyResult {
	s.AppGroupId = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetEnabled(v bool) *CreateAppGroupCredentialsResponseBodyResult {
	s.Enabled = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetToken(v string) *CreateAppGroupCredentialsResponseBodyResult {
	s.Token = &v
	return s
}

func (s *CreateAppGroupCredentialsResponseBodyResult) SetType(v string) *CreateAppGroupCredentialsResponseBodyResult {
	s.Type = &v
	return s
}

type CreateAppGroupCredentialsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAppGroupCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAppGroupCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAppGroupCredentialsResponse) GoString() string {
	return s.String()
}

func (s *CreateAppGroupCredentialsResponse) SetHeaders(v map[string]*string) *CreateAppGroupCredentialsResponse {
	s.Headers = v
	return s
}

func (s *CreateAppGroupCredentialsResponse) SetStatusCode(v int32) *CreateAppGroupCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAppGroupCredentialsResponse) SetBody(v *CreateAppGroupCredentialsResponseBody) *CreateAppGroupCredentialsResponse {
	s.Body = v
	return s
}

type CreateFirstRankRequest struct {
	// The request body that contains the parameters of the rough sort expression.
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The content of the expression.
	Meta []*CreateFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
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
	//
	// example:
	//
	// 1
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature functions, or field to be searched for.
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: [-100000,100000]. The value cannot be 0.
	//
	// example:
	//
	// 10
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The parameters used to create the instance.
	//
	// example:
	//
	// [   { "name": "param1", "value": "val1"   } ]
	CreateParameters []*CreateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The CRON expression used to schedule periodic training, in the format of Minutes Hours DayofMonth Month DayofWeek. The default value is empty, which specifies that no periodic training is performed. A value of 0 for DayofWeek specifies Sunday.
	//
	// example:
	//
	// 0 0 ? 	- 0,1,2,3,4,5,6
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description.
	//
	// example:
	//
	// test instance
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The feature type.
	//
	// 	- Default value: PAAS. Training is required before you can use the feature.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The instance name. The name must be 1 to 30 characters in length and can contain letters, digits, and underscores (_). The name is case-sensitive and must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// ctr_test
	InstanceName *string `json:"instanceName,omitempty" xml:"instanceName,omitempty"`
	// The model type. The value varies based on the model.
	//
	// 	- Click-through rate (CTR) model: tf_checkpoint
	//
	// 	- Popularity model: pop
	//
	// 	- Category model: offline_inference
	//
	// 	- Hotword model: offline_inference
	//
	// 	- Hint model: offline_inference
	//
	// 	- Hotword model for real-time top searches: near_realtime
	//
	// 	- Personalized hint model: near_realtime
	//
	// 	- Drop-down suggestion model: offline_inference
	//
	// 	- Tokenization model: text
	//
	// 	- Term weight model: tf_checkpoint
	//
	// 	- Synonym model: offline_inference
	//
	// This parameter is required.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The parameters used to use the instance.
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
	// The parameter name.
	//
	// example:
	//
	// title_field
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// title
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
	// The parameter name.
	//
	// example:
	//
	// allow_dict_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// Version.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateFunctionResourceRequest struct {
	// The resource data. The data structure varies with the resource type.
	Data *CreateFunctionResourceRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the resource.
	//
	// example:
	//
	// fg_jsoon
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- feature_generator
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw_file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// feature_generator
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateFunctionResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequest) SetData(v *CreateFunctionResourceRequestData) *CreateFunctionResourceRequest {
	s.Data = v
	return s
}

func (s *CreateFunctionResourceRequest) SetDescription(v string) *CreateFunctionResourceRequest {
	s.Description = &v
	return s
}

func (s *CreateFunctionResourceRequest) SetResourceName(v string) *CreateFunctionResourceRequest {
	s.ResourceName = &v
	return s
}

func (s *CreateFunctionResourceRequest) SetResourceType(v string) *CreateFunctionResourceRequest {
	s.ResourceType = &v
	return s
}

type CreateFunctionResourceRequestData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// "abc"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*CreateFunctionResourceRequestDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s CreateFunctionResourceRequestData) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceRequestData) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestData) SetContent(v string) *CreateFunctionResourceRequestData {
	s.Content = &v
	return s
}

func (s *CreateFunctionResourceRequestData) SetGenerators(v []*CreateFunctionResourceRequestDataGenerators) *CreateFunctionResourceRequestData {
	s.Generators = v
	return s
}

type CreateFunctionResourceRequestDataGenerators struct {
	// The type of the feature generator.
	//
	// Valid values:
	//
	// 	- lookup
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- sequence
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- overlap
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- combo
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- id
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// id
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *CreateFunctionResourceRequestDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// item_id_feature
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s CreateFunctionResourceRequestDataGenerators) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGenerators) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGenerators) SetGenerator(v string) *CreateFunctionResourceRequestDataGenerators {
	s.Generator = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGenerators) SetInput(v *CreateFunctionResourceRequestDataGeneratorsInput) *CreateFunctionResourceRequestDataGenerators {
	s.Input = v
	return s
}

func (s *CreateFunctionResourceRequestDataGenerators) SetOutput(v string) *CreateFunctionResourceRequestDataGenerators {
	s.Output = &v
	return s
}

type CreateFunctionResourceRequestDataGeneratorsInput struct {
	// The input features.
	Features []*CreateFunctionResourceRequestDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s CreateFunctionResourceRequestDataGeneratorsInput) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGeneratorsInput) SetFeatures(v []*CreateFunctionResourceRequestDataGeneratorsInputFeatures) *CreateFunctionResourceRequestDataGeneratorsInput {
	s.Features = v
	return s
}

type CreateFunctionResourceRequestDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// Valid values:
	//
	// 	- item
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateFunctionResourceRequestDataGeneratorsInputFeatures) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceRequestDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) SetName(v string) *CreateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *CreateFunctionResourceRequestDataGeneratorsInputFeatures) SetType(v string) *CreateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

type CreateFunctionResourceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// ""
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. Valid values:
	//
	// 	- OK
	//
	// 	- FAIL
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateFunctionResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceResponseBody) SetCode(v string) *CreateFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetHttpCode(v int64) *CreateFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetLatency(v float64) *CreateFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetMessage(v string) *CreateFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetRequestId(v string) *CreateFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFunctionResourceResponseBody) SetStatus(v string) *CreateFunctionResourceResponseBody {
	s.Status = &v
	return s
}

type CreateFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFunctionResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateFunctionResourceResponse) SetHeaders(v map[string]*string) *CreateFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateFunctionResourceResponse) SetStatusCode(v int32) *CreateFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFunctionResourceResponse) SetBody(v *CreateFunctionResourceResponseBody) *CreateFunctionResourceResponse {
	s.Body = v
	return s
}

type CreateFunctionTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Task.IsRunning
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157990724
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The type of the analyzer. Valid values:
	//
	// 	- MODEL: model-based custom analyzer.
	//
	// 	- SYSTEM: system analyzer.
	//
	// 	- USER: custom analyzer.
	//
	// example:
	//
	// SYSTEM
	AnalyzerType *string `json:"analyzerType,omitempty" xml:"analyzerType,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// ner_dict_ec
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering.
	//
	// 	- synonym: an intervention dictionary for synonym configuration.
	//
	// 	- correction: an intervention dictionary for spelling correction.
	//
	// 	- category_prediction: an intervention dictionary for category prediction.
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER).
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis.
	//
	// 	- suggest_allowlist: a drop-down suggestion whitelist.
	//
	// 	- suggest_denylist: a drop-down suggestion blacklist.
	//
	// 	- hot_allowlist: a top search whitelist.
	//
	// 	- hot_denylist: a top search blacklist.
	//
	// 	- hint_allowlist: a hint whitelist.
	//
	// 	- hint_denylist: a hint blacklist.
	//
	// example:
	//
	// ner
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 80326EFE-485F-46E7-B291-5A1DD08D2198
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
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
	// The custom analyzer.
	//
	// example:
	//
	// dianshang
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the test scenario was created.
	//
	// example:
	//
	// 1591086323
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// testb
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// ner
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1591086323
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	//
	// example:
	//
	// {}
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	//
	// example:
	//
	// 1587398402
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry to which the query analysis rule was applied. Valid values:
	//
	// 	- GENERAL: general.
	//
	// 	- ECOMMERCE: e-commerce.
	//
	// 	- IT_CONTENT: IT content.
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule was applied.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// query_filter
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	//
	// For more information, see [QueryProcessor](https://help.aliyun.com/document_detail/170014.html).
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last modified.
	//
	// example:
	//
	// 1587398402
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// 请求体
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
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The query policy.
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
	// The ID of the request.
	//
	// example:
	//
	// "abc123"
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body. For more information, see [SecondRank](https://help.aliyun.com/document_detail/170008.html).
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the fine sort expression.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type CreateSortScriptRequest struct {
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The script name.
	//
	// example:
	//
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The script type. Set the value to cava_script.
	//
	// example:
	//
	// cava_script
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSortScriptRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptRequest) GoString() string {
	return s.String()
}

func (s *CreateSortScriptRequest) SetScope(v string) *CreateSortScriptRequest {
	s.Scope = &v
	return s
}

func (s *CreateSortScriptRequest) SetScriptName(v string) *CreateSortScriptRequest {
	s.ScriptName = &v
	return s
}

func (s *CreateSortScriptRequest) SetType(v string) *CreateSortScriptRequest {
	s.Type = &v
	return s
}

type CreateSortScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	Result *CreateSortScriptResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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

func (s *CreateSortScriptResponseBody) SetResult(v *CreateSortScriptResponseBodyResult) *CreateSortScriptResponseBody {
	s.Result = v
	return s
}

type CreateSortScriptResponseBodyResult struct {
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The script name.
	//
	// example:
	//
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The script type.
	//
	// example:
	//
	// cava_script
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateSortScriptResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s CreateSortScriptResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateSortScriptResponseBodyResult) SetScope(v string) *CreateSortScriptResponseBodyResult {
	s.Scope = &v
	return s
}

func (s *CreateSortScriptResponseBodyResult) SetScriptName(v string) *CreateSortScriptResponseBodyResult {
	s.ScriptName = &v
	return s
}

func (s *CreateSortScriptResponseBodyResult) SetType(v string) *CreateSortScriptResponseBodyResult {
	s.Type = &v
	return s
}

type CreateSortScriptResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The basic analyzer.
	//
	// example:
	//
	// chn_standard
	Business *string `json:"business,omitempty" xml:"business,omitempty"`
	// The application ID of the custom analyzer.
	//
	// example:
	//
	// 110123123
	BusinessAppGroupId *string `json:"businessAppGroupId,omitempty" xml:"businessAppGroupId,omitempty"`
	// The basic analyzer type. Valid values: AUTO, MODEL, SYSTEM, and USER.
	//
	// example:
	//
	// AUTO
	BusinessType *string `json:"businessType,omitempty" xml:"businessType,omitempty"`
	// The analyzer name.
	//
	// example:
	//
	// jmbon_analyzer
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The engine type. Valid values: HA3 and ES.
	//
	// example:
	//
	// HA3
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result that was returned.
	//
	// example:
	//
	// {}
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// "Instance.NotExist"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 10
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// "instance not exist."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "1081EB05-473C-5BF4-95BE-6D7CAD5E2213"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// 	- OK: The request is successful.
	//
	// 	- FAIL: The request fails.
	//
	// example:
	//
	// "OK"
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DeleteFunctionResourceResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
	//
	// example:
	//
	// ""
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status code. Valid values:
	//
	// 	- OK
	//
	// 	- FAIL
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteFunctionResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResourceResponseBody) SetCode(v string) *DeleteFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetHttpCode(v int64) *DeleteFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetLatency(v float64) *DeleteFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetMessage(v string) *DeleteFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetRequestId(v string) *DeleteFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFunctionResourceResponseBody) SetStatus(v string) *DeleteFunctionResourceResponseBody {
	s.Status = &v
	return s
}

type DeleteFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFunctionResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFunctionResourceResponse) SetHeaders(v map[string]*string) *DeleteFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFunctionResourceResponse) SetStatusCode(v int32) *DeleteFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFunctionResourceResponse) SetBody(v *DeleteFunctionResourceResponseBody) *DeleteFunctionResourceResponse {
	s.Body = v
	return s
}

type DeleteFunctionTaskResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// Task.UnableDelete
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "1081EB05-473C-5BF4-95BE-6D7CAD5E2213"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test.
	//
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// 	- true: in effect
	//
	// 	- false: not in effect
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the test.
	Params *DescribeABTestExperimentResponseBodyResultParams `json:"params,omitempty" xml:"params,omitempty" type:"Struct"`
	// The percentage of traffic that is routed to the test.
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	//
	// example:
	//
	// 1588842080
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
	//
	// example:
	//
	// default
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
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 1588839490
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 13466
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// Group_2020-5-7_15:23:3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	//
	// example:
	//
	// 1588839490
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 1596527691
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test scenario.
	//
	// example:
	//
	// 20614
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test scenario.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// 	- 0: The test is stopped.
	//
	// 	- 1: The test is started.
	//
	// example:
	//
	// 0
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test was last modified.
	//
	// example:
	//
	// 1596527691
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The indicators of the test scenarios.
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 33477D76-C380-1D84-A4AD-043F52876CB1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the version.
	//
	// example:
	//
	// {}
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
	// Indicates whether the version is automatically published to the online environment.
	//
	// example:
	//
	// true
	AutoSwitch *bool                                 `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster    *DescribeAppResponseBodyResultCluster `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	// The name of the cluster.
	//
	// example:
	//
	// -
	ClusterName *string                                     `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	DataSources []*DescribeAppResponseBodyResultDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The description of the version.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *DescribeAppResponseBodyResultDomain `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	// The default display fields.
	//
	// example:
	//
	// []
	FetchFields []*string                                  `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks  []*DescribeAppResponseBodyResultFirstRanks `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	// The ID of the version.
	//
	// example:
	//
	// 100303063
	Id              *string                `json:"id,omitempty" xml:"id,omitempty"`
	Interpretations map[string]interface{} `json:"interpretations,omitempty" xml:"interpretations,omitempty"`
	IsCurrent       *bool                  `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	// The progress of data import, in percentage. For example, a value of 83 indicates 83%.
	//
	// example:
	//
	// 100
	ProgressPercent *int32                                          `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Prompts         []map[string]interface{}                        `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	QueryProcessors []*DescribeAppResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	// The quota information about the version.
	//
	// example:
	//
	// {}
	Quota *DescribeAppResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The application schema.
	//
	// example:
	//
	// {}
	Schema      *DescribeAppResponseBodyResultSchema        `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	Schemas     []*DescribeAppResponseBodyResultSchemas     `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks []*DescribeAppResponseBodyResultSecondRanks `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	// The status of the version. Valid values:
	//
	// 	- ok
	//
	// 	- stopped
	//
	// 	- frozen
	//
	// 	- initializing
	//
	// 	- unavailable
	//
	// 	- data_waiting
	//
	// 	- data_preparing
	//
	// example:
	//
	// ok
	Status    *string                                   `json:"status,omitempty" xml:"status,omitempty"`
	Summaries []*DescribeAppResponseBodyResultSummaries `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard: a standard application.
	//
	// 	- advance: an advanced application which is of an old application type. New applications cannot be of this type.
	//
	// 	- enhanced: an advanced application which is of a new application type.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetCluster(v *DescribeAppResponseBodyResultCluster) *DescribeAppResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetClusterName(v string) *DescribeAppResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetDataSources(v []*DescribeAppResponseBodyResultDataSources) *DescribeAppResponseBodyResult {
	s.DataSources = v
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

func (s *DescribeAppResponseBodyResult) SetFirstRanks(v []*DescribeAppResponseBodyResultFirstRanks) *DescribeAppResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetId(v string) *DescribeAppResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetInterpretations(v map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetIsCurrent(v bool) *DescribeAppResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetProgressPercent(v int32) *DescribeAppResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetPrompts(v []map[string]interface{}) *DescribeAppResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQueryProcessors(v []*DescribeAppResponseBodyResultQueryProcessors) *DescribeAppResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetQuota(v *DescribeAppResponseBodyResultQuota) *DescribeAppResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchema(v *DescribeAppResponseBodyResultSchema) *DescribeAppResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSchemas(v []*DescribeAppResponseBodyResultSchemas) *DescribeAppResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSecondRanks(v []*DescribeAppResponseBodyResultSecondRanks) *DescribeAppResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetStatus(v string) *DescribeAppResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppResponseBodyResult) SetSummaries(v []*DescribeAppResponseBodyResultSummaries) *DescribeAppResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *DescribeAppResponseBodyResult) SetType(v string) *DescribeAppResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeAppResponseBodyResultCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s DescribeAppResponseBodyResultCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *DescribeAppResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *DescribeAppResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *DescribeAppResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

type DescribeAppResponseBodyResultDataSources struct {
	Fields     []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                  `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}   `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]interface{}   `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResultDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetKeyField(v string) *DescribeAppResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *DescribeAppResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetSchemaName(v string) *DescribeAppResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetTableName(v string) *DescribeAppResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *DescribeAppResponseBodyResultDataSources) SetType(v string) *DescribeAppResponseBodyResultDataSources {
	s.Type = &v
	return s
}

type DescribeAppResponseBodyResultDomain struct {
	// The category. By default, this parameter is left empty.
	//
	// example:
	//
	// -
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// search functions
	//
	// example:
	//
	// {}
	Functions *DescribeAppResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	// The name
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// Algorithm structure
	//
	// example:
	//
	// []
	Algo []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	// Queryprocessor description
	//
	// example:
	//
	// []
	Qp []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	// Function description
	//
	// example:
	//
	// []
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

type DescribeAppResponseBodyResultFirstRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppResponseBodyResultFirstRanks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetActive(v bool) *DescribeAppResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetDescription(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetMeta(v interface{}) *DescribeAppResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetName(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultFirstRanks) SetType(v string) *DescribeAppResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

type DescribeAppResponseBodyResultQueryProcessors struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s DescribeAppResponseBodyResultQueryProcessors) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetActive(v bool) *DescribeAppResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetCategory(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetDomain(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetIndexes(v []*string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetName(v string) *DescribeAppResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *DescribeAppResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

type DescribeAppResponseBodyResultQuota struct {
	// The computing resources. Unit: logical computing units (LCUs).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The number of search requests per second. You are charged based on the number of search requests per second in the earlier billing model.
	//
	// example:
	//
	// 5
	Qps *int32 `json:"qps,omitempty" xml:"qps,omitempty"`
	// The specifications of the application. Valid values:
	//
	// 	- opensearch.share.junior: basic
	//
	// 	- opensearch.share.common: shared general-purpose
	//
	// 	- opensearch.share.compute: shared computing
	//
	// 	- opensearch.share.storage: shared storage
	//
	// 	- opensearch.private.common: exclusive general-purpose
	//
	// 	- opensearch.private.compute: exclusive computing
	//
	// 	- opensearch.private.storage: exclusive storage
	//
	// example:
	//
	// opensearch.share.common
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

type DescribeAppResponseBodyResultSchema struct {
	IndexSortConfig  []*DescribeAppResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *DescribeAppResponseBodyResultSchemaIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                               `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                               `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                             `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                               `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                                `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *DescribeAppResponseBodyResultSchemaTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBodyResultSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchema) SetIndexSortConfig(v []*DescribeAppResponseBodyResultSchemaIndexSortConfig) *DescribeAppResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetIndexes(v *DescribeAppResponseBodyResultSchemaIndexes) *DescribeAppResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetName(v string) *DescribeAppResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetRouteField(v string) *DescribeAppResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetRouteFieldValues(v []*string) *DescribeAppResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetSecondRouteField(v string) *DescribeAppResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetTables(v map[string]interface{}) *DescribeAppResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *DescribeAppResponseBodyResultSchema) SetTtlField(v *DescribeAppResponseBodyResultSchemaTtlField) *DescribeAppResponseBodyResultSchema {
	s.TtlField = v
	return s
}

type DescribeAppResponseBodyResultSchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemaIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *DescribeAppResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexSortConfig) SetField(v string) *DescribeAppResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

type DescribeAppResponseBodyResultSchemaIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemaIndexes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *DescribeAppResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

type DescribeAppResponseBodyResultSchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemaTtlField) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) SetName(v string) *DescribeAppResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemaTtlField) SetTtl(v int64) *DescribeAppResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

type DescribeAppResponseBodyResultSchemas struct {
	IndexSortConfig  []*DescribeAppResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *DescribeAppResponseBodyResultSchemasIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                                `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                              `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                                `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                                 `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *DescribeAppResponseBodyResultSchemasTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppResponseBodyResultSchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemas) SetIndexSortConfig(v []*DescribeAppResponseBodyResultSchemasIndexSortConfig) *DescribeAppResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetIndexes(v *DescribeAppResponseBodyResultSchemasIndexes) *DescribeAppResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetName(v string) *DescribeAppResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetRouteField(v string) *DescribeAppResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *DescribeAppResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetSecondRouteField(v string) *DescribeAppResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetTables(v map[string]interface{}) *DescribeAppResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemas) SetTtlField(v *DescribeAppResponseBodyResultSchemasTtlField) *DescribeAppResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

type DescribeAppResponseBodyResultSchemasIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemasIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *DescribeAppResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexSortConfig) SetField(v string) *DescribeAppResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

type DescribeAppResponseBodyResultSchemasIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemasIndexes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *DescribeAppResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

type DescribeAppResponseBodyResultSchemasTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s DescribeAppResponseBodyResultSchemasTtlField) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) SetName(v string) *DescribeAppResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppResponseBodyResultSchemasTtlField) SetTtl(v int64) *DescribeAppResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

type DescribeAppResponseBodyResultSecondRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultSecondRanks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetActive(v bool) *DescribeAppResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetDescription(v string) *DescribeAppResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetMeta(v interface{}) *DescribeAppResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultSecondRanks) SetName(v string) *DescribeAppResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

type DescribeAppResponseBodyResultSummaries struct {
	Meta []*DescribeAppResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name *string                                       `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppResponseBodyResultSummaries) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSummaries) SetMeta(v []*DescribeAppResponseBodyResultSummariesMeta) *DescribeAppResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *DescribeAppResponseBodyResultSummaries) SetName(v string) *DescribeAppResponseBodyResultSummaries {
	s.Name = &v
	return s
}

type DescribeAppResponseBodyResultSummariesMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s DescribeAppResponseBodyResultSummariesMeta) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetElement(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetEllipsis(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetField(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetLen(v int32) *DescribeAppResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *DescribeAppResponseBodyResultSummariesMeta) SetSnippet(v string) *DescribeAppResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

type DescribeAppResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billable item. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: queries per second (QPS).
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1575442875
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 110116134
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry of the application.
	//
	// example:
	//
	// ecommerce
	Domain     *string `json:"domain,omitempty" xml:"domain,omitempty"`
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// -
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the created rough sort expression.
	//
	// example:
	//
	// 0
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval state of the quotas. Valid values:
	//
	// 	- 0: The application is in service.
	//
	// 	- 1: The quotas are being reviewed.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 110116134
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// -
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock state. Valid values:
	//
	// 	- Unlock: The instance is unlocked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	//
	// example:
	//
	// 0
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The application name.
	//
	// example:
	//
	// os_function_test_v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	//
	// example:
	//
	// 0
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is not complete.
	//
	// example:
	//
	// -
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the application is created. Valid values:
	//
	// 	- 0: The application is being created.
	//
	// 	- 1: The application is created.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// -
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota *DescribeAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmoiyerh6nzly
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The ID of the created fine sort expression.
	//
	// example:
	//
	// 0
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The state of the application. Valid values:
	//
	// 	- producing: The application is being created.
	//
	// 	- review_pending: The application is being reviewed.
	//
	// 	- config_pending: The application is to be configured.
	//
	// 	- normal: The application is in service.
	//
	// 	- frozen: The application is frozen.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 0
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The application tags.
	Tags []*DescribeAppGroupResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// *
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1578916076
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

func (s *DescribeAppGroupResponseBodyResult) SetEngineType(v string) *DescribeAppGroupResponseBodyResult {
	s.EngineType = &v
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
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic.
	//
	// 	- opensearch.share.common: shared general-purpose.
	//
	// 	- opensearch.share.compute: shared computing.
	//
	// 	- opensearch.share.storage: shared storage.
	//
	// 	- opensearch.private.common: exclusive general-purpose.
	//
	// 	- opensearch.private.compute: exclusive computing.
	//
	// 	- opensearch.private.storage: exclusive storage.
	//
	// example:
	//
	// opensearch.share.common
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
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value
	//
	// example:
	//
	// bar
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 76FC45F1-4167-D3CD-6737-4F97BA503FA0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The statistics.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// 77CAA411-0010-4DB9-82E2-1C384E590AFF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each version.
	Result []*DescribeAppsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
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

func (s *DescribeAppsResponseBody) SetResult(v []*DescribeAppsResponseBodyResult) *DescribeAppsResponseBody {
	s.Result = v
	return s
}

type DescribeAppsResponseBodyResult struct {
	AutoSwitch      *bool                                            `json:"autoSwitch,omitempty" xml:"autoSwitch,omitempty"`
	Cluster         *DescribeAppsResponseBodyResultCluster           `json:"cluster,omitempty" xml:"cluster,omitempty" type:"Struct"`
	ClusterName     *string                                          `json:"clusterName,omitempty" xml:"clusterName,omitempty"`
	DataSources     []*DescribeAppsResponseBodyResultDataSources     `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	Description     *string                                          `json:"description,omitempty" xml:"description,omitempty"`
	Domain          *DescribeAppsResponseBodyResultDomain            `json:"domain,omitempty" xml:"domain,omitempty" type:"Struct"`
	FetchFields     []*string                                        `json:"fetchFields,omitempty" xml:"fetchFields,omitempty" type:"Repeated"`
	FirstRanks      []*DescribeAppsResponseBodyResultFirstRanks      `json:"firstRanks,omitempty" xml:"firstRanks,omitempty" type:"Repeated"`
	Id              *string                                          `json:"id,omitempty" xml:"id,omitempty"`
	Interpretations map[string]interface{}                           `json:"interpretations,omitempty" xml:"interpretations,omitempty"`
	IsCurrent       *bool                                            `json:"isCurrent,omitempty" xml:"isCurrent,omitempty"`
	ProgressPercent *int32                                           `json:"progressPercent,omitempty" xml:"progressPercent,omitempty"`
	Prompts         []map[string]interface{}                         `json:"prompts,omitempty" xml:"prompts,omitempty" type:"Repeated"`
	QueryProcessors []*DescribeAppsResponseBodyResultQueryProcessors `json:"queryProcessors,omitempty" xml:"queryProcessors,omitempty" type:"Repeated"`
	Quota           *DescribeAppsResponseBodyResultQuota             `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	Schema          *DescribeAppsResponseBodyResultSchema            `json:"schema,omitempty" xml:"schema,omitempty" type:"Struct"`
	Schemas         []*DescribeAppsResponseBodyResultSchemas         `json:"schemas,omitempty" xml:"schemas,omitempty" type:"Repeated"`
	SecondRanks     []*DescribeAppsResponseBodyResultSecondRanks     `json:"secondRanks,omitempty" xml:"secondRanks,omitempty" type:"Repeated"`
	Status          *string                                          `json:"status,omitempty" xml:"status,omitempty"`
	Summaries       []*DescribeAppsResponseBodyResultSummaries       `json:"summaries,omitempty" xml:"summaries,omitempty" type:"Repeated"`
	Type            *string                                          `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppsResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResult) SetAutoSwitch(v bool) *DescribeAppsResponseBodyResult {
	s.AutoSwitch = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetCluster(v *DescribeAppsResponseBodyResultCluster) *DescribeAppsResponseBodyResult {
	s.Cluster = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetClusterName(v string) *DescribeAppsResponseBodyResult {
	s.ClusterName = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDataSources(v []*DescribeAppsResponseBodyResultDataSources) *DescribeAppsResponseBodyResult {
	s.DataSources = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDescription(v string) *DescribeAppsResponseBodyResult {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetDomain(v *DescribeAppsResponseBodyResultDomain) *DescribeAppsResponseBodyResult {
	s.Domain = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetFetchFields(v []*string) *DescribeAppsResponseBodyResult {
	s.FetchFields = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetFirstRanks(v []*DescribeAppsResponseBodyResultFirstRanks) *DescribeAppsResponseBodyResult {
	s.FirstRanks = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetId(v string) *DescribeAppsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetInterpretations(v map[string]interface{}) *DescribeAppsResponseBodyResult {
	s.Interpretations = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetIsCurrent(v bool) *DescribeAppsResponseBodyResult {
	s.IsCurrent = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetProgressPercent(v int32) *DescribeAppsResponseBodyResult {
	s.ProgressPercent = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetPrompts(v []map[string]interface{}) *DescribeAppsResponseBodyResult {
	s.Prompts = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetQueryProcessors(v []*DescribeAppsResponseBodyResultQueryProcessors) *DescribeAppsResponseBodyResult {
	s.QueryProcessors = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetQuota(v *DescribeAppsResponseBodyResultQuota) *DescribeAppsResponseBodyResult {
	s.Quota = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSchema(v *DescribeAppsResponseBodyResultSchema) *DescribeAppsResponseBodyResult {
	s.Schema = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSchemas(v []*DescribeAppsResponseBodyResultSchemas) *DescribeAppsResponseBodyResult {
	s.Schemas = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSecondRanks(v []*DescribeAppsResponseBodyResultSecondRanks) *DescribeAppsResponseBodyResult {
	s.SecondRanks = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetStatus(v string) *DescribeAppsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetSummaries(v []*DescribeAppsResponseBodyResultSummaries) *DescribeAppsResponseBodyResult {
	s.Summaries = v
	return s
}

func (s *DescribeAppsResponseBodyResult) SetType(v string) *DescribeAppsResponseBodyResult {
	s.Type = &v
	return s
}

type DescribeAppsResponseBodyResultCluster struct {
	MaxQueryClauseLength *int32 `json:"maxQueryClauseLength,omitempty" xml:"maxQueryClauseLength,omitempty"`
	MaxTimeoutMS         *int32 `json:"maxTimeoutMS,omitempty" xml:"maxTimeoutMS,omitempty"`
}

func (s DescribeAppsResponseBodyResultCluster) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultCluster) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultCluster) SetMaxQueryClauseLength(v int32) *DescribeAppsResponseBodyResultCluster {
	s.MaxQueryClauseLength = &v
	return s
}

func (s *DescribeAppsResponseBodyResultCluster) SetMaxTimeoutMS(v int32) *DescribeAppsResponseBodyResultCluster {
	s.MaxTimeoutMS = &v
	return s
}

type DescribeAppsResponseBodyResultDataSources struct {
	Fields     []map[string]interface{} `json:"fields,omitempty" xml:"fields,omitempty" type:"Repeated"`
	KeyField   *string                  `json:"keyField,omitempty" xml:"keyField,omitempty"`
	Parameters map[string]interface{}   `json:"parameters,omitempty" xml:"parameters,omitempty"`
	Plugins    map[string]interface{}   `json:"plugins,omitempty" xml:"plugins,omitempty"`
	SchemaName *string                  `json:"schemaName,omitempty" xml:"schemaName,omitempty"`
	TableName  *string                  `json:"tableName,omitempty" xml:"tableName,omitempty"`
	Type       *string                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppsResponseBodyResultDataSources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDataSources) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDataSources) SetFields(v []map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Fields = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetKeyField(v string) *DescribeAppsResponseBodyResultDataSources {
	s.KeyField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetParameters(v map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Parameters = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetPlugins(v map[string]interface{}) *DescribeAppsResponseBodyResultDataSources {
	s.Plugins = v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetSchemaName(v string) *DescribeAppsResponseBodyResultDataSources {
	s.SchemaName = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetTableName(v string) *DescribeAppsResponseBodyResultDataSources {
	s.TableName = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDataSources) SetType(v string) *DescribeAppsResponseBodyResultDataSources {
	s.Type = &v
	return s
}

type DescribeAppsResponseBodyResultDomain struct {
	Category  *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	Functions *DescribeAppsResponseBodyResultDomainFunctions `json:"functions,omitempty" xml:"functions,omitempty" type:"Struct"`
	Name      *string                                        `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppsResponseBodyResultDomain) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDomain) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDomain) SetCategory(v string) *DescribeAppsResponseBodyResultDomain {
	s.Category = &v
	return s
}

func (s *DescribeAppsResponseBodyResultDomain) SetFunctions(v *DescribeAppsResponseBodyResultDomainFunctions) *DescribeAppsResponseBodyResultDomain {
	s.Functions = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomain) SetName(v string) *DescribeAppsResponseBodyResultDomain {
	s.Name = &v
	return s
}

type DescribeAppsResponseBodyResultDomainFunctions struct {
	Algo    []*string `json:"algo,omitempty" xml:"algo,omitempty" type:"Repeated"`
	Qp      []*string `json:"qp,omitempty" xml:"qp,omitempty" type:"Repeated"`
	Service []*string `json:"service,omitempty" xml:"service,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyResultDomainFunctions) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultDomainFunctions) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetAlgo(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Algo = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetQp(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Qp = v
	return s
}

func (s *DescribeAppsResponseBodyResultDomainFunctions) SetService(v []*string) *DescribeAppsResponseBodyResultDomainFunctions {
	s.Service = v
	return s
}

type DescribeAppsResponseBodyResultFirstRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
	Type        *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeAppsResponseBodyResultFirstRanks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultFirstRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetActive(v bool) *DescribeAppsResponseBodyResultFirstRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetDescription(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetMeta(v interface{}) *DescribeAppsResponseBodyResultFirstRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetName(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultFirstRanks) SetType(v string) *DescribeAppsResponseBodyResultFirstRanks {
	s.Type = &v
	return s
}

type DescribeAppsResponseBodyResultQueryProcessors struct {
	Active     *bool                    `json:"active,omitempty" xml:"active,omitempty"`
	Category   *string                  `json:"category,omitempty" xml:"category,omitempty"`
	Domain     *string                  `json:"domain,omitempty" xml:"domain,omitempty"`
	Indexes    []*string                `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	Name       *string                  `json:"name,omitempty" xml:"name,omitempty"`
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
}

func (s DescribeAppsResponseBodyResultQueryProcessors) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultQueryProcessors) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetActive(v bool) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetCategory(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Category = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetDomain(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Domain = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetIndexes(v []*string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetName(v string) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQueryProcessors) SetProcessors(v []map[string]interface{}) *DescribeAppsResponseBodyResultQueryProcessors {
	s.Processors = v
	return s
}

type DescribeAppsResponseBodyResultQuota struct {
	ComputeResource *int32  `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	DocSize         *int32  `json:"docSize,omitempty" xml:"docSize,omitempty"`
	Qps             *int32  `json:"qps,omitempty" xml:"qps,omitempty"`
	Spec            *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s DescribeAppsResponseBodyResultQuota) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultQuota) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultQuota) SetComputeResource(v int32) *DescribeAppsResponseBodyResultQuota {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetDocSize(v int32) *DescribeAppsResponseBodyResultQuota {
	s.DocSize = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetQps(v int32) *DescribeAppsResponseBodyResultQuota {
	s.Qps = &v
	return s
}

func (s *DescribeAppsResponseBodyResultQuota) SetSpec(v string) *DescribeAppsResponseBodyResultQuota {
	s.Spec = &v
	return s
}

type DescribeAppsResponseBodyResultSchema struct {
	IndexSortConfig  []*DescribeAppsResponseBodyResultSchemaIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *DescribeAppsResponseBodyResultSchemaIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                                `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                                `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                              `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                                `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                                 `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *DescribeAppsResponseBodyResultSchemaTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBodyResultSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchema) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchema) SetIndexSortConfig(v []*DescribeAppsResponseBodyResultSchemaIndexSortConfig) *DescribeAppsResponseBodyResultSchema {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetIndexes(v *DescribeAppsResponseBodyResultSchemaIndexes) *DescribeAppsResponseBodyResultSchema {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetName(v string) *DescribeAppsResponseBodyResultSchema {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetRouteField(v string) *DescribeAppsResponseBodyResultSchema {
	s.RouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetRouteFieldValues(v []*string) *DescribeAppsResponseBodyResultSchema {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetSecondRouteField(v string) *DescribeAppsResponseBodyResultSchema {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetTables(v map[string]interface{}) *DescribeAppsResponseBodyResultSchema {
	s.Tables = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchema) SetTtlField(v *DescribeAppsResponseBodyResultSchemaTtlField) *DescribeAppsResponseBodyResultSchema {
	s.TtlField = v
	return s
}

type DescribeAppsResponseBodyResultSchemaIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemaIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) SetDirection(v string) *DescribeAppsResponseBodyResultSchemaIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexSortConfig) SetField(v string) *DescribeAppsResponseBodyResultSchemaIndexSortConfig {
	s.Field = &v
	return s
}

type DescribeAppsResponseBodyResultSchemaIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemaIndexes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) SetFilterFields(v []*string) *DescribeAppsResponseBodyResultSchemaIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemaIndexes {
	s.SearchFields = v
	return s
}

type DescribeAppsResponseBodyResultSchemaTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemaTtlField) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemaTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) SetName(v string) *DescribeAppsResponseBodyResultSchemaTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemaTtlField) SetTtl(v int64) *DescribeAppsResponseBodyResultSchemaTtlField {
	s.Ttl = &v
	return s
}

type DescribeAppsResponseBodyResultSchemas struct {
	IndexSortConfig  []*DescribeAppsResponseBodyResultSchemasIndexSortConfig `json:"indexSortConfig,omitempty" xml:"indexSortConfig,omitempty" type:"Repeated"`
	Indexes          *DescribeAppsResponseBodyResultSchemasIndexes           `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Struct"`
	Name             *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	RouteField       *string                                                 `json:"routeField,omitempty" xml:"routeField,omitempty"`
	RouteFieldValues []*string                                               `json:"routeFieldValues,omitempty" xml:"routeFieldValues,omitempty" type:"Repeated"`
	SecondRouteField *string                                                 `json:"secondRouteField,omitempty" xml:"secondRouteField,omitempty"`
	Tables           map[string]interface{}                                  `json:"tables,omitempty" xml:"tables,omitempty"`
	TtlField         *DescribeAppsResponseBodyResultSchemasTtlField          `json:"ttlField,omitempty" xml:"ttlField,omitempty" type:"Struct"`
}

func (s DescribeAppsResponseBodyResultSchemas) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemas) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemas) SetIndexSortConfig(v []*DescribeAppsResponseBodyResultSchemasIndexSortConfig) *DescribeAppsResponseBodyResultSchemas {
	s.IndexSortConfig = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetIndexes(v *DescribeAppsResponseBodyResultSchemasIndexes) *DescribeAppsResponseBodyResultSchemas {
	s.Indexes = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetName(v string) *DescribeAppsResponseBodyResultSchemas {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetRouteField(v string) *DescribeAppsResponseBodyResultSchemas {
	s.RouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetRouteFieldValues(v []*string) *DescribeAppsResponseBodyResultSchemas {
	s.RouteFieldValues = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetSecondRouteField(v string) *DescribeAppsResponseBodyResultSchemas {
	s.SecondRouteField = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetTables(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemas {
	s.Tables = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemas) SetTtlField(v *DescribeAppsResponseBodyResultSchemasTtlField) *DescribeAppsResponseBodyResultSchemas {
	s.TtlField = v
	return s
}

type DescribeAppsResponseBodyResultSchemasIndexSortConfig struct {
	Direction *string `json:"direction,omitempty" xml:"direction,omitempty"`
	Field     *string `json:"field,omitempty" xml:"field,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemasIndexSortConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasIndexSortConfig) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) SetDirection(v string) *DescribeAppsResponseBodyResultSchemasIndexSortConfig {
	s.Direction = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexSortConfig) SetField(v string) *DescribeAppsResponseBodyResultSchemasIndexSortConfig {
	s.Field = &v
	return s
}

type DescribeAppsResponseBodyResultSchemasIndexes struct {
	FilterFields []*string              `json:"filterFields,omitempty" xml:"filterFields,omitempty" type:"Repeated"`
	SearchFields map[string]interface{} `json:"searchFields,omitempty" xml:"searchFields,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemasIndexes) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) SetFilterFields(v []*string) *DescribeAppsResponseBodyResultSchemasIndexes {
	s.FilterFields = v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasIndexes) SetSearchFields(v map[string]interface{}) *DescribeAppsResponseBodyResultSchemasIndexes {
	s.SearchFields = v
	return s
}

type DescribeAppsResponseBodyResultSchemasTtlField struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	Ttl  *int64  `json:"ttl,omitempty" xml:"ttl,omitempty"`
}

func (s DescribeAppsResponseBodyResultSchemasTtlField) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSchemasTtlField) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) SetName(v string) *DescribeAppsResponseBodyResultSchemasTtlField {
	s.Name = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSchemasTtlField) SetTtl(v int64) *DescribeAppsResponseBodyResultSchemasTtlField {
	s.Ttl = &v
	return s
}

type DescribeAppsResponseBodyResultSecondRanks struct {
	Active      *bool       `json:"active,omitempty" xml:"active,omitempty"`
	Description *string     `json:"description,omitempty" xml:"description,omitempty"`
	Meta        interface{} `json:"meta,omitempty" xml:"meta,omitempty"`
	Name        *string     `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppsResponseBodyResultSecondRanks) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSecondRanks) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetActive(v bool) *DescribeAppsResponseBodyResultSecondRanks {
	s.Active = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetDescription(v string) *DescribeAppsResponseBodyResultSecondRanks {
	s.Description = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetMeta(v interface{}) *DescribeAppsResponseBodyResultSecondRanks {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultSecondRanks) SetName(v string) *DescribeAppsResponseBodyResultSecondRanks {
	s.Name = &v
	return s
}

type DescribeAppsResponseBodyResultSummaries struct {
	Meta []*DescribeAppsResponseBodyResultSummariesMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	Name *string                                        `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DescribeAppsResponseBodyResultSummaries) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSummaries) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSummaries) SetMeta(v []*DescribeAppsResponseBodyResultSummariesMeta) *DescribeAppsResponseBodyResultSummaries {
	s.Meta = v
	return s
}

func (s *DescribeAppsResponseBodyResultSummaries) SetName(v string) *DescribeAppsResponseBodyResultSummaries {
	s.Name = &v
	return s
}

type DescribeAppsResponseBodyResultSummariesMeta struct {
	Element  *string `json:"element,omitempty" xml:"element,omitempty"`
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	Field    *string `json:"field,omitempty" xml:"field,omitempty"`
	Len      *int32  `json:"len,omitempty" xml:"len,omitempty"`
	Snippet  *string `json:"snippet,omitempty" xml:"snippet,omitempty"`
}

func (s DescribeAppsResponseBodyResultSummariesMeta) String() string {
	return tea.Prettify(s)
}

func (s DescribeAppsResponseBodyResultSummariesMeta) GoString() string {
	return s.String()
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetElement(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Element = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetEllipsis(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Ellipsis = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetField(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Field = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetLen(v int32) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Len = &v
	return s
}

func (s *DescribeAppsResponseBodyResultSummariesMeta) SetSnippet(v string) *DescribeAppsResponseBodyResultSummariesMeta {
	s.Snippet = &v
	return s
}

type DescribeAppsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 72FAD77B-83F9-F393-BA8E-5834E2427BF8
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
	// The time when the task was created.
	//
	// example:
	//
	// 1581065837
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of data collected. Valid values:
	//
	// 	- behavior: behavioral data.
	//
	// 	- item_info: project information.
	//
	// 	- industry_specific: industry-specific data.
	//
	// example:
	//
	// BEHAVIOR
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	// The ID of the data collection task.
	//
	// example:
	//
	// 286
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry name. Valid values:
	//
	// 	- general
	//
	// 	- ecommerce
	//
	// example:
	//
	// GENERAL
	IndustryName *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	// The name of the data collection task.
	//
	// example:
	//
	// os_function_test_v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data collection feature. Valid values:
	//
	// 	- 0: The feature is disabled.
	//
	// 	- 1: The feature is being enabled.
	//
	// 	- 2: The feature is enabled.
	//
	// 	- 3: The feature failed to be enabled.
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The sundial ID.
	//
	// example:
	//
	// 1755
	SundialId *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	// The type of the source from which data was collected. Valid values:
	//
	// 	- server
	//
	// 	- web
	//
	// 	- app Note: Only server is supported.
	//
	// example:
	//
	// server
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the data collection task was updated.
	//
	// example:
	//
	// 1581065904
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataCollctionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// false
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*DescribeFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// ar_wear_edit_time
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
	//
	// example:
	//
	// ar_edit_time
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// example:
	//
	// timeliness_ms()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: [-100000,100000] (excluding 0).
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D7CCF454-472A-030E-F254-604520B028AA
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
	//
	// example:
	//
	// -
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	//
	// example:
	//
	// 1536233287
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// category_prediction
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1536233287
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
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	//
	// example:
	//
	// 1587398402
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule applies.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last updated.
	//
	// example:
	//
	// 1587398402
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type DescribeRegionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3B7E42BD-1D63-2F6B-C8E0-41BACEA76EEB
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
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
	// The console URL.
	//
	// example:
	//
	// https://opensearch-cn-hangzhou.console.aliyun.com
	ConsoleUrl *string `json:"consoleUrl,omitempty" xml:"consoleUrl,omitempty"`
	// The endpoint.
	//
	// example:
	//
	// opensearch.cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The region name.
	//
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"localName,omitempty" xml:"localName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 922DC0D9-31B5-45F9-47B7-37DC678D61A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the scheduled task.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 1587052801
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// -
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	//
	// example:
	//
	// 89047
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression.
	//
	// You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// random()+now()
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the expression.
	//
	// example:
	//
	// tests
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was last updated.
	//
	// example:
	//
	// 1587052801
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 5C1C1C45-C64A-AD30-565F-140871D57E5E
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
	//
	// example:
	//
	// 100298370
	AppGroupId *string `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The network type of the slow query optimization service. Valid values:
	//
	// 	- outer: the Internet
	//
	// 	- internal: the internal network
	//
	// example:
	//
	// internal
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The status of the slow query optimization service. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// 	- n/a
	//
	// example:
	//
	// disabled
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlowQueryStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The Associated information,output properties based on hierarchy.
	//
	// 	- **all**: Outputs associated app information
	//
	// example:
	//
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
	//
	// example:
	//
	// FFAEF396-10EF-53C7-1F51-518853880729
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the custom analyzer.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 728E89C6-8673-D39B-39A1-5BA2B56D448F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 728E89C6-8673-D39B-39A1-5BA2B56D448F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *EnableSlowQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body parameters.
	Body *Schema `json:"body,omitempty" xml:"body,omitempty"`
	// The specifications of the OpenSearch instance. This parameter is used to check whether the OpenSearch instance meets the special limits on an exclusive instance.
	//
	// Default value: opensearch.share.common.
	//
	// For more information, see the description of the spec field in the Quota topic.
	//
	// example:
	//
	// "opensearch.share.common"
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The response parameters.
	Result *GenerateMergedTableResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
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
	// The tables on which the JOIN operation is performed.
	//
	// example:
	//
	// -
	FromTable map[string]interface{} `json:"fromTable,omitempty" xml:"fromTable,omitempty"`
	// The wide table that is generated after the JOIN operation is performed on multiple tables.
	//
	// example:
	//
	// -
	MergeTable map[string]interface{} `json:"mergeTable,omitempty" xml:"mergeTable,omitempty"`
	// The primary key.
	//
	// example:
	//
	// -
	PrimaryKey *string `json:"primaryKey,omitempty" xml:"primaryKey,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateMergedTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The name or ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// -
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// general
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The industry. By default, this parameter is left empty, which indicates General-purpose Edition.
	//
	// example:
	//
	// ecommerce
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- PAAS. This is the default value.
	//
	// 	- SAAS.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model. The following features correspond to different model types:
	//
	// 	- CTR model: tf_checkpoint
	//
	// 	- Popularity model: pop
	//
	// 	- Category model: offline_inference
	//
	// 	- Hotword model: offline_inference
	//
	// 	- Shading model: offline_inference
	//
	// 	- Drop-down suggestion model: offline_inference
	//
	// 	- Word segmentation model: text
	//
	// 	- Word weight model: tf_checkpoint
	//
	// This parameter is required.
	//
	// example:
	//
	// tf_checkpoint
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
	//
	// example:
	//
	// Version.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionCurrentVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- PAAS
	//
	// 	- SAAS
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The configuration information about the instance.
	VersionConfig *GetFunctionCurrentVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	// The ID of the version.
	//
	// example:
	//
	// 101
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The name of the version.
	//
	// example:
	//
	// v1
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
	//
	// example:
	//
	// [                 {                     "name": "params1",                     "required": "true",                     "formItemProps": "{\\"required\\": true, \\"pattern?\\": \\"/^[a-zA-Z][a-zA-Z0-9_]{0,29}$/\\"}",                     "componentProps": "{\\"component\\": \\"Input\\", \\"attributes\\": {\\"defaultValue\\": \\"value1\\"}}"                 }             ]
	CreateParameters []*GetFunctionCurrentVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionCurrentVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used to use the instance online.
	//
	// example:
	//
	// []
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
	//
	// example:
	//
	// params1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// ""
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The dependency.
	//
	// example:
	//
	// ""
	Dependency *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	// The description.
	//
	// example:
	//
	// ""
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
	//
	// example:
	//
	// ""
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// ""
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionCurrentVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// DefaultInstance.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// cdn_waf
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// sh-bp1oi31w1jn4ctdyv
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The default running time.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// default instance not set.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 062BA91F-A568-5779-8A5B-9E62C9BB3DC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// {\\"Pagination\\": {\\"TotalCount\\": 0, \\"PageNumber\\": 1, \\"PageSize\\": 10}, \\"AntConsortiums\\": []}
	Result *GetFunctionDefaultInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	//
	// example:
	//
	// model1
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
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// 	- simple: displays only the basic information.
	//
	// 	- normal: displays information such as createParameters and cron. This is the default value.
	//
	// 	- detail: returns the details of the training task.
	//
	// example:
	//
	// detail
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
	//
	// example:
	//
	// Instance.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// instance not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 68ED4E1B-92B8-5821-A886-9C90686139EB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the instance.
	//
	// example:
	//
	// {}
	Result *GetFunctionInstanceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression used to schedule training, in the format of (Minutes Hours DayofMonth Month DayofWeek). If the value is empty, it indicates that no periodic training is performed.
	//
	// example:
	//
	// 0 3 ? \\	- 0,1,3,5 (at 3 a.m. on Sunday, Monday, Wednesday, and Friday)
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// instance descriptions
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string.
	//
	// example:
	//
	// {\\"dataReport\\":{},\\"errors\\":{}}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// ctr_test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// 	- user: The instance is created by user.
	//
	// 	- builtin: The instance is created by the system.
	//
	// example:
	//
	// user
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The status of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	//
	// 2.  available: Models can be used.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The information about the training task. This parameter is not displayed if no task is available.
	Task *GetFunctionInstanceResponseBodyResultTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Struct"`
	// The parameters that are used.
	UsageParameters []*GetFunctionInstanceResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	//
	// example:
	//
	// 101
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
	//
	// example:
	//
	// general
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	//
	// example:
	//
	// ecommerce
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	//
	// example:
	//
	// zh
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
	//
	// example:
	//
	// param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// value1
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
	// 	- success: succeeded
	//
	// 	- failed: failed
	//
	// 	- untrained: to be trained
	//
	// 	- pending: being scheduled
	//
	// 	- running: being trained
	//
	// example:
	//
	// success
	DagStatus *string `json:"DagStatus,omitempty" xml:"DagStatus,omitempty"`
	// The time consumed for the most recent run, in milliseconds.
	//
	// example:
	//
	// 1234
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
	//
	// example:
	//
	// use_param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// value1
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetFunctionResourceRequest struct {
	// The output level.
	//
	// Valid values:
	//
	// 	- simple
	//
	// 	- normal
	//
	// 	- detail
	//
	// example:
	//
	// detail
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
}

func (s GetFunctionResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceRequest) SetOutput(v string) *GetFunctionResourceRequest {
	s.Output = &v
	return s
}

type GetFunctionResourceResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
	//
	// example:
	//
	// Resource.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the API request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Resource not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E215C843-0795-5293-AC9A-14FE0723E890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	Result *GetFunctionResourceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// 	- OK
	//
	// 	- FAIL
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetFunctionResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBody) SetCode(v string) *GetFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetHttpCode(v int64) *GetFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetLatency(v float64) *GetFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetMessage(v string) *GetFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetRequestId(v string) *GetFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFunctionResourceResponseBody) SetResult(v *GetFunctionResourceResponseBodyResult) *GetFunctionResourceResponseBody {
	s.Result = v
	return s
}

func (s *GetFunctionResourceResponseBody) SetStatus(v string) *GetFunctionResourceResponseBody {
	s.Status = &v
	return s
}

type GetFunctionResourceResponseBodyResult struct {
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource data. The data structure varies with the resource type.
	Data *GetFunctionResourceResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// ""
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// rank
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The time when the resource was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The algorithm instances that are referenced.
	ReferencedInstances []*string `json:"ReferencedInstances,omitempty" xml:"ReferencedInstances,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// fg_json
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// raw_file
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s GetFunctionResourceResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResult) SetCreateTime(v int64) *GetFunctionResourceResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetData(v *GetFunctionResourceResponseBodyResultData) *GetFunctionResourceResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetDescription(v string) *GetFunctionResourceResponseBodyResult {
	s.Description = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetFunctionName(v string) *GetFunctionResourceResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetModifyTime(v int64) *GetFunctionResourceResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetReferencedInstances(v []*string) *GetFunctionResourceResponseBodyResult {
	s.ReferencedInstances = v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetResourceName(v string) *GetFunctionResourceResponseBodyResult {
	s.ResourceName = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResult) SetResourceType(v string) *GetFunctionResourceResponseBodyResult {
	s.ResourceType = &v
	return s
}

type GetFunctionResourceResponseBodyResultData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// abc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*GetFunctionResourceResponseBodyResultDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s GetFunctionResourceResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultData) SetContent(v string) *GetFunctionResourceResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultData) SetGenerators(v []*GetFunctionResourceResponseBodyResultDataGenerators) *GetFunctionResourceResponseBodyResultData {
	s.Generators = v
	return s
}

type GetFunctionResourceResponseBodyResultDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// id
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *GetFunctionResourceResponseBodyResultDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// output_feature_name
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s GetFunctionResourceResponseBodyResultDataGenerators) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGenerators) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetGenerator(v string) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Generator = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetInput(v *GetFunctionResourceResponseBodyResultDataGeneratorsInput) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Input = v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGenerators) SetOutput(v string) *GetFunctionResourceResponseBodyResultDataGenerators {
	s.Output = &v
	return s
}

type GetFunctionResourceResponseBodyResultDataGeneratorsInput struct {
	// The input features.
	Features []*GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInput) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInput) SetFeatures(v []*GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) *GetFunctionResourceResponseBodyResultDataGeneratorsInput {
	s.Features = v
	return s
}

type GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) SetName(v string) *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures) SetType(v string) *GetFunctionResourceResponseBodyResultDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

type GetFunctionResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFunctionResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *GetFunctionResourceResponse) SetHeaders(v map[string]*string) *GetFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *GetFunctionResourceResponse) SetStatusCode(v int32) *GetFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFunctionResourceResponse) SetBody(v *GetFunctionResourceResponseBody) *GetFunctionResourceResponse {
	s.Body = v
	return s
}

type GetFunctionTaskResponseBody struct {
	// The error code.
	//
	// example:
	//
	// Task.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	Result *GetFunctionTaskResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	//
	// example:
	//
	// 1647213406267
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The extended information, which is a JSON string.
	//
	// example:
	//
	// {\\"recall\\":91,\\"errors\\":[]}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The number of iterations.
	//
	// example:
	//
	// 1
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The progress. 90 indicates 90%.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// trigger__2021-03-05T12:18:41
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The timestamp that indicates the start time of the task. Unit: milliseconds.
	//
	// example:
	//
	// 1647212220000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// 	- running
	//
	// example:
	//
	// success
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// Version.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The maximum duration for which a task can be executed.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// version not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result body.
	//
	// example:
	//
	// []
	Result *GetFunctionVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature. Valid values:
	//
	// 	- PAAS
	//
	// 	- SAAS
	//
	// example:
	//
	// PAAS
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// The configuration information.
	VersionConfig *GetFunctionVersionResponseBodyResultVersionConfig `json:"VersionConfig,omitempty" xml:"VersionConfig,omitempty" type:"Struct"`
	// The ID of the version.
	//
	// example:
	//
	// 101
	VersionId *int64 `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
	// The name of the version.
	//
	// example:
	//
	// v1
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
	//
	// example:
	//
	// [                 {                     "name": "params1",                     "required": "true",                     "formItemProps": "{\\"required\\": true, \\"pattern?\\": \\"/^[a-zA-Z][a-zA-Z0-9_]{0,29}$/\\"}",                     "componentProps": "{\\"component\\": \\"Input\\", \\"attributes\\": {\\"defaultValue\\": \\"value1\\"}}"                 }             ]
	CreateParameters []*GetFunctionVersionResponseBodyResultVersionConfigCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The dependencies of the instance.
	Depends []*GetFunctionVersionResponseBodyResultVersionConfigDepends `json:"Depends,omitempty" xml:"Depends,omitempty" type:"Repeated"`
	// The parameters that are used during online use of the instance.
	//
	// example:
	//
	// []
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
	//
	// example:
	//
	// params1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// ""
	Condition *string `json:"Condition,omitempty" xml:"Condition,omitempty"`
	// The dependency.
	//
	// example:
	//
	// ""
	Dependency *string `json:"Dependency,omitempty" xml:"Dependency,omitempty"`
	// The description.
	//
	// example:
	//
	// ""
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
	//
	// example:
	//
	// ""
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the parameter is required.
	//
	// example:
	//
	// ""
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFunctionVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type GetScriptFileNamesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
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
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The name of the script file.
	//
	// example:
	//
	// my_cava_script.cava
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The time when the script file was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The path name of the script file.
	//
	// example:
	//
	// src
	PathName *string `json:"pathName,omitempty" xml:"pathName,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScriptFileNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 5C1C1C45-C64A-AD30-565F-140871D57E5E
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the script.
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
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The name of the script.
	//
	// example:
	//
	// rank_cava_20230606_v7
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The status of the script. For more information, see the description of the status response parameter in the ListSortScripts topic.
	//
	// example:
	//
	// released
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the script.
	//
	// example:
	//
	// cava_script
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

func (s *GetSortScriptResponseBodyResult) SetScriptName(v string) *GetSortScriptResponseBodyResult {
	s.ScriptName = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
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
	//
	// example:
	//
	// YWJjZGVmZw==
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The time when the script was created.
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The version of the script content.
	//
	// example:
	//
	// 123456
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The experiment details.\\
	//
	// For more information, see [ABTestExperiment](https://help.aliyun.com/document_detail/173617.html).
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
	// The time when the experiment was created.
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The experiment ID.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The group alias.
	//
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the experiment is in effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The experiment parameters.
	//
	// example:
	//
	// 1
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the experiment.
	//
	// Valid values: [0,100]
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the experiment was last modified.
	//
	// example:
	//
	// 1588842080
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestExperimentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The test groups.
	//
	// For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
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
	//
	// example:
	//
	// 1588839490
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 13466
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// Group_2020-5-7_15:23:3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last modified.
	//
	// example:
	//
	// 1588839490
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenarios.
	//
	// For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html).
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
	//
	// example:
	//
	// 1588836130
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 20404
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test group.
	//
	// example:
	//
	// kevintest_2020-5-7_15:21:482
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test scenario. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test scenario was last modified.
	//
	// example:
	//
	// 1588836129
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListABTestScenesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the instance. Exact match is used.
	//
	// example:
	//
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// "110123123"
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The method based on which applications are sorted. Valid values:
	//
	// 	- 0: sorts applications in descending order by creation time.
	//
	// 	- 1: sorts applications in descending order by modification time.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	SortBy *int32 `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The tags.
	Tags []*ListAppGroupsRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
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
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bar
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
	// The ID of the instance. Exact match is used.
	//
	// example:
	//
	// ops-cn-xxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// my_name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// "110123123"
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The method based on which applications are sorted. Valid values:
	//
	// 	- 0: sorts applications in descending order by creation time.
	//
	// 	- 1: sorts applications in descending order by modification time.
	//
	// Default value: 0.
	//
	// example:
	//
	// 0
	SortBy *int32 `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The tags.
	TagsShrink *string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the application.
	//
	// For more information, see [AppGroup](https://help.aliyun.com/document_detail/170000.html).
	//
	// example:
	//
	// []
	Result []*ListAppGroupsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billable item. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: queries per second (QPS).
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1575442875
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 110116134
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// "xxx"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The industry of the application.
	//
	// example:
	//
	// ""
	Domain     *string `json:"domain,omitempty" xml:"domain,omitempty"`
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The time when the application expired.
	//
	// example:
	//
	// "xxx"
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval state of the quotas. Valid values:
	//
	// 	- 0: The application is in service.
	//
	// 	- 1: The quotas are being reviewed.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 110116134
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// "xxx"
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock state. Valid values:
	//
	// 	- Unlock: The instance is unlocked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires.
	//
	// example:
	//
	// 0
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The application name.
	//
	// example:
	//
	// os_function_test_v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the application is created. Valid values:
	//
	// 	- 0: The application is being created.
	//
	// 	- 1: The application is created.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// "xxx"
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application. For more information, see [Quota](https://help.aliyun.com/document_detail/170001.html).
	//
	// example:
	//
	// {}
	Quota *ListAppGroupsResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The state of the application. Valid values:
	//
	// 	- producing: The application is being created.
	//
	// 	- review_pending: The application is being reviewed.
	//
	// 	- config_pending: The application is to be configured.
	//
	// 	- normal: The application is in service.
	//
	// 	- frozen: The application is frozen.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 0
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The application tags.
	Tags []*ListAppGroupsResponseBodyResultTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The type of the application. Valid values:
	//
	// 	- standard: a High-performance Search Edition application.
	//
	// *
	//
	// 	- enhanced: an Industry Algorithm Edition application.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1578916076
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

func (s *ListAppGroupsResponseBodyResult) SetEngineType(v string) *ListAppGroupsResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ListAppGroupsResponseBodyResult) SetExpireOn(v string) *ListAppGroupsResponseBodyResult {
	s.ExpireOn = &v
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
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic.
	//
	// 	- opensearch.share.common: shared general-purpose.
	//
	// 	- opensearch.share.compute: shared computing.
	//
	// 	- opensearch.share.storage: shared storage.
	//
	// 	- opensearch.private.common: exclusive general-purpose.
	//
	// 	- opensearch.private.compute: exclusive computing.
	//
	// 	- opensearch.private.storage: exclusive storage.
	//
	// example:
	//
	// opensearch.share.common
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
	// The tag key.
	//
	// example:
	//
	// foo
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// bar
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListDataCollectionsRequest struct {
	// 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	//
	// example:
	//
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
	//
	// example:
	//
	// 959D8782-B130-95EB-86CC-1F6ED447981F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the data collection tasks.
	//
	// For more information, see [DataCollection](https://help.aliyun.com/document_detail/173605.html).
	Result []*ListDataCollectionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the returned data collection tasks.
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// 1581065837
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the data that is collected by the task. Valid values:
	//
	// 	- behavior: behavioral data
	//
	// 	- item_info: project data
	//
	// 	- industry_specific: industry-specific data
	//
	// example:
	//
	// BEHAVIOR
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	// The ID of the data collection task.
	//
	// example:
	//
	// 286
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry to which the data collection task applies. Valid values:
	//
	// 	- general
	//
	// 	- ecommerce
	//
	// example:
	//
	// GENERAL
	IndustryName *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	// The name of the data collection task.
	//
	// example:
	//
	// os_function_test_v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data collection task. Valid values:
	//
	// 	- 0: disabled
	//
	// 	- 1: being enabled
	//
	// 	- 2: enabled
	//
	// 	- 3: failed to be enabled
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The ID of the sundial.
	//
	// example:
	//
	// 1755
	SundialId *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	// The type of the data source. Valid values:
	//
	// 	- server
	//
	// 	- web
	//
	// 	- app
	//
	// Note: Only server is supported.
	//
	// example:
	//
	// server
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the data collection task was updated.
	//
	// example:
	//
	// 1581065904
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCollectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The parameters of the data source. The value of the params parameter is a JSON string. The value must be URL-encoded.
	//
	// Different types of data sources use different parameters. For more information, see the following sections of the "DataSource" topic:
	//
	// 	- [rds](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [polardb](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [odps](https://help.aliyun.com/document_detail/170005.html)
	//
	// 	- [mysql](https://help.aliyun.com/document_detail/173627.html)
	//
	// 	- [drds](https://help.aliyun.com/document_detail/173627.html)
	//
	// This parameter is required.
	//
	// example:
	//
	// {}
	Params *string `json:"params,omitempty" xml:"params,omitempty"`
	// Specifies whether to return the original field types of the data source.
	//
	// example:
	//
	// false
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTableFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// This parameter is required.
	//
	// example:
	//
	// -
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataSourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each rough sort expression.
	//
	// For more information, see [FirstRank](https://help.aliyun.com/document_detail/170007.html).
	//
	// example:
	//
	// []
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	//
	// example:
	//
	// []
	Meta []*ListFirstRanksResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the cluster was updated.
	//
	// example:
	//
	// 0
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
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
	// For more information, see [Rough sort functions](https://help.aliyun.com/document_detail/180765.html).
	//
	// example:
	//
	// ""
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see [Rough sort functions](https://help.aliyun.com/document_detail/180765.html).
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: [-100000,100000] (excluding 0).
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFirstRanksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// "PAAS"
	FunctionType *string `json:"functionType,omitempty" xml:"functionType,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// tf_checkpoint
	ModelType *string `json:"modelType,omitempty" xml:"modelType,omitempty"`
	// The richness of the returned information. Valid values:
	//
	// 	- normal: displays information such as createParameters and cron. This is the default value.
	//
	// 	- simple: displays only the basic information.
	//
	// 	- detail: returns the details of the training task.
	//
	// example:
	//
	// normal
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// How the instance is created. Valid values:
	//
	// 	- builtin: The instance is created by system.
	//
	// 	- user: The instance is created by user. This is the default value.
	//
	// 	- all: all instances
	//
	// example:
	//
	// user
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
	//
	// example:
	//
	// Instance.NotExist
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message. If no error occurs, the parameter is left empty.
	//
	// example:
	//
	// instance not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the instances.
	//
	// example:
	//
	// []
	Result []*ListFunctionInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	//
	// example:
	//
	// "OK"
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
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
	//
	// example:
	//
	// {}
	Belongs *ListFunctionInstancesResponseBodyResultBelongs `json:"Belongs,omitempty" xml:"Belongs,omitempty" type:"Struct"`
	// The parameters of the instance.
	//
	// example:
	//
	// []
	CreateParameters []*ListFunctionInstancesResponseBodyResultCreateParameters `json:"CreateParameters,omitempty" xml:"CreateParameters,omitempty" type:"Repeated"`
	// The time when the instance was created.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cron expression used to schedule training, in the format of (Minutes Hours DayofMonth Month DayofWeek). If the value is empty, it indicates that no periodic training is performed.
	//
	// example:
	//
	// 0 3 ? \\	- 0,1,3,5 (at 3 a.m. on Sunday, Monday, Wednesday, and Friday)
	Cron *string `json:"Cron,omitempty" xml:"Cron,omitempty"`
	// The description.
	//
	// example:
	//
	// " "
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extended information, which is a JSON string. It includes model evaluation information and error information.
	//
	// example:
	//
	// "{\\"dataReport\\":{},\\"errors\\":{}}"
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// "ctr"
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// "PAAS"
	FunctionType *string `json:"FunctionType,omitempty" xml:"FunctionType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// "ctr_test"
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The type of the model.
	//
	// example:
	//
	// "tf_checkpoint"
	ModelType *string `json:"ModelType,omitempty" xml:"ModelType,omitempty"`
	// How the instance is created. Valid values:
	//
	// 	- user: The instance is created by user.
	//
	// 	- builtin: The instance is created by system.
	//
	// example:
	//
	// "user"
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The state of the instance. Valid values:
	//
	// 1.  unavailable: No model is available. Models must be trained before you can use them.
	//
	// 2.  available: Models can be used.
	//
	// example:
	//
	// available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters that are used.
	UsageParameters []*ListFunctionInstancesResponseBodyResultUsageParameters `json:"UsageParameters,omitempty" xml:"UsageParameters,omitempty" type:"Repeated"`
	// The ID of the version.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// "general"
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The industry.
	//
	// example:
	//
	// "ecommerce"
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The abbreviation of the language that applies.
	//
	// example:
	//
	// "zh"
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
	//
	// example:
	//
	// "param1"
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// "value1"
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
	//
	// example:
	//
	// use_param1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// value1
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
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListFunctionResourcesRequest struct {
	// The output level.
	//
	// Valid values:
	//
	// 	- simple
	//
	// 	- normal
	//
	// 	- detail
	//
	// example:
	//
	// detail
	Output *string `json:"output,omitempty" xml:"output,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of the resource.
	//
	// Valid values:
	//
	// 	- feature_generator
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- raw_file
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// feature_generator
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListFunctionResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesRequest) SetOutput(v string) *ListFunctionResourcesRequest {
	s.Output = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetPageNumber(v int32) *ListFunctionResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetPageSize(v int32) *ListFunctionResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListFunctionResourcesRequest) SetResourceType(v string) *ListFunctionResourcesRequest {
	s.ResourceType = &v
	return s
}

type ListFunctionResourcesResponseBody struct {
	// The error code returned. If no error occurs, this value is empty.
	//
	// example:
	//
	// Resource.InvalidResourceName
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The amount of time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Invalid resource name.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// "3A809095-C554-5CF5-8FCE-BE19D4673790"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The results returned.
	Result []*ListFunctionResourcesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request. Valid values: OK and FAIL.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of records that meet the requirements.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFunctionResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBody) SetCode(v string) *ListFunctionResourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetHttpCode(v int64) *ListFunctionResourcesResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetLatency(v float64) *ListFunctionResourcesResponseBody {
	s.Latency = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetMessage(v string) *ListFunctionResourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetRequestId(v string) *ListFunctionResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetResult(v []*ListFunctionResourcesResponseBodyResult) *ListFunctionResourcesResponseBody {
	s.Result = v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetStatus(v string) *ListFunctionResourcesResponseBody {
	s.Status = &v
	return s
}

func (s *ListFunctionResourcesResponseBody) SetTotalCount(v int64) *ListFunctionResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListFunctionResourcesResponseBodyResult struct {
	// The time when the resource was created. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The resource data. The data structure varies with the resource type.
	Data *ListFunctionResourcesResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// resource description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// rank
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The time when the resource was modified. Unit: milliseconds.
	//
	// example:
	//
	// 1234
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The algorithm instances that are referenced.
	ReferencedInstances []*string `json:"ReferencedInstances,omitempty" xml:"ReferencedInstances,omitempty" type:"Repeated"`
	// The name of the resource.
	//
	// example:
	//
	// fg_json
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The type of the resource.
	//
	// example:
	//
	// feature_generator
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResult) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResult) SetCreateTime(v int64) *ListFunctionResourcesResponseBodyResult {
	s.CreateTime = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetData(v *ListFunctionResourcesResponseBodyResultData) *ListFunctionResourcesResponseBodyResult {
	s.Data = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetDescription(v string) *ListFunctionResourcesResponseBodyResult {
	s.Description = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetFunctionName(v string) *ListFunctionResourcesResponseBodyResult {
	s.FunctionName = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetModifyTime(v int64) *ListFunctionResourcesResponseBodyResult {
	s.ModifyTime = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetReferencedInstances(v []*string) *ListFunctionResourcesResponseBodyResult {
	s.ReferencedInstances = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetResourceName(v string) *ListFunctionResourcesResponseBodyResult {
	s.ResourceName = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResult) SetResourceType(v string) *ListFunctionResourcesResponseBodyResult {
	s.ResourceType = &v
	return s
}

type ListFunctionResourcesResponseBodyResultData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// "abc"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*ListFunctionResourcesResponseBodyResultDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s ListFunctionResourcesResponseBodyResultData) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultData) SetContent(v string) *ListFunctionResourcesResponseBodyResultData {
	s.Content = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultData) SetGenerators(v []*ListFunctionResourcesResponseBodyResultDataGenerators) *ListFunctionResourcesResponseBodyResultData {
	s.Generators = v
	return s
}

type ListFunctionResourcesResponseBodyResultDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// combo
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *ListFunctionResourcesResponseBodyResultDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// feature1
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResultDataGenerators) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGenerators) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetGenerator(v string) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Generator = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetInput(v *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Input = v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGenerators) SetOutput(v string) *ListFunctionResourcesResponseBodyResultDataGenerators {
	s.Output = &v
	return s
}

type ListFunctionResourcesResponseBodyResultDataGeneratorsInput struct {
	// The input features.
	Features []*ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInput) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInput) SetFeatures(v []*ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) *ListFunctionResourcesResponseBodyResultDataGeneratorsInput {
	s.Features = v
	return s
}

type ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// Valid values:
	//
	// 	- item
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) SetName(v string) *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures) SetType(v string) *ListFunctionResourcesResponseBodyResultDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

type ListFunctionResourcesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFunctionResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFunctionResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListFunctionResourcesResponse) SetHeaders(v map[string]*string) *ListFunctionResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListFunctionResourcesResponse) SetStatusCode(v int32) *ListFunctionResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFunctionResourcesResponse) SetBody(v *ListFunctionResourcesResponseBody) *ListFunctionResourcesResponse {
	s.Body = v
	return s
}

type ListFunctionTasksRequest struct {
	// The end time is less than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1582646399
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 1.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The start time is greater than the specified time. Specify the time in the UNIX timestamp format. Unit: milliseconds.
	//
	// example:
	//
	// 1582214400
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// 	- running
	//
	// example:
	//
	// success
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
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// fail
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1638157479281
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// [         {             "functionName": "ctr",             "progress": 100,             "status": "success",             "startTime": 100010,             "endTime": 200020,             "extendInfo": "{\\"recall\\":91,\\"errors\\":[]}",             "runId": "trigger__2021-03-05T12:18:41"         }     ]
	Result []*ListFunctionTasksResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The status of the request.
	//
	// example:
	//
	// 200
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of records that meet the requirements.
	//
	// example:
	//
	// 2
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
	//
	// example:
	//
	// 100010
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The value is a JSON string. It includes model evaluation information and training error information.
	//
	// example:
	//
	// {\\"recall\\":91,\\"errors\\":[]}
	ExtendInfo *string `json:"ExtendInfo,omitempty" xml:"ExtendInfo,omitempty"`
	// The name of the feature.
	//
	// example:
	//
	// ctr
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The number of iterations.
	//
	// example:
	//
	// 2
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The progress. 90 indicates 90%.
	//
	// example:
	//
	// 90
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// trigger__2021-03-05T12:18:41
	RunId *string `json:"RunId,omitempty" xml:"RunId,omitempty"`
	// The timestamp that indicates the start time. Unit: milliseconds.
	//
	// example:
	//
	// 100010
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- success
	//
	// 	- failed
	//
	// 	- running
	//
	// example:
	//
	// success
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFunctionTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The number of the page to return. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// ["synonym"]
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each intervention dictionary.
	//
	// For more information, see [InterventionDictionary](https://help.aliyun.com/document_detail/173608.html).
	Result []*ListInterventionDictionariesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
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
	//
	// example:
	//
	// ""
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	//
	// example:
	//
	// 1539158325
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the intervention dictionary.
	//
	// example:
	//
	// 1
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// synonym
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1539158313
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries returned per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The intervention entry.
	//
	// example:
	//
	// test
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
	// The request ID.
	//
	// example:
	//
	// 516A02B7-2167-8D92-12D0-B639A2A0F3C5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about intervention entries.
	//
	// For more information, see [InterventionDictionaryEntry](https://help.aliyun.com/document_detail/173606.html).
	Result []*ListInterventionDictionaryEntriesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 8
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
	// The command. Valid values:
	//
	// 	- add
	//
	// 	- delete
	//
	// example:
	//
	// add
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// The timestamp when the intervention entry was created.
	//
	// example:
	//
	// 1536690285
	Created *int64 `json:"created,omitempty" xml:"created,omitempty"`
	// The content of an intervention entry for category prediction. The field value consists of key-value pairs. The key in a key-value pair indicates the ID of the category. The value in a key-value pair indicates the relevance to the category. A value of 0 indicates irrelevant. A value of 1 indicates slightly relevant. A value of 2 indicates relevant. Example: {"2":1, "100":0}
	//
	// example:
	//
	// {                 "100": "0",                 "200": "2"             }
	Relevance map[string]interface{} `json:"relevance,omitempty" xml:"relevance,omitempty"`
	// The status of the intervention entry. Valid value:
	//
	// 	- ACTIVE: The intervention entry takes effect.
	//
	// example:
	//
	// ACTIVE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The content of the intervention entry for term weight analysis.
	Tokens []*ListInterventionDictionaryEntriesResponseBodyResultTokens `json:"tokens,omitempty" xml:"tokens,omitempty" type:"Repeated"`
	// The timestamp when the intervention entry was last updated.
	//
	// example:
	//
	// 1537348987
	Updated *int64 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The intervention entry.
	//
	// example:
	//
	// \\u8fc7\\u513f
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
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The internal name of the identified entity type. Valid values:
	//
	// 	- brand
	//
	// 	- category
	//
	// 	- material
	//
	// 	- element
	//
	// 	- style
	//
	// 	- color
	//
	// 	- function
	//
	// 	- scenario
	//
	// 	- people
	//
	// 	- season
	//
	// 	- model
	//
	// 	- region
	//
	// 	- name
	//
	// 	- adjective
	//
	// 	- category-modifier
	//
	// 	- size
	//
	// 	- quality
	//
	// 	- suit
	//
	// 	- new-release
	//
	// 	- series
	//
	// 	- marketing
	//
	// 	- entertainment
	//
	// 	- organization
	//
	// 	- movie
	//
	// 	- game
	//
	// 	- number
	//
	// 	- unit
	//
	// 	- common
	//
	// 	- new-word
	//
	// 	- proper-noun
	//
	// 	- symbol
	//
	// 	- prefix
	//
	// 	- suffix
	//
	// 	- gift
	//
	// 	- negative
	//
	// 	- agent
	//
	// example:
	//
	// category
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The description of the internal name of the identified entity type.
	//
	// example:
	//
	// category
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The entity.
	//
	// example:
	//
	// category
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// Query keywords.
	//
	// This parameter is required.
	//
	// example:
	//
	// "hello world"
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
	//
	// example:
	//
	// 8F780CA8-D4D4-2FFE-B8AC-42040822C554
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The NER result.
	//
	// For more information, see [InterventionDictionaryEntry](https://help.aliyun.com/document_detail/173606.html).
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
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The internal name of the identified entity type. Valid values:
	//
	// 	- brand
	//
	// 	- category
	//
	// 	- material
	//
	// 	- element
	//
	// 	- style
	//
	// 	- color
	//
	// 	- function
	//
	// 	- scenario
	//
	// 	- people
	//
	// 	- season
	//
	// 	- model
	//
	// 	- region
	//
	// 	- name
	//
	// 	- adjective
	//
	// 	- category-modifier
	//
	// 	- size
	//
	// 	- quality
	//
	// 	- suit
	//
	// 	- new-release
	//
	// 	- series
	//
	// 	- marketing
	//
	// 	- entertainment
	//
	// 	- organization
	//
	// 	- movie
	//
	// 	- game
	//
	// 	- number
	//
	// 	- unit
	//
	// 	- common
	//
	// 	- new-word
	//
	// 	- proper-noun
	//
	// 	- symbol
	//
	// 	- prefix
	//
	// 	- suffix
	//
	// 	- gift
	//
	// 	- negative
	//
	// 	- agent
	//
	// example:
	//
	// category
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
	// The description of the internal name of the identified entity type.
	//
	// example:
	//
	// category
	TagLabel *string `json:"tagLabel,omitempty" xml:"tagLabel,omitempty"`
	// The entity.
	//
	// example:
	//
	// eaa73f35-007a-4be7-88c7-37dca4a04ab7
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
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryNerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterventionDictionaryRelatedEntitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ListProceedingsRequest struct {
	// Specifies whether the filtering is complete.
	//
	// example:
	//
	// true
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
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProceedingsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// This parameter is required.
	//
	// example:
	//
	// "abcde"
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
	//
	// example:
	//
	// 98724351-D6B2-5D8A-B089-7FFD1821A7E9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorAnalyzerResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The type of the industry.
	//
	// 	- ECOMMERCE
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The priority settings of entity types.
	//
	// For more information, see [Priority settings of entity types](https://help.aliyun.com/document_detail/173616.html).
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
	//
	// example:
	//
	// brand
	Label *string `json:"label,omitempty" xml:"label,omitempty"`
	// The priority of an entity type among entity types that have the same priority level. A smaller value indicates a higher priority. Default value: 0.
	//
	// example:
	//
	// 1
	Order *int32 `json:"order,omitempty" xml:"order,omitempty"`
	// The priority level of the entity type. Valid values:
	//
	// 	- HIGH
	//
	// 	- MIDDLE
	//
	// 	- LOW
	//
	// example:
	//
	// HIGH
	Priority *string `json:"priority,omitempty" xml:"priority,omitempty"`
	// The internal name of the entity type.
	//
	// example:
	//
	// brand
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorNersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The scope of query analysis rules to be queried. Default value: 0. Valid values:
	//
	// 	- 0: queries all query analysis rules.
	//
	// 	- 1: queries the default query analysis rules.
	//
	// 	- 2: queries the query analysis rules that are not the default rules.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	//
	// For more information, see [QueryProcessor](https://help.aliyun.com/document_detail/170014.html).
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
	// Indicates whether the query analysis rule is a default rule.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the query analysis rule was created.
	//
	// example:
	//
	// 1587398402
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry to which the query analysis rule is applied. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule is applied.
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// ner
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The features that are used in the query analysis rule.
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the query analysis rule was last modified.
	//
	// example:
	//
	// 1587398402
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// "3351A21F-705B-508C-9450-DA65A681547F"
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the tickets. For more information, see [QuotaReviewTask](https://help.aliyun.com/document_detail/173609.html).
	//
	// example:
	//
	// []
	Result []*ListQuotaReviewTasksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 500
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
	// The application ID.
	//
	// example:
	//
	// 120123456
	AppGroupId *int32 `json:"appGroupId,omitempty" xml:"appGroupId,omitempty"`
	// The application name.
	//
	// example:
	//
	// "td_test_os"
	AppGroupName *string `json:"appGroupName,omitempty" xml:"appGroupName,omitempty"`
	// The application type.
	//
	// example:
	//
	// "standard"
	AppGroupType *string `json:"appGroupType,omitempty" xml:"appGroupType,omitempty"`
	// Indicates whether the ticket is approved.
	//
	// example:
	//
	// true
	Approved *bool `json:"approved,omitempty" xml:"approved,omitempty"`
	// Indicates whether the application is available.
	//
	// example:
	//
	// true
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The time when the ticket was created.
	//
	// example:
	//
	// "2020-04-08T08:29:45+0000"
	GmtCreate *string `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The time when the ticket was last updated.
	//
	// example:
	//
	// "2020-04-08T08:36:36+0000"
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The ticket ID.
	//
	// example:
	//
	// 142
	Id *int32 `json:"id,omitempty" xml:"id,omitempty"`
	// The remarks.
	//
	// example:
	//
	// null
	Memo *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// The computing resource quota that is applied for.
	//
	// example:
	//
	// 6000
	NewComputeResource *int32 `json:"newComputeResource,omitempty" xml:"newComputeResource,omitempty"`
	// The storage capacity quota that is applied for.
	//
	// example:
	//
	// 1100
	NewSocSize *int32 `json:"newSocSize,omitempty" xml:"newSocSize,omitempty"`
	// The application specifications that are applied for.
	//
	// example:
	//
	// "opensearch.private.common"
	NewSpec *string `json:"newSpec,omitempty" xml:"newSpec,omitempty"`
	// The original quota of computing resources.
	//
	// example:
	//
	// 500
	OldComputeResource *int32 `json:"oldComputeResource,omitempty" xml:"oldComputeResource,omitempty"`
	// The original quota of storage capacity.
	//
	// example:
	//
	// 900
	OldDocSize *int32 `json:"oldDocSize,omitempty" xml:"oldDocSize,omitempty"`
	// The original specifications of the application.
	//
	// example:
	//
	// "opensearch.private.common"
	OldSpec *string `json:"oldSpec,omitempty" xml:"oldSpec,omitempty"`
	// Indicates whether the ticket is pending.
	//
	// example:
	//
	// false
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQuotaReviewTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The scheduled task type. Valid values:
	//
	// 	- wipe: data cleaning.
	//
	// 	- fork: reindexing.
	//
	// 	- check-status: application status check.
	//
	// 	- index: reindexing.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledTasksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 9C6351F5-2E2E-5249-888B-88A74E1B8A65
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSearchStrategiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about each fine sort expression.
	//
	// For more information, see [SecondRank](https://help.aliyun.com/document_detail/170008.html).
	Result []*ListSecondRanksResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// false
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the expression. This parameter appears only in the response.
	//
	// example:
	//
	// 890473
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter appears only in the response. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression.
	//
	// You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// random()+now()
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The name of the expression.
	//
	// example:
	//
	// tests
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was last updated.
	//
	// example:
	//
	// 1587052801
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecondRanksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 4406F40B-A0A2-9D5D-531F-3B6936567584
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
	// 	- PENDING: preparing
	//
	// 	- SUCCESS: succeeded
	//
	// 	- RUNNING: running
	//
	// 	- FAILED: failed
	//
	// 	- N/A: unknown
	//
	// example:
	//
	// "PENDING"
	AnalyzeStatus *string `json:"analyzeStatus,omitempty" xml:"analyzeStatus,omitempty"`
	// The timestamp that indicates the end of the time range to query.
	//
	// example:
	//
	// 1589990340
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The timestamp that indicates the beginning of the time range to query.
	//
	// example:
	//
	// 1589986800
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlowQueryCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// EB250CA0-ACFD-C5DE-17CD-01445BFE8AE5
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
	//
	// example:
	//
	// no data
	AppQuery *string `json:"appQuery,omitempty" xml:"appQuery,omitempty"`
	// The end of the time range that was queried.
	//
	// example:
	//
	// 1589990340
	End *int32 `json:"end,omitempty" xml:"end,omitempty"`
	// The ID of the optimization suggestion.
	//
	// example:
	//
	// 0
	Index *int32 `json:"index,omitempty" xml:"index,omitempty"`
	// The beginning of the time range that was queried.
	//
	// example:
	//
	// 1589986800
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSlowQueryQueriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the rough sort or fine sort expressions that are returned.
	//
	// For more information, see [FirstRank](https://help.aliyun.com/document_detail/170007.html) and [SecondRank](https://help.aliyun.com/document_detail/170008.html).
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The timestamp when the sort expression was created.
	//
	// example:
	//
	// 1655793690
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the sort expression.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the sort expression.
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The timestamp when the sort expression was updated.
	//
	// example:
	//
	// 1655793690
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSortExpressionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The scripts.
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
	//
	// example:
	//
	// 2020-04-02 20:21:14
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The time when the script was last modified.
	//
	// example:
	//
	// 2020-04-02 21:21:14
	ModifyTime *string `json:"modifyTime,omitempty" xml:"modifyTime,omitempty"`
	// The sort phase to which the script applies.
	//
	// example:
	//
	// second_rank
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The name of the script.
	//
	// example:
	//
	// test
	ScriptName *string `json:"scriptName,omitempty" xml:"scriptName,omitempty"`
	// The status of the script. Valid values:
	//
	// 	- configurable: The script is created, but no script files are uploaded.
	//
	// 	- not compiled: The script is not compiled.
	//
	// 	- compile failed: The compilation of the script failed.
	//
	// 	- compile successful: The script is compiled.
	//
	// 	- released: The script is published.
	//
	// example:
	//
	// released
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the script.
	//
	// example:
	//
	// cava_script
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSortScriptsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The fields to query. Format: columns=wordsTopPv.
	//
	// For more information, see [Metrics in statistical reports](https://help.aliyun.com/document_detail/187665.html).
	//
	// example:
	//
	// wordsTopPv
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// Specifies whether to use the distinct clause.
	//
	// example:
	//
	// true
	Distinct *bool `json:"distinct,omitempty" xml:"distinct,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The content of the query clause.
	//
	// example:
	//
	// "default:\\"OpenSearch\\""
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The content of the sort clause.
	//
	// example:
	//
	// "-id"
	SortBy *string `json:"sortBy,omitempty" xml:"sortBy,omitempty"`
	// The beginning of the time range to query. The default value is the timestamp of 00:00:00 on the current day.
	//
	// example:
	//
	// 1582214400
	StartTime *int32 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The end of the time range to query. The default value is the timestamp of 24:00:00 on the current day.
	//
	// example:
	//
	// 1682222400
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
	// The request ID.
	//
	// example:
	//
	// F76ACE3D-E510-EE2C-B7B1-39B3136A61EE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result. For more information, see
	//
	// 	- [Parameters of hotwords rankings](https://help.aliyun.com/document_detail/421248.html).
	//
	// example:
	//
	// []
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStatisticLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// pv,uv
	Columns *string `json:"columns,omitempty" xml:"columns,omitempty"`
	// 1582646399
	//
	// example:
	//
	// 1582646399
	EndTime *int32 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// bizType:test,sceneTag:myTag
	//
	// example:
	//
	// bizType:test,sceneTag:myTag
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// 1582214400
	//
	// example:
	//
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
	//
	// example:
	//
	// F65C8BB2-C14F-5983-888B-41C4E082D3BC
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The queried reports. Valid values:
	//
	// For more information about the metrics in data quality reports, see the Upload behavioral data section of [Data collection 2.0](https://help.aliyun.com/document_detail/131547.html).
	//
	// For more information about the metrics in application and A/B test reports, see the Core metrics section of [Metrics of statistical reports](https://help.aliyun.com/document_detail/187654.html).
	//
	// For more information about the metrics in query analysis reports, see the Query analysis metrics section of [Metrics of statistical reports](https://help.aliyun.com/document_detail/187654.html).
	//
	// example:
	//
	// []
	Result []map[string]interface{} `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of the queried reports.
	//
	// example:
	//
	// 43
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListStatisticReportResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 60
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIGDATA
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
	//
	// example:
	//
	// bm
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// Uefi
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
	//
	// example:
	//
	// 60
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 resource IDs.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIGDATA
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
	//
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 54041
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// hostGroup
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The key of the tag.
	//
	// example:
	//
	// GENIE_FUNCTION
	TagKey *string `json:"tagKey,omitempty" xml:"tagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// ALLOW
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The key to be used to query entries.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 516A02B7-2167-8D92-12D0-B639A2A0F3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The entries of the custom analyzer. For more information, see [UserAnalyzerEntry](https://www.alibabacloud.com/help/en/open-search/industry-algorithm-edition/useranalyzerentry).
	//
	// example:
	//
	// []
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries to return on each page. Default value: 10.
	//
	// example:
	//
	// 10
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// For more information, see [UserAnalyzer](https://help.aliyun.com/document_detail/178934.html).
	Result []*ListUserAnalyzersResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number.
	//
	// example:
	//
	// 1
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
	//
	// example:
	//
	// false
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The basic analyzer. Valid values:
	//
	// 	- chn_standard: [a common analyzer in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- chn_scene_name: an analyzer for person names in Chinese
	//
	// 	- chn_ecommerce: [an analyzer for E-commerce in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- chn_it_content: [an analyzer for IT content in Chinese](https://help.aliyun.com/document_detail/179424.html)
	//
	// 	- en_min: a small-granularity analyzer in English
	//
	// 	- th_standard: a common analyzer in Thai
	//
	// 	- th_ecommerce: an analyzer for E-commerce in Thai
	//
	// 	- vn_standard: a common analyzer in Vietnamese
	//
	// 	- chn_community_it: an analyzer for IT community content in Chinese
	//
	// 	- chn_ecommerce_general: a common analyzer for the E-commerce industry in Chinese
	//
	// 	- chn_esports_general: a common analyzer for the gaming industry in Chinese
	//
	// 	- chn_edu_question: an analyzer for question search of the education industry in Chinese
	//
	// example:
	//
	// chn_standard
	Business *string `json:"business,omitempty" xml:"business,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The dictionaries that are used by the custom analyzer.
	//
	// For more information, see [UserDict](https://help.aliyun.com/document_detail/178933.html).
	Dicts []*ListUserAnalyzersResponseBodyResultDicts `json:"dicts,omitempty" xml:"dicts,omitempty" type:"Repeated"`
	// The ID of the custom analyzer.
	//
	// example:
	//
	// 1234
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the custom analyzer.
	//
	// example:
	//
	// kevin_test2
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1588054131
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
	//
	// example:
	//
	// false
	Available *bool `json:"available,omitempty" xml:"available,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The number of intervention entries.
	//
	// example:
	//
	// -1
	EntriesCount *int32 `json:"entriesCount,omitempty" xml:"entriesCount,omitempty"`
	// The maximum number of intervention entries that can be created in the dictionary.
	//
	// example:
	//
	// 4
	EntriesLimit *int32 `json:"entriesLimit,omitempty" xml:"entriesLimit,omitempty"`
	// The ID of the dictionary.
	//
	// example:
	//
	// 123
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The type. Valid value:
	//
	// 	- segment
	//
	// example:
	//
	// segment
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1588054131
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAnalyzersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The online version of the application.
	//
	// example:
	//
	// 1223232
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// "test"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- general: general.
	//
	// 	- ecommerce: e-commerce.
	//
	// 	- education: education.
	//
	// 	- esports: electronic sports.
	//
	// 	- community: content community.
	//
	// example:
	//
	// "ecommerce"
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// rg-****
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// Specifies whether to verify the application before modification. Valid values: true and false.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned data.
	//
	// example:
	//
	// {}
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
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billable item. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: QPS.
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1590139524
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The type of the industry. Valid values:
	//
	// 	- GENERAL: general.
	//
	// 	- ECOMMERCE: e-commerce.
	//
	// 	- IT_CONTENT: IT content.
	//
	// example:
	//
	// GENERAL
	Domain     *string `json:"domain,omitempty" xml:"domain,omitempty"`
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The time when the application expired.
	//
	// example:
	//
	// 1
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// 	- 0: normal.
	//
	// 	- 1: being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 10030288
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock status. Valid values:
	//
	// 	- Unlock: The instance is unlocked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// 	- 0: The order is in progress.
	//
	// 	- 1: The order is complete.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// 1
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	//
	// example:
	//
	// {}
	Quota           *ModifyAppGroupResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId *string                                `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The state of the application. Valid values:
	//
	// 	- producing: being produced.
	//
	// 	- review_pending: being approved.
	//
	// 	- config_pending: to be configured.
	//
	// 	- normal: normal.
	//
	// 	- frozen: frozen.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard: a standard edition application.
	//
	// 	- advance: an advanced edition application of an old version. New versions are not supported for this edition.
	//
	// 	- enhanced: an advanced edition application of a new version.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last modified.
	//
	// example:
	//
	// 1590978265
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

func (s *ModifyAppGroupResponseBodyResult) SetEngineType(v string) *ModifyAppGroupResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ModifyAppGroupResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupResponseBodyResult {
	s.ExpireOn = &v
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

func (s *ModifyAppGroupResponseBodyResult) SetName(v string) *ModifyAppGroupResponseBodyResult {
	s.Name = &v
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

func (s *ModifyAppGroupResponseBodyResult) SetResourceGroupId(v string) *ModifyAppGroupResponseBodyResult {
	s.ResourceGroupId = &v
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
	// The computing resources. Unit: logical computing unit (LCU).
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications. Valid values:
	//
	// 	- opensearch.share.junior: basic.
	//
	// 	- opensearch.share.common: shared general-purpose.
	//
	// 	- opensearch.share.compute: shared computing.
	//
	// 	- opensearch.share.storage: shared storage.
	//
	// 	- opensearch.private.common: exclusive general-purpose.
	//
	// 	- opensearch.private.compute: exclusive computing.
	//
	// 	- opensearch.private.storage: exclusive storage.
	//
	// example:
	//
	// opensearch.share.common
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body *Quota `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to verify the application before modification. Valid values: true and false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	// 	- POSTPAY: pay-as-you-go
	//
	// 	- PREPAY: subscription
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing model. Valid values:
	//
	// 	- 1: computing resources
	//
	// 	- 2: queries per second (QPS)
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1590139542
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	EngineType  *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 1
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The approval status of the quotas. Valid values:
	//
	// 	- 0: The quotas are approved.
	//
	// 	- 1: The quotas are being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The ID of the application.
	//
	// example:
	//
	// 100302881
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 1
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock mode of the instance. Valid values:
	//
	// 	- Unlock: The instance is not locked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// lsh_test_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the order is complete. Valid values:
	//
	// 	- 0: The order is in progress.
	//
	// 	- 1: The order is complete.
	//
	// example:
	//
	// 1
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// 1000
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The information about the quotas of the application.
	Quota           *ModifyAppGroupQuotaResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	ResourceGroupId *string                                     `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// The status of the application. Valid values:
	//
	// 	- producing
	//
	// 	- review_pending
	//
	// 	- config_pending
	//
	// 	- normal
	//
	// 	- frozen
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application. Valid values:
	//
	// 	- standard: a standard application.
	//
	// 	- advance: an advanced application which is of an old application type. New applications cannot be of this type.
	//
	// 	- enhanced: an advanced application which is of a new application type.
	//
	// example:
	//
	// enhanced
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was last updated.
	//
	// example:
	//
	// 1590978265
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

func (s *ModifyAppGroupQuotaResponseBodyResult) SetEngineType(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.EngineType = &v
	return s
}

func (s *ModifyAppGroupQuotaResponseBodyResult) SetExpireOn(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ExpireOn = &v
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

func (s *ModifyAppGroupQuotaResponseBodyResult) SetName(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.Name = &v
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

func (s *ModifyAppGroupQuotaResponseBodyResult) SetResourceGroupId(v string) *ModifyAppGroupQuotaResponseBodyResult {
	s.ResourceGroupId = &v
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
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity. Unit: GB.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications of the application. Valid values:
	//
	// 	- opensearch.share.junior: basic
	//
	// 	- opensearch.share.common: shared general-purpose
	//
	// 	- opensearch.share.compute: shared computing
	//
	// 	- opensearch.share.storage: shared storage
	//
	// 	- opensearch.private.common: exclusive general-purpose
	//
	// 	- opensearch.private.compute: exclusive computing
	//
	// 	- opensearch.private.storage: exclusive storage
	//
	// example:
	//
	// opensearch.share.common
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAppGroupQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body *FirstRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the rough sort expression.
	//
	// example:
	//
	// 1
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The information about the expression.
	Meta []*ModifyFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
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
	//
	// example:
	//
	// “1 ”
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight. Valid values: -100000 to 100000. The value cannot be 0.
	//
	// example:
	//
	// 10
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//     "domain": "GENERAL",
	//
	//     "category": "",
	//
	//     "processors": [
	//
	//         {
	//
	//             "name": "synonym",
	//
	//             "useSystemDictionary": true
	//
	//         },
	//
	//         {
	//
	//             "name": "stop_word",
	//
	//             "useSystemDictionary": true
	//
	//         }
	//
	//     ]
	//
	// }
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the query analysis rule.
	//
	// example:
	//
	// {}
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
	// Indicates whether the query analysis rule is a default rule.
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the rule was created.
	//
	// example:
	//
	// 0
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of the industry to which the query analysis rule is applied. Valid values:
	//
	// 	- GENERAL
	//
	// 	- ECOMMERCE
	//
	// 	- IT_CONTENT
	//
	// example:
	//
	// GENERAL
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The indexes to which the query analysis rule is applied.
	//
	// example:
	//
	// ["default"]
	Indexes []*string `json:"indexes,omitempty" xml:"indexes,omitempty" type:"Repeated"`
	// The name of the query analysis rule.
	//
	// example:
	//
	// synonym
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis rule.
	//
	// example:
	//
	// []
	Processors []map[string]interface{} `json:"processors,omitempty" xml:"processors,omitempty" type:"Repeated"`
	// The time when the rule was updated.
	//
	// example:
	//
	// 1
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request parameters.
	//
	// example:
	//
	// The request parameters.
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the scheduled task.
	//
	// example:
	//
	// Array
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request parameters.
	Body *SecondRank `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
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
	// The request ID.
	//
	// example:
	//
	// C5E2F73C-8241-81F8-3A62-65478C5A3111
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the fine sort expression.
	//
	// example:
	//
	// {}
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The time when the expression was created.
	//
	// example:
	//
	// 1
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// "11"
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expression ID. This parameter is displayed only in the response.
	//
	// example:
	//
	// 890473
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Indicates whether the expression is the default one. This parameter is displayed only in the response. Valid values:
	//
	// 	- true: the expression is the default one.
	//
	// 	- false: the expression is not the default one.
	//
	// example:
	//
	// true
	IsDefault *string `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
	// Indicates whether the expression is a system expression. This parameter is displayed only in the response. Valid values:
	//
	// 	- true: The expression is a system expression.
	//
	// 	- false:The expression is not a system expression
	//
	// example:
	//
	// false
	IsSys *string `json:"isSys,omitempty" xml:"isSys,omitempty"`
	// The content of the fine sort expression. You can define an expression that consists of fields, feature functions, and mathematical functions to implement complex sort logic.
	//
	// example:
	//
	// cate_id > 0 and cate_id < 1000
	Meta *string `json:"meta,omitempty" xml:"meta,omitempty"`
	// The expression name.
	//
	// example:
	//
	// lsh_second_1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The time when the expression was updated.
	//
	// example:
	//
	// 1
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type PushInterventionDictionaryEntriesRequest struct {
	// The request body.
	Body []map[string]interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// Specifies whether to check the validity of input parameters. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**: checks only the validity of input parameters.
	//
	// 	- **false**: checks the validity of input parameters and creates an attribution configuration.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
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
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushInterventionDictionaryEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The entries of the custom analyzer.
	Entries []*PushUserAnalyzerEntriesRequestEntries `json:"entries,omitempty" xml:"entries,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The operation to be performed on the entries.
	//
	// Valid values:
	//
	// 	- add
	//
	// 	- delete
	//
	// example:
	//
	// "add"
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// The key to be used to query entries.
	//
	// example:
	//
	// "testvalue"
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Specifies whether to further analyze the terms that are generated after the search query is analyzed.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	SplitEnabled *bool `json:"splitEnabled,omitempty" xml:"splitEnabled,omitempty"`
	// The analysis result.
	//
	// example:
	//
	// "test value"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result returned.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type ReleaseSortScriptResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 33477D76-C380-2D84-A4AD-043F52876CB1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The return result.
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 3AA29D02-54F3-8569-F71A-90E1B7BE4E7E
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// —
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveDataCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// E676FAB6-A0AC-64D9-F9D7-D0D33C930CFF
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
	//
	// example:
	//
	// true
	Active *bool `json:"active,omitempty" xml:"active,omitempty"`
	// The description of the expression.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The content of the expression.
	Meta []*RemoveFirstRankResponseBodyResultMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Repeated"`
	// The name of the expression.
	//
	// example:
	//
	// default
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
	//
	// example:
	//
	// ""
	Arg *string `json:"arg,omitempty" xml:"arg,omitempty"`
	// The attribute, feature function, or field to be searched for.
	//
	// For more information about supported feature functions, see Rough sort functions.
	//
	// example:
	//
	// static_bm25()
	Attribute *string `json:"attribute,omitempty" xml:"attribute,omitempty"`
	// The weight.
	//
	// Valid values: [-100000,100000] (excluding 0).
	//
	// example:
	//
	// 10
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveFirstRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 06BBD41A-5F72-34E4-2DAF-E43B0526051D
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
	//
	// example:
	//
	// ""
	Analyzer *string `json:"analyzer,omitempty" xml:"analyzer,omitempty"`
	// The time when the intervention dictionary was created.
	//
	// example:
	//
	// 1539158313
	Created *string `json:"created,omitempty" xml:"created,omitempty"`
	// The name of the intervention dictionary.
	//
	// example:
	//
	// tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type of the intervention dictionary. Valid values:
	//
	// 	- stopword: an intervention dictionary for stop word filtering
	//
	// 	- synonym: an intervention dictionary for synonym configuration
	//
	// 	- correction: an intervention dictionary for spelling correction
	//
	// 	- category_prediction: an intervention dictionary for category prediction
	//
	// 	- ner: an intervention dictionary for named entity recognition (NER)
	//
	// 	- term_weighting: an intervention dictionary for term weight analysis
	//
	// example:
	//
	// synonym
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the intervention dictionary was last updated.
	//
	// example:
	//
	// 1539158313
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveInterventionDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	//
	// example:
	//
	// []
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveQueryProcessorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveScheduledTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// F5099063-6B86-F398-D843-905F9EFB683A
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveSecondRankResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The ID of the request.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The renewal request body.
	Body *PrepayOrderInfo `json:"body,omitempty" xml:"body,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 74db41d8cd3c784209093aa76afbe89e
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the application was renewed.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewAppGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// AC5F78BA-66B9-545B-9CF1-8F542E682E1F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// {}
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
	// The billing method. Valid values:
	//
	// 	- POSTPAY: pay-as-you-go.
	//
	// 	- PREPAY: subscription.
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// The billing type. Valid values:
	//
	// 	- 1: computing resources.
	//
	// 	- 2: queries per second (QPS).
	//
	// example:
	//
	// 1
	ChargingWay *int32 `json:"chargingWay,omitempty" xml:"chargingWay,omitempty"`
	// The code of the commodity.
	//
	// example:
	//
	// opensearch
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// The timestamp when the application was created.
	//
	// example:
	//
	// 1588054131
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the current online version.
	//
	// example:
	//
	// 100302903
	CurrentVersion *string `json:"currentVersion,omitempty" xml:"currentVersion,omitempty"`
	// The description of the application.
	//
	// example:
	//
	// ""
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// -
	ExpireOn *string `json:"expireOn,omitempty" xml:"expireOn,omitempty"`
	// The ID of the rough sort expression.
	//
	// example:
	//
	// 0
	FirstRankAlgoDeploymentId *int32 `json:"firstRankAlgoDeploymentId,omitempty" xml:"firstRankAlgoDeploymentId,omitempty"`
	// The approval state of the quotas. Valid values:
	//
	// 	- 0: The approval status is normal.
	//
	// 	- 1: The quotas are being approved.
	//
	// example:
	//
	// 0
	HasPendingQuotaReviewTask *int32 `json:"hasPendingQuotaReviewTask,omitempty" xml:"hasPendingQuotaReviewTask,omitempty"`
	// The application ID.
	//
	// example:
	//
	// -
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// -
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The lock state. Valid values:
	//
	// 	- Unlock: The instance is unlocked.
	//
	// 	- LockByExpiration: The instance is automatically locked after it expires.
	//
	// 	- ManualLock: The instance is manually locked.
	//
	// example:
	//
	// Unlock
	LockMode *string `json:"lockMode,omitempty" xml:"lockMode,omitempty"`
	// Indicates whether the instance is automatically locked after it expires. Valid values:
	//
	// 	- 0: The instance is not automatically locked after it expires.
	//
	// 	- 1: The instance is automatically locked after it expires.
	//
	// example:
	//
	// 0
	LockedByExpiration *int32 `json:"lockedByExpiration,omitempty" xml:"lockedByExpiration,omitempty"`
	// The name of the order.
	//
	// example:
	//
	// -
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The ID of the fine sort expression that is being created.
	//
	// example:
	//
	// 0
	PendingSecondRankAlgoDeploymentId *int32 `json:"pendingSecondRankAlgoDeploymentId,omitempty" xml:"pendingSecondRankAlgoDeploymentId,omitempty"`
	// The ID of the order that is in progress.
	//
	// example:
	//
	// -
	ProcessingOrderId *string `json:"processingOrderId,omitempty" xml:"processingOrderId,omitempty"`
	// Indicates whether the order is produced.
	//
	// example:
	//
	// 0
	Produced *int32 `json:"produced,omitempty" xml:"produced,omitempty"`
	// The name of the A/B test group.
	//
	// example:
	//
	// -
	ProjectId *string `json:"projectId,omitempty" xml:"projectId,omitempty"`
	// The configuration information.
	Quota *ReplaceAppGroupCommodityCodeResponseBodyResultQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// The ID of the fine sort expression.
	//
	// example:
	//
	// 0
	SecondRankAlgoDeploymentId *int32 `json:"secondRankAlgoDeploymentId,omitempty" xml:"secondRankAlgoDeploymentId,omitempty"`
	// The status of the application.
	//
	// example:
	//
	// normal
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The timestamp when the current online version was published.
	//
	// example:
	//
	// 1590486386
	SwitchedTime *int32 `json:"switchedTime,omitempty" xml:"switchedTime,omitempty"`
	// The type of the application.
	//
	// example:
	//
	// -
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The timestamp when the application was updated.
	//
	// example:
	//
	// 1581065904
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
	// The versions.
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
	// The number of computing resources configured.
	//
	// example:
	//
	// 20
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// The storage capacity.
	//
	// example:
	//
	// 1
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// The specifications configured.
	//
	// example:
	//
	// -
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
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceAppGroupCommodityCodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The script content that is encoded in the Base64 format.
	//
	// example:
	//
	// 4769#0: *28194492991 a client request body is buffered to a temporary file /usr/local/webserver/openresty/nginx/client_body_temp/0000624474,
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The version number of the script content.
	//
	// example:
	//
	// 2022-12-01
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveSortScriptFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// N/A
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSlowQueryAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// This parameter is required.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProductVersion
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// The tags. You can specify a maximum number of 20 tags.
	//
	// This parameter is required.
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
	//
	// example:
	//
	// cloud_manage
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 31261301
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
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request parameters.
	//
	// example:
	//
	// {
	//
	//   "name": "kevintest-analyzer"
	//
	// }
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
	// The ID of the request.
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The custom analyzer.
	//
	// example:
	//
	// []
	Result map[string]interface{} `json:"result,omitempty" xml:"result,omitempty"`
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindESUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 0A6EB64B-B4C8-CF02-810F-E660812972FF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The data returned.
	//
	// example:
	//
	// {}
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindEsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	//
	// This parameter is required.
	ResourceId []*string `json:"resourceId,omitempty" xml:"resourceId,omitempty" type:"Repeated"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProductVersion
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
	//
	// example:
	//
	// true
	All *bool `json:"all,omitempty" xml:"all,omitempty"`
	// The resource IDs. You can specify a maximum number of 50 IDs.
	//
	// This parameter is required.
	ResourceIdShrink *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// The resource type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ProductVersion
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
	//
	// example:
	//
	// 1-A-0-B
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body. For more information, see [ABTestExperiment](https://help.aliyun.com/document_detail/173617.html).
	Body *ABTestExperiment `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. No endpoint is created. The system checks whether your AccessKey is valid, whether Resource Access Management (RAM) users are authorized, and whether the required parameters are set.
	//
	// 	- false (default): creates an endpoint immediately.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
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
	//
	// example:
	//
	// 1588842080
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The test ID.
	//
	// example:
	//
	// 12888
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The alias of the test.
	//
	// example:
	//
	// test1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Indicates whether the test is in effect. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The test parameters.
	//
	// example:
	//
	// {}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is routed to the test. Valid values: [0,100]
	//
	// example:
	//
	// 30
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test was last modified.
	//
	// example:
	//
	// 1588842080
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The returned results.
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestFixedFlowDividersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body. For more information, see [ABTestGroup](https://help.aliyun.com/document_detail/178935.html).
	Body *ABTestGroup `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. No endpoint is created. The system checks whether your AccessKey is valid, whether Resource Access Management (RAM) users are authorized, and whether the required parameters are set.
	//
	// 	- false (default): creates an endpoint immediately.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// \\"\\"1111\\"\\"
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
	//
	// example:
	//
	// 1588839490
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test group.
	//
	// example:
	//
	// 13466
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The alias of the test group.
	//
	// example:
	//
	// Group_2020-5-7_15:23:3
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the test group was last updated.
	//
	// example:
	//
	// 1588839490
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the test scenario. For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html).
	//
	// example:
	//
	// {}
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
	//
	// example:
	//
	// 1596527691
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The ID of the test scenario.
	//
	// example:
	//
	// 20614
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The name of the test scenario.
	//
	// example:
	//
	// kevintest22
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test. Valid values:
	//
	// 	- true: The test is started.
	//
	// 	- false: The test is stopped.
	//
	// example:
	//
	// true
	Online *bool `json:"online,omitempty" xml:"online,omitempty"`
	// The parameters of the A/B test.
	//
	// example:
	//
	// {}
	Params map[string]interface{} `json:"params,omitempty" xml:"params,omitempty"`
	// The percentage of traffic that is allocated to the A/B test. Valid values: 0 to 100.
	//
	// example:
	//
	// 111
	Traffic *int32 `json:"traffic,omitempty" xml:"traffic,omitempty"`
	// The time when the test scenario was last modified.
	//
	// example:
	//
	// 1596527691
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body []*string `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// true
	//
	// example:
	//
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
	//
	// example:
	//
	// D77D0DAF-790D-F5F5-A9C0-133738165014
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFetchFieldsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// "pop_test"
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
	//
	// example:
	//
	// DefaultInstance.SetFail
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 123
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A4D487A9-A456-5AA5-A9C6-B7BF2889CF74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
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
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionDefaultInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// {             "name": "title_field",             "value": "title"         }
	CreateParameters []*UpdateFunctionInstanceRequestCreateParameters `json:"createParameters,omitempty" xml:"createParameters,omitempty" type:"Repeated"`
	// The cron expression used to schedule periodic training, in the format of (Minutes Hours DayofMonth Month DayofWeek). The default value is empty, which indicates that no periodic training is performed. DayofWeek 0 indicates Sunday.
	//
	// example:
	//
	// "0 3 ? 	- 0,1,3,5"
	Cron *string `json:"cron,omitempty" xml:"cron,omitempty"`
	// The description of the instance.
	//
	// example:
	//
	// test instance
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
	//
	// example:
	//
	// title_field
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// title
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
	//
	// example:
	//
	// allow_dict_id
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// 123
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
	//
	// example:
	//
	// "Instance.NotExist"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request, in milliseconds.
	//
	// example:
	//
	// 10
	Latency *int64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message.
	//
	// example:
	//
	// "instance not exist."
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// "3A809095-C554-5CF5-8FCE-BE19D4673790"
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request. Valid values:
	//
	// 	- OK: The request was successful.
	//
	// 	- FAIL: The request failed.
	//
	// example:
	//
	// OK
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
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

type UpdateFunctionResourceRequest struct {
	// The resource data. The data structure varies with the resource type.
	Data *UpdateFunctionResourceRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the resource.
	//
	// example:
	//
	// updated description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s UpdateFunctionResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequest) SetData(v *UpdateFunctionResourceRequestData) *UpdateFunctionResourceRequest {
	s.Data = v
	return s
}

func (s *UpdateFunctionResourceRequest) SetDescription(v string) *UpdateFunctionResourceRequest {
	s.Description = &v
	return s
}

type UpdateFunctionResourceRequestData struct {
	// The content of the file that corresponds to a resource of the raw_file type.
	//
	// example:
	//
	// abc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The feature generators that correspond to resources of the feature_generator type.
	Generators []*UpdateFunctionResourceRequestDataGenerators `json:"Generators,omitempty" xml:"Generators,omitempty" type:"Repeated"`
}

func (s UpdateFunctionResourceRequestData) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceRequestData) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestData) SetContent(v string) *UpdateFunctionResourceRequestData {
	s.Content = &v
	return s
}

func (s *UpdateFunctionResourceRequestData) SetGenerators(v []*UpdateFunctionResourceRequestDataGenerators) *UpdateFunctionResourceRequestData {
	s.Generators = v
	return s
}

type UpdateFunctionResourceRequestDataGenerators struct {
	// The type of the feature generator.
	//
	// example:
	//
	// combo
	Generator *string `json:"Generator,omitempty" xml:"Generator,omitempty"`
	// The input.
	Input *UpdateFunctionResourceRequestDataGeneratorsInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The name of the output feature.
	//
	// example:
	//
	// feature1
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
}

func (s UpdateFunctionResourceRequestDataGenerators) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGenerators) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetGenerator(v string) *UpdateFunctionResourceRequestDataGenerators {
	s.Generator = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetInput(v *UpdateFunctionResourceRequestDataGeneratorsInput) *UpdateFunctionResourceRequestDataGenerators {
	s.Input = v
	return s
}

func (s *UpdateFunctionResourceRequestDataGenerators) SetOutput(v string) *UpdateFunctionResourceRequestDataGenerators {
	s.Output = &v
	return s
}

type UpdateFunctionResourceRequestDataGeneratorsInput struct {
	// The input features.
	Features []*UpdateFunctionResourceRequestDataGeneratorsInputFeatures `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
}

func (s UpdateFunctionResourceRequestDataGeneratorsInput) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGeneratorsInput) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInput) SetFeatures(v []*UpdateFunctionResourceRequestDataGeneratorsInputFeatures) *UpdateFunctionResourceRequestDataGeneratorsInput {
	s.Features = v
	return s
}

type UpdateFunctionResourceRequestDataGeneratorsInputFeatures struct {
	// The name of the feature.
	//
	// example:
	//
	// system_item_id
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the feature.
	//
	// example:
	//
	// item
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateFunctionResourceRequestDataGeneratorsInputFeatures) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceRequestDataGeneratorsInputFeatures) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) SetName(v string) *UpdateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Name = &v
	return s
}

func (s *UpdateFunctionResourceRequestDataGeneratorsInputFeatures) SetType(v string) *UpdateFunctionResourceRequestDataGeneratorsInputFeatures {
	s.Type = &v
	return s
}

type UpdateFunctionResourceResponseBody struct {
	// The error code. If no error occurs, this parameter is left empty.
	//
	// example:
	//
	// InvalidRequest
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpCode *int64 `json:"HttpCode,omitempty" xml:"HttpCode,omitempty"`
	// The time consumed for the request. Unit: milliseconds.
	//
	// example:
	//
	// 123
	Latency *float64 `json:"Latency,omitempty" xml:"Latency,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Invalid request.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7E375703-5B12-5466-8B48-C4D01AE1291A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the request.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateFunctionResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceResponseBody) SetCode(v string) *UpdateFunctionResourceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetHttpCode(v int64) *UpdateFunctionResourceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetLatency(v float64) *UpdateFunctionResourceResponseBody {
	s.Latency = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetMessage(v string) *UpdateFunctionResourceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetRequestId(v string) *UpdateFunctionResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateFunctionResourceResponseBody) SetStatus(v string) *UpdateFunctionResourceResponseBody {
	s.Status = &v
	return s
}

type UpdateFunctionResourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFunctionResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFunctionResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFunctionResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateFunctionResourceResponse) SetHeaders(v map[string]*string) *UpdateFunctionResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateFunctionResourceResponse) SetStatusCode(v int32) *UpdateFunctionResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFunctionResourceResponse) SetBody(v *UpdateFunctionResourceResponseBody) *UpdateFunctionResourceResponse {
	s.Body = v
	return s
}

type UpdateSearchStrategyRequest struct {
	// The request body.
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
	// The ID of the request.
	//
	// example:
	//
	// ABCDEFGH
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
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSearchStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	//
	// example:
	//
	// 9F165784-5507-5342-ABF3-545293F9808A
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSortScriptResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request body.
	Body []*UpdateSummariesRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
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
	// The HTML tag that is used to highlight terms in red.
	//
	// example:
	//
	// "em"
	Element *string `json:"element,omitempty" xml:"element,omitempty"`
	// The connector that is used to connect segments.
	//
	// example:
	//
	// "..."
	Ellipsis *string `json:"ellipsis,omitempty" xml:"ellipsis,omitempty"`
	// The field.
	//
	// example:
	//
	// "title"
	Field *string `json:"field,omitempty" xml:"field,omitempty"`
	// The length of a segment.
	//
	// example:
	//
	// 50
	Len *int32 `json:"len,omitempty" xml:"len,omitempty"`
	// The number of segments.
	//
	// example:
	//
	// 1
	Snippet *int32 `json:"snippet,omitempty" xml:"snippet,omitempty"`
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
	// The request ID.
	//
	// example:
	//
	// 7A389E09-7964-5A2B-FE9D-F6CFA7162852
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSummariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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
	// The request parameter. For more information, see [DataSource](https://help.aliyun.com/document_detail/170005.html).
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
	// Id of the request
	//
	// example:
	//
	// 8FA2B338-AFDC-46B4-A132-B5487820C2BF
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The result returned.
	//
	// example:
	//
	// []
	Result []*ValidateDataSourcesResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
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
	// The code returned for the request.
	//
	// example:
	//
	// SUCCEED
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The data source.
	//
	// example:
	//
	// {}
	DataSource *ValidateDataSourcesResponseBodyResultDataSource `json:"dataSource,omitempty" xml:"dataSource,omitempty" type:"Struct"`
	// The status of the execution.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
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
	// The parameters of the data source.
	//
	// example:
	//
	// {}
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// user_activity_decision
	TableName *string `json:"tableName,omitempty" xml:"tableName,omitempty"`
	// The type of the data source.
	//
	// example:
	//
	// rds
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

// Summary:
//
// Binds a custom analyzer to an Elasticsearch instance.
//
// @param request - BindESUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindESUserAnalyzerResponse
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

// Summary:
//
// Binds a custom analyzer to an Elasticsearch instance.
//
// @param request - BindESUserAnalyzerRequest
//
// @return BindESUserAnalyzerResponse
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

// Summary:
//
// Binds an Elasticsearch instance.
//
// @param request - BindEsInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BindEsInstanceResponse
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

// Summary:
//
// Binds an Elasticsearch instance.
//
// @param request - BindEsInstanceRequest
//
// @return BindEsInstanceResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompileSortScriptResponse
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

// @return CompileSortScriptResponse
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

// Summary:
//
// Creates an experiment.
//
// @param request - CreateABTestExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestExperimentResponse
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

// Summary:
//
// Creates an experiment.
//
// @param request - CreateABTestExperimentRequest
//
// @return CreateABTestExperimentResponse
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

// Summary:
//
// Creates a test group.
//
// @param request - CreateABTestGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestGroupResponse
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

// Summary:
//
// Creates a test group.
//
// @param request - CreateABTestGroupRequest
//
// @return CreateABTestGroupResponse
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

// @param request - CreateABTestSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateABTestSceneResponse
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

// @param request - CreateABTestSceneRequest
//
// @return CreateABTestSceneResponse
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

// Summary:
//
// Creates a version for an OpenSearch application.
//
// Description:
//
//   When you create a standard application, a new version of the application is created if the specified application name already exists.
//
// 	- When you create a version of an existing application, you must specify the autoSwitch and realtimeShared parameters.
//
// 	- When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
//
// 	- When you create a version of an existing application, the modification of the value of the quota parameter does not take effect.
//
// @param request - CreateAppRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppResponse
func (client *Client) CreateAppWithOptions(appGroupIdentity *string, request *CreateAppRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoSwitch)) {
		body["autoSwitch"] = request.AutoSwitch
	}

	if !tea.BoolValue(util.IsUnset(request.Cluster)) {
		body["cluster"] = request.Cluster
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["dataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Domain)) {
		body["domain"] = request.Domain
	}

	if !tea.BoolValue(util.IsUnset(request.FetchFields)) {
		body["fetchFields"] = request.FetchFields
	}

	if !tea.BoolValue(util.IsUnset(request.FirstRanks)) {
		body["firstRanks"] = request.FirstRanks
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		body["networkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.QueryProcessors)) {
		body["queryProcessors"] = request.QueryProcessors
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		body["schema"] = request.Schema
	}

	if !tea.BoolValue(util.IsUnset(request.Schemas)) {
		body["schemas"] = request.Schemas
	}

	if !tea.BoolValue(util.IsUnset(request.SecondRanks)) {
		body["secondRanks"] = request.SecondRanks
	}

	if !tea.BoolValue(util.IsUnset(request.Summaries)) {
		body["summaries"] = request.Summaries
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
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

// Summary:
//
// Creates a version for an OpenSearch application.
//
// Description:
//
//   When you create a standard application, a new version of the application is created if the specified application name already exists.
//
// 	- When you create a version of an existing application, you must specify the autoSwitch and realtimeShared parameters.
//
// 	- When you create a version of an existing application, the value of the quota parameter is the same as that of the quota parameter in the previous version of the application.
//
// 	- When you create a version of an existing application, the modification of the value of the quota parameter does not take effect.
//
// @param request - CreateAppRequest
//
// @return CreateAppResponse
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

// Summary:
//
// Create an OpenSearch application.
//
// @param request - CreateAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupResponse
func (client *Client) CreateAppGroupWithOptions(request *CreateAppGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["chargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Quota)) {
		body["quota"] = request.Quota
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["resourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
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

// Summary:
//
// Create an OpenSearch application.
//
// @param request - CreateAppGroupRequest
//
// @return CreateAppGroupResponse
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

// @param request - CreateAppGroupCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateAppGroupCredentialsResponse
func (client *Client) CreateAppGroupCredentialsWithOptions(appGroupIdentity *string, request *CreateAppGroupCredentialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppGroupCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DryRun)) {
		query["dryRun"] = request.DryRun
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateAppGroupCredentials"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/credentials"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateAppGroupCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - CreateAppGroupCredentialsRequest
//
// @return CreateAppGroupCredentialsResponse
func (client *Client) CreateAppGroupCredentials(appGroupIdentity *string, request *CreateAppGroupCredentialsRequest) (_result *CreateAppGroupCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppGroupCredentialsResponse{}
	_body, _err := client.CreateAppGroupCredentialsWithOptions(appGroupIdentity, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a rough sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified rough sort expression. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateFirstRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFirstRankResponse
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

// Summary:
//
// Creates a rough sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified rough sort expression. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateFirstRankRequest
//
// @return CreateFirstRankResponse
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

// Summary:
//
// Creates an algorithm instance of a feature.
//
// Description:
//
// You can call the [GetFunctionCurrentVersion](https://help.aliyun.com/document_detail/421377.html) operation to query the latest version of a feature. The response of the operation includes the createParameters parameter that is used to create an algorithm instance, the usageParameters parameter, and the requirements for setting these parameters.
//
// @param request - CreateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionInstanceResponse
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

// Summary:
//
// Creates an algorithm instance of a feature.
//
// Description:
//
// You can call the [GetFunctionCurrentVersion](https://help.aliyun.com/document_detail/421377.html) operation to query the latest version of a feature. The response of the operation includes the createParameters parameter that is used to create an algorithm instance, the usageParameters parameter, and the requirements for setting these parameters.
//
// @param request - CreateFunctionInstanceRequest
//
// @return CreateFunctionInstanceResponse
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

// Summary:
//
// Creates an algorithm resource for a specific feature.
//
// @param request - CreateFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionResourceResponse
func (client *Client) CreateFunctionResourceWithOptions(appGroupIdentity *string, functionName *string, request *CreateFunctionResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFunctionResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["ResourceName"] = request.ResourceName
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		body["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFunctionResource"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFunctionResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates an algorithm resource for a specific feature.
//
// @param request - CreateFunctionResourceRequest
//
// @return CreateFunctionResourceResponse
func (client *Client) CreateFunctionResource(appGroupIdentity *string, functionName *string, request *CreateFunctionResourceRequest) (_result *CreateFunctionResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFunctionResourceResponse{}
	_body, _err := client.CreateFunctionResourceWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Starts a training task for an algorithm instance.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFunctionTaskResponse
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

// Summary:
//
// Starts a training task for an algorithm instance.
//
// @return CreateFunctionTaskResponse
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

// Summary:
//
// Create an intervention dictionary.
//
// @param request - CreateInterventionDictionaryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateInterventionDictionaryResponse
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

// Summary:
//
// Create an intervention dictionary.
//
// @param request - CreateInterventionDictionaryRequest
//
// @return CreateInterventionDictionaryResponse
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

// Summary:
//
// Creates a query analysis rule. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateQueryProcessorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateQueryProcessorResponse
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

// Summary:
//
// Creates a query analysis rule. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not set this parameter.
//
// @param request - CreateQueryProcessorRequest
//
// @return CreateQueryProcessorResponse
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

// @param request - CreateScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateScheduledTaskResponse
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

// @param request - CreateScheduledTaskRequest
//
// @return CreateScheduledTaskResponse
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

// Summary:
//
// Creates a query policy.
//
// @param request - CreateSearchStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSearchStrategyResponse
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

// Summary:
//
// Creates a query policy.
//
// @param request - CreateSearchStrategyRequest
//
// @return CreateSearchStrategyResponse
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

// Summary:
//
// Creates a fine sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified fine sort expression. The default value of dryRun is false if you do not set this parameter.
//
// @param request - CreateSecondRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSecondRankResponse
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

// Summary:
//
// Creates a fine sort expression for a version of an OpenSearch application. If you set dryRun to true, this operation checks the specified fine sort expression. The default value of dryRun is false if you do not set this parameter.
//
// @param request - CreateSecondRankRequest
//
// @return CreateSecondRankResponse
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

// Summary:
//
// Creates a sort script.
//
// @param request - CreateSortScriptRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSortScriptResponse
func (client *Client) CreateSortScriptWithOptions(appGroupIdentity *string, appVersionId *string, request *CreateSortScriptRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSortScriptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Scope)) {
		body["scope"] = request.Scope
	}

	if !tea.BoolValue(util.IsUnset(request.ScriptName)) {
		body["scriptName"] = request.ScriptName
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
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

// Summary:
//
// Creates a sort script.
//
// @param request - CreateSortScriptRequest
//
// @return CreateSortScriptResponse
func (client *Client) CreateSortScript(appGroupIdentity *string, appVersionId *string, request *CreateSortScriptRequest) (_result *CreateSortScriptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSortScriptResponse{}
	_body, _err := client.CreateSortScriptWithOptions(appGroupIdentity, appVersionId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// @param request - CreateUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserAnalyzerResponse
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

// @param request - CreateUserAnalyzerRequest
//
// @return CreateUserAnalyzerResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestExperimentResponse
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

// @return DeleteABTestExperimentResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestGroupResponse
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

// @return DeleteABTestGroupResponse
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

// Summary:
//
// Deletes an A/B test scenario.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteABTestSceneResponse
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

// Summary:
//
// Deletes an A/B test scenario.
//
// @return DeleteABTestSceneResponse
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

// Summary:
//
// Deletes an algorithm instance. Before you delete an instance, make sure that it is not in use to prevent service interruptions.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionInstanceResponse
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

// Summary:
//
// Deletes an algorithm instance. Before you delete an instance, make sure that it is not in use to prevent service interruptions.
//
// @return DeleteFunctionInstanceResponse
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

// Summary:
//
// Deletes an algorithm resource.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionResourceResponse
func (client *Client) DeleteFunctionResourceWithOptions(appGroupIdentity *string, functionName *string, resourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFunctionResourceResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFunctionResource"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(resourceName))),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFunctionResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes an algorithm resource.
//
// @return DeleteFunctionResourceResponse
func (client *Client) DeleteFunctionResource(appGroupIdentity *string, functionName *string, resourceName *string) (_result *DeleteFunctionResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFunctionResourceResponse{}
	_body, _err := client.DeleteFunctionResourceWithOptions(appGroupIdentity, functionName, resourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a training task. The training task in progress cannot be deleted.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFunctionTaskResponse
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

// Summary:
//
// Deletes a training task. The training task in progress cannot be deleted.
//
// @return DeleteFunctionTaskResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSortScriptResponse
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

// @return DeleteSortScriptResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSortScriptFileResponse
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

// @return DeleteSortScriptFileResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestExperimentResponse
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

// @return DescribeABTestExperimentResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestGroupResponse
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

// @return DescribeABTestGroupResponse
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

// Summary:
//
// Queries the information about an A/B test scenario.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeABTestSceneResponse
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

// Summary:
//
// Queries the information about an A/B test scenario.
//
// @return DescribeABTestSceneResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppResponse
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

// @return DescribeAppResponse
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

// Summary:
//
// Queries the details of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppGroupResponse
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

// Summary:
//
// Queries the details of an OpenSearch application.
//
// @return DescribeAppGroupResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppStatisticsResponse
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

// @return DescribeAppStatisticsResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeAppsResponse
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

// @return DescribeAppsResponse
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

// Summary:
//
// Queries the details of a data collection task of an application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDataCollctionResponse
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

// Summary:
//
// Queries the details of a data collection task of an application.
//
// @return DescribeDataCollctionResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeFirstRankResponse
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

// @return DescribeFirstRankResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInterventionDictionaryResponse
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

// @return DescribeInterventionDictionaryResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeQueryProcessorResponse
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

// @return DescribeQueryProcessorResponse
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

// Summary:
//
// Queries the endpoints of all regions that support OpenSearch.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
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

// Summary:
//
// Queries the endpoints of all regions that support OpenSearch.
//
// @return DescribeRegionsResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeScheduledTaskResponse
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

// @return DescribeScheduledTaskResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSecondRankResponse
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

// @return DescribeSecondRankResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSlowQueryStatusResponse
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

// @return DescribeSlowQueryStatusResponse
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

// @param request - DescribeUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeUserAnalyzerResponse
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

// @param request - DescribeUserAnalyzerRequest
//
// @return DescribeUserAnalyzerResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DisableSlowQueryResponse
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

// @return DisableSlowQueryResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return EnableSlowQueryResponse
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

// @return EnableSlowQueryResponse
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

// Summary:
//
// Queries the information about a wide table that is generated after a JOIN operation is performed on multiple tables.
//
// @param request - GenerateMergedTableRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GenerateMergedTableResponse
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

// Summary:
//
// Queries the information about a wide table that is generated after a JOIN operation is performed on multiple tables.
//
// @param request - GenerateMergedTableRequest
//
// @return GenerateMergedTableResponse
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

// Summary:
//
// Queries the type of an industry.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
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

// Summary:
//
// Queries the type of an industry.
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
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

// Summary:
//
// Queries the version information about the current feature when you create an instance.
//
// @param request - GetFunctionCurrentVersionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionCurrentVersionResponse
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

// Summary:
//
// Queries the version information about the current feature when you create an instance.
//
// @param request - GetFunctionCurrentVersionRequest
//
// @return GetFunctionCurrentVersionResponse
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

// Summary:
//
// Queries the algorithm instance that an application uses by default.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionDefaultInstanceResponse
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

// Summary:
//
// Queries the algorithm instance that an application uses by default.
//
// @return GetFunctionDefaultInstanceResponse
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

// Summary:
//
// Queries the details of an algorithm instance by instance name.
//
// @param request - GetFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionInstanceResponse
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

// Summary:
//
// Queries the details of an algorithm instance by instance name.
//
// @param request - GetFunctionInstanceRequest
//
// @return GetFunctionInstanceResponse
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

// Summary:
//
// Queries an algorithm resource.
//
// @param request - GetFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionResourceResponse
func (client *Client) GetFunctionResourceWithOptions(appGroupIdentity *string, functionName *string, resourceName *string, request *GetFunctionResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFunctionResourceResponse, _err error) {
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
		Action:      tea.String("GetFunctionResource"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(resourceName))),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFunctionResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries an algorithm resource.
//
// @param request - GetFunctionResourceRequest
//
// @return GetFunctionResourceResponse
func (client *Client) GetFunctionResource(appGroupIdentity *string, functionName *string, resourceName *string, request *GetFunctionResourceRequest) (_result *GetFunctionResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFunctionResourceResponse{}
	_body, _err := client.GetFunctionResourceWithOptions(appGroupIdentity, functionName, resourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the details of a training task.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionTaskResponse
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

// Summary:
//
// Queries the details of a training task.
//
// @return GetFunctionTaskResponse
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

// Summary:
//
// Queries version information by version ID.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFunctionVersionResponse
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

// Summary:
//
// Queries version information by version ID.
//
// @return GetFunctionVersionResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetScriptFileNamesResponse
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

// @return GetScriptFileNamesResponse
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

// Summary:
//
// Queries the details of a query policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSearchStrategyResponse
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

// Summary:
//
// Queries the details of a query policy.
//
// @return GetSearchStrategyResponse
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

// Summary:
//
// Queries the details of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSortScriptResponse
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

// Summary:
//
// Queries the details of a sort script.
//
// @return GetSortScriptResponse
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

// Summary:
//
// Queries the content of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSortScriptFileResponse
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

// Summary:
//
// Queries the content of a sort script.
//
// @return GetSortScriptFileResponse
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

// Summary:
//
// Queries a list of experiments.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestExperimentsResponse
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

// Summary:
//
// Queries a list of experiments.
//
// @return ListABTestExperimentsResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestFixedFlowDividersResponse
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

// @return ListABTestFixedFlowDividersResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestGroupsResponse
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

// @return ListABTestGroupsResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListABTestScenesResponse
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

// @return ListABTestScenesResponse
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

// Summary:
//
// Queries a list of OpenSearch applications.
//
// Description:
//
//   This operation allows you to query applications by application name, instance ID, and application type.
//
// 	- This operation allows you to sort the applications based on their creation time.
//
// 	- This operation supports the parameters for paging.
//
// @param tmpReq - ListAppGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAppGroupsResponse
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

// Summary:
//
// Queries a list of OpenSearch applications.
//
// Description:
//
//   This operation allows you to query applications by application name, instance ID, and application type.
//
// 	- This operation allows you to sort the applications based on their creation time.
//
// 	- This operation supports the parameters for paging.
//
// @param request - ListAppGroupsRequest
//
// @return ListAppGroupsResponse
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

// @param request - ListDataCollectionsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataCollectionsResponse
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

// @param request - ListDataCollectionsRequest
//
// @return ListDataCollectionsResponse
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

// Summary:
//
// Queries all fields in a table of a data source. This operation is for internal use only.
//
// @param request - ListDataSourceTableFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTableFieldsResponse
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

// Summary:
//
// Queries all fields in a table of a data source. This operation is for internal use only.
//
// @param request - ListDataSourceTableFieldsRequest
//
// @return ListDataSourceTableFieldsResponse
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

// @param request - ListDataSourceTablesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDataSourceTablesResponse
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

// @param request - ListDataSourceTablesRequest
//
// @return ListDataSourceTablesResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFirstRanksResponse
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

// @return ListFirstRanksResponse
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

// Summary:
//
// Queries all algorithm instances of a user, which meet specified conditions.
//
// @param request - ListFunctionInstancesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionInstancesResponse
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

// Summary:
//
// Queries all algorithm instances of a user, which meet specified conditions.
//
// @param request - ListFunctionInstancesRequest
//
// @return ListFunctionInstancesResponse
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

// Summary:
//
// Queries algorithm resources.
//
// @param request - ListFunctionResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionResourcesResponse
func (client *Client) ListFunctionResourcesWithOptions(appGroupIdentity *string, functionName *string, request *ListFunctionResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFunctionResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Output)) {
		query["output"] = request.Output
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["pageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["pageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["resourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFunctionResources"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFunctionResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries algorithm resources.
//
// @param request - ListFunctionResourcesRequest
//
// @return ListFunctionResourcesResponse
func (client *Client) ListFunctionResources(appGroupIdentity *string, functionName *string, request *ListFunctionResourcesRequest) (_result *ListFunctionResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFunctionResourcesResponse{}
	_body, _err := client.ListFunctionResourcesWithOptions(appGroupIdentity, functionName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the training tasks. The returned results are sorted by start time in descending order.
//
// @param request - ListFunctionTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFunctionTasksResponse
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

// Summary:
//
// Queries the training tasks. The returned results are sorted by start time in descending order.
//
// @param request - ListFunctionTasksRequest
//
// @return ListFunctionTasksResponse
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

// @param request - ListInterventionDictionariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionariesResponse
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

// @param request - ListInterventionDictionariesRequest
//
// @return ListInterventionDictionariesResponse
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

// Summary:
//
// Queries the intervention entries in an intervention dictionary.
//
// @param request - ListInterventionDictionaryEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryEntriesResponse
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

// Summary:
//
// Queries the intervention entries in an intervention dictionary.
//
// @param request - ListInterventionDictionaryEntriesRequest
//
// @return ListInterventionDictionaryEntriesResponse
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

// @param request - ListInterventionDictionaryNerResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryNerResultsResponse
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

// @param request - ListInterventionDictionaryNerResultsRequest
//
// @return ListInterventionDictionaryNerResultsResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListInterventionDictionaryRelatedEntitiesResponse
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

// @return ListInterventionDictionaryRelatedEntitiesResponse
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

// Summary:
//
// 查看当前的处理流
//
// @param request - ListProceedingsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListProceedingsResponse
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

// Summary:
//
// 查看当前的处理流
//
// @param request - ListProceedingsRequest
//
// @return ListProceedingsResponse
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

// Summary:
//
// Queries the results of a query analysis test. This API operation is available only to existing applications of OpenSearch Open Source Compatible Edition.
//
// @param request - ListQueryProcessorAnalyzerResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorAnalyzerResultsResponse
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

// Summary:
//
// Queries the results of a query analysis test. This API operation is available only to existing applications of OpenSearch Open Source Compatible Edition.
//
// @param request - ListQueryProcessorAnalyzerResultsRequest
//
// @return ListQueryProcessorAnalyzerResultsResponse
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

// Summary:
//
// Queries the recommended priority settings of entity types for named entity recognition (NER).
//
// @param request - ListQueryProcessorNersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorNersResponse
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

// Summary:
//
// Queries the recommended priority settings of entity types for named entity recognition (NER).
//
// @param request - ListQueryProcessorNersRequest
//
// @return ListQueryProcessorNersResponse
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

// Summary:
//
// Queries a list of query analysis rules that are configured for a version of an OpenSearch application.
//
// @param request - ListQueryProcessorsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQueryProcessorsResponse
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

// Summary:
//
// Queries a list of query analysis rules that are configured for a version of an OpenSearch application.
//
// @param request - ListQueryProcessorsRequest
//
// @return ListQueryProcessorsResponse
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

// Summary:
//
// Queries tickets that are submitted to apply for quotas for an OpenSearch application.
//
// @param request - ListQuotaReviewTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListQuotaReviewTasksResponse
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

// Summary:
//
// Queries tickets that are submitted to apply for quotas for an OpenSearch application.
//
// @param request - ListQuotaReviewTasksRequest
//
// @return ListQuotaReviewTasksResponse
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

// Summary:
//
// Queries a list of scheduled tasks of an OpenSearch application.
//
// @param request - ListScheduledTasksRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListScheduledTasksResponse
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

// Summary:
//
// Queries a list of scheduled tasks of an OpenSearch application.
//
// @param request - ListScheduledTasksRequest
//
// @return ListScheduledTasksResponse
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

// Summary:
//
// Queries the details of query policies.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSearchStrategiesResponse
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

// Summary:
//
// Queries the details of query policies.
//
// @return ListSearchStrategiesResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSecondRanksResponse
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

// @return ListSecondRanksResponse
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

// Summary:
//
// Queries the suggestions that are provided by Optimization Master for slow queries.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSlowQueryCategoriesResponse
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

// Summary:
//
// Queries the suggestions that are provided by Optimization Master for slow queries.
//
// @return ListSlowQueryCategoriesResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSlowQueryQueriesResponse
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

// @return ListSlowQueryQueriesResponse
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

// Summary:
//
// Queries a list of sort expressions that are configured for a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSortExpressionsResponse
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

// Summary:
//
// Queries a list of sort expressions that are configured for a version of an OpenSearch application.
//
// @return ListSortExpressionsResponse
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

// Summary:
//
// Queries all sort scripts of an application version.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSortScriptsResponse
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

// Summary:
//
// Queries all sort scripts of an application version.
//
// @return ListSortScriptsResponse
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

// Summary:
//
// Queries log statistics, such as application error logs, hotword rankings, and slow query logs.
//
// @param request - ListStatisticLogsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatisticLogsResponse
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

// Summary:
//
// Queries log statistics, such as application error logs, hotword rankings, and slow query logs.
//
// @param request - ListStatisticLogsRequest
//
// @return ListStatisticLogsResponse
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

// @param request - ListStatisticReportRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListStatisticReportResponse
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

// @param request - ListStatisticReportRequest
//
// @return ListStatisticReportResponse
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

// Summary:
//
// Queries tagged resources.
//
// @param tmpReq - ListTagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries tagged resources.
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
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

// Summary:
//
// Queries the entries of a custom analyzer.
//
// @param request - ListUserAnalyzerEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAnalyzerEntriesResponse
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

// Summary:
//
// Queries the entries of a custom analyzer.
//
// @param request - ListUserAnalyzerEntriesRequest
//
// @return ListUserAnalyzerEntriesResponse
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

// Summary:
//
// Queries the custom analyzers that belong to the current account.
//
// @param request - ListUserAnalyzersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserAnalyzersResponse
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

// Summary:
//
// Queries the custom analyzers that belong to the current account.
//
// @param request - ListUserAnalyzersRequest
//
// @return ListUserAnalyzersResponse
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

// Summary:
//
// Modifies the properties of an OpenSearch application or sets the online version of an OpenSearch application.
//
// @param request - ModifyAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppGroupResponse
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

// Summary:
//
// Modifies the properties of an OpenSearch application or sets the online version of an OpenSearch application.
//
// @param request - ModifyAppGroupRequest
//
// @return ModifyAppGroupResponse
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

// @param request - ModifyAppGroupQuotaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyAppGroupQuotaResponse
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

// @param request - ModifyAppGroupQuotaRequest
//
// @return ModifyAppGroupQuotaResponse
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

// Summary:
//
// Modifies a rough sort expression for an OpenSearch application. If you set dryRun to true, this operation checks the rough sort expression after the expression is modified. If you do not specify this parameter, false is used by default.
//
// @param request - ModifyFirstRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyFirstRankResponse
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

// Summary:
//
// Modifies a rough sort expression for an OpenSearch application. If you set dryRun to true, this operation checks the rough sort expression after the expression is modified. If you do not specify this parameter, false is used by default.
//
// @param request - ModifyFirstRankRequest
//
// @return ModifyFirstRankResponse
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

// Summary:
//
// Modifies a query analysis rule for a specific application version. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifyQueryProcessorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyQueryProcessorResponse
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

// Summary:
//
// Modifies a query analysis rule for a specific application version. If you set dryRun to true, this operation checks the specified query analysis rule. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifyQueryProcessorRequest
//
// @return ModifyQueryProcessorResponse
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

// Summary:
//
// Modifies a scheduled task.
//
// @param request - ModifyScheduledTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifyScheduledTaskResponse
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

// Summary:
//
// Modifies a scheduled task.
//
// @param request - ModifyScheduledTaskRequest
//
// @return ModifyScheduledTaskResponse
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

// Summary:
//
// Modifies a fine sort expression that is configured for a specific OpenSearch application version. If you set dryRun to true, the specified fine sort expression is checked after the expression is modified. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifySecondRankRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ModifySecondRankResponse
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

// Summary:
//
// Modifies a fine sort expression that is configured for a specific OpenSearch application version. If you set dryRun to true, the specified fine sort expression is checked after the expression is modified. By default, the value of dryRun is false if you do not specify this parameter.
//
// @param request - ModifySecondRankRequest
//
// @return ModifySecondRankResponse
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

// Summary:
//
// Accepts the changes in intervention entries.
//
// @param request - PushInterventionDictionaryEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushInterventionDictionaryEntriesResponse
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

// Summary:
//
// Accepts the changes in intervention entries.
//
// @param request - PushInterventionDictionaryEntriesRequest
//
// @return PushInterventionDictionaryEntriesResponse
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

// Summary:
//
// Accepts the changes in the entries of a custom analyzer.
//
// @param request - PushUserAnalyzerEntriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return PushUserAnalyzerEntriesResponse
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

// Summary:
//
// Accepts the changes in the entries of a custom analyzer.
//
// @param request - PushUserAnalyzerEntriesRequest
//
// @return PushUserAnalyzerEntriesResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReleaseSortScriptResponse
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

// @return ReleaseSortScriptResponse
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

// Summary:
//
// Deletes a version of an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppResponse
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

// Summary:
//
// Deletes a version of an OpenSearch application.
//
// @return RemoveAppResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveAppGroupResponse
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

// @return RemoveAppGroupResponse
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

// Summary:
//
// Disables data collection.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveDataCollectionResponse
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

// Summary:
//
// Disables data collection.
//
// @return RemoveDataCollectionResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFirstRankResponse
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

// @return RemoveFirstRankResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveInterventionDictionaryResponse
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

// @return RemoveInterventionDictionaryResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveQueryProcessorResponse
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

// @return RemoveQueryProcessorResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveScheduledTaskResponse
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

// @return RemoveScheduledTaskResponse
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

// Summary:
//
// Deletes a query policy.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSearchStrategyResponse
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

// Summary:
//
// Deletes a query policy.
//
// @return RemoveSearchStrategyResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveSecondRankResponse
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

// @return RemoveSecondRankResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveUserAnalyzerResponse
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

// @return RemoveUserAnalyzerResponse
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

// Summary:
//
// Renews an application. This operation is not available now. You must renew an application in the OpenSearch console.
//
// @param request - RenewAppGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RenewAppGroupResponse
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

// Summary:
//
// Renews an application. This operation is not available now. You must renew an application in the OpenSearch console.
//
// @param request - RenewAppGroupRequest
//
// @return RenewAppGroupResponse
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

// Summary:
//
// Converts a service-based application to an instance-based application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReplaceAppGroupCommodityCodeResponse
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

// Summary:
//
// Converts a service-based application to an instance-based application.
//
// @return ReplaceAppGroupCommodityCodeResponse
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

// Summary:
//
// Uploads a sort script.
//
// @param request - SaveSortScriptFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SaveSortScriptFileResponse
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

// Summary:
//
// Uploads a sort script.
//
// @param request - SaveSortScriptFileRequest
//
// @return SaveSortScriptFileResponse
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

// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StartSlowQueryAnalyzerResponse
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

// @return StartSlowQueryAnalyzerResponse
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

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
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

// Summary:
//
// Adds tags to resources.
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
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

// Summary:
//
// Unbinds a custom analyzer from an Elasticsearch instance.
//
// Description:
//
// You can call this operation to unbind a custom analyzer from an Elasticsearch instance.
//
// @param request - UnbindESUserAnalyzerRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindESUserAnalyzerResponse
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

// Summary:
//
// Unbinds a custom analyzer from an Elasticsearch instance.
//
// Description:
//
// You can call this operation to unbind a custom analyzer from an Elasticsearch instance.
//
// @param request - UnbindESUserAnalyzerRequest
//
// @return UnbindESUserAnalyzerResponse
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

// Summary:
//
// Unbinds an Elasticsearch instance from an OpenSearch application.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnbindEsInstanceResponse
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

// Summary:
//
// Unbinds an Elasticsearch instance from an OpenSearch application.
//
// @return UnbindEsInstanceResponse
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

// Summary:
//
// Remove tags from resources.
//
// @param tmpReq - UntagResourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
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

// Summary:
//
// Remove tags from resources.
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
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

// Summary:
//
// Modifies the parameters of an A/B test.
//
// @param request - UpdateABTestExperimentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestExperimentResponse
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

// Summary:
//
// Modifies the parameters of an A/B test.
//
// @param request - UpdateABTestExperimentRequest
//
// @return UpdateABTestExperimentResponse
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

// Summary:
//
// Modifies whitelists.
//
// @param request - UpdateABTestFixedFlowDividersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestFixedFlowDividersResponse
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

// Summary:
//
// Modifies whitelists.
//
// @param request - UpdateABTestFixedFlowDividersRequest
//
// @return UpdateABTestFixedFlowDividersResponse
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

// Summary:
//
// Modifies a test group.
//
// @param request - UpdateABTestGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestGroupResponse
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

// Summary:
//
// Modifies a test group.
//
// @param request - UpdateABTestGroupRequest
//
// @return UpdateABTestGroupResponse
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

// Summary:
//
// Modifies an A/B test scenario.
//
// @param request - UpdateABTestSceneRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateABTestSceneResponse
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

// Summary:
//
// Modifies an A/B test scenario.
//
// @param request - UpdateABTestSceneRequest
//
// @return UpdateABTestSceneResponse
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

// @param request - UpdateFetchFieldsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFetchFieldsResponse
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

// @param request - UpdateFetchFieldsRequest
//
// @return UpdateFetchFieldsResponse
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

// Summary:
//
// Sets the default algorithm instance used by the specified application. The new algorithm instance automatically overwrites the most recently set default instance. If no instance is set, the default instance is canceled.
//
// @param request - UpdateFunctionDefaultInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionDefaultInstanceResponse
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

// Summary:
//
// Sets the default algorithm instance used by the specified application. The new algorithm instance automatically overwrites the most recently set default instance. If no instance is set, the default instance is canceled.
//
// @param request - UpdateFunctionDefaultInstanceRequest
//
// @return UpdateFunctionDefaultInstanceResponse
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

// Summary:
//
// Updates an algorithm instance.
//
// @param request - UpdateFunctionInstanceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionInstanceResponse
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

// Summary:
//
// Updates an algorithm instance.
//
// @param request - UpdateFunctionInstanceRequest
//
// @return UpdateFunctionInstanceResponse
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

// Summary:
//
// Updates an algorithm resource.
//
// Description:
//
// You can call this operation to update the information about resources by resource name. You can modify only the values of data and description.
//
// @param request - UpdateFunctionResourceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFunctionResourceResponse
func (client *Client) UpdateFunctionResourceWithOptions(appGroupIdentity *string, functionName *string, resourceName *string, request *UpdateFunctionResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFunctionResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Data)) {
		body["Data"] = request.Data
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFunctionResource"),
		Version:     tea.String("2017-12-25"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v4/openapi/app-groups/" + tea.StringValue(openapiutil.GetEncodeParam(appGroupIdentity)) + "/functions/" + tea.StringValue(openapiutil.GetEncodeParam(functionName)) + "/resources/" + tea.StringValue(openapiutil.GetEncodeParam(resourceName))),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFunctionResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates an algorithm resource.
//
// Description:
//
// You can call this operation to update the information about resources by resource name. You can modify only the values of data and description.
//
// @param request - UpdateFunctionResourceRequest
//
// @return UpdateFunctionResourceResponse
func (client *Client) UpdateFunctionResource(appGroupIdentity *string, functionName *string, resourceName *string, request *UpdateFunctionResourceRequest) (_result *UpdateFunctionResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFunctionResourceResponse{}
	_body, _err := client.UpdateFunctionResourceWithOptions(appGroupIdentity, functionName, resourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a query policy.
//
// @param request - UpdateSearchStrategyRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSearchStrategyResponse
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

// Summary:
//
// Modifies a query policy.
//
// @param request - UpdateSearchStrategyRequest
//
// @return UpdateSearchStrategyResponse
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

// Summary:
//
// Modifies the description of a sort script.
//
// Description:
//
// You can call this operation to modify the description of a sort script.
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSortScriptResponse
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

// Summary:
//
// Modifies the description of a sort script.
//
// Description:
//
// You can call this operation to modify the description of a sort script.
//
// @return UpdateSortScriptResponse
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

// Summary:
//
// Updates summaries. A dry run is supported.
//
// @param request - UpdateSummariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateSummariesResponse
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

// Summary:
//
// Updates summaries. A dry run is supported.
//
// @param request - UpdateSummariesRequest
//
// @return UpdateSummariesResponse
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

// Summary:
//
// Verifies data sources.
//
// @param request - ValidateDataSourcesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ValidateDataSourcesResponse
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

// Summary:
//
// Verifies data sources.
//
// @param request - ValidateDataSourcesRequest
//
// @return ValidateDataSourcesResponse
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
