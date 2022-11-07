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

type CancelTaskRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CancelTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskRequest) GoString() string {
	return s.String()
}

func (s *CancelTaskRequest) SetOwnerAccount(v string) *CancelTaskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelTaskRequest) SetOwnerId(v int64) *CancelTaskRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetRegionId(v string) *CancelTaskRequest {
	s.RegionId = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerAccount(v string) *CancelTaskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelTaskRequest) SetResourceOwnerId(v int64) *CancelTaskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelTaskRequest) SetTaskId(v string) *CancelTaskRequest {
	s.TaskId = &v
	return s
}

type CancelTaskResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CancelTaskResponseBody) SetRequestId(v string) *CancelTaskResponseBody {
	s.RequestId = &v
	return s
}

type CancelTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CancelTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CancelTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelTaskResponse) GoString() string {
	return s.String()
}

func (s *CancelTaskResponse) SetHeaders(v map[string]*string) *CancelTaskResponse {
	s.Headers = v
	return s
}

func (s *CancelTaskResponse) SetStatusCode(v int32) *CancelTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelTaskResponse) SetBody(v *CancelTaskResponseBody) *CancelTaskResponse {
	s.Body = v
	return s
}

type CopyImageRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DestinationRegionId  *string `json:"DestinationRegionId,omitempty" xml:"DestinationRegionId,omitempty"`
	ImageId              *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName            *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CopyImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyImageRequest) GoString() string {
	return s.String()
}

func (s *CopyImageRequest) SetClientToken(v string) *CopyImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CopyImageRequest) SetDescription(v string) *CopyImageRequest {
	s.Description = &v
	return s
}

func (s *CopyImageRequest) SetDestinationRegionId(v string) *CopyImageRequest {
	s.DestinationRegionId = &v
	return s
}

func (s *CopyImageRequest) SetImageId(v string) *CopyImageRequest {
	s.ImageId = &v
	return s
}

func (s *CopyImageRequest) SetImageName(v string) *CopyImageRequest {
	s.ImageName = &v
	return s
}

func (s *CopyImageRequest) SetOwnerAccount(v string) *CopyImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CopyImageRequest) SetOwnerId(v int64) *CopyImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CopyImageRequest) SetRegionId(v string) *CopyImageRequest {
	s.RegionId = &v
	return s
}

func (s *CopyImageRequest) SetResourceOwnerAccount(v string) *CopyImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CopyImageRequest) SetResourceOwnerId(v int64) *CopyImageRequest {
	s.ResourceOwnerId = &v
	return s
}

type CopyImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyImageResponseBody) GoString() string {
	return s.String()
}

func (s *CopyImageResponseBody) SetImageId(v string) *CopyImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CopyImageResponseBody) SetRequestId(v string) *CopyImageResponseBody {
	s.RequestId = &v
	return s
}

type CopyImageResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CopyImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyImageResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyImageResponse) GoString() string {
	return s.String()
}

func (s *CopyImageResponse) SetHeaders(v map[string]*string) *CopyImageResponse {
	s.Headers = v
	return s
}

func (s *CopyImageResponse) SetStatusCode(v int32) *CopyImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyImageResponse) SetBody(v *CopyImageResponseBody) *CopyImageResponse {
	s.Body = v
	return s
}

type CreateImageRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageName            *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateImageRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s *CreateImageRequest) SetClientToken(v string) *CreateImageRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageRequest) SetDescription(v string) *CreateImageRequest {
	s.Description = &v
	return s
}

func (s *CreateImageRequest) SetImageName(v string) *CreateImageRequest {
	s.ImageName = &v
	return s
}

func (s *CreateImageRequest) SetInstanceId(v string) *CreateImageRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateImageRequest) SetOwnerAccount(v string) *CreateImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetOwnerId(v int64) *CreateImageRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageRequest) SetRegionId(v string) *CreateImageRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerAccount(v string) *CreateImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageRequest) SetResourceOwnerId(v int64) *CreateImageRequest {
	s.ResourceOwnerId = &v
	return s
}

type CreateImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateImageResponseBody) SetImageId(v string) *CreateImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *CreateImageResponseBody) SetRequestId(v string) *CreateImageResponseBody {
	s.RequestId = &v
	return s
}

type CreateImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateImageResponse) SetStatusCode(v int32) *CreateImageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}

type DeleteImagesRequest struct {
	Force                *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	ImageId              []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteImagesRequest) SetForce(v bool) *DeleteImagesRequest {
	s.Force = &v
	return s
}

func (s *DeleteImagesRequest) SetImageId(v []*string) *DeleteImagesRequest {
	s.ImageId = v
	return s
}

func (s *DeleteImagesRequest) SetOwnerAccount(v string) *DeleteImagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteImagesRequest) SetOwnerId(v int64) *DeleteImagesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteImagesRequest) SetRegionId(v string) *DeleteImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImagesRequest) SetResourceOwnerAccount(v string) *DeleteImagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteImagesRequest) SetResourceOwnerId(v int64) *DeleteImagesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteImagesResponseBody struct {
	ImageResponses *DeleteImagesResponseBodyImageResponses `json:"ImageResponses,omitempty" xml:"ImageResponses,omitempty" type:"Struct"`
	RequestId      *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBody) SetImageResponses(v *DeleteImagesResponseBodyImageResponses) *DeleteImagesResponseBody {
	s.ImageResponses = v
	return s
}

func (s *DeleteImagesResponseBody) SetRequestId(v string) *DeleteImagesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteImagesResponseBodyImageResponses struct {
	ImageResponse []*DeleteImagesResponseBodyImageResponsesImageResponse `json:"ImageResponse,omitempty" xml:"ImageResponse,omitempty" type:"Repeated"`
}

func (s DeleteImagesResponseBodyImageResponses) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBodyImageResponses) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBodyImageResponses) SetImageResponse(v []*DeleteImagesResponseBodyImageResponsesImageResponse) *DeleteImagesResponseBodyImageResponses {
	s.ImageResponse = v
	return s
}

type DeleteImagesResponseBodyImageResponsesImageResponse struct {
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s DeleteImagesResponseBodyImageResponsesImageResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponseBodyImageResponsesImageResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponseBodyImageResponsesImageResponse) SetCode(v string) *DeleteImagesResponseBodyImageResponsesImageResponse {
	s.Code = &v
	return s
}

func (s *DeleteImagesResponseBodyImageResponsesImageResponse) SetImageId(v string) *DeleteImagesResponseBodyImageResponsesImageResponse {
	s.ImageId = &v
	return s
}

func (s *DeleteImagesResponseBodyImageResponsesImageResponse) SetMessage(v string) *DeleteImagesResponseBodyImageResponsesImageResponse {
	s.Message = &v
	return s
}

type DeleteImagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) SetHeaders(v map[string]*string) *DeleteImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesResponse) SetStatusCode(v int32) *DeleteImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagesResponse) SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse {
	s.Body = v
	return s
}

type DeleteInstancesRequest struct {
	Force                *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) SetForce(v bool) *DeleteInstancesRequest {
	s.Force = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceId(v []*string) *DeleteInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerAccount(v string) *DeleteInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerId(v int64) *DeleteInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstancesRequest) SetRegionId(v string) *DeleteInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerAccount(v string) *DeleteInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerId(v int64) *DeleteInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponse) SetHeaders(v map[string]*string) *DeleteInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstancesResponse) SetStatusCode(v int32) *DeleteInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstancesResponse) SetBody(v *DeleteInstancesResponseBody) *DeleteInstancesResponse {
	s.Body = v
	return s
}

type DeleteKeyPairsRequest struct {
	KeyPairName          []*string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsRequest) SetKeyPairName(v []*string) *DeleteKeyPairsRequest {
	s.KeyPairName = v
	return s
}

func (s *DeleteKeyPairsRequest) SetOwnerAccount(v string) *DeleteKeyPairsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetOwnerId(v int64) *DeleteKeyPairsRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetRegionId(v string) *DeleteKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetResourceOwnerAccount(v string) *DeleteKeyPairsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteKeyPairsRequest) SetResourceOwnerId(v int64) *DeleteKeyPairsRequest {
	s.ResourceOwnerId = &v
	return s
}

type DeleteKeyPairsResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponseBody) SetRequestId(v string) *DeleteKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteKeyPairsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *DeleteKeyPairsResponse) SetHeaders(v map[string]*string) *DeleteKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *DeleteKeyPairsResponse) SetStatusCode(v int32) *DeleteKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteKeyPairsResponse) SetBody(v *DeleteKeyPairsResponseBody) *DeleteKeyPairsResponse {
	s.Body = v
	return s
}

type FetchFileRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OssBucket            *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssObject            *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Path                 *string `json:"Path,omitempty" xml:"Path,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s FetchFileRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchFileRequest) GoString() string {
	return s.String()
}

func (s *FetchFileRequest) SetInstanceId(v string) *FetchFileRequest {
	s.InstanceId = &v
	return s
}

func (s *FetchFileRequest) SetOssBucket(v string) *FetchFileRequest {
	s.OssBucket = &v
	return s
}

func (s *FetchFileRequest) SetOssObject(v string) *FetchFileRequest {
	s.OssObject = &v
	return s
}

func (s *FetchFileRequest) SetOwnerAccount(v string) *FetchFileRequest {
	s.OwnerAccount = &v
	return s
}

func (s *FetchFileRequest) SetOwnerId(v int64) *FetchFileRequest {
	s.OwnerId = &v
	return s
}

func (s *FetchFileRequest) SetPath(v string) *FetchFileRequest {
	s.Path = &v
	return s
}

func (s *FetchFileRequest) SetRegionId(v string) *FetchFileRequest {
	s.RegionId = &v
	return s
}

func (s *FetchFileRequest) SetResourceOwnerAccount(v string) *FetchFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *FetchFileRequest) SetResourceOwnerId(v int64) *FetchFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type FetchFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s FetchFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchFileResponseBody) GoString() string {
	return s.String()
}

func (s *FetchFileResponseBody) SetRequestId(v string) *FetchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchFileResponseBody) SetTaskId(v string) *FetchFileResponseBody {
	s.TaskId = &v
	return s
}

type FetchFileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FetchFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchFileResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchFileResponse) GoString() string {
	return s.String()
}

func (s *FetchFileResponse) SetHeaders(v map[string]*string) *FetchFileResponse {
	s.Headers = v
	return s
}

func (s *FetchFileResponse) SetStatusCode(v int32) *FetchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *FetchFileResponse) SetBody(v *FetchFileResponseBody) *FetchFileResponse {
	s.Body = v
	return s
}

type ImportImageRequest struct {
	ClientToken          *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Format               *string `json:"Format,omitempty" xml:"Format,omitempty"`
	ImageName            *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	OssBucket            *string `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssObject            *string `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Platform             *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ImportImageRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportImageRequest) GoString() string {
	return s.String()
}

func (s *ImportImageRequest) SetClientToken(v string) *ImportImageRequest {
	s.ClientToken = &v
	return s
}

func (s *ImportImageRequest) SetDescription(v string) *ImportImageRequest {
	s.Description = &v
	return s
}

func (s *ImportImageRequest) SetFormat(v string) *ImportImageRequest {
	s.Format = &v
	return s
}

func (s *ImportImageRequest) SetImageName(v string) *ImportImageRequest {
	s.ImageName = &v
	return s
}

func (s *ImportImageRequest) SetOssBucket(v string) *ImportImageRequest {
	s.OssBucket = &v
	return s
}

func (s *ImportImageRequest) SetOssObject(v string) *ImportImageRequest {
	s.OssObject = &v
	return s
}

func (s *ImportImageRequest) SetOwnerAccount(v string) *ImportImageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ImportImageRequest) SetOwnerId(v int64) *ImportImageRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportImageRequest) SetPlatform(v string) *ImportImageRequest {
	s.Platform = &v
	return s
}

func (s *ImportImageRequest) SetRegionId(v string) *ImportImageRequest {
	s.RegionId = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerAccount(v string) *ImportImageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportImageRequest) SetResourceOwnerId(v int64) *ImportImageRequest {
	s.ResourceOwnerId = &v
	return s
}

type ImportImageResponseBody struct {
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportImageResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportImageResponseBody) GoString() string {
	return s.String()
}

func (s *ImportImageResponseBody) SetImageId(v string) *ImportImageResponseBody {
	s.ImageId = &v
	return s
}

func (s *ImportImageResponseBody) SetRequestId(v string) *ImportImageResponseBody {
	s.RequestId = &v
	return s
}

type ImportImageResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportImageResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportImageResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportImageResponse) GoString() string {
	return s.String()
}

func (s *ImportImageResponse) SetHeaders(v map[string]*string) *ImportImageResponse {
	s.Headers = v
	return s
}

func (s *ImportImageResponse) SetStatusCode(v int32) *ImportImageResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportImageResponse) SetBody(v *ImportImageResponseBody) *ImportImageResponse {
	s.Body = v
	return s
}

type ImportKeyPairRequest struct {
	KeyPairName          *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PublicKeyBody        *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ImportKeyPairRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetOwnerAccount(v string) *ImportKeyPairRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ImportKeyPairRequest) SetOwnerId(v int64) *ImportKeyPairRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

func (s *ImportKeyPairRequest) SetRegionId(v string) *ImportKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *ImportKeyPairRequest) SetResourceOwnerAccount(v string) *ImportKeyPairRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportKeyPairRequest) SetResourceOwnerId(v int64) *ImportKeyPairRequest {
	s.ResourceOwnerId = &v
	return s
}

type ImportKeyPairResponseBody struct {
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportKeyPairResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponseBody) SetKeyPairFingerPrint(v string) *ImportKeyPairResponseBody {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetKeyPairName(v string) *ImportKeyPairResponseBody {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairResponseBody) SetRequestId(v string) *ImportKeyPairResponseBody {
	s.RequestId = &v
	return s
}

type ImportKeyPairResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ImportKeyPairResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ImportKeyPairResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportKeyPairResponse) GoString() string {
	return s.String()
}

func (s *ImportKeyPairResponse) SetHeaders(v map[string]*string) *ImportKeyPairResponse {
	s.Headers = v
	return s
}

func (s *ImportKeyPairResponse) SetStatusCode(v int32) *ImportKeyPairResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportKeyPairResponse) SetBody(v *ImportKeyPairResponseBody) *ImportKeyPairResponse {
	s.Body = v
	return s
}

type InstallApplicationRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OssBucket            *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssObject            *string   `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s InstallApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s InstallApplicationRequest) GoString() string {
	return s.String()
}

func (s *InstallApplicationRequest) SetInstanceId(v []*string) *InstallApplicationRequest {
	s.InstanceId = v
	return s
}

func (s *InstallApplicationRequest) SetOssBucket(v string) *InstallApplicationRequest {
	s.OssBucket = &v
	return s
}

func (s *InstallApplicationRequest) SetOssObject(v string) *InstallApplicationRequest {
	s.OssObject = &v
	return s
}

func (s *InstallApplicationRequest) SetOwnerAccount(v string) *InstallApplicationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *InstallApplicationRequest) SetOwnerId(v int64) *InstallApplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *InstallApplicationRequest) SetRegionId(v string) *InstallApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *InstallApplicationRequest) SetResourceOwnerAccount(v string) *InstallApplicationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InstallApplicationRequest) SetResourceOwnerId(v int64) *InstallApplicationRequest {
	s.ResourceOwnerId = &v
	return s
}

type InstallApplicationResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *InstallApplicationResponseBodyTaskId `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Struct"`
}

func (s InstallApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InstallApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *InstallApplicationResponseBody) SetRequestId(v string) *InstallApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallApplicationResponseBody) SetTaskId(v *InstallApplicationResponseBodyTaskId) *InstallApplicationResponseBody {
	s.TaskId = v
	return s
}

type InstallApplicationResponseBodyTaskId struct {
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s InstallApplicationResponseBodyTaskId) String() string {
	return tea.Prettify(s)
}

func (s InstallApplicationResponseBodyTaskId) GoString() string {
	return s.String()
}

func (s *InstallApplicationResponseBodyTaskId) SetTaskId(v []*string) *InstallApplicationResponseBodyTaskId {
	s.TaskId = v
	return s
}

type InstallApplicationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InstallApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InstallApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s InstallApplicationResponse) GoString() string {
	return s.String()
}

func (s *InstallApplicationResponse) SetHeaders(v map[string]*string) *InstallApplicationResponse {
	s.Headers = v
	return s
}

func (s *InstallApplicationResponse) SetStatusCode(v int32) *InstallApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallApplicationResponse) SetBody(v *InstallApplicationResponseBody) *InstallApplicationResponse {
	s.Body = v
	return s
}

type ListImageSharePermissionRequest struct {
	ImageId              *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListImageSharePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImageSharePermissionRequest) GoString() string {
	return s.String()
}

func (s *ListImageSharePermissionRequest) SetImageId(v string) *ListImageSharePermissionRequest {
	s.ImageId = &v
	return s
}

func (s *ListImageSharePermissionRequest) SetOwnerAccount(v string) *ListImageSharePermissionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListImageSharePermissionRequest) SetOwnerId(v int64) *ListImageSharePermissionRequest {
	s.OwnerId = &v
	return s
}

func (s *ListImageSharePermissionRequest) SetRegionId(v string) *ListImageSharePermissionRequest {
	s.RegionId = &v
	return s
}

func (s *ListImageSharePermissionRequest) SetResourceOwnerAccount(v string) *ListImageSharePermissionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListImageSharePermissionRequest) SetResourceOwnerId(v int64) *ListImageSharePermissionRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListImageSharePermissionResponseBody struct {
	Accounts  *ListImageSharePermissionResponseBodyAccounts `json:"Accounts,omitempty" xml:"Accounts,omitempty" type:"Struct"`
	RequestId *string                                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImageSharePermissionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImageSharePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageSharePermissionResponseBody) SetAccounts(v *ListImageSharePermissionResponseBodyAccounts) *ListImageSharePermissionResponseBody {
	s.Accounts = v
	return s
}

func (s *ListImageSharePermissionResponseBody) SetRequestId(v string) *ListImageSharePermissionResponseBody {
	s.RequestId = &v
	return s
}

type ListImageSharePermissionResponseBodyAccounts struct {
	Account []*ListImageSharePermissionResponseBodyAccountsAccount `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
}

func (s ListImageSharePermissionResponseBodyAccounts) String() string {
	return tea.Prettify(s)
}

func (s ListImageSharePermissionResponseBodyAccounts) GoString() string {
	return s.String()
}

func (s *ListImageSharePermissionResponseBodyAccounts) SetAccount(v []*ListImageSharePermissionResponseBodyAccountsAccount) *ListImageSharePermissionResponseBodyAccounts {
	s.Account = v
	return s
}

type ListImageSharePermissionResponseBodyAccountsAccount struct {
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
}

func (s ListImageSharePermissionResponseBodyAccountsAccount) String() string {
	return tea.Prettify(s)
}

func (s ListImageSharePermissionResponseBodyAccountsAccount) GoString() string {
	return s.String()
}

func (s *ListImageSharePermissionResponseBodyAccountsAccount) SetAliyunId(v string) *ListImageSharePermissionResponseBodyAccountsAccount {
	s.AliyunId = &v
	return s
}

type ListImageSharePermissionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImageSharePermissionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImageSharePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImageSharePermissionResponse) GoString() string {
	return s.String()
}

func (s *ListImageSharePermissionResponse) SetHeaders(v map[string]*string) *ListImageSharePermissionResponse {
	s.Headers = v
	return s
}

func (s *ListImageSharePermissionResponse) SetStatusCode(v int32) *ListImageSharePermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImageSharePermissionResponse) SetBody(v *ListImageSharePermissionResponseBody) *ListImageSharePermissionResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	ImageCategory        *string   `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageId              []*string `json:"ImageId,omitempty" xml:"ImageId,omitempty" type:"Repeated"`
	ImageName            *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	InstanceType         *string   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	MaxResults           *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	Status               *string   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageCategory(v string) *ListImagesRequest {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesRequest) SetImageId(v []*string) *ListImagesRequest {
	s.ImageId = v
	return s
}

func (s *ListImagesRequest) SetImageName(v string) *ListImagesRequest {
	s.ImageName = &v
	return s
}

func (s *ListImagesRequest) SetInstanceType(v string) *ListImagesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListImagesRequest) SetMaxResults(v int32) *ListImagesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListImagesRequest) SetNextToken(v string) *ListImagesRequest {
	s.NextToken = &v
	return s
}

