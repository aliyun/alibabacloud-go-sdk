// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DeleteModelServiceRequest
	GetGwClusterId() *string
	SetModelName(v string) *DeleteModelServiceRequest
	GetModelName() *string
	SetRegionId(v string) *DeleteModelServiceRequest
	GetRegionId() *string
}

type DeleteModelServiceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelServiceRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteModelServiceRequest) GetModelName() *string {
	return s.ModelName
}

func (s *DeleteModelServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteModelServiceRequest) SetGwClusterId(v string) *DeleteModelServiceRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteModelServiceRequest) SetModelName(v string) *DeleteModelServiceRequest {
	s.ModelName = &v
	return s
}

func (s *DeleteModelServiceRequest) SetRegionId(v string) *DeleteModelServiceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
