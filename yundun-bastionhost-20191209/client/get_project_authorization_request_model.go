// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *GetProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *GetProjectAuthorizationRequest
	GetInstanceId() *string
	SetRegionId(v string) *GetProjectAuthorizationRequest
	GetRegionId() *string
}

type GetProjectAuthorizationRequest struct {
	// The ID of the project authorization to query.
	//
	// > You can call the ListProjectAuthorizations operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *GetProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *GetProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetProjectAuthorizationRequest) SetAuthorizationId(v int64) *GetProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *GetProjectAuthorizationRequest) SetInstanceId(v string) *GetProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetProjectAuthorizationRequest) SetRegionId(v string) *GetProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *GetProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
