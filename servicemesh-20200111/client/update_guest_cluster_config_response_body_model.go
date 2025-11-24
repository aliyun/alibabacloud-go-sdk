// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGuestClusterConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateGuestClusterConfigResponseBody
	GetRequestId() *string
}

type UpdateGuestClusterConfigResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 71680038-8009-5073-B43E-C057E9******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateGuestClusterConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateGuestClusterConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateGuestClusterConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateGuestClusterConfigResponseBody) SetRequestId(v string) *UpdateGuestClusterConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateGuestClusterConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
