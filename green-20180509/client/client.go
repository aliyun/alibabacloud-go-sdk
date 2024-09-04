// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AddFacesRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddFacesRequest) GoString() string {
	return s.String()
}

func (s *AddFacesRequest) SetClientInfo(v string) *AddFacesRequest {
	s.ClientInfo = &v
	return s
}

func (s *AddFacesRequest) SetRegionId(v string) *AddFacesRequest {
	s.RegionId = &v
	return s
}

type AddFacesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddFacesResponseBody) GoString() string {
	return s.String()
}

func (s *AddFacesResponseBody) SetRequestId(v string) *AddFacesResponseBody {
	s.RequestId = &v
	return s
}

type AddFacesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddFacesResponse) SetStatusCode(v int32) *AddFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFacesResponse) SetBody(v *AddFacesResponseBody) *AddFacesResponse {
	s.Body = v
	return s
}

type AddGroupsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupsRequest) GoString() string {
	return s.String()
}

func (s *AddGroupsRequest) SetClientInfo(v string) *AddGroupsRequest {
	s.ClientInfo = &v
	return s
}

func (s *AddGroupsRequest) SetRegionId(v string) *AddGroupsRequest {
	s.RegionId = &v
	return s
}

type AddGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *AddGroupsResponseBody) SetRequestId(v string) *AddGroupsResponseBody {
	s.RequestId = &v
	return s
}

type AddGroupsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddGroupsResponse) SetStatusCode(v int32) *AddGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGroupsResponse) SetBody(v *AddGroupsResponseBody) *AddGroupsResponse {
	s.Body = v
	return s
}

type AddPersonRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s AddPersonRequest) GoString() string {
	return s.String()
}

func (s *AddPersonRequest) SetClientInfo(v string) *AddPersonRequest {
	s.ClientInfo = &v
	return s
}

func (s *AddPersonRequest) SetRegionId(v string) *AddPersonRequest {
	s.RegionId = &v
	return s
}

type AddPersonResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddPersonResponseBody) GoString() string {
	return s.String()
}

func (s *AddPersonResponseBody) SetRequestId(v string) *AddPersonResponseBody {
	s.RequestId = &v
	return s
}

type AddPersonResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddPersonResponse) SetStatusCode(v int32) *AddPersonResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPersonResponse) SetBody(v *AddPersonResponseBody) *AddPersonResponse {
	s.Body = v
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

type AddSimilarityImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSimilarityImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityImageResponseBody) GoString() string {
	return s.String()
}

func (s *AddSimilarityImageResponseBody) SetRequestId(v string) *AddSimilarityImageResponseBody {
	s.RequestId = &v
	return s
}

type AddSimilarityImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSimilarityImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddSimilarityImageResponse) SetStatusCode(v int32) *AddSimilarityImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSimilarityImageResponse) SetBody(v *AddSimilarityImageResponseBody) *AddSimilarityImageResponse {
	s.Body = v
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

type AddSimilarityLibraryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSimilarityLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddSimilarityLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *AddSimilarityLibraryResponseBody) SetRequestId(v string) *AddSimilarityLibraryResponseBody {
	s.RequestId = &v
	return s
}

type AddSimilarityLibraryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSimilarityLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddSimilarityLibraryResponse) SetStatusCode(v int32) *AddSimilarityLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSimilarityLibraryResponse) SetBody(v *AddSimilarityLibraryResponseBody) *AddSimilarityLibraryResponse {
	s.Body = v
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

type AddVideoDnaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVideoDnaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaResponseBody) GoString() string {
	return s.String()
}

func (s *AddVideoDnaResponseBody) SetRequestId(v string) *AddVideoDnaResponseBody {
	s.RequestId = &v
	return s
}

type AddVideoDnaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVideoDnaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddVideoDnaResponse) SetStatusCode(v int32) *AddVideoDnaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVideoDnaResponse) SetBody(v *AddVideoDnaResponseBody) *AddVideoDnaResponse {
	s.Body = v
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

type AddVideoDnaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVideoDnaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddVideoDnaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddVideoDnaGroupResponseBody) SetRequestId(v string) *AddVideoDnaGroupResponseBody {
	s.RequestId = &v
	return s
}

type AddVideoDnaGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddVideoDnaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *AddVideoDnaGroupResponse) SetStatusCode(v int32) *AddVideoDnaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddVideoDnaGroupResponse) SetBody(v *AddVideoDnaGroupResponseBody) *AddVideoDnaGroupResponse {
	s.Body = v
	return s
}

type DeleteFacesRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFacesRequest) SetClientInfo(v string) *DeleteFacesRequest {
	s.ClientInfo = &v
	return s
}

func (s *DeleteFacesRequest) SetRegionId(v string) *DeleteFacesRequest {
	s.RegionId = &v
	return s
}

type DeleteFacesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFacesResponseBody) SetRequestId(v string) *DeleteFacesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteFacesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteFacesResponse) SetStatusCode(v int32) *DeleteFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFacesResponse) SetBody(v *DeleteFacesResponseBody) *DeleteFacesResponse {
	s.Body = v
	return s
}

type DeleteGroupsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupsRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupsRequest) SetClientInfo(v string) *DeleteGroupsRequest {
	s.ClientInfo = &v
	return s
}

func (s *DeleteGroupsRequest) SetRegionId(v string) *DeleteGroupsRequest {
	s.RegionId = &v
	return s
}

type DeleteGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupsResponseBody) SetRequestId(v string) *DeleteGroupsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteGroupsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteGroupsResponse) SetStatusCode(v int32) *DeleteGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupsResponse) SetBody(v *DeleteGroupsResponseBody) *DeleteGroupsResponse {
	s.Body = v
	return s
}

type DeletePersonRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePersonRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePersonRequest) GoString() string {
	return s.String()
}

func (s *DeletePersonRequest) SetClientInfo(v string) *DeletePersonRequest {
	s.ClientInfo = &v
	return s
}

func (s *DeletePersonRequest) SetRegionId(v string) *DeletePersonRequest {
	s.RegionId = &v
	return s
}

type DeletePersonResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePersonResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePersonResponseBody) SetRequestId(v string) *DeletePersonResponseBody {
	s.RequestId = &v
	return s
}

type DeletePersonResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeletePersonResponse) SetStatusCode(v int32) *DeletePersonResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePersonResponse) SetBody(v *DeletePersonResponseBody) *DeletePersonResponse {
	s.Body = v
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

type DeleteSimilarityImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSimilarityImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityImageResponseBody) SetRequestId(v string) *DeleteSimilarityImageResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSimilarityImageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSimilarityImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteSimilarityImageResponse) SetStatusCode(v int32) *DeleteSimilarityImageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSimilarityImageResponse) SetBody(v *DeleteSimilarityImageResponseBody) *DeleteSimilarityImageResponse {
	s.Body = v
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

type DeleteSimilarityLibraryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSimilarityLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteSimilarityLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSimilarityLibraryResponseBody) SetRequestId(v string) *DeleteSimilarityLibraryResponseBody {
	s.RequestId = &v
	return s
}

type DeleteSimilarityLibraryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSimilarityLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteSimilarityLibraryResponse) SetStatusCode(v int32) *DeleteSimilarityLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSimilarityLibraryResponse) SetBody(v *DeleteSimilarityLibraryResponseBody) *DeleteSimilarityLibraryResponse {
	s.Body = v
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

type DeleteVideoDnaResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoDnaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaResponseBody) SetRequestId(v string) *DeleteVideoDnaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVideoDnaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVideoDnaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteVideoDnaResponse) SetStatusCode(v int32) *DeleteVideoDnaResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoDnaResponse) SetBody(v *DeleteVideoDnaResponseBody) *DeleteVideoDnaResponse {
	s.Body = v
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

type DeleteVideoDnaGroupResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoDnaGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVideoDnaGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoDnaGroupResponseBody) SetRequestId(v string) *DeleteVideoDnaGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVideoDnaGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVideoDnaGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DeleteVideoDnaGroupResponse) SetStatusCode(v int32) *DeleteVideoDnaGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoDnaGroupResponse) SetBody(v *DeleteVideoDnaGroupResponseBody) *DeleteVideoDnaGroupResponse {
	s.Body = v
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

type DetectFaceResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DetectFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DetectFaceResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceResponseBody) SetRequestId(v string) *DetectFaceResponseBody {
	s.RequestId = &v
	return s
}

type DetectFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *DetectFaceResponse) SetStatusCode(v int32) *DetectFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectFaceResponse) SetBody(v *DetectFaceResponseBody) *DetectFaceResponse {
	s.Body = v
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

type FileAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FileAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *FileAsyncScanResponseBody) SetRequestId(v string) *FileAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type FileAsyncScanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *FileAsyncScanResponse) SetStatusCode(v int32) *FileAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *FileAsyncScanResponse) SetBody(v *FileAsyncScanResponseBody) *FileAsyncScanResponse {
	s.Body = v
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

type FileAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FileAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *FileAsyncScanResultsResponseBody) SetRequestId(v string) *FileAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type FileAsyncScanResultsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *FileAsyncScanResultsResponse) SetStatusCode(v int32) *FileAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *FileAsyncScanResultsResponse) SetBody(v *FileAsyncScanResultsResponseBody) *FileAsyncScanResultsResponse {
	s.Body = v
	return s
}

type FileAsyncScanV2ResponseBody struct {
	// example:
	//
	// DA36A1DA-C466-538D-AD52-E64D75597750
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s FileAsyncScanV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanV2ResponseBody) GoString() string {
	return s.String()
}

func (s *FileAsyncScanV2ResponseBody) SetRequestId(v string) *FileAsyncScanV2ResponseBody {
	s.RequestId = &v
	return s
}

type FileAsyncScanV2Response struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FileAsyncScanV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FileAsyncScanV2Response) String() string {
	return tea.Prettify(s)
}

func (s FileAsyncScanV2Response) GoString() string {
	return s.String()
}

func (s *FileAsyncScanV2Response) SetHeaders(v map[string]*string) *FileAsyncScanV2Response {
	s.Headers = v
	return s
}

func (s *FileAsyncScanV2Response) SetStatusCode(v int32) *FileAsyncScanV2Response {
	s.StatusCode = &v
	return s
}

func (s *FileAsyncScanV2Response) SetBody(v *FileAsyncScanV2ResponseBody) *FileAsyncScanV2Response {
	s.Body = v
	return s
}

type GetFacesRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFacesRequest) GoString() string {
	return s.String()
}

func (s *GetFacesRequest) SetClientInfo(v string) *GetFacesRequest {
	s.ClientInfo = &v
	return s
}

func (s *GetFacesRequest) SetRegionId(v string) *GetFacesRequest {
	s.RegionId = &v
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetFacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetFacesResponse) SetStatusCode(v int32) *GetFacesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFacesResponse) SetBody(v *GetFacesResponseBody) *GetFacesResponse {
	s.Body = v
	return s
}

type GetGroupsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupsRequest) GoString() string {
	return s.String()
}

func (s *GetGroupsRequest) SetClientInfo(v string) *GetGroupsRequest {
	s.ClientInfo = &v
	return s
}

func (s *GetGroupsRequest) SetRegionId(v string) *GetGroupsRequest {
	s.RegionId = &v
	return s
}

type GetGroupsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *GetGroupsResponseBody) SetRequestId(v string) *GetGroupsResponseBody {
	s.RequestId = &v
	return s
}

type GetGroupsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetGroupsResponse) SetStatusCode(v int32) *GetGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupsResponse) SetBody(v *GetGroupsResponseBody) *GetGroupsResponse {
	s.Body = v
	return s
}

type GetPersonRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonRequest) GoString() string {
	return s.String()
}

func (s *GetPersonRequest) SetClientInfo(v string) *GetPersonRequest {
	s.ClientInfo = &v
	return s
}

func (s *GetPersonRequest) SetRegionId(v string) *GetPersonRequest {
	s.RegionId = &v
	return s
}

type GetPersonResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonResponseBody) SetRequestId(v string) *GetPersonResponseBody {
	s.RequestId = &v
	return s
}

type GetPersonResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetPersonResponse) SetStatusCode(v int32) *GetPersonResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPersonResponse) SetBody(v *GetPersonResponseBody) *GetPersonResponse {
	s.Body = v
	return s
}

type GetPersonsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetPersonsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPersonsRequest) GoString() string {
	return s.String()
}

func (s *GetPersonsRequest) SetClientInfo(v string) *GetPersonsRequest {
	s.ClientInfo = &v
	return s
}

func (s *GetPersonsRequest) SetRegionId(v string) *GetPersonsRequest {
	s.RegionId = &v
	return s
}

type GetPersonsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPersonsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPersonsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPersonsResponseBody) SetRequestId(v string) *GetPersonsResponseBody {
	s.RequestId = &v
	return s
}

type GetPersonsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPersonsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetPersonsResponse) SetStatusCode(v int32) *GetPersonsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPersonsResponse) SetBody(v *GetPersonsResponseBody) *GetPersonsResponse {
	s.Body = v
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

type GetSimilarityImageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSimilarityImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityImageResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimilarityImageResponseBody) SetRequestId(v string) *GetSimilarityImageResponseBody {
	s.RequestId = &v
	return s
}

type GetSimilarityImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSimilarityImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetSimilarityImageResponse) SetStatusCode(v int32) *GetSimilarityImageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSimilarityImageResponse) SetBody(v *GetSimilarityImageResponseBody) *GetSimilarityImageResponse {
	s.Body = v
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

type GetSimilarityLibraryResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSimilarityLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarityLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimilarityLibraryResponseBody) SetRequestId(v string) *GetSimilarityLibraryResponseBody {
	s.RequestId = &v
	return s
}

type GetSimilarityLibraryResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSimilarityLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *GetSimilarityLibraryResponse) SetStatusCode(v int32) *GetSimilarityLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSimilarityLibraryResponse) SetBody(v *GetSimilarityLibraryResponseBody) *GetSimilarityLibraryResponse {
	s.Body = v
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

type ImageAsyncManualScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncManualScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanResponseBody) SetRequestId(v string) *ImageAsyncManualScanResponseBody {
	s.RequestId = &v
	return s
}

type ImageAsyncManualScanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncManualScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageAsyncManualScanResponse) SetStatusCode(v int32) *ImageAsyncManualScanResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncManualScanResponse) SetBody(v *ImageAsyncManualScanResponseBody) *ImageAsyncManualScanResponse {
	s.Body = v
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

type ImageAsyncManualScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncManualScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncManualScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncManualScanResultsResponseBody) SetRequestId(v string) *ImageAsyncManualScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type ImageAsyncManualScanResultsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncManualScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageAsyncManualScanResultsResponse) SetStatusCode(v int32) *ImageAsyncManualScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncManualScanResultsResponse) SetBody(v *ImageAsyncManualScanResultsResponseBody) *ImageAsyncManualScanResultsResponse {
	s.Body = v
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

type ImageAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanResponseBody) SetRequestId(v string) *ImageAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type ImageAsyncScanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageAsyncScanResponse) SetStatusCode(v int32) *ImageAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncScanResponse) SetBody(v *ImageAsyncScanResponseBody) *ImageAsyncScanResponse {
	s.Body = v
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

type ImageAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ImageAsyncScanResultsResponseBody) SetRequestId(v string) *ImageAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type ImageAsyncScanResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageAsyncScanResultsResponse) SetStatusCode(v int32) *ImageAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageAsyncScanResultsResponse) SetBody(v *ImageAsyncScanResultsResponseBody) *ImageAsyncScanResultsResponse {
	s.Body = v
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

type ImageScanFeedbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageScanFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageScanFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *ImageScanFeedbackResponseBody) SetRequestId(v string) *ImageScanFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type ImageScanFeedbackResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageScanFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageScanFeedbackResponse) SetStatusCode(v int32) *ImageScanFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageScanFeedbackResponse) SetBody(v *ImageScanFeedbackResponseBody) *ImageScanFeedbackResponse {
	s.Body = v
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

type ImageSyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageSyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImageSyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *ImageSyncScanResponseBody) SetRequestId(v string) *ImageSyncScanResponseBody {
	s.RequestId = &v
	return s
}

type ImageSyncScanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImageSyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ImageSyncScanResponse) SetStatusCode(v int32) *ImageSyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *ImageSyncScanResponse) SetBody(v *ImageSyncScanResponseBody) *ImageSyncScanResponse {
	s.Body = v
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

type ListSimilarityImagesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSimilarityImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimilarityImagesResponseBody) SetRequestId(v string) *ListSimilarityImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListSimilarityImagesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSimilarityImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListSimilarityImagesResponse) SetStatusCode(v int32) *ListSimilarityImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSimilarityImagesResponse) SetBody(v *ListSimilarityImagesResponseBody) *ListSimilarityImagesResponse {
	s.Body = v
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

type ListSimilarityLibrariesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSimilarityLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSimilarityLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSimilarityLibrariesResponseBody) SetRequestId(v string) *ListSimilarityLibrariesResponseBody {
	s.RequestId = &v
	return s
}

type ListSimilarityLibrariesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSimilarityLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ListSimilarityLibrariesResponse) SetStatusCode(v int32) *ListSimilarityLibrariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSimilarityLibrariesResponse) SetBody(v *ListSimilarityLibrariesResponseBody) *ListSimilarityLibrariesResponse {
	s.Body = v
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

type LiveStreamAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LiveStreamAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanResponseBody) SetRequestId(v string) *LiveStreamAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type LiveStreamAsyncScanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiveStreamAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *LiveStreamAsyncScanResponse) SetStatusCode(v int32) *LiveStreamAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *LiveStreamAsyncScanResponse) SetBody(v *LiveStreamAsyncScanResponseBody) *LiveStreamAsyncScanResponse {
	s.Body = v
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

type LiveStreamAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LiveStreamAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *LiveStreamAsyncScanResultsResponseBody) SetRequestId(v string) *LiveStreamAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type LiveStreamAsyncScanResultsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiveStreamAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *LiveStreamAsyncScanResultsResponse) SetStatusCode(v int32) *LiveStreamAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *LiveStreamAsyncScanResultsResponse) SetBody(v *LiveStreamAsyncScanResultsResponseBody) *LiveStreamAsyncScanResultsResponse {
	s.Body = v
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

type LiveStreamCancelScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s LiveStreamCancelScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LiveStreamCancelScanResponseBody) GoString() string {
	return s.String()
}

func (s *LiveStreamCancelScanResponseBody) SetRequestId(v string) *LiveStreamCancelScanResponseBody {
	s.RequestId = &v
	return s
}

type LiveStreamCancelScanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LiveStreamCancelScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *LiveStreamCancelScanResponse) SetStatusCode(v int32) *LiveStreamCancelScanResponse {
	s.StatusCode = &v
	return s
}

func (s *LiveStreamCancelScanResponse) SetBody(v *LiveStreamCancelScanResponseBody) *LiveStreamCancelScanResponse {
	s.Body = v
	return s
}

type SetPersonRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetPersonRequest) String() string {
	return tea.Prettify(s)
}

func (s SetPersonRequest) GoString() string {
	return s.String()
}

func (s *SetPersonRequest) SetClientInfo(v string) *SetPersonRequest {
	s.ClientInfo = &v
	return s
}

func (s *SetPersonRequest) SetRegionId(v string) *SetPersonRequest {
	s.RegionId = &v
	return s
}

type SetPersonResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPersonResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetPersonResponseBody) GoString() string {
	return s.String()
}

func (s *SetPersonResponseBody) SetRequestId(v string) *SetPersonResponseBody {
	s.RequestId = &v
	return s
}

type SetPersonResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetPersonResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *SetPersonResponse) SetStatusCode(v int32) *SetPersonResponse {
	s.StatusCode = &v
	return s
}

func (s *SetPersonResponse) SetBody(v *SetPersonResponseBody) *SetPersonResponse {
	s.Body = v
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

type TextAsyncManualScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextAsyncManualScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanResponseBody) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanResponseBody) SetRequestId(v string) *TextAsyncManualScanResponseBody {
	s.RequestId = &v
	return s
}

type TextAsyncManualScanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextAsyncManualScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TextAsyncManualScanResponse) SetStatusCode(v int32) *TextAsyncManualScanResponse {
	s.StatusCode = &v
	return s
}

func (s *TextAsyncManualScanResponse) SetBody(v *TextAsyncManualScanResponseBody) *TextAsyncManualScanResponse {
	s.Body = v
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

type TextAsyncManualScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextAsyncManualScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextAsyncManualScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *TextAsyncManualScanResultsResponseBody) SetRequestId(v string) *TextAsyncManualScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type TextAsyncManualScanResultsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextAsyncManualScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TextAsyncManualScanResultsResponse) SetStatusCode(v int32) *TextAsyncManualScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *TextAsyncManualScanResultsResponse) SetBody(v *TextAsyncManualScanResultsResponseBody) *TextAsyncManualScanResultsResponse {
	s.Body = v
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

type TextFeedbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *TextFeedbackResponseBody) SetRequestId(v string) *TextFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type TextFeedbackResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TextFeedbackResponse) SetStatusCode(v int32) *TextFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *TextFeedbackResponse) SetBody(v *TextFeedbackResponseBody) *TextFeedbackResponse {
	s.Body = v
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

type TextScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TextScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TextScanResponseBody) GoString() string {
	return s.String()
}

