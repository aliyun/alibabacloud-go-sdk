// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigData(v map[string]interface{}) *CreateConfigRequest
	GetConfigData() map[string]interface{}
	SetDryRun(v bool) *CreateConfigRequest
	GetDryRun() *bool
}

type CreateConfigRequest struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// Specifies whether the request is a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateConfigRequest) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *CreateConfigRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateConfigRequest) SetConfigData(v map[string]interface{}) *CreateConfigRequest {
	s.ConfigData = v
	return s
}

func (s *CreateConfigRequest) SetDryRun(v bool) *CreateConfigRequest {
	s.DryRun = &v
	return s
}

func (s *CreateConfigRequest) Validate() error {
	return dara.Validate(s)
}