func (s *ListImagesRequest) SetOwnerAccount(v string) *ListImagesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListImagesRequest) SetOwnerId(v int64) *ListImagesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListImagesRequest) SetRegionId(v string) *ListImagesRequest {
	s.RegionId = &v
	return s
}

func (s *ListImagesRequest) SetResourceOwnerAccount(v string) *ListImagesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListImagesRequest) SetResourceOwnerId(v int64) *ListImagesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListImagesRequest) SetStatus(v string) *ListImagesRequest {
	s.Status = &v
	return s
}

type ListImagesResponseBody struct {
	Images     *ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Struct"`
	MaxResults *int32                        `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v *ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetMaxResults(v int32) *ListImagesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListImagesResponseBody) SetNextToken(v string) *ListImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImagesResponseBody) SetRegionId(v string) *ListImagesResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int32) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesResponseBodyImages struct {
	Image []*ListImagesResponseBodyImagesImage `json:"Image,omitempty" xml:"Image,omitempty" type:"Repeated"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetImage(v []*ListImagesResponseBodyImagesImage) *ListImagesResponseBodyImages {
	s.Image = v
	return s
}

type ListImagesResponseBodyImagesImage struct {
	CreationTime  *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageCategory *string `json:"ImageCategory,omitempty" xml:"ImageCategory,omitempty"`
	ImageId       *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName     *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	IsSelfShared  *bool   `json:"IsSelfShared,omitempty" xml:"IsSelfShared,omitempty"`
	OSName        *string `json:"OSName,omitempty" xml:"OSName,omitempty"`
	OSNameEn      *string `json:"OSNameEn,omitempty" xml:"OSNameEn,omitempty"`
	OSType        *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	Platform      *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	Progress      *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	Status        *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Usage         *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListImagesResponseBodyImagesImage) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImagesImage) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImagesImage) SetCreationTime(v string) *ListImagesResponseBodyImagesImage {
	s.CreationTime = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetDescription(v string) *ListImagesResponseBodyImagesImage {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetImageCategory(v string) *ListImagesResponseBodyImagesImage {
	s.ImageCategory = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetImageId(v string) *ListImagesResponseBodyImagesImage {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetImageName(v string) *ListImagesResponseBodyImagesImage {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetIsSelfShared(v bool) *ListImagesResponseBodyImagesImage {
	s.IsSelfShared = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetOSName(v string) *ListImagesResponseBodyImagesImage {
	s.OSName = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetOSNameEn(v string) *ListImagesResponseBodyImagesImage {
	s.OSNameEn = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetOSType(v string) *ListImagesResponseBodyImagesImage {
	s.OSType = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetPlatform(v string) *ListImagesResponseBodyImagesImage {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetProgress(v string) *ListImagesResponseBodyImagesImage {
	s.Progress = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetStatus(v string) *ListImagesResponseBodyImagesImage {
	s.Status = &v
	return s
}

func (s *ListImagesResponseBodyImagesImage) SetUsage(v string) *ListImagesResponseBodyImagesImage {
	s.Usage = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListInstanceTypesRequest struct {
	InstanceType         []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
	InstanceTypeFamily   *string   `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListInstanceTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesRequest) SetInstanceType(v []*string) *ListInstanceTypesRequest {
	s.InstanceType = v
	return s
}

func (s *ListInstanceTypesRequest) SetInstanceTypeFamily(v string) *ListInstanceTypesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *ListInstanceTypesRequest) SetOwnerAccount(v string) *ListInstanceTypesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListInstanceTypesRequest) SetOwnerId(v int64) *ListInstanceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListInstanceTypesRequest) SetRegionId(v string) *ListInstanceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceTypesRequest) SetResourceOwnerAccount(v string) *ListInstanceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListInstanceTypesRequest) SetResourceOwnerId(v int64) *ListInstanceTypesRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListInstanceTypesResponseBody struct {
	InstanceTypes *ListInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBody) SetInstanceTypes(v *ListInstanceTypesResponseBodyInstanceTypes) *ListInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *ListInstanceTypesResponseBody) SetRequestId(v string) *ListInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListInstanceTypesResponseBodyInstanceTypes struct {
	InstanceType []*ListInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s ListInstanceTypesResponseBodyInstanceTypes) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*ListInstanceTypesResponseBodyInstanceTypesInstanceType) *ListInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

type ListInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	CpuCoreCount       *int32                                                             `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	DefaultResolution  *string                                                            `json:"DefaultResolution,omitempty" xml:"DefaultResolution,omitempty"`
	InstanceType       *string                                                            `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	InstanceTypeFamily *string                                                            `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	MemorySize         *string                                                            `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	Name               *string                                                            `json:"Name,omitempty" xml:"Name,omitempty"`
	NameEn             *string                                                            `json:"NameEn,omitempty" xml:"NameEn,omitempty"`
	Resolutions        *ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions `json:"Resolutions,omitempty" xml:"Resolutions,omitempty" type:"Struct"`
}

func (s ListInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetDefaultResolution(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.DefaultResolution = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceType(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceType = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeFamily(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetName(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.Name = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetNameEn(v string) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.NameEn = &v
	return s
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceType) SetResolutions(v *ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions) *ListInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.Resolutions = v
	return s
}

type ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions struct {
	Resolution []*string `json:"Resolution,omitempty" xml:"Resolution,omitempty" type:"Repeated"`
}

func (s ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions) SetResolution(v []*string) *ListInstanceTypesResponseBodyInstanceTypesInstanceTypeResolutions {
	s.Resolution = v
	return s
}

type ListInstanceTypesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceTypesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceTypesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceTypesResponse) SetHeaders(v map[string]*string) *ListInstanceTypesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceTypesResponse) SetStatusCode(v int32) *ListInstanceTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceTypesResponse) SetBody(v *ListInstanceTypesResponseBody) *ListInstanceTypesResponse {
	s.Body = v
	return s
}

type ListInstanceVncUrlRequest struct {
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListInstanceVncUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceVncUrlRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceVncUrlRequest) SetInstanceId(v string) *ListInstanceVncUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceVncUrlRequest) SetOwnerAccount(v string) *ListInstanceVncUrlRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListInstanceVncUrlRequest) SetOwnerId(v int64) *ListInstanceVncUrlRequest {
	s.OwnerId = &v
	return s
}

func (s *ListInstanceVncUrlRequest) SetRegionId(v string) *ListInstanceVncUrlRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstanceVncUrlRequest) SetResourceOwnerAccount(v string) *ListInstanceVncUrlRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListInstanceVncUrlRequest) SetResourceOwnerId(v int64) *ListInstanceVncUrlRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListInstanceVncUrlResponseBody struct {
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VncUrl      *string `json:"VncUrl,omitempty" xml:"VncUrl,omitempty"`
	WebRtcToken *string `json:"WebRtcToken,omitempty" xml:"WebRtcToken,omitempty"`
}

func (s ListInstanceVncUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceVncUrlResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceVncUrlResponseBody) SetRequestId(v string) *ListInstanceVncUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceVncUrlResponseBody) SetVncUrl(v string) *ListInstanceVncUrlResponseBody {
	s.VncUrl = &v
	return s
}

func (s *ListInstanceVncUrlResponseBody) SetWebRtcToken(v string) *ListInstanceVncUrlResponseBody {
	s.WebRtcToken = &v
	return s
}

type ListInstanceVncUrlResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstanceVncUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListInstanceVncUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s ListInstanceVncUrlResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceVncUrlResponse) SetHeaders(v map[string]*string) *ListInstanceVncUrlResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceVncUrlResponse) SetStatusCode(v int32) *ListInstanceVncUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceVncUrlResponse) SetBody(v *ListInstanceVncUrlResponseBody) *ListInstanceVncUrlResponse {
	s.Body = v
	return s
}

type ListInstancesRequest struct {
	ChargeType           *string                    `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ImageId              *string                    `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId           []*string                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	InstanceName         *string                    `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType         *string                    `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KeyPairName          *string                    `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	MaxResults           *int32                     `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string                    `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                    `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                     `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resolution           *string                    `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ResourceOwnerAccount *string                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                     `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	ShowWebRtcToken      *bool                      `json:"ShowWebRtcToken,omitempty" xml:"ShowWebRtcToken,omitempty"`
	Status               *string                    `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                  []*ListInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	ZoneId               *string                    `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) SetChargeType(v string) *ListInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesRequest) SetImageId(v string) *ListInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceId(v []*string) *ListInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *ListInstancesRequest) SetInstanceName(v string) *ListInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesRequest) SetInstanceType(v string) *ListInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesRequest) SetKeyPairName(v string) *ListInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *ListInstancesRequest) SetMaxResults(v int32) *ListInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesRequest) SetNextToken(v string) *ListInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListInstancesRequest) SetOwnerAccount(v string) *ListInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListInstancesRequest) SetOwnerId(v int64) *ListInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListInstancesRequest) SetRegionId(v string) *ListInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListInstancesRequest) SetResolution(v string) *ListInstancesRequest {
	s.Resolution = &v
	return s
}

func (s *ListInstancesRequest) SetResourceOwnerAccount(v string) *ListInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListInstancesRequest) SetResourceOwnerId(v int64) *ListInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListInstancesRequest) SetShowWebRtcToken(v bool) *ListInstancesRequest {
	s.ShowWebRtcToken = &v
	return s
}

func (s *ListInstancesRequest) SetStatus(v string) *ListInstancesRequest {
	s.Status = &v
	return s
}

func (s *ListInstancesRequest) SetTag(v []*ListInstancesRequestTag) *ListInstancesRequest {
	s.Tag = v
	return s
}

func (s *ListInstancesRequest) SetZoneId(v string) *ListInstancesRequest {
	s.ZoneId = &v
	return s
}

type ListInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *ListInstancesRequestTag) SetKey(v string) *ListInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *ListInstancesRequestTag) SetValue(v string) *ListInstancesRequestTag {
	s.Value = &v
	return s
}

type ListInstancesResponseBody struct {
	Instances  *ListInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	MaxResults *int32                              `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                             `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) SetInstances(v *ListInstancesResponseBodyInstances) *ListInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListInstancesResponseBody) SetMaxResults(v int32) *ListInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstancesResponseBody) SetNextToken(v string) *ListInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int32) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListInstancesResponseBodyInstances struct {
	Instance []*ListInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstances) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstances) SetInstance(v []*ListInstancesResponseBodyInstancesInstance) *ListInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

