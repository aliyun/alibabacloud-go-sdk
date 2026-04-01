// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncRCKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *SyncRCKeyPairRequest
	GetKeyPairName() *string
	SetRegionId(v string) *SyncRCKeyPairRequest
	GetRegionId() *string
	SetSyncMode(v bool) *SyncRCKeyPairRequest
	GetSyncMode() *bool
}

type SyncRCKeyPairRequest struct {
	// The name of the key pair.
	//
	// example:
	//
	// customer_keypairs
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SyncMode *bool   `json:"SyncMode,omitempty" xml:"SyncMode,omitempty"`
}

func (s SyncRCKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s SyncRCKeyPairRequest) GoString() string {
	return s.String()
}

func (s *SyncRCKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *SyncRCKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SyncRCKeyPairRequest) GetSyncMode() *bool {
	return s.SyncMode
}

func (s *SyncRCKeyPairRequest) SetKeyPairName(v string) *SyncRCKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *SyncRCKeyPairRequest) SetRegionId(v string) *SyncRCKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *SyncRCKeyPairRequest) SetSyncMode(v bool) *SyncRCKeyPairRequest {
	s.SyncMode = &v
	return s
}

func (s *SyncRCKeyPairRequest) Validate() error {
	return dara.Validate(s)
}
