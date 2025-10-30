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
	SetResourceGroupId(v string) *CreateExternalCACertificateShrinkRequest
	GetResourceGroupId() *string
	SetTags(v []*CreateExternalCACertificateShrinkRequestTags) *CreateExternalCACertificateShrinkRequest
	GetTags() []*CreateExternalCACertificateShrinkRequestTags
	SetValidity(v string) *CreateExternalCACertificateShrinkRequest
	GetValidity() *string
}

type CreateExternalCACertificateShrinkRequest struct {
	ApiPassthroughShrink *string `json:"ApiPassthrough,omitempty" xml:"ApiPassthrough,omitempty"`
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
	// example:
	//
	// cas_deposit-cn-1234abcd
	InstanceId      *string                                         `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ResourceGroupId *string                                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*CreateExternalCACertificateShrinkRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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

func (s *CreateExternalCACertificateShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateExternalCACertificateShrinkRequest) GetTags() []*CreateExternalCACertificateShrinkRequestTags {
	return s.Tags
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

func (s *CreateExternalCACertificateShrinkRequest) SetResourceGroupId(v string) *CreateExternalCACertificateShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) SetTags(v []*CreateExternalCACertificateShrinkRequestTags) *CreateExternalCACertificateShrinkRequest {
	s.Tags = v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) SetValidity(v string) *CreateExternalCACertificateShrinkRequest {
	s.Validity = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateExternalCACertificateShrinkRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateExternalCACertificateShrinkRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateShrinkRequestTags) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateShrinkRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreateExternalCACertificateShrinkRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreateExternalCACertificateShrinkRequestTags) SetKey(v string) *CreateExternalCACertificateShrinkRequestTags {
	s.Key = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequestTags) SetValue(v string) *CreateExternalCACertificateShrinkRequestTags {
	s.Value = &v
	return s
}

func (s *CreateExternalCACertificateShrinkRequestTags) Validate() error {
	return dara.Validate(s)
}
