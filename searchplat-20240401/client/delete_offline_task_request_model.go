// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOfflineTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteOfflineTaskRequest
	GetRegionId() *string
}

type DeleteOfflineTaskRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteOfflineTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOfflineTaskRequest) GoString() string {
	return s.String()
}

func (s *DeleteOfflineTaskRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteOfflineTaskRequest) SetRegionId(v string) *DeleteOfflineTaskRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteOfflineTaskRequest) Validate() error {
	return dara.Validate(s)
}
