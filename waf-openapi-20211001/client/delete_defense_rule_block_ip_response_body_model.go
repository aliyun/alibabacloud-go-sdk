// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDefenseRuleBlockIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDefenseRuleBlockIpResponseBody
	GetRequestId() *string
}

type DeleteDefenseRuleBlockIpResponseBody struct {
	// example:
	//
	// 9D11AC3A-A10C-56E7-A342-E87EC892BAE2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDefenseRuleBlockIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDefenseRuleBlockIpResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDefenseRuleBlockIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDefenseRuleBlockIpResponseBody) SetRequestId(v string) *DeleteDefenseRuleBlockIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDefenseRuleBlockIpResponseBody) Validate() error {
	return dara.Validate(s)
}
