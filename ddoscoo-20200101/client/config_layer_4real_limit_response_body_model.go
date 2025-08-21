// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigLayer4RealLimitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigLayer4RealLimitResponseBody
	GetRequestId() *string
}

type ConfigLayer4RealLimitResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CFCF71BD-680E-5A20-8847-174CEC62E67D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigLayer4RealLimitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigLayer4RealLimitResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigLayer4RealLimitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigLayer4RealLimitResponseBody) SetRequestId(v string) *ConfigLayer4RealLimitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigLayer4RealLimitResponseBody) Validate() error {
	return dara.Validate(s)
}
