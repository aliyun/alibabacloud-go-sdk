// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *CreateModelServiceRequest
	GetApiKey() *string
	SetBaseUrl(v string) *CreateModelServiceRequest
	GetBaseUrl() *string
	SetGwClusterId(v string) *CreateModelServiceRequest
	GetGwClusterId() *string
	SetInputCostPointsPerMillion(v string) *CreateModelServiceRequest
	GetInputCostPointsPerMillion() *string
	SetModelCategory(v string) *CreateModelServiceRequest
	GetModelCategory() *string
	SetName(v string) *CreateModelServiceRequest
	GetName() *string
	SetOutputCostPointsPerMillion(v string) *CreateModelServiceRequest
	GetOutputCostPointsPerMillion() *string
	SetProtocol(v string) *CreateModelServiceRequest
	GetProtocol() *string
	SetRegionId(v string) *CreateModelServiceRequest
	GetRegionId() *string
	SetRequestCostPoints(v string) *CreateModelServiceRequest
	GetRequestCostPoints() *string
	SetVendor(v string) *CreateModelServiceRequest
	GetVendor() *string
}

type CreateModelServiceRequest struct {
	// The API key of the model service.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6c4b1f0317cd4fd7a5b446d3503d**
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The address of the upstream service.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// The gateway instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The number of input points.
	//
	// example:
	//
	// 10
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// The model category.
	//
	// This parameter is required.
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// The model service name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of output points.
	//
	// example:
	//
	// 10
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// The protocol type.
	//
	// This parameter is required.
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of requests.
	//
	// example:
	//
	// 10
	RequestCostPoints *string `json:"RequestCostPoints,omitempty" xml:"RequestCostPoints,omitempty"`
	// The model service provider.
	//
	// This parameter is required.
	//
	// example:
	//
	// bailian
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s CreateModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateModelServiceRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateModelServiceRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *CreateModelServiceRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateModelServiceRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *CreateModelServiceRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *CreateModelServiceRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelServiceRequest) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *CreateModelServiceRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateModelServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateModelServiceRequest) GetRequestCostPoints() *string {
	return s.RequestCostPoints
}

func (s *CreateModelServiceRequest) GetVendor() *string {
	return s.Vendor
}

func (s *CreateModelServiceRequest) SetApiKey(v string) *CreateModelServiceRequest {
	s.ApiKey = &v
	return s
}

func (s *CreateModelServiceRequest) SetBaseUrl(v string) *CreateModelServiceRequest {
	s.BaseUrl = &v
	return s
}

func (s *CreateModelServiceRequest) SetGwClusterId(v string) *CreateModelServiceRequest {
	s.GwClusterId = &v
	return s
}

func (s *CreateModelServiceRequest) SetInputCostPointsPerMillion(v string) *CreateModelServiceRequest {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *CreateModelServiceRequest) SetModelCategory(v string) *CreateModelServiceRequest {
	s.ModelCategory = &v
	return s
}

func (s *CreateModelServiceRequest) SetName(v string) *CreateModelServiceRequest {
	s.Name = &v
	return s
}

func (s *CreateModelServiceRequest) SetOutputCostPointsPerMillion(v string) *CreateModelServiceRequest {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *CreateModelServiceRequest) SetProtocol(v string) *CreateModelServiceRequest {
	s.Protocol = &v
	return s
}

func (s *CreateModelServiceRequest) SetRegionId(v string) *CreateModelServiceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateModelServiceRequest) SetRequestCostPoints(v string) *CreateModelServiceRequest {
	s.RequestCostPoints = &v
	return s
}

func (s *CreateModelServiceRequest) SetVendor(v string) *CreateModelServiceRequest {
	s.Vendor = &v
	return s
}

func (s *CreateModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
