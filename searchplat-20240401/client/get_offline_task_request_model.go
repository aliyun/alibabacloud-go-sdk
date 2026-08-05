// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetOfflineTaskRequest
	GetRegionId() *string
}

type GetOfflineTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOfflineTaskRequest) SetRegionId(v string) *GetOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *GetOfflineTaskRequest) Validate() error {
	return dara.Validate(s)
}
