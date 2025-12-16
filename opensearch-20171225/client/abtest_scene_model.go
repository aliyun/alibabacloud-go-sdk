// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iABTestScene interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ABTestScene
	GetName() *string
	SetStatus(v int32) *ABTestScene
	GetStatus() *int32
	SetValues(v []*string) *ABTestScene
	GetValues() []*string
}

type ABTestScene struct {
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 0 停止实验 1 开通实验
	Status *int32    `json:"status,omitempty" xml:"status,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ABTestScene) String() string {
	return dara.Prettify(s)
}

func (s ABTestScene) GoString() string {
	return s.String()
}

func (s *ABTestScene) GetName() *string {
	return s.Name
}

func (s *ABTestScene) GetStatus() *int32 {
	return s.Status
}

func (s *ABTestScene) GetValues() []*string {
	return s.Values
}

func (s *ABTestScene) SetName(v string) *ABTestScene {
	s.Name = &v
	return s
}

func (s *ABTestScene) SetStatus(v int32) *ABTestScene {
	s.Status = &v
	return s
}

func (s *ABTestScene) SetValues(v []*string) *ABTestScene {
	s.Values = v
	return s
}

func (s *ABTestScene) Validate() error {
	return dara.Validate(s)
}
