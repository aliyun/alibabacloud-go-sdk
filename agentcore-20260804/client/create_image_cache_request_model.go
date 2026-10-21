// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *CreateImageCacheRequestBody) *CreateImageCacheRequest
	GetBody() *CreateImageCacheRequestBody
	SetClientToken(v string) *CreateImageCacheRequest
	GetClientToken() *string
	SetRegionId(v string) *CreateImageCacheRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *CreateImageCacheRequest
	GetWorkspaceId() *string
}

type CreateImageCacheRequest struct {
	// The business parameters for creating an image cache.
	//
	// This parameter is required.
	Body *CreateImageCacheRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
	// This parameter is required and cannot be empty. It is reserved to be compatible with the API. Currently, it does not provide the ability to deduplicate requests or recover the original result based on this field. If a timeout occurs or the result is unconfirmed, query the query cache status first before deciding whether to recommit.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The product entry region, which is determined by the selected region endpoint.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// The current routed workspace. The server verifies the account, region, and resource ownership. This parameter cannot be omitted or used to access other workspaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// ws-example
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequest) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequest) GetBody() *CreateImageCacheRequestBody {
	return s.Body
}

func (s *CreateImageCacheRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageCacheRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateImageCacheRequest) SetBody(v *CreateImageCacheRequestBody) *CreateImageCacheRequest {
	s.Body = v
	return s
}

func (s *CreateImageCacheRequest) SetClientToken(v string) *CreateImageCacheRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageCacheRequest) SetRegionId(v string) *CreateImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageCacheRequest) SetWorkspaceId(v string) *CreateImageCacheRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateImageCacheRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateImageCacheRequestBody struct {
	// Specifies whether to allow the current source reading and transfer to use the public network. At least one of this parameter and accessVpc must be set to true. The server verifies the access conditions of the current workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AccessInternet *bool `json:"accessInternet,omitempty" xml:"accessInternet,omitempty"`
	// Specifies whether to allow the current source reading and transfer to use a VPC. At least one of this parameter and accessInternet must be set to true. The server parses the trusted network of the current workspace and does not modify the agent runtime network.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	AccessVpc *bool `json:"accessVpc,omitempty" xml:"accessVpc,omitempty"`
	// The ID of the source ACR Enterprise instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-example
	AcrInstanceId *string `json:"acrInstanceId,omitempty" xml:"acrInstanceId,omitempty"`
	// The namespace of the source ACR.
	//
	// This parameter is required.
	//
	// example:
	//
	// team
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The name of the repository under the namespace. Do not include the namespace repeatedly.
	//
	// This parameter is required.
	//
	// example:
	//
	// app
	Repository *string `json:"repository,omitempty" xml:"repository,omitempty"`
	// The selected tag of the source image.
	//
	// This parameter is required.
	//
	// example:
	//
	// v1
	Tag *string `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CreateImageCacheRequestBody) String() string {
	return dara.Prettify(s)
}

func (s CreateImageCacheRequestBody) GoString() string {
	return s.String()
}

func (s *CreateImageCacheRequestBody) GetAccessInternet() *bool {
	return s.AccessInternet
}

func (s *CreateImageCacheRequestBody) GetAccessVpc() *bool {
	return s.AccessVpc
}

func (s *CreateImageCacheRequestBody) GetAcrInstanceId() *string {
	return s.AcrInstanceId
}

func (s *CreateImageCacheRequestBody) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateImageCacheRequestBody) GetRepository() *string {
	return s.Repository
}

func (s *CreateImageCacheRequestBody) GetTag() *string {
	return s.Tag
}

func (s *CreateImageCacheRequestBody) SetAccessInternet(v bool) *CreateImageCacheRequestBody {
	s.AccessInternet = &v
	return s
}

func (s *CreateImageCacheRequestBody) SetAccessVpc(v bool) *CreateImageCacheRequestBody {
	s.AccessVpc = &v
	return s
}

func (s *CreateImageCacheRequestBody) SetAcrInstanceId(v string) *CreateImageCacheRequestBody {
	s.AcrInstanceId = &v
	return s
}

func (s *CreateImageCacheRequestBody) SetNamespace(v string) *CreateImageCacheRequestBody {
	s.Namespace = &v
	return s
}

func (s *CreateImageCacheRequestBody) SetRepository(v string) *CreateImageCacheRequestBody {
	s.Repository = &v
	return s
}

func (s *CreateImageCacheRequestBody) SetTag(v string) *CreateImageCacheRequestBody {
	s.Tag = &v
	return s
}

func (s *CreateImageCacheRequestBody) Validate() error {
	return dara.Validate(s)
}
