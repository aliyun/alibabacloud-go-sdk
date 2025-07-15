// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllInstanceIdListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetAllInstanceIdListRequest
	GetRegionId() *string
}

type GetAllInstanceIdListRequest struct {
	// The region ID of the instance. This parameter is reserved.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetAllInstanceIdListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAllInstanceIdListRequest) GoString() string {
	return s.String()
}

func (s *GetAllInstanceIdListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAllInstanceIdListRequest) SetRegionId(v string) *GetAllInstanceIdListRequest {
	s.RegionId = &v
	return s
}

func (s *GetAllInstanceIdListRequest) Validate() error {
	return dara.Validate(s)
}
