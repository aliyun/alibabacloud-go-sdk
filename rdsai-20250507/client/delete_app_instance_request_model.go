// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteAppInstanceRequest
	GetClientToken() *string
	SetDeleteDBInstance(v bool) *DeleteAppInstanceRequest
	GetDeleteDBInstance() *bool
	SetInstanceName(v string) *DeleteAppInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *DeleteAppInstanceRequest
	GetRegionId() *string
}

type DeleteAppInstanceRequest struct {
	// The client token that is used to ensure the idempotence of the request. The client generates the value of this parameter to prevent duplicate requests from being submitted.
	//
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to delete the corresponding database instance.
	DeleteDBInstance *bool `json:"DeleteDBInstance,omitempty" xml:"DeleteDBInstance,omitempty"`
	// The instance ID of the AI application.
	//
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAppInstanceRequest) GetDeleteDBInstance() *bool {
	return s.DeleteDBInstance
}

func (s *DeleteAppInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DeleteAppInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAppInstanceRequest) SetClientToken(v string) *DeleteAppInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAppInstanceRequest) SetDeleteDBInstance(v bool) *DeleteAppInstanceRequest {
	s.DeleteDBInstance = &v
	return s
}

func (s *DeleteAppInstanceRequest) SetInstanceName(v string) *DeleteAppInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *DeleteAppInstanceRequest) SetRegionId(v string) *DeleteAppInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
