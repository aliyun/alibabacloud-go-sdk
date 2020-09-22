// This file is auto-generated, don't edit it. Thanks.
package client

import (
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	roa "github.com/alibabacloud-go/tea-roa/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type GetProjectEventsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetProjectEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetProjectEventsRequest) GoString() string {
	return s.String()
}

func (s *GetProjectEventsRequest) SetHeaders(v map[string]*string) *GetProjectEventsRequest {
	s.Headers = v
	return s
}

type GetProjectEventsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetProjectEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetProjectEventsResponse) GoString() string {
	return s.String()
}

func (s *GetProjectEventsResponse) SetHeaders(v map[string]*string) *GetProjectEventsResponse {
	s.Headers = v
	return s
}

type DescribeKubernetesTemplateRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeKubernetesTemplateRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesTemplateRequest) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesTemplateRequest) SetHeaders(v map[string]*string) *DescribeKubernetesTemplateRequest {
	s.Headers = v
	return s
}

type DescribeKubernetesTemplateResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeKubernetesTemplateResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesTemplateResponse) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesTemplateResponse) SetHeaders(v map[string]*string) *DescribeKubernetesTemplateResponse {
	s.Headers = v
	return s
}

type DescribeAgilityTunnelCertsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeAgilityTunnelCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgilityTunnelCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgilityTunnelCertsRequest) SetHeaders(v map[string]*string) *DescribeAgilityTunnelCertsRequest {
	s.Headers = v
	return s
}

type DescribeAgilityTunnelCertsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeAgilityTunnelCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgilityTunnelCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgilityTunnelCertsResponse) SetHeaders(v map[string]*string) *DescribeAgilityTunnelCertsResponse {
	s.Headers = v
	return s
}

type DescribeClusterTokensRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterTokensRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTokensRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterTokensRequest) SetHeaders(v map[string]*string) *DescribeClusterTokensRequest {
	s.Headers = v
	return s
}

type DescribeClusterTokensResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterTokensResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterTokensResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterTokensResponse) SetHeaders(v map[string]*string) *DescribeClusterTokensResponse {
	s.Headers = v
	return s
}

type DownloadClusterNodeCertsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DownloadClusterNodeCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadClusterNodeCertsRequest) GoString() string {
	return s.String()
}

func (s *DownloadClusterNodeCertsRequest) SetHeaders(v map[string]*string) *DownloadClusterNodeCertsRequest {
	s.Headers = v
	return s
}

type DownloadClusterNodeCertsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DownloadClusterNodeCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadClusterNodeCertsResponse) GoString() string {
	return s.String()
}

func (s *DownloadClusterNodeCertsResponse) SetHeaders(v map[string]*string) *DownloadClusterNodeCertsResponse {
	s.Headers = v
	return s
}

type DescribeServiceContainersRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeServiceContainersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceContainersRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceContainersRequest) SetHeaders(v map[string]*string) *DescribeServiceContainersRequest {
	s.Headers = v
	return s
}

type DescribeServiceContainersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeServiceContainersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceContainersResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceContainersResponse) SetHeaders(v map[string]*string) *DescribeServiceContainersResponse {
	s.Headers = v
	return s
}

type GatherLogsTokenRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GatherLogsTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GatherLogsTokenRequest) GoString() string {
	return s.String()
}

func (s *GatherLogsTokenRequest) SetHeaders(v map[string]*string) *GatherLogsTokenRequest {
	s.Headers = v
	return s
}

type GatherLogsTokenResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GatherLogsTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GatherLogsTokenResponse) GoString() string {
	return s.String()
}

func (s *GatherLogsTokenResponse) SetHeaders(v map[string]*string) *GatherLogsTokenResponse {
	s.Headers = v
	return s
}

type GetClusterProjectsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetClusterProjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetClusterProjectsRequest) GoString() string {
	return s.String()
}

func (s *GetClusterProjectsRequest) SetHeaders(v map[string]*string) *GetClusterProjectsRequest {
	s.Headers = v
	return s
}

type GetClusterProjectsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetClusterProjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetClusterProjectsResponse) GoString() string {
	return s.String()
}

func (s *GetClusterProjectsResponse) SetHeaders(v map[string]*string) *GetClusterProjectsResponse {
	s.Headers = v
	return s
}

type DescribeApiVersionRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeApiVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiVersionRequest) GoString() string {
	return s.String()
}

func (s *DescribeApiVersionRequest) SetHeaders(v map[string]*string) *DescribeApiVersionRequest {
	s.Headers = v
	return s
}

type DescribeApiVersionResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeApiVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeApiVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiVersionResponse) SetHeaders(v map[string]*string) *DescribeApiVersionResponse {
	s.Headers = v
	return s
}

type DescribeTaskInfoRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeTaskInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoRequest) SetHeaders(v map[string]*string) *DescribeTaskInfoRequest {
	s.Headers = v
	return s
}

type DescribeTaskInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeTaskInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskInfoResponse) SetHeaders(v map[string]*string) *DescribeTaskInfoResponse {
	s.Headers = v
	return s
}

type DescribeClusterNodesQuery struct {
	PageSize   *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
}

func (s DescribeClusterNodesQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesQuery) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesQuery) SetPageSize(v string) *DescribeClusterNodesQuery {
	s.PageSize = &v
	return s
}

func (s *DescribeClusterNodesQuery) SetPageNumber(v string) *DescribeClusterNodesQuery {
	s.PageNumber = &v
	return s
}

type DescribeClusterNodesRequest struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClusterNodesQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesRequest) SetHeaders(v map[string]*string) *DescribeClusterNodesRequest {
	s.Headers = v
	return s
}

func (s *DescribeClusterNodesRequest) SetQuery(v *DescribeClusterNodesQuery) *DescribeClusterNodesRequest {
	s.Query = v
	return s
}

type DescribeClusterNodesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodesResponse) SetHeaders(v map[string]*string) *DescribeClusterNodesResponse {
	s.Headers = v
	return s
}

type DescribeAgilityTunnelAgentInfoRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeAgilityTunnelAgentInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgilityTunnelAgentInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeAgilityTunnelAgentInfoRequest) SetHeaders(v map[string]*string) *DescribeAgilityTunnelAgentInfoRequest {
	s.Headers = v
	return s
}

type DescribeAgilityTunnelAgentInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeAgilityTunnelAgentInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeAgilityTunnelAgentInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgilityTunnelAgentInfoResponse) SetHeaders(v map[string]*string) *DescribeAgilityTunnelAgentInfoResponse {
	s.Headers = v
	return s
}

type DeleteClusterNodeQuery struct {
	Force           *string `json:"force,omitempty" xml:"force,omitempty"`
	ReleaseInstance *string `json:"releaseInstance,omitempty" xml:"releaseInstance,omitempty"`
}

func (s DeleteClusterNodeQuery) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodeQuery) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodeQuery) SetForce(v string) *DeleteClusterNodeQuery {
	s.Force = &v
	return s
}

func (s *DeleteClusterNodeQuery) SetReleaseInstance(v string) *DeleteClusterNodeQuery {
	s.ReleaseInstance = &v
	return s
}

type DeleteClusterNodeRequest struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DeleteClusterNodeQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DeleteClusterNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodeRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodeRequest) SetHeaders(v map[string]*string) *DeleteClusterNodeRequest {
	s.Headers = v
	return s
}

func (s *DeleteClusterNodeRequest) SetQuery(v *DeleteClusterNodeQuery) *DeleteClusterNodeRequest {
	s.Query = v
	return s
}

type DeleteClusterNodeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteClusterNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterNodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterNodeResponse) SetHeaders(v map[string]*string) *DeleteClusterNodeResponse {
	s.Headers = v
	return s
}

type DescribeClusterNodeInfoWithInstanceRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterNodeInfoWithInstanceRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodeInfoWithInstanceRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodeInfoWithInstanceRequest) SetHeaders(v map[string]*string) *DescribeClusterNodeInfoWithInstanceRequest {
	s.Headers = v
	return s
}

type DescribeClusterNodeInfoWithInstanceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNodeInfoWithInstanceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodeInfoWithInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodeInfoWithInstanceResponse) SetHeaders(v map[string]*string) *DescribeClusterNodeInfoWithInstanceResponse {
	s.Headers = v
	return s
}

