// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterRuleAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateTrafficMirrorFilterRuleAttributeResponseBody
	GetRequestId() *string
}

type UpdateTrafficMirrorFilterRuleAttributeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 02EB8585-D4DC-4E29-A0F4-7C588C82863C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTrafficMirrorFilterRuleAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterRuleAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponseBody) SetRequestId(v string) *UpdateTrafficMirrorFilterRuleAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterRuleAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
