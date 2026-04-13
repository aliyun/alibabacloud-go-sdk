// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDetectConfigResponseBody
	GetRequestId() *string
}

type UpdateDetectConfigResponseBody struct {
	// example:
	//
	// valueA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateDetectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDetectConfigResponseBody) SetRequestId(v string) *UpdateDetectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDetectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
