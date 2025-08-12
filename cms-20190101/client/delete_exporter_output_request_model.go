// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExporterOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestName(v string) *DeleteExporterOutputRequest
	GetDestName() *string
	SetRegionId(v string) *DeleteExporterOutputRequest
	GetRegionId() *string
}

type DeleteExporterOutputRequest struct {
	// The name of the configuration set.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	DestName *string `json:"DestName,omitempty" xml:"DestName,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteExporterOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExporterOutputRequest) GoString() string {
	return s.String()
}

func (s *DeleteExporterOutputRequest) GetDestName() *string {
	return s.DestName
}

func (s *DeleteExporterOutputRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExporterOutputRequest) SetDestName(v string) *DeleteExporterOutputRequest {
	s.DestName = &v
	return s
}

func (s *DeleteExporterOutputRequest) SetRegionId(v string) *DeleteExporterOutputRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExporterOutputRequest) Validate() error {
	return dara.Validate(s)
}
