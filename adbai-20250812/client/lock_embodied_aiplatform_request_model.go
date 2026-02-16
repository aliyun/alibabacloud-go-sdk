// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *LockEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetPlatformName(v string) *LockEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRegionId(v string) *LockEmbodiedAIPlatformRequest
	GetRegionId() *string
}

type LockEmbodiedAIPlatformRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// platform1
	PlatformName *string `json:"PlatformName,omitempty" xml:"PlatformName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s LockEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s LockEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *LockEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *LockEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *LockEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LockEmbodiedAIPlatformRequest) SetDBClusterId(v string) *LockEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *LockEmbodiedAIPlatformRequest) SetPlatformName(v string) *LockEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *LockEmbodiedAIPlatformRequest) SetRegionId(v string) *LockEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *LockEmbodiedAIPlatformRequest) Validate() error {
	return dara.Validate(s)
}
