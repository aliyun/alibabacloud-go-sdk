// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeResponseBody
	GetRenderingInstanceId() *string
	SetRequestId(v string) *ModifyRenderingChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyRenderingChargeTypeResponseBody struct {
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRenderingChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRenderingChargeTypeResponseBody) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *ModifyRenderingChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRenderingChargeTypeResponseBody) SetRenderingInstanceId(v string) *ModifyRenderingChargeTypeResponseBody {
	s.RenderingInstanceId = &v
	return s
}

func (s *ModifyRenderingChargeTypeResponseBody) SetRequestId(v string) *ModifyRenderingChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRenderingChargeTypeResponseBody) Validate() error {
	return dara.Validate(s)
}
