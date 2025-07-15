// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteParameterRequest
	GetName() *string
	SetRegionId(v string) *DeleteParameterRequest
	GetRegionId() *string
}

type DeleteParameterRequest struct {
	// The name of the common parameter. The name can be up to 180 characters in length and can contain only letters, digits, hyphens (-), and underscores (_). It cannot start with aliyun, acs, alibaba, alicloud, or oos.
	//
	// This parameter is required.
	//
	// example:
	//
	// MyParameter
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterRequest) GetName() *string {
	return s.Name
}

func (s *DeleteParameterRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteParameterRequest) SetName(v string) *DeleteParameterRequest {
	s.Name = &v
	return s
}

func (s *DeleteParameterRequest) SetRegionId(v string) *DeleteParameterRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteParameterRequest) Validate() error {
	return dara.Validate(s)
}
