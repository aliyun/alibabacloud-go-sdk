// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeBaaSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *AuthorizeBaaSRequest
	GetRegionId() *string
}

type AuthorizeBaaSRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AuthorizeBaaSRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeBaaSRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeBaaSRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeBaaSRequest) SetRegionId(v string) *AuthorizeBaaSRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeBaaSRequest) Validate() error {
	return dara.Validate(s)
}
