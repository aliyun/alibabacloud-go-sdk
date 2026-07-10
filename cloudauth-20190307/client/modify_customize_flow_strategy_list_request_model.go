// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomizeFlowStrategyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetProductType(v string) *ModifyCustomizeFlowStrategyListRequest
	GetProductType() *string
	SetStrategyObject(v []*ModifyCustomizeFlowStrategyListRequestStrategyObject) *ModifyCustomizeFlowStrategyListRequest
	GetStrategyObject() []*ModifyCustomizeFlowStrategyListRequestStrategyObject
}

type ModifyCustomizeFlowStrategyListRequest struct {
	// The product type. Currently, only **ANT_CLOUD_AUTH*	- (financial-grade ID Verification) is supported. All other types have been discontinued.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The policy list.
	StrategyObject []*ModifyCustomizeFlowStrategyListRequestStrategyObject `json:"StrategyObject,omitempty" xml:"StrategyObject,omitempty" type:"Repeated"`
}

func (s ModifyCustomizeFlowStrategyListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizeFlowStrategyListRequest) GoString() string {
	return s.String()
}

func (s *ModifyCustomizeFlowStrategyListRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyCustomizeFlowStrategyListRequest) GetStrategyObject() []*ModifyCustomizeFlowStrategyListRequestStrategyObject {
	return s.StrategyObject
}

func (s *ModifyCustomizeFlowStrategyListRequest) SetProductType(v string) *ModifyCustomizeFlowStrategyListRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequest) SetStrategyObject(v []*ModifyCustomizeFlowStrategyListRequestStrategyObject) *ModifyCustomizeFlowStrategyListRequest {
	s.StrategyObject = v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequest) Validate() error {
	if s.StrategyObject != nil {
		for _, item := range s.StrategyObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCustomizeFlowStrategyListRequestStrategyObject struct {
	// AccumulateKey
	//
	// example:
	//
	// -
	AccumulateKey *string `json:"AccumulateKey,omitempty" xml:"AccumulateKey,omitempty"`
	// The size of the rate limiting statistical window, in minutes.
	//
	// example:
	//
	// 60
	AccumulateWindow *int64 `json:"AccumulateWindow,omitempty" xml:"AccumulateWindow,omitempty"`
	// The API name, which is the same as **ProductCode**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// The flow type. Valid values:
	//
	// - **ACCUMULATE**: repeated occurrence of an ID card.
	//
	// - **PASSED_RATE**: pass rate is less than.
	//
	// - **SUB_CODE_205**: authentication failed and the proportion of liveness attack 205 is greater than.
	//
	// - **SUB_CODE_206**: authentication failed and the proportion of liveness attack 206 is greater than.
	//
	// example:
	//
	// ACCUMULATE
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The operation.
	//
	// example:
	//
	// -
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The status. Valid values:
	//
	// - **disabled**: disabled.
	//
	// - **normal**: enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The rate limiting threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 126005125163xxxx
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ModifyCustomizeFlowStrategyListRequestStrategyObject) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomizeFlowStrategyListRequestStrategyObject) GoString() string {
	return s.String()
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetAccumulateKey() *string {
	return s.AccumulateKey
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetAccumulateWindow() *int64 {
	return s.AccumulateWindow
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetApiName() *string {
	return s.ApiName
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetFlowType() *string {
	return s.FlowType
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetId() *int64 {
	return s.Id
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetOperation() *string {
	return s.Operation
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetStatus() *string {
	return s.Status
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetThreshold() *int32 {
	return s.Threshold
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) GetUserId() *int64 {
	return s.UserId
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetAccumulateKey(v string) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.AccumulateKey = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetAccumulateWindow(v int64) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.AccumulateWindow = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetApiName(v string) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.ApiName = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetFlowType(v string) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.FlowType = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetId(v int64) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.Id = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetOperation(v string) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.Operation = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetStatus(v string) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.Status = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetThreshold(v int32) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.Threshold = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) SetUserId(v int64) *ModifyCustomizeFlowStrategyListRequestStrategyObject {
	s.UserId = &v
	return s
}

func (s *ModifyCustomizeFlowStrategyListRequestStrategyObject) Validate() error {
	return dara.Validate(s)
}
