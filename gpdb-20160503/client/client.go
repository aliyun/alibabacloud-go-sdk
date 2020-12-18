// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type DescribeUserEncryptionKeyListRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeUserEncryptionKeyListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListRequest) SetRegionId(v string) *DescribeUserEncryptionKeyListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageNumber(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeUserEncryptionKeyListRequest) SetPageSize(v string) *DescribeUserEncryptionKeyListRequest {
	s.PageSize = &v
	return s
}

type DescribeUserEncryptionKeyListResponse struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	KmsKeys   []*DescribeUserEncryptionKeyListResponseKmsKeys `json:"KmsKeys,omitempty" xml:"KmsKeys,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeUserEncryptionKeyListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponse) SetRequestId(v string) *DescribeUserEncryptionKeyListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponse) SetKmsKeys(v []*DescribeUserEncryptionKeyListResponseKmsKeys) *DescribeUserEncryptionKeyListResponse {
	s.KmsKeys = v
	return s
}

type DescribeUserEncryptionKeyListResponseKmsKeys struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty" require:"true"`
}

func (s DescribeUserEncryptionKeyListResponseKmsKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseKmsKeys) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseKmsKeys) SetKeyId(v string) *DescribeUserEncryptionKeyListResponseKmsKeys {
	s.KeyId = &v
	return s
}

type DescribeModifyParameterLogRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s DescribeModifyParameterLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogRequest) SetDBInstanceId(v string) *DescribeModifyParameterLogRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetStartTime(v string) *DescribeModifyParameterLogRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeModifyParameterLogRequest) SetEndTime(v string) *DescribeModifyParameterLogRequest {
	s.EndTime = &v
	return s
}

type DescribeModifyParameterLogResponse struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Changelogs []*DescribeModifyParameterLogResponseChangelogs `json:"Changelogs,omitempty" xml:"Changelogs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeModifyParameterLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponse) SetRequestId(v string) *DescribeModifyParameterLogResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeModifyParameterLogResponse) SetChangelogs(v []*DescribeModifyParameterLogResponseChangelogs) *DescribeModifyParameterLogResponse {
	s.Changelogs = v
	return s
}

type DescribeModifyParameterLogResponseChangelogs struct {
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty" require:"true"`
	ParameterValueBefore *string `json:"ParameterValueBefore,omitempty" xml:"ParameterValueBefore,omitempty" require:"true"`
	ParameterValueAfter  *string `json:"ParameterValueAfter,omitempty" xml:"ParameterValueAfter,omitempty" require:"true"`
	ParameterValid       *string `json:"ParameterValid,omitempty" xml:"ParameterValid,omitempty" require:"true"`
	EffectTime           *string `json:"EffectTime,omitempty" xml:"EffectTime,omitempty" require:"true"`
}

func (s DescribeModifyParameterLogResponseChangelogs) String() string {
	return tea.Prettify(s)
}

func (s DescribeModifyParameterLogResponseChangelogs) GoString() string {
	return s.String()
}

func (s *DescribeModifyParameterLogResponseChangelogs) SetParameterName(v string) *DescribeModifyParameterLogResponseChangelogs {
	s.ParameterName = &v
	return s
}

func (s *DescribeModifyParameterLogResponseChangelogs) SetParameterValueBefore(v string) *DescribeModifyParameterLogResponseChangelogs {
	s.ParameterValueBefore = &v
	return s
}

func (s *DescribeModifyParameterLogResponseChangelogs) SetParameterValueAfter(v string) *DescribeModifyParameterLogResponseChangelogs {
	s.ParameterValueAfter = &v
	return s
}

func (s *DescribeModifyParameterLogResponseChangelogs) SetParameterValid(v string) *DescribeModifyParameterLogResponseChangelogs {
	s.ParameterValid = &v
	return s
}

func (s *DescribeModifyParameterLogResponseChangelogs) SetEffectTime(v string) *DescribeModifyParameterLogResponseChangelogs {
	s.EffectTime = &v
	return s
}

type DescribeParametersRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDBInstanceId(v string) *DescribeParametersRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeParametersResponse struct {
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Parameters []*DescribeParametersResponseParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetRequestId(v string) *DescribeParametersResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersResponse) SetParameters(v []*DescribeParametersResponseParameters) *DescribeParametersResponse {
	s.Parameters = v
	return s
}

type DescribeParametersResponseParameters struct {
	ParameterName        *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty" require:"true"`
	ParameterValue       *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty" require:"true"`
	CurrentValue         *string `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty" require:"true"`
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty" require:"true"`
	ForceRestartInstance *string `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty" require:"true"`
	IsChangeableConfig   *string `json:"IsChangeableConfig,omitempty" xml:"IsChangeableConfig,omitempty" require:"true"`
	OptionalRange        *string `json:"OptionalRange,omitempty" xml:"OptionalRange,omitempty" require:"true"`
}

func (s DescribeParametersResponseParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseParameters) SetParameterName(v string) *DescribeParametersResponseParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetParameterValue(v string) *DescribeParametersResponseParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetCurrentValue(v string) *DescribeParametersResponseParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetParameterDescription(v string) *DescribeParametersResponseParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetForceRestartInstance(v string) *DescribeParametersResponseParameters {
	s.ForceRestartInstance = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetIsChangeableConfig(v string) *DescribeParametersResponseParameters {
	s.IsChangeableConfig = &v
	return s
}

func (s *DescribeParametersResponseParameters) SetOptionalRange(v string) *DescribeParametersResponseParameters {
	s.OptionalRange = &v
	return s
}

type ModifyParametersRequest struct {
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty" require:"true"`
	ForceRestartInstance *bool   `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParametersRequest) SetForceRestartInstance(v bool) *ModifyParametersRequest {
	s.ForceRestartInstance = &v
	return s
}

type ModifyParametersResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) SetRequestId(v string) *ModifyParametersResponse {
	s.RequestId = &v
	return s
}

type CreateServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleRequest) SetRegionId(v string) *CreateServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CreateServiceLinkedRoleResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponse) SetRequestId(v string) *CreateServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

type CheckServiceLinkedRoleRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckServiceLinkedRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleRequest) SetRegionId(v string) *CheckServiceLinkedRoleRequest {
	s.RegionId = &v
	return s
}

type CheckServiceLinkedRoleResponse struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty" require:"true"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
}

func (s CheckServiceLinkedRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckServiceLinkedRoleResponse) GoString() string {
	return s.String()
}

func (s *CheckServiceLinkedRoleResponse) SetRequestId(v string) *CheckServiceLinkedRoleResponse {
	s.RequestId = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetHasServiceLinkedRole(v string) *CheckServiceLinkedRoleResponse {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckServiceLinkedRoleResponse) SetRegionId(v string) *CheckServiceLinkedRoleResponse {
	s.RegionId = &v
	return s
}

type DescribeSQLLogCountRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s DescribeSQLLogCountRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountRequest) SetDBInstanceId(v string) *DescribeSQLLogCountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetQueryKeywords(v string) *DescribeSQLLogCountRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetStartTime(v string) *DescribeSQLLogCountRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetDatabase(v string) *DescribeSQLLogCountRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetUser(v string) *DescribeSQLLogCountRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetEndTime(v string) *DescribeSQLLogCountRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteCost(v string) *DescribeSQLLogCountRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetSourceIP(v string) *DescribeSQLLogCountRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetExecuteState(v string) *DescribeSQLLogCountRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationClass(v string) *DescribeSQLLogCountRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogCountRequest) SetOperationType(v string) *DescribeSQLLogCountRequest {
	s.OperationType = &v
	return s
}

type DescribeSQLLogCountResponse struct {
	RequestId   *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBClusterId *string                             `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty" require:"true"`
	StartTime   *string                             `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime     *string                             `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	Items       []*DescribeSQLLogCountResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogCountResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponse) SetRequestId(v string) *DescribeSQLLogCountResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetDBClusterId(v string) *DescribeSQLLogCountResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetStartTime(v string) *DescribeSQLLogCountResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetEndTime(v string) *DescribeSQLLogCountResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogCountResponse) SetItems(v []*DescribeSQLLogCountResponseItems) *DescribeSQLLogCountResponse {
	s.Items = v
	return s
}

type DescribeSQLLogCountResponseItems struct {
	Name   *string                                   `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Series []*DescribeSQLLogCountResponseItemsSeries `json:"Series,omitempty" xml:"Series,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseItems) SetName(v string) *DescribeSQLLogCountResponseItems {
	s.Name = &v
	return s
}

func (s *DescribeSQLLogCountResponseItems) SetSeries(v []*DescribeSQLLogCountResponseItemsSeries) *DescribeSQLLogCountResponseItems {
	s.Series = v
	return s
}

type DescribeSQLLogCountResponseItemsSeries struct {
	Values []*DescribeSQLLogCountResponseItemsSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseItemsSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseItemsSeries) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseItemsSeries) SetValues(v []*DescribeSQLLogCountResponseItemsSeriesValues) *DescribeSQLLogCountResponseItemsSeries {
	s.Values = v
	return s
}

type DescribeSQLLogCountResponseItemsSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogCountResponseItemsSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogCountResponseItemsSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogCountResponseItemsSeriesValues) SetPoint(v []*string) *DescribeSQLLogCountResponseItemsSeriesValues {
	s.Point = v
	return s
}

type DescribeSQLLogsRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	QueryKeywords  *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Database       *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User           *string `json:"User,omitempty" xml:"User,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize       *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ExecuteCost    *string `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty"`
	SourceIP       *string `json:"SourceIP,omitempty" xml:"SourceIP,omitempty"`
	ExecuteState   *string `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty"`
	OperationClass *string `json:"OperationClass,omitempty" xml:"OperationClass,omitempty"`
	OperationType  *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s DescribeSQLLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsRequest) SetDBInstanceId(v string) *DescribeSQLLogsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetQueryKeywords(v string) *DescribeSQLLogsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetStartTime(v string) *DescribeSQLLogsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetDatabase(v string) *DescribeSQLLogsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetUser(v string) *DescribeSQLLogsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetEndTime(v string) *DescribeSQLLogsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageSize(v int) *DescribeSQLLogsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetPageNumber(v int) *DescribeSQLLogsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteCost(v string) *DescribeSQLLogsRequest {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetSourceIP(v string) *DescribeSQLLogsRequest {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetExecuteState(v string) *DescribeSQLLogsRequest {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationClass(v string) *DescribeSQLLogsRequest {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsRequest) SetOperationType(v string) *DescribeSQLLogsRequest {
	s.OperationType = &v
	return s
}

type DescribeSQLLogsResponse struct {
	RequestId       *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber      *int                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageRecordCount *int                            `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty" require:"true"`
	Items           []*DescribeSQLLogsResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponse) SetRequestId(v string) *DescribeSQLLogsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogsResponse) SetPageNumber(v int) *DescribeSQLLogsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogsResponse) SetPageRecordCount(v int) *DescribeSQLLogsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogsResponse) SetItems(v []*DescribeSQLLogsResponseItems) *DescribeSQLLogsResponse {
	s.Items = v
	return s
}

type DescribeSQLLogsResponseItems struct {
	DBName               *string  `json:"DBName,omitempty" xml:"DBName,omitempty" require:"true"`
	AccountName          *string  `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	OperationExecuteTime *string  `json:"OperationExecuteTime,omitempty" xml:"OperationExecuteTime,omitempty" require:"true"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty" require:"true"`
	ReturnRowCounts      *int64   `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty" require:"true"`
	ExecuteCost          *float32 `json:"ExecuteCost,omitempty" xml:"ExecuteCost,omitempty" require:"true"`
	DBRole               *string  `json:"DBRole,omitempty" xml:"DBRole,omitempty" require:"true"`
	SourceIP             *string  `json:"SourceIP,omitempty" xml:"SourceIP,omitempty" require:"true"`
	SourcePort           *int     `json:"SourcePort,omitempty" xml:"SourcePort,omitempty" require:"true"`
	ExecuteState         *string  `json:"ExecuteState,omitempty" xml:"ExecuteState,omitempty" require:"true"`
	OperationClass       *string  `json:"OperationClass,omitempty" xml:"OperationClass,omitempty" require:"true"`
	OperationType        *string  `json:"OperationType,omitempty" xml:"OperationType,omitempty" require:"true"`
	ScanRowCounts        *int64   `json:"ScanRowCounts,omitempty" xml:"ScanRowCounts,omitempty" require:"true"`
	SQLPlan              *string  `json:"SQLPlan,omitempty" xml:"SQLPlan,omitempty" require:"true"`
}

