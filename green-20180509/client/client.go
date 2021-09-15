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

type AddFacesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFacesRequest) GoString() string {
	return s.String()
}

func (s *AddFacesRequest) SetRegionId(v string) *AddFacesRequest {
	s.RegionId = &v
	return s
}

func (s *AddFacesRequest) SetClientInfo(v string) *AddFacesRequest {
	s.ClientInfo = &v
	return s
}

type AddFacesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddFacesResponse) GoString() string {
	return s.String()
}

func (s *AddFacesResponse) SetHeaders(v map[string]*string) *AddFacesResponse {
	s.Headers = v
	return s
}

type AddGroupsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddGroupsRequest) SetRegionId(v string) *AddGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *AddGroupsRequest) SetClientInfo(v string) *AddGroupsRequest {
	s.ClientInfo = &v
	return s
}

type AddGroupsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupsResponse) GoString() string {
	return s.String()
}

func (s *AddGroupsResponse) SetHeaders(v map[string]*string) *AddGroupsResponse {
	s.Headers = v
	return s
}

type AddPersonRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPersonRequest) GoString() string {
	return s.String()
}

func (s *AddPersonRequest) SetRegionId(v string) *AddPersonRequest {
	s.RegionId = &v
	return s
}

func (s *AddPersonRequest) SetClientInfo(v string) *AddPersonRequest {
	s.ClientInfo = &v
	return s
}

type AddPersonResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s AddPersonResponse) GoString() string {
	return s.String()
}

func (s *AddPersonResponse) SetHeaders(v map[string]*string) *AddPersonResponse {
	s.Headers = v
	return s
}

type AddSimilarityImageRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddSimilarityImageRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityImageRequest) GoString() string {
	return s.String()
}

func (s *AddSimilarityImageRequest) SetClientInfo(v string) *AddSimilarityImageRequest {
	s.ClientInfo = &v
	return s
}

type AddSimilarityImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddSimilarityImageResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityImageResponse) GoString() string {
	return s.String()
}

func (s *AddSimilarityImageResponse) SetHeaders(v map[string]*string) *AddSimilarityImageResponse {
	s.Headers = v
	return s
}

type AddSimilarityLibraryRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddSimilarityLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityLibraryRequest) GoString() string {
	return s.String()
}

func (s *AddSimilarityLibraryRequest) SetClientInfo(v string) *AddSimilarityLibraryRequest {
	s.ClientInfo = &v
	return s
}

type AddSimilarityLibraryResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddSimilarityLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityLibraryResponse) GoString() string {
	return s.String()
}

func (s *AddSimilarityLibraryResponse) SetHeaders(v map[string]*string) *AddSimilarityLibraryResponse {
	s.Headers = v
	return s
}

type AddVideoDnaRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddVideoDnaRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaRequest) GoString() string {
	return s.String()
}

func (s *AddVideoDnaRequest) SetClientInfo(v string) *AddVideoDnaRequest {
	s.ClientInfo = &v
	return s
}

type AddVideoDnaResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddVideoDnaResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaResponse) GoString() string {
	return s.String()
}

func (s *AddVideoDnaResponse) SetHeaders(v map[string]*string) *AddVideoDnaResponse {
	s.Headers = v
	return s
}

type AddVideoDnaGroupRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s AddVideoDnaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaGroupRequest) GoString() string {
	return s.String()
}

func (s *AddVideoDnaGroupRequest) SetClientInfo(v string) *AddVideoDnaGroupRequest {
	s.ClientInfo = &v
	return s
}

type AddVideoDnaGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s AddVideoDnaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaGroupResponse) GoString() string {
	return s.String()
}

func (s *AddVideoDnaGroupResponse) SetHeaders(v map[string]*string) *AddVideoDnaGroupResponse {
	s.Headers = v
	return s
}

type DeleteFacesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFacesRequest) SetRegionId(v string) *DeleteFacesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteFacesRequest) SetClientInfo(v string) *DeleteFacesRequest {
	s.ClientInfo = &v
	return s
}

type DeleteFacesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFacesResponse) SetHeaders(v map[string]*string) *DeleteFacesResponse {
	s.Headers = v
	return s
}

type DeleteGroupsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupsRequest) SetRegionId(v string) *DeleteGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGroupsRequest) SetClientInfo(v string) *DeleteGroupsRequest {
	s.ClientInfo = &v
	return s
}

type DeleteGroupsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupsResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupsResponse) SetHeaders(v map[string]*string) *DeleteGroupsResponse {
	s.Headers = v
	return s
}

type DeletePersonRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeletePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePersonRequest) GoString() string {
	return s.String()
}

func (s *DeletePersonRequest) SetRegionId(v string) *DeletePersonRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePersonRequest) SetClientInfo(v string) *DeletePersonRequest {
	s.ClientInfo = &v
	return s
}

type DeletePersonResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeletePersonResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePersonResponse) GoString() string {
	return s.String()
}

func (s *DeletePersonResponse) SetHeaders(v map[string]*string) *DeletePersonResponse {
	s.Headers = v
	return s
}

type DeleteSimilarityImageRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteSimilarityImageRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityImageRequest) SetClientInfo(v string) *DeleteSimilarityImageRequest {
	s.ClientInfo = &v
	return s
}

type DeleteSimilarityImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSimilarityImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityImageResponse) SetHeaders(v map[string]*string) *DeleteSimilarityImageResponse {
	s.Headers = v
	return s
}

type DeleteSimilarityLibraryRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteSimilarityLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityLibraryRequest) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityLibraryRequest) SetClientInfo(v string) *DeleteSimilarityLibraryRequest {
	s.ClientInfo = &v
	return s
}

type DeleteSimilarityLibraryResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteSimilarityLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityLibraryResponse) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityLibraryResponse) SetHeaders(v map[string]*string) *DeleteSimilarityLibraryResponse {
	s.Headers = v
	return s
}

type DeleteVideoDnaRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteVideoDnaRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaRequest) SetClientInfo(v string) *DeleteVideoDnaRequest {
	s.ClientInfo = &v
	return s
}

type DeleteVideoDnaResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteVideoDnaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaResponse) SetHeaders(v map[string]*string) *DeleteVideoDnaResponse {
	s.Headers = v
	return s
}

type DeleteVideoDnaGroupRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DeleteVideoDnaGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaGroupRequest) SetClientInfo(v string) *DeleteVideoDnaGroupRequest {
	s.ClientInfo = &v
	return s
}

type DeleteVideoDnaGroupResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DeleteVideoDnaGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaGroupResponse) SetHeaders(v map[string]*string) *DeleteVideoDnaGroupResponse {
	s.Headers = v
	return s
}

type DetectFaceRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s DetectFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceRequest) GoString() string {
	return s.String()
}

func (s *DetectFaceRequest) SetClientInfo(v string) *DetectFaceRequest {
	s.ClientInfo = &v
	return s
}

type DetectFaceResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s DetectFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponse) GoString() string {
	return s.String()
}

func (s *DetectFaceResponse) SetHeaders(v map[string]*string) *DetectFaceResponse {
	s.Headers = v
	return s
}

type FileAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s FileAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *FileAsyncScanRequest) SetClientInfo(v string) *FileAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type FileAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s FileAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *FileAsyncScanResponse) SetHeaders(v map[string]*string) *FileAsyncScanResponse {
	s.Headers = v
	return s
}

type FileAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s FileAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *FileAsyncScanResultsRequest) SetClientInfo(v string) *FileAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type FileAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s FileAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *FileAsyncScanResultsResponse) SetHeaders(v map[string]*string) *FileAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type GetAddVideoDnaResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetAddVideoDnaResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAddVideoDnaResultsRequest) GoString() string {
	return s.String()
}

func (s *GetAddVideoDnaResultsRequest) SetClientInfo(v string) *GetAddVideoDnaResultsRequest {
	s.ClientInfo = &v
	return s
}

type GetAddVideoDnaResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetAddVideoDnaResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAddVideoDnaResultsResponse) GoString() string {
	return s.String()
}

func (s *GetAddVideoDnaResultsResponse) SetHeaders(v map[string]*string) *GetAddVideoDnaResultsResponse {
	s.Headers = v
	return s
}

type GetFacesRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFacesRequest) GoString() string {
	return s.String()
}

func (s *GetFacesRequest) SetRegionId(v string) *GetFacesRequest {
	s.RegionId = &v
	return s
}

func (s *GetFacesRequest) SetClientInfo(v string) *GetFacesRequest {
	s.ClientInfo = &v
	return s
}

type GetFacesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFacesResponseBody) GoString() string {
	return s.String()
}

func (s *GetFacesResponseBody) SetRequestId(v string) *GetFacesResponseBody {
	s.RequestId = &v
	return s
}

type GetFacesResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFacesResponse) GoString() string {
	return s.String()
}

func (s *GetFacesResponse) SetHeaders(v map[string]*string) *GetFacesResponse {
	s.Headers = v
	return s
}

func (s *GetFacesResponse) SetBody(v *GetFacesResponseBody) *GetFacesResponse {
	s.Body = v
	return s
}

type GetGroupsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetGroupsRequest) SetRegionId(v string) *GetGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *GetGroupsRequest) SetClientInfo(v string) *GetGroupsRequest {
	s.ClientInfo = &v
	return s
}

type GetGroupsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupsResponse) GoString() string {
	return s.String()
}

func (s *GetGroupsResponse) SetHeaders(v map[string]*string) *GetGroupsResponse {
	s.Headers = v
	return s
}

type GetPersonRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonRequest) GoString() string {
	return s.String()
}

func (s *GetPersonRequest) SetRegionId(v string) *GetPersonRequest {
	s.RegionId = &v
	return s
}

func (s *GetPersonRequest) SetClientInfo(v string) *GetPersonRequest {
	s.ClientInfo = &v
	return s
}

type GetPersonResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonResponse) GoString() string {
	return s.String()
}

func (s *GetPersonResponse) SetHeaders(v map[string]*string) *GetPersonResponse {
	s.Headers = v
	return s
}

type GetPersonsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetPersonsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonsRequest) GoString() string {
	return s.String()
}

func (s *GetPersonsRequest) SetRegionId(v string) *GetPersonsRequest {
	s.RegionId = &v
	return s
}

func (s *GetPersonsRequest) SetClientInfo(v string) *GetPersonsRequest {
	s.ClientInfo = &v
	return s
}

type GetPersonsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetPersonsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPersonsResponse) GoString() string {
	return s.String()
}

func (s *GetPersonsResponse) SetHeaders(v map[string]*string) *GetPersonsResponse {
	s.Headers = v
	return s
}

type GetSimilarityImageRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetSimilarityImageRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityImageRequest) GoString() string {
	return s.String()
}

func (s *GetSimilarityImageRequest) SetClientInfo(v string) *GetSimilarityImageRequest {
	s.ClientInfo = &v
	return s
}

type GetSimilarityImageResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetSimilarityImageResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityImageResponse) GoString() string {
	return s.String()
}

func (s *GetSimilarityImageResponse) SetHeaders(v map[string]*string) *GetSimilarityImageResponse {
	s.Headers = v
	return s
}

type GetSimilarityLibraryRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s GetSimilarityLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetSimilarityLibraryRequest) SetClientInfo(v string) *GetSimilarityLibraryRequest {
	s.ClientInfo = &v
	return s
}

type GetSimilarityLibraryResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s GetSimilarityLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetSimilarityLibraryResponse) SetHeaders(v map[string]*string) *GetSimilarityLibraryResponse {
	s.Headers = v
	return s
}

type ImageAsyncManualScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageAsyncManualScanRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanRequest) SetClientInfo(v string) *ImageAsyncManualScanRequest {
	s.ClientInfo = &v
	return s
}

type ImageAsyncManualScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageAsyncManualScanResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanResponse) SetHeaders(v map[string]*string) *ImageAsyncManualScanResponse {
	s.Headers = v
	return s
}

type ImageAsyncManualScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageAsyncManualScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanResultsRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanResultsRequest) SetClientInfo(v string) *ImageAsyncManualScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type ImageAsyncManualScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageAsyncManualScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanResultsResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanResultsResponse) SetHeaders(v map[string]*string) *ImageAsyncManualScanResultsResponse {
	s.Headers = v
	return s
}

type ImageAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanRequest) SetClientInfo(v string) *ImageAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type ImageAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanResponse) SetHeaders(v map[string]*string) *ImageAsyncScanResponse {
	s.Headers = v
	return s
}

type ImageAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanResultsRequest) SetClientInfo(v string) *ImageAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type ImageAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanResultsResponse) SetHeaders(v map[string]*string) *ImageAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type ImageScanFeedbackRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageScanFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageScanFeedbackRequest) GoString() string {
	return s.String()
}

func (s *ImageScanFeedbackRequest) SetClientInfo(v string) *ImageScanFeedbackRequest {
	s.ClientInfo = &v
	return s
}

type ImageScanFeedbackResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageScanFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageScanFeedbackResponse) GoString() string {
	return s.String()
}

func (s *ImageScanFeedbackResponse) SetHeaders(v map[string]*string) *ImageScanFeedbackResponse {
	s.Headers = v
	return s
}

type ImageSyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ImageSyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s ImageSyncScanRequest) GoString() string {
	return s.String()
}

func (s *ImageSyncScanRequest) SetClientInfo(v string) *ImageSyncScanRequest {
	s.ClientInfo = &v
	return s
}

type ImageSyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ImageSyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s ImageSyncScanResponse) GoString() string {
	return s.String()
}

func (s *ImageSyncScanResponse) SetHeaders(v map[string]*string) *ImageSyncScanResponse {
	s.Headers = v
	return s
}

type ListSimilarityImagesRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ListSimilarityImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityImagesRequest) GoString() string {
	return s.String()
}

func (s *ListSimilarityImagesRequest) SetClientInfo(v string) *ListSimilarityImagesRequest {
	s.ClientInfo = &v
	return s
}

type ListSimilarityImagesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListSimilarityImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityImagesResponse) GoString() string {
	return s.String()
}

func (s *ListSimilarityImagesResponse) SetHeaders(v map[string]*string) *ListSimilarityImagesResponse {
	s.Headers = v
	return s
}

type ListSimilarityLibrariesRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s ListSimilarityLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityLibrariesRequest) GoString() string {
	return s.String()
}

func (s *ListSimilarityLibrariesRequest) SetClientInfo(v string) *ListSimilarityLibrariesRequest {
	s.ClientInfo = &v
	return s
}

type ListSimilarityLibrariesResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ListSimilarityLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityLibrariesResponse) GoString() string {
	return s.String()
}

func (s *ListSimilarityLibrariesResponse) SetHeaders(v map[string]*string) *ListSimilarityLibrariesResponse {
	s.Headers = v
	return s
}

type LiveStreamAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s LiveStreamAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanRequest) SetClientInfo(v string) *LiveStreamAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type LiveStreamAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s LiveStreamAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanResponse) SetHeaders(v map[string]*string) *LiveStreamAsyncScanResponse {
	s.Headers = v
	return s
}

type LiveStreamAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s LiveStreamAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanResultsRequest) SetClientInfo(v string) *LiveStreamAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type LiveStreamAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s LiveStreamAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanResultsResponse) SetHeaders(v map[string]*string) *LiveStreamAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type LiveStreamCancelScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s LiveStreamCancelScanRequest) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamCancelScanRequest) GoString() string {
	return s.String()
}

func (s *LiveStreamCancelScanRequest) SetClientInfo(v string) *LiveStreamCancelScanRequest {
	s.ClientInfo = &v
	return s
}

type LiveStreamCancelScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s LiveStreamCancelScanResponse) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamCancelScanResponse) GoString() string {
	return s.String()
}

func (s *LiveStreamCancelScanResponse) SetHeaders(v map[string]*string) *LiveStreamCancelScanResponse {
	s.Headers = v
	return s
}

type PostAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s PostAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s PostAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *PostAsyncScanRequest) SetClientInfo(v string) *PostAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type PostAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PostAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s PostAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *PostAsyncScanResponse) SetHeaders(v map[string]*string) *PostAsyncScanResponse {
	s.Headers = v
	return s
}

type PostAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s PostAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s PostAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *PostAsyncScanResultsRequest) SetClientInfo(v string) *PostAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type PostAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s PostAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s PostAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *PostAsyncScanResultsResponse) SetHeaders(v map[string]*string) *PostAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type SearchPersonRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s SearchPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPersonRequest) GoString() string {
	return s.String()
}

func (s *SearchPersonRequest) SetRegionId(v string) *SearchPersonRequest {
	s.RegionId = &v
	return s
}

func (s *SearchPersonRequest) SetClientInfo(v string) *SearchPersonRequest {
	s.ClientInfo = &v
	return s
}

type SearchPersonResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SearchPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPersonResponse) GoString() string {
	return s.String()
}

func (s *SearchPersonResponse) SetHeaders(v map[string]*string) *SearchPersonResponse {
	s.Headers = v
	return s
}

type SetPersonRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s SetPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPersonRequest) GoString() string {
	return s.String()
}

func (s *SetPersonRequest) SetRegionId(v string) *SetPersonRequest {
	s.RegionId = &v
	return s
}

func (s *SetPersonRequest) SetClientInfo(v string) *SetPersonRequest {
	s.ClientInfo = &v
	return s
}

type SetPersonResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SetPersonResponse) String() string {
	return tea.Prettify(s)
}

func (s SetPersonResponse) GoString() string {
	return s.String()
}

func (s *SetPersonResponse) SetHeaders(v map[string]*string) *SetPersonResponse {
	s.Headers = v
	return s
}

type TextAsyncManualScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s TextAsyncManualScanRequest) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanRequest) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanRequest) SetClientInfo(v string) *TextAsyncManualScanRequest {
	s.ClientInfo = &v
	return s
}

type TextAsyncManualScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TextAsyncManualScanResponse) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanResponse) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanResponse) SetHeaders(v map[string]*string) *TextAsyncManualScanResponse {
	s.Headers = v
	return s
}

type TextAsyncManualScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s TextAsyncManualScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanResultsRequest) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanResultsRequest) SetClientInfo(v string) *TextAsyncManualScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type TextAsyncManualScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TextAsyncManualScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanResultsResponse) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanResultsResponse) SetHeaders(v map[string]*string) *TextAsyncManualScanResultsResponse {
	s.Headers = v
	return s
}

type TextFeedbackRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s TextFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s TextFeedbackRequest) GoString() string {
	return s.String()
}

func (s *TextFeedbackRequest) SetClientInfo(v string) *TextFeedbackRequest {
	s.ClientInfo = &v
	return s
}

type TextFeedbackResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TextFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s TextFeedbackResponse) GoString() string {
	return s.String()
}

func (s *TextFeedbackResponse) SetHeaders(v map[string]*string) *TextFeedbackResponse {
	s.Headers = v
	return s
}

type TextScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s TextScanRequest) String() string {
	return tea.Prettify(s)
}

func (s TextScanRequest) GoString() string {
	return s.String()
}

func (s *TextScanRequest) SetClientInfo(v string) *TextScanRequest {
	s.ClientInfo = &v
	return s
}

type TextScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TextScanResponse) String() string {
	return tea.Prettify(s)
}

