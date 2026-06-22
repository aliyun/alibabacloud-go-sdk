// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNodeGroupAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeGroupAttributesResponseBody
	GetRequestId() *string
}

type UpdateNodeGroupAttributesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNodeGroupAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeGroupAttributesResponseBody) SetRequestId(v string) *UpdateNodeGroupAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
