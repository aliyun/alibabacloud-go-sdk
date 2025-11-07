// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackListStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackListStrategy(v *ModifyBlackListStrategyRequestBlackListStrategy) *ModifyBlackListStrategyRequest
	GetBlackListStrategy() *ModifyBlackListStrategyRequestBlackListStrategy
	SetRegionId(v string) *ModifyBlackListStrategyRequest
	GetRegionId() *string
}

type ModifyBlackListStrategyRequest struct {
	// Blacklist rule.
	BlackListStrategy *ModifyBlackListStrategyRequestBlackListStrategy `json:"BlackListStrategy,omitempty" xml:"BlackListStrategy,omitempty" type:"Struct"`
	// Region ID
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyBlackListStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackListStrategyRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlackListStrategyRequest) GetBlackListStrategy() *ModifyBlackListStrategyRequestBlackListStrategy {
	return s.BlackListStrategy
}

func (s *ModifyBlackListStrategyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyBlackListStrategyRequest) SetBlackListStrategy(v *ModifyBlackListStrategyRequestBlackListStrategy) *ModifyBlackListStrategyRequest {
	s.BlackListStrategy = v
	return s
}

func (s *ModifyBlackListStrategyRequest) SetRegionId(v string) *ModifyBlackListStrategyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyBlackListStrategyRequest) Validate() error {
	if s.BlackListStrategy != nil {
		if err := s.BlackListStrategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBlackListStrategyRequestBlackListStrategy struct {
	// Blacklist string, with each blacklist entry separated by commas.
	//
	// example:
	//
	// 127.0.0.1,127.0.0.2
	BizContent *string `json:"BizContent,omitempty" xml:"BizContent,omitempty"`
	// List Type:
	//
	// - **mobile**: Mobile number blacklist
	//
	// - **ip**: IP blacklist
	//
	// - **identifyNum**: ID number blacklist
	//
	// - **bankCard**: Bank card blacklist
	//
	// example:
	//
	// ip
	BizKey *string `json:"BizKey,omitempty" xml:"BizKey,omitempty"`
	// Rule ID:
	//
	// - **Empty**: Add a new rule
	//
	// - **Non-empty**: Modify an existing rule
	//
	// example:
	//
	// 38
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Product Name:
	//
	// - **id2meta**: ID card two-factor verification
	//
	// - **mobile3Meta**: Mobile number factor verification
	//
	// - **bankcardMeta**: Bank card factor verification
	//
	// example:
	//
	// id2meta
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// Verification Status:
	//
	// - **1**: Verification passed
	//
	// - **2**: Verification failed
	//
	// example:
	//
	// 3
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyBlackListStrategyRequestBlackListStrategy) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackListStrategyRequestBlackListStrategy) GoString() string {
	return s.String()
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) GetBizContent() *string {
	return s.BizContent
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) GetBizKey() *string {
	return s.BizKey
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) GetId() *int64 {
	return s.Id
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) GetProductName() *string {
	return s.ProductName
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) GetStatus() *string {
	return s.Status
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) SetBizContent(v string) *ModifyBlackListStrategyRequestBlackListStrategy {
	s.BizContent = &v
	return s
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) SetBizKey(v string) *ModifyBlackListStrategyRequestBlackListStrategy {
	s.BizKey = &v
	return s
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) SetId(v int64) *ModifyBlackListStrategyRequestBlackListStrategy {
	s.Id = &v
	return s
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) SetProductName(v string) *ModifyBlackListStrategyRequestBlackListStrategy {
	s.ProductName = &v
	return s
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) SetStatus(v string) *ModifyBlackListStrategyRequestBlackListStrategy {
	s.Status = &v
	return s
}

func (s *ModifyBlackListStrategyRequestBlackListStrategy) Validate() error {
	return dara.Validate(s)
}
