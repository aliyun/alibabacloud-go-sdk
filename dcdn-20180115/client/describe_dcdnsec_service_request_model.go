// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnsecServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnsecServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnsecServiceRequest
	GetSecurityToken() *string
}

type DescribeDcdnsecServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnsecServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnsecServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnsecServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnsecServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnsecServiceRequest) SetOwnerId(v int64) *DescribeDcdnsecServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnsecServiceRequest) SetSecurityToken(v string) *DescribeDcdnsecServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnsecServiceRequest) Validate() error {
	return dara.Validate(s)
}
