// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNames(v string) *GetParametersRequest
	GetNames() *string
	SetRegionId(v string) *GetParametersRequest
	GetRegionId() *string
}

type GetParametersRequest struct {
	// The names of the common parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["parameter1","parameter2"]
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParametersRequest) GoString() string {
	return s.String()
}

func (s *GetParametersRequest) GetNames() *string {
	return s.Names
}

func (s *GetParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetParametersRequest) SetNames(v string) *GetParametersRequest {
	s.Names = &v
	return s
}

func (s *GetParametersRequest) SetRegionId(v string) *GetParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetParametersRequest) Validate() error {
	return dara.Validate(s)
}
