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

type CancelImageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CancelImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelImageRequest) GoString() string {
	return s.String()
}

func (s *CancelImageRequest) SetRegionId(v string) *CancelImageRequest {
	s.RegionId = &v
	return s
}

type CancelImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CancelImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelImageResponse) GoString() string {
	return s.String()
}

func (s *CancelImageResponse) SetHeaders(v map[string]*string) *CancelImageResponse {
	s.Headers = v
	return s
}

type ChangeJobPriorityRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ChangeJobPriorityRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeJobPriorityRequest) GoString() string {
	return s.String()
}

func (s *ChangeJobPriorityRequest) SetRegionId(v string) *ChangeJobPriorityRequest {
	s.RegionId = &v
	return s
}

type ChangeJobPriorityResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ChangeJobPriorityResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeJobPriorityResponse) GoString() string {
	return s.String()
}

func (s *ChangeJobPriorityResponse) SetHeaders(v map[string]*string) *ChangeJobPriorityResponse {
	s.Headers = v
	return s
}

type CreateAppResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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

type CreateClusterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetRegionId(v string) *CreateClusterRequest {
	s.RegionId = &v
	return s
}

type CreateClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

type CreateImageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

type CreateImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CreateImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s *CreateImageResponse) SetHeaders(v map[string]*string) *CreateImageResponse {
	s.Headers = v
	return s
}

type CreateJobRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetRegionId(v string) *CreateJobRequest {
	s.RegionId = &v
	return s
}

type CreateJobResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

type DeleteAppResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteAppResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAppResponse) GoString() string {
	return s.String()
}

func (s *DeleteAppResponse) SetHeaders(v map[string]*string) *DeleteAppResponse {
	s.Headers = v
	return s
}

type DeleteClusterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetRegionId(v string) *DeleteClusterRequest {
	s.RegionId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

type DeleteClusterInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteClusterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterInstanceRequest) SetRegionId(v string) *DeleteClusterInstanceRequest {
	s.RegionId = &v
	return s
}

type DeleteClusterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterInstanceResponse) SetHeaders(v map[string]*string) *DeleteClusterInstanceResponse {
	s.Headers = v
	return s
}

type DeleteImageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) SetRegionId(v string) *DeleteImageRequest {
	s.RegionId = &v
	return s
}

type DeleteImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImageResponse) SetHeaders(v map[string]*string) *DeleteImageResponse {
	s.Headers = v
	return s
}

type DeleteJobRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteJobRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobRequest) SetRegionId(v string) *DeleteJobRequest {
	s.RegionId = &v
	return s
}

type DeleteJobResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

type DeleteProjectRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteProjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteProjectRequest) SetRegionId(v string) *DeleteProjectRequest {
	s.RegionId = &v
	return s
}

type DeleteProjectResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteProjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteProjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponse) SetHeaders(v map[string]*string) *DeleteProjectResponse {
	s.Headers = v
	return s
}

type GetAppResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetAppResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAppResponse) GoString() string {
	return s.String()
}

func (s *GetAppResponse) SetHeaders(v map[string]*string) *GetAppResponse {
	s.Headers = v
	return s
}

type GetClusterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterRequest) GoString() string {
	return s.String()
}

func (s *GetClusterRequest) SetRegionId(v string) *GetClusterRequest {
	s.RegionId = &v
	return s
}

type GetClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterResponse) GoString() string {
	return s.String()
}

func (s *GetClusterResponse) SetHeaders(v map[string]*string) *GetClusterResponse {
	s.Headers = v
	return s
}

type GetClusterInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetClusterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetClusterInstanceRequest) SetRegionId(v string) *GetClusterInstanceRequest {
	s.RegionId = &v
	return s
}

type GetClusterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetClusterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetClusterInstanceResponse) SetHeaders(v map[string]*string) *GetClusterInstanceResponse {
	s.Headers = v
	return s
}

type GetImageRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s *GetImageRequest) SetRegionId(v string) *GetImageRequest {
	s.RegionId = &v
	return s
}

type GetImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s *GetImageResponse) SetHeaders(v map[string]*string) *GetImageResponse {
	s.Headers = v
	return s
}

type GetInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceRequest) SetRegionId(v string) *GetInstanceRequest {
	s.RegionId = &v
	return s
}

type GetInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceResponse) SetHeaders(v map[string]*string) *GetInstanceResponse {
	s.Headers = v
	return s
}

type GetJobRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetJobRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) SetRegionId(v string) *GetJobRequest {
	s.RegionId = &v
	return s
}

type GetJobResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

type GetJobDescriptionRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetJobDescriptionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobDescriptionRequest) GoString() string {
	return s.String()
}

func (s *GetJobDescriptionRequest) SetRegionId(v string) *GetJobDescriptionRequest {
	s.RegionId = &v
	return s
}

type GetJobDescriptionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetJobDescriptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobDescriptionResponse) GoString() string {
	return s.String()
}

func (s *GetJobDescriptionResponse) SetHeaders(v map[string]*string) *GetJobDescriptionResponse {
	s.Headers = v
	return s
}

type GetQuotaRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetRegionId(v string) *GetQuotaRequest {
	s.RegionId = &v
	return s
}

type GetQuotaResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

type GetTaskRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskRequest) GoString() string {
	return s.String()
}

func (s *GetTaskRequest) SetRegionId(v string) *GetTaskRequest {
	s.RegionId = &v
	return s
}

type GetTaskResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskResponse) GoString() string {
	return s.String()
}

func (s *GetTaskResponse) SetHeaders(v map[string]*string) *GetTaskResponse {
	s.Headers = v
	return s
}

type ListAppsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
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

type ListAvailableInstanceTypeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListAvailableInstanceTypeResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAvailableInstanceTypeResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableInstanceTypeResponse) SetHeaders(v map[string]*string) *ListAvailableInstanceTypeResponse {
	s.Headers = v
	return s
}

type ListClusterInstancesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClusterInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterInstancesRequest) SetRegionId(v string) *ListClusterInstancesRequest {
	s.RegionId = &v
	return s
}

type ListClusterInstancesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListClusterInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterInstancesResponse) SetHeaders(v map[string]*string) *ListClusterInstancesResponse {
	s.Headers = v
	return s
}

type ListClustersRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetRegionId(v string) *ListClustersRequest {
	s.RegionId = &v
	return s
}

type ListClustersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

type ListImagesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

type ListInstancesRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

type ListInstancesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesResponse) SetHeaders(v map[string]*string) *ListInstancesResponse {
	s.Headers = v
	return s
}

type ListJobsRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetRegionId(v string) *ListJobsRequest {
	s.RegionId = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

type ListRegionsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponse) GoString() string {
	return s.String()
}

func (s *ListRegionsResponse) SetHeaders(v map[string]*string) *ListRegionsResponse {
	s.Headers = v
	return s
}

type ListTasksRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetRegionId(v string) *ListTasksRequest {
	s.RegionId = &v
	return s
}

type ListTasksResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListTasksResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponse) GoString() string {
	return s.String()
}

func (s *ListTasksResponse) SetHeaders(v map[string]*string) *ListTasksResponse {
	s.Headers = v
	return s
}

type ModifyAppResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyAppResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyAppResponse) GoString() string {
	return s.String()
}

func (s *ModifyAppResponse) SetHeaders(v map[string]*string) *ModifyAppResponse {
	s.Headers = v
	return s
}

type ModifyClusterRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterRequest) SetRegionId(v string) *ModifyClusterRequest {
	s.RegionId = &v
	return s
}

type ModifyClusterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ModifyClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ModifyClusterResponse) GoString() string {
	return s.String()
}

func (s *ModifyClusterResponse) SetHeaders(v map[string]*string) *ModifyClusterResponse {
	s.Headers = v
	return s
}

type PollForTaskRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s PollForTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s PollForTaskRequest) GoString() string {
	return s.String()
}

func (s *PollForTaskRequest) SetRegionId(v string) *PollForTaskRequest {
	s.RegionId = &v
	return s
}

type PollForTaskResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PollForTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s PollForTaskResponse) GoString() string {
	return s.String()
}

func (s *PollForTaskResponse) SetHeaders(v map[string]*string) *PollForTaskResponse {
	s.Headers = v
	return s
}

type RecreateClusterInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RecreateClusterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RecreateClusterInstanceRequest) GoString() string {
	return s.String()
}

