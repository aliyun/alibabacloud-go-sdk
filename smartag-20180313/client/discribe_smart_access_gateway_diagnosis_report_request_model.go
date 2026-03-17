// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiscribeSmartAccessGatewayDiagnosisReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest
	GetSmartAGSn() *string
}

type DiscribeSmartAccessGatewayDiagnosisReportRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device.
	//
	// This parameter is required.
	//
	// example:
	//
	// sage62x022502****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DiscribeSmartAccessGatewayDiagnosisReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DiscribeSmartAccessGatewayDiagnosisReportRequest) GoString() string {
	return s.String()
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetOwnerAccount(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetOwnerId(v int64) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.OwnerId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetRegionId(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.RegionId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetResourceOwnerAccount(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetResourceOwnerId(v int64) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetSmartAGId(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.SmartAGId = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) SetSmartAGSn(v string) *DiscribeSmartAccessGatewayDiagnosisReportRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DiscribeSmartAccessGatewayDiagnosisReportRequest) Validate() error {
	return dara.Validate(s)
}
