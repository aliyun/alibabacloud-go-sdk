// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVodDefaultRoleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v string) *CheckVodDefaultRoleRequest
	GetOwnerId() *string
	SetResourceRealOwnerId(v int64) *CheckVodDefaultRoleRequest
	GetResourceRealOwnerId() *int64
}

type CheckVodDefaultRoleRequest struct {
	OwnerId             *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceRealOwnerId *int64  `json:"ResourceRealOwnerId,omitempty" xml:"ResourceRealOwnerId,omitempty"`
}

func (s CheckVodDefaultRoleRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckVodDefaultRoleRequest) GoString() string {
	return s.String()
}

func (s *CheckVodDefaultRoleRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *CheckVodDefaultRoleRequest) GetResourceRealOwnerId() *int64 {
	return s.ResourceRealOwnerId
}

func (s *CheckVodDefaultRoleRequest) SetOwnerId(v string) *CheckVodDefaultRoleRequest {
	s.OwnerId = &v
	return s
}

func (s *CheckVodDefaultRoleRequest) SetResourceRealOwnerId(v int64) *CheckVodDefaultRoleRequest {
	s.ResourceRealOwnerId = &v
	return s
}

func (s *CheckVodDefaultRoleRequest) Validate() error {
	return dara.Validate(s)
}
