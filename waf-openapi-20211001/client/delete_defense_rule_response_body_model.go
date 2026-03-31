// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDefenseRuleResponseBody
	GetRequestId() *string
}

type DeleteDefenseRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 2EC856FE-6D31-5861-8275-E5DEDB539089
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDefenseRuleResponseBody) SetRequestId(v string) *DeleteDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
