// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyServiceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CopyServiceConfigRequest
	GetRegionId() *string
	SetResourceType(v string) *CopyServiceConfigRequest
	GetResourceType() *string
	SetServiceCode(v string) *CopyServiceConfigRequest
	GetServiceCode() *string
	SetServiceDesc(v string) *CopyServiceConfigRequest
	GetServiceDesc() *string
	SetServiceName(v string) *CopyServiceConfigRequest
	GetServiceName() *string
}

type CopyServiceConfigRequest struct {
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
	ServiceDesc *string `json:"ServiceDesc,omitempty" xml:"ServiceDesc,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s CopyServiceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyServiceConfigRequest) GoString() string {
	return s.String()
}

func (s *CopyServiceConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CopyServiceConfigRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CopyServiceConfigRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *CopyServiceConfigRequest) GetServiceDesc() *string {
	return s.ServiceDesc
}

func (s *CopyServiceConfigRequest) GetServiceName() *string {
	return s.ServiceName
}

func (s *CopyServiceConfigRequest) SetRegionId(v string) *CopyServiceConfigRequest {
	s.RegionId = &v
	return s
}

func (s *CopyServiceConfigRequest) SetResourceType(v string) *CopyServiceConfigRequest {
	s.ResourceType = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceCode(v string) *CopyServiceConfigRequest {
	s.ServiceCode = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceDesc(v string) *CopyServiceConfigRequest {
	s.ServiceDesc = &v
	return s
}

func (s *CopyServiceConfigRequest) SetServiceName(v string) *CopyServiceConfigRequest {
	s.ServiceName = &v
	return s
}

func (s *CopyServiceConfigRequest) Validate() error {
	return dara.Validate(s)
}
