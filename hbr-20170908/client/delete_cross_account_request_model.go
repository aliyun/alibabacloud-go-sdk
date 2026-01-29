// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrossAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *DeleteCrossAccountRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DeleteCrossAccountRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DeleteCrossAccountRequest
	GetCrossAccountUserId() *int64
}

type DeleteCrossAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1841xxxxx3649795
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s DeleteCrossAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrossAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrossAccountRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DeleteCrossAccountRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DeleteCrossAccountRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DeleteCrossAccountRequest) SetCrossAccountRoleName(v string) *DeleteCrossAccountRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DeleteCrossAccountRequest) SetCrossAccountType(v string) *DeleteCrossAccountRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DeleteCrossAccountRequest) SetCrossAccountUserId(v int64) *DeleteCrossAccountRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DeleteCrossAccountRequest) Validate() error {
	return dara.Validate(s)
}
