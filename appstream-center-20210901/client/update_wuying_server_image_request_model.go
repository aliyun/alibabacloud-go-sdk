// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWuyingServerImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *UpdateWuyingServerImageRequest
	GetImageId() *string
	SetProductType(v string) *UpdateWuyingServerImageRequest
	GetProductType() *string
	SetWuyingServerId(v string) *UpdateWuyingServerImageRequest
	GetWuyingServerId() *string
}

type UpdateWuyingServerImageRequest struct {
	// example:
	//
	// imgc-0aae4rgk9****6e8p
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s UpdateWuyingServerImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWuyingServerImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateWuyingServerImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateWuyingServerImageRequest) GetProductType() *string {
	return s.ProductType
}

func (s *UpdateWuyingServerImageRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *UpdateWuyingServerImageRequest) SetImageId(v string) *UpdateWuyingServerImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateWuyingServerImageRequest) SetProductType(v string) *UpdateWuyingServerImageRequest {
	s.ProductType = &v
	return s
}

func (s *UpdateWuyingServerImageRequest) SetWuyingServerId(v string) *UpdateWuyingServerImageRequest {
	s.WuyingServerId = &v
	return s
}

func (s *UpdateWuyingServerImageRequest) Validate() error {
	return dara.Validate(s)
}
