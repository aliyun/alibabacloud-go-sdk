// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenArmsDefaultSLRRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *OpenArmsDefaultSLRRequest
	GetRegionId() *string
}

type OpenArmsDefaultSLRRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s OpenArmsDefaultSLRRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenArmsDefaultSLRRequest) GoString() string {
	return s.String()
}

func (s *OpenArmsDefaultSLRRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenArmsDefaultSLRRequest) SetRegionId(v string) *OpenArmsDefaultSLRRequest {
	s.RegionId = &v
	return s
}

func (s *OpenArmsDefaultSLRRequest) Validate() error {
	return dara.Validate(s)
}
