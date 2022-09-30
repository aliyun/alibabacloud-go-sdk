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

type CreateDatabaseRequest struct {
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Collation    *string `json:"Collation,omitempty" xml:"Collation,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Encoding     *string `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateDatabaseRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseRequest) GoString() string {
	return s.String()
}

func (s *CreateDatabaseRequest) SetClientToken(v string) *CreateDatabaseRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDatabaseRequest) SetCollation(v string) *CreateDatabaseRequest {
	s.Collation = &v
	return s
}

func (s *CreateDatabaseRequest) SetDatabaseName(v string) *CreateDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseRequest) SetDescription(v string) *CreateDatabaseRequest {
	s.Description = &v
	return s
}

func (s *CreateDatabaseRequest) SetEncoding(v string) *CreateDatabaseRequest {
	s.Encoding = &v
	return s
}

func (s *CreateDatabaseRequest) SetInstanceId(v string) *CreateDatabaseRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateDatabaseRequest) SetTenantId(v string) *CreateDatabaseRequest {
	s.TenantId = &v
	return s
}

type CreateDatabaseResponseBody struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) SetDatabaseName(v string) *CreateDatabaseResponseBody {
	s.DatabaseName = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
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

type CreateInstanceRequest struct {
	AutoRenew       *bool   `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	AutoRenewPeriod *int64  `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	ChargeType      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	DiskSize        *int64  `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType        *string `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	InstanceClass   *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ObVersion       *string `json:"ObVersion,omitempty" xml:"ObVersion,omitempty"`
	Period          *int64  `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit      *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Series          *string `json:"Series,omitempty" xml:"Series,omitempty"`
	Zones           *string `json:"Zones,omitempty" xml:"Zones,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int64) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetChargeType(v string) *CreateInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetDiskSize(v int64) *CreateInstanceRequest {
	s.DiskSize = &v
	return s
}

func (s *CreateInstanceRequest) SetDiskType(v string) *CreateInstanceRequest {
	s.DiskType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceClass(v string) *CreateInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetObVersion(v string) *CreateInstanceRequest {
	s.ObVersion = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int64) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSeries(v string) *CreateInstanceRequest {
	s.Series = &v
	return s
}

func (s *CreateInstanceRequest) SetZones(v string) *CreateInstanceRequest {
	s.Zones = &v
	return s
}

type CreateInstanceResponseBody struct {
	Data      []*CreateInstanceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) SetData(v []*CreateInstanceResponseBodyData) *CreateInstanceResponseBody {
	s.Data = v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

type CreateInstanceResponseBodyData struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OrderId         *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s CreateInstanceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyData) SetInstanceId(v string) *CreateInstanceResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetOrderId(v string) *CreateInstanceResponseBodyData {
	s.OrderId = &v
	return s
}

func (s *CreateInstanceResponseBodyData) SetResourceGroupId(v string) *CreateInstanceResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

type CreateInstanceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponse) SetHeaders(v map[string]*string) *CreateInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceResponse) SetStatusCode(v int32) *CreateInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceResponse) SetBody(v *CreateInstanceResponseBody) *CreateInstanceResponse {
	s.Body = v
	return s
}

type CreateOmsMysqlDataSourceRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DgDatabaseId *string `json:"DgDatabaseId,omitempty" xml:"DgDatabaseId,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Ip           *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Schema       *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Username     *string `json:"Username,omitempty" xml:"Username,omitempty"`
	VpcId        *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateOmsMysqlDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsMysqlDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateOmsMysqlDataSourceRequest) SetDescription(v string) *CreateOmsMysqlDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetDgDatabaseId(v string) *CreateOmsMysqlDataSourceRequest {
	s.DgDatabaseId = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetInstanceId(v string) *CreateOmsMysqlDataSourceRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetIp(v string) *CreateOmsMysqlDataSourceRequest {
	s.Ip = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetName(v string) *CreateOmsMysqlDataSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetPassword(v string) *CreateOmsMysqlDataSourceRequest {
	s.Password = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetPort(v string) *CreateOmsMysqlDataSourceRequest {
	s.Port = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetSchema(v string) *CreateOmsMysqlDataSourceRequest {
	s.Schema = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetType(v string) *CreateOmsMysqlDataSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetUsername(v string) *CreateOmsMysqlDataSourceRequest {
	s.Username = &v
	return s
}

func (s *CreateOmsMysqlDataSourceRequest) SetVpcId(v string) *CreateOmsMysqlDataSourceRequest {
	s.VpcId = &v
	return s
}

type CreateOmsMysqlDataSourceResponseBody struct {
	Data      *CreateOmsMysqlDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOmsMysqlDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsMysqlDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOmsMysqlDataSourceResponseBody) SetData(v *CreateOmsMysqlDataSourceResponseBodyData) *CreateOmsMysqlDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *CreateOmsMysqlDataSourceResponseBody) SetRequestId(v string) *CreateOmsMysqlDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateOmsMysqlDataSourceResponseBodyData struct {
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
}

func (s CreateOmsMysqlDataSourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsMysqlDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOmsMysqlDataSourceResponseBodyData) SetEndpointId(v string) *CreateOmsMysqlDataSourceResponseBodyData {
	s.EndpointId = &v
	return s
}

type CreateOmsMysqlDataSourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOmsMysqlDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOmsMysqlDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsMysqlDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateOmsMysqlDataSourceResponse) SetHeaders(v map[string]*string) *CreateOmsMysqlDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateOmsMysqlDataSourceResponse) SetStatusCode(v int32) *CreateOmsMysqlDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOmsMysqlDataSourceResponse) SetBody(v *CreateOmsMysqlDataSourceResponseBody) *CreateOmsMysqlDataSourceResponse {
	s.Body = v
	return s
}

type CreateOmsOpenAPIProjectRequest struct {
	BusinessName       *string                                           `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	DestConfig         *CreateOmsOpenAPIProjectRequestDestConfig         `json:"DestConfig,omitempty" xml:"DestConfig,omitempty" type:"Struct"`
	LabelIds           []*string                                         `json:"LabelIds,omitempty" xml:"LabelIds,omitempty" type:"Repeated"`
	PageNumber         *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectName        *string                                           `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceConfig       *CreateOmsOpenAPIProjectRequestSourceConfig       `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty" type:"Struct"`
	TransferMapping    *CreateOmsOpenAPIProjectRequestTransferMapping    `json:"TransferMapping,omitempty" xml:"TransferMapping,omitempty" type:"Struct"`
	TransferStepConfig *CreateOmsOpenAPIProjectRequestTransferStepConfig `json:"TransferStepConfig,omitempty" xml:"TransferStepConfig,omitempty" type:"Struct"`
	WorkerGradeId      *string                                           `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequest) SetBusinessName(v string) *CreateOmsOpenAPIProjectRequest {
	s.BusinessName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetDestConfig(v *CreateOmsOpenAPIProjectRequestDestConfig) *CreateOmsOpenAPIProjectRequest {
	s.DestConfig = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetLabelIds(v []*string) *CreateOmsOpenAPIProjectRequest {
	s.LabelIds = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetPageNumber(v int32) *CreateOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetPageSize(v int32) *CreateOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetProjectName(v string) *CreateOmsOpenAPIProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetSourceConfig(v *CreateOmsOpenAPIProjectRequestSourceConfig) *CreateOmsOpenAPIProjectRequest {
	s.SourceConfig = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetTransferMapping(v *CreateOmsOpenAPIProjectRequestTransferMapping) *CreateOmsOpenAPIProjectRequest {
	s.TransferMapping = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetTransferStepConfig(v *CreateOmsOpenAPIProjectRequestTransferStepConfig) *CreateOmsOpenAPIProjectRequest {
	s.TransferStepConfig = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *CreateOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type CreateOmsOpenAPIProjectRequestDestConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestDestConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestDestConfig) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetEnableMsgTrace(v bool) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetEndpointId(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.EndpointId = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetEndpointType(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.EndpointType = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetMsgTags(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.MsgTags = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetPartition(v int32) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.Partition = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetPartitionMode(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.PartitionMode = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetProducerGroup(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.ProducerGroup = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetSendMsgTimeout(v int64) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetSequenceEnable(v bool) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.SequenceEnable = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetSequenceStartTimestamp(v int64) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetSerializerType(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.SerializerType = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestDestConfig) SetTopicType(v string) *CreateOmsOpenAPIProjectRequestDestConfig {
	s.TopicType = &v
	return s
}

type CreateOmsOpenAPIProjectRequestSourceConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestSourceConfig) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetEnableMsgTrace(v bool) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetEndpointId(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.EndpointId = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetEndpointType(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.EndpointType = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetMsgTags(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.MsgTags = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetPartition(v int32) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.Partition = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetPartitionMode(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.PartitionMode = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetProducerGroup(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.ProducerGroup = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetSendMsgTimeout(v int64) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetSequenceEnable(v bool) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.SequenceEnable = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetSequenceStartTimestamp(v int64) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetSerializerType(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.SerializerType = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestSourceConfig) SetTopicType(v string) *CreateOmsOpenAPIProjectRequestSourceConfig {
	s.TopicType = &v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferMapping struct {
	Databases []*CreateOmsOpenAPIProjectRequestTransferMappingDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Mode      *string                                                   `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestTransferMapping) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferMapping) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferMapping) SetDatabases(v []*CreateOmsOpenAPIProjectRequestTransferMappingDatabases) *CreateOmsOpenAPIProjectRequestTransferMapping {
	s.Databases = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMapping) SetMode(v string) *CreateOmsOpenAPIProjectRequestTransferMapping {
	s.Mode = &v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferMappingDatabases struct {
	DatabaseId   *string                                                         `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	DatabaseName *string                                                         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	MappedName   *string                                                         `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	Tables       []*CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TenantName   *string                                                         `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	Type         *string                                                         `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabases) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabases) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetDatabaseId(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.DatabaseId = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetDatabaseName(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.DatabaseName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetMappedName(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.MappedName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetTables(v []*CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.Tables = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetTenantName(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.TenantName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabases) SetType(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabases {
	s.Type = &v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables struct {
	AdbTableSchema *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema `json:"AdbTableSchema,omitempty" xml:"AdbTableSchema,omitempty" type:"Struct"`
	FilterColumns  []*string                                                                   `json:"FilterColumns,omitempty" xml:"FilterColumns,omitempty" type:"Repeated"`
	MappedName     *string                                                                     `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	ShardColumns   []*string                                                                   `json:"ShardColumns,omitempty" xml:"ShardColumns,omitempty" type:"Repeated"`
	TableId        *string                                                                     `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName      *string                                                                     `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type           *string                                                                     `json:"Type,omitempty" xml:"Type,omitempty"`
	WhereClause    *string                                                                     `json:"WhereClause,omitempty" xml:"WhereClause,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetAdbTableSchema(v *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.AdbTableSchema = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetFilterColumns(v []*string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.FilterColumns = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetMappedName(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.MappedName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetShardColumns(v []*string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.ShardColumns = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetTableId(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.TableId = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetTableName(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.TableName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetType(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.Type = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables) SetWhereClause(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTables {
	s.WhereClause = &v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema struct {
	DistributedKeys    []*string `json:"DistributedKeys,omitempty" xml:"DistributedKeys,omitempty" type:"Repeated"`
	PartitionLifeCycle *int32    `json:"PartitionLifeCycle,omitempty" xml:"PartitionLifeCycle,omitempty"`
	PartitionStatement *string   `json:"PartitionStatement,omitempty" xml:"PartitionStatement,omitempty"`
	PrimaryKeys        []*string `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty" type:"Repeated"`
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) SetDistributedKeys(v []*string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema {
	s.DistributedKeys = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) SetPartitionLifeCycle(v int32) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionLifeCycle = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) SetPartitionStatement(v string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionStatement = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema) SetPrimaryKeys(v []*string) *CreateOmsOpenAPIProjectRequestTransferMappingDatabasesTablesAdbTableSchema {
	s.PrimaryKeys = v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferStepConfig struct {
	EnableFullSync             *bool                                                                       `json:"EnableFullSync,omitempty" xml:"EnableFullSync,omitempty"`
	EnableIncrSync             *bool                                                                       `json:"EnableIncrSync,omitempty" xml:"EnableIncrSync,omitempty"`
	EnableStructSync           *bool                                                                       `json:"EnableStructSync,omitempty" xml:"EnableStructSync,omitempty"`
	IncrSyncStepTransferConfig *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig `json:"IncrSyncStepTransferConfig,omitempty" xml:"IncrSyncStepTransferConfig,omitempty" type:"Struct"`
}

func (s CreateOmsOpenAPIProjectRequestTransferStepConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferStepConfig) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfig) SetEnableFullSync(v bool) *CreateOmsOpenAPIProjectRequestTransferStepConfig {
	s.EnableFullSync = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfig) SetEnableIncrSync(v bool) *CreateOmsOpenAPIProjectRequestTransferStepConfig {
	s.EnableIncrSync = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfig) SetEnableStructSync(v bool) *CreateOmsOpenAPIProjectRequestTransferStepConfig {
	s.EnableStructSync = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfig) SetIncrSyncStepTransferConfig(v *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) *CreateOmsOpenAPIProjectRequestTransferStepConfig {
	s.IncrSyncStepTransferConfig = v
	return s
}

type CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig struct {
	RecordTypeList          []*string `json:"RecordTypeList,omitempty" xml:"RecordTypeList,omitempty" type:"Repeated"`
	StartTimestamp          *int64    `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StoreLogKeptHour        *int64    `json:"StoreLogKeptHour,omitempty" xml:"StoreLogKeptHour,omitempty"`
	StoreTransactionEnabled *bool     `json:"StoreTransactionEnabled,omitempty" xml:"StoreTransactionEnabled,omitempty"`
	TransferStepType        *string   `json:"TransferStepType,omitempty" xml:"TransferStepType,omitempty"`
}

func (s CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) SetRecordTypeList(v []*string) *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig {
	s.RecordTypeList = v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) SetStartTimestamp(v int64) *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig {
	s.StartTimestamp = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) SetStoreLogKeptHour(v int64) *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreLogKeptHour = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) SetStoreTransactionEnabled(v bool) *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreTransactionEnabled = &v
	return s
}

func (s *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig) SetTransferStepType(v string) *CreateOmsOpenAPIProjectRequestTransferStepConfigIncrSyncStepTransferConfig {
	s.TransferStepType = &v
	return s
}

type CreateOmsOpenAPIProjectShrinkRequest struct {
	BusinessName             *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	DestConfigShrink         *string `json:"DestConfig,omitempty" xml:"DestConfig,omitempty"`
	LabelIdsShrink           *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	PageNumber               *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectName              *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	SourceConfigShrink       *string `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty"`
	TransferMappingShrink    *string `json:"TransferMapping,omitempty" xml:"TransferMapping,omitempty"`
	TransferStepConfigShrink *string `json:"TransferStepConfig,omitempty" xml:"TransferStepConfig,omitempty"`
	WorkerGradeId            *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s CreateOmsOpenAPIProjectShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetBusinessName(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.BusinessName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetDestConfigShrink(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.DestConfigShrink = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetLabelIdsShrink(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.LabelIdsShrink = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetPageNumber(v int32) *CreateOmsOpenAPIProjectShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetPageSize(v int32) *CreateOmsOpenAPIProjectShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetProjectName(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetSourceConfigShrink(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.SourceConfigShrink = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetTransferMappingShrink(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.TransferMappingShrink = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetTransferStepConfigShrink(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.TransferStepConfigShrink = &v
	return s
}

func (s *CreateOmsOpenAPIProjectShrinkRequest) SetWorkerGradeId(v string) *CreateOmsOpenAPIProjectShrinkRequest {
	s.WorkerGradeId = &v
	return s
}

type CreateOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                         `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                         `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *string                                         `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *CreateOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetAdvice(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetCode(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetCost(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetData(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetErrorDetail(v *CreateOmsOpenAPIProjectResponseBodyErrorDetail) *CreateOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetMessage(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *CreateOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *CreateOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetRequestId(v string) *CreateOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *CreateOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *CreateOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type CreateOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s CreateOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *CreateOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *CreateOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *CreateOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *CreateOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type CreateOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *CreateOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateOmsOpenAPIProjectResponse) SetStatusCode(v int32) *CreateOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOmsOpenAPIProjectResponse) SetBody(v *CreateOmsOpenAPIProjectResponseBody) *CreateOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type CreateSecurityIpGroupRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIps         *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s CreateSecurityIpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupRequest) SetInstanceId(v string) *CreateSecurityIpGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSecurityIpGroupRequest) SetSecurityIpGroupName(v string) *CreateSecurityIpGroupRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *CreateSecurityIpGroupRequest) SetSecurityIps(v string) *CreateSecurityIpGroupRequest {
	s.SecurityIps = &v
	return s
}

type CreateSecurityIpGroupResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroup *CreateSecurityIpGroupResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s CreateSecurityIpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponseBody) SetRequestId(v string) *CreateSecurityIpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBody) SetSecurityIpGroup(v *CreateSecurityIpGroupResponseBodySecurityIpGroup) *CreateSecurityIpGroupResponseBody {
	s.SecurityIpGroup = v
	return s
}

type CreateSecurityIpGroupResponseBodySecurityIpGroup struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIps         *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s CreateSecurityIpGroupResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetInstanceId(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *CreateSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIps(v string) *CreateSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIps = &v
	return s
}

type CreateSecurityIpGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSecurityIpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSecurityIpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSecurityIpGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSecurityIpGroupResponse) SetHeaders(v map[string]*string) *CreateSecurityIpGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSecurityIpGroupResponse) SetStatusCode(v int32) *CreateSecurityIpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSecurityIpGroupResponse) SetBody(v *CreateSecurityIpGroupResponseBody) *CreateSecurityIpGroupResponse {
	s.Body = v
	return s
}

type CreateTenantRequest struct {
	Charset       *string `json:"Charset,omitempty" xml:"Charset,omitempty"`
	Cpu           *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Memory        *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	PrimaryZone   *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	TenantMode    *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	TenantName    *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	TimeZone      *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
	UnitNum       *int32  `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	UserVSwitchId *string `json:"UserVSwitchId,omitempty" xml:"UserVSwitchId,omitempty"`
	UserVpcId     *string `json:"UserVpcId,omitempty" xml:"UserVpcId,omitempty"`
}

func (s CreateTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantRequest) SetCharset(v string) *CreateTenantRequest {
	s.Charset = &v
	return s
}

func (s *CreateTenantRequest) SetCpu(v int32) *CreateTenantRequest {
	s.Cpu = &v
	return s
}

func (s *CreateTenantRequest) SetDescription(v string) *CreateTenantRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantRequest) SetInstanceId(v string) *CreateTenantRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantRequest) SetMemory(v int32) *CreateTenantRequest {
	s.Memory = &v
	return s
}

func (s *CreateTenantRequest) SetPrimaryZone(v string) *CreateTenantRequest {
	s.PrimaryZone = &v
	return s
}

func (s *CreateTenantRequest) SetTenantMode(v string) *CreateTenantRequest {
	s.TenantMode = &v
	return s
}

func (s *CreateTenantRequest) SetTenantName(v string) *CreateTenantRequest {
	s.TenantName = &v
	return s
}

func (s *CreateTenantRequest) SetTimeZone(v string) *CreateTenantRequest {
	s.TimeZone = &v
	return s
}

func (s *CreateTenantRequest) SetUnitNum(v int32) *CreateTenantRequest {
	s.UnitNum = &v
	return s
}

func (s *CreateTenantRequest) SetUserVSwitchId(v string) *CreateTenantRequest {
	s.UserVSwitchId = &v
	return s
}

func (s *CreateTenantRequest) SetUserVpcId(v string) *CreateTenantRequest {
	s.UserVpcId = &v
	return s
}

type CreateTenantResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s CreateTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantResponseBody) SetRequestId(v string) *CreateTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantResponseBody) SetTenantId(v string) *CreateTenantResponseBody {
	s.TenantId = &v
	return s
}

type CreateTenantResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantResponse) SetHeaders(v map[string]*string) *CreateTenantResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantResponse) SetStatusCode(v int32) *CreateTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantResponse) SetBody(v *CreateTenantResponseBody) *CreateTenantResponse {
	s.Body = v
	return s
}

type CreateTenantReadOnlyConnectionRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateTenantReadOnlyConnectionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionRequest) SetInstanceId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionRequest) SetTenantId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionRequest) SetZoneId(v string) *CreateTenantReadOnlyConnectionRequest {
	s.ZoneId = &v
	return s
}

type CreateTenantReadOnlyConnectionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTenantReadOnlyConnectionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionResponseBody) SetRequestId(v string) *CreateTenantReadOnlyConnectionResponseBody {
	s.RequestId = &v
	return s
}

type CreateTenantReadOnlyConnectionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTenantReadOnlyConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantReadOnlyConnectionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantReadOnlyConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantReadOnlyConnectionResponse) SetHeaders(v map[string]*string) *CreateTenantReadOnlyConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantReadOnlyConnectionResponse) SetStatusCode(v int32) *CreateTenantReadOnlyConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantReadOnlyConnectionResponse) SetBody(v *CreateTenantReadOnlyConnectionResponseBody) *CreateTenantReadOnlyConnectionResponse {
	s.Body = v
	return s
}

type CreateTenantUserRequest struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Roles        *string `json:"Roles,omitempty" xml:"Roles,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
	UserType     *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreateTenantUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserRequest) GoString() string {
	return s.String()
}

func (s *CreateTenantUserRequest) SetDescription(v string) *CreateTenantUserRequest {
	s.Description = &v
	return s
}

func (s *CreateTenantUserRequest) SetInstanceId(v string) *CreateTenantUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTenantUserRequest) SetRoles(v string) *CreateTenantUserRequest {
	s.Roles = &v
	return s
}

func (s *CreateTenantUserRequest) SetTenantId(v string) *CreateTenantUserRequest {
	s.TenantId = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserName(v string) *CreateTenantUserRequest {
	s.UserName = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserPassword(v string) *CreateTenantUserRequest {
	s.UserPassword = &v
	return s
}

func (s *CreateTenantUserRequest) SetUserType(v string) *CreateTenantUserRequest {
	s.UserType = &v
	return s
}

type CreateTenantUserResponseBody struct {
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantUser []*CreateTenantUserResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Repeated"`
}

func (s CreateTenantUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBody) SetRequestId(v string) *CreateTenantUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTenantUserResponseBody) SetTenantUser(v []*CreateTenantUserResponseBodyTenantUser) *CreateTenantUserResponseBody {
	s.TenantUser = v
	return s
}

type CreateTenantUserResponseBodyTenantUser struct {
	Roles      []*CreateTenantUserResponseBodyTenantUserRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	UserName   *string                                        `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserStatus *string                                        `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	UserType   *string                                        `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s CreateTenantUserResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBodyTenantUser) SetRoles(v []*CreateTenantUserResponseBodyTenantUserRoles) *CreateTenantUserResponseBodyTenantUser {
	s.Roles = v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserName(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserStatus(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserStatus = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUser) SetUserType(v string) *CreateTenantUserResponseBodyTenantUser {
	s.UserType = &v
	return s
}

type CreateTenantUserResponseBodyTenantUserRoles struct {
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s CreateTenantUserResponseBodyTenantUserRoles) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponseBodyTenantUserRoles) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponseBodyTenantUserRoles) SetDatabase(v string) *CreateTenantUserResponseBodyTenantUserRoles {
	s.Database = &v
	return s
}

func (s *CreateTenantUserResponseBodyTenantUserRoles) SetRole(v string) *CreateTenantUserResponseBodyTenantUserRoles {
	s.Role = &v
	return s
}

type CreateTenantUserResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTenantUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTenantUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTenantUserResponse) GoString() string {
	return s.String()
}

func (s *CreateTenantUserResponse) SetHeaders(v map[string]*string) *CreateTenantUserResponse {
	s.Headers = v
	return s
}

func (s *CreateTenantUserResponse) SetStatusCode(v int32) *CreateTenantUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTenantUserResponse) SetBody(v *CreateTenantUserResponseBody) *CreateTenantUserResponse {
	s.Body = v
	return s
}

type DeleteDatabasesRequest struct {
	DatabaseNames *string `json:"DatabaseNames,omitempty" xml:"DatabaseNames,omitempty"`
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId      *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesRequest) SetDatabaseNames(v string) *DeleteDatabasesRequest {
	s.DatabaseNames = &v
	return s
}

func (s *DeleteDatabasesRequest) SetInstanceId(v string) *DeleteDatabasesRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteDatabasesRequest) SetTenantId(v string) *DeleteDatabasesRequest {
	s.TenantId = &v
	return s
}

type DeleteDatabasesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesResponseBody) SetRequestId(v string) *DeleteDatabasesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDatabasesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DeleteDatabasesResponse) SetHeaders(v map[string]*string) *DeleteDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DeleteDatabasesResponse) SetStatusCode(v int32) *DeleteDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDatabasesResponse) SetBody(v *DeleteDatabasesResponseBody) *DeleteDatabasesResponse {
	s.Body = v
	return s
}

type DeleteInstancesRequest struct {
	BackupRetainMode *string `json:"BackupRetainMode,omitempty" xml:"BackupRetainMode,omitempty"`
	InstanceIds      *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
}

func (s DeleteInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) SetBackupRetainMode(v string) *DeleteInstancesRequest {
	s.BackupRetainMode = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceIds(v string) *DeleteInstancesRequest {
	s.InstanceIds = &v
	return s
}

type DeleteInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponse) SetHeaders(v map[string]*string) *DeleteInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancesResponse) SetStatusCode(v int32) *DeleteInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstancesResponse) SetBody(v *DeleteInstancesResponseBody) *DeleteInstancesResponse {
	s.Body = v
	return s
}

type DeleteOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s DeleteOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteOmsOpenAPIProjectRequest) SetPageNumber(v int32) *DeleteOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectRequest) SetPageSize(v int32) *DeleteOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectRequest) SetProjectId(v string) *DeleteOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *DeleteOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type DeleteOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                         `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                         `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                           `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *DeleteOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DeleteOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetAdvice(v string) *DeleteOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetCode(v string) *DeleteOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetCost(v string) *DeleteOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetData(v bool) *DeleteOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetErrorDetail(v *DeleteOmsOpenAPIProjectResponseBodyErrorDetail) *DeleteOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetMessage(v string) *DeleteOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *DeleteOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *DeleteOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetRequestId(v string) *DeleteOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *DeleteOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *DeleteOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type DeleteOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s DeleteOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s DeleteOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *DeleteOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *DeleteOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *DeleteOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *DeleteOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *DeleteOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type DeleteOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *DeleteOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponse) SetStatusCode(v int32) *DeleteOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteOmsOpenAPIProjectResponse) SetBody(v *DeleteOmsOpenAPIProjectResponseBody) *DeleteOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type DeleteSecurityIpGroupRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
}

func (s DeleteSecurityIpGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupRequest) SetInstanceId(v string) *DeleteSecurityIpGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSecurityIpGroupRequest) SetSecurityIpGroupName(v string) *DeleteSecurityIpGroupRequest {
	s.SecurityIpGroupName = &v
	return s
}

type DeleteSecurityIpGroupResponseBody struct {
	RequestId       *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroup *DeleteSecurityIpGroupResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s DeleteSecurityIpGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponseBody) SetRequestId(v string) *DeleteSecurityIpGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityIpGroupResponseBody) SetSecurityIpGroup(v *DeleteSecurityIpGroupResponseBodySecurityIpGroup) *DeleteSecurityIpGroupResponseBody {
	s.SecurityIpGroup = v
	return s
}

type DeleteSecurityIpGroupResponseBodySecurityIpGroup struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
}

func (s DeleteSecurityIpGroupResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponseBodySecurityIpGroup) SetInstanceId(v string) *DeleteSecurityIpGroupResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *DeleteSecurityIpGroupResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *DeleteSecurityIpGroupResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

type DeleteSecurityIpGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSecurityIpGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSecurityIpGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSecurityIpGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIpGroupResponse) SetHeaders(v map[string]*string) *DeleteSecurityIpGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSecurityIpGroupResponse) SetStatusCode(v int32) *DeleteSecurityIpGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSecurityIpGroupResponse) SetBody(v *DeleteSecurityIpGroupResponseBody) *DeleteSecurityIpGroupResponse {
	s.Body = v
	return s
}

type DeleteTenantUsersRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Users      *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s DeleteTenantUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersRequest) SetInstanceId(v string) *DeleteTenantUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTenantUsersRequest) SetTenantId(v string) *DeleteTenantUsersRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteTenantUsersRequest) SetUsers(v string) *DeleteTenantUsersRequest {
	s.Users = &v
	return s
}

type DeleteTenantUsersResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTenantUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersResponseBody) SetRequestId(v string) *DeleteTenantUsersResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTenantUsersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTenantUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTenantUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantUsersResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantUsersResponse) SetHeaders(v map[string]*string) *DeleteTenantUsersResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantUsersResponse) SetStatusCode(v int32) *DeleteTenantUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTenantUsersResponse) SetBody(v *DeleteTenantUsersResponseBody) *DeleteTenantUsersResponse {
	s.Body = v
	return s
}

type DeleteTenantsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantIds  *string `json:"TenantIds,omitempty" xml:"TenantIds,omitempty"`
}

func (s DeleteTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsRequest) GoString() string {
	return s.String()
}

