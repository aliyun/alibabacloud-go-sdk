// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationName(v string) *GetApplicationRequest
	GetApplicationName() *string
	SetClusterId(v string) *GetApplicationRequest
	GetClusterId() *string
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
	// HDFS
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-e6a9d46e92675****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
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

func (s *GetApplicationRequest) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *GetApplicationRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetApplicationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApplicationRequest) SetApplicationName(v string) *GetApplicationRequest {
	s.ApplicationName = &v
	return s
}

func (s *GetApplicationRequest) SetClusterId(v string) *GetApplicationRequest {
	s.ClusterId = &v
	return s
}

func (s *GetApplicationRequest) SetRegionId(v string) *GetApplicationRequest {
	s.RegionId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
