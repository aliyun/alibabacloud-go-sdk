// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSolutionInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSolutionInstanceAttributeResponseBody
	GetRequestId() *string
}

type UpdateSolutionInstanceAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSolutionInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSolutionInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSolutionInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSolutionInstanceAttributeResponseBody) SetRequestId(v string) *UpdateSolutionInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSolutionInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
