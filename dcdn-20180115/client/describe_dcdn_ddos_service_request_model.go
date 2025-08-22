// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDdosServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeDcdnDdosServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeDcdnDdosServiceRequest
	GetSecurityToken() *string
}

type DescribeDcdnDdosServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeDcdnDdosServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDdosServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDdosServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDcdnDdosServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeDcdnDdosServiceRequest) SetOwnerId(v int64) *DescribeDcdnDdosServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDcdnDdosServiceRequest) SetSecurityToken(v string) *DescribeDcdnDdosServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeDcdnDdosServiceRequest) Validate() error {
	return dara.Validate(s)
}
