// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSliderCaptchaPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteSliderCaptchaPageRequest
	GetId() *int64
}

type DeleteSliderCaptchaPageRequest struct {
	// The ID of the custom slider CAPTCHA page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteSliderCaptchaPageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSliderCaptchaPageRequest) GoString() string {
	return s.String()
}

func (s *DeleteSliderCaptchaPageRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteSliderCaptchaPageRequest) SetId(v int64) *DeleteSliderCaptchaPageRequest {
	s.Id = &v
	return s
}

func (s *DeleteSliderCaptchaPageRequest) Validate() error {
	return dara.Validate(s)
}
