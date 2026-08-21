// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelDeploymentResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMemberMatches(v []*GetModelDeploymentResourcesResponseBodyMemberMatches) *GetModelDeploymentResourcesResponseBody
	GetMemberMatches() []*GetModelDeploymentResourcesResponseBodyMemberMatches
	SetRequestId(v string) *GetModelDeploymentResourcesResponseBody
	GetRequestId() *string
}

type GetModelDeploymentResourcesResponseBody struct {
	MemberMatches []*GetModelDeploymentResourcesResponseBodyMemberMatches `json:"MemberMatches,omitempty" xml:"MemberMatches,omitempty" type:"Repeated"`
	// example:
	//
	// B6B54325-C98C-5937-87A3-2F96C07652EC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetModelDeploymentResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponseBody) GetMemberMatches() []*GetModelDeploymentResourcesResponseBodyMemberMatches {
	return s.MemberMatches
}

func (s *GetModelDeploymentResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetModelDeploymentResourcesResponseBody) SetMemberMatches(v []*GetModelDeploymentResourcesResponseBodyMemberMatches) *GetModelDeploymentResourcesResponseBody {
	s.MemberMatches = v
	return s
}

func (s *GetModelDeploymentResourcesResponseBody) SetRequestId(v string) *GetModelDeploymentResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetModelDeploymentResourcesResponseBody) Validate() error {
	if s.MemberMatches != nil {
		for _, item := range s.MemberMatches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelDeploymentResourcesResponseBodyMemberMatches struct {
	DedicatedResources []*GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources `json:"DedicatedResources,omitempty" xml:"DedicatedResources,omitempty" type:"Repeated"`
	// example:
	//
	// Default
	MemberType      *string                                                                `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
	PublicResources []*GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources `json:"PublicResources,omitempty" xml:"PublicResources,omitempty" type:"Repeated"`
	Quotas          []*GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas          `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatches) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatches) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) GetDedicatedResources() []*GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources {
	return s.DedicatedResources
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) GetMemberType() *string {
	return s.MemberType
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) GetPublicResources() []*GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources {
	return s.PublicResources
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) GetQuotas() []*GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas {
	return s.Quotas
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) SetDedicatedResources(v []*GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) *GetModelDeploymentResourcesResponseBodyMemberMatches {
	s.DedicatedResources = v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) SetMemberType(v string) *GetModelDeploymentResourcesResponseBodyMemberMatches {
	s.MemberType = &v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) SetPublicResources(v []*GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) *GetModelDeploymentResourcesResponseBodyMemberMatches {
	s.PublicResources = v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) SetQuotas(v []*GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) *GetModelDeploymentResourcesResponseBodyMemberMatches {
	s.Quotas = v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatches) Validate() error {
	if s.DedicatedResources != nil {
		for _, item := range s.DedicatedResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PublicResources != nil {
		for _, item := range s.PublicResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Quotas != nil {
		for _, item := range s.Quotas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources struct {
	// example:
	//
	// eas-r-lq9p****ao9m2
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) SetResourceId(v string) *GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources {
	s.ResourceId = &v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesDedicatedResources) Validate() error {
	return dara.Validate(s)
}

type GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources struct {
	// example:
	//
	// ml.gx9cf.8.62xlarg
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) SetInstanceType(v string) *GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources {
	s.InstanceType = &v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesPublicResources) Validate() error {
	return dara.Validate(s)
}

type GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas struct {
	// example:
	//
	// quotagn***bb68
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) String() string {
	return dara.Prettify(s)
}

func (s GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) GoString() string {
	return s.String()
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) GetQuotaId() *string {
	return s.QuotaId
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) SetQuotaId(v string) *GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas {
	s.QuotaId = &v
	return s
}

func (s *GetModelDeploymentResourcesResponseBodyMemberMatchesQuotas) Validate() error {
	return dara.Validate(s)
}
