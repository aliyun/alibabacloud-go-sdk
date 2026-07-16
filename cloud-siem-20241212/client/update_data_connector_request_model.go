// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataConnectorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfigId(v string) *UpdateDataConnectorRequest
	GetAuthConfigId() *string
	SetAuthConfigProduct(v string) *UpdateDataConnectorRequest
	GetAuthConfigProduct() *string
	SetAuthConfigVendor(v string) *UpdateDataConnectorRequest
	GetAuthConfigVendor() *string
	SetDataConnectorConfig(v string) *UpdateDataConnectorRequest
	GetDataConnectorConfig() *string
	SetDataConnectorId(v string) *UpdateDataConnectorRequest
	GetDataConnectorId() *string
	SetDataConnectorStatus(v string) *UpdateDataConnectorRequest
	GetDataConnectorStatus() *string
	SetLang(v string) *UpdateDataConnectorRequest
	GetLang() *string
	SetRegionId(v string) *UpdateDataConnectorRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataConnectorRequest
	GetRoleFor() *int64
}

type UpdateDataConnectorRequest struct {
	// The configuration item ID of the collector access object in the multi-cloud configuration.
	//
	// example:
	//
	// Opera20_Salesforce_Prod
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// The cloud service to which the authentication configuration belongs.
	//
	// example:
	//
	// salesForceRestAPI
	AuthConfigProduct *string `json:"AuthConfigProduct,omitempty" xml:"AuthConfigProduct,omitempty"`
	// The authentication vendor name.
	//
	// example:
	//
	// SALESFORCE
	AuthConfigVendor *string `json:"AuthConfigVendor,omitempty" xml:"AuthConfigVendor,omitempty"`
	// The configuration information of the collector.
	//
	// example:
	//
	// {\\"regionId\\":\\"cn-hangzhou\\",\\"bucket\\":\\"actiontrail-logs-1481501495248334-d776c375\\",\\"format\\":{\\"type\\":\\"JSON\\"},\\"encoding\\":\\"UTF-8\\",\\"compressFormat\\":\\"autoDetect\\",\\"interval\\":\\"5m\\"}
	DataConnectorConfig *string `json:"DataConnectorConfig,omitempty" xml:"DataConnectorConfig,omitempty"`
	// The collector ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dc-07423146117d77db266f78bc41f4fd80
	DataConnectorId *string `json:"DataConnectorId,omitempty" xml:"DataConnectorId,omitempty"`
	// The status of the collector. Valid values:
	//
	// - enabled: enabled.
	//
	// - disabled: disabled.
	//
	// example:
	//
	// enabled
	DataConnectorStatus *string `json:"DataConnectorStatus,omitempty" xml:"DataConnectorStatus,omitempty"`
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The data management center of the threat analysis feature. Specify this parameter based on the region where your assets reside. Valid values:
	//
	// - cn-hangzhou: Your assets belong to the Chinese mainland and Hong Kong (China).
	//
	// - ap-southeast-1: Your assets belong to regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the member account that the administrator switches to.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataConnectorRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataConnectorRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataConnectorRequest) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *UpdateDataConnectorRequest) GetAuthConfigProduct() *string {
	return s.AuthConfigProduct
}

func (s *UpdateDataConnectorRequest) GetAuthConfigVendor() *string {
	return s.AuthConfigVendor
}

func (s *UpdateDataConnectorRequest) GetDataConnectorConfig() *string {
	return s.DataConnectorConfig
}

func (s *UpdateDataConnectorRequest) GetDataConnectorId() *string {
	return s.DataConnectorId
}

func (s *UpdateDataConnectorRequest) GetDataConnectorStatus() *string {
	return s.DataConnectorStatus
}

func (s *UpdateDataConnectorRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataConnectorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataConnectorRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataConnectorRequest) SetAuthConfigId(v string) *UpdateDataConnectorRequest {
	s.AuthConfigId = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetAuthConfigProduct(v string) *UpdateDataConnectorRequest {
	s.AuthConfigProduct = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetAuthConfigVendor(v string) *UpdateDataConnectorRequest {
	s.AuthConfigVendor = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetDataConnectorConfig(v string) *UpdateDataConnectorRequest {
	s.DataConnectorConfig = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetDataConnectorId(v string) *UpdateDataConnectorRequest {
	s.DataConnectorId = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetDataConnectorStatus(v string) *UpdateDataConnectorRequest {
	s.DataConnectorStatus = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetLang(v string) *UpdateDataConnectorRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetRegionId(v string) *UpdateDataConnectorRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataConnectorRequest) SetRoleFor(v int64) *UpdateDataConnectorRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataConnectorRequest) Validate() error {
	return dara.Validate(s)
}
