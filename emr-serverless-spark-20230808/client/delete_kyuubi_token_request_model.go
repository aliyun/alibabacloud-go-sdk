// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKyuubiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteKyuubiTokenRequest
	GetRegionId() *string
}

type DeleteKyuubiTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteKyuubiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *DeleteKyuubiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteKyuubiTokenRequest) SetRegionId(v string) *DeleteKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteKyuubiTokenRequest) Validate() error {
	return dara.Validate(s)
}
