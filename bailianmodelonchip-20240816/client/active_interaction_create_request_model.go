// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveInteractionCreateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImage(v string) *ActiveInteractionCreateRequest
	GetImage() *string
}

type ActiveInteractionCreateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://linkscreen-daily.oss-cn-hangzhou.aliyuncs.com/mock/251023/a_person_1.jpg
	Image *string `json:"image,omitempty" xml:"image,omitempty"`
}

func (s ActiveInteractionCreateRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveInteractionCreateRequest) GoString() string {
	return s.String()
}

func (s *ActiveInteractionCreateRequest) GetImage() *string {
	return s.Image
}

func (s *ActiveInteractionCreateRequest) SetImage(v string) *ActiveInteractionCreateRequest {
	s.Image = &v
	return s
}

func (s *ActiveInteractionCreateRequest) Validate() error {
	return dara.Validate(s)
}
