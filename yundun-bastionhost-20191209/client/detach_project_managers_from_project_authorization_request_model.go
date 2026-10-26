// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachProjectManagersFromProjectAuthorizationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationId(v int64) *DetachProjectManagersFromProjectAuthorizationRequest
	GetAuthorizationId() *int64
	SetInstanceId(v string) *DetachProjectManagersFromProjectAuthorizationRequest
	GetInstanceId() *string
	SetPrincipalIds(v string) *DetachProjectManagersFromProjectAuthorizationRequest
	GetPrincipalIds() *string
	SetRegionId(v string) *DetachProjectManagersFromProjectAuthorizationRequest
	GetRegionId() *string
}

type DetachProjectManagersFromProjectAuthorizationRequest struct {
	// The project authorization ID.
	//
	// > You can call the ListProjectAuthorizations operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AuthorizationId *int64 `json:"AuthorizationId,omitempty" xml:"AuthorizationId,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zpr3h2zo60l
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The array of project administrator IDs to remove.
	//
	// > You can call the GetProjectAuthorization operation to query this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2036****6767
	PrincipalIds *string `json:"PrincipalIds,omitempty" xml:"PrincipalIds,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DetachProjectManagersFromProjectAuthorizationRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachProjectManagersFromProjectAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) GetAuthorizationId() *int64 {
	return s.AuthorizationId
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) GetPrincipalIds() *string {
	return s.PrincipalIds
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) SetAuthorizationId(v int64) *DetachProjectManagersFromProjectAuthorizationRequest {
	s.AuthorizationId = &v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) SetInstanceId(v string) *DetachProjectManagersFromProjectAuthorizationRequest {
	s.InstanceId = &v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) SetPrincipalIds(v string) *DetachProjectManagersFromProjectAuthorizationRequest {
	s.PrincipalIds = &v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) SetRegionId(v string) *DetachProjectManagersFromProjectAuthorizationRequest {
	s.RegionId = &v
	return s
}

func (s *DetachProjectManagersFromProjectAuthorizationRequest) Validate() error {
	return dara.Validate(s)
}
