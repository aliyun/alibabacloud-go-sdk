// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyConfigurationAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEgressProxyConfigurationAttributeResponseBody
	GetRequestId() *string
}

type UpdateEgressProxyConfigurationAttributeResponseBody struct {
	// example:
	//
	// 019F3A67-2D1B-537B-B08D-B2DEB69E8E8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEgressProxyConfigurationAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyConfigurationAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyConfigurationAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEgressProxyConfigurationAttributeResponseBody) SetRequestId(v string) *UpdateEgressProxyConfigurationAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEgressProxyConfigurationAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
