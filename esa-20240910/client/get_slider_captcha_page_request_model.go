// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSliderCaptchaPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetSliderCaptchaPageRequest
	GetId() *int64
}

type GetSliderCaptchaPageRequest struct {
	// The ID of the custom slider page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSliderCaptchaPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSliderCaptchaPageRequest) GoString() string {
	return s.String()
}

func (s *GetSliderCaptchaPageRequest) GetId() *int64 {
	return s.Id
}

func (s *GetSliderCaptchaPageRequest) SetId(v int64) *GetSliderCaptchaPageRequest {
	s.Id = &v
	return s
}

func (s *GetSliderCaptchaPageRequest) Validate() error {
	return dara.Validate(s)
}
