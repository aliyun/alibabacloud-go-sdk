// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrainingJobLatestMetricsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNames(v string) *GetTrainingJobLatestMetricsRequest
	GetNames() *string
}

type GetTrainingJobLatestMetricsRequest struct {
	// example:
	//
	// loss
	Names *string `json:"Names,omitempty" xml:"Names,omitempty"`
}

func (s GetTrainingJobLatestMetricsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTrainingJobLatestMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetTrainingJobLatestMetricsRequest) GetNames() *string {
	return s.Names
}

func (s *GetTrainingJobLatestMetricsRequest) SetNames(v string) *GetTrainingJobLatestMetricsRequest {
	s.Names = &v
	return s
}

func (s *GetTrainingJobLatestMetricsRequest) Validate() error {
	return dara.Validate(s)
}
