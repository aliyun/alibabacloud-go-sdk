// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebRuleResponseBody
	GetRequestId() *string
}

type DeleteWebRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9EC62E89-BD30-4FCD-9CB8-FA53865FF0D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebRuleResponseBody) SetRequestId(v string) *DeleteWebRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
