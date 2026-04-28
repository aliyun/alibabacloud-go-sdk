// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiKey(v string) *ModifyModelServiceRequest
	GetApiKey() *string
	SetBaseUrl(v string) *ModifyModelServiceRequest
	GetBaseUrl() *string
	SetGwClusterId(v string) *ModifyModelServiceRequest
	GetGwClusterId() *string
	SetInputCostPointsPerMillion(v string) *ModifyModelServiceRequest
	GetInputCostPointsPerMillion() *string
	SetModelCategory(v string) *ModifyModelServiceRequest
	GetModelCategory() *string
	SetModelServiceId(v string) *ModifyModelServiceRequest
	GetModelServiceId() *string
	SetName(v string) *ModifyModelServiceRequest
	GetName() *string
	SetOutputCostPointsPerMillion(v string) *ModifyModelServiceRequest
	GetOutputCostPointsPerMillion() *string
	SetProtocol(v string) *ModifyModelServiceRequest
	GetProtocol() *string
	SetRegionId(v string) *ModifyModelServiceRequest
	GetRegionId() *string
	SetRequestCostPoints(v string) *ModifyModelServiceRequest
	GetRequestCostPoints() *string
}

type ModifyModelServiceRequest struct {
	// example:
	//
	// xxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// example:
	//
	// https://dashscope.aliyuncs.com/compatible-mode/v1
	BaseUrl *string `json:"BaseUrl,omitempty" xml:"BaseUrl,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// example:
	//
	// 10
	InputCostPointsPerMillion *string `json:"InputCostPointsPerMillion,omitempty" xml:"InputCostPointsPerMillion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// text
	ModelCategory *string `json:"ModelCategory,omitempty" xml:"ModelCategory,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ms-xxxxxx
	ModelServiceId *string `json:"ModelServiceId,omitempty" xml:"ModelServiceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	OutputCostPointsPerMillion *string `json:"OutputCostPointsPerMillion,omitempty" xml:"OutputCostPointsPerMillion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// openai
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 10
	RequestCostPoints *string `json:"RequestCostPoints,omitempty" xml:"RequestCostPoints,omitempty"`
}

func (s ModifyModelServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyModelServiceRequest) GetApiKey() *string {
	return s.ApiKey
}

func (s *ModifyModelServiceRequest) GetBaseUrl() *string {
	return s.BaseUrl
}

func (s *ModifyModelServiceRequest) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *ModifyModelServiceRequest) GetInputCostPointsPerMillion() *string {
	return s.InputCostPointsPerMillion
}

func (s *ModifyModelServiceRequest) GetModelCategory() *string {
	return s.ModelCategory
}

func (s *ModifyModelServiceRequest) GetModelServiceId() *string {
	return s.ModelServiceId
}

func (s *ModifyModelServiceRequest) GetName() *string {
	return s.Name
}

func (s *ModifyModelServiceRequest) GetOutputCostPointsPerMillion() *string {
	return s.OutputCostPointsPerMillion
}

func (s *ModifyModelServiceRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyModelServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyModelServiceRequest) GetRequestCostPoints() *string {
	return s.RequestCostPoints
}

func (s *ModifyModelServiceRequest) SetApiKey(v string) *ModifyModelServiceRequest {
	s.ApiKey = &v
	return s
}

func (s *ModifyModelServiceRequest) SetBaseUrl(v string) *ModifyModelServiceRequest {
	s.BaseUrl = &v
	return s
}

func (s *ModifyModelServiceRequest) SetGwClusterId(v string) *ModifyModelServiceRequest {
	s.GwClusterId = &v
	return s
}

func (s *ModifyModelServiceRequest) SetInputCostPointsPerMillion(v string) *ModifyModelServiceRequest {
	s.InputCostPointsPerMillion = &v
	return s
}

func (s *ModifyModelServiceRequest) SetModelCategory(v string) *ModifyModelServiceRequest {
	s.ModelCategory = &v
	return s
}

func (s *ModifyModelServiceRequest) SetModelServiceId(v string) *ModifyModelServiceRequest {
	s.ModelServiceId = &v
	return s
}

func (s *ModifyModelServiceRequest) SetName(v string) *ModifyModelServiceRequest {
	s.Name = &v
	return s
}

func (s *ModifyModelServiceRequest) SetOutputCostPointsPerMillion(v string) *ModifyModelServiceRequest {
	s.OutputCostPointsPerMillion = &v
	return s
}

func (s *ModifyModelServiceRequest) SetProtocol(v string) *ModifyModelServiceRequest {
	s.Protocol = &v
	return s
}

func (s *ModifyModelServiceRequest) SetRegionId(v string) *ModifyModelServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyModelServiceRequest) SetRequestCostPoints(v string) *ModifyModelServiceRequest {
	s.RequestCostPoints = &v
	return s
}

func (s *ModifyModelServiceRequest) Validate() error {
	return dara.Validate(s)
}
