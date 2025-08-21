// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeAICInstanceImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpgradeAICInstanceImageRequest
	GetImageId() *string
	SetServerIds(v []*string) *UpgradeAICInstanceImageRequest
	GetServerIds() []*string
	SetTimeout(v int32) *UpgradeAICInstanceImageRequest
	GetTimeout() *int32
}

type UpgradeAICInstanceImageRequest struct {
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
	ServerIds []*string `json:"ServerIds,omitempty" xml:"ServerIds,omitempty" type:"Repeated"`
	// The timeout period of the update. Unit: seconds.
	//
	// example:
	//
	// 1800
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s UpgradeAICInstanceImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeAICInstanceImageRequest) GoString() string {
	return s.String()
}

func (s *UpgradeAICInstanceImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpgradeAICInstanceImageRequest) GetServerIds() []*string {
	return s.ServerIds
}

func (s *UpgradeAICInstanceImageRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpgradeAICInstanceImageRequest) SetImageId(v string) *UpgradeAICInstanceImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpgradeAICInstanceImageRequest) SetServerIds(v []*string) *UpgradeAICInstanceImageRequest {
	s.ServerIds = v
	return s
}

func (s *UpgradeAICInstanceImageRequest) SetTimeout(v int32) *UpgradeAICInstanceImageRequest {
	s.Timeout = &v
	return s
}

func (s *UpgradeAICInstanceImageRequest) Validate() error {
	return dara.Validate(s)
}
