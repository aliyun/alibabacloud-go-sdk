// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *ConfigRuntimeModelTemplateRequest
	GetModelTemplateId() *string
	SetRuntimeIds(v []*string) *ConfigRuntimeModelTemplateRequest
	GetRuntimeIds() []*string
	SetRuntimeType(v string) *ConfigRuntimeModelTemplateRequest
	GetRuntimeType() *string
}

type ConfigRuntimeModelTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxxxxxxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
	// This parameter is required.
	RuntimeIds []*string `json:"RuntimeIds,omitempty" xml:"RuntimeIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// CloudDesktop
	RuntimeType *string `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
}

func (s ConfigRuntimeModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *ConfigRuntimeModelTemplateRequest) GetRuntimeIds() []*string {
	return s.RuntimeIds
}

func (s *ConfigRuntimeModelTemplateRequest) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *ConfigRuntimeModelTemplateRequest) SetModelTemplateId(v string) *ConfigRuntimeModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *ConfigRuntimeModelTemplateRequest) SetRuntimeIds(v []*string) *ConfigRuntimeModelTemplateRequest {
	s.RuntimeIds = v
	return s
}

func (s *ConfigRuntimeModelTemplateRequest) SetRuntimeType(v string) *ConfigRuntimeModelTemplateRequest {
	s.RuntimeType = &v
	return s
}

func (s *ConfigRuntimeModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
