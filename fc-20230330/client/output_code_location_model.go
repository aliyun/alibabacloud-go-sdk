// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOutputCodeLocation interface {
	dara.Model
	String() string
	GoString() string
	SetLocation(v string) *OutputCodeLocation
	GetLocation() *string
	SetRepositoryType(v string) *OutputCodeLocation
	GetRepositoryType() *string
}

type OutputCodeLocation struct {
	// example:
	//
	// https://xyz.oss-cn-shanghai.aliyuncs.com/xxx/xxx/xxx
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// OSS
	RepositoryType *string `json:"repositoryType,omitempty" xml:"repositoryType,omitempty"`
}

func (s OutputCodeLocation) String() string {
	return dara.Prettify(s)
}

func (s OutputCodeLocation) GoString() string {
	return s.String()
}

func (s *OutputCodeLocation) GetLocation() *string {
	return s.Location
}

func (s *OutputCodeLocation) GetRepositoryType() *string {
	return s.RepositoryType
}

func (s *OutputCodeLocation) SetLocation(v string) *OutputCodeLocation {
	s.Location = &v
	return s
}

func (s *OutputCodeLocation) SetRepositoryType(v string) *OutputCodeLocation {
	s.RepositoryType = &v
	return s
}

func (s *OutputCodeLocation) Validate() error {
	return dara.Validate(s)
}
