// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RollbackImageRequest
	GetId() *string
	SetImageVersion(v string) *RollbackImageRequest
	GetImageVersion() *string
}

type RollbackImageRequest struct {
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// Custom_image_xxxx_xxxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The image version to roll back to.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
}

func (s RollbackImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackImageRequest) GoString() string {
	return s.String()
}

func (s *RollbackImageRequest) GetId() *string {
	return s.Id
}

func (s *RollbackImageRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *RollbackImageRequest) SetId(v string) *RollbackImageRequest {
	s.Id = &v
	return s
}

func (s *RollbackImageRequest) SetImageVersion(v string) *RollbackImageRequest {
	s.ImageVersion = &v
	return s
}

func (s *RollbackImageRequest) Validate() error {
	return dara.Validate(s)
}