func (s *RecreateClusterInstanceRequest) SetRegionId(v string) *RecreateClusterInstanceRequest {
	s.RegionId = &v
	return s
}

type RecreateClusterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RecreateClusterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RecreateClusterInstanceResponse) GoString() string {
	return s.String()
}

func (s *RecreateClusterInstanceResponse) SetHeaders(v map[string]*string) *RecreateClusterInstanceResponse {
	s.Headers = v
	return s
}

type RenewClusterInstanceRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RenewClusterInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewClusterInstanceRequest) GoString() string {
	return s.String()
}

func (s *RenewClusterInstanceRequest) SetRegionId(v string) *RenewClusterInstanceRequest {
	s.RegionId = &v
	return s
}

type RenewClusterInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s RenewClusterInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewClusterInstanceResponse) GoString() string {
	return s.String()
}

func (s *RenewClusterInstanceResponse) SetHeaders(v map[string]*string) *RenewClusterInstanceResponse {
	s.Headers = v
	return s
}

type ReportTaskStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReportTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportTaskStatusRequest) SetRegionId(v string) *ReportTaskStatusRequest {
	s.RegionId = &v
	return s
}

type ReportTaskStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ReportTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportTaskStatusResponse) SetHeaders(v map[string]*string) *ReportTaskStatusResponse {
	s.Headers = v
	return s
}

type ReportWorkerStatusRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReportWorkerStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s ReportWorkerStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportWorkerStatusRequest) SetRegionId(v string) *ReportWorkerStatusRequest {
	s.RegionId = &v
	return s
}

type ReportWorkerStatusResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ReportWorkerStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportWorkerStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportWorkerStatusResponse) SetHeaders(v map[string]*string) *ReportWorkerStatusResponse {
	s.Headers = v
	return s
}

type StartJobRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StartJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StartJobRequest) GoString() string {
	return s.String()
}

func (s *StartJobRequest) SetRegionId(v string) *StartJobRequest {
	s.RegionId = &v
	return s
}

type StartJobResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s StartJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StartJobResponse) GoString() string {
	return s.String()
}

func (s *StartJobResponse) SetHeaders(v map[string]*string) *StartJobResponse {
	s.Headers = v
	return s
}

type StopJobRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopJobRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobRequest) GoString() string {
	return s.String()
}

func (s *StopJobRequest) SetRegionId(v string) *StopJobRequest {
	s.RegionId = &v
	return s
}

type StopJobResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s StopJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
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
	client.Endpoint, _err = client.GetEndpoint(tea.String("batchcompute"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelImage(ResourceName *string, request *CancelImageRequest) (_result *CancelImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelImageResponse{}
	_body, _err := client.CancelImageWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CancelImageWithOptions(ResourceName *string, request *CancelImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CancelImageResponse{}
	_body, _err := client.DoROARequest(tea.String("CancelImage"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/images/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ChangeJobPriority(ResourceName *string, request *ChangeJobPriorityRequest) (_result *ChangeJobPriorityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ChangeJobPriorityResponse{}
	_body, _err := client.ChangeJobPriorityWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ChangeJobPriorityWithOptions(ResourceName *string, request *ChangeJobPriorityRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ChangeJobPriorityResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ChangeJobPriorityResponse{}
	_body, _err := client.DoROARequest(tea.String("ChangeJobPriority"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateApp() (_result *CreateAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAppResponse{}
	_body, _err := client.CreateAppWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAppWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &CreateAppResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateApp"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/apps"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCluster"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateImage"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/images"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateJob"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/jobs"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteApp(ResourceName *string) (_result *DeleteAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteAppResponse{}
	_body, _err := client.DeleteAppWithOptions(ResourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAppWithOptions(ResourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteAppResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteApp"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/apps/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(ResourceName *string, request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(ResourceName *string, request *DeleteClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteCluster"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterInstance(ClusterId *string, GroupName *string, InstanceId *string, request *DeleteClusterInstanceRequest) (_result *DeleteClusterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteClusterInstanceResponse{}
	_body, _err := client.DeleteClusterInstanceWithOptions(ClusterId, GroupName, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterInstanceWithOptions(ClusterId *string, GroupName *string, InstanceId *string, request *DeleteClusterInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteClusterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteClusterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteClusterInstance"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/groups/{GroupName}/instances/{InstanceId}"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImage(ResourceName *string, request *DeleteImageRequest) (_result *DeleteImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteImageResponse{}
	_body, _err := client.DeleteImageWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImageWithOptions(ResourceName *string, request *DeleteImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteImageResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteImage"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/images/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(ResourceName *string, request *DeleteJobRequest) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(ResourceName *string, request *DeleteJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteJob"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteProject(ProjectName *string, request *DeleteProjectRequest) (_result *DeleteProjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteProjectResponse{}
	_body, _err := client.DeleteProjectWithOptions(ProjectName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteProjectWithOptions(ProjectName *string, request *DeleteProjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteProjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteProjectResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteProject"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/projects/"+tea.StringValue(ProjectName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetApp(ResourceName *string) (_result *GetAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAppResponse{}
	_body, _err := client.GetAppWithOptions(ResourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAppWithOptions(ResourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetAppResponse{}
	_body, _err := client.DoROARequest(tea.String("GetApp"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/apps/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCluster(ResourceName *string, request *GetClusterRequest) (_result *GetClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterResponse{}
	_body, _err := client.GetClusterWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterWithOptions(ResourceName *string, request *GetClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetClusterResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCluster"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterInstance(ClusterId *string, GroupName *string, InstanceId *string, request *GetClusterInstanceRequest) (_result *GetClusterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetClusterInstanceResponse{}
	_body, _err := client.GetClusterInstanceWithOptions(ClusterId, GroupName, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterInstanceWithOptions(ClusterId *string, GroupName *string, InstanceId *string, request *GetClusterInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetClusterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetClusterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetClusterInstance"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/groups/{GroupName}/instances/{InstanceId}"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetImage(ResourceName *string, request *GetImageRequest) (_result *GetImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetImageResponse{}
	_body, _err := client.GetImageWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetImageWithOptions(ResourceName *string, request *GetImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetImageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetImage"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/images/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetInstance(ResourceName *string, TaskName *string, InstanceId *string, request *GetInstanceRequest) (_result *GetInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetInstanceResponse{}
	_body, _err := client.GetInstanceWithOptions(ResourceName, TaskName, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetInstanceWithOptions(ResourceName *string, TaskName *string, InstanceId *string, request *GetInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetInstance"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/tasks/{TaskName}/instances/{InstanceId}"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(ResourceName *string, request *GetJobRequest) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(ResourceName *string, request *GetJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJob"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobDescription(ResourceName *string, request *GetJobDescriptionRequest) (_result *GetJobDescriptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobDescriptionResponse{}
	_body, _err := client.GetJobDescriptionWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobDescriptionWithOptions(ResourceName *string, request *GetJobDescriptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobDescriptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobDescriptionResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobDescription"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"?description"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuota(request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaWithOptions(request *GetQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("GetQuota"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/quotas"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTask(ResourceName *string, TaskName *string, request *GetTaskRequest) (_result *GetTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskResponse{}
	_body, _err := client.GetTaskWithOptions(ResourceName, TaskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTaskWithOptions(ResourceName *string, TaskName *string, request *GetTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTask"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/tasks/{TaskName}"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListApps() (_result *ListAppsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAppsResponse{}
	_body, _err := client.ListAppsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAppsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAppsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListAppsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListApps"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/apps"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAvailableInstanceType() (_result *ListAvailableInstanceTypeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAvailableInstanceTypeResponse{}
	_body, _err := client.ListAvailableInstanceTypeWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAvailableInstanceTypeWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAvailableInstanceTypeResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListAvailableInstanceTypeResponse{}
	_body, _err := client.DoROARequest(tea.String("ListAvailableInstanceType"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/available"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusterInstances(ClusterId *string, GroupName *string, request *ListClusterInstancesRequest) (_result *ListClusterInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClusterInstancesResponse{}
	_body, _err := client.ListClusterInstancesWithOptions(ClusterId, GroupName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClusterInstancesWithOptions(ClusterId *string, GroupName *string, request *ListClusterInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClusterInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListClusterInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListClusterInstances"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/groups/{GroupName}/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListClustersWithOptions(request *ListClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.DoROARequest(tea.String("ListClusters"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListImages"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/images"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(ResourceName *string, TaskName *string, request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(ResourceName, TaskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(ResourceName *string, TaskName *string, request *ListInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListInstances"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/tasks/{TaskName}/instances"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(request *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListJobs"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions() (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListRegions"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/regions"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(ResourceName *string, request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(ResourceName *string, request *ListTasksRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTasks"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/tasks"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyApp(ResourceName *string) (_result *ModifyAppResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyAppResponse{}
	_body, _err := client.ModifyAppWithOptions(ResourceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyAppWithOptions(ResourceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyAppResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ModifyAppResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyApp"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/apps/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ModifyCluster(ResourceName *string, request *ModifyClusterRequest) (_result *ModifyClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ModifyClusterResponse{}
	_body, _err := client.ModifyClusterWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ModifyClusterWithOptions(ResourceName *string, request *ModifyClusterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ModifyClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ModifyClusterResponse{}
	_body, _err := client.DoROARequest(tea.String("ModifyCluster"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ResourceName)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PollForTask(ClusterId *string, WorkerId *string, request *PollForTaskRequest) (_result *PollForTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PollForTaskResponse{}
	_body, _err := client.PollForTaskWithOptions(ClusterId, WorkerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PollForTaskWithOptions(ClusterId *string, WorkerId *string, request *PollForTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PollForTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &PollForTaskResponse{}
	_body, _err := client.DoROARequest(tea.String("PollForTask"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/workers/{WorkerId}/fetchTask"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RecreateClusterInstance(ClusterId *string, GroupName *string, InstanceId *string, request *RecreateClusterInstanceRequest) (_result *RecreateClusterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RecreateClusterInstanceResponse{}
	_body, _err := client.RecreateClusterInstanceWithOptions(ClusterId, GroupName, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RecreateClusterInstanceWithOptions(ClusterId *string, GroupName *string, InstanceId *string, request *RecreateClusterInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RecreateClusterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &RecreateClusterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("RecreateClusterInstance"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/groups/{GroupName}/instances/{InstanceId}/recreate"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewClusterInstance(ClusterId *string, GroupName *string, InstanceId *string, request *RenewClusterInstanceRequest) (_result *RenewClusterInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RenewClusterInstanceResponse{}
	_body, _err := client.RenewClusterInstanceWithOptions(ClusterId, GroupName, InstanceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewClusterInstanceWithOptions(ClusterId *string, GroupName *string, InstanceId *string, request *RenewClusterInstanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RenewClusterInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &RenewClusterInstanceResponse{}
	_body, _err := client.DoROARequest(tea.String("RenewClusterInstance"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/groups/{GroupName}/instances/{InstanceId}/renew"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportTaskStatus(ClusterId *string, WorkerId *string, request *ReportTaskStatusRequest) (_result *ReportTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportTaskStatusResponse{}
	_body, _err := client.ReportTaskStatusWithOptions(ClusterId, WorkerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportTaskStatusWithOptions(ClusterId *string, WorkerId *string, request *ReportTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReportTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ReportTaskStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("ReportTaskStatus"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/workers/{WorkerId}/updateTaskStatus"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportWorkerStatus(ClusterId *string, WorkerId *string, request *ReportWorkerStatusRequest) (_result *ReportWorkerStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportWorkerStatusResponse{}
	_body, _err := client.ReportWorkerStatusWithOptions(ClusterId, WorkerId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportWorkerStatusWithOptions(ClusterId *string, WorkerId *string, request *ReportWorkerStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReportWorkerStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ReportWorkerStatusResponse{}
	_body, _err := client.DoROARequest(tea.String("ReportWorkerStatus"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(ClusterId)+"/workers/{WorkerId}/updateStatus"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartJob(ResourceName *string, request *StartJobRequest) (_result *StartJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartJobResponse{}
	_body, _err := client.StartJobWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartJobWithOptions(ResourceName *string, request *StartJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &StartJobResponse{}
	_body, _err := client.DoROARequest(tea.String("StartJob"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/start"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJob(ResourceName *string, request *StopJobRequest) (_result *StopJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(ResourceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobWithOptions(ResourceName *string, request *StopJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &StopJobResponse{}
	_body, _err := client.DoROARequest(tea.String("StopJob"), tea.String("2015-11-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/jobs/"+tea.StringValue(ResourceName)+"/stop"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
