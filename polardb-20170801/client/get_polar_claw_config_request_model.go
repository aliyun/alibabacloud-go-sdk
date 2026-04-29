// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPolarClawConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetPolarClawConfigRequest
	GetApplicationId() *string
	SetConfigPath(v string) *GetPolarClawConfigRequest
	GetConfigPath() *string
}

type GetPolarClawConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// models
	ConfigPath *string `json:"ConfigPath,omitempty" xml:"ConfigPath,omitempty"`
}

func (s GetPolarClawConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPolarClawConfigRequest) GoString() string {
	return s.String()
}

func (s *GetPolarClawConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetPolarClawConfigRequest) GetConfigPath() *string {
	return s.ConfigPath
}

func (s *GetPolarClawConfigRequest) SetApplicationId(v string) *GetPolarClawConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetPolarClawConfigRequest) SetConfigPath(v string) *GetPolarClawConfigRequest {
	s.ConfigPath = &v
	return s
}

func (s *GetPolarClawConfigRequest) Validate() error {
	return dara.Validate(s)
}
