// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIds(v []*string) *DescribeAppsRequest
	GetAppIds() []*string
	SetAppNames(v []*string) *DescribeAppsRequest
	GetAppNames() []*string
	SetClientToken(v string) *DescribeAppsRequest
	GetClientToken() *string
	SetOwner(v string) *DescribeAppsRequest
	GetOwner() *string
	SetRegionId(v string) *DescribeAppsRequest
	GetRegionId() *string
}

type DescribeAppsRequest struct {
	AppIds []*string `json:"AppIds,omitempty" xml:"AppIds,omitempty" type:"Repeated"`
	// example:
	//
	// App1,App2
	AppNames []*string `json:"AppNames,omitempty" xml:"AppNames,omitempty" type:"Repeated"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1485558146415628
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeAppsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppsRequest) GetAppIds() []*string {
	return s.AppIds
}

func (s *DescribeAppsRequest) GetAppNames() []*string {
	return s.AppNames
}

func (s *DescribeAppsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAppsRequest) GetOwner() *string {
	return s.Owner
}

func (s *DescribeAppsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAppsRequest) SetAppIds(v []*string) *DescribeAppsRequest {
	s.AppIds = v
	return s
}

func (s *DescribeAppsRequest) SetAppNames(v []*string) *DescribeAppsRequest {
	s.AppNames = v
	return s
}

func (s *DescribeAppsRequest) SetClientToken(v string) *DescribeAppsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAppsRequest) SetOwner(v string) *DescribeAppsRequest {
	s.Owner = &v
	return s
}

func (s *DescribeAppsRequest) SetRegionId(v string) *DescribeAppsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAppsRequest) Validate() error {
	return dara.Validate(s)
}
