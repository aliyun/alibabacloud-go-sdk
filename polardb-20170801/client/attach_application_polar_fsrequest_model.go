// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplicationPolarFSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AttachApplicationPolarFSRequest
	GetApplicationId() *string
	SetPolarFSAccessKeyId(v string) *AttachApplicationPolarFSRequest
	GetPolarFSAccessKeyId() *string
	SetPolarFSAccessKeySecret(v string) *AttachApplicationPolarFSRequest
	GetPolarFSAccessKeySecret() *string
	SetPolarFSInstanceId(v string) *AttachApplicationPolarFSRequest
	GetPolarFSInstanceId() *string
}

type AttachApplicationPolarFSRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// LT**********************
	PolarFSAccessKeyId *string `json:"PolarFSAccessKeyId,omitempty" xml:"PolarFSAccessKeyId,omitempty"`
	// example:
	//
	// H3****************************
	PolarFSAccessKeySecret *string `json:"PolarFSAccessKeySecret,omitempty" xml:"PolarFSAccessKeySecret,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pcs-**************
	PolarFSInstanceId *string `json:"PolarFSInstanceId,omitempty" xml:"PolarFSInstanceId,omitempty"`
}

func (s AttachApplicationPolarFSRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachApplicationPolarFSRequest) GoString() string {
	return s.String()
}

func (s *AttachApplicationPolarFSRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSAccessKeyId() *string {
	return s.PolarFSAccessKeyId
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSAccessKeySecret() *string {
	return s.PolarFSAccessKeySecret
}

func (s *AttachApplicationPolarFSRequest) GetPolarFSInstanceId() *string {
	return s.PolarFSInstanceId
}

func (s *AttachApplicationPolarFSRequest) SetApplicationId(v string) *AttachApplicationPolarFSRequest {
	s.ApplicationId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSAccessKeyId(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSAccessKeyId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSAccessKeySecret(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSAccessKeySecret = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) SetPolarFSInstanceId(v string) *AttachApplicationPolarFSRequest {
	s.PolarFSInstanceId = &v
	return s
}

func (s *AttachApplicationPolarFSRequest) Validate() error {
	return dara.Validate(s)
}
