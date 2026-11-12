// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCustomImagesRequest
	GetClientToken() *string
	SetImageIds(v string) *DeleteCustomImagesRequest
	GetImageIds() *string
	SetRegionId(v string) *DeleteCustomImagesRequest
	GetRegionId() *string
}

type DeleteCustomImagesRequest struct {
	// Ensures the idempotence of the request. Generate the value from your client and ensure that it is unique across different requests. **ClientToken*	- supports only ASCII characters and can be up to 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance IDs of the simple application servers. The value can be a JSON array that consists of multiple instance IDs. You can specify up to 15 IDs. Separate multiple IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["m-2zehv38jjmwva1ee****", "m-bp1hj0zhmheyq2kz****"]
	ImageIds *string `json:"ImageIds,omitempty" xml:"ImageIds,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCustomImagesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImagesRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCustomImagesRequest) GetImageIds() *string {
	return s.ImageIds
}

func (s *DeleteCustomImagesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomImagesRequest) SetClientToken(v string) *DeleteCustomImagesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCustomImagesRequest) SetImageIds(v string) *DeleteCustomImagesRequest {
	s.ImageIds = &v
	return s
}

func (s *DeleteCustomImagesRequest) SetRegionId(v string) *DeleteCustomImagesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomImagesRequest) Validate() error {
	return dara.Validate(s)
}
