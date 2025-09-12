// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalCACertificateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiPassthroughShrink(v string) *CreateExternalCACertificateShrinkRequest
	GetApiPassthroughShrink() *string
	SetCsr(v string) *CreateExternalCACertificateShrinkRequest
	GetCsr() *string
	SetInstanceId(v string) *CreateExternalCACertificateShrinkRequest
	GetInstanceId() *string
	SetValidity(v string) *CreateExternalCACertificateShrinkRequest
	GetValidity() *string
}

type CreateExternalCACertificateShrinkRequest struct {
	ApiPassthroughShrink *string `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// vbIgMQIhAKHDWD6/WAMbtezAt4bysJ/BZIDz1jPWuUR5GV4TJ/mS
	//
	// -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cas_deposit-cn-1234abcd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10y
	Validity *string `json:"Validity,omitempty" xml:"Validity,omitempty"`
}

func (s CreateExternalCACertificateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateShrinkRequest) GetApiPassthroughShrink() *string {
	return s.ApiPassthroughShrink
}

func (s *CreateExternalCACertificateShrinkRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateExternalCACertificateShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateExternalCACertificateShrinkRequest) GetValidity() *string {
	return s.Validity
}

func (s *CreateExternalCACertificateShrinkRequest) SetApiPassthroughShrink(v string) *CreateExternalCACertificateShrinkRequest {
	s.ApiPassthroughShrink = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) SetCsr(v string) *CreateExternalCACertificateShrinkRequest {
	s.Csr = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) SetInstanceId(v string) *CreateExternalCACertificateShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) SetValidity(v string) *CreateExternalCACertificateShrinkRequest {
	s.Validity = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
