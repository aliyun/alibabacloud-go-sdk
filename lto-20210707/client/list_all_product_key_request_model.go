// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProductKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListAllProductKeyRequest
	GetRegionId() *string
}

type ListAllProductKeyRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListAllProductKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAllProductKeyRequest) GoString() string {
	return s.String()
}

func (s *ListAllProductKeyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAllProductKeyRequest) SetRegionId(v string) *ListAllProductKeyRequest {
	s.RegionId = &v
	return s
}

func (s *ListAllProductKeyRequest) Validate() error {
	return dara.Validate(s)
}
