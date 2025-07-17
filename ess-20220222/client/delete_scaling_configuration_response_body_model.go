// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScalingConfigurationResponseBody
	GetRequestId() *string
}

type DeleteScalingConfigurationResponseBody struct {
	// The ID of the request. The request ID is returned regardless of whether the call is successful.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScalingConfigurationResponseBody) SetRequestId(v string) *DeleteScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
