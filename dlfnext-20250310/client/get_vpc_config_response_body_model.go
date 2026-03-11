// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetTrustedVpcs(v []*GetVpcConfigResponseBodyTrustedVpcs) *GetVpcConfigResponseBody
	GetTrustedVpcs() []*GetVpcConfigResponseBodyTrustedVpcs
}

type GetVpcConfigResponseBody struct {
	TrustedVpcs []*GetVpcConfigResponseBodyTrustedVpcs `json:"trustedVpcs,omitempty" xml:"trustedVpcs,omitempty" type:"Repeated"`
}

func (s GetVpcConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVpcConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcConfigResponseBody) GetTrustedVpcs() []*GetVpcConfigResponseBodyTrustedVpcs {
	return s.TrustedVpcs
}

func (s *GetVpcConfigResponseBody) SetTrustedVpcs(v []*GetVpcConfigResponseBodyTrustedVpcs) *GetVpcConfigResponseBody {
	s.TrustedVpcs = v
	return s
}

func (s *GetVpcConfigResponseBody) Validate() error {
	if s.TrustedVpcs != nil {
		for _, item := range s.TrustedVpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetVpcConfigResponseBodyTrustedVpcs struct {
	// example:
	//
	// 1744970111419
	CreatedAt       *int64             `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	ExtendedOptions map[string]*string `json:"extendedOptions,omitempty" xml:"extendedOptions,omitempty"`
	// VPC ID。
	//
	// example:
	//
	// vpc-uf67xxxx
	VpcId *string `json:"vpcId,omitempty" xml:"vpcId,omitempty"`
}

func (s GetVpcConfigResponseBodyTrustedVpcs) String() string {
	return dara.Prettify(s)
}

func (s GetVpcConfigResponseBodyTrustedVpcs) GoString() string {
	return s.String()
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) GetExtendedOptions() map[string]*string {
	return s.ExtendedOptions
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) SetCreatedAt(v int64) *GetVpcConfigResponseBodyTrustedVpcs {
	s.CreatedAt = &v
	return s
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) SetExtendedOptions(v map[string]*string) *GetVpcConfigResponseBodyTrustedVpcs {
	s.ExtendedOptions = v
	return s
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) SetVpcId(v string) *GetVpcConfigResponseBodyTrustedVpcs {
	s.VpcId = &v
	return s
}

func (s *GetVpcConfigResponseBodyTrustedVpcs) Validate() error {
	return dara.Validate(s)
}
