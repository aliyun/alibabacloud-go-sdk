// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDemonstrationForCustomizedVoiceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetScenario(v string) *GetDemonstrationForCustomizedVoiceJobRequest
	GetScenario() *string
}

type GetDemonstrationForCustomizedVoiceJobRequest struct {
	// The demonstration scenario.
	//
	// Valid values:
	//
	// 	- **story**
	//
	// 	- **interaction**
	//
	// 	- **navigation**
	//
	// This parameter is required.
	//
	// example:
	//
	// story
	Scenario *string `json:"Scenario,omitempty" xml:"Scenario,omitempty"`
}

func (s GetDemonstrationForCustomizedVoiceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDemonstrationForCustomizedVoiceJobRequest) GoString() string {
	return s.String()
}

func (s *GetDemonstrationForCustomizedVoiceJobRequest) GetScenario() *string {
	return s.Scenario
}

func (s *GetDemonstrationForCustomizedVoiceJobRequest) SetScenario(v string) *GetDemonstrationForCustomizedVoiceJobRequest {
	s.Scenario = &v
	return s
}

func (s *GetDemonstrationForCustomizedVoiceJobRequest) Validate() error {
	return dara.Validate(s)
}