func (s TextScanResponse) GoString() string {
	return s.String()
}

func (s *TextScanResponse) SetHeaders(v map[string]*string) *TextScanResponse {
	s.Headers = v
	return s
}

type UploadCredentialsRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s UploadCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCredentialsRequest) GoString() string {
	return s.String()
}

func (s *UploadCredentialsRequest) SetRegionId(v string) *UploadCredentialsRequest {
	s.RegionId = &v
	return s
}

func (s *UploadCredentialsRequest) SetClientInfo(v string) *UploadCredentialsRequest {
	s.ClientInfo = &v
	return s
}

type UploadCredentialsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s UploadCredentialsResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCredentialsResponse) GoString() string {
	return s.String()
}

func (s *UploadCredentialsResponse) SetHeaders(v map[string]*string) *UploadCredentialsResponse {
	s.Headers = v
	return s
}

type VideoAsyncManualScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoAsyncManualScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanRequest) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanRequest) SetClientInfo(v string) *VideoAsyncManualScanRequest {
	s.ClientInfo = &v
	return s
}

type VideoAsyncManualScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoAsyncManualScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanResponse) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanResponse) SetHeaders(v map[string]*string) *VideoAsyncManualScanResponse {
	s.Headers = v
	return s
}

type VideoAsyncManualScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoAsyncManualScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanResultsRequest) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanResultsRequest) SetClientInfo(v string) *VideoAsyncManualScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type VideoAsyncManualScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoAsyncManualScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanResultsResponse) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanResultsResponse) SetHeaders(v map[string]*string) *VideoAsyncManualScanResultsResponse {
	s.Headers = v
	return s
}

type VideoAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanRequest) SetClientInfo(v string) *VideoAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type VideoAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanResponse) SetHeaders(v map[string]*string) *VideoAsyncScanResponse {
	s.Headers = v
	return s
}

type VideoAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanResultsRequest) SetClientInfo(v string) *VideoAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type VideoAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanResultsResponse) SetHeaders(v map[string]*string) *VideoAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type VideoCancelScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoCancelScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoCancelScanRequest) GoString() string {
	return s.String()
}

func (s *VideoCancelScanRequest) SetClientInfo(v string) *VideoCancelScanRequest {
	s.ClientInfo = &v
	return s
}

type VideoCancelScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoCancelScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoCancelScanResponse) GoString() string {
	return s.String()
}

func (s *VideoCancelScanResponse) SetHeaders(v map[string]*string) *VideoCancelScanResponse {
	s.Headers = v
	return s
}

type VideoFeedbackRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoFeedbackRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoFeedbackRequest) GoString() string {
	return s.String()
}

func (s *VideoFeedbackRequest) SetClientInfo(v string) *VideoFeedbackRequest {
	s.ClientInfo = &v
	return s
}

type VideoFeedbackResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoFeedbackResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoFeedbackResponse) GoString() string {
	return s.String()
}

func (s *VideoFeedbackResponse) SetHeaders(v map[string]*string) *VideoFeedbackResponse {
	s.Headers = v
	return s
}

type VideoSyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VideoSyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VideoSyncScanRequest) GoString() string {
	return s.String()
}

func (s *VideoSyncScanRequest) SetClientInfo(v string) *VideoSyncScanRequest {
	s.ClientInfo = &v
	return s
}

type VideoSyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VideoSyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VideoSyncScanResponse) GoString() string {
	return s.String()
}

func (s *VideoSyncScanResponse) SetHeaders(v map[string]*string) *VideoSyncScanResponse {
	s.Headers = v
	return s
}

type VodAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VodAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *VodAsyncScanRequest) SetClientInfo(v string) *VodAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type VodAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VodAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *VodAsyncScanResponse) SetHeaders(v map[string]*string) *VodAsyncScanResponse {
	s.Headers = v
	return s
}

type VodAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VodAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *VodAsyncScanResultsRequest) SetClientInfo(v string) *VodAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type VodAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VodAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *VodAsyncScanResultsResponse) SetHeaders(v map[string]*string) *VodAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type VoiceAsyncManualScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceAsyncManualScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanRequest) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanRequest) SetClientInfo(v string) *VoiceAsyncManualScanRequest {
	s.ClientInfo = &v
	return s
}

type VoiceAsyncManualScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceAsyncManualScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanResponse) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanResponse) SetHeaders(v map[string]*string) *VoiceAsyncManualScanResponse {
	s.Headers = v
	return s
}

type VoiceAsyncManualScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceAsyncManualScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanResultsRequest) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanResultsRequest) SetClientInfo(v string) *VoiceAsyncManualScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type VoiceAsyncManualScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceAsyncManualScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanResultsResponse) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanResultsResponse) SetHeaders(v map[string]*string) *VoiceAsyncManualScanResultsResponse {
	s.Headers = v
	return s
}

type VoiceAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanRequest) SetClientInfo(v string) *VoiceAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type VoiceAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanResponse) SetHeaders(v map[string]*string) *VoiceAsyncScanResponse {
	s.Headers = v
	return s
}

type VoiceAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanResultsRequest) SetClientInfo(v string) *VoiceAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type VoiceAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanResultsResponse) SetHeaders(v map[string]*string) *VoiceAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type VoiceCancelScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceCancelScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceCancelScanRequest) GoString() string {
	return s.String()
}

func (s *VoiceCancelScanRequest) SetClientInfo(v string) *VoiceCancelScanRequest {
	s.ClientInfo = &v
	return s
}

type VoiceCancelScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceCancelScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceCancelScanResponse) GoString() string {
	return s.String()
}

func (s *VoiceCancelScanResponse) SetHeaders(v map[string]*string) *VoiceCancelScanResponse {
	s.Headers = v
	return s
}

type VoiceIdentityCheckRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceIdentityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityCheckRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityCheckRequest) SetRegionId(v string) *VoiceIdentityCheckRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceIdentityCheckRequest) SetClientInfo(v string) *VoiceIdentityCheckRequest {
	s.ClientInfo = &v
	return s
}

type VoiceIdentityCheckResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceIdentityCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityCheckResponse) GoString() string {
	return s.String()
}

func (s *VoiceIdentityCheckResponse) SetHeaders(v map[string]*string) *VoiceIdentityCheckResponse {
	s.Headers = v
	return s
}

type VoiceIdentityRegisterRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceIdentityRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityRegisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityRegisterRequest) SetRegionId(v string) *VoiceIdentityRegisterRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceIdentityRegisterRequest) SetClientInfo(v string) *VoiceIdentityRegisterRequest {
	s.ClientInfo = &v
	return s
}

type VoiceIdentityRegisterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceIdentityRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityRegisterResponse) GoString() string {
	return s.String()
}

func (s *VoiceIdentityRegisterResponse) SetHeaders(v map[string]*string) *VoiceIdentityRegisterResponse {
	s.Headers = v
	return s
}

type VoiceIdentityStartCheckRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceIdentityStartCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartCheckRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartCheckRequest) SetRegionId(v string) *VoiceIdentityStartCheckRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceIdentityStartCheckRequest) SetClientInfo(v string) *VoiceIdentityStartCheckRequest {
	s.ClientInfo = &v
	return s
}

type VoiceIdentityStartCheckResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceIdentityStartCheckResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartCheckResponse) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartCheckResponse) SetHeaders(v map[string]*string) *VoiceIdentityStartCheckResponse {
	s.Headers = v
	return s
}

type VoiceIdentityStartRegisterRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceIdentityStartRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartRegisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartRegisterRequest) SetRegionId(v string) *VoiceIdentityStartRegisterRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceIdentityStartRegisterRequest) SetClientInfo(v string) *VoiceIdentityStartRegisterRequest {
	s.ClientInfo = &v
	return s
}

type VoiceIdentityStartRegisterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceIdentityStartRegisterResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartRegisterResponse) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartRegisterResponse) SetHeaders(v map[string]*string) *VoiceIdentityStartRegisterResponse {
	s.Headers = v
	return s
}

type VoiceIdentityUnregisterRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceIdentityUnregisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityUnregisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityUnregisterRequest) SetRegionId(v string) *VoiceIdentityUnregisterRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceIdentityUnregisterRequest) SetClientInfo(v string) *VoiceIdentityUnregisterRequest {
	s.ClientInfo = &v
	return s
}

type VoiceIdentityUnregisterResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceIdentityUnregisterResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityUnregisterResponse) GoString() string {
	return s.String()
}

func (s *VoiceIdentityUnregisterResponse) SetHeaders(v map[string]*string) *VoiceIdentityUnregisterResponse {
	s.Headers = v
	return s
}

type VoiceSyncScanRequest struct {
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s VoiceSyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceSyncScanRequest) GoString() string {
	return s.String()
}

func (s *VoiceSyncScanRequest) SetRegionId(v string) *VoiceSyncScanRequest {
	s.RegionId = &v
	return s
}

func (s *VoiceSyncScanRequest) SetClientInfo(v string) *VoiceSyncScanRequest {
	s.ClientInfo = &v
	return s
}

type VoiceSyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s VoiceSyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s VoiceSyncScanResponse) GoString() string {
	return s.String()
}

func (s *VoiceSyncScanResponse) SetHeaders(v map[string]*string) *VoiceSyncScanResponse {
	s.Headers = v
	return s
}

type WebpageAsyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s WebpageAsyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanRequest) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanRequest) SetClientInfo(v string) *WebpageAsyncScanRequest {
	s.ClientInfo = &v
	return s
}

type WebpageAsyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WebpageAsyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanResponse) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanResponse) SetHeaders(v map[string]*string) *WebpageAsyncScanResponse {
	s.Headers = v
	return s
}

type WebpageAsyncScanResultsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s WebpageAsyncScanResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanResultsRequest) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanResultsRequest) SetClientInfo(v string) *WebpageAsyncScanResultsRequest {
	s.ClientInfo = &v
	return s
}

type WebpageAsyncScanResultsResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WebpageAsyncScanResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanResultsResponse) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanResultsResponse) SetHeaders(v map[string]*string) *WebpageAsyncScanResultsResponse {
	s.Headers = v
	return s
}

type WebpageSyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
}

func (s WebpageSyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s WebpageSyncScanRequest) GoString() string {
	return s.String()
}

func (s *WebpageSyncScanRequest) SetClientInfo(v string) *WebpageSyncScanRequest {
	s.ClientInfo = &v
	return s
}

type WebpageSyncScanResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WebpageSyncScanResponse) String() string {
	return tea.Prettify(s)
}

func (s WebpageSyncScanResponse) GoString() string {
	return s.String()
}

