// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEciScalingConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEciScalingConfigurationResponseBody
	GetRequestId() *string
}

type DeleteEciScalingConfigurationResponseBody struct {
	// The request ID.
	//
	// The request ID is consistently returned in the response, irrespective of whether the request is executed successfully or encounters an error.
	//
	// example:
	//
	// 45D5B0AD-3B00-4A9B-9911-6D5303B06712
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEciScalingConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEciScalingConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEciScalingConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEciScalingConfigurationResponseBody) SetRequestId(v string) *DeleteEciScalingConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEciScalingConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
