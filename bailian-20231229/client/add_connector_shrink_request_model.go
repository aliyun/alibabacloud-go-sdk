// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddConnectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorName(v string) *AddConnectorShrinkRequest
	GetConnectorName() *string
	SetConnectorType(v string) *AddConnectorShrinkRequest
	GetConnectorType() *string
	SetDescription(v string) *AddConnectorShrinkRequest
	GetDescription() *string
	SetFileConnectorConfigShrink(v string) *AddConnectorShrinkRequest
	GetFileConnectorConfigShrink() *string
}

type AddConnectorShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-connector
	ConnectorName *string `json:"ConnectorName,omitempty" xml:"ConnectorName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// FILE
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// This parameter is required.
	Description               *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileConnectorConfigShrink *string `json:"FileConnectorConfig,omitempty" xml:"FileConnectorConfig,omitempty"`
}

func (s AddConnectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddConnectorShrinkRequest) GetConnectorName() *string {
	return s.ConnectorName
}

func (s *AddConnectorShrinkRequest) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *AddConnectorShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *AddConnectorShrinkRequest) GetFileConnectorConfigShrink() *string {
	return s.FileConnectorConfigShrink
}

func (s *AddConnectorShrinkRequest) SetConnectorName(v string) *AddConnectorShrinkRequest {
	s.ConnectorName = &v
	return s
}

func (s *AddConnectorShrinkRequest) SetConnectorType(v string) *AddConnectorShrinkRequest {
	s.ConnectorType = &v
	return s
}

func (s *AddConnectorShrinkRequest) SetDescription(v string) *AddConnectorShrinkRequest {
	s.Description = &v
	return s
}

func (s *AddConnectorShrinkRequest) SetFileConnectorConfigShrink(v string) *AddConnectorShrinkRequest {
	s.FileConnectorConfigShrink = &v
	return s
}

func (s *AddConnectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
