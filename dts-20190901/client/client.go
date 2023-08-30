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

type ConfigureSynchronizationJobRequest struct {
	DestinationEndpoint      *ConfigureSynchronizationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	PartitionKey             *ConfigureSynchronizationJobRequestPartitionKey        `json:"PartitionKey,omitempty" xml:"PartitionKey,omitempty" type:"Struct"`
	SourceEndpoint           *ConfigureSynchronizationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	Checkpoint               *string                                                `json:"Checkpoint,omitempty" xml:"Checkpoint,omitempty"`
	DataInitialization       *bool                                                  `json:"DataInitialization,omitempty" xml:"DataInitialization,omitempty"`
	MigrationReserved        *string                                                `json:"MigrationReserved,omitempty" xml:"MigrationReserved,omitempty"`
	OwnerId                  *string                                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	StructureInitialization  *bool                                                  `json:"StructureInitialization,omitempty" xml:"StructureInitialization,omitempty"`
	SynchronizationDirection *string                                                `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string                                                `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName   *string                                                `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
	SynchronizationObjects   *string                                                `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
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

func (s *ConfigureSynchronizationJobRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobRequest {
	s.SynchronizationDirection = &v
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

type ConfigureSynchronizationJobAlertRequest struct {
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s ConfigureSynchronizationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetDelayOverSeconds(v string) *ConfigureSynchronizationJobAlertRequest {
	s.DelayOverSeconds = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertPhone(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertPhone = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetErrorAlertStatus(v string) *ConfigureSynchronizationJobAlertRequest {
	s.ErrorAlertStatus = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetOwnerId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *ConfigureSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

type ConfigureSynchronizationJobAlertResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureSynchronizationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ConfigureSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrCode(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetErrMessage(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetRequestId(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponseBody) SetSuccess(v string) *ConfigureSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

type ConfigureSynchronizationJobAlertResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ConfigureSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ConfigureSynchronizationJobAlertResponse) SetStatusCode(v int32) *ConfigureSynchronizationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureSynchronizationJobAlertResponse) SetBody(v *ConfigureSynchronizationJobAlertResponseBody) *ConfigureSynchronizationJobAlertResponse {
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
	Topology                *string                                             `json:"Topology,omitempty" xml:"Topology,omitempty"`
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

func (s *CreateSynchronizationJobRequest) SetTopology(v string) *CreateSynchronizationJobRequest {
	s.Topology = &v
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

type DescribeEndpointSwitchStatusRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId     *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	TaskId      *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeEndpointSwitchStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointSwitchStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusRequest) SetClientToken(v string) *DescribeEndpointSwitchStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetOwnerId(v string) *DescribeEndpointSwitchStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusRequest) SetTaskId(v string) *DescribeEndpointSwitchStatusRequest {
	s.TaskId = &v
	return s
}

type DescribeEndpointSwitchStatusResponseBody struct {
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeEndpointSwitchStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeEndpointSwitchStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetErrorMessage(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetRequestId(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponseBody) SetStatus(v string) *DescribeEndpointSwitchStatusResponseBody {
	s.Status = &v
	return s
}

type DescribeEndpointSwitchStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeEndpointSwitchStatusResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeEndpointSwitchStatusResponse) SetStatusCode(v int32) *DescribeEndpointSwitchStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndpointSwitchStatusResponse) SetBody(v *DescribeEndpointSwitchStatusResponseBody) *DescribeEndpointSwitchStatusResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobAlertRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s DescribeSynchronizationJobAlertRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertRequest) SetClientToken(v string) *DescribeSynchronizationJobAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetOwnerId(v string) *DescribeSynchronizationJobAlertRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *DescribeSynchronizationJobAlertRequest) SetSynchronizationJobId(v string) *DescribeSynchronizationJobAlertRequest {
	s.SynchronizationJobId = &v
	return s
}

type DescribeSynchronizationJobAlertResponseBody struct {
	DelayAlertPhone          *string `json:"DelayAlertPhone,omitempty" xml:"DelayAlertPhone,omitempty"`
	DelayAlertStatus         *string `json:"DelayAlertStatus,omitempty" xml:"DelayAlertStatus,omitempty"`
	DelayOverSeconds         *string `json:"DelayOverSeconds,omitempty" xml:"DelayOverSeconds,omitempty"`
	ErrCode                  *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage               *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	ErrorAlertPhone          *string `json:"ErrorAlertPhone,omitempty" xml:"ErrorAlertPhone,omitempty"`
	ErrorAlertStatus         *string `json:"ErrorAlertStatus,omitempty" xml:"ErrorAlertStatus,omitempty"`
	RequestId                *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                  *string `json:"Success,omitempty" xml:"Success,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationJobName   *string `json:"SynchronizationJobName,omitempty" xml:"SynchronizationJobName,omitempty"`
}

func (s DescribeSynchronizationJobAlertResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobAlertResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetDelayOverSeconds(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.DelayOverSeconds = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrCode(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrMessage(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertPhone(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertPhone = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetErrorAlertStatus(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.ErrorAlertStatus = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetRequestId(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSuccess(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobAlertResponseBody {
	s.SynchronizationDirection = &v
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

type DescribeSynchronizationJobAlertResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronizationJobAlertResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DescribeSynchronizationJobAlertResponse) SetStatusCode(v int32) *DescribeSynchronizationJobAlertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronizationJobAlertResponse) SetBody(v *DescribeSynchronizationJobAlertResponseBody) *DescribeSynchronizationJobAlertResponse {
	s.Body = v
	return s
}

type DescribeSynchronizationJobStatusRequest struct {
	ClientToken              *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
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

func (s *DescribeSynchronizationJobStatusRequest) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusRequest {
	s.SynchronizationDirection = &v
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
	SynchronizationDirection      *string                                                                    `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
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

func (s *DescribeSynchronizationJobStatusResponseBody) SetSynchronizationDirection(v string) *DescribeSynchronizationJobStatusResponseBody {
	s.SynchronizationDirection = &v
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
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodyDestinationEndpoint {
	s.InstanceId = &v
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
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobStatusResponseBodySourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetIP(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobStatusResponseBodySourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobStatusResponseBodySourceEndpoint {
	s.InstanceId = &v
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
	SynchronizationDirection      *string                                                                                       `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
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

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstances) SetSynchronizationDirection(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstances {
	s.SynchronizationDirection = &v
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
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesDestinationEndpoint {
	s.InstanceId = &v
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
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	UserName     *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) GoString() string {
	return s.String()
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetEngineName(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.EngineName = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetIP(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.IP = &v
	return s
}

func (s *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint) SetInstanceId(v string) *DescribeSynchronizationJobsResponseBodySynchronizationInstancesSourceEndpoint {
	s.InstanceId = &v
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

type ModifySynchronizationObjectRequest struct {
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
	SynchronizationObjects   *string `json:"SynchronizationObjects,omitempty" xml:"SynchronizationObjects,omitempty"`
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

func (s *ModifySynchronizationObjectRequest) SetSynchronizationDirection(v string) *ModifySynchronizationObjectRequest {
	s.SynchronizationDirection = &v
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

type ResetSynchronizationJobRequest struct {
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s ResetSynchronizationJobRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetSynchronizationJobRequest) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobRequest) SetOwnerId(v string) *ResetSynchronizationJobRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationDirection(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *ResetSynchronizationJobRequest) SetSynchronizationJobId(v string) *ResetSynchronizationJobRequest {
	s.SynchronizationJobId = &v
	return s
}

type ResetSynchronizationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetSynchronizationJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobResponseBody) SetErrCode(v string) *ResetSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetErrMessage(v string) *ResetSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetRequestId(v string) *ResetSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetSuccess(v string) *ResetSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

type ResetSynchronizationJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetSynchronizationJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ResetSynchronizationJobResponse) SetStatusCode(v int32) *ResetSynchronizationJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetSynchronizationJobResponse) SetBody(v *ResetSynchronizationJobResponseBody) *ResetSynchronizationJobResponse {
	s.Body = v
	return s
}

type StartSynchronizationJobRequest struct {
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
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

func (s *StartSynchronizationJobRequest) SetSynchronizationDirection(v string) *StartSynchronizationJobRequest {
	s.SynchronizationDirection = &v
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

type SuspendSynchronizationJobRequest struct {
	OwnerId                  *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
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

func (s *SuspendSynchronizationJobRequest) SetSynchronizationDirection(v string) *SuspendSynchronizationJobRequest {
	s.SynchronizationDirection = &v
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

type SwitchSynchronizationEndpointRequest struct {
	Endpoint                 *SwitchSynchronizationEndpointRequestEndpoint `json:"Endpoint,omitempty" xml:"Endpoint,omitempty" type:"Struct"`
	OwnerId                  *string                                       `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SynchronizationDirection *string                                       `json:"SynchronizationDirection,omitempty" xml:"SynchronizationDirection,omitempty"`
	SynchronizationJobId     *string                                       `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
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

func (s *SwitchSynchronizationEndpointRequest) SetOwnerId(v string) *SwitchSynchronizationEndpointRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationDirection(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationDirection = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequest) SetSynchronizationJobId(v string) *SwitchSynchronizationEndpointRequest {
	s.SynchronizationJobId = &v
	return s
}

type SwitchSynchronizationEndpointRequestEndpoint struct {
	IP           *string `json:"IP,omitempty" xml:"IP,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SwitchSynchronizationEndpointRequestEndpoint) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointRequestEndpoint) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetIP(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.IP = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceId(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceId = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetInstanceType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.InstanceType = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetPort(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Port = &v
	return s
}

func (s *SwitchSynchronizationEndpointRequestEndpoint) SetType(v string) *SwitchSynchronizationEndpointRequestEndpoint {
	s.Type = &v
	return s
}

type SwitchSynchronizationEndpointResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SwitchSynchronizationEndpointResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SwitchSynchronizationEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrCode(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetErrMessage(v string) *SwitchSynchronizationEndpointResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetRequestId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetSuccess(v string) *SwitchSynchronizationEndpointResponseBody {
	s.Success = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponseBody) SetTaskId(v string) *SwitchSynchronizationEndpointResponseBody {
	s.TaskId = &v
	return s
}

type SwitchSynchronizationEndpointResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SwitchSynchronizationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *SwitchSynchronizationEndpointResponse) SetStatusCode(v int32) *SwitchSynchronizationEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *SwitchSynchronizationEndpointResponse) SetBody(v *SwitchSynchronizationEndpointResponseBody) *SwitchSynchronizationEndpointResponse {
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

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
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
		Version:     tea.String("2019-09-01"),
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

func (client *Client) ConfigureSynchronizationJobAlertWithOptions(request *ConfigureSynchronizationJobAlertRequest, runtime *util.RuntimeOptions) (_result *ConfigureSynchronizationJobAlertResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DelayAlertPhone)) {
		query["DelayAlertPhone"] = request.DelayAlertPhone
	}

	if !tea.BoolValue(util.IsUnset(request.DelayAlertStatus)) {
		query["DelayAlertStatus"] = request.DelayAlertStatus
	}

	if !tea.BoolValue(util.IsUnset(request.DelayOverSeconds)) {
		query["DelayOverSeconds"] = request.DelayOverSeconds
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorAlertPhone)) {
		query["ErrorAlertPhone"] = request.ErrorAlertPhone
	}

	if !tea.BoolValue(util.IsUnset(request.ErrorAlertStatus)) {
		query["ErrorAlertStatus"] = request.ErrorAlertStatus
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigureSynchronizationJobAlert"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ConfigureSynchronizationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Topology)) {
		query["Topology"] = request.Topology
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
		Version:     tea.String("2019-09-01"),
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
		Version:     tea.String("2019-09-01"),
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

func (client *Client) DescribeEndpointSwitchStatusWithOptions(request *DescribeEndpointSwitchStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeEndpointSwitchStatusResponse, _err error) {
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
		Action:      tea.String("DescribeEndpointSwitchStatus"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeEndpointSwitchStatusResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) DescribeSynchronizationJobAlertWithOptions(request *DescribeSynchronizationJobAlertRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronizationJobAlertResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronizationJobAlert"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronizationJobAlertResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronizationJobStatus"),
		Version:     tea.String("2019-09-01"),
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
		Version:     tea.String("2019-09-01"),
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
		Version:     tea.String("2019-09-01"),
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

func (client *Client) ModifySynchronizationObjectWithOptions(request *ModifySynchronizationObjectRequest, runtime *util.RuntimeOptions) (_result *ModifySynchronizationObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
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
		Version:     tea.String("2019-09-01"),
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

func (client *Client) ResetSynchronizationJobWithOptions(request *ResetSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *ResetSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetSynchronizationJob"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetSynchronizationJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) StartSynchronizationJobWithOptions(request *StartSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *StartSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSynchronizationJob"),
		Version:     tea.String("2019-09-01"),
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

func (client *Client) SuspendSynchronizationJobWithOptions(request *SuspendSynchronizationJobRequest, runtime *util.RuntimeOptions) (_result *SuspendSynchronizationJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendSynchronizationJob"),
		Version:     tea.String("2019-09-01"),
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

func (client *Client) SwitchSynchronizationEndpointWithOptions(request *SwitchSynchronizationEndpointRequest, runtime *util.RuntimeOptions) (_result *SwitchSynchronizationEndpointResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationDirection)) {
		query["SynchronizationDirection"] = request.SynchronizationDirection
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronizationJobId)) {
		query["SynchronizationJobId"] = request.SynchronizationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		query["Endpoint"] = request.Endpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SwitchSynchronizationEndpoint"),
		Version:     tea.String("2019-09-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SwitchSynchronizationEndpointResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
