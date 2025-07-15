// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorSessionAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficMirrorSessionAttributeResponseBody
	GetRequestId() *string
}

type UpdateTrafficMirrorSessionAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 64DCAF03-E2C7-479A-ACEA-38B79876B006
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficMirrorSessionAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorSessionAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorSessionAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficMirrorSessionAttributeResponseBody) SetRequestId(v string) *UpdateTrafficMirrorSessionAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficMirrorSessionAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
