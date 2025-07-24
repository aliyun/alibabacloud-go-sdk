// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetUserInformationRequest
	GetRegionId() *string
}

type GetUserInformationRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetUserInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserInformationRequest) GoString() string {
	return s.String()
}

func (s *GetUserInformationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserInformationRequest) SetRegionId(v string) *GetUserInformationRequest {
	s.RegionId = &v
	return s
}

func (s *GetUserInformationRequest) Validate() error {
	return dara.Validate(s)
}
