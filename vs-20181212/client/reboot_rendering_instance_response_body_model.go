// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootRenderingInstanceResponseBody
	GetRequestId() *string
}

type RebootRenderingInstanceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootRenderingInstanceResponseBody) SetRequestId(v string) *RebootRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
