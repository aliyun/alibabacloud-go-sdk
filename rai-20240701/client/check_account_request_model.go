// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CheckAccountRequest
	GetRegionId() *string
}

type CheckAccountRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CheckAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckAccountRequest) GoString() string {
	return s.String()
}

func (s *CheckAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CheckAccountRequest) SetRegionId(v string) *CheckAccountRequest {
	s.RegionId = &v
	return s
}

func (s *CheckAccountRequest) Validate() error {
	return dara.Validate(s)
}
