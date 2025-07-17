// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEciScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEciScalingConfigurationResponseBody
	GetRequestId() *string
	SetScalingConfigurationId(v string) *CreateEciScalingConfigurationResponseBody
	GetScalingConfigurationId() *string
}

type CreateEciScalingConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 89945DD3-9072-47D0-A318-353284CF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the elastic container instance (container group).
	//
	// example:
	//
	// eci-uf6fonnghi50u374****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s CreateEciScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEciScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEciScalingConfigurationResponseBody) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *CreateEciScalingConfigurationResponseBody) SetRequestId(v string) *CreateEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEciScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateEciScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

func (s *CreateEciScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