type ListInstancesResponseBodyInstancesInstance struct {
	AutoRenew       *bool                                                    `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType      *string                                                  `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	CreationTime    *string                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description     *string                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	EipAddress      *ListInstancesResponseBodyInstancesInstanceEipAddress    `json:"EipAddress,omitempty" xml:"EipAddress,omitempty" type:"Struct"`
	ExpiredTime     *string                                                  `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	ImageId         *string                                                  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId      *string                                                  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName    *string                                                  `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType    *string                                                  `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KeyPairName     *string                                                  `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OsName          *string                                                  `json:"OsName,omitempty" xml:"OsName,omitempty"`
	OsNameEn        *string                                                  `json:"OsNameEn,omitempty" xml:"OsNameEn,omitempty"`
	RegionId        *string                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resolution      *string                                                  `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	SecurityGroupId *string                                                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Status          *string                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            *ListInstancesResponseBodyInstancesInstanceTags          `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcAttributes   *ListInstancesResponseBodyInstancesInstanceVpcAttributes `json:"VpcAttributes,omitempty" xml:"VpcAttributes,omitempty" type:"Struct"`
	WebRtcToken     *string                                                  `json:"WebRtcToken,omitempty" xml:"WebRtcToken,omitempty"`
	ZoneId          *string                                                  `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstance) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstance) SetAutoRenew(v bool) *ListInstancesResponseBodyInstancesInstance {
	s.AutoRenew = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetChargeType(v string) *ListInstancesResponseBodyInstancesInstance {
	s.ChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetCreationTime(v string) *ListInstancesResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetDescription(v string) *ListInstancesResponseBodyInstancesInstance {
	s.Description = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetEipAddress(v *ListInstancesResponseBodyInstancesInstanceEipAddress) *ListInstancesResponseBodyInstancesInstance {
	s.EipAddress = v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetExpiredTime(v string) *ListInstancesResponseBodyInstancesInstance {
	s.ExpiredTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetImageId(v string) *ListInstancesResponseBodyInstancesInstance {
	s.ImageId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetInstanceType(v string) *ListInstancesResponseBodyInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetKeyPairName(v string) *ListInstancesResponseBodyInstancesInstance {
	s.KeyPairName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetOsName(v string) *ListInstancesResponseBodyInstancesInstance {
	s.OsName = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetOsNameEn(v string) *ListInstancesResponseBodyInstancesInstance {
	s.OsNameEn = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetRegionId(v string) *ListInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetResolution(v string) *ListInstancesResponseBodyInstancesInstance {
	s.Resolution = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetSecurityGroupId(v string) *ListInstancesResponseBodyInstancesInstance {
	s.SecurityGroupId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetStatus(v string) *ListInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetTags(v *ListInstancesResponseBodyInstancesInstanceTags) *ListInstancesResponseBodyInstancesInstance {
	s.Tags = v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetVpcAttributes(v *ListInstancesResponseBodyInstancesInstanceVpcAttributes) *ListInstancesResponseBodyInstancesInstance {
	s.VpcAttributes = v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetWebRtcToken(v string) *ListInstancesResponseBodyInstancesInstance {
	s.WebRtcToken = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstance) SetZoneId(v string) *ListInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

type ListInstancesResponseBodyInstancesInstanceEipAddress struct {
	AllocationId       *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	Bandwidth          *int32  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	IpAddress          *string `json:"IpAddress,omitempty" xml:"IpAddress,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceEipAddress) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceEipAddress) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceEipAddress) SetAllocationId(v string) *ListInstancesResponseBodyInstancesInstanceEipAddress {
	s.AllocationId = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceEipAddress) SetBandwidth(v int32) *ListInstancesResponseBodyInstancesInstanceEipAddress {
	s.Bandwidth = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceEipAddress) SetInternetChargeType(v string) *ListInstancesResponseBodyInstancesInstanceEipAddress {
	s.InternetChargeType = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceEipAddress) SetIpAddress(v string) *ListInstancesResponseBodyInstancesInstanceEipAddress {
	s.IpAddress = &v
	return s
}

type ListInstancesResponseBodyInstancesInstanceTags struct {
	Tag []*ListInstancesResponseBodyInstancesInstanceTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListInstancesResponseBodyInstancesInstanceTags) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceTags) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceTags) SetTag(v []*ListInstancesResponseBodyInstancesInstanceTagsTag) *ListInstancesResponseBodyInstancesInstanceTags {
	s.Tag = v
	return s
}

type ListInstancesResponseBodyInstancesInstanceTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceTagsTag) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceTagsTag) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceTagsTag) SetKey(v string) *ListInstancesResponseBodyInstancesInstanceTagsTag {
	s.Key = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceTagsTag) SetValue(v string) *ListInstancesResponseBodyInstancesInstanceTagsTag {
	s.Value = &v
	return s
}

type ListInstancesResponseBodyInstancesInstanceVpcAttributes struct {
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	VSwitchId        *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s ListInstancesResponseBodyInstancesInstanceVpcAttributes) String() string {
	return tea.Prettify(s)
}

func (s ListInstancesResponseBodyInstancesInstanceVpcAttributes) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstancesInstanceVpcAttributes) SetPrivateIpAddress(v string) *ListInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.PrivateIpAddress = &v
	return s
}

func (s *ListInstancesResponseBodyInstancesInstanceVpcAttributes) SetVSwitchId(v string) *ListInstancesResponseBodyInstancesInstanceVpcAttributes {
	s.VSwitchId = &v
	return s
}

type ListInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListInstancesResponse) SetStatusCode(v int32) *ListInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesResponse) SetBody(v *ListInstancesResponseBody) *ListInstancesResponse {
	s.Body = v
	return s
}

type ListKeyPairsRequest struct {
	KeyPairFingerPrint   *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairName          *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	MaxResults           *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListKeyPairsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsRequest) GoString() string {
	return s.String()
}

func (s *ListKeyPairsRequest) SetKeyPairFingerPrint(v string) *ListKeyPairsRequest {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *ListKeyPairsRequest) SetKeyPairName(v string) *ListKeyPairsRequest {
	s.KeyPairName = &v
	return s
}

func (s *ListKeyPairsRequest) SetMaxResults(v int32) *ListKeyPairsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListKeyPairsRequest) SetNextToken(v string) *ListKeyPairsRequest {
	s.NextToken = &v
	return s
}

func (s *ListKeyPairsRequest) SetOwnerAccount(v string) *ListKeyPairsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListKeyPairsRequest) SetOwnerId(v int64) *ListKeyPairsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListKeyPairsRequest) SetRegionId(v string) *ListKeyPairsRequest {
	s.RegionId = &v
	return s
}

func (s *ListKeyPairsRequest) SetResourceOwnerAccount(v string) *ListKeyPairsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListKeyPairsRequest) SetResourceOwnerId(v int64) *ListKeyPairsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListKeyPairsResponseBody struct {
	KeyPairs   *ListKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Struct"`
	MaxResults *int32                            `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                           `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeyPairsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBody) SetKeyPairs(v *ListKeyPairsResponseBodyKeyPairs) *ListKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *ListKeyPairsResponseBody) SetMaxResults(v int32) *ListKeyPairsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetNextToken(v string) *ListKeyPairsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetRequestId(v string) *ListKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetTotalCount(v int32) *ListKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

type ListKeyPairsResponseBodyKeyPairs struct {
	KeyPair []*ListKeyPairsResponseBodyKeyPairsKeyPair `json:"KeyPair,omitempty" xml:"KeyPair,omitempty" type:"Repeated"`
}

func (s ListKeyPairsResponseBodyKeyPairs) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetKeyPair(v []*ListKeyPairsResponseBodyKeyPairsKeyPair) *ListKeyPairsResponseBodyKeyPairs {
	s.KeyPair = v
	return s
}

type ListKeyPairsResponseBodyKeyPairsKeyPair struct {
	CreationTime       *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	KeyPairFingerPrint *string `json:"KeyPairFingerPrint,omitempty" xml:"KeyPairFingerPrint,omitempty"`
	KeyPairName        *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
}

func (s ListKeyPairsResponseBodyKeyPairsKeyPair) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponseBodyKeyPairsKeyPair) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBodyKeyPairsKeyPair) SetCreationTime(v string) *ListKeyPairsResponseBodyKeyPairsKeyPair {
	s.CreationTime = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairFingerPrint(v string) *ListKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairFingerPrint = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairsKeyPair) SetKeyPairName(v string) *ListKeyPairsResponseBodyKeyPairsKeyPair {
	s.KeyPairName = &v
	return s
}

type ListKeyPairsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListKeyPairsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponse) SetHeaders(v map[string]*string) *ListKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *ListKeyPairsResponse) SetStatusCode(v int32) *ListKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeyPairsResponse) SetBody(v *ListKeyPairsResponseBody) *ListKeyPairsResponse {
	s.Body = v
	return s
}

type ListRegionsRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsRequest) GoString() string {
	return s.String()
}

func (s *ListRegionsRequest) SetOwnerAccount(v string) *ListRegionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListRegionsRequest) SetOwnerId(v int64) *ListRegionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ListRegionsRequest) SetRegionId(v string) *ListRegionsRequest {
	s.RegionId = &v
	return s
}

