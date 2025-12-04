// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErRequest interface {
	dara.Model
	String() string
	GoString() string
	SetErId(v string) *DeleteErRequest
	GetErId() *string
	SetRegionId(v string) *DeleteErRequest
	GetRegionId() *string
}

type DeleteErRequest struct {
	// Lingjun HUB Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// er-kkopgtne
	ErId *string `json:"ErId,omitempty" xml:"ErId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteErRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteErRequest) GoString() string {
	return s.String()
}

func (s *DeleteErRequest) GetErId() *string {
	return s.ErId
}

func (s *DeleteErRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteErRequest) SetErId(v string) *DeleteErRequest {
	s.ErId = &v
	return s
}

func (s *DeleteErRequest) SetRegionId(v string) *DeleteErRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteErRequest) Validate() error {
	return dara.Validate(s)
}
