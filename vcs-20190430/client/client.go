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

type CreateComputeTaskRequest struct {
	AlgorithmCodeList *string `json:"AlgorithmCodeList,omitempty" xml:"AlgorithmCodeList,omitempty"`
	DeviceCodeList    *string `json:"DeviceCodeList,omitempty" xml:"DeviceCodeList,omitempty"`
	ProjectId         *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TaskName          *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	VcsId             *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s CreateComputeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskRequest) SetAlgorithmCodeList(v string) *CreateComputeTaskRequest {
	s.AlgorithmCodeList = &v
	return s
}

func (s *CreateComputeTaskRequest) SetDeviceCodeList(v string) *CreateComputeTaskRequest {
	s.DeviceCodeList = &v
	return s
}

func (s *CreateComputeTaskRequest) SetProjectId(v string) *CreateComputeTaskRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateComputeTaskRequest) SetRegionId(v string) *CreateComputeTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CreateComputeTaskRequest) SetTaskName(v string) *CreateComputeTaskRequest {
	s.TaskName = &v
	return s
}

func (s *CreateComputeTaskRequest) SetVcsId(v string) *CreateComputeTaskRequest {
	s.VcsId = &v
	return s
}

type CreateComputeTaskResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateComputeTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateComputeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskResponseBody) SetCode(v string) *CreateComputeTaskResponseBody {
	s.Code = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetData(v *CreateComputeTaskResponseBodyData) *CreateComputeTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateComputeTaskResponseBody) SetMessage(v string) *CreateComputeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateComputeTaskResponseBody) SetRequestId(v string) *CreateComputeTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateComputeTaskResponseBodyData struct {
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateComputeTaskResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskResponseBodyData) SetTaskId(v string) *CreateComputeTaskResponseBodyData {
	s.TaskId = &v
	return s
}

type CreateComputeTaskResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateComputeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateComputeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateComputeTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateComputeTaskResponse) SetHeaders(v map[string]*string) *CreateComputeTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateComputeTaskResponse) SetBody(v *CreateComputeTaskResponseBody) *CreateComputeTaskResponse {
	s.Body = v
	return s
}

type CreateProjectRequest struct {
	AreaCode     *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	ProjectName  *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId     *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TimeZoneCode *string `json:"TimeZoneCode,omitempty" xml:"TimeZoneCode,omitempty"`
	VcsId        *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s CreateProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectRequest) SetAreaCode(v string) *CreateProjectRequest {
	s.AreaCode = &v
	return s
}

func (s *CreateProjectRequest) SetProjectName(v string) *CreateProjectRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateProjectRequest) SetRegionId(v string) *CreateProjectRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProjectRequest) SetTimeZoneCode(v string) *CreateProjectRequest {
	s.TimeZoneCode = &v
	return s
}

func (s *CreateProjectRequest) SetVcsId(v string) *CreateProjectRequest {
	s.VcsId = &v
	return s
}

type CreateProjectResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *CreateProjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) SetCode(v string) *CreateProjectResponseBody {
	s.Code = &v
	return s
}

func (s *CreateProjectResponseBody) SetData(v *CreateProjectResponseBodyData) *CreateProjectResponseBody {
	s.Data = v
	return s
}

func (s *CreateProjectResponseBody) SetMessage(v string) *CreateProjectResponseBody {
	s.Message = &v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

type CreateProjectResponseBodyData struct {
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateProjectResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyData) SetProjectId(v string) *CreateProjectResponseBodyData {
	s.ProjectId = &v
	return s
}

type CreateProjectResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateProjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateProjectResponse) GoString() string {
	return s.String()
}

func (s *CreateProjectResponse) SetHeaders(v map[string]*string) *CreateProjectResponse {
	s.Headers = v
	return s
}

func (s *CreateProjectResponse) SetBody(v *CreateProjectResponseBody) *CreateProjectResponse {
	s.Body = v
	return s
}

