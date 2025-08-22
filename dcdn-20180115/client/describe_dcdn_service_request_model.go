// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnServiceRequest
	GetSecurityToken() *string
}

type DescribeDcdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnServiceRequest) SetOwnerId(v int64) *DescribeDcdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnServiceRequest) SetSecurityToken(v string) *DescribeDcdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnServiceRequest) Validate() error {
	return dara.Validate(s)
}
