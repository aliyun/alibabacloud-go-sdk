// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLabelConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetServiceLabelConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *GetServiceLabelConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *GetServiceLabelConfigRequest
	GetServiceCode() *string
}

type GetServiceLabelConfigRequest struct {
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// text
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// nickname_detection
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s GetServiceLabelConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLabelConfigRequest) GoString() string {
	return s.String()
}

func (s *GetServiceLabelConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetServiceLabelConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetServiceLabelConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *GetServiceLabelConfigRequest) SetRegionId(v string) *GetServiceLabelConfigRequest {
	s.RegionId = &v
	return s
}

func (s *GetServiceLabelConfigRequest) SetResourceType(v string) *GetServiceLabelConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *GetServiceLabelConfigRequest) SetServiceCode(v string) *GetServiceLabelConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *GetServiceLabelConfigRequest) Validate() error {
	return dara.Validate(s)
}
