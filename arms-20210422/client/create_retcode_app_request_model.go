// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRetcodeAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *CreateRetcodeAppRequest
	GetRegionId() *string
	SetRetcodeAppName(v string) *CreateRetcodeAppRequest
	GetRetcodeAppName() *string
	SetRetcodeAppType(v string) *CreateRetcodeAppRequest
	GetRetcodeAppType() *string
}

type CreateRetcodeAppRequest struct {
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	RetcodeAppName *string `json:"RetcodeAppName,omitempty" xml:"RetcodeAppName,omitempty"`
	// This parameter is required.
	RetcodeAppType *string `json:"RetcodeAppType,omitempty" xml:"RetcodeAppType,omitempty"`
}

func (s CreateRetcodeAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRetcodeAppRequest) GoString() string {
	return s.String()
}

func (s *CreateRetcodeAppRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRetcodeAppRequest) GetRetcodeAppName() *string {
	return s.RetcodeAppName
}

func (s *CreateRetcodeAppRequest) GetRetcodeAppType() *string {
	return s.RetcodeAppType
}

func (s *CreateRetcodeAppRequest) SetRegionId(v string) *CreateRetcodeAppRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppName(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppName = &v
	return s
}

func (s *CreateRetcodeAppRequest) SetRetcodeAppType(v string) *CreateRetcodeAppRequest {
	s.RetcodeAppType = &v
	return s
}

func (s *CreateRetcodeAppRequest) Validate() error {
	return dara.Validate(s)
}
