// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCostCenterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCostCenterDtoList(v []*CreateCostCenterResponseBodyCostCenterDtoList) *CreateCostCenterResponseBody
	GetCostCenterDtoList() []*CreateCostCenterResponseBodyCostCenterDtoList
	SetMetadata(v interface{}) *CreateCostCenterResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *CreateCostCenterResponseBody
	GetRequestId() *string
}

type CreateCostCenterResponseBody struct {
	CostCenterDtoList []*CreateCostCenterResponseBodyCostCenterDtoList `json:"CostCenterDtoList,omitempty" xml:"CostCenterDtoList,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// C1BD134E-D914-6AE0-1901-AEB2A99FA205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCostCenterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponseBody) GetCostCenterDtoList() []*CreateCostCenterResponseBodyCostCenterDtoList {
	return s.CostCenterDtoList
}

func (s *CreateCostCenterResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *CreateCostCenterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCostCenterResponseBody) SetCostCenterDtoList(v []*CreateCostCenterResponseBodyCostCenterDtoList) *CreateCostCenterResponseBody {
	s.CostCenterDtoList = v
	return s
}

func (s *CreateCostCenterResponseBody) SetMetadata(v interface{}) *CreateCostCenterResponseBody {
	s.Metadata = v
	return s
}

func (s *CreateCostCenterResponseBody) SetRequestId(v string) *CreateCostCenterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCostCenterResponseBody) Validate() error {
	if s.CostCenterDtoList != nil {
		for _, item := range s.CostCenterDtoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCostCenterResponseBodyCostCenterDtoList struct {
	// example:
	//
	// 485938
	CostCenterId   *int64  `json:"CostCenterId,omitempty" xml:"CostCenterId,omitempty"`
	CostCenterName *string `json:"CostCenterName,omitempty" xml:"CostCenterName,omitempty"`
	// example:
	//
	// 1314839403940987
	OwnerAccountId *int64 `json:"OwnerAccountId,omitempty" xml:"OwnerAccountId,omitempty"`
	// example:
	//
	// -1
	ParentCostCenterId *int64 `json:"ParentCostCenterId,omitempty" xml:"ParentCostCenterId,omitempty"`
}

func (s CreateCostCenterResponseBodyCostCenterDtoList) String() string {
	return dara.Prettify(s)
}

func (s CreateCostCenterResponseBodyCostCenterDtoList) GoString() string {
	return s.String()
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) GetCostCenterId() *int64 {
	return s.CostCenterId
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) GetOwnerAccountId() *int64 {
	return s.OwnerAccountId
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) GetParentCostCenterId() *int64 {
	return s.ParentCostCenterId
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetCostCenterId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterId = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetCostCenterName(v string) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.CostCenterName = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetOwnerAccountId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.OwnerAccountId = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) SetParentCostCenterId(v int64) *CreateCostCenterResponseBodyCostCenterDtoList {
	s.ParentCostCenterId = &v
	return s
}

func (s *CreateCostCenterResponseBodyCostCenterDtoList) Validate() error {
	return dara.Validate(s)
}
