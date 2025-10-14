// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGdnInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGdnInstanceName(v string) *DeleteGdnInstanceRequest
	GetGdnInstanceName() *string
	SetRegionId(v string) *DeleteGdnInstanceRequest
	GetRegionId() *string
}

type DeleteGdnInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gdn-yq****dorhkzttoi
	GdnInstanceName *string `json:"GdnInstanceName,omitempty" xml:"GdnInstanceName,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteGdnInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteGdnInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteGdnInstanceRequest) GetGdnInstanceName() *string {
	return s.GdnInstanceName
}

func (s *DeleteGdnInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteGdnInstanceRequest) SetGdnInstanceName(v string) *DeleteGdnInstanceRequest {
	s.GdnInstanceName = &v
	return s
}

func (s *DeleteGdnInstanceRequest) SetRegionId(v string) *DeleteGdnInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteGdnInstanceRequest) Validate() error {
	return dara.Validate(s)
}
