// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelServiceRegistrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelServiceRegistrationRequest
	GetClientToken() *string
	SetRegionId(v string) *CancelServiceRegistrationRequest
	GetRegionId() *string
	SetRegistrationId(v string) *CancelServiceRegistrationRequest
	GetRegistrationId() *string
}

type CancelServiceRegistrationRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique value from your client. This value must be unique for each request and can contain only ASCII characters.
	//
	// example:
	//
	// AAAAAWns8w4MmhzeptXVRG0PUEU=
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the service review request.
	//
	// This parameter is required.
	//
	// example:
	//
	// sr-540930183f93xxxxxx
	RegistrationId *string `json:"RegistrationId,omitempty" xml:"RegistrationId,omitempty"`
}

func (s CancelServiceRegistrationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelServiceRegistrationRequest) GoString() string {
	return s.String()
}

func (s *CancelServiceRegistrationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelServiceRegistrationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelServiceRegistrationRequest) GetRegistrationId() *string {
	return s.RegistrationId
}

func (s *CancelServiceRegistrationRequest) SetClientToken(v string) *CancelServiceRegistrationRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetRegionId(v string) *CancelServiceRegistrationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelServiceRegistrationRequest) SetRegistrationId(v string) *CancelServiceRegistrationRequest {
	s.RegistrationId = &v
	return s
}

func (s *CancelServiceRegistrationRequest) Validate() error {
	return dara.Validate(s)
}
