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

type ConfigureMigrationJobRequest struct {
	DestinationEndpoint *ConfigureMigrationJobRequestDestinationEndpoint `json:"DestinationEndpoint,omitempty" xml:"DestinationEndpoint,omitempty" type:"Struct"`
	MigrationMode       *ConfigureMigrationJobRequestMigrationMode       `json:"MigrationMode,omitempty" xml:"MigrationMode,omitempty" type:"Struct"`
	SourceEndpoint      *ConfigureMigrationJobRequestSourceEndpoint      `json:"SourceEndpoint,omitempty" xml:"SourceEndpoint,omitempty" type:"Struct"`
	MigrationJobId      *string                                          `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	MigrationJobName    *string                                          `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
	MigrationObject     *string                                          `json:"MigrationObject,omitempty" xml:"MigrationObject,omitempty"`
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

type ConfigureMigrationJobRequestDestinationEndpoint struct {
	DataBaseName *string `json:"DataBaseName,omitempty" xml:"DataBaseName,omitempty"`
	EngineName   *string `json:"EngineName,omitempty" xml:"EngineName,omitempty"`
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
	Password     *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Port         *string `json:"Port,omitempty" xml:"Port,omitempty"`
	Region       *string `json:"Region,omitempty" xml:"Region,omitempty"`
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

func (s *ConfigureMigrationJobRequestSourceEndpoint) SetUserName(v string) *ConfigureMigrationJobRequestSourceEndpoint {
	s.UserName = &v
	return s
}

type ConfigureMigrationJobResponseBody struct {
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
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

type CreateMigrationJobRequest struct {
	ClientToken       *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	MigrationJobClass *string `json:"MigrationJobClass,omitempty" xml:"MigrationJobClass,omitempty"`
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

func (s *CreateMigrationJobRequest) SetRegion(v string) *CreateMigrationJobRequest {
	s.Region = &v
	return s
}

type CreateMigrationJobResponseBody struct {
	ErrCode        *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage     *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type CreateSynchronousJobRequest struct {
	DestinationInstanceId  *string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	FullDataIntialization  *bool   `json:"FullDataIntialization,omitempty" xml:"FullDataIntialization,omitempty"`
	SourceInstanceId       *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	StructureIntialization *bool   `json:"StructureIntialization,omitempty" xml:"StructureIntialization,omitempty"`
	SynchronousJobName     *string `json:"SynchronousJobName,omitempty" xml:"SynchronousJobName,omitempty"`
	SynchronousObjectList  *string `json:"SynchronousObjectList,omitempty" xml:"SynchronousObjectList,omitempty"`
	SynchronousSpeedLimit  *string `json:"SynchronousSpeedLimit,omitempty" xml:"SynchronousSpeedLimit,omitempty"`
}

func (s CreateSynchronousJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronousJobRequest) GoString() string {
	return s.String()
}

func (s *CreateSynchronousJobRequest) SetDestinationInstanceId(v string) *CreateSynchronousJobRequest {
	s.DestinationInstanceId = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetFullDataIntialization(v bool) *CreateSynchronousJobRequest {
	s.FullDataIntialization = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetSourceInstanceId(v string) *CreateSynchronousJobRequest {
	s.SourceInstanceId = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetStructureIntialization(v bool) *CreateSynchronousJobRequest {
	s.StructureIntialization = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetSynchronousJobName(v string) *CreateSynchronousJobRequest {
	s.SynchronousJobName = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetSynchronousObjectList(v string) *CreateSynchronousJobRequest {
	s.SynchronousObjectList = &v
	return s
}

func (s *CreateSynchronousJobRequest) SetSynchronousSpeedLimit(v string) *CreateSynchronousJobRequest {
	s.SynchronousSpeedLimit = &v
	return s
}

type CreateSynchronousJobResponseBody struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s CreateSynchronousJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronousJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynchronousJobResponseBody) SetSynchronousJobId(v string) *CreateSynchronousJobResponseBody {
	s.SynchronousJobId = &v
	return s
}

type CreateSynchronousJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateSynchronousJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSynchronousJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSynchronousJobResponse) GoString() string {
	return s.String()
}

func (s *CreateSynchronousJobResponse) SetHeaders(v map[string]*string) *CreateSynchronousJobResponse {
	s.Headers = v
	return s
}

func (s *CreateSynchronousJobResponse) SetStatusCode(v int32) *CreateSynchronousJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSynchronousJobResponse) SetBody(v *CreateSynchronousJobResponseBody) *CreateSynchronousJobResponse {
	s.Body = v
	return s
}

type DeleteMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type DeleteMigrationJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type DeleteSynchronousJobRequest struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s DeleteSynchronousJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronousJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSynchronousJobRequest) SetSynchronousJobId(v string) *DeleteSynchronousJobRequest {
	s.SynchronousJobId = &v
	return s
}

type DeleteSynchronousJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteSynchronousJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSynchronousJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteSynchronousJobResponse) SetHeaders(v map[string]*string) *DeleteSynchronousJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteSynchronousJobResponse) SetStatusCode(v int32) *DeleteSynchronousJobResponse {
	s.StatusCode = &v
	return s
}

type DescirbeMigrationJobsRequest struct {
	MigrationJobName *string `json:"MigrationJobName,omitempty" xml:"MigrationJobName,omitempty"`
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

type DescribeMigrationJobStatusRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type DescribeSynchronousJobConfigurationRequest struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s DescribeSynchronousJobConfigurationRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationRequest) SetSynchronousJobId(v string) *DescribeSynchronousJobConfigurationRequest {
	s.SynchronousJobId = &v
	return s
}

type DescribeSynchronousJobConfigurationResponseBody struct {
	CreateTime             *string                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DestinationDBType      *string                                                               `json:"DestinationDBType,omitempty" xml:"DestinationDBType,omitempty"`
	DestinationInstanceId  *string                                                               `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	FullDataIntialization  *string                                                               `json:"FullDataIntialization,omitempty" xml:"FullDataIntialization,omitempty"`
	SourceDBType           *string                                                               `json:"SourceDBType,omitempty" xml:"SourceDBType,omitempty"`
	SourceInstanceId       *string                                                               `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	StructureIntialization *string                                                               `json:"StructureIntialization,omitempty" xml:"StructureIntialization,omitempty"`
	SynchronousJobId       *string                                                               `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
	SynchronousJobName     *string                                                               `json:"SynchronousJobName,omitempty" xml:"SynchronousJobName,omitempty"`
	SynchronousObjectList  *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList `json:"SynchronousObjectList,omitempty" xml:"SynchronousObjectList,omitempty" type:"Struct"`
	SynchronousSpeedLimit  *string                                                               `json:"SynchronousSpeedLimit,omitempty" xml:"SynchronousSpeedLimit,omitempty"`
}

func (s DescribeSynchronousJobConfigurationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetCreateTime(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetDestinationDBType(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.DestinationDBType = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetDestinationInstanceId(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.DestinationInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetFullDataIntialization(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.FullDataIntialization = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSourceDBType(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.SourceDBType = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSourceInstanceId(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetStructureIntialization(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.StructureIntialization = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSynchronousJobId(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.SynchronousJobId = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSynchronousJobName(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.SynchronousJobName = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSynchronousObjectList(v *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList) *DescribeSynchronousJobConfigurationResponseBody {
	s.SynchronousObjectList = v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBody) SetSynchronousSpeedLimit(v string) *DescribeSynchronousJobConfigurationResponseBody {
	s.SynchronousSpeedLimit = &v
	return s
}

type DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList struct {
	SynchronousObject []*DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject `json:"SynchronousObject,omitempty" xml:"SynchronousObject,omitempty" type:"Repeated"`
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList) SetSynchronousObject(v []*DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectList {
	s.SynchronousObject = v
	return s
}

type DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject struct {
	DatabaseName  *string                                                                                         `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	TableList     *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList `json:"TableList,omitempty" xml:"TableList,omitempty" type:"Struct"`
	WholeDatabase *string                                                                                         `json:"WholeDatabase,omitempty" xml:"WholeDatabase,omitempty"`
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) SetDatabaseName(v string) *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject {
	s.DatabaseName = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) SetTableList(v *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList) *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject {
	s.TableList = v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject) SetWholeDatabase(v string) *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObject {
	s.WholeDatabase = &v
	return s
}

type DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList struct {
	Table []*string `json:"Table,omitempty" xml:"Table,omitempty" type:"Repeated"`
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList) SetTable(v []*string) *DescribeSynchronousJobConfigurationResponseBodySynchronousObjectListSynchronousObjectTableList {
	s.Table = v
	return s
}

type DescribeSynchronousJobConfigurationResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronousJobConfigurationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronousJobConfigurationResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobConfigurationResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobConfigurationResponse) SetHeaders(v map[string]*string) *DescribeSynchronousJobConfigurationResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponse) SetStatusCode(v int32) *DescribeSynchronousJobConfigurationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronousJobConfigurationResponse) SetBody(v *DescribeSynchronousJobConfigurationResponseBody) *DescribeSynchronousJobConfigurationResponse {
	s.Body = v
	return s
}

type DescribeSynchronousJobDetailsRequest struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s DescribeSynchronousJobDetailsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobDetailsRequest) SetSynchronousJobId(v string) *DescribeSynchronousJobDetailsRequest {
	s.SynchronousJobId = &v
	return s
}