func (s DescribeSQLLogsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogsResponseItems) SetDBName(v string) *DescribeSQLLogsResponseItems {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetAccountName(v string) *DescribeSQLLogsResponseItems {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetOperationExecuteTime(v string) *DescribeSQLLogsResponseItems {
	s.OperationExecuteTime = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetSQLText(v string) *DescribeSQLLogsResponseItems {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetReturnRowCounts(v int64) *DescribeSQLLogsResponseItems {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetExecuteCost(v float32) *DescribeSQLLogsResponseItems {
	s.ExecuteCost = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetDBRole(v string) *DescribeSQLLogsResponseItems {
	s.DBRole = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetSourceIP(v string) *DescribeSQLLogsResponseItems {
	s.SourceIP = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetSourcePort(v int) *DescribeSQLLogsResponseItems {
	s.SourcePort = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetExecuteState(v string) *DescribeSQLLogsResponseItems {
	s.ExecuteState = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetOperationClass(v string) *DescribeSQLLogsResponseItems {
	s.OperationClass = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetOperationType(v string) *DescribeSQLLogsResponseItems {
	s.OperationType = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetScanRowCounts(v int64) *DescribeSQLLogsResponseItems {
	s.ScanRowCounts = &v
	return s
}

func (s *DescribeSQLLogsResponseItems) SetSQLPlan(v string) *DescribeSQLLogsResponseItems {
	s.SQLPlan = &v
	return s
}

type CreateECSDBInstanceRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" require:"true"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	InstanceSpec          *string `json:"InstanceSpec,omitempty" xml:"InstanceSpec,omitempty" require:"true"`
	SegNodeNum            *int    `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty" require:"true"`
	SegStorageType        *string `json:"SegStorageType,omitempty" xml:"SegStorageType,omitempty" require:"true"`
	StorageSize           *int    `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" require:"true"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime              *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress      *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	EncryptionKey         *string `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty"`
	EncryptionType        *string `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty"`
}

func (s CreateECSDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceRequest) SetOwnerId(v int64) *CreateECSDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetRegionId(v string) *CreateECSDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetZoneId(v string) *CreateECSDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngineVersion(v string) *CreateECSDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEngine(v string) *CreateECSDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceSpec(v string) *CreateECSDBInstanceRequest {
	s.InstanceSpec = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegNodeNum(v int) *CreateECSDBInstanceRequest {
	s.SegNodeNum = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSegStorageType(v string) *CreateECSDBInstanceRequest {
	s.SegStorageType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetStorageSize(v int) *CreateECSDBInstanceRequest {
	s.StorageSize = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetDBInstanceDescription(v string) *CreateECSDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetSecurityIPList(v string) *CreateECSDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPayType(v string) *CreateECSDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPeriod(v string) *CreateECSDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetUsedTime(v string) *CreateECSDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetClientToken(v string) *CreateECSDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetInstanceNetworkType(v string) *CreateECSDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVPCId(v string) *CreateECSDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetVSwitchId(v string) *CreateECSDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetPrivateIpAddress(v string) *CreateECSDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionKey(v string) *CreateECSDBInstanceRequest {
	s.EncryptionKey = &v
	return s
}

func (s *CreateECSDBInstanceRequest) SetEncryptionType(v string) *CreateECSDBInstanceRequest {
	s.EncryptionType = &v
	return s
}

type CreateECSDBInstanceResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty" require:"true"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
}

func (s CreateECSDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateECSDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateECSDBInstanceResponse) SetRequestId(v string) *CreateECSDBInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateECSDBInstanceResponse) SetDBInstanceId(v string) *CreateECSDBInstanceResponse {
	s.DBInstanceId = &v
	return s
}

func (s *CreateECSDBInstanceResponse) SetOrderId(v string) *CreateECSDBInstanceResponse {
	s.OrderId = &v
	return s
}

func (s *CreateECSDBInstanceResponse) SetConnectionString(v string) *CreateECSDBInstanceResponse {
	s.ConnectionString = &v
	return s
}

func (s *CreateECSDBInstanceResponse) SetPort(v string) *CreateECSDBInstanceResponse {
	s.Port = &v
	return s
}

type DescribeDBClusterPerformanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeDBClusterPerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceRequest) SetDBInstanceId(v string) *DescribeDBClusterPerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetKey(v string) *DescribeDBClusterPerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetStartTime(v string) *DescribeDBClusterPerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceRequest) SetEndTime(v string) *DescribeDBClusterPerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBClusterPerformanceResponse struct {
	RequestId       *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBClusterId     *string                                                `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty" require:"true"`
	StartTime       *string                                                `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string                                                `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PerformanceKeys []*DescribeDBClusterPerformanceResponsePerformanceKeys `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponse) SetRequestId(v string) *DescribeDBClusterPerformanceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetDBClusterId(v string) *DescribeDBClusterPerformanceResponse {
	s.DBClusterId = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetStartTime(v string) *DescribeDBClusterPerformanceResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetEndTime(v string) *DescribeDBClusterPerformanceResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponse) SetPerformanceKeys(v []*DescribeDBClusterPerformanceResponsePerformanceKeys) *DescribeDBClusterPerformanceResponse {
	s.PerformanceKeys = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeys struct {
	Unit   *string                                                      `json:"Unit,omitempty" xml:"Unit,omitempty" require:"true"`
	Name   *string                                                      `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Series []*DescribeDBClusterPerformanceResponsePerformanceKeysSeries `json:"Series,omitempty" xml:"Series,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeys) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeys) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeys) SetUnit(v string) *DescribeDBClusterPerformanceResponsePerformanceKeys {
	s.Unit = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeys) SetName(v string) *DescribeDBClusterPerformanceResponsePerformanceKeys {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeys) SetSeries(v []*DescribeDBClusterPerformanceResponsePerformanceKeysSeries) *DescribeDBClusterPerformanceResponsePerformanceKeys {
	s.Series = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeysSeries struct {
	Name   *string                                                            `json:"Name,omitempty" xml:"Name,omitempty" require:"true"`
	Role   *string                                                            `json:"Role,omitempty" xml:"Role,omitempty" require:"true"`
	Values []*DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues `json:"Values,omitempty" xml:"Values,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysSeries) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysSeries) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysSeries) SetName(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysSeries {
	s.Name = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysSeries) SetRole(v string) *DescribeDBClusterPerformanceResponsePerformanceKeysSeries {
	s.Role = &v
	return s
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysSeries) SetValues(v []*DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues) *DescribeDBClusterPerformanceResponsePerformanceKeysSeries {
	s.Values = v
	return s
}

type DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues struct {
	Point []*string `json:"Point,omitempty" xml:"Point,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues) SetPoint(v []*string) *DescribeDBClusterPerformanceResponsePerformanceKeysSeriesValues {
	s.Point = v
	return s
}

type DescribeDBInstanceOnECSAttributeRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceOnECSAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceOnECSAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceOnECSAttributeResponse struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Items     *DescribeDBInstanceOnECSAttributeResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceOnECSAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetRequestId(v string) *DescribeDBInstanceOnECSAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponse) SetItems(v *DescribeDBInstanceOnECSAttributeResponseItems) *DescribeDBInstanceOnECSAttributeResponse {
	s.Items = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseItems struct {
	DBInstanceAttribute []*DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseItems) SetDBInstanceAttribute(v []*DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) *DescribeDBInstanceOnECSAttributeResponseItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute struct {
	DBInstanceId          *string                                                               `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	PayType               *string                                                               `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	RegionId              *string                                                               `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Engine                *string                                                               `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	EngineVersion         *string                                                               `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" require:"true"`
	DBInstanceClass       *string                                                               `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" require:"true"`
	DBInstanceStatus      *string                                                               `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty" require:"true"`
	DBInstanceDescription *string                                                               `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty" require:"true"`
	LockMode              *string                                                               `json:"LockMode,omitempty" xml:"LockMode,omitempty" require:"true"`
	CreationTime          *string                                                               `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ExpireTime            *string                                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	ZoneId                *string                                                               `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	InstanceNetworkType   *string                                                               `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty" require:"true"`
	VpcId                 *string                                                               `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	ConnectionMode        *string                                                               `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty" require:"true"`
	StorageType           *string                                                               `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	InstanceDeployType    *string                                                               `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty" require:"true"`
	SegNodeNum            *int                                                                  `json:"SegNodeNum,omitempty" xml:"SegNodeNum,omitempty" require:"true"`
	MemorySize            *int                                                                  `json:"MemorySize,omitempty" xml:"MemorySize,omitempty" require:"true"`
	CpuCores              *int                                                                  `json:"CpuCores,omitempty" xml:"CpuCores,omitempty" require:"true"`
	StorageSize           *int                                                                  `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" require:"true"`
	VSwitchId             *string                                                               `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	ConnectionString      *string                                                               `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty" require:"true"`
	Port                  *string                                                               `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	EncryptionType        *string                                                               `json:"EncryptionType,omitempty" xml:"EncryptionType,omitempty" require:"true"`
	EncryptionKey         *string                                                               `json:"EncryptionKey,omitempty" xml:"EncryptionKey,omitempty" require:"true"`
	Tags                  *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetInstanceDeployType(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetSegNodeNum(v int) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.SegNodeNum = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetMemorySize(v int) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.MemorySize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetCpuCores(v int) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.CpuCores = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetStorageSize(v int) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.StorageSize = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetVSwitchId(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetEncryptionType(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.EncryptionType = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetEncryptionKey(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.EncryptionKey = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceOnECSAttributeResponseItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeAvailableResourcesRequest struct {
	Region     *string `json:"Region,omitempty" xml:"Region,omitempty" require:"true"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
}

func (s DescribeAvailableResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesRequest) SetRegion(v string) *DescribeAvailableResourcesRequest {
	s.Region = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetZoneId(v string) *DescribeAvailableResourcesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesRequest) SetChargeType(v string) *DescribeAvailableResourcesRequest {
	s.ChargeType = &v
	return s
}

type DescribeAvailableResourcesResponse struct {
	RequestId *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	RegionId  *string                                        `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Resources []*DescribeAvailableResourcesResponseResources `json:"Resources,omitempty" xml:"Resources,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponse) SetRequestId(v string) *DescribeAvailableResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetRegionId(v string) *DescribeAvailableResourcesResponse {
	s.RegionId = &v
	return s
}

func (s *DescribeAvailableResourcesResponse) SetResources(v []*DescribeAvailableResourcesResponseResources) *DescribeAvailableResourcesResponse {
	s.Resources = v
	return s
}

type DescribeAvailableResourcesResponseResources struct {
	ZoneId           *string                                                        `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	SupportedEngines []*DescribeAvailableResourcesResponseResourcesSupportedEngines `json:"SupportedEngines,omitempty" xml:"SupportedEngines,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseResources) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseResources) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseResources) SetZoneId(v string) *DescribeAvailableResourcesResponseResources {
	s.ZoneId = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResources) SetSupportedEngines(v []*DescribeAvailableResourcesResponseResourcesSupportedEngines) *DescribeAvailableResourcesResponseResources {
	s.SupportedEngines = v
	return s
}

type DescribeAvailableResourcesResponseResourcesSupportedEngines struct {
	SupportedEngineVersion   *string                                                                                `json:"SupportedEngineVersion,omitempty" xml:"SupportedEngineVersion,omitempty" require:"true"`
	Mode                     *string                                                                                `json:"Mode,omitempty" xml:"Mode,omitempty" require:"true"`
	SupportedInstanceClasses []*DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses `json:"SupportedInstanceClasses,omitempty" xml:"SupportedInstanceClasses,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEngines) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEngines) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEngines) SetSupportedEngineVersion(v string) *DescribeAvailableResourcesResponseResourcesSupportedEngines {
	s.SupportedEngineVersion = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEngines) SetMode(v string) *DescribeAvailableResourcesResponseResourcesSupportedEngines {
	s.Mode = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEngines) SetSupportedInstanceClasses(v []*DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) *DescribeAvailableResourcesResponseResourcesSupportedEngines {
	s.SupportedInstanceClasses = v
	return s
}

type DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses struct {
	DisplayClass  *string                                                                                         `json:"DisplayClass,omitempty" xml:"DisplayClass,omitempty" require:"true"`
	InstanceClass *string                                                                                         `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty" require:"true"`
	Description   *string                                                                                         `json:"Description,omitempty" xml:"Description,omitempty" require:"true"`
	StorageType   *string                                                                                         `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	NodeCount     *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount   `json:"NodeCount,omitempty" xml:"NodeCount,omitempty" require:"true" type:"Struct"`
	StorageSize   *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize `json:"StorageSize,omitempty" xml:"StorageSize,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetDisplayClass(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.DisplayClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetInstanceClass(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.InstanceClass = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetDescription(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.Description = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetStorageType(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageType = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetNodeCount(v *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.NodeCount = v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses) SetStorageSize(v *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClasses {
	s.StorageSize = v
	return s
}

type DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount struct {
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty" require:"true"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty" require:"true"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMinCount(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetMaxCount(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount) SetStep(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesNodeCount {
	s.Step = &v
	return s
}

type DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize struct {
	MinCount *string `json:"MinCount,omitempty" xml:"MinCount,omitempty" require:"true"`
	MaxCount *string `json:"MaxCount,omitempty" xml:"MaxCount,omitempty" require:"true"`
	Step     *string `json:"Step,omitempty" xml:"Step,omitempty" require:"true"`
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMinCount(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MinCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetMaxCount(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.MaxCount = &v
	return s
}

func (s *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize) SetStep(v string) *DescribeAvailableResourcesResponseResourcesSupportedEnginesSupportedInstanceClassesStorageSize {
	s.Step = &v
	return s
}

type DescribeDBInstanceSSLRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLRequest) SetDBInstanceId(v string) *DescribeDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceSSLResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty" require:"true"`
	CertCommonName *string `json:"CertCommonName,omitempty" xml:"CertCommonName,omitempty" require:"true"`
	SSLExpiredTime *string `json:"SSLExpiredTime,omitempty" xml:"SSLExpiredTime,omitempty" require:"true"`
	SSLEnabled     *bool   `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty" require:"true"`
}

func (s DescribeDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceSSLResponse) SetRequestId(v string) *DescribeDBInstanceSSLResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetDBInstanceId(v string) *DescribeDBInstanceSSLResponse {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetDBInstanceName(v string) *DescribeDBInstanceSSLResponse {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetCertCommonName(v string) *DescribeDBInstanceSSLResponse {
	s.CertCommonName = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetSSLExpiredTime(v string) *DescribeDBInstanceSSLResponse {
	s.SSLExpiredTime = &v
	return s
}

func (s *DescribeDBInstanceSSLResponse) SetSSLEnabled(v bool) *DescribeDBInstanceSSLResponse {
	s.SSLEnabled = &v
	return s
}

type ModifyDBInstanceSSLRequest struct {
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	SSLEnabled       *int    `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty" require:"true"`
}

func (s ModifyDBInstanceSSLRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLRequest) SetDBInstanceId(v string) *ModifyDBInstanceSSLRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetConnectionString(v string) *ModifyDBInstanceSSLRequest {
	s.ConnectionString = &v
	return s
}

func (s *ModifyDBInstanceSSLRequest) SetSSLEnabled(v int) *ModifyDBInstanceSSLRequest {
	s.SSLEnabled = &v
	return s
}

type ModifyDBInstanceSSLResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceSSLResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceSSLResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceSSLResponse) SetRequestId(v string) *ModifyDBInstanceSSLResponse {
	s.RequestId = &v
	return s
}

type DescribeTagsRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
}

func (s DescribeTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTagsRequest) SetOwnerId(v int64) *DescribeTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerAccount(v string) *DescribeTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceOwnerId(v int64) *DescribeTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTagsRequest) SetOwnerAccount(v string) *DescribeTagsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTagsRequest) SetRegionId(v string) *DescribeTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTagsRequest) SetResourceType(v string) *DescribeTagsRequest {
	s.ResourceType = &v
	return s
}

type DescribeTagsResponse struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Tags      []*DescribeTagsResponseTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponse) SetRequestId(v string) *DescribeTagsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponse) SetTags(v []*DescribeTagsResponseTags) *DescribeTagsResponse {
	s.Tags = v
	return s
}

type DescribeTagsResponseTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
}

func (s DescribeTagsResponseTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeTagsResponseTags) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseTags) SetTagKey(v string) *DescribeTagsResponseTags {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseTags) SetTagValue(v string) *DescribeTagsResponseTags {
	s.TagValue = &v
	return s
}

type DescribeSpecificationRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	StorageType  *string `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	CpuCores     *int    `json:"CpuCores,omitempty" xml:"CpuCores,omitempty" require:"true"`
	TotalNodeNum *int    `json:"TotalNodeNum,omitempty" xml:"TotalNodeNum,omitempty" require:"true"`
}

func (s DescribeSpecificationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationRequest) SetOwnerId(v int64) *DescribeSpecificationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetDBInstanceId(v string) *DescribeSpecificationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSpecificationRequest) SetStorageType(v string) *DescribeSpecificationRequest {
	s.StorageType = &v
	return s
}

func (s *DescribeSpecificationRequest) SetCpuCores(v int) *DescribeSpecificationRequest {
	s.CpuCores = &v
	return s
}

func (s *DescribeSpecificationRequest) SetTotalNodeNum(v int) *DescribeSpecificationRequest {
	s.TotalNodeNum = &v
	return s
}

type DescribeSpecificationResponse struct {
	RequestId            *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceClass      []*DescribeSpecificationResponseDBInstanceClass      `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" require:"true" type:"Repeated"`
	DBInstanceGroupCount []*DescribeSpecificationResponseDBInstanceGroupCount `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" require:"true" type:"Repeated"`
	StorageNotice        []*DescribeSpecificationResponseStorageNotice        `json:"StorageNotice,omitempty" xml:"StorageNotice,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSpecificationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponse) SetRequestId(v string) *DescribeSpecificationResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSpecificationResponse) SetDBInstanceClass(v []*DescribeSpecificationResponseDBInstanceClass) *DescribeSpecificationResponse {
	s.DBInstanceClass = v
	return s
}

func (s *DescribeSpecificationResponse) SetDBInstanceGroupCount(v []*DescribeSpecificationResponseDBInstanceGroupCount) *DescribeSpecificationResponse {
	s.DBInstanceGroupCount = v
	return s
}

func (s *DescribeSpecificationResponse) SetStorageNotice(v []*DescribeSpecificationResponseStorageNotice) *DescribeSpecificationResponse {
	s.StorageNotice = v
	return s
}

type DescribeSpecificationResponseDBInstanceClass struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeSpecificationResponseDBInstanceClass) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseDBInstanceClass) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseDBInstanceClass) SetText(v string) *DescribeSpecificationResponseDBInstanceClass {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseDBInstanceClass) SetValue(v string) *DescribeSpecificationResponseDBInstanceClass {
	s.Value = &v
	return s
}

type DescribeSpecificationResponseDBInstanceGroupCount struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeSpecificationResponseDBInstanceGroupCount) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseDBInstanceGroupCount) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseDBInstanceGroupCount) SetText(v string) *DescribeSpecificationResponseDBInstanceGroupCount {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseDBInstanceGroupCount) SetValue(v string) *DescribeSpecificationResponseDBInstanceGroupCount {
	s.Value = &v
	return s
}

type DescribeSpecificationResponseStorageNotice struct {
	Text  *string `json:"Text,omitempty" xml:"Text,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeSpecificationResponseStorageNotice) String() string {
	return tea.Prettify(s)
}

func (s DescribeSpecificationResponseStorageNotice) GoString() string {
	return s.String()
}

func (s *DescribeSpecificationResponseStorageNotice) SetText(v string) *DescribeSpecificationResponseStorageNotice {
	s.Text = &v
	return s
}

func (s *DescribeSpecificationResponseStorageNotice) SetValue(v string) *DescribeSpecificationResponseStorageNotice {
	s.Value = &v
	return s
}

type UpgradeDBVersionRequest struct {
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	MinorVersion   *string `json:"MinorVersion,omitempty" xml:"MinorVersion,omitempty"`
	MajorVersion   *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	SwitchTimeMode *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
	SwitchTime     *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
}

func (s UpgradeDBVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionRequest) SetOwnerId(v int64) *UpgradeDBVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetRegionId(v string) *UpgradeDBVersionRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetDBInstanceId(v string) *UpgradeDBVersionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMinorVersion(v string) *UpgradeDBVersionRequest {
	s.MinorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetMajorVersion(v string) *UpgradeDBVersionRequest {
	s.MajorVersion = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTimeMode(v string) *UpgradeDBVersionRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *UpgradeDBVersionRequest) SetSwitchTime(v string) *UpgradeDBVersionRequest {
	s.SwitchTime = &v
	return s
}

type UpgradeDBVersionResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty" require:"true"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty" require:"true"`
}

func (s UpgradeDBVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBVersionResponse) SetRequestId(v string) *UpgradeDBVersionResponse {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetDBInstanceId(v string) *UpgradeDBVersionResponse {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetDBInstanceName(v string) *UpgradeDBVersionResponse {
	s.DBInstanceName = &v
	return s
}

func (s *UpgradeDBVersionResponse) SetTaskId(v string) *UpgradeDBVersionResponse {
	s.TaskId = &v
	return s
}

type UpgradeDBInstanceRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DBInstanceClass      *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" require:"true"`
	DBInstanceGroupCount *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" require:"true"`
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	PayType              *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
}

func (s UpgradeDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceRequest) SetOwnerId(v int64) *UpgradeDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetRegionId(v string) *UpgradeDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceClass(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceGroupCount(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetDBInstanceId(v string) *UpgradeDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceRequest) SetPayType(v string) *UpgradeDBInstanceRequest {
	s.PayType = &v
	return s
}

type UpgradeDBInstanceResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	OrderId      *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
}

func (s UpgradeDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceResponse) SetRequestId(v string) *UpgradeDBInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceResponse) SetDBInstanceId(v string) *UpgradeDBInstanceResponse {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceResponse) SetOrderId(v string) *UpgradeDBInstanceResponse {
	s.OrderId = &v
	return s
}

type UntagResourcesRequest struct {
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerId(v int64) *UntagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

type UntagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetRequestId(v string) *UntagResourcesResponse {
	s.RequestId = &v
	return s
}

type TagResourcesRequest struct {
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true" type:"Repeated"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerId(v int64) *TagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type TagResourcesResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetRequestId(v string) *TagResourcesResponse {
	s.RequestId = &v
	return s
}

type ListTagResourcesRequest struct {
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                        `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerId(v int64) *ListTagResourcesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

type ListTagResourcesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
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

type ListTagResourcesResponse struct {
	RequestId    *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	NextToken    *string                               `json:"NextToken,omitempty" xml:"NextToken,omitempty" require:"true"`
	TagResources *ListTagResourcesResponseTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" require:"true" type:"Struct"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetRequestId(v string) *ListTagResourcesResponse {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponse) SetNextToken(v string) *ListTagResourcesResponse {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponse) SetTagResources(v *ListTagResourcesResponseTagResources) *ListTagResourcesResponse {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseTagResources struct {
	TagResource []*ListTagResourcesResponseTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" require:"true" type:"Repeated"`
}

func (s ListTagResourcesResponseTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResources) SetTagResource(v []*ListTagResourcesResponseTagResourcesTagResource) *ListTagResourcesResponseTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseTagResourcesTagResource struct {
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty" require:"true"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty" require:"true"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty" require:"true"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" require:"true"`
}

func (s ListTagResourcesResponseTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

type DescribeRdsVSwitchsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVSwitchsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsRequest) SetSecurityToken(v string) *DescribeRdsVSwitchsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetResourceOwnerId(v int64) *DescribeRdsVSwitchsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetOwnerAccount(v string) *DescribeRdsVSwitchsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetVpcId(v string) *DescribeRdsVSwitchsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVSwitchsRequest) SetZoneId(v string) *DescribeRdsVSwitchsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVSwitchsResponse struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	VSwitches *DescribeRdsVSwitchsResponseVSwitches `json:"VSwitches,omitempty" xml:"VSwitches,omitempty" require:"true" type:"Struct"`
}

func (s DescribeRdsVSwitchsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponse) SetRequestId(v string) *DescribeRdsVSwitchsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponse) SetVSwitches(v *DescribeRdsVSwitchsResponseVSwitches) *DescribeRdsVSwitchsResponse {
	s.VSwitches = v
	return s
}

type DescribeRdsVSwitchsResponseVSwitches struct {
	VSwitch []*DescribeRdsVSwitchsResponseVSwitchesVSwitch `json:"VSwitch,omitempty" xml:"VSwitch,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRdsVSwitchsResponseVSwitches) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseVSwitches) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseVSwitches) SetVSwitch(v []*DescribeRdsVSwitchsResponseVSwitchesVSwitch) *DescribeRdsVSwitchsResponseVSwitches {
	s.VSwitch = v
	return s
}

type DescribeRdsVSwitchsResponseVSwitchesVSwitch struct {
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty" require:"true"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty" require:"true"`
	Bid         *string `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	AliUid      *string `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	RegionNo    *string `json:"RegionNo,omitempty" xml:"RegionNo,omitempty" require:"true"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
}

func (s DescribeRdsVSwitchsResponseVSwitchesVSwitch) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVSwitchsResponseVSwitchesVSwitch) GoString() string {
	return s.String()
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetVSwitchId(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetVSwitchName(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.VSwitchName = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetIzNo(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetBid(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetAliUid(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetRegionNo(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetCidrBlock(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetIsDefault(v bool) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetStatus(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.Status = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetGmtCreate(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVSwitchsResponseVSwitchesVSwitch) SetGmtModified(v string) *DescribeRdsVSwitchsResponseVSwitchesVSwitch {
	s.GmtModified = &v
	return s
}

type DescribeRdsVpcsRequest struct {
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	ZoneId               *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRdsVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsRequest) SetSecurityToken(v string) *DescribeRdsVpcsRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetResourceOwnerId(v int64) *DescribeRdsVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetOwnerAccount(v string) *DescribeRdsVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRdsVpcsRequest) SetZoneId(v string) *DescribeRdsVpcsRequest {
	s.ZoneId = &v
	return s
}

type DescribeRdsVpcsResponse struct {
	RequestId *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Vpcs      *DescribeRdsVpcsResponseVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" require:"true" type:"Struct"`
}

func (s DescribeRdsVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponse) SetRequestId(v string) *DescribeRdsVpcsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRdsVpcsResponse) SetVpcs(v *DescribeRdsVpcsResponseVpcs) *DescribeRdsVpcsResponse {
	s.Vpcs = v
	return s
}

type DescribeRdsVpcsResponseVpcs struct {
	Vpc []*DescribeRdsVpcsResponseVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRdsVpcsResponseVpcs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseVpcs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseVpcs) SetVpc(v []*DescribeRdsVpcsResponseVpcsVpc) *DescribeRdsVpcsResponseVpcs {
	s.Vpc = v
	return s
}

type DescribeRdsVpcsResponseVpcsVpc struct {
	VpcId       *string                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VpcName     *string                                   `json:"VpcName,omitempty" xml:"VpcName,omitempty" require:"true"`
	Bid         *string                                   `json:"Bid,omitempty" xml:"Bid,omitempty" require:"true"`
	AliUid      *string                                   `json:"AliUid,omitempty" xml:"AliUid,omitempty" require:"true"`
	RegionNo    *string                                   `json:"RegionNo,omitempty" xml:"RegionNo,omitempty" require:"true"`
	CidrBlock   *string                                   `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	IsDefault   *bool                                     `json:"IsDefault,omitempty" xml:"IsDefault,omitempty" require:"true"`
	Status      *string                                   `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	GmtCreate   *string                                   `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	GmtModified *string                                   `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
	VSwitchs    []*DescribeRdsVpcsResponseVpcsVpcVSwitchs `json:"VSwitchs,omitempty" xml:"VSwitchs,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRdsVpcsResponseVpcsVpc) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetVpcId(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetVpcName(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetBid(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.Bid = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetAliUid(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.AliUid = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetRegionNo(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.RegionNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetCidrBlock(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetIsDefault(v bool) *DescribeRdsVpcsResponseVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetStatus(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetGmtCreate(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetGmtModified(v string) *DescribeRdsVpcsResponseVpcsVpc {
	s.GmtModified = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpc) SetVSwitchs(v []*DescribeRdsVpcsResponseVpcsVpcVSwitchs) *DescribeRdsVpcsResponseVpcsVpc {
	s.VSwitchs = v
	return s
}

type DescribeRdsVpcsResponseVpcsVpcVSwitchs struct {
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty" require:"true"`
	IzNo        *string `json:"IzNo,omitempty" xml:"IzNo,omitempty" require:"true"`
	CidrBlock   *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty" require:"true"`
	IsDefault   *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty" require:"true"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty" require:"true"`
	GmtCreate   *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty" require:"true"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty" require:"true"`
}

func (s DescribeRdsVpcsResponseVpcsVpcVSwitchs) String() string {
	return tea.Prettify(s)
}

func (s DescribeRdsVpcsResponseVpcsVpcVSwitchs) GoString() string {
	return s.String()
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetVSwitchId(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.VSwitchId = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetVSwitchName(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.VSwitchName = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetIzNo(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.IzNo = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetCidrBlock(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.CidrBlock = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetIsDefault(v bool) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.IsDefault = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetStatus(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.Status = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetGmtCreate(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.GmtCreate = &v
	return s
}

func (s *DescribeRdsVpcsResponseVpcsVpcVSwitchs) SetGmtModified(v string) *DescribeRdsVpcsResponseVpcsVpcVSwitchs {
	s.GmtModified = &v
	return s
}

type AddBuDBInstanceRelationRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	BusinessUnit *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty" require:"true"`
}

func (s AddBuDBInstanceRelationRequest) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationRequest) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationRequest) SetOwnerId(v int64) *AddBuDBInstanceRelationRequest {
	s.OwnerId = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetDBInstanceId(v string) *AddBuDBInstanceRelationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AddBuDBInstanceRelationRequest) SetBusinessUnit(v string) *AddBuDBInstanceRelationRequest {
	s.BusinessUnit = &v
	return s
}

type AddBuDBInstanceRelationResponse struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	BusinessUnit   *string `json:"BusinessUnit,omitempty" xml:"BusinessUnit,omitempty" require:"true"`
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty" require:"true"`
}

func (s AddBuDBInstanceRelationResponse) String() string {
	return tea.Prettify(s)
}

func (s AddBuDBInstanceRelationResponse) GoString() string {
	return s.String()
}

func (s *AddBuDBInstanceRelationResponse) SetRequestId(v string) *AddBuDBInstanceRelationResponse {
	s.RequestId = &v
	return s
}

func (s *AddBuDBInstanceRelationResponse) SetBusinessUnit(v string) *AddBuDBInstanceRelationResponse {
	s.BusinessUnit = &v
	return s
}

func (s *AddBuDBInstanceRelationResponse) SetDBInstanceName(v string) *AddBuDBInstanceRelationResponse {
	s.DBInstanceName = &v
	return s
}

type DescribeSQLLogRecordsRequest struct {
	DBInstanceId  *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	QueryKeywords *string `json:"QueryKeywords,omitempty" xml:"QueryKeywords,omitempty"`
	StartTime     *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	Database      *string `json:"Database,omitempty" xml:"Database,omitempty"`
	User          *string `json:"User,omitempty" xml:"User,omitempty"`
	Form          *string `json:"Form,omitempty" xml:"Form,omitempty"`
	EndTime       *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PageSize      *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber    *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSQLLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsRequest) SetDBInstanceId(v string) *DescribeSQLLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetQueryKeywords(v string) *DescribeSQLLogRecordsRequest {
	s.QueryKeywords = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetStartTime(v string) *DescribeSQLLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetDatabase(v string) *DescribeSQLLogRecordsRequest {
	s.Database = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetUser(v string) *DescribeSQLLogRecordsRequest {
	s.User = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetForm(v string) *DescribeSQLLogRecordsRequest {
	s.Form = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetEndTime(v string) *DescribeSQLLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageSize(v int) *DescribeSQLLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogRecordsRequest) SetPageNumber(v int) *DescribeSQLLogRecordsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSQLLogRecordsResponse struct {
	RequestId        *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalRecordCount *int                                `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty" require:"true"`
	PageNumber       *int                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageRecordCount  *int                                `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty" require:"true"`
	Items            *DescribeSQLLogRecordsResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSQLLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponse) SetRequestId(v string) *DescribeSQLLogRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetTotalRecordCount(v int) *DescribeSQLLogRecordsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetPageNumber(v int) *DescribeSQLLogRecordsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetPageRecordCount(v int) *DescribeSQLLogRecordsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogRecordsResponse) SetItems(v *DescribeSQLLogRecordsResponseItems) *DescribeSQLLogRecordsResponse {
	s.Items = v
	return s
}

type DescribeSQLLogRecordsResponseItems struct {
	SQLRecord []*DescribeSQLLogRecordsResponseItemsSQLRecord `json:"SQLRecord,omitempty" xml:"SQLRecord,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogRecordsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseItems) SetSQLRecord(v []*DescribeSQLLogRecordsResponseItemsSQLRecord) *DescribeSQLLogRecordsResponseItems {
	s.SQLRecord = v
	return s
}

type DescribeSQLLogRecordsResponseItemsSQLRecord struct {
	DBName              *string `json:"DBName,omitempty" xml:"DBName,omitempty" require:"true"`
	AccountName         *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	HostAddress         *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty" require:"true"`
	SQLText             *string `json:"SQLText,omitempty" xml:"SQLText,omitempty" require:"true"`
	TotalExecutionTimes *int64  `json:"TotalExecutionTimes,omitempty" xml:"TotalExecutionTimes,omitempty" require:"true"`
	ReturnRowCounts     *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty" require:"true"`
	ExecuteTime         *string `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty" require:"true"`
	ThreadID            *string `json:"ThreadID,omitempty" xml:"ThreadID,omitempty" require:"true"`
}

func (s DescribeSQLLogRecordsResponseItemsSQLRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogRecordsResponseItemsSQLRecord) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetDBName(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetAccountName(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.AccountName = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetHostAddress(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetSQLText(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetTotalExecutionTimes(v int64) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.TotalExecutionTimes = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetReturnRowCounts(v int64) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetExecuteTime(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLLogRecordsResponseItemsSQLRecord) SetThreadID(v string) *DescribeSQLLogRecordsResponseItemsSQLRecord {
	s.ThreadID = &v
	return s
}

type ModifySQLCollectorPolicyRequest struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty" require:"true"`
}

func (s ModifySQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyRequest) SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest {
	s.SQLCollectorStatus = &v
	return s
}

type ModifySQLCollectorPolicyResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifySQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyResponse) SetRequestId(v string) *ModifySQLCollectorPolicyResponse {
	s.RequestId = &v
	return s
}

type DescribeSQLLogFilesRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSQLLogFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesRequest) SetDBInstanceId(v string) *DescribeSQLLogFilesRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetFileName(v string) *DescribeSQLLogFilesRequest {
	s.FileName = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageSize(v int) *DescribeSQLLogFilesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLLogFilesRequest) SetPageNumber(v int) *DescribeSQLLogFilesRequest {
	s.PageNumber = &v
	return s
}

type DescribeSQLLogFilesResponse struct {
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	TotalRecordCount *int                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty" require:"true"`
	PageNumber       *int                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageRecordCount  *int                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty" require:"true"`
	Items            *DescribeSQLLogFilesResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSQLLogFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponse) SetRequestId(v string) *DescribeSQLLogFilesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetTotalRecordCount(v int) *DescribeSQLLogFilesResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetPageNumber(v int) *DescribeSQLLogFilesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetPageRecordCount(v int) *DescribeSQLLogFilesResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSQLLogFilesResponse) SetItems(v *DescribeSQLLogFilesResponseItems) *DescribeSQLLogFilesResponse {
	s.Items = v
	return s
}

type DescribeSQLLogFilesResponseItems struct {
	LogFile []*DescribeSQLLogFilesResponseItemsLogFile `json:"LogFile,omitempty" xml:"LogFile,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSQLLogFilesResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseItems) SetLogFile(v []*DescribeSQLLogFilesResponseItemsLogFile) *DescribeSQLLogFilesResponseItems {
	s.LogFile = v
	return s
}

type DescribeSQLLogFilesResponseItemsLogFile struct {
	FileID         *string `json:"FileID,omitempty" xml:"FileID,omitempty" require:"true"`
	LogStatus      *string `json:"LogStatus,omitempty" xml:"LogStatus,omitempty" require:"true"`
	LogDownloadURL *string `json:"LogDownloadURL,omitempty" xml:"LogDownloadURL,omitempty" require:"true"`
	LogSize        *string `json:"LogSize,omitempty" xml:"LogSize,omitempty" require:"true"`
	LogStartTime   *string `json:"LogStartTime,omitempty" xml:"LogStartTime,omitempty" require:"true"`
	LogEndTime     *string `json:"LogEndTime,omitempty" xml:"LogEndTime,omitempty" require:"true"`
}

func (s DescribeSQLLogFilesResponseItemsLogFile) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLLogFilesResponseItemsLogFile) GoString() string {
	return s.String()
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetFileID(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.FileID = &v
	return s
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetLogStatus(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.LogStatus = &v
	return s
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetLogDownloadURL(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.LogDownloadURL = &v
	return s
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetLogSize(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.LogSize = &v
	return s
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetLogStartTime(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.LogStartTime = &v
	return s
}

func (s *DescribeSQLLogFilesResponseItemsLogFile) SetLogEndTime(v string) *DescribeSQLLogFilesResponseItemsLogFile {
	s.LogEndTime = &v
	return s
}

type DescribeSQLCollectorPolicyRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeSQLCollectorPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyRequest) SetDBInstanceId(v string) *DescribeSQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeSQLCollectorPolicyResponse struct {
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty" require:"true"`
}

func (s DescribeSQLCollectorPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLCollectorPolicyResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLCollectorPolicyResponse) SetRequestId(v string) *DescribeSQLCollectorPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLCollectorPolicyResponse) SetSQLCollectorStatus(v string) *DescribeSQLCollectorPolicyResponse {
	s.SQLCollectorStatus = &v
	return s
}

type DescribeSlowLogRecordsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	SQLId        *int64  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	PageSize     *int    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber   *int    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSlowLogRecordsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsRequest) SetDBInstanceId(v string) *DescribeSlowLogRecordsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetSQLId(v int64) *DescribeSlowLogRecordsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetStartTime(v string) *DescribeSlowLogRecordsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetEndTime(v string) *DescribeSlowLogRecordsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetDBName(v string) *DescribeSlowLogRecordsRequest {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageSize(v int) *DescribeSlowLogRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowLogRecordsRequest) SetPageNumber(v int) *DescribeSlowLogRecordsRequest {
	s.PageNumber = &v
	return s
}

type DescribeSlowLogRecordsResponse struct {
	RequestId        *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Engine           *string                              `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	TotalRecordCount *int                                 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty" require:"true"`
	PageNumber       *int                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	PageRecordCount  *int                                 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty" require:"true"`
	Items            *DescribeSlowLogRecordsResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeSlowLogRecordsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponse) SetRequestId(v string) *DescribeSlowLogRecordsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetEngine(v string) *DescribeSlowLogRecordsResponse {
	s.Engine = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetTotalRecordCount(v int) *DescribeSlowLogRecordsResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetPageNumber(v int) *DescribeSlowLogRecordsResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetPageRecordCount(v int) *DescribeSlowLogRecordsResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSlowLogRecordsResponse) SetItems(v *DescribeSlowLogRecordsResponseItems) *DescribeSlowLogRecordsResponse {
	s.Items = v
	return s
}

type DescribeSlowLogRecordsResponseItems struct {
	SQLSlowRecord []*DescribeSlowLogRecordsResponseItemsSQLSlowRecord `json:"SQLSlowRecord,omitempty" xml:"SQLSlowRecord,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeSlowLogRecordsResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseItems) SetSQLSlowRecord(v []*DescribeSlowLogRecordsResponseItemsSQLSlowRecord) *DescribeSlowLogRecordsResponseItems {
	s.SQLSlowRecord = v
	return s
}

type DescribeSlowLogRecordsResponseItemsSQLSlowRecord struct {
	HostAddress        *string `json:"HostAddress,omitempty" xml:"HostAddress,omitempty" require:"true"`
	DBName             *string `json:"DBName,omitempty" xml:"DBName,omitempty" require:"true"`
	SQLText            *string `json:"SQLText,omitempty" xml:"SQLText,omitempty" require:"true"`
	QueryTimes         *int64  `json:"QueryTimes,omitempty" xml:"QueryTimes,omitempty" require:"true"`
	LockTimes          *int64  `json:"LockTimes,omitempty" xml:"LockTimes,omitempty" require:"true"`
	ParseRowCounts     *int64  `json:"ParseRowCounts,omitempty" xml:"ParseRowCounts,omitempty" require:"true"`
	ReturnRowCounts    *int64  `json:"ReturnRowCounts,omitempty" xml:"ReturnRowCounts,omitempty" require:"true"`
	ExecutionStartTime *string `json:"ExecutionStartTime,omitempty" xml:"ExecutionStartTime,omitempty" require:"true"`
}

func (s DescribeSlowLogRecordsResponseItemsSQLSlowRecord) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowLogRecordsResponseItemsSQLSlowRecord) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetHostAddress(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.HostAddress = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetDBName(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.DBName = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetSQLText(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetQueryTimes(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.QueryTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetLockTimes(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.LockTimes = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetParseRowCounts(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ParseRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetReturnRowCounts(v int64) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ReturnRowCounts = &v
	return s
}

func (s *DescribeSlowLogRecordsResponseItemsSQLSlowRecord) SetExecutionStartTime(v string) *DescribeSlowLogRecordsResponseItemsSQLSlowRecord {
	s.ExecutionStartTime = &v
	return s
}

type SwitchDBInstanceNetTypeRequest struct {
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty" require:"true"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
}

func (s SwitchDBInstanceNetTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeRequest) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeRequest) SetDBInstanceId(v string) *SwitchDBInstanceNetTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetConnectionStringPrefix(v string) *SwitchDBInstanceNetTypeRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *SwitchDBInstanceNetTypeRequest) SetPort(v string) *SwitchDBInstanceNetTypeRequest {
	s.Port = &v
	return s
}

type SwitchDBInstanceNetTypeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s SwitchDBInstanceNetTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchDBInstanceNetTypeResponse) GoString() string {
	return s.String()
}

func (s *SwitchDBInstanceNetTypeResponse) SetRequestId(v string) *SwitchDBInstanceNetTypeResponse {
	s.RequestId = &v
	return s
}

type RestartDBInstanceRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s RestartDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceRequest) SetClientToken(v string) *RestartDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RestartDBInstanceRequest) SetDBInstanceId(v string) *RestartDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type RestartDBInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s RestartDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RestartDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartDBInstanceResponse) SetRequestId(v string) *RestartDBInstanceResponse {
	s.RequestId = &v
	return s
}

type ResetAccountPasswordRequest struct {
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	AccountName     *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	AccountPassword *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty" require:"true"`
}

func (s ResetAccountPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordRequest) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordRequest) SetDBInstanceId(v string) *ResetAccountPasswordRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountName(v string) *ResetAccountPasswordRequest {
	s.AccountName = &v
	return s
}

func (s *ResetAccountPasswordRequest) SetAccountPassword(v string) *ResetAccountPasswordRequest {
	s.AccountPassword = &v
	return s
}

type ResetAccountPasswordResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ResetAccountPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetAccountPasswordResponse) SetRequestId(v string) *ResetAccountPasswordResponse {
	s.RequestId = &v
	return s
}

type ReleaseInstancePublicConnectionRequest struct {
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty" require:"true"`
}

func (s ReleaseInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionRequest) SetDBInstanceId(v string) *ReleaseInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ReleaseInstancePublicConnectionRequest) SetCurrentConnectionString(v string) *ReleaseInstancePublicConnectionRequest {
	s.CurrentConnectionString = &v
	return s
}

type ReleaseInstancePublicConnectionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ReleaseInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseInstancePublicConnectionResponse) SetRequestId(v string) *ReleaseInstancePublicConnectionResponse {
	s.RequestId = &v
	return s
}

type ModifySecurityIpsRequest struct {
	DBInstanceId               *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty" require:"true"`
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

type ModifySecurityIpsResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifySecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) SetRequestId(v string) *ModifySecurityIpsResponse {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceNetworkTypeRequest struct {
	DBInstanceId        *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	InstanceNetworkType *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty" require:"true"`
	VPCId               *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId           *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress    *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s ModifyDBInstanceNetworkTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetDBInstanceId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetInstanceNetworkType(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVPCId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VPCId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetVSwitchId(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.VSwitchId = &v
	return s
}

func (s *ModifyDBInstanceNetworkTypeRequest) SetPrivateIpAddress(v string) *ModifyDBInstanceNetworkTypeRequest {
	s.PrivateIpAddress = &v
	return s
}

type ModifyDBInstanceNetworkTypeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceNetworkTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceNetworkTypeResponse) SetRequestId(v string) *ModifyDBInstanceNetworkTypeResponse {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceMaintainTimeRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s ModifyDBInstanceMaintainTimeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetDBInstanceId(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetStartTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyDBInstanceMaintainTimeRequest) SetEndTime(v string) *ModifyDBInstanceMaintainTimeRequest {
	s.EndTime = &v
	return s
}

type ModifyDBInstanceMaintainTimeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceMaintainTimeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceMaintainTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceMaintainTimeResponse) SetRequestId(v string) *ModifyDBInstanceMaintainTimeResponse {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceDescriptionRequest struct {
	DBInstanceId          *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty" require:"true"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

type ModifyDBInstanceDescriptionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionResponse) SetRequestId(v string) *ModifyDBInstanceDescriptionResponse {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionStringRequest struct {
	DBInstanceId            *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	ConnectionStringPrefix  *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty" require:"true"`
	Port                    *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	CurrentConnectionString *string `json:"CurrentConnectionString,omitempty" xml:"CurrentConnectionString,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionStringRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionStringRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetConnectionStringPrefix(v string) *ModifyDBInstanceConnectionStringRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetPort(v string) *ModifyDBInstanceConnectionStringRequest {
	s.Port = &v
	return s
}

func (s *ModifyDBInstanceConnectionStringRequest) SetCurrentConnectionString(v string) *ModifyDBInstanceConnectionStringRequest {
	s.CurrentConnectionString = &v
	return s
}

type ModifyDBInstanceConnectionStringResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionStringResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionStringResponse) SetRequestId(v string) *ModifyDBInstanceConnectionStringResponse {
	s.RequestId = &v
	return s
}

type ModifyDBInstanceConnectionModeRequest struct {
	DBInstanceId   *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	ConnectionMode *string `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionModeRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeRequest) SetDBInstanceId(v string) *ModifyDBInstanceConnectionModeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConnectionModeRequest) SetConnectionMode(v string) *ModifyDBInstanceConnectionModeRequest {
	s.ConnectionMode = &v
	return s
}

type ModifyDBInstanceConnectionModeResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyDBInstanceConnectionModeResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDBInstanceConnectionModeResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConnectionModeResponse) SetRequestId(v string) *ModifyDBInstanceConnectionModeResponse {
	s.RequestId = &v
	return s
}

type ModifyAccountDescriptionRequest struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionRequest) SetDBInstanceId(v string) *ModifyAccountDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountName(v string) *ModifyAccountDescriptionRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountDescriptionRequest) SetAccountDescription(v string) *ModifyAccountDescriptionRequest {
	s.AccountDescription = &v
	return s
}

type ModifyAccountDescriptionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s ModifyAccountDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAccountDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyAccountDescriptionResponse) SetRequestId(v string) *ModifyAccountDescriptionResponse {
	s.RequestId = &v
	return s
}

type DescribeResourceUsageRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeResourceUsageRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageRequest) SetDBInstanceId(v string) *DescribeResourceUsageRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeResourceUsageResponse struct {
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	DiskUsed     *int64  `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty" require:"true"`
	DataSize     *int64  `json:"DataSize,omitempty" xml:"DataSize,omitempty" require:"true"`
	LogSize      *int64  `json:"LogSize,omitempty" xml:"LogSize,omitempty" require:"true"`
	BackupSize   *int64  `json:"BackupSize,omitempty" xml:"BackupSize,omitempty" require:"true"`
}

func (s DescribeResourceUsageResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceUsageResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceUsageResponse) SetRequestId(v string) *DescribeResourceUsageResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetDBInstanceId(v string) *DescribeResourceUsageResponse {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetEngine(v string) *DescribeResourceUsageResponse {
	s.Engine = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetDiskUsed(v int64) *DescribeResourceUsageResponse {
	s.DiskUsed = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetDataSize(v int64) *DescribeResourceUsageResponse {
	s.DataSize = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetLogSize(v int64) *DescribeResourceUsageResponse {
	s.LogSize = &v
	return s
}

func (s *DescribeResourceUsageResponse) SetBackupSize(v int64) *DescribeResourceUsageResponse {
	s.BackupSize = &v
	return s
}

type DescribeRegionsRequest struct {
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegion(v string) *DescribeRegionsRequest {
	s.Region = &v
	return s
}

type DescribeRegionsResponse struct {
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Regions   *DescribeRegionsResponseRegions `json:"Regions,omitempty" xml:"Regions,omitempty" require:"true" type:"Struct"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetRequestId(v string) *DescribeRegionsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponse) SetRegions(v *DescribeRegionsResponseRegions) *DescribeRegionsResponse {
	s.Regions = v
	return s
}

type DescribeRegionsResponseRegions struct {
	Region []*DescribeRegionsResponseRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponseRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegions) SetRegion(v []*DescribeRegionsResponseRegionsRegion) *DescribeRegionsResponseRegions {
	s.Region = v
	return s
}

type DescribeRegionsResponseRegionsRegion struct {
	RegionId *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	Zones    *DescribeRegionsResponseRegionsRegionZones `json:"Zones,omitempty" xml:"Zones,omitempty" require:"true" type:"Struct"`
}

func (s DescribeRegionsResponseRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegion) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegion) SetRegionId(v string) *DescribeRegionsResponseRegionsRegion {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseRegionsRegion) SetZones(v *DescribeRegionsResponseRegionsRegionZones) *DescribeRegionsResponseRegionsRegion {
	s.Zones = v
	return s
}

type DescribeRegionsResponseRegionsRegionZones struct {
	Zone []*DescribeRegionsResponseRegionsRegionZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeRegionsResponseRegionsRegionZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegionZones) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegionZones) SetZone(v []*DescribeRegionsResponseRegionsRegionZonesZone) *DescribeRegionsResponseRegionsRegionZones {
	s.Zone = v
	return s
}

type DescribeRegionsResponseRegionsRegionZonesZone struct {
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	VpcEnabled *bool   `json:"VpcEnabled,omitempty" xml:"VpcEnabled,omitempty" require:"true"`
}

func (s DescribeRegionsResponseRegionsRegionZonesZone) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseRegionsRegionZonesZone) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseRegionsRegionZonesZone) SetZoneId(v string) *DescribeRegionsResponseRegionsRegionZonesZone {
	s.ZoneId = &v
	return s
}

func (s *DescribeRegionsResponseRegionsRegionZonesZone) SetVpcEnabled(v bool) *DescribeRegionsResponseRegionsRegionZonesZone {
	s.VpcEnabled = &v
	return s
}

type DescribeDBInstancesRequest struct {
	OwnerId               *int64                           `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string                          `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	DBInstanceDescription *string                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	InstanceNetworkType   *string                          `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	DBInstanceIds         *string                          `json:"DBInstanceIds,omitempty" xml:"DBInstanceIds,omitempty"`
	Tag                   []*DescribeDBInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	PageSize              *int                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber            *int                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeDBInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequest) SetOwnerId(v int64) *DescribeDBInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetRegionId(v string) *DescribeDBInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceDescription(v string) *DescribeDBInstancesRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetInstanceNetworkType(v string) *DescribeDBInstancesRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetDBInstanceIds(v string) *DescribeDBInstancesRequest {
	s.DBInstanceIds = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetTag(v []*DescribeDBInstancesRequestTag) *DescribeDBInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageSize(v int) *DescribeDBInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDBInstancesRequest) SetPageNumber(v int) *DescribeDBInstancesRequest {
	s.PageNumber = &v
	return s
}

type DescribeDBInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDBInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesRequestTag) SetKey(v string) *DescribeDBInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesRequestTag) SetValue(v string) *DescribeDBInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeDBInstancesResponse struct {
	RequestId        *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	PageNumber       *int                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty" require:"true"`
	TotalRecordCount *int                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty" require:"true"`
	PageRecordCount  *int                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty" require:"true"`
	Items            *DescribeDBInstancesResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponse) SetRequestId(v string) *DescribeDBInstancesResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetPageNumber(v int) *DescribeDBInstancesResponse {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetTotalRecordCount(v int) *DescribeDBInstancesResponse {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetPageRecordCount(v int) *DescribeDBInstancesResponse {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDBInstancesResponse) SetItems(v *DescribeDBInstancesResponseItems) *DescribeDBInstancesResponse {
	s.Items = v
	return s
}

type DescribeDBInstancesResponseItems struct {
	DBInstance []*DescribeDBInstancesResponseItemsDBInstance `json:"DBInstance,omitempty" xml:"DBInstance,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstancesResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseItems) SetDBInstance(v []*DescribeDBInstancesResponseItemsDBInstance) *DescribeDBInstancesResponseItems {
	s.DBInstance = v
	return s
}

type DescribeDBInstancesResponseItemsDBInstance struct {
	DBInstanceId          *string                                         `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DBInstanceDescription *string                                         `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty" require:"true"`
	PayType               *string                                         `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	InstanceNetworkType   *string                                         `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty" require:"true"`
	ConnectionMode        *string                                         `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty" require:"true"`
	RegionId              *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ZoneId                *string                                         `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	ExpireTime            *string                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	DBInstanceStatus      *string                                         `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty" require:"true"`
	Engine                *string                                         `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	EngineVersion         *string                                         `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" require:"true"`
	DBInstanceNetType     *string                                         `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty" require:"true"`
	LockMode              *string                                         `json:"LockMode,omitempty" xml:"LockMode,omitempty" require:"true"`
	LockReason            *string                                         `json:"LockReason,omitempty" xml:"LockReason,omitempty" require:"true"`
	CreateTime            *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty" require:"true"`
	VpcId                 *string                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	VSwitchId             *string                                         `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	InstanceDeployType    *string                                         `json:"InstanceDeployType,omitempty" xml:"InstanceDeployType,omitempty" require:"true"`
	Tags                  *DescribeDBInstancesResponseItemsDBInstanceTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstancesResponseItemsDBInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseItemsDBInstance) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetDBInstanceId(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetDBInstanceDescription(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetPayType(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetInstanceNetworkType(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetConnectionMode(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetRegionId(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetZoneId(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetExpireTime(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetDBInstanceStatus(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetEngine(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetEngineVersion(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetDBInstanceNetType(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetLockMode(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetLockReason(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetCreateTime(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetVpcId(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetVSwitchId(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetInstanceDeployType(v string) *DescribeDBInstancesResponseItemsDBInstance {
	s.InstanceDeployType = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstance) SetTags(v *DescribeDBInstancesResponseItemsDBInstanceTags) *DescribeDBInstancesResponseItemsDBInstance {
	s.Tags = v
	return s
}

type DescribeDBInstancesResponseItemsDBInstanceTags struct {
	Tag []*DescribeDBInstancesResponseItemsDBInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstancesResponseItemsDBInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseItemsDBInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseItemsDBInstanceTags) SetTag(v []*DescribeDBInstancesResponseItemsDBInstanceTagsTag) *DescribeDBInstancesResponseItemsDBInstanceTags {
	s.Tag = v
	return s
}

type DescribeDBInstancesResponseItemsDBInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeDBInstancesResponseItemsDBInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancesResponseItemsDBInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancesResponseItemsDBInstanceTagsTag) SetKey(v string) *DescribeDBInstancesResponseItemsDBInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancesResponseItemsDBInstanceTagsTag) SetValue(v string) *DescribeDBInstancesResponseItemsDBInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeDBInstancePerformanceRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
}

func (s DescribeDBInstancePerformanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceRequest) SetDBInstanceId(v string) *DescribeDBInstancePerformanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetKey(v string) *DescribeDBInstancePerformanceRequest {
	s.Key = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetStartTime(v string) *DescribeDBInstancePerformanceRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceRequest) SetEndTime(v string) *DescribeDBInstancePerformanceRequest {
	s.EndTime = &v
	return s
}

type DescribeDBInstancePerformanceResponse struct {
	RequestId       *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId    *string   `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	Engine          *string   `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	StartTime       *string   `json:"StartTime,omitempty" xml:"StartTime,omitempty" require:"true"`
	EndTime         *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty" require:"true"`
	PerformanceKeys []*string `json:"PerformanceKeys,omitempty" xml:"PerformanceKeys,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstancePerformanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstancePerformanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePerformanceResponse) SetRequestId(v string) *DescribeDBInstancePerformanceResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetDBInstanceId(v string) *DescribeDBInstancePerformanceResponse {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetEngine(v string) *DescribeDBInstancePerformanceResponse {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetStartTime(v string) *DescribeDBInstancePerformanceResponse {
	s.StartTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetEndTime(v string) *DescribeDBInstancePerformanceResponse {
	s.EndTime = &v
	return s
}

func (s *DescribeDBInstancePerformanceResponse) SetPerformanceKeys(v []*string) *DescribeDBInstancePerformanceResponse {
	s.PerformanceKeys = v
	return s
}

type DescribeDBInstanceNetInfoRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceNetInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoRequest) SetDBInstanceId(v string) *DescribeDBInstanceNetInfoRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceNetInfoResponse struct {
	RequestId           *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	InstanceNetworkType *string                                              `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty" require:"true"`
	DBInstanceNetInfos  *DescribeDBInstanceNetInfoResponseDBInstanceNetInfos `json:"DBInstanceNetInfos,omitempty" xml:"DBInstanceNetInfos,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceNetInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponse) SetRequestId(v string) *DescribeDBInstanceNetInfoResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetInstanceNetworkType(v string) *DescribeDBInstanceNetInfoResponse {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponse) SetDBInstanceNetInfos(v *DescribeDBInstanceNetInfoResponseDBInstanceNetInfos) *DescribeDBInstanceNetInfoResponse {
	s.DBInstanceNetInfos = v
	return s
}

type DescribeDBInstanceNetInfoResponseDBInstanceNetInfos struct {
	DBInstanceNetInfo []*DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo `json:"DBInstanceNetInfo,omitempty" xml:"DBInstanceNetInfo,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceNetInfoResponseDBInstanceNetInfos) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseDBInstanceNetInfos) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfos) SetDBInstanceNetInfo(v []*DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfos {
	s.DBInstanceNetInfo = v
	return s
}

type DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo struct {
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty" require:"true"`
	IPAddress        *string `json:"IPAddress,omitempty" xml:"IPAddress,omitempty" require:"true"`
	IPType           *string `json:"IPType,omitempty" xml:"IPType,omitempty" require:"true"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	VPCId            *string `json:"VPCId,omitempty" xml:"VPCId,omitempty" require:"true"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" require:"true"`
	VpcInstanceId    *string `json:"VpcInstanceId,omitempty" xml:"VpcInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetConnectionString(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetIPAddress(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.IPAddress = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetIPType(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.IPType = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetPort(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetVPCId(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.VPCId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetVSwitchId(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.VSwitchId = &v
	return s
}

func (s *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo) SetVpcInstanceId(v string) *DescribeDBInstanceNetInfoResponseDBInstanceNetInfosDBInstanceNetInfo {
	s.VpcInstanceId = &v
	return s
}

type DescribeDBInstanceIPArrayListRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceIPArrayListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListRequest) SetDBInstanceId(v string) *DescribeDBInstanceIPArrayListRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceIPArrayListResponse struct {
	RequestId *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Items     *DescribeDBInstanceIPArrayListResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceIPArrayListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponse) SetRequestId(v string) *DescribeDBInstanceIPArrayListResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponse) SetItems(v *DescribeDBInstanceIPArrayListResponseItems) *DescribeDBInstanceIPArrayListResponse {
	s.Items = v
	return s
}

type DescribeDBInstanceIPArrayListResponseItems struct {
	DBInstanceIPArray []*DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray `json:"DBInstanceIPArray,omitempty" xml:"DBInstanceIPArray,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceIPArrayListResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseItems) SetDBInstanceIPArray(v []*DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) *DescribeDBInstanceIPArrayListResponseItems {
	s.DBInstanceIPArray = v
	return s
}

type DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray struct {
	DBInstanceIPArrayName      *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty" require:"true"`
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty" require:"true"`
	SecurityIPList             *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty" require:"true"`
}

func (s DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) SetDBInstanceIPArrayName(v string) *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) SetDBInstanceIPArrayAttribute(v string) *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray) SetSecurityIPList(v string) *DescribeDBInstanceIPArrayListResponseItemsDBInstanceIPArray {
	s.SecurityIPList = &v
	return s
}

type DescribeDBInstanceAttributeRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DescribeDBInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeRequest) SetOwnerId(v int64) *DescribeDBInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstanceAttributeRequest) SetDBInstanceId(v string) *DescribeDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

type DescribeDBInstanceAttributeResponse struct {
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Items     *DescribeDBInstanceAttributeResponseItems `json:"Items,omitempty" xml:"Items,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponse) SetRequestId(v string) *DescribeDBInstanceAttributeResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponse) SetItems(v *DescribeDBInstanceAttributeResponseItems) *DescribeDBInstanceAttributeResponse {
	s.Items = v
	return s
}

type DescribeDBInstanceAttributeResponseItems struct {
	DBInstanceAttribute []*DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute `json:"DBInstanceAttribute,omitempty" xml:"DBInstanceAttribute,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseItems) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseItems) SetDBInstanceAttribute(v []*DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) *DescribeDBInstanceAttributeResponseItems {
	s.DBInstanceAttribute = v
	return s
}

type DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute struct {
	DBInstanceId          *string                                                          `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	PayType               *string                                                          `json:"PayType,omitempty" xml:"PayType,omitempty" require:"true"`
	DBInstanceClassType   *string                                                          `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty" require:"true"`
	RegionId              *string                                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ConnectionString      *string                                                          `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty" require:"true"`
	Port                  *string                                                          `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
	Engine                *string                                                          `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	EngineVersion         *string                                                          `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" require:"true"`
	DBInstanceClass       *string                                                          `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" require:"true"`
	DBInstanceCpuCores    *int                                                             `json:"DBInstanceCpuCores,omitempty" xml:"DBInstanceCpuCores,omitempty" require:"true"`
	DBInstanceMemory      *int64                                                           `json:"DBInstanceMemory,omitempty" xml:"DBInstanceMemory,omitempty" require:"true"`
	DBInstanceStorage     *int64                                                           `json:"DBInstanceStorage,omitempty" xml:"DBInstanceStorage,omitempty" require:"true"`
	DBInstanceDiskMBPS    *int64                                                           `json:"DBInstanceDiskMBPS,omitempty" xml:"DBInstanceDiskMBPS,omitempty" require:"true"`
	HostType              *string                                                          `json:"HostType,omitempty" xml:"HostType,omitempty" require:"true"`
	DBInstanceGroupCount  *string                                                          `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" require:"true"`
	DBInstanceNetType     *string                                                          `json:"DBInstanceNetType,omitempty" xml:"DBInstanceNetType,omitempty" require:"true"`
	DBInstanceStatus      *string                                                          `json:"DBInstanceStatus,omitempty" xml:"DBInstanceStatus,omitempty" require:"true"`
	DBInstanceDescription *string                                                          `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty" require:"true"`
	LockMode              *string                                                          `json:"LockMode,omitempty" xml:"LockMode,omitempty" require:"true"`
	LockReason            *string                                                          `json:"LockReason,omitempty" xml:"LockReason,omitempty" require:"true"`
	ReadDelayTime         *string                                                          `json:"ReadDelayTime,omitempty" xml:"ReadDelayTime,omitempty" require:"true"`
	CreationTime          *string                                                          `json:"CreationTime,omitempty" xml:"CreationTime,omitempty" require:"true"`
	ExpireTime            *string                                                          `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty" require:"true"`
	MaintainStartTime     *string                                                          `json:"MaintainStartTime,omitempty" xml:"MaintainStartTime,omitempty" require:"true"`
	MaintainEndTime       *string                                                          `json:"MaintainEndTime,omitempty" xml:"MaintainEndTime,omitempty" require:"true"`
	AvailabilityValue     *string                                                          `json:"AvailabilityValue,omitempty" xml:"AvailabilityValue,omitempty" require:"true"`
	MaxConnections        *int                                                             `json:"MaxConnections,omitempty" xml:"MaxConnections,omitempty" require:"true"`
	SecurityIPList        *string                                                          `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty" require:"true"`
	ZoneId                *string                                                          `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	InstanceNetworkType   *string                                                          `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty" require:"true"`
	VpcId                 *string                                                          `json:"VpcId,omitempty" xml:"VpcId,omitempty" require:"true"`
	ConnectionMode        *string                                                          `json:"ConnectionMode,omitempty" xml:"ConnectionMode,omitempty" require:"true"`
	StorageType           *string                                                          `json:"StorageType,omitempty" xml:"StorageType,omitempty" require:"true"`
	CpuCoresPerNode       *int                                                             `json:"CpuCoresPerNode,omitempty" xml:"CpuCoresPerNode,omitempty" require:"true"`
	SegmentCounts         *int                                                             `json:"SegmentCounts,omitempty" xml:"SegmentCounts,omitempty" require:"true"`
	StoragePerNode        *int                                                             `json:"StoragePerNode,omitempty" xml:"StoragePerNode,omitempty" require:"true"`
	MemoryPerNode         *int                                                             `json:"MemoryPerNode,omitempty" xml:"MemoryPerNode,omitempty" require:"true"`
	StorageUnit           *string                                                          `json:"StorageUnit,omitempty" xml:"StorageUnit,omitempty" require:"true"`
	MemoryUnit            *string                                                          `json:"MemoryUnit,omitempty" xml:"MemoryUnit,omitempty" require:"true"`
	Tags                  *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags `json:"Tags,omitempty" xml:"Tags,omitempty" require:"true" type:"Struct"`
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceId(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetPayType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.PayType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceClassType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetRegionId(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.RegionId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetConnectionString(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.ConnectionString = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetPort(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.Port = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetEngine(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.Engine = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetEngineVersion(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.EngineVersion = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceClass(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceCpuCores(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceCpuCores = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceMemory(v int64) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceMemory = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceStorage(v int64) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceStorage = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceDiskMBPS(v int64) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceDiskMBPS = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetHostType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.HostType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceGroupCount(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceNetType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceNetType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceStatus(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceStatus = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetDBInstanceDescription(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.DBInstanceDescription = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetLockMode(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.LockMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetLockReason(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.LockReason = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetReadDelayTime(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.ReadDelayTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetCreationTime(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.CreationTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetExpireTime(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetMaintainStartTime(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.MaintainStartTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetMaintainEndTime(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.MaintainEndTime = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetAvailabilityValue(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.AvailabilityValue = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetMaxConnections(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.MaxConnections = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetSecurityIPList(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetZoneId(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.ZoneId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetInstanceNetworkType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.InstanceNetworkType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetVpcId(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.VpcId = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetConnectionMode(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.ConnectionMode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetStorageType(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetCpuCoresPerNode(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.CpuCoresPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetSegmentCounts(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.SegmentCounts = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetStoragePerNode(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.StoragePerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetMemoryPerNode(v int) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.MemoryPerNode = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetStorageUnit(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.StorageUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetMemoryUnit(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.MemoryUnit = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute) SetTags(v *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttribute {
	s.Tags = v
	return s
}

type DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags struct {
	Tag []*DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags) SetTag(v []*DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTags {
	s.Tag = v
	return s
}

type DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty" require:"true"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty" require:"true"`
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag) SetKey(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag) SetValue(v string) *DescribeDBInstanceAttributeResponseItemsDBInstanceAttributeTagsTag {
	s.Value = &v
	return s
}

type DescribeAccountsRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	AccountName  *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
}

func (s DescribeAccountsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountsRequest) SetDBInstanceId(v string) *DescribeAccountsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsRequest) SetAccountName(v string) *DescribeAccountsRequest {
	s.AccountName = &v
	return s
}

type DescribeAccountsResponse struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	Accounts  *DescribeAccountsResponseAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" require:"true" type:"Struct"`
}

func (s DescribeAccountsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponse) SetRequestId(v string) *DescribeAccountsResponse {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountsResponse) SetAccounts(v *DescribeAccountsResponseAccounts) *DescribeAccountsResponse {
	s.Accounts = v
	return s
}

type DescribeAccountsResponseAccounts struct {
	DBInstanceAccount []*DescribeAccountsResponseAccountsDBInstanceAccount `json:"DBInstanceAccount,omitempty" xml:"DBInstanceAccount,omitempty" require:"true" type:"Repeated"`
}

func (s DescribeAccountsResponseAccounts) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseAccounts) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseAccounts) SetDBInstanceAccount(v []*DescribeAccountsResponseAccountsDBInstanceAccount) *DescribeAccountsResponseAccounts {
	s.DBInstanceAccount = v
	return s
}

type DescribeAccountsResponseAccountsDBInstanceAccount struct {
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	AccountStatus      *string `json:"AccountStatus,omitempty" xml:"AccountStatus,omitempty" require:"true"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty" require:"true"`
}

func (s DescribeAccountsResponseAccountsDBInstanceAccount) String() string {
	return tea.Prettify(s)
}

func (s DescribeAccountsResponseAccountsDBInstanceAccount) GoString() string {
	return s.String()
}

func (s *DescribeAccountsResponseAccountsDBInstanceAccount) SetDBInstanceId(v string) *DescribeAccountsResponseAccountsDBInstanceAccount {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAccountsResponseAccountsDBInstanceAccount) SetAccountName(v string) *DescribeAccountsResponseAccountsDBInstanceAccount {
	s.AccountName = &v
	return s
}

func (s *DescribeAccountsResponseAccountsDBInstanceAccount) SetAccountStatus(v string) *DescribeAccountsResponseAccountsDBInstanceAccount {
	s.AccountStatus = &v
	return s
}

func (s *DescribeAccountsResponseAccountsDBInstanceAccount) SetAccountDescription(v string) *DescribeAccountsResponseAccountsDBInstanceAccount {
	s.AccountDescription = &v
	return s
}

type DeleteDBInstanceRequest struct {
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
}

func (s DeleteDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceRequest) SetOwnerId(v int64) *DeleteDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetClientToken(v string) *DeleteDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteDBInstanceRequest) SetDBInstanceId(v string) *DeleteDBInstanceRequest {
	s.DBInstanceId = &v
	return s
}

type DeleteDBInstanceResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBInstanceResponse) SetRequestId(v string) *DeleteDBInstanceResponse {
	s.RequestId = &v
	return s
}

type DeleteDatabaseRequest struct {
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DBName       *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s DeleteDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseRequest) SetDBInstanceId(v string) *DeleteDatabaseRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DeleteDatabaseRequest) SetDBName(v string) *DeleteDatabaseRequest {
	s.DBName = &v
	return s
}

type DeleteDatabaseResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s DeleteDatabaseResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabaseResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabaseResponse) SetRequestId(v string) *DeleteDatabaseResponse {
	s.RequestId = &v
	return s
}

type CreateDBInstanceRequest struct {
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty" require:"true"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" require:"true"`
	EngineVersion         *string `json:"EngineVersion,omitempty" xml:"EngineVersion,omitempty" require:"true"`
	Engine                *string `json:"Engine,omitempty" xml:"Engine,omitempty" require:"true"`
	DBInstanceClass       *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty" require:"true"`
	DBInstanceGroupCount  *string `json:"DBInstanceGroupCount,omitempty" xml:"DBInstanceGroupCount,omitempty" require:"true"`
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	SecurityIPList        *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty" require:"true"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                *string `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime              *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken           *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty" require:"true"`
	InstanceNetworkType   *string `json:"InstanceNetworkType,omitempty" xml:"InstanceNetworkType,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	PrivateIpAddress      *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
}

func (s CreateDBInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceRequest) SetOwnerId(v int64) *CreateDBInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetRegionId(v string) *CreateDBInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetZoneId(v string) *CreateDBInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngineVersion(v string) *CreateDBInstanceRequest {
	s.EngineVersion = &v
	return s
}

func (s *CreateDBInstanceRequest) SetEngine(v string) *CreateDBInstanceRequest {
	s.Engine = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceClass(v string) *CreateDBInstanceRequest {
	s.DBInstanceClass = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceGroupCount(v string) *CreateDBInstanceRequest {
	s.DBInstanceGroupCount = &v
	return s
}

func (s *CreateDBInstanceRequest) SetDBInstanceDescription(v string) *CreateDBInstanceRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *CreateDBInstanceRequest) SetSecurityIPList(v string) *CreateDBInstanceRequest {
	s.SecurityIPList = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPayType(v string) *CreateDBInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPeriod(v string) *CreateDBInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDBInstanceRequest) SetUsedTime(v string) *CreateDBInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDBInstanceRequest) SetClientToken(v string) *CreateDBInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDBInstanceRequest) SetInstanceNetworkType(v string) *CreateDBInstanceRequest {
	s.InstanceNetworkType = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVPCId(v string) *CreateDBInstanceRequest {
	s.VPCId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetVSwitchId(v string) *CreateDBInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateDBInstanceRequest) SetPrivateIpAddress(v string) *CreateDBInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

type CreateDBInstanceResponse struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
	DBInstanceId     *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	OrderId          *string `json:"OrderId,omitempty" xml:"OrderId,omitempty" require:"true"`
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty" require:"true"`
	Port             *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
}

func (s CreateDBInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponse) SetRequestId(v string) *CreateDBInstanceResponse {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponse) SetDBInstanceId(v string) *CreateDBInstanceResponse {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponse) SetOrderId(v string) *CreateDBInstanceResponse {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponse) SetConnectionString(v string) *CreateDBInstanceResponse {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponse) SetPort(v string) *CreateDBInstanceResponse {
	s.Port = &v
	return s
}

type CreateAccountRequest struct {
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DBInstanceId       *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	DatabaseName       *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	AccountName        *string `json:"AccountName,omitempty" xml:"AccountName,omitempty" require:"true"`
	AccountPassword    *string `json:"AccountPassword,omitempty" xml:"AccountPassword,omitempty" require:"true"`
	AccountDescription *string `json:"AccountDescription,omitempty" xml:"AccountDescription,omitempty"`
}

func (s CreateAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountRequest) GoString() string {
	return s.String()
}

func (s *CreateAccountRequest) SetOwnerId(v int64) *CreateAccountRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateAccountRequest) SetDBInstanceId(v string) *CreateAccountRequest {
	s.DBInstanceId = &v
	return s
}

func (s *CreateAccountRequest) SetDatabaseName(v string) *CreateAccountRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountName(v string) *CreateAccountRequest {
	s.AccountName = &v
	return s
}

func (s *CreateAccountRequest) SetAccountPassword(v string) *CreateAccountRequest {
	s.AccountPassword = &v
	return s
}

func (s *CreateAccountRequest) SetAccountDescription(v string) *CreateAccountRequest {
	s.AccountDescription = &v
	return s
}

type CreateAccountResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s CreateAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateAccountResponse) SetRequestId(v string) *CreateAccountResponse {
	s.RequestId = &v
	return s
}

type AllocateInstancePublicConnectionRequest struct {
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount   *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId        *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	DBInstanceId           *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty" require:"true"`
	ConnectionStringPrefix *string `json:"ConnectionStringPrefix,omitempty" xml:"ConnectionStringPrefix,omitempty" require:"true"`
	Port                   *string `json:"Port,omitempty" xml:"Port,omitempty" require:"true"`
}

func (s AllocateInstancePublicConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionRequest) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionRequest) SetOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerAccount(v string) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetResourceOwnerId(v int64) *AllocateInstancePublicConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetDBInstanceId(v string) *AllocateInstancePublicConnectionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetConnectionStringPrefix(v string) *AllocateInstancePublicConnectionRequest {
	s.ConnectionStringPrefix = &v
	return s
}

func (s *AllocateInstancePublicConnectionRequest) SetPort(v string) *AllocateInstancePublicConnectionRequest {
	s.Port = &v
	return s
}

type AllocateInstancePublicConnectionResponse struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty" require:"true"`
}

func (s AllocateInstancePublicConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s AllocateInstancePublicConnectionResponse) GoString() string {
	return s.String()
}

func (s *AllocateInstancePublicConnectionResponse) SetRequestId(v string) *AllocateInstancePublicConnectionResponse {
	s.RequestId = &v
	return s
}

type Client struct {
	rpc.Client
}

func NewClient(config *rpc.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *rpc.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou":           tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai":           tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen":           tea.String("gpdb.aliyuncs.com"),
		"cn-hongkong":           tea.String("gpdb.aliyuncs.com"),
		"ap-southeast-1":        tea.String("gpdb.aliyuncs.com"),
		"us-west-1":             tea.String("gpdb.aliyuncs.com"),
		"us-east-1":             tea.String("gpdb.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("gpdb.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("gpdb.aliyuncs.com"),
		"cn-qingdao":            tea.String("gpdb.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("gpdb.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("gpdb"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) DescribeUserEncryptionKeyListWithOptions(request *DescribeUserEncryptionKeyListRequest, runtime *util.RuntimeOptions) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeUserEncryptionKeyList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserEncryptionKeyList(request *DescribeUserEncryptionKeyListRequest) (_result *DescribeUserEncryptionKeyListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserEncryptionKeyListResponse{}
	_body, _err := client.DescribeUserEncryptionKeyListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeModifyParameterLogWithOptions(request *DescribeModifyParameterLogRequest, runtime *util.RuntimeOptions) (_result *DescribeModifyParameterLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeModifyParameterLog"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeModifyParameterLog(request *DescribeModifyParameterLogRequest) (_result *DescribeModifyParameterLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeModifyParameterLogResponse{}
	_body, _err := client.DescribeModifyParameterLogWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeParametersWithOptions(request *DescribeParametersRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeParameters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParameters(request *DescribeParametersRequest) (_result *DescribeParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersResponse{}
	_body, _err := client.DescribeParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyParametersWithOptions(request *ModifyParametersRequest, runtime *util.RuntimeOptions) (_result *ModifyParametersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyParameters"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyParameters(request *ModifyParametersRequest) (_result *ModifyParametersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyParametersResponse{}
	_body, _err := client.ModifyParametersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceLinkedRoleWithOptions(request *CreateServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CreateServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CreateServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceLinkedRole(request *CreateServiceLinkedRoleRequest) (_result *CreateServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateServiceLinkedRoleResponse{}
	_body, _err := client.CreateServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CheckServiceLinkedRoleWithOptions(request *CheckServiceLinkedRoleRequest, runtime *util.RuntimeOptions) (_result *CheckServiceLinkedRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.DoRequest(tea.String("CheckServiceLinkedRole"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CheckServiceLinkedRole(request *CheckServiceLinkedRoleRequest) (_result *CheckServiceLinkedRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CheckServiceLinkedRoleResponse{}
	_body, _err := client.CheckServiceLinkedRoleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogCountWithOptions(request *DescribeSQLLogCountRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogCountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSQLLogCount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogCount(request *DescribeSQLLogCountRequest) (_result *DescribeSQLLogCountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogCountResponse{}
	_body, _err := client.DescribeSQLLogCountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogsWithOptions(request *DescribeSQLLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSQLLogs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogs(request *DescribeSQLLogsRequest) (_result *DescribeSQLLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogsResponse{}
	_body, _err := client.DescribeSQLLogsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateECSDBInstanceWithOptions(request *CreateECSDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateECSDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateECSDBInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateECSDBInstance(request *CreateECSDBInstanceRequest) (_result *CreateECSDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateECSDBInstanceResponse{}
	_body, _err := client.CreateECSDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformanceWithOptions(request *DescribeDBClusterPerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBClusterPerformance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBClusterPerformance(request *DescribeDBClusterPerformanceRequest) (_result *DescribeDBClusterPerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBClusterPerformanceResponse{}
	_body, _err := client.DescribeDBClusterPerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceOnECSAttributeWithOptions(request *DescribeDBInstanceOnECSAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstanceOnECSAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceOnECSAttribute(request *DescribeDBInstanceOnECSAttributeRequest) (_result *DescribeDBInstanceOnECSAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceOnECSAttributeResponse{}
	_body, _err := client.DescribeDBInstanceOnECSAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableResourcesWithOptions(request *DescribeAvailableResourcesRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAvailableResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableResources(request *DescribeAvailableResourcesRequest) (_result *DescribeAvailableResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableResourcesResponse{}
	_body, _err := client.DescribeAvailableResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSLWithOptions(request *DescribeDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstanceSSL"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceSSL(request *DescribeDBInstanceSSLRequest) (_result *DescribeDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceSSLResponse{}
	_body, _err := client.DescribeDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSLWithOptions(request *ModifyDBInstanceSSLRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceSSLResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceSSL"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceSSL(request *ModifyDBInstanceSSLRequest) (_result *ModifyDBInstanceSSLResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceSSLResponse{}
	_body, _err := client.ModifyDBInstanceSSLWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTagsWithOptions(request *DescribeTagsRequest, runtime *util.RuntimeOptions) (_result *DescribeTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeTags"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTags(request *DescribeTagsRequest) (_result *DescribeTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTagsResponse{}
	_body, _err := client.DescribeTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSpecificationWithOptions(request *DescribeSpecificationRequest, runtime *util.RuntimeOptions) (_result *DescribeSpecificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSpecification"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSpecification(request *DescribeSpecificationRequest) (_result *DescribeSpecificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSpecificationResponse{}
	_body, _err := client.DescribeSpecificationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBVersionWithOptions(request *UpgradeDBVersionRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.DoRequest(tea.String("UpgradeDBVersion"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBVersion(request *UpgradeDBVersionRequest) (_result *UpgradeDBVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBVersionResponse{}
	_body, _err := client.UpgradeDBVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpgradeDBInstanceWithOptions(request *UpgradeDBInstanceRequest, runtime *util.RuntimeOptions) (_result *UpgradeDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("UpgradeDBInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeDBInstance(request *UpgradeDBInstanceRequest) (_result *UpgradeDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeDBInstanceResponse{}
	_body, _err := client.UpgradeDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("UntagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("TagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRequest(tea.String("ListTagResources"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsVSwitchsWithOptions(request *DescribeRdsVSwitchsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVSwitchsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRdsVSwitchs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsVSwitchs(request *DescribeRdsVSwitchsRequest) (_result *DescribeRdsVSwitchsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVSwitchsResponse{}
	_body, _err := client.DescribeRdsVSwitchsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRdsVpcsWithOptions(request *DescribeRdsVpcsRequest, runtime *util.RuntimeOptions) (_result *DescribeRdsVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRdsVpcs"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRdsVpcs(request *DescribeRdsVpcsRequest) (_result *DescribeRdsVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRdsVpcsResponse{}
	_body, _err := client.DescribeRdsVpcsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddBuDBInstanceRelationWithOptions(request *AddBuDBInstanceRelationRequest, runtime *util.RuntimeOptions) (_result *AddBuDBInstanceRelationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.DoRequest(tea.String("AddBuDBInstanceRelation"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddBuDBInstanceRelation(request *AddBuDBInstanceRelationRequest) (_result *AddBuDBInstanceRelationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddBuDBInstanceRelationResponse{}
	_body, _err := client.AddBuDBInstanceRelationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogRecordsWithOptions(request *DescribeSQLLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSQLLogRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogRecords(request *DescribeSQLLogRecordsRequest) (_result *DescribeSQLLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogRecordsResponse{}
	_body, _err := client.DescribeSQLLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySQLCollectorPolicyWithOptions(request *ModifySQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("ModifySQLCollectorPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySQLCollectorPolicy(request *ModifySQLCollectorPolicyRequest) (_result *ModifySQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySQLCollectorPolicyResponse{}
	_body, _err := client.ModifySQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLLogFilesWithOptions(request *DescribeSQLLogFilesRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLLogFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSQLLogFiles"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLLogFiles(request *DescribeSQLLogFilesRequest) (_result *DescribeSQLLogFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLLogFilesResponse{}
	_body, _err := client.DescribeSQLLogFilesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicyWithOptions(request *DescribeSQLCollectorPolicyRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSQLCollectorPolicy"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLCollectorPolicy(request *DescribeSQLCollectorPolicyRequest) (_result *DescribeSQLCollectorPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLCollectorPolicyResponse{}
	_body, _err := client.DescribeSQLCollectorPolicyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowLogRecordsWithOptions(request *DescribeSlowLogRecordsRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowLogRecordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeSlowLogRecords"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowLogRecords(request *DescribeSlowLogRecordsRequest) (_result *DescribeSlowLogRecordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowLogRecordsResponse{}
	_body, _err := client.DescribeSlowLogRecordsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetTypeWithOptions(request *SwitchDBInstanceNetTypeRequest, runtime *util.RuntimeOptions) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.DoRequest(tea.String("SwitchDBInstanceNetType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchDBInstanceNetType(request *SwitchDBInstanceNetTypeRequest) (_result *SwitchDBInstanceNetTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchDBInstanceNetTypeResponse{}
	_body, _err := client.SwitchDBInstanceNetTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestartDBInstanceWithOptions(request *RestartDBInstanceRequest, runtime *util.RuntimeOptions) (_result *RestartDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("RestartDBInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (_result *RestartDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RestartDBInstanceResponse{}
	_body, _err := client.RestartDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetAccountPasswordWithOptions(request *ResetAccountPasswordRequest, runtime *util.RuntimeOptions) (_result *ResetAccountPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.DoRequest(tea.String("ResetAccountPassword"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetAccountPassword(request *ResetAccountPasswordRequest) (_result *ResetAccountPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetAccountPasswordResponse{}
	_body, _err := client.ResetAccountPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnectionWithOptions(request *ReleaseInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("ReleaseInstancePublicConnection"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseInstancePublicConnection(request *ReleaseInstancePublicConnectionRequest) (_result *ReleaseInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseInstancePublicConnectionResponse{}
	_body, _err := client.ReleaseInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.DoRequest(tea.String("ModifySecurityIps"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySecurityIps(request *ModifySecurityIpsRequest) (_result *ModifySecurityIpsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.ModifySecurityIpsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkTypeWithOptions(request *ModifyDBInstanceNetworkTypeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceNetworkType"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceNetworkType(request *ModifyDBInstanceNetworkTypeRequest) (_result *ModifyDBInstanceNetworkTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceNetworkTypeResponse{}
	_body, _err := client.ModifyDBInstanceNetworkTypeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTimeWithOptions(request *ModifyDBInstanceMaintainTimeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceMaintainTime"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceMaintainTime(request *ModifyDBInstanceMaintainTimeRequest) (_result *ModifyDBInstanceMaintainTimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceMaintainTimeResponse{}
	_body, _err := client.ModifyDBInstanceMaintainTimeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceDescriptionWithOptions(request *ModifyDBInstanceDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceDescription(request *ModifyDBInstanceDescriptionRequest) (_result *ModifyDBInstanceDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceDescriptionResponse{}
	_body, _err := client.ModifyDBInstanceDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionStringWithOptions(request *ModifyDBInstanceConnectionStringRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceConnectionString"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionString(request *ModifyDBInstanceConnectionStringRequest) (_result *ModifyDBInstanceConnectionStringResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionStringResponse{}
	_body, _err := client.ModifyDBInstanceConnectionStringWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionModeWithOptions(request *ModifyDBInstanceConnectionModeRequest, runtime *util.RuntimeOptions) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyDBInstanceConnectionMode"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDBInstanceConnectionMode(request *ModifyDBInstanceConnectionModeRequest) (_result *ModifyDBInstanceConnectionModeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDBInstanceConnectionModeResponse{}
	_body, _err := client.ModifyDBInstanceConnectionModeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAccountDescriptionWithOptions(request *ModifyAccountDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyAccountDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.DoRequest(tea.String("ModifyAccountDescription"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyAccountDescription(request *ModifyAccountDescriptionRequest) (_result *ModifyAccountDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyAccountDescriptionResponse{}
	_body, _err := client.ModifyAccountDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceUsageWithOptions(request *DescribeResourceUsageRequest, runtime *util.RuntimeOptions) (_result *DescribeResourceUsageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeResourceUsage"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceUsage(request *DescribeResourceUsageRequest) (_result *DescribeResourceUsageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeResourceUsageResponse{}
	_body, _err := client.DescribeResourceUsageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeRegions"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancesWithOptions(request *DescribeDBInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstances"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstances(request *DescribeDBInstancesRequest) (_result *DescribeDBInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancesResponse{}
	_body, _err := client.DescribeDBInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformanceWithOptions(request *DescribeDBInstancePerformanceRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstancePerformance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstancePerformance(request *DescribeDBInstancePerformanceRequest) (_result *DescribeDBInstancePerformanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstancePerformanceResponse{}
	_body, _err := client.DescribeDBInstancePerformanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfoWithOptions(request *DescribeDBInstanceNetInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstanceNetInfo"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceNetInfo(request *DescribeDBInstanceNetInfoRequest) (_result *DescribeDBInstanceNetInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceNetInfoResponse{}
	_body, _err := client.DescribeDBInstanceNetInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceIPArrayListWithOptions(request *DescribeDBInstanceIPArrayListRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstanceIPArrayList"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceIPArrayList(request *DescribeDBInstanceIPArrayListRequest) (_result *DescribeDBInstanceIPArrayListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceIPArrayListResponse{}
	_body, _err := client.DescribeDBInstanceIPArrayListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDBInstanceAttributeWithOptions(request *DescribeDBInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeDBInstanceAttribute"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDBInstanceAttribute(request *DescribeDBInstanceAttributeRequest) (_result *DescribeDBInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDBInstanceAttributeResponse{}
	_body, _err := client.DescribeDBInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAccountsWithOptions(request *DescribeAccountsRequest, runtime *util.RuntimeOptions) (_result *DescribeAccountsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DoRequest(tea.String("DescribeAccounts"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAccounts(request *DescribeAccountsRequest) (_result *DescribeAccountsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAccountsResponse{}
	_body, _err := client.DescribeAccountsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDBInstanceWithOptions(request *DeleteDBInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDBInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDBInstance(request *DeleteDBInstanceRequest) (_result *DeleteDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDBInstanceResponse{}
	_body, _err := client.DeleteDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabaseWithOptions(request *DeleteDatabaseRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DoRequest(tea.String("DeleteDatabase"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabase(request *DeleteDatabaseRequest) (_result *DeleteDatabaseResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabaseResponse{}
	_body, _err := client.DeleteDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDBInstanceWithOptions(request *CreateDBInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDBInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.DoRequest(tea.String("CreateDBInstance"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDBInstance(request *CreateDBInstanceRequest) (_result *CreateDBInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDBInstanceResponse{}
	_body, _err := client.CreateDBInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAccountWithOptions(request *CreateAccountRequest, runtime *util.RuntimeOptions) (_result *CreateAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAccountResponse{}
	_body, _err := client.DoRequest(tea.String("CreateAccount"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAccount(request *CreateAccountRequest) (_result *CreateAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAccountResponse{}
	_body, _err := client.CreateAccountWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnectionWithOptions(request *AllocateInstancePublicConnectionRequest, runtime *util.RuntimeOptions) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.DoRequest(tea.String("AllocateInstancePublicConnection"), tea.String("HTTPS"), tea.String("POST"), tea.String("2016-05-03"), tea.String("AK"), nil, tea.ToMap(request), runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AllocateInstancePublicConnection(request *AllocateInstancePublicConnectionRequest) (_result *AllocateInstancePublicConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AllocateInstancePublicConnectionResponse{}
	_body, _err := client.AllocateInstancePublicConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
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
