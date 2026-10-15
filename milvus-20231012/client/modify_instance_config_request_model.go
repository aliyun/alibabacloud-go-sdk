// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceConfigRequest
	GetInstanceId() *string
	SetReason(v string) *ModifyInstanceConfigRequest
	GetReason() *string
	SetUserConfig(v string) *ModifyInstanceConfigRequest
	GetUserConfig() *string
	SetBody(v *ModifyInstanceConfigRequestBody) *ModifyInstanceConfigRequest
	GetBody() *ModifyInstanceConfigRequestBody
}

type ModifyInstanceConfigRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The reason for the update.
	//
	// This parameter is required.
	//
	// example:
	//
	// for test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The custom configuration.
	//
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"UserConfig,omitempty" xml:"UserConfig,omitempty"`
	// Submits the instance configurations to be updated as a JSON object.
	Body *ModifyInstanceConfigRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ModifyInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceConfigRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyInstanceConfigRequest) GetUserConfig() *string {
	return s.UserConfig
}

func (s *ModifyInstanceConfigRequest) GetBody() *ModifyInstanceConfigRequestBody {
	return s.Body
}

func (s *ModifyInstanceConfigRequest) SetInstanceId(v string) *ModifyInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetReason(v string) *ModifyInstanceConfigRequest {
	s.Reason = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetUserConfig(v string) *ModifyInstanceConfigRequest {
	s.UserConfig = &v
	return s
}

func (s *ModifyInstanceConfigRequest) SetBody(v *ModifyInstanceConfigRequestBody) *ModifyInstanceConfigRequest {
	s.Body = v
	return s
}

func (s *ModifyInstanceConfigRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceConfigRequestBody struct {
	// The configuration content to write to the instance, formatted as a YAML string.
	//
	// example:
	//
	// dataCoord:\\n  segment:\\n    maxSize: 1024
	UserConfig *string `json:"userConfig,omitempty" xml:"userConfig,omitempty"`
}

func (s ModifyInstanceConfigRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceConfigRequestBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceConfigRequestBody) GetUserConfig() *string {
	return s.UserConfig
}

func (s *ModifyInstanceConfigRequestBody) SetUserConfig(v string) *ModifyInstanceConfigRequestBody {
	s.UserConfig = &v
	return s
}

func (s *ModifyInstanceConfigRequestBody) Validate() error {
	return dara.Validate(s)
}