func (s *DeleteTenantsRequest) SetInstanceId(v string) *DeleteTenantsRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteTenantsRequest) SetTenantIds(v string) *DeleteTenantsRequest {
	s.TenantIds = &v
	return s
}

type DeleteTenantsResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantIds []*string `json:"TenantIds,omitempty" xml:"TenantIds,omitempty" type:"Repeated"`
}

func (s DeleteTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTenantsResponseBody) SetRequestId(v string) *DeleteTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTenantsResponseBody) SetTenantIds(v []*string) *DeleteTenantsResponseBody {
	s.TenantIds = v
	return s
}

type DeleteTenantsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTenantsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTenantsResponse) SetHeaders(v map[string]*string) *DeleteTenantsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTenantsResponse) SetStatusCode(v int32) *DeleteTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTenantsResponse) SetBody(v *DeleteTenantsResponseBody) *DeleteTenantsResponse {
	s.Body = v
	return s
}

type DescribeAnomalySQLListRequest struct {
	AcceptLanguage  *string                `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	DbName          *string                `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime         *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp          *string                `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber      *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId           *string                `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord   *string                `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter *string                `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule      *string                `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue     *string                `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn      *string                `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder       *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime       *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId        *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAnomalySQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListRequest) SetAcceptLanguage(v string) *DescribeAnomalySQLListRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetDbName(v string) *DescribeAnomalySQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetEndTime(v string) *DescribeAnomalySQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeAnomalySQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetNodeIp(v string) *DescribeAnomalySQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetPageNumber(v int32) *DescribeAnomalySQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetPageSize(v int32) *DescribeAnomalySQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSQLId(v string) *DescribeAnomalySQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchKeyWord(v string) *DescribeAnomalySQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchParameter(v string) *DescribeAnomalySQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchRule(v string) *DescribeAnomalySQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSearchValue(v string) *DescribeAnomalySQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSortColumn(v string) *DescribeAnomalySQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetSortOrder(v string) *DescribeAnomalySQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetStartTime(v string) *DescribeAnomalySQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAnomalySQLListRequest) SetTenantId(v string) *DescribeAnomalySQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeAnomalySQLListShrinkRequest struct {
	AcceptLanguage        *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	DbName                *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp                *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId                 *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord         *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter       *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule            *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue           *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn            *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder             *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId              *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAnomalySQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListShrinkRequest) SetAcceptLanguage(v string) *DescribeAnomalySQLListShrinkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetDbName(v string) *DescribeAnomalySQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetEndTime(v string) *DescribeAnomalySQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeAnomalySQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetNodeIp(v string) *DescribeAnomalySQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetPageNumber(v int32) *DescribeAnomalySQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetPageSize(v int32) *DescribeAnomalySQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSQLId(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchParameter(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchRule(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSearchValue(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSortColumn(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetSortOrder(v string) *DescribeAnomalySQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetStartTime(v string) *DescribeAnomalySQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAnomalySQLListShrinkRequest) SetTenantId(v string) *DescribeAnomalySQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeAnomalySQLListResponseBody struct {
	AnomalySQLList []*DescribeAnomalySQLListResponseBodyAnomalySQLList `json:"AnomalySQLList,omitempty" xml:"AnomalySQLList,omitempty" type:"Repeated"`
	RequestId      *string                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount     *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAnomalySQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponseBody) SetAnomalySQLList(v []*DescribeAnomalySQLListResponseBodyAnomalySQLList) *DescribeAnomalySQLListResponseBody {
	s.AnomalySQLList = v
	return s
}

func (s *DescribeAnomalySQLListResponseBody) SetRequestId(v string) *DescribeAnomalySQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBody) SetTotalCount(v int64) *DescribeAnomalySQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeAnomalySQLListResponseBodyAnomalySQLList struct {
	CpuTime              *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	Diagnosis            *string  `json:"Diagnosis,omitempty" xml:"Diagnosis,omitempty"`
	DiagnosisRule        *string  `json:"DiagnosisRule,omitempty" xml:"DiagnosisRule,omitempty"`
	Executions           *int64   `json:"Executions,omitempty" xml:"Executions,omitempty"`
	Key                  *int64   `json:"Key,omitempty" xml:"Key,omitempty"`
	RequestTime          *float32 `json:"RequestTime,omitempty" xml:"RequestTime,omitempty"`
	RequestTimeUTCString *string  `json:"RequestTimeUTCString,omitempty" xml:"RequestTimeUTCString,omitempty"`
	SQLId                *string  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	Suggestion           *string  `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeAnomalySQLListResponseBodyAnomalySQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponseBodyAnomalySQLList) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetCpuTime(v float32) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDbName(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.DbName = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDiagnosis(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Diagnosis = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetDiagnosisRule(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.DiagnosisRule = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetExecutions(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Executions = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetKey(v int64) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Key = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetRequestTime(v float32) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.RequestTime = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetRequestTimeUTCString(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.RequestTimeUTCString = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSQLId(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSQLText(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetSuggestion(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.Suggestion = &v
	return s
}

func (s *DescribeAnomalySQLListResponseBodyAnomalySQLList) SetUserName(v string) *DescribeAnomalySQLListResponseBodyAnomalySQLList {
	s.UserName = &v
	return s
}

type DescribeAnomalySQLListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAnomalySQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAnomalySQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAnomalySQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAnomalySQLListResponse) SetHeaders(v map[string]*string) *DescribeAnomalySQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAnomalySQLListResponse) SetStatusCode(v int32) *DescribeAnomalySQLListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAnomalySQLListResponse) SetBody(v *DescribeAnomalySQLListResponseBody) *DescribeAnomalySQLListResponse {
	s.Body = v
	return s
}

type DescribeAvailableCpuResourceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeAvailableCpuResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceRequest) SetInstanceId(v string) *DescribeAvailableCpuResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableCpuResourceRequest) SetModifyType(v string) *DescribeAvailableCpuResourceRequest {
	s.ModifyType = &v
	return s
}

func (s *DescribeAvailableCpuResourceRequest) SetTenantId(v string) *DescribeAvailableCpuResourceRequest {
	s.TenantId = &v
	return s
}

type DescribeAvailableCpuResourceResponseBody struct {
	Data      []*DescribeAvailableCpuResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableCpuResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponseBody) SetData(v []*DescribeAvailableCpuResourceResponseBodyData) *DescribeAvailableCpuResourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBody) SetRequestId(v string) *DescribeAvailableCpuResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableCpuResourceResponseBodyData struct {
	MaxCpu  *int64 `json:"MaxCpu,omitempty" xml:"MaxCpu,omitempty"`
	MinCpu  *int64 `json:"MinCpu,omitempty" xml:"MinCpu,omitempty"`
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeAvailableCpuResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetMaxCpu(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.MaxCpu = &v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetMinCpu(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.MinCpu = &v
	return s
}

func (s *DescribeAvailableCpuResourceResponseBodyData) SetUnitNum(v int64) *DescribeAvailableCpuResourceResponseBodyData {
	s.UnitNum = &v
	return s
}

type DescribeAvailableCpuResourceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableCpuResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableCpuResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableCpuResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCpuResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableCpuResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableCpuResourceResponse) SetStatusCode(v int32) *DescribeAvailableCpuResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableCpuResourceResponse) SetBody(v *DescribeAvailableCpuResourceResponseBody) *DescribeAvailableCpuResourceResponse {
	s.Body = v
	return s
}

type DescribeAvailableMemResourceRequest struct {
	CpuNum     *int64  `json:"CpuNum,omitempty" xml:"CpuNum,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UnitNum    *int64  `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeAvailableMemResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceRequest) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceRequest) SetCpuNum(v int64) *DescribeAvailableMemResourceRequest {
	s.CpuNum = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetInstanceId(v string) *DescribeAvailableMemResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetTenantId(v string) *DescribeAvailableMemResourceRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeAvailableMemResourceRequest) SetUnitNum(v int64) *DescribeAvailableMemResourceRequest {
	s.UnitNum = &v
	return s
}

type DescribeAvailableMemResourceResponseBody struct {
	Data      *DescribeAvailableMemResourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAvailableMemResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponseBody) SetData(v *DescribeAvailableMemResourceResponseBodyData) *DescribeAvailableMemResourceResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableMemResourceResponseBody) SetRequestId(v string) *DescribeAvailableMemResourceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeAvailableMemResourceResponseBodyData struct {
	MaxMem  *int64 `json:"MaxMem,omitempty" xml:"MaxMem,omitempty"`
	MinMem  *int64 `json:"MinMem,omitempty" xml:"MinMem,omitempty"`
	UsedMem *int64 `json:"UsedMem,omitempty" xml:"UsedMem,omitempty"`
}

func (s DescribeAvailableMemResourceResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetMaxMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.MaxMem = &v
	return s
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetMinMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.MinMem = &v
	return s
}

func (s *DescribeAvailableMemResourceResponseBodyData) SetUsedMem(v int64) *DescribeAvailableMemResourceResponseBodyData {
	s.UsedMem = &v
	return s
}

type DescribeAvailableMemResourceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeAvailableMemResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeAvailableMemResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAvailableMemResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableMemResourceResponse) SetHeaders(v map[string]*string) *DescribeAvailableMemResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableMemResourceResponse) SetStatusCode(v int32) *DescribeAvailableMemResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableMemResourceResponse) SetBody(v *DescribeAvailableMemResourceResponseBody) *DescribeAvailableMemResourceResponse {
	s.Body = v
	return s
}

type DescribeCharsetRequest struct {
	TenantMode *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
}

func (s DescribeCharsetRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharsetRequest) SetTenantMode(v string) *DescribeCharsetRequest {
	s.TenantMode = &v
	return s
}

type DescribeCharsetResponseBody struct {
	Charset   []*DescribeCharsetResponseBodyCharset `json:"Charset,omitempty" xml:"Charset,omitempty" type:"Repeated"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCharsetResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponseBody) SetCharset(v []*DescribeCharsetResponseBodyCharset) *DescribeCharsetResponseBody {
	s.Charset = v
	return s
}

func (s *DescribeCharsetResponseBody) SetRequestId(v string) *DescribeCharsetResponseBody {
	s.RequestId = &v
	return s
}

type DescribeCharsetResponseBodyCharset struct {
	Charset    *string   `json:"Charset,omitempty" xml:"Charset,omitempty"`
	Collations []*string `json:"Collations,omitempty" xml:"Collations,omitempty" type:"Repeated"`
}

func (s DescribeCharsetResponseBodyCharset) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponseBodyCharset) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponseBodyCharset) SetCharset(v string) *DescribeCharsetResponseBodyCharset {
	s.Charset = &v
	return s
}

func (s *DescribeCharsetResponseBodyCharset) SetCollations(v []*string) *DescribeCharsetResponseBodyCharset {
	s.Collations = v
	return s
}

type DescribeCharsetResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeCharsetResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeCharsetResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeCharsetResponse) GoString() string {
	return s.String()
}

func (s *DescribeCharsetResponse) SetHeaders(v map[string]*string) *DescribeCharsetResponse {
	s.Headers = v
	return s
}

func (s *DescribeCharsetResponse) SetStatusCode(v int32) *DescribeCharsetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCharsetResponse) SetBody(v *DescribeCharsetResponseBody) *DescribeCharsetResponse {
	s.Body = v
	return s
}

type DescribeDatabasesRequest struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey    *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	WithTables   *bool   `json:"WithTables,omitempty" xml:"WithTables,omitempty"`
}

func (s DescribeDatabasesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesRequest) SetDatabaseName(v string) *DescribeDatabasesRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageNumber(v int32) *DescribeDatabasesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDatabasesRequest) SetPageSize(v int32) *DescribeDatabasesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDatabasesRequest) SetSearchKey(v string) *DescribeDatabasesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDatabasesRequest) SetTenantId(v string) *DescribeDatabasesRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeDatabasesRequest) SetWithTables(v bool) *DescribeDatabasesRequest {
	s.WithTables = &v
	return s
}

type DescribeDatabasesResponseBody struct {
	Databases  []*DescribeDatabasesResponseBodyDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDatabasesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBody) SetDatabases(v []*DescribeDatabasesResponseBodyDatabases) *DescribeDatabasesResponseBody {
	s.Databases = v
	return s
}

func (s *DescribeDatabasesResponseBody) SetRequestId(v string) *DescribeDatabasesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDatabasesResponseBody) SetTotalCount(v int32) *DescribeDatabasesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDatabasesResponseBodyDatabases struct {
	Collation    *string                                         `json:"Collation,omitempty" xml:"Collation,omitempty"`
	CreateTime   *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataSize     *float64                                        `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	DatabaseName *string                                         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	DbType       *string                                         `json:"DbType,omitempty" xml:"DbType,omitempty"`
	Description  *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	Encoding     *string                                         `json:"Encoding,omitempty" xml:"Encoding,omitempty"`
	RequiredSize *float64                                        `json:"RequiredSize,omitempty" xml:"RequiredSize,omitempty"`
	Status       *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tables       []*DescribeDatabasesResponseBodyDatabasesTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TenantId     *string                                         `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Users        []*DescribeDatabasesResponseBodyDatabasesUsers  `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s DescribeDatabasesResponseBodyDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabases) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabases) SetCollation(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Collation = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetCreateTime(v string) *DescribeDatabasesResponseBodyDatabases {
	s.CreateTime = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDataSize(v float64) *DescribeDatabasesResponseBodyDatabases {
	s.DataSize = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDatabaseName(v string) *DescribeDatabasesResponseBodyDatabases {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDbType(v string) *DescribeDatabasesResponseBodyDatabases {
	s.DbType = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetDescription(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Description = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetEncoding(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Encoding = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetRequiredSize(v float64) *DescribeDatabasesResponseBodyDatabases {
	s.RequiredSize = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetStatus(v string) *DescribeDatabasesResponseBodyDatabases {
	s.Status = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetTables(v []*DescribeDatabasesResponseBodyDatabasesTables) *DescribeDatabasesResponseBodyDatabases {
	s.Tables = v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetTenantId(v string) *DescribeDatabasesResponseBodyDatabases {
	s.TenantId = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabases) SetUsers(v []*DescribeDatabasesResponseBodyDatabasesUsers) *DescribeDatabasesResponseBodyDatabases {
	s.Users = v
	return s
}

type DescribeDatabasesResponseBodyDatabasesTables struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesTables) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesTables) SetTableName(v string) *DescribeDatabasesResponseBodyDatabasesTables {
	s.TableName = &v
	return s
}

type DescribeDatabasesResponseBodyDatabasesUsers struct {
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserType *string `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeDatabasesResponseBodyDatabasesUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponseBodyDatabasesUsers) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetRole(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.Role = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetUserName(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.UserName = &v
	return s
}

func (s *DescribeDatabasesResponseBodyDatabasesUsers) SetUserType(v string) *DescribeDatabasesResponseBodyDatabasesUsers {
	s.UserType = &v
	return s
}

type DescribeDatabasesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDatabasesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDatabasesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDatabasesResponse) SetHeaders(v map[string]*string) *DescribeDatabasesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDatabasesResponse) SetStatusCode(v int32) *DescribeDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDatabasesResponse) SetBody(v *DescribeDatabasesResponseBody) *DescribeDatabasesResponse {
	s.Body = v
	return s
}

type DescribeInstanceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRequest) SetInstanceId(v string) *DescribeInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageNumber(v int32) *DescribeInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceRequest) SetPageSize(v int32) *DescribeInstanceRequest {
	s.PageSize = &v
	return s
}

type DescribeInstanceResponseBody struct {
	Instance  *DescribeInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBody) SetInstance(v *DescribeInstanceResponseBodyInstance) *DescribeInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeInstanceResponseBody) SetRequestId(v string) *DescribeInstanceResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceResponseBodyInstance struct {
	AutoRenewal          *bool                                         `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	AutoUpgradeObVersion *bool                                         `json:"AutoUpgradeObVersion,omitempty" xml:"AutoUpgradeObVersion,omitempty"`
	AvailableZones       []*string                                     `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	CreateTime           *string                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DataMergeTime        *string                                       `json:"DataMergeTime,omitempty" xml:"DataMergeTime,omitempty"`
	DeployMode           *string                                       `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	DeployType           *string                                       `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskType             *string                                       `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EnableUpgradeLogDisk *bool                                         `json:"EnableUpgradeLogDisk,omitempty" xml:"EnableUpgradeLogDisk,omitempty"`
	ExpireTime           *string                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceClass        *string                                       `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId           *string                                       `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                                       `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	IsLatestObVersion    *bool                                         `json:"IsLatestObVersion,omitempty" xml:"IsLatestObVersion,omitempty"`
	IsTrustEcs           *bool                                         `json:"IsTrustEcs,omitempty" xml:"IsTrustEcs,omitempty"`
	MaintainTime         *string                                       `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	ObRpmVersion         *string                                       `json:"ObRpmVersion,omitempty" xml:"ObRpmVersion,omitempty"`
	PayType              *string                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Resource             *DescribeInstanceResponseBodyInstanceResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	Series               *string                                       `json:"Series,omitempty" xml:"Series,omitempty"`
	Status               *string                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	Version              *string                                       `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeInstanceResponseBodyInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstance) SetAutoRenewal(v bool) *DescribeInstanceResponseBodyInstance {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetAutoUpgradeObVersion(v bool) *DescribeInstanceResponseBodyInstance {
	s.AutoUpgradeObVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetAvailableZones(v []*string) *DescribeInstanceResponseBodyInstance {
	s.AvailableZones = v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetCreateTime(v string) *DescribeInstanceResponseBodyInstance {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDataMergeTime(v string) *DescribeInstanceResponseBodyInstance {
	s.DataMergeTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDeployMode(v string) *DescribeInstanceResponseBodyInstance {
	s.DeployMode = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDeployType(v string) *DescribeInstanceResponseBodyInstance {
	s.DeployType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetDiskType(v string) *DescribeInstanceResponseBodyInstance {
	s.DiskType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetEnableUpgradeLogDisk(v bool) *DescribeInstanceResponseBodyInstance {
	s.EnableUpgradeLogDisk = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetExpireTime(v string) *DescribeInstanceResponseBodyInstance {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceClass(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceId(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetInstanceName(v string) *DescribeInstanceResponseBodyInstance {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetIsLatestObVersion(v bool) *DescribeInstanceResponseBodyInstance {
	s.IsLatestObVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetIsTrustEcs(v bool) *DescribeInstanceResponseBodyInstance {
	s.IsTrustEcs = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetMaintainTime(v string) *DescribeInstanceResponseBodyInstance {
	s.MaintainTime = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetObRpmVersion(v string) *DescribeInstanceResponseBodyInstance {
	s.ObRpmVersion = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetPayType(v string) *DescribeInstanceResponseBodyInstance {
	s.PayType = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetResource(v *DescribeInstanceResponseBodyInstanceResource) *DescribeInstanceResponseBodyInstance {
	s.Resource = v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetSeries(v string) *DescribeInstanceResponseBodyInstance {
	s.Series = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetStatus(v string) *DescribeInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstance) SetVersion(v string) *DescribeInstanceResponseBodyInstance {
	s.Version = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResource struct {
	Cpu         *DescribeInstanceResponseBodyInstanceResourceCpu         `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	DiskSize    *DescribeInstanceResponseBodyInstanceResourceDiskSize    `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	LogDiskSize *DescribeInstanceResponseBodyInstanceResourceLogDiskSize `json:"LogDiskSize,omitempty" xml:"LogDiskSize,omitempty" type:"Struct"`
	Memory      *DescribeInstanceResponseBodyInstanceResourceMemory      `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	UnitCount   *int64                                                   `json:"UnitCount,omitempty" xml:"UnitCount,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResource) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetCpu(v *DescribeInstanceResponseBodyInstanceResourceCpu) *DescribeInstanceResponseBodyInstanceResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetDiskSize(v *DescribeInstanceResponseBodyInstanceResourceDiskSize) *DescribeInstanceResponseBodyInstanceResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetLogDiskSize(v *DescribeInstanceResponseBodyInstanceResourceLogDiskSize) *DescribeInstanceResponseBodyInstanceResource {
	s.LogDiskSize = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetMemory(v *DescribeInstanceResponseBodyInstanceResourceMemory) *DescribeInstanceResponseBodyInstanceResource {
	s.Memory = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResource) SetUnitCount(v int64) *DescribeInstanceResponseBodyInstanceResource {
	s.UnitCount = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceCpu struct {
	TotalCpu *int64 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	UnitCpu  *int64 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	UsedCpu  *int64 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetTotalCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetUnitCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceCpu) SetUsedCpu(v int64) *DescribeInstanceResponseBodyInstanceResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceDiskSize struct {
	DataUsedSize        *float64  `json:"DataUsedSize,omitempty" xml:"DataUsedSize,omitempty"`
	MaxDiskUsedObServer []*string `json:"MaxDiskUsedObServer,omitempty" xml:"MaxDiskUsedObServer,omitempty" type:"Repeated"`
	MaxDiskUsedPercent  *float64  `json:"MaxDiskUsedPercent,omitempty" xml:"MaxDiskUsedPercent,omitempty"`
	TotalDiskSize       *int64    `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	UnitDiskSize        *int64    `json:"UnitDiskSize,omitempty" xml:"UnitDiskSize,omitempty"`
	UsedDiskSize        *int64    `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetDataUsedSize(v float64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.DataUsedSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetMaxDiskUsedObServer(v []*string) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.MaxDiskUsedObServer = v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetMaxDiskUsedPercent(v float64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.MaxDiskUsedPercent = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetTotalDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetUnitDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.UnitDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceDiskSize) SetUsedDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceLogDiskSize struct {
	TotalDiskSize *int64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	UnitDiskSize  *int64 `json:"UnitDiskSize,omitempty" xml:"UnitDiskSize,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceLogDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceLogDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceLogDiskSize) SetTotalDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceLogDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceLogDiskSize) SetUnitDiskSize(v int64) *DescribeInstanceResponseBodyInstanceResourceLogDiskSize {
	s.UnitDiskSize = &v
	return s
}

type DescribeInstanceResponseBodyInstanceResourceMemory struct {
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UnitMemory  *int64 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	UsedMemory  *int64 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstanceResponseBodyInstanceResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponseBodyInstanceResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetTotalMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetUnitMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstanceResponseBodyInstanceResourceMemory) SetUsedMemory(v int64) *DescribeInstanceResponseBodyInstanceResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceResponse) SetHeaders(v map[string]*string) *DescribeInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceResponse) SetStatusCode(v int32) *DescribeInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceResponse) SetBody(v *DescribeInstanceResponseBody) *DescribeInstanceResponse {
	s.Body = v
	return s
}

type DescribeInstanceCreatableZoneRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceCreatableZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneRequest) SetInstanceId(v string) *DescribeInstanceCreatableZoneRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceCreatableZoneResponseBody struct {
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ZoneList  []*DescribeInstanceCreatableZoneResponseBodyZoneList `json:"ZoneList,omitempty" xml:"ZoneList,omitempty" type:"Repeated"`
}

func (s DescribeInstanceCreatableZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponseBody) SetRequestId(v string) *DescribeInstanceCreatableZoneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceCreatableZoneResponseBody) SetZoneList(v []*DescribeInstanceCreatableZoneResponseBodyZoneList) *DescribeInstanceCreatableZoneResponseBody {
	s.ZoneList = v
	return s
}

type DescribeInstanceCreatableZoneResponseBodyZoneList struct {
	IsInCluster *bool   `json:"IsInCluster,omitempty" xml:"IsInCluster,omitempty"`
	Zone        *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeInstanceCreatableZoneResponseBodyZoneList) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponseBodyZoneList) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponseBodyZoneList) SetIsInCluster(v bool) *DescribeInstanceCreatableZoneResponseBodyZoneList {
	s.IsInCluster = &v
	return s
}

func (s *DescribeInstanceCreatableZoneResponseBodyZoneList) SetZone(v string) *DescribeInstanceCreatableZoneResponseBodyZoneList {
	s.Zone = &v
	return s
}

type DescribeInstanceCreatableZoneResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceCreatableZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceCreatableZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceCreatableZoneResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceCreatableZoneResponse) SetHeaders(v map[string]*string) *DescribeInstanceCreatableZoneResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceCreatableZoneResponse) SetStatusCode(v int32) *DescribeInstanceCreatableZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceCreatableZoneResponse) SetBody(v *DescribeInstanceCreatableZoneResponseBody) *DescribeInstanceCreatableZoneResponse {
	s.Body = v
	return s
}

type DescribeInstanceTenantModesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceTenantModesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesRequest) SetInstanceId(v string) *DescribeInstanceTenantModesRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceTenantModesResponseBody struct {
	InstanceModes []*string `json:"InstanceModes,omitempty" xml:"InstanceModes,omitempty" type:"Repeated"`
	RequestId     *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTenantModesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesResponseBody) SetInstanceModes(v []*string) *DescribeInstanceTenantModesResponseBody {
	s.InstanceModes = v
	return s
}

func (s *DescribeInstanceTenantModesResponseBody) SetRequestId(v string) *DescribeInstanceTenantModesResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTenantModesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceTenantModesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTenantModesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTenantModesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTenantModesResponse) SetHeaders(v map[string]*string) *DescribeInstanceTenantModesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTenantModesResponse) SetStatusCode(v int32) *DescribeInstanceTenantModesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTenantModesResponse) SetBody(v *DescribeInstanceTenantModesResponseBody) *DescribeInstanceTenantModesResponse {
	s.Body = v
	return s
}

type DescribeInstanceTopologyRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeInstanceTopologyRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyRequest) SetInstanceId(v string) *DescribeInstanceTopologyRequest {
	s.InstanceId = &v
	return s
}

type DescribeInstanceTopologyResponseBody struct {
	InstanceTopology *DescribeInstanceTopologyResponseBodyInstanceTopology `json:"InstanceTopology,omitempty" xml:"InstanceTopology,omitempty" type:"Struct"`
	RequestId        *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTopologyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBody) SetInstanceTopology(v *DescribeInstanceTopologyResponseBodyInstanceTopology) *DescribeInstanceTopologyResponseBody {
	s.InstanceTopology = v
	return s
}

