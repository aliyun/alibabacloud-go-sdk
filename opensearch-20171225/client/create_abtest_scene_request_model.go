// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestSceneRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ABTestScene) *CreateABTestSceneRequest
	GetBody() *ABTestScene
	SetDryRun(v bool) *CreateABTestSceneRequest
	GetDryRun() *bool
}

type CreateABTestSceneRequest struct {
	// The A/B test scene. For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html).
	Body *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to validate only the request parameters. The default value is false.
	//
	// Values:
	//
	// - **true**: Validates only the request parameters.
	//
	// - **false**: Validates the request parameters and creates the attribution configuration.
	//
	// example:
	//
	// false
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateABTestSceneRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestSceneRequest) GoString() string {
	return s.String()
}

func (s *CreateABTestSceneRequest) GetBody() *ABTestScene {
	return s.Body
}

func (s *CreateABTestSceneRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateABTestSceneRequest) SetBody(v *ABTestScene) *CreateABTestSceneRequest {
	s.Body = v
	return s
}

func (s *CreateABTestSceneRequest) SetDryRun(v bool) *CreateABTestSceneRequest {
	s.DryRun = &v
	return s
}

func (s *CreateABTestSceneRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
