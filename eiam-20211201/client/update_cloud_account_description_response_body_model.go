// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateCloudAccountDescriptionResponseBody
	GetRequestId() *string
}

type UpdateCloudAccountDescriptionResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCloudAccountDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCloudAccountDescriptionResponseBody) SetRequestId(v string) *UpdateCloudAccountDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCloudAccountDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
