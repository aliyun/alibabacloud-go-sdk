// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentAddonsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironmentId(v string) *ListEnvironmentAddonsRequest
	GetEnvironmentId() *string
	SetRegionId(v string) *ListEnvironmentAddonsRequest
	GetRegionId() *string
}

type ListEnvironmentAddonsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// env-xxx
	EnvironmentId *string `json:"EnvironmentId,omitempty" xml:"EnvironmentId,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListEnvironmentAddonsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentAddonsRequest) GoString() string {
	return s.String()
}

func (s *ListEnvironmentAddonsRequest) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *ListEnvironmentAddonsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListEnvironmentAddonsRequest) SetEnvironmentId(v string) *ListEnvironmentAddonsRequest {
	s.EnvironmentId = &v
	return s
}

func (s *ListEnvironmentAddonsRequest) SetRegionId(v string) *ListEnvironmentAddonsRequest {
	s.RegionId = &v
	return s
}

func (s *ListEnvironmentAddonsRequest) Validate() error {
	return dara.Validate(s)
}
