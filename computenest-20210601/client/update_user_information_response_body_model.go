// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInformationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateUserInformationResponseBody
	GetRequestId() *string
}

type UpdateUserInformationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51945B04-6AA6-410D-93BA-236E0248B104
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateUserInformationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInformationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserInformationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserInformationResponseBody) SetRequestId(v string) *UpdateUserInformationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserInformationResponseBody) Validate() error {
	return dara.Validate(s)
}
