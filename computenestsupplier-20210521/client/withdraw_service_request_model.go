// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWithdrawServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *WithdrawServiceRequest
	GetClientToken() *string
	SetRegionId(v string) *WithdrawServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *WithdrawServiceRequest
	GetServiceId() *string
	SetServiceVersion(v string) *WithdrawServiceRequest
	GetServiceVersion() *string
}

type WithdrawServiceRequest struct {
	// Client token, used to ensure the idempotence of requests. Generate a unique value for this parameter from your client to ensure it is unique across different requests. ClientToken supports only ASCII characters.
	//
	// example:
	//
	// 788E7CP0EN9D51P
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// service-0e6fca6a51a544xxxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// Service version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	ServiceVersion *string `json:"ServiceVersion,omitempty" xml:"ServiceVersion,omitempty"`
}

func (s WithdrawServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s WithdrawServiceRequest) GoString() string {
	return s.String()
}

func (s *WithdrawServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *WithdrawServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *WithdrawServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *WithdrawServiceRequest) GetServiceVersion() *string {
	return s.ServiceVersion
}

func (s *WithdrawServiceRequest) SetClientToken(v string) *WithdrawServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *WithdrawServiceRequest) SetRegionId(v string) *WithdrawServiceRequest {
	s.RegionId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceId(v string) *WithdrawServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *WithdrawServiceRequest) SetServiceVersion(v string) *WithdrawServiceRequest {
	s.ServiceVersion = &v
	return s
}

func (s *WithdrawServiceRequest) Validate() error {
	return dara.Validate(s)
}
