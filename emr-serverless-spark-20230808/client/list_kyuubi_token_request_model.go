// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKyuubiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ListKyuubiTokenRequest
	GetRegionId() *string
}

type ListKyuubiTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListKyuubiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *ListKyuubiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKyuubiTokenRequest) SetRegionId(v string) *ListKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *ListKyuubiTokenRequest) Validate() error {
	return dara.Validate(s)
}
