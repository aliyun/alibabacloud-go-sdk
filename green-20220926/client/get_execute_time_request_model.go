// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetExecuteTimeRequest
	GetRegionId() *string
}

type GetExecuteTimeRequest struct {
	// Region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetExecuteTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteTimeRequest) GoString() string {
	return s.String()
}

func (s *GetExecuteTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetExecuteTimeRequest) SetRegionId(v string) *GetExecuteTimeRequest {
	s.RegionId = &v
	return s
}

func (s *GetExecuteTimeRequest) Validate() error {
	return dara.Validate(s)
}
