// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomImageRequest
	GetClientToken() *string
	SetImageId(v string) *DeleteCustomImageRequest
	GetImageId() *string
	SetRegionId(v string) *DeleteCustomImageRequest
	GetRegionId() *string
}

type DeleteCustomImageRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value on your client to ensure that the value is unique across different requests. **ClientToken*	- supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomImageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DeleteCustomImageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomImageRequest) SetClientToken(v string) *DeleteCustomImageRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomImageRequest) SetImageId(v string) *DeleteCustomImageRequest {
	s.ImageId = &v
	return s
}

func (s *DeleteCustomImageRequest) SetRegionId(v string) *DeleteCustomImageRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomImageRequest) Validate() error {
	return dara.Validate(s)
}
