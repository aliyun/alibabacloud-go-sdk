// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionEuCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v string) *ActiveInteractionEuCreateRequest
	GetImage() *string
}

type ActiveInteractionEuCreateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://linkscreen-daily.oss-cn-hangzhou.aliyuncs.com/mock/251023/a_person_12.jpg
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
}

func (s ActiveInteractionEuCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionEuCreateRequest) GoString() string {
	return s.String()
}

func (s *ActiveInteractionEuCreateRequest) GetImage() *string {
	return s.Image
}

func (s *ActiveInteractionEuCreateRequest) SetImage(v string) *ActiveInteractionEuCreateRequest {
	s.Image = &v
	return s
}

func (s *ActiveInteractionEuCreateRequest) Validate() error {
	return dara.Validate(s)
}
