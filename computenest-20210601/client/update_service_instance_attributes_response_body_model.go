// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceInstanceAttributesResponseBody
	GetRequestId() *string
}

type UpdateServiceInstanceAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 06BF8F22-02DC-4750-83DF-3FFC11C065EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceInstanceAttributesResponseBody) SetRequestId(v string) *UpdateServiceInstanceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceInstanceAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
