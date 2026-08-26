// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteImageRequest
	GetId() *string
}

type DeleteImageRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageRequest) GetId() *string {
	return s.Id
}

func (s *DeleteImageRequest) SetId(v string) *DeleteImageRequest {
	s.Id = &v
	return s
}

func (s *DeleteImageRequest) Validate() error {
	return dara.Validate(s)
}