func (s *ListRegionsRequest) SetResourceOwnerAccount(v string) *ListRegionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListRegionsRequest) SetResourceOwnerId(v int64) *ListRegionsRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListRegionsResponseBody struct {
	Regions   *ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) SetRegions(v *ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

type ListRegionsResponseBodyRegions struct {
	Region []*ListRegionsResponseBodyRegionsRegion `json:"Region,omitempty" xml:"Region,omitempty" type:"Repeated"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) SetRegion(v []*ListRegionsResponseBodyRegionsRegion) *ListRegionsResponseBodyRegions {
	s.Region = v
	return s
}

type ListRegionsResponseBodyRegionsRegion struct {
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	RegionId       *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegionsRegion) String() string {
	return tea.Prettify(s)
}

func (s ListRegionsResponseBodyRegionsRegion) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegionsRegion) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegionsRegion {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegionsRegion) SetRegionId(v string) *ListRegionsResponseBodyRegionsRegion {
	s.RegionId = &v
	return s
}

type ListRegionsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListRegionsResponse) SetStatusCode(v int32) *ListRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRegionsResponse) SetBody(v *ListRegionsResponseBody) *ListRegionsResponse {
	s.Body = v
	return s
}

type ListTagKeysRequest struct {
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagKeysRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysRequest) GoString() string {
	return s.String()
}

func (s *ListTagKeysRequest) SetNextToken(v string) *ListTagKeysRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerAccount(v string) *ListTagKeysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetOwnerId(v int64) *ListTagKeysRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagKeysRequest) SetRegionId(v string) *ListTagKeysRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceOwnerAccount(v string) *ListTagKeysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagKeysRequest) SetResourceType(v string) *ListTagKeysRequest {
	s.ResourceType = &v
	return s
}

type ListTagKeysResponseBody struct {
	Keys       *ListTagKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Struct"`
	MaxResults *int32                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTagKeysResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) SetKeys(v *ListTagKeysResponseBodyKeys) *ListTagKeysResponseBody {
	s.Keys = v
	return s
}

func (s *ListTagKeysResponseBody) SetMaxResults(v int32) *ListTagKeysResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

type ListTagKeysResponseBodyKeys struct {
	Key []*string `json:"Key,omitempty" xml:"Key,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBodyKeys) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyKeys) SetKey(v []*string) *ListTagKeysResponseBodyKeys {
	s.Key = v
	return s
}

type ListTagKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagKeysResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagKeysResponse) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponse) SetHeaders(v map[string]*string) *ListTagKeysResponse {
	s.Headers = v
	return s
}

func (s *ListTagKeysResponse) SetStatusCode(v int32) *ListTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagKeysResponse) SetBody(v *ListTagKeysResponseBody) *ListTagKeysResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	NextToken            *string                       `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string                       `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                        `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                       `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                       `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                       `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerAccount(v string) *ListTagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetOwnerId(v int64) *ListTagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceOwnerAccount(v string) *ListTagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
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
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey       *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue     *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListTagValuesRequest struct {
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	NextToken            *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) SetKey(v string) *ListTagValuesRequest {
	s.Key = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerAccount(v string) *ListTagValuesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetOwnerId(v int64) *ListTagValuesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTagValuesRequest) SetRegionId(v string) *ListTagValuesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceOwnerAccount(v string) *ListTagValuesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTagValuesRequest) SetResourceType(v string) *ListTagValuesRequest {
	s.ResourceType = &v
	return s
}

type ListTagValuesResponseBody struct {
	MaxResults *int32                           `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                          `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Values     *ListTagValuesResponseBodyValues `json:"Values,omitempty" xml:"Values,omitempty" type:"Struct"`
}

func (s ListTagValuesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) SetMaxResults(v int32) *ListTagValuesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetValues(v *ListTagValuesResponseBodyValues) *ListTagValuesResponseBody {
	s.Values = v
	return s
}

type ListTagValuesResponseBodyValues struct {
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBodyValues) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponseBodyValues) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyValues) SetValue(v []*string) *ListTagValuesResponseBodyValues {
	s.Value = v
	return s
}

type ListTagValuesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagValuesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagValuesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagValuesResponse) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponse) SetHeaders(v map[string]*string) *ListTagValuesResponse {
	s.Headers = v
	return s
}

func (s *ListTagValuesResponse) SetStatusCode(v int32) *ListTagValuesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagValuesResponse) SetBody(v *ListTagValuesResponseBody) *ListTagValuesResponse {
	s.Body = v
	return s
}

type ListTasksRequest struct {
	InstanceId           *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxResults           *int32    `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken            *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	TaskId               []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
	TaskStatus           *string   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType             *string   `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTasksRequest) GoString() string {
	return s.String()
}

func (s *ListTasksRequest) SetInstanceId(v string) *ListTasksRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTasksRequest) SetMaxResults(v int32) *ListTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTasksRequest) SetNextToken(v string) *ListTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListTasksRequest) SetOwnerAccount(v string) *ListTasksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListTasksRequest) SetOwnerId(v int64) *ListTasksRequest {
	s.OwnerId = &v
	return s
}

func (s *ListTasksRequest) SetRegionId(v string) *ListTasksRequest {
	s.RegionId = &v
	return s
}

func (s *ListTasksRequest) SetResourceOwnerAccount(v string) *ListTasksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListTasksRequest) SetResourceOwnerId(v int64) *ListTasksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListTasksRequest) SetTaskId(v []*string) *ListTasksRequest {
	s.TaskId = v
	return s
}

func (s *ListTasksRequest) SetTaskStatus(v string) *ListTasksRequest {
	s.TaskStatus = &v
	return s
}

func (s *ListTasksRequest) SetTaskType(v string) *ListTasksRequest {
	s.TaskType = &v
	return s
}

type ListTasksResponseBody struct {
	MaxResults *int32                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RegionId   *string                     `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks      *ListTasksResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Struct"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTasksResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBody) SetMaxResults(v int32) *ListTasksResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListTasksResponseBody) SetNextToken(v string) *ListTasksResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTasksResponseBody) SetRegionId(v string) *ListTasksResponseBody {
	s.RegionId = &v
	return s
}

func (s *ListTasksResponseBody) SetRequestId(v string) *ListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTasksResponseBody) SetTasks(v *ListTasksResponseBodyTasks) *ListTasksResponseBody {
	s.Tasks = v
	return s
}

func (s *ListTasksResponseBody) SetTotalCount(v int32) *ListTasksResponseBody {
	s.TotalCount = &v
	return s
}

type ListTasksResponseBodyTasks struct {
	Task []*ListTasksResponseBodyTasksTask `json:"Task,omitempty" xml:"Task,omitempty" type:"Repeated"`
}

func (s ListTasksResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasks) SetTask(v []*ListTasksResponseBodyTasksTask) *ListTasksResponseBodyTasks {
	s.Task = v
	return s
}

type ListTasksResponseBodyTasksTask struct {
	CreateTime   *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExecuteMsg   *string `json:"ExecuteMsg,omitempty" xml:"ExecuteMsg,omitempty"`
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Progress     *string `json:"Progress,omitempty" xml:"Progress,omitempty"`
	TaskId       *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus   *string `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	TaskType     *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s ListTasksResponseBodyTasksTask) String() string {
	return tea.Prettify(s)
}

func (s ListTasksResponseBodyTasksTask) GoString() string {
	return s.String()
}

func (s *ListTasksResponseBodyTasksTask) SetCreateTime(v string) *ListTasksResponseBodyTasksTask {
	s.CreateTime = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetExecuteMsg(v string) *ListTasksResponseBodyTasksTask {
	s.ExecuteMsg = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetFinishedTime(v string) *ListTasksResponseBodyTasksTask {
	s.FinishedTime = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetInstanceId(v string) *ListTasksResponseBodyTasksTask {
	s.InstanceId = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetProgress(v string) *ListTasksResponseBodyTasksTask {
	s.Progress = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetTaskId(v string) *ListTasksResponseBodyTasksTask {
	s.TaskId = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetTaskStatus(v string) *ListTasksResponseBodyTasksTask {
	s.TaskStatus = &v
	return s
}

func (s *ListTasksResponseBodyTasksTask) SetTaskType(v string) *ListTasksResponseBodyTasksTask {
	s.TaskType = &v
	return s
}

type ListTasksResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTasksResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTasksResponse) SetStatusCode(v int32) *ListTasksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTasksResponse) SetBody(v *ListTasksResponseBody) *ListTasksResponse {
	s.Body = v
	return s
}

type ListZonesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ListZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListZonesRequest) GoString() string {
	return s.String()
}

func (s *ListZonesRequest) SetOwnerAccount(v string) *ListZonesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListZonesRequest) SetOwnerId(v int64) *ListZonesRequest {
	s.OwnerId = &v
	return s
}

func (s *ListZonesRequest) SetRegionId(v string) *ListZonesRequest {
	s.RegionId = &v
	return s
}

func (s *ListZonesRequest) SetResourceOwnerAccount(v string) *ListZonesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListZonesRequest) SetResourceOwnerId(v int64) *ListZonesRequest {
	s.ResourceOwnerId = &v
	return s
}

type ListZonesResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Zones     *ListZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Struct"`
}

func (s ListZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBody) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBody) SetRequestId(v string) *ListZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListZonesResponseBody) SetZones(v *ListZonesResponseBodyZones) *ListZonesResponseBody {
	s.Zones = v
	return s
}

type ListZonesResponseBodyZones struct {
	Zone []*ListZonesResponseBodyZonesZone `json:"Zone,omitempty" xml:"Zone,omitempty" type:"Repeated"`
}

func (s ListZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyZones) SetZone(v []*ListZonesResponseBodyZonesZone) *ListZonesResponseBodyZones {
	s.Zone = v
	return s
}

type ListZonesResponseBodyZonesZone struct {
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListZonesResponseBodyZonesZone) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponseBodyZonesZone) GoString() string {
	return s.String()
}

func (s *ListZonesResponseBodyZonesZone) SetZoneId(v string) *ListZonesResponseBodyZonesZone {
	s.ZoneId = &v
	return s
}

type ListZonesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListZonesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListZonesResponse) GoString() string {
	return s.String()
}

func (s *ListZonesResponse) SetHeaders(v map[string]*string) *ListZonesResponse {
	s.Headers = v
	return s
}

func (s *ListZonesResponse) SetStatusCode(v int32) *ListZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListZonesResponse) SetBody(v *ListZonesResponseBody) *ListZonesResponse {
	s.Body = v
	return s
}

type RebootInstancesRequest struct {
	Force                *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RebootInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesRequest) SetForce(v bool) *RebootInstancesRequest {
	s.Force = &v
	return s
}

func (s *RebootInstancesRequest) SetInstanceId(v []*string) *RebootInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *RebootInstancesRequest) SetOwnerAccount(v string) *RebootInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootInstancesRequest) SetOwnerId(v int64) *RebootInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootInstancesRequest) SetRegionId(v string) *RebootInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebootInstancesRequest) SetResourceOwnerAccount(v string) *RebootInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootInstancesRequest) SetResourceOwnerId(v int64) *RebootInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type RebootInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponseBody) SetRequestId(v string) *RebootInstancesResponseBody {
	s.RequestId = &v
	return s
}

type RebootInstancesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RebootInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RebootInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootInstancesResponse) GoString() string {
	return s.String()
}

func (s *RebootInstancesResponse) SetHeaders(v map[string]*string) *RebootInstancesResponse {
	s.Headers = v
	return s
}

