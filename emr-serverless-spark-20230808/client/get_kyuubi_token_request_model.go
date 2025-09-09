// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKyuubiTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetKyuubiTokenRequest
	GetRegionId() *string
}

type GetKyuubiTokenRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetKyuubiTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKyuubiTokenRequest) GoString() string {
	return s.String()
}

func (s *GetKyuubiTokenRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKyuubiTokenRequest) SetRegionId(v string) *GetKyuubiTokenRequest {
	s.RegionId = &v
	return s
}

func (s *GetKyuubiTokenRequest) Validate() error {
	return dara.Validate(s)
}
