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
	// Product type, currently only supports **ANT_CLOUD_AUTH*	- (Financial-grade real person), all others have been phased out.
	//
	// example:
	//
	// ANT_CLOUD_AUTH
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Strategy list.
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
	// Flow control statistical window size, in minutes.
	//
	// example:
	//
	// 60
	AccumulateWindow *int64 `json:"AccumulateWindow,omitempty" xml:"AccumulateWindow,omitempty"`
	// API name, same as **ProductCode**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ID_PRO
	ApiName *string `json:"ApiName,omitempty" xml:"ApiName,omitempty"`
	// Flow type:
	//
	// - **ACCUMULATE**: Repeated appearance of ID card
	//
	// - **PASSED_RATE**: Pass rate less than
	//
	// - **SUB_CODE_205**: Authentication failed and liveness attack 205 ratio greater than
	//
	// - **SUB_CODE_206**: Authentication failed and liveness attack 206 ratio greater than
	//
	// example:
	//
	// ACCUMULATE
	FlowType *string `json:"FlowType,omitempty" xml:"FlowType,omitempty"`
	// Rule ID.
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Operation.
	//
	// example:
	//
	// -
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// Status:
	//
	// - **disabled**: Disabled
	//
	// - **normal**: Enabled
	//
	// This parameter is required.
	//
	// example:
	//
	// disabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Flow control threshold.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Threshold *int32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// User ID.
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
