// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingInstanceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRenderingInstanceConfigurationResponseBody
	GetRequestId() *string
}

type UpdateRenderingInstanceConfigurationResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRenderingInstanceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingInstanceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRenderingInstanceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRenderingInstanceConfigurationResponseBody) SetRequestId(v string) *UpdateRenderingInstanceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRenderingInstanceConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
