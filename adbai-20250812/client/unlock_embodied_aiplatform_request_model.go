// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockEmbodiedAIPlatformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *UnlockEmbodiedAIPlatformRequest
	GetDBClusterId() *string
	SetPlatformName(v string) *UnlockEmbodiedAIPlatformRequest
	GetPlatformName() *string
	SetRegionId(v string) *UnlockEmbodiedAIPlatformRequest
	GetRegionId() *string
}

type UnlockEmbodiedAIPlatformRequest struct {
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

func (s UnlockEmbodiedAIPlatformRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockEmbodiedAIPlatformRequest) GoString() string {
	return s.String()
}

func (s *UnlockEmbodiedAIPlatformRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UnlockEmbodiedAIPlatformRequest) GetPlatformName() *string {
	return s.PlatformName
}

func (s *UnlockEmbodiedAIPlatformRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnlockEmbodiedAIPlatformRequest) SetDBClusterId(v string) *UnlockEmbodiedAIPlatformRequest {
	s.DBClusterId = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) SetPlatformName(v string) *UnlockEmbodiedAIPlatformRequest {
	s.PlatformName = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) SetRegionId(v string) *UnlockEmbodiedAIPlatformRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockEmbodiedAIPlatformRequest) Validate() error {
	return dara.Validate(s)
}
