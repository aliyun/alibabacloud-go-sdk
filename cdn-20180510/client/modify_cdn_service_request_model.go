// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetChargeType(v string) *ModifyCdnServiceRequest
	GetInternetChargeType() *string
	SetOwnerId(v int64) *ModifyCdnServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *ModifyCdnServiceRequest
	GetSecurityToken() *string
}

type ModifyCdnServiceRequest struct {
	// The new metering method for Alibaba Cloud CDN. Valid values:
	//
	// 	- **PayByTraffic**: pay-by-data-transfer
	//
	// 	- **PayByBandwidth**: pay-by-bandwidth
	//
	// This parameter is required.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken      *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyCdnServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnServiceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *ModifyCdnServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCdnServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyCdnServiceRequest) SetInternetChargeType(v string) *ModifyCdnServiceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *ModifyCdnServiceRequest) SetOwnerId(v int64) *ModifyCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCdnServiceRequest) SetSecurityToken(v string) *ModifyCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyCdnServiceRequest) Validate() error {
	return dara.Validate(s)
}
