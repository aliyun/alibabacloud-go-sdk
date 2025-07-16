// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeCdnServiceRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *DescribeCdnServiceRequest
	GetSecurityToken() *string
}

type DescribeCdnServiceRequest struct {
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeCdnServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnServiceRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCdnServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeCdnServiceRequest) SetOwnerId(v int64) *DescribeCdnServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCdnServiceRequest) SetSecurityToken(v string) *DescribeCdnServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeCdnServiceRequest) Validate() error {
	return dara.Validate(s)
}
