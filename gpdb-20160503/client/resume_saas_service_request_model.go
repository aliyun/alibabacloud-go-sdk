// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeSaasServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *ResumeSaasServiceRequest
	GetRegionId() *string
	SetServiceId(v string) *ResumeSaasServiceRequest
	GetServiceId() *string
}

type ResumeSaasServiceRequest struct {
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// agdb-xxxxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ResumeSaasServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeSaasServiceRequest) GoString() string {
	return s.String()
}

func (s *ResumeSaasServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResumeSaasServiceRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *ResumeSaasServiceRequest) SetRegionId(v string) *ResumeSaasServiceRequest {
	s.RegionId = &v
	return s
}

func (s *ResumeSaasServiceRequest) SetServiceId(v string) *ResumeSaasServiceRequest {
	s.ServiceId = &v
	return s
}

func (s *ResumeSaasServiceRequest) Validate() error {
	return dara.Validate(s)
}