func (s *RebootInstancesResponse) SetStatusCode(v int32) *RebootInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootInstancesResponse) SetBody(v *RebootInstancesResponseBody) *RebootInstancesResponse {
	s.Body = v
	return s
}

type RenewInstancesRequest struct {
	AutoPay              *bool     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	ClientToken          *string   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period               *int32    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RenewInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RenewInstancesRequest) GoString() string {
	return s.String()
}

func (s *RenewInstancesRequest) SetAutoPay(v bool) *RenewInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RenewInstancesRequest) SetClientToken(v string) *RenewInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RenewInstancesRequest) SetInstanceId(v []*string) *RenewInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *RenewInstancesRequest) SetOwnerAccount(v string) *RenewInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RenewInstancesRequest) SetOwnerId(v int64) *RenewInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RenewInstancesRequest) SetPeriod(v int32) *RenewInstancesRequest {
	s.Period = &v
	return s
}

func (s *RenewInstancesRequest) SetPeriodUnit(v string) *RenewInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RenewInstancesRequest) SetRegionId(v string) *RenewInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RenewInstancesRequest) SetResourceOwnerAccount(v string) *RenewInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RenewInstancesRequest) SetResourceOwnerId(v int64) *RenewInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type RenewInstancesResponseBody struct {
	InstanceIds *RenewInstancesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	OrderId     *string                                `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenewInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewInstancesResponseBody) SetInstanceIds(v *RenewInstancesResponseBodyInstanceIds) *RenewInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *RenewInstancesResponseBody) SetOrderId(v string) *RenewInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewInstancesResponseBody) SetRequestId(v string) *RenewInstancesResponseBody {
	s.RequestId = &v
	return s
}

type RenewInstancesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s RenewInstancesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s RenewInstancesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *RenewInstancesResponseBodyInstanceIds) SetInstanceId(v []*string) *RenewInstancesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type RenewInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RenewInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenewInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RenewInstancesResponse) GoString() string {
	return s.String()
}

func (s *RenewInstancesResponse) SetHeaders(v map[string]*string) *RenewInstancesResponse {
	s.Headers = v
	return s
}

func (s *RenewInstancesResponse) SetStatusCode(v int32) *RenewInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewInstancesResponse) SetBody(v *RenewInstancesResponseBody) *RenewInstancesResponse {
	s.Body = v
	return s
}

type ResetInstancesRequest struct {
	ImageId              *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ResetInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ResetInstancesRequest) GoString() string {
	return s.String()
}

func (s *ResetInstancesRequest) SetImageId(v string) *ResetInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *ResetInstancesRequest) SetInstanceId(v []*string) *ResetInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *ResetInstancesRequest) SetOwnerAccount(v string) *ResetInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ResetInstancesRequest) SetOwnerId(v int64) *ResetInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ResetInstancesRequest) SetRegionId(v string) *ResetInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ResetInstancesRequest) SetResourceOwnerAccount(v string) *ResetInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ResetInstancesRequest) SetResourceOwnerId(v int64) *ResetInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type ResetInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ResetInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ResetInstancesResponseBody) SetRequestId(v string) *ResetInstancesResponseBody {
	s.RequestId = &v
	return s
}

type ResetInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ResetInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ResetInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ResetInstancesResponse) GoString() string {
	return s.String()
}

func (s *ResetInstancesResponse) SetHeaders(v map[string]*string) *ResetInstancesResponse {
	s.Headers = v
	return s
}

func (s *ResetInstancesResponse) SetStatusCode(v int32) *ResetInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstancesResponse) SetBody(v *ResetInstancesResponseBody) *ResetInstancesResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	Command              *string   `json:"Command,omitempty" xml:"Command,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetCommand(v string) *RunCommandRequest {
	s.Command = &v
	return s
}

func (s *RunCommandRequest) SetInstanceId(v []*string) *RunCommandRequest {
	s.InstanceId = v
	return s
}

func (s *RunCommandRequest) SetOwnerAccount(v string) *RunCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunCommandRequest) SetOwnerId(v int64) *RunCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *RunCommandRequest) SetRegionId(v string) *RunCommandRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandRequest) SetResourceOwnerAccount(v string) *RunCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunCommandRequest) SetResourceOwnerId(v int64) *RunCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

type RunCommandResponseBody struct {
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *RunCommandResponseBodyTaskId `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Struct"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponseBody) SetTaskId(v *RunCommandResponseBodyTaskId) *RunCommandResponseBody {
	s.TaskId = v
	return s
}

type RunCommandResponseBodyTaskId struct {
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s RunCommandResponseBodyTaskId) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBodyTaskId) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBodyTaskId) SetTaskId(v []*string) *RunCommandResponseBodyTaskId {
	s.TaskId = v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type RunInstancesRequest struct {
	Amount               *int32                    `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoPay              *bool                     `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	AutoRenew            *bool                     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType           *string                   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ClientToken          *string                   `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description          *string                   `json:"Description,omitempty" xml:"Description,omitempty"`
	EipBandwidth         *int32                    `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	ImageId              *string                   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceName         *string                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceType         *string                   `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	KeyPairName          *string                   `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period               *int64                    `json:"Period,omitempty" xml:"Period,omitempty"`
	PeriodUnit           *string                   `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	PrivateIpAddress     *string                   `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resolution           *string                   `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	SecurityGroupId      *string                   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Tag                  []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VSwitchId            *string                   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s RunInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s RunInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesRequest) SetAmount(v int32) *RunInstancesRequest {
	s.Amount = &v
	return s
}

func (s *RunInstancesRequest) SetAutoPay(v bool) *RunInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenew(v bool) *RunInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunInstancesRequest) SetChargeType(v string) *RunInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetClientToken(v string) *RunInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RunInstancesRequest) SetDescription(v string) *RunInstancesRequest {
	s.Description = &v
	return s
}

func (s *RunInstancesRequest) SetEipBandwidth(v int32) *RunInstancesRequest {
	s.EipBandwidth = &v
	return s
}

func (s *RunInstancesRequest) SetImageId(v string) *RunInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceName(v string) *RunInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceType(v string) *RunInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequest) SetKeyPairName(v string) *RunInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunInstancesRequest) SetOwnerAccount(v string) *RunInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetOwnerId(v int64) *RunInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RunInstancesRequest) SetPeriod(v int64) *RunInstancesRequest {
	s.Period = &v
	return s
}

func (s *RunInstancesRequest) SetPeriodUnit(v string) *RunInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunInstancesRequest) SetPrivateIpAddress(v string) *RunInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunInstancesRequest) SetRegionId(v string) *RunInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RunInstancesRequest) SetResolution(v string) *RunInstancesRequest {
	s.Resolution = &v
	return s
}

func (s *RunInstancesRequest) SetResourceOwnerAccount(v string) *RunInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityGroupId(v string) *RunInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunInstancesRequest) SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest {
	s.Tag = v
	return s
}

func (s *RunInstancesRequest) SetVSwitchId(v string) *RunInstancesRequest {
	s.VSwitchId = &v
	return s
}

type RunInstancesRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunInstancesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s RunInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestTag) SetKey(v string) *RunInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *RunInstancesRequestTag) SetValue(v string) *RunInstancesRequestTag {
	s.Value = &v
	return s
}

type RunInstancesResponseBody struct {
	InstanceIds *RunInstancesResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	OrderId     *string                              `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	RequestId   *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TradePrice  *float32                             `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s RunInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RunInstancesResponseBody) SetInstanceIds(v *RunInstancesResponseBodyInstanceIds) *RunInstancesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *RunInstancesResponseBody) SetOrderId(v string) *RunInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RunInstancesResponseBody) SetRequestId(v string) *RunInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunInstancesResponseBody) SetTradePrice(v float32) *RunInstancesResponseBody {
	s.TradePrice = &v
	return s
}

type RunInstancesResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s RunInstancesResponseBodyInstanceIds) String() string {
	return tea.Prettify(s)
}

func (s RunInstancesResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *RunInstancesResponseBodyInstanceIds) SetInstanceId(v []*string) *RunInstancesResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

type RunInstancesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RunInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RunInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s RunInstancesResponse) GoString() string {
	return s.String()
}

func (s *RunInstancesResponse) SetHeaders(v map[string]*string) *RunInstancesResponse {
	s.Headers = v
	return s
}

func (s *RunInstancesResponse) SetStatusCode(v int32) *RunInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *RunInstancesResponse) SetBody(v *RunInstancesResponseBody) *RunInstancesResponse {
	s.Body = v
	return s
}

type SendFileRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OssBucket            *string   `json:"OssBucket,omitempty" xml:"OssBucket,omitempty"`
	OssObject            *string   `json:"OssObject,omitempty" xml:"OssObject,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Path                 *string   `json:"Path,omitempty" xml:"Path,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SendFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) SetInstanceId(v []*string) *SendFileRequest {
	s.InstanceId = v
	return s
}

func (s *SendFileRequest) SetOssBucket(v string) *SendFileRequest {
	s.OssBucket = &v
	return s
}

func (s *SendFileRequest) SetOssObject(v string) *SendFileRequest {
	s.OssObject = &v
	return s
}

func (s *SendFileRequest) SetOwnerAccount(v string) *SendFileRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetOwnerId(v int64) *SendFileRequest {
	s.OwnerId = &v
	return s
}

func (s *SendFileRequest) SetPath(v string) *SendFileRequest {
	s.Path = &v
	return s
}

func (s *SendFileRequest) SetRegionId(v string) *SendFileRequest {
	s.RegionId = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerAccount(v string) *SendFileRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SendFileRequest) SetResourceOwnerId(v int64) *SendFileRequest {
	s.ResourceOwnerId = &v
	return s
}

type SendFileResponseBody struct {
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *SendFileResponseBodyTaskId `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Struct"`
}

func (s SendFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFileResponseBody) SetTaskId(v *SendFileResponseBodyTaskId) *SendFileResponseBody {
	s.TaskId = v
	return s
}

type SendFileResponseBodyTaskId struct {
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s SendFileResponseBodyTaskId) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBodyTaskId) GoString() string {
	return s.String()
}

func (s *SendFileResponseBodyTaskId) SetTaskId(v []*string) *SendFileResponseBodyTaskId {
	s.TaskId = v
	return s
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SendFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) SetHeaders(v map[string]*string) *SendFileResponse {
	s.Headers = v
	return s
}

