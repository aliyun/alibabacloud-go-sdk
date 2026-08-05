// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigData(v map[string]interface{}) *UpdateConfigRequest
	GetConfigData() map[string]interface{}
	SetDryRun(v bool) *UpdateConfigRequest
	GetDryRun() *bool
}

type UpdateConfigRequest struct {
	// The configuration content.
	ConfigData map[string]interface{} `json:"configData,omitempty" xml:"configData,omitempty"`
	// Specifies whether this is a dry run request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigRequest) GetConfigData() map[string]interface{} {
	return s.ConfigData
}

func (s *UpdateConfigRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateConfigRequest) SetConfigData(v map[string]interface{}) *UpdateConfigRequest {
	s.ConfigData = v
	return s
}

func (s *UpdateConfigRequest) SetDryRun(v bool) *UpdateConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateConfigRequest) Validate() error {
	return dara.Validate(s)
}
