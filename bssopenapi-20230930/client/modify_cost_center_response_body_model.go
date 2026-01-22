// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCostCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterOperateDto(v []*ModifyCostCenterResponseBodyCostCenterOperateDto) *ModifyCostCenterResponseBody
	GetCostCenterOperateDto() []*ModifyCostCenterResponseBodyCostCenterOperateDto
	SetMetadata(v interface{}) *ModifyCostCenterResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *ModifyCostCenterResponseBody
	GetRequestId() *string
}

type ModifyCostCenterResponseBody struct {
	CostCenterOperateDto []*ModifyCostCenterResponseBodyCostCenterOperateDto `json:"CostCenterOperateDto,omitempty" xml:"CostCenterOperateDto,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCostCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponseBody) GetCostCenterOperateDto() []*ModifyCostCenterResponseBodyCostCenterOperateDto {
	return s.CostCenterOperateDto
}

func (s *ModifyCostCenterResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ModifyCostCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCostCenterResponseBody) SetCostCenterOperateDto(v []*ModifyCostCenterResponseBodyCostCenterOperateDto) *ModifyCostCenterResponseBody {
	s.CostCenterOperateDto = v
	return s
}

func (s *ModifyCostCenterResponseBody) SetMetadata(v interface{}) *ModifyCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *ModifyCostCenterResponseBody) SetRequestId(v string) *ModifyCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCostCenterResponseBody) Validate() error {
	if s.CostCenterOperateDto != nil {
		for _, item := range s.CostCenterOperateDto {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCostCenterResponseBodyCostCenterOperateDto struct {
	// example:
	//
	// 485938
	CostCenterId *int64 `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	// example:
	//
	// True
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
}

func (s ModifyCostCenterResponseBodyCostCenterOperateDto) String() string {
	return dara.Prettify(s)
}

func (s ModifyCostCenterResponseBodyCostCenterOperateDto) GoString() string {
	return s.String()
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetCostCenterId(v int64) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.CostCenterId = &v
	return s
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetIsSuccess(v bool) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.IsSuccess = &v
	return s
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) SetOwnerAccountId(v int64) *ModifyCostCenterResponseBodyCostCenterOperateDto {
	s.OwnerAccountId = &v
	return s
}

func (s *ModifyCostCenterResponseBodyCostCenterOperateDto) Validate() error {
	return dara.Validate(s)
}
