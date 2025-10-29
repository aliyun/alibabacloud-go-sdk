// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHaVipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHaVipIds(v []*string) *DeleteHaVipsRequest
	GetHaVipIds() []*string
}

type DeleteHaVipsRequest struct {
	// The IDs of high-availability virtual IP addresses (HAVIPs).
	//
	// This parameter is required.
	HaVipIds []*string `json:"HaVipIds,omitempty" xml:"HaVipIds,omitempty" type:"Repeated"`
}

func (s DeleteHaVipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteHaVipsRequest) GoString() string {
	return s.String()
}

func (s *DeleteHaVipsRequest) GetHaVipIds() []*string {
	return s.HaVipIds
}

func (s *DeleteHaVipsRequest) SetHaVipIds(v []*string) *DeleteHaVipsRequest {
	s.HaVipIds = v
	return s
}

func (s *DeleteHaVipsRequest) Validate() error {
	return dara.Validate(s)
}
