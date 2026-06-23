// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomPrivacyPoliciesToBrandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *AddCustomPrivacyPoliciesToBrandRequest
	GetBrandId() *string
	SetCustomPrivacyPolicyIds(v []*string) *AddCustomPrivacyPoliciesToBrandRequest
	GetCustomPrivacyPolicyIds() []*string
	SetInstanceId(v string) *AddCustomPrivacyPoliciesToBrandRequest
	GetInstanceId() *string
}

type AddCustomPrivacyPoliciesToBrandRequest struct {
	// The brand ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// brand_xxxxx
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// The list of term IDs.
	//
	// This parameter is required.
	CustomPrivacyPolicyIds []*string `json:"CustomPrivacyPolicyIds,omitempty" xml:"CustomPrivacyPolicyIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s AddCustomPrivacyPoliciesToBrandRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomPrivacyPoliciesToBrandRequest) GoString() string {
	return s.String()
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) GetCustomPrivacyPolicyIds() []*string {
	return s.CustomPrivacyPolicyIds
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) SetBrandId(v string) *AddCustomPrivacyPoliciesToBrandRequest {
	s.BrandId = &v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) SetCustomPrivacyPolicyIds(v []*string) *AddCustomPrivacyPoliciesToBrandRequest {
	s.CustomPrivacyPolicyIds = v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) SetInstanceId(v string) *AddCustomPrivacyPoliciesToBrandRequest {
	s.InstanceId = &v
	return s
}

func (s *AddCustomPrivacyPoliciesToBrandRequest) Validate() error {
	return dara.Validate(s)
}
