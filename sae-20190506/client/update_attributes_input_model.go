// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAttributesInput interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateAttributesInput
	GetDescription() *string
	SetHttpTriggerConfig(v *HTTPTriggerConfig) *UpdateAttributesInput
	GetHttpTriggerConfig() *HTTPTriggerConfig
	SetVersionID(v string) *UpdateAttributesInput
	GetVersionID() *string
}

type UpdateAttributesInput struct {
	Description       *string            `json:"description,omitempty" xml:"description,omitempty"`
	HttpTriggerConfig *HTTPTriggerConfig `json:"httpTriggerConfig,omitempty" xml:"httpTriggerConfig,omitempty"`
	VersionID         *string            `json:"versionID,omitempty" xml:"versionID,omitempty"`
}

func (s UpdateAttributesInput) String() string {
	return dara.Prettify(s)
}

func (s UpdateAttributesInput) GoString() string {
	return s.String()
}

func (s *UpdateAttributesInput) GetDescription() *string {
	return s.Description
}

func (s *UpdateAttributesInput) GetHttpTriggerConfig() *HTTPTriggerConfig {
	return s.HttpTriggerConfig
}

func (s *UpdateAttributesInput) GetVersionID() *string {
	return s.VersionID
}

func (s *UpdateAttributesInput) SetDescription(v string) *UpdateAttributesInput {
	s.Description = &v
	return s
}

func (s *UpdateAttributesInput) SetHttpTriggerConfig(v *HTTPTriggerConfig) *UpdateAttributesInput {
	s.HttpTriggerConfig = v
	return s
}

func (s *UpdateAttributesInput) SetVersionID(v string) *UpdateAttributesInput {
	s.VersionID = &v
	return s
}

func (s *UpdateAttributesInput) Validate() error {
	return dara.Validate(s)
}
