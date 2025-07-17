// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyScalingConfigurationResponseBody
	GetRequestId() *string
}

type ModifyScalingConfigurationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScalingConfigurationResponseBody) SetRequestId(v string) *ModifyScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
