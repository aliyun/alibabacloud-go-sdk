// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBasicAcceleratorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateBasicAcceleratorResponseBody
	GetRequestId() *string
}

type UpdateBasicAcceleratorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBasicAcceleratorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBasicAcceleratorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBasicAcceleratorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBasicAcceleratorResponseBody) SetRequestId(v string) *UpdateBasicAcceleratorResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBasicAcceleratorResponseBody) Validate() error {
	return dara.Validate(s)
}
