// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIntegrationPolicyResponseBody
	GetRequestId() *string
}

type DeleteIntegrationPolicyResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CD8BA7D6-995D-578D-9941-78B0FECD14B5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteIntegrationPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntegrationPolicyResponseBody) SetRequestId(v string) *DeleteIntegrationPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntegrationPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
