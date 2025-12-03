// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckVersionsOfComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComponents(v *CheckVersionsOfComponentsResponseBodyComponents) *CheckVersionsOfComponentsResponseBody
	GetComponents() *CheckVersionsOfComponentsResponseBodyComponents
	SetRequestId(v string) *CheckVersionsOfComponentsResponseBody
	GetRequestId() *string
}

type CheckVersionsOfComponentsResponseBody struct {
	Components *CheckVersionsOfComponentsResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Struct"`
	RequestId  *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckVersionsOfComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBody) GetComponents() *CheckVersionsOfComponentsResponseBodyComponents {
	return s.Components
}

func (s *CheckVersionsOfComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckVersionsOfComponentsResponseBody) SetComponents(v *CheckVersionsOfComponentsResponseBodyComponents) *CheckVersionsOfComponentsResponseBody {
	s.Components = v
	return s
}

func (s *CheckVersionsOfComponentsResponseBody) SetRequestId(v string) *CheckVersionsOfComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckVersionsOfComponentsResponseBody) Validate() error {
	if s.Components != nil {
		if err := s.Components.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckVersionsOfComponentsResponseBodyComponents struct {
	Components []*CheckVersionsOfComponentsResponseBodyComponentsComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
}

func (s CheckVersionsOfComponentsResponseBodyComponents) String() string {
	return dara.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBodyComponents) GetComponents() []*CheckVersionsOfComponentsResponseBodyComponentsComponents {
	return s.Components
}

func (s *CheckVersionsOfComponentsResponseBodyComponents) SetComponents(v []*CheckVersionsOfComponentsResponseBodyComponentsComponents) *CheckVersionsOfComponentsResponseBodyComponents {
	s.Components = v
	return s
}

func (s *CheckVersionsOfComponentsResponseBodyComponents) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckVersionsOfComponentsResponseBodyComponentsComponents struct {
	Component       *string `json:"Component,omitempty" xml:"Component,omitempty"`
	IsLatestVersion *string `json:"IsLatestVersion,omitempty" xml:"IsLatestVersion,omitempty"`
}

func (s CheckVersionsOfComponentsResponseBodyComponentsComponents) String() string {
	return dara.Prettify(s)
}

func (s CheckVersionsOfComponentsResponseBodyComponentsComponents) GoString() string {
	return s.String()
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) GetComponent() *string {
	return s.Component
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) GetIsLatestVersion() *string {
	return s.IsLatestVersion
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) SetComponent(v string) *CheckVersionsOfComponentsResponseBodyComponentsComponents {
	s.Component = &v
	return s
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) SetIsLatestVersion(v string) *CheckVersionsOfComponentsResponseBodyComponentsComponents {
	s.IsLatestVersion = &v
	return s
}

func (s *CheckVersionsOfComponentsResponseBodyComponentsComponents) Validate() error {
	return dara.Validate(s)
}
