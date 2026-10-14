// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBizBandWidthModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBizBandWidthModeResponseBody
	GetRequestId() *string
}

type ModifyBizBandWidthModeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBizBandWidthModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBizBandWidthModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBizBandWidthModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBizBandWidthModeResponseBody) SetRequestId(v string) *ModifyBizBandWidthModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBizBandWidthModeResponseBody) Validate() error {
	return dara.Validate(s)
}