func (s *TextScanResponseBody) SetRequestId(v string) *TextScanResponseBody {
	s.RequestId = &v
	return s
}

type TextScanResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *TextScanResponse) SetStatusCode(v int32) *TextScanResponse {
	s.StatusCode = &v
	return s
}

func (s *TextScanResponse) SetBody(v *TextScanResponseBody) *TextScanResponse {
	s.Body = v
	return s
}

type UploadCredentialsRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UploadCredentialsRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCredentialsRequest) GoString() string {
	return s.String()
}

func (s *UploadCredentialsRequest) SetClientInfo(v string) *UploadCredentialsRequest {
	s.ClientInfo = &v
	return s
}

func (s *UploadCredentialsRequest) SetRegionId(v string) *UploadCredentialsRequest {
	s.RegionId = &v
	return s
}

type UploadCredentialsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadCredentialsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *UploadCredentialsResponseBody) SetRequestId(v string) *UploadCredentialsResponseBody {
	s.RequestId = &v
	return s
}

type UploadCredentialsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *UploadCredentialsResponse) SetStatusCode(v int32) *UploadCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadCredentialsResponse) SetBody(v *UploadCredentialsResponseBody) *UploadCredentialsResponse {
	s.Body = v
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

type VideoAsyncManualScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoAsyncManualScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanResponseBody) SetRequestId(v string) *VideoAsyncManualScanResponseBody {
	s.RequestId = &v
	return s
}

type VideoAsyncManualScanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAsyncManualScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoAsyncManualScanResponse) SetStatusCode(v int32) *VideoAsyncManualScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAsyncManualScanResponse) SetBody(v *VideoAsyncManualScanResponseBody) *VideoAsyncManualScanResponse {
	s.Body = v
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

type VideoAsyncManualScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoAsyncManualScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncManualScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAsyncManualScanResultsResponseBody) SetRequestId(v string) *VideoAsyncManualScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type VideoAsyncManualScanResultsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAsyncManualScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoAsyncManualScanResultsResponse) SetStatusCode(v int32) *VideoAsyncManualScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAsyncManualScanResultsResponse) SetBody(v *VideoAsyncManualScanResultsResponseBody) *VideoAsyncManualScanResultsResponse {
	s.Body = v
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

type VideoAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanResponseBody) SetRequestId(v string) *VideoAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type VideoAsyncScanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoAsyncScanResponse) SetStatusCode(v int32) *VideoAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAsyncScanResponse) SetBody(v *VideoAsyncScanResponseBody) *VideoAsyncScanResponse {
	s.Body = v
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

type VideoAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *VideoAsyncScanResultsResponseBody) SetRequestId(v string) *VideoAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type VideoAsyncScanResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoAsyncScanResultsResponse) SetStatusCode(v int32) *VideoAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoAsyncScanResultsResponse) SetBody(v *VideoAsyncScanResultsResponseBody) *VideoAsyncScanResultsResponse {
	s.Body = v
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

type VideoCancelScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoCancelScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoCancelScanResponseBody) GoString() string {
	return s.String()
}

func (s *VideoCancelScanResponseBody) SetRequestId(v string) *VideoCancelScanResponseBody {
	s.RequestId = &v
	return s
}

type VideoCancelScanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoCancelScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoCancelScanResponse) SetStatusCode(v int32) *VideoCancelScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoCancelScanResponse) SetBody(v *VideoCancelScanResponseBody) *VideoCancelScanResponse {
	s.Body = v
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

type VideoFeedbackResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoFeedbackResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *VideoFeedbackResponseBody) SetRequestId(v string) *VideoFeedbackResponseBody {
	s.RequestId = &v
	return s
}

type VideoFeedbackResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoFeedbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoFeedbackResponse) SetStatusCode(v int32) *VideoFeedbackResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoFeedbackResponse) SetBody(v *VideoFeedbackResponseBody) *VideoFeedbackResponse {
	s.Body = v
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

type VideoSyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VideoSyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VideoSyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *VideoSyncScanResponseBody) SetRequestId(v string) *VideoSyncScanResponseBody {
	s.RequestId = &v
	return s
}

type VideoSyncScanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VideoSyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VideoSyncScanResponse) SetStatusCode(v int32) *VideoSyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VideoSyncScanResponse) SetBody(v *VideoSyncScanResponseBody) *VideoSyncScanResponse {
	s.Body = v
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

type VodAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VodAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *VodAsyncScanResponseBody) SetRequestId(v string) *VodAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type VodAsyncScanResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VodAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VodAsyncScanResponse) SetStatusCode(v int32) *VodAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VodAsyncScanResponse) SetBody(v *VodAsyncScanResponseBody) *VodAsyncScanResponse {
	s.Body = v
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

type VodAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VodAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VodAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *VodAsyncScanResultsResponseBody) SetRequestId(v string) *VodAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type VodAsyncScanResultsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VodAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VodAsyncScanResultsResponse) SetStatusCode(v int32) *VodAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *VodAsyncScanResultsResponse) SetBody(v *VodAsyncScanResultsResponseBody) *VodAsyncScanResultsResponse {
	s.Body = v
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

type VoiceAsyncManualScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceAsyncManualScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanResponseBody) SetRequestId(v string) *VoiceAsyncManualScanResponseBody {
	s.RequestId = &v
	return s
}

type VoiceAsyncManualScanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceAsyncManualScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceAsyncManualScanResponse) SetStatusCode(v int32) *VoiceAsyncManualScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceAsyncManualScanResponse) SetBody(v *VoiceAsyncManualScanResponseBody) *VoiceAsyncManualScanResponse {
	s.Body = v
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

type VoiceAsyncManualScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceAsyncManualScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncManualScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceAsyncManualScanResultsResponseBody) SetRequestId(v string) *VoiceAsyncManualScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type VoiceAsyncManualScanResultsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceAsyncManualScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceAsyncManualScanResultsResponse) SetStatusCode(v int32) *VoiceAsyncManualScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceAsyncManualScanResultsResponse) SetBody(v *VoiceAsyncManualScanResultsResponseBody) *VoiceAsyncManualScanResultsResponse {
	s.Body = v
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

type VoiceAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanResponseBody) SetRequestId(v string) *VoiceAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type VoiceAsyncScanResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceAsyncScanResponse) SetStatusCode(v int32) *VoiceAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceAsyncScanResponse) SetBody(v *VoiceAsyncScanResponseBody) *VoiceAsyncScanResponse {
	s.Body = v
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

type VoiceAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceAsyncScanResultsResponseBody) SetRequestId(v string) *VoiceAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type VoiceAsyncScanResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceAsyncScanResultsResponse) SetStatusCode(v int32) *VoiceAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceAsyncScanResultsResponse) SetBody(v *VoiceAsyncScanResultsResponseBody) *VoiceAsyncScanResultsResponse {
	s.Body = v
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

type VoiceCancelScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceCancelScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceCancelScanResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceCancelScanResponseBody) SetRequestId(v string) *VoiceCancelScanResponseBody {
	s.RequestId = &v
	return s
}

type VoiceCancelScanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceCancelScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceCancelScanResponse) SetStatusCode(v int32) *VoiceCancelScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceCancelScanResponse) SetBody(v *VoiceCancelScanResponseBody) *VoiceCancelScanResponse {
	s.Body = v
	return s
}

type VoiceIdentityCheckRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceIdentityCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityCheckRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityCheckRequest) SetClientInfo(v string) *VoiceIdentityCheckRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceIdentityCheckRequest) SetRegionId(v string) *VoiceIdentityCheckRequest {
	s.RegionId = &v
	return s
}

type VoiceIdentityCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceIdentityCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityCheckResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceIdentityCheckResponseBody) SetRequestId(v string) *VoiceIdentityCheckResponseBody {
	s.RequestId = &v
	return s
}

type VoiceIdentityCheckResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceIdentityCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceIdentityCheckResponse) SetStatusCode(v int32) *VoiceIdentityCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceIdentityCheckResponse) SetBody(v *VoiceIdentityCheckResponseBody) *VoiceIdentityCheckResponse {
	s.Body = v
	return s
}

type VoiceIdentityRegisterRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceIdentityRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityRegisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityRegisterRequest) SetClientInfo(v string) *VoiceIdentityRegisterRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceIdentityRegisterRequest) SetRegionId(v string) *VoiceIdentityRegisterRequest {
	s.RegionId = &v
	return s
}

type VoiceIdentityRegisterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceIdentityRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceIdentityRegisterResponseBody) SetRequestId(v string) *VoiceIdentityRegisterResponseBody {
	s.RequestId = &v
	return s
}

type VoiceIdentityRegisterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceIdentityRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceIdentityRegisterResponse) SetStatusCode(v int32) *VoiceIdentityRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceIdentityRegisterResponse) SetBody(v *VoiceIdentityRegisterResponseBody) *VoiceIdentityRegisterResponse {
	s.Body = v
	return s
}

type VoiceIdentityStartCheckRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceIdentityStartCheckRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartCheckRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartCheckRequest) SetClientInfo(v string) *VoiceIdentityStartCheckRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceIdentityStartCheckRequest) SetRegionId(v string) *VoiceIdentityStartCheckRequest {
	s.RegionId = &v
	return s
}

type VoiceIdentityStartCheckResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceIdentityStartCheckResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartCheckResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartCheckResponseBody) SetRequestId(v string) *VoiceIdentityStartCheckResponseBody {
	s.RequestId = &v
	return s
}

type VoiceIdentityStartCheckResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceIdentityStartCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceIdentityStartCheckResponse) SetStatusCode(v int32) *VoiceIdentityStartCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceIdentityStartCheckResponse) SetBody(v *VoiceIdentityStartCheckResponseBody) *VoiceIdentityStartCheckResponse {
	s.Body = v
	return s
}

type VoiceIdentityStartRegisterRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceIdentityStartRegisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartRegisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartRegisterRequest) SetClientInfo(v string) *VoiceIdentityStartRegisterRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceIdentityStartRegisterRequest) SetRegionId(v string) *VoiceIdentityStartRegisterRequest {
	s.RegionId = &v
	return s
}

type VoiceIdentityStartRegisterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceIdentityStartRegisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityStartRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceIdentityStartRegisterResponseBody) SetRequestId(v string) *VoiceIdentityStartRegisterResponseBody {
	s.RequestId = &v
	return s
}

type VoiceIdentityStartRegisterResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceIdentityStartRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceIdentityStartRegisterResponse) SetStatusCode(v int32) *VoiceIdentityStartRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceIdentityStartRegisterResponse) SetBody(v *VoiceIdentityStartRegisterResponseBody) *VoiceIdentityStartRegisterResponse {
	s.Body = v
	return s
}

type VoiceIdentityUnregisterRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceIdentityUnregisterRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityUnregisterRequest) GoString() string {
	return s.String()
}

func (s *VoiceIdentityUnregisterRequest) SetClientInfo(v string) *VoiceIdentityUnregisterRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceIdentityUnregisterRequest) SetRegionId(v string) *VoiceIdentityUnregisterRequest {
	s.RegionId = &v
	return s
}

type VoiceIdentityUnregisterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceIdentityUnregisterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceIdentityUnregisterResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceIdentityUnregisterResponseBody) SetRequestId(v string) *VoiceIdentityUnregisterResponseBody {
	s.RequestId = &v
	return s
}

type VoiceIdentityUnregisterResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceIdentityUnregisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceIdentityUnregisterResponse) SetStatusCode(v int32) *VoiceIdentityUnregisterResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceIdentityUnregisterResponse) SetBody(v *VoiceIdentityUnregisterResponseBody) *VoiceIdentityUnregisterResponse {
	s.Body = v
	return s
}

