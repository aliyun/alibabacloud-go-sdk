// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCrossBorderApprovalStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *QueryCrossBorderApprovalStatusRequest
	GetRegionId() *string
}

type QueryCrossBorderApprovalStatusRequest struct {
	// The region ID of the GA instance. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s QueryCrossBorderApprovalStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCrossBorderApprovalStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryCrossBorderApprovalStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCrossBorderApprovalStatusRequest) SetRegionId(v string) *QueryCrossBorderApprovalStatusRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCrossBorderApprovalStatusRequest) Validate() error {
	return dara.Validate(s)
}
