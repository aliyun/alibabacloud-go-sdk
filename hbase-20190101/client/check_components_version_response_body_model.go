// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckComponentsVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v *CheckComponentsVersionResponseBodyComponents) *CheckComponentsVersionResponseBody
	GetComponents() *CheckComponentsVersionResponseBodyComponents
	SetRequestId(v string) *CheckComponentsVersionResponseBody
	GetRequestId() *string
}

type CheckComponentsVersionResponseBody struct {
	Components *CheckComponentsVersionResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	// example:
	//
	// E3537EB4-1100-41CA-A147-C74CCC8BB12C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckComponentsVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckComponentsVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBody) GetComponents() *CheckComponentsVersionResponseBodyComponents {
	return s.Components
}

func (s *CheckComponentsVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckComponentsVersionResponseBody) SetComponents(v *CheckComponentsVersionResponseBodyComponents) *CheckComponentsVersionResponseBody {
	s.Components = v
	return s
}

func (s *CheckComponentsVersionResponseBody) SetRequestId(v string) *CheckComponentsVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckComponentsVersionResponseBody) Validate() error {
	if s.Components != nil {
		if err := s.Components.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckComponentsVersionResponseBodyComponents struct {
	Component []*CheckComponentsVersionResponseBodyComponentsComponent `json:"Component,omitempty" xml:"Component,omitempty" type:"Repeated"`
}

func (s CheckComponentsVersionResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponents) GetComponent() []*CheckComponentsVersionResponseBodyComponentsComponent {
	return s.Component
}

func (s *CheckComponentsVersionResponseBodyComponents) SetComponent(v []*CheckComponentsVersionResponseBodyComponentsComponent) *CheckComponentsVersionResponseBodyComponents {
	s.Component = v
	return s
}

func (s *CheckComponentsVersionResponseBodyComponents) Validate() error {
	if s.Component != nil {
		for _, item := range s.Component {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckComponentsVersionResponseBodyComponentsComponent struct {
	// example:
	//
	// HBASE
	Component *string `json:"Component,omitempty" xml:"Component,omitempty"`
	// example:
	//
	// true
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) String() string {
	return dara.Prettify(s)
}

func (s CheckComponentsVersionResponseBodyComponentsComponent) GoString() string {
	return s.String()
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) GetComponent() *string {
	return s.Component
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) GetIsLatestVersion() *string {
	return s.IsLatestVersion
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetComponent(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.Component = &v
	return s
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) SetIsLatestVersion(v string) *CheckComponentsVersionResponseBodyComponentsComponent {
	s.IsLatestVersion = &v
	return s
}

func (s *CheckComponentsVersionResponseBodyComponentsComponent) Validate() error {
	return dara.Validate(s)
}
