// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCostCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *DeleteCostCenterRuleResponseBody
	GetData() *int64
	SetMetadata(v interface{}) *DeleteCostCenterRuleResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *DeleteCostCenterRuleResponseBody
	GetRequestId() *string
}

type DeleteCostCenterRuleResponseBody struct {
	// example:
	//
	// 37404
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// UUID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCostCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCostCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCostCenterRuleResponseBody) GetData() *int64 {
	return s.Data
}

func (s *DeleteCostCenterRuleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *DeleteCostCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCostCenterRuleResponseBody) SetData(v int64) *DeleteCostCenterRuleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteCostCenterRuleResponseBody) SetMetadata(v interface{}) *DeleteCostCenterRuleResponseBody {
	s.Metadata = v
	return s
}

func (s *DeleteCostCenterRuleResponseBody) SetRequestId(v string) *DeleteCostCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCostCenterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
