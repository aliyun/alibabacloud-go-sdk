// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCatalogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v string) *ListCatalogsRequest
	GetEnvironment() *string
	SetRegionId(v string) *ListCatalogsRequest
	GetRegionId() *string
}

type ListCatalogsRequest struct {
	// example:
	//
	// dev
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListCatalogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCatalogsRequest) GoString() string {
	return s.String()
}

func (s *ListCatalogsRequest) GetEnvironment() *string {
	return s.Environment
}

func (s *ListCatalogsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCatalogsRequest) SetEnvironment(v string) *ListCatalogsRequest {
	s.Environment = &v
	return s
}

func (s *ListCatalogsRequest) SetRegionId(v string) *ListCatalogsRequest {
	s.RegionId = &v
	return s
}

func (s *ListCatalogsRequest) Validate() error {
	return dara.Validate(s)
}
