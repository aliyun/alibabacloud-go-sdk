// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyEciScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ApplyEciScalingConfigurationResponseBody
	GetRequestId() *string
	SetScalingConfigurationId(v string) *ApplyEciScalingConfigurationResponseBody
	GetScalingConfigurationId() *string
}

type ApplyEciScalingConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the scaling configuration.
	//
	// example:
	//
	// asc-bp16har3jpj6fjbx****
	ScalingConfigurationId *string `json:"ScalingConfigurationId,omitempty" xml:"ScalingConfigurationId,omitempty"`
}

func (s ApplyEciScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ApplyEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ApplyEciScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApplyEciScalingConfigurationResponseBody) GetScalingConfigurationId() *string {
	return s.ScalingConfigurationId
}

func (s *ApplyEciScalingConfigurationResponseBody) SetRequestId(v string) *ApplyEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApplyEciScalingConfigurationResponseBody) SetScalingConfigurationId(v string) *ApplyEciScalingConfigurationResponseBody {
	s.ScalingConfigurationId = &v
	return s
}

func (s *ApplyEciScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
