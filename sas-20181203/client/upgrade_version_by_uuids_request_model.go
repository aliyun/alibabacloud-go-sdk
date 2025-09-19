// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeVersionByUuidsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetUpgradeVersion(v string) *UpgradeVersionByUuidsRequest
	GetUpgradeVersion() *string
	SetUuidList(v []*string) *UpgradeVersionByUuidsRequest
	GetUuidList() []*string
}

type UpgradeVersionByUuidsRequest struct {
	// The version to which you want to upgrade the client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.9
	UpgradeVersion *string `json:"UpgradeVersion,omitempty" xml:"UpgradeVersion,omitempty"`
	// The UUIDs of the assets on which you want to run the detection task.
	//
	// This parameter is required.
	UuidList []*string `json:"UuidList,omitempty" xml:"UuidList,omitempty" type:"Repeated"`
}

func (s UpgradeVersionByUuidsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeVersionByUuidsRequest) GoString() string {
	return s.String()
}

func (s *UpgradeVersionByUuidsRequest) GetUpgradeVersion() *string {
	return s.UpgradeVersion
}

func (s *UpgradeVersionByUuidsRequest) GetUuidList() []*string {
	return s.UuidList
}

func (s *UpgradeVersionByUuidsRequest) SetUpgradeVersion(v string) *UpgradeVersionByUuidsRequest {
	s.UpgradeVersion = &v
	return s
}

func (s *UpgradeVersionByUuidsRequest) SetUuidList(v []*string) *UpgradeVersionByUuidsRequest {
	s.UuidList = v
	return s
}

func (s *UpgradeVersionByUuidsRequest) Validate() error {
	return dara.Validate(s)
}
