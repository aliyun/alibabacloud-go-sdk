// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRenderingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRenderingProjectResponseBody
	GetRequestId() *string
}

type DeleteRenderingProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRenderingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRenderingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRenderingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRenderingProjectResponseBody) SetRequestId(v string) *DeleteRenderingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRenderingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
