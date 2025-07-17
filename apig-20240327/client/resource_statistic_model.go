// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResourceStatistic interface {
	dara.Model
	String() string
	GoString() string
	SetResourceCount(v int32) *ResourceStatistic
	GetResourceCount() *int32
	SetResourceType(v string) *ResourceStatistic
	GetResourceType() *string
}

type ResourceStatistic struct {
	ResourceCount *int32  `json:"resourceCount,omitempty" xml:"resourceCount,omitempty"`
	ResourceType  *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ResourceStatistic) String() string {
	return dara.Prettify(s)
}

func (s ResourceStatistic) GoString() string {
	return s.String()
}

func (s *ResourceStatistic) GetResourceCount() *int32 {
	return s.ResourceCount
}

func (s *ResourceStatistic) GetResourceType() *string {
	return s.ResourceType
}

func (s *ResourceStatistic) SetResourceCount(v int32) *ResourceStatistic {
	s.ResourceCount = &v
	return s
}

func (s *ResourceStatistic) SetResourceType(v string) *ResourceStatistic {
	s.ResourceType = &v
	return s
}

func (s *ResourceStatistic) Validate() error {
	return dara.Validate(s)
}
