// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachApplication2ConnectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIdsShrink(v string) *AttachApplication2ConnectorShrinkRequest
	GetApplicationIdsShrink() *string
	SetConnectorId(v string) *AttachApplication2ConnectorShrinkRequest
	GetConnectorId() *string
}

type AttachApplication2ConnectorShrinkRequest struct {
	// This parameter is required.
	ApplicationIdsShrink *string `json:"ApplicationIds,omitempty" xml:"ApplicationIds,omitempty"`
	// ConnectorID。
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-94db94e06b98****
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
}

func (s AttachApplication2ConnectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachApplication2ConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *AttachApplication2ConnectorShrinkRequest) GetApplicationIdsShrink() *string {
	return s.ApplicationIdsShrink
}

func (s *AttachApplication2ConnectorShrinkRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *AttachApplication2ConnectorShrinkRequest) SetApplicationIdsShrink(v string) *AttachApplication2ConnectorShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *AttachApplication2ConnectorShrinkRequest) SetConnectorId(v string) *AttachApplication2ConnectorShrinkRequest {
	s.ConnectorId = &v
	return s
}

func (s *AttachApplication2ConnectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
