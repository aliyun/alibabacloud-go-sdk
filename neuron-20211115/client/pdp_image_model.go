// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPdpImage interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *PdpImage
	GetImageId() *string
	SetImageTag(v string) *PdpImage
	GetImageTag() *string
}

type PdpImage struct {
	// This parameter is required.
	//
	// example:
	//
	// i-v12wpq
	ImageId *string `json:"imageId,omitempty" xml:"imageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20220421100717_prepub
	ImageTag *string `json:"imageTag,omitempty" xml:"imageTag,omitempty"`
}

func (s PdpImage) String() string {
	return dara.Prettify(s)
}

func (s PdpImage) GoString() string {
	return s.String()
}

func (s *PdpImage) GetImageId() *string {
	return s.ImageId
}

func (s *PdpImage) GetImageTag() *string {
	return s.ImageTag
}

func (s *PdpImage) SetImageId(v string) *PdpImage {
	s.ImageId = &v
	return s
}

func (s *PdpImage) SetImageTag(v string) *PdpImage {
	s.ImageTag = &v
	return s
}

func (s *PdpImage) Validate() error {
	return dara.Validate(s)
}
