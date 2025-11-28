// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllAdminRequest
	GetRegionId() *string
}

type ListAllAdminRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllAdminRequest) GoString() string {
	return s.String()
}

func (s *ListAllAdminRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllAdminRequest) SetRegionId(v string) *ListAllAdminRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllAdminRequest) Validate() error {
	return dara.Validate(s)
}
