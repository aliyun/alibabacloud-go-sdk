// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOriginProtectionIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateOriginProtectionIpWhiteListResponseBody
	GetRequestId() *string
}

type UpdateOriginProtectionIpWhiteListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateOriginProtectionIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOriginProtectionIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOriginProtectionIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOriginProtectionIpWhiteListResponseBody) SetRequestId(v string) *UpdateOriginProtectionIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOriginProtectionIpWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
