// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChargeType(v string) *CreateWorkspaceRequest
	GetChargeType() *string
	SetEngineType(v string) *CreateWorkspaceRequest
	GetEngineType() *string
	SetName(v string) *CreateWorkspaceRequest
	GetName() *string
	SetQuota(v *CreateWorkspaceRequestQuota) *CreateWorkspaceRequest
	GetQuota() *CreateWorkspaceRequestQuota
	SetType(v string) *CreateWorkspaceRequest
	GetType() *string
}

type CreateWorkspaceRequest struct {
	// Billing type
	//
	// - POSTPAY: Pay-as-you-go
	//
	// example:
	//
	// POSTPAY
	ChargeType *string `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	// Engine type
	//
	// - rag
	//
	// example:
	//
	// rag
	EngineType *string `json:"engineType,omitempty" xml:"engineType,omitempty"`
	// Workspace name
	//
	// example:
	//
	// default
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Quota
	Quota *CreateWorkspaceRequestQuota `json:"quota,omitempty" xml:"quota,omitempty" type:"Struct"`
	// Type
	//
	// - standard
	//
	// example:
	//
	// standard
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateWorkspaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequest) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateWorkspaceRequest) GetEngineType() *string {
	return s.EngineType
}

func (s *CreateWorkspaceRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkspaceRequest) GetQuota() *CreateWorkspaceRequestQuota {
	return s.Quota
}

func (s *CreateWorkspaceRequest) GetType() *string {
	return s.Type
}

func (s *CreateWorkspaceRequest) SetChargeType(v string) *CreateWorkspaceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetEngineType(v string) *CreateWorkspaceRequest {
	s.EngineType = &v
	return s
}

func (s *CreateWorkspaceRequest) SetName(v string) *CreateWorkspaceRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkspaceRequest) SetQuota(v *CreateWorkspaceRequestQuota) *CreateWorkspaceRequest {
	s.Quota = v
	return s
}

func (s *CreateWorkspaceRequest) SetType(v string) *CreateWorkspaceRequest {
	s.Type = &v
	return s
}

func (s *CreateWorkspaceRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceRequestQuota struct {
	// Compute resource (unit: LCU)
	//
	// example:
	//
	// 0
	ComputeResource *int32 `json:"computeResource,omitempty" xml:"computeResource,omitempty"`
	// Storage capacity (unit: GB)
	//
	// example:
	//
	// 0
	DocSize *int32 `json:"docSize,omitempty" xml:"docSize,omitempty"`
	// Specification
	//
	// - rag.share.common
	//
	// example:
	//
	// rag.share.common
	Spec *string `json:"spec,omitempty" xml:"spec,omitempty"`
}

func (s CreateWorkspaceRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceRequestQuota) GetComputeResource() *int32 {
	return s.ComputeResource
}

func (s *CreateWorkspaceRequestQuota) GetDocSize() *int32 {
	return s.DocSize
}

func (s *CreateWorkspaceRequestQuota) GetSpec() *string {
	return s.Spec
}

func (s *CreateWorkspaceRequestQuota) SetComputeResource(v int32) *CreateWorkspaceRequestQuota {
	s.ComputeResource = &v
	return s
}

func (s *CreateWorkspaceRequestQuota) SetDocSize(v int32) *CreateWorkspaceRequestQuota {
	s.DocSize = &v
	return s
}

func (s *CreateWorkspaceRequestQuota) SetSpec(v string) *CreateWorkspaceRequestQuota {
	s.Spec = &v
	return s
}

func (s *CreateWorkspaceRequestQuota) Validate() error {
	return dara.Validate(s)
}
