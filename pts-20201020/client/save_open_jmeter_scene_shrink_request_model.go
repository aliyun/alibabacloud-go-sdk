// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveOpenJMeterSceneShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpenJMeterSceneShrink(v string) *SaveOpenJMeterSceneShrinkRequest
	GetOpenJMeterSceneShrink() *string
}

type SaveOpenJMeterSceneShrinkRequest struct {
	// The details of the scenario.
	//
	// This parameter is required.
	OpenJMeterSceneShrink *string `json:"OpenJMeterScene,omitempty" xml:"OpenJMeterScene,omitempty"`
}

func (s SaveOpenJMeterSceneShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveOpenJMeterSceneShrinkRequest) GoString() string {
	return s.String()
}

func (s *SaveOpenJMeterSceneShrinkRequest) GetOpenJMeterSceneShrink() *string {
	return s.OpenJMeterSceneShrink
}

func (s *SaveOpenJMeterSceneShrinkRequest) SetOpenJMeterSceneShrink(v string) *SaveOpenJMeterSceneShrinkRequest {
	s.OpenJMeterSceneShrink = &v
	return s
}

func (s *SaveOpenJMeterSceneShrinkRequest) Validate() error {
	return dara.Validate(s)
}
