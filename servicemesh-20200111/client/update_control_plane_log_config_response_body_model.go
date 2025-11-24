// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateControlPlaneLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateControlPlaneLogConfigResponseBody
	GetRequestId() *string
}

type UpdateControlPlaneLogConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 488F046B-63D2-5D96-9A70-E00C3685D49F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateControlPlaneLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateControlPlaneLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateControlPlaneLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateControlPlaneLogConfigResponseBody) SetRequestId(v string) *UpdateControlPlaneLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateControlPlaneLogConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
