// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVsPullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateVsPullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type UpdateVsPullStreamInfoConfigResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateVsPullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVsPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVsPullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVsPullStreamInfoConfigResponseBody) SetRequestId(v string) *UpdateVsPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVsPullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
