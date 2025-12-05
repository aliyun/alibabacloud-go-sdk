// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeploymentsByNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreJobSummary(v bool) *GetDeploymentsByNameRequest
	GetIgnoreJobSummary() *bool
	SetIgnoreResourceSetting(v bool) *GetDeploymentsByNameRequest
	GetIgnoreResourceSetting() *bool
}

type GetDeploymentsByNameRequest struct {
	// example:
	//
	// true
	IgnoreJobSummary *bool `json:"ignoreJobSummary,omitempty" xml:"ignoreJobSummary,omitempty"`
	// example:
	//
	// true
	IgnoreResourceSetting *bool `json:"ignoreResourceSetting,omitempty" xml:"ignoreResourceSetting,omitempty"`
}

func (s GetDeploymentsByNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeploymentsByNameRequest) GoString() string {
	return s.String()
}

func (s *GetDeploymentsByNameRequest) GetIgnoreJobSummary() *bool {
	return s.IgnoreJobSummary
}

func (s *GetDeploymentsByNameRequest) GetIgnoreResourceSetting() *bool {
	return s.IgnoreResourceSetting
}

func (s *GetDeploymentsByNameRequest) SetIgnoreJobSummary(v bool) *GetDeploymentsByNameRequest {
	s.IgnoreJobSummary = &v
	return s
}

func (s *GetDeploymentsByNameRequest) SetIgnoreResourceSetting(v bool) *GetDeploymentsByNameRequest {
	s.IgnoreResourceSetting = &v
	return s
}

func (s *GetDeploymentsByNameRequest) Validate() error {
	return dara.Validate(s)
}
