// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *GetApplicationRequest
	GetName() *string
	SetRegionId(v string) *GetApplicationRequest
	GetRegionId() *string
}

type GetApplicationRequest struct {
	// The application name.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyApplication
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID. Set the value to cn-hangzhou.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetName() *string {
	return s.Name
}

func (s *GetApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationRequest) SetName(v string) *GetApplicationRequest {
	s.Name = &v
	return s
}

func (s *GetApplicationRequest) SetRegionId(v string) *GetApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
