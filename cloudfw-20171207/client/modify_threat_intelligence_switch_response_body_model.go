// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyThreatIntelligenceSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyThreatIntelligenceSwitchResponseBody
	GetRequestId() *string
}

type ModifyThreatIntelligenceSwitchResponseBody struct {
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-0009012****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyThreatIntelligenceSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyThreatIntelligenceSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyThreatIntelligenceSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyThreatIntelligenceSwitchResponseBody) SetRequestId(v string) *ModifyThreatIntelligenceSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyThreatIntelligenceSwitchResponseBody) Validate() error {
	return dara.Validate(s)
}
