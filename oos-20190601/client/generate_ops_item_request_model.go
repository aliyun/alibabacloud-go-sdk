// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOpsItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GenerateOpsItemRequest
	GetClientToken() *string
	SetConfigurationId(v string) *GenerateOpsItemRequest
	GetConfigurationId() *string
	SetData(v string) *GenerateOpsItemRequest
	GetData() *string
	SetDataSource(v string) *GenerateOpsItemRequest
	GetDataSource() *string
	SetRegionId(v string) *GenerateOpsItemRequest
	GetRegionId() *string
}

type GenerateOpsItemRequest struct {
	// The token that is used to ensure the idempotency.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the O\\&M item.
	//
	// example:
	//
	// oic-ae4f******2c682e3
	ConfigurationId *string `json:"ConfigurationId,omitempty" xml:"ConfigurationId,omitempty"`
	// The source system data.
	//
	// This parameter is required.
	//
	// example:
	//
	// /aliyun/ecs
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The data source system.
	//
	// example:
	//
	// /aliyun/eventbridge/event
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GenerateOpsItemRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateOpsItemRequest) GoString() string {
	return s.String()
}

func (s *GenerateOpsItemRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GenerateOpsItemRequest) GetConfigurationId() *string {
	return s.ConfigurationId
}

func (s *GenerateOpsItemRequest) GetData() *string {
	return s.Data
}

func (s *GenerateOpsItemRequest) GetDataSource() *string {
	return s.DataSource
}

func (s *GenerateOpsItemRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GenerateOpsItemRequest) SetClientToken(v string) *GenerateOpsItemRequest {
	s.ClientToken = &v
	return s
}

func (s *GenerateOpsItemRequest) SetConfigurationId(v string) *GenerateOpsItemRequest {
	s.ConfigurationId = &v
	return s
}

func (s *GenerateOpsItemRequest) SetData(v string) *GenerateOpsItemRequest {
	s.Data = &v
	return s
}

func (s *GenerateOpsItemRequest) SetDataSource(v string) *GenerateOpsItemRequest {
	s.DataSource = &v
	return s
}

func (s *GenerateOpsItemRequest) SetRegionId(v string) *GenerateOpsItemRequest {
	s.RegionId = &v
	return s
}

func (s *GenerateOpsItemRequest) Validate() error {
	return dara.Validate(s)
}