type DescribeUserContainersQuery struct {
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s DescribeUserContainersQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserContainersQuery) GoString() string {
	return s.String()
}

func (s *DescribeUserContainersQuery) SetServiceId(v string) *DescribeUserContainersQuery {
	s.ServiceId = &v
	return s
}

type DescribeUserContainersRequest struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeUserContainersQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeUserContainersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserContainersRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserContainersRequest) SetHeaders(v map[string]*string) *DescribeUserContainersRequest {
	s.Headers = v
	return s
}

func (s *DescribeUserContainersRequest) SetQuery(v *DescribeUserContainersQuery) *DescribeUserContainersRequest {
	s.Query = v
	return s
}

type DescribeUserContainersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeUserContainersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeUserContainersResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserContainersResponse) SetHeaders(v map[string]*string) *DescribeUserContainersResponse {
	s.Headers = v
	return s
}

type CreateClusterTokenRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CreateClusterTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterTokenRequest) SetHeaders(v map[string]*string) *CreateClusterTokenRequest {
	s.Headers = v
	return s
}

type CreateClusterTokenResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CreateClusterTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterTokenResponse) SetHeaders(v map[string]*string) *CreateClusterTokenResponse {
	s.Headers = v
	return s
}

type DescribeClusterHostsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterHostsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostsRequest) SetHeaders(v map[string]*string) *DescribeClusterHostsRequest {
	s.Headers = v
	return s
}

type DescribeClusterHostsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterHostsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterHostsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterHostsResponse) SetHeaders(v map[string]*string) *DescribeClusterHostsResponse {
	s.Headers = v
	return s
}

type DescribeKubernetesTemplatesQuery struct {
	KubernetesVersion *string `json:"KubernetesVersion,omitempty" xml:"KubernetesVersion,omitempty"`
	Region            *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DescribeKubernetesTemplatesQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesTemplatesQuery) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesTemplatesQuery) SetKubernetesVersion(v string) *DescribeKubernetesTemplatesQuery {
	s.KubernetesVersion = &v
	return s
}

func (s *DescribeKubernetesTemplatesQuery) SetRegion(v string) *DescribeKubernetesTemplatesQuery {
	s.Region = &v
	return s
}

type DescribeKubernetesTemplatesRequest struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeKubernetesTemplatesQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeKubernetesTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesTemplatesRequest) SetHeaders(v map[string]*string) *DescribeKubernetesTemplatesRequest {
	s.Headers = v
	return s
}

func (s *DescribeKubernetesTemplatesRequest) SetQuery(v *DescribeKubernetesTemplatesQuery) *DescribeKubernetesTemplatesRequest {
	s.Query = v
	return s
}

type DescribeKubernetesTemplatesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeKubernetesTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeKubernetesTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeKubernetesTemplatesResponse) SetHeaders(v map[string]*string) *DescribeKubernetesTemplatesResponse {
	s.Headers = v
	return s
}

type DescribeTemplatesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeTemplatesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesRequest) SetHeaders(v map[string]*string) *DescribeTemplatesRequest {
	s.Headers = v
	return s
}

type DescribeTemplatesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeTemplatesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplatesResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatesResponse) SetHeaders(v map[string]*string) *DescribeTemplatesResponse {
	s.Headers = v
	return s
}

type DescribeClusterScaledNodeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterScaledNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterScaledNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterScaledNodeRequest) SetHeaders(v map[string]*string) *DescribeClusterScaledNodeRequest {
	s.Headers = v
	return s
}

type DescribeClusterScaledNodeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterScaledNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterScaledNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterScaledNodeResponse) SetHeaders(v map[string]*string) *DescribeClusterScaledNodeResponse {
	s.Headers = v
	return s
}

type CallbackClusterTokenRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CallbackClusterTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s CallbackClusterTokenRequest) GoString() string {
	return s.String()
}

func (s *CallbackClusterTokenRequest) SetHeaders(v map[string]*string) *CallbackClusterTokenRequest {
	s.Headers = v
	return s
}

type CallbackClusterTokenResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s CallbackClusterTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s CallbackClusterTokenResponse) GoString() string {
	return s.String()
}

func (s *CallbackClusterTokenResponse) SetHeaders(v map[string]*string) *CallbackClusterTokenResponse {
	s.Headers = v
	return s
}

