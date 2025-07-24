// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetJobRunRequest
	GetRegionId() *string
}

type GetJobRunRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetJobRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRunRequest) GoString() string {
	return s.String()
}

func (s *GetJobRunRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetJobRunRequest) SetRegionId(v string) *GetJobRunRequest {
	s.RegionId = &v
	return s
}

func (s *GetJobRunRequest) Validate() error {
	return dara.Validate(s)
}