func (s *DescribeInstanceTopologyResponseBody) SetRequestId(v string) *DescribeInstanceTopologyResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopology struct {
	Tenants []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenants `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	Zones   []*DescribeInstanceTopologyResponseBodyInstanceTopologyZones   `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopology) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopology) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopology) SetTenants(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) *DescribeInstanceTopologyResponseBodyInstanceTopology {
	s.Tenants = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopology) SetZones(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZones) *DescribeInstanceTopologyResponseBodyInstanceTopology {
	s.Zones = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenants struct {
	PrimaryZoneDeployType *string                                                                   `json:"PrimaryZoneDeployType,omitempty" xml:"PrimaryZoneDeployType,omitempty"`
	TenantCpu             *float32                                                                  `json:"TenantCpu,omitempty" xml:"TenantCpu,omitempty"`
	TenantDeployType      *string                                                                   `json:"TenantDeployType,omitempty" xml:"TenantDeployType,omitempty"`
	TenantId              *string                                                                   `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantMemory          *float32                                                                  `json:"TenantMemory,omitempty" xml:"TenantMemory,omitempty"`
	TenantMode            *string                                                                   `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	TenantName            *string                                                                   `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	TenantStatus          *string                                                                   `json:"TenantStatus,omitempty" xml:"TenantStatus,omitempty"`
	TenantUnitNum         *int32                                                                    `json:"TenantUnitNum,omitempty" xml:"TenantUnitNum,omitempty"`
	TenantZones           []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetPrimaryZoneDeployType(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.PrimaryZoneDeployType = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantDeployType(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantDeployType = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantMode(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantMode = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantName(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantName = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantStatus = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantUnitNum(v int32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantUnitNum = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants) SetTenantZones(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenants {
	s.TenantZones = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones struct {
	IsPrimaryTenantZone *string                                                                        `json:"IsPrimaryTenantZone,omitempty" xml:"IsPrimaryTenantZone,omitempty"`
	TenantZoneId        *string                                                                        `json:"TenantZoneId,omitempty" xml:"TenantZoneId,omitempty"`
	TenantZoneRole      *string                                                                        `json:"TenantZoneRole,omitempty" xml:"TenantZoneRole,omitempty"`
	Units               []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits `json:"Units,omitempty" xml:"Units,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetIsPrimaryTenantZone(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.IsPrimaryTenantZone = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetTenantZoneId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.TenantZoneId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetTenantZoneRole(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.TenantZoneRole = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones) SetUnits(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZones {
	s.Units = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits struct {
	EnableCancelMigrateUnit *bool    `json:"EnableCancelMigrateUnit,omitempty" xml:"EnableCancelMigrateUnit,omitempty"`
	EnableMigrateUnit       *bool    `json:"EnableMigrateUnit,omitempty" xml:"EnableMigrateUnit,omitempty"`
	ManualMigrate           *bool    `json:"ManualMigrate,omitempty" xml:"ManualMigrate,omitempty"`
	NodeId                  *string  `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	UnitCpu                 *float32 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	UnitDataSize            *int64   `json:"UnitDataSize,omitempty" xml:"UnitDataSize,omitempty"`
	UnitId                  *string  `json:"UnitId,omitempty" xml:"UnitId,omitempty"`
	UnitMemory              *float32 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	UnitStatus              *string  `json:"UnitStatus,omitempty" xml:"UnitStatus,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetEnableCancelMigrateUnit(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.EnableCancelMigrateUnit = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetEnableMigrateUnit(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.EnableMigrateUnit = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetManualMigrate(v bool) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.ManualMigrate = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetNodeId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitDataSize(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitDataSize = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits) SetUnitStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyTenantsTenantZonesUnits {
	s.UnitStatus = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZones struct {
	Nodes        []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes      `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Region       *string                                                                `json:"Region,omitempty" xml:"Region,omitempty"`
	ZoneDisk     *string                                                                `json:"ZoneDisk,omitempty" xml:"ZoneDisk,omitempty"`
	ZoneId       *string                                                                `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneResource *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource `json:"ZoneResource,omitempty" xml:"ZoneResource,omitempty" type:"Struct"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZones) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetNodes(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.Nodes = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetRegion(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.Region = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetZoneDisk(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.ZoneDisk = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetZoneId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZones) SetZoneResource(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource) *DescribeInstanceTopologyResponseBodyInstanceTopologyZones {
	s.ZoneResource = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes struct {
	NodeCopyId   *int64                                                                        `json:"NodeCopyId,omitempty" xml:"NodeCopyId,omitempty"`
	NodeId       *string                                                                       `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeResource []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource `json:"NodeResource,omitempty" xml:"NodeResource,omitempty" type:"Repeated"`
	NodeStatus   *string                                                                       `json:"NodeStatus,omitempty" xml:"NodeStatus,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeCopyId(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeCopyId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeId(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeId = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeResource(v []*DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeResource = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes) SetNodeStatus(v string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodes {
	s.NodeStatus = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource struct {
	Cpu      *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu      `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	DiskSize *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	Memory   *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory   `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetCpu(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetDiskSize(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource) SetMemory(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResource {
	s.Memory = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu struct {
	TotalCpu *int32   `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	UsedCpu  *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) SetTotalCpu(v int32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu) SetUsedCpu(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize struct {
	TotalDiskSize *float64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	UsedDiskSize  *float64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) SetTotalDiskSize(v float64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize) SetUsedDiskSize(v float64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory struct {
	TotalMemory *int64   `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UsedMemory  *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) SetTotalMemory(v int64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory) SetUsedMemory(v float32) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesNodesNodeResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource struct {
	DiskSize *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource) SetDiskSize(v *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResource {
	s.DiskSize = v
	return s
}

type DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize struct {
	MaxDiskUsedObServer []*string `json:"MaxDiskUsedObServer,omitempty" xml:"MaxDiskUsedObServer,omitempty" type:"Repeated"`
	MaxDiskUsedPercent  *float64  `json:"MaxDiskUsedPercent,omitempty" xml:"MaxDiskUsedPercent,omitempty"`
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize) SetMaxDiskUsedObServer(v []*string) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize {
	s.MaxDiskUsedObServer = v
	return s
}

func (s *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize) SetMaxDiskUsedPercent(v float64) *DescribeInstanceTopologyResponseBodyInstanceTopologyZonesZoneResourceDiskSize {
	s.MaxDiskUsedPercent = &v
	return s
}

type DescribeInstanceTopologyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstanceTopologyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstanceTopologyResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstanceTopologyResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTopologyResponse) SetHeaders(v map[string]*string) *DescribeInstanceTopologyResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceTopologyResponse) SetStatusCode(v int32) *DescribeInstanceTopologyResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceTopologyResponse) SetBody(v *DescribeInstanceTopologyResponseBody) *DescribeInstanceTopologyResponse {
	s.Body = v
	return s
}

type DescribeInstancesRequest struct {
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SearchKey       *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
}

func (s DescribeInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstancesRequest) SetInstanceId(v string) *DescribeInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesRequest) SetInstanceName(v string) *DescribeInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageNumber(v int32) *DescribeInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesRequest) SetPageSize(v int32) *DescribeInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesRequest) SetResourceGroupId(v string) *DescribeInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesRequest) SetSearchKey(v string) *DescribeInstancesRequest {
	s.SearchKey = &v
	return s
}

type DescribeInstancesResponseBody struct {
	Instances  []*DescribeInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) SetInstances(v []*DescribeInstancesResponseBodyInstances) *DescribeInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeInstancesResponseBodyInstances struct {
	AvailableZones     []*string                                       `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	CommodityCode      *string                                         `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	Cpu                *int32                                          `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime         *string                                         `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode         *string                                         `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	DeployType         *string                                         `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	DiskSize           *string                                         `json:"DiskSize,omitempty" xml:"DiskSize,omitempty"`
	DiskType           *string                                         `json:"DiskType,omitempty" xml:"DiskType,omitempty"`
	EnableUpgradeNodes *bool                                           `json:"EnableUpgradeNodes,omitempty" xml:"EnableUpgradeNodes,omitempty"`
	ExpireSeconds      *int32                                          `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	ExpireTime         *string                                         `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	InstanceClass      *string                                         `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	InstanceId         *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName       *string                                         `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	MaintainTime       *string                                         `json:"MaintainTime,omitempty" xml:"MaintainTime,omitempty"`
	Mem                *int64                                          `json:"Mem,omitempty" xml:"Mem,omitempty"`
	PayType            *string                                         `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Resource           *DescribeInstancesResponseBodyInstancesResource `json:"Resource,omitempty" xml:"Resource,omitempty" type:"Struct"`
	ResourceGroupId    *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityIps        []*string                                       `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty" type:"Repeated"`
	Series             *string                                         `json:"Series,omitempty" xml:"Series,omitempty"`
	State              *string                                         `json:"State,omitempty" xml:"State,omitempty"`
	UsedDiskSize       *int64                                          `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
	Version            *string                                         `json:"Version,omitempty" xml:"Version,omitempty"`
	VpcId              *string                                         `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstances) SetAvailableZones(v []*string) *DescribeInstancesResponseBodyInstances {
	s.AvailableZones = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCommodityCode(v string) *DescribeInstancesResponseBodyInstances {
	s.CommodityCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCpu(v int32) *DescribeInstancesResponseBodyInstances {
	s.Cpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetCreateTime(v string) *DescribeInstancesResponseBodyInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDeployMode(v string) *DescribeInstancesResponseBodyInstances {
	s.DeployMode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDeployType(v string) *DescribeInstancesResponseBodyInstances {
	s.DeployType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDiskSize(v string) *DescribeInstancesResponseBodyInstances {
	s.DiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetDiskType(v string) *DescribeInstancesResponseBodyInstances {
	s.DiskType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetEnableUpgradeNodes(v bool) *DescribeInstancesResponseBodyInstances {
	s.EnableUpgradeNodes = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireSeconds(v int32) *DescribeInstancesResponseBodyInstances {
	s.ExpireSeconds = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetExpireTime(v string) *DescribeInstancesResponseBodyInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceClass(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceId(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetInstanceName(v string) *DescribeInstancesResponseBodyInstances {
	s.InstanceName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMaintainTime(v string) *DescribeInstancesResponseBodyInstances {
	s.MaintainTime = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetMem(v int64) *DescribeInstancesResponseBodyInstances {
	s.Mem = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetPayType(v string) *DescribeInstancesResponseBodyInstances {
	s.PayType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResource(v *DescribeInstancesResponseBodyInstancesResource) *DescribeInstancesResponseBodyInstances {
	s.Resource = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetResourceGroupId(v string) *DescribeInstancesResponseBodyInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSecurityIps(v []*string) *DescribeInstancesResponseBodyInstances {
	s.SecurityIps = v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetSeries(v string) *DescribeInstancesResponseBodyInstances {
	s.Series = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetState(v string) *DescribeInstancesResponseBodyInstances {
	s.State = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetUsedDiskSize(v int64) *DescribeInstancesResponseBodyInstances {
	s.UsedDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVersion(v string) *DescribeInstancesResponseBodyInstances {
	s.Version = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstances) SetVpcId(v string) *DescribeInstancesResponseBodyInstances {
	s.VpcId = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResource struct {
	Cpu       *DescribeInstancesResponseBodyInstancesResourceCpu      `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	DiskSize  *DescribeInstancesResponseBodyInstancesResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	Memory    *DescribeInstancesResponseBodyInstancesResourceMemory   `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	UnitCount *int64                                                  `json:"UnitCount,omitempty" xml:"UnitCount,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResource) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetCpu(v *DescribeInstancesResponseBodyInstancesResourceCpu) *DescribeInstancesResponseBodyInstancesResource {
	s.Cpu = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetDiskSize(v *DescribeInstancesResponseBodyInstancesResourceDiskSize) *DescribeInstancesResponseBodyInstancesResource {
	s.DiskSize = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetMemory(v *DescribeInstancesResponseBodyInstancesResourceMemory) *DescribeInstancesResponseBodyInstancesResource {
	s.Memory = v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResource) SetUnitCount(v int64) *DescribeInstancesResponseBodyInstancesResource {
	s.UnitCount = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceCpu struct {
	TotalCpu *int64 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	UnitCpu  *int64 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	UsedCpu  *int64 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetTotalCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetUnitCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceCpu) SetUsedCpu(v int64) *DescribeInstancesResponseBodyInstancesResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceDiskSize struct {
	TotalDiskSize *int64 `json:"TotalDiskSize,omitempty" xml:"TotalDiskSize,omitempty"`
	UnitDiskSize  *int64 `json:"UnitDiskSize,omitempty" xml:"UnitDiskSize,omitempty"`
	UsedDiskSize  *int64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetTotalDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.TotalDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetUnitDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.UnitDiskSize = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceDiskSize) SetUsedDiskSize(v int64) *DescribeInstancesResponseBodyInstancesResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeInstancesResponseBodyInstancesResourceMemory struct {
	TotalMemory *int64 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UnitMemory  *int64 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	UsedMemory  *int64 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeInstancesResponseBodyInstancesResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstancesResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetTotalMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetUnitMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstancesResourceMemory) SetUsedMemory(v int64) *DescribeInstancesResponseBodyInstancesResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeInstancesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponse) SetHeaders(v map[string]*string) *DescribeInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstancesResponse) SetStatusCode(v int32) *DescribeInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstancesResponse) SetBody(v *DescribeInstancesResponseBody) *DescribeInstancesResponse {
	s.Body = v
	return s
}

type DescribeNodeMetricsRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Metrics    *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	NodeIdList *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	NodeName   *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeNodeMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsRequest) SetEndTime(v string) *DescribeNodeMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetInstanceId(v string) *DescribeNodeMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetMetrics(v string) *DescribeNodeMetricsRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetNodeIdList(v string) *DescribeNodeMetricsRequest {
	s.NodeIdList = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetNodeName(v string) *DescribeNodeMetricsRequest {
	s.NodeName = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetPageNumber(v int32) *DescribeNodeMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetPageSize(v int32) *DescribeNodeMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetStartTime(v string) *DescribeNodeMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNodeMetricsRequest) SetTenantId(v string) *DescribeNodeMetricsRequest {
	s.TenantId = &v
	return s
}

type DescribeNodeMetricsResponseBody struct {
	NodeMetrics *string `json:"NodeMetrics,omitempty" xml:"NodeMetrics,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNodeMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsResponseBody) SetNodeMetrics(v string) *DescribeNodeMetricsResponseBody {
	s.NodeMetrics = &v
	return s
}

func (s *DescribeNodeMetricsResponseBody) SetRequestId(v string) *DescribeNodeMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeMetricsResponseBody) SetTotalCount(v int32) *DescribeNodeMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeNodeMetricsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeNodeMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeNodeMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeMetricsResponse) SetHeaders(v map[string]*string) *DescribeNodeMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeMetricsResponse) SetStatusCode(v int32) *DescribeNodeMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeMetricsResponse) SetBody(v *DescribeNodeMetricsResponseBody) *DescribeNodeMetricsResponse {
	s.Body = v
	return s
}

type DescribeOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s DescribeOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectRequest) SetPageNumber(v int32) *DescribeOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectRequest) SetPageSize(v int32) *DescribeOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectRequest) SetProjectId(v string) *DescribeOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *DescribeOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                           `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                           `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *DescribeOmsOpenAPIProjectResponseBodyData        `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	ErrorDetail *DescribeOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                           `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetAdvice(v string) *DescribeOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetCode(v string) *DescribeOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetCost(v string) *DescribeOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetData(v *DescribeOmsOpenAPIProjectResponseBodyData) *DescribeOmsOpenAPIProjectResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetErrorDetail(v *DescribeOmsOpenAPIProjectResponseBodyErrorDetail) *DescribeOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetMessage(v string) *DescribeOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *DescribeOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *DescribeOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetRequestId(v string) *DescribeOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *DescribeOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *DescribeOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyData struct {
	BusinessName       *string                                                      `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	DestConfig         *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig         `json:"DestConfig,omitempty" xml:"DestConfig,omitempty" type:"Struct"`
	Labels             []*DescribeOmsOpenAPIProjectResponseBodyDataLabels           `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ProjectId          *string                                                      `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string                                                      `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectOwner       *string                                                      `json:"ProjectOwner,omitempty" xml:"ProjectOwner,omitempty"`
	SourceConfig       *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig       `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty" type:"Struct"`
	Steps              []*DescribeOmsOpenAPIProjectResponseBodyDataSteps            `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	TransferMapping    *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping    `json:"TransferMapping,omitempty" xml:"TransferMapping,omitempty" type:"Struct"`
	TransferStepConfig *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig `json:"TransferStepConfig,omitempty" xml:"TransferStepConfig,omitempty" type:"Struct"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetBusinessName(v string) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetDestConfig(v *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.DestConfig = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetLabels(v []*DescribeOmsOpenAPIProjectResponseBodyDataLabels) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.Labels = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetProjectId(v string) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetProjectName(v string) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetProjectOwner(v string) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.ProjectOwner = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetSourceConfig(v *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.SourceConfig = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetSteps(v []*DescribeOmsOpenAPIProjectResponseBodyDataSteps) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.Steps = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetTransferMapping(v *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.TransferMapping = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyData) SetTransferStepConfig(v *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) *DescribeOmsOpenAPIProjectResponseBodyData {
	s.TransferStepConfig = v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataDestConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetEnableMsgTrace(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetEndpointId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.EndpointId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetEndpointType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.EndpointType = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetMsgTags(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.MsgTags = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetPartition(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.Partition = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetPartitionMode(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.PartitionMode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetProducerGroup(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.ProducerGroup = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetSendMsgTimeout(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetSequenceEnable(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.SequenceEnable = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetSequenceStartTimestamp(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetSerializerType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.SerializerType = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig) SetTopicType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataDestConfig {
	s.TopicType = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataLabels struct {
	Count   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataLabels) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataLabels) SetCount(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataLabels {
	s.Count = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataLabels) SetCreator(v string) *DescribeOmsOpenAPIProjectResponseBodyDataLabels {
	s.Creator = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataLabels) SetId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataLabels {
	s.Id = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataLabels) SetName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataLabels {
	s.Name = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetEnableMsgTrace(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetEndpointId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.EndpointId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetEndpointType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.EndpointType = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetMsgTags(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.MsgTags = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetPartition(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetPartitionMode(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.PartitionMode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetProducerGroup(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.ProducerGroup = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetSendMsgTimeout(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetSequenceEnable(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.SequenceEnable = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetSequenceStartTimestamp(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetSerializerType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.SerializerType = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig) SetTopicType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSourceConfig {
	s.TopicType = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataSteps struct {
	EstimatedRemainingSeconds *int64                                                   `json:"EstimatedRemainingSeconds,omitempty" xml:"EstimatedRemainingSeconds,omitempty"`
	ExtraInfo                 *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	FinishTime                *string                                                  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Interactive               *bool                                                    `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	StartTime                 *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepDescription           *string                                                  `json:"StepDescription,omitempty" xml:"StepDescription,omitempty"`
	StepInfo                  *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo  `json:"StepInfo,omitempty" xml:"StepInfo,omitempty" type:"Struct"`
	StepName                  *string                                                  `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepOrder                 *int32                                                   `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	StepProgress              *int32                                                   `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	StepStatus                *string                                                  `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetEstimatedRemainingSeconds(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.EstimatedRemainingSeconds = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetExtraInfo(v *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.ExtraInfo = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetFinishTime(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.FinishTime = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetInteractive(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.Interactive = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStartTime(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StartTime = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepDescription(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepDescription = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepInfo(v *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepInfo = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepOrder(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepOrder = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepProgress(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepProgress = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataSteps) SetStepStatus(v string) *DescribeOmsOpenAPIProjectResponseBodyDataSteps {
	s.StepStatus = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo struct {
	ErrorCode    *string                                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorDetails []*DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	ErrorMsg     *string                                                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ErrorParam   map[string]*string                                                     `json:"ErrorParam,omitempty" xml:"ErrorParam,omitempty"`
	FailedTime   *string                                                                `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) SetErrorCode(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo {
	s.ErrorCode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) SetErrorDetails(v []*DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo {
	s.ErrorDetails = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) SetErrorMsg(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) SetErrorParam(v map[string]*string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo {
	s.ErrorParam = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo) SetFailedTime(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfo {
	s.FailedTime = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) SetCode(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) SetLevel(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails {
	s.Level = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) SetMessage(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails) SetProposal(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsExtraInfoErrorDetails {
	s.Proposal = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo struct {
	Capacity                      *int64                                                                               `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Checkpoint                    *string                                                                              `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	ConnectorFullProgressOverview *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview `json:"ConnectorFullProgressOverview,omitempty" xml:"ConnectorFullProgressOverview,omitempty" type:"Struct"`
	DeployId                      *string                                                                              `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	DstIops                       *int64                                                                               `json:"DstIops,omitempty" xml:"DstIops,omitempty"`
	DstRps                        *int64                                                                               `json:"DstRps,omitempty" xml:"DstRps,omitempty"`
	DstRpsRef                     *int64                                                                               `json:"DstRpsRef,omitempty" xml:"DstRpsRef,omitempty"`
	DstRt                         *int64                                                                               `json:"DstRt,omitempty" xml:"DstRt,omitempty"`
	DstRtRef                      *int64                                                                               `json:"DstRtRef,omitempty" xml:"DstRtRef,omitempty"`
	Gmt                           *int64                                                                               `json:"Gmt,omitempty" xml:"Gmt,omitempty"`
	Inconsistencies               *int64                                                                               `json:"Inconsistencies,omitempty" xml:"Inconsistencies,omitempty"`
	IncrTimestampCheckpoint       *int64                                                                               `json:"IncrTimestampCheckpoint,omitempty" xml:"IncrTimestampCheckpoint,omitempty"`
	JobId                         *string                                                                              `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProcessedRecords              *int64                                                                               `json:"ProcessedRecords,omitempty" xml:"ProcessedRecords,omitempty"`
	Skipped                       *bool                                                                                `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	SrcIops                       *int64                                                                               `json:"SrcIops,omitempty" xml:"SrcIops,omitempty"`
	SrcIopsRef                    *int64                                                                               `json:"SrcIopsRef,omitempty" xml:"SrcIopsRef,omitempty"`
	SrcRps                        *int64                                                                               `json:"SrcRps,omitempty" xml:"SrcRps,omitempty"`
	SrcRpsRef                     *int64                                                                               `json:"SrcRpsRef,omitempty" xml:"SrcRpsRef,omitempty"`
	SrcRt                         *int64                                                                               `json:"SrcRt,omitempty" xml:"SrcRt,omitempty"`
	SrcRtRef                      *int64                                                                               `json:"SrcRtRef,omitempty" xml:"SrcRtRef,omitempty"`
	Validated                     *bool                                                                                `json:"Validated,omitempty" xml:"Validated,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetCapacity(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetCheckpoint(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Checkpoint = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetConnectorFullProgressOverview(v *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.ConnectorFullProgressOverview = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDeployId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DeployId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDstIops(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DstIops = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDstRps(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DstRps = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDstRpsRef(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DstRpsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDstRt(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DstRt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetDstRtRef(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.DstRtRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetGmt(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Gmt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetInconsistencies(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Inconsistencies = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetIncrTimestampCheckpoint(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.IncrTimestampCheckpoint = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetJobId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.JobId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetProcessedRecords(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.ProcessedRecords = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSkipped(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Skipped = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcIops(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcIops = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcIopsRef(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcIopsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcRps(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcRps = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcRpsRef(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcRpsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcRt(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcRt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetSrcRtRef(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.SrcRtRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo) SetValidated(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfo {
	s.Validated = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview struct {
	EstimatedRemainingTimeOfSec *int64 `json:"EstimatedRemainingTimeOfSec,omitempty" xml:"EstimatedRemainingTimeOfSec,omitempty"`
	EstimatedTotalCount         *int64 `json:"EstimatedTotalCount,omitempty" xml:"EstimatedTotalCount,omitempty"`
	FinishedCount               *int64 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	Progress                    *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetEstimatedRemainingTimeOfSec(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.EstimatedRemainingTimeOfSec = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetEstimatedTotalCount(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.EstimatedTotalCount = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetFinishedCount(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.FinishedCount = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetProgress(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.Progress = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping struct {
	Databases []*DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Mode      *string                                                              `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping) SetDatabases(v []*DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping {
	s.Databases = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping) SetMode(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMapping {
	s.Mode = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases struct {
	DatabaseId   *string                                                                    `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	DatabaseName *string                                                                    `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	MappedName   *string                                                                    `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	Tables       []*DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TenantName   *string                                                                    `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	Type         *string                                                                    `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetDatabaseId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.DatabaseId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetDatabaseName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.DatabaseName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetMappedName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.MappedName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetTables(v []*DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.Tables = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetTenantName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.TenantName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases) SetType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabases {
	s.Type = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables struct {
	AdbTableSchema *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema `json:"AdbTableSchema,omitempty" xml:"AdbTableSchema,omitempty" type:"Struct"`
	FilterColumns  []*string                                                                              `json:"FilterColumns,omitempty" xml:"FilterColumns,omitempty" type:"Repeated"`
	MappedName     *string                                                                                `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	ShardColumns   []*string                                                                              `json:"ShardColumns,omitempty" xml:"ShardColumns,omitempty" type:"Repeated"`
	TableId        *string                                                                                `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName      *string                                                                                `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type           *string                                                                                `json:"Type,omitempty" xml:"Type,omitempty"`
	WhereClause    *string                                                                                `json:"WhereClause,omitempty" xml:"WhereClause,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetAdbTableSchema(v *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.AdbTableSchema = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetFilterColumns(v []*string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.FilterColumns = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetMappedName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.MappedName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetShardColumns(v []*string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.ShardColumns = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetTableId(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.TableId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetTableName(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.TableName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.Type = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables) SetWhereClause(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTables {
	s.WhereClause = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema struct {
	DistributedKeys    []*string `json:"DistributedKeys,omitempty" xml:"DistributedKeys,omitempty" type:"Repeated"`
	PartitionLifeCycle *int32    `json:"PartitionLifeCycle,omitempty" xml:"PartitionLifeCycle,omitempty"`
	PartitionStatement *string   `json:"PartitionStatement,omitempty" xml:"PartitionStatement,omitempty"`
	PrimaryKeys        []*string `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty" type:"Repeated"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetDistributedKeys(v []*string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.DistributedKeys = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPartitionLifeCycle(v int32) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionLifeCycle = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPartitionStatement(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionStatement = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPrimaryKeys(v []*string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PrimaryKeys = v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig struct {
	EnableFullSync             *bool                                                                                  `json:"EnableFullSync,omitempty" xml:"EnableFullSync,omitempty"`
	EnableIncrSync             *bool                                                                                  `json:"EnableIncrSync,omitempty" xml:"EnableIncrSync,omitempty"`
	EnableStructSync           *bool                                                                                  `json:"EnableStructSync,omitempty" xml:"EnableStructSync,omitempty"`
	IncrSyncStepTransferConfig *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig `json:"IncrSyncStepTransferConfig,omitempty" xml:"IncrSyncStepTransferConfig,omitempty" type:"Struct"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) SetEnableFullSync(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig {
	s.EnableFullSync = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) SetEnableIncrSync(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig {
	s.EnableIncrSync = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) SetEnableStructSync(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig {
	s.EnableStructSync = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig) SetIncrSyncStepTransferConfig(v *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfig {
	s.IncrSyncStepTransferConfig = v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig struct {
	RecordTypeList          []*string `json:"RecordTypeList,omitempty" xml:"RecordTypeList,omitempty" type:"Repeated"`
	StartTimestamp          *int64    `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StoreLogKeptHour        *int64    `json:"StoreLogKeptHour,omitempty" xml:"StoreLogKeptHour,omitempty"`
	StoreTransactionEnabled *bool     `json:"StoreTransactionEnabled,omitempty" xml:"StoreTransactionEnabled,omitempty"`
	TransferStepType        *string   `json:"TransferStepType,omitempty" xml:"TransferStepType,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetRecordTypeList(v []*string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.RecordTypeList = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStartTimestamp(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStoreLogKeptHour(v int64) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreLogKeptHour = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStoreTransactionEnabled(v bool) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreTransactionEnabled = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetTransferStepType(v string) *DescribeOmsOpenAPIProjectResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.TransferStepType = &v
	return s
}

type DescribeOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s DescribeOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *DescribeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *DescribeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *DescribeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *DescribeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type DescribeOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *DescribeOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponse) SetStatusCode(v int32) *DescribeOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectResponse) SetBody(v *DescribeOmsOpenAPIProjectResponseBody) *DescribeOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type DescribeOmsOpenAPIProjectStepsRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsRequest) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsRequest) SetPageNumber(v int32) *DescribeOmsOpenAPIProjectStepsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsRequest) SetPageSize(v int32) *DescribeOmsOpenAPIProjectStepsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsRequest) SetProjectId(v string) *DescribeOmsOpenAPIProjectStepsRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsRequest) SetWorkerGradeId(v string) *DescribeOmsOpenAPIProjectStepsRequest {
	s.WorkerGradeId = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBody struct {
	Advice      *string                                                `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                                `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        []*DescribeOmsOpenAPIProjectStepsResponseBodyData      `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorDetail *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetAdvice(v string) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Advice = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetCode(v string) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetCost(v string) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Cost = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetData(v []*DescribeOmsOpenAPIProjectStepsResponseBodyData) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetErrorDetail(v *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetMessage(v string) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetPageNumber(v int32) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetPageSize(v int32) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetRequestId(v string) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetSuccess(v bool) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBody) SetTotalCount(v int64) *DescribeOmsOpenAPIProjectStepsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyData struct {
	EstimatedRemainingSeconds *int64                                                   `json:"EstimatedRemainingSeconds,omitempty" xml:"EstimatedRemainingSeconds,omitempty"`
	ExtraInfo                 *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	FinishTime                *string                                                  `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Interactive               *bool                                                    `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	StartTime                 *string                                                  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepDescription           *string                                                  `json:"StepDescription,omitempty" xml:"StepDescription,omitempty"`
	StepInfo                  *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo  `json:"StepInfo,omitempty" xml:"StepInfo,omitempty" type:"Struct"`
	StepName                  *string                                                  `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepOrder                 *int32                                                   `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	StepProgress              *int32                                                   `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	StepStatus                *string                                                  `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetEstimatedRemainingSeconds(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.EstimatedRemainingSeconds = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetExtraInfo(v *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.ExtraInfo = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetFinishTime(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetInteractive(v bool) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.Interactive = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStartTime(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepDescription(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepDescription = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepInfo(v *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepInfo = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepName(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepName = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepOrder(v int32) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepOrder = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepProgress(v int32) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepProgress = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyData) SetStepStatus(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyData {
	s.StepStatus = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo struct {
	ErrorCode    *string                                                                `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorDetails []*DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	ErrorMsg     *string                                                                `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ErrorParam   map[string]*string                                                     `json:"ErrorParam,omitempty" xml:"ErrorParam,omitempty"`
	FailedTime   *string                                                                `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) SetErrorCode(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo {
	s.ErrorCode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) SetErrorDetails(v []*DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo {
	s.ErrorDetails = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) SetErrorMsg(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo {
	s.ErrorMsg = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) SetErrorParam(v map[string]*string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo {
	s.ErrorParam = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo) SetFailedTime(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfo {
	s.FailedTime = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) SetCode(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) SetLevel(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails {
	s.Level = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) SetMessage(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails) SetProposal(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataExtraInfoErrorDetails {
	s.Proposal = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo struct {
	Capacity                      *int64                                                                               `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Checkpoint                    *string                                                                              `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	ConnectorFullProgressOverview *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview `json:"ConnectorFullProgressOverview,omitempty" xml:"ConnectorFullProgressOverview,omitempty" type:"Struct"`
	DeployId                      *string                                                                              `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	DstIops                       *int64                                                                               `json:"DstIops,omitempty" xml:"DstIops,omitempty"`
	DstRps                        *int64                                                                               `json:"DstRps,omitempty" xml:"DstRps,omitempty"`
	DstRpsRef                     *int64                                                                               `json:"DstRpsRef,omitempty" xml:"DstRpsRef,omitempty"`
	DstRt                         *int64                                                                               `json:"DstRt,omitempty" xml:"DstRt,omitempty"`
	DstRtRef                      *int64                                                                               `json:"DstRtRef,omitempty" xml:"DstRtRef,omitempty"`
	Gmt                           *int64                                                                               `json:"Gmt,omitempty" xml:"Gmt,omitempty"`
	Inconsistencies               *int64                                                                               `json:"Inconsistencies,omitempty" xml:"Inconsistencies,omitempty"`
	IncrTimestampCheckpoint       *int64                                                                               `json:"IncrTimestampCheckpoint,omitempty" xml:"IncrTimestampCheckpoint,omitempty"`
	JobId                         *string                                                                              `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProcessedRecords              *int64                                                                               `json:"ProcessedRecords,omitempty" xml:"ProcessedRecords,omitempty"`
	Skipped                       *bool                                                                                `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	SrcIops                       *int64                                                                               `json:"SrcIops,omitempty" xml:"SrcIops,omitempty"`
	SrcIopsRef                    *int64                                                                               `json:"SrcIopsRef,omitempty" xml:"SrcIopsRef,omitempty"`
	SrcRps                        *int64                                                                               `json:"SrcRps,omitempty" xml:"SrcRps,omitempty"`
	SrcRpsRef                     *int64                                                                               `json:"SrcRpsRef,omitempty" xml:"SrcRpsRef,omitempty"`
	SrcRt                         *int64                                                                               `json:"SrcRt,omitempty" xml:"SrcRt,omitempty"`
	SrcRtRef                      *int64                                                                               `json:"SrcRtRef,omitempty" xml:"SrcRtRef,omitempty"`
	Validated                     *bool                                                                                `json:"Validated,omitempty" xml:"Validated,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetCapacity(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Capacity = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetCheckpoint(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Checkpoint = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetConnectorFullProgressOverview(v *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.ConnectorFullProgressOverview = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDeployId(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DeployId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDstIops(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DstIops = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDstRps(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DstRps = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDstRpsRef(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DstRpsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDstRt(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DstRt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetDstRtRef(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.DstRtRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetGmt(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Gmt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetInconsistencies(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Inconsistencies = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetIncrTimestampCheckpoint(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.IncrTimestampCheckpoint = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetJobId(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.JobId = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetProcessedRecords(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.ProcessedRecords = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSkipped(v bool) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Skipped = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcIops(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcIops = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcIopsRef(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcIopsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcRps(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcRps = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcRpsRef(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcRpsRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcRt(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcRt = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetSrcRtRef(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.SrcRtRef = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo) SetValidated(v bool) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfo {
	s.Validated = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview struct {
	EstimatedRemainingTimeOfSec *int64 `json:"EstimatedRemainingTimeOfSec,omitempty" xml:"EstimatedRemainingTimeOfSec,omitempty"`
	EstimatedTotalCount         *int64 `json:"EstimatedTotalCount,omitempty" xml:"EstimatedTotalCount,omitempty"`
	FinishedCount               *int64 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	Progress                    *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) SetEstimatedRemainingTimeOfSec(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview {
	s.EstimatedRemainingTimeOfSec = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) SetEstimatedTotalCount(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview {
	s.EstimatedTotalCount = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) SetFinishedCount(v int64) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview {
	s.FinishedCount = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview) SetProgress(v int32) *DescribeOmsOpenAPIProjectStepsResponseBodyDataStepInfoConnectorFullProgressOverview {
	s.Progress = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) SetCode(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) SetLevel(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) SetMessage(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail) SetProposal(v string) *DescribeOmsOpenAPIProjectStepsResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type DescribeOmsOpenAPIProjectStepsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOmsOpenAPIProjectStepsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOmsOpenAPIProjectStepsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOmsOpenAPIProjectStepsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOmsOpenAPIProjectStepsResponse) SetHeaders(v map[string]*string) *DescribeOmsOpenAPIProjectStepsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponse) SetStatusCode(v int32) *DescribeOmsOpenAPIProjectStepsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOmsOpenAPIProjectStepsResponse) SetBody(v *DescribeOmsOpenAPIProjectStepsResponseBody) *DescribeOmsOpenAPIProjectStepsResponse {
	s.Body = v
	return s
}

type DescribeOutlineBindingRequest struct {
	DatabaseName      *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	InstanceId        *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IsConcurrentLimit *bool   `json:"IsConcurrentLimit,omitempty" xml:"IsConcurrentLimit,omitempty"`
	SQLId             *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	TableName         *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TenantId          *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeOutlineBindingRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingRequest) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingRequest) SetDatabaseName(v string) *DescribeOutlineBindingRequest {
	s.DatabaseName = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetInstanceId(v string) *DescribeOutlineBindingRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetIsConcurrentLimit(v bool) *DescribeOutlineBindingRequest {
	s.IsConcurrentLimit = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetSQLId(v string) *DescribeOutlineBindingRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetTableName(v string) *DescribeOutlineBindingRequest {
	s.TableName = &v
	return s
}

func (s *DescribeOutlineBindingRequest) SetTenantId(v string) *DescribeOutlineBindingRequest {
	s.TenantId = &v
	return s
}

type DescribeOutlineBindingResponseBody struct {
	OutlineBinding *DescribeOutlineBindingResponseBodyOutlineBinding `json:"OutlineBinding,omitempty" xml:"OutlineBinding,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOutlineBindingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponseBody) SetOutlineBinding(v *DescribeOutlineBindingResponseBodyOutlineBinding) *DescribeOutlineBindingResponseBody {
	s.OutlineBinding = v
	return s
}

func (s *DescribeOutlineBindingResponseBody) SetRequestId(v string) *DescribeOutlineBindingResponseBody {
	s.RequestId = &v
	return s
}

type DescribeOutlineBindingResponseBodyOutlineBinding struct {
	BindIndex     *string `json:"BindIndex,omitempty" xml:"BindIndex,omitempty"`
	BindPlan      *string `json:"BindPlan,omitempty" xml:"BindPlan,omitempty"`
	MaxConcurrent *int32  `json:"MaxConcurrent,omitempty" xml:"MaxConcurrent,omitempty"`
	OutlineId     *int64  `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
}

func (s DescribeOutlineBindingResponseBodyOutlineBinding) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponseBodyOutlineBinding) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetBindIndex(v string) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.BindIndex = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetBindPlan(v string) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.BindPlan = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetMaxConcurrent(v int32) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.MaxConcurrent = &v
	return s
}

func (s *DescribeOutlineBindingResponseBodyOutlineBinding) SetOutlineId(v int64) *DescribeOutlineBindingResponseBodyOutlineBinding {
	s.OutlineId = &v
	return s
}

type DescribeOutlineBindingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeOutlineBindingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeOutlineBindingResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeOutlineBindingResponse) GoString() string {
	return s.String()
}

func (s *DescribeOutlineBindingResponse) SetHeaders(v map[string]*string) *DescribeOutlineBindingResponse {
	s.Headers = v
	return s
}

func (s *DescribeOutlineBindingResponse) SetStatusCode(v int32) *DescribeOutlineBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOutlineBindingResponse) SetBody(v *DescribeOutlineBindingResponseBody) *DescribeOutlineBindingResponse {
	s.Body = v
	return s
}

type DescribeParametersRequest struct {
	Dimension      *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersRequest) SetDimension(v string) *DescribeParametersRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeParametersRequest) SetDimensionValue(v string) *DescribeParametersRequest {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersRequest) SetInstanceId(v string) *DescribeParametersRequest {
	s.InstanceId = &v
	return s
}

type DescribeParametersResponseBody struct {
	Parameters []*DescribeParametersResponseBodyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBody) SetParameters(v []*DescribeParametersResponseBodyParameters) *DescribeParametersResponseBody {
	s.Parameters = v
	return s
}

func (s *DescribeParametersResponseBody) SetRequestId(v string) *DescribeParametersResponseBody {
	s.RequestId = &v
	return s
}

type DescribeParametersResponseBodyParameters struct {
	AcceptableValue []*string `json:"AcceptableValue,omitempty" xml:"AcceptableValue,omitempty" type:"Repeated"`
	CurrentValue    *string   `json:"CurrentValue,omitempty" xml:"CurrentValue,omitempty"`
	DefaultValue    *string   `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	Description     *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	Name            *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	NeedReboot      *bool     `json:"NeedReboot,omitempty" xml:"NeedReboot,omitempty"`
	RejectedValue   []*string `json:"RejectedValue,omitempty" xml:"RejectedValue,omitempty" type:"Repeated"`
	ValueType       *string   `json:"ValueType,omitempty" xml:"ValueType,omitempty"`
}

func (s DescribeParametersResponseBodyParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponseBodyParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponseBodyParameters) SetAcceptableValue(v []*string) *DescribeParametersResponseBodyParameters {
	s.AcceptableValue = v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetCurrentValue(v string) *DescribeParametersResponseBodyParameters {
	s.CurrentValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetDefaultValue(v string) *DescribeParametersResponseBodyParameters {
	s.DefaultValue = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetDescription(v string) *DescribeParametersResponseBodyParameters {
	s.Description = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetName(v string) *DescribeParametersResponseBodyParameters {
	s.Name = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetNeedReboot(v bool) *DescribeParametersResponseBodyParameters {
	s.NeedReboot = &v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetRejectedValue(v []*string) *DescribeParametersResponseBodyParameters {
	s.RejectedValue = v
	return s
}

func (s *DescribeParametersResponseBodyParameters) SetValueType(v string) *DescribeParametersResponseBodyParameters {
	s.ValueType = &v
	return s
}

type DescribeParametersResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersResponse) SetHeaders(v map[string]*string) *DescribeParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersResponse) SetStatusCode(v int32) *DescribeParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersResponse) SetBody(v *DescribeParametersResponseBody) *DescribeParametersResponse {
	s.Body = v
	return s
}

type DescribeParametersHistoryRequest struct {
	Dimension      *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime      *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeParametersHistoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryRequest) SetDimension(v string) *DescribeParametersHistoryRequest {
	s.Dimension = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetDimensionValue(v string) *DescribeParametersHistoryRequest {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetEndTime(v string) *DescribeParametersHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetInstanceId(v string) *DescribeParametersHistoryRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetPageNumber(v int32) *DescribeParametersHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetPageSize(v int32) *DescribeParametersHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeParametersHistoryRequest) SetStartTime(v string) *DescribeParametersHistoryRequest {
	s.StartTime = &v
	return s
}

type DescribeParametersHistoryResponseBody struct {
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Respond   []*DescribeParametersHistoryResponseBodyRespond `json:"Respond,omitempty" xml:"Respond,omitempty" type:"Repeated"`
}

func (s DescribeParametersHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBody) SetRequestId(v string) *DescribeParametersHistoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeParametersHistoryResponseBody) SetRespond(v []*DescribeParametersHistoryResponseBodyRespond) *DescribeParametersHistoryResponseBody {
	s.Respond = v
	return s
}

type DescribeParametersHistoryResponseBodyRespond struct {
	PageNumber *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	Parameters []*DescribeParametersHistoryResponseBodyRespondParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeParametersHistoryResponseBodyRespond) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBodyRespond) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetPageNumber(v int32) *DescribeParametersHistoryResponseBodyRespond {
	s.PageNumber = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetParameters(v []*DescribeParametersHistoryResponseBodyRespondParameters) *DescribeParametersHistoryResponseBodyRespond {
	s.Parameters = v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespond) SetTotalCount(v int32) *DescribeParametersHistoryResponseBodyRespond {
	s.TotalCount = &v
	return s
}

type DescribeParametersHistoryResponseBodyRespondParameters struct {
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	NewValue       *string `json:"NewValue,omitempty" xml:"NewValue,omitempty"`
	OldValue       *string `json:"OldValue,omitempty" xml:"OldValue,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeParametersHistoryResponseBodyRespondParameters) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponseBodyRespondParameters) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetCreateTime(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.CreateTime = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetDimensionValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.DimensionValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetName(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.Name = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetNewValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.NewValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetOldValue(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.OldValue = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetStatus(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.Status = &v
	return s
}

func (s *DescribeParametersHistoryResponseBodyRespondParameters) SetUpdateTime(v string) *DescribeParametersHistoryResponseBodyRespondParameters {
	s.UpdateTime = &v
	return s
}

type DescribeParametersHistoryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeParametersHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeParametersHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeParametersHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeParametersHistoryResponse) SetHeaders(v map[string]*string) *DescribeParametersHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeParametersHistoryResponse) SetStatusCode(v int32) *DescribeParametersHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParametersHistoryResponse) SetBody(v *DescribeParametersHistoryResponseBody) *DescribeParametersHistoryResponse {
	s.Body = v
	return s
}

type DescribeRecommendIndexRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SQLId      *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeRecommendIndexRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexRequest) SetInstanceId(v string) *DescribeRecommendIndexRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeRecommendIndexRequest) SetSQLId(v string) *DescribeRecommendIndexRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeRecommendIndexRequest) SetTenantId(v string) *DescribeRecommendIndexRequest {
	s.TenantId = &v
	return s
}

type DescribeRecommendIndexResponseBody struct {
	RecommendIndex *DescribeRecommendIndexResponseBodyRecommendIndex `json:"RecommendIndex,omitempty" xml:"RecommendIndex,omitempty" type:"Struct"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRecommendIndexResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponseBody) SetRecommendIndex(v *DescribeRecommendIndexResponseBodyRecommendIndex) *DescribeRecommendIndexResponseBody {
	s.RecommendIndex = v
	return s
}

func (s *DescribeRecommendIndexResponseBody) SetRequestId(v string) *DescribeRecommendIndexResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRecommendIndexResponseBodyRecommendIndex struct {
	SuggestIndex *string `json:"SuggestIndex,omitempty" xml:"SuggestIndex,omitempty"`
	TableList    *string `json:"TableList,omitempty" xml:"TableList,omitempty"`
	TenantMode   *string `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
}

func (s DescribeRecommendIndexResponseBodyRecommendIndex) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponseBodyRecommendIndex) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetSuggestIndex(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.SuggestIndex = &v
	return s
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetTableList(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.TableList = &v
	return s
}

func (s *DescribeRecommendIndexResponseBodyRecommendIndex) SetTenantMode(v string) *DescribeRecommendIndexResponseBodyRecommendIndex {
	s.TenantMode = &v
	return s
}

type DescribeRecommendIndexResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRecommendIndexResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRecommendIndexResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRecommendIndexResponse) GoString() string {
	return s.String()
}

func (s *DescribeRecommendIndexResponse) SetHeaders(v map[string]*string) *DescribeRecommendIndexResponse {
	s.Headers = v
	return s
}

func (s *DescribeRecommendIndexResponse) SetStatusCode(v int32) *DescribeRecommendIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRecommendIndexResponse) SetBody(v *DescribeRecommendIndexResponseBody) *DescribeRecommendIndexResponse {
	s.Body = v
	return s
}

type DescribeSQLDetailsRequest struct {
	SQLId    *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsRequest) SetSQLId(v string) *DescribeSQLDetailsRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLDetailsRequest) SetTenantId(v string) *DescribeSQLDetailsRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLDetailsResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLDetails []*DescribeSQLDetailsResponseBodySQLDetails `json:"SQLDetails,omitempty" xml:"SQLDetails,omitempty" type:"Repeated"`
}

func (s DescribeSQLDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponseBody) SetRequestId(v string) *DescribeSQLDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLDetailsResponseBody) SetSQLDetails(v []*DescribeSQLDetailsResponseBodySQLDetails) *DescribeSQLDetailsResponseBody {
	s.SQLDetails = v
	return s
}

type DescribeSQLDetailsResponseBodySQLDetails struct {
	DbName   *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	SQLText  *string `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSQLDetailsResponseBodySQLDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponseBodySQLDetails) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetDbName(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.DbName = &v
	return s
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetSQLText(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.SQLText = &v
	return s
}

func (s *DescribeSQLDetailsResponseBodySQLDetails) SetUserName(v string) *DescribeSQLDetailsResponseBodySQLDetails {
	s.UserName = &v
	return s
}

type DescribeSQLDetailsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLDetailsResponse) SetHeaders(v map[string]*string) *DescribeSQLDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLDetailsResponse) SetStatusCode(v int32) *DescribeSQLDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLDetailsResponse) SetBody(v *DescribeSQLDetailsResponseBody) *DescribeSQLDetailsResponse {
	s.Body = v
	return s
}

type DescribeSQLHistoryListRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId      *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListRequest) SetEndTime(v string) *DescribeSQLHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetPageNumber(v int32) *DescribeSQLHistoryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetPageSize(v int32) *DescribeSQLHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetSQLId(v string) *DescribeSQLHistoryListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetStartTime(v string) *DescribeSQLHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSQLHistoryListRequest) SetTenantId(v string) *DescribeSQLHistoryListRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLHistoryListResponseBody struct {
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLHistoryList *DescribeSQLHistoryListResponseBodySQLHistoryList `json:"SQLHistoryList,omitempty" xml:"SQLHistoryList,omitempty" type:"Struct"`
}

func (s DescribeSQLHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBody) SetRequestId(v string) *DescribeSQLHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBody) SetSQLHistoryList(v *DescribeSQLHistoryListResponseBodySQLHistoryList) *DescribeSQLHistoryListResponseBody {
	s.SQLHistoryList = v
	return s
}

type DescribeSQLHistoryListResponseBodySQLHistoryList struct {
	Count *int64                                                  `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*DescribeSQLHistoryListResponseBodySQLHistoryListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryList) SetCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryList {
	s.Count = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryList) SetList(v []*DescribeSQLHistoryListResponseBodySQLHistoryListList) *DescribeSQLHistoryListResponseBodySQLHistoryList {
	s.List = v
	return s
}

type DescribeSQLHistoryListResponseBodySQLHistoryListList struct {
	AffectedRows         *int64   `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *int64   `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *int64   `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *int64   `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *int64   `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	EndTime              *int64   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	EndTimeUTCString     *string  `json:"EndTimeUTCString,omitempty" xml:"EndTimeUTCString,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *int64   `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *int64   `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *int64   `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetPlanTime          *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	IOWaitTime           *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	LogicalRead          *int64   `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *int64   `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *int64   `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetWaitTime          *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *int64   `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *int64   `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *int64   `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *int64   `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *int64   `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	ScheduleTime         *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SsstoreReadRowCount  *int64   `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TotalWaitTime        *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponseBodySQLHistoryListList) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetAffectedRows(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetAppWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBlockCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBlockIndexCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetBloomFilterCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetClientIp(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetConcurrencyWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetCpuTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDbName(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DbName = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDecodeTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetDiskRead(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetElapsedTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEndTime(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.EndTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEndTimeUTCString(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.EndTimeUTCString = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetEvent(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.Event = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecPerSecond(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecuteTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetExecutions(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.Executions = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetFailTimes(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetGetPlanTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetIOWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetLogicalRead(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMaxCpuTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMaxElapsedTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMemstoreReadRowCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetMissPlans(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetNetWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetNodeIp(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetQueueTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRPCCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRemotePlans(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRetryCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetReturnRows(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetRowCacheHit(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetScheduleTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetSsstoreReadRowCount(v int64) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetTotalWaitTime(v float32) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSQLHistoryListResponseBodySQLHistoryListList) SetUserName(v string) *DescribeSQLHistoryListResponseBodySQLHistoryListList {
	s.UserName = &v
	return s
}

type DescribeSQLHistoryListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLHistoryListResponse) SetHeaders(v map[string]*string) *DescribeSQLHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLHistoryListResponse) SetStatusCode(v int32) *DescribeSQLHistoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLHistoryListResponse) SetBody(v *DescribeSQLHistoryListResponseBody) *DescribeSQLHistoryListResponse {
	s.Body = v
	return s
}

type DescribeSQLPlansRequest struct {
	SQLId    *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSQLPlansRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansRequest) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansRequest) SetSQLId(v string) *DescribeSQLPlansRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSQLPlansRequest) SetTenantId(v string) *DescribeSQLPlansRequest {
	s.TenantId = &v
	return s
}

type DescribeSQLPlansResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SQLPlans  []*DescribeSQLPlansResponseBodySQLPlans `json:"SQLPlans,omitempty" xml:"SQLPlans,omitempty" type:"Repeated"`
}

func (s DescribeSQLPlansResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponseBody) SetRequestId(v string) *DescribeSQLPlansResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLPlansResponseBody) SetSQLPlans(v []*DescribeSQLPlansResponseBodySQLPlans) *DescribeSQLPlansResponseBody {
	s.SQLPlans = v
	return s
}

type DescribeSQLPlansResponseBodySQLPlans struct {
	AvgExecutionMS         *float32 `json:"AvgExecutionMS,omitempty" xml:"AvgExecutionMS,omitempty"`
	AvgExecutionTimeMS     *int64   `json:"AvgExecutionTimeMS,omitempty" xml:"AvgExecutionTimeMS,omitempty"`
	FirstLoadTime          *int64   `json:"FirstLoadTime,omitempty" xml:"FirstLoadTime,omitempty"`
	FirstLoadTimeUTCString *string  `json:"FirstLoadTimeUTCString,omitempty" xml:"FirstLoadTimeUTCString,omitempty"`
	HitCount               *int32   `json:"HitCount,omitempty" xml:"HitCount,omitempty"`
	MergedVersion          *int32   `json:"MergedVersion,omitempty" xml:"MergedVersion,omitempty"`
	NodeIp                 *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	OutlineData            *string  `json:"OutlineData,omitempty" xml:"OutlineData,omitempty"`
	OutlineId              *int64   `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	OutlineTime            *int64   `json:"OutlineTime,omitempty" xml:"OutlineTime,omitempty"`
	OutlineTimeUTCString   *string  `json:"OutlineTimeUTCString,omitempty" xml:"OutlineTimeUTCString,omitempty"`
	PlanFull               *string  `json:"PlanFull,omitempty" xml:"PlanFull,omitempty"`
	PlanId                 *int32   `json:"PlanId,omitempty" xml:"PlanId,omitempty"`
	PlanInfo               *string  `json:"PlanInfo,omitempty" xml:"PlanInfo,omitempty"`
	PlanUnionHash          *string  `json:"PlanUnionHash,omitempty" xml:"PlanUnionHash,omitempty"`
	QuerySQL               *string  `json:"QuerySQL,omitempty" xml:"QuerySQL,omitempty"`
}

func (s DescribeSQLPlansResponseBodySQLPlans) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponseBodySQLPlans) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetAvgExecutionMS(v float32) *DescribeSQLPlansResponseBodySQLPlans {
	s.AvgExecutionMS = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetAvgExecutionTimeMS(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.AvgExecutionTimeMS = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetFirstLoadTime(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.FirstLoadTime = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetFirstLoadTimeUTCString(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.FirstLoadTimeUTCString = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetHitCount(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.HitCount = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetMergedVersion(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.MergedVersion = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetNodeIp(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.NodeIp = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineData(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineData = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineId(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineId = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineTime(v int64) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineTime = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetOutlineTimeUTCString(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.OutlineTimeUTCString = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanFull(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanFull = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanId(v int32) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanId = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanInfo(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanInfo = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetPlanUnionHash(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.PlanUnionHash = &v
	return s
}

func (s *DescribeSQLPlansResponseBodySQLPlans) SetQuerySQL(v string) *DescribeSQLPlansResponseBodySQLPlans {
	s.QuerySQL = &v
	return s
}

type DescribeSQLPlansResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSQLPlansResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSQLPlansResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSQLPlansResponse) GoString() string {
	return s.String()
}

func (s *DescribeSQLPlansResponse) SetHeaders(v map[string]*string) *DescribeSQLPlansResponse {
	s.Headers = v
	return s
}

func (s *DescribeSQLPlansResponse) SetStatusCode(v int32) *DescribeSQLPlansResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSQLPlansResponse) SetBody(v *DescribeSQLPlansResponseBody) *DescribeSQLPlansResponse {
	s.Body = v
	return s
}

type DescribeSecurityIpGroupsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeSecurityIpGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsRequest) SetInstanceId(v string) *DescribeSecurityIpGroupsRequest {
	s.InstanceId = &v
	return s
}

type DescribeSecurityIpGroupsResponseBody struct {
	RequestId        *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroups []*DescribeSecurityIpGroupsResponseBodySecurityIpGroups `json:"SecurityIpGroups,omitempty" xml:"SecurityIpGroups,omitempty" type:"Repeated"`
	TotalCount       *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityIpGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponseBody) SetRequestId(v string) *DescribeSecurityIpGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBody) SetSecurityIpGroups(v []*DescribeSecurityIpGroupsResponseBodySecurityIpGroups) *DescribeSecurityIpGroupsResponseBody {
	s.SecurityIpGroups = v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBody) SetTotalCount(v int32) *DescribeSecurityIpGroupsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSecurityIpGroupsResponseBodySecurityIpGroups struct {
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIps         *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s DescribeSecurityIpGroupsResponseBodySecurityIpGroups) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponseBodySecurityIpGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponseBodySecurityIpGroups) SetSecurityIpGroupName(v string) *DescribeSecurityIpGroupsResponseBodySecurityIpGroups {
	s.SecurityIpGroupName = &v
	return s
}

func (s *DescribeSecurityIpGroupsResponseBodySecurityIpGroups) SetSecurityIps(v string) *DescribeSecurityIpGroupsResponseBodySecurityIpGroups {
	s.SecurityIps = &v
	return s
}

type DescribeSecurityIpGroupsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSecurityIpGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSecurityIpGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSecurityIpGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpGroupsResponse) SetHeaders(v map[string]*string) *DescribeSecurityIpGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecurityIpGroupsResponse) SetStatusCode(v int32) *DescribeSecurityIpGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecurityIpGroupsResponse) SetBody(v *DescribeSecurityIpGroupsResponseBody) *DescribeSecurityIpGroupsResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLHistoryListRequest struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId      *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLHistoryListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListRequest) SetEndTime(v string) *DescribeSlowSQLHistoryListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetPageNumber(v int32) *DescribeSlowSQLHistoryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetPageSize(v int32) *DescribeSlowSQLHistoryListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetSQLId(v string) *DescribeSlowSQLHistoryListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetStartTime(v string) *DescribeSlowSQLHistoryListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListRequest) SetTenantId(v string) *DescribeSlowSQLHistoryListRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLHistoryListResponseBody struct {
	RequestId          *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowSQLHistoryList *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList `json:"SlowSQLHistoryList,omitempty" xml:"SlowSQLHistoryList,omitempty" type:"Struct"`
}

func (s DescribeSlowSQLHistoryListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBody) SetRequestId(v string) *DescribeSlowSQLHistoryListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBody) SetSlowSQLHistoryList(v *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) *DescribeSlowSQLHistoryListResponseBody {
	s.SlowSQLHistoryList = v
	return s
}

type DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList struct {
	Count *int64                                                          `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) SetCount(v int64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList {
	s.Count = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList) SetList(v []*DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryList {
	s.List = v
	return s
}

type DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList struct {
	AffectedRows         *float64 `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float64 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *float64 `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *float64 `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *float64 `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float64 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float64 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float64 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *float64 `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float64 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	EndTimeUTCString     *string  `json:"EndTimeUTCString,omitempty" xml:"EndTimeUTCString,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *float64 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float64 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *float64 `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *float64 `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetplanTime          *float64 `json:"GetplanTime,omitempty" xml:"GetplanTime,omitempty"`
	IOWaitTime           *float64 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	LogicalRead          *float64 `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float64 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float64 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *float64 `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *float64 `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetwaitTime          *float64 `json:"NetwaitTime,omitempty" xml:"NetwaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float64 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *float64 `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *float64 `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *float64 `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *float64 `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *float64 `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	ScheduleTime         *float64 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SqlId                *string  `json:"SqlId,omitempty" xml:"SqlId,omitempty"`
	SqlType              *string  `json:"SqlType,omitempty" xml:"SqlType,omitempty"`
	SsstoreReadRowCount  *float64 `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TenantName           *string  `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	TotalWaitTime        *float64 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetAffectedRows(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetAppWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBlockCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBlockIndexCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetBloomFilterCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetClientIp(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetConcurrencyWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetCpuTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDbName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDecodeTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetDiskRead(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetElapsedTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetEndTimeUTCString(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.EndTimeUTCString = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetEvent(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.Event = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecPerSecond(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecuteTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetExecutions(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.Executions = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetFailTimes(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetGetplanTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.GetplanTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetIOWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetLogicalRead(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMaxCpuTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMaxElapsedTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMemstoreReadRowCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetMissPlans(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetNetwaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.NetwaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetNodeIp(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetQueueTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRPCCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRemotePlans(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRetryCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetReturnRows(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetRowCacheHit(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetScheduleTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSqlId(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SqlId = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSqlType(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SqlType = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetSsstoreReadRowCount(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetTenantName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.TenantName = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetTotalWaitTime(v float64) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList) SetUserName(v string) *DescribeSlowSQLHistoryListResponseBodySlowSQLHistoryListList {
	s.UserName = &v
	return s
}

type DescribeSlowSQLHistoryListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowSQLHistoryListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLHistoryListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLHistoryListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLHistoryListResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLHistoryListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLHistoryListResponse) SetStatusCode(v int32) *DescribeSlowSQLHistoryListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowSQLHistoryListResponse) SetBody(v *DescribeSlowSQLHistoryListResponseBody) *DescribeSlowSQLHistoryListResponse {
	s.Body = v
	return s
}

type DescribeSlowSQLListRequest struct {
	DbName          *string                `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime         *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp          *string                `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber      *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId           *string                `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord   *string                `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter *string                `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule      *string                `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue     *string                `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn      *string                `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder       *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime       *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId        *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListRequest) SetDbName(v string) *DescribeSlowSQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetEndTime(v string) *DescribeSlowSQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeSlowSQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeSlowSQLListRequest) SetNodeIp(v string) *DescribeSlowSQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetPageNumber(v int32) *DescribeSlowSQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetPageSize(v int32) *DescribeSlowSQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSQLId(v string) *DescribeSlowSQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchKeyWord(v string) *DescribeSlowSQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchParameter(v string) *DescribeSlowSQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchRule(v string) *DescribeSlowSQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSearchValue(v string) *DescribeSlowSQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSortColumn(v string) *DescribeSlowSQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetSortOrder(v string) *DescribeSlowSQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetStartTime(v string) *DescribeSlowSQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLListRequest) SetTenantId(v string) *DescribeSlowSQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLListShrinkRequest struct {
	DbName                *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp                *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId                 *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord         *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter       *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule            *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue           *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn            *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder             *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId              *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeSlowSQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListShrinkRequest) SetDbName(v string) *DescribeSlowSQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetEndTime(v string) *DescribeSlowSQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeSlowSQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetNodeIp(v string) *DescribeSlowSQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetPageNumber(v int32) *DescribeSlowSQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetPageSize(v int32) *DescribeSlowSQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSQLId(v string) *DescribeSlowSQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchParameter(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchRule(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSearchValue(v string) *DescribeSlowSQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSortColumn(v string) *DescribeSlowSQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetSortOrder(v string) *DescribeSlowSQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetStartTime(v string) *DescribeSlowSQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowSQLListShrinkRequest) SetTenantId(v string) *DescribeSlowSQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeSlowSQLListResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlowSQLList []*DescribeSlowSQLListResponseBodySlowSQLList `json:"SlowSQLList,omitempty" xml:"SlowSQLList,omitempty" type:"Repeated"`
	TotalCount  *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSlowSQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponseBody) SetRequestId(v string) *DescribeSlowSQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlowSQLListResponseBody) SetSlowSQLList(v []*DescribeSlowSQLListResponseBodySlowSQLList) *DescribeSlowSQLListResponseBody {
	s.SlowSQLList = v
	return s
}

func (s *DescribeSlowSQLListResponseBody) SetTotalCount(v int64) *DescribeSlowSQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSlowSQLListResponseBodySlowSQLList struct {
	AffectedRows         *int64   `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *int64   `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *int64   `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *int64   `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *int64   `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *float32 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *int64   `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *int64   `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetPlanTime          *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	IOWaitTime           *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	Key                  *int64   `json:"Key,omitempty" xml:"Key,omitempty"`
	LogicalRead          *int64   `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *int64   `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *int64   `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetWaitTime          *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *int64   `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *int64   `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *int64   `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *int64   `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *int64   `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	SQLId                *string  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SQLType              *int64   `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	ScheduleTime         *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SsstoreReadRowCount  *int64   `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TotalWaitTime        *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSlowSQLListResponseBodySlowSQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponseBodySlowSQLList) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetAffectedRows(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetAppWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBlockCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBlockIndexCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetBloomFilterCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetClientIp(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ClientIp = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetConcurrencyWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetCpuTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDbName(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DbName = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDecodeTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetDiskRead(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.DiskRead = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetElapsedTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetEvent(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Event = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecPerSecond(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecuteTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetExecutions(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Executions = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetFailTimes(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.FailTimes = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetGetPlanTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetIOWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetKey(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.Key = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetLogicalRead(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMaxCpuTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMaxElapsedTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMemstoreReadRowCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetMissPlans(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.MissPlans = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetNetWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetNodeIp(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.NodeIp = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetQueueTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.QueueTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRPCCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RPCCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRemotePlans(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRetryCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RetryCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetReturnRows(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetRowCacheHit(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLId(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLText(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSQLType(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SQLType = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetScheduleTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetSsstoreReadRowCount(v int64) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetTotalWaitTime(v float32) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeSlowSQLListResponseBodySlowSQLList) SetUserName(v string) *DescribeSlowSQLListResponseBodySlowSQLList {
	s.UserName = &v
	return s
}

type DescribeSlowSQLListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSlowSQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSlowSQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSlowSQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlowSQLListResponse) SetHeaders(v map[string]*string) *DescribeSlowSQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlowSQLListResponse) SetStatusCode(v int32) *DescribeSlowSQLListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlowSQLListResponse) SetBody(v *DescribeSlowSQLListResponseBody) *DescribeSlowSQLListResponse {
	s.Body = v
	return s
}

type DescribeTenantRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTenantRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantRequest) SetInstanceId(v string) *DescribeTenantRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantRequest) SetTenantId(v string) *DescribeTenantRequest {
	s.TenantId = &v
	return s
}

type DescribeTenantResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tenant    *DescribeTenantResponseBodyTenant `json:"Tenant,omitempty" xml:"Tenant,omitempty" type:"Struct"`
}

func (s DescribeTenantResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBody) SetRequestId(v string) *DescribeTenantResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantResponseBody) SetTenant(v *DescribeTenantResponseBodyTenant) *DescribeTenantResponseBody {
	s.Tenant = v
	return s
}

type DescribeTenantResponseBodyTenant struct {
	Charset                      *string                                              `json:"Charset,omitempty" xml:"Charset,omitempty"`
	ClogServiceStatus            *string                                              `json:"ClogServiceStatus,omitempty" xml:"ClogServiceStatus,omitempty"`
	Collation                    *string                                              `json:"Collation,omitempty" xml:"Collation,omitempty"`
	CreateTime                   *string                                              `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode                   *string                                              `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	DeployType                   *string                                              `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Description                  *string                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	EnableClogService            *bool                                                `json:"EnableClogService,omitempty" xml:"EnableClogService,omitempty"`
	EnableInternetAddressService *bool                                                `json:"EnableInternetAddressService,omitempty" xml:"EnableInternetAddressService,omitempty"`
	MasterIntranetAddressZone    *string                                              `json:"MasterIntranetAddressZone,omitempty" xml:"MasterIntranetAddressZone,omitempty"`
	PrimaryZone                  *string                                              `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	PrimaryZoneDeployType        *string                                              `json:"PrimaryZoneDeployType,omitempty" xml:"PrimaryZoneDeployType,omitempty"`
	Status                       *string                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantConnections            []*DescribeTenantResponseBodyTenantTenantConnections `json:"TenantConnections,omitempty" xml:"TenantConnections,omitempty" type:"Repeated"`
	TenantId                     *string                                              `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantMode                   *string                                              `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	TenantName                   *string                                              `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	TenantResource               *DescribeTenantResponseBodyTenantTenantResource      `json:"TenantResource,omitempty" xml:"TenantResource,omitempty" type:"Struct"`
	TenantZones                  []*DescribeTenantResponseBodyTenantTenantZones       `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
	VpcId                        *string                                              `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantResponseBodyTenant) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenant) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenant) SetCharset(v string) *DescribeTenantResponseBodyTenant {
	s.Charset = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetClogServiceStatus(v string) *DescribeTenantResponseBodyTenant {
	s.ClogServiceStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetCollation(v string) *DescribeTenantResponseBodyTenant {
	s.Collation = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetCreateTime(v string) *DescribeTenantResponseBodyTenant {
	s.CreateTime = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDeployMode(v string) *DescribeTenantResponseBodyTenant {
	s.DeployMode = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDeployType(v string) *DescribeTenantResponseBodyTenant {
	s.DeployType = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetDescription(v string) *DescribeTenantResponseBodyTenant {
	s.Description = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetEnableClogService(v bool) *DescribeTenantResponseBodyTenant {
	s.EnableClogService = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetEnableInternetAddressService(v bool) *DescribeTenantResponseBodyTenant {
	s.EnableInternetAddressService = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetMasterIntranetAddressZone(v string) *DescribeTenantResponseBodyTenant {
	s.MasterIntranetAddressZone = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetPrimaryZone(v string) *DescribeTenantResponseBodyTenant {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetPrimaryZoneDeployType(v string) *DescribeTenantResponseBodyTenant {
	s.PrimaryZoneDeployType = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetStatus(v string) *DescribeTenantResponseBodyTenant {
	s.Status = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantConnections(v []*DescribeTenantResponseBodyTenantTenantConnections) *DescribeTenantResponseBodyTenant {
	s.TenantConnections = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantId(v string) *DescribeTenantResponseBodyTenant {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantMode(v string) *DescribeTenantResponseBodyTenant {
	s.TenantMode = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantName(v string) *DescribeTenantResponseBodyTenant {
	s.TenantName = &v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantResource(v *DescribeTenantResponseBodyTenantTenantResource) *DescribeTenantResponseBodyTenant {
	s.TenantResource = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetTenantZones(v []*DescribeTenantResponseBodyTenantTenantZones) *DescribeTenantResponseBodyTenant {
	s.TenantZones = v
	return s
}

func (s *DescribeTenantResponseBodyTenant) SetVpcId(v string) *DescribeTenantResponseBodyTenant {
	s.VpcId = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantConnections struct {
	ConnectionRole              *string   `json:"ConnectionRole,omitempty" xml:"ConnectionRole,omitempty"`
	ConnectionZones             []*string `json:"ConnectionZones,omitempty" xml:"ConnectionZones,omitempty" type:"Repeated"`
	InternetAddress             *string   `json:"InternetAddress,omitempty" xml:"InternetAddress,omitempty"`
	InternetAddressStatus       *string   `json:"InternetAddressStatus,omitempty" xml:"InternetAddressStatus,omitempty"`
	InternetPort                *int32    `json:"InternetPort,omitempty" xml:"InternetPort,omitempty"`
	IntranetAddress             *string   `json:"IntranetAddress,omitempty" xml:"IntranetAddress,omitempty"`
	IntranetAddressMasterZoneId *string   `json:"IntranetAddressMasterZoneId,omitempty" xml:"IntranetAddressMasterZoneId,omitempty"`
	IntranetAddressSlaveZoneId  *string   `json:"IntranetAddressSlaveZoneId,omitempty" xml:"IntranetAddressSlaveZoneId,omitempty"`
	IntranetAddressStatus       *string   `json:"IntranetAddressStatus,omitempty" xml:"IntranetAddressStatus,omitempty"`
	IntranetPort                *int32    `json:"IntranetPort,omitempty" xml:"IntranetPort,omitempty"`
	VSwitchId                   *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId                       *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantConnections) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantConnections) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetConnectionRole(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.ConnectionRole = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetConnectionZones(v []*string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.ConnectionZones = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetAddress(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetAddress = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetAddressStatus(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetAddressStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetInternetPort(v int32) *DescribeTenantResponseBodyTenantTenantConnections {
	s.InternetPort = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddress(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddress = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressMasterZoneId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressMasterZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressSlaveZoneId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressSlaveZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetAddressStatus(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetAddressStatus = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetIntranetPort(v int32) *DescribeTenantResponseBodyTenantTenantConnections {
	s.IntranetPort = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetVSwitchId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.VSwitchId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantConnections) SetVpcId(v string) *DescribeTenantResponseBodyTenantTenantConnections {
	s.VpcId = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResource struct {
	Cpu      *DescribeTenantResponseBodyTenantTenantResourceCpu      `json:"Cpu,omitempty" xml:"Cpu,omitempty" type:"Struct"`
	DiskSize *DescribeTenantResponseBodyTenantTenantResourceDiskSize `json:"DiskSize,omitempty" xml:"DiskSize,omitempty" type:"Struct"`
	Memory   *DescribeTenantResponseBodyTenantTenantResourceMemory   `json:"Memory,omitempty" xml:"Memory,omitempty" type:"Struct"`
	UnitNum  *int32                                                  `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResource) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResource) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetCpu(v *DescribeTenantResponseBodyTenantTenantResourceCpu) *DescribeTenantResponseBodyTenantTenantResource {
	s.Cpu = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetDiskSize(v *DescribeTenantResponseBodyTenantTenantResourceDiskSize) *DescribeTenantResponseBodyTenantTenantResource {
	s.DiskSize = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetMemory(v *DescribeTenantResponseBodyTenantTenantResourceMemory) *DescribeTenantResponseBodyTenantTenantResource {
	s.Memory = v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResource) SetUnitNum(v int32) *DescribeTenantResponseBodyTenantTenantResource {
	s.UnitNum = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceCpu struct {
	TotalCpu *float32 `json:"TotalCpu,omitempty" xml:"TotalCpu,omitempty"`
	UnitCpu  *float32 `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	UsedCpu  *float32 `json:"UsedCpu,omitempty" xml:"UsedCpu,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceCpu) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceCpu) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetTotalCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.TotalCpu = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetUnitCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.UnitCpu = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceCpu) SetUsedCpu(v float32) *DescribeTenantResponseBodyTenantTenantResourceCpu {
	s.UsedCpu = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceDiskSize struct {
	UsedDiskSize *float32 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceDiskSize) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceDiskSize) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceDiskSize) SetUsedDiskSize(v float32) *DescribeTenantResponseBodyTenantTenantResourceDiskSize {
	s.UsedDiskSize = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantResourceMemory struct {
	TotalMemory *float32 `json:"TotalMemory,omitempty" xml:"TotalMemory,omitempty"`
	UnitMemory  *float32 `json:"UnitMemory,omitempty" xml:"UnitMemory,omitempty"`
	UsedMemory  *float32 `json:"UsedMemory,omitempty" xml:"UsedMemory,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantResourceMemory) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantResourceMemory) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetTotalMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.TotalMemory = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetUnitMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.UnitMemory = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantResourceMemory) SetUsedMemory(v float32) *DescribeTenantResponseBodyTenantTenantResourceMemory {
	s.UsedMemory = &v
	return s
}

type DescribeTenantResponseBodyTenantTenantZones struct {
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	TenantZoneId   *string `json:"TenantZoneId,omitempty" xml:"TenantZoneId,omitempty"`
	TenantZoneRole *string `json:"TenantZoneRole,omitempty" xml:"TenantZoneRole,omitempty"`
}

func (s DescribeTenantResponseBodyTenantTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponseBodyTenantTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetRegion(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.Region = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetTenantZoneId(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.TenantZoneId = &v
	return s
}

func (s *DescribeTenantResponseBodyTenantTenantZones) SetTenantZoneRole(v string) *DescribeTenantResponseBodyTenantTenantZones {
	s.TenantZoneRole = &v
	return s
}

type DescribeTenantResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantResponse) SetHeaders(v map[string]*string) *DescribeTenantResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantResponse) SetStatusCode(v int32) *DescribeTenantResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantResponse) SetBody(v *DescribeTenantResponseBody) *DescribeTenantResponse {
	s.Body = v
	return s
}

type DescribeTenantMetricsRequest struct {
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Metrics      *string `json:"Metrics,omitempty" xml:"Metrics,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantIdList *string `json:"TenantIdList,omitempty" xml:"TenantIdList,omitempty"`
	TenantName   *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeTenantMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsRequest) SetEndTime(v string) *DescribeTenantMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetInstanceId(v string) *DescribeTenantMetricsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetMetrics(v string) *DescribeTenantMetricsRequest {
	s.Metrics = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetPageNumber(v int32) *DescribeTenantMetricsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetPageSize(v int32) *DescribeTenantMetricsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetStartTime(v string) *DescribeTenantMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantId(v string) *DescribeTenantMetricsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantIdList(v string) *DescribeTenantMetricsRequest {
	s.TenantIdList = &v
	return s
}

func (s *DescribeTenantMetricsRequest) SetTenantName(v string) *DescribeTenantMetricsRequest {
	s.TenantName = &v
	return s
}

type DescribeTenantMetricsResponseBody struct {
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantMetrics *string `json:"TenantMetrics,omitempty" xml:"TenantMetrics,omitempty"`
	TotalCount    *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsResponseBody) SetRequestId(v string) *DescribeTenantMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantMetricsResponseBody) SetTenantMetrics(v string) *DescribeTenantMetricsResponseBody {
	s.TenantMetrics = &v
	return s
}

func (s *DescribeTenantMetricsResponseBody) SetTotalCount(v int32) *DescribeTenantMetricsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantMetricsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantMetricsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantMetricsResponse) SetHeaders(v map[string]*string) *DescribeTenantMetricsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantMetricsResponse) SetStatusCode(v int32) *DescribeTenantMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantMetricsResponse) SetBody(v *DescribeTenantMetricsResponseBody) *DescribeTenantMetricsResponse {
	s.Body = v
	return s
}

type DescribeTenantUserRolesResponseBody struct {
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Role      []*string `json:"Role,omitempty" xml:"Role,omitempty" type:"Repeated"`
}

func (s DescribeTenantUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantUserRolesResponseBody) SetRequestId(v string) *DescribeTenantUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantUserRolesResponseBody) SetRole(v []*string) *DescribeTenantUserRolesResponseBody {
	s.Role = v
	return s
}

type DescribeTenantUserRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUserRolesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantUserRolesResponse) SetHeaders(v map[string]*string) *DescribeTenantUserRolesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantUserRolesResponse) SetStatusCode(v int32) *DescribeTenantUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantUserRolesResponse) SetBody(v *DescribeTenantUserRolesResponseBody) *DescribeTenantUserRolesResponse {
	s.Body = v
	return s
}

type DescribeTenantUsersRequest struct {
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeTenantUsersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersRequest) SetPageNumber(v int32) *DescribeTenantUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetPageSize(v int32) *DescribeTenantUsersRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetSearchKey(v string) *DescribeTenantUsersRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetTenantId(v string) *DescribeTenantUsersRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantUsersRequest) SetUserName(v string) *DescribeTenantUsersRequest {
	s.UserName = &v
	return s
}

type DescribeTenantUsersResponseBody struct {
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantUsers []*DescribeTenantUsersResponseBodyTenantUsers `json:"TenantUsers,omitempty" xml:"TenantUsers,omitempty" type:"Repeated"`
	TotalCount  *int32                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantUsersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBody) SetRequestId(v string) *DescribeTenantUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantUsersResponseBody) SetTenantUsers(v []*DescribeTenantUsersResponseBodyTenantUsers) *DescribeTenantUsersResponseBody {
	s.TenantUsers = v
	return s
}

func (s *DescribeTenantUsersResponseBody) SetTotalCount(v int32) *DescribeTenantUsersResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantUsersResponseBodyTenantUsers struct {
	Databases   []*DescribeTenantUsersResponseBodyTenantUsersDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Description *string                                                `json:"Description,omitempty" xml:"Description,omitempty"`
	UserName    *string                                                `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserStatus  *string                                                `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
	UserType    *string                                                `json:"UserType,omitempty" xml:"UserType,omitempty"`
}

func (s DescribeTenantUsersResponseBodyTenantUsers) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBodyTenantUsers) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetDatabases(v []*DescribeTenantUsersResponseBodyTenantUsersDatabases) *DescribeTenantUsersResponseBodyTenantUsers {
	s.Databases = v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetDescription(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.Description = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserName(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserName = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserStatus(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserStatus = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsers) SetUserType(v string) *DescribeTenantUsersResponseBodyTenantUsers {
	s.UserType = &v
	return s
}

type DescribeTenantUsersResponseBodyTenantUsersDatabases struct {
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Table    *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeTenantUsersResponseBodyTenantUsersDatabases) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponseBodyTenantUsersDatabases) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetDatabase(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Database = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetRole(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Role = &v
	return s
}

func (s *DescribeTenantUsersResponseBodyTenantUsersDatabases) SetTable(v string) *DescribeTenantUsersResponseBodyTenantUsersDatabases {
	s.Table = &v
	return s
}

type DescribeTenantUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantUsersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantUsersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantUsersResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantUsersResponse) SetHeaders(v map[string]*string) *DescribeTenantUsersResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantUsersResponse) SetStatusCode(v int32) *DescribeTenantUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantUsersResponse) SetBody(v *DescribeTenantUsersResponseBody) *DescribeTenantUsersResponse {
	s.Body = v
	return s
}

type DescribeTenantZonesReadRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTenantZonesReadRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadRequest) SetInstanceId(v string) *DescribeTenantZonesReadRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantZonesReadRequest) SetTenantId(v string) *DescribeTenantZonesReadRequest {
	s.TenantId = &v
	return s
}

type DescribeTenantZonesReadResponseBody struct {
	RequestId   *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantZones []*DescribeTenantZonesReadResponseBodyTenantZones `json:"TenantZones,omitempty" xml:"TenantZones,omitempty" type:"Repeated"`
}

func (s DescribeTenantZonesReadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponseBody) SetRequestId(v string) *DescribeTenantZonesReadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBody) SetTenantZones(v []*DescribeTenantZonesReadResponseBodyTenantZones) *DescribeTenantZonesReadResponseBody {
	s.TenantZones = v
	return s
}

type DescribeTenantZonesReadResponseBodyTenantZones struct {
	IsElectable             *bool   `json:"IsElectable,omitempty" xml:"IsElectable,omitempty"`
	IsPrimary               *bool   `json:"IsPrimary,omitempty" xml:"IsPrimary,omitempty"`
	IsReadOnlyAddressMaster *bool   `json:"IsReadOnlyAddressMaster,omitempty" xml:"IsReadOnlyAddressMaster,omitempty"`
	IsReadable              *string `json:"IsReadable,omitempty" xml:"IsReadable,omitempty"`
	Zone                    *string `json:"Zone,omitempty" xml:"Zone,omitempty"`
}

func (s DescribeTenantZonesReadResponseBodyTenantZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponseBodyTenantZones) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsElectable(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsElectable = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsPrimary(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsPrimary = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsReadOnlyAddressMaster(v bool) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsReadOnlyAddressMaster = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetIsReadable(v string) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.IsReadable = &v
	return s
}

func (s *DescribeTenantZonesReadResponseBodyTenantZones) SetZone(v string) *DescribeTenantZonesReadResponseBodyTenantZones {
	s.Zone = &v
	return s
}

type DescribeTenantZonesReadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantZonesReadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantZonesReadResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantZonesReadResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantZonesReadResponse) SetHeaders(v map[string]*string) *DescribeTenantZonesReadResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantZonesReadResponse) SetStatusCode(v int32) *DescribeTenantZonesReadResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantZonesReadResponse) SetBody(v *DescribeTenantZonesReadResponseBody) *DescribeTenantZonesReadResponse {
	s.Body = v
	return s
}

type DescribeTenantsRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNumber *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey  *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantName *string `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
}

func (s DescribeTenantsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTenantsRequest) SetInstanceId(v string) *DescribeTenantsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTenantsRequest) SetPageNumber(v int32) *DescribeTenantsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTenantsRequest) SetPageSize(v int32) *DescribeTenantsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTenantsRequest) SetSearchKey(v string) *DescribeTenantsRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeTenantsRequest) SetTenantId(v string) *DescribeTenantsRequest {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantsRequest) SetTenantName(v string) *DescribeTenantsRequest {
	s.TenantName = &v
	return s
}

type DescribeTenantsResponseBody struct {
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tenants    []*DescribeTenantsResponseBodyTenants `json:"Tenants,omitempty" xml:"Tenants,omitempty" type:"Repeated"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTenantsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponseBody) SetRequestId(v string) *DescribeTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTenantsResponseBody) SetTenants(v []*DescribeTenantsResponseBodyTenants) *DescribeTenantsResponseBody {
	s.Tenants = v
	return s
}

func (s *DescribeTenantsResponseBody) SetTotalCount(v int32) *DescribeTenantsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTenantsResponseBodyTenants struct {
	Charset      *string  `json:"Charset,omitempty" xml:"Charset,omitempty"`
	Collation    *string  `json:"Collation,omitempty" xml:"Collation,omitempty"`
	Cpu          *int32   `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	CreateTime   *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DeployMode   *string  `json:"DeployMode,omitempty" xml:"DeployMode,omitempty"`
	DeployType   *string  `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Description  *string  `json:"Description,omitempty" xml:"Description,omitempty"`
	Mem          *int32   `json:"Mem,omitempty" xml:"Mem,omitempty"`
	PrimaryZone  *string  `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	Status       *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	TenantId     *string  `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	TenantMode   *string  `json:"TenantMode,omitempty" xml:"TenantMode,omitempty"`
	TenantName   *string  `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	UnitCpu      *int32   `json:"UnitCpu,omitempty" xml:"UnitCpu,omitempty"`
	UnitMem      *int32   `json:"UnitMem,omitempty" xml:"UnitMem,omitempty"`
	UnitNum      *int32   `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
	UsedDiskSize *float64 `json:"UsedDiskSize,omitempty" xml:"UsedDiskSize,omitempty"`
	VpcId        *string  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeTenantsResponseBodyTenants) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponseBodyTenants) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponseBodyTenants) SetCharset(v string) *DescribeTenantsResponseBodyTenants {
	s.Charset = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetCollation(v string) *DescribeTenantsResponseBodyTenants {
	s.Collation = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetCpu(v int32) *DescribeTenantsResponseBodyTenants {
	s.Cpu = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetCreateTime(v string) *DescribeTenantsResponseBodyTenants {
	s.CreateTime = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDeployMode(v string) *DescribeTenantsResponseBodyTenants {
	s.DeployMode = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDeployType(v string) *DescribeTenantsResponseBodyTenants {
	s.DeployType = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetDescription(v string) *DescribeTenantsResponseBodyTenants {
	s.Description = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetMem(v int32) *DescribeTenantsResponseBodyTenants {
	s.Mem = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetPrimaryZone(v string) *DescribeTenantsResponseBodyTenants {
	s.PrimaryZone = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetStatus(v string) *DescribeTenantsResponseBodyTenants {
	s.Status = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantId(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantId = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantMode(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantMode = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetTenantName(v string) *DescribeTenantsResponseBodyTenants {
	s.TenantName = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitCpu(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitCpu = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitMem(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitMem = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUnitNum(v int32) *DescribeTenantsResponseBodyTenants {
	s.UnitNum = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetUsedDiskSize(v float64) *DescribeTenantsResponseBodyTenants {
	s.UsedDiskSize = &v
	return s
}

func (s *DescribeTenantsResponseBodyTenants) SetVpcId(v string) *DescribeTenantsResponseBodyTenants {
	s.VpcId = &v
	return s
}

type DescribeTenantsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTenantsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTenantsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTenantsResponse) GoString() string {
	return s.String()
}

func (s *DescribeTenantsResponse) SetHeaders(v map[string]*string) *DescribeTenantsResponse {
	s.Headers = v
	return s
}

func (s *DescribeTenantsResponse) SetStatusCode(v int32) *DescribeTenantsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTenantsResponse) SetBody(v *DescribeTenantsResponseBody) *DescribeTenantsResponse {
	s.Body = v
	return s
}

type DescribeTimeZonesResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TimeZones *DescribeTimeZonesResponseBodyTimeZones `json:"TimeZones,omitempty" xml:"TimeZones,omitempty" type:"Struct"`
}

func (s DescribeTimeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBody) SetRequestId(v string) *DescribeTimeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTimeZonesResponseBody) SetTimeZones(v *DescribeTimeZonesResponseBodyTimeZones) *DescribeTimeZonesResponseBody {
	s.TimeZones = v
	return s
}

type DescribeTimeZonesResponseBodyTimeZones struct {
	Default *string                                       `json:"Default,omitempty" xml:"Default,omitempty"`
	List    []*DescribeTimeZonesResponseBodyTimeZonesList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTimeZonesResponseBodyTimeZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBodyTimeZones) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBodyTimeZones) SetDefault(v string) *DescribeTimeZonesResponseBodyTimeZones {
	s.Default = &v
	return s
}

func (s *DescribeTimeZonesResponseBodyTimeZones) SetList(v []*DescribeTimeZonesResponseBodyTimeZonesList) *DescribeTimeZonesResponseBodyTimeZones {
	s.List = v
	return s
}

type DescribeTimeZonesResponseBodyTimeZonesList struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	TimeZone    *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s DescribeTimeZonesResponseBodyTimeZonesList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponseBodyTimeZonesList) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponseBodyTimeZonesList) SetDescription(v string) *DescribeTimeZonesResponseBodyTimeZonesList {
	s.Description = &v
	return s
}

func (s *DescribeTimeZonesResponseBodyTimeZonesList) SetTimeZone(v string) *DescribeTimeZonesResponseBodyTimeZonesList {
	s.TimeZone = &v
	return s
}

type DescribeTimeZonesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTimeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTimeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTimeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTimeZonesResponse) SetHeaders(v map[string]*string) *DescribeTimeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeTimeZonesResponse) SetStatusCode(v int32) *DescribeTimeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTimeZonesResponse) SetBody(v *DescribeTimeZonesResponseBody) *DescribeTimeZonesResponse {
	s.Body = v
	return s
}

type DescribeTopSQLListRequest struct {
	DbName          *string                `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime         *string                `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterCondition map[string]interface{} `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp          *string                `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber      *int32                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize        *int32                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId           *string                `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord   *string                `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter *string                `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule      *string                `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue     *string                `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn      *string                `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder       *string                `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime       *string                `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId        *string                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTopSQLListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListRequest) SetDbName(v string) *DescribeTopSQLListRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetEndTime(v string) *DescribeTopSQLListRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetFilterCondition(v map[string]interface{}) *DescribeTopSQLListRequest {
	s.FilterCondition = v
	return s
}

func (s *DescribeTopSQLListRequest) SetNodeIp(v string) *DescribeTopSQLListRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetPageNumber(v int32) *DescribeTopSQLListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetPageSize(v int32) *DescribeTopSQLListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSQLId(v string) *DescribeTopSQLListRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchKeyWord(v string) *DescribeTopSQLListRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchParameter(v string) *DescribeTopSQLListRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchRule(v string) *DescribeTopSQLListRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSearchValue(v string) *DescribeTopSQLListRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSortColumn(v string) *DescribeTopSQLListRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetSortOrder(v string) *DescribeTopSQLListRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetStartTime(v string) *DescribeTopSQLListRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopSQLListRequest) SetTenantId(v string) *DescribeTopSQLListRequest {
	s.TenantId = &v
	return s
}

type DescribeTopSQLListShrinkRequest struct {
	DbName                *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	EndTime               *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterConditionShrink *string `json:"FilterCondition,omitempty" xml:"FilterCondition,omitempty"`
	NodeIp                *string `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	PageNumber            *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize              *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SQLId                 *string `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SearchKeyWord         *string `json:"SearchKeyWord,omitempty" xml:"SearchKeyWord,omitempty"`
	SearchParameter       *string `json:"SearchParameter,omitempty" xml:"SearchParameter,omitempty"`
	SearchRule            *string `json:"SearchRule,omitempty" xml:"SearchRule,omitempty"`
	SearchValue           *string `json:"SearchValue,omitempty" xml:"SearchValue,omitempty"`
	SortColumn            *string `json:"SortColumn,omitempty" xml:"SortColumn,omitempty"`
	SortOrder             *string `json:"SortOrder,omitempty" xml:"SortOrder,omitempty"`
	StartTime             *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	TenantId              *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DescribeTopSQLListShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListShrinkRequest) SetDbName(v string) *DescribeTopSQLListShrinkRequest {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetEndTime(v string) *DescribeTopSQLListShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetFilterConditionShrink(v string) *DescribeTopSQLListShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetNodeIp(v string) *DescribeTopSQLListShrinkRequest {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetPageNumber(v int32) *DescribeTopSQLListShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetPageSize(v int32) *DescribeTopSQLListShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSQLId(v string) *DescribeTopSQLListShrinkRequest {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchKeyWord(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchKeyWord = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchParameter(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchParameter = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchRule(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchRule = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSearchValue(v string) *DescribeTopSQLListShrinkRequest {
	s.SearchValue = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSortColumn(v string) *DescribeTopSQLListShrinkRequest {
	s.SortColumn = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetSortOrder(v string) *DescribeTopSQLListShrinkRequest {
	s.SortOrder = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetStartTime(v string) *DescribeTopSQLListShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeTopSQLListShrinkRequest) SetTenantId(v string) *DescribeTopSQLListShrinkRequest {
	s.TenantId = &v
	return s
}

type DescribeTopSQLListResponseBody struct {
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TopSQLList []*DescribeTopSQLListResponseBodyTopSQLList `json:"TopSQLList,omitempty" xml:"TopSQLList,omitempty" type:"Repeated"`
	TotalCount *int64                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTopSQLListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponseBody) SetRequestId(v string) *DescribeTopSQLListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTopSQLListResponseBody) SetTopSQLList(v []*DescribeTopSQLListResponseBodyTopSQLList) *DescribeTopSQLListResponseBody {
	s.TopSQLList = v
	return s
}

func (s *DescribeTopSQLListResponseBody) SetTotalCount(v int64) *DescribeTopSQLListResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeTopSQLListResponseBodyTopSQLList struct {
	AffectedRows         *int64   `json:"AffectedRows,omitempty" xml:"AffectedRows,omitempty"`
	AppWaitTime          *float32 `json:"AppWaitTime,omitempty" xml:"AppWaitTime,omitempty"`
	BlockCacheHit        *int64   `json:"BlockCacheHit,omitempty" xml:"BlockCacheHit,omitempty"`
	BlockIndexCacheHit   *int64   `json:"BlockIndexCacheHit,omitempty" xml:"BlockIndexCacheHit,omitempty"`
	BloomFilterCacheHit  *int64   `json:"BloomFilterCacheHit,omitempty" xml:"BloomFilterCacheHit,omitempty"`
	ClientIp             *string  `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	ConcurrencyWaitTime  *float32 `json:"ConcurrencyWaitTime,omitempty" xml:"ConcurrencyWaitTime,omitempty"`
	CpuTime              *float32 `json:"CpuTime,omitempty" xml:"CpuTime,omitempty"`
	DbName               *string  `json:"DbName,omitempty" xml:"DbName,omitempty"`
	DecodeTime           *float32 `json:"DecodeTime,omitempty" xml:"DecodeTime,omitempty"`
	DiskRead             *int64   `json:"DiskRead,omitempty" xml:"DiskRead,omitempty"`
	ElapsedTime          *float32 `json:"ElapsedTime,omitempty" xml:"ElapsedTime,omitempty"`
	Event                *string  `json:"Event,omitempty" xml:"Event,omitempty"`
	ExecPerSecond        *float32 `json:"ExecPerSecond,omitempty" xml:"ExecPerSecond,omitempty"`
	ExecuteTime          *float32 `json:"ExecuteTime,omitempty" xml:"ExecuteTime,omitempty"`
	Executions           *int64   `json:"Executions,omitempty" xml:"Executions,omitempty"`
	FailTimes            *int64   `json:"FailTimes,omitempty" xml:"FailTimes,omitempty"`
	GetPlanTime          *float32 `json:"GetPlanTime,omitempty" xml:"GetPlanTime,omitempty"`
	IOWaitTime           *float32 `json:"IOWaitTime,omitempty" xml:"IOWaitTime,omitempty"`
	Key                  *int64   `json:"Key,omitempty" xml:"Key,omitempty"`
	LogicalRead          *int64   `json:"LogicalRead,omitempty" xml:"LogicalRead,omitempty"`
	MaxCpuTime           *float32 `json:"MaxCpuTime,omitempty" xml:"MaxCpuTime,omitempty"`
	MaxElapsedTime       *float32 `json:"MaxElapsedTime,omitempty" xml:"MaxElapsedTime,omitempty"`
	MemstoreReadRowCount *int64   `json:"MemstoreReadRowCount,omitempty" xml:"MemstoreReadRowCount,omitempty"`
	MissPlans            *int64   `json:"MissPlans,omitempty" xml:"MissPlans,omitempty"`
	NetWaitTime          *float32 `json:"NetWaitTime,omitempty" xml:"NetWaitTime,omitempty"`
	NodeIp               *string  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	QueueTime            *float32 `json:"QueueTime,omitempty" xml:"QueueTime,omitempty"`
	RPCCount             *int64   `json:"RPCCount,omitempty" xml:"RPCCount,omitempty"`
	RemotePlans          *int64   `json:"RemotePlans,omitempty" xml:"RemotePlans,omitempty"`
	RetryCount           *int64   `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	ReturnRows           *int64   `json:"ReturnRows,omitempty" xml:"ReturnRows,omitempty"`
	RowCacheHit          *int64   `json:"RowCacheHit,omitempty" xml:"RowCacheHit,omitempty"`
	SQLId                *string  `json:"SQLId,omitempty" xml:"SQLId,omitempty"`
	SQLText              *string  `json:"SQLText,omitempty" xml:"SQLText,omitempty"`
	SQLType              *int64   `json:"SQLType,omitempty" xml:"SQLType,omitempty"`
	ScheduleTime         *float32 `json:"ScheduleTime,omitempty" xml:"ScheduleTime,omitempty"`
	SsstoreReadRowCount  *int64   `json:"SsstoreReadRowCount,omitempty" xml:"SsstoreReadRowCount,omitempty"`
	TotalWaitTime        *float32 `json:"TotalWaitTime,omitempty" xml:"TotalWaitTime,omitempty"`
	UserName             *string  `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeTopSQLListResponseBodyTopSQLList) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponseBodyTopSQLList) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetAffectedRows(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.AffectedRows = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetAppWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.AppWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBlockCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BlockCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBlockIndexCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BlockIndexCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetBloomFilterCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.BloomFilterCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetClientIp(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ClientIp = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetConcurrencyWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ConcurrencyWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetCpuTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.CpuTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDbName(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DbName = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDecodeTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DecodeTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetDiskRead(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.DiskRead = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetElapsedTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ElapsedTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetEvent(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Event = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecPerSecond(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ExecPerSecond = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecuteTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ExecuteTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetExecutions(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Executions = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetFailTimes(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.FailTimes = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetGetPlanTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.GetPlanTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetIOWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.IOWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetKey(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.Key = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetLogicalRead(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.LogicalRead = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMaxCpuTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MaxCpuTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMaxElapsedTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MaxElapsedTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMemstoreReadRowCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MemstoreReadRowCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetMissPlans(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.MissPlans = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetNetWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.NetWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetNodeIp(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.NodeIp = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetQueueTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.QueueTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRPCCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RPCCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRemotePlans(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RemotePlans = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRetryCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RetryCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetReturnRows(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ReturnRows = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetRowCacheHit(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.RowCacheHit = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLId(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLId = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLText(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLText = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSQLType(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SQLType = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetScheduleTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.ScheduleTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetSsstoreReadRowCount(v int64) *DescribeTopSQLListResponseBodyTopSQLList {
	s.SsstoreReadRowCount = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetTotalWaitTime(v float32) *DescribeTopSQLListResponseBodyTopSQLList {
	s.TotalWaitTime = &v
	return s
}

func (s *DescribeTopSQLListResponseBodyTopSQLList) SetUserName(v string) *DescribeTopSQLListResponseBodyTopSQLList {
	s.UserName = &v
	return s
}

type DescribeTopSQLListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeTopSQLListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeTopSQLListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTopSQLListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTopSQLListResponse) SetHeaders(v map[string]*string) *DescribeTopSQLListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTopSQLListResponse) SetStatusCode(v int32) *DescribeTopSQLListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTopSQLListResponse) SetBody(v *DescribeTopSQLListResponseBody) *DescribeTopSQLListResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Series     *string `json:"Series,omitempty" xml:"Series,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetDeployType(v string) *DescribeZonesRequest {
	s.DeployType = &v
	return s
}

func (s *DescribeZonesRequest) SetSeries(v string) *DescribeZonesRequest {
	s.Series = &v
	return s
}

type DescribeZonesResponseBody struct {
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	DeployType *string `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	Series     *string `json:"Series,omitempty" xml:"Series,omitempty"`
	ZoneId     *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName   *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetDeployType(v string) *DescribeZonesResponseBodyZones {
	s.DeployType = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetSeries(v string) *DescribeZonesResponseBodyZones {
	s.Series = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneName(v string) *DescribeZonesResponseBodyZones {
	s.ZoneName = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
}

type ModifyDatabaseDescriptionRequest struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyDatabaseDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionRequest) SetDatabaseName(v string) *ModifyDatabaseDescriptionRequest {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetDescription(v string) *ModifyDatabaseDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetInstanceId(v string) *ModifyDatabaseDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseDescriptionRequest) SetTenantId(v string) *ModifyDatabaseDescriptionRequest {
	s.TenantId = &v
	return s
}

type ModifyDatabaseDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDatabaseDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponseBody) SetRequestId(v string) *ModifyDatabaseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyDatabaseDescriptionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseDescriptionResponse) SetHeaders(v map[string]*string) *ModifyDatabaseDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetStatusCode(v int32) *ModifyDatabaseDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseDescriptionResponse) SetBody(v *ModifyDatabaseDescriptionResponseBody) *ModifyDatabaseDescriptionResponse {
	s.Body = v
	return s
}

type ModifyDatabaseUserRolesRequest struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Users        *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s ModifyDatabaseUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesRequest) SetDatabaseName(v string) *ModifyDatabaseUserRolesRequest {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetInstanceId(v string) *ModifyDatabaseUserRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetTenantId(v string) *ModifyDatabaseUserRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyDatabaseUserRolesRequest) SetUsers(v string) *ModifyDatabaseUserRolesRequest {
	s.Users = &v
	return s
}

type ModifyDatabaseUserRolesResponseBody struct {
	RequestId  *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantUser *ModifyDatabaseUserRolesResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Struct"`
}

func (s ModifyDatabaseUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBody) SetRequestId(v string) *ModifyDatabaseUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBody) SetTenantUser(v *ModifyDatabaseUserRolesResponseBodyTenantUser) *ModifyDatabaseUserRolesResponseBody {
	s.TenantUser = v
	return s
}

type ModifyDatabaseUserRolesResponseBodyTenantUser struct {
	DatabaseName *string                                               `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TenantId     *string                                               `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	Users        []*ModifyDatabaseUserRolesResponseBodyTenantUserUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetDatabaseName(v string) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.DatabaseName = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetTenantId(v string) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUser) SetUsers(v []*ModifyDatabaseUserRolesResponseBodyTenantUserUsers) *ModifyDatabaseUserRolesResponseBodyTenantUser {
	s.Users = v
	return s
}

type ModifyDatabaseUserRolesResponseBodyTenantUserUsers struct {
	Role     *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUserUsers) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponseBodyTenantUserUsers) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUserUsers) SetRole(v string) *ModifyDatabaseUserRolesResponseBodyTenantUserUsers {
	s.Role = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponseBodyTenantUserUsers) SetUserName(v string) *ModifyDatabaseUserRolesResponseBodyTenantUserUsers {
	s.UserName = &v
	return s
}

type ModifyDatabaseUserRolesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyDatabaseUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDatabaseUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDatabaseUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseUserRolesResponse) SetHeaders(v map[string]*string) *ModifyDatabaseUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseUserRolesResponse) SetStatusCode(v int32) *ModifyDatabaseUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseUserRolesResponse) SetBody(v *ModifyDatabaseUserRolesResponseBody) *ModifyDatabaseUserRolesResponse {
	s.Body = v
	return s
}

type ModifyInstanceNameRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) SetInstanceId(v string) *ModifyInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetInstanceName(v string) *ModifyInstanceNameRequest {
	s.InstanceName = &v
	return s
}

type ModifyInstanceNameResponseBody struct {
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponseBody) SetInstanceName(v string) *ModifyInstanceNameResponseBody {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceNameResponseBody) SetRequestId(v string) *ModifyInstanceNameResponseBody {
	s.RequestId = &v
	return s
}

type ModifyInstanceNameResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyInstanceNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyInstanceNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyInstanceNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameResponse) SetHeaders(v map[string]*string) *ModifyInstanceNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceNameResponse) SetStatusCode(v int32) *ModifyInstanceNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceNameResponse) SetBody(v *ModifyInstanceNameResponseBody) *ModifyInstanceNameResponse {
	s.Body = v
	return s
}

type ModifyParametersRequest struct {
	Dimension      *string `json:"Dimension,omitempty" xml:"Dimension,omitempty"`
	DimensionValue *string `json:"DimensionValue,omitempty" xml:"DimensionValue,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Parameters     *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) SetDimension(v string) *ModifyParametersRequest {
	s.Dimension = &v
	return s
}

func (s *ModifyParametersRequest) SetDimensionValue(v string) *ModifyParametersRequest {
	s.DimensionValue = &v
	return s
}

func (s *ModifyParametersRequest) SetInstanceId(v string) *ModifyParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

type ModifyParametersResponseBody struct {
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *ModifyParametersResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
}

func (s ModifyParametersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBody) SetRequestId(v string) *ModifyParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParametersResponseBody) SetResults(v *ModifyParametersResponseBodyResults) *ModifyParametersResponseBody {
	s.Results = v
	return s
}

type ModifyParametersResponseBodyResults struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Success *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyParametersResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponseBodyResults) SetMessage(v string) *ModifyParametersResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ModifyParametersResponseBodyResults) SetSuccess(v bool) *ModifyParametersResponseBodyResults {
	s.Success = &v
	return s
}

type ModifyParametersResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyParametersResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyParametersResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyParametersResponse) GoString() string {
	return s.String()
}

func (s *ModifyParametersResponse) SetHeaders(v map[string]*string) *ModifyParametersResponse {
	s.Headers = v
	return s
}

func (s *ModifyParametersResponse) SetStatusCode(v int32) *ModifyParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyParametersResponse) SetBody(v *ModifyParametersResponseBody) *ModifyParametersResponse {
	s.Body = v
	return s
}

type ModifySecurityIpsRequest struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIps         *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) SetInstanceId(v string) *ModifySecurityIpsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIpGroupName(v string) *ModifySecurityIpsRequest {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIps(v string) *ModifySecurityIpsRequest {
	s.SecurityIps = &v
	return s
}

type ModifySecurityIpsResponseBody struct {
	RequestId       *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityIpGroup *ModifySecurityIpsResponseBodySecurityIpGroup `json:"SecurityIpGroup,omitempty" xml:"SecurityIpGroup,omitempty" type:"Struct"`
}

func (s ModifySecurityIpsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBody) SetRequestId(v string) *ModifySecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIpsResponseBody) SetSecurityIpGroup(v *ModifySecurityIpsResponseBodySecurityIpGroup) *ModifySecurityIpsResponseBody {
	s.SecurityIpGroup = v
	return s
}

type ModifySecurityIpsResponseBodySecurityIpGroup struct {
	InstanceId          *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	SecurityIpGroupName *string `json:"SecurityIpGroupName,omitempty" xml:"SecurityIpGroupName,omitempty"`
	SecurityIps         *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifySecurityIpsResponseBodySecurityIpGroup) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponseBodySecurityIpGroup) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetInstanceId(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.InstanceId = &v
	return s
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetSecurityIpGroupName(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.SecurityIpGroupName = &v
	return s
}

func (s *ModifySecurityIpsResponseBodySecurityIpGroup) SetSecurityIps(v string) *ModifySecurityIpsResponseBodySecurityIpGroup {
	s.SecurityIps = &v
	return s
}

type ModifySecurityIpsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySecurityIpsResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsResponse) SetHeaders(v map[string]*string) *ModifySecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *ModifySecurityIpsResponse) SetStatusCode(v int32) *ModifySecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySecurityIpsResponse) SetBody(v *ModifySecurityIpsResponseBody) *ModifySecurityIpsResponse {
	s.Body = v
	return s
}

type ModifyTenantPrimaryZoneRequest struct {
	InstanceId                *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MasterIntranetAddressZone *string `json:"MasterIntranetAddressZone,omitempty" xml:"MasterIntranetAddressZone,omitempty"`
	ModifyType                *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	PrimaryZone               *string `json:"PrimaryZone,omitempty" xml:"PrimaryZone,omitempty"`
	PrimaryZoneDeployType     *string `json:"PrimaryZoneDeployType,omitempty" xml:"PrimaryZoneDeployType,omitempty"`
	TenantId                  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserVSwitchId             *string `json:"UserVSwitchId,omitempty" xml:"UserVSwitchId,omitempty"`
}

func (s ModifyTenantPrimaryZoneRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneRequest) SetInstanceId(v string) *ModifyTenantPrimaryZoneRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetMasterIntranetAddressZone(v string) *ModifyTenantPrimaryZoneRequest {
	s.MasterIntranetAddressZone = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetModifyType(v string) *ModifyTenantPrimaryZoneRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetPrimaryZone(v string) *ModifyTenantPrimaryZoneRequest {
	s.PrimaryZone = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetPrimaryZoneDeployType(v string) *ModifyTenantPrimaryZoneRequest {
	s.PrimaryZoneDeployType = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetTenantId(v string) *ModifyTenantPrimaryZoneRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantPrimaryZoneRequest) SetUserVSwitchId(v string) *ModifyTenantPrimaryZoneRequest {
	s.UserVSwitchId = &v
	return s
}

type ModifyTenantPrimaryZoneResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantPrimaryZoneResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneResponseBody) SetRequestId(v string) *ModifyTenantPrimaryZoneResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantPrimaryZoneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantPrimaryZoneResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantPrimaryZoneResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantPrimaryZoneResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantPrimaryZoneResponse) SetHeaders(v map[string]*string) *ModifyTenantPrimaryZoneResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantPrimaryZoneResponse) SetStatusCode(v int32) *ModifyTenantPrimaryZoneResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantPrimaryZoneResponse) SetBody(v *ModifyTenantPrimaryZoneResponseBody) *ModifyTenantPrimaryZoneResponse {
	s.Body = v
	return s
}

type ModifyTenantResourceRequest struct {
	Cpu        *int32  `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Memory     *int32  `json:"Memory,omitempty" xml:"Memory,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyTenantResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceRequest) SetCpu(v int32) *ModifyTenantResourceRequest {
	s.Cpu = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetInstanceId(v string) *ModifyTenantResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetMemory(v int32) *ModifyTenantResourceRequest {
	s.Memory = &v
	return s
}

func (s *ModifyTenantResourceRequest) SetTenantId(v string) *ModifyTenantResourceRequest {
	s.TenantId = &v
	return s
}

type ModifyTenantResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantId  *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ModifyTenantResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceResponseBody) SetRequestId(v string) *ModifyTenantResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantResourceResponseBody) SetTenantId(v string) *ModifyTenantResourceResponseBody {
	s.TenantId = &v
	return s
}

type ModifyTenantResourceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantResourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantResourceResponse) SetHeaders(v map[string]*string) *ModifyTenantResourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantResourceResponse) SetStatusCode(v int32) *ModifyTenantResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantResourceResponse) SetBody(v *ModifyTenantResourceResponseBody) *ModifyTenantResourceResponse {
	s.Body = v
	return s
}

type ModifyTenantUserDescriptionRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId    *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName    *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ModifyTenantUserDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionRequest) SetDescription(v string) *ModifyTenantUserDescriptionRequest {
	s.Description = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetInstanceId(v string) *ModifyTenantUserDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetTenantId(v string) *ModifyTenantUserDescriptionRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserDescriptionRequest) SetUserName(v string) *ModifyTenantUserDescriptionRequest {
	s.UserName = &v
	return s
}

type ModifyTenantUserDescriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantUserDescriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionResponseBody) SetRequestId(v string) *ModifyTenantUserDescriptionResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantUserDescriptionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantUserDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserDescriptionResponse) SetHeaders(v map[string]*string) *ModifyTenantUserDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserDescriptionResponse) SetStatusCode(v int32) *ModifyTenantUserDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantUserDescriptionResponse) SetBody(v *ModifyTenantUserDescriptionResponseBody) *ModifyTenantUserDescriptionResponse {
	s.Body = v
	return s
}

type ModifyTenantUserPasswordRequest struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId     *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserPassword *string `json:"UserPassword,omitempty" xml:"UserPassword,omitempty"`
}

func (s ModifyTenantUserPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordRequest) SetInstanceId(v string) *ModifyTenantUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetTenantId(v string) *ModifyTenantUserPasswordRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetUserName(v string) *ModifyTenantUserPasswordRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserPasswordRequest) SetUserPassword(v string) *ModifyTenantUserPasswordRequest {
	s.UserPassword = &v
	return s
}

type ModifyTenantUserPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTenantUserPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordResponseBody) SetRequestId(v string) *ModifyTenantUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

type ModifyTenantUserPasswordResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantUserPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserPasswordResponse) SetHeaders(v map[string]*string) *ModifyTenantUserPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserPasswordResponse) SetStatusCode(v int32) *ModifyTenantUserPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantUserPasswordResponse) SetBody(v *ModifyTenantUserPasswordResponseBody) *ModifyTenantUserPasswordResponse {
	s.Body = v
	return s
}

type ModifyTenantUserRolesRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ModifyType *string `json:"ModifyType,omitempty" xml:"ModifyType,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserRole   *string `json:"UserRole,omitempty" xml:"UserRole,omitempty"`
}

func (s ModifyTenantUserRolesRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesRequest) SetInstanceId(v string) *ModifyTenantUserRolesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetModifyType(v string) *ModifyTenantUserRolesRequest {
	s.ModifyType = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetTenantId(v string) *ModifyTenantUserRolesRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetUserName(v string) *ModifyTenantUserRolesRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserRolesRequest) SetUserRole(v string) *ModifyTenantUserRolesRequest {
	s.UserRole = &v
	return s
}

type ModifyTenantUserRolesResponseBody struct {
	RequestId  *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantUser *ModifyTenantUserRolesResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Struct"`
}

func (s ModifyTenantUserRolesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBody) SetRequestId(v string) *ModifyTenantUserRolesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBody) SetTenantUser(v *ModifyTenantUserRolesResponseBodyTenantUser) *ModifyTenantUserRolesResponseBody {
	s.TenantUser = v
	return s
}

type ModifyTenantUserRolesResponseBodyTenantUser struct {
	TenantId *string                                                `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName *string                                                `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserRole []*ModifyTenantUserRolesResponseBodyTenantUserUserRole `json:"UserRole,omitempty" xml:"UserRole,omitempty" type:"Repeated"`
}

func (s ModifyTenantUserRolesResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetTenantId(v string) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetUserName(v string) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUser) SetUserRole(v []*ModifyTenantUserRolesResponseBodyTenantUserUserRole) *ModifyTenantUserRolesResponseBodyTenantUser {
	s.UserRole = v
	return s
}

type ModifyTenantUserRolesResponseBodyTenantUserUserRole struct {
	Database  *string `json:"Database,omitempty" xml:"Database,omitempty"`
	IsSuccess *bool   `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	Role      *string `json:"Role,omitempty" xml:"Role,omitempty"`
	Table     *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ModifyTenantUserRolesResponseBodyTenantUserUserRole) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponseBodyTenantUserUserRole) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetDatabase(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Database = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetIsSuccess(v bool) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.IsSuccess = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetRole(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Role = &v
	return s
}

func (s *ModifyTenantUserRolesResponseBodyTenantUserUserRole) SetTable(v string) *ModifyTenantUserRolesResponseBodyTenantUserUserRole {
	s.Table = &v
	return s
}

type ModifyTenantUserRolesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantUserRolesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserRolesResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserRolesResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserRolesResponse) SetHeaders(v map[string]*string) *ModifyTenantUserRolesResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserRolesResponse) SetStatusCode(v int32) *ModifyTenantUserRolesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantUserRolesResponse) SetBody(v *ModifyTenantUserRolesResponseBody) *ModifyTenantUserRolesResponse {
	s.Body = v
	return s
}

type ModifyTenantUserStatusRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyTenantUserStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusRequest) SetInstanceId(v string) *ModifyTenantUserStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetTenantId(v string) *ModifyTenantUserStatusRequest {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetUserName(v string) *ModifyTenantUserStatusRequest {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserStatusRequest) SetUserStatus(v string) *ModifyTenantUserStatusRequest {
	s.UserStatus = &v
	return s
}

type ModifyTenantUserStatusResponseBody struct {
	RequestId  *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TenantUser []*ModifyTenantUserStatusResponseBodyTenantUser `json:"TenantUser,omitempty" xml:"TenantUser,omitempty" type:"Repeated"`
}

func (s ModifyTenantUserStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponseBody) SetRequestId(v string) *ModifyTenantUserStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBody) SetTenantUser(v []*ModifyTenantUserStatusResponseBodyTenantUser) *ModifyTenantUserStatusResponseBody {
	s.TenantUser = v
	return s
}

type ModifyTenantUserStatusResponseBodyTenantUser struct {
	TenantId   *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	UserName   *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	UserStatus *string `json:"UserStatus,omitempty" xml:"UserStatus,omitempty"`
}

func (s ModifyTenantUserStatusResponseBodyTenantUser) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponseBodyTenantUser) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetTenantId(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.TenantId = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetUserName(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.UserName = &v
	return s
}

func (s *ModifyTenantUserStatusResponseBodyTenantUser) SetUserStatus(v string) *ModifyTenantUserStatusResponseBodyTenantUser {
	s.UserStatus = &v
	return s
}

type ModifyTenantUserStatusResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyTenantUserStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyTenantUserStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyTenantUserStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyTenantUserStatusResponse) SetHeaders(v map[string]*string) *ModifyTenantUserStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyTenantUserStatusResponse) SetStatusCode(v int32) *ModifyTenantUserStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTenantUserStatusResponse) SetBody(v *ModifyTenantUserStatusResponseBody) *ModifyTenantUserStatusResponse {
	s.Body = v
	return s
}

type ReleaseOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s ReleaseOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *ReleaseOmsOpenAPIProjectRequest) SetPageNumber(v int32) *ReleaseOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectRequest) SetPageSize(v int32) *ReleaseOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectRequest) SetProjectId(v string) *ReleaseOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *ReleaseOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type ReleaseOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                          `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                          `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                            `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ReleaseOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetAdvice(v string) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetCode(v string) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetCost(v string) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetData(v bool) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetErrorDetail(v *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) *ReleaseOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetMessage(v string) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *ReleaseOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *ReleaseOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetRequestId(v string) *ReleaseOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *ReleaseOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *ReleaseOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type ReleaseOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *ReleaseOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type ReleaseOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *ReleaseOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *ReleaseOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponse) SetStatusCode(v int32) *ReleaseOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseOmsOpenAPIProjectResponse) SetBody(v *ReleaseOmsOpenAPIProjectResponseBody) *ReleaseOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type ResetOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s ResetOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *ResetOmsOpenAPIProjectRequest) SetPageNumber(v int32) *ResetOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ResetOmsOpenAPIProjectRequest) SetPageSize(v int32) *ResetOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ResetOmsOpenAPIProjectRequest) SetProjectId(v string) *ResetOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ResetOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *ResetOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type ResetOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                        `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                        `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                          `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *ResetOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ResetOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetAdvice(v string) *ResetOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetCode(v string) *ResetOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetCost(v string) *ResetOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetData(v bool) *ResetOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetErrorDetail(v *ResetOmsOpenAPIProjectResponseBodyErrorDetail) *ResetOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetMessage(v string) *ResetOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *ResetOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *ResetOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetRequestId(v string) *ResetOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *ResetOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *ResetOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type ResetOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s ResetOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s ResetOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *ResetOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *ResetOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *ResetOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *ResetOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *ResetOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type ResetOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *ResetOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *ResetOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *ResetOmsOpenAPIProjectResponse) SetStatusCode(v int32) *ResetOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetOmsOpenAPIProjectResponse) SetBody(v *ResetOmsOpenAPIProjectResponseBody) *ResetOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type ResumeOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s ResumeOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ResumeOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *ResumeOmsOpenAPIProjectRequest) SetPageNumber(v int32) *ResumeOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectRequest) SetPageSize(v int32) *ResumeOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectRequest) SetProjectId(v string) *ResumeOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *ResumeOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type ResumeOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                         `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                         `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                           `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *ResumeOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                           `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ResumeOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResumeOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetAdvice(v string) *ResumeOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetCode(v string) *ResumeOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetCost(v string) *ResumeOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetData(v bool) *ResumeOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetErrorDetail(v *ResumeOmsOpenAPIProjectResponseBodyErrorDetail) *ResumeOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetMessage(v string) *ResumeOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *ResumeOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *ResumeOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetRequestId(v string) *ResumeOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *ResumeOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *ResumeOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type ResumeOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s ResumeOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s ResumeOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *ResumeOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *ResumeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *ResumeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *ResumeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *ResumeOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type ResumeOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResumeOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResumeOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ResumeOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *ResumeOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *ResumeOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponse) SetStatusCode(v int32) *ResumeOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeOmsOpenAPIProjectResponse) SetBody(v *ResumeOmsOpenAPIProjectResponseBody) *ResumeOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type SearchOmsOpenAPIMonitorMetricRequest struct {
	BeginTime     *int64  `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	EndTime       *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	MaxPointNum   *int64  `json:"MaxPointNum,omitempty" xml:"MaxPointNum,omitempty"`
	Metric        *string `json:"Metric,omitempty" xml:"Metric,omitempty"`
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s SearchOmsOpenAPIMonitorMetricRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricRequest) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetBeginTime(v int64) *SearchOmsOpenAPIMonitorMetricRequest {
	s.BeginTime = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetEndTime(v int64) *SearchOmsOpenAPIMonitorMetricRequest {
	s.EndTime = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetMaxPointNum(v int64) *SearchOmsOpenAPIMonitorMetricRequest {
	s.MaxPointNum = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetMetric(v string) *SearchOmsOpenAPIMonitorMetricRequest {
	s.Metric = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetPageNumber(v int32) *SearchOmsOpenAPIMonitorMetricRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetPageSize(v int32) *SearchOmsOpenAPIMonitorMetricRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetProjectId(v string) *SearchOmsOpenAPIMonitorMetricRequest {
	s.ProjectId = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricRequest) SetWorkerGradeId(v string) *SearchOmsOpenAPIMonitorMetricRequest {
	s.WorkerGradeId = &v
	return s
}

type SearchOmsOpenAPIMonitorMetricResponseBody struct {
	Advice      *string                                               `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                               `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        []*SearchOmsOpenAPIMonitorMetricResponseBodyData      `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorDetail *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchOmsOpenAPIMonitorMetricResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetAdvice(v string) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Advice = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetCode(v string) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Code = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetCost(v string) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Cost = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetData(v []*SearchOmsOpenAPIMonitorMetricResponseBodyData) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Data = v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetErrorDetail(v *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetMessage(v string) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Message = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetPageNumber(v int32) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetPageSize(v int32) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetRequestId(v string) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetSuccess(v bool) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.Success = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBody) SetTotalCount(v int64) *SearchOmsOpenAPIMonitorMetricResponseBody {
	s.TotalCount = &v
	return s
}

type SearchOmsOpenAPIMonitorMetricResponseBodyData struct {
	DataPoints []*SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints `json:"DataPoints,omitempty" xml:"DataPoints,omitempty" type:"Repeated"`
	Metric     *string                                                    `json:"Metric,omitempty" xml:"Metric,omitempty"`
	Tags       map[string]*string                                         `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyData) SetDataPoints(v []*SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints) *SearchOmsOpenAPIMonitorMetricResponseBodyData {
	s.DataPoints = v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyData) SetMetric(v string) *SearchOmsOpenAPIMonitorMetricResponseBodyData {
	s.Metric = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyData) SetTags(v map[string]*string) *SearchOmsOpenAPIMonitorMetricResponseBodyData {
	s.Tags = v
	return s
}

type SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints struct {
	Timestamp *int64   `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Value     *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints) SetTimestamp(v int64) *SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints {
	s.Timestamp = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints) SetValue(v float64) *SearchOmsOpenAPIMonitorMetricResponseBodyDataDataPoints {
	s.Value = &v
	return s
}

type SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) SetCode(v string) *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) SetLevel(v string) *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) SetMessage(v string) *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail) SetProposal(v string) *SearchOmsOpenAPIMonitorMetricResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type SearchOmsOpenAPIMonitorMetricResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchOmsOpenAPIMonitorMetricResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOmsOpenAPIMonitorMetricResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIMonitorMetricResponse) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIMonitorMetricResponse) SetHeaders(v map[string]*string) *SearchOmsOpenAPIMonitorMetricResponse {
	s.Headers = v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponse) SetStatusCode(v int32) *SearchOmsOpenAPIMonitorMetricResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchOmsOpenAPIMonitorMetricResponse) SetBody(v *SearchOmsOpenAPIMonitorMetricResponseBody) *SearchOmsOpenAPIMonitorMetricResponse {
	s.Body = v
	return s
}

type SearchOmsOpenAPIProjectsRequest struct {
	DestDbTypes   []*string `json:"DestDbTypes,omitempty" xml:"DestDbTypes,omitempty" type:"Repeated"`
	LabelIds      []*string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty" type:"Repeated"`
	PageNumber    *int32    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey     *string   `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SourceDbTypes []*string `json:"SourceDbTypes,omitempty" xml:"SourceDbTypes,omitempty" type:"Repeated"`
	StatusList    []*string `json:"StatusList,omitempty" xml:"StatusList,omitempty" type:"Repeated"`
	WorkerGradeId *string   `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s SearchOmsOpenAPIProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsRequest) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsRequest) SetDestDbTypes(v []*string) *SearchOmsOpenAPIProjectsRequest {
	s.DestDbTypes = v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetLabelIds(v []*string) *SearchOmsOpenAPIProjectsRequest {
	s.LabelIds = v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetPageNumber(v int32) *SearchOmsOpenAPIProjectsRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetPageSize(v int32) *SearchOmsOpenAPIProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetSearchKey(v string) *SearchOmsOpenAPIProjectsRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetSourceDbTypes(v []*string) *SearchOmsOpenAPIProjectsRequest {
	s.SourceDbTypes = v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetStatusList(v []*string) *SearchOmsOpenAPIProjectsRequest {
	s.StatusList = v
	return s
}

func (s *SearchOmsOpenAPIProjectsRequest) SetWorkerGradeId(v string) *SearchOmsOpenAPIProjectsRequest {
	s.WorkerGradeId = &v
	return s
}

type SearchOmsOpenAPIProjectsShrinkRequest struct {
	DestDbTypesShrink   *string `json:"DestDbTypes,omitempty" xml:"DestDbTypes,omitempty"`
	LabelIdsShrink      *string `json:"LabelIds,omitempty" xml:"LabelIds,omitempty"`
	PageNumber          *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize            *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SearchKey           *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	SourceDbTypesShrink *string `json:"SourceDbTypes,omitempty" xml:"SourceDbTypes,omitempty"`
	StatusListShrink    *string `json:"StatusList,omitempty" xml:"StatusList,omitempty"`
	WorkerGradeId       *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s SearchOmsOpenAPIProjectsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsShrinkRequest) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetDestDbTypesShrink(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.DestDbTypesShrink = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetLabelIdsShrink(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.LabelIdsShrink = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetPageNumber(v int32) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetPageSize(v int32) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetSearchKey(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.SearchKey = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetSourceDbTypesShrink(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.SourceDbTypesShrink = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetStatusListShrink(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.StatusListShrink = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsShrinkRequest) SetWorkerGradeId(v string) *SearchOmsOpenAPIProjectsShrinkRequest {
	s.WorkerGradeId = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBody struct {
	Advice      *string                                          `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                          `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        []*SearchOmsOpenAPIProjectsResponseBodyData      `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	ErrorDetail *SearchOmsOpenAPIProjectsResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                           `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                           `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetAdvice(v string) *SearchOmsOpenAPIProjectsResponseBody {
	s.Advice = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetCode(v string) *SearchOmsOpenAPIProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetCost(v string) *SearchOmsOpenAPIProjectsResponseBody {
	s.Cost = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetData(v []*SearchOmsOpenAPIProjectsResponseBodyData) *SearchOmsOpenAPIProjectsResponseBody {
	s.Data = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetErrorDetail(v *SearchOmsOpenAPIProjectsResponseBodyErrorDetail) *SearchOmsOpenAPIProjectsResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetMessage(v string) *SearchOmsOpenAPIProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetPageNumber(v int32) *SearchOmsOpenAPIProjectsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetPageSize(v int32) *SearchOmsOpenAPIProjectsResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetRequestId(v string) *SearchOmsOpenAPIProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetSuccess(v bool) *SearchOmsOpenAPIProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBody) SetTotalCount(v int64) *SearchOmsOpenAPIProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyData struct {
	BusinessName       *string                                                     `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	DestConfig         *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig         `json:"DestConfig,omitempty" xml:"DestConfig,omitempty" type:"Struct"`
	Labels             []*SearchOmsOpenAPIProjectsResponseBodyDataLabels           `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	ProjectId          *string                                                     `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName        *string                                                     `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	ProjectOwner       *string                                                     `json:"ProjectOwner,omitempty" xml:"ProjectOwner,omitempty"`
	SourceConfig       *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig       `json:"SourceConfig,omitempty" xml:"SourceConfig,omitempty" type:"Struct"`
	Steps              []*SearchOmsOpenAPIProjectsResponseBodyDataSteps            `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	TransferMapping    *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping    `json:"TransferMapping,omitempty" xml:"TransferMapping,omitempty" type:"Struct"`
	TransferStepConfig *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig `json:"TransferStepConfig,omitempty" xml:"TransferStepConfig,omitempty" type:"Struct"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetBusinessName(v string) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetDestConfig(v *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.DestConfig = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetLabels(v []*SearchOmsOpenAPIProjectsResponseBodyDataLabels) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.Labels = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetProjectId(v string) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetProjectName(v string) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetProjectOwner(v string) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.ProjectOwner = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetSourceConfig(v *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.SourceConfig = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetSteps(v []*SearchOmsOpenAPIProjectsResponseBodyDataSteps) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.Steps = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetTransferMapping(v *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.TransferMapping = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyData) SetTransferStepConfig(v *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) *SearchOmsOpenAPIProjectsResponseBodyData {
	s.TransferStepConfig = v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataDestConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetEnableMsgTrace(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetEndpointId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.EndpointId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetEndpointType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.EndpointType = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetMsgTags(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.MsgTags = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetPartition(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.Partition = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetPartitionMode(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.PartitionMode = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetProducerGroup(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.ProducerGroup = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetSendMsgTimeout(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetSequenceEnable(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.SequenceEnable = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetSequenceStartTimestamp(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetSerializerType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.SerializerType = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig) SetTopicType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataDestConfig {
	s.TopicType = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataLabels struct {
	Count   *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataLabels) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataLabels) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataLabels) SetCount(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataLabels {
	s.Count = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataLabels) SetCreator(v string) *SearchOmsOpenAPIProjectsResponseBodyDataLabels {
	s.Creator = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataLabels) SetId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataLabels {
	s.Id = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataLabels) SetName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataLabels {
	s.Name = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig struct {
	EnableMsgTrace         *bool   `json:"EnableMsgTrace,omitempty" xml:"EnableMsgTrace,omitempty"`
	EndpointId             *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	EndpointType           *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	MsgTags                *string `json:"MsgTags,omitempty" xml:"MsgTags,omitempty"`
	Partition              *int32  `json:"Partition,omitempty" xml:"Partition,omitempty"`
	PartitionMode          *string `json:"PartitionMode,omitempty" xml:"PartitionMode,omitempty"`
	ProducerGroup          *string `json:"ProducerGroup,omitempty" xml:"ProducerGroup,omitempty"`
	SendMsgTimeout         *int64  `json:"SendMsgTimeout,omitempty" xml:"SendMsgTimeout,omitempty"`
	SequenceEnable         *bool   `json:"SequenceEnable,omitempty" xml:"SequenceEnable,omitempty"`
	SequenceStartTimestamp *int64  `json:"SequenceStartTimestamp,omitempty" xml:"SequenceStartTimestamp,omitempty"`
	SerializerType         *string `json:"SerializerType,omitempty" xml:"SerializerType,omitempty"`
	TopicType              *string `json:"TopicType,omitempty" xml:"TopicType,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetEnableMsgTrace(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.EnableMsgTrace = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetEndpointId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.EndpointId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetEndpointType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.EndpointType = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetMsgTags(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.MsgTags = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetPartition(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.Partition = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetPartitionMode(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.PartitionMode = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetProducerGroup(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.ProducerGroup = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetSendMsgTimeout(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.SendMsgTimeout = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetSequenceEnable(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.SequenceEnable = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetSequenceStartTimestamp(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.SequenceStartTimestamp = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetSerializerType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.SerializerType = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig) SetTopicType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSourceConfig {
	s.TopicType = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataSteps struct {
	EstimatedRemainingSeconds *int64                                                  `json:"EstimatedRemainingSeconds,omitempty" xml:"EstimatedRemainingSeconds,omitempty"`
	ExtraInfo                 *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty" type:"Struct"`
	FinishTime                *string                                                 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	Interactive               *bool                                                   `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	StartTime                 *string                                                 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StepDescription           *string                                                 `json:"StepDescription,omitempty" xml:"StepDescription,omitempty"`
	StepInfo                  *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo  `json:"StepInfo,omitempty" xml:"StepInfo,omitempty" type:"Struct"`
	StepName                  *string                                                 `json:"StepName,omitempty" xml:"StepName,omitempty"`
	StepOrder                 *int32                                                  `json:"StepOrder,omitempty" xml:"StepOrder,omitempty"`
	StepProgress              *int32                                                  `json:"StepProgress,omitempty" xml:"StepProgress,omitempty"`
	StepStatus                *string                                                 `json:"StepStatus,omitempty" xml:"StepStatus,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataSteps) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataSteps) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetEstimatedRemainingSeconds(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.EstimatedRemainingSeconds = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetExtraInfo(v *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.ExtraInfo = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetFinishTime(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.FinishTime = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetInteractive(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.Interactive = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStartTime(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StartTime = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepDescription(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepDescription = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepInfo(v *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepInfo = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepOrder(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepOrder = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepProgress(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepProgress = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataSteps) SetStepStatus(v string) *SearchOmsOpenAPIProjectsResponseBodyDataSteps {
	s.StepStatus = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo struct {
	ErrorCode    *string                                                               `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorDetails []*SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails `json:"ErrorDetails,omitempty" xml:"ErrorDetails,omitempty" type:"Repeated"`
	ErrorMsg     *string                                                               `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ErrorParam   map[string]*string                                                    `json:"ErrorParam,omitempty" xml:"ErrorParam,omitempty"`
	FailedTime   *string                                                               `json:"FailedTime,omitempty" xml:"FailedTime,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) SetErrorCode(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo {
	s.ErrorCode = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) SetErrorDetails(v []*SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo {
	s.ErrorDetails = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) SetErrorMsg(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo {
	s.ErrorMsg = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) SetErrorParam(v map[string]*string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo {
	s.ErrorParam = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo) SetFailedTime(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfo {
	s.FailedTime = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) SetCode(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails {
	s.Code = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) SetLevel(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails {
	s.Level = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) SetMessage(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails {
	s.Message = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails) SetProposal(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsExtraInfoErrorDetails {
	s.Proposal = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo struct {
	Capacity                      *int64                                                                              `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	Checkpoint                    *string                                                                             `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	ConnectorFullProgressOverview *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview `json:"ConnectorFullProgressOverview,omitempty" xml:"ConnectorFullProgressOverview,omitempty" type:"Struct"`
	DeployId                      *string                                                                             `json:"DeployId,omitempty" xml:"DeployId,omitempty"`
	DstIops                       *int64                                                                              `json:"DstIops,omitempty" xml:"DstIops,omitempty"`
	DstRps                        *int64                                                                              `json:"DstRps,omitempty" xml:"DstRps,omitempty"`
	DstRpsRef                     *int64                                                                              `json:"DstRpsRef,omitempty" xml:"DstRpsRef,omitempty"`
	DstRt                         *int64                                                                              `json:"DstRt,omitempty" xml:"DstRt,omitempty"`
	DstRtRef                      *int64                                                                              `json:"DstRtRef,omitempty" xml:"DstRtRef,omitempty"`
	Gmt                           *int64                                                                              `json:"Gmt,omitempty" xml:"Gmt,omitempty"`
	Inconsistencies               *int64                                                                              `json:"Inconsistencies,omitempty" xml:"Inconsistencies,omitempty"`
	IncrTimestampCheckpoint       *int64                                                                              `json:"IncrTimestampCheckpoint,omitempty" xml:"IncrTimestampCheckpoint,omitempty"`
	JobId                         *string                                                                             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProcessedRecords              *int64                                                                              `json:"ProcessedRecords,omitempty" xml:"ProcessedRecords,omitempty"`
	Skipped                       *bool                                                                               `json:"Skipped,omitempty" xml:"Skipped,omitempty"`
	SrcIops                       *int64                                                                              `json:"SrcIops,omitempty" xml:"SrcIops,omitempty"`
	SrcIopsRef                    *int64                                                                              `json:"SrcIopsRef,omitempty" xml:"SrcIopsRef,omitempty"`
	SrcRps                        *int64                                                                              `json:"SrcRps,omitempty" xml:"SrcRps,omitempty"`
	SrcRpsRef                     *int64                                                                              `json:"SrcRpsRef,omitempty" xml:"SrcRpsRef,omitempty"`
	SrcRt                         *int64                                                                              `json:"SrcRt,omitempty" xml:"SrcRt,omitempty"`
	SrcRtRef                      *int64                                                                              `json:"SrcRtRef,omitempty" xml:"SrcRtRef,omitempty"`
	Validated                     *bool                                                                               `json:"Validated,omitempty" xml:"Validated,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetCapacity(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Capacity = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetCheckpoint(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Checkpoint = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetConnectorFullProgressOverview(v *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.ConnectorFullProgressOverview = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDeployId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DeployId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDstIops(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DstIops = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDstRps(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DstRps = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDstRpsRef(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DstRpsRef = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDstRt(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DstRt = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetDstRtRef(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.DstRtRef = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetGmt(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Gmt = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetInconsistencies(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Inconsistencies = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetIncrTimestampCheckpoint(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.IncrTimestampCheckpoint = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetJobId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.JobId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetProcessedRecords(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.ProcessedRecords = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSkipped(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Skipped = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcIops(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcIops = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcIopsRef(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcIopsRef = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcRps(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcRps = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcRpsRef(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcRpsRef = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcRt(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcRt = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetSrcRtRef(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.SrcRtRef = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo) SetValidated(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfo {
	s.Validated = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview struct {
	EstimatedRemainingTimeOfSec *int64 `json:"EstimatedRemainingTimeOfSec,omitempty" xml:"EstimatedRemainingTimeOfSec,omitempty"`
	EstimatedTotalCount         *int64 `json:"EstimatedTotalCount,omitempty" xml:"EstimatedTotalCount,omitempty"`
	FinishedCount               *int64 `json:"FinishedCount,omitempty" xml:"FinishedCount,omitempty"`
	Progress                    *int32 `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetEstimatedRemainingTimeOfSec(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.EstimatedRemainingTimeOfSec = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetEstimatedTotalCount(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.EstimatedTotalCount = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetFinishedCount(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.FinishedCount = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview) SetProgress(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataStepsStepInfoConnectorFullProgressOverview {
	s.Progress = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping struct {
	Databases []*SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases `json:"Databases,omitempty" xml:"Databases,omitempty" type:"Repeated"`
	Mode      *string                                                             `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping) SetDatabases(v []*SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping {
	s.Databases = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping) SetMode(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMapping {
	s.Mode = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases struct {
	DatabaseId   *string                                                                   `json:"DatabaseId,omitempty" xml:"DatabaseId,omitempty"`
	DatabaseName *string                                                                   `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	MappedName   *string                                                                   `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	Tables       []*SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
	TenantName   *string                                                                   `json:"TenantName,omitempty" xml:"TenantName,omitempty"`
	Type         *string                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetDatabaseId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.DatabaseId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetDatabaseName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.DatabaseName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetMappedName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.MappedName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetTables(v []*SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.Tables = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetTenantName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.TenantName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases) SetType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabases {
	s.Type = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables struct {
	AdbTableSchema *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema `json:"AdbTableSchema,omitempty" xml:"AdbTableSchema,omitempty" type:"Struct"`
	FilterColumns  []*string                                                                             `json:"FilterColumns,omitempty" xml:"FilterColumns,omitempty" type:"Repeated"`
	MappedName     *string                                                                               `json:"MappedName,omitempty" xml:"MappedName,omitempty"`
	ShardColumns   []*string                                                                             `json:"ShardColumns,omitempty" xml:"ShardColumns,omitempty" type:"Repeated"`
	TableId        *string                                                                               `json:"TableId,omitempty" xml:"TableId,omitempty"`
	TableName      *string                                                                               `json:"TableName,omitempty" xml:"TableName,omitempty"`
	Type           *string                                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	WhereClause    *string                                                                               `json:"WhereClause,omitempty" xml:"WhereClause,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetAdbTableSchema(v *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.AdbTableSchema = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetFilterColumns(v []*string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.FilterColumns = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetMappedName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.MappedName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetShardColumns(v []*string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.ShardColumns = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetTableId(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.TableId = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetTableName(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.TableName = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.Type = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables) SetWhereClause(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTables {
	s.WhereClause = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema struct {
	DistributedKeys    []*string `json:"DistributedKeys,omitempty" xml:"DistributedKeys,omitempty" type:"Repeated"`
	PartitionLifeCycle *int32    `json:"PartitionLifeCycle,omitempty" xml:"PartitionLifeCycle,omitempty"`
	PartitionStatement *string   `json:"PartitionStatement,omitempty" xml:"PartitionStatement,omitempty"`
	PrimaryKeys        []*string `json:"PrimaryKeys,omitempty" xml:"PrimaryKeys,omitempty" type:"Repeated"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetDistributedKeys(v []*string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.DistributedKeys = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPartitionLifeCycle(v int32) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionLifeCycle = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPartitionStatement(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PartitionStatement = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema) SetPrimaryKeys(v []*string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferMappingDatabasesTablesAdbTableSchema {
	s.PrimaryKeys = v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig struct {
	EnableFullSync             *bool                                                                                 `json:"EnableFullSync,omitempty" xml:"EnableFullSync,omitempty"`
	EnableIncrSync             *bool                                                                                 `json:"EnableIncrSync,omitempty" xml:"EnableIncrSync,omitempty"`
	EnableStructSync           *bool                                                                                 `json:"EnableStructSync,omitempty" xml:"EnableStructSync,omitempty"`
	IncrSyncStepTransferConfig *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig `json:"IncrSyncStepTransferConfig,omitempty" xml:"IncrSyncStepTransferConfig,omitempty" type:"Struct"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) SetEnableFullSync(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig {
	s.EnableFullSync = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) SetEnableIncrSync(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig {
	s.EnableIncrSync = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) SetEnableStructSync(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig {
	s.EnableStructSync = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig) SetIncrSyncStepTransferConfig(v *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfig {
	s.IncrSyncStepTransferConfig = v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig struct {
	RecordTypeList          []*string `json:"RecordTypeList,omitempty" xml:"RecordTypeList,omitempty" type:"Repeated"`
	StartTimestamp          *int64    `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	StoreLogKeptHour        *int64    `json:"StoreLogKeptHour,omitempty" xml:"StoreLogKeptHour,omitempty"`
	StoreTransactionEnabled *bool     `json:"StoreTransactionEnabled,omitempty" xml:"StoreTransactionEnabled,omitempty"`
	TransferStepType        *string   `json:"TransferStepType,omitempty" xml:"TransferStepType,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetRecordTypeList(v []*string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.RecordTypeList = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStartTimestamp(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StartTimestamp = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStoreLogKeptHour(v int64) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreLogKeptHour = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetStoreTransactionEnabled(v bool) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.StoreTransactionEnabled = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig) SetTransferStepType(v string) *SearchOmsOpenAPIProjectsResponseBodyDataTransferStepConfigIncrSyncStepTransferConfig {
	s.TransferStepType = &v
	return s
}

type SearchOmsOpenAPIProjectsResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s SearchOmsOpenAPIProjectsResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponseBodyErrorDetail) SetCode(v string) *SearchOmsOpenAPIProjectsResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyErrorDetail) SetLevel(v string) *SearchOmsOpenAPIProjectsResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyErrorDetail) SetMessage(v string) *SearchOmsOpenAPIProjectsResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponseBodyErrorDetail) SetProposal(v string) *SearchOmsOpenAPIProjectsResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type SearchOmsOpenAPIProjectsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchOmsOpenAPIProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchOmsOpenAPIProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchOmsOpenAPIProjectsResponse) GoString() string {
	return s.String()
}

func (s *SearchOmsOpenAPIProjectsResponse) SetHeaders(v map[string]*string) *SearchOmsOpenAPIProjectsResponse {
	s.Headers = v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponse) SetStatusCode(v int32) *SearchOmsOpenAPIProjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchOmsOpenAPIProjectsResponse) SetBody(v *SearchOmsOpenAPIProjectsResponseBody) *SearchOmsOpenAPIProjectsResponse {
	s.Body = v
	return s
}

type StartOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s StartOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s StartOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *StartOmsOpenAPIProjectRequest) SetPageNumber(v int32) *StartOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *StartOmsOpenAPIProjectRequest) SetPageSize(v int32) *StartOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *StartOmsOpenAPIProjectRequest) SetProjectId(v string) *StartOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *StartOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *StartOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type StartOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                        `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                        `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                          `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *StartOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                        `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                          `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                         `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s StartOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *StartOmsOpenAPIProjectResponseBody) SetAdvice(v string) *StartOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetCode(v string) *StartOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetCost(v string) *StartOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetData(v bool) *StartOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetErrorDetail(v *StartOmsOpenAPIProjectResponseBodyErrorDetail) *StartOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetMessage(v string) *StartOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *StartOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *StartOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetRequestId(v string) *StartOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *StartOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *StartOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type StartOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s StartOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s StartOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *StartOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *StartOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *StartOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *StartOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *StartOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type StartOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s StartOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *StartOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *StartOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *StartOmsOpenAPIProjectResponse) SetStatusCode(v int32) *StartOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *StartOmsOpenAPIProjectResponse) SetBody(v *StartOmsOpenAPIProjectResponseBody) *StartOmsOpenAPIProjectResponse {
	s.Body = v
	return s
}

type StopOmsOpenAPIProjectRequest struct {
	PageNumber    *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize      *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId     *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	WorkerGradeId *string `json:"WorkerGradeId,omitempty" xml:"WorkerGradeId,omitempty"`
}

func (s StopOmsOpenAPIProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s StopOmsOpenAPIProjectRequest) GoString() string {
	return s.String()
}

func (s *StopOmsOpenAPIProjectRequest) SetPageNumber(v int32) *StopOmsOpenAPIProjectRequest {
	s.PageNumber = &v
	return s
}

func (s *StopOmsOpenAPIProjectRequest) SetPageSize(v int32) *StopOmsOpenAPIProjectRequest {
	s.PageSize = &v
	return s
}

func (s *StopOmsOpenAPIProjectRequest) SetProjectId(v string) *StopOmsOpenAPIProjectRequest {
	s.ProjectId = &v
	return s
}

func (s *StopOmsOpenAPIProjectRequest) SetWorkerGradeId(v string) *StopOmsOpenAPIProjectRequest {
	s.WorkerGradeId = &v
	return s
}

type StopOmsOpenAPIProjectResponseBody struct {
	Advice      *string                                       `json:"Advice,omitempty" xml:"Advice,omitempty"`
	Code        *string                                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Cost        *string                                       `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Data        *bool                                         `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrorDetail *StopOmsOpenAPIProjectResponseBodyErrorDetail `json:"ErrorDetail,omitempty" xml:"ErrorDetail,omitempty" type:"Struct"`
	Message     *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber  *int32                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success     *bool                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalCount  *int64                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s StopOmsOpenAPIProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopOmsOpenAPIProjectResponseBody) GoString() string {
	return s.String()
}

func (s *StopOmsOpenAPIProjectResponseBody) SetAdvice(v string) *StopOmsOpenAPIProjectResponseBody {
	s.Advice = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetCode(v string) *StopOmsOpenAPIProjectResponseBody {
	s.Code = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetCost(v string) *StopOmsOpenAPIProjectResponseBody {
	s.Cost = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetData(v bool) *StopOmsOpenAPIProjectResponseBody {
	s.Data = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetErrorDetail(v *StopOmsOpenAPIProjectResponseBodyErrorDetail) *StopOmsOpenAPIProjectResponseBody {
	s.ErrorDetail = v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetMessage(v string) *StopOmsOpenAPIProjectResponseBody {
	s.Message = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetPageNumber(v int32) *StopOmsOpenAPIProjectResponseBody {
	s.PageNumber = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetPageSize(v int32) *StopOmsOpenAPIProjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetRequestId(v string) *StopOmsOpenAPIProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetSuccess(v bool) *StopOmsOpenAPIProjectResponseBody {
	s.Success = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBody) SetTotalCount(v int64) *StopOmsOpenAPIProjectResponseBody {
	s.TotalCount = &v
	return s
}

type StopOmsOpenAPIProjectResponseBodyErrorDetail struct {
	Code     *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Level    *string `json:"Level,omitempty" xml:"Level,omitempty"`
	Message  *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Proposal *string `json:"Proposal,omitempty" xml:"Proposal,omitempty"`
}

func (s StopOmsOpenAPIProjectResponseBodyErrorDetail) String() string {
	return tea.Prettify(s)
}

func (s StopOmsOpenAPIProjectResponseBodyErrorDetail) GoString() string {
	return s.String()
}

func (s *StopOmsOpenAPIProjectResponseBodyErrorDetail) SetCode(v string) *StopOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Code = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBodyErrorDetail) SetLevel(v string) *StopOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Level = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBodyErrorDetail) SetMessage(v string) *StopOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Message = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponseBodyErrorDetail) SetProposal(v string) *StopOmsOpenAPIProjectResponseBodyErrorDetail {
	s.Proposal = &v
	return s
}

type StopOmsOpenAPIProjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopOmsOpenAPIProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopOmsOpenAPIProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s StopOmsOpenAPIProjectResponse) GoString() string {
	return s.String()
}

func (s *StopOmsOpenAPIProjectResponse) SetHeaders(v map[string]*string) *StopOmsOpenAPIProjectResponse {
	s.Headers = v
	return s
}

func (s *StopOmsOpenAPIProjectResponse) SetStatusCode(v int32) *StopOmsOpenAPIProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *StopOmsOpenAPIProjectResponse) SetBody(v *StopOmsOpenAPIProjectResponseBody) *StopOmsOpenAPIProjectResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("oceanbasepro"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateDatabaseWithOptions(request *CreateDatabaseRequest, runtime *util.RuntimeOptions) (_result *CreateDatabaseResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Collation)) {
		body["Collation"] = request.Collation
	}

	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Encoding)) {
		body["Encoding"] = request.Encoding
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDatabase"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
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
	_result = &CreateDatabaseResponse{}
	_body, _err := client.CreateDatabaseWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateInstanceWithOptions(request *CreateInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenewPeriod)) {
		body["AutoRenewPeriod"] = request.AutoRenewPeriod
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.DiskSize)) {
		body["DiskSize"] = request.DiskSize
	}

	if !tea.BoolValue(util.IsUnset(request.DiskType)) {
		body["DiskType"] = request.DiskType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceClass)) {
		body["InstanceClass"] = request.InstanceClass
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.ObVersion)) {
		body["ObVersion"] = request.ObVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		body["Series"] = request.Series
	}

	if !tea.BoolValue(util.IsUnset(request.Zones)) {
		body["Zones"] = request.Zones
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateInstance"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateInstance(request *CreateInstanceRequest) (_result *CreateInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateInstanceResponse{}
	_body, _err := client.CreateInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOmsMysqlDataSourceWithOptions(request *CreateOmsMysqlDataSourceRequest, runtime *util.RuntimeOptions) (_result *CreateOmsMysqlDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DgDatabaseId)) {
		body["DgDatabaseId"] = request.DgDatabaseId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		body["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Password)) {
		body["Password"] = request.Password
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.Schema)) {
		body["Schema"] = request.Schema
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["Type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		body["VpcId"] = request.VpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOmsMysqlDataSource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOmsMysqlDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOmsMysqlDataSource(request *CreateOmsMysqlDataSourceRequest) (_result *CreateOmsMysqlDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOmsMysqlDataSourceResponse{}
	_body, _err := client.CreateOmsMysqlDataSourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateOmsOpenAPIProjectWithOptions(tmpReq *CreateOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *CreateOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateOmsOpenAPIProjectShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.DestConfig))) {
		request.DestConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.DestConfig), tea.String("DestConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LabelIds)) {
		request.LabelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelIds, tea.String("LabelIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.SourceConfig))) {
		request.SourceConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.SourceConfig), tea.String("SourceConfig"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.TransferMapping))) {
		request.TransferMappingShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.TransferMapping), tea.String("TransferMapping"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(tmpReq.TransferStepConfig))) {
		request.TransferStepConfigShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tea.ToMap(tmpReq.TransferStepConfig), tea.String("TransferStepConfig"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BusinessName)) {
		body["BusinessName"] = request.BusinessName
	}

	if !tea.BoolValue(util.IsUnset(request.DestConfigShrink)) {
		body["DestConfig"] = request.DestConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LabelIdsShrink)) {
		body["LabelIds"] = request.LabelIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.SourceConfigShrink)) {
		body["SourceConfig"] = request.SourceConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TransferMappingShrink)) {
		body["TransferMapping"] = request.TransferMappingShrink
	}

	if !tea.BoolValue(util.IsUnset(request.TransferStepConfigShrink)) {
		body["TransferStepConfig"] = request.TransferStepConfigShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateOmsOpenAPIProject(request *CreateOmsOpenAPIProjectRequest) (_result *CreateOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateOmsOpenAPIProjectResponse{}
	_body, _err := client.CreateOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSecurityIpGroupWithOptions(request *CreateSecurityIpGroupRequest, runtime *util.RuntimeOptions) (_result *CreateSecurityIpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		body["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSecurityIpGroup"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSecurityIpGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSecurityIpGroup(request *CreateSecurityIpGroupRequest) (_result *CreateSecurityIpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSecurityIpGroupResponse{}
	_body, _err := client.CreateSecurityIpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantWithOptions(request *CreateTenantRequest, runtime *util.RuntimeOptions) (_result *CreateTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Charset)) {
		body["Charset"] = request.Charset
	}

	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZone)) {
		body["PrimaryZone"] = request.PrimaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.TenantMode)) {
		body["TenantMode"] = request.TenantMode
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZone)) {
		body["TimeZone"] = request.TimeZone
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		body["UnitNum"] = request.UnitNum
	}

	if !tea.BoolValue(util.IsUnset(request.UserVSwitchId)) {
		body["UserVSwitchId"] = request.UserVSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.UserVpcId)) {
		body["UserVpcId"] = request.UserVpcId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenant"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenant(request *CreateTenantRequest) (_result *CreateTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantResponse{}
	_body, _err := client.CreateTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantReadOnlyConnectionWithOptions(request *CreateTenantReadOnlyConnectionRequest, runtime *util.RuntimeOptions) (_result *CreateTenantReadOnlyConnectionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		body["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantReadOnlyConnection"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantReadOnlyConnectionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenantReadOnlyConnection(request *CreateTenantReadOnlyConnectionRequest) (_result *CreateTenantReadOnlyConnectionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantReadOnlyConnectionResponse{}
	_body, _err := client.CreateTenantReadOnlyConnectionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTenantUserWithOptions(request *CreateTenantUserRequest, runtime *util.RuntimeOptions) (_result *CreateTenantUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Roles)) {
		body["Roles"] = request.Roles
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPassword)) {
		body["UserPassword"] = request.UserPassword
	}

	if !tea.BoolValue(util.IsUnset(request.UserType)) {
		body["UserType"] = request.UserType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTenantUser"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTenantUserResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTenantUser(request *CreateTenantUserRequest) (_result *CreateTenantUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTenantUserResponse{}
	_body, _err := client.CreateTenantUserWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDatabasesWithOptions(request *DeleteDatabasesRequest, runtime *util.RuntimeOptions) (_result *DeleteDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseNames)) {
		body["DatabaseNames"] = request.DatabaseNames
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDatabases"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDatabases(request *DeleteDatabasesRequest) (_result *DeleteDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDatabasesResponse{}
	_body, _err := client.DeleteDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstancesWithOptions(request *DeleteInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BackupRetainMode)) {
		body["BackupRetainMode"] = request.BackupRetainMode
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceIds)) {
		body["InstanceIds"] = request.InstanceIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstances"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstances(request *DeleteInstancesRequest) (_result *DeleteInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.DeleteInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteOmsOpenAPIProjectWithOptions(request *DeleteOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *DeleteOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteOmsOpenAPIProject(request *DeleteOmsOpenAPIProjectRequest) (_result *DeleteOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteOmsOpenAPIProjectResponse{}
	_body, _err := client.DeleteOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSecurityIpGroupWithOptions(request *DeleteSecurityIpGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteSecurityIpGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSecurityIpGroup"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSecurityIpGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSecurityIpGroup(request *DeleteSecurityIpGroupRequest) (_result *DeleteSecurityIpGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSecurityIpGroupResponse{}
	_body, _err := client.DeleteSecurityIpGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTenantUsersWithOptions(request *DeleteTenantUsersRequest, runtime *util.RuntimeOptions) (_result *DeleteTenantUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTenantUsers"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTenantUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTenantUsers(request *DeleteTenantUsersRequest) (_result *DeleteTenantUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTenantUsersResponse{}
	_body, _err := client.DeleteTenantUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTenantsWithOptions(request *DeleteTenantsRequest, runtime *util.RuntimeOptions) (_result *DeleteTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantIds)) {
		body["TenantIds"] = request.TenantIds
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteTenants"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTenants(request *DeleteTenantsRequest) (_result *DeleteTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteTenantsResponse{}
	_body, _err := client.DeleteTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAnomalySQLListWithOptions(tmpReq *DescribeAnomalySQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeAnomalySQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeAnomalySQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAnomalySQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAnomalySQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAnomalySQLList(request *DescribeAnomalySQLListRequest) (_result *DescribeAnomalySQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAnomalySQLListResponse{}
	_body, _err := client.DescribeAnomalySQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableCpuResourceWithOptions(request *DescribeAvailableCpuResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableCpuResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableCpuResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableCpuResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableCpuResource(request *DescribeAvailableCpuResourceRequest) (_result *DescribeAvailableCpuResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableCpuResourceResponse{}
	_body, _err := client.DescribeAvailableCpuResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAvailableMemResourceWithOptions(request *DescribeAvailableMemResourceRequest, runtime *util.RuntimeOptions) (_result *DescribeAvailableMemResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CpuNum)) {
		body["CpuNum"] = request.CpuNum
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UnitNum)) {
		body["UnitNum"] = request.UnitNum
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeAvailableMemResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeAvailableMemResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAvailableMemResource(request *DescribeAvailableMemResourceRequest) (_result *DescribeAvailableMemResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAvailableMemResourceResponse{}
	_body, _err := client.DescribeAvailableMemResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeCharsetWithOptions(request *DescribeCharsetRequest, runtime *util.RuntimeOptions) (_result *DescribeCharsetResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TenantMode)) {
		body["TenantMode"] = request.TenantMode
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCharset"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeCharsetResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeCharset(request *DescribeCharsetRequest) (_result *DescribeCharsetResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeCharsetResponse{}
	_body, _err := client.DescribeCharsetWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDatabasesWithOptions(request *DescribeDatabasesRequest, runtime *util.RuntimeOptions) (_result *DescribeDatabasesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.WithTables)) {
		body["WithTables"] = request.WithTables
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDatabases"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDatabases(request *DescribeDatabasesRequest) (_result *DescribeDatabasesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDatabasesResponse{}
	_body, _err := client.DescribeDatabasesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceWithOptions(request *DescribeInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstance"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstance(request *DescribeInstanceRequest) (_result *DescribeInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceResponse{}
	_body, _err := client.DescribeInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceCreatableZoneWithOptions(request *DescribeInstanceCreatableZoneRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceCreatableZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceCreatableZone"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceCreatableZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceCreatableZone(request *DescribeInstanceCreatableZoneRequest) (_result *DescribeInstanceCreatableZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceCreatableZoneResponse{}
	_body, _err := client.DescribeInstanceCreatableZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTenantModesWithOptions(request *DescribeInstanceTenantModesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTenantModesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceTenantModes"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTenantModesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTenantModes(request *DescribeInstanceTenantModesRequest) (_result *DescribeInstanceTenantModesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTenantModesResponse{}
	_body, _err := client.DescribeInstanceTenantModesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstanceTopologyWithOptions(request *DescribeInstanceTopologyRequest, runtime *util.RuntimeOptions) (_result *DescribeInstanceTopologyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstanceTopology"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstanceTopologyResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstanceTopology(request *DescribeInstanceTopologyRequest) (_result *DescribeInstanceTopologyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstanceTopologyResponse{}
	_body, _err := client.DescribeInstanceTopologyWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInstancesWithOptions(request *DescribeInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInstances"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInstances(request *DescribeInstancesRequest) (_result *DescribeInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInstancesResponse{}
	_body, _err := client.DescribeInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeNodeMetricsWithOptions(request *DescribeNodeMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdList)) {
		body["NodeIdList"] = request.NodeIdList
	}

	if !tea.BoolValue(util.IsUnset(request.NodeName)) {
		body["NodeName"] = request.NodeName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNodeMetrics"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeNodeMetrics(request *DescribeNodeMetricsRequest) (_result *DescribeNodeMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeMetricsResponse{}
	_body, _err := client.DescribeNodeMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOmsOpenAPIProjectWithOptions(request *DescribeOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *DescribeOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOmsOpenAPIProject(request *DescribeOmsOpenAPIProjectRequest) (_result *DescribeOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOmsOpenAPIProjectResponse{}
	_body, _err := client.DescribeOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOmsOpenAPIProjectStepsWithOptions(request *DescribeOmsOpenAPIProjectStepsRequest, runtime *util.RuntimeOptions) (_result *DescribeOmsOpenAPIProjectStepsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOmsOpenAPIProjectSteps"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOmsOpenAPIProjectStepsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOmsOpenAPIProjectSteps(request *DescribeOmsOpenAPIProjectStepsRequest) (_result *DescribeOmsOpenAPIProjectStepsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOmsOpenAPIProjectStepsResponse{}
	_body, _err := client.DescribeOmsOpenAPIProjectStepsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeOutlineBindingWithOptions(request *DescribeOutlineBindingRequest, runtime *util.RuntimeOptions) (_result *DescribeOutlineBindingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.IsConcurrentLimit)) {
		body["IsConcurrentLimit"] = request.IsConcurrentLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TableName)) {
		body["TableName"] = request.TableName
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeOutlineBinding"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeOutlineBindingResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeOutlineBinding(request *DescribeOutlineBindingRequest) (_result *DescribeOutlineBindingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeOutlineBindingResponse{}
	_body, _err := client.DescribeOutlineBindingWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParameters"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeParametersHistoryWithOptions(request *DescribeParametersHistoryRequest, runtime *util.RuntimeOptions) (_result *DescribeParametersHistoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeParametersHistory"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeParametersHistoryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeParametersHistory(request *DescribeParametersHistoryRequest) (_result *DescribeParametersHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeParametersHistoryResponse{}
	_body, _err := client.DescribeParametersHistoryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRecommendIndexWithOptions(request *DescribeRecommendIndexRequest, runtime *util.RuntimeOptions) (_result *DescribeRecommendIndexResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRecommendIndex"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRecommendIndexResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRecommendIndex(request *DescribeRecommendIndexRequest) (_result *DescribeRecommendIndexResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRecommendIndexResponse{}
	_body, _err := client.DescribeRecommendIndexWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLDetailsWithOptions(request *DescribeSQLDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLDetails"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLDetails(request *DescribeSQLDetailsRequest) (_result *DescribeSQLDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLDetailsResponse{}
	_body, _err := client.DescribeSQLDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLHistoryListWithOptions(request *DescribeSQLHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLHistoryList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLHistoryList(request *DescribeSQLHistoryListRequest) (_result *DescribeSQLHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLHistoryListResponse{}
	_body, _err := client.DescribeSQLHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSQLPlansWithOptions(request *DescribeSQLPlansRequest, runtime *util.RuntimeOptions) (_result *DescribeSQLPlansResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSQLPlans"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSQLPlansResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSQLPlans(request *DescribeSQLPlansRequest) (_result *DescribeSQLPlansResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSQLPlansResponse{}
	_body, _err := client.DescribeSQLPlansWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSecurityIpGroupsWithOptions(request *DescribeSecurityIpGroupsRequest, runtime *util.RuntimeOptions) (_result *DescribeSecurityIpGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSecurityIpGroups"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSecurityIpGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSecurityIpGroups(request *DescribeSecurityIpGroupsRequest) (_result *DescribeSecurityIpGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSecurityIpGroupsResponse{}
	_body, _err := client.DescribeSecurityIpGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLHistoryListWithOptions(request *DescribeSlowSQLHistoryListRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLHistoryListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowSQLHistoryList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowSQLHistoryListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLHistoryList(request *DescribeSlowSQLHistoryListRequest) (_result *DescribeSlowSQLHistoryListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLHistoryListResponse{}
	_body, _err := client.DescribeSlowSQLHistoryListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSlowSQLListWithOptions(tmpReq *DescribeSlowSQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeSlowSQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSlowSQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSlowSQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSlowSQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSlowSQLList(request *DescribeSlowSQLListRequest) (_result *DescribeSlowSQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSlowSQLListResponse{}
	_body, _err := client.DescribeSlowSQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantWithOptions(request *DescribeTenantRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenant"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenant(request *DescribeTenantRequest) (_result *DescribeTenantResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantResponse{}
	_body, _err := client.DescribeTenantWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantMetricsWithOptions(request *DescribeTenantMetricsRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Metrics)) {
		body["Metrics"] = request.Metrics
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantIdList)) {
		body["TenantIdList"] = request.TenantIdList
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantMetrics"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantMetrics(request *DescribeTenantMetricsRequest) (_result *DescribeTenantMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantMetricsResponse{}
	_body, _err := client.DescribeTenantMetricsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantUserRolesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTenantUserRolesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantUserRoles() (_result *DescribeTenantUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantUserRolesResponse{}
	_body, _err := client.DescribeTenantUserRolesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantUsersWithOptions(request *DescribeTenantUsersRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantUsersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantUsers"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantUsersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantUsers(request *DescribeTenantUsersRequest) (_result *DescribeTenantUsersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantUsersResponse{}
	_body, _err := client.DescribeTenantUsersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantZonesReadWithOptions(request *DescribeTenantZonesReadRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantZonesReadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenantZonesRead"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantZonesReadResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenantZonesRead(request *DescribeTenantZonesReadRequest) (_result *DescribeTenantZonesReadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantZonesReadResponse{}
	_body, _err := client.DescribeTenantZonesReadWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTenantsWithOptions(request *DescribeTenantsRequest, runtime *util.RuntimeOptions) (_result *DescribeTenantsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantName)) {
		body["TenantName"] = request.TenantName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTenants"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTenantsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTenants(request *DescribeTenantsRequest) (_result *DescribeTenantsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTenantsResponse{}
	_body, _err := client.DescribeTenantsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTimeZonesWithOptions(runtime *util.RuntimeOptions) (_result *DescribeTimeZonesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("DescribeTimeZones"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTimeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTimeZones() (_result *DescribeTimeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTimeZonesResponse{}
	_body, _err := client.DescribeTimeZonesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTopSQLListWithOptions(tmpReq *DescribeTopSQLListRequest, runtime *util.RuntimeOptions) (_result *DescribeTopSQLListResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeTopSQLListShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.FilterCondition)) {
		request.FilterConditionShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.FilterCondition, tea.String("FilterCondition"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DbName)) {
		body["DbName"] = request.DbName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FilterConditionShrink)) {
		body["FilterCondition"] = request.FilterConditionShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIp)) {
		body["NodeIp"] = request.NodeIp
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SQLId)) {
		body["SQLId"] = request.SQLId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKeyWord)) {
		body["SearchKeyWord"] = request.SearchKeyWord
	}

	if !tea.BoolValue(util.IsUnset(request.SearchParameter)) {
		body["SearchParameter"] = request.SearchParameter
	}

	if !tea.BoolValue(util.IsUnset(request.SearchRule)) {
		body["SearchRule"] = request.SearchRule
	}

	if !tea.BoolValue(util.IsUnset(request.SearchValue)) {
		body["SearchValue"] = request.SearchValue
	}

	if !tea.BoolValue(util.IsUnset(request.SortColumn)) {
		body["SortColumn"] = request.SortColumn
	}

	if !tea.BoolValue(util.IsUnset(request.SortOrder)) {
		body["SortOrder"] = request.SortOrder
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTopSQLList"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTopSQLListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTopSQLList(request *DescribeTopSQLListRequest) (_result *DescribeTopSQLListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTopSQLListResponse{}
	_body, _err := client.DescribeTopSQLListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeployType)) {
		body["DeployType"] = request.DeployType
	}

	if !tea.BoolValue(util.IsUnset(request.Series)) {
		body["Series"] = request.Series
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseDescriptionWithOptions(request *ModifyDatabaseDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseDescription"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseDescription(request *ModifyDatabaseDescriptionRequest) (_result *ModifyDatabaseDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseDescriptionResponse{}
	_body, _err := client.ModifyDatabaseDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDatabaseUserRolesWithOptions(request *ModifyDatabaseUserRolesRequest, runtime *util.RuntimeOptions) (_result *ModifyDatabaseUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DatabaseName)) {
		body["DatabaseName"] = request.DatabaseName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.Users)) {
		body["Users"] = request.Users
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyDatabaseUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyDatabaseUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDatabaseUserRoles(request *ModifyDatabaseUserRolesRequest) (_result *ModifyDatabaseUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDatabaseUserRolesResponse{}
	_body, _err := client.ModifyDatabaseUserRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyInstanceNameWithOptions(request *ModifyInstanceNameRequest, runtime *util.RuntimeOptions) (_result *ModifyInstanceNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		body["InstanceName"] = request.InstanceName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyInstanceName"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyInstanceName(request *ModifyInstanceNameRequest) (_result *ModifyInstanceNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyInstanceNameResponse{}
	_body, _err := client.ModifyInstanceNameWithOptions(request, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Dimension)) {
		body["Dimension"] = request.Dimension
	}

	if !tea.BoolValue(util.IsUnset(request.DimensionValue)) {
		body["DimensionValue"] = request.DimensionValue
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Parameters)) {
		body["Parameters"] = request.Parameters
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyParameters"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyParametersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifySecurityIpsWithOptions(request *ModifySecurityIpsRequest, runtime *util.RuntimeOptions) (_result *ModifySecurityIpsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIpGroupName)) {
		body["SecurityIpGroupName"] = request.SecurityIpGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityIps)) {
		body["SecurityIps"] = request.SecurityIps
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySecurityIps"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySecurityIpsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ModifyTenantPrimaryZoneWithOptions(request *ModifyTenantPrimaryZoneRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantPrimaryZoneResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MasterIntranetAddressZone)) {
		body["MasterIntranetAddressZone"] = request.MasterIntranetAddressZone
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZone)) {
		body["PrimaryZone"] = request.PrimaryZone
	}

	if !tea.BoolValue(util.IsUnset(request.PrimaryZoneDeployType)) {
		body["PrimaryZoneDeployType"] = request.PrimaryZoneDeployType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserVSwitchId)) {
		body["UserVSwitchId"] = request.UserVSwitchId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantPrimaryZone"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantPrimaryZoneResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantPrimaryZone(request *ModifyTenantPrimaryZoneRequest) (_result *ModifyTenantPrimaryZoneResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantPrimaryZoneResponse{}
	_body, _err := client.ModifyTenantPrimaryZoneWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantResourceWithOptions(request *ModifyTenantResourceRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cpu)) {
		body["Cpu"] = request.Cpu
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.Memory)) {
		body["Memory"] = request.Memory
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantResource"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantResource(request *ModifyTenantResourceRequest) (_result *ModifyTenantResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantResourceResponse{}
	_body, _err := client.ModifyTenantResourceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserDescriptionWithOptions(request *ModifyTenantUserDescriptionRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserDescription"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserDescriptionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserDescription(request *ModifyTenantUserDescriptionRequest) (_result *ModifyTenantUserDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserDescriptionResponse{}
	_body, _err := client.ModifyTenantUserDescriptionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserPasswordWithOptions(request *ModifyTenantUserPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserPassword)) {
		body["UserPassword"] = request.UserPassword
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserPassword"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserPasswordResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserPassword(request *ModifyTenantUserPasswordRequest) (_result *ModifyTenantUserPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserPasswordResponse{}
	_body, _err := client.ModifyTenantUserPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserRolesWithOptions(request *ModifyTenantUserRolesRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserRolesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.ModifyType)) {
		body["ModifyType"] = request.ModifyType
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserRole)) {
		body["UserRole"] = request.UserRole
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserRoles"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserRolesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserRoles(request *ModifyTenantUserRolesRequest) (_result *ModifyTenantUserRolesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserRolesResponse{}
	_body, _err := client.ModifyTenantUserRolesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyTenantUserStatusWithOptions(request *ModifyTenantUserStatusRequest, runtime *util.RuntimeOptions) (_result *ModifyTenantUserStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.TenantId)) {
		body["TenantId"] = request.TenantId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["UserName"] = request.UserName
	}

	if !tea.BoolValue(util.IsUnset(request.UserStatus)) {
		body["UserStatus"] = request.UserStatus
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyTenantUserStatus"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyTenantUserStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyTenantUserStatus(request *ModifyTenantUserStatusRequest) (_result *ModifyTenantUserStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyTenantUserStatusResponse{}
	_body, _err := client.ModifyTenantUserStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseOmsOpenAPIProjectWithOptions(request *ReleaseOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *ReleaseOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseOmsOpenAPIProject(request *ReleaseOmsOpenAPIProjectRequest) (_result *ReleaseOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReleaseOmsOpenAPIProjectResponse{}
	_body, _err := client.ReleaseOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetOmsOpenAPIProjectWithOptions(request *ResetOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *ResetOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetOmsOpenAPIProject(request *ResetOmsOpenAPIProjectRequest) (_result *ResetOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetOmsOpenAPIProjectResponse{}
	_body, _err := client.ResetOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResumeOmsOpenAPIProjectWithOptions(request *ResumeOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *ResumeOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ResumeOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResumeOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResumeOmsOpenAPIProject(request *ResumeOmsOpenAPIProjectRequest) (_result *ResumeOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResumeOmsOpenAPIProjectResponse{}
	_body, _err := client.ResumeOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOmsOpenAPIMonitorMetricWithOptions(request *SearchOmsOpenAPIMonitorMetricRequest, runtime *util.RuntimeOptions) (_result *SearchOmsOpenAPIMonitorMetricResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		body["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		body["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxPointNum)) {
		body["MaxPointNum"] = request.MaxPointNum
	}

	if !tea.BoolValue(util.IsUnset(request.Metric)) {
		body["Metric"] = request.Metric
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchOmsOpenAPIMonitorMetric"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchOmsOpenAPIMonitorMetricResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOmsOpenAPIMonitorMetric(request *SearchOmsOpenAPIMonitorMetricRequest) (_result *SearchOmsOpenAPIMonitorMetricResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchOmsOpenAPIMonitorMetricResponse{}
	_body, _err := client.SearchOmsOpenAPIMonitorMetricWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchOmsOpenAPIProjectsWithOptions(tmpReq *SearchOmsOpenAPIProjectsRequest, runtime *util.RuntimeOptions) (_result *SearchOmsOpenAPIProjectsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SearchOmsOpenAPIProjectsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DestDbTypes)) {
		request.DestDbTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DestDbTypes, tea.String("DestDbTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.LabelIds)) {
		request.LabelIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.LabelIds, tea.String("LabelIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.SourceDbTypes)) {
		request.SourceDbTypesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SourceDbTypes, tea.String("SourceDbTypes"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.StatusList)) {
		request.StatusListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.StatusList, tea.String("StatusList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestDbTypesShrink)) {
		body["DestDbTypes"] = request.DestDbTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.LabelIdsShrink)) {
		body["LabelIds"] = request.LabelIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		body["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceDbTypesShrink)) {
		body["SourceDbTypes"] = request.SourceDbTypesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.StatusListShrink)) {
		body["StatusList"] = request.StatusListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchOmsOpenAPIProjects"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchOmsOpenAPIProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchOmsOpenAPIProjects(request *SearchOmsOpenAPIProjectsRequest) (_result *SearchOmsOpenAPIProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchOmsOpenAPIProjectsResponse{}
	_body, _err := client.SearchOmsOpenAPIProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartOmsOpenAPIProjectWithOptions(request *StartOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *StartOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StartOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartOmsOpenAPIProject(request *StartOmsOpenAPIProjectRequest) (_result *StartOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartOmsOpenAPIProjectResponse{}
	_body, _err := client.StartOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopOmsOpenAPIProjectWithOptions(request *StopOmsOpenAPIProjectRequest, runtime *util.RuntimeOptions) (_result *StopOmsOpenAPIProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		body["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		body["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		body["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkerGradeId)) {
		body["WorkerGradeId"] = request.WorkerGradeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopOmsOpenAPIProject"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopOmsOpenAPIProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopOmsOpenAPIProject(request *StopOmsOpenAPIProjectRequest) (_result *StopOmsOpenAPIProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopOmsOpenAPIProjectResponse{}
	_body, _err := client.StopOmsOpenAPIProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
