// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRenderingInstanceBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyRenderingInstanceBandwidthResponseBody
	GetRequestId() *string
}

type ModifyRenderingInstanceBandwidthResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRenderingInstanceBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRenderingInstanceBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRenderingInstanceBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRenderingInstanceBandwidthResponseBody) SetRequestId(v string) *ModifyRenderingInstanceBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRenderingInstanceBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
