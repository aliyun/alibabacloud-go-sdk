// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContainerConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v string) *ContainerConfiguration
	GetImage() *string
}

type ContainerConfiguration struct {
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
}

func (s ContainerConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ContainerConfiguration) GoString() string {
	return s.String()
}

func (s *ContainerConfiguration) GetImage() *string {
	return s.Image
}

func (s *ContainerConfiguration) SetImage(v string) *ContainerConfiguration {
	s.Image = &v
	return s
}

func (s *ContainerConfiguration) Validate() error {
	return dara.Validate(s)
}
