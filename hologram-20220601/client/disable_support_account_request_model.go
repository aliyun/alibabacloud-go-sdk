// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSupportAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DisableSupportAccountRequest
	GetRegionId() *string
}

type DisableSupportAccountRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableSupportAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSupportAccountRequest) GoString() string {
	return s.String()
}

func (s *DisableSupportAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableSupportAccountRequest) SetRegionId(v string) *DisableSupportAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DisableSupportAccountRequest) Validate() error {
	return dara.Validate(s)
}
