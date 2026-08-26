// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarOSSAuthorizedAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizedUserIds(v string) *AddPolarOSSAuthorizedAccountRequest
	GetAuthorizedUserIds() *string
	SetDBClusterId(v string) *AddPolarOSSAuthorizedAccountRequest
	GetDBClusterId() *string
	SetPfsInstanceId(v string) *AddPolarOSSAuthorizedAccountRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *AddPolarOSSAuthorizedAccountRequest
	GetRegionId() *string
}

type AddPolarOSSAuthorizedAccountRequest struct {
	// The list of authorized accounts to add, separated by commas. You can pass in UIDs and RAM role ARNs in mixed parameter notation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890,acs:ram::123456:role/myrole
	AuthorizedUserIds *string `json:"AuthorizedUserIds,omitempty" xml:"AuthorizedUserIds,omitempty"`
	// The ID of the PolarDB cluster.
	//
	// example:
	//
	// pc-xxxxxxxxxxxxxxxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The cold storage instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxxxxxxxxxxxxxxxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddPolarOSSAuthorizedAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPolarOSSAuthorizedAccountRequest) GoString() string {
	return s.String()
}

func (s *AddPolarOSSAuthorizedAccountRequest) GetAuthorizedUserIds() *string {
	return s.AuthorizedUserIds
}

func (s *AddPolarOSSAuthorizedAccountRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *AddPolarOSSAuthorizedAccountRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *AddPolarOSSAuthorizedAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddPolarOSSAuthorizedAccountRequest) SetAuthorizedUserIds(v string) *AddPolarOSSAuthorizedAccountRequest {
	s.AuthorizedUserIds = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountRequest) SetDBClusterId(v string) *AddPolarOSSAuthorizedAccountRequest {
	s.DBClusterId = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountRequest) SetPfsInstanceId(v string) *AddPolarOSSAuthorizedAccountRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountRequest) SetRegionId(v string) *AddPolarOSSAuthorizedAccountRequest {
	s.RegionId = &v
	return s
}

func (s *AddPolarOSSAuthorizedAccountRequest) Validate() error {
	return dara.Validate(s)
}
