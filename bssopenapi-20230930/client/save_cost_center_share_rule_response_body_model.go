// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveCostCenterShareRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetadata(v interface{}) *SaveCostCenterShareRuleResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *SaveCostCenterShareRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveCostCenterShareRuleResponseBody
	GetSuccess() *bool
}

type SaveCostCenterShareRuleResponseBody struct {
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveCostCenterShareRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveCostCenterShareRuleResponseBody) GoString() string {
	return s.String()
}

func (s *SaveCostCenterShareRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *SaveCostCenterShareRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveCostCenterShareRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveCostCenterShareRuleResponseBody) SetMetadata(v interface{}) *SaveCostCenterShareRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *SaveCostCenterShareRuleResponseBody) SetRequestId(v string) *SaveCostCenterShareRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveCostCenterShareRuleResponseBody) SetSuccess(v bool) *SaveCostCenterShareRuleResponseBody {
	s.Success = &v
	return s
}

func (s *SaveCostCenterShareRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
