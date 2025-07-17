// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeletePromInstance(v bool) *DeleteEnvironmentRequest
	GetDeletePromInstance() *bool
	SetEnvironmentId(v string) *DeleteEnvironmentRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *DeleteEnvironmentRequest
	GetRegionId() *string
}

type DeleteEnvironmentRequest struct {
	// Specifies whether to delete the related Prometheus instance.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	DeletePromInstance *bool `json:"DeletePromInstance,omitempty" xml:"DeletePromInstance,omitempty"`
	// The ID of the environment instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// env-xxxxxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteEnvironmentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentRequest) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentRequest) GetDeletePromInstance() *bool {
	return s.DeletePromInstance
}

func (s *DeleteEnvironmentRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *DeleteEnvironmentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteEnvironmentRequest) SetDeletePromInstance(v bool) *DeleteEnvironmentRequest {
	s.DeletePromInstance = &v
	return s
}

func (s *DeleteEnvironmentRequest) SetEnvironmentId(v string) *DeleteEnvironmentRequest {
	s.EnvironmentId = &v
	return s
}

func (s *DeleteEnvironmentRequest) SetRegionId(v string) *DeleteEnvironmentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteEnvironmentRequest) Validate() error {
	return dara.Validate(s)
}
