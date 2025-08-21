// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAICInstanceImageShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeAICInstanceImageShrinkRequest
	GetImageId() *string
	SetServerIdsShrink(v string) *UpgradeAICInstanceImageShrinkRequest
	GetServerIdsShrink() *string
	SetTimeout(v int32) *UpgradeAICInstanceImageShrinkRequest
	GetTimeout() *int32
}

type UpgradeAICInstanceImageShrinkRequest struct {
	// The ID of the AIC image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The IDs of the servers.
	//
	// This parameter is required.
	ServerIdsShrink *string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty"`
	// The timeout period of the update. Unit: seconds.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpgradeAICInstanceImageShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAICInstanceImageShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAICInstanceImageShrinkRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeAICInstanceImageShrinkRequest) GetServerIdsShrink() *string {
	return s.ServerIdsShrink
}

func (s *UpgradeAICInstanceImageShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpgradeAICInstanceImageShrinkRequest) SetImageId(v string) *UpgradeAICInstanceImageShrinkRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeAICInstanceImageShrinkRequest) SetServerIdsShrink(v string) *UpgradeAICInstanceImageShrinkRequest {
	s.ServerIdsShrink = &v
	return s
}

func (s *UpgradeAICInstanceImageShrinkRequest) SetTimeout(v int32) *UpgradeAICInstanceImageShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *UpgradeAICInstanceImageShrinkRequest) Validate() error {
	return dara.Validate(s)
}
