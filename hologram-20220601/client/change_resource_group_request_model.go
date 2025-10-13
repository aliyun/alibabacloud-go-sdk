// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ChangeResourceGroupRequest
	GetInstanceId() *string
	SetNewResourceGroupId(v string) *ChangeResourceGroupRequest
	GetNewResourceGroupId() *string
}

type ChangeResourceGroupRequest struct {
	// The instance ID.
	//
	// example:
	//
	// hgprecn-cn-zvp25ysv3006
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// new resource group id
	//
	// example:
	//
	// rg-acfmxwerqwerasfd
	NewResourceGroupId *string `json:"newResourceGroupId,omitempty" xml:"newResourceGroupId,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *ChangeResourceGroupRequest) SetInstanceId(v string) *ChangeResourceGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetNewResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