type VoiceSyncScanRequest struct {
	ClientInfo *string `json:"ClientInfo,omitempty" xml:"ClientInfo,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s VoiceSyncScanRequest) String() string {
	return tea.Prettify(s)
}

func (s VoiceSyncScanRequest) GoString() string {
	return s.String()
}

func (s *VoiceSyncScanRequest) SetClientInfo(v string) *VoiceSyncScanRequest {
	s.ClientInfo = &v
	return s
}

func (s *VoiceSyncScanRequest) SetRegionId(v string) *VoiceSyncScanRequest {
	s.RegionId = &v
	return s
}

type VoiceSyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VoiceSyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s VoiceSyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *VoiceSyncScanResponseBody) SetRequestId(v string) *VoiceSyncScanResponseBody {
	s.RequestId = &v
	return s
}

type VoiceSyncScanResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceSyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *VoiceSyncScanResponse) SetStatusCode(v int32) *VoiceSyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceSyncScanResponse) SetBody(v *VoiceSyncScanResponseBody) *VoiceSyncScanResponse {
	s.Body = v
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

type WebpageAsyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebpageAsyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanResponseBody) SetRequestId(v string) *WebpageAsyncScanResponseBody {
	s.RequestId = &v
	return s
}

type WebpageAsyncScanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebpageAsyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *WebpageAsyncScanResponse) SetStatusCode(v int32) *WebpageAsyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *WebpageAsyncScanResponse) SetBody(v *WebpageAsyncScanResponseBody) *WebpageAsyncScanResponse {
	s.Body = v
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

type WebpageAsyncScanResultsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebpageAsyncScanResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WebpageAsyncScanResultsResponseBody) GoString() string {
	return s.String()
}

func (s *WebpageAsyncScanResultsResponseBody) SetRequestId(v string) *WebpageAsyncScanResultsResponseBody {
	s.RequestId = &v
	return s
}

type WebpageAsyncScanResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebpageAsyncScanResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *WebpageAsyncScanResultsResponse) SetStatusCode(v int32) *WebpageAsyncScanResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *WebpageAsyncScanResultsResponse) SetBody(v *WebpageAsyncScanResultsResponseBody) *WebpageAsyncScanResultsResponse {
	s.Body = v
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

type WebpageSyncScanResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s WebpageSyncScanResponseBody) String() string {
	return tea.Prettify(s)
}

func (s WebpageSyncScanResponseBody) GoString() string {
	return s.String()
}

func (s *WebpageSyncScanResponseBody) SetRequestId(v string) *WebpageSyncScanResponseBody {
	s.RequestId = &v
	return s
}

type WebpageSyncScanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *WebpageSyncScanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *WebpageSyncScanResponse) SetStatusCode(v int32) *WebpageSyncScanResponse {
	s.StatusCode = &v
	return s
}

func (s *WebpageSyncScanResponse) SetBody(v *WebpageSyncScanResponseBody) *WebpageSyncScanResponse {
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

// Summary:
//
// 添加人脸
//
// @param request - AddFacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddFacesResponse
func (client *Client) AddFacesWithOptions(request *AddFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddFaces"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/face/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加人脸
//
// @param request - AddFacesRequest
//
// @return AddFacesResponse
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

// Summary:
//
// 添加分组
//
// @param request - AddGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupsResponse
func (client *Client) AddGroupsWithOptions(request *AddGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGroups"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person/groups/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加分组
//
// @param request - AddGroupsRequest
//
// @return AddGroupsResponse
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

// Summary:
//
// 添加个体
//
// @param request - AddPersonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddPersonResponse
func (client *Client) AddPersonWithOptions(request *AddPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AddPerson"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddPersonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加个体
//
// @param request - AddPersonRequest
//
// @return AddPersonResponse
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

// Summary:
//
// 添加相似图
//
// @param request - AddSimilarityImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSimilarityImageResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddSimilarityImage"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/image/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSimilarityImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加相似图
//
// @param request - AddSimilarityImageRequest
//
// @return AddSimilarityImageResponse
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

// Summary:
//
// 添加相似图库
//
// @param request - AddSimilarityLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddSimilarityLibraryResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddSimilarityLibrary"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/library/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddSimilarityLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加相似图库
//
// @param request - AddSimilarityLibraryRequest
//
// @return AddSimilarityLibraryResponse
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

// Summary:
//
// 添加视频Dna
//
// @param request - AddVideoDnaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVideoDnaResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddVideoDna"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/dna/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVideoDnaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加视频Dna
//
// @param request - AddVideoDnaRequest
//
// @return AddVideoDnaResponse
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

// Summary:
//
// 添加视频Dna分组
//
// @param request - AddVideoDnaGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddVideoDnaGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("AddVideoDnaGroup"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/dna/group/add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddVideoDnaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 添加视频Dna分组
//
// @param request - AddVideoDnaGroupRequest
//
// @return AddVideoDnaGroupResponse
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

// Summary:
//
// 删除人脸
//
// @param request - DeleteFacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFacesResponse
func (client *Client) DeleteFacesWithOptions(request *DeleteFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFaces"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/face/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除人脸
//
// @param request - DeleteFacesRequest
//
// @return DeleteFacesResponse
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

// Summary:
//
// 删除分组
//
// @param request - DeleteGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupsResponse
func (client *Client) DeleteGroupsWithOptions(request *DeleteGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroups"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person/groups/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除分组
//
// @param request - DeleteGroupsRequest
//
// @return DeleteGroupsResponse
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

// Summary:
//
// 删除个体
//
// @param request - DeletePersonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeletePersonResponse
func (client *Client) DeletePersonWithOptions(request *DeletePersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeletePersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeletePerson"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeletePersonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除个体
//
// @param request - DeletePersonRequest
//
// @return DeletePersonResponse
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

// Summary:
//
// 删除相似图
//
// @param request - DeleteSimilarityImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSimilarityImageResponse
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
	params := &openapi.Params{
		Action:      tea.String("DeleteSimilarityImage"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/image/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSimilarityImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除相似图
//
// @param request - DeleteSimilarityImageRequest
//
// @return DeleteSimilarityImageResponse
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

// Summary:
//
// 删除相似图库
//
// @param request - DeleteSimilarityLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteSimilarityLibraryResponse
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
	params := &openapi.Params{
		Action:      tea.String("DeleteSimilarityLibrary"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/library/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteSimilarityLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除相似图库
//
// @param request - DeleteSimilarityLibraryRequest
//
// @return DeleteSimilarityLibraryResponse
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

// Summary:
//
// 删除视频Dna
//
// @param request - DeleteVideoDnaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideoDnaResponse
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
	params := &openapi.Params{
		Action:      tea.String("DeleteVideoDna"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/dna/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVideoDnaResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除视频Dna
//
// @param request - DeleteVideoDnaRequest
//
// @return DeleteVideoDnaResponse
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

// Summary:
//
// 删除视频Dna分组
//
// @param request - DeleteVideoDnaGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVideoDnaGroupResponse
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
	params := &openapi.Params{
		Action:      tea.String("DeleteVideoDnaGroup"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/dna/group/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVideoDnaGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除视频Dna分组
//
// @param request - DeleteVideoDnaGroupRequest
//
// @return DeleteVideoDnaGroupResponse
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

// Summary:
//
// 人脸属性检测
//
// @param request - DetectFaceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DetectFaceResponse
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
	params := &openapi.Params{
		Action:      tea.String("DetectFace"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/face/detect"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DetectFaceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 人脸属性检测
//
// @param request - DetectFaceRequest
//
// @return DetectFaceResponse
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

// Summary:
//
// 文件异步检测
//
// @param request - FileAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("FileAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/file/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件异步检测
//
// @param request - FileAsyncScanRequest
//
// @return FileAsyncScanResponse
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

// Summary:
//
// 文件异步检测结果
//
// @param request - FileAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("FileAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/file/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件异步检测结果
//
// @param request - FileAsyncScanResultsRequest
//
// @return FileAsyncScanResultsResponse
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

// Summary:
//
// 文件检测新版本
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileAsyncScanV2Response
func (client *Client) FileAsyncScanV2WithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileAsyncScanV2Response, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("FileAsyncScanV2"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/file/asyncscanv2"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileAsyncScanV2Response{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文件检测新版本
//
// @return FileAsyncScanV2Response
func (client *Client) FileAsyncScanV2() (_result *FileAsyncScanV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileAsyncScanV2Response{}
	_body, _err := client.FileAsyncScanV2WithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取人脸列表
//
// @param request - GetFacesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFacesResponse
func (client *Client) GetFacesWithOptions(request *GetFacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFaces"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/faces"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFacesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取人脸列表
//
// @param request - GetFacesRequest
//
// @return GetFacesResponse
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

// Summary:
//
// 获取组列表
//
// @param request - GetGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupsResponse
func (client *Client) GetGroupsWithOptions(request *GetGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroups"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取组列表
//
// @param request - GetGroupsRequest
//
// @return GetGroupsResponse
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

// Summary:
//
// 获取单个个体
//
// @param request - GetPersonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPersonResponse
func (client *Client) GetPersonWithOptions(request *GetPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPerson"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPersonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个个体
//
// @param request - GetPersonRequest
//
// @return GetPersonResponse
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

// Summary:
//
// 获取个体列表
//
// @param request - GetPersonsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetPersonsResponse
func (client *Client) GetPersonsWithOptions(request *GetPersonsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPersonsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPersons"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/group/persons"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPersonsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取个体列表
//
// @param request - GetPersonsRequest
//
// @return GetPersonsResponse
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

// Summary:
//
// 获取相似图
//
// @param request - GetSimilarityImageRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSimilarityImageResponse
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
	params := &openapi.Params{
		Action:      tea.String("GetSimilarityImage"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/image/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimilarityImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图
//
// @param request - GetSimilarityImageRequest
//
// @return GetSimilarityImageResponse
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

// Summary:
//
// 获取相似图库
//
// @param request - GetSimilarityLibraryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetSimilarityLibraryResponse
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
	params := &openapi.Params{
		Action:      tea.String("GetSimilarityLibrary"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/library/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetSimilarityLibraryResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图库
//
// @param request - GetSimilarityLibraryRequest
//
// @return GetSimilarityLibraryResponse
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

// Summary:
//
// 图片人工异步审核
//
// @param request - ImageAsyncManualScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncManualScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageAsyncManualScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/manual/asyncScan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageAsyncManualScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片人工异步审核
//
// @param request - ImageAsyncManualScanRequest
//
// @return ImageAsyncManualScanResponse
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

// Summary:
//
// 图片人工异步审核结果
//
// @param request - ImageAsyncManualScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncManualScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageAsyncManualScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/manual/scan/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageAsyncManualScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片人工异步审核结果
//
// @param request - ImageAsyncManualScanResultsRequest
//
// @return ImageAsyncManualScanResultsResponse
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

// Summary:
//
// 图片异步检测
//
// @param request - ImageAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片异步检测
//
// @param request - ImageAsyncScanRequest
//
// @return ImageAsyncScanResponse
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

// Summary:
//
// 图片异步检测结果
//
// @param request - ImageAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片异步检测结果
//
// @param request - ImageAsyncScanResultsRequest
//
// @return ImageAsyncScanResultsResponse
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

// Summary:
//
// 图片检测反馈
//
// @param request - ImageScanFeedbackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageScanFeedbackResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageScanFeedback"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/feedback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageScanFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片检测反馈
//
// @param request - ImageScanFeedbackRequest
//
// @return ImageScanFeedbackResponse
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

// Summary:
//
// 图片同步检测
//
// @param request - ImageSyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImageSyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("ImageSyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/image/scan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImageSyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 图片同步检测
//
// @param request - ImageSyncScanRequest
//
// @return ImageSyncScanResponse
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

// Summary:
//
// 获取相似图
//
// @param request - ListSimilarityImagesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSimilarityImagesResponse
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
	params := &openapi.Params{
		Action:      tea.String("ListSimilarityImages"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/image/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSimilarityImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图
//
// @param request - ListSimilarityImagesRequest
//
// @return ListSimilarityImagesResponse
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

// Summary:
//
// 获取相似图库
//
// @param request - ListSimilarityLibrariesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSimilarityLibrariesResponse
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
	params := &openapi.Params{
		Action:      tea.String("ListSimilarityLibraries"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/similarity/library/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListSimilarityLibrariesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图库
//
// @param request - ListSimilarityLibrariesRequest
//
// @return ListSimilarityLibrariesResponse
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

// Summary:
//
// 直播流异步检测
//
// @param request - LiveStreamAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiveStreamAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("LiveStreamAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/livestream/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LiveStreamAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播流异步检测
//
// @param request - LiveStreamAsyncScanRequest
//
// @return LiveStreamAsyncScanResponse
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

// Summary:
//
// 直播流异步检测结果
//
// @param request - LiveStreamAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiveStreamAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("LiveStreamAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/livestream/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LiveStreamAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播流异步检测结果
//
// @param request - LiveStreamAsyncScanResultsRequest
//
// @return LiveStreamAsyncScanResultsResponse
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

// Summary:
//
// 直播流取消检测
//
// @param request - LiveStreamCancelScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LiveStreamCancelScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("LiveStreamCancelScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/livestream/cancelscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LiveStreamCancelScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 直播流取消检测
//
// @param request - LiveStreamCancelScanRequest
//
// @return LiveStreamCancelScanResponse
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

// Summary:
//
// 设置个体
//
// @param request - SetPersonRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SetPersonResponse
func (client *Client) SetPersonWithOptions(request *SetPersonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SetPersonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SetPerson"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/sface/person/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SetPersonResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 设置个体
//
// @param request - SetPersonRequest
//
// @return SetPersonResponse
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

// Summary:
//
// 文本异步人工审核
//
// @param request - TextAsyncManualScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextAsyncManualScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("TextAsyncManualScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/text/manual/asyncScan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TextAsyncManualScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本异步人工审核
//
// @param request - TextAsyncManualScanRequest
//
// @return TextAsyncManualScanResponse
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

// Summary:
//
// 文本异步人工审核结果
//
// @param request - TextAsyncManualScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextAsyncManualScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("TextAsyncManualScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/text/manual/scan/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TextAsyncManualScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本异步人工审核结果
//
// @param request - TextAsyncManualScanResultsRequest
//
// @return TextAsyncManualScanResultsResponse
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

// Summary:
//
// 文本结果反馈
//
// @param request - TextFeedbackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextFeedbackResponse
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
	params := &openapi.Params{
		Action:      tea.String("TextFeedback"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/text/feedback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TextFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 文本结果反馈
//
// @param request - TextFeedbackRequest
//
// @return TextFeedbackResponse
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

// @param request - TextScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TextScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("TextScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/text/scan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TextScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// @param request - TextScanRequest
//
// @return TextScanResponse
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

// Summary:
//
// 获取上传证书
//
// @param request - UploadCredentialsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UploadCredentialsResponse
func (client *Client) UploadCredentialsWithOptions(request *UploadCredentialsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCredentialsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadCredentials"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/credentials/uploadcredentials"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UploadCredentialsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取上传证书
//
// @param request - UploadCredentialsRequest
//
// @return UploadCredentialsResponse
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

// Summary:
//
// 视频异步人工审核
//
// @param request - VideoAsyncManualScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAsyncManualScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoAsyncManualScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/manual/asyncScan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoAsyncManualScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频异步人工审核
//
// @param request - VideoAsyncManualScanRequest
//
// @return VideoAsyncManualScanResponse
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

// Summary:
//
// 视频异步人工审核结果
//
// @param request - VideoAsyncManualScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAsyncManualScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoAsyncManualScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/manual/scan/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoAsyncManualScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频异步人工审核结果
//
// @param request - VideoAsyncManualScanResultsRequest
//
// @return VideoAsyncManualScanResultsResponse
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

// Summary:
//
// 视频异步检测
//
// @param request - VideoAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频异步检测
//
// @param request - VideoAsyncScanRequest
//
// @return VideoAsyncScanResponse
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

// Summary:
//
// 视频异步检测结果
//
// @param request - VideoAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频异步检测结果
//
// @param request - VideoAsyncScanResultsRequest
//
// @return VideoAsyncScanResultsResponse
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

// Summary:
//
// 视频取消检测
//
// @param request - VideoCancelScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoCancelScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoCancelScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/cancelscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoCancelScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频取消检测
//
// @param request - VideoCancelScanRequest
//
// @return VideoCancelScanResponse
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

// Summary:
//
// 视频结果反馈
//
// @param request - VideoFeedbackRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoFeedbackResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoFeedback"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/feedback"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoFeedbackResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频结果反馈
//
// @param request - VideoFeedbackRequest
//
// @return VideoFeedbackResponse
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

// Summary:
//
// 视频同步检测
//
// @param request - VideoSyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VideoSyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VideoSyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/video/syncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VideoSyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频同步检测
//
// @param request - VideoSyncScanRequest
//
// @return VideoSyncScanResponse
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

// Summary:
//
// 视频点播异步检测
//
// @param request - VodAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VodAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VodAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/vod/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VodAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频点播异步检测
//
// @param request - VodAsyncScanRequest
//
// @return VodAsyncScanResponse
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

// Summary:
//
// 视频点播异步检测结果
//
// @param request - VodAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VodAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("VodAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/vod/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VodAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 视频点播异步检测结果
//
// @param request - VodAsyncScanResultsRequest
//
// @return VodAsyncScanResultsResponse
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

// Summary:
//
// 语音异步人工审核
//
// @param request - VoiceAsyncManualScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceAsyncManualScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VoiceAsyncManualScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/manual/asyncScan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceAsyncManualScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音异步人工审核
//
// @param request - VoiceAsyncManualScanRequest
//
// @return VoiceAsyncManualScanResponse
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

// Summary:
//
// 语音异步人工审核结果
//
// @param request - VoiceAsyncManualScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceAsyncManualScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("VoiceAsyncManualScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/manual/scan/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceAsyncManualScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音异步人工审核结果
//
// @param request - VoiceAsyncManualScanResultsRequest
//
// @return VoiceAsyncManualScanResultsResponse
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

// Summary:
//
// 语音异步检测
//
// @param request - VoiceAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VoiceAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音异步检测
//
// @param request - VoiceAsyncScanRequest
//
// @return VoiceAsyncScanResponse
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

// Summary:
//
// 语音异步检测结果
//
// @param request - VoiceAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("VoiceAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音异步检测结果
//
// @param request - VoiceAsyncScanResultsRequest
//
// @return VoiceAsyncScanResultsResponse
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

// Summary:
//
// 语音取消检测
//
// @param request - VoiceCancelScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceCancelScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("VoiceCancelScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/cancelscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceCancelScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音取消检测
//
// @param request - VoiceCancelScanRequest
//
// @return VoiceCancelScanResponse
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

// Summary:
//
// 声纹比对
//
// @param request - VoiceIdentityCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceIdentityCheckResponse
func (client *Client) VoiceIdentityCheckWithOptions(request *VoiceIdentityCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceIdentityCheck"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/auth/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceIdentityCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 声纹比对
//
// @param request - VoiceIdentityCheckRequest
//
// @return VoiceIdentityCheckResponse
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

// Summary:
//
// 声纹注册
//
// @param request - VoiceIdentityRegisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceIdentityRegisterResponse
func (client *Client) VoiceIdentityRegisterWithOptions(request *VoiceIdentityRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceIdentityRegister"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/auth/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceIdentityRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 声纹注册
//
// @param request - VoiceIdentityRegisterRequest
//
// @return VoiceIdentityRegisterResponse
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

// Summary:
//
// 声纹开始比对
//
// @param request - VoiceIdentityStartCheckRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceIdentityStartCheckResponse
func (client *Client) VoiceIdentityStartCheckWithOptions(request *VoiceIdentityStartCheckRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityStartCheckResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceIdentityStartCheck"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/auth/start/check"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceIdentityStartCheckResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 声纹开始比对
//
// @param request - VoiceIdentityStartCheckRequest
//
// @return VoiceIdentityStartCheckResponse
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

// Summary:
//
// 声纹开始注册
//
// @param request - VoiceIdentityStartRegisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceIdentityStartRegisterResponse
func (client *Client) VoiceIdentityStartRegisterWithOptions(request *VoiceIdentityStartRegisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityStartRegisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceIdentityStartRegister"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/auth/start/register"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceIdentityStartRegisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 声纹开始注册
//
// @param request - VoiceIdentityStartRegisterRequest
//
// @return VoiceIdentityStartRegisterResponse
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

// Summary:
//
// 声纹注销
//
// @param request - VoiceIdentityUnregisterRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceIdentityUnregisterResponse
func (client *Client) VoiceIdentityUnregisterWithOptions(request *VoiceIdentityUnregisterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceIdentityUnregisterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceIdentityUnregister"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/auth/unregister"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceIdentityUnregisterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 声纹注销
//
// @param request - VoiceIdentityUnregisterRequest
//
// @return VoiceIdentityUnregisterResponse
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

// Summary:
//
// 语音同步检测
//
// @param request - VoiceSyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return VoiceSyncScanResponse
func (client *Client) VoiceSyncScanWithOptions(request *VoiceSyncScanRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VoiceSyncScanResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientInfo)) {
		query["ClientInfo"] = request.ClientInfo
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("VoiceSyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/voice/syncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &VoiceSyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 语音同步检测
//
// @param request - VoiceSyncScanRequest
//
// @return VoiceSyncScanResponse
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

// Summary:
//
// 站点异步检测
//
// @param request - WebpageAsyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebpageAsyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("WebpageAsyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/webpage/asyncscan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WebpageAsyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 站点异步检测
//
// @param request - WebpageAsyncScanRequest
//
// @return WebpageAsyncScanResponse
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

// Summary:
//
// 站点异步检测结果
//
// @param request - WebpageAsyncScanResultsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebpageAsyncScanResultsResponse
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
	params := &openapi.Params{
		Action:      tea.String("WebpageAsyncScanResults"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/webpage/results"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WebpageAsyncScanResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 站点异步检测结果
//
// @param request - WebpageAsyncScanResultsRequest
//
// @return WebpageAsyncScanResultsResponse
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

// Summary:
//
// 站点同步检测
//
// @param request - WebpageSyncScanRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return WebpageSyncScanResponse
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
	params := &openapi.Params{
		Action:      tea.String("WebpageSyncScan"),
		Version:     tea.String("2018-05-09"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/green/webpage/scan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &WebpageSyncScanResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 站点同步检测
//
// @param request - WebpageSyncScanRequest
//
// @return WebpageSyncScanResponse
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
