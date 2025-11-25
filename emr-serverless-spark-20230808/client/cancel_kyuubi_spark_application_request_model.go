// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKyuubiSparkApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CancelKyuubiSparkApplicationRequest
	GetRegionId() *string
}

type CancelKyuubiSparkApplicationRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CancelKyuubiSparkApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelKyuubiSparkApplicationRequest) GoString() string {
	return s.String()
}

func (s *CancelKyuubiSparkApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelKyuubiSparkApplicationRequest) SetRegionId(v string) *CancelKyuubiSparkApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *CancelKyuubiSparkApplicationRequest) Validate() error {
	return dara.Validate(s)
}
