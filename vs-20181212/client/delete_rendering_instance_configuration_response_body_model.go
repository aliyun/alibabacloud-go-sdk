// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRenderingInstanceConfigurationResponseBody
	GetRequestId() *string
}

type DeleteRenderingInstanceConfigurationResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingInstanceConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRenderingInstanceConfigurationResponseBody) SetRequestId(v string) *DeleteRenderingInstanceConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRenderingInstanceConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
