// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSavePtsSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSceneShrink(v string) *SavePtsSceneShrinkRequest
	GetSceneShrink() *string
}

type SavePtsSceneShrinkRequest struct {
	// The information about the scenario.
	//
	// This parameter is required.
	SceneShrink *string `json:"Scene,omitempty" xml:"Scene,omitempty"`
}

func (s SavePtsSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SavePtsSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *SavePtsSceneShrinkRequest) GetSceneShrink() *string {
	return s.SceneShrink
}

func (s *SavePtsSceneShrinkRequest) SetSceneShrink(v string) *SavePtsSceneShrinkRequest {
	s.SceneShrink = &v
	return s
}

func (s *SavePtsSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
