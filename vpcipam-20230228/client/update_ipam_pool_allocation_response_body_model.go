// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamPoolAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpamPoolAllocationResponseBody
	GetRequestId() *string
}

type UpdateIpamPoolAllocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F4650E33-895C-53F0-9CD5-D1338F322DE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamPoolAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamPoolAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpamPoolAllocationResponseBody) SetRequestId(v string) *UpdateIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpamPoolAllocationResponseBody) Validate() error {
	return dara.Validate(s)
}
