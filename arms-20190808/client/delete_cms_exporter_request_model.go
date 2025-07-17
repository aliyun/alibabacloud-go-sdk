// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCmsExporterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteCmsExporterRequest
	GetClusterId() *string
	SetRegionId(v string) *DeleteCmsExporterRequest
	GetRegionId() *string
}

type DeleteCmsExporterRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cc7a37ee31aea4ed1a059eff8034b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCmsExporterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCmsExporterRequest) GoString() string {
	return s.String()
}

func (s *DeleteCmsExporterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteCmsExporterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCmsExporterRequest) SetClusterId(v string) *DeleteCmsExporterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteCmsExporterRequest) SetRegionId(v string) *DeleteCmsExporterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCmsExporterRequest) Validate() error {
	return dara.Validate(s)
}
