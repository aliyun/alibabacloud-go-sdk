// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingInstanceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRenderingInstanceSettingsResponseBody
	GetRequestId() *string
}

type DeleteRenderingInstanceSettingsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingInstanceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingInstanceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingInstanceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRenderingInstanceSettingsResponseBody) SetRequestId(v string) *DeleteRenderingInstanceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRenderingInstanceSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}
