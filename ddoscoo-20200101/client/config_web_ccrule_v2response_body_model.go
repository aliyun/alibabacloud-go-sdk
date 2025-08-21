// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebCCRuleV2ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigWebCCRuleV2ResponseBody
	GetRequestId() *string
}

type ConfigWebCCRuleV2ResponseBody struct {
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigWebCCRuleV2ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebCCRuleV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigWebCCRuleV2ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigWebCCRuleV2ResponseBody) SetRequestId(v string) *ConfigWebCCRuleV2ResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigWebCCRuleV2ResponseBody) Validate() error {
	return dara.Validate(s)
}
