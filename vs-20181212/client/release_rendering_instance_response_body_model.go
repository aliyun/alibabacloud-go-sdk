// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseRenderingInstanceResponseBody
	GetRequestId() *string
}

type ReleaseRenderingInstanceResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseRenderingInstanceResponseBody) SetRequestId(v string) *ReleaseRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
