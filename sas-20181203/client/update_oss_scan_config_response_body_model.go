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
	// The request ID.
	//
	// example:
	//
	// FFA14F61-4E2F-54C7-9276-81C60BAC****
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
