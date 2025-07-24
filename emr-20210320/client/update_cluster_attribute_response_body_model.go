// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClusterAttributeResponseBody
	GetRequestId() *string
}

type UpdateClusterAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateClusterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterAttributeResponseBody) SetRequestId(v string) *UpdateClusterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
