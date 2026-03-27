// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntegrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIntegrationPolicyResponseBody
	GetRequestId() *string
}

type UpdateIntegrationPolicyResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 0CEC5375-C554-562B-A65F-9A629907C1F0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateIntegrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntegrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntegrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIntegrationPolicyResponseBody) SetRequestId(v string) *UpdateIntegrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntegrationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
