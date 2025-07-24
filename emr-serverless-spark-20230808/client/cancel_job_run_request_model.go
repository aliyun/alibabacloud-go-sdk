// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelJobRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CancelJobRunRequest
	GetRegionId() *string
}

type CancelJobRunRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s CancelJobRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelJobRunRequest) GoString() string {
	return s.String()
}

func (s *CancelJobRunRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelJobRunRequest) SetRegionId(v string) *CancelJobRunRequest {
	s.RegionId = &v
	return s
}

func (s *CancelJobRunRequest) Validate() error {
	return dara.Validate(s)
}
