// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigHealthCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigHealthCheckResponseBody
	GetRequestId() *string
}

type ConfigHealthCheckResponseBody struct {
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigHealthCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigHealthCheckResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigHealthCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigHealthCheckResponseBody) SetRequestId(v string) *ConfigHealthCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigHealthCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
