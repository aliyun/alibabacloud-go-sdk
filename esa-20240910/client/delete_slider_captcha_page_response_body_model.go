// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSliderCaptchaPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSliderCaptchaPageResponseBody
	GetRequestId() *string
}

type DeleteSliderCaptchaPageResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSliderCaptchaPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSliderCaptchaPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSliderCaptchaPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSliderCaptchaPageResponseBody) SetRequestId(v string) *DeleteSliderCaptchaPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSliderCaptchaPageResponseBody) Validate() error {
	return dara.Validate(s)
}
