// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRenderingInstanceAttributeResponseBody
	GetRequestId() *string
}

type ModifyRenderingInstanceAttributeResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRenderingInstanceAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRenderingInstanceAttributeResponseBody) SetRequestId(v string) *ModifyRenderingInstanceAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRenderingInstanceAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