func (s *WebpageSyncScanResponse) SetHeaders(v map[string]*string) *WebpageSyncScanResponse {
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
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"ap-northeast-1":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-south-1":            tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-2":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-3":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("green.aliyuncs.com"),
		"cn-hongkong":           tea.String("green.aliyuncs.com"),
		"cn-huhehaote":          tea.String("green.aliyuncs.com"),
		"cn-qingdao":            tea.String("green.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("green.aliyuncs.com"),
		"eu-central-1":          tea.String("green.ap-southeast-1.aliyuncs.com"),
		"eu-west-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"me-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"us-east-1":             tea.String("green.ap-southeast-1.aliyuncs.com"),
		"cn-hangzhou-finance":   tea.String("green.aliyuncs.com"),
		"cn-shenzhen-finance-1": tea.String("green.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("green.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("green.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("green"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) AddFaces(request *AddFacesRequest) (_result *AddFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddFacesResponse{}
	_body, _err := client.AddFacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddFacesWithOptions(request *AddFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddFacesResponse{}
	_body, _err := client.DoROARequest(tea.String("AddFaces"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/face/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddGroups(request *AddGroupsRequest) (_result *AddGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGroupsResponse{}
	_body, _err := client.AddGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddGroupsWithOptions(request *AddGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("AddGroups"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person/groups/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddPerson(request *AddPersonRequest) (_result *AddPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddPersonResponse{}
	_body, _err := client.AddPersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddPersonWithOptions(request *AddPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddPersonResponse{}
	_body, _err := client.DoROARequest(tea.String("AddPerson"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSimilarityImage(request *AddSimilarityImageRequest) (_result *AddSimilarityImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddSimilarityImageResponse{}
	_body, _err := client.AddSimilarityImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSimilarityImageWithOptions(request *AddSimilarityImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddSimilarityImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddSimilarityImageResponse{}
	_body, _err := client.DoROARequest(tea.String("AddSimilarityImage"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/image/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddSimilarityLibrary(request *AddSimilarityLibraryRequest) (_result *AddSimilarityLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddSimilarityLibraryResponse{}
	_body, _err := client.AddSimilarityLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddSimilarityLibraryWithOptions(request *AddSimilarityLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddSimilarityLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddSimilarityLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("AddSimilarityLibrary"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/library/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVideoDna(request *AddVideoDnaRequest) (_result *AddVideoDnaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddVideoDnaResponse{}
	_body, _err := client.AddVideoDnaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVideoDnaWithOptions(request *AddVideoDnaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddVideoDnaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddVideoDnaResponse{}
	_body, _err := client.DoROARequest(tea.String("AddVideoDna"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/dna/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddVideoDnaGroup(request *AddVideoDnaGroupRequest) (_result *AddVideoDnaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddVideoDnaGroupResponse{}
	_body, _err := client.AddVideoDnaGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddVideoDnaGroupWithOptions(request *AddVideoDnaGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddVideoDnaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &AddVideoDnaGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("AddVideoDnaGroup"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/dna/group/add"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaces(request *DeleteFacesRequest) (_result *DeleteFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFacesResponse{}
	_body, _err := client.DeleteFacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFacesWithOptions(request *DeleteFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteFacesResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteFaces"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/face/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteGroups(request *DeleteGroupsRequest) (_result *DeleteGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupsResponse{}
	_body, _err := client.DeleteGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteGroupsWithOptions(request *DeleteGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteGroups"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person/groups/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePerson(request *DeletePersonRequest) (_result *DeletePersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeletePersonResponse{}
	_body, _err := client.DeletePersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePersonWithOptions(request *DeletePersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeletePersonResponse{}
	_body, _err := client.DoROARequest(tea.String("DeletePerson"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSimilarityImage(request *DeleteSimilarityImageRequest) (_result *DeleteSimilarityImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSimilarityImageResponse{}
	_body, _err := client.DeleteSimilarityImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSimilarityImageWithOptions(request *DeleteSimilarityImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSimilarityImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteSimilarityImageResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSimilarityImage"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/image/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteSimilarityLibrary(request *DeleteSimilarityLibraryRequest) (_result *DeleteSimilarityLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteSimilarityLibraryResponse{}
	_body, _err := client.DeleteSimilarityLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteSimilarityLibraryWithOptions(request *DeleteSimilarityLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteSimilarityLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteSimilarityLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteSimilarityLibrary"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/library/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideoDna(request *DeleteVideoDnaRequest) (_result *DeleteVideoDnaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVideoDnaResponse{}
	_body, _err := client.DeleteVideoDnaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoDnaWithOptions(request *DeleteVideoDnaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVideoDnaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteVideoDnaResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteVideoDna"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/dna/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteVideoDnaGroup(request *DeleteVideoDnaGroupRequest) (_result *DeleteVideoDnaGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteVideoDnaGroupResponse{}
	_body, _err := client.DeleteVideoDnaGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteVideoDnaGroupWithOptions(request *DeleteVideoDnaGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteVideoDnaGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteVideoDnaGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteVideoDnaGroup"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/dna/group/delete"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DetectFace(request *DetectFaceRequest) (_result *DetectFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DetectFaceResponse{}
	_body, _err := client.DetectFaceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DetectFaceWithOptions(request *DetectFaceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DetectFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.DoROARequest(tea.String("DetectFace"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/face/detect"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileAsyncScan(request *FileAsyncScanRequest) (_result *FileAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileAsyncScanResponse{}
	_body, _err := client.FileAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileAsyncScanWithOptions(request *FileAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FileAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("FileAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/file/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FileAsyncScanResults(request *FileAsyncScanResultsRequest) (_result *FileAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileAsyncScanResultsResponse{}
	_body, _err := client.FileAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FileAsyncScanResultsWithOptions(request *FileAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &FileAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("FileAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/file/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAddVideoDnaResults(request *GetAddVideoDnaResultsRequest) (_result *GetAddVideoDnaResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAddVideoDnaResultsResponse{}
	_body, _err := client.GetAddVideoDnaResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAddVideoDnaResultsWithOptions(request *GetAddVideoDnaResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAddVideoDnaResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetAddVideoDnaResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAddVideoDnaResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/dna/add/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFaces(request *GetFacesRequest) (_result *GetFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFacesResponse{}
	_body, _err := client.GetFacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFacesWithOptions(request *GetFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetFacesResponse{}
	_body, _err := client.DoROARequest(tea.String("GetFaces"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/faces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetGroups(request *GetGroupsRequest) (_result *GetGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupsResponse{}
	_body, _err := client.GetGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetGroupsWithOptions(request *GetGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetGroups"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/groups"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPerson(request *GetPersonRequest) (_result *GetPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPersonResponse{}
	_body, _err := client.GetPersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPersonWithOptions(request *GetPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPersonResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPerson"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPersons(request *GetPersonsRequest) (_result *GetPersonsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPersonsResponse{}
	_body, _err := client.GetPersonsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPersonsWithOptions(request *GetPersonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPersonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPersonsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPersons"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/group/persons"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimilarityImage(request *GetSimilarityImageRequest) (_result *GetSimilarityImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSimilarityImageResponse{}
	_body, _err := client.GetSimilarityImageWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimilarityImageWithOptions(request *GetSimilarityImageRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSimilarityImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSimilarityImageResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSimilarityImage"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/image/get"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimilarityLibrary(request *GetSimilarityLibraryRequest) (_result *GetSimilarityLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSimilarityLibraryResponse{}
	_body, _err := client.GetSimilarityLibraryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimilarityLibraryWithOptions(request *GetSimilarityLibraryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSimilarityLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetSimilarityLibraryResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSimilarityLibrary"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/library/get"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageAsyncManualScan(request *ImageAsyncManualScanRequest) (_result *ImageAsyncManualScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageAsyncManualScanResponse{}
	_body, _err := client.ImageAsyncManualScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageAsyncManualScanWithOptions(request *ImageAsyncManualScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageAsyncManualScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageAsyncManualScanResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageAsyncManualScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/manual/asyncScan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageAsyncManualScanResults(request *ImageAsyncManualScanResultsRequest) (_result *ImageAsyncManualScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageAsyncManualScanResultsResponse{}
	_body, _err := client.ImageAsyncManualScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageAsyncManualScanResultsWithOptions(request *ImageAsyncManualScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageAsyncManualScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageAsyncManualScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageAsyncManualScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/manual/scan/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageAsyncScan(request *ImageAsyncScanRequest) (_result *ImageAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageAsyncScanResponse{}
	_body, _err := client.ImageAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageAsyncScanWithOptions(request *ImageAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageAsyncScanResults(request *ImageAsyncScanResultsRequest) (_result *ImageAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageAsyncScanResultsResponse{}
	_body, _err := client.ImageAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageAsyncScanResultsWithOptions(request *ImageAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageScanFeedback(request *ImageScanFeedbackRequest) (_result *ImageScanFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageScanFeedbackResponse{}
	_body, _err := client.ImageScanFeedbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageScanFeedbackWithOptions(request *ImageScanFeedbackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageScanFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageScanFeedbackResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageScanFeedback"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/feedback"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImageSyncScan(request *ImageSyncScanRequest) (_result *ImageSyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImageSyncScanResponse{}
	_body, _err := client.ImageSyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImageSyncScanWithOptions(request *ImageSyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImageSyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ImageSyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("ImageSyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/image/scan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSimilarityImages(request *ListSimilarityImagesRequest) (_result *ListSimilarityImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSimilarityImagesResponse{}
	_body, _err := client.ListSimilarityImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSimilarityImagesWithOptions(request *ListSimilarityImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSimilarityImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSimilarityImagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSimilarityImages"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/image/list"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSimilarityLibraries(request *ListSimilarityLibrariesRequest) (_result *ListSimilarityLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSimilarityLibrariesResponse{}
	_body, _err := client.ListSimilarityLibrariesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSimilarityLibrariesWithOptions(request *ListSimilarityLibrariesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSimilarityLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSimilarityLibrariesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSimilarityLibraries"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/similarity/library/list"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LiveStreamAsyncScan(request *LiveStreamAsyncScanRequest) (_result *LiveStreamAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LiveStreamAsyncScanResponse{}
	_body, _err := client.LiveStreamAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LiveStreamAsyncScanWithOptions(request *LiveStreamAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LiveStreamAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &LiveStreamAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("LiveStreamAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/livestream/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LiveStreamAsyncScanResults(request *LiveStreamAsyncScanResultsRequest) (_result *LiveStreamAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LiveStreamAsyncScanResultsResponse{}
	_body, _err := client.LiveStreamAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LiveStreamAsyncScanResultsWithOptions(request *LiveStreamAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LiveStreamAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &LiveStreamAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("LiveStreamAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/livestream/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LiveStreamCancelScan(request *LiveStreamCancelScanRequest) (_result *LiveStreamCancelScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LiveStreamCancelScanResponse{}
	_body, _err := client.LiveStreamCancelScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LiveStreamCancelScanWithOptions(request *LiveStreamCancelScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LiveStreamCancelScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &LiveStreamCancelScanResponse{}
	_body, _err := client.DoROARequest(tea.String("LiveStreamCancelScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/livestream/cancelscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostAsyncScan(request *PostAsyncScanRequest) (_result *PostAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostAsyncScanResponse{}
	_body, _err := client.PostAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostAsyncScanWithOptions(request *PostAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PostAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &PostAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("PostAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/post/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostAsyncScanResults(request *PostAsyncScanResultsRequest) (_result *PostAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostAsyncScanResultsResponse{}
	_body, _err := client.PostAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostAsyncScanResultsWithOptions(request *PostAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PostAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &PostAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("PostAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/post/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchPerson(request *SearchPersonRequest) (_result *SearchPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchPersonResponse{}
	_body, _err := client.SearchPersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchPersonWithOptions(request *SearchPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &SearchPersonResponse{}
	_body, _err := client.DoROARequest(tea.String("SearchPerson"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/search"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetPerson(request *SetPersonRequest) (_result *SetPersonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SetPersonResponse{}
	_body, _err := client.SetPersonWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetPersonWithOptions(request *SetPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SetPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &SetPersonResponse{}
	_body, _err := client.DoROARequest(tea.String("SetPerson"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/sface/person/update"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextAsyncManualScan(request *TextAsyncManualScanRequest) (_result *TextAsyncManualScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TextAsyncManualScanResponse{}
	_body, _err := client.TextAsyncManualScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextAsyncManualScanWithOptions(request *TextAsyncManualScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TextAsyncManualScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &TextAsyncManualScanResponse{}
	_body, _err := client.DoROARequest(tea.String("TextAsyncManualScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/text/manual/asyncScan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextAsyncManualScanResults(request *TextAsyncManualScanResultsRequest) (_result *TextAsyncManualScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TextAsyncManualScanResultsResponse{}
	_body, _err := client.TextAsyncManualScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextAsyncManualScanResultsWithOptions(request *TextAsyncManualScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TextAsyncManualScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &TextAsyncManualScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("TextAsyncManualScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/text/manual/scan/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextFeedback(request *TextFeedbackRequest) (_result *TextFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TextFeedbackResponse{}
	_body, _err := client.TextFeedbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextFeedbackWithOptions(request *TextFeedbackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TextFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &TextFeedbackResponse{}
	_body, _err := client.DoROARequest(tea.String("TextFeedback"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/text/feedback"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextScan(request *TextScanRequest) (_result *TextScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TextScanResponse{}
	_body, _err := client.TextScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextScanWithOptions(request *TextScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TextScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &TextScanResponse{}
	_body, _err := client.DoROARequest(tea.String("TextScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/text/scan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadCredentials(request *UploadCredentialsRequest) (_result *UploadCredentialsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCredentialsResponse{}
	_body, _err := client.UploadCredentialsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadCredentialsWithOptions(request *UploadCredentialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UploadCredentialsResponse{}
	_body, _err := client.DoROARequest(tea.String("UploadCredentials"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/credentials/uploadcredentials"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoAsyncManualScan(request *VideoAsyncManualScanRequest) (_result *VideoAsyncManualScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoAsyncManualScanResponse{}
	_body, _err := client.VideoAsyncManualScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoAsyncManualScanWithOptions(request *VideoAsyncManualScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoAsyncManualScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoAsyncManualScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoAsyncManualScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/manual/asyncScan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoAsyncManualScanResults(request *VideoAsyncManualScanResultsRequest) (_result *VideoAsyncManualScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoAsyncManualScanResultsResponse{}
	_body, _err := client.VideoAsyncManualScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoAsyncManualScanResultsWithOptions(request *VideoAsyncManualScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoAsyncManualScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoAsyncManualScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoAsyncManualScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/manual/scan/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoAsyncScan(request *VideoAsyncScanRequest) (_result *VideoAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoAsyncScanResponse{}
	_body, _err := client.VideoAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoAsyncScanWithOptions(request *VideoAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoAsyncScanResults(request *VideoAsyncScanResultsRequest) (_result *VideoAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoAsyncScanResultsResponse{}
	_body, _err := client.VideoAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoAsyncScanResultsWithOptions(request *VideoAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoCancelScan(request *VideoCancelScanRequest) (_result *VideoCancelScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoCancelScanResponse{}
	_body, _err := client.VideoCancelScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoCancelScanWithOptions(request *VideoCancelScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoCancelScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoCancelScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoCancelScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/cancelscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoFeedback(request *VideoFeedbackRequest) (_result *VideoFeedbackResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoFeedbackResponse{}
	_body, _err := client.VideoFeedbackWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoFeedbackWithOptions(request *VideoFeedbackRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoFeedbackResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoFeedbackResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoFeedback"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/feedback"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VideoSyncScan(request *VideoSyncScanRequest) (_result *VideoSyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VideoSyncScanResponse{}
	_body, _err := client.VideoSyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VideoSyncScanWithOptions(request *VideoSyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VideoSyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VideoSyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VideoSyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/video/syncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VodAsyncScan(request *VodAsyncScanRequest) (_result *VodAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VodAsyncScanResponse{}
	_body, _err := client.VodAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VodAsyncScanWithOptions(request *VodAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VodAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VodAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VodAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/vod/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VodAsyncScanResults(request *VodAsyncScanResultsRequest) (_result *VodAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VodAsyncScanResultsResponse{}
	_body, _err := client.VodAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VodAsyncScanResultsWithOptions(request *VodAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VodAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VodAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("VodAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/vod/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceAsyncManualScan(request *VoiceAsyncManualScanRequest) (_result *VoiceAsyncManualScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceAsyncManualScanResponse{}
	_body, _err := client.VoiceAsyncManualScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceAsyncManualScanWithOptions(request *VoiceAsyncManualScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceAsyncManualScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceAsyncManualScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceAsyncManualScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/manual/asyncScan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceAsyncManualScanResults(request *VoiceAsyncManualScanResultsRequest) (_result *VoiceAsyncManualScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceAsyncManualScanResultsResponse{}
	_body, _err := client.VoiceAsyncManualScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceAsyncManualScanResultsWithOptions(request *VoiceAsyncManualScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceAsyncManualScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceAsyncManualScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceAsyncManualScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/manual/scan/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceAsyncScan(request *VoiceAsyncScanRequest) (_result *VoiceAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceAsyncScanResponse{}
	_body, _err := client.VoiceAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceAsyncScanWithOptions(request *VoiceAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceAsyncScanResults(request *VoiceAsyncScanResultsRequest) (_result *VoiceAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceAsyncScanResultsResponse{}
	_body, _err := client.VoiceAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceAsyncScanResultsWithOptions(request *VoiceAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceCancelScan(request *VoiceCancelScanRequest) (_result *VoiceCancelScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceCancelScanResponse{}
	_body, _err := client.VoiceCancelScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceCancelScanWithOptions(request *VoiceCancelScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceCancelScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceCancelScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceCancelScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/cancelscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceIdentityCheck(request *VoiceIdentityCheckRequest) (_result *VoiceIdentityCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceIdentityCheckResponse{}
	_body, _err := client.VoiceIdentityCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceIdentityCheckWithOptions(request *VoiceIdentityCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceIdentityCheckResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceIdentityCheck"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/auth/check"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceIdentityRegister(request *VoiceIdentityRegisterRequest) (_result *VoiceIdentityRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceIdentityRegisterResponse{}
	_body, _err := client.VoiceIdentityRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceIdentityRegisterWithOptions(request *VoiceIdentityRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceIdentityRegisterResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceIdentityRegister"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/auth/register"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceIdentityStartCheck(request *VoiceIdentityStartCheckRequest) (_result *VoiceIdentityStartCheckResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceIdentityStartCheckResponse{}
	_body, _err := client.VoiceIdentityStartCheckWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceIdentityStartCheckWithOptions(request *VoiceIdentityStartCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityStartCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceIdentityStartCheckResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceIdentityStartCheck"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/auth/start/check"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceIdentityStartRegister(request *VoiceIdentityStartRegisterRequest) (_result *VoiceIdentityStartRegisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceIdentityStartRegisterResponse{}
	_body, _err := client.VoiceIdentityStartRegisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceIdentityStartRegisterWithOptions(request *VoiceIdentityStartRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityStartRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceIdentityStartRegisterResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceIdentityStartRegister"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/auth/start/register"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceIdentityUnregister(request *VoiceIdentityUnregisterRequest) (_result *VoiceIdentityUnregisterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceIdentityUnregisterResponse{}
	_body, _err := client.VoiceIdentityUnregisterWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceIdentityUnregisterWithOptions(request *VoiceIdentityUnregisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityUnregisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceIdentityUnregisterResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceIdentityUnregister"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/auth/unregister"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) VoiceSyncScan(request *VoiceSyncScanRequest) (_result *VoiceSyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VoiceSyncScanResponse{}
	_body, _err := client.VoiceSyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) VoiceSyncScanWithOptions(request *VoiceSyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceSyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &VoiceSyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("VoiceSyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/voice/syncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WebpageAsyncScan(request *WebpageAsyncScanRequest) (_result *WebpageAsyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebpageAsyncScanResponse{}
	_body, _err := client.WebpageAsyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WebpageAsyncScanWithOptions(request *WebpageAsyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WebpageAsyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &WebpageAsyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("WebpageAsyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/webpage/asyncscan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WebpageAsyncScanResults(request *WebpageAsyncScanResultsRequest) (_result *WebpageAsyncScanResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebpageAsyncScanResultsResponse{}
	_body, _err := client.WebpageAsyncScanResultsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WebpageAsyncScanResultsWithOptions(request *WebpageAsyncScanResultsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WebpageAsyncScanResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &WebpageAsyncScanResultsResponse{}
	_body, _err := client.DoROARequest(tea.String("WebpageAsyncScanResults"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/webpage/results"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WebpageSyncScan(request *WebpageSyncScanRequest) (_result *WebpageSyncScanResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WebpageSyncScanResponse{}
	_body, _err := client.WebpageSyncScanWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WebpageSyncScanWithOptions(request *WebpageSyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WebpageSyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &WebpageSyncScanResponse{}
	_body, _err := client.DoROARequest(tea.String("WebpageSyncScan"), tea.String("2018-05-09"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/green/webpage/scan"), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
