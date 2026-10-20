// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSliderCaptchaPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSliderCaptchaPageResponseBody
	GetRequestId() *string
}

type UpdateSliderCaptchaPageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSliderCaptchaPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSliderCaptchaPageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSliderCaptchaPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSliderCaptchaPageResponseBody) SetRequestId(v string) *UpdateSliderCaptchaPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSliderCaptchaPageResponseBody) Validate() error {
	return dara.Validate(s)
}
