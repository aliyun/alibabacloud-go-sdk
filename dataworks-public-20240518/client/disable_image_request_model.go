// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DisableImageRequest
	GetId() *string
}

type DisableImageRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DisableImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableImageRequest) GoString() string {
	return s.String()
}

func (s *DisableImageRequest) GetId() *string {
	return s.Id
}

func (s *DisableImageRequest) SetId(v string) *DisableImageRequest {
	s.Id = &v
	return s
}

func (s *DisableImageRequest) Validate() error {
	return dara.Validate(s)
}
