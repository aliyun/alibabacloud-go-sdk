// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetV2Request interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Resource) *UpdateDeploymentTargetV2Request
	GetBody() *Resource
}

type UpdateDeploymentTargetV2Request struct {
	Body *Resource `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetV2Request) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetV2Request) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetV2Request) GetBody() *Resource {
	return s.Body
}

func (s *UpdateDeploymentTargetV2Request) SetBody(v *Resource) *UpdateDeploymentTargetV2Request {
	s.Body = v
	return s
}

func (s *UpdateDeploymentTargetV2Request) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
