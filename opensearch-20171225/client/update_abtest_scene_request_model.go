// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestScene) *UpdateABTestSceneRequest
	GetBody() *ABTestScene
	SetDryRun(v bool) *UpdateABTestSceneRequest
	GetDryRun() *bool
}

type UpdateABTestSceneRequest struct {
	// The request body.
	Body *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateABTestSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestSceneRequest) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneRequest) GetBody() *ABTestScene {
	return s.Body
}

func (s *UpdateABTestSceneRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateABTestSceneRequest) SetBody(v *ABTestScene) *UpdateABTestSceneRequest {
	s.Body = v
	return s
}

func (s *UpdateABTestSceneRequest) SetDryRun(v bool) *UpdateABTestSceneRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateABTestSceneRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
