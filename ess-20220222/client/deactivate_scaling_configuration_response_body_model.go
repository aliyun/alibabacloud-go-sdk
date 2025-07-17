// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeactivateScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeactivateScalingConfigurationResponseBody
	GetRequestId() *string
}

type DeactivateScalingConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D09C9E64-7D19-4E1C-8CD6-123D34A4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeactivateScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeactivateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeactivateScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeactivateScalingConfigurationResponseBody) SetRequestId(v string) *DeactivateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeactivateScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
