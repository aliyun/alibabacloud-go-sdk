// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRenderingInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResetRenderingInstanceResponseBody
	GetRequestId() *string
}

type ResetRenderingInstanceResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetRenderingInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetRenderingInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ResetRenderingInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetRenderingInstanceResponseBody) SetRequestId(v string) *ResetRenderingInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetRenderingInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
