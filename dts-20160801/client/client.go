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

type ConfigurationSynchronizationJobRequest struct {
	DestinationEndpoint    *ConfigurationSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	Initialization         *ConfigurationSynchronizationJobRequestInitialization      `json:"Initialization,omitempty" xml:"Initialization,omitempty" type:"Struct"`
	SourceEndpoint         *ConfigurationSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	OwnerId                *string                                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId   *string                                                    `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName *string                                                    `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationObject  *string                                                    `json:"SynchronizationObject,omitempty" xml:"SynchronizationObject,omitempty"`
}

func (s ConfigurationSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobRequest) SetDestinationEndpoint(v *ConfigurationSynchronizationJobRequestDestinationEndpoint) *ConfigurationSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetInitialization(v *ConfigurationSynchronizationJobRequestInitialization) *ConfigurationSynchronizationJobRequest {
	s.Initialization = v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetSourceEndpoint(v *ConfigurationSynchronizationJobRequestSourceEndpoint) *ConfigurationSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetOwnerId(v string) *ConfigurationSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetSynchronizationJobId(v string) *ConfigurationSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetSynchronizationJobName(v string) *ConfigurationSynchronizationJobRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequest) SetSynchronizationObject(v string) *ConfigurationSynchronizationJobRequest {
	s.SynchronizationObject = &v
	return s
}

type ConfigurationSynchronizationJobRequestDestinationEndpoint struct {
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ConfigurationSynchronizationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobRequestDestinationEndpoint) SetInstanceID(v string) *ConfigurationSynchronizationJobRequestDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigurationSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

type ConfigurationSynchronizationJobRequestInitialization struct {
	DataLoad      *bool `json:"DataLoad,omitempty" xml:"DataLoad,omitempty"`
	StructureLoad *bool `json:"StructureLoad,omitempty" xml:"StructureLoad,omitempty"`
}

func (s ConfigurationSynchronizationJobRequestInitialization) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobRequestInitialization) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobRequestInitialization) SetDataLoad(v bool) *ConfigurationSynchronizationJobRequestInitialization {
	s.DataLoad = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequestInitialization) SetStructureLoad(v bool) *ConfigurationSynchronizationJobRequestInitialization {
	s.StructureLoad = &v
	return s
}

type ConfigurationSynchronizationJobRequestSourceEndpoint struct {
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ConfigurationSynchronizationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobRequestSourceEndpoint) SetInstanceID(v string) *ConfigurationSynchronizationJobRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigurationSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigurationSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

type ConfigurationSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigurationSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobResponseBody) SetErrCode(v string) *ConfigurationSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigurationSynchronizationJobResponseBody) SetErrMessage(v string) *ConfigurationSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigurationSynchronizationJobResponseBody) SetRequestId(v string) *ConfigurationSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigurationSynchronizationJobResponseBody) SetSuccess(v string) *ConfigurationSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type ConfigurationSynchronizationJobResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigurationSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigurationSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigurationSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigurationSynchronizationJobResponse) SetHeaders(v map[string]*string) *ConfigurationSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigurationSynchronizationJobResponse) SetStatusCode(v int32) *ConfigurationSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigurationSynchronizationJobResponse) SetBody(v *ConfigurationSynchronizationJobResponseBody) *ConfigurationSynchronizationJobResponse {
	s.Body = v
	return s
}

