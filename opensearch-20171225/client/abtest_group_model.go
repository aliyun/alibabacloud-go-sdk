// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iABTestGroup interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ABTestGroup
	GetName() *string
	SetStatus(v int32) *ABTestGroup
	GetStatus() *int32
}

type ABTestGroup struct {
	// The alias of the test group.
	//
	// example:
	//
	// “kevin_test”
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the test group. Valid values:
	//
	// 	- 0: not in effect
	//
	// 	- 1: in effect
	//
	// example:
	//
	// 1
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ABTestGroup) String() string {
	return dara.Prettify(s)
}

func (s ABTestGroup) GoString() string {
	return s.String()
}

func (s *ABTestGroup) GetName() *string {
	return s.Name
}

func (s *ABTestGroup) GetStatus() *int32 {
	return s.Status
}

func (s *ABTestGroup) SetName(v string) *ABTestGroup {
	s.Name = &v
	return s
}

func (s *ABTestGroup) SetStatus(v int32) *ABTestGroup {
	s.Status = &v
	return s
}

func (s *ABTestGroup) Validate() error {
	return dara.Validate(s)
}