type DescribeSynchronousJobDetailsResponseBody struct {
	CreateTime                       *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	DelaySecond                      *string                                                   `json:"DelaySecond,omitempty" xml:"DelaySecond,omitempty"`
	DestinationInstanceId            *string                                                   `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	ErrorMessage                     *string                                                   `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FullDataIntializatioPercent      *string                                                   `json:"FullDataIntializatioPercent,omitempty" xml:"FullDataIntializatioPercent,omitempty"`
	FullDataIntializatioProgress     *string                                                   `json:"FullDataIntializatioProgress,omitempty" xml:"FullDataIntializatioProgress,omitempty"`
	FullDataIntializationStatus      *string                                                   `json:"FullDataIntializationStatus,omitempty" xml:"FullDataIntializationStatus,omitempty"`
	IncrementDataIntializatioDelay   *string                                                   `json:"IncrementDataIntializatioDelay,omitempty" xml:"IncrementDataIntializatioDelay,omitempty"`
	IncrementDataIntializatioPercent *string                                                   `json:"IncrementDataIntializatioPercent,omitempty" xml:"IncrementDataIntializatioPercent,omitempty"`
	IncrementDataIntializationStatus *string                                                   `json:"IncrementDataIntializationStatus,omitempty" xml:"IncrementDataIntializationStatus,omitempty"`
	PrecheckDetails                  *DescribeSynchronousJobDetailsResponseBodyPrecheckDetails `json:"PrecheckDetails,omitempty" xml:"PrecheckDetails,omitempty" type:"Struct"`
	PrecheckPercent                  *string                                                   `json:"PrecheckPercent,omitempty" xml:"PrecheckPercent,omitempty"`
	PrecheckStatus                   *string                                                   `json:"PrecheckStatus,omitempty" xml:"PrecheckStatus,omitempty"`
	SourceInstanceId                 *string                                                   `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	StructureIntializationPercent    *string                                                   `json:"StructureIntializationPercent,omitempty" xml:"StructureIntializationPercent,omitempty"`
	StructureIntializationProgress   *string                                                   `json:"StructureIntializationProgress,omitempty" xml:"StructureIntializationProgress,omitempty"`
	StructureIntializationStatus     *string                                                   `json:"StructureIntializationStatus,omitempty" xml:"StructureIntializationStatus,omitempty"`
	SynchronousFlow                  *string                                                   `json:"SynchronousFlow,omitempty" xml:"SynchronousFlow,omitempty"`
	SynchronousJobId                 *string                                                   `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
	SynchronousJobName               *string                                                   `json:"SynchronousJobName,omitempty" xml:"SynchronousJobName,omitempty"`
	SynchronousStatus                *string                                                   `json:"SynchronousStatus,omitempty" xml:"SynchronousStatus,omitempty"`
	SynchronousTPS                   *string                                                   `json:"SynchronousTPS,omitempty" xml:"SynchronousTPS,omitempty"`
}

func (s DescribeSynchronousJobDetailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetCreateTime(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetDelaySecond(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.DelaySecond = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetDestinationInstanceId(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.DestinationInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetErrorMessage(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetFullDataIntializatioPercent(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.FullDataIntializatioPercent = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetFullDataIntializatioProgress(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.FullDataIntializatioProgress = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetFullDataIntializationStatus(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.FullDataIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetIncrementDataIntializatioDelay(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.IncrementDataIntializatioDelay = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetIncrementDataIntializatioPercent(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.IncrementDataIntializatioPercent = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetIncrementDataIntializationStatus(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.IncrementDataIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetPrecheckDetails(v *DescribeSynchronousJobDetailsResponseBodyPrecheckDetails) *DescribeSynchronousJobDetailsResponseBody {
	s.PrecheckDetails = v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetPrecheckPercent(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.PrecheckPercent = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetPrecheckStatus(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.PrecheckStatus = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSourceInstanceId(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetStructureIntializationPercent(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.StructureIntializationPercent = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetStructureIntializationProgress(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.StructureIntializationProgress = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetStructureIntializationStatus(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.StructureIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSynchronousFlow(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SynchronousFlow = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSynchronousJobId(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SynchronousJobId = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSynchronousJobName(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SynchronousJobName = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSynchronousStatus(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SynchronousStatus = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBody) SetSynchronousTPS(v string) *DescribeSynchronousJobDetailsResponseBody {
	s.SynchronousTPS = &v
	return s
}

type DescribeSynchronousJobDetailsResponseBodyPrecheckDetails struct {
	PrecheckDetail []*DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail `json:"PrecheckDetail,omitempty" xml:"PrecheckDetail,omitempty" type:"Repeated"`
}

func (s DescribeSynchronousJobDetailsResponseBodyPrecheckDetails) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobDetailsResponseBodyPrecheckDetails) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetails) SetPrecheckDetail(v []*DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetails {
	s.PrecheckDetail = v
	return s
}

type DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail struct {
	CheckItem            *string `json:"CheckItem,omitempty" xml:"CheckItem,omitempty"`
	CheckItemDescription *string `json:"CheckItemDescription,omitempty" xml:"CheckItemDescription,omitempty"`
	CheckResult          *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	FailedReson          *string `json:"FailedReson,omitempty" xml:"FailedReson,omitempty"`
	RepairMethod         *string `json:"RepairMethod,omitempty" xml:"RepairMethod,omitempty"`
}

func (s DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) SetCheckItem(v string) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail {
	s.CheckItem = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) SetCheckItemDescription(v string) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail {
	s.CheckItemDescription = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) SetCheckResult(v string) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail {
	s.CheckResult = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) SetFailedReson(v string) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail {
	s.FailedReson = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail) SetRepairMethod(v string) *DescribeSynchronousJobDetailsResponseBodyPrecheckDetailsPrecheckDetail {
	s.RepairMethod = &v
	return s
}

type DescribeSynchronousJobDetailsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronousJobDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronousJobDetailsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobDetailsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobDetailsResponse) SetHeaders(v map[string]*string) *DescribeSynchronousJobDetailsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronousJobDetailsResponse) SetStatusCode(v int32) *DescribeSynchronousJobDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronousJobDetailsResponse) SetBody(v *DescribeSynchronousJobDetailsResponseBody) *DescribeSynchronousJobDetailsResponse {
	s.Body = v
	return s
}

type DescribeSynchronousJobListRequest struct {
	InstanceId         *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PageNum            *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize           *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SynchronousJobName *string `json:"SynchronousJobName,omitempty" xml:"SynchronousJobName,omitempty"`
}

func (s DescribeSynchronousJobListRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobListRequest) SetInstanceId(v string) *DescribeSynchronousJobListRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSynchronousJobListRequest) SetPageNum(v int32) *DescribeSynchronousJobListRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeSynchronousJobListRequest) SetPageSize(v int32) *DescribeSynchronousJobListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSynchronousJobListRequest) SetSynchronousJobName(v string) *DescribeSynchronousJobListRequest {
	s.SynchronousJobName = &v
	return s
}

type DescribeSynchronousJobListResponseBody struct {
	PageNumber         *int32                                                    `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageRecordCount    *int32                                                    `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	SynchronousJobList *DescribeSynchronousJobListResponseBodySynchronousJobList `json:"SynchronousJobList,omitempty" xml:"SynchronousJobList,omitempty" type:"Struct"`
	TotalRecordCount   *int64                                                    `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeSynchronousJobListResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobListResponseBody) SetPageNumber(v int32) *DescribeSynchronousJobListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBody) SetPageRecordCount(v int32) *DescribeSynchronousJobListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBody) SetSynchronousJobList(v *DescribeSynchronousJobListResponseBodySynchronousJobList) *DescribeSynchronousJobListResponseBody {
	s.SynchronousJobList = v
	return s
}

func (s *DescribeSynchronousJobListResponseBody) SetTotalRecordCount(v int64) *DescribeSynchronousJobListResponseBody {
	s.TotalRecordCount = &v
	return s
}

type DescribeSynchronousJobListResponseBodySynchronousJobList struct {
	SynchronousJob []*DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob `json:"SynchronousJob,omitempty" xml:"SynchronousJob,omitempty" type:"Repeated"`
}

func (s DescribeSynchronousJobListResponseBodySynchronousJobList) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobListResponseBodySynchronousJobList) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobList) SetSynchronousJob(v []*DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) *DescribeSynchronousJobListResponseBodySynchronousJobList {
	s.SynchronousJob = v
	return s
}

type DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob struct {
	DelaySecond                      *string `json:"DelaySecond,omitempty" xml:"DelaySecond,omitempty"`
	DestinationInstanceId            *string `json:"DestinationInstanceId,omitempty" xml:"DestinationInstanceId,omitempty"`
	ErrorMessage                     *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	FullDataIntializatioPercent      *string `json:"FullDataIntializatioPercent,omitempty" xml:"FullDataIntializatioPercent,omitempty"`
	FullDataIntializatioProgress     *string `json:"FullDataIntializatioProgress,omitempty" xml:"FullDataIntializatioProgress,omitempty"`
	FullDataIntializationStatus      *string `json:"FullDataIntializationStatus,omitempty" xml:"FullDataIntializationStatus,omitempty"`
	IncrementDataIntializatioDelay   *string `json:"IncrementDataIntializatioDelay,omitempty" xml:"IncrementDataIntializatioDelay,omitempty"`
	IncrementDataIntializatioPercent *string `json:"IncrementDataIntializatioPercent,omitempty" xml:"IncrementDataIntializatioPercent,omitempty"`
	IncrementDataIntializationStatus *string `json:"IncrementDataIntializationStatus,omitempty" xml:"IncrementDataIntializationStatus,omitempty"`
	SourceInstanceId                 *string `json:"SourceInstanceId,omitempty" xml:"SourceInstanceId,omitempty"`
	StructureIntializationPercent    *string `json:"StructureIntializationPercent,omitempty" xml:"StructureIntializationPercent,omitempty"`
	StructureIntializationProgress   *string `json:"StructureIntializationProgress,omitempty" xml:"StructureIntializationProgress,omitempty"`
	StructureIntializationStatus     *string `json:"StructureIntializationStatus,omitempty" xml:"StructureIntializationStatus,omitempty"`
	SynchronousFlow                  *string `json:"SynchronousFlow,omitempty" xml:"SynchronousFlow,omitempty"`
	SynchronousJobId                 *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
	SynchronousJobName               *string `json:"SynchronousJobName,omitempty" xml:"SynchronousJobName,omitempty"`
	SynchronousStatus                *string `json:"SynchronousStatus,omitempty" xml:"SynchronousStatus,omitempty"`
	SynchronousTPS                   *string `json:"SynchronousTPS,omitempty" xml:"SynchronousTPS,omitempty"`
}

func (s DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetDelaySecond(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.DelaySecond = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetDestinationInstanceId(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.DestinationInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetErrorMessage(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetFullDataIntializatioPercent(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.FullDataIntializatioPercent = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetFullDataIntializatioProgress(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.FullDataIntializatioProgress = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetFullDataIntializationStatus(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.FullDataIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetIncrementDataIntializatioDelay(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.IncrementDataIntializatioDelay = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetIncrementDataIntializatioPercent(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.IncrementDataIntializatioPercent = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetIncrementDataIntializationStatus(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.IncrementDataIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSourceInstanceId(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SourceInstanceId = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetStructureIntializationPercent(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.StructureIntializationPercent = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetStructureIntializationProgress(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.StructureIntializationProgress = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetStructureIntializationStatus(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.StructureIntializationStatus = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSynchronousFlow(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SynchronousFlow = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSynchronousJobId(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SynchronousJobId = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSynchronousJobName(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SynchronousJobName = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSynchronousStatus(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SynchronousStatus = &v
	return s
}

func (s *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob) SetSynchronousTPS(v string) *DescribeSynchronousJobListResponseBodySynchronousJobListSynchronousJob {
	s.SynchronousTPS = &v
	return s
}

type DescribeSynchronousJobListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeSynchronousJobListResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeSynchronousJobListResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSynchronousJobListResponse) GoString() string {
	return s.String()
}

func (s *DescribeSynchronousJobListResponse) SetHeaders(v map[string]*string) *DescribeSynchronousJobListResponse {
	s.Headers = v
	return s
}

func (s *DescribeSynchronousJobListResponse) SetStatusCode(v int32) *DescribeSynchronousJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSynchronousJobListResponse) SetBody(v *DescribeSynchronousJobListResponseBody) *DescribeSynchronousJobListResponse {
	s.Body = v
	return s
}

type StartMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type StartMigrationJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type StartSynchronousJobRequest struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s StartSynchronousJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronousJobRequest) GoString() string {
	return s.String()
}

func (s *StartSynchronousJobRequest) SetSynchronousJobId(v string) *StartSynchronousJobRequest {
	s.SynchronousJobId = &v
	return s
}

type StartSynchronousJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s StartSynchronousJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartSynchronousJobResponse) GoString() string {
	return s.String()
}

func (s *StartSynchronousJobResponse) SetHeaders(v map[string]*string) *StartSynchronousJobResponse {
	s.Headers = v
	return s
}

func (s *StartSynchronousJobResponse) SetStatusCode(v int32) *StartSynchronousJobResponse {
	s.StatusCode = &v
	return s
}

type StopMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type StopMigrationJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type SuspendMigrationJobRequest struct {
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
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

type SuspendMigrationJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type SuspendSynchronousJobRequest struct {
	SynchronousJobId *string `json:"SynchronousJobId,omitempty" xml:"SynchronousJobId,omitempty"`
}

func (s SuspendSynchronousJobRequest) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronousJobRequest) GoString() string {
	return s.String()
}

func (s *SuspendSynchronousJobRequest) SetSynchronousJobId(v string) *SuspendSynchronousJobRequest {
	s.SynchronousJobId = &v
	return s
}

type SuspendSynchronousJobResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s SuspendSynchronousJobResponse) String() string {
	return tea.Prettify(s)
}

func (s SuspendSynchronousJobResponse) GoString() string {
	return s.String()
}

func (s *SuspendSynchronousJobResponse) SetHeaders(v map[string]*string) *SuspendSynchronousJobResponse {
	s.Headers = v
	return s
}

func (s *SuspendSynchronousJobResponse) SetStatusCode(v int32) *SuspendSynchronousJobResponse {
	s.StatusCode = &v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("dtsrs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationJobName)) {
		query["MigrationJobName"] = request.MigrationJobName
	}

	if !tea.BoolValue(util.IsUnset(request.MigrationObject)) {
		query["MigrationObject"] = request.MigrationObject
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.DestinationEndpoint))) {
		query["DestinationEndpoint"] = request.DestinationEndpoint
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.MigrationMode))) {
		query["MigrationMode"] = request.MigrationMode
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.SourceEndpoint))) {
		query["SourceEndpoint"] = request.SourceEndpoint
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ConfigureMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
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

	if !tea.BoolValue(util.IsUnset(request.Region)) {
		query["Region"] = request.Region
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
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

func (client *Client) CreateSynchronousJobWithOptions(request *CreateSynchronousJobRequest, runtime *util.RuntimeOptions) (_result *CreateSynchronousJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationInstanceId)) {
		query["DestinationInstanceId"] = request.DestinationInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.FullDataIntialization)) {
		query["FullDataIntialization"] = request.FullDataIntialization
	}

	if !tea.BoolValue(util.IsUnset(request.SourceInstanceId)) {
		query["SourceInstanceId"] = request.SourceInstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.StructureIntialization)) {
		query["StructureIntialization"] = request.StructureIntialization
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronousJobName)) {
		query["SynchronousJobName"] = request.SynchronousJobName
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronousObjectList)) {
		query["SynchronousObjectList"] = request.SynchronousObjectList
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronousSpeedLimit)) {
		query["SynchronousSpeedLimit"] = request.SynchronousSpeedLimit
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSynchronousJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSynchronousJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSynchronousJob(request *CreateSynchronousJobRequest) (_result *CreateSynchronousJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSynchronousJobResponse{}
	_body, _err := client.CreateSynchronousJobWithOptions(request, runtime)
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
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

func (client *Client) DeleteSynchronousJobWithOptions(request *DeleteSynchronousJobRequest, runtime *util.RuntimeOptions) (_result *DeleteSynchronousJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynchronousJobId)) {
		query["SynchronousJobId"] = request.SynchronousJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteSynchronousJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteSynchronousJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSynchronousJob(request *DeleteSynchronousJobRequest) (_result *DeleteSynchronousJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteSynchronousJobResponse{}
	_body, _err := client.DeleteSynchronousJobWithOptions(request, runtime)
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
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
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

func (client *Client) DescribeMigrationJobStatusWithOptions(request *DescribeMigrationJobStatusRequest, runtime *util.RuntimeOptions) (_result *DescribeMigrationJobStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeMigrationJobStatus"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
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

func (client *Client) DescribeSynchronousJobConfigurationWithOptions(request *DescribeSynchronousJobConfigurationRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronousJobConfigurationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynchronousJobId)) {
		query["SynchronousJobId"] = request.SynchronousJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronousJobConfiguration"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronousJobConfigurationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronousJobConfiguration(request *DescribeSynchronousJobConfigurationRequest) (_result *DescribeSynchronousJobConfigurationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronousJobConfigurationResponse{}
	_body, _err := client.DescribeSynchronousJobConfigurationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronousJobDetailsWithOptions(request *DescribeSynchronousJobDetailsRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronousJobDetailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynchronousJobId)) {
		query["SynchronousJobId"] = request.SynchronousJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronousJobDetails"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronousJobDetailsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronousJobDetails(request *DescribeSynchronousJobDetailsRequest) (_result *DescribeSynchronousJobDetailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronousJobDetailsResponse{}
	_body, _err := client.DescribeSynchronousJobDetailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeSynchronousJobListWithOptions(request *DescribeSynchronousJobListRequest, runtime *util.RuntimeOptions) (_result *DescribeSynchronousJobListResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SynchronousJobName)) {
		query["SynchronousJobName"] = request.SynchronousJobName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSynchronousJobList"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSynchronousJobListResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeSynchronousJobList(request *DescribeSynchronousJobListRequest) (_result *DescribeSynchronousJobListResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSynchronousJobListResponse{}
	_body, _err := client.DescribeSynchronousJobListWithOptions(request, runtime)
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

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
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

func (client *Client) StartSynchronousJobWithOptions(request *StartSynchronousJobRequest, runtime *util.RuntimeOptions) (_result *StartSynchronousJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynchronousJobId)) {
		query["SynchronousJobId"] = request.SynchronousJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartSynchronousJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &StartSynchronousJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartSynchronousJob(request *StartSynchronousJobRequest) (_result *StartSynchronousJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartSynchronousJobResponse{}
	_body, _err := client.StartSynchronousJobWithOptions(request, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
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
	if !tea.BoolValue(util.IsUnset(request.MigrationJobId)) {
		query["MigrationJobId"] = request.MigrationJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendMigrationJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
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

func (client *Client) SuspendSynchronousJobWithOptions(request *SuspendSynchronousJobRequest, runtime *util.RuntimeOptions) (_result *SuspendSynchronousJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SynchronousJobId)) {
		query["SynchronousJobId"] = request.SynchronousJobId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SuspendSynchronousJob"),
		Version:     tea.String("2015-11-24"),
		Protocol:    tea.String("HTTP"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("none"),
	}
	_result = &SuspendSynchronousJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SuspendSynchronousJob(request *SuspendSynchronousJobRequest) (_result *SuspendSynchronousJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SuspendSynchronousJobResponse{}
	_body, _err := client.SuspendSynchronousJobWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
