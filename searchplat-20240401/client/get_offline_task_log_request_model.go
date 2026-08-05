// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOfflineTaskLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetOfflineTaskLogRequest
	GetRegionId() *string
}

type GetOfflineTaskLogRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s GetOfflineTaskLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOfflineTaskLogRequest) GoString() string {
	return s.String()
}

func (s *GetOfflineTaskLogRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetOfflineTaskLogRequest) SetRegionId(v string) *GetOfflineTaskLogRequest {
	s.RegionId = &v
	return s
}

func (s *GetOfflineTaskLogRequest) Validate() error {
	return dara.Validate(s)
}
