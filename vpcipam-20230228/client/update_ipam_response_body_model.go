// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpamResponseBody
	GetRequestId() *string
}

type UpdateIpamResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F4650E33-895C-53F0-9CD5-D1338F322DE8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpamResponseBody) SetRequestId(v string) *UpdateIpamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpamResponseBody) Validate() error {
	return dara.Validate(s)
}
