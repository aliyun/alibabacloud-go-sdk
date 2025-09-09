// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateServiceInstanceAttributeResponseBody
	GetRequestId() *string
}

type UpdateServiceInstanceAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0CB2E0A9-B4DF-5C16-86AD-C511C483144B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceInstanceAttributeResponseBody) SetRequestId(v string) *UpdateServiceInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
