// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebPreciseAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebPreciseAccessRuleResponseBody
	GetRequestId() *string
}

type DeleteWebPreciseAccessRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebPreciseAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebPreciseAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebPreciseAccessRuleResponseBody) SetRequestId(v string) *DeleteWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebPreciseAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
