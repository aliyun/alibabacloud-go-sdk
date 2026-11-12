// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLakebaseMountParameterTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeLakebaseMountParameterTemplateListRequest
	GetEngine() *string
	SetParamGroupCategory(v string) *DescribeLakebaseMountParameterTemplateListRequest
	GetParamGroupCategory() *string
	SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterTemplateListRequest
	GetPolarFsInstanceId() *string
	SetRegionId(v string) *DescribeLakebaseMountParameterTemplateListRequest
	GetRegionId() *string
}

type DescribeLakebaseMountParameterTemplateListRequest struct {
	// The engine. Set the value to polar_3fs.
	//
	// This parameter is required.
	//
	// example:
	//
	// polar_3fs
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The instance category. Valid values: basic (Basic Edition) and high_performance (High-performance Edition).
	//
	// example:
	//
	// basic
	ParamGroupCategory *string `json:"ParamGroupCategory,omitempty" xml:"ParamGroupCategory,omitempty"`
	// The PolarFS instance ID. You can call DescribePolarFs to query available PolarFS instances and use the returned PolarFsInstanceId.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PolarFsInstanceId *string `json:"PolarFsInstanceId,omitempty" xml:"PolarFsInstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLakebaseMountParameterTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLakebaseMountParameterTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) GetParamGroupCategory() *string {
	return s.ParamGroupCategory
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) GetPolarFsInstanceId() *string {
	return s.PolarFsInstanceId
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) SetEngine(v string) *DescribeLakebaseMountParameterTemplateListRequest {
	s.Engine = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) SetParamGroupCategory(v string) *DescribeLakebaseMountParameterTemplateListRequest {
	s.ParamGroupCategory = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) SetPolarFsInstanceId(v string) *DescribeLakebaseMountParameterTemplateListRequest {
	s.PolarFsInstanceId = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) SetRegionId(v string) *DescribeLakebaseMountParameterTemplateListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLakebaseMountParameterTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
