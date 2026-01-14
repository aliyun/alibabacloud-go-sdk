// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInvalidDomainCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetInvalidDomainCountRequest
	GetRegionId() *string
}

type GetInvalidDomainCountRequest struct {
	// The ID of the region where the Global Accelerator (GA) instance is deployed. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetInvalidDomainCountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInvalidDomainCountRequest) GoString() string {
	return s.String()
}

func (s *GetInvalidDomainCountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInvalidDomainCountRequest) SetRegionId(v string) *GetInvalidDomainCountRequest {
	s.RegionId = &v
	return s
}

func (s *GetInvalidDomainCountRequest) Validate() error {
	return dara.Validate(s)
}
