// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenCdnServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInternetChargeType(v string) *OpenCdnServiceRequest
	GetInternetChargeType() *string
	SetOwnerId(v int64) *OpenCdnServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *OpenCdnServiceRequest
	GetSecurityToken() *string
}

type OpenCdnServiceRequest struct {
	// The metering method of Alibaba Cloud CDN. A value of **PayByTraffic*	- indicates that the metering method is pay-by-data-transfer.
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

func (s OpenCdnServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenCdnServiceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *OpenCdnServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenCdnServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *OpenCdnServiceRequest) SetInternetChargeType(v string) *OpenCdnServiceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *OpenCdnServiceRequest) SetOwnerId(v int64) *OpenCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenCdnServiceRequest) SetSecurityToken(v string) *OpenCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *OpenCdnServiceRequest) Validate() error {
	return dara.Validate(s)
}
