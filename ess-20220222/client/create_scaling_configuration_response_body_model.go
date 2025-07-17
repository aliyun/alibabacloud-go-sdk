// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateScalingConfigurationResponseBody
	GetRequestId() *string
	SetScalingConfigurationId(v string) *CreateScalingConfigurationResponseBody
	GetScalingConfigurationId() *string
}

type CreateScalingConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-bp1ffogfdauy0nu5****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s CreateScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScalingConfigurationResponseBody) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *CreateScalingConfigurationResponseBody) SetRequestId(v string) *CreateScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *CreateScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

func (s *CreateScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