type DescribeComputeTasksRequest struct {
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	VcsId     *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s DescribeComputeTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeTasksRequest) GoString() string {
	return s.String()
}

func (s *DescribeComputeTasksRequest) SetPageNum(v int32) *DescribeComputeTasksRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeComputeTasksRequest) SetPageSize(v int32) *DescribeComputeTasksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeComputeTasksRequest) SetProjectId(v string) *DescribeComputeTasksRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeComputeTasksRequest) SetRegionId(v string) *DescribeComputeTasksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeComputeTasksRequest) SetSearchKey(v string) *DescribeComputeTasksRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeComputeTasksRequest) SetVcsId(v string) *DescribeComputeTasksRequest {
	s.VcsId = &v
	return s
}

type DescribeComputeTasksResponseBody struct {
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeComputeTasksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeComputeTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeTasksResponseBody) SetCode(v string) *DescribeComputeTasksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeComputeTasksResponseBody) SetData(v []*DescribeComputeTasksResponseBodyData) *DescribeComputeTasksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeComputeTasksResponseBody) SetMessage(v string) *DescribeComputeTasksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeComputeTasksResponseBody) SetRequestId(v string) *DescribeComputeTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComputeTasksResponseBody) SetTotalCount(v int32) *DescribeComputeTasksResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeComputeTasksResponseBodyData struct {
	AlgorithmName  *string  `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	DeviceNum      *int32   `json:"DeviceNum,omitempty" xml:"DeviceNum,omitempty"`
	ImageSize      *float32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	Runtime        *string  `json:"Runtime,omitempty" xml:"Runtime,omitempty"`
	StructuredSize *float32 `json:"StructuredSize,omitempty" xml:"StructuredSize,omitempty"`
	TaskId         *string  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName       *string  `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	TaskStatus     *string  `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	VectorSize     *float32 `json:"VectorSize,omitempty" xml:"VectorSize,omitempty"`
}

func (s DescribeComputeTasksResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeTasksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeComputeTasksResponseBodyData) SetAlgorithmName(v string) *DescribeComputeTasksResponseBodyData {
	s.AlgorithmName = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetDeviceNum(v int32) *DescribeComputeTasksResponseBodyData {
	s.DeviceNum = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetImageSize(v float32) *DescribeComputeTasksResponseBodyData {
	s.ImageSize = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetRuntime(v string) *DescribeComputeTasksResponseBodyData {
	s.Runtime = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetStructuredSize(v float32) *DescribeComputeTasksResponseBodyData {
	s.StructuredSize = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetTaskId(v string) *DescribeComputeTasksResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetTaskName(v string) *DescribeComputeTasksResponseBodyData {
	s.TaskName = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetTaskStatus(v string) *DescribeComputeTasksResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *DescribeComputeTasksResponseBodyData) SetVectorSize(v float32) *DescribeComputeTasksResponseBodyData {
	s.VectorSize = &v
	return s
}

type DescribeComputeTasksResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeComputeTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeComputeTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeComputeTasksResponse) GoString() string {
	return s.String()
}

func (s *DescribeComputeTasksResponse) SetHeaders(v map[string]*string) *DescribeComputeTasksResponse {
	s.Headers = v
	return s
}

func (s *DescribeComputeTasksResponse) SetBody(v *DescribeComputeTasksResponseBody) *DescribeComputeTasksResponse {
	s.Body = v
	return s
}

type DescribeDevicesRequest struct {
	FilterKey *string `json:"FilterKey,omitempty" xml:"FilterKey,omitempty"`
	PageNum   *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SearchKey *string `json:"SearchKey,omitempty" xml:"SearchKey,omitempty"`
	VcsId     *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s DescribeDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDevicesRequest) SetFilterKey(v string) *DescribeDevicesRequest {
	s.FilterKey = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageNum(v int32) *DescribeDevicesRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeDevicesRequest) SetPageSize(v int32) *DescribeDevicesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDevicesRequest) SetProjectId(v string) *DescribeDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *DescribeDevicesRequest) SetRegionId(v string) *DescribeDevicesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDevicesRequest) SetSearchKey(v string) *DescribeDevicesRequest {
	s.SearchKey = &v
	return s
}

func (s *DescribeDevicesRequest) SetVcsId(v string) *DescribeDevicesRequest {
	s.VcsId = &v
	return s
}

type DescribeDevicesResponseBody struct {
	Code       *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                             `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBody) SetCode(v string) *DescribeDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetData(v []*DescribeDevicesResponseBodyData) *DescribeDevicesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDevicesResponseBody) SetMessage(v string) *DescribeDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetRequestId(v string) *DescribeDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDevicesResponseBody) SetTotalCount(v int32) *DescribeDevicesResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeDevicesResponseBodyData struct {
	DeviceCode       *string  `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	DeviceId         *string  `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	DeviceName       *string  `json:"DeviceName,omitempty" xml:"DeviceName,omitempty"`
	DeviceType       *string  `json:"DeviceType,omitempty" xml:"DeviceType,omitempty"`
	FrameRate        *string  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	ImageSize        *float32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	Location         *string  `json:"Location,omitempty" xml:"Location,omitempty"`
	Owner            *string  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	PullStreamStatus *string  `json:"PullStreamStatus,omitempty" xml:"PullStreamStatus,omitempty"`
	RealTimeDataRate *string  `json:"RealTimeDataRate,omitempty" xml:"RealTimeDataRate,omitempty"`
	StructuredSize   *float32 `json:"StructuredSize,omitempty" xml:"StructuredSize,omitempty"`
	VectorSize       *float32 `json:"VectorSize,omitempty" xml:"VectorSize,omitempty"`
}