func (s *SendFileResponse) SetStatusCode(v int32) *SendFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFileResponse) SetBody(v *SendFileResponseBody) *SendFileResponse {
	s.Body = v
	return s
}

type StartInstancesRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesRequest) SetInstanceId(v []*string) *StartInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StartInstancesRequest) SetOwnerAccount(v string) *StartInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetOwnerId(v int64) *StartInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StartInstancesRequest) SetRegionId(v string) *StartInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerAccount(v string) *StartInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerId(v int64) *StartInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type StartInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstancesResponseBody) SetRequestId(v string) *StartInstancesResponseBody {
	s.RequestId = &v
	return s
}

type StartInstancesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s StartInstancesResponse) GoString() string {
	return s.String()
}

func (s *StartInstancesResponse) SetHeaders(v map[string]*string) *StartInstancesResponse {
	s.Headers = v
	return s
}

func (s *StartInstancesResponse) SetStatusCode(v int32) *StartInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StartInstancesResponse) SetBody(v *StartInstancesResponseBody) *StartInstancesResponse {
	s.Body = v
	return s
}

type StopInstancesRequest struct {
	Force                *bool     `json:"Force,omitempty" xml:"Force,omitempty"`
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StopInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) SetForce(v bool) *StopInstancesRequest {
	s.Force = &v
	return s
}

func (s *StopInstancesRequest) SetInstanceId(v []*string) *StopInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StopInstancesRequest) SetOwnerAccount(v string) *StopInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetOwnerId(v int64) *StopInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInstancesRequest) SetRegionId(v string) *StopInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerAccount(v string) *StopInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerId(v int64) *StopInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

type StopInstancesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *StopInstancesResponseBody) SetRequestId(v string) *StopInstancesResponseBody {
	s.RequestId = &v
	return s
}

type StopInstancesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInstancesResponse) GoString() string {
	return s.String()
}

func (s *StopInstancesResponse) SetHeaders(v map[string]*string) *StopInstancesResponse {
	s.Headers = v
	return s
}

func (s *StopInstancesResponse) SetStatusCode(v int32) *StopInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInstancesResponse) SetBody(v *StopInstancesResponseBody) *StopInstancesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	OwnerAccount         *string                   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string                 `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string                   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Tag                  []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetOwnerAccount(v string) *TagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetOwnerId(v int64) *TagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceOwnerAccount(v string) *TagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UninstallApplicationRequest struct {
	InstanceId           []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PackageName          *string   `json:"PackageName,omitempty" xml:"PackageName,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UninstallApplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s UninstallApplicationRequest) GoString() string {
	return s.String()
}

func (s *UninstallApplicationRequest) SetInstanceId(v []*string) *UninstallApplicationRequest {
	s.InstanceId = v
	return s
}

func (s *UninstallApplicationRequest) SetOwnerAccount(v string) *UninstallApplicationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UninstallApplicationRequest) SetOwnerId(v int64) *UninstallApplicationRequest {
	s.OwnerId = &v
	return s
}

func (s *UninstallApplicationRequest) SetPackageName(v string) *UninstallApplicationRequest {
	s.PackageName = &v
	return s
}

func (s *UninstallApplicationRequest) SetRegionId(v string) *UninstallApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *UninstallApplicationRequest) SetResourceOwnerAccount(v string) *UninstallApplicationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UninstallApplicationRequest) SetResourceOwnerId(v int64) *UninstallApplicationRequest {
	s.ResourceOwnerId = &v
	return s
}

type UninstallApplicationResponseBody struct {
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId    *UninstallApplicationResponseBodyTaskId `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Struct"`
}

func (s UninstallApplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UninstallApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *UninstallApplicationResponseBody) SetRequestId(v string) *UninstallApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UninstallApplicationResponseBody) SetTaskId(v *UninstallApplicationResponseBodyTaskId) *UninstallApplicationResponseBody {
	s.TaskId = v
	return s
}

type UninstallApplicationResponseBodyTaskId struct {
	TaskId []*string `json:"TaskId,omitempty" xml:"TaskId,omitempty" type:"Repeated"`
}

func (s UninstallApplicationResponseBodyTaskId) String() string {
	return tea.Prettify(s)
}

func (s UninstallApplicationResponseBodyTaskId) GoString() string {
	return s.String()
}

func (s *UninstallApplicationResponseBodyTaskId) SetTaskId(v []*string) *UninstallApplicationResponseBodyTaskId {
	s.TaskId = v
	return s
}

type UninstallApplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UninstallApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UninstallApplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s UninstallApplicationResponse) GoString() string {
	return s.String()
}

func (s *UninstallApplicationResponse) SetHeaders(v map[string]*string) *UninstallApplicationResponse {
	s.Headers = v
	return s
}

func (s *UninstallApplicationResponse) SetStatusCode(v int32) *UninstallApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *UninstallApplicationResponse) SetBody(v *UninstallApplicationResponseBody) *UninstallApplicationResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	All                  *bool     `json:"All,omitempty" xml:"All,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceId           []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceType         *string   `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	TagKey               []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerAccount(v string) *UntagResourcesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetOwnerId(v int64) *UntagResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceOwnerAccount(v string) *UntagResourcesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateImageAttributeRequest struct {
	AddAccount           []*string `json:"AddAccount,omitempty" xml:"AddAccount,omitempty" type:"Repeated"`
	Description          *string   `json:"Description,omitempty" xml:"Description,omitempty"`
	ImageId              *string   `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName            *string   `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	OwnerAccount         *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RemoveAccount        []*string `json:"RemoveAccount,omitempty" xml:"RemoveAccount,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s UpdateImageAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateImageAttributeRequest) SetAddAccount(v []*string) *UpdateImageAttributeRequest {
	s.AddAccount = v
	return s
}

func (s *UpdateImageAttributeRequest) SetDescription(v string) *UpdateImageAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetImageId(v string) *UpdateImageAttributeRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetImageName(v string) *UpdateImageAttributeRequest {
	s.ImageName = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetOwnerAccount(v string) *UpdateImageAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetOwnerId(v int64) *UpdateImageAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetRegionId(v string) *UpdateImageAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetRemoveAccount(v []*string) *UpdateImageAttributeRequest {
	s.RemoveAccount = v
	return s
}

func (s *UpdateImageAttributeRequest) SetResourceOwnerAccount(v string) *UpdateImageAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateImageAttributeRequest) SetResourceOwnerId(v int64) *UpdateImageAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

type UpdateImageAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateImageAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateImageAttributeResponseBody) SetRequestId(v string) *UpdateImageAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateImageAttributeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateImageAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateImageAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateImageAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateImageAttributeResponse) SetHeaders(v map[string]*string) *UpdateImageAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateImageAttributeResponse) SetStatusCode(v int32) *UpdateImageAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateImageAttributeResponse) SetBody(v *UpdateImageAttributeResponseBody) *UpdateImageAttributeResponse {
	s.Body = v
	return s
}

type UpdateInstanceAttributeRequest struct {
	Description          *string                              `json:"Description,omitempty" xml:"Description,omitempty"`
	InstanceId           *string                              `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName         *string                              `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	KeyPairName          *string                              `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerAccount         *string                              `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                               `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                              `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Resolution           *string                              `json:"Resolution,omitempty" xml:"Resolution,omitempty"`
	ResourceOwnerAccount *string                              `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	Tag                  []*UpdateInstanceAttributeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VncPassword          *string                              `json:"VncPassword,omitempty" xml:"VncPassword,omitempty"`
}

func (s UpdateInstanceAttributeRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeRequest) SetDescription(v string) *UpdateInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceId(v string) *UpdateInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetInstanceName(v string) *UpdateInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetKeyPairName(v string) *UpdateInstanceAttributeRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetOwnerAccount(v string) *UpdateInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetOwnerId(v int64) *UpdateInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetRegionId(v string) *UpdateInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetResolution(v string) *UpdateInstanceAttributeRequest {
	s.Resolution = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetResourceOwnerAccount(v string) *UpdateInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetTag(v []*UpdateInstanceAttributeRequestTag) *UpdateInstanceAttributeRequest {
	s.Tag = v
	return s
}

func (s *UpdateInstanceAttributeRequest) SetVncPassword(v string) *UpdateInstanceAttributeRequest {
	s.VncPassword = &v
	return s
}

type UpdateInstanceAttributeRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateInstanceAttributeRequestTag) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeRequestTag) SetKey(v string) *UpdateInstanceAttributeRequestTag {
	s.Key = &v
	return s
}

func (s *UpdateInstanceAttributeRequestTag) SetValue(v string) *UpdateInstanceAttributeRequestTag {
	s.Value = &v
	return s
}

type UpdateInstanceAttributeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateInstanceAttributeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponseBody) SetRequestId(v string) *UpdateInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

type UpdateInstanceAttributeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateInstanceAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateInstanceAttributeResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInstanceAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAttributeResponse) SetHeaders(v map[string]*string) *UpdateInstanceAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetStatusCode(v int32) *UpdateInstanceAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAttributeResponse) SetBody(v *UpdateInstanceAttributeResponseBody) *UpdateInstanceAttributeResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudphone"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
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

