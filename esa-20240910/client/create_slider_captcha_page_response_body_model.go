// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSliderCaptchaPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateSliderCaptchaPageResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateSliderCaptchaPageResponseBody
	GetRequestId() *string
}

type CreateSliderCaptchaPageResponseBody struct {
	// The ID of the slider CAPTCHA page.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSliderCaptchaPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSliderCaptchaPageResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSliderCaptchaPageResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateSliderCaptchaPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSliderCaptchaPageResponseBody) SetId(v int64) *CreateSliderCaptchaPageResponseBody {
	s.Id = &v
	return s
}

func (s *CreateSliderCaptchaPageResponseBody) SetRequestId(v string) *CreateSliderCaptchaPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSliderCaptchaPageResponseBody) Validate() error {
	return dara.Validate(s)
}