func (s DescribeDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponseBodyData) SetDeviceCode(v string) *DescribeDevicesResponseBodyData {
	s.DeviceCode = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetDeviceId(v string) *DescribeDevicesResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetDeviceName(v string) *DescribeDevicesResponseBodyData {
	s.DeviceName = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetDeviceType(v string) *DescribeDevicesResponseBodyData {
	s.DeviceType = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetFrameRate(v string) *DescribeDevicesResponseBodyData {
	s.FrameRate = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetImageSize(v float32) *DescribeDevicesResponseBodyData {
	s.ImageSize = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetLocation(v string) *DescribeDevicesResponseBodyData {
	s.Location = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetOwner(v string) *DescribeDevicesResponseBodyData {
	s.Owner = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetPullStreamStatus(v string) *DescribeDevicesResponseBodyData {
	s.PullStreamStatus = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetRealTimeDataRate(v string) *DescribeDevicesResponseBodyData {
	s.RealTimeDataRate = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetStructuredSize(v float32) *DescribeDevicesResponseBodyData {
	s.StructuredSize = &v
	return s
}

func (s *DescribeDevicesResponseBodyData) SetVectorSize(v float32) *DescribeDevicesResponseBodyData {
	s.VectorSize = &v
	return s
}

type DescribeDevicesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDevicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDevicesResponse) SetHeaders(v map[string]*string) *DescribeDevicesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDevicesResponse) SetBody(v *DescribeDevicesResponseBody) *DescribeDevicesResponse {
	s.Body = v
	return s
}

type DescribeProjectsRequest struct {
	PageNum     *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VcsId       *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s DescribeProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectsRequest) GoString() string {
	return s.String()
}

func (s *DescribeProjectsRequest) SetPageNum(v int32) *DescribeProjectsRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeProjectsRequest) SetPageSize(v int32) *DescribeProjectsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeProjectsRequest) SetProjectName(v string) *DescribeProjectsRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeProjectsRequest) SetRegionId(v string) *DescribeProjectsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeProjectsRequest) SetVcsId(v string) *DescribeProjectsRequest {
	s.VcsId = &v
	return s
}

type DescribeProjectsResponseBody struct {
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeProjectsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeProjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBody) SetCode(v string) *DescribeProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetData(v []*DescribeProjectsResponseBodyData) *DescribeProjectsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeProjectsResponseBody) SetMessage(v string) *DescribeProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetRequestId(v string) *DescribeProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProjectsResponseBody) SetTotalCount(v int32) *DescribeProjectsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeProjectsResponseBodyData struct {
	AlgorithmName  *string  `json:"AlgorithmName,omitempty" xml:"AlgorithmName,omitempty"`
	GbId           *string  `json:"GbId,omitempty" xml:"GbId,omitempty"`
	GbIp           *string  `json:"GbIp,omitempty" xml:"GbIp,omitempty"`
	GbPort         *string  `json:"GbPort,omitempty" xml:"GbPort,omitempty"`
	ImageSize      *float32 `json:"ImageSize,omitempty" xml:"ImageSize,omitempty"`
	PackagePattern *string  `json:"PackagePattern,omitempty" xml:"PackagePattern,omitempty"`
	ProjectCode    *string  `json:"ProjectCode,omitempty" xml:"ProjectCode,omitempty"`
	ProjectId      *string  `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	ProjectName    *string  `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Protocol       *string  `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	Status         *string  `json:"Status,omitempty" xml:"Status,omitempty"`
	StructuredSize *float32 `json:"StructuredSize,omitempty" xml:"StructuredSize,omitempty"`
	VectorSize     *float32 `json:"VectorSize,omitempty" xml:"VectorSize,omitempty"`
}

func (s DescribeProjectsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponseBodyData) SetAlgorithmName(v string) *DescribeProjectsResponseBodyData {
	s.AlgorithmName = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetGbId(v string) *DescribeProjectsResponseBodyData {
	s.GbId = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetGbIp(v string) *DescribeProjectsResponseBodyData {
	s.GbIp = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetGbPort(v string) *DescribeProjectsResponseBodyData {
	s.GbPort = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetImageSize(v float32) *DescribeProjectsResponseBodyData {
	s.ImageSize = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetPackagePattern(v string) *DescribeProjectsResponseBodyData {
	s.PackagePattern = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetProjectCode(v string) *DescribeProjectsResponseBodyData {
	s.ProjectCode = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetProjectId(v string) *DescribeProjectsResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetProjectName(v string) *DescribeProjectsResponseBodyData {
	s.ProjectName = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetProtocol(v string) *DescribeProjectsResponseBodyData {
	s.Protocol = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetStatus(v string) *DescribeProjectsResponseBodyData {
	s.Status = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetStructuredSize(v float32) *DescribeProjectsResponseBodyData {
	s.StructuredSize = &v
	return s
}

func (s *DescribeProjectsResponseBodyData) SetVectorSize(v float32) *DescribeProjectsResponseBodyData {
	s.VectorSize = &v
	return s
}

type DescribeProjectsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DescribeProjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeProjectsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectsResponse) SetHeaders(v map[string]*string) *DescribeProjectsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectsResponse) SetBody(v *DescribeProjectsResponseBody) *DescribeProjectsResponse {
	s.Body = v
	return s
}

type GetPictureSearchResultsRequest struct {
	AlgorithmType   *string `json:"AlgorithmType,omitempty" xml:"AlgorithmType,omitempty"`
	BeginTime       *string `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	DeviceList      *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	EndTime         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PageNum         *int32  `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	PictureContents *string `json:"PictureContents,omitempty" xml:"PictureContents,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TopNum          *int32  `json:"TopNum,omitempty" xml:"TopNum,omitempty"`
	VcsId           *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s GetPictureSearchResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchResultsRequest) GoString() string {
	return s.String()
}

func (s *GetPictureSearchResultsRequest) SetAlgorithmType(v string) *GetPictureSearchResultsRequest {
	s.AlgorithmType = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetBeginTime(v string) *GetPictureSearchResultsRequest {
	s.BeginTime = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetDeviceList(v string) *GetPictureSearchResultsRequest {
	s.DeviceList = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetEndTime(v string) *GetPictureSearchResultsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetPageNum(v int32) *GetPictureSearchResultsRequest {
	s.PageNum = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetPictureContents(v string) *GetPictureSearchResultsRequest {
	s.PictureContents = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetRegionId(v string) *GetPictureSearchResultsRequest {
	s.RegionId = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetTopNum(v int32) *GetPictureSearchResultsRequest {
	s.TopNum = &v
	return s
}

func (s *GetPictureSearchResultsRequest) SetVcsId(v string) *GetPictureSearchResultsRequest {
	s.VcsId = &v
	return s
}

type GetPictureSearchResultsResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      []*GetPictureSearchResultsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPictureSearchResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchResultsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPictureSearchResultsResponseBody) SetCode(v string) *GetPictureSearchResultsResponseBody {
	s.Code = &v
	return s
}

func (s *GetPictureSearchResultsResponseBody) SetData(v []*GetPictureSearchResultsResponseBodyData) *GetPictureSearchResultsResponseBody {
	s.Data = v
	return s
}

func (s *GetPictureSearchResultsResponseBody) SetMessage(v string) *GetPictureSearchResultsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPictureSearchResultsResponseBody) SetRequestId(v string) *GetPictureSearchResultsResponseBody {
	s.RequestId = &v
	return s
}

type GetPictureSearchResultsResponseBodyData struct {
	DeviceId                    *string  `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
	LeftUpperCornerXCoordinate  *string  `json:"LeftUpperCornerXCoordinate,omitempty" xml:"LeftUpperCornerXCoordinate,omitempty"`
	LeftUpperCornerYCoordinate  *string  `json:"LeftUpperCornerYCoordinate,omitempty" xml:"LeftUpperCornerYCoordinate,omitempty"`
	LocationMarkTime            *string  `json:"LocationMarkTime,omitempty" xml:"LocationMarkTime,omitempty"`
	PictureId1                  *string  `json:"PictureId1,omitempty" xml:"PictureId1,omitempty"`
	PictureId2                  *string  `json:"PictureId2,omitempty" xml:"PictureId2,omitempty"`
	PictureUrl1                 *string  `json:"PictureUrl1,omitempty" xml:"PictureUrl1,omitempty"`
	PictureUrl2                 *string  `json:"PictureUrl2,omitempty" xml:"PictureUrl2,omitempty"`
	RightLowerCornerXCoordinate *string  `json:"RightLowerCornerXCoordinate,omitempty" xml:"RightLowerCornerXCoordinate,omitempty"`
	RightLowerCornerYCoordinate *string  `json:"RightLowerCornerYCoordinate,omitempty" xml:"RightLowerCornerYCoordinate,omitempty"`
	Similarity                  *float32 `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s GetPictureSearchResultsResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchResultsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPictureSearchResultsResponseBodyData) SetDeviceId(v string) *GetPictureSearchResultsResponseBodyData {
	s.DeviceId = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetLeftUpperCornerXCoordinate(v string) *GetPictureSearchResultsResponseBodyData {
	s.LeftUpperCornerXCoordinate = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetLeftUpperCornerYCoordinate(v string) *GetPictureSearchResultsResponseBodyData {
	s.LeftUpperCornerYCoordinate = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetLocationMarkTime(v string) *GetPictureSearchResultsResponseBodyData {
	s.LocationMarkTime = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetPictureId1(v string) *GetPictureSearchResultsResponseBodyData {
	s.PictureId1 = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetPictureId2(v string) *GetPictureSearchResultsResponseBodyData {
	s.PictureId2 = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetPictureUrl1(v string) *GetPictureSearchResultsResponseBodyData {
	s.PictureUrl1 = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetPictureUrl2(v string) *GetPictureSearchResultsResponseBodyData {
	s.PictureUrl2 = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetRightLowerCornerXCoordinate(v string) *GetPictureSearchResultsResponseBodyData {
	s.RightLowerCornerXCoordinate = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetRightLowerCornerYCoordinate(v string) *GetPictureSearchResultsResponseBodyData {
	s.RightLowerCornerYCoordinate = &v
	return s
}

func (s *GetPictureSearchResultsResponseBodyData) SetSimilarity(v float32) *GetPictureSearchResultsResponseBodyData {
	s.Similarity = &v
	return s
}

type GetPictureSearchResultsResponse struct {
	Headers map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPictureSearchResultsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPictureSearchResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPictureSearchResultsResponse) GoString() string {
	return s.String()
}

func (s *GetPictureSearchResultsResponse) SetHeaders(v map[string]*string) *GetPictureSearchResultsResponse {
	s.Headers = v
	return s
}

func (s *GetPictureSearchResultsResponse) SetBody(v *GetPictureSearchResultsResponseBody) *GetPictureSearchResultsResponse {
	s.Body = v
	return s
}

type ImportDevicesRequest struct {
	DeviceList *string `json:"DeviceList,omitempty" xml:"DeviceList,omitempty"`
	ProjectId  *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VcsId      *string `json:"VcsId,omitempty" xml:"VcsId,omitempty"`
}

func (s ImportDevicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesRequest) GoString() string {
	return s.String()
}

func (s *ImportDevicesRequest) SetDeviceList(v string) *ImportDevicesRequest {
	s.DeviceList = &v
	return s
}

func (s *ImportDevicesRequest) SetProjectId(v string) *ImportDevicesRequest {
	s.ProjectId = &v
	return s
}

func (s *ImportDevicesRequest) SetRegionId(v string) *ImportDevicesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportDevicesRequest) SetVcsId(v string) *ImportDevicesRequest {
	s.VcsId = &v
	return s
}

type ImportDevicesResponseBody struct {
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *ImportDevicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportDevicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportDevicesResponseBody) SetCode(v string) *ImportDevicesResponseBody {
	s.Code = &v
	return s
}

func (s *ImportDevicesResponseBody) SetData(v *ImportDevicesResponseBodyData) *ImportDevicesResponseBody {
	s.Data = v
	return s
}

func (s *ImportDevicesResponseBody) SetMessage(v string) *ImportDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportDevicesResponseBody) SetRequestId(v string) *ImportDevicesResponseBody {
	s.RequestId = &v
	return s
}

type ImportDevicesResponseBodyData struct {
	Failure []*ImportDevicesResponseBodyDataFailure `json:"Failure,omitempty" xml:"Failure,omitempty" type:"Repeated"`
	Success []*ImportDevicesResponseBodyDataSuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Repeated"`
}

func (s ImportDevicesResponseBodyData) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportDevicesResponseBodyData) SetFailure(v []*ImportDevicesResponseBodyDataFailure) *ImportDevicesResponseBodyData {
	s.Failure = v
	return s
}

func (s *ImportDevicesResponseBodyData) SetSuccess(v []*ImportDevicesResponseBodyDataSuccess) *ImportDevicesResponseBodyData {
	s.Success = v
	return s
}

type ImportDevicesResponseBodyDataFailure struct {
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s ImportDevicesResponseBodyDataFailure) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesResponseBodyDataFailure) GoString() string {
	return s.String()
}

func (s *ImportDevicesResponseBodyDataFailure) SetDeviceCode(v string) *ImportDevicesResponseBodyDataFailure {
	s.DeviceCode = &v
	return s
}

func (s *ImportDevicesResponseBodyDataFailure) SetDeviceId(v string) *ImportDevicesResponseBodyDataFailure {
	s.DeviceId = &v
	return s
}

type ImportDevicesResponseBodyDataSuccess struct {
	DeviceCode *string `json:"DeviceCode,omitempty" xml:"DeviceCode,omitempty"`
	DeviceId   *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s ImportDevicesResponseBodyDataSuccess) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesResponseBodyDataSuccess) GoString() string {
	return s.String()
}

func (s *ImportDevicesResponseBodyDataSuccess) SetDeviceCode(v string) *ImportDevicesResponseBodyDataSuccess {
	s.DeviceCode = &v
	return s
}

func (s *ImportDevicesResponseBodyDataSuccess) SetDeviceId(v string) *ImportDevicesResponseBodyDataSuccess {
	s.DeviceId = &v
	return s
}

type ImportDevicesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ImportDevicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportDevicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportDevicesResponse) GoString() string {
	return s.String()
}

func (s *ImportDevicesResponse) SetHeaders(v map[string]*string) *ImportDevicesResponse {
	s.Headers = v
	return s
}

func (s *ImportDevicesResponse) SetBody(v *ImportDevicesResponseBody) *ImportDevicesResponse {
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("vcs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CreateComputeTaskWithOptions(request *CreateComputeTaskRequest, runtime *util.RuntimeOptions) (_result *CreateComputeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmCodeList)) {
		query["AlgorithmCodeList"] = request.AlgorithmCodeList
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceCodeList)) {
		query["DeviceCodeList"] = request.DeviceCodeList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskName)) {
		query["TaskName"] = request.TaskName
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateComputeTask"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateComputeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateComputeTask(request *CreateComputeTaskRequest) (_result *CreateComputeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateComputeTaskResponse{}
	_body, _err := client.CreateComputeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateProjectWithOptions(request *CreateProjectRequest, runtime *util.RuntimeOptions) (_result *CreateProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AreaCode)) {
		query["AreaCode"] = request.AreaCode
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TimeZoneCode)) {
		query["TimeZoneCode"] = request.TimeZoneCode
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateProject"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateProjectResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateProject(request *CreateProjectRequest) (_result *CreateProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateProjectResponse{}
	_body, _err := client.CreateProjectWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeComputeTasksWithOptions(request *DescribeComputeTasksRequest, runtime *util.RuntimeOptions) (_result *DescribeComputeTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeComputeTasks"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeComputeTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeComputeTasks(request *DescribeComputeTasksRequest) (_result *DescribeComputeTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeComputeTasksResponse{}
	_body, _err := client.DescribeComputeTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeDevicesWithOptions(request *DescribeDevicesRequest, runtime *util.RuntimeOptions) (_result *DescribeDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FilterKey)) {
		query["FilterKey"] = request.FilterKey
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.SearchKey)) {
		query["SearchKey"] = request.SearchKey
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDevices"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeDevices(request *DescribeDevicesRequest) (_result *DescribeDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDevicesResponse{}
	_body, _err := client.DescribeDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeProjectsWithOptions(request *DescribeProjectsRequest, runtime *util.RuntimeOptions) (_result *DescribeProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeProjects"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeProjectsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeProjects(request *DescribeProjectsRequest) (_result *DescribeProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeProjectsResponse{}
	_body, _err := client.DescribeProjectsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPictureSearchResultsWithOptions(request *GetPictureSearchResultsRequest, runtime *util.RuntimeOptions) (_result *GetPictureSearchResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgorithmType)) {
		query["AlgorithmType"] = request.AlgorithmType
	}

	if !tea.BoolValue(util.IsUnset(request.BeginTime)) {
		query["BeginTime"] = request.BeginTime
	}

	if !tea.BoolValue(util.IsUnset(request.DeviceList)) {
		query["DeviceList"] = request.DeviceList
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PictureContents)) {
		query["PictureContents"] = request.PictureContents
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.TopNum)) {
		query["TopNum"] = request.TopNum
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPictureSearchResults"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPictureSearchResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPictureSearchResults(request *GetPictureSearchResultsRequest) (_result *GetPictureSearchResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPictureSearchResultsResponse{}
	_body, _err := client.GetPictureSearchResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportDevicesWithOptions(request *ImportDevicesRequest, runtime *util.RuntimeOptions) (_result *ImportDevicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DeviceList)) {
		query["DeviceList"] = request.DeviceList
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectId)) {
		query["ProjectId"] = request.ProjectId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.VcsId)) {
		query["VcsId"] = request.VcsId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportDevices"),
		Version:     tea.String("2019-04-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportDevicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportDevices(request *ImportDevicesRequest) (_result *ImportDevicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportDevicesResponse{}
	_body, _err := client.ImportDevicesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
