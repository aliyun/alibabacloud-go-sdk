// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEciScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEciScalingConfigurationResponseBody
	GetRequestId() *string
}

type ModifyEciScalingConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEciScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEciScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEciScalingConfigurationResponseBody) SetRequestId(v string) *ModifyEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEciScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
