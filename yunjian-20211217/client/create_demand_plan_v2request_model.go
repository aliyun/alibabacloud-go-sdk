// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemandPlanV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *CreateDemandPlanV2Request
	GetAccountId() *string
	SetDescription(v string) *CreateDemandPlanV2Request
	GetDescription() *string
	SetName(v string) *CreateDemandPlanV2Request
	GetName() *string
	SetProductType(v string) *CreateDemandPlanV2Request
	GetProductType() *string
	SetTargetCid(v int64) *CreateDemandPlanV2Request
	GetTargetCid() *int64
	SetType(v string) *CreateDemandPlanV2Request
	GetType() *string
	SetUserId(v string) *CreateDemandPlanV2Request
	GetUserId() *string
}

type CreateDemandPlanV2Request struct {
	// This parameter is required.
	AccountId   *string `json:"accountId,omitempty" xml:"accountId,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// This parameter is required.
	Name        *string `json:"name,omitempty" xml:"name,omitempty"`
	ProductType *string `json:"productType,omitempty" xml:"productType,omitempty"`
	TargetCid   *int64  `json:"targetCid,omitempty" xml:"targetCid,omitempty"`
	// This parameter is required.
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// This parameter is required.
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreateDemandPlanV2Request) String() string {
	return dara.Prettify(s)
}

func (s CreateDemandPlanV2Request) GoString() string {
	return s.String()
}

func (s *CreateDemandPlanV2Request) GetAccountId() *string {
	return s.AccountId
}

func (s *CreateDemandPlanV2Request) GetDescription() *string {
	return s.Description
}

func (s *CreateDemandPlanV2Request) GetName() *string {
	return s.Name
}

func (s *CreateDemandPlanV2Request) GetProductType() *string {
	return s.ProductType
}

func (s *CreateDemandPlanV2Request) GetTargetCid() *int64 {
	return s.TargetCid
}

func (s *CreateDemandPlanV2Request) GetType() *string {
	return s.Type
}

func (s *CreateDemandPlanV2Request) GetUserId() *string {
	return s.UserId
}

func (s *CreateDemandPlanV2Request) SetAccountId(v string) *CreateDemandPlanV2Request {
	s.AccountId = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetDescription(v string) *CreateDemandPlanV2Request {
	s.Description = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetName(v string) *CreateDemandPlanV2Request {
	s.Name = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetProductType(v string) *CreateDemandPlanV2Request {
	s.ProductType = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetTargetCid(v int64) *CreateDemandPlanV2Request {
	s.TargetCid = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetType(v string) *CreateDemandPlanV2Request {
	s.Type = &v
	return s
}

func (s *CreateDemandPlanV2Request) SetUserId(v string) *CreateDemandPlanV2Request {
	s.UserId = &v
	return s
}

func (s *CreateDemandPlanV2Request) Validate() error {
	return dara.Validate(s)
}
