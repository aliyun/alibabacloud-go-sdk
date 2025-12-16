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
	// The ABTest scenario. For more information, see [ABTestScene](https://help.aliyun.com/document_detail/173618.html)
	Body *ABTestScene `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to check the validity of input parameters. Default value: false.
	//
	// Valid values:
	//
	// 	- **true**: checks only the validity of input parameters.
	//
	// 	- **false**: checks the validity of input parameters and creates an attribution configuration.
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
