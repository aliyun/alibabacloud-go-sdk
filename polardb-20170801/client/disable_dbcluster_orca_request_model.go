// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterOrcaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCleanData(v string) *DisableDBClusterOrcaRequest
	GetCleanData() *string
	SetDBClusterId(v string) *DisableDBClusterOrcaRequest
	GetDBClusterId() *string
	SetOwnerAccount(v string) *DisableDBClusterOrcaRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DisableDBClusterOrcaRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DisableDBClusterOrcaRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DisableDBClusterOrcaRequest
	GetResourceOwnerId() *int64
}

type DisableDBClusterOrcaRequest struct {
	// example:
	//
	// Enable
	CleanData *string `json:"CleanData,omitempty" xml:"CleanData,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId          *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DisableDBClusterOrcaRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterOrcaRequest) GoString() string {
	return s.String()
}

func (s *DisableDBClusterOrcaRequest) GetCleanData() *string {
	return s.CleanData
}

func (s *DisableDBClusterOrcaRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DisableDBClusterOrcaRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DisableDBClusterOrcaRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DisableDBClusterOrcaRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DisableDBClusterOrcaRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DisableDBClusterOrcaRequest) SetCleanData(v string) *DisableDBClusterOrcaRequest {
	s.CleanData = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) SetDBClusterId(v string) *DisableDBClusterOrcaRequest {
	s.DBClusterId = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) SetOwnerAccount(v string) *DisableDBClusterOrcaRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) SetOwnerId(v int64) *DisableDBClusterOrcaRequest {
	s.OwnerId = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) SetResourceOwnerAccount(v string) *DisableDBClusterOrcaRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) SetResourceOwnerId(v int64) *DisableDBClusterOrcaRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DisableDBClusterOrcaRequest) Validate() error {
	return dara.Validate(s)
}
