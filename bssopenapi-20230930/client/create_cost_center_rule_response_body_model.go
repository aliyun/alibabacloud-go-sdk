// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *CreateCostCenterRuleResponseBody
	GetData() *int64
	SetMetadata(v interface{}) *CreateCostCenterRuleResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CreateCostCenterRuleResponseBody
	GetRequestId() *string
}

type CreateCostCenterRuleResponseBody struct {
	// example:
	//
	// 1111
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 79EE7556-0CFD-44EB-9CD6-B3B526E3A85F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCostCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostCenterRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateCostCenterRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateCostCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCostCenterRuleResponseBody) SetData(v int64) *CreateCostCenterRuleResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCostCenterRuleResponseBody) SetMetadata(v interface{}) *CreateCostCenterRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateCostCenterRuleResponseBody) SetRequestId(v string) *CreateCostCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostCenterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
