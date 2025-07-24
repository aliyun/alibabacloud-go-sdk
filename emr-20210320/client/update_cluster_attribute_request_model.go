// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateClusterAttributeRequest
	GetClusterId() *string
	SetClusterName(v string) *UpdateClusterAttributeRequest
	GetClusterName() *string
	SetDeletionProtection(v bool) *UpdateClusterAttributeRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *UpdateClusterAttributeRequest
	GetDescription() *string
	SetRegionId(v string) *UpdateClusterAttributeRequest
	GetRegionId() *string
}

type UpdateClusterAttributeRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// emrtest
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Specifies whether release protection is enabled.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The cluster description.
	//
	// example:
	//
	// Emr cluster for ETL
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s UpdateClusterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterAttributeRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterAttributeRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *UpdateClusterAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *UpdateClusterAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateClusterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateClusterAttributeRequest) SetClusterId(v string) *UpdateClusterAttributeRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterAttributeRequest) SetClusterName(v string) *UpdateClusterAttributeRequest {
	s.ClusterName = &v
	return s
}

func (s *UpdateClusterAttributeRequest) SetDeletionProtection(v bool) *UpdateClusterAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *UpdateClusterAttributeRequest) SetDescription(v string) *UpdateClusterAttributeRequest {
	s.Description = &v
	return s
}

func (s *UpdateClusterAttributeRequest) SetRegionId(v string) *UpdateClusterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateClusterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
