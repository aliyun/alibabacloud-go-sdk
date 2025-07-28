// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceElasticVCUUpperLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateInstanceElasticVCUUpperLimitResponseBody
	GetRequestId() *string
}

type UpdateInstanceElasticVCUUpperLimitResponseBody struct {
	// The request ID, which can be used to troubleshoot issues.
	//
	// example:
	//
	// B37BBA04-D827-55C8-B901-5264B904E8C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateInstanceElasticVCUUpperLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceElasticVCUUpperLimitResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstanceElasticVCUUpperLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstanceElasticVCUUpperLimitResponseBody) SetRequestId(v string) *UpdateInstanceElasticVCUUpperLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstanceElasticVCUUpperLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
