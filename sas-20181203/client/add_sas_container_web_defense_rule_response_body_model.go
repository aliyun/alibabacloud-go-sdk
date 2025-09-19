// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSasContainerWebDefenseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *AddSasContainerWebDefenseRuleResponseBody
	GetData() *int64
	SetRequestId(v string) *AddSasContainerWebDefenseRuleResponseBody
	GetRequestId() *string
}

type AddSasContainerWebDefenseRuleResponseBody struct {
	// The unique value of the created rule.
	//
	// example:
	//
	// 200634
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8B4B6E6D-B0B0-5F05-A14E-82917D9648EE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddSasContainerWebDefenseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddSasContainerWebDefenseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *AddSasContainerWebDefenseRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddSasContainerWebDefenseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddSasContainerWebDefenseRuleResponseBody) SetData(v int64) *AddSasContainerWebDefenseRuleResponseBody {
	s.Data = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleResponseBody) SetRequestId(v string) *AddSasContainerWebDefenseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddSasContainerWebDefenseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