func (client *Client) CancelTaskWithOptions(request *CancelTaskRequest, runtime *util.RuntimeOptions) (_result *CancelTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelTask"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CancelTask(request *CancelTaskRequest) (_result *CancelTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CancelTaskResponse{}
	_body, _err := client.CancelTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyImageWithOptions(request *CopyImageRequest, runtime *util.RuntimeOptions) (_result *CopyImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DestinationRegionId)) {
		query["DestinationRegionId"] = request.DestinationRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyImage"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyImage(request *CopyImageRequest) (_result *CopyImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CopyImageResponse{}
	_body, _err := client.CopyImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateImageWithOptions(request *CreateImageRequest, runtime *util.RuntimeOptions) (_result *CreateImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateImage"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateImage(request *CreateImageRequest) (_result *CreateImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateImageResponse{}
	_body, _err := client.CreateImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteImagesWithOptions(request *DeleteImagesRequest, runtime *util.RuntimeOptions) (_result *DeleteImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteImages"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteImages(request *DeleteImagesRequest) (_result *DeleteImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteImagesResponse{}
	_body, _err := client.DeleteImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteInstancesWithOptions(request *DeleteInstancesRequest, runtime *util.RuntimeOptions) (_result *DeleteInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteInstances(request *DeleteInstancesRequest) (_result *DeleteInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteInstancesResponse{}
	_body, _err := client.DeleteInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteKeyPairsWithOptions(request *DeleteKeyPairsRequest, runtime *util.RuntimeOptions) (_result *DeleteKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteKeyPairs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteKeyPairs(request *DeleteKeyPairsRequest) (_result *DeleteKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteKeyPairsResponse{}
	_body, _err := client.DeleteKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchFileWithOptions(request *FetchFileRequest, runtime *util.RuntimeOptions) (_result *FetchFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssObject)) {
		query["OssObject"] = request.OssObject
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("FetchFile"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &FetchFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchFile(request *FetchFileRequest) (_result *FetchFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchFileResponse{}
	_body, _err := client.FetchFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportImageWithOptions(request *ImportImageRequest, runtime *util.RuntimeOptions) (_result *ImportImageResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Format)) {
		query["Format"] = request.Format
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssObject)) {
		query["OssObject"] = request.OssObject
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		query["Platform"] = request.Platform
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportImage"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportImageResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportImage(request *ImportImageRequest) (_result *ImportImageResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportImageResponse{}
	_body, _err := client.ImportImageWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ImportKeyPairWithOptions(request *ImportKeyPairRequest, runtime *util.RuntimeOptions) (_result *ImportKeyPairResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PublicKeyBody)) {
		query["PublicKeyBody"] = request.PublicKeyBody
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportKeyPair"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ImportKeyPair(request *ImportKeyPairRequest) (_result *ImportKeyPairResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ImportKeyPairResponse{}
	_body, _err := client.ImportKeyPairWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InstallApplicationWithOptions(request *InstallApplicationRequest, runtime *util.RuntimeOptions) (_result *InstallApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssObject)) {
		query["OssObject"] = request.OssObject
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InstallApplication"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &InstallApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InstallApplication(request *InstallApplicationRequest) (_result *InstallApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InstallApplicationResponse{}
	_body, _err := client.InstallApplicationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImageSharePermissionWithOptions(request *ListImageSharePermissionRequest, runtime *util.RuntimeOptions) (_result *ListImageSharePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImageSharePermission"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImageSharePermissionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImageSharePermission(request *ListImageSharePermissionRequest) (_result *ListImageSharePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImageSharePermissionResponse{}
	_body, _err := client.ListImageSharePermissionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageCategory)) {
		query["ImageCategory"] = request.ImageCategory
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceTypesWithOptions(request *ListInstanceTypesRequest, runtime *util.RuntimeOptions) (_result *ListInstanceTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceTypeFamily)) {
		query["InstanceTypeFamily"] = request.InstanceTypeFamily
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceTypes"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceTypes(request *ListInstanceTypesRequest) (_result *ListInstanceTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceTypesResponse{}
	_body, _err := client.ListInstanceTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstanceVncUrlWithOptions(request *ListInstanceVncUrlRequest, runtime *util.RuntimeOptions) (_result *ListInstanceVncUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstanceVncUrl"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstanceVncUrlResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstanceVncUrl(request *ListInstanceVncUrlRequest) (_result *ListInstanceVncUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstanceVncUrlResponse{}
	_body, _err := client.ListInstanceVncUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListInstancesWithOptions(request *ListInstancesRequest, runtime *util.RuntimeOptions) (_result *ListInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resolution)) {
		query["Resolution"] = request.Resolution
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowWebRtcToken)) {
		query["ShowWebRtcToken"] = request.ShowWebRtcToken
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.ZoneId)) {
		query["ZoneId"] = request.ZoneId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListInstances(request *ListInstancesRequest) (_result *ListInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListInstancesResponse{}
	_body, _err := client.ListInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListKeyPairsWithOptions(request *ListKeyPairsRequest, runtime *util.RuntimeOptions) (_result *ListKeyPairsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.KeyPairFingerPrint)) {
		query["KeyPairFingerPrint"] = request.KeyPairFingerPrint
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListKeyPairs"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListKeyPairsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListKeyPairs(request *ListKeyPairsRequest) (_result *ListKeyPairsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListKeyPairsResponse{}
	_body, _err := client.ListKeyPairsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegionsWithOptions(request *ListRegionsRequest, runtime *util.RuntimeOptions) (_result *ListRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRegions"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegions(request *ListRegionsRequest) (_result *ListRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegionsResponse{}
	_body, _err := client.ListRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagKeysWithOptions(request *ListTagKeysRequest, runtime *util.RuntimeOptions) (_result *ListTagKeysResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagKeys"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagKeysResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagKeys(request *ListTagKeysRequest) (_result *ListTagKeysResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagKeysResponse{}
	_body, _err := client.ListTagKeysWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) ListTagValuesWithOptions(request *ListTagValuesRequest, runtime *util.RuntimeOptions) (_result *ListTagValuesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Key)) {
		query["Key"] = request.Key
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagValues"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagValuesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagValues(request *ListTagValuesRequest) (_result *ListTagValuesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagValuesResponse{}
	_body, _err := client.ListTagValuesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTasksWithOptions(request *ListTasksRequest, runtime *util.RuntimeOptions) (_result *ListTasksResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		query["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		query["TaskId"] = request.TaskId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskStatus)) {
		query["TaskStatus"] = request.TaskStatus
	}

	if !tea.BoolValue(util.IsUnset(request.TaskType)) {
		query["TaskType"] = request.TaskType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTasks"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTasksResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTasks(request *ListTasksRequest) (_result *ListTasksResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTasksResponse{}
	_body, _err := client.ListTasksWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListZonesWithOptions(request *ListZonesRequest, runtime *util.RuntimeOptions) (_result *ListZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListZones"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListZones(request *ListZonesRequest) (_result *ListZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListZonesResponse{}
	_body, _err := client.ListZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RebootInstancesWithOptions(request *RebootInstancesRequest, runtime *util.RuntimeOptions) (_result *RebootInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RebootInstances(request *RebootInstancesRequest) (_result *RebootInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootInstancesResponse{}
	_body, _err := client.RebootInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenewInstancesWithOptions(request *RenewInstancesRequest, runtime *util.RuntimeOptions) (_result *RenewInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RenewInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RenewInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenewInstances(request *RenewInstancesRequest) (_result *RenewInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenewInstancesResponse{}
	_body, _err := client.RenewInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ResetInstancesWithOptions(request *ResetInstancesRequest, runtime *util.RuntimeOptions) (_result *ResetInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ResetInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ResetInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ResetInstances(request *ResetInstancesRequest) (_result *ResetInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ResetInstancesResponse{}
	_body, _err := client.ResetInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunCommandWithOptions(request *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Command)) {
		query["Command"] = request.Command
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RunInstancesWithOptions(request *RunInstancesRequest, runtime *util.RuntimeOptions) (_result *RunInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		query["Amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		query["AutoPay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		query["AutoRenew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EipBandwidth)) {
		query["EipBandwidth"] = request.EipBandwidth
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceType)) {
		query["InstanceType"] = request.InstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		query["Period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		query["PeriodUnit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.PrivateIpAddress)) {
		query["PrivateIpAddress"] = request.PrivateIpAddress
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resolution)) {
		query["Resolution"] = request.Resolution
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		query["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		query["VSwitchId"] = request.VSwitchId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("RunInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RunInstances(request *RunInstancesRequest) (_result *RunInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunInstancesResponse{}
	_body, _err := client.RunInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SendFileWithOptions(request *SendFileRequest, runtime *util.RuntimeOptions) (_result *SendFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OssBucket)) {
		query["OssBucket"] = request.OssBucket
	}

	if !tea.BoolValue(util.IsUnset(request.OssObject)) {
		query["OssObject"] = request.OssObject
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		query["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("SendFile"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendFileResponse{}
	_body, _err := client.SendFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartInstancesWithOptions(request *StartInstancesRequest, runtime *util.RuntimeOptions) (_result *StartInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StartInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StartInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartInstances(request *StartInstancesRequest) (_result *StartInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StartInstancesResponse{}
	_body, _err := client.StartInstancesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopInstancesWithOptions(request *StopInstancesRequest, runtime *util.RuntimeOptions) (_result *StopInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Force)) {
		query["Force"] = request.Force
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInstances"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopInstances(request *StopInstancesRequest) (_result *StopInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInstancesResponse{}
	_body, _err := client.StopInstancesWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UninstallApplicationWithOptions(request *UninstallApplicationRequest, runtime *util.RuntimeOptions) (_result *UninstallApplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.PackageName)) {
		query["PackageName"] = request.PackageName
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UninstallApplication"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UninstallApplicationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UninstallApplication(request *UninstallApplicationRequest) (_result *UninstallApplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UninstallApplicationResponse{}
	_body, _err := client.UninstallApplicationWithOptions(request, runtime)
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
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

func (client *Client) UpdateImageAttributeWithOptions(request *UpdateImageAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateImageAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddAccount)) {
		query["AddAccount"] = request.AddAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.ImageId)) {
		query["ImageId"] = request.ImageId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageName)) {
		query["ImageName"] = request.ImageName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RemoveAccount)) {
		query["RemoveAccount"] = request.RemoveAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerId)) {
		query["ResourceOwnerId"] = request.ResourceOwnerId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateImageAttribute"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateImageAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateImageAttribute(request *UpdateImageAttributeRequest) (_result *UpdateImageAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateImageAttributeResponse{}
	_body, _err := client.UpdateImageAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateInstanceAttributeWithOptions(request *UpdateInstanceAttributeRequest, runtime *util.RuntimeOptions) (_result *UpdateInstanceAttributeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		query["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		query["InstanceId"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceName)) {
		query["InstanceName"] = request.InstanceName
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		query["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerAccount)) {
		query["OwnerAccount"] = request.OwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerId)) {
		query["OwnerId"] = request.OwnerId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.Resolution)) {
		query["Resolution"] = request.Resolution
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceOwnerAccount)) {
		query["ResourceOwnerAccount"] = request.ResourceOwnerAccount
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VncPassword)) {
		query["VncPassword"] = request.VncPassword
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateInstanceAttribute"),
		Version:     tea.String("2020-12-30"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateInstanceAttribute(request *UpdateInstanceAttributeRequest) (_result *UpdateInstanceAttributeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateInstanceAttributeResponse{}
	_body, _err := client.UpdateInstanceAttributeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
