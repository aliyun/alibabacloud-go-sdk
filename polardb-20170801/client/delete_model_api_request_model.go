// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGwClusterId(v string) *DeleteModelApiRequest
	GetGwClusterId() *string
	SetModelApiId(v string) *DeleteModelApiRequest
	GetModelApiId() *string
	SetRegionId(v string) *DeleteModelApiRequest
	GetRegionId() *string
}

type DeleteModelApiRequest struct {
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
	// mi-xxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteModelApiRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelApiRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelApiRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DeleteModelApiRequest) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *DeleteModelApiRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteModelApiRequest) SetGwClusterId(v string) *DeleteModelApiRequest {
	s.GwClusterId = &v
	return s
}

func (s *DeleteModelApiRequest) SetModelApiId(v string) *DeleteModelApiRequest {
	s.ModelApiId = &v
	return s
}

func (s *DeleteModelApiRequest) SetRegionId(v string) *DeleteModelApiRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteModelApiRequest) Validate() error {
	return dara.Validate(s)
}
