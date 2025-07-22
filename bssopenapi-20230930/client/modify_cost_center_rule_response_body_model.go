// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *ModifyCostCenterRuleResponseBody
	GetData() *int64
	SetMetadata(v interface{}) *ModifyCostCenterRuleResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *ModifyCostCenterRuleResponseBody
	GetRequestId() *string
}

type ModifyCostCenterRuleResponseBody struct {
	// example:
	//
	// 5632
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCostCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *ModifyCostCenterRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ModifyCostCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCostCenterRuleResponseBody) SetData(v int64) *ModifyCostCenterRuleResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyCostCenterRuleResponseBody) SetMetadata(v interface{}) *ModifyCostCenterRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *ModifyCostCenterRuleResponseBody) SetRequestId(v string) *ModifyCostCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostCenterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
