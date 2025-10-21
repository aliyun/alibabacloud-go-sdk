// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBriefDeploymentTarget interface {
	dara.Model
	String() string
	GoString() string
	SetMode(v string) *BriefDeploymentTarget
	GetMode() *string
	SetName(v string) *BriefDeploymentTarget
	GetName() *string
}

type BriefDeploymentTarget struct {
	Mode *string `json:"mode,omitempty" xml:"mode,omitempty"`
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s BriefDeploymentTarget) String() string {
	return dara.Prettify(s)
}

func (s BriefDeploymentTarget) GoString() string {
	return s.String()
}

func (s *BriefDeploymentTarget) GetMode() *string {
	return s.Mode
}

func (s *BriefDeploymentTarget) GetName() *string {
	return s.Name
}

func (s *BriefDeploymentTarget) SetMode(v string) *BriefDeploymentTarget {
	s.Mode = &v
	return s
}

func (s *BriefDeploymentTarget) SetName(v string) *BriefDeploymentTarget {
	s.Name = &v
	return s
}

func (s *BriefDeploymentTarget) Validate() error {
	return dara.Validate(s)
}
