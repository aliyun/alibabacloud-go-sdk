// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOssScanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOssScanConfigResponseBody
	GetRequestId() *string
}

type UpdateOssScanConfigResponseBody struct {
	// example:
	//
	// E99D386F-5546-5BCD-BADD-F4EF4B******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOssScanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOssScanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOssScanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOssScanConfigResponseBody) SetRequestId(v string) *UpdateOssScanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOssScanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
