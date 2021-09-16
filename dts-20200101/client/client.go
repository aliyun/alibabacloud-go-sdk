// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigureDtsJobRequest struct {
	DtsJobName                      *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	SourceEndpointInstanceType      *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointInstanceID        *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointEngineName        *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	SourceEndpointRegion            *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointIP                *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort              *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointOracleSID         *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	SourceEndpointDatabaseName      *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	SourceEndpointUserName          *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	SourceEndpointPassword          *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	SourceEndpointOwnerID           *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	SourceEndpointRole              *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	DestinationEndpointInstanceID   *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	DestinationEndpointEngineName   *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	DestinationEndpointRegion       *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	DestinationEndpointIP           *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	DestinationEndpointPort         *string `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	DestinationEndpointDataBaseName *string `json:"DestinationEndpointDataBaseName,omitempty" xml:"DestinationEndpointDataBaseName,omitempty"`
	DestinationEndpointUserName     *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	DestinationEndpointPassword     *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	StructureInitialization         *bool   `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	DataInitialization              *bool   `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization             *bool   `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	DbList                          *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	Reserve                         *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	Checkpoint                      *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	OwnerId                         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	DestinationEndpointOracleSID    *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	JobType                         *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	DtsJobId                        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId                   *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DelayPhone                      *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	DelayRuleTime                   *int64  `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	DelayNotice                     *bool   `json:"DelayNotice,omitempty" xml:"DelayNotice,omitempty"`
	ErrorPhone                      *string `json:"ErrorPhone,omitempty" xml:"ErrorPhone,omitempty"`
	ErrorNotice                     *bool   `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
	SynchronizationDirection        *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ConfigureDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobRequest) SetDtsJobName(v string) *ConfigureDtsJobRequest {
	s.DtsJobName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointInstanceType(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointInstanceID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointEngineName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointRegion(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointIP(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointPort(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointOracleSID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointDatabaseName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointUserName(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointPassword(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointOwnerID(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointOwnerID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSourceEndpointRole(v string) *ConfigureDtsJobRequest {
	s.SourceEndpointRole = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointInstanceType(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointInstanceID(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointEngineName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointRegion(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointIP(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointPort(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointDataBaseName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointDataBaseName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointUserName(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointPassword(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetStructureInitialization(v bool) *ConfigureDtsJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDataInitialization(v bool) *ConfigureDtsJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDataSynchronization(v bool) *ConfigureDtsJobRequest {
	s.DataSynchronization = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDbList(v string) *ConfigureDtsJobRequest {
	s.DbList = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetReserve(v string) *ConfigureDtsJobRequest {
	s.Reserve = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetCheckpoint(v string) *ConfigureDtsJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetOwnerId(v string) *ConfigureDtsJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDestinationEndpointOracleSID(v string) *ConfigureDtsJobRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetJobType(v string) *ConfigureDtsJobRequest {
	s.JobType = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsJobId(v string) *ConfigureDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDtsInstanceId(v string) *ConfigureDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayPhone(v string) *ConfigureDtsJobRequest {
	s.DelayPhone = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayRuleTime(v int64) *ConfigureDtsJobRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetDelayNotice(v bool) *ConfigureDtsJobRequest {
	s.DelayNotice = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetErrorPhone(v string) *ConfigureDtsJobRequest {
	s.ErrorPhone = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetErrorNotice(v bool) *ConfigureDtsJobRequest {
	s.ErrorNotice = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetSynchronizationDirection(v string) *ConfigureDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureDtsJobRequest) SetRegionId(v string) *ConfigureDtsJobRequest {
	s.RegionId = &v
	return s
}

type ConfigureDtsJobResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId  *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobResponseBody) SetHttpStatusCode(v string) *ConfigureDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetRequestId(v string) *ConfigureDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetErrCode(v string) *ConfigureDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetDtsJobId(v string) *ConfigureDtsJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetDtsInstanceId(v string) *ConfigureDtsJobResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetSuccess(v string) *ConfigureDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureDtsJobResponseBody) SetErrMessage(v string) *ConfigureDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureDtsJobResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ConfigureDtsJobResponse) SetHeaders(v map[string]*string) *ConfigureDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ConfigureDtsJobResponse) SetBody(v *ConfigureDtsJobResponseBody) *ConfigureDtsJobResponse {
	s.Body = v
	return s
}

type ConfigureMigrationJobRequest struct {
	SourceEndpoint      *ConfigureMigrationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	DestinationEndpoint *ConfigureMigrationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode       *ConfigureMigrationJobRequestMigrationMode       `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	RegionId            *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId      *string                                          `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	MigrationJobName    *string                                          `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationObject     *string                                          `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	MigrationReserved   *string                                          `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	Checkpoint          *string                                          `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	OwnerId             *string                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId           *string                                          `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequest) SetSourceEndpoint(v *ConfigureMigrationJobRequestSourceEndpoint) *ConfigureMigrationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetDestinationEndpoint(v *ConfigureMigrationJobRequestDestinationEndpoint) *ConfigureMigrationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetMigrationMode(v *ConfigureMigrationJobRequestMigrationMode) *ConfigureMigrationJobRequest {
	s.MigrationMode = v
	return s
}

func (s *ConfigureMigrationJobRequest) SetRegionId(v string) *ConfigureMigrationJobRequest {
	s.RegionId = &v
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

func (s *ConfigureMigrationJobRequest) SetCheckpoint(v string) *ConfigureMigrationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetOwnerId(v string) *ConfigureMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureMigrationJobRequest) SetAccountId(v string) *ConfigureMigrationJobRequest {
	s.AccountId = &v
	return s
}

type ConfigureMigrationJobRequestSourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ConfigureMigrationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetIP(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPort(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOracleSID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetRole(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

type ConfigureMigrationJobRequestDestinationEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetInstanceID(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetEngineName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetRegion(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPort(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetDataBaseName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.DataBaseName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureMigrationJobRequestDestinationEndpoint) SetOracleSID(v string) *ConfigureMigrationJobRequestDestinationEndpoint {
	s.OracleSID = &v
	return s
}

type ConfigureMigrationJobRequestMigrationMode struct {
	StructureIntialization *bool `json:"StructureIntialization,omitempty" xml:"StructureIntialization,omitempty"`
	DataIntialization      *bool `json:"DataIntialization,omitempty" xml:"DataIntialization,omitempty"`
	DataSynchronization    *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
}

func (s ConfigureMigrationJobRequestMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetStructureIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.StructureIntialization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataIntialization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataIntialization = &v
	return s
}

func (s *ConfigureMigrationJobRequestMigrationMode) SetDataSynchronization(v bool) *ConfigureMigrationJobRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

type ConfigureMigrationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobResponseBody) SetRequestId(v string) *ConfigureMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetErrCode(v string) *ConfigureMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetSuccess(v string) *ConfigureMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetErrMessage(v string) *ConfigureMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureMigrationJobResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigureMigrationJobResponse) SetBody(v *ConfigureMigrationJobResponseBody) *ConfigureMigrationJobResponse {
	s.Body = v
	return s
}

type ConfigureMigrationJobAlertRequest struct {
	RegionId         *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId   *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayAlertPhone  *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrorAlertPhone  *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	OwnerId          *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId        *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureMigrationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertRequest) SetRegionId(v string) *ConfigureMigrationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetMigrationJobId(v string) *ConfigureMigrationJobAlertRequest {
	s.MigrationJobId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayAlertStatus(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayAlertPhone(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetErrorAlertStatus(v string) *ConfigureMigrationJobAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetErrorAlertPhone(v string) *ConfigureMigrationJobAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetDelayOverSeconds(v string) *ConfigureMigrationJobAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetOwnerId(v string) *ConfigureMigrationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureMigrationJobAlertRequest) SetAccountId(v string) *ConfigureMigrationJobAlertRequest {
	s.AccountId = &v
	return s
}

type ConfigureMigrationJobAlertResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureMigrationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertResponseBody) SetRequestId(v string) *ConfigureMigrationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetErrCode(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetSuccess(v string) *ConfigureMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetErrMessage(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureMigrationJobAlertResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureMigrationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureMigrationJobAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureMigrationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureMigrationJobAlertResponse) SetHeaders(v map[string]*string) *ConfigureMigrationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureMigrationJobAlertResponse) SetBody(v *ConfigureMigrationJobAlertResponseBody) *ConfigureMigrationJobAlertResponse {
	s.Body = v
	return s
}

type ConfigureSubscriptionRequest struct {
	RegionId                        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsJobName                      *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	DtsInstanceId                   *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId                        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	SourceEndpointEngineName        *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	SourceEndpointInstanceType      *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointRegion            *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointInstanceID        *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointIP                *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort              *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointOracleSID         *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	SourceEndpointDatabaseName      *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	SourceEndpointUserName          *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	SourceEndpointPassword          *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	SourceEndpointOwnerID           *string `json:"SourceEndpointOwnerID,omitempty" xml:"SourceEndpointOwnerID,omitempty"`
	SourceEndpointRole              *string `json:"SourceEndpointRole,omitempty" xml:"SourceEndpointRole,omitempty"`
	DbList                          *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	Reserve                         *string `json:"Reserve,omitempty" xml:"Reserve,omitempty"`
	Checkpoint                      *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	SubscriptionInstanceNetworkType *string `json:"SubscriptionInstanceNetworkType,omitempty" xml:"SubscriptionInstanceNetworkType,omitempty"`
	SubscriptionInstanceVPCId       *string `json:"SubscriptionInstanceVPCId,omitempty" xml:"SubscriptionInstanceVPCId,omitempty"`
	SubscriptionInstanceVSwitchId   *string `json:"SubscriptionInstanceVSwitchId,omitempty" xml:"SubscriptionInstanceVSwitchId,omitempty"`
	SubscriptionDataTypeDDL         *bool   `json:"SubscriptionDataTypeDDL,omitempty" xml:"SubscriptionDataTypeDDL,omitempty"`
	SubscriptionDataTypeDML         *bool   `json:"SubscriptionDataTypeDML,omitempty" xml:"SubscriptionDataTypeDML,omitempty"`
	DelayPhone                      *string `json:"DelayPhone,omitempty" xml:"DelayPhone,omitempty"`
	DelayRuleTime                   *int64  `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	DelayNotice                     *bool   `json:"DelayNotice,omitempty" xml:"DelayNotice,omitempty"`
	ErrorPhone                      *string `json:"ErrorPhone,omitempty" xml:"ErrorPhone,omitempty"`
	ErrorNotice                     *bool   `json:"ErrorNotice,omitempty" xml:"ErrorNotice,omitempty"`
}

func (s ConfigureSubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionRequest) SetRegionId(v string) *ConfigureSubscriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsJobName(v string) *ConfigureSubscriptionRequest {
	s.DtsJobName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsInstanceId(v string) *ConfigureSubscriptionRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDtsJobId(v string) *ConfigureSubscriptionRequest {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointEngineName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointInstanceType(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointRegion(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointInstanceID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointIP(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointPort(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointOracleSID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointDatabaseName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointUserName(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointPassword(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointOwnerID(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointOwnerID = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSourceEndpointRole(v string) *ConfigureSubscriptionRequest {
	s.SourceEndpointRole = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDbList(v string) *ConfigureSubscriptionRequest {
	s.DbList = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetReserve(v string) *ConfigureSubscriptionRequest {
	s.Reserve = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetCheckpoint(v string) *ConfigureSubscriptionRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceNetworkType = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceVPCId(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceVPCId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionInstanceVSwitchId(v string) *ConfigureSubscriptionRequest {
	s.SubscriptionInstanceVSwitchId = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionDataTypeDDL(v bool) *ConfigureSubscriptionRequest {
	s.SubscriptionDataTypeDDL = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetSubscriptionDataTypeDML(v bool) *ConfigureSubscriptionRequest {
	s.SubscriptionDataTypeDML = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayPhone(v string) *ConfigureSubscriptionRequest {
	s.DelayPhone = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayRuleTime(v int64) *ConfigureSubscriptionRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetDelayNotice(v bool) *ConfigureSubscriptionRequest {
	s.DelayNotice = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetErrorPhone(v string) *ConfigureSubscriptionRequest {
	s.ErrorPhone = &v
	return s
}

func (s *ConfigureSubscriptionRequest) SetErrorNotice(v bool) *ConfigureSubscriptionRequest {
	s.ErrorNotice = &v
	return s
}

type ConfigureSubscriptionResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	DtsInstanceId  *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionResponseBody) SetHttpStatusCode(v string) *ConfigureSubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetRequestId(v string) *ConfigureSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetErrCode(v string) *ConfigureSubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetDtsJobId(v string) *ConfigureSubscriptionResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetSuccess(v string) *ConfigureSubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetDtsInstanceId(v string) *ConfigureSubscriptionResponseBody {
	s.DtsInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionResponseBody) SetErrMessage(v string) *ConfigureSubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSubscriptionResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionResponse) SetBody(v *ConfigureSubscriptionResponseBody) *ConfigureSubscriptionResponse {
	s.Body = v
	return s
}

type ConfigureSubscriptionInstanceRequest struct {
	SourceEndpoint                  *ConfigureSubscriptionInstanceRequestSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType            *ConfigureSubscriptionInstanceRequestSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstance            *ConfigureSubscriptionInstanceRequestSubscriptionInstance `json:"SubscriptionInstance,omitempty" xml:"SubscriptionInstance,omitempty" type:"Struct"`
	RegionId                        *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId          *string                                                   `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	SubscriptionInstanceName        *string                                                   `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionObject              *string                                                   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
	SubscriptionInstanceNetworkType *string                                                   `json:"SubscriptionInstanceNetworkType,omitempty" xml:"SubscriptionInstanceNetworkType,omitempty"`
	OwnerId                         *string                                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                       *string                                                   `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstance(v *ConfigureSubscriptionInstanceRequestSubscriptionInstance) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstance = v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetRegionId(v string) *ConfigureSubscriptionInstanceRequest {
	s.RegionId = &v
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

func (s *ConfigureSubscriptionInstanceRequest) SetSubscriptionInstanceNetworkType(v string) *ConfigureSubscriptionInstanceRequest {
	s.SubscriptionInstanceNetworkType = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetOwnerId(v string) *ConfigureSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequest) SetAccountId(v string) *ConfigureSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

type ConfigureSubscriptionInstanceRequestSourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetInstanceID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetIP(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetPort(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetUserName(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetPassword(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetOracleSID(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureSubscriptionInstanceRequestSourceEndpoint {
	s.DatabaseName = &v
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

type ConfigureSubscriptionInstanceRequestSubscriptionInstance struct {
	VPCId     *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionInstance) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceRequestSubscriptionInstance) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) SetVPCId(v string) *ConfigureSubscriptionInstanceRequestSubscriptionInstance {
	s.VPCId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceRequestSubscriptionInstance) SetVSwitchId(v string) *ConfigureSubscriptionInstanceRequestSubscriptionInstance {
	s.VSwitchId = &v
	return s
}

type ConfigureSubscriptionInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetRequestId(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSubscriptionInstanceResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigureSubscriptionInstanceResponse) SetBody(v *ConfigureSubscriptionInstanceResponseBody) *ConfigureSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type ConfigureSubscriptionInstanceAlertRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	DelayAlertStatus       *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayAlertPhone        *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	ErrorAlertStatus       *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrorAlertPhone        *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	DelayOverSeconds       *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureSubscriptionInstanceAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetRegionId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetSubscriptionInstanceId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetErrorAlertStatus(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetErrorAlertPhone(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetDelayOverSeconds(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetOwnerId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertRequest) SetAccountId(v string) *ConfigureSubscriptionInstanceAlertRequest {
	s.AccountId = &v
	return s
}

type ConfigureSubscriptionInstanceAlertResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSubscriptionInstanceAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetRequestId(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSubscriptionInstanceAlertResponse struct {
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSubscriptionInstanceAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSubscriptionInstanceAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSubscriptionInstanceAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSubscriptionInstanceAlertResponse) SetHeaders(v map[string]*string) *ConfigureSubscriptionInstanceAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponse) SetBody(v *ConfigureSubscriptionInstanceAlertResponseBody) *ConfigureSubscriptionInstanceAlertResponse {
	s.Body = v
	return s
}

type ConfigureSynchronizationJobRequest struct {
	SourceEndpoint           *ConfigureSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	DestinationEndpoint      *ConfigureSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	PartitionKey             *ConfigureSynchronizationJobRequestPartitionKey        `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty" type:"Struct"`
	RegionId                 *string                                                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobName   *string                                                `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationJobId     *string                                                `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string                                                `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	StructureInitialization  *bool                                                  `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	DataInitialization       *bool                                                  `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	SynchronizationObjects   *string                                                `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
	MigrationReserved        *string                                                `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	Checkpoint               *string                                                `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	OwnerId                  *string                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string                                                `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequest) SetSourceEndpoint(v *ConfigureSynchronizationJobRequestSourceEndpoint) *ConfigureSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetDestinationEndpoint(v *ConfigureSynchronizationJobRequestDestinationEndpoint) *ConfigureSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetPartitionKey(v *ConfigureSynchronizationJobRequestPartitionKey) *ConfigureSynchronizationJobRequest {
	s.PartitionKey = v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetRegionId(v string) *ConfigureSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobName(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetStructureInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.StructureInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetDataInitialization(v bool) *ConfigureSynchronizationJobRequest {
	s.DataInitialization = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationObjects(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationObjects = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetMigrationReserved(v string) *ConfigureSynchronizationJobRequest {
	s.MigrationReserved = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetCheckpoint(v string) *ConfigureSynchronizationJobRequest {
	s.Checkpoint = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetOwnerId(v string) *ConfigureSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequest) SetAccountId(v string) *ConfigureSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

type ConfigureSynchronizationJobRequestSourceEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	OwnerID      *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Role         *string `json:"Role,omitempty" xml:"Role,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.IP = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPort(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Port = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetOwnerID(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetRole(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.Role = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestSourceEndpoint) SetDatabaseName(v string) *ConfigureSynchronizationJobRequestSourceEndpoint {
	s.DatabaseName = &v
	return s
}

type ConfigureSynchronizationJobRequestDestinationEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceId(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetInstanceType(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetIP(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.IP = &v
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

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetPassword(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.Password = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestDestinationEndpoint) SetDataBaseName(v string) *ConfigureSynchronizationJobRequestDestinationEndpoint {
	s.DataBaseName = &v
	return s
}

type ConfigureSynchronizationJobRequestPartitionKey struct {
	ModifyTimeYear   *bool `json:"ModifyTime_Year,omitempty" xml:"ModifyTime_Year,omitempty"`
	ModifyTimeMonth  *bool `json:"ModifyTime_Month,omitempty" xml:"ModifyTime_Month,omitempty"`
	ModifyTimeDay    *bool `json:"ModifyTime_Day,omitempty" xml:"ModifyTime_Day,omitempty"`
	ModifyTimeHour   *bool `json:"ModifyTime_Hour,omitempty" xml:"ModifyTime_Hour,omitempty"`
	ModifyTimeMinute *bool `json:"ModifyTime_Minute,omitempty" xml:"ModifyTime_Minute,omitempty"`
}

func (s ConfigureSynchronizationJobRequestPartitionKey) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobRequestPartitionKey) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeYear(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeYear = &v
	return s
}

func (s *ConfigureSynchronizationJobRequestPartitionKey) SetModifyTimeMonth(v bool) *ConfigureSynchronizationJobRequestPartitionKey {
	s.ModifyTimeMonth = &v
	return s
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

type ConfigureSynchronizationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSynchronizationJobResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigureSynchronizationJobResponse) SetBody(v *ConfigureSynchronizationJobResponseBody) *ConfigureSynchronizationJobResponse {
	s.Body = v
	return s
}

type ConfigureSynchronizationJobAlertRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureSynchronizationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertRequest) SetRegionId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayOverSeconds(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetOwnerId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetAccountId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.AccountId = &v
	return s
}

type ConfigureSynchronizationJobAlertResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSynchronizationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSynchronizationJobAlertResponse struct {
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSynchronizationJobAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponse) SetBody(v *ConfigureSynchronizationJobAlertResponseBody) *ConfigureSynchronizationJobAlertResponse {
	s.Body = v
	return s
}

type ConfigureSynchronizationJobReplicatorCompareRequest struct {
	RegionId                               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId                   *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection               *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationReplicatorCompareEnable *bool   `json:"SynchronizationReplicatorCompareEnable,omitempty" xml:"SynchronizationReplicatorCompareEnable,omitempty"`
	ClientToken                            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ConfigureSynchronizationJobReplicatorCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetRegionId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetSynchronizationReplicatorCompareEnable(v bool) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationReplicatorCompareEnable = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetClientToken(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.ClientToken = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetOwnerId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareRequest) SetAccountId(v string) *ConfigureSynchronizationJobReplicatorCompareRequest {
	s.AccountId = &v
	return s
}

type ConfigureSynchronizationJobReplicatorCompareResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ConfigureSynchronizationJobReplicatorCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

type ConfigureSynchronizationJobReplicatorCompareResponse struct {
	Headers map[string]*string                                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ConfigureSynchronizationJobReplicatorCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ConfigureSynchronizationJobReplicatorCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobReplicatorCompareResponse) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) SetHeaders(v map[string]*string) *ConfigureSynchronizationJobReplicatorCompareResponse {
	s.Headers = v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponse) SetBody(v *ConfigureSynchronizationJobReplicatorCompareResponseBody) *ConfigureSynchronizationJobReplicatorCompareResponse {
	s.Body = v
	return s
}

type CreateConsumerChannelRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId         *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId              *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ConsumerGroupName     *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	ConsumerGroupPassword *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
}

func (s CreateConsumerChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelRequest) SetRegionId(v string) *CreateConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetDtsInstanceId(v string) *CreateConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetDtsJobId(v string) *CreateConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupName(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupPassword(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *CreateConsumerChannelRequest) SetConsumerGroupUserName(v string) *CreateConsumerChannelRequest {
	s.ConsumerGroupUserName = &v
	return s
}

type CreateConsumerChannelResponseBody struct {
	HttpStatusCode  *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode         *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success         *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage      *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
}

func (s CreateConsumerChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelResponseBody) SetHttpStatusCode(v string) *CreateConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetRequestId(v string) *CreateConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetErrCode(v string) *CreateConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetSuccess(v string) *CreateConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetErrMessage(v string) *CreateConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateConsumerChannelResponseBody) SetConsumerGroupID(v string) *CreateConsumerChannelResponseBody {
	s.ConsumerGroupID = &v
	return s
}

type CreateConsumerChannelResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConsumerChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerChannelResponse) SetHeaders(v map[string]*string) *CreateConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerChannelResponse) SetBody(v *CreateConsumerChannelResponseBody) *CreateConsumerChannelResponse {
	s.Body = v
	return s
}

type CreateConsumerGroupRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ConsumerGroupName      *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	ConsumerGroupUserName  *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	ConsumerGroupPassword  *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CreateConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupRequest) SetRegionId(v string) *CreateConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetSubscriptionInstanceId(v string) *CreateConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupName(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupUserName(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetConsumerGroupPassword(v string) *CreateConsumerGroupRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetOwnerId(v string) *CreateConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateConsumerGroupRequest) SetAccountId(v string) *CreateConsumerGroupRequest {
	s.AccountId = &v
	return s
}

type CreateConsumerGroupResponseBody struct {
	ConsumerGroupID *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode         *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success         *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage      *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CreateConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponseBody) SetConsumerGroupID(v string) *CreateConsumerGroupResponseBody {
	s.ConsumerGroupID = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetRequestId(v string) *CreateConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetErrCode(v string) *CreateConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v string) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetErrMessage(v string) *CreateConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

type CreateConsumerGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateConsumerGroupResponse) SetHeaders(v map[string]*string) *CreateConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateConsumerGroupResponse) SetBody(v *CreateConsumerGroupResponseBody) *CreateConsumerGroupResponse {
	s.Body = v
	return s
}

type CreateDtsInstanceRequest struct {
	JobId                         *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	InstanceClass                 *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	PayType                       *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                        *string `json:"Period,omitempty" xml:"Period,omitempty"`
	SyncArchitecture              *string `json:"SyncArchitecture,omitempty" xml:"SyncArchitecture,omitempty"`
	AutoStart                     *bool   `json:"AutoStart,omitempty" xml:"AutoStart,omitempty"`
	UsedTime                      *int32  `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	Quantity                      *int32  `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
	AutoPay                       *bool   `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	Type                          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DatabaseCount                 *int32  `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	SourceRegion                  *string `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	DestinationRegion             *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	SourceEndpointEngineName      *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	DestinationEndpointEngineName *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	ComputeUnit                   *int32  `json:"ComputeUnit,omitempty" xml:"ComputeUnit,omitempty"`
	RegionId                      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateDtsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDtsInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceRequest) SetJobId(v string) *CreateDtsInstanceRequest {
	s.JobId = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetInstanceClass(v string) *CreateDtsInstanceRequest {
	s.InstanceClass = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetPayType(v string) *CreateDtsInstanceRequest {
	s.PayType = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetPeriod(v string) *CreateDtsInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSyncArchitecture(v string) *CreateDtsInstanceRequest {
	s.SyncArchitecture = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetAutoStart(v bool) *CreateDtsInstanceRequest {
	s.AutoStart = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetUsedTime(v int32) *CreateDtsInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetQuantity(v int32) *CreateDtsInstanceRequest {
	s.Quantity = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetAutoPay(v bool) *CreateDtsInstanceRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetType(v string) *CreateDtsInstanceRequest {
	s.Type = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDatabaseCount(v int32) *CreateDtsInstanceRequest {
	s.DatabaseCount = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSourceRegion(v string) *CreateDtsInstanceRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDestinationRegion(v string) *CreateDtsInstanceRequest {
	s.DestinationRegion = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetSourceEndpointEngineName(v string) *CreateDtsInstanceRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetDestinationEndpointEngineName(v string) *CreateDtsInstanceRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetComputeUnit(v int32) *CreateDtsInstanceRequest {
	s.ComputeUnit = &v
	return s
}

func (s *CreateDtsInstanceRequest) SetRegionId(v string) *CreateDtsInstanceRequest {
	s.RegionId = &v
	return s
}

type CreateDtsInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	JobId      *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateDtsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDtsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceResponseBody) SetRequestId(v string) *CreateDtsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetErrCode(v string) *CreateDtsInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetSuccess(v string) *CreateDtsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetJobId(v string) *CreateDtsInstanceResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetErrMessage(v string) *CreateDtsInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetInstanceId(v string) *CreateDtsInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type CreateDtsInstanceResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDtsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDtsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDtsInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceResponse) SetHeaders(v map[string]*string) *CreateDtsInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDtsInstanceResponse) SetBody(v *CreateDtsInstanceResponseBody) *CreateDtsInstanceResponse {
	s.Body = v
	return s
}

type CreateJobMonitorRuleRequest struct {
	DtsJobId      *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	Phone         *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	DelayRuleTime *int64  `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateJobMonitorRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleRequest) SetDtsJobId(v string) *CreateJobMonitorRuleRequest {
	s.DtsJobId = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetType(v string) *CreateJobMonitorRuleRequest {
	s.Type = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetState(v string) *CreateJobMonitorRuleRequest {
	s.State = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetPhone(v string) *CreateJobMonitorRuleRequest {
	s.Phone = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetDelayRuleTime(v int64) *CreateJobMonitorRuleRequest {
	s.DelayRuleTime = &v
	return s
}

func (s *CreateJobMonitorRuleRequest) SetRegionId(v string) *CreateJobMonitorRuleRequest {
	s.RegionId = &v
	return s
}

type CreateJobMonitorRuleResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
}

func (s CreateJobMonitorRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleResponseBody) SetHttpStatusCode(v int32) *CreateJobMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetRequestId(v string) *CreateJobMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetDtsJobId(v string) *CreateJobMonitorRuleResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetErrCode(v string) *CreateJobMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetSuccess(v bool) *CreateJobMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetErrMessage(v string) *CreateJobMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetCode(v string) *CreateJobMonitorRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateJobMonitorRuleResponseBody) SetDynamicMessage(v string) *CreateJobMonitorRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

type CreateJobMonitorRuleResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobMonitorRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateJobMonitorRuleResponse) SetHeaders(v map[string]*string) *CreateJobMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateJobMonitorRuleResponse) SetBody(v *CreateJobMonitorRuleResponseBody) *CreateJobMonitorRuleResponse {
	s.Body = v
	return s
}

type CreateMigrationJobRequest struct {
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	OwnerId           *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	AccountId         *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CreateMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobRequest) SetRegionId(v string) *CreateMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateMigrationJobRequest) SetRegion(v string) *CreateMigrationJobRequest {
	s.Region = &v
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

func (s *CreateMigrationJobRequest) SetClientToken(v string) *CreateMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateMigrationJobRequest) SetAccountId(v string) *CreateMigrationJobRequest {
	s.AccountId = &v
	return s
}

type CreateMigrationJobResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CreateMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponseBody) SetRequestId(v string) *CreateMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetMigrationJobId(v string) *CreateMigrationJobResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetErrCode(v string) *CreateMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetSuccess(v string) *CreateMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetErrMessage(v string) *CreateMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type CreateMigrationJobResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateMigrationJobResponse) SetBody(v *CreateMigrationJobResponseBody) *CreateMigrationJobResponse {
	s.Body = v
	return s
}

type CreateSubscriptionInstanceRequest struct {
	SourceEndpoint *CreateSubscriptionInstanceRequestSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	RegionId       *string                                          `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Region         *string                                          `json:"Region,omitempty" xml:"Region,omitempty"`
	PayType        *string                                          `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period         *string                                          `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime       *int32                                           `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken    *string                                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string                                          `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string                                          `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s CreateSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceRequest) SetSourceEndpoint(v *CreateSubscriptionInstanceRequestSourceEndpoint) *CreateSubscriptionInstanceRequest {
	s.SourceEndpoint = v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetRegionId(v string) *CreateSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetRegion(v string) *CreateSubscriptionInstanceRequest {
	s.Region = &v
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

func (s *CreateSubscriptionInstanceRequest) SetUsedTime(v int32) *CreateSubscriptionInstanceRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetClientToken(v string) *CreateSubscriptionInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetOwnerId(v string) *CreateSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSubscriptionInstanceRequest) SetAccountId(v string) *CreateSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

type CreateSubscriptionInstanceRequestSourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s CreateSubscriptionInstanceRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceRequestSourceEndpoint) SetInstanceType(v string) *CreateSubscriptionInstanceRequestSourceEndpoint {
	s.InstanceType = &v
	return s
}

type CreateSubscriptionInstanceResponseBody struct {
	RequestId              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ErrCode                *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage             *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CreateSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscriptionInstanceResponseBody) SetRequestId(v string) *CreateSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSubscriptionInstanceId(v string) *CreateSubscriptionInstanceResponseBody {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrCode(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSuccess(v string) *CreateSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrMessage(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

type CreateSubscriptionInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSubscriptionInstanceResponse) SetBody(v *CreateSubscriptionInstanceResponseBody) *CreateSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type CreateSynchronizationJobRequest struct {
	SourceEndpoint          *CreateSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	DestinationEndpoint     *CreateSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	RegionId                *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceRegion            *string                                             `json:"SourceRegion,omitempty" xml:"SourceRegion,omitempty"`
	DestRegion              *string                                             `json:"DestRegion,omitempty" xml:"DestRegion,omitempty"`
	Topology                *string                                             `json:"Topology,omitempty" xml:"Topology,omitempty"`
	SynchronizationJobClass *string                                             `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	PayType                 *string                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Period                  *string                                             `json:"Period,omitempty" xml:"Period,omitempty"`
	UsedTime                *int32                                              `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	ClientToken             *string                                             `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	NetworkType             *string                                             `json:"networkType,omitempty" xml:"networkType,omitempty"`
	OwnerId                 *string                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId               *string                                             `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DBInstanceCount         *int32                                              `json:"DBInstanceCount,omitempty" xml:"DBInstanceCount,omitempty"`
}

func (s CreateSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobRequest) SetSourceEndpoint(v *CreateSynchronizationJobRequestSourceEndpoint) *CreateSynchronizationJobRequest {
	s.SourceEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDestinationEndpoint(v *CreateSynchronizationJobRequestDestinationEndpoint) *CreateSynchronizationJobRequest {
	s.DestinationEndpoint = v
	return s
}

func (s *CreateSynchronizationJobRequest) SetRegionId(v string) *CreateSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSourceRegion(v string) *CreateSynchronizationJobRequest {
	s.SourceRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDestRegion(v string) *CreateSynchronizationJobRequest {
	s.DestRegion = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetTopology(v string) *CreateSynchronizationJobRequest {
	s.Topology = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetSynchronizationJobClass(v string) *CreateSynchronizationJobRequest {
	s.SynchronizationJobClass = &v
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

func (s *CreateSynchronizationJobRequest) SetUsedTime(v int32) *CreateSynchronizationJobRequest {
	s.UsedTime = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetClientToken(v string) *CreateSynchronizationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetNetworkType(v string) *CreateSynchronizationJobRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetOwnerId(v string) *CreateSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetAccountId(v string) *CreateSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

func (s *CreateSynchronizationJobRequest) SetDBInstanceCount(v int32) *CreateSynchronizationJobRequest {
	s.DBInstanceCount = &v
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

type CreateSynchronizationJobResponseBody struct {
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode              *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success              *string `json:"Success,omitempty" xml:"Success,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	ErrMessage           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s CreateSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponseBody) SetRequestId(v string) *CreateSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetErrCode(v string) *CreateSynchronizationJobResponseBody {
	s.ErrCode = &v
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

func (s *CreateSynchronizationJobResponseBody) SetErrMessage(v string) *CreateSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type CreateSynchronizationJobResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateSynchronizationJobResponse) SetBody(v *CreateSynchronizationJobResponseBody) *CreateSynchronizationJobResponse {
	s.Body = v
	return s
}

type DeleteConsumerChannelRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId   *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
}

func (s DeleteConsumerChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelRequest) SetRegionId(v string) *DeleteConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetDtsInstanceId(v string) *DeleteConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetDtsJobId(v string) *DeleteConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *DeleteConsumerChannelRequest) SetConsumerGroupId(v string) *DeleteConsumerChannelRequest {
	s.ConsumerGroupId = &v
	return s
}

type DeleteConsumerChannelResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteConsumerChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelResponseBody) SetHttpStatusCode(v string) *DeleteConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetRequestId(v string) *DeleteConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetErrCode(v string) *DeleteConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetSuccess(v string) *DeleteConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerChannelResponseBody) SetErrMessage(v string) *DeleteConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteConsumerChannelResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConsumerChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerChannelResponse) SetHeaders(v map[string]*string) *DeleteConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerChannelResponse) SetBody(v *DeleteConsumerChannelResponseBody) *DeleteConsumerChannelResponse {
	s.Body = v
	return s
}

type DeleteConsumerGroupRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ConsumerGroupID        *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupRequest) SetRegionId(v string) *DeleteConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetSubscriptionInstanceId(v string) *DeleteConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetConsumerGroupID(v string) *DeleteConsumerGroupRequest {
	s.ConsumerGroupID = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetOwnerId(v string) *DeleteConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteConsumerGroupRequest) SetAccountId(v string) *DeleteConsumerGroupRequest {
	s.AccountId = &v
	return s
}

type DeleteConsumerGroupResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetErrCode(v string) *DeleteConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v string) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetErrMessage(v string) *DeleteConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteConsumerGroupResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponse) SetHeaders(v map[string]*string) *DeleteConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteConsumerGroupResponse) SetBody(v *DeleteConsumerGroupResponseBody) *DeleteConsumerGroupResponse {
	s.Body = v
	return s
}

type DeleteDtsJobRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDtsJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobRequest) SetDtsJobId(v string) *DeleteDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetDtsInstanceId(v string) *DeleteDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DeleteDtsJobRequest) SetSynchronizationDirection(v string) *DeleteDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DeleteDtsJobRequest) SetRegionId(v string) *DeleteDtsJobRequest {
	s.RegionId = &v
	return s
}

type DeleteDtsJobResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s DeleteDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobResponseBody) SetHttpStatusCode(v int32) *DeleteDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetRequestId(v string) *DeleteDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrCode(v string) *DeleteDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetSuccess(v bool) *DeleteDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetErrMessage(v string) *DeleteDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetDynamicMessage(v string) *DeleteDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DeleteDtsJobResponseBody) SetDynamicCode(v string) *DeleteDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

type DeleteDtsJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDtsJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDtsJobResponse) SetHeaders(v map[string]*string) *DeleteDtsJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDtsJobResponse) SetBody(v *DeleteDtsJobResponseBody) *DeleteDtsJobResponse {
	s.Body = v
	return s
}

type DeleteMigrationJobRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobRequest) SetRegionId(v string) *DeleteMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetMigrationJobId(v string) *DeleteMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetOwnerId(v string) *DeleteMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteMigrationJobRequest) SetAccountId(v string) *DeleteMigrationJobRequest {
	s.AccountId = &v
	return s
}

type DeleteMigrationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponseBody) SetRequestId(v string) *DeleteMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetErrCode(v string) *DeleteMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetSuccess(v string) *DeleteMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetErrMessage(v string) *DeleteMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteMigrationJobResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteMigrationJobResponse) SetBody(v *DeleteMigrationJobResponseBody) *DeleteMigrationJobResponse {
	s.Body = v
	return s
}

type DeleteSubscriptionInstanceRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceRequest) SetRegionId(v string) *DeleteSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *DeleteSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetOwnerId(v string) *DeleteSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSubscriptionInstanceRequest) SetAccountId(v string) *DeleteSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

type DeleteSubscriptionInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceResponseBody) SetRequestId(v string) *DeleteSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrCode(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetSuccess(v string) *DeleteSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrMessage(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteSubscriptionInstanceResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSubscriptionInstanceResponse) SetBody(v *DeleteSubscriptionInstanceResponseBody) *DeleteSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type DeleteSynchronizationJobRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DeleteSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobRequest) SetRegionId(v string) *DeleteSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetSynchronizationJobId(v string) *DeleteSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetOwnerId(v string) *DeleteSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSynchronizationJobRequest) SetAccountId(v string) *DeleteSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

type DeleteSynchronizationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DeleteSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobResponseBody) SetRequestId(v string) *DeleteSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetErrCode(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetSuccess(v string) *DeleteSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetErrMessage(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type DeleteSynchronizationJobResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteSynchronizationJobResponse) SetBody(v *DeleteSynchronizationJobResponseBody) *DeleteSynchronizationJobResponse {
	s.Body = v
	return s
}

type DescribeConnectionStatusRequest struct {
	SourceEndpointArchitecture      *string `json:"SourceEndpointArchitecture,omitempty" xml:"SourceEndpointArchitecture,omitempty"`
	SourceEndpointInstanceType      *string `json:"SourceEndpointInstanceType,omitempty" xml:"SourceEndpointInstanceType,omitempty"`
	SourceEndpointInstanceID        *string `json:"SourceEndpointInstanceID,omitempty" xml:"SourceEndpointInstanceID,omitempty"`
	SourceEndpointEngineName        *string `json:"SourceEndpointEngineName,omitempty" xml:"SourceEndpointEngineName,omitempty"`
	SourceEndpointRegion            *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	SourceEndpointIP                *string `json:"SourceEndpointIP,omitempty" xml:"SourceEndpointIP,omitempty"`
	SourceEndpointPort              *string `json:"SourceEndpointPort,omitempty" xml:"SourceEndpointPort,omitempty"`
	SourceEndpointOracleSID         *string `json:"SourceEndpointOracleSID,omitempty" xml:"SourceEndpointOracleSID,omitempty"`
	SourceEndpointDatabaseName      *string `json:"SourceEndpointDatabaseName,omitempty" xml:"SourceEndpointDatabaseName,omitempty"`
	SourceEndpointUserName          *string `json:"SourceEndpointUserName,omitempty" xml:"SourceEndpointUserName,omitempty"`
	SourceEndpointPassword          *string `json:"SourceEndpointPassword,omitempty" xml:"SourceEndpointPassword,omitempty"`
	DestinationEndpointInstanceType *string `json:"DestinationEndpointInstanceType,omitempty" xml:"DestinationEndpointInstanceType,omitempty"`
	DestinationEndpointInstanceID   *string `json:"DestinationEndpointInstanceID,omitempty" xml:"DestinationEndpointInstanceID,omitempty"`
	DestinationEndpointEngineName   *string `json:"DestinationEndpointEngineName,omitempty" xml:"DestinationEndpointEngineName,omitempty"`
	DestinationEndpointRegion       *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
	DestinationEndpointIP           *string `json:"DestinationEndpointIP,omitempty" xml:"DestinationEndpointIP,omitempty"`
	DestinationEndpointPort         *string `json:"DestinationEndpointPort,omitempty" xml:"DestinationEndpointPort,omitempty"`
	DestinationEndpointDatabaseName *string `json:"DestinationEndpointDatabaseName,omitempty" xml:"DestinationEndpointDatabaseName,omitempty"`
	DestinationEndpointUserName     *string `json:"DestinationEndpointUserName,omitempty" xml:"DestinationEndpointUserName,omitempty"`
	DestinationEndpointPassword     *string `json:"DestinationEndpointPassword,omitempty" xml:"DestinationEndpointPassword,omitempty"`
	DestinationEndpointOracleSID    *string `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	DestinationEndpointArchitecture *string `json:"DestinationEndpointArchitecture,omitempty" xml:"DestinationEndpointArchitecture,omitempty"`
	RegionId                        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeConnectionStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointArchitecture(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointArchitecture = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointInstanceType(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointInstanceType = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointInstanceID(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointInstanceID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointEngineName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointEngineName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointRegion(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointIP(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointIP = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointPort(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointPort = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointOracleSID(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointOracleSID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointDatabaseName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointDatabaseName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointUserName(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointUserName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetSourceEndpointPassword(v string) *DescribeConnectionStatusRequest {
	s.SourceEndpointPassword = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointInstanceType(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointInstanceType = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointInstanceID(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointInstanceID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointEngineName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointEngineName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointRegion(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointRegion = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointIP(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointIP = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointPort(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointPort = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointDatabaseName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointDatabaseName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointUserName(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointUserName = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointPassword(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointPassword = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointOracleSID(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointArchitecture(v string) *DescribeConnectionStatusRequest {
	s.DestinationEndpointArchitecture = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetRegionId(v string) *DescribeConnectionStatusRequest {
	s.RegionId = &v
	return s
}

type DescribeConnectionStatusResponseBody struct {
	SourceConnectionStatus      map[string]interface{} `json:"SourceConnectionStatus,omitempty" xml:"SourceConnectionStatus,omitempty"`
	DestinationConnectionStatus map[string]interface{} `json:"DestinationConnectionStatus,omitempty" xml:"DestinationConnectionStatus,omitempty"`
	RequestId                   *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                     *string                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                     *string                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage                  *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DescribeConnectionStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusResponseBody) SetSourceConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody {
	s.SourceConnectionStatus = v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetDestinationConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody {
	s.DestinationConnectionStatus = v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetRequestId(v string) *DescribeConnectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrCode(v string) *DescribeConnectionStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetSuccess(v string) *DescribeConnectionStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrMessage(v string) *DescribeConnectionStatusResponseBody {
	s.ErrMessage = &v
	return s
}

type DescribeConnectionStatusResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConnectionStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConnectionStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConnectionStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeConnectionStatusResponse) SetHeaders(v map[string]*string) *DescribeConnectionStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeConnectionStatusResponse) SetBody(v *DescribeConnectionStatusResponseBody) *DescribeConnectionStatusResponse {
	s.Body = v
	return s
}

type DescribeConsumerChannelRequest struct {
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId   *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	PageSize        *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber      *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ParentChannelId *string `json:"ParentChannelId,omitempty" xml:"ParentChannelId,omitempty"`
}

func (s DescribeConsumerChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelRequest) SetRegionId(v string) *DescribeConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetDtsInstanceId(v string) *DescribeConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetDtsJobId(v string) *DescribeConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetPageSize(v int32) *DescribeConsumerChannelRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetPageNumber(v int32) *DescribeConsumerChannelRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerChannelRequest) SetParentChannelId(v string) *DescribeConsumerChannelRequest {
	s.ParentChannelId = &v
	return s
}

type DescribeConsumerChannelResponseBody struct {
	HttpStatusCode   *string                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId        *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode          *string                                                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PageRecordCount  *int32                                                 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Success          *string                                                `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalRecordCount *int64                                                 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage       *string                                                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber       *int32                                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConsumerChannels []*DescribeConsumerChannelResponseBodyConsumerChannels `json:"ConsumerChannels,omitempty" xml:"ConsumerChannels,omitempty" type:"Repeated"`
}

func (s DescribeConsumerChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponseBody) SetHttpStatusCode(v string) *DescribeConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetRequestId(v string) *DescribeConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetErrCode(v string) *DescribeConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetPageRecordCount(v int32) *DescribeConsumerChannelResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetSuccess(v string) *DescribeConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetTotalRecordCount(v int64) *DescribeConsumerChannelResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetErrMessage(v string) *DescribeConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetPageNumber(v int32) *DescribeConsumerChannelResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerChannelResponseBody) SetConsumerChannels(v []*DescribeConsumerChannelResponseBodyConsumerChannels) *DescribeConsumerChannelResponseBody {
	s.ConsumerChannels = v
	return s
}

type DescribeConsumerChannelResponseBodyConsumerChannels struct {
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	ConsumerGroupId       *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	MessageDelay          *int64  `json:"MessageDelay,omitempty" xml:"MessageDelay,omitempty"`
	ConsumerGroupName     *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	UnconsumedData        *int64  `json:"UnconsumedData,omitempty" xml:"UnconsumedData,omitempty"`
}

func (s DescribeConsumerChannelResponseBodyConsumerChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerChannelResponseBodyConsumerChannels) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupUserName(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupId(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetMessageDelay(v int64) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.MessageDelay = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumerGroupName(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetConsumptionCheckpoint(v string) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeConsumerChannelResponseBodyConsumerChannels) SetUnconsumedData(v int64) *DescribeConsumerChannelResponseBodyConsumerChannels {
	s.UnconsumedData = &v
	return s
}

type DescribeConsumerChannelResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConsumerChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumerChannelResponse) SetHeaders(v map[string]*string) *DescribeConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumerChannelResponse) SetBody(v *DescribeConsumerChannelResponseBody) *DescribeConsumerChannelResponse {
	s.Body = v
	return s
}

type DescribeConsumerGroupRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum                *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeConsumerGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupRequest) SetRegionId(v string) *DescribeConsumerGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetPageSize(v int32) *DescribeConsumerGroupRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetPageNum(v int32) *DescribeConsumerGroupRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetSubscriptionInstanceId(v string) *DescribeConsumerGroupRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetOwnerId(v string) *DescribeConsumerGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeConsumerGroupRequest) SetAccountId(v string) *DescribeConsumerGroupRequest {
	s.AccountId = &v
	return s
}

type DescribeConsumerGroupResponseBody struct {
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode          *string                                            `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success          *string                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	PageRecordCount  *int32                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage       *string                                            `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber       *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ConsumerChannels *DescribeConsumerGroupResponseBodyConsumerChannels `json:"ConsumerChannels,omitempty" xml:"ConsumerChannels,omitempty" type:"Struct"`
}

func (s DescribeConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBody) SetRequestId(v string) *DescribeConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrCode(v string) *DescribeConsumerGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetSuccess(v string) *DescribeConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetTotalRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrMessage(v string) *DescribeConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageNumber(v int32) *DescribeConsumerGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetConsumerChannels(v *DescribeConsumerGroupResponseBodyConsumerChannels) *DescribeConsumerGroupResponseBody {
	s.ConsumerChannels = v
	return s
}

type DescribeConsumerGroupResponseBodyConsumerChannels struct {
	DescribeConsumerChannel []*DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel `json:"DescribeConsumerChannel,omitempty" xml:"DescribeConsumerChannel,omitempty" type:"Repeated"`
}

func (s DescribeConsumerGroupResponseBodyConsumerChannels) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupResponseBodyConsumerChannels) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannels) SetDescribeConsumerChannel(v []*DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) *DescribeConsumerGroupResponseBodyConsumerChannels {
	s.DescribeConsumerChannel = v
	return s
}

type DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel struct {
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	ConsumerGroupID       *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	MessageDelay          *int64  `json:"MessageDelay,omitempty" xml:"MessageDelay,omitempty"`
	ConsumerGroupName     *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	ConsumptionCheckpoint *string `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	UnconsumedData        *int64  `json:"UnconsumedData,omitempty" xml:"UnconsumedData,omitempty"`
}

func (s DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupUserName(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupID(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupID = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetMessageDelay(v int64) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.MessageDelay = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumerGroupName(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetConsumptionCheckpoint(v string) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel) SetUnconsumedData(v int64) *DescribeConsumerGroupResponseBodyConsumerChannelsDescribeConsumerChannel {
	s.UnconsumedData = &v
	return s
}

type DescribeConsumerGroupResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeConsumerGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeConsumerGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponse) SetHeaders(v map[string]*string) *DescribeConsumerGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsumerGroupResponse) SetBody(v *DescribeConsumerGroupResponseBody) *DescribeConsumerGroupResponse {
	s.Body = v
	return s
}

type DescribeDTSIPRequest struct {
	RegionId                  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SourceEndpointRegion      *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
}

func (s DescribeDTSIPRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDTSIPRequest) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPRequest) SetRegionId(v string) *DescribeDTSIPRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDTSIPRequest) SetSourceEndpointRegion(v string) *DescribeDTSIPRequest {
	s.SourceEndpointRegion = &v
	return s
}

func (s *DescribeDTSIPRequest) SetDestinationEndpointRegion(v string) *DescribeDTSIPRequest {
	s.DestinationEndpointRegion = &v
	return s
}

type DescribeDTSIPResponseBody struct {
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s DescribeDTSIPResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDTSIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPResponseBody) SetRequestId(v string) *DescribeDTSIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrCode(v string) *DescribeDTSIPResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetSuccess(v string) *DescribeDTSIPResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrMessage(v string) *DescribeDTSIPResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetDynamicMessage(v string) *DescribeDTSIPResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetDynamicCode(v string) *DescribeDTSIPResponseBody {
	s.DynamicCode = &v
	return s
}

type DescribeDTSIPResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDTSIPResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDTSIPResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDTSIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPResponse) SetHeaders(v map[string]*string) *DescribeDTSIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeDTSIPResponse) SetBody(v *DescribeDTSIPResponseBody) *DescribeDTSIPResponse {
	s.Body = v
	return s
}

type DescribeDtsJobDetailRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceID            *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeDtsJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailRequest) SetDtsJobId(v string) *DescribeDtsJobDetailRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetDtsInstanceID(v string) *DescribeDtsJobDetailRequest {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetSynchronizationDirection(v string) *DescribeDtsJobDetailRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailRequest) SetRegionId(v string) *DescribeDtsJobDetailRequest {
	s.RegionId = &v
	return s
}

type DescribeDtsJobDetailResponseBody struct {
	Status                   *string                                               `json:"Status,omitempty" xml:"Status,omitempty"`
	DtsJobName               *string                                               `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	FinishTime               *string                                               `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	ErrorMessage             *string                                               `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DtsJobId                 *string                                               `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	CreateTime               *string                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                  *string                                               `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reserved                 *string                                               `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	DatabaseCount            *int32                                                `json:"DatabaseCount,omitempty" xml:"DatabaseCount,omitempty"`
	DtsJobClass              *string                                               `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	EndTimestamp             *string                                               `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	AppName                  *string                                               `json:"AppName,omitempty" xml:"AppName,omitempty"`
	DestNetType              *string                                               `json:"DestNetType,omitempty" xml:"DestNetType,omitempty"`
	SubscribeTopic           *string                                               `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	DtsInstanceID            *string                                               `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	RequestId                *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code                     *int32                                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Checkpoint               *int32                                                `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	Delay                    *int32                                                `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ExpireTime               *string                                               `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ErrCode                  *string                                               `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                  *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage               *string                                               `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ConsumptionClient        *string                                               `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	DbObject                 *string                                               `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	DynamicMessage           *string                                               `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ConsumptionCheckpoint    *string                                               `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	EtlCalculator            *string                                               `json:"EtlCalculator,omitempty" xml:"EtlCalculator,omitempty"`
	HttpStatusCode           *int32                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	BeginTimestamp           *string                                               `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	GroupId                  *string                                               `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	SynchronizationDirection *string                                               `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	DtsJobDirection          *string                                               `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	SourceEndpoint           *DescribeDtsJobDetailResponseBodySourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	DestinationEndpoint      *DescribeDtsJobDetailResponseBodyDestinationEndpoint  `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode            *DescribeDtsJobDetailResponseBodyMigrationMode        `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	SubscriptionHost         *DescribeDtsJobDetailResponseBodySubscriptionHost     `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
	SubscriptionDataType     *DescribeDtsJobDetailResponseBodySubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
}

func (s DescribeDtsJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBody) SetStatus(v string) *DescribeDtsJobDetailResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobName(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetFinishTime(v string) *DescribeDtsJobDetailResponseBody {
	s.FinishTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrorMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobId(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCreateTime(v string) *DescribeDtsJobDetailResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetPayType(v string) *DescribeDtsJobDetailResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetReserved(v string) *DescribeDtsJobDetailResponseBody {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDatabaseCount(v int32) *DescribeDtsJobDetailResponseBody {
	s.DatabaseCount = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobClass(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetEndTimestamp(v string) *DescribeDtsJobDetailResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetAppName(v string) *DescribeDtsJobDetailResponseBody {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDestNetType(v string) *DescribeDtsJobDetailResponseBody {
	s.DestNetType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscribeTopic(v string) *DescribeDtsJobDetailResponseBody {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsInstanceID(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetRequestId(v string) *DescribeDtsJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCode(v int32) *DescribeDtsJobDetailResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetCheckpoint(v int32) *DescribeDtsJobDetailResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDelay(v int32) *DescribeDtsJobDetailResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetExpireTime(v string) *DescribeDtsJobDetailResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrCode(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSuccess(v bool) *DescribeDtsJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetErrMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetConsumptionClient(v string) *DescribeDtsJobDetailResponseBody {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDbObject(v string) *DescribeDtsJobDetailResponseBody {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDynamicMessage(v string) *DescribeDtsJobDetailResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetConsumptionCheckpoint(v string) *DescribeDtsJobDetailResponseBody {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetEtlCalculator(v string) *DescribeDtsJobDetailResponseBody {
	s.EtlCalculator = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetHttpStatusCode(v int32) *DescribeDtsJobDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetBeginTimestamp(v string) *DescribeDtsJobDetailResponseBody {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetGroupId(v string) *DescribeDtsJobDetailResponseBody {
	s.GroupId = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSynchronizationDirection(v string) *DescribeDtsJobDetailResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDtsJobDirection(v string) *DescribeDtsJobDetailResponseBody {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSourceEndpoint(v *DescribeDtsJobDetailResponseBodySourceEndpoint) *DescribeDtsJobDetailResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetDestinationEndpoint(v *DescribeDtsJobDetailResponseBodyDestinationEndpoint) *DescribeDtsJobDetailResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetMigrationMode(v *DescribeDtsJobDetailResponseBodyMigrationMode) *DescribeDtsJobDetailResponseBody {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscriptionHost(v *DescribeDtsJobDetailResponseBodySubscriptionHost) *DescribeDtsJobDetailResponseBody {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeDtsJobDetailResponseBody) SetSubscriptionDataType(v *DescribeDtsJobDetailResponseBodySubscriptionDataType) *DescribeDtsJobDetailResponseBody {
	s.SubscriptionDataType = v
	return s
}

type DescribeDtsJobDetailResponseBodySourceEndpoint struct {
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	RoleName        *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	AliyunUid       *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetRoleName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.RoleName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetAliyunUid(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.AliyunUid = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySourceEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

type DescribeDtsJobDetailResponseBodyDestinationEndpoint struct {
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetRegion(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetIp(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetPort(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobDetailResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

type DescribeDtsJobDetailResponseBodyMigrationMode struct {
	DataExtractTransformLoad *bool `json:"DataExtractTransformLoad,omitempty" xml:"DataExtractTransformLoad,omitempty"`
	DataInitialization       *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization      *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization  *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodyMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodyMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataExtractTransformLoad(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataExtractTransformLoad = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodyMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobDetailResponseBodyMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeDtsJobDetailResponseBodySubscriptionHost struct {
	VpcHost     *string `json:"VpcHost,omitempty" xml:"VpcHost,omitempty"`
	PublicHost  *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubscriptionHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetVpcHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.VpcHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetPublicHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionHost) SetPrivateHost(v string) *DescribeDtsJobDetailResponseBodySubscriptionHost {
	s.PrivateHost = &v
	return s
}

type DescribeDtsJobDetailResponseBodySubscriptionDataType struct {
	Dml *bool `json:"Dml,omitempty" xml:"Dml,omitempty"`
	Ddl *bool `json:"Ddl,omitempty" xml:"Ddl,omitempty"`
}

func (s DescribeDtsJobDetailResponseBodySubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponseBodySubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) SetDml(v bool) *DescribeDtsJobDetailResponseBodySubscriptionDataType {
	s.Dml = &v
	return s
}

func (s *DescribeDtsJobDetailResponseBodySubscriptionDataType) SetDdl(v bool) *DescribeDtsJobDetailResponseBodySubscriptionDataType {
	s.Ddl = &v
	return s
}

type DescribeDtsJobDetailResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDtsJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDtsJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobDetailResponse) SetHeaders(v map[string]*string) *DescribeDtsJobDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsJobDetailResponse) SetBody(v *DescribeDtsJobDetailResponseBody) *DescribeDtsJobDetailResponse {
	s.Body = v
	return s
}

type DescribeDtsJobsRequest struct {
	Type           *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Params         *string `json:"Params,omitempty" xml:"Params,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	OrderDirection *string `json:"OrderDirection,omitempty" xml:"OrderDirection,omitempty"`
	OrderColumn    *string `json:"OrderColumn,omitempty" xml:"OrderColumn,omitempty"`
	Tags           *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	JobType        *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	GroupId        *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s DescribeDtsJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsRequest) SetType(v string) *DescribeDtsJobsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetParams(v string) *DescribeDtsJobsRequest {
	s.Params = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetRegion(v string) *DescribeDtsJobsRequest {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetStatus(v string) *DescribeDtsJobsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetOrderDirection(v string) *DescribeDtsJobsRequest {
	s.OrderDirection = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetOrderColumn(v string) *DescribeDtsJobsRequest {
	s.OrderColumn = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetTags(v string) *DescribeDtsJobsRequest {
	s.Tags = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetPageSize(v int32) *DescribeDtsJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetPageNumber(v int32) *DescribeDtsJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetJobType(v string) *DescribeDtsJobsRequest {
	s.JobType = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetRegionId(v string) *DescribeDtsJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobsRequest) SetGroupId(v string) *DescribeDtsJobsRequest {
	s.GroupId = &v
	return s
}

type DescribeDtsJobsResponseBody struct {
	HttpStatusCode   *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId        *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode          *string                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success          *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	PageRecordCount  *int32                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount *int32                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage       *string                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage   *string                                  `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	PageNumber       *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DynamicCode      *string                                  `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DtsJobList       []*DescribeDtsJobsResponseBodyDtsJobList `json:"DtsJobList,omitempty" xml:"DtsJobList,omitempty" type:"Repeated"`
}

func (s DescribeDtsJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBody) SetHttpStatusCode(v int32) *DescribeDtsJobsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetRequestId(v string) *DescribeDtsJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetErrCode(v string) *DescribeDtsJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetSuccess(v bool) *DescribeDtsJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetPageRecordCount(v int32) *DescribeDtsJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetTotalRecordCount(v int32) *DescribeDtsJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetErrMessage(v string) *DescribeDtsJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetDynamicMessage(v string) *DescribeDtsJobsResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetPageNumber(v int32) *DescribeDtsJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetDynamicCode(v string) *DescribeDtsJobsResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDtsJobsResponseBody) SetDtsJobList(v []*DescribeDtsJobsResponseBodyDtsJobList) *DescribeDtsJobsResponseBody {
	s.DtsJobList = v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobList struct {
	Status                        *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	DtsJobName                    *string                                                             `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	Delay                         *int32                                                              `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage                  *string                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                             `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	DtsJobId                      *string                                                             `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	CreateTime                    *string                                                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                       *string                                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reserved                      *string                                                             `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	ConsumptionClient             *string                                                             `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	DbObject                      *string                                                             `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	DtsJobClass                   *string                                                             `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	ConsumptionCheckpoint         *string                                                             `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	EndTimestamp                  *string                                                             `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	AppName                       *string                                                             `json:"AppName,omitempty" xml:"AppName,omitempty"`
	BeginTimestamp                *string                                                             `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	DtsInstanceID                 *string                                                             `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	DtsJobDirection               *string                                                             `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	Checkpoint                    *string                                                             `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	TagList                       []*DescribeDtsJobsResponseBodyDtsJobListTagList                     `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	DataInitializationStatus      *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DataEtlStatus                 *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus                 `json:"DataEtlStatus,omitempty" xml:"DataEtlStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode                 *DescribeDtsJobsResponseBodyDtsJobListMigrationMode                 `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	Performance                   *DescribeDtsJobsResponseBodyDtsJobListPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	ReverseJob                    *DescribeDtsJobsResponseBodyDtsJobListReverseJob                    `json:"ReverseJob,omitempty" xml:"ReverseJob,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	RetryState                    *DescribeDtsJobsResponseBodyDtsJobListRetryState                    `json:"RetryState,omitempty" xml:"RetryState,omitempty" type:"Struct"`
}

func (s DescribeDtsJobsResponseBodyDtsJobList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDelay(v int32) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetExpireTime(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetCreateTime(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPayType(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetReserved(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetConsumptionClient(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDbObject(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetConsumptionCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetEndTimestamp(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetAppName(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.AppName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetBeginTimestamp(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetTagList(v []*DescribeDtsJobsResponseBodyDtsJobListTagList) *DescribeDtsJobsResponseBodyDtsJobList {
	s.TagList = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDataEtlStatus(v *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DataEtlStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) *DescribeDtsJobsResponseBodyDtsJobList {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetMigrationMode(v *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) *DescribeDtsJobsResponseBodyDtsJobList {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPerformance(v *DescribeDtsJobsResponseBodyDtsJobListPerformance) *DescribeDtsJobsResponseBodyDtsJobList {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetReverseJob(v *DescribeDtsJobsResponseBodyDtsJobListReverseJob) *DescribeDtsJobsResponseBodyDtsJobList {
	s.ReverseJob = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) *DescribeDtsJobsResponseBodyDtsJobList {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobList {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobList) SetRetryState(v *DescribeDtsJobsResponseBodyDtsJobListRetryState) *DescribeDtsJobsResponseBodyDtsJobList {
	s.RetryState = v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListTagList struct {
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListTagList) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListTagList) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) SetTagValue(v string) *DescribeDtsJobsResponseBodyDtsJobListTagList {
	s.TagValue = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListTagList) SetTagKey(v string) *DescribeDtsJobsResponseBodyDtsJobListTagList {
	s.TagKey = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	NeedUpgrade  *bool   `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListDataEtlStatus {
	s.Progress = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint struct {
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListDestinationEndpoint {
	s.EngineName = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListPerformance struct {
	Rps  *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyDtsJobListPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyDtsJobListPerformance {
	s.Flow = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus struct {
	Status       *string                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string                                                      `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string                                                      `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Detail       []*DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatus {
	s.Detail = v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail struct {
	CheckResult          *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	CheckItem            *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	RepairMethod         *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	FailedReason         *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyDtsJobListPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJob struct {
	Status                        *string                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	DtsJobName                    *string                                                                       `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	Delay                         *int32                                                                        `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage                  *string                                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DtsJobId                      *string                                                                       `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ExpireTime                    *string                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime                    *string                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                       *string                                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Reserved                      *string                                                                       `json:"Reserved,omitempty" xml:"Reserved,omitempty"`
	DbObject                      *string                                                                       `json:"DbObject,omitempty" xml:"DbObject,omitempty"`
	DtsJobClass                   *string                                                                       `json:"DtsJobClass,omitempty" xml:"DtsJobClass,omitempty"`
	DtsInstanceID                 *string                                                                       `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	DtsJobDirection               *string                                                                       `json:"DtsJobDirection,omitempty" xml:"DtsJobDirection,omitempty"`
	Checkpoint                    *string                                                                       `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DataInitializationStatus      *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode                 *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode                 `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	Performance                   *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJob) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDelay(v int32) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Delay = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobId(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetExpireTime(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetCreateTime(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.CreateTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPayType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.PayType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetReserved(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Reserved = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDbObject(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DbObject = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobClass(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobClass = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDtsJobDirection(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DtsJobDirection = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetCheckpoint(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Checkpoint = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDataInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDataSynchronizationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetDestinationEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetMigrationMode(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.MigrationMode = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPerformance(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.Performance = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetPrecheckStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetSourceEndpoint(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJob) SetStructureInitializationStatus(v *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) *DescribeDtsJobsResponseBodyDtsJobListReverseJob {
	s.StructureInitializationStatus = v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	NeedUpgrade  *bool   `json:"NeedUpgrade,omitempty" xml:"NeedUpgrade,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetNeedUpgrade(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.NeedUpgrade = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.Progress = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint struct {
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobDestinationEndpoint {
	s.EngineName = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode struct {
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetDataInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetDataSynchronization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.DataSynchronization = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode) SetStructureInitialization(v bool) *DescribeDtsJobsResponseBodyDtsJobListReverseJobMigrationMode {
	s.StructureInitialization = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance struct {
	Rps  *string `json:"Rps,omitempty" xml:"Rps,omitempty"`
	Flow *string `json:"Flow,omitempty" xml:"Flow,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) SetRps(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance {
	s.Rps = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance) SetFlow(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPerformance {
	s.Flow = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus struct {
	Status       *string                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string                                                                `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string                                                                `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Detail       []*DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus) SetDetail(v []*DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatus {
	s.Detail = v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail struct {
	CheckResult          *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	CheckItem            *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	RepairMethod         *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	FailedReason         *string `json:"FailedReason,omitempty" xml:"FailedReason,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckResult(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItemDescription(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetCheckItem(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetRepairMethod(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.RepairMethod = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail) SetFailedReason(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobPrecheckStatusDetail {
	s.FailedReason = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint struct {
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobSourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListReverseJobStructureInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint struct {
	SslSolutionEnum *string `json:"SslSolutionEnum,omitempty" xml:"SslSolutionEnum,omitempty"`
	OracleSID       *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DatabaseName    *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	Ip              *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	InstanceID      *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port            *string `json:"Port,omitempty" xml:"Port,omitempty"`
	InstanceType    *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName        *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName      *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetSslSolutionEnum(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.SslSolutionEnum = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetOracleSID(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetRegion(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Region = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetDatabaseName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetIp(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Ip = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetInstanceID(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetPort(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetInstanceType(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetUserName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint) SetEngineName(v string) *DescribeDtsJobsResponseBodyDtsJobListSourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetStatus(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetPercent(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetErrorMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus) SetProgress(v string) *DescribeDtsJobsResponseBodyDtsJobListStructureInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeDtsJobsResponseBodyDtsJobListRetryState struct {
	RetryCount   *int32  `json:"RetryCount,omitempty" xml:"RetryCount,omitempty"`
	MaxRetryTime *int32  `json:"MaxRetryTime,omitempty" xml:"MaxRetryTime,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RetryTarget  *string `json:"RetryTarget,omitempty" xml:"RetryTarget,omitempty"`
	RetryTime    *int32  `json:"RetryTime,omitempty" xml:"RetryTime,omitempty"`
}

func (s DescribeDtsJobsResponseBodyDtsJobListRetryState) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponseBodyDtsJobListRetryState) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryCount(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryCount = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetMaxRetryTime(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.MaxRetryTime = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetErrMessage(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryTarget(v string) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryTarget = &v
	return s
}

func (s *DescribeDtsJobsResponseBodyDtsJobListRetryState) SetRetryTime(v int32) *DescribeDtsJobsResponseBodyDtsJobListRetryState {
	s.RetryTime = &v
	return s
}

type DescribeDtsJobsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDtsJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDtsJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDtsJobsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobsResponse) SetHeaders(v map[string]*string) *DescribeDtsJobsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDtsJobsResponse) SetBody(v *DescribeDtsJobsResponseBody) *DescribeDtsJobsResponse {
	s.Body = v
	return s
}

type DescribeEndpointSwitchStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeEndpointSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusRequest) SetRegionId(v string) *DescribeEndpointSwitchStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetTaskId(v string) *DescribeEndpointSwitchStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetClientToken(v string) *DescribeEndpointSwitchStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetOwnerId(v string) *DescribeEndpointSwitchStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetAccountId(v string) *DescribeEndpointSwitchStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeEndpointSwitchStatusResponseBody struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode      *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success      *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DescribeEndpointSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetStatus(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrorMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetRequestId(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrCode(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetSuccess(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrMessage = &v
	return s
}

type DescribeEndpointSwitchStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeEndpointSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeEndpointSwitchStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointSwitchStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusResponse) SetHeaders(v map[string]*string) *DescribeEndpointSwitchStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndpointSwitchStatusResponse) SetBody(v *DescribeEndpointSwitchStatusResponseBody) *DescribeEndpointSwitchStatusResponse {
	s.Body = v
	return s
}

type DescribeInitializationStatusRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum              *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	OwnerId              *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId            *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeInitializationStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusRequest) SetRegionId(v string) *DescribeInitializationStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetSynchronizationJobId(v string) *DescribeInitializationStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageSize(v int32) *DescribeInitializationStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetPageNum(v int32) *DescribeInitializationStatusRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetOwnerId(v string) *DescribeInitializationStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInitializationStatusRequest) SetAccountId(v string) *DescribeInitializationStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeInitializationStatusResponseBody struct {
	RequestId                      *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                        *string                                                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                        *string                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage                     *string                                                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

func (s *DescribeInitializationStatusResponseBody) SetRequestId(v string) *DescribeInitializationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrCode(v string) *DescribeInitializationStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetSuccess(v string) *DescribeInitializationStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrMessage(v string) *DescribeInitializationStatusResponseBody {
	s.ErrMessage = &v
	return s
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
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UsedTime               *string `json:"UsedTime,omitempty" xml:"UsedTime,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishRowNum           *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	TotalRowNum            *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetUsedTime(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.UsedTime = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TableName = &v
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

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataInitializationDetails) SetTotalRowNum(v string) *DescribeInitializationStatusResponseBodyDataInitializationDetails {
	s.TotalRowNum = &v
	return s
}

type DescribeInitializationStatusResponseBodyDataSynchronizationDetails struct {
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyDataSynchronizationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetTableName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.TableName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyDataSynchronizationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyDataSynchronizationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetails struct {
	Status                 *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string                                                                              `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	ObjectDefinition       *string                                                                              `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string                                                                              `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string                                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string                                                                              `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string                                                                              `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	Constraints            []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Repeated"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetails) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.ObjectName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetConstraints(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Constraints = v
	return s
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints struct {
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	ObjectDefinition       *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetStatus(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.Status = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetSourceOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectDefinition(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectType(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectType = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetErrorMessage(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetDestinationOwnerDBName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) SetObjectName(v string) *DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints {
	s.ObjectName = &v
	return s
}

type DescribeInitializationStatusResponse struct {
	Headers map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeInitializationStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeInitializationStatusResponse) SetBody(v *DescribeInitializationStatusResponseBody) *DescribeInitializationStatusResponse {
	s.Body = v
	return s
}

type DescribeJobMonitorRuleRequest struct {
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeJobMonitorRuleRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMonitorRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleRequest) SetDtsJobId(v string) *DescribeJobMonitorRuleRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeJobMonitorRuleRequest) SetRegionId(v string) *DescribeJobMonitorRuleRequest {
	s.RegionId = &v
	return s
}

type DescribeJobMonitorRuleResponseBody struct {
	HttpStatusCode *int32                                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DtsJobId       *string                                           `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ErrCode        *string                                           `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string                                           `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string                                           `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string                                           `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	MonitorRules   []*DescribeJobMonitorRuleResponseBodyMonitorRules `json:"MonitorRules,omitempty" xml:"MonitorRules,omitempty" type:"Repeated"`
}

func (s DescribeJobMonitorRuleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMonitorRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponseBody) SetHttpStatusCode(v int32) *DescribeJobMonitorRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetRequestId(v string) *DescribeJobMonitorRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetDtsJobId(v string) *DescribeJobMonitorRuleResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetErrCode(v string) *DescribeJobMonitorRuleResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetSuccess(v bool) *DescribeJobMonitorRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetErrMessage(v string) *DescribeJobMonitorRuleResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetCode(v string) *DescribeJobMonitorRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetDynamicMessage(v string) *DescribeJobMonitorRuleResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBody) SetMonitorRules(v []*DescribeJobMonitorRuleResponseBodyMonitorRules) *DescribeJobMonitorRuleResponseBody {
	s.MonitorRules = v
	return s
}

type DescribeJobMonitorRuleResponseBodyMonitorRules struct {
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
	DelayRuleTime *int64  `json:"DelayRuleTime,omitempty" xml:"DelayRuleTime,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	Phone         *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
}

func (s DescribeJobMonitorRuleResponseBodyMonitorRules) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMonitorRuleResponseBodyMonitorRules) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetType(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Type = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetDelayRuleTime(v int64) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.DelayRuleTime = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetState(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.State = &v
	return s
}

func (s *DescribeJobMonitorRuleResponseBodyMonitorRules) SetPhone(v string) *DescribeJobMonitorRuleResponseBodyMonitorRules {
	s.Phone = &v
	return s
}

type DescribeJobMonitorRuleResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeJobMonitorRuleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeJobMonitorRuleResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeJobMonitorRuleResponse) GoString() string {
	return s.String()
}

func (s *DescribeJobMonitorRuleResponse) SetHeaders(v map[string]*string) *DescribeJobMonitorRuleResponse {
	s.Headers = v
	return s
}

func (s *DescribeJobMonitorRuleResponse) SetBody(v *DescribeJobMonitorRuleResponseBody) *DescribeJobMonitorRuleResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobAlertRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeMigrationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertRequest) SetRegionId(v string) *DescribeMigrationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetMigrationJobId(v string) *DescribeMigrationJobAlertRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetClientToken(v string) *DescribeMigrationJobAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetOwnerId(v string) *DescribeMigrationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobAlertRequest) SetAccountId(v string) *DescribeMigrationJobAlertRequest {
	s.AccountId = &v
	return s
}

type DescribeMigrationJobAlertResponseBody struct {
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrCode          *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success          *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorAlertPhone  *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	ErrMessage       *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayAlertPhone  *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	MigrationJobId   *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
}

func (s DescribeMigrationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertResponseBody) SetRequestId(v string) *DescribeMigrationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetMigrationJobName(v string) *DescribeMigrationJobAlertResponseBody {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrCode(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetSuccess(v string) *DescribeMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrMessage(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobAlertResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

type DescribeMigrationJobAlertResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMigrationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeMigrationJobAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertResponse) SetHeaders(v map[string]*string) *DescribeMigrationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeMigrationJobAlertResponse) SetBody(v *DescribeMigrationJobAlertResponseBody) *DescribeMigrationJobAlertResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobDetailRequest struct {
	MigrationMode  *DescribeMigrationJobDetailRequestMigrationMode `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	RegionId       *string                                         `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize       *int32                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum        *int32                                          `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	MigrationJobId *string                                         `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ClientToken    *string                                         `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string                                         `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
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

func (s *DescribeMigrationJobDetailRequest) SetRegionId(v string) *DescribeMigrationJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageSize(v int32) *DescribeMigrationJobDetailRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetPageNum(v int32) *DescribeMigrationJobDetailRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetMigrationJobId(v string) *DescribeMigrationJobDetailRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetClientToken(v string) *DescribeMigrationJobDetailRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetOwnerId(v string) *DescribeMigrationJobDetailRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobDetailRequest) SetAccountId(v string) *DescribeMigrationJobDetailRequest {
	s.AccountId = &v
	return s
}

type DescribeMigrationJobDetailRequestMigrationMode struct {
	StructureInitialization *bool `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	DataInitialization      *bool `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	DataSynchronization     *bool `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty"`
}

func (s DescribeMigrationJobDetailRequestMigrationMode) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailRequestMigrationMode) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataInitialization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataInitialization = &v
	return s
}

func (s *DescribeMigrationJobDetailRequestMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobDetailRequestMigrationMode {
	s.DataSynchronization = &v
	return s
}

type DescribeMigrationJobDetailResponseBody struct {
	RequestId                         *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                           *string                                                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PageRecordCount                   *int32                                                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Success                           *string                                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalRecordCount                  *int64                                                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage                        *string                                                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber                        *int32                                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DataInitializationDetailList      *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList      `json:"DataInitializationDetailList,omitempty" xml:"DataInitializationDetailList,omitempty" type:"Struct"`
	DataSynchronizationDetailList     *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList     `json:"DataSynchronizationDetailList,omitempty" xml:"DataSynchronizationDetailList,omitempty" type:"Struct"`
	StructureInitializationDetailList *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList `json:"StructureInitializationDetailList,omitempty" xml:"StructureInitializationDetailList,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBody) SetRequestId(v string) *DescribeMigrationJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetErrCode(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetSuccess(v string) *DescribeMigrationJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobDetailResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetErrMessage(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageNumber(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataSynchronizationDetailList(v *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataSynchronizationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetStructureInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.StructureInitializationDetailList = v
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
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FinishRowNum           *string `json:"FinishRowNum,omitempty" xml:"FinishRowNum,omitempty"`
	MigrationTime          *string `json:"MigrationTime,omitempty" xml:"MigrationTime,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	TotalRowNum            *string `json:"TotalRowNum,omitempty" xml:"TotalRowNum,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) GoString() string {
	return s.String()
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

func (s *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataInitializationDetailListDataInitializationDetail {
	s.DestinationOwnerDBName = &v
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
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	TableName              *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetTableName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.TableName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailListDataSynchronizationDetail {
	s.DestinationOwnerDBName = &v
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
	Status                 *string                                                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string                                                                                                             `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	ObjectDefinition       *string                                                                                                             `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string                                                                                                             `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string                                                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string                                                                                                             `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string                                                                                                             `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	ConstraintList         *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList `json:"ConstraintList,omitempty" xml:"ConstraintList,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetConstraintList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ConstraintList = v
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
	Status                 *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	ObjectDefinition       *string `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetStatus(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetSourceOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.SourceOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectDefinition(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectDefinition = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectType(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectType = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetErrorMessage(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetDestinationOwnerDBName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.DestinationOwnerDBName = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail) SetObjectName(v string) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintListStructureInitializationDetail {
	s.ObjectName = &v
	return s
}

type DescribeMigrationJobDetailResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMigrationJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMigrationJobDetailResponse) SetBody(v *DescribeMigrationJobDetailResponseBody) *DescribeMigrationJobDetailResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobsRequest struct {
	RegionId         *string                            `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize         *int32                             `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum          *int32                             `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	MigrationJobName *string                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	OwnerId          *string                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId        *string                            `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Tag              []*DescribeMigrationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsRequest) SetRegionId(v string) *DescribeMigrationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageSize(v int32) *DescribeMigrationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetPageNum(v int32) *DescribeMigrationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetMigrationJobName(v string) *DescribeMigrationJobsRequest {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetOwnerId(v string) *DescribeMigrationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetAccountId(v string) *DescribeMigrationJobsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeMigrationJobsRequest) SetTag(v []*DescribeMigrationJobsRequestTag) *DescribeMigrationJobsRequest {
	s.Tag = v
	return s
}

type DescribeMigrationJobsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMigrationJobsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsRequestTag) SetKey(v string) *DescribeMigrationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeMigrationJobsRequestTag) SetValue(v string) *DescribeMigrationJobsRequestTag {
	s.Value = &v
	return s
}

type DescribeMigrationJobsResponseBody struct {
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode          *string                                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Success          *string                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalRecordCount *int64                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage       *string                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	MigrationJobs    *DescribeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBody) SetRequestId(v string) *DescribeMigrationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetErrCode(v string) *DescribeMigrationJobsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetSuccess(v string) *DescribeMigrationJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetErrMessage(v string) *DescribeMigrationJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageNumber(v int32) *DescribeMigrationJobsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetMigrationJobs(v *DescribeMigrationJobsResponseBodyMigrationJobs) *DescribeMigrationJobsResponseBody {
	s.MigrationJobs = v
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
	MigrationJobName        *string                                                                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus      *string                                                                            `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	PayType                 *string                                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	MigrationJobClass       *string                                                                            `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	InstanceCreateTime      *string                                                                            `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	MigrationJobID          *string                                                                            `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
	JobCreateTime           *string                                                                            `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	MigrationObject         *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject         `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	Tags                    *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	DataInitialization      *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization      `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty" type:"Struct"`
	DataSynchronization     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization     `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	DestinationEndpoint     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint     `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode           *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode           `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobClass(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobClass = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetInstanceCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetJobCreateTime(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetTags(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Tags = v
	return s
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
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
	WholeDatabase *string                                                                                              `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
	TableList     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetTableList(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.TableList = v
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

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags struct {
	Tag []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) SetTag(v []*DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags {
	s.Tag = v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetKey(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag) SetValue(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTagsTag {
	s.Value = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.Progress = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint struct {
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint {
	s.EngineName = &v
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

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck struct {
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck {
	s.Percent = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint struct {
	OracleSID    *string `json:"OracleSID,omitempty" xml:"OracleSID,omitempty"`
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetOracleSID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.OracleSID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.DatabaseName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceID(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceID = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetPort(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetIP(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetInstanceType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetUserName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) SetEngineName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetProgress(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.Progress = &v
	return s
}

type DescribeMigrationJobsResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMigrationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMigrationJobsResponse) SetBody(v *DescribeMigrationJobsResponseBody) *DescribeMigrationJobsResponse {
	s.Body = v
	return s
}

type DescribeMigrationJobStatusRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeMigrationJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusRequest) SetRegionId(v string) *DescribeMigrationJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetMigrationJobId(v string) *DescribeMigrationJobStatusRequest {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetClientToken(v string) *DescribeMigrationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetOwnerId(v string) *DescribeMigrationJobStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeMigrationJobStatusRequest) SetAccountId(v string) *DescribeMigrationJobStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeMigrationJobStatusResponseBody struct {
	RequestId                     *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MigrationJobName              *string                                                              `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationJobStatus            *string                                                              `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	ErrCode                       *string                                                              `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                       *string                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	PayType                       *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ErrMessage                    *string                                                              `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	MigrationObject               *string                                                              `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	MigrationJobId                *string                                                              `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	TaskId                        *string                                                              `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	MigrationJobClass             *string                                                              `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	DataInitializationStatus      *DescribeMigrationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeMigrationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode                 *DescribeMigrationJobStatusResponseBodyMigrationMode                 `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
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

func (s *DescribeMigrationJobStatusResponseBody) SetRequestId(v string) *DescribeMigrationJobStatusResponseBody {
	s.RequestId = &v
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

func (s *DescribeMigrationJobStatusResponseBody) SetErrCode(v string) *DescribeMigrationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetSuccess(v string) *DescribeMigrationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPayType(v string) *DescribeMigrationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetErrMessage(v string) *DescribeMigrationJobStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationObject(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationObject = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetTaskId(v string) *DescribeMigrationJobStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobClass(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobClass = &v
	return s
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

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationMode(v *DescribeMigrationJobStatusResponseBodyMigrationMode) *DescribeMigrationJobStatusResponseBody {
	s.MigrationMode = v
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
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Checkpoint   *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
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

func (s *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyDestinationEndpoint struct {
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	OracleSID    *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
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

func (s *DescribeMigrationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
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
	Status  *string                                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent *string                                                     `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Detail  *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyPrecheckStatus) SetDetail(v *DescribeMigrationJobStatusResponseBodyPrecheckStatusDetail) *DescribeMigrationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
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
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	OracleSID    *string `json:"oracleSID,omitempty" xml:"oracleSID,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
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

func (s *DescribeMigrationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeMigrationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeMigrationJobStatusResponseBodyStructureInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeMigrationJobStatusResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeMigrationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeMigrationJobStatusResponse) SetBody(v *DescribeMigrationJobStatusResponseBody) *DescribeMigrationJobStatusResponse {
	s.Body = v
	return s
}

type DescribePreCheckStatusRequest struct {
	DtsJobId    *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	PageNo      *string `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	JobCode     *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	StructType  *string `json:"StructType,omitempty" xml:"StructType,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	StructPhase *string `json:"StructPhase,omitempty" xml:"StructPhase,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePreCheckStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusRequest) SetDtsJobId(v string) *DescribePreCheckStatusRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetPageNo(v string) *DescribePreCheckStatusRequest {
	s.PageNo = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetPageSize(v string) *DescribePreCheckStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetJobCode(v string) *DescribePreCheckStatusRequest {
	s.JobCode = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetStructType(v string) *DescribePreCheckStatusRequest {
	s.StructType = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetRegionId(v string) *DescribePreCheckStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetStructPhase(v string) *DescribePreCheckStatusRequest {
	s.StructPhase = &v
	return s
}

func (s *DescribePreCheckStatusRequest) SetName(v string) *DescribePreCheckStatusRequest {
	s.Name = &v
	return s
}

type DescribePreCheckStatusResponseBody struct {
	State                   *string                                                      `json:"State,omitempty" xml:"State,omitempty"`
	Success                 *bool                                                        `json:"Success,omitempty" xml:"Success,omitempty"`
	PageRecordCount         *int64                                                       `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Total                   *int32                                                       `json:"Total,omitempty" xml:"Total,omitempty"`
	HttpStatusCode          *int32                                                       `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId               *string                                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	JobName                 *string                                                      `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobId                   *string                                                      `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ErrorItem               *int32                                                       `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	TotalRecordCount        *int64                                                       `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	Code                    *string                                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	PageNumber              *int64                                                       `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	JobProgress             []*DescribePreCheckStatusResponseBodyJobProgress             `json:"JobProgress,omitempty" xml:"JobProgress,omitempty" type:"Repeated"`
	SubDistributedJobStatus []*DescribePreCheckStatusResponseBodySubDistributedJobStatus `json:"SubDistributedJobStatus,omitempty" xml:"SubDistributedJobStatus,omitempty" type:"Repeated"`
}

func (s DescribePreCheckStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBody) SetState(v string) *DescribePreCheckStatusResponseBody {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetSuccess(v bool) *DescribePreCheckStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetPageRecordCount(v int64) *DescribePreCheckStatusResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetTotal(v int32) *DescribePreCheckStatusResponseBody {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetHttpStatusCode(v int32) *DescribePreCheckStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetRequestId(v string) *DescribePreCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobName(v string) *DescribePreCheckStatusResponseBody {
	s.JobName = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobId(v string) *DescribePreCheckStatusResponseBody {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetErrorItem(v int32) *DescribePreCheckStatusResponseBody {
	s.ErrorItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetTotalRecordCount(v int64) *DescribePreCheckStatusResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetCode(v string) *DescribePreCheckStatusResponseBody {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetPageNumber(v int64) *DescribePreCheckStatusResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetJobProgress(v []*DescribePreCheckStatusResponseBodyJobProgress) *DescribePreCheckStatusResponseBody {
	s.JobProgress = v
	return s
}

func (s *DescribePreCheckStatusResponseBody) SetSubDistributedJobStatus(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatus) *DescribePreCheckStatusResponseBody {
	s.SubDistributedJobStatus = v
	return s
}

type DescribePreCheckStatusResponseBodyJobProgress struct {
	Skip         *bool                                                `json:"Skip,omitempty" xml:"Skip,omitempty"`
	FinishTime   *string                                              `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	IgnoreFlag   *string                                              `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	DelaySeconds *int32                                               `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	DdlSql       *string                                              `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	State        *string                                              `json:"State,omitempty" xml:"State,omitempty"`
	BootTime     *string                                              `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	Item         *string                                              `json:"Item,omitempty" xml:"Item,omitempty"`
	Sub          *string                                              `json:"Sub,omitempty" xml:"Sub,omitempty"`
	Total        *int32                                               `json:"Total,omitempty" xml:"Total,omitempty"`
	TargetNames  *string                                              `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	CanSkip      *bool                                                `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	Names        *string                                              `json:"Names,omitempty" xml:"Names,omitempty"`
	ErrDetail    *string                                              `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	DiffRow      *int64                                               `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	JobId        *string                                              `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SourceSchema *string                                              `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	DestSchema   *string                                              `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	ParentObj    *string                                              `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	ErrMsg       *string                                              `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	OrderNum     *int32                                               `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	RepairMethod *string                                              `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	Id           *string                                              `json:"Id,omitempty" xml:"Id,omitempty"`
	Logs         []*DescribePreCheckStatusResponseBodyJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s DescribePreCheckStatusResponseBodyJobProgress) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodyJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodyJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodyJobProgressLogs) *DescribePreCheckStatusResponseBodyJobProgress {
	s.Logs = v
	return s
}

type DescribePreCheckStatusResponseBodyJobProgressLogs struct {
	ErrData  *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	ErrMsg   *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	ErrType  *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
}

func (s DescribePreCheckStatusResponseBodyJobProgressLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodyJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodyJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodyJobProgressLogs {
	s.ErrType = &v
	return s
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatus struct {
	JobName     *string                                                                 `json:"JobName,omitempty" xml:"JobName,omitempty"`
	State       *string                                                                 `json:"State,omitempty" xml:"State,omitempty"`
	JobId       *string                                                                 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ErrorItem   *int32                                                                  `json:"ErrorItem,omitempty" xml:"ErrorItem,omitempty"`
	Code        *string                                                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Total       *int32                                                                  `json:"Total,omitempty" xml:"Total,omitempty"`
	JobProgress []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress `json:"JobProgress,omitempty" xml:"JobProgress,omitempty" type:"Repeated"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatus) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobName(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobName = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetState(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetErrorItem(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.ErrorItem = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetCode(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.Code = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetTotal(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatus) SetJobProgress(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) *DescribePreCheckStatusResponseBodySubDistributedJobStatus {
	s.JobProgress = v
	return s
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress struct {
	Skip         *bool                                                                       `json:"Skip,omitempty" xml:"Skip,omitempty"`
	FinishTime   *string                                                                     `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	IgnoreFlag   *string                                                                     `json:"IgnoreFlag,omitempty" xml:"IgnoreFlag,omitempty"`
	DelaySeconds *int32                                                                      `json:"DelaySeconds,omitempty" xml:"DelaySeconds,omitempty"`
	DdlSql       *string                                                                     `json:"DdlSql,omitempty" xml:"DdlSql,omitempty"`
	State        *string                                                                     `json:"State,omitempty" xml:"State,omitempty"`
	BootTime     *string                                                                     `json:"BootTime,omitempty" xml:"BootTime,omitempty"`
	Item         *string                                                                     `json:"Item,omitempty" xml:"Item,omitempty"`
	Sub          *string                                                                     `json:"Sub,omitempty" xml:"Sub,omitempty"`
	Total        *int32                                                                      `json:"Total,omitempty" xml:"Total,omitempty"`
	TargetNames  *string                                                                     `json:"TargetNames,omitempty" xml:"TargetNames,omitempty"`
	CanSkip      *bool                                                                       `json:"CanSkip,omitempty" xml:"CanSkip,omitempty"`
	Names        *string                                                                     `json:"Names,omitempty" xml:"Names,omitempty"`
	ErrDetail    *string                                                                     `json:"ErrDetail,omitempty" xml:"ErrDetail,omitempty"`
	DiffRow      *int64                                                                      `json:"DiffRow,omitempty" xml:"DiffRow,omitempty"`
	JobId        *string                                                                     `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SourceSchema *string                                                                     `json:"SourceSchema,omitempty" xml:"SourceSchema,omitempty"`
	DestSchema   *string                                                                     `json:"DestSchema,omitempty" xml:"DestSchema,omitempty"`
	ParentObj    *string                                                                     `json:"ParentObj,omitempty" xml:"ParentObj,omitempty"`
	ErrMsg       *string                                                                     `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	OrderNum     *int32                                                                      `json:"OrderNum,omitempty" xml:"OrderNum,omitempty"`
	RepairMethod *string                                                                     `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
	Id           *string                                                                     `json:"Id,omitempty" xml:"Id,omitempty"`
	Logs         []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSkip(v bool) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Skip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetFinishTime(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.FinishTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetIgnoreFlag(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.IgnoreFlag = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDelaySeconds(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DelaySeconds = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDdlSql(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DdlSql = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetState(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.State = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetBootTime(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.BootTime = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetItem(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Item = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSub(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Sub = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetTotal(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Total = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetTargetNames(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.TargetNames = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetCanSkip(v bool) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.CanSkip = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetNames(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Names = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetErrDetail(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ErrDetail = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDiffRow(v int64) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DiffRow = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetJobId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.JobId = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetSourceSchema(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.SourceSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetDestSchema(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.DestSchema = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetParentObj(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ParentObj = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetErrMsg(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetOrderNum(v int32) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.OrderNum = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetRepairMethod(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.RepairMethod = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetId(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Id = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress) SetLogs(v []*DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgress {
	s.Logs = v
	return s
}

type DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs struct {
	ErrData  *string `json:"ErrData,omitempty" xml:"ErrData,omitempty"`
	ErrMsg   *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	LogLevel *string `json:"LogLevel,omitempty" xml:"LogLevel,omitempty"`
	ErrType  *string `json:"ErrType,omitempty" xml:"ErrType,omitempty"`
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrData(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrData = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrMsg(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrMsg = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetLogLevel(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.LogLevel = &v
	return s
}

func (s *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs) SetErrType(v string) *DescribePreCheckStatusResponseBodySubDistributedJobStatusJobProgressLogs {
	s.ErrType = &v
	return s
}

type DescribePreCheckStatusResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribePreCheckStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribePreCheckStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribePreCheckStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckStatusResponse) SetHeaders(v map[string]*string) *DescribePreCheckStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckStatusResponse) SetBody(v *DescribePreCheckStatusResponseBody) *DescribePreCheckStatusResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInstanceAlertRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ClientToken            *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetRegionId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetClientToken(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetOwnerId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertRequest) SetAccountId(v string) *DescribeSubscriptionInstanceAlertRequest {
	s.AccountId = &v
	return s
}

type DescribeSubscriptionInstanceAlertResponseBody struct {
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrCode                  *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	ErrMessage               *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	SubscriptionInstanceID   *string `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetRequestId(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

type DescribeSubscriptionInstanceAlertResponse struct {
	Headers map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubscriptionInstanceAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionInstanceAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionInstanceAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponse) SetBody(v *DescribeSubscriptionInstanceAlertResponseBody) *DescribeSubscriptionInstanceAlertResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInstancesRequest struct {
	RegionId                 *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize                 *int32                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum                  *int32                                     `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	SubscriptionInstanceName *string                                    `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	ClientToken              *string                                    `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string                                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string                                    `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Tag                      []*DescribeSubscriptionInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesRequest) SetRegionId(v string) *DescribeSubscriptionInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageSize(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetPageNum(v int32) *DescribeSubscriptionInstancesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesRequest {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetClientToken(v string) *DescribeSubscriptionInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetOwnerId(v string) *DescribeSubscriptionInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetAccountId(v string) *DescribeSubscriptionInstancesRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequest) SetTag(v []*DescribeSubscriptionInstancesRequestTag) *DescribeSubscriptionInstancesRequest {
	s.Tag = v
	return s
}

type DescribeSubscriptionInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSubscriptionInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesRequestTag) SetKey(v string) *DescribeSubscriptionInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionInstancesRequestTag) SetValue(v string) *DescribeSubscriptionInstancesRequestTag {
	s.Value = &v
	return s
}

type DescribeSubscriptionInstancesResponseBody struct {
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode               *string                                                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PageRecordCount       *int32                                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Success               *string                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalRecordCount      *int64                                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage            *string                                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SubscriptionInstances *DescribeSubscriptionInstancesResponseBodySubscriptionInstances `json:"SubscriptionInstances,omitempty" xml:"SubscriptionInstances,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBody) SetRequestId(v string) *DescribeSubscriptionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrCode(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSuccess(v string) *DescribeSubscriptionInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeSubscriptionInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSubscriptionInstances(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) *DescribeSubscriptionInstancesResponseBody {
	s.SubscriptionInstances = v
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
	Status                   *string                                                                                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage             *string                                                                                                 `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	PayType                  *string                                                                                                 `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ConsumptionClient        *string                                                                                                 `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	ConsumptionCheckpoint    *string                                                                                                 `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	EndTimestamp             *string                                                                                                 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	InstanceCreateTime       *string                                                                                                 `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	BeginTimestamp           *string                                                                                                 `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	SubscribeTopic           *string                                                                                                 `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	SubscriptionInstanceName *string                                                                                                 `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionInstanceID   *string                                                                                                 `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	JobCreateTime            *string                                                                                                 `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	SubscriptionObject       *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
	Tags                     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	SourceEndpoint           *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionHost         *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost     `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetStatus(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Status = &v
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

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionClient(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetEndTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetInstanceCreateTime(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetBeginTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscribeTopic(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetJobCreateTime(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionObject(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionObject = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetTags(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Tags = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSourceEndpoint(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionDataType(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionHost(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionHost = v
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
	WholeDatabase *string                                                                                                                         `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
	TableList     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
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

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetTableList(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.TableList = v
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

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags struct {
	Tag []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) SetTag(v []*DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags {
	s.Tag = v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) SetKey(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag) SetValue(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTagsTag {
	s.Value = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceType(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint) SetInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint {
	s.InstanceID = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType struct {
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDML(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DML = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType) SetDDL(v bool) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType {
	s.DDL = &v
	return s
}

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost struct {
	VPCHost     *string `json:"VPCHost,omitempty" xml:"VPCHost,omitempty"`
	PublicHost  *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetVPCHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.VPCHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPublicHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPrivateHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PrivateHost = &v
	return s
}

type DescribeSubscriptionInstancesResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubscriptionInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSubscriptionInstancesResponse) SetBody(v *DescribeSubscriptionInstancesResponseBody) *DescribeSubscriptionInstancesResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionInstanceStatusRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetRegionId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetSubscriptionInstanceId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetOwnerId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusRequest) SetAccountId(v string) *DescribeSubscriptionInstanceStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBody struct {
	Status                   *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage             *string                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrCode                  *string                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                  *string                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage               *string                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PayType                  *string                                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ConsumptionClient        *string                                                             `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	ConsumptionCheckpoint    *string                                                             `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	EndTimestamp             *string                                                             `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	BeginTimestamp           *string                                                             `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubscribeTopic           *string                                                             `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	SubscriptionInstanceName *string                                                             `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionInstanceID   *string                                                             `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	TaskId                   *string                                                             `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	SourceEndpoint           *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType     *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionHost         *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost     `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
	SubscriptionObject       *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstanceStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetStatus(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrorMessage(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrCode(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSuccess(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetPayType(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetConsumptionClient(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetEndTimestamp(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetBeginTimestamp(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.BeginTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetRequestId(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscribeTopic(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetTaskId(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSourceEndpoint(v *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionDataType(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionDataType = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionHost(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionObject(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionObject = v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint struct {
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceID   *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint) SetInstanceID(v string) *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint {
	s.InstanceID = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType struct {
	DML *bool `json:"DML,omitempty" xml:"DML,omitempty"`
	DDL *bool `json:"DDL,omitempty" xml:"DDL,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) SetDML(v bool) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType {
	s.DML = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType) SetDDL(v bool) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType {
	s.DDL = &v
	return s
}

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost struct {
	VPCHost     *string `json:"VPCHost,omitempty" xml:"VPCHost,omitempty"`
	PublicHost  *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetVPCHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.VPCHost = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetPublicHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetPrivateHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.PrivateHost = &v
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
	WholeDatabase *string                                                                                     `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
	TableList     *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
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

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject) SetTableList(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObjectTableList) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObjectSynchronousObject {
	s.TableList = v
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
	Headers map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubscriptionInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSubscriptionInstanceStatusResponse) SetBody(v *DescribeSubscriptionInstanceStatusResponseBody) *DescribeSubscriptionInstanceStatusResponse {
	s.Body = v
	return s
}

type DescribeSubscriptionMetaRequest struct {
	Sid                *string                `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DtsInstanceId      *string                `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SubMigrationJobIds map[string]interface{} `json:"SubMigrationJobIds,omitempty" xml:"SubMigrationJobIds,omitempty"`
	Topics             map[string]interface{} `json:"Topics,omitempty" xml:"Topics,omitempty"`
	RegionId           *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSubscriptionMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionMetaRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaRequest) SetSid(v string) *DescribeSubscriptionMetaRequest {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetDtsInstanceId(v string) *DescribeSubscriptionMetaRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetSubMigrationJobIds(v map[string]interface{}) *DescribeSubscriptionMetaRequest {
	s.SubMigrationJobIds = v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetTopics(v map[string]interface{}) *DescribeSubscriptionMetaRequest {
	s.Topics = v
	return s
}

func (s *DescribeSubscriptionMetaRequest) SetRegionId(v string) *DescribeSubscriptionMetaRequest {
	s.RegionId = &v
	return s
}

type DescribeSubscriptionMetaShrinkRequest struct {
	Sid                      *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SubMigrationJobIdsShrink *string `json:"SubMigrationJobIds,omitempty" xml:"SubMigrationJobIds,omitempty"`
	TopicsShrink             *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSubscriptionMetaShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionMetaShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetSid(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetDtsInstanceId(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetSubMigrationJobIdsShrink(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.SubMigrationJobIdsShrink = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetTopicsShrink(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *DescribeSubscriptionMetaShrinkRequest) SetRegionId(v string) *DescribeSubscriptionMetaShrinkRequest {
	s.RegionId = &v
	return s
}

type DescribeSubscriptionMetaResponseBody struct {
	HttpStatusCode       *string                                                     `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId            *string                                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode              *string                                                     `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success              *string                                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage           *string                                                     `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SubscriptionMetaList []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList `json:"SubscriptionMetaList,omitempty" xml:"SubscriptionMetaList,omitempty" type:"Repeated"`
}

func (s DescribeSubscriptionMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponseBody) SetHttpStatusCode(v string) *DescribeSubscriptionMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetRequestId(v string) *DescribeSubscriptionMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetErrCode(v string) *DescribeSubscriptionMetaResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetSuccess(v string) *DescribeSubscriptionMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetErrMessage(v string) *DescribeSubscriptionMetaResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBody) SetSubscriptionMetaList(v []*DescribeSubscriptionMetaResponseBodySubscriptionMetaList) *DescribeSubscriptionMetaResponseBody {
	s.SubscriptionMetaList = v
	return s
}

type DescribeSubscriptionMetaResponseBodySubscriptionMetaList struct {
	Sid        *string `json:"Sid,omitempty" xml:"Sid,omitempty"`
	Checkpoint *int64  `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DProxyUrl  *string `json:"DProxyUrl,omitempty" xml:"DProxyUrl,omitempty"`
	DBList     *string `json:"DBList,omitempty" xml:"DBList,omitempty"`
	Topic      *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s DescribeSubscriptionMetaResponseBodySubscriptionMetaList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionMetaResponseBodySubscriptionMetaList) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetSid(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Sid = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetCheckpoint(v int64) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetDProxyUrl(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.DProxyUrl = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetDBList(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.DBList = &v
	return s
}

func (s *DescribeSubscriptionMetaResponseBodySubscriptionMetaList) SetTopic(v string) *DescribeSubscriptionMetaResponseBodySubscriptionMetaList {
	s.Topic = &v
	return s
}

type DescribeSubscriptionMetaResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSubscriptionMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSubscriptionMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionMetaResponse) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionMetaResponse) SetHeaders(v map[string]*string) *DescribeSubscriptionMetaResponse {
	s.Headers = v
	return s
}

func (s *DescribeSubscriptionMetaResponse) SetBody(v *DescribeSubscriptionMetaResponseBody) *DescribeSubscriptionMetaResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobAlertRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSynchronizationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertRequest) SetRegionId(v string) *DescribeSynchronizationJobAlertRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetClientToken(v string) *DescribeSynchronizationJobAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetOwnerId(v string) *DescribeSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetAccountId(v string) *DescribeSynchronizationJobAlertRequest {
	s.AccountId = &v
	return s
}

type DescribeSynchronizationJobAlertResponseBody struct {
	SynchronizationJobName   *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrCode                  *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	ErrMessage               *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s DescribeSynchronizationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrCode(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSuccess(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetRequestId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationDirection = &v
	return s
}

type DescribeSynchronizationJobAlertResponse struct {
	Headers map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationJobAlertResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobAlertResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobAlertResponse) SetBody(v *DescribeSynchronizationJobAlertResponseBody) *DescribeSynchronizationJobAlertResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobReplicatorCompareRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetRegionId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetClientToken(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetOwnerId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareRequest) SetAccountId(v string) *DescribeSynchronizationJobReplicatorCompareRequest {
	s.AccountId = &v
	return s
}

type DescribeSynchronizationJobReplicatorCompareResponseBody struct {
	SynchronizationReplicatorCompareEnable *bool   `json:"SynchronizationReplicatorCompareEnable,omitempty" xml:"SynchronizationReplicatorCompareEnable,omitempty"`
	RequestId                              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                                *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                                *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage                             *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSynchronizationReplicatorCompareEnable(v bool) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.SynchronizationReplicatorCompareEnable = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetRequestId(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

type DescribeSynchronizationJobReplicatorCompareResponse struct {
	Headers map[string]*string                                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationJobReplicatorCompareResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationJobReplicatorCompareResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobReplicatorCompareResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponse) SetBody(v *DescribeSynchronizationJobReplicatorCompareResponseBody) *DescribeSynchronizationJobReplicatorCompareResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobsRequest struct {
	RegionId               *string                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	PageSize               *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PageNum                *int32                                   `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	SynchronizationJobName *string                                  `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	ClientToken            *string                                  `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                *string                                  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string                                  `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Tag                    []*DescribeSynchronizationJobsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsRequest) SetRegionId(v string) *DescribeSynchronizationJobsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageSize(v int32) *DescribeSynchronizationJobsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetPageNum(v int32) *DescribeSynchronizationJobsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsRequest {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetClientToken(v string) *DescribeSynchronizationJobsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetOwnerId(v string) *DescribeSynchronizationJobsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetAccountId(v string) *DescribeSynchronizationJobsRequest {
	s.AccountId = &v
	return s
}

func (s *DescribeSynchronizationJobsRequest) SetTag(v []*DescribeSynchronizationJobsRequestTag) *DescribeSynchronizationJobsRequest {
	s.Tag = v
	return s
}

type DescribeSynchronizationJobsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSynchronizationJobsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsRequestTag) SetKey(v string) *DescribeSynchronizationJobsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSynchronizationJobsRequestTag) SetValue(v string) *DescribeSynchronizationJobsRequestTag {
	s.Value = &v
	return s
}

type DescribeSynchronizationJobsResponseBody struct {
	PageNumber               *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageRecordCount          *int32                                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	TotalRecordCount         *int64                                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	SynchronizationInstances []*DescribeSynchronizationJobsResponseBodySynchronizationInstances `json:"SynchronizationInstances,omitempty" xml:"SynchronizationInstances,omitempty" type:"Repeated"`
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

func (s *DescribeSynchronizationJobsResponseBody) SetRequestId(v string) *DescribeSynchronizationJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetSynchronizationInstances(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstances) *DescribeSynchronizationJobsResponseBody {
	s.SynchronizationInstances = v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstances struct {
	DataInitialization            *string                                                                                       `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	Status                        *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	SynchronizationJobName        *string                                                                                       `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	Delay                         *string                                                                                       `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage                  *string                                                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	CreateTime                    *string                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                       *string                                                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	StructureInitialization       *string                                                                                       `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	SynchronizationJobClass       *string                                                                                       `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	InstanceCreateTime            *string                                                                                       `json:"InstanceCreateTime,omitempty" xml:"InstanceCreateTime,omitempty"`
	SynchronizationJobId          *string                                                                                       `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection      *string                                                                                       `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	JobCreateTime                 *string                                                                                       `json:"JobCreateTime,omitempty" xml:"JobCreateTime,omitempty"`
	Tags                          []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	SynchronizationObjects        []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
	DataInitializationStatus      *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	Performance                   *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Delay = &v
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPayType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetInstanceCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.InstanceCreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationDirection(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetJobCreateTime(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.JobCreateTime = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetTags(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Tags = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationObjects(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationObjects = v
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDestinationEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DestinationEndpoint = v
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitializationStatus = v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) SetKey(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags {
	s.Key = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) SetValue(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags {
	s.Value = &v
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

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetPort(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetUserName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.EngineName = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance struct {
	RPS  *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetRPS(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.RPS = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) SetFLOW(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance {
	s.FLOW = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus struct {
	Status  *string                                                                                `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent *string                                                                                `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Detail  []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatusDetail) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus {
	s.Detail = v
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
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetPort(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetUserName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeSynchronizationJobsResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSynchronizationJobsResponse) SetBody(v *DescribeSynchronizationJobsResponseBody) *DescribeSynchronizationJobsResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobStatusRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSynchronizationJobStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusRequest) SetRegionId(v string) *DescribeSynchronizationJobStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetClientToken(v string) *DescribeSynchronizationJobStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetOwnerId(v string) *DescribeSynchronizationJobStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusRequest) SetAccountId(v string) *DescribeSynchronizationJobStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBody struct {
	SynchronizationJobName        *string                                                                    `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	Status                        *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	DataInitialization            *string                                                                    `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	Delay                         *string                                                                    `json:"Delay,omitempty" xml:"Delay,omitempty"`
	ErrorMessage                  *string                                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	ErrCode                       *string                                                                    `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                       *string                                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	StructureInitialization       *string                                                                    `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	PayType                       *string                                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ErrMessage                    *string                                                                    `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SynchronizationJobClass       *string                                                                    `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationJobId          *string                                                                    `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	DelayMillis                   *int64                                                                     `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	SynchronizationDirection      *string                                                                    `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	TaskId                        *string                                                                    `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Checkpoint                    *string                                                                    `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DataInitializationStatus      *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	Performance                   *DescribeSynchronizationJobStatusResponseBodyPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeSynchronizationJobStatusResponseBodySourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	SynchronizationObjects        []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Delay = &v
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPayType(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBody {
	s.DelayMillis = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetTaskId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Checkpoint = &v
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.DestinationEndpoint = v
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetSourceEndpoint(v *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationObjects(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationObjects = v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DelayMillis  *int64  `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	Checkpoint   *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.DelayMillis = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus {
	s.Checkpoint = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint struct {
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetPort(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetUserName(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyPerformance struct {
	RPS  *string `json:"RPS,omitempty" xml:"RPS,omitempty"`
	FLOW *string `json:"FLOW,omitempty" xml:"FLOW,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPerformance) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetRPS(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.RPS = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPerformance) SetFLOW(v string) *DescribeSynchronizationJobStatusResponseBodyPerformance {
	s.FLOW = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyPrecheckStatus struct {
	Status  *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent *string                                                             `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Detail  []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationJobStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus {
	s.Detail = v
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
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceType(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceType = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetPort(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.Port = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetUserName(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.UserName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

type DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
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
	Headers map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationJobStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSynchronizationJobStatusResponse) SetBody(v *DescribeSynchronizationJobStatusResponseBody) *DescribeSynchronizationJobStatusResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobStatusListRequest struct {
	RegionId                        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobIdListJsonStr *string `json:"SynchronizationJobIdListJsonStr,omitempty" xml:"SynchronizationJobIdListJsonStr,omitempty"`
	ClientToken                     *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                       *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSynchronizationJobStatusListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListRequest) SetRegionId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetSynchronizationJobIdListJsonStr(v string) *DescribeSynchronizationJobStatusListRequest {
	s.SynchronizationJobIdListJsonStr = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetClientToken(v string) *DescribeSynchronizationJobStatusListRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetOwnerId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListRequest) SetAccountId(v string) *DescribeSynchronizationJobStatusListRequest {
	s.AccountId = &v
	return s
}

type DescribeSynchronizationJobStatusListResponseBody struct {
	RequestId                        *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                          *string                                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PageRecordCount                  *int32                                                                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	Success                          *string                                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	TotalRecordCount                 *int64                                                                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	ErrMessage                       *string                                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PageNumber                       *int32                                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	SynchronizationJobListStatusList []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList `json:"SynchronizationJobListStatusList,omitempty" xml:"SynchronizationJobListStatusList,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobStatusListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSynchronizationJobListStatusList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) *DescribeSynchronizationJobStatusListResponseBody {
	s.SynchronizationJobListStatusList = v
	return s
}

type DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList struct {
	SynchronizationJobId             *string                                                                                                             `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirectionInfoList []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList `json:"SynchronizationDirectionInfoList,omitempty" xml:"SynchronizationDirectionInfoList,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) SetSynchronizationDirectionInfoList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList {
	s.SynchronizationDirectionInfoList = v
	return s
}

type DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList struct {
	Status                   *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	Checkpoint               *string `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetStatus(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList) SetCheckpoint(v string) *DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusListSynchronizationDirectionInfoList {
	s.Checkpoint = &v
	return s
}

type DescribeSynchronizationJobStatusListResponse struct {
	Headers map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationJobStatusListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronizationJobStatusListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponse) SetHeaders(v map[string]*string) *DescribeSynchronizationJobStatusListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponse) SetBody(v *DescribeSynchronizationJobStatusListResponseBody) *DescribeSynchronizationJobStatusListResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationObjectModifyStatusRequest struct {
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId   *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetRegionId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetTaskId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetClientToken(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetOwnerId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusRequest) SetAccountId(v string) *DescribeSynchronizationObjectModifyStatusRequest {
	s.AccountId = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBody struct {
	Status                        *string                                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	ErrorMessage                  *string                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId                     *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode                       *string                                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                       *string                                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage                    *string                                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DataInitializationStatus      *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	PrecheckStatus                *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus struct {
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetDelay(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus {
	s.ErrorMessage = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus struct {
	Status  *string                                                                      `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent *string                                                                      `json:"Percent,omitempty" xml:"Percent,omitempty"`
	Detail  []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) SetDetail(v []*DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatusDetail) *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus {
	s.Detail = v
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
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetStatus(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetPercent(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Percent = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) SetProgress(v string) *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus {
	s.Progress = &v
	return s
}

type DescribeSynchronizationObjectModifyStatusResponse struct {
	Headers map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeSynchronizationObjectModifyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSynchronizationObjectModifyStatusResponse) SetBody(v *DescribeSynchronizationObjectModifyStatusResponseBody) *DescribeSynchronizationObjectModifyStatusResponse {
	s.Body = v
	return s
}

type IgnoreJobDetailRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	JobCode                  *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	Id                       *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
}

func (s IgnoreJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s IgnoreJobDetailRequest) GoString() string {
	return s.String()
}

func (s *IgnoreJobDetailRequest) SetDtsJobId(v string) *IgnoreJobDetailRequest {
	s.DtsJobId = &v
	return s
}

func (s *IgnoreJobDetailRequest) SetJobCode(v string) *IgnoreJobDetailRequest {
	s.JobCode = &v
	return s
}

func (s *IgnoreJobDetailRequest) SetId(v string) *IgnoreJobDetailRequest {
	s.Id = &v
	return s
}

func (s *IgnoreJobDetailRequest) SetRegionId(v string) *IgnoreJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *IgnoreJobDetailRequest) SetSynchronizationDirection(v string) *IgnoreJobDetailRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *IgnoreJobDetailRequest) SetDtsInstanceId(v string) *IgnoreJobDetailRequest {
	s.DtsInstanceId = &v
	return s
}

type IgnoreJobDetailResponseBody struct {
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s IgnoreJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s IgnoreJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreJobDetailResponseBody) SetCode(v string) *IgnoreJobDetailResponseBody {
	s.Code = &v
	return s
}

func (s *IgnoreJobDetailResponseBody) SetHttpStatusCode(v int32) *IgnoreJobDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *IgnoreJobDetailResponseBody) SetRequestId(v string) *IgnoreJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreJobDetailResponseBody) SetSuccess(v bool) *IgnoreJobDetailResponseBody {
	s.Success = &v
	return s
}

type IgnoreJobDetailResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *IgnoreJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s IgnoreJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s IgnoreJobDetailResponse) GoString() string {
	return s.String()
}

func (s *IgnoreJobDetailResponse) SetHeaders(v map[string]*string) *IgnoreJobDetailResponse {
	s.Headers = v
	return s
}

func (s *IgnoreJobDetailResponse) SetBody(v *IgnoreJobDetailResponseBody) *IgnoreJobDetailResponse {
	s.Body = v
	return s
}

type InitDtsRdsInstanceRequest struct {
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	EndpointInstanceType *string `json:"EndpointInstanceType,omitempty" xml:"EndpointInstanceType,omitempty"`
	EndpointRegion       *string `json:"EndpointRegion,omitempty" xml:"EndpointRegion,omitempty"`
	EndpointInstanceId   *string `json:"EndpointInstanceId,omitempty" xml:"EndpointInstanceId,omitempty"`
	EndpointCenId        *string `json:"EndpointCenId,omitempty" xml:"EndpointCenId,omitempty"`
}

func (s InitDtsRdsInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s InitDtsRdsInstanceRequest) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceRequest) SetRegionId(v string) *InitDtsRdsInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointInstanceType(v string) *InitDtsRdsInstanceRequest {
	s.EndpointInstanceType = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointRegion(v string) *InitDtsRdsInstanceRequest {
	s.EndpointRegion = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointInstanceId(v string) *InitDtsRdsInstanceRequest {
	s.EndpointInstanceId = &v
	return s
}

func (s *InitDtsRdsInstanceRequest) SetEndpointCenId(v string) *InitDtsRdsInstanceRequest {
	s.EndpointCenId = &v
	return s
}

type InitDtsRdsInstanceResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	AdminAccount   *string `json:"AdminAccount,omitempty" xml:"AdminAccount,omitempty"`
	AdminPassword  *string `json:"AdminPassword,omitempty" xml:"AdminPassword,omitempty"`
}

func (s InitDtsRdsInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitDtsRdsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceResponseBody) SetHttpStatusCode(v string) *InitDtsRdsInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetRequestId(v string) *InitDtsRdsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetErrCode(v string) *InitDtsRdsInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetSuccess(v string) *InitDtsRdsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetErrMessage(v string) *InitDtsRdsInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetAdminAccount(v string) *InitDtsRdsInstanceResponseBody {
	s.AdminAccount = &v
	return s
}

func (s *InitDtsRdsInstanceResponseBody) SetAdminPassword(v string) *InitDtsRdsInstanceResponseBody {
	s.AdminPassword = &v
	return s
}

type InitDtsRdsInstanceResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InitDtsRdsInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitDtsRdsInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s InitDtsRdsInstanceResponse) GoString() string {
	return s.String()
}

func (s *InitDtsRdsInstanceResponse) SetHeaders(v map[string]*string) *InitDtsRdsInstanceResponse {
	s.Headers = v
	return s
}

func (s *InitDtsRdsInstanceResponse) SetBody(v *InitDtsRdsInstanceResponseBody) *InitDtsRdsInstanceResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	RegionId     *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	NextToken    *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ResourceId   []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
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

type ListTagResourcesResponseBody struct {
	NextToken    *string                                   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId    *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode      *string                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage   *string                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
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

func (s *ListTagResourcesResponseBody) SetErrCode(v string) *ListTagResourcesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrMessage(v string) *ListTagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ModifyConsumerChannelRequest struct {
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId         *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId              *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ConsumerGroupId       *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	ConsumerGroupPassword *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	ConsumerGroupUserName *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	ConsumerGroupName     *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
}

func (s ModifyConsumerChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerChannelRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelRequest) SetRegionId(v string) *ModifyConsumerChannelRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetDtsInstanceId(v string) *ModifyConsumerChannelRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetDtsJobId(v string) *ModifyConsumerChannelRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupId(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupId = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupPassword(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupUserName(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *ModifyConsumerChannelRequest) SetConsumerGroupName(v string) *ModifyConsumerChannelRequest {
	s.ConsumerGroupName = &v
	return s
}

type ModifyConsumerChannelResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifyConsumerChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelResponseBody) SetHttpStatusCode(v string) *ModifyConsumerChannelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetRequestId(v string) *ModifyConsumerChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetErrCode(v string) *ModifyConsumerChannelResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetSuccess(v string) *ModifyConsumerChannelResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerChannelResponseBody) SetErrMessage(v string) *ModifyConsumerChannelResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifyConsumerChannelResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyConsumerChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyConsumerChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerChannelResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerChannelResponse) SetHeaders(v map[string]*string) *ModifyConsumerChannelResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerChannelResponse) SetBody(v *ModifyConsumerChannelResponseBody) *ModifyConsumerChannelResponse {
	s.Body = v
	return s
}

type ModifyConsumerGroupPasswordRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId   *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ConsumerGroupID          *string `json:"ConsumerGroupID,omitempty" xml:"ConsumerGroupID,omitempty"`
	ConsumerGroupName        *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	ConsumerGroupUserName    *string `json:"ConsumerGroupUserName,omitempty" xml:"ConsumerGroupUserName,omitempty"`
	ConsumerGroupPassword    *string `json:"ConsumerGroupPassword,omitempty" xml:"ConsumerGroupPassword,omitempty"`
	ConsumerGroupNewPassword *string `json:"consumerGroupNewPassword,omitempty" xml:"consumerGroupNewPassword,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ModifyConsumerGroupPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerGroupPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordRequest) SetRegionId(v string) *ModifyConsumerGroupPasswordRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetSubscriptionInstanceId(v string) *ModifyConsumerGroupPasswordRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupID(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupID = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupName(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupName = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupUserName(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupUserName = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupPassword(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupPassword = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetConsumerGroupNewPassword(v string) *ModifyConsumerGroupPasswordRequest {
	s.ConsumerGroupNewPassword = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetOwnerId(v string) *ModifyConsumerGroupPasswordRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordRequest) SetAccountId(v string) *ModifyConsumerGroupPasswordRequest {
	s.AccountId = &v
	return s
}

type ModifyConsumerGroupPasswordResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifyConsumerGroupPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerGroupPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetRequestId(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifyConsumerGroupPasswordResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyConsumerGroupPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyConsumerGroupPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumerGroupPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyConsumerGroupPasswordResponse) SetHeaders(v map[string]*string) *ModifyConsumerGroupPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyConsumerGroupPasswordResponse) SetBody(v *ModifyConsumerGroupPasswordResponseBody) *ModifyConsumerGroupPasswordResponse {
	s.Body = v
	return s
}

type ModifyConsumptionTimestampRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	ConsumptionTimestamp   *string `json:"ConsumptionTimestamp,omitempty" xml:"ConsumptionTimestamp,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ModifyConsumptionTimestampRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumptionTimestampRequest) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampRequest) SetRegionId(v string) *ModifyConsumptionTimestampRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetSubscriptionInstanceId(v string) *ModifyConsumptionTimestampRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetConsumptionTimestamp(v string) *ModifyConsumptionTimestampRequest {
	s.ConsumptionTimestamp = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetOwnerId(v string) *ModifyConsumptionTimestampRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyConsumptionTimestampRequest) SetAccountId(v string) *ModifyConsumptionTimestampRequest {
	s.AccountId = &v
	return s
}

type ModifyConsumptionTimestampResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifyConsumptionTimestampResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyConsumptionTimestampResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConsumptionTimestampResponseBody) SetRequestId(v string) *ModifyConsumptionTimestampResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrCode(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetSuccess(v string) *ModifyConsumptionTimestampResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrMessage(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifyConsumptionTimestampResponse struct {
	Headers map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyConsumptionTimestampResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifyConsumptionTimestampResponse) SetBody(v *ModifyConsumptionTimestampResponseBody) *ModifyConsumptionTimestampResponse {
	s.Body = v
	return s
}

type ModifyDtsJobRequest struct {
	RegionId                   *string                `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId              *string                `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection   *string                `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ClientToken                *string                `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbList                     map[string]interface{} `json:"DbList,omitempty" xml:"DbList,omitempty"`
	EtlOperatorColumnReference *string                `json:"EtlOperatorColumnReference,omitempty" xml:"EtlOperatorColumnReference,omitempty"`
}

func (s ModifyDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobRequest) SetRegionId(v string) *ModifyDtsJobRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDtsInstanceId(v string) *ModifyDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobRequest) SetSynchronizationDirection(v string) *ModifyDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobRequest) SetClientToken(v string) *ModifyDtsJobRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDtsJobRequest) SetDbList(v map[string]interface{}) *ModifyDtsJobRequest {
	s.DbList = v
	return s
}

func (s *ModifyDtsJobRequest) SetEtlOperatorColumnReference(v string) *ModifyDtsJobRequest {
	s.EtlOperatorColumnReference = &v
	return s
}

type ModifyDtsJobShrinkRequest struct {
	RegionId                   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId              *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection   *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ClientToken                *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	DbListShrink               *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	EtlOperatorColumnReference *string `json:"EtlOperatorColumnReference,omitempty" xml:"EtlOperatorColumnReference,omitempty"`
}

func (s ModifyDtsJobShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobShrinkRequest) SetRegionId(v string) *ModifyDtsJobShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDtsInstanceId(v string) *ModifyDtsJobShrinkRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetSynchronizationDirection(v string) *ModifyDtsJobShrinkRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetClientToken(v string) *ModifyDtsJobShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetDbListShrink(v string) *ModifyDtsJobShrinkRequest {
	s.DbListShrink = &v
	return s
}

func (s *ModifyDtsJobShrinkRequest) SetEtlOperatorColumnReference(v string) *ModifyDtsJobShrinkRequest {
	s.EtlOperatorColumnReference = &v
	return s
}

type ModifyDtsJobResponseBody struct {
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DtsJobId   *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *bool   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifyDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobResponseBody) SetStatus(v string) *ModifyDtsJobResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetRequestId(v string) *ModifyDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetDtsJobId(v string) *ModifyDtsJobResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetErrCode(v string) *ModifyDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetSuccess(v bool) *ModifyDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobResponseBody) SetErrMessage(v bool) *ModifyDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifyDtsJobResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobResponse) SetHeaders(v map[string]*string) *ModifyDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobResponse) SetBody(v *ModifyDtsJobResponseBody) *ModifyDtsJobResponse {
	s.Body = v
	return s
}

type ModifyDtsJobNameRequest struct {
	DtsJobId   *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsJobName *string `json:"DtsJobName,omitempty" xml:"DtsJobName,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDtsJobNameRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameRequest) SetDtsJobId(v string) *ModifyDtsJobNameRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetDtsJobName(v string) *ModifyDtsJobNameRequest {
	s.DtsJobName = &v
	return s
}

func (s *ModifyDtsJobNameRequest) SetRegionId(v string) *ModifyDtsJobNameRequest {
	s.RegionId = &v
	return s
}

type ModifyDtsJobNameResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
}

func (s ModifyDtsJobNameResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameResponseBody) SetHttpStatusCode(v int32) *ModifyDtsJobNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetRequestId(v string) *ModifyDtsJobNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetErrCode(v string) *ModifyDtsJobNameResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetSuccess(v bool) *ModifyDtsJobNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetErrMessage(v string) *ModifyDtsJobNameResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetCode(v string) *ModifyDtsJobNameResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobNameResponseBody) SetDynamicMessage(v string) *ModifyDtsJobNameResponseBody {
	s.DynamicMessage = &v
	return s
}

type ModifyDtsJobNameResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDtsJobNameResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDtsJobNameResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobNameResponse) SetHeaders(v map[string]*string) *ModifyDtsJobNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobNameResponse) SetBody(v *ModifyDtsJobNameResponseBody) *ModifyDtsJobNameResponse {
	s.Body = v
	return s
}

type ModifyDtsJobPasswordRequest struct {
	DtsJobId *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDtsJobPasswordRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobPasswordRequest) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordRequest) SetDtsJobId(v string) *ModifyDtsJobPasswordRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetUserName(v string) *ModifyDtsJobPasswordRequest {
	s.UserName = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetPassword(v string) *ModifyDtsJobPasswordRequest {
	s.Password = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetEndpoint(v string) *ModifyDtsJobPasswordRequest {
	s.Endpoint = &v
	return s
}

func (s *ModifyDtsJobPasswordRequest) SetRegionId(v string) *ModifyDtsJobPasswordRequest {
	s.RegionId = &v
	return s
}

type ModifyDtsJobPasswordResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
}

func (s ModifyDtsJobPasswordResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordResponseBody) SetHttpStatusCode(v int32) *ModifyDtsJobPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetRequestId(v string) *ModifyDtsJobPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetErrCode(v string) *ModifyDtsJobPasswordResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetSuccess(v bool) *ModifyDtsJobPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetErrMessage(v string) *ModifyDtsJobPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetCode(v string) *ModifyDtsJobPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDtsJobPasswordResponseBody) SetDynamicMessage(v string) *ModifyDtsJobPasswordResponseBody {
	s.DynamicMessage = &v
	return s
}

type ModifyDtsJobPasswordResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifyDtsJobPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifyDtsJobPasswordResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyDtsJobPasswordResponse) GoString() string {
	return s.String()
}

func (s *ModifyDtsJobPasswordResponse) SetHeaders(v map[string]*string) *ModifyDtsJobPasswordResponse {
	s.Headers = v
	return s
}

func (s *ModifyDtsJobPasswordResponse) SetBody(v *ModifyDtsJobPasswordResponseBody) *ModifyDtsJobPasswordResponse {
	s.Body = v
	return s
}

type ModifySubscriptionRequest struct {
	RegionId                *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId           *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DbList                  *string `json:"DbList,omitempty" xml:"DbList,omitempty"`
	DtsJobId                *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	SubscriptionDataTypeDDL *bool   `json:"SubscriptionDataTypeDDL,omitempty" xml:"SubscriptionDataTypeDDL,omitempty"`
	SubscriptionDataTypeDML *bool   `json:"SubscriptionDataTypeDML,omitempty" xml:"SubscriptionDataTypeDML,omitempty"`
}

func (s ModifySubscriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionRequest) SetRegionId(v string) *ModifySubscriptionRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetDtsInstanceId(v string) *ModifySubscriptionRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetDbList(v string) *ModifySubscriptionRequest {
	s.DbList = &v
	return s
}

func (s *ModifySubscriptionRequest) SetDtsJobId(v string) *ModifySubscriptionRequest {
	s.DtsJobId = &v
	return s
}

func (s *ModifySubscriptionRequest) SetSubscriptionDataTypeDDL(v bool) *ModifySubscriptionRequest {
	s.SubscriptionDataTypeDDL = &v
	return s
}

func (s *ModifySubscriptionRequest) SetSubscriptionDataTypeDML(v bool) *ModifySubscriptionRequest {
	s.SubscriptionDataTypeDML = &v
	return s
}

type ModifySubscriptionResponseBody struct {
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifySubscriptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionResponseBody) SetHttpStatusCode(v string) *ModifySubscriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetRequestId(v string) *ModifySubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetErrCode(v string) *ModifySubscriptionResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetSuccess(v string) *ModifySubscriptionResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySubscriptionResponseBody) SetErrMessage(v string) *ModifySubscriptionResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifySubscriptionResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ModifySubscriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionResponse) SetBody(v *ModifySubscriptionResponseBody) *ModifySubscriptionResponse {
	s.Body = v
	return s
}

type ModifySubscriptionObjectRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	SubscriptionObject     *string `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ModifySubscriptionObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectRequest) SetRegionId(v string) *ModifySubscriptionObjectRequest {
	s.RegionId = &v
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

func (s *ModifySubscriptionObjectRequest) SetOwnerId(v string) *ModifySubscriptionObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySubscriptionObjectRequest) SetAccountId(v string) *ModifySubscriptionObjectRequest {
	s.AccountId = &v
	return s
}

type ModifySubscriptionObjectResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifySubscriptionObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySubscriptionObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectResponseBody) SetRequestId(v string) *ModifySubscriptionObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetErrCode(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetSuccess(v string) *ModifySubscriptionObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetErrMessage(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifySubscriptionObjectResponse struct {
	Headers map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySubscriptionObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySubscriptionObjectResponse) SetBody(v *ModifySubscriptionObjectResponseBody) *ModifySubscriptionObjectResponse {
	s.Body = v
	return s
}

type ModifySynchronizationObjectRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationObjects   *string `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ModifySynchronizationObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectRequest) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectRequest) SetRegionId(v string) *ModifySynchronizationObjectRequest {
	s.RegionId = &v
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

func (s *ModifySynchronizationObjectRequest) SetSynchronizationDirection(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetOwnerId(v string) *ModifySynchronizationObjectRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySynchronizationObjectRequest) SetAccountId(v string) *ModifySynchronizationObjectRequest {
	s.AccountId = &v
	return s
}

type ModifySynchronizationObjectResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ModifySynchronizationObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponseBody) SetRequestId(v string) *ModifySynchronizationObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrCode(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetTaskId(v string) *ModifySynchronizationObjectResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetSuccess(v string) *ModifySynchronizationObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrMessage(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrMessage = &v
	return s
}

type ModifySynchronizationObjectResponse struct {
	Headers map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ModifySynchronizationObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ModifySynchronizationObjectResponse) SetBody(v *ModifySynchronizationObjectResponseBody) *ModifySynchronizationObjectResponse {
	s.Body = v
	return s
}

type RenewInstanceRequest struct {
	DtsJobId   *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	BuyCount   *string `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewInstanceRequest) SetDtsJobId(v string) *RenewInstanceRequest {
	s.DtsJobId = &v
	return s
}

func (s *RenewInstanceRequest) SetChargeType(v string) *RenewInstanceRequest {
	s.ChargeType = &v
	return s
}

func (s *RenewInstanceRequest) SetPeriod(v string) *RenewInstanceRequest {
	s.Period = &v
	return s
}

func (s *RenewInstanceRequest) SetBuyCount(v string) *RenewInstanceRequest {
	s.BuyCount = &v
	return s
}

func (s *RenewInstanceRequest) SetRegionId(v string) *RenewInstanceRequest {
	s.RegionId = &v
	return s
}

type RenewInstanceResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RenewInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponseBody) SetHttpStatusCode(v int32) *RenewInstanceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetEndTime(v string) *RenewInstanceResponseBody {
	s.EndTime = &v
	return s
}

func (s *RenewInstanceResponseBody) SetRequestId(v string) *RenewInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrCode(v string) *RenewInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RenewInstanceResponseBody) SetDtsJobId(v string) *RenewInstanceResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *RenewInstanceResponseBody) SetSuccess(v bool) *RenewInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *RenewInstanceResponseBody) SetErrMessage(v string) *RenewInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetChargeType(v string) *RenewInstanceResponseBody {
	s.ChargeType = &v
	return s
}

func (s *RenewInstanceResponseBody) SetCode(v string) *RenewInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *RenewInstanceResponseBody) SetDynamicMessage(v string) *RenewInstanceResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *RenewInstanceResponseBody) SetInstanceId(v string) *RenewInstanceResponseBody {
	s.InstanceId = &v
	return s
}

type RenewInstanceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenewInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewInstanceResponse) SetHeaders(v map[string]*string) *RenewInstanceResponse {
	s.Headers = v
	return s
}

func (s *RenewInstanceResponse) SetBody(v *RenewInstanceResponseBody) *RenewInstanceResponse {
	s.Body = v
	return s
}

type ResetDtsJobRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ResetDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetDtsJobRequest) GoString() string {
	return s.String()
}

func (s *ResetDtsJobRequest) SetDtsJobId(v string) *ResetDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *ResetDtsJobRequest) SetDtsInstanceId(v string) *ResetDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ResetDtsJobRequest) SetSynchronizationDirection(v string) *ResetDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ResetDtsJobRequest) SetRegionId(v string) *ResetDtsJobRequest {
	s.RegionId = &v
	return s
}

type ResetDtsJobResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s ResetDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetDtsJobResponseBody) SetHttpStatusCode(v int32) *ResetDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetRequestId(v string) *ResetDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetErrCode(v string) *ResetDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetSuccess(v bool) *ResetDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetErrMessage(v string) *ResetDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetDynamicMessage(v string) *ResetDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *ResetDtsJobResponseBody) SetDynamicCode(v string) *ResetDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

type ResetDtsJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetDtsJobResponse) GoString() string {
	return s.String()
}

func (s *ResetDtsJobResponse) SetHeaders(v map[string]*string) *ResetDtsJobResponse {
	s.Headers = v
	return s
}

func (s *ResetDtsJobResponse) SetBody(v *ResetDtsJobResponseBody) *ResetDtsJobResponse {
	s.Body = v
	return s
}

type ResetSynchronizationJobRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s ResetSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobRequest) SetRegionId(v string) *ResetSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationJobId(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationDirection(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetOwnerId(v string) *ResetSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetAccountId(v string) *ResetSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

type ResetSynchronizationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ResetSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobResponseBody) SetRequestId(v string) *ResetSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetErrCode(v string) *ResetSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetSuccess(v string) *ResetSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetErrMessage(v string) *ResetSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type ResetSynchronizationJobResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ResetSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetSynchronizationJobResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetSynchronizationJobResponse) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobResponse) SetHeaders(v map[string]*string) *ResetSynchronizationJobResponse {
	s.Headers = v
	return s
}

func (s *ResetSynchronizationJobResponse) SetBody(v *ResetSynchronizationJobResponseBody) *ResetSynchronizationJobResponse {
	s.Body = v
	return s
}

type ShieldPrecheckRequest struct {
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	PrecheckItems *string `json:"PrecheckItems,omitempty" xml:"PrecheckItems,omitempty"`
}

func (s ShieldPrecheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ShieldPrecheckRequest) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckRequest) SetRegionId(v string) *ShieldPrecheckRequest {
	s.RegionId = &v
	return s
}

func (s *ShieldPrecheckRequest) SetDtsInstanceId(v string) *ShieldPrecheckRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *ShieldPrecheckRequest) SetPrecheckItems(v string) *ShieldPrecheckRequest {
	s.PrecheckItems = &v
	return s
}

type ShieldPrecheckResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s ShieldPrecheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShieldPrecheckResponseBody) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckResponseBody) SetRequestId(v string) *ShieldPrecheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetErrCode(v string) *ShieldPrecheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetSuccess(v bool) *ShieldPrecheckResponseBody {
	s.Success = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetErrMessage(v string) *ShieldPrecheckResponseBody {
	s.ErrMessage = &v
	return s
}

type ShieldPrecheckResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ShieldPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ShieldPrecheckResponse) String() string {
	return tea.Prettify(s)
}

func (s ShieldPrecheckResponse) GoString() string {
	return s.String()
}

func (s *ShieldPrecheckResponse) SetHeaders(v map[string]*string) *ShieldPrecheckResponse {
	s.Headers = v
	return s
}

func (s *ShieldPrecheckResponse) SetBody(v *ShieldPrecheckResponseBody) *ShieldPrecheckResponse {
	s.Body = v
	return s
}

type SkipPreCheckRequest struct {
	JobId             *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	Skip              *bool   `json:"Skip,omitempty" xml:"Skip,omitempty"`
	SkipPreCheckItems *string `json:"SkipPreCheckItems,omitempty" xml:"SkipPreCheckItems,omitempty"`
	SkipPreCheckNames *string `json:"SkipPreCheckNames,omitempty" xml:"SkipPreCheckNames,omitempty"`
	DtsJobId          *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SkipPreCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s SkipPreCheckRequest) GoString() string {
	return s.String()
}

func (s *SkipPreCheckRequest) SetJobId(v string) *SkipPreCheckRequest {
	s.JobId = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkip(v bool) *SkipPreCheckRequest {
	s.Skip = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkipPreCheckItems(v string) *SkipPreCheckRequest {
	s.SkipPreCheckItems = &v
	return s
}

func (s *SkipPreCheckRequest) SetSkipPreCheckNames(v string) *SkipPreCheckRequest {
	s.SkipPreCheckNames = &v
	return s
}

func (s *SkipPreCheckRequest) SetDtsJobId(v string) *SkipPreCheckRequest {
	s.DtsJobId = &v
	return s
}

func (s *SkipPreCheckRequest) SetRegionId(v string) *SkipPreCheckRequest {
	s.RegionId = &v
	return s
}

type SkipPreCheckResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SkipNames      *string `json:"SkipNames,omitempty" xml:"SkipNames,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ScheduleJobId  *string `json:"ScheduleJobId,omitempty" xml:"ScheduleJobId,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	SkipItems      *string `json:"SkipItems,omitempty" xml:"SkipItems,omitempty"`
}

func (s SkipPreCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SkipPreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SkipPreCheckResponseBody) SetHttpStatusCode(v int32) *SkipPreCheckResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetRequestId(v string) *SkipPreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetErrCode(v string) *SkipPreCheckResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSuccess(v bool) *SkipPreCheckResponseBody {
	s.Success = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetErrMessage(v string) *SkipPreCheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSkipNames(v string) *SkipPreCheckResponseBody {
	s.SkipNames = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetCode(v string) *SkipPreCheckResponseBody {
	s.Code = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetScheduleJobId(v string) *SkipPreCheckResponseBody {
	s.ScheduleJobId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetDynamicMessage(v string) *SkipPreCheckResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetMigrationJobId(v string) *SkipPreCheckResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *SkipPreCheckResponseBody) SetSkipItems(v string) *SkipPreCheckResponseBody {
	s.SkipItems = &v
	return s
}

type SkipPreCheckResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SkipPreCheckResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SkipPreCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s SkipPreCheckResponse) GoString() string {
	return s.String()
}

func (s *SkipPreCheckResponse) SetHeaders(v map[string]*string) *SkipPreCheckResponse {
	s.Headers = v
	return s
}

func (s *SkipPreCheckResponse) SetBody(v *SkipPreCheckResponseBody) *SkipPreCheckResponse {
	s.Body = v
	return s
}

type StartDtsJobRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartDtsJobRequest) GoString() string {
	return s.String()
}

func (s *StartDtsJobRequest) SetDtsJobId(v string) *StartDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *StartDtsJobRequest) SetDtsInstanceId(v string) *StartDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *StartDtsJobRequest) SetSynchronizationDirection(v string) *StartDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StartDtsJobRequest) SetRegionId(v string) *StartDtsJobRequest {
	s.RegionId = &v
	return s
}

type StartDtsJobResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s StartDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartDtsJobResponseBody) SetHttpStatusCode(v int32) *StartDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StartDtsJobResponseBody) SetRequestId(v string) *StartDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDtsJobResponseBody) SetErrCode(v string) *StartDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartDtsJobResponseBody) SetSuccess(v bool) *StartDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartDtsJobResponseBody) SetErrMessage(v string) *StartDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartDtsJobResponseBody) SetDynamicMessage(v string) *StartDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StartDtsJobResponseBody) SetDynamicCode(v string) *StartDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

type StartDtsJobResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartDtsJobResponse) GoString() string {
	return s.String()
}

func (s *StartDtsJobResponse) SetHeaders(v map[string]*string) *StartDtsJobResponse {
	s.Headers = v
	return s
}

func (s *StartDtsJobResponse) SetBody(v *StartDtsJobResponseBody) *StartDtsJobResponse {
	s.Body = v
	return s
}

type StartMigrationJobRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s StartMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StartMigrationJobRequest) SetRegionId(v string) *StartMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartMigrationJobRequest) SetMigrationJobId(v string) *StartMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StartMigrationJobRequest) SetOwnerId(v string) *StartMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartMigrationJobRequest) SetAccountId(v string) *StartMigrationJobRequest {
	s.AccountId = &v
	return s
}

type StartMigrationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s StartMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartMigrationJobResponseBody) SetRequestId(v string) *StartMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetErrCode(v string) *StartMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetSuccess(v string) *StartMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetErrMessage(v string) *StartMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type StartMigrationJobResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartMigrationJobResponse) SetBody(v *StartMigrationJobResponseBody) *StartMigrationJobResponse {
	s.Body = v
	return s
}

type StartSubscriptionInstanceRequest struct {
	RegionId               *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SubscriptionInstanceId *string `json:"SubscriptionInstanceId,omitempty" xml:"SubscriptionInstanceId,omitempty"`
	OwnerId                *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId              *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s StartSubscriptionInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceRequest) SetRegionId(v string) *StartSubscriptionInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetSubscriptionInstanceId(v string) *StartSubscriptionInstanceRequest {
	s.SubscriptionInstanceId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetOwnerId(v string) *StartSubscriptionInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSubscriptionInstanceRequest) SetAccountId(v string) *StartSubscriptionInstanceRequest {
	s.AccountId = &v
	return s
}

type StartSubscriptionInstanceResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s StartSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponseBody) SetRequestId(v string) *StartSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrCode(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetTaskId(v string) *StartSubscriptionInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetSuccess(v string) *StartSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrMessage(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

type StartSubscriptionInstanceResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartSubscriptionInstanceResponse) SetBody(v *StartSubscriptionInstanceResponseBody) *StartSubscriptionInstanceResponse {
	s.Body = v
	return s
}

type StartSynchronizationJobRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s StartSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobRequest) SetRegionId(v string) *StartSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetSynchronizationJobId(v string) *StartSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetSynchronizationDirection(v string) *StartSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetOwnerId(v string) *StartSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StartSynchronizationJobRequest) SetAccountId(v string) *StartSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

type StartSynchronizationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s StartSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobResponseBody) SetRequestId(v string) *StartSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetErrCode(v string) *StartSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetSuccess(v string) *StartSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetErrMessage(v string) *StartSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type StartSynchronizationJobResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartSynchronizationJobResponse) SetBody(v *StartSynchronizationJobResponseBody) *StartSynchronizationJobResponse {
	s.Body = v
	return s
}

type StopDtsJobRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopDtsJobRequest) GoString() string {
	return s.String()
}

func (s *StopDtsJobRequest) SetDtsJobId(v string) *StopDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *StopDtsJobRequest) SetDtsInstanceId(v string) *StopDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *StopDtsJobRequest) SetSynchronizationDirection(v string) *StopDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *StopDtsJobRequest) SetRegionId(v string) *StopDtsJobRequest {
	s.RegionId = &v
	return s
}

type StopDtsJobResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s StopDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopDtsJobResponseBody) SetHttpStatusCode(v int32) *StopDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *StopDtsJobResponseBody) SetRequestId(v string) *StopDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDtsJobResponseBody) SetErrCode(v string) *StopDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopDtsJobResponseBody) SetSuccess(v bool) *StopDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopDtsJobResponseBody) SetErrMessage(v string) *StopDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopDtsJobResponseBody) SetDynamicMessage(v string) *StopDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *StopDtsJobResponseBody) SetDynamicCode(v string) *StopDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

type StopDtsJobResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopDtsJobResponse) GoString() string {
	return s.String()
}

func (s *StopDtsJobResponse) SetHeaders(v map[string]*string) *StopDtsJobResponse {
	s.Headers = v
	return s
}

func (s *StopDtsJobResponse) SetBody(v *StopDtsJobResponseBody) *StopDtsJobResponse {
	s.Body = v
	return s
}

type StopMigrationJobRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s StopMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *StopMigrationJobRequest) SetRegionId(v string) *StopMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *StopMigrationJobRequest) SetMigrationJobId(v string) *StopMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *StopMigrationJobRequest) SetClientToken(v string) *StopMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *StopMigrationJobRequest) SetOwnerId(v string) *StopMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *StopMigrationJobRequest) SetAccountId(v string) *StopMigrationJobRequest {
	s.AccountId = &v
	return s
}

type StopMigrationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s StopMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopMigrationJobResponseBody) SetRequestId(v string) *StopMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetErrCode(v string) *StopMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetSuccess(v string) *StopMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetErrMessage(v string) *StopMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type StopMigrationJobResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopMigrationJobResponse) SetBody(v *StopMigrationJobResponseBody) *StopMigrationJobResponse {
	s.Body = v
	return s
}

type SummaryJobDetailRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	JobCode                  *string `json:"JobCode,omitempty" xml:"JobCode,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
}

func (s SummaryJobDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s SummaryJobDetailRequest) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailRequest) SetDtsJobId(v string) *SummaryJobDetailRequest {
	s.DtsJobId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetJobCode(v string) *SummaryJobDetailRequest {
	s.JobCode = &v
	return s
}

func (s *SummaryJobDetailRequest) SetSynchronizationDirection(v string) *SummaryJobDetailRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SummaryJobDetailRequest) SetRegionId(v string) *SummaryJobDetailRequest {
	s.RegionId = &v
	return s
}

func (s *SummaryJobDetailRequest) SetDtsInstanceId(v string) *SummaryJobDetailRequest {
	s.DtsInstanceId = &v
	return s
}

type SummaryJobDetailResponseBody struct {
	HttpStatusCode         *int32                                                `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Code                   *string                                               `json:"Code,omitempty" xml:"Code,omitempty"`
	RequestId              *string                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	JobId                  *string                                               `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ProgressSummaryDetails []*SummaryJobDetailResponseBodyProgressSummaryDetails `json:"ProgressSummaryDetails,omitempty" xml:"ProgressSummaryDetails,omitempty" type:"Repeated"`
}

func (s SummaryJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SummaryJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponseBody) SetHttpStatusCode(v int32) *SummaryJobDetailResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetCode(v string) *SummaryJobDetailResponseBody {
	s.Code = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetRequestId(v string) *SummaryJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetSuccess(v bool) *SummaryJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetJobId(v string) *SummaryJobDetailResponseBody {
	s.JobId = &v
	return s
}

func (s *SummaryJobDetailResponseBody) SetProgressSummaryDetails(v []*SummaryJobDetailResponseBodyProgressSummaryDetails) *SummaryJobDetailResponseBody {
	s.ProgressSummaryDetails = v
	return s
}

type SummaryJobDetailResponseBodyProgressSummaryDetails struct {
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
	TotalCount *int64  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	State      *int32  `json:"State,omitempty" xml:"State,omitempty"`
}

func (s SummaryJobDetailResponseBodyProgressSummaryDetails) String() string {
	return tea.Prettify(s)
}

func (s SummaryJobDetailResponseBodyProgressSummaryDetails) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetKey(v string) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.Key = &v
	return s
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetTotalCount(v int64) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.TotalCount = &v
	return s
}

func (s *SummaryJobDetailResponseBodyProgressSummaryDetails) SetState(v int32) *SummaryJobDetailResponseBodyProgressSummaryDetails {
	s.State = &v
	return s
}

type SummaryJobDetailResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SummaryJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SummaryJobDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s SummaryJobDetailResponse) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponse) SetHeaders(v map[string]*string) *SummaryJobDetailResponse {
	s.Headers = v
	return s
}

func (s *SummaryJobDetailResponse) SetBody(v *SummaryJobDetailResponseBody) *SummaryJobDetailResponse {
	s.Body = v
	return s
}

type SuspendDtsJobRequest struct {
	DtsJobId                 *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	DtsInstanceId            *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SuspendDtsJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendDtsJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobRequest) SetDtsJobId(v string) *SuspendDtsJobRequest {
	s.DtsJobId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetDtsInstanceId(v string) *SuspendDtsJobRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *SuspendDtsJobRequest) SetSynchronizationDirection(v string) *SuspendDtsJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SuspendDtsJobRequest) SetRegionId(v string) *SuspendDtsJobRequest {
	s.RegionId = &v
	return s
}

type SuspendDtsJobResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s SuspendDtsJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendDtsJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobResponseBody) SetHttpStatusCode(v int32) *SuspendDtsJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetRequestId(v string) *SuspendDtsJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetErrCode(v string) *SuspendDtsJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetSuccess(v bool) *SuspendDtsJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetErrMessage(v string) *SuspendDtsJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetDynamicMessage(v string) *SuspendDtsJobResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *SuspendDtsJobResponseBody) SetDynamicCode(v string) *SuspendDtsJobResponseBody {
	s.DynamicCode = &v
	return s
}

type SuspendDtsJobResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendDtsJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SuspendDtsJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendDtsJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendDtsJobResponse) SetHeaders(v map[string]*string) *SuspendDtsJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendDtsJobResponse) SetBody(v *SuspendDtsJobResponseBody) *SuspendDtsJobResponse {
	s.Body = v
	return s
}

type SuspendMigrationJobRequest struct {
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ClientToken    *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId        *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId      *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s SuspendMigrationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendMigrationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobRequest) SetRegionId(v string) *SuspendMigrationJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetMigrationJobId(v string) *SuspendMigrationJobRequest {
	s.MigrationJobId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetClientToken(v string) *SuspendMigrationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetOwnerId(v string) *SuspendMigrationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendMigrationJobRequest) SetAccountId(v string) *SuspendMigrationJobRequest {
	s.AccountId = &v
	return s
}

type SuspendMigrationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SuspendMigrationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobResponseBody) SetRequestId(v string) *SuspendMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetErrCode(v string) *SuspendMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetSuccess(v string) *SuspendMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetErrMessage(v string) *SuspendMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type SuspendMigrationJobResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendMigrationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SuspendMigrationJobResponse) SetBody(v *SuspendMigrationJobResponseBody) *SuspendMigrationJobResponse {
	s.Body = v
	return s
}

type SuspendSynchronizationJobRequest struct {
	RegionId                 *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s SuspendSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobRequest) SetRegionId(v string) *SuspendSynchronizationJobRequest {
	s.RegionId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetSynchronizationJobId(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetSynchronizationDirection(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetOwnerId(v string) *SuspendSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *SuspendSynchronizationJobRequest) SetAccountId(v string) *SuspendSynchronizationJobRequest {
	s.AccountId = &v
	return s
}

type SuspendSynchronizationJobResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SuspendSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobResponseBody) SetRequestId(v string) *SuspendSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetErrCode(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetSuccess(v string) *SuspendSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetErrMessage(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

type SuspendSynchronizationJobResponse struct {
	Headers map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SuspendSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SuspendSynchronizationJobResponse) SetBody(v *SuspendSynchronizationJobResponseBody) *SuspendSynchronizationJobResponse {
	s.Body = v
	return s
}

type SwitchSynchronizationEndpointRequest struct {
	Endpoint                 *SwitchSynchronizationEndpointRequestEndpoint       `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	SourceEndpoint           *SwitchSynchronizationEndpointRequestSourceEndpoint `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	RegionId                 *string                                             `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SynchronizationJobId     *string                                             `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection *string                                             `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	OwnerId                  *string                                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	AccountId                *string                                             `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
}

func (s SwitchSynchronizationEndpointRequest) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequest) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequest) SetEndpoint(v *SwitchSynchronizationEndpointRequestEndpoint) *SwitchSynchronizationEndpointRequest {
	s.Endpoint = v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSourceEndpoint(v *SwitchSynchronizationEndpointRequestSourceEndpoint) *SwitchSynchronizationEndpointRequest {
	s.SourceEndpoint = v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetRegionId(v string) *SwitchSynchronizationEndpointRequest {
	s.RegionId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationJobId(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationJobId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationDirection(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetOwnerId(v string) *SwitchSynchronizationEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetAccountId(v string) *SwitchSynchronizationEndpointRequest {
	s.AccountId = &v
	return s
}

type SwitchSynchronizationEndpointRequestEndpoint struct {
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s SwitchSynchronizationEndpointRequestEndpoint) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequestEndpoint) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Type = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceType = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceId(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetIP(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.IP = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetPort(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Port = &v
	return s
}

type SwitchSynchronizationEndpointRequestSourceEndpoint struct {
	OwnerID *string `json:"OwnerID,omitempty" xml:"OwnerID,omitempty"`
	Role    *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s SwitchSynchronizationEndpointRequestSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequestSourceEndpoint) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) SetOwnerID(v string) *SwitchSynchronizationEndpointRequestSourceEndpoint {
	s.OwnerID = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestSourceEndpoint) SetRole(v string) *SwitchSynchronizationEndpointRequestSourceEndpoint {
	s.Role = &v
	return s
}

type SwitchSynchronizationEndpointResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s SwitchSynchronizationEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponseBody) SetRequestId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrCode(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetTaskId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetSuccess(v string) *SwitchSynchronizationEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrMessage(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

type SwitchSynchronizationEndpointResponse struct {
	Headers map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SwitchSynchronizationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SwitchSynchronizationEndpointResponse) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponse) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponse) SetHeaders(v map[string]*string) *SwitchSynchronizationEndpointResponse {
	s.Headers = v
	return s
}

func (s *SwitchSynchronizationEndpointResponse) SetBody(v *SwitchSynchronizationEndpointResponseBody) *SwitchSynchronizationEndpointResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	RegionId     *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ResourceId   []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	Tag          []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
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

type TagResourcesResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

func (s *TagResourcesResponseBody) SetErrCode(v string) *TagResourcesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetErrMessage(v string) *TagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

type TagResourcesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type TransferInstanceClassRequest struct {
	OrderType     *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	DtsJobId      *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
}

func (s TransferInstanceClassRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferInstanceClassRequest) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassRequest) SetOrderType(v string) *TransferInstanceClassRequest {
	s.OrderType = &v
	return s
}

func (s *TransferInstanceClassRequest) SetInstanceClass(v string) *TransferInstanceClassRequest {
	s.InstanceClass = &v
	return s
}

func (s *TransferInstanceClassRequest) SetRegionId(v string) *TransferInstanceClassRequest {
	s.RegionId = &v
	return s
}

func (s *TransferInstanceClassRequest) SetDtsJobId(v string) *TransferInstanceClassRequest {
	s.DtsJobId = &v
	return s
}

type TransferInstanceClassResponseBody struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TransferInstanceClassResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferInstanceClassResponseBody) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassResponseBody) SetEndTime(v string) *TransferInstanceClassResponseBody {
	s.EndTime = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetHttpStatusCode(v int32) *TransferInstanceClassResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetRequestId(v string) *TransferInstanceClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetDtsJobId(v string) *TransferInstanceClassResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetErrCode(v string) *TransferInstanceClassResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetSuccess(v bool) *TransferInstanceClassResponseBody {
	s.Success = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetChargeType(v string) *TransferInstanceClassResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetErrMessage(v string) *TransferInstanceClassResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetCode(v string) *TransferInstanceClassResponseBody {
	s.Code = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetDynamicMessage(v string) *TransferInstanceClassResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TransferInstanceClassResponseBody) SetInstanceId(v string) *TransferInstanceClassResponseBody {
	s.InstanceId = &v
	return s
}

type TransferInstanceClassResponse struct {
	Headers map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferInstanceClassResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferInstanceClassResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferInstanceClassResponse) GoString() string {
	return s.String()
}

func (s *TransferInstanceClassResponse) SetHeaders(v map[string]*string) *TransferInstanceClassResponse {
	s.Headers = v
	return s
}

func (s *TransferInstanceClassResponse) SetBody(v *TransferInstanceClassResponseBody) *TransferInstanceClassResponse {
	s.Body = v
	return s
}

type TransferPayTypeRequest struct {
	DtsJobId   *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Period     *string `json:"Period,omitempty" xml:"Period,omitempty"`
	BuyCount   *string `json:"BuyCount,omitempty" xml:"BuyCount,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s TransferPayTypeRequest) String() string {
	return tea.Prettify(s)
}

func (s TransferPayTypeRequest) GoString() string {
	return s.String()
}

func (s *TransferPayTypeRequest) SetDtsJobId(v string) *TransferPayTypeRequest {
	s.DtsJobId = &v
	return s
}

func (s *TransferPayTypeRequest) SetChargeType(v string) *TransferPayTypeRequest {
	s.ChargeType = &v
	return s
}

func (s *TransferPayTypeRequest) SetPeriod(v string) *TransferPayTypeRequest {
	s.Period = &v
	return s
}

func (s *TransferPayTypeRequest) SetBuyCount(v string) *TransferPayTypeRequest {
	s.BuyCount = &v
	return s
}

func (s *TransferPayTypeRequest) SetRegionId(v string) *TransferPayTypeRequest {
	s.RegionId = &v
	return s
}

type TransferPayTypeResponseBody struct {
	EndTime        *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DtsJobId       *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ChargeType     *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	InstanceId     *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s TransferPayTypeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TransferPayTypeResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPayTypeResponseBody) SetEndTime(v string) *TransferPayTypeResponseBody {
	s.EndTime = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetHttpStatusCode(v int32) *TransferPayTypeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetRequestId(v string) *TransferPayTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetDtsJobId(v string) *TransferPayTypeResponseBody {
	s.DtsJobId = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetErrCode(v string) *TransferPayTypeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetSuccess(v bool) *TransferPayTypeResponseBody {
	s.Success = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetChargeType(v string) *TransferPayTypeResponseBody {
	s.ChargeType = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetErrMessage(v string) *TransferPayTypeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetCode(v string) *TransferPayTypeResponseBody {
	s.Code = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetDynamicMessage(v string) *TransferPayTypeResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *TransferPayTypeResponseBody) SetInstanceId(v string) *TransferPayTypeResponseBody {
	s.InstanceId = &v
	return s
}

type TransferPayTypeResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TransferPayTypeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TransferPayTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s TransferPayTypeResponse) GoString() string {
	return s.String()
}

func (s *TransferPayTypeResponse) SetHeaders(v map[string]*string) *TransferPayTypeResponse {
	s.Headers = v
	return s
}

func (s *TransferPayTypeResponse) SetBody(v *TransferPayTypeResponseBody) *TransferPayTypeResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	RegionId     *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceType *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	All          *bool     `json:"All,omitempty" xml:"All,omitempty"`
	ResourceId   []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	TagKey       []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
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

type UntagResourcesResponseBody struct {
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrCode(v string) *UntagResourcesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrMessage(v string) *UntagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

type UntagResourcesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpgradeTwoWayRequest struct {
	InstanceId    *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpgradeTwoWayRequest) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTwoWayRequest) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayRequest) SetInstanceId(v string) *UpgradeTwoWayRequest {
	s.InstanceId = &v
	return s
}

func (s *UpgradeTwoWayRequest) SetInstanceClass(v string) *UpgradeTwoWayRequest {
	s.InstanceClass = &v
	return s
}

func (s *UpgradeTwoWayRequest) SetRegionId(v string) *UpgradeTwoWayRequest {
	s.RegionId = &v
	return s
}

type UpgradeTwoWayResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s UpgradeTwoWayResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTwoWayResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayResponseBody) SetHttpStatusCode(v int32) *UpgradeTwoWayResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetRequestId(v string) *UpgradeTwoWayResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetErrCode(v string) *UpgradeTwoWayResponseBody {
	s.ErrCode = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetSuccess(v bool) *UpgradeTwoWayResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetErrMessage(v string) *UpgradeTwoWayResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetDynamicMessage(v string) *UpgradeTwoWayResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *UpgradeTwoWayResponseBody) SetDynamicCode(v string) *UpgradeTwoWayResponseBody {
	s.DynamicCode = &v
	return s
}

type UpgradeTwoWayResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpgradeTwoWayResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpgradeTwoWayResponse) String() string {
	return tea.Prettify(s)
}

func (s UpgradeTwoWayResponse) GoString() string {
	return s.String()
}

func (s *UpgradeTwoWayResponse) SetHeaders(v map[string]*string) *UpgradeTwoWayResponse {
	s.Headers = v
	return s
}

func (s *UpgradeTwoWayResponse) SetBody(v *UpgradeTwoWayResponseBody) *UpgradeTwoWayResponse {
	s.Body = v
	return s
}

type WhiteIpListRequest struct {
	Type              *string `json:"Type,omitempty" xml:"Type,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
	DestinationRegion *string `json:"DestinationRegion,omitempty" xml:"DestinationRegion,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s WhiteIpListRequest) String() string {
	return tea.Prettify(s)
}

func (s WhiteIpListRequest) GoString() string {
	return s.String()
}

func (s *WhiteIpListRequest) SetType(v string) *WhiteIpListRequest {
	s.Type = &v
	return s
}

func (s *WhiteIpListRequest) SetRegion(v string) *WhiteIpListRequest {
	s.Region = &v
	return s
}

func (s *WhiteIpListRequest) SetDestinationRegion(v string) *WhiteIpListRequest {
	s.DestinationRegion = &v
	return s
}

func (s *WhiteIpListRequest) SetRegionId(v string) *WhiteIpListRequest {
	s.RegionId = &v
	return s
}

type WhiteIpListResponseBody struct {
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success        *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	IpList         *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
}

func (s WhiteIpListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WhiteIpListResponseBody) GoString() string {
	return s.String()
}

func (s *WhiteIpListResponseBody) SetHttpStatusCode(v int32) *WhiteIpListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WhiteIpListResponseBody) SetRequestId(v string) *WhiteIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *WhiteIpListResponseBody) SetErrCode(v string) *WhiteIpListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *WhiteIpListResponseBody) SetSuccess(v bool) *WhiteIpListResponseBody {
	s.Success = &v
	return s
}

func (s *WhiteIpListResponseBody) SetErrMessage(v string) *WhiteIpListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *WhiteIpListResponseBody) SetIpList(v string) *WhiteIpListResponseBody {
	s.IpList = &v
	return s
}

func (s *WhiteIpListResponseBody) SetDynamicMessage(v string) *WhiteIpListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *WhiteIpListResponseBody) SetDynamicCode(v string) *WhiteIpListResponseBody {
	s.DynamicCode = &v
	return s
}

type WhiteIpListResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *WhiteIpListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s WhiteIpListResponse) String() string {
	return tea.Prettify(s)
}

func (s WhiteIpListResponse) GoString() string {
	return s.String()
}

func (s *WhiteIpListResponse) SetHeaders(v map[string]*string) *WhiteIpListResponse {
	s.Headers = v
	return s
}

func (s *WhiteIpListResponse) SetBody(v *WhiteIpListResponseBody) *WhiteIpListResponse {
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

func (client *Client) ConfigureDtsJobWithOptions(request *ConfigureDtsJobRequest, runtime *util.RuntimeOptions) (_result *ConfigureDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureDtsJob(request *ConfigureDtsJobRequest) (_result *ConfigureDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureDtsJobResponse{}
	_body, _err := client.ConfigureDtsJobWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConfigureMigrationJobAlertWithOptions(request *ConfigureMigrationJobAlertRequest, runtime *util.RuntimeOptions) (_result *ConfigureMigrationJobAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureMigrationJobAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureMigrationJobAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureMigrationJobAlert(request *ConfigureMigrationJobAlertRequest) (_result *ConfigureMigrationJobAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureMigrationJobAlertResponse{}
	_body, _err := client.ConfigureMigrationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureSubscriptionWithOptions(request *ConfigureSubscriptionRequest, runtime *util.RuntimeOptions) (_result *ConfigureSubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSubscription"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSubscription(request *ConfigureSubscriptionRequest) (_result *ConfigureSubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSubscriptionResponse{}
	_body, _err := client.ConfigureSubscriptionWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSubscriptionInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSubscriptionInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConfigureSubscriptionInstanceAlertWithOptions(request *ConfigureSubscriptionInstanceAlertRequest, runtime *util.RuntimeOptions) (_result *ConfigureSubscriptionInstanceAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSubscriptionInstanceAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSubscriptionInstanceAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSubscriptionInstanceAlert(request *ConfigureSubscriptionInstanceAlertRequest) (_result *ConfigureSubscriptionInstanceAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSubscriptionInstanceAlertResponse{}
	_body, _err := client.ConfigureSubscriptionInstanceAlertWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ConfigureSynchronizationJobAlertWithOptions(request *ConfigureSynchronizationJobAlertRequest, runtime *util.RuntimeOptions) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSynchronizationJobAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSynchronizationJobAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSynchronizationJobAlert(request *ConfigureSynchronizationJobAlertRequest) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobAlertResponse{}
	_body, _err := client.ConfigureSynchronizationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ConfigureSynchronizationJobReplicatorCompareWithOptions(request *ConfigureSynchronizationJobReplicatorCompareRequest, runtime *util.RuntimeOptions) (_result *ConfigureSynchronizationJobReplicatorCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ConfigureSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ConfigureSynchronizationJobReplicatorCompare"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ConfigureSynchronizationJobReplicatorCompare(request *ConfigureSynchronizationJobReplicatorCompareRequest) (_result *ConfigureSynchronizationJobReplicatorCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ConfigureSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.ConfigureSynchronizationJobReplicatorCompareWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerChannelWithOptions(request *CreateConsumerChannelRequest, runtime *util.RuntimeOptions) (_result *CreateConsumerChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConsumerChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConsumerChannel"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsumerChannel(request *CreateConsumerChannelRequest) (_result *CreateConsumerChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConsumerChannelResponse{}
	_body, _err := client.CreateConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateConsumerGroupWithOptions(request *CreateConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *CreateConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateConsumerGroup"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateConsumerGroup(request *CreateConsumerGroupRequest) (_result *CreateConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateConsumerGroupResponse{}
	_body, _err := client.CreateConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDtsInstanceWithOptions(request *CreateDtsInstanceRequest, runtime *util.RuntimeOptions) (_result *CreateDtsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateDtsInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateDtsInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDtsInstance(request *CreateDtsInstanceRequest) (_result *CreateDtsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDtsInstanceResponse{}
	_body, _err := client.CreateDtsInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobMonitorRuleWithOptions(request *CreateJobMonitorRuleRequest, runtime *util.RuntimeOptions) (_result *CreateJobMonitorRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateJobMonitorRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateJobMonitorRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJobMonitorRule(request *CreateJobMonitorRuleRequest) (_result *CreateJobMonitorRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateJobMonitorRuleResponse{}
	_body, _err := client.CreateJobMonitorRuleWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSubscriptionInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSubscriptionInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DeleteConsumerChannelWithOptions(request *DeleteConsumerChannelRequest, runtime *util.RuntimeOptions) (_result *DeleteConsumerChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConsumerChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConsumerChannel"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerChannel(request *DeleteConsumerChannelRequest) (_result *DeleteConsumerChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConsumerChannelResponse{}
	_body, _err := client.DeleteConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteConsumerGroupWithOptions(request *DeleteConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteConsumerGroup"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteConsumerGroup(request *DeleteConsumerGroupRequest) (_result *DeleteConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteConsumerGroupResponse{}
	_body, _err := client.DeleteConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDtsJobWithOptions(request *DeleteDtsJobRequest, runtime *util.RuntimeOptions) (_result *DeleteDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDtsJob(request *DeleteDtsJobRequest) (_result *DeleteDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteDtsJobResponse{}
	_body, _err := client.DeleteDtsJobWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSubscriptionInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSubscriptionInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeConnectionStatusWithOptions(request *DescribeConnectionStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeConnectionStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConnectionStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConnectionStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConnectionStatus(request *DescribeConnectionStatusRequest) (_result *DescribeConnectionStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConnectionStatusResponse{}
	_body, _err := client.DescribeConnectionStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConsumerChannelWithOptions(request *DescribeConsumerChannelRequest, runtime *util.RuntimeOptions) (_result *DescribeConsumerChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConsumerChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConsumerChannel"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConsumerChannel(request *DescribeConsumerChannelRequest) (_result *DescribeConsumerChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConsumerChannelResponse{}
	_body, _err := client.DescribeConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeConsumerGroupWithOptions(request *DescribeConsumerGroupRequest, runtime *util.RuntimeOptions) (_result *DescribeConsumerGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeConsumerGroupResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeConsumerGroup"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeConsumerGroup(request *DescribeConsumerGroupRequest) (_result *DescribeConsumerGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeConsumerGroupResponse{}
	_body, _err := client.DescribeConsumerGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDTSIPWithOptions(request *DescribeDTSIPRequest, runtime *util.RuntimeOptions) (_result *DescribeDTSIPResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDTSIPResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDTSIP"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDTSIP(request *DescribeDTSIPRequest) (_result *DescribeDTSIPResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDTSIPResponse{}
	_body, _err := client.DescribeDTSIPWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDtsJobDetailWithOptions(request *DescribeDtsJobDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeDtsJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDtsJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDtsJobDetail"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDtsJobDetail(request *DescribeDtsJobDetailRequest) (_result *DescribeDtsJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDtsJobDetailResponse{}
	_body, _err := client.DescribeDtsJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDtsJobsWithOptions(request *DescribeDtsJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeDtsJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeDtsJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeDtsJobs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDtsJobs(request *DescribeDtsJobsRequest) (_result *DescribeDtsJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDtsJobsResponse{}
	_body, _err := client.DescribeDtsJobsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeEndpointSwitchStatusWithOptions(request *DescribeEndpointSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeEndpointSwitchStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeEndpointSwitchStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeEndpointSwitchStatus(request *DescribeEndpointSwitchStatusRequest) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeEndpointSwitchStatusResponse{}
	_body, _err := client.DescribeEndpointSwitchStatusWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeInitializationStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeInitializationStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeJobMonitorRuleWithOptions(request *DescribeJobMonitorRuleRequest, runtime *util.RuntimeOptions) (_result *DescribeJobMonitorRuleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeJobMonitorRuleResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeJobMonitorRule"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeJobMonitorRule(request *DescribeJobMonitorRuleRequest) (_result *DescribeJobMonitorRuleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeJobMonitorRuleResponse{}
	_body, _err := client.DescribeJobMonitorRuleWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeMigrationJobAlertWithOptions(request *DescribeMigrationJobAlertRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMigrationJobAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMigrationJobAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeMigrationJobAlert(request *DescribeMigrationJobAlertRequest) (_result *DescribeMigrationJobAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeMigrationJobAlertResponse{}
	_body, _err := client.DescribeMigrationJobAlertWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMigrationJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMigrationJobDetail"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeMigrationJobsWithOptions(request *DescribeMigrationJobsRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMigrationJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMigrationJobs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeMigrationJobStatusWithOptions(request *DescribeMigrationJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeMigrationJobStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeMigrationJobStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribePreCheckStatusWithOptions(request *DescribePreCheckStatusRequest, runtime *util.RuntimeOptions) (_result *DescribePreCheckStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribePreCheckStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribePreCheckStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribePreCheckStatus(request *DescribePreCheckStatusRequest) (_result *DescribePreCheckStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribePreCheckStatusResponse{}
	_body, _err := client.DescribePreCheckStatusWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstanceAlertWithOptions(request *DescribeSubscriptionInstanceAlertRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionInstanceAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubscriptionInstanceAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubscriptionInstanceAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionInstanceAlert(request *DescribeSubscriptionInstanceAlertRequest) (_result *DescribeSubscriptionInstanceAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionInstanceAlertResponse{}
	_body, _err := client.DescribeSubscriptionInstanceAlertWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubscriptionInstancesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubscriptionInstances"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSubscriptionInstanceStatusWithOptions(request *DescribeSubscriptionInstanceStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionInstanceStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubscriptionInstanceStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubscriptionInstanceStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSubscriptionMetaWithOptions(tmpReq *DescribeSubscriptionMetaRequest, runtime *util.RuntimeOptions) (_result *DescribeSubscriptionMetaResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &DescribeSubscriptionMetaShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.SubMigrationJobIds)) {
		request.SubMigrationJobIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.SubMigrationJobIds, tea.String("SubMigrationJobIds"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Topics)) {
		request.TopicsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Topics, tea.String("Topics"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSubscriptionMetaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSubscriptionMeta"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSubscriptionMeta(request *DescribeSubscriptionMetaRequest) (_result *DescribeSubscriptionMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSubscriptionMetaResponse{}
	_body, _err := client.DescribeSubscriptionMetaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobAlertWithOptions(request *DescribeSynchronizationJobAlertRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationJobAlertResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationJobAlert"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobAlert(request *DescribeSynchronizationJobAlertRequest) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationJobAlertResponse{}
	_body, _err := client.DescribeSynchronizationJobAlertWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobReplicatorCompareWithOptions(request *DescribeSynchronizationJobReplicatorCompareRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobReplicatorCompareResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationJobReplicatorCompare"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobReplicatorCompare(request *DescribeSynchronizationJobReplicatorCompareRequest) (_result *DescribeSynchronizationJobReplicatorCompareResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationJobReplicatorCompareResponse{}
	_body, _err := client.DescribeSynchronizationJobReplicatorCompareWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationJobsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationJobs"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSynchronizationJobStatusWithOptions(request *DescribeSynchronizationJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationJobStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationJobStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) DescribeSynchronizationJobStatusListWithOptions(request *DescribeSynchronizationJobStatusListRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobStatusListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationJobStatusListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationJobStatusList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronizationJobStatusList(request *DescribeSynchronizationJobStatusListRequest) (_result *DescribeSynchronizationJobStatusListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronizationJobStatusListResponse{}
	_body, _err := client.DescribeSynchronizationJobStatusListWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DescribeSynchronizationObjectModifyStatusResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DescribeSynchronizationObjectModifyStatus"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) IgnoreJobDetailWithOptions(request *IgnoreJobDetailRequest, runtime *util.RuntimeOptions) (_result *IgnoreJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &IgnoreJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("IgnoreJobDetail"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IgnoreJobDetail(request *IgnoreJobDetailRequest) (_result *IgnoreJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &IgnoreJobDetailResponse{}
	_body, _err := client.IgnoreJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitDtsRdsInstanceWithOptions(request *InitDtsRdsInstanceRequest, runtime *util.RuntimeOptions) (_result *InitDtsRdsInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InitDtsRdsInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InitDtsRdsInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitDtsRdsInstance(request *InitDtsRdsInstanceRequest) (_result *InitDtsRdsInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InitDtsRdsInstanceResponse{}
	_body, _err := client.InitDtsRdsInstanceWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyConsumerChannelWithOptions(request *ModifyConsumerChannelRequest, runtime *util.RuntimeOptions) (_result *ModifyConsumerChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyConsumerChannelResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyConsumerChannel"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyConsumerChannel(request *ModifyConsumerChannelRequest) (_result *ModifyConsumerChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyConsumerChannelResponse{}
	_body, _err := client.ModifyConsumerChannelWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyConsumerGroupPasswordWithOptions(request *ModifyConsumerGroupPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyConsumerGroupPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyConsumerGroupPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyConsumerGroupPassword"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyConsumerGroupPassword(request *ModifyConsumerGroupPasswordRequest) (_result *ModifyConsumerGroupPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyConsumerGroupPasswordResponse{}
	_body, _err := client.ModifyConsumerGroupPasswordWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyConsumptionTimestampResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyConsumptionTimestamp"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) ModifyDtsJobWithOptions(tmpReq *ModifyDtsJobRequest, runtime *util.RuntimeOptions) (_result *ModifyDtsJobResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ModifyDtsJobShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.DbList)) {
		request.DbListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DbList, tea.String("DbList"), tea.String("json"))
	}

	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDtsJob(request *ModifyDtsJobRequest) (_result *ModifyDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDtsJobResponse{}
	_body, _err := client.ModifyDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDtsJobNameWithOptions(request *ModifyDtsJobNameRequest, runtime *util.RuntimeOptions) (_result *ModifyDtsJobNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDtsJobNameResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDtsJobName"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDtsJobName(request *ModifyDtsJobNameRequest) (_result *ModifyDtsJobNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDtsJobNameResponse{}
	_body, _err := client.ModifyDtsJobNameWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyDtsJobPasswordWithOptions(request *ModifyDtsJobPasswordRequest, runtime *util.RuntimeOptions) (_result *ModifyDtsJobPasswordResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifyDtsJobPasswordResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifyDtsJobPassword"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyDtsJobPassword(request *ModifyDtsJobPasswordRequest) (_result *ModifyDtsJobPasswordResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifyDtsJobPasswordResponse{}
	_body, _err := client.ModifyDtsJobPasswordWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifySubscriptionWithOptions(request *ModifySubscriptionRequest, runtime *util.RuntimeOptions) (_result *ModifySubscriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySubscriptionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySubscription"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifySubscription(request *ModifySubscriptionRequest) (_result *ModifySubscriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ModifySubscriptionResponse{}
	_body, _err := client.ModifySubscriptionWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySubscriptionObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySubscriptionObject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ModifySynchronizationObjectResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ModifySynchronizationObject"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) RenewInstanceWithOptions(request *RenewInstanceRequest, runtime *util.RuntimeOptions) (_result *RenewInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenewInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenewInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstance(request *RenewInstanceRequest) (_result *RenewInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstanceResponse{}
	_body, _err := client.RenewInstanceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetDtsJobWithOptions(request *ResetDtsJobRequest, runtime *util.RuntimeOptions) (_result *ResetDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetDtsJob(request *ResetDtsJobRequest) (_result *ResetDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetDtsJobResponse{}
	_body, _err := client.ResetDtsJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetSynchronizationJobWithOptions(request *ResetSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *ResetSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ResetSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ResetSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetSynchronizationJob(request *ResetSynchronizationJobRequest) (_result *ResetSynchronizationJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetSynchronizationJobResponse{}
	_body, _err := client.ResetSynchronizationJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ShieldPrecheckWithOptions(request *ShieldPrecheckRequest, runtime *util.RuntimeOptions) (_result *ShieldPrecheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ShieldPrecheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ShieldPrecheck"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ShieldPrecheck(request *ShieldPrecheckRequest) (_result *ShieldPrecheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ShieldPrecheckResponse{}
	_body, _err := client.ShieldPrecheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SkipPreCheckWithOptions(request *SkipPreCheckRequest, runtime *util.RuntimeOptions) (_result *SkipPreCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SkipPreCheckResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SkipPreCheck"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SkipPreCheck(request *SkipPreCheckRequest) (_result *SkipPreCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SkipPreCheckResponse{}
	_body, _err := client.SkipPreCheckWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartDtsJobWithOptions(request *StartDtsJobRequest, runtime *util.RuntimeOptions) (_result *StartDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartDtsJob(request *StartDtsJobRequest) (_result *StartDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartDtsJobResponse{}
	_body, _err := client.StartDtsJobWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartSubscriptionInstanceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartSubscriptionInstance"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StartSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StartSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) StopDtsJobWithOptions(request *StopDtsJobRequest, runtime *util.RuntimeOptions) (_result *StopDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopDtsJob(request *StopDtsJobRequest) (_result *StopDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopDtsJobResponse{}
	_body, _err := client.StopDtsJobWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &StopMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("StopMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SummaryJobDetailWithOptions(request *SummaryJobDetailRequest, runtime *util.RuntimeOptions) (_result *SummaryJobDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SummaryJobDetailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SummaryJobDetail"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SummaryJobDetail(request *SummaryJobDetailRequest) (_result *SummaryJobDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SummaryJobDetailResponse{}
	_body, _err := client.SummaryJobDetailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SuspendDtsJobWithOptions(request *SuspendDtsJobRequest, runtime *util.RuntimeOptions) (_result *SuspendDtsJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendDtsJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendDtsJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendDtsJob(request *SuspendDtsJobRequest) (_result *SuspendDtsJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendDtsJobResponse{}
	_body, _err := client.SuspendDtsJobWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendMigrationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendMigrationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SuspendSynchronizationJobResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SuspendSynchronizationJob"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) SwitchSynchronizationEndpointWithOptions(request *SwitchSynchronizationEndpointRequest, runtime *util.RuntimeOptions) (_result *SwitchSynchronizationEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SwitchSynchronizationEndpointResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SwitchSynchronizationEndpoint"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SwitchSynchronizationEndpoint(request *SwitchSynchronizationEndpointRequest) (_result *SwitchSynchronizationEndpointResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SwitchSynchronizationEndpointResponse{}
	_body, _err := client.SwitchSynchronizationEndpointWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) TransferInstanceClassWithOptions(request *TransferInstanceClassRequest, runtime *util.RuntimeOptions) (_result *TransferInstanceClassResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferInstanceClassResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferInstanceClass"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferInstanceClass(request *TransferInstanceClassRequest) (_result *TransferInstanceClassResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferInstanceClassResponse{}
	_body, _err := client.TransferInstanceClassWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TransferPayTypeWithOptions(request *TransferPayTypeRequest, runtime *util.RuntimeOptions) (_result *TransferPayTypeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TransferPayTypeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TransferPayType"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TransferPayType(request *TransferPayTypeRequest) (_result *TransferPayTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TransferPayTypeResponse{}
	_body, _err := client.TransferPayTypeWithOptions(request, runtime)
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
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UntagResources"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
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

func (client *Client) UpgradeTwoWayWithOptions(request *UpgradeTwoWayRequest, runtime *util.RuntimeOptions) (_result *UpgradeTwoWayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &UpgradeTwoWayResponse{}
	_body, _err := client.DoRPCRequest(tea.String("UpgradeTwoWay"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpgradeTwoWay(request *UpgradeTwoWayRequest) (_result *UpgradeTwoWayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpgradeTwoWayResponse{}
	_body, _err := client.UpgradeTwoWayWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WhiteIpListWithOptions(request *WhiteIpListRequest, runtime *util.RuntimeOptions) (_result *WhiteIpListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &WhiteIpListResponse{}
	_body, _err := client.DoRPCRequest(tea.String("WhiteIpList"), tea.String("2020-01-01"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WhiteIpList(request *WhiteIpListRequest) (_result *WhiteIpListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &WhiteIpListResponse{}
	_body, _err := client.WhiteIpListWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
