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
	SetInstanceName(v string) *DeleteAppInstanceRequest
	GetInstanceName() *string
	SetRegionId(v string) *DeleteAppInstanceRequest
	GetRegionId() *string
}

type DeleteAppInstanceRequest struct {
	// example:
	//
	// ETnLKlblzczshOTUbOCz****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// ra-supabase-8moov5lxba****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
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
