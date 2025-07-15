// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficMirrorFilterAttributeResponseBody
	GetRequestId() *string
}

type UpdateTrafficMirrorFilterAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5816D35F-94D5-48CE-838F-2327C8EE8D49
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficMirrorFilterAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficMirrorFilterAttributeResponseBody) SetRequestId(v string) *UpdateTrafficMirrorFilterAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
