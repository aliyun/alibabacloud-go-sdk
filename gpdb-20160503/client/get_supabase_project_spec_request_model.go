// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSupabaseProjectSpecRequest
	GetRegionId() *string
}

type GetSupabaseProjectSpecRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupabaseProjectSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectSpecRequest) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectSpecRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupabaseProjectSpecRequest) SetRegionId(v string) *GetSupabaseProjectSpecRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupabaseProjectSpecRequest) Validate() error {
	return dara.Validate(s)
}
