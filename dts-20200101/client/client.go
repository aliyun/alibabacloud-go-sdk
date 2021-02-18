// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ConfigureMigrationJobRequest struct {
	SourceEndpoint      *ConfigureMigrationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	DestinationEndpoint *ConfigureMigrationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode       *ConfigureMigrationJobRequestMigrationMode       `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureMigrationJobResponseBody) SetErrMessage(v string) *ConfigureMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetSuccess(v string) *ConfigureMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobResponseBody) SetErrCode(v string) *ConfigureMigrationJobResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureMigrationJobAlertResponseBody) SetErrMessage(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetSuccess(v string) *ConfigureMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureMigrationJobAlertResponseBody) SetErrCode(v string) *ConfigureMigrationJobAlertResponseBody {
	s.ErrCode = &v
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

type ConfigureSubscriptionInstanceRequest struct {
	SourceEndpoint                  *ConfigureSubscriptionInstanceRequestSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType            *ConfigureSubscriptionInstanceRequestSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstance            *ConfigureSubscriptionInstanceRequestSubscriptionInstance `json:"SubscriptionInstance,omitempty" xml:"SubscriptionInstance,omitempty" type:"Struct"`
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *ConfigureSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureSynchronizationJobResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
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

type CreateConsumerGroupRequest struct {
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
	ErrMessage      *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success         *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode         *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *CreateConsumerGroupResponseBody) SetErrMessage(v string) *CreateConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetSuccess(v string) *CreateConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateConsumerGroupResponseBody) SetErrCode(v string) *CreateConsumerGroupResponseBody {
	s.ErrCode = &v
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

type CreateMigrationJobRequest struct {
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
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *CreateMigrationJobResponseBody) SetErrMessage(v string) *CreateMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetSuccess(v string) *CreateMigrationJobResponseBody {
	s.Success = &v
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
	ErrMessage             *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *CreateSubscriptionInstanceResponseBody) SetErrMessage(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetSuccess(v string) *CreateSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSubscriptionInstanceResponseBody) SetErrCode(v string) *CreateSubscriptionInstanceResponseBody {
	s.ErrCode = &v
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
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	RequestId            *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrMessage           *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success              *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode              *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s CreateSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponseBody) SetSynchronizationJobId(v string) *CreateSynchronizationJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetRequestId(v string) *CreateSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetErrMessage(v string) *CreateSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetSuccess(v string) *CreateSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetErrCode(v string) *CreateSynchronizationJobResponseBody {
	s.ErrCode = &v
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

type DeleteConsumerGroupRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DeleteConsumerGroupResponseBody) SetErrMessage(v string) *DeleteConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetSuccess(v string) *DeleteConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) SetErrCode(v string) *DeleteConsumerGroupResponseBody {
	s.ErrCode = &v
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

type DeleteMigrationJobRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DeleteMigrationJobResponseBody) SetErrMessage(v string) *DeleteMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetSuccess(v string) *DeleteMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetErrCode(v string) *DeleteMigrationJobResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DeleteSubscriptionInstanceResponseBody) SetErrMessage(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetSuccess(v string) *DeleteSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponseBody) SetErrCode(v string) *DeleteSubscriptionInstanceResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DeleteSynchronizationJobResponseBody) SetErrMessage(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetSuccess(v string) *DeleteSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetErrCode(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrCode = &v
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
	DestinationEndpointOracleSID    *bool   `json:"DestinationEndpointOracleSID,omitempty" xml:"DestinationEndpointOracleSID,omitempty"`
	DestinationEndpointArchitecture *bool   `json:"DestinationEndpointArchitecture,omitempty" xml:"DestinationEndpointArchitecture,omitempty"`
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

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointOracleSID(v bool) *DescribeConnectionStatusRequest {
	s.DestinationEndpointOracleSID = &v
	return s
}

func (s *DescribeConnectionStatusRequest) SetDestinationEndpointArchitecture(v bool) *DescribeConnectionStatusRequest {
	s.DestinationEndpointArchitecture = &v
	return s
}

type DescribeConnectionStatusResponseBody struct {
	SourceConnectionStatus      map[string]interface{} `json:"SourceConnectionStatus,omitempty" xml:"SourceConnectionStatus,omitempty"`
	RequestId                   *string                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DestinationConnectionStatus map[string]interface{} `json:"DestinationConnectionStatus,omitempty" xml:"DestinationConnectionStatus,omitempty"`
	ErrMessage                  *string                `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                     *string                `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                     *string                `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DescribeConnectionStatusResponseBody) SetRequestId(v string) *DescribeConnectionStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetDestinationConnectionStatus(v map[string]interface{}) *DescribeConnectionStatusResponseBody {
	s.DestinationConnectionStatus = v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrMessage(v string) *DescribeConnectionStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetSuccess(v string) *DescribeConnectionStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConnectionStatusResponseBody) SetErrCode(v string) *DescribeConnectionStatusResponseBody {
	s.ErrCode = &v
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

type DescribeConsumerGroupRequest struct {
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
	TotalRecordCount *int32                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ErrMessage       *string                                            `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ConsumerChannels *DescribeConsumerGroupResponseBodyConsumerChannels `json:"ConsumerChannels,omitempty" xml:"ConsumerChannels,omitempty" type:"Struct"`
	Success          *string                                            `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode          *string                                            `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeConsumerGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupResponseBody) SetTotalRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageRecordCount(v int32) *DescribeConsumerGroupResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetRequestId(v string) *DescribeConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetPageNumber(v int32) *DescribeConsumerGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrMessage(v string) *DescribeConsumerGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetConsumerChannels(v *DescribeConsumerGroupResponseBodyConsumerChannels) *DescribeConsumerGroupResponseBody {
	s.ConsumerChannels = v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetSuccess(v string) *DescribeConsumerGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeConsumerGroupResponseBody) SetErrCode(v string) *DescribeConsumerGroupResponseBody {
	s.ErrCode = &v
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
	SourceEndpointRegion      *string `json:"SourceEndpointRegion,omitempty" xml:"SourceEndpointRegion,omitempty"`
	DestinationEndpointRegion *string `json:"DestinationEndpointRegion,omitempty" xml:"DestinationEndpointRegion,omitempty"`
}

func (s DescribeDTSIPRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDTSIPRequest) GoString() string {
	return s.String()
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
	DynamicCode    *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success        *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DescribeDTSIPResponseBody) SetDynamicCode(v string) *DescribeDTSIPResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetDynamicMessage(v string) *DescribeDTSIPResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrMessage(v string) *DescribeDTSIPResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetSuccess(v string) *DescribeDTSIPResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrCode(v string) *DescribeDTSIPResponseBody {
	s.ErrCode = &v
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

type DescribeEndpointSwitchStatusRequest struct {
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
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrMessage   *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success      *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode      *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *DescribeEndpointSwitchStatusResponseBody) SetRequestId(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrorMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetSuccess(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrCode(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrCode = &v
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
	StructureInitializationDetails []*DescribeInitializationStatusResponseBodyStructureInitializationDetails `json:"StructureInitializationDetails,omitempty" xml:"StructureInitializationDetails,omitempty" type:"Repeated"`
	RequestId                      *string                                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DataInitializationDetails      []*DescribeInitializationStatusResponseBodyDataInitializationDetails      `json:"DataInitializationDetails,omitempty" xml:"DataInitializationDetails,omitempty" type:"Repeated"`
	ErrMessage                     *string                                                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                        *string                                                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                        *string                                                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	DataSynchronizationDetails     []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails     `json:"DataSynchronizationDetails,omitempty" xml:"DataSynchronizationDetails,omitempty" type:"Repeated"`
}

func (s DescribeInitializationStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInitializationStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInitializationStatusResponseBody) SetStructureInitializationDetails(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.StructureInitializationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetRequestId(v string) *DescribeInitializationStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetDataInitializationDetails(v []*DescribeInitializationStatusResponseBodyDataInitializationDetails) *DescribeInitializationStatusResponseBody {
	s.DataInitializationDetails = v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrMessage(v string) *DescribeInitializationStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetSuccess(v string) *DescribeInitializationStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetErrCode(v string) *DescribeInitializationStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeInitializationStatusResponseBody) SetDataSynchronizationDetails(v []*DescribeInitializationStatusResponseBodyDataSynchronizationDetails) *DescribeInitializationStatusResponseBody {
	s.DataSynchronizationDetails = v
	return s
}

type DescribeInitializationStatusResponseBodyStructureInitializationDetails struct {
	Status                 *string                                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	SourceOwnerDBName      *string                                                                              `json:"SourceOwnerDBName,omitempty" xml:"SourceOwnerDBName,omitempty"`
	ObjectDefinition       *string                                                                              `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string                                                                              `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string                                                                              `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Constraints            []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints `json:"Constraints,omitempty" xml:"Constraints,omitempty" type:"Repeated"`
	DestinationOwnerDBName *string                                                                              `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string                                                                              `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
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

func (s *DescribeInitializationStatusResponseBodyStructureInitializationDetails) SetConstraints(v []*DescribeInitializationStatusResponseBodyStructureInitializationDetailsConstraints) *DescribeInitializationStatusResponseBodyStructureInitializationDetails {
	s.Constraints = v
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

type DescribeMigrationJobAlertRequest struct {
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
	ErrorAlertPhone  *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	RequestId        *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DelayAlertPhone  *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	ErrorAlertStatus *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrMessage       *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DelayAlertStatus *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	Success          *string `json:"Success,omitempty" xml:"Success,omitempty"`
	DelayOverSeconds *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	ErrCode          *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	MigrationJobId   *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
}

func (s DescribeMigrationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetRequestId(v string) *DescribeMigrationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertPhone = &v
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

func (s *DescribeMigrationJobAlertResponseBody) SetErrMessage(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetSuccess(v string) *DescribeMigrationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeMigrationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetErrCode(v string) *DescribeMigrationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobAlertResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobAlertResponseBody {
	s.MigrationJobId = &v
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
	TotalRecordCount                  *int64                                                                   `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount                   *int32                                                                   `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	DataSynchronizationDetailList     *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList     `json:"DataSynchronizationDetailList,omitempty" xml:"DataSynchronizationDetailList,omitempty" type:"Struct"`
	RequestId                         *string                                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber                        *int32                                                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	DataInitializationDetailList      *DescribeMigrationJobDetailResponseBodyDataInitializationDetailList      `json:"DataInitializationDetailList,omitempty" xml:"DataInitializationDetailList,omitempty" type:"Struct"`
	ErrMessage                        *string                                                                  `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                           *string                                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
	StructureInitializationDetailList *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList `json:"StructureInitializationDetailList,omitempty" xml:"StructureInitializationDetailList,omitempty" type:"Struct"`
	ErrCode                           *string                                                                  `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeMigrationJobDetailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobDetailResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobDetailResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobDetailResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetDataSynchronizationDetailList(v *DescribeMigrationJobDetailResponseBodyDataSynchronizationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.DataSynchronizationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetRequestId(v string) *DescribeMigrationJobDetailResponseBody {
	s.RequestId = &v
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

func (s *DescribeMigrationJobDetailResponseBody) SetErrMessage(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetSuccess(v string) *DescribeMigrationJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetStructureInitializationDetailList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailList) *DescribeMigrationJobDetailResponseBody {
	s.StructureInitializationDetailList = v
	return s
}

func (s *DescribeMigrationJobDetailResponseBody) SetErrCode(v string) *DescribeMigrationJobDetailResponseBody {
	s.ErrCode = &v
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
	ConstraintList         *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList `json:"ConstraintList,omitempty" xml:"ConstraintList,omitempty" type:"Struct"`
	ObjectDefinition       *string                                                                                                             `json:"ObjectDefinition,omitempty" xml:"ObjectDefinition,omitempty"`
	ObjectType             *string                                                                                                             `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	ErrorMessage           *string                                                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	DestinationOwnerDBName *string                                                                                                             `json:"DestinationOwnerDBName,omitempty" xml:"DestinationOwnerDBName,omitempty"`
	ObjectName             *string                                                                                                             `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
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

func (s *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail) SetConstraintList(v *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetailConstraintList) *DescribeMigrationJobDetailResponseBodyStructureInitializationDetailListStructureInitializationDetail {
	s.ConstraintList = v
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
	TotalRecordCount *int64                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount  *int32                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId        *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber       *int32                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	MigrationJobs    *DescribeMigrationJobsResponseBodyMigrationJobs `json:"MigrationJobs,omitempty" xml:"MigrationJobs,omitempty" type:"Struct"`
	ErrMessage       *string                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success          *string                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode          *string                                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeMigrationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeMigrationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetPageRecordCount(v int32) *DescribeMigrationJobsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetRequestId(v string) *DescribeMigrationJobsResponseBody {
	s.RequestId = &v
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

func (s *DescribeMigrationJobsResponseBody) SetErrMessage(v string) *DescribeMigrationJobsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetSuccess(v string) *DescribeMigrationJobsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobsResponseBody) SetErrCode(v string) *DescribeMigrationJobsResponseBody {
	s.ErrCode = &v
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
	Precheck                *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck                `json:"Precheck,omitempty" xml:"Precheck,omitempty" type:"Struct"`
	MigrationJobName        *string                                                                            `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	StructureInitialization *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty" type:"Struct"`
	PayType                 *string                                                                            `json:"PayType,omitempty" xml:"PayType,omitempty"`
	Tags                    *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags                    `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	MigrationObject         *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject         `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty" type:"Struct"`
	DataSynchronization     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization     `json:"DataSynchronization,omitempty" xml:"DataSynchronization,omitempty" type:"Struct"`
	DestinationEndpoint     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDestinationEndpoint     `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationJobStatus      *string                                                                            `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	SourceEndpoint          *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint          `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	MigrationMode           *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode           `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	MigrationJobClass       *string                                                                            `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	MigrationJobID          *string                                                                            `json:"MigrationJobID,omitempty" xml:"MigrationJobID,omitempty"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPrecheck(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobPrecheck) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Precheck = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobName = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetStructureInitialization(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.StructureInitialization = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetPayType(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.PayType = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetTags(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobTags) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.Tags = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationObject(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObject) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationObject = v
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationJobStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetSourceEndpoint(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobSourceEndpoint) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob) SetMigrationMode(v *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationMode) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJob {
	s.MigrationMode = v
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

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization struct {
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataInitialization {
	s.ErrorMessage = &v
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

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization struct {
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status       *string `json:"status,omitempty" xml:"status,omitempty"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobStructureInitialization {
	s.ErrorMessage = &v
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
	WholeDatabase *string                                                                                              `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
	DatabaseName  *string                                                                                              `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetWholeDatabase(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject) SetDatabaseName(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobMigrationObjectSynchronousObject {
	s.DatabaseName = &v
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

type DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization struct {
	Delay        *string `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Percent      *string `json:"Percent,omitempty" xml:"Percent,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
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

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetPercent(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Percent = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetErrorMessage(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization) SetStatus(v string) *DescribeMigrationJobsResponseBodyMigrationJobsMigrationJobDataSynchronization {
	s.Status = &v
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
	DataInitializationStatus      *DescribeMigrationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	RequestId                     *string                                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	MigrationJobName              *string                                                              `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	PayType                       *string                                                              `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ErrMessage                    *string                                                              `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	MigrationJobStatus            *string                                                              `json:"MigrationJobStatus,omitempty" xml:"MigrationJobStatus,omitempty"`
	Success                       *string                                                              `json:"Success,omitempty" xml:"Success,omitempty"`
	MigrationMode                 *DescribeMigrationJobStatusResponseBodyMigrationMode                 `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	ErrCode                       *string                                                              `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	MigrationJobId                *string                                                              `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	PrecheckStatus                *DescribeMigrationJobStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	MigrationObject               *string                                                              `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
	DestinationEndpoint           *DescribeMigrationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationJobClass             *string                                                              `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
	SourceEndpoint                *DescribeMigrationJobStatusResponseBodySourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
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

func (s *DescribeMigrationJobStatusResponseBody) SetRequestId(v string) *DescribeMigrationJobStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobName(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobName = &v
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

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobStatus(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobStatus = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetSuccess(v string) *DescribeMigrationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationMode(v *DescribeMigrationJobStatusResponseBodyMigrationMode) *DescribeMigrationJobStatusResponseBody {
	s.MigrationMode = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetErrCode(v string) *DescribeMigrationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationJobId(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetPrecheckStatus(v *DescribeMigrationJobStatusResponseBodyPrecheckStatus) *DescribeMigrationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetMigrationObject(v string) *DescribeMigrationJobStatusResponseBody {
	s.MigrationObject = &v
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

func (s *DescribeMigrationJobStatusResponseBody) SetSourceEndpoint(v *DescribeMigrationJobStatusResponseBodySourceEndpoint) *DescribeMigrationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeMigrationJobStatusResponseBodyStructureInitializationStatus) *DescribeMigrationJobStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeMigrationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeMigrationJobStatusResponseBodyDataSynchronizationStatus) *DescribeMigrationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
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

type DescribeMigrationJobStatusResponseBodyMigrationMode struct {
	DataInitialization      *bool `json:"dataInitialization,omitempty" xml:"dataInitialization,omitempty"`
	StructureInitialization *bool `json:"structureInitialization,omitempty" xml:"structureInitialization,omitempty"`
	DataSynchronization     *bool `json:"dataSynchronization,omitempty" xml:"dataSynchronization,omitempty"`
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

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetStructureInitialization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.StructureInitialization = &v
	return s
}

func (s *DescribeMigrationJobStatusResponseBodyMigrationMode) SetDataSynchronization(v bool) *DescribeMigrationJobStatusResponseBodyMigrationMode {
	s.DataSynchronization = &v
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

type DescribeSubscriptionInstanceAlertRequest struct {
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
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	SubscriptionInstanceName *string `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrMessage               *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SubscriptionInstanceID   *string `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	Success                  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	ErrCode                  *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeSubscriptionInstanceAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetRequestId(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.SubscriptionInstanceID = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetSuccess(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSubscriptionInstanceAlertResponseBody) SetErrCode(v string) *DescribeSubscriptionInstanceAlertResponseBody {
	s.ErrCode = &v
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
	TotalRecordCount      *int64                                                          `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount       *int32                                                          `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId             *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber            *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ErrMessage            *string                                                         `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success               *string                                                         `json:"Success,omitempty" xml:"Success,omitempty"`
	SubscriptionInstances *DescribeSubscriptionInstancesResponseBodySubscriptionInstances `json:"SubscriptionInstances,omitempty" xml:"SubscriptionInstances,omitempty" type:"Struct"`
	ErrCode               *string                                                         `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBody) SetTotalRecordCount(v int64) *DescribeSubscriptionInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageRecordCount(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetRequestId(v string) *DescribeSubscriptionInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetPageNumber(v int32) *DescribeSubscriptionInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSuccess(v string) *DescribeSubscriptionInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetSubscriptionInstances(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstances) *DescribeSubscriptionInstancesResponseBody {
	s.SubscriptionInstances = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBody) SetErrCode(v string) *DescribeSubscriptionInstancesResponseBody {
	s.ErrCode = &v
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
	Tags                     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags                 `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	ConsumptionClient        *string                                                                                                 `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	SubscriptionObject       *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
	SubscriptionHost         *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost     `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
	EndTimestamp             *string                                                                                                 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	ConsumptionCheckpoint    *string                                                                                                 `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	SubscribeTopic           *string                                                                                                 `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	BeginTimestamp           *string                                                                                                 `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	SourceEndpoint           *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionInstanceName *string                                                                                                 `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	SubscriptionInstanceID   *string                                                                                                 `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
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

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetTags(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceTags) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.Tags = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionClient(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionClient = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionObject(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObject) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionObject = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionHost(v *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionHost = v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetEndTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetConsumptionCheckpoint(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.ConsumptionCheckpoint = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscribeTopic(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscribeTopic = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetBeginTimestamp(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.BeginTimestamp = &v
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

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceName = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance) SetSubscriptionInstanceID(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstance {
	s.SubscriptionInstanceID = &v
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
	WholeDatabase *string                                                                                                                         `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
	DatabaseName  *string                                                                                                                         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetWholeDatabase(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.WholeDatabase = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject) SetDatabaseName(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionObjectSynchronousObject {
	s.DatabaseName = &v
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

type DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost struct {
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	VPCHost     *string `json:"VPCHost,omitempty" xml:"VPCHost,omitempty"`
	PublicHost  *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPrivateHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetVPCHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.VPCHost = &v
	return s
}

func (s *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost) SetPublicHost(v string) *DescribeSubscriptionInstancesResponseBodySubscriptionInstancesSubscriptionInstanceSubscriptionHost {
	s.PublicHost = &v
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
	BeginTimestamp           *string                                                             `json:"BeginTimestamp,omitempty" xml:"BeginTimestamp,omitempty"`
	ConsumptionCheckpoint    *string                                                             `json:"ConsumptionCheckpoint,omitempty" xml:"ConsumptionCheckpoint,omitempty"`
	ConsumptionClient        *string                                                             `json:"ConsumptionClient,omitempty" xml:"ConsumptionClient,omitempty"`
	EndTimestamp             *string                                                             `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	ErrMessage               *string                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	PayType                  *string                                                             `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RequestId                *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status                   *string                                                             `json:"Status,omitempty" xml:"Status,omitempty"`
	SubscribeTopic           *string                                                             `json:"SubscribeTopic,omitempty" xml:"SubscribeTopic,omitempty"`
	SubscriptionInstanceID   *string                                                             `json:"SubscriptionInstanceID,omitempty" xml:"SubscriptionInstanceID,omitempty"`
	SubscriptionInstanceName *string                                                             `json:"SubscriptionInstanceName,omitempty" xml:"SubscriptionInstanceName,omitempty"`
	ErrCode                  *string                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	Success                  *string                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrorMessage             *string                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	SubscriptionObject       *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject   `json:"SubscriptionObject,omitempty" xml:"SubscriptionObject,omitempty" type:"Struct"`
	SourceEndpoint           *DescribeSubscriptionInstanceStatusResponseBodySourceEndpoint       `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SubscriptionDataType     *DescribeSubscriptionInstanceStatusResponseBodySubscriptionDataType `json:"SubscriptionDataType,omitempty" xml:"SubscriptionDataType,omitempty" type:"Struct"`
	SubscriptionHost         *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost     `json:"SubscriptionHost,omitempty" xml:"SubscriptionHost,omitempty" type:"Struct"`
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

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrMessage(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetPayType(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.PayType = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetRequestId(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetStatus(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscribeTopic(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscribeTopic = &v
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

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrCode(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSuccess(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetErrorMessage(v string) *DescribeSubscriptionInstanceStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBody) SetSubscriptionObject(v *DescribeSubscriptionInstanceStatusResponseBodySubscriptionObject) *DescribeSubscriptionInstanceStatusResponseBody {
	s.SubscriptionObject = v
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

type DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost struct {
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	PublicHost  *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	VPCHost     *string `json:"VPCHost,omitempty" xml:"VPCHost,omitempty"`
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) String() string {
	return tea.Prettify(s)
}

func (s DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) GoString() string {
	return s.String()
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetPrivateHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.PrivateHost = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetPublicHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.PublicHost = &v
	return s
}

func (s *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost) SetVPCHost(v string) *DescribeSubscriptionInstanceStatusResponseBodySubscriptionHost {
	s.VPCHost = &v
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

type DescribeSynchronizationJobAlertRequest struct {
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
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName   *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	ErrMessage               *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	Success                  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	ErrCode                  *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeSynchronizationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetRequestId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
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

func (s *DescribeSynchronizationJobAlertResponseBody) SetSuccess(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrCode(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
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
	RequestId                              *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SynchronizationReplicatorCompareEnable *bool   `json:"SynchronizationReplicatorCompareEnable,omitempty" xml:"SynchronizationReplicatorCompareEnable,omitempty"`
	ErrMessage                             *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                                *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                                *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobReplicatorCompareResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetRequestId(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSynchronizationReplicatorCompareEnable(v bool) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.SynchronizationReplicatorCompareEnable = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetSuccess(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobReplicatorCompareResponseBody) SetErrCode(v string) *DescribeSynchronizationJobReplicatorCompareResponseBody {
	s.ErrCode = &v
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
	TotalRecordCount         *int64                                                             `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	SynchronizationInstances []*DescribeSynchronizationJobsResponseBodySynchronizationInstances `json:"SynchronizationInstances,omitempty" xml:"SynchronizationInstances,omitempty" type:"Repeated"`
	PageRecordCount          *int32                                                             `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId                *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber               *int32                                                             `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBody) SetSynchronizationInstances(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstances) *DescribeSynchronizationJobsResponseBody {
	s.SynchronizationInstances = v
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

func (s *DescribeSynchronizationJobsResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobsResponseBody {
	s.PageNumber = &v
	return s
}

type DescribeSynchronizationJobsResponseBodySynchronizationInstances struct {
	SynchronizationJobName        *string                                                                                       `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	Status                        *string                                                                                       `json:"Status,omitempty" xml:"Status,omitempty"`
	DataInitialization            *string                                                                                       `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	Performance                   *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	Delay                         *string                                                                                       `json:"Delay,omitempty" xml:"Delay,omitempty"`
	PrecheckStatus                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	StructureInitializationStatus *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	ErrorMessage                  *string                                                                                       `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                                       `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	SynchronizationObjects        []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
	CreateTime                    *string                                                                                       `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	PayType                       *string                                                                                       `json:"PayType,omitempty" xml:"PayType,omitempty"`
	StructureInitialization       *string                                                                                       `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	SynchronizationJobClass       *string                                                                                       `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	Tags                          []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags                        `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	DataInitializationStatus      *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	DestinationEndpoint           *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	SynchronizationJobId          *string                                                                                       `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationDirection      *string                                                                                       `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstances) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationJobName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStatus(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitialization(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPerformance(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPerformance) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDelay(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetPrecheckStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesPrecheckStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetStructureInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesStructureInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.StructureInitializationStatus = v
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationObjects(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationObjects = v
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetTags(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesTags) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.Tags = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataInitializationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataInitializationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDestinationEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetDataSynchronizationStatus(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDataSynchronizationStatus) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSourceEndpoint(v *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SourceEndpoint = v
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

type DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects struct {
	NewSchemaName *string                                                                                               `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	TableIncludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
	TableExcludes []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	SchemaName    *string                                                                                               `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableIncludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableIncludes = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjectsTableExcludes) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSynchronizationObjects {
	s.SchemaName = &v
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
	DataInitializationStatus      *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	SynchronizationObjects        []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects      `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty" type:"Repeated"`
	Delay                         *string                                                                    `json:"Delay,omitempty" xml:"Delay,omitempty"`
	Success                       *string                                                                    `json:"Success,omitempty" xml:"Success,omitempty"`
	DelayMillis                   *int64                                                                     `json:"DelayMillis,omitempty" xml:"DelayMillis,omitempty"`
	DataInitialization            *string                                                                    `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	SynchronizationJobClass       *string                                                                    `json:"SynchronizationJobClass,omitempty" xml:"SynchronizationJobClass,omitempty"`
	DataSynchronizationStatus     *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
	Status                        *string                                                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	SynchronizationJobName        *string                                                                    `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	RequestId                     *string                                                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PayType                       *string                                                                    `json:"PayType,omitempty" xml:"PayType,omitempty"`
	ErrMessage                    *string                                                                    `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ErrCode                       *string                                                                    `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	PrecheckStatus                *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	SynchronizationJobId          *string                                                                    `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	Checkpoint                    *string                                                                    `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DestinationEndpoint           *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint           `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	SourceEndpoint                *DescribeSynchronizationJobStatusResponseBodySourceEndpoint                `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	StructureInitialization       *string                                                                    `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	ErrorMessage                  *string                                                                    `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ExpireTime                    *string                                                                    `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Performance                   *DescribeSynchronizationJobStatusResponseBodyPerformance                   `json:"Performance,omitempty" xml:"Performance,omitempty" type:"Struct"`
	SynchronizationDirection      *string                                                                    `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	StructureInitializationStatus *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
}

func (s DescribeSynchronizationJobStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationObjects(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationObjects = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelay(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Delay = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDelayMillis(v int64) *DescribeSynchronizationJobStatusResponseBody {
	s.DelayMillis = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.DataInitialization = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobClass(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobClass = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationJobStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.DataSynchronizationStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStatus(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobName(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.RequestId = &v
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationJobStatusResponseBodyPrecheckStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationJobId(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetCheckpoint(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.Checkpoint = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetDestinationEndpoint(v *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.DestinationEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSourceEndpoint(v *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) *DescribeSynchronizationJobStatusResponseBody {
	s.SourceEndpoint = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitialization(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitialization = &v
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetPerformance(v *DescribeSynchronizationJobStatusResponseBodyPerformance) *DescribeSynchronizationJobStatusResponseBody {
	s.Performance = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationJobStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationJobStatusResponseBody {
	s.StructureInitializationStatus = v
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

type DescribeSynchronizationJobStatusResponseBodySynchronizationObjects struct {
	NewSchemaName *string                                                                            `json:"NewSchemaName,omitempty" xml:"NewSchemaName,omitempty"`
	TableIncludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes `json:"TableIncludes,omitempty" xml:"TableIncludes,omitempty" type:"Repeated"`
	TableExcludes []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes `json:"TableExcludes,omitempty" xml:"TableExcludes,omitempty" type:"Repeated"`
	SchemaName    *string                                                                            `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
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

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableIncludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableIncludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableIncludes = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetTableExcludes(v []*DescribeSynchronizationJobStatusResponseBodySynchronizationObjectsTableExcludes) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.TableExcludes = v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects) SetSchemaName(v string) *DescribeSynchronizationJobStatusResponseBodySynchronizationObjects {
	s.SchemaName = &v
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
	TotalRecordCount                 *int64                                                                              `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
	PageRecordCount                  *int32                                                                              `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	RequestId                        *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PageNumber                       *int32                                                                              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	ErrMessage                       *string                                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	SynchronizationJobListStatusList []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList `json:"SynchronizationJobListStatusList,omitempty" xml:"SynchronizationJobListStatusList,omitempty" type:"Repeated"`
	Success                          *string                                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                          *string                                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s DescribeSynchronizationJobStatusListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronizationJobStatusListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageRecordCount(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetRequestId(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetPageNumber(v int32) *DescribeSynchronizationJobStatusListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSynchronizationJobListStatusList(v []*DescribeSynchronizationJobStatusListResponseBodySynchronizationJobListStatusList) *DescribeSynchronizationJobStatusListResponseBody {
	s.SynchronizationJobListStatusList = v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetSuccess(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobStatusListResponseBody) SetErrCode(v string) *DescribeSynchronizationJobStatusListResponseBody {
	s.ErrCode = &v
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
	PrecheckStatus                *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus                `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty" type:"Struct"`
	DataInitializationStatus      *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus      `json:"DataInitializationStatus,omitempty" xml:"DataInitializationStatus,omitempty" type:"Struct"`
	RequestId                     *string                                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrorMessage                  *string                                                                             `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	ErrMessage                    *string                                                                             `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success                       *string                                                                             `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode                       *string                                                                             `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	StructureInitializationStatus *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus `json:"StructureInitializationStatus,omitempty" xml:"StructureInitializationStatus,omitempty" type:"Struct"`
	DataSynchronizationStatus     *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus     `json:"DataSynchronizationStatus,omitempty" xml:"DataSynchronizationStatus,omitempty" type:"Struct"`
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

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetPrecheckStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyPrecheckStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.PrecheckStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetRequestId(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrorMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrMessage(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetSuccess(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetErrCode(v string) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetStructureInitializationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyStructureInitializationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.StructureInitializationStatus = v
	return s
}

func (s *DescribeSynchronizationObjectModifyStatusResponseBody) SetDataSynchronizationStatus(v *DescribeSynchronizationObjectModifyStatusResponseBodyDataSynchronizationStatus) *DescribeSynchronizationObjectModifyStatusResponseBody {
	s.DataSynchronizationStatus = v
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
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
	ErrMessage   *string                                   `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success      *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode      *string                                   `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrMessage(v string) *ListTagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetSuccess(v bool) *ListTagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetErrCode(v string) *ListTagResourcesResponseBody {
	s.ErrCode = &v
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
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
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

type ModifyConsumerGroupPasswordRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrMessage(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetSuccess(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumerGroupPasswordResponseBody) SetErrCode(v string) *ModifyConsumerGroupPasswordResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ModifyConsumptionTimestampResponseBody) SetErrMessage(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetSuccess(v string) *ModifyConsumptionTimestampResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyConsumptionTimestampResponseBody) SetErrCode(v string) *ModifyConsumptionTimestampResponseBody {
	s.ErrCode = &v
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

type ModifySubscriptionObjectRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ModifySubscriptionObjectResponseBody) SetErrMessage(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetSuccess(v string) *ModifySubscriptionObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySubscriptionObjectResponseBody) SetErrCode(v string) *ModifySubscriptionObjectResponseBody {
	s.ErrCode = &v
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
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s ModifySynchronizationObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ModifySynchronizationObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySynchronizationObjectResponseBody) SetTaskId(v string) *ModifySynchronizationObjectResponseBody {
	s.TaskId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetRequestId(v string) *ModifySynchronizationObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrMessage(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetSuccess(v string) *ModifySynchronizationObjectResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySynchronizationObjectResponseBody) SetErrCode(v string) *ModifySynchronizationObjectResponseBody {
	s.ErrCode = &v
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

type ResetSynchronizationJobRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ResetSynchronizationJobResponseBody) SetErrMessage(v string) *ResetSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetSuccess(v string) *ResetSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetErrCode(v string) *ResetSynchronizationJobResponseBody {
	s.ErrCode = &v
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
	DtsInstanceId *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	PrecheckItems *string `json:"PrecheckItems,omitempty" xml:"PrecheckItems,omitempty"`
}

func (s ShieldPrecheckRequest) String() string {
	return tea.Prettify(s)
}

func (s ShieldPrecheckRequest) GoString() string {
	return s.String()
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *ShieldPrecheckResponseBody) SetErrMessage(v string) *ShieldPrecheckResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetSuccess(v bool) *ShieldPrecheckResponseBody {
	s.Success = &v
	return s
}

func (s *ShieldPrecheckResponseBody) SetErrCode(v string) *ShieldPrecheckResponseBody {
	s.ErrCode = &v
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

type StartMigrationJobRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *StartMigrationJobResponseBody) SetErrMessage(v string) *StartMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetSuccess(v string) *StartMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetErrCode(v string) *StartMigrationJobResponseBody {
	s.ErrCode = &v
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
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s StartSubscriptionInstanceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartSubscriptionInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartSubscriptionInstanceResponseBody) SetTaskId(v string) *StartSubscriptionInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetRequestId(v string) *StartSubscriptionInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrMessage(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetSuccess(v string) *StartSubscriptionInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *StartSubscriptionInstanceResponseBody) SetErrCode(v string) *StartSubscriptionInstanceResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *StartSynchronizationJobResponseBody) SetErrMessage(v string) *StartSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetSuccess(v string) *StartSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetErrCode(v string) *StartSynchronizationJobResponseBody {
	s.ErrCode = &v
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

type StopMigrationJobRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *StopMigrationJobResponseBody) SetErrMessage(v string) *StopMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetSuccess(v string) *StopMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StopMigrationJobResponseBody) SetErrCode(v string) *StopMigrationJobResponseBody {
	s.ErrCode = &v
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

type SuspendMigrationJobRequest struct {
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *SuspendMigrationJobResponseBody) SetErrMessage(v string) *SuspendMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetSuccess(v string) *SuspendMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetErrCode(v string) *SuspendMigrationJobResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *SuspendSynchronizationJobResponseBody) SetErrMessage(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetSuccess(v string) *SuspendSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetErrCode(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrCode = &v
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
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
}

func (s SwitchSynchronizationEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponseBody) SetTaskId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.TaskId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetRequestId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrMessage(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetSuccess(v string) *SwitchSynchronizationEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrCode(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *TagResourcesResponseBody) SetErrMessage(v string) *TagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *TagResourcesResponseBody) SetSuccess(v bool) *TagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *TagResourcesResponseBody) SetErrCode(v string) *TagResourcesResponseBody {
	s.ErrCode = &v
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
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Success    *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
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

func (s *UntagResourcesResponseBody) SetErrMessage(v string) *UntagResourcesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *UntagResourcesResponseBody) SetSuccess(v bool) *UntagResourcesResponseBody {
	s.Success = &v
	return s
}

func (s *UntagResourcesResponseBody) SetErrCode(v string) *UntagResourcesResponseBody {
	s.ErrCode = &v
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