type ConfigureMigrationJobRequest struct {
	DestinationEndpoint *ConfigureMigrationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode       *ConfigureMigrationJobRequestMigrationMode       `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	SourceEndpoint      *ConfigureMigrationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Checkpoint          *string                                          `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	MigrationJobId      *string                                          `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	MigrationJobName    *string                                          `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationObject     *string                                          `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	MigrationReserved   *string                                          `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId             *string                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ConfigureMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequest) SetDestinationEndpoint(v *ConfigureMigrationJobRequestDestinationEndpoint) *ConfigureMigrationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationMode(v *ConfigureMigrationJobRequestMigrationMode) *ConfigureMigrationJobRequest {
	s.MigrationMode = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetSourceEndpoint(v *ConfigureMigrationJobRequestSourceEndpoint) *ConfigureMigrationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetCheckpoint(v string) *ConfigureMigrationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationJobId(v string) *ConfigureMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationJobName(v string) *ConfigureMigrationJobRequest {
	s.MigrationJobName = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationObject(v string) *ConfigureMigrationJobRequest {
	s.MigrationObject = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationReserved(v string) *ConfigureMigrationJobRequest {
	s.MigrationReserved = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetOwnerId(v string) *ConfigureMigrationJobRequest {
	s.OwnerId = &v
	return s
}

type ConfigureMigrationJobRequestDestinationEndpoint struct {
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetDataBaseName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.DataBaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPort(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.UserName = &v
	return s
}

type ConfigureMigrationJobRequestMigrationMode struct {
	DataIntialization      *bool `json:"DataIntialization,omitempty" xml:"DataIntialization,omitempty"`
	DataSynchronization    *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureIntialization *bool `json:"StructureIntialization,omitempty" xml:"StructureIntialization,omitempty"`
}

func (s ConfigureMigrationJobRequestMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataIntialization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataSynchronization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetStructureIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.StructureIntialization = &v
	return s
}

type ConfigureMigrationJobRequestSourceEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureMigrationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetIP(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOracleSID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPort(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRole(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

type ConfigureMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobResponseBody) SetErrCode(v string) *ConfigureMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetErrMessage(v string) *ConfigureMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetRequestId(v string) *ConfigureMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetSuccess(v string) *ConfigureMigrationJobResponseBody {
	s.Success = &v
	return s
}

type ConfigureMigrationJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigureMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobResponse) SetHeaders(v map[string]*string) *ConfigureMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureMigrationJobResponse) SetStatusCode(v int32) *ConfigureMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureMigrationJobResponse) SetBody(v *ConfigureMigrationJobResponseBody) *ConfigureMigrationJobResponse {
	s.Body = v
	return s
}

type ConfigureSubscriptionInstanceRequest struct {
	SourceEndpoint           *ConfigureSubscriptionInstanceRequestSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType     *ConfigureSubscriptionInstanceRequestSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	OwnerId                  *string                                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId   *string                                                   `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	SubscriptionInstanceName *string                                                   `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionObject       *string                                                   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequest) SetSourceEndpoint(v *ConfigureSubscriptionInstanceRequestSourceEndpoint) *ConfigureSubscriptionInstanceRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionDataType(v *ConfigureSubscriptionInstanceRequestSubscriptionDataType) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionDataType = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetOwnerId(v string) *ConfigureSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceName(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionObject(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionObject = &v
	return s
}

type ConfigureSubscriptionInstanceRequestSourceEndpoint struct {
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetOwnerID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetRole(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Role = &v
	return s
}

type ConfigureSubscriptionInstanceRequestSubscriptionDataType struct {
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) SetDDL(v bool) *ConfigureSubscriptionInstanceRequestSubscriptionDataType {
	s.DDL = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionDataType) SetDML(v bool) *ConfigureSubscriptionInstanceRequestSubscriptionDataType {
	s.DML = &v
	return s
}

type ConfigureSubscriptionInstanceResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetRequestId(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

type ConfigureSubscriptionInstanceResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigureSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSubscriptionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionInstanceResponse) SetStatusCode(v int32) *ConfigureSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponse) SetBody(v *ConfigureSubscriptionInstanceResponseBody) *ConfigureSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type ConfigureSynchronizationJobRequest struct {
	DestinationEndpoint     *ConfigureSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	PartitionKey            *ConfigureSynchronizationJobRequestPartitionKey        `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty" type:"Struct"`
	SourceEndpoint          *ConfigureSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Checkpoint              *string                                                `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DataInitialization      *bool                                                  `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	MigrationReserved       *string                                                `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId                 *string                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StructureInitialization *bool                                                  `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	SynchronizationJobId    *string                                                `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName  *string                                                `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationObjects  *string                                                `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
}

func (s ConfigureSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequest) SetDestinationEndpoint(v *ConfigureSynchronizationJobRequestDestinationEndpoint) *ConfigureSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetPartitionKey(v *ConfigureSynchronizationJobRequestPartitionKey) *ConfigureSynchronizationJobRequest {
	s.PartitionKey = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSourceEndpoint(v *ConfigureSynchronizationJobRequestSourceEndpoint) *ConfigureSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetCheckpoint(v string) *ConfigureSynchronizationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetDataInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetMigrationReserved(v string) *ConfigureSynchronizationJobRequest {
	s.MigrationReserved = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetOwnerId(v string) *ConfigureSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetStructureInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobName(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationObjects(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationObjects = &v
	return s
}

type ConfigureSynchronizationJobRequestDestinationEndpoint struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetPort(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetUserName(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.UserName = &v
	return s
}

type ConfigureSynchronizationJobRequestPartitionKey struct {
	ModifyTimeDay    *bool `json:"ModifyTime_Day,omitempty" xml:"ModifyTime_Day,omitempty"`
	ModifyTimeHour   *bool `json:"ModifyTime_Hour,omitempty" xml:"ModifyTime_Hour,omitempty"`
	ModifyTimeMinute *bool `json:"ModifyTime_Minute,omitempty" xml:"ModifyTime_Minute,omitempty"`
	ModifyTimeMonth  *bool `json:"ModifyTime_Month,omitempty" xml:"ModifyTime_Month,omitempty"`
	ModifyTimeYear   *bool `json:"ModifyTime_Year,omitempty" xml:"ModifyTime_Year,omitempty"`
}

func (s ConfigureSynchronizationJobRequestPartitionKey) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestPartitionKey) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeDay(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeDay = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeHour(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeHour = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeMinute(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeMinute = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeMonth(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeMonth = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeYear(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeYear = &v
	return s
}

type ConfigureSynchronizationJobRequestSourceEndpoint struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPort(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetRole(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

type ConfigureSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type ConfigureSynchronizationJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigureSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobResponse) SetStatusCode(v int32) *ConfigureSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSynchronizationJobResponse) SetBody(v *ConfigureSynchronizationJobResponseBody) *ConfigureSynchronizationJobResponse {
	s.Body = v
	return s
}

type CreateMigrationJobRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s CreateMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobRequest) SetClientToken(v string) *CreateMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMigrationJobRequest) SetMigrationJobClass(v string) *CreateMigrationJobRequest {
	s.MigrationJobClass = &v
	return s
}

func (s *CreateMigrationJobRequest) SetOwnerId(v string) *CreateMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateMigrationJobRequest) SetRegion(v string) *CreateMigrationJobRequest {
	s.Region = &v
	return s
}

type CreateMigrationJobResponseBody struct {
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponseBody) SetErrCode(v string) *CreateMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetErrMessage(v string) *CreateMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetMigrationJobId(v string) *CreateMigrationJobResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetRequestId(v string) *CreateMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetSuccess(v string) *CreateMigrationJobResponseBody {
	s.Success = &v
	return s
}

type CreateMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponse) SetHeaders(v map[string]*string) *CreateMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateMigrationJobResponse) SetStatusCode(v int32) *CreateMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMigrationJobResponse) SetBody(v *CreateMigrationJobResponseBody) *CreateMigrationJobResponse {
	s.Body = v
	return s
}

type CreateSubscriptionInstanceRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType     *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period      *string `json:"Period,omitempty" xml:"Period,omitempty"`
	Region      *string `json:"Region,omitempty" xml:"Region,omitempty"`
	UsedTime    *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s CreateSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceRequest) SetClientToken(v string) *CreateSubscriptionInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetOwnerId(v string) *CreateSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetPayType(v string) *CreateSubscriptionInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetPeriod(v string) *CreateSubscriptionInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetRegion(v string) *CreateSubscriptionInstanceRequest {
	s.Region = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetUsedTime(v int32) *CreateSubscriptionInstanceRequest {
	s.UsedTime = &v
	return s
}

type CreateSubscriptionInstanceResponseBody struct {
	ErrCode                *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage             *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	Success                *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrCode(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrMessage(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetRequestId(v string) *CreateSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSubscriptionInstanceId(v string) *CreateSubscriptionInstanceResponseBody {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSuccess(v string) *CreateSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

type CreateSubscriptionInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSubscriptionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *CreateSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateSubscriptionInstanceResponse) SetStatusCode(v int32) *CreateSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSubscriptionInstanceResponse) SetBody(v *CreateSubscriptionInstanceResponseBody) *CreateSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type CreateSynchronizationJobRequest struct {
	DestinationEndpoint     *CreateSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	SourceEndpoint          *CreateSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	ClientToken             *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DestRegion              *string                                             `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	OwnerId                 *string                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PayType                 *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                  *string                                             `json:"Period,omitempty" xml:"Period,omitempty"`
	SourceRegion            *string                                             `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	SynchronizationJobClass *string                                             `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	UsedTime                *int32                                              `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	NetworkType             *string                                             `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s CreateSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequest) SetDestinationEndpoint(v *CreateSynchronizationJobRequestDestinationEndpoint) *CreateSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSourceEndpoint(v *CreateSynchronizationJobRequestSourceEndpoint) *CreateSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetClientToken(v string) *CreateSynchronizationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDestRegion(v string) *CreateSynchronizationJobRequest {
	s.DestRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetOwnerId(v string) *CreateSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetPayType(v string) *CreateSynchronizationJobRequest {
	s.PayType = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetPeriod(v string) *CreateSynchronizationJobRequest {
	s.Period = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSourceRegion(v string) *CreateSynchronizationJobRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSynchronizationJobClass(v string) *CreateSynchronizationJobRequest {
	s.SynchronizationJobClass = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetUsedTime(v int32) *CreateSynchronizationJobRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetNetworkType(v string) *CreateSynchronizationJobRequest {
	s.NetworkType = &v
	return s
}

type CreateSynchronizationJobRequestDestinationEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSynchronizationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *CreateSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

type CreateSynchronizationJobRequestSourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSynchronizationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *CreateSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

type CreateSynchronizationJobResponseBody struct {
	ErrCode              *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success              *string `json:"Success,omitempty" xml:"Success,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s CreateSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponseBody) SetErrCode(v string) *CreateSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetErrMessage(v string) *CreateSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetRequestId(v string) *CreateSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetSuccess(v string) *CreateSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetSynchronizationJobId(v string) *CreateSynchronizationJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

type CreateSynchronizationJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponse) SetHeaders(v map[string]*string) *CreateSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSynchronizationJobResponse) SetStatusCode(v int32) *CreateSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSynchronizationJobResponse) SetBody(v *CreateSynchronizationJobResponseBody) *CreateSynchronizationJobResponse {
	s.Body = v
	return s
}

type DeleteMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DeleteMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobRequest) SetMigrationJobId(v string) *DeleteMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetOwnerId(v string) *DeleteMigrationJobRequest {
	s.OwnerId = &v
	return s
}

type DeleteMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponseBody) SetErrCode(v string) *DeleteMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetErrMessage(v string) *DeleteMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetRequestId(v string) *DeleteMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetSuccess(v string) *DeleteMigrationJobResponseBody {
	s.Success = &v
	return s
}

type DeleteMigrationJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponse) SetHeaders(v map[string]*string) *DeleteMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteMigrationJobResponse) SetStatusCode(v int32) *DeleteMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMigrationJobResponse) SetBody(v *DeleteMigrationJobResponseBody) *DeleteMigrationJobResponse {
	s.Body = v
	return s
}

type DeleteSubscriptionInstanceRequest struct {
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DeleteSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceRequest) SetOwnerId(v string) *DeleteSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *DeleteSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

type DeleteSubscriptionInstanceResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrCode(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrMessage(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetRequestId(v string) *DeleteSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetSuccess(v string) *DeleteSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

type DeleteSubscriptionInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSubscriptionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *DeleteSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscriptionInstanceResponse) SetStatusCode(v int32) *DeleteSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponse) SetBody(v *DeleteSubscriptionInstanceResponseBody) *DeleteSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type DeleteSynchronizationJobRequest struct {
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DeleteSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobRequest) SetOwnerId(v string) *DeleteSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetSynchronizationJobId(v string) *DeleteSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

type DeleteSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobResponseBody) SetErrCode(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetErrMessage(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetRequestId(v string) *DeleteSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetSuccess(v string) *DeleteSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type DeleteSynchronizationJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobResponse) SetHeaders(v map[string]*string) *DeleteSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSynchronizationJobResponse) SetStatusCode(v int32) *DeleteSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSynchronizationJobResponse) SetBody(v *DeleteSynchronizationJobResponseBody) *DeleteSynchronizationJobResponse {
	s.Body = v
	return s
}

type DescirbeMigrationJobsRequest struct {
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescirbeMigrationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsRequest) SetMigrationJobName(v string) *DescirbeMigrationJobsRequest {
	s.MigrationJobName = &v
	return s
}

func (s *DescirbeMigrationJobsRequest) SetOwnerId(v string) *DescirbeMigrationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescirbeMigrationJobsRequest) SetPageNum(v int32) *DescirbeMigrationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescirbeMigrationJobsRequest) SetPageSize(v int32) *DescirbeMigrationJobsRequest {
	s.PageSize = &v
	return s
}

type DescirbeMigrationJobsResponseBody struct {
	MigrationJobs    *DescirbeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int64                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescirbeMigrationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBody) SetMigrationJobs(v *DescirbeMigrationJobsResponseBodyMigrationJobs) *DescirbeMigrationJobsResponseBody {
	s.MigrationJobs = v
	return s
}

func (s *DescirbeMigrationJobsResponseBody) SetPageNumber(v int32) *DescirbeMigrationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescirbeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescirbeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobs struct {
	MigrationJob []*DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob `json:"MigrationJob,omitempty" xml:"MigrationJob,omitempty" type:"Repeated"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobs) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobs) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobs) SetMigrationJob(v []*DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) *DescirbeMigrationJobsResponseBodyMigrationJobs {
	s.MigrationJob = v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob struct {
	DataInitialization      *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization      `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty" type:"Struct"`
	DataSynchronization     *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization     `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	DestinationEndpoint     *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint     `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationJobClass       *string                                                                            `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	MigrationJobID          *string                                                                            `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
	MigrationJobName        *string                                                                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus      *string                                                                            `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	MigrationMode           *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode           `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	MigrationObject         *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject         `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	PayType                 *string                                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Precheck                *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck                `json:"Precheck,omitempty" xml:"Precheck,omitempty" type:"Struct"`
	SourceEndpoint          *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint          `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitialization *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty" type:"Struct"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataInitialization(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataInitialization = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataSynchronization(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataSynchronization = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDestinationEndpoint(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobClass(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobClass = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobID(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobID = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPrecheck(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Precheck = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetSourceEndpoint(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetStructureInitialization(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.StructureInitialization = v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetPercent(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Percent = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetProgress(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Progress = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetStatus(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Status = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetDelay(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Delay = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetDatabaseName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetEngineName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetIP(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceID(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceType(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetOracleSID(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetPort(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetUserName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.UserName = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataInitialization(v bool) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataSynchronization(v bool) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetStructureInitialization(v bool) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject struct {
	SynchronousObject []*DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) SetSynchronousObject(v []*DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	s.SynchronousObject = v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject struct {
	DatabaseName  *string                                                                                              `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                              `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetDatabaseName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetTableList(v *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) SetTable(v []*string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck struct {
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetPercent(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Percent = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetStatus(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Status = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetDatabaseName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetEngineName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetIP(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceID(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceType(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetOracleSID(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetPort(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetUserName(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.UserName = &v
	return s
}

type DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetPercent(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Percent = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetProgress(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Progress = &v
	return s
}

func (s *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetStatus(v string) *DescirbeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Status = &v
	return s
}

type DescirbeMigrationJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescirbeMigrationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescirbeMigrationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescirbeMigrationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescirbeMigrationJobsResponse) SetHeaders(v map[string]*string) *DescirbeMigrationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescirbeMigrationJobsResponse) SetStatusCode(v int32) *DescirbeMigrationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescirbeMigrationJobsResponse) SetBody(v *DescirbeMigrationJobsResponseBody) *DescirbeMigrationJobsResponse {
	s.Body = v
	return s
}

type DescribeInitializationStatusRequest struct {
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum              *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeInitializationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusRequest) SetOwnerId(v string) *DescribeInitializationStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageNum(v int32) *DescribeInitializationStatusRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageSize(v int32) *DescribeInitializationStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetSynchronizationJobId(v string) *DescribeInitializationStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

type DescribeInitializationStatusResponseBody struct {
	DataInitializationDetails      []*DescribeInitializationStatusResponseBodyDataInitializationDetails      `json:"DataInitializationDetails,omitempty" xml:"DataInitializationDetails,omitempty" type:"Repeated"`
	DataSynchronizationDetails     []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails     `json:"DataSynchronizationDetails,omitempty" xml:"DataSynchronizationDetails,omitempty" type:"Repeated"`
	StructureInitializationDetails []*DescribeInitializationStatusResponseBodyStructureInitializationDetails `json:"StructureInitializationDetails,omitempty" xml:"StructureInitializationDetails,omitempty" type:"Repeated"`
}

func (s DescribeInitializationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBody) SetDataInitializationDetails(v []*DescribeInitializationStatusResponseBodyDataInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.DataInitializationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetDataSynchronizationDetails(v []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails) *DescribeInitializationStatusResponseBody {
	s.DataSynchronizationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetStructureInitializationDetails(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.StructureInitializationDetails = v
	return s
}

type DescribeInitializationStatusResponseBodyDataInitializationDetails struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishRowNum           *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TotalRowNum            *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
	UsedTime               *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetFinishRowNum(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.FinishRowNum = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TableName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTotalRowNum(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TotalRowNum = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetUsedTime(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.UsedTime = &v
	return s
}

type DescribeInitializationStatusResponseBodyDataSynchronizationDetails struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.TableName = &v
	return s
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetails struct {
	Constraints            []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Repeated"`
	DestinationOwnerDBName *string                                                                              `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string                                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string                                                                              `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string                                                                              `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string                                                                              `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string                                                                              `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetConstraints(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Constraints = v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Status = &v
	return s
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.Status = &v
	return s
}

type DescribeInitializationStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeInitializationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeInitializationStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponse) SetHeaders(v map[string]*string) *DescribeInitializationStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeInitializationStatusResponse) SetStatusCode(v int32) *DescribeInitializationStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInitializationStatusResponse) SetBody(v *DescribeInitializationStatusResponseBody) *DescribeInitializationStatusResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobDetailRequest struct {
	MigrationMode  *DescribeMigrationJobDetailRequestMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	ClientToken    *string                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobId *string                                         `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum        *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMigrationJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailRequest) SetMigrationMode(v *DescribeMigrationJobDetailRequestMigrationMode) *DescribeMigrationJobDetailRequest {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetClientToken(v string) *DescribeMigrationJobDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetMigrationJobId(v string) *DescribeMigrationJobDetailRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetOwnerId(v string) *DescribeMigrationJobDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageNum(v int32) *DescribeMigrationJobDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageSize(v int32) *DescribeMigrationJobDetailRequest {
	s.PageSize = &v
	return s
}

type DescribeMigrationJobDetailRequestMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeMigrationJobDetailRequestMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeMigrationJobDetailResponseBody struct {
	DataInitializationDetailList      *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList      `json:"DataInitializationDetailList,omitempty" xml:"DataInitializationDetailList,omitempty" type:"Struct"`
	DataSynchronizationDetailList     *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList     `json:"DataSynchronizationDetailList,omitempty" xml:"DataSynchronizationDetailList,omitempty" type:"Struct"`
	PageNumber                        *int32                                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount                   *int32                                                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	StructureInitializationDetailList *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList `json:"StructureInitializationDetailList,omitempty" xml:"StructureInitializationDetailList,omitempty" type:"Struct"`
	TotalRecordCount                  *int64                                                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataSynchronizationDetailList(v *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataSynchronizationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageNumber(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetStructureInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.StructureInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobDetailResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeMigrationJobDetailResponseBodyDataInitializationDetailList struct {
	DataInitializationDetail []*DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail `json:"DataInitializationDetail,omitempty" xml:"DataInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) SetDataInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList {
	s.DataInitializationDetail = v
	return s
}

type DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishRowNum           *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	MigrationTime          *string `json:"MigrationTime,omitempty" xml:"MigrationTime,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	TotalRowNum            *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetFinishRowNum(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.FinishRowNum = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetMigrationTime(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.MigrationTime = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetTableName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.TableName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetTotalRowNum(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.TotalRowNum = &v
	return s
}

type DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList struct {
	DataSynchronizationDetail []*DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail `json:"DataSynchronizationDetail,omitempty" xml:"DataSynchronizationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) SetDataSynchronizationDetail(v []*DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList {
	s.DataSynchronizationDetail = v
	return s
}

type DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetTableName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.TableName = &v
	return s
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList struct {
	StructureInitializationDetail []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail `json:"StructureInitializationDetail,omitempty" xml:"StructureInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) SetStructureInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList {
	s.StructureInitializationDetail = v
	return s
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail struct {
	ConstraintList         *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList `json:"ConstraintList,omitempty" xml:"ConstraintList,omitempty" type:"Struct"`
	DestinationOwnerDBName *string                                                                                                             `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string                                                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string                                                                                                             `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string                                                                                                             `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string                                                                                                             `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string                                                                                                             `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string                                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetConstraintList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ConstraintList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.Status = &v
	return s
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList struct {
	StructureInitializationDetail []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail `json:"StructureInitializationDetail,omitempty" xml:"StructureInitializationDetail,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) SetStructureInitializationDetail(v []*DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList {
	s.StructureInitializationDetail = v
	return s
}

type DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail struct {
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ObjectDefinition       *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectName             *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ObjectType             *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.Status = &v
	return s
}

type DescribeMigrationJobDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMigrationJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobDetailResponse) SetStatusCode(v int32) *DescribeMigrationJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobDetailResponse) SetBody(v *DescribeMigrationJobDetailResponseBody) *DescribeMigrationJobDetailResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobStatusRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeMigrationJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusRequest) SetClientToken(v string) *DescribeMigrationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetMigrationJobId(v string) *DescribeMigrationJobStatusRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetOwnerId(v string) *DescribeMigrationJobStatusRequest {
	s.OwnerId = &v
	return s
}

type DescribeMigrationJobStatusResponseBody struct {
	DataInitializationStatus      *DescribeMigrationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeMigrationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationJobClass             *string                                                              `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	MigrationJobId                *string                                                              `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	MigrationJobName              *string                                                              `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus            *string                                                              `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	MigrationMode                 *DescribeMigrationJobStatusResponseBodyMigrationMode                 `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	MigrationObject               *string                                                              `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	PayType                       *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	PrecheckStatus                *DescribeMigrationJobStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeMigrationJobStatusResponseBodySourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBody) SetDataInitializationStatus(v *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) *DescribeMigrationJobStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) *DescribeMigrationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) *DescribeMigrationJobStatusResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobClass(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobName(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobStatus(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationMode(v *DescribeMigrationJobStatusResponseBodyMigrationMode) *DescribeMigrationJobStatusResponseBody {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationObject(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationObject = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPayType(v string) *DescribeMigrationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPrecheckStatus(v *DescribeMigrationJobStatusResponseBodyPrecheckStatus) *DescribeMigrationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetSourceEndpoint(v *DescribeMigrationJobStatusResponseBodySourceEndpoint) *DescribeMigrationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) *DescribeMigrationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

type DescribeMigrationJobStatusResponseBodyDataInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus struct {
	Checkpoint   *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyDestinationEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	OracleSID    *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetPort(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.OracleSID = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyMigrationMode struct {
	DataInitialization      *bool `json:"dataInitialization,omitempty" xml:"dataInitialization,omitempty"`
	DataSynchronization     *bool `json:"dataSynchronization,omitempty" xml:"dataSynchronization,omitempty"`
	StructureInitialization *bool `json:"structureInitialization,omitempty" xml:"structureInitialization,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatus struct {
	Detail  *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	Percent *string                                                     `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetDetail(v *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail struct {
	CheckItem []*DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem `json:"CheckItem,omitempty" xml:"CheckItem,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) SetCheckItem(v []*DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail {
	s.CheckItem = v
	return s
}

type DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ItemName     *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetCheckStatus(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.CheckStatus = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetItemName(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.ItemName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem) SetRepairMethod(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetailCheckItem {
	s.RepairMethod = &v
	return s
}

type DescribeMigrationJobStatusResponseBodySourceEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	OracleSID    *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetPort(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetUserName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.OracleSID = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyStructureInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

type DescribeMigrationJobStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMigrationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobStatusResponse) SetStatusCode(v int32) *DescribeMigrationJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobStatusResponse) SetBody(v *DescribeMigrationJobStatusResponseBody) *DescribeMigrationJobStatusResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobsRequest struct {
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum          *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize         *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeMigrationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsRequest) SetMigrationJobName(v string) *DescribeMigrationJobsRequest {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetOwnerId(v string) *DescribeMigrationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageNum(v int32) *DescribeMigrationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageSize(v int32) *DescribeMigrationJobsRequest {
	s.PageSize = &v
	return s
}

type DescribeMigrationJobsResponseBody struct {
	MigrationJobs    *DescribeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int64                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeMigrationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBody) SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody {
	s.MigrationJobs = v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageNumber(v int32) *DescribeMigrationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobs struct {
	MigrationJob []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob `json:"MigrationJob,omitempty" xml:"MigrationJob,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobs) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobs) SetMigrationJob(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) *DescribeMigrationJobsResponseBodyMigrationJobs {
	s.MigrationJob = v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob struct {
	DataInitialization      *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization      `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty" type:"Struct"`
	DataSynchronization     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization     `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	DestinationEndpoint     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint     `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationJobClass       *string                                                                            `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	MigrationJobID          *string                                                                            `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
	MigrationJobName        *string                                                                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus      *string                                                                            `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	MigrationMode           *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode           `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	MigrationObject         *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject         `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	PayType                 *string                                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Precheck                *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck                `json:"Precheck,omitempty" xml:"Precheck,omitempty" type:"Struct"`
	SourceEndpoint          *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint          `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitialization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDataSynchronization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DataSynchronization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetDestinationEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobClass(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPrecheck(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Precheck = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetSourceEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetStructureInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.StructureInitialization = v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Status = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetDelay(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Delay = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.UserName = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject struct {
	SynchronousObject []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) SetSynchronousObject(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject {
	s.SynchronousObject = v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject struct {
	DatabaseName  *string                                                                                              `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                              `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetTableList(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck struct {
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Status = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.UserName = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Progress = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Status = &v
	return s
}

type DescribeMigrationJobsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeMigrationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobsResponse) SetStatusCode(v int32) *DescribeMigrationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMigrationJobsResponse) SetBody(v *DescribeMigrationJobsResponseBody) *DescribeMigrationJobsResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInstanceStatusRequest struct {
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetOwnerId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.SubscriptionInstanceId = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBody struct {
	BeginTimestamp           *string                                                             `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	ConsumptionCheckpoint    *string                                                             `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	ConsumptionClient        *string                                                             `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	EndTimestamp             *string                                                             `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	ErrorMessage             *string                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PayType                  *string                                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	SourceEndpoint           *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Status                   *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionDataType     *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstanceID   *string                                                             `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	SubscriptionInstanceName *string                                                             `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionObject       *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetBeginTimestamp(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetConsumptionClient(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetEndTimestamp(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrorMessage(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetPayType(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSourceEndpoint(v *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetStatus(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionDataType(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionObject(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionObject = v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint struct {
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) SetInstanceID(v string) *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType struct {
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) SetDDL(v bool) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType {
	s.DDL = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) SetDML(v bool) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType {
	s.DML = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject struct {
	SynchronousObject []*DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) SetSynchronousObject(v []*DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject {
	s.SynchronousObject = v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject struct {
	DatabaseName  *string                                                                                     `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                     `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) SetDatabaseName(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) SetTableList(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

type DescribeSubscriptionInstanceStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSubscriptionInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionInstanceStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetStatusCode(v int32) *DescribeSubscriptionInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponse) SetBody(v *DescribeSubscriptionInstanceStatusResponseBody) *DescribeSubscriptionInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInstancesRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum                  *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize                 *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
}

func (s DescribeSubscriptionInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesRequest) SetClientToken(v string) *DescribeSubscriptionInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetOwnerId(v string) *DescribeSubscriptionInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageNum(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageSize(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesRequest {
	s.SubscriptionInstanceName = &v
	return s
}

type DescribeSubscriptionInstancesResponseBody struct {
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount       *int32                                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	SubscriptionInstances *DescribeSubscriptionInstancesResponseBodySubscriptionInstances `json:"SubscriptionInstances,omitempty" xml:"SubscriptionInstances,omitempty" type:"Struct"`
	TotalRecordCount      *int64                                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSubscriptionInstances(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) *DescribeSubscriptionInstancesResponseBody {
	s.SubscriptionInstances = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeSubscriptionInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstances struct {
	SubscriptionInstance []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance `json:"SubscriptionInstance,omitempty" xml:"SubscriptionInstance,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstances) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) SetSubscriptionInstance(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) *DescribeSubscriptionInstancesResponseBodySubscriptionInstances {
	s.SubscriptionInstance = v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance struct {
	BeginTimestamp           *string                                                                                                 `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	ConsumptionCheckpoint    *string                                                                                                 `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	ConsumptionClient        *string                                                                                                 `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	EndTimestamp             *string                                                                                                 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	ErrorMessage             *string                                                                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PayType                  *string                                                                                                 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	SourceEndpoint           *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Status                   *string                                                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscriptionDataType     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstanceID   *string                                                                                                 `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	SubscriptionInstanceName *string                                                                                                 `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionObject       *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetBeginTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionClient(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetEndTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetErrorMessage(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetPayType(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.PayType = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSourceEndpoint(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetStatus(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionDataType(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionObject(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionObject = v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint struct {
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceType(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType struct {
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDDL(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DDL = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDML(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DML = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject struct {
	SynchronousObject []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) SetSynchronousObject(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject {
	s.SynchronousObject = v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject struct {
	DatabaseName  *string                                                                                                                         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                                                         `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetDatabaseName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetTableList(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) SetTable(v []*string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList {
	s.Table = v
	return s
}

type DescribeSubscriptionInstancesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSubscriptionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstancesResponse) SetStatusCode(v int32) *DescribeSubscriptionInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponse) SetBody(v *DescribeSubscriptionInstancesResponseBody) *DescribeSubscriptionInstancesResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionObjectModifyStatusRequest struct {
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s DescribeSubscriptionObjectModifyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionObjectModifyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionObjectModifyStatusRequest) SetClientToken(v string) *DescribeSubscriptionObjectModifyStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusRequest) SetOwnerId(v string) *DescribeSubscriptionObjectModifyStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionObjectModifyStatusRequest {
	s.SubscriptionInstanceId = &v
	return s
}

type DescribeSubscriptionObjectModifyStatusResponseBody struct {
	Detail    *DescribeSubscriptionObjectModifyStatusResponseBodyDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
	Percent   *string                                                   `json:"Percent,omitempty" xml:"Percent,omitempty"`
	RequestId *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSubscriptionObjectModifyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionObjectModifyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBody) SetDetail(v *DescribeSubscriptionObjectModifyStatusResponseBodyDetail) *DescribeSubscriptionObjectModifyStatusResponseBody {
	s.Detail = v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBody) SetPercent(v string) *DescribeSubscriptionObjectModifyStatusResponseBody {
	s.Percent = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBody) SetRequestId(v string) *DescribeSubscriptionObjectModifyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBody) SetStatus(v string) *DescribeSubscriptionObjectModifyStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeSubscriptionObjectModifyStatusResponseBodyDetail struct {
	CheckItem []*DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem `json:"CheckItem,omitempty" xml:"CheckItem,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionObjectModifyStatusResponseBodyDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionObjectModifyStatusResponseBodyDetail) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBodyDetail) SetCheckItem(v []*DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) *DescribeSubscriptionObjectModifyStatusResponseBodyDetail {
	s.CheckItem = v
	return s
}

type DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ItemName     *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) SetCheckStatus(v string) *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) SetErrorMessage(v string) *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) SetItemName(v string) *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem {
	s.ItemName = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem) SetRepairMethod(v string) *DescribeSubscriptionObjectModifyStatusResponseBodyDetailCheckItem {
	s.RepairMethod = &v
	return s
}

type DescribeSubscriptionObjectModifyStatusResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSubscriptionObjectModifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionObjectModifyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionObjectModifyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionObjectModifyStatusResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionObjectModifyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponse) SetStatusCode(v int32) *DescribeSubscriptionObjectModifyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSubscriptionObjectModifyStatusResponse) SetBody(v *DescribeSubscriptionObjectModifyStatusResponseBody) *DescribeSubscriptionObjectModifyStatusResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobStatusRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeSynchronizationJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusRequest) SetClientToken(v string) *DescribeSynchronizationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetOwnerId(v string) *DescribeSynchronizationJobStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBody struct {
	Checkpoint                    *string                                                                    `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DataInitialization            *string                                                                    `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataInitializationStatus      *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	Delay                         *string                                                                    `json:"Delay,omitempty" xml:"Delay,omitempty"`
	DestinationEndpoint           *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	ErrorMessage                  *string                                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType                       *string                                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Performance                   *DescribeSynchronizationJobStatusResponseBodyPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SourceEndpoint                *DescribeSynchronizationJobStatusResponseBodySourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Status                        *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureInitialization       *string                                                                    `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	StructureInitializationStatus *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	SynchronizationJobClass       *string                                                                    `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	SynchronizationJobId          *string                                                                    `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName        *string                                                                    `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationObjects        []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetExpireTime(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPayType(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPerformance(v *DescribeSynchronizationJobStatusResponseBodyPerformance) *DescribeSynchronizationJobStatusResponseBody {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSourceEndpoint(v *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationObjects(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationObjects = v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus struct {
	Checkpoint   *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyPerformance struct {
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
	RPS  *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetFLOW(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.FLOW = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetRPS(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.RPS = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatus struct {
	Detail  []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	Percent *string                                                             `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ItemName     *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodySourceEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjects struct {
	NewSchemaName *string                                                                            `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	SchemaName    *string                                                                            `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableExcludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	TableIncludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetNewSchemaName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.NewSchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.SchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableIncludes = v
	return s
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) SetTableName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes {
	s.TableName = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) SetTableName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes {
	s.TableName = &v
	return s
}

type DescribeSynchronizationJobStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronizationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationJobStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponse) SetStatusCode(v int32) *DescribeSynchronizationJobStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponse) SetBody(v *DescribeSynchronizationJobStatusResponseBody) *DescribeSynchronizationJobStatusResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobsRequest struct {
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNum                *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SynchronizationJobName *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
}

func (s DescribeSynchronizationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsRequest) SetClientToken(v string) *DescribeSynchronizationJobsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetOwnerId(v string) *DescribeSynchronizationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageNum(v int32) *DescribeSynchronizationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageSize(v int32) *DescribeSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsRequest {
	s.SynchronizationJobName = &v
	return s
}

type DescribeSynchronizationJobsResponseBody struct {
	PageNumber               *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount          *int32                                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationInstances []*DescribeSynchronizationJobsResponseBodySynchronizationInstances `json:"SynchronizationInstances,omitempty" xml:"SynchronizationInstances,omitempty" type:"Repeated"`
	TotalRecordCount         *int64                                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetRequestId(v string) *DescribeSynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetSynchronizationInstances(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstances) *DescribeSynchronizationJobsResponseBody {
	s.SynchronizationInstances = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstances struct {
	DataInitialization            *string                                                                                       `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataInitializationStatus      *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	Delay                         *string                                                                                       `json:"Delay,omitempty" xml:"Delay,omitempty"`
	DestinationEndpoint           *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	ErrorMessage                  *string                                                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	PayType                       *string                                                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Performance                   *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Status                        *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureInitialization       *string                                                                                       `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	StructureInitializationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	SynchronizationJobClass       *string                                                                                       `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	SynchronizationJobId          *string                                                                                       `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName        *string                                                                                       `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationObjects        []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataSynchronizationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDestinationEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetExpireTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPayType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPerformance(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPrecheckStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSourceEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationObjects(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationObjects = v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance struct {
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
	RPS  *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetFLOW(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.FLOW = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetRPS(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.RPS = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus struct {
	Detail  []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	Percent *string                                                                                `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ItemName     *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceType = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects struct {
	NewSchemaName *string                                                                                               `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	SchemaName    *string                                                                                               `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	TableExcludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	TableIncludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetNewSchemaName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.NewSchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.SchemaName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableIncludes = v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) SetTableName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes {
	s.TableName = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes struct {
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) SetTableName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes {
	s.TableName = &v
	return s
}

type DescribeSynchronizationJobsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobsResponse) SetStatusCode(v int32) *DescribeSynchronizationJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobsResponse) SetBody(v *DescribeSynchronizationJobsResponseBody) *DescribeSynchronizationJobsResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationObjectModifyStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetClientToken(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetOwnerId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetTaskId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.TaskId = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBody struct {
	DataInitializationStatus      *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	ErrorMessage                  *string                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PrecheckStatus                *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	RequestId                     *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                        *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	StructureInitializationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus struct {
	Detail  []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	Percent *string                                                                      `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status  *string                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Detail = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail struct {
	CheckStatus  *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ItemName     *string `json:"ItemName,omitempty" xml:"ItemName,omitempty"`
	RepairMethod *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetCheckStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.CheckStatus = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetItemName(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.ItemName = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) SetRepairMethod(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronizationObjectModifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationObjectModifyStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationObjectModifyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetStatusCode(v int32) *DescribeSynchronizationObjectModifyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetBody(v *DescribeSynchronizationObjectModifyStatusResponseBody) *DescribeSynchronizationObjectModifyStatusResponse {
	s.Body = v
	return s
}

type ModifyConsumptionTimestampRequest struct {
	ConsumptionTimestamp   *string `json:"ConsumptionTimestamp,omitempty" xml:"ConsumptionTimestamp,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s ModifyConsumptionTimestampRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumptionTimestampRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampRequest) SetConsumptionTimestamp(v string) *ModifyConsumptionTimestampRequest {
	s.ConsumptionTimestamp = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetOwnerId(v string) *ModifyConsumptionTimestampRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetSubscriptionInstanceId(v string) *ModifyConsumptionTimestampRequest {
	s.SubscriptionInstanceId = &v
	return s
}

type ModifyConsumptionTimestampResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyConsumptionTimestampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumptionTimestampResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrCode(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrMessage(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetRequestId(v string) *ModifyConsumptionTimestampResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetSuccess(v string) *ModifyConsumptionTimestampResponseBody {
	s.Success = &v
	return s
}

type ModifyConsumptionTimestampResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyConsumptionTimestampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyConsumptionTimestampResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumptionTimestampResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampResponse) SetHeaders(v map[string]*string) *ModifyConsumptionTimestampResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumptionTimestampResponse) SetStatusCode(v int32) *ModifyConsumptionTimestampResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyConsumptionTimestampResponse) SetBody(v *ModifyConsumptionTimestampResponseBody) *ModifyConsumptionTimestampResponse {
	s.Body = v
	return s
}

type ModifyMigrationObjectRequest struct {
	ClientToken     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobId  *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	MigrationObject *string `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyMigrationObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyMigrationObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifyMigrationObjectRequest) SetClientToken(v string) *ModifyMigrationObjectRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyMigrationObjectRequest) SetMigrationJobId(v string) *ModifyMigrationObjectRequest {
	s.MigrationJobId = &v
	return s
}

func (s *ModifyMigrationObjectRequest) SetMigrationObject(v string) *ModifyMigrationObjectRequest {
	s.MigrationObject = &v
	return s
}

func (s *ModifyMigrationObjectRequest) SetOwnerId(v string) *ModifyMigrationObjectRequest {
	s.OwnerId = &v
	return s
}

type ModifyMigrationObjectResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyMigrationObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyMigrationObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMigrationObjectResponseBody) SetErrCode(v string) *ModifyMigrationObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyMigrationObjectResponseBody) SetErrMessage(v string) *ModifyMigrationObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyMigrationObjectResponseBody) SetRequestId(v string) *ModifyMigrationObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMigrationObjectResponseBody) SetSuccess(v string) *ModifyMigrationObjectResponseBody {
	s.Success = &v
	return s
}

type ModifyMigrationObjectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifyMigrationObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyMigrationObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyMigrationObjectResponse) GoString() string {
	return s.String()
}

func (s *ModifyMigrationObjectResponse) SetHeaders(v map[string]*string) *ModifyMigrationObjectResponse {
	s.Headers = v
	return s
}

func (s *ModifyMigrationObjectResponse) SetStatusCode(v int32) *ModifyMigrationObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyMigrationObjectResponse) SetBody(v *ModifyMigrationObjectResponseBody) *ModifyMigrationObjectResponse {
	s.Body = v
	return s
}

type ModifySubscriptionObjectRequest struct {
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	SubscriptionObject     *string `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
}

func (s ModifySubscriptionObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectRequest) SetOwnerId(v string) *ModifySubscriptionObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetSubscriptionInstanceId(v string) *ModifySubscriptionObjectRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetSubscriptionObject(v string) *ModifySubscriptionObjectRequest {
	s.SubscriptionObject = &v
	return s
}

type ModifySubscriptionObjectResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySubscriptionObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectResponseBody) SetErrCode(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetErrMessage(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetRequestId(v string) *ModifySubscriptionObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetSuccess(v string) *ModifySubscriptionObjectResponseBody {
	s.Success = &v
	return s
}

type ModifySubscriptionObjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySubscriptionObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySubscriptionObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionObjectResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectResponse) SetHeaders(v map[string]*string) *ModifySubscriptionObjectResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionObjectResponse) SetStatusCode(v int32) *ModifySubscriptionObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionObjectResponse) SetBody(v *ModifySubscriptionObjectResponseBody) *ModifySubscriptionObjectResponse {
	s.Body = v
	return s
}

type ModifySynchronizationObjectRequest struct {
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId   *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationObjects *string `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
}

func (s ModifySynchronizationObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectRequest) SetOwnerId(v string) *ModifySynchronizationObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetSynchronizationJobId(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetSynchronizationObjects(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationObjects = &v
	return s
}

type ModifySynchronizationObjectResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifySynchronizationObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponseBody) SetErrCode(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrMessage(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetRequestId(v string) *ModifySynchronizationObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetSuccess(v string) *ModifySynchronizationObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetTaskId(v string) *ModifySynchronizationObjectResponseBody {
	s.TaskId = &v
	return s
}

type ModifySynchronizationObjectResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ModifySynchronizationObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySynchronizationObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectResponse) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponse) SetHeaders(v map[string]*string) *ModifySynchronizationObjectResponse {
	s.Headers = v
	return s
}

func (s *ModifySynchronizationObjectResponse) SetStatusCode(v int32) *ModifySynchronizationObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySynchronizationObjectResponse) SetBody(v *ModifySynchronizationObjectResponseBody) *ModifySynchronizationObjectResponse {
	s.Body = v
	return s
}

type StartMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StartMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StartMigrationJobRequest) SetMigrationJobId(v string) *StartMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StartMigrationJobRequest) SetOwnerId(v string) *StartMigrationJobRequest {
	s.OwnerId = &v
	return s
}

type StartMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartMigrationJobResponseBody) SetErrCode(v string) *StartMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetErrMessage(v string) *StartMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetRequestId(v string) *StartMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetSuccess(v string) *StartMigrationJobResponseBody {
	s.Success = &v
	return s
}

type StartMigrationJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *StartMigrationJobResponse) SetHeaders(v map[string]*string) *StartMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *StartMigrationJobResponse) SetStatusCode(v int32) *StartMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartMigrationJobResponse) SetBody(v *StartMigrationJobResponseBody) *StartMigrationJobResponse {
	s.Body = v
	return s
}

type StartSubscriptionInstanceRequest struct {
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
}

func (s StartSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceRequest) SetOwnerId(v string) *StartSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *StartSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

type StartSubscriptionInstanceResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponseBody) SetErrCode(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrMessage(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetRequestId(v string) *StartSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetSuccess(v string) *StartSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetTaskId(v string) *StartSubscriptionInstanceResponseBody {
	s.TaskId = &v
	return s
}

type StartSubscriptionInstanceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartSubscriptionInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *StartSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartSubscriptionInstanceResponse) SetStatusCode(v int32) *StartSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSubscriptionInstanceResponse) SetBody(v *StartSubscriptionInstanceResponseBody) *StartSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type StartSynchronizationJobRequest struct {
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s StartSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobRequest) SetOwnerId(v string) *StartSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetSynchronizationJobId(v string) *StartSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

type StartSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobResponseBody) SetErrCode(v string) *StartSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetErrMessage(v string) *StartSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetRequestId(v string) *StartSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetSuccess(v string) *StartSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type StartSynchronizationJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobResponse) SetHeaders(v map[string]*string) *StartSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *StartSynchronizationJobResponse) SetStatusCode(v int32) *StartSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSynchronizationJobResponse) SetBody(v *StartSynchronizationJobResponseBody) *StartSynchronizationJobResponse {
	s.Body = v
	return s
}

type StopMigrationJobRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s StopMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StopMigrationJobRequest) SetClientToken(v string) *StopMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StopMigrationJobRequest) SetMigrationJobId(v string) *StopMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StopMigrationJobRequest) SetOwnerId(v string) *StopMigrationJobRequest {
	s.OwnerId = &v
	return s
}

type StopMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopMigrationJobResponseBody) SetErrCode(v string) *StopMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetErrMessage(v string) *StopMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetRequestId(v string) *StopMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetSuccess(v string) *StopMigrationJobResponseBody {
	s.Success = &v
	return s
}

type StopMigrationJobResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *StopMigrationJobResponse) SetHeaders(v map[string]*string) *StopMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *StopMigrationJobResponse) SetStatusCode(v int32) *StopMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopMigrationJobResponse) SetBody(v *StopMigrationJobResponseBody) *StopMigrationJobResponse {
	s.Body = v
	return s
}

type SuspendMigrationJobRequest struct {
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s SuspendMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobRequest) SetClientToken(v string) *SuspendMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetMigrationJobId(v string) *SuspendMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetOwnerId(v string) *SuspendMigrationJobRequest {
	s.OwnerId = &v
	return s
}

type SuspendMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobResponseBody) SetErrCode(v string) *SuspendMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetErrMessage(v string) *SuspendMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetRequestId(v string) *SuspendMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetSuccess(v string) *SuspendMigrationJobResponseBody {
	s.Success = &v
	return s
}

type SuspendMigrationJobResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendMigrationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendMigrationJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobResponse) SetHeaders(v map[string]*string) *SuspendMigrationJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendMigrationJobResponse) SetStatusCode(v int32) *SuspendMigrationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendMigrationJobResponse) SetBody(v *SuspendMigrationJobResponseBody) *SuspendMigrationJobResponse {
	s.Body = v
	return s
}

type SuspendSynchronizationJobRequest struct {
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s SuspendSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobRequest) SetOwnerId(v string) *SuspendSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetSynchronizationJobId(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

type SuspendSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobResponseBody) SetErrCode(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetErrMessage(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetRequestId(v string) *SuspendSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetSuccess(v string) *SuspendSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type SuspendSynchronizationJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SuspendSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobResponse) SetHeaders(v map[string]*string) *SuspendSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendSynchronizationJobResponse) SetStatusCode(v int32) *SuspendSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendSynchronizationJobResponse) SetBody(v *SuspendSynchronizationJobResponseBody) *SuspendSynchronizationJobResponse {
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
		"cn-qingdao":                  tea.String("dts.aliyuncs.com"),
		"cn-beijing":                  tea.String("dts.aliyuncs.com"),
		"cn-zhangjiakou":              tea.String("dts.aliyuncs.com"),
		"cn-huhehaote":                tea.String("dts.aliyuncs.com"),
		"cn-hangzhou":                 tea.String("dts.aliyuncs.com"),
		"cn-shanghai":                 tea.String("dts.aliyuncs.com"),
		"cn-shenzhen":                 tea.String("dts.aliyuncs.com"),
		"cn-hongkong":                 tea.String("dts.aliyuncs.com"),
		"ap-southeast-1":              tea.String("dts.aliyuncs.com"),
		"ap-southeast-2":              tea.String("dts.aliyuncs.com"),
		"ap-southeast-3":              tea.String("dts.aliyuncs.com"),
		"ap-southeast-5":              tea.String("dts.aliyuncs.com"),
		"eu-west-1":                   tea.String("dts.aliyuncs.com"),
		"us-west-1":                   tea.String("dts.aliyuncs.com"),
		"us-east-1":                   tea.String("dts.aliyuncs.com"),
		"eu-central-1":                tea.String("dts.aliyuncs.com"),
		"me-east-1":                   tea.String("dts.aliyuncs.com"),
		"ap-south-1":                  tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("dts.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("dts.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("dts.aliyuncs.com"),
		"cn-north-2-gov-1":            tea.String("dts.aliyuncs.com"),
		"ap-northeast-2-pop":          tea.String("dts.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("dts.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("dts.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("dts.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("dts.aliyuncs.com"),
		"cn-chengdu":                  tea.String("dts.aliyuncs.com"),
		"cn-edge-1":                   tea.String("dts.aliyuncs.com"),
		"cn-fujian":                   tea.String("dts.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("dts.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("dts.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("dts.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("dts.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("dts.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("dts.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("dts.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("dts.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("dts.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("dts.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("dts.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("dts.aliyuncs.com"),
		"cn-wuhan":                    tea.String("dts.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("dts.aliyuncs.com"),
		"cn-yushanfang":               tea.String("dts.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("dts.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("dts.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("dts.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("dts.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("dts.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("dts"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) ConfigurationSynchronizationJobWithOptions(request *ConfigurationSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *ConfigurationSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobName)) {
		query["SynchronizationJobName"] = request.SynchronizationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationObject)) {
		query["SynchronizationObject"] = request.SynchronizationObject
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationEndpoint)) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Initialization)) {
		query["Initialization"] = request.Initialization
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndpoint)) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigurationSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigurationSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigurationSynchronizationJob(request *ConfigurationSynchronizationJobRequest) (_result *ConfigurationSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigurationSynchronizationJobResponse{}
	_body, _err := client.ConfigurationSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureMigrationJobWithOptions(request *ConfigureMigrationJobRequest, runtime *util.RuntimeOptions) (_result *ConfigureMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Checkpoint)) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobName)) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationObject)) {
		query["MigrationObject"] = request.MigrationObject
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationReserved)) {
		query["MigrationReserved"] = request.MigrationReserved
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationEndpoint)) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationMode)) {
		query["MigrationMode"] = request.MigrationMode
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndpoint)) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigureMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigureMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureMigrationJob(request *ConfigureMigrationJobRequest) (_result *ConfigureMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureMigrationJobResponse{}
	_body, _err := client.ConfigureMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureSubscriptionInstanceWithOptions(request *ConfigureSubscriptionInstanceRequest, runtime *util.RuntimeOptions) (_result *ConfigureSubscriptionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceName)) {
		query["SubscriptionInstanceName"] = request.SubscriptionInstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionObject)) {
		query["SubscriptionObject"] = request.SubscriptionObject
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndpoint)) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionDataType)) {
		query["SubscriptionDataType"] = request.SubscriptionDataType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigureSubscriptionInstance"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigureSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSubscriptionInstance(request *ConfigureSubscriptionInstanceRequest) (_result *ConfigureSubscriptionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSubscriptionInstanceResponse{}
	_body, _err := client.ConfigureSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureSynchronizationJobWithOptions(request *ConfigureSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *ConfigureSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Checkpoint)) {
		query["Checkpoint"] = request.Checkpoint
	}

	if !tea.BoolValue(util.IsUnset(request.DataInitialization)) {
		query["DataInitialization"] = request.DataInitialization
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationReserved)) {
		query["MigrationReserved"] = request.MigrationReserved
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.StructureInitialization)) {
		query["StructureInitialization"] = request.StructureInitialization
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobName)) {
		query["SynchronizationJobName"] = request.SynchronizationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationObjects)) {
		query["SynchronizationObjects"] = request.SynchronizationObjects
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationEndpoint)) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.PartitionKey)) {
		query["PartitionKey"] = request.PartitionKey
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndpoint)) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigureSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigureSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSynchronizationJob(request *ConfigureSynchronizationJobRequest) (_result *ConfigureSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobResponse{}
	_body, _err := client.ConfigureSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateMigrationJobWithOptions(request *CreateMigrationJobRequest, runtime *util.RuntimeOptions) (_result *CreateMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobClass)) {
		query["MigrationJobClass"] = request.MigrationJobClass
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateMigrationJob(request *CreateMigrationJobRequest) (_result *CreateMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.CreateMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSubscriptionInstanceWithOptions(request *CreateSubscriptionInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateSubscriptionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSubscriptionInstance"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSubscriptionInstance(request *CreateSubscriptionInstanceRequest) (_result *CreateSubscriptionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSubscriptionInstanceResponse{}
	_body, _err := client.CreateSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSynchronizationJobWithOptions(request *CreateSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *CreateSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.DestRegion)) {
		query["DestRegion"] = request.DestRegion
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PayType)) {
		query["PayType"] = request.PayType
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.SourceRegion)) {
		query["SourceRegion"] = request.SourceRegion
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobClass)) {
		query["SynchronizationJobClass"] = request.SynchronizationJobClass
	}

	if !tea.BoolValue(util.IsUnset(request.UsedTime)) {
		query["UsedTime"] = request.UsedTime
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkType)) {
		query["networkType"] = request.NetworkType
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationEndpoint)) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !tea.BoolValue(util.IsUnset(request.SourceEndpoint)) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSynchronizationJob(request *CreateSynchronizationJobRequest) (_result *CreateSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSynchronizationJobResponse{}
	_body, _err := client.CreateSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMigrationJobWithOptions(request *DeleteMigrationJobRequest, runtime *util.RuntimeOptions) (_result *DeleteMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMigrationJob(request *DeleteMigrationJobRequest) (_result *DeleteMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.DeleteMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSubscriptionInstanceWithOptions(request *DeleteSubscriptionInstanceRequest, runtime *util.RuntimeOptions) (_result *DeleteSubscriptionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSubscriptionInstance"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSubscriptionInstance(request *DeleteSubscriptionInstanceRequest) (_result *DeleteSubscriptionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSubscriptionInstanceResponse{}
	_body, _err := client.DeleteSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSynchronizationJobWithOptions(request *DeleteSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *DeleteSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSynchronizationJob(request *DeleteSynchronizationJobRequest) (_result *DeleteSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSynchronizationJobResponse{}
	_body, _err := client.DeleteSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescirbeMigrationJobsWithOptions(request *DescirbeMigrationJobsRequest, runtime *util.RuntimeOptions) (_result *DescirbeMigrationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobName)) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescirbeMigrationJobs"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescirbeMigrationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescirbeMigrationJobs(request *DescirbeMigrationJobsRequest) (_result *DescirbeMigrationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescirbeMigrationJobsResponse{}
	_body, _err := client.DescirbeMigrationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeInitializationStatusWithOptions(request *DescribeInitializationStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeInitializationStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInitializationStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInitializationStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeInitializationStatus(request *DescribeInitializationStatusRequest) (_result *DescribeInitializationStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInitializationStatusResponse{}
	_body, _err := client.DescribeInitializationStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobDetailWithOptions(request *DescribeMigrationJobDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationMode)) {
		query["MigrationMode"] = request.MigrationMode
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobDetail"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMigrationJobDetailResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobDetail(request *DescribeMigrationJobDetailRequest) (_result *DescribeMigrationJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMigrationJobDetailResponse{}
	_body, _err := client.DescribeMigrationJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobStatusWithOptions(request *DescribeMigrationJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMigrationJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobStatus(request *DescribeMigrationJobStatusRequest) (_result *DescribeMigrationJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMigrationJobStatusResponse{}
	_body, _err := client.DescribeMigrationJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobsWithOptions(request *DescribeMigrationJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobName)) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobs"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeMigrationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobs(request *DescribeMigrationJobsRequest) (_result *DescribeMigrationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMigrationJobsResponse{}
	_body, _err := client.DescribeMigrationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstanceStatusWithOptions(request *DescribeSubscriptionInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionInstanceStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionInstanceStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstanceStatus(request *DescribeSubscriptionInstanceStatusRequest) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionInstanceStatusResponse{}
	_body, _err := client.DescribeSubscriptionInstanceStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstancesWithOptions(request *DescribeSubscriptionInstancesRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceName)) {
		query["SubscriptionInstanceName"] = request.SubscriptionInstanceName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionInstances"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstances(request *DescribeSubscriptionInstancesRequest) (_result *DescribeSubscriptionInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionInstancesResponse{}
	_body, _err := client.DescribeSubscriptionInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionObjectModifyStatusWithOptions(request *DescribeSubscriptionObjectModifyStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionObjectModifyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSubscriptionObjectModifyStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSubscriptionObjectModifyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionObjectModifyStatus(request *DescribeSubscriptionObjectModifyStatusRequest) (_result *DescribeSubscriptionObjectModifyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionObjectModifyStatusResponse{}
	_body, _err := client.DescribeSubscriptionObjectModifyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobStatusWithOptions(request *DescribeSynchronizationJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronizationJobStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronizationJobStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobStatus(request *DescribeSynchronizationJobStatusRequest) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationJobStatusResponse{}
	_body, _err := client.DescribeSynchronizationJobStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobsWithOptions(request *DescribeSynchronizationJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobName)) {
		query["SynchronizationJobName"] = request.SynchronizationJobName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronizationJobs"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronizationJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobs(request *DescribeSynchronizationJobsRequest) (_result *DescribeSynchronizationJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationJobsResponse{}
	_body, _err := client.DescribeSynchronizationJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronizationObjectModifyStatusWithOptions(request *DescribeSynchronizationObjectModifyStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationObjectModifyStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronizationObjectModifyStatus"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronizationObjectModifyStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationObjectModifyStatus(request *DescribeSynchronizationObjectModifyStatusRequest) (_result *DescribeSynchronizationObjectModifyStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationObjectModifyStatusResponse{}
	_body, _err := client.DescribeSynchronizationObjectModifyStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyConsumptionTimestampWithOptions(request *ModifyConsumptionTimestampRequest, runtime *util.RuntimeOptions) (_result *ModifyConsumptionTimestampResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ConsumptionTimestamp)) {
		query["ConsumptionTimestamp"] = request.ConsumptionTimestamp
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyConsumptionTimestamp"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyConsumptionTimestampResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyConsumptionTimestamp(request *ModifyConsumptionTimestampRequest) (_result *ModifyConsumptionTimestampResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyConsumptionTimestampResponse{}
	_body, _err := client.ModifyConsumptionTimestampWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyMigrationObjectWithOptions(request *ModifyMigrationObjectRequest, runtime *util.RuntimeOptions) (_result *ModifyMigrationObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationObject)) {
		query["MigrationObject"] = request.MigrationObject
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifyMigrationObject"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifyMigrationObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyMigrationObject(request *ModifyMigrationObjectRequest) (_result *ModifyMigrationObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyMigrationObjectResponse{}
	_body, _err := client.ModifyMigrationObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubscriptionObjectWithOptions(request *ModifySubscriptionObjectRequest, runtime *util.RuntimeOptions) (_result *ModifySubscriptionObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionObject)) {
		query["SubscriptionObject"] = request.SubscriptionObject
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySubscriptionObject"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySubscriptionObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubscriptionObject(request *ModifySubscriptionObjectRequest) (_result *ModifySubscriptionObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubscriptionObjectResponse{}
	_body, _err := client.ModifySubscriptionObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySynchronizationObjectWithOptions(request *ModifySynchronizationObjectRequest, runtime *util.RuntimeOptions) (_result *ModifySynchronizationObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationObjects)) {
		query["SynchronizationObjects"] = request.SynchronizationObjects
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ModifySynchronizationObject"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ModifySynchronizationObjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySynchronizationObject(request *ModifySynchronizationObjectRequest) (_result *ModifySynchronizationObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySynchronizationObjectResponse{}
	_body, _err := client.ModifySynchronizationObjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartMigrationJobWithOptions(request *StartMigrationJobRequest, runtime *util.RuntimeOptions) (_result *StartMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartMigrationJob(request *StartMigrationJobRequest) (_result *StartMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartMigrationJobResponse{}
	_body, _err := client.StartMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSubscriptionInstanceWithOptions(request *StartSubscriptionInstanceRequest, runtime *util.RuntimeOptions) (_result *StartSubscriptionInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SubscriptionInstanceId)) {
		query["SubscriptionInstanceId"] = request.SubscriptionInstanceId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSubscriptionInstance"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSubscriptionInstanceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSubscriptionInstance(request *StartSubscriptionInstanceRequest) (_result *StartSubscriptionInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSubscriptionInstanceResponse{}
	_body, _err := client.StartSubscriptionInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartSynchronizationJobWithOptions(request *StartSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *StartSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSynchronizationJob(request *StartSynchronizationJobRequest) (_result *StartSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSynchronizationJobResponse{}
	_body, _err := client.StartSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopMigrationJobWithOptions(request *StopMigrationJobRequest, runtime *util.RuntimeOptions) (_result *StopMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopMigrationJob(request *StopMigrationJobRequest) (_result *StopMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopMigrationJobResponse{}
	_body, _err := client.StopMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendMigrationJobWithOptions(request *SuspendMigrationJobRequest, runtime *util.RuntimeOptions) (_result *SuspendMigrationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendMigrationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendMigrationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendMigrationJob(request *SuspendMigrationJobRequest) (_result *SuspendMigrationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendMigrationJobResponse{}
	_body, _err := client.SuspendMigrationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendSynchronizationJobWithOptions(request *SuspendSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *SuspendSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendSynchronizationJob"),
		Version:     tea.String("2016-08-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SuspendSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendSynchronizationJob(request *SuspendSynchronizationJobRequest) (_result *SuspendSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendSynchronizationJobResponse{}
	_body, _err := client.SuspendSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
