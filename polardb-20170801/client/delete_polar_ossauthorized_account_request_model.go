// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarOSSAuthorizedAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserIds(v string) *DeletePolarOSSAuthorizedAccountRequest
	GetAuthorizedUserIds() *string
	SetDBClusterId(v string) *DeletePolarOSSAuthorizedAccountRequest
	GetDBClusterId() *string
	SetPfsInstanceId(v string) *DeletePolarOSSAuthorizedAccountRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *DeletePolarOSSAuthorizedAccountRequest
	GetRegionId() *string
}

type DeletePolarOSSAuthorizedAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1234567890,acs:ram::123456:role/myrole
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	// example:
	//
	// pc-xxxxxxxxxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxxxxxxxxxxxxxxxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeletePolarOSSAuthorizedAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarOSSAuthorizedAccountRequest) GoString() string {
	return s.String()
}

func (s *DeletePolarOSSAuthorizedAccountRequest) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *DeletePolarOSSAuthorizedAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeletePolarOSSAuthorizedAccountRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *DeletePolarOSSAuthorizedAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePolarOSSAuthorizedAccountRequest) SetAuthorizedUserIds(v string) *DeletePolarOSSAuthorizedAccountRequest {
	s.AuthorizedUserIds = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountRequest) SetDBClusterId(v string) *DeletePolarOSSAuthorizedAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountRequest) SetPfsInstanceId(v string) *DeletePolarOSSAuthorizedAccountRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountRequest) SetRegionId(v string) *DeletePolarOSSAuthorizedAccountRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePolarOSSAuthorizedAccountRequest) Validate() error {
	return dara.Validate(s)
}
