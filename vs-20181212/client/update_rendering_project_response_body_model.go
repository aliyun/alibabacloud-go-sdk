// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRenderingProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRenderingProjectResponseBody
	GetRequestId() *string
}

type UpdateRenderingProjectResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRenderingProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRenderingProjectResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRenderingProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRenderingProjectResponseBody) SetRequestId(v string) *UpdateRenderingProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRenderingProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