type DescribeImagesQuery struct {
	RegionID      *string `json:"RegionID,omitempty" xml:"RegionID,omitempty"`
	DockerVersion *string `json:"DockerVersion,omitempty" xml:"DockerVersion,omitempty"`
	ImageName     *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
}

func (s DescribeImagesQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesQuery) GoString() string {
	return s.String()
}

func (s *DescribeImagesQuery) SetRegionID(v string) *DescribeImagesQuery {
	s.RegionID = &v
	return s
}

func (s *DescribeImagesQuery) SetDockerVersion(v string) *DescribeImagesQuery {
	s.DockerVersion = &v
	return s
}

func (s *DescribeImagesQuery) SetImageName(v string) *DescribeImagesQuery {
	s.ImageName = &v
	return s
}

type DescribeImagesRequest struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeImagesQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImagesRequest) SetHeaders(v map[string]*string) *DescribeImagesRequest {
	s.Headers = v
	return s
}

func (s *DescribeImagesRequest) SetQuery(v *DescribeImagesQuery) *DescribeImagesRequest {
	s.Query = v
	return s
}

type DescribeImagesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeImagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeImagesResponse) SetHeaders(v map[string]*string) *DescribeImagesResponse {
	s.Headers = v
	return s
}

type DescribeClusterLogsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsRequest) SetHeaders(v map[string]*string) *DescribeClusterLogsRequest {
	s.Headers = v
	return s
}

type DescribeClusterLogsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterLogsResponse) SetHeaders(v map[string]*string) *DescribeClusterLogsResponse {
	s.Headers = v
	return s
}

type DescribeClusterServicesRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServicesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterServicesRequest) SetHeaders(v map[string]*string) *DescribeClusterServicesRequest {
	s.Headers = v
	return s
}

type DescribeClusterServicesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterServicesResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterServicesResponse) SetHeaders(v map[string]*string) *DescribeClusterServicesResponse {
	s.Headers = v
	return s
}

type GetTriggerHookRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s GetTriggerHookRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerHookRequest) GoString() string {
	return s.String()
}

func (s *GetTriggerHookRequest) SetHeaders(v map[string]*string) *GetTriggerHookRequest {
	s.Headers = v
	return s
}

type GetTriggerHookResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetTriggerHookResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTriggerHookResponse) GoString() string {
	return s.String()
}

func (s *GetTriggerHookResponse) SetHeaders(v map[string]*string) *GetTriggerHookResponse {
	s.Headers = v
	return s
}

type DescribeTemplateAttributeRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeTemplateAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeRequest) SetHeaders(v map[string]*string) *DescribeTemplateAttributeRequest {
	s.Headers = v
	return s
}

type DescribeTemplateAttributeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeTemplateAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTemplateAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplateAttributeResponse) SetHeaders(v map[string]*string) *DescribeTemplateAttributeResponse {
	s.Headers = v
	return s
}

type DescribeClusterCertsRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterCertsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterCertsRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterCertsRequest) SetHeaders(v map[string]*string) *DescribeClusterCertsRequest {
	s.Headers = v
	return s
}

type DescribeClusterCertsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterCertsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterCertsResponse) SetHeaders(v map[string]*string) *DescribeClusterCertsResponse {
	s.Headers = v
	return s
}

type DescribeClusterNodeInfoRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterNodeInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodeInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodeInfoRequest) SetHeaders(v map[string]*string) *DescribeClusterNodeInfoRequest {
	s.Headers = v
	return s
}

type DescribeClusterNodeInfoResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterNodeInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterNodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterNodeInfoResponse) SetHeaders(v map[string]*string) *DescribeClusterNodeInfoResponse {
	s.Headers = v
	return s
}

type DeleteClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetHeaders(v map[string]*string) *DeleteClusterRequest {
	s.Headers = v
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

type CreateClusterRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s CreateClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) SetHeaders(v map[string]*string) *CreateClusterRequest {
	s.Headers = v
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

type DescribeClusterDetailRequest struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
}

func (s DescribeClusterDetailRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailRequest) SetHeaders(v map[string]*string) *DescribeClusterDetailRequest {
	s.Headers = v
	return s
}

type DescribeClusterDetailResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClusterDetailResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterDetailResponse) SetHeaders(v map[string]*string) *DescribeClusterDetailResponse {
	s.Headers = v
	return s
}

type DescribeClustersQuery struct {
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ClusterType *string `json:"clusterType,omitempty" xml:"clusterType,omitempty"`
}

func (s DescribeClustersQuery) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersQuery) GoString() string {
	return s.String()
}

func (s *DescribeClustersQuery) SetName(v string) *DescribeClustersQuery {
	s.Name = &v
	return s
}

func (s *DescribeClustersQuery) SetClusterType(v string) *DescribeClustersQuery {
	s.ClusterType = &v
	return s
}

type DescribeClustersRequest struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	Query   *DescribeClustersQuery `json:"query,omitempty" xml:"query,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) SetHeaders(v map[string]*string) *DescribeClustersRequest {
	s.Headers = v
	return s
}

func (s *DescribeClustersRequest) SetQuery(v *DescribeClustersQuery) *DescribeClustersRequest {
	s.Query = v
	return s
}

type DescribeClustersResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DescribeClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClustersResponse) GoString() string {
	return s.String()
}

func (s *DescribeClustersResponse) SetHeaders(v map[string]*string) *DescribeClustersResponse {
	s.Headers = v
	return s
}

type Client struct {
	roa.Client
}

func NewClient(config *roa.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *roa.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-2-pop":          tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-1":        tea.String("cs.aliyuncs.com"),
		"cn-beijing-finance-pop":      tea.String("cs.aliyuncs.com"),
		"cn-beijing-gov-1":            tea.String("cs.aliyuncs.com"),
		"cn-beijing-nu16-b01":         tea.String("cs.aliyuncs.com"),
		"cn-edge-1":                   tea.String("cs.aliyuncs.com"),
		"cn-fujian":                   tea.String("cs.aliyuncs.com"),
		"cn-haidian-cm12-c01":         tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-bj-b01":          tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-finance":         tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-prod-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-2": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-internal-test-3": tea.String("cs.aliyuncs.com"),
		"cn-hangzhou-test-306":        tea.String("cs.aliyuncs.com"),
		"cn-hongkong-finance-pop":     tea.String("cs.aliyuncs.com"),
		"cn-huhehaote-nebula-1":       tea.String("cs.aliyuncs.com"),
		"cn-qingdao-nebula":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et15-b01":        tea.String("cs.aliyuncs.com"),
		"cn-shanghai-et2-b01":         tea.String("cs.aliyuncs.com"),
		"cn-shanghai-finance-1":       tea.String("cs.aliyuncs.com"),
		"cn-shanghai-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shanghai-internal-test-1": tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-finance-1":       tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-inner":           tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-st4-d01":         tea.String("cs.aliyuncs.com"),
		"cn-shenzhen-su18-b01":        tea.String("cs.aliyuncs.com"),
		"cn-wuhan":                    tea.String("cs.aliyuncs.com"),
		"cn-wulanchabu":               tea.String("cs.aliyuncs.com"),
		"cn-yushanfang":               tea.String("cs.aliyuncs.com"),
		"cn-zhangbei-na61-b01":        tea.String("cs.aliyuncs.com"),
		"cn-zhangjiakou-na62-a01":     tea.String("cs.aliyuncs.com"),
		"cn-zhengzhou-nebula-1":       tea.String("cs.aliyuncs.com"),
		"eu-west-1-oxs":               tea.String("cs.aliyuncs.com"),
		"rus-west-1-pop":              tea.String("cs.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.EndpointHost, _err = client.GetEndpoint(tea.String("cs"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.EndpointHost)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetProjectEventsWithOptions(clusterId *string, projectId *string, request *GetProjectEventsRequest, runtime *util.RuntimeOptions) (_result *GetProjectEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetProjectEventsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetProjectEvents"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/projects/"+tea.StringValue(projectId)+"/events"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetProjectEvents(clusterId *string, projectId *string, request *GetProjectEventsRequest) (_result *GetProjectEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetProjectEventsResponse{}
	_body, _err := client.GetProjectEventsWithOptions(clusterId, projectId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKubernetesTemplateWithOptions(clusterId *string, request *DescribeKubernetesTemplateRequest, runtime *util.RuntimeOptions) (_result *DescribeKubernetesTemplateResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKubernetesTemplateResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeKubernetesTemplate"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/templates/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKubernetesTemplate(clusterId *string, request *DescribeKubernetesTemplateRequest) (_result *DescribeKubernetesTemplateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKubernetesTemplateResponse{}
	_body, _err := client.DescribeKubernetesTemplateWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAgilityTunnelCertsWithOptions(token *string, request *DescribeAgilityTunnelCertsRequest, runtime *util.RuntimeOptions) (_result *DescribeAgilityTunnelCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAgilityTunnelCertsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeAgilityTunnelCerts"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("Anonymous"), tea.String("/agility/"+tea.StringValue(token)+"/agent_certs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAgilityTunnelCerts(token *string, request *DescribeAgilityTunnelCertsRequest) (_result *DescribeAgilityTunnelCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAgilityTunnelCertsResponse{}
	_body, _err := client.DescribeAgilityTunnelCertsWithOptions(token, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterTokensWithOptions(clusterId *string, request *DescribeClusterTokensRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterTokensResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterTokensResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterTokens"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/tokens"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterTokens(clusterId *string, request *DescribeClusterTokensRequest) (_result *DescribeClusterTokensResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterTokensResponse{}
	_body, _err := client.DescribeClusterTokensWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DownloadClusterNodeCertsWithOptions(token *string, nodeId *string, request *DownloadClusterNodeCertsRequest, runtime *util.RuntimeOptions) (_result *DownloadClusterNodeCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DownloadClusterNodeCertsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DownloadClusterNodeCerts"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("Anonymous"), tea.String("/token/"+tea.StringValue(token)+"/nodes/"+tea.StringValue(nodeId)+"/certs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DownloadClusterNodeCerts(token *string, nodeId *string, request *DownloadClusterNodeCertsRequest) (_result *DownloadClusterNodeCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DownloadClusterNodeCertsResponse{}
	_body, _err := client.DownloadClusterNodeCertsWithOptions(token, nodeId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceContainersWithOptions(clusterId *string, serviceId *string, request *DescribeServiceContainersRequest, runtime *util.RuntimeOptions) (_result *DescribeServiceContainersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeServiceContainersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeServiceContainers"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/services/"+tea.StringValue(serviceId)+"/containers"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceContainers(clusterId *string, serviceId *string, request *DescribeServiceContainersRequest) (_result *DescribeServiceContainersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeServiceContainersResponse{}
	_body, _err := client.DescribeServiceContainersWithOptions(clusterId, serviceId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GatherLogsTokenWithOptions(token *string, request *GatherLogsTokenRequest, runtime *util.RuntimeOptions) (_result *GatherLogsTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GatherLogsTokenResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GatherLogsToken"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("Anonymous"), tea.String("/token/"+tea.StringValue(token)+"/gather_logs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GatherLogsToken(token *string, request *GatherLogsTokenRequest) (_result *GatherLogsTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GatherLogsTokenResponse{}
	_body, _err := client.GatherLogsTokenWithOptions(token, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetClusterProjectsWithOptions(clusterId *string, request *GetClusterProjectsRequest, runtime *util.RuntimeOptions) (_result *GetClusterProjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetClusterProjectsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetClusterProjects"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/projects"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetClusterProjects(clusterId *string, request *GetClusterProjectsRequest) (_result *GetClusterProjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetClusterProjectsResponse{}
	_body, _err := client.GetClusterProjectsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeApiVersionWithOptions(request *DescribeApiVersionRequest, runtime *util.RuntimeOptions) (_result *DescribeApiVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeApiVersionResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeApiVersion"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/version"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeApiVersion(request *DescribeApiVersionRequest) (_result *DescribeApiVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeApiVersionResponse{}
	_body, _err := client.DescribeApiVersionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTaskInfoWithOptions(taskId *string, request *DescribeTaskInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTaskInfo"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/tasks/"+tea.StringValue(taskId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTaskInfo(taskId *string, request *DescribeTaskInfoRequest) (_result *DescribeTaskInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskInfoResponse{}
	_body, _err := client.DescribeTaskInfoWithOptions(taskId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodesWithOptions(clusterId *string, request *DescribeClusterNodesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodes"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/nodes"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodes(clusterId *string, request *DescribeClusterNodesRequest) (_result *DescribeClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodesResponse{}
	_body, _err := client.DescribeClusterNodesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeAgilityTunnelAgentInfoWithOptions(token *string, request *DescribeAgilityTunnelAgentInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeAgilityTunnelAgentInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeAgilityTunnelAgentInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeAgilityTunnelAgentInfo"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("Anonymous"), tea.String("/agility/"+tea.StringValue(token)+"/agent_info"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeAgilityTunnelAgentInfo(token *string, request *DescribeAgilityTunnelAgentInfoRequest) (_result *DescribeAgilityTunnelAgentInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeAgilityTunnelAgentInfoResponse{}
	_body, _err := client.DescribeAgilityTunnelAgentInfoWithOptions(token, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterNodeWithOptions(clusterId *string, ip *string, request *DeleteClusterNodeRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterNodeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteClusterNode"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/ip/"+tea.StringValue(ip)), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteClusterNode(clusterId *string, ip *string, request *DeleteClusterNodeRequest) (_result *DeleteClusterNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterNodeResponse{}
	_body, _err := client.DeleteClusterNodeWithOptions(clusterId, ip, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodeInfoWithInstanceWithOptions(token *string, instanceId *string, request *DescribeClusterNodeInfoWithInstanceRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodeInfoWithInstanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodeInfoWithInstanceResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodeInfoWithInstance"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("Anonymous"), tea.String("/token/"+tea.StringValue(token)+"/instance/"+tea.StringValue(instanceId)+"/node_info"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodeInfoWithInstance(token *string, instanceId *string, request *DescribeClusterNodeInfoWithInstanceRequest) (_result *DescribeClusterNodeInfoWithInstanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodeInfoWithInstanceResponse{}
	_body, _err := client.DescribeClusterNodeInfoWithInstanceWithOptions(token, instanceId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeUserContainersWithOptions(regionId *string, request *DescribeUserContainersRequest, runtime *util.RuntimeOptions) (_result *DescribeUserContainersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeUserContainersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeUserContainers"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/region/"+tea.StringValue(regionId)+"/containers"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeUserContainers(regionId *string, request *DescribeUserContainersRequest) (_result *DescribeUserContainersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeUserContainersResponse{}
	_body, _err := client.DescribeUserContainersWithOptions(regionId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterTokenWithOptions(clusterId *string, request *CreateClusterTokenRequest, runtime *util.RuntimeOptions) (_result *CreateClusterTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateClusterTokenResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateClusterToken"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/token"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateClusterToken(clusterId *string, request *CreateClusterTokenRequest) (_result *CreateClusterTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterTokenResponse{}
	_body, _err := client.CreateClusterTokenWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterHostsWithOptions(clusterId *string, request *DescribeClusterHostsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterHostsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterHostsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterHosts"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/hosts"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterHosts(clusterId *string, request *DescribeClusterHostsRequest) (_result *DescribeClusterHostsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterHostsResponse{}
	_body, _err := client.DescribeClusterHostsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeKubernetesTemplatesWithOptions(request *DescribeKubernetesTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeKubernetesTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeKubernetesTemplatesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeKubernetesTemplates"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/k8s/templates"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeKubernetesTemplates(request *DescribeKubernetesTemplatesRequest) (_result *DescribeKubernetesTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeKubernetesTemplatesResponse{}
	_body, _err := client.DescribeKubernetesTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplatesWithOptions(request *DescribeTemplatesRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplatesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTemplates"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/templates"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplates(request *DescribeTemplatesRequest) (_result *DescribeTemplatesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplatesResponse{}
	_body, _err := client.DescribeTemplatesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterScaledNodeWithOptions(clusterId *string, request *DescribeClusterScaledNodeRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterScaledNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterScaledNodeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterScaledNode"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/scaled_nodes/"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterScaledNode(clusterId *string, request *DescribeClusterScaledNodeRequest) (_result *DescribeClusterScaledNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterScaledNodeResponse{}
	_body, _err := client.DescribeClusterScaledNodeWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CallbackClusterTokenWithOptions(token *string, reqOnce *string, request *CallbackClusterTokenRequest, runtime *util.RuntimeOptions) (_result *CallbackClusterTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CallbackClusterTokenResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CallbackClusterToken"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("Anonymous"), tea.String("/token/"+tea.StringValue(token)+"/req_once/"+tea.StringValue(reqOnce)+"/callback"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CallbackClusterToken(token *string, reqOnce *string, request *CallbackClusterTokenRequest) (_result *CallbackClusterTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CallbackClusterTokenResponse{}
	_body, _err := client.CallbackClusterTokenWithOptions(token, reqOnce, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeImagesWithOptions(request *DescribeImagesRequest, runtime *util.RuntimeOptions) (_result *DescribeImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeImages"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/images"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeImages(request *DescribeImagesRequest) (_result *DescribeImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeImagesResponse{}
	_body, _err := client.DescribeImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterLogsWithOptions(clusterId *string, request *DescribeClusterLogsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterLogs"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/logs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterLogs(clusterId *string, request *DescribeClusterLogsRequest) (_result *DescribeClusterLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterLogsResponse{}
	_body, _err := client.DescribeClusterLogsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterServicesWithOptions(clusterId *string, request *DescribeClusterServicesRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterServicesResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterServices"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/services"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterServices(clusterId *string, request *DescribeClusterServicesRequest) (_result *DescribeClusterServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterServicesResponse{}
	_body, _err := client.DescribeClusterServicesWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTriggerHookWithOptions(clusterId *string, projectId *string, request *GetTriggerHookRequest, runtime *util.RuntimeOptions) (_result *GetTriggerHookResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetTriggerHookResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("GetTriggerHook"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/hook/trigger/"+tea.StringValue(clusterId)+"/"+tea.StringValue(projectId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTriggerHook(clusterId *string, projectId *string, request *GetTriggerHookRequest) (_result *GetTriggerHookResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetTriggerHookResponse{}
	_body, _err := client.GetTriggerHookWithOptions(clusterId, projectId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeTemplateAttributeWithOptions(templateId *string, request *DescribeTemplateAttributeRequest, runtime *util.RuntimeOptions) (_result *DescribeTemplateAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeTemplateAttribute"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/templates/"+tea.StringValue(templateId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeTemplateAttribute(templateId *string, request *DescribeTemplateAttributeRequest) (_result *DescribeTemplateAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTemplateAttributeResponse{}
	_body, _err := client.DescribeTemplateAttributeWithOptions(templateId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterCertsWithOptions(clusterId *string, request *DescribeClusterCertsRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterCertsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterCertsResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterCerts"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)+"/certs"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterCerts(clusterId *string, request *DescribeClusterCertsRequest) (_result *DescribeClusterCertsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterCertsResponse{}
	_body, _err := client.DescribeClusterCertsWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterNodeInfoWithOptions(token *string, request *DescribeClusterNodeInfoRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterNodeInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterNodeInfoResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterNodeInfo"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("Anonymous"), tea.String("/token/"+tea.StringValue(token)+"/node_info"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterNodeInfo(token *string, request *DescribeClusterNodeInfoRequest) (_result *DescribeClusterNodeInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterNodeInfoResponse{}
	_body, _err := client.DescribeClusterNodeInfoWithOptions(token, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteClusterWithOptions(clusterId *string, request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DeleteCluster"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCluster(clusterId *string, request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateClusterWithOptions(request *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("CreateCluster"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/clusters"), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClusterDetailWithOptions(clusterId *string, request *DescribeClusterDetailRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterDetailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusterDetail"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters/"+tea.StringValue(clusterId)), nil, request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusterDetail(clusterId *string, request *DescribeClusterDetailRequest) (_result *DescribeClusterDetailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterDetailResponse{}
	_body, _err := client.DescribeClusterDetailWithOptions(clusterId, request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeClustersWithOptions(request *DescribeClustersRequest, runtime *util.RuntimeOptions) (_result *DescribeClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DoRequestWithAction(tea.String("DescribeClusters"), tea.String("2018-04-18"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/clusters"), util.StringifyMapValue(tea.ToMap(request.Query)), request.Headers, nil, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeClusters(request *DescribeClustersRequest) (_result *DescribeClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClustersResponse{}
	_body, _err := client.DescribeClustersWithOptions(request, runtime)
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
