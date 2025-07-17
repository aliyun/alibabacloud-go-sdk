// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenXtraceDefaultSLRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *OpenXtraceDefaultSLRRequest
	GetRegionId() *string
}

type OpenXtraceDefaultSLRRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenXtraceDefaultSLRRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenXtraceDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenXtraceDefaultSLRRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenXtraceDefaultSLRRequest) SetRegionId(v string) *OpenXtraceDefaultSLRRequest {
	s.RegionId = &v
	return s
}

func (s *OpenXtraceDefaultSLRRequest) Validate() error {
	return dara.Validate(s)
}
