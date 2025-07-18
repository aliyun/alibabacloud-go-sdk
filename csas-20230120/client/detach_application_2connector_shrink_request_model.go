// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachApplication2ConnectorShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationIdsShrink(v string) *DetachApplication2ConnectorShrinkRequest
	GetApplicationIdsShrink() *string
	SetConnectorId(v string) *DetachApplication2ConnectorShrinkRequest
	GetConnectorId() *string
}

type DetachApplication2ConnectorShrinkRequest struct {
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

func (s DetachApplication2ConnectorShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachApplication2ConnectorShrinkRequest) GoString() string {
	return s.String()
}

func (s *DetachApplication2ConnectorShrinkRequest) GetApplicationIdsShrink() *string {
	return s.ApplicationIdsShrink
}

func (s *DetachApplication2ConnectorShrinkRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DetachApplication2ConnectorShrinkRequest) SetApplicationIdsShrink(v string) *DetachApplication2ConnectorShrinkRequest {
	s.ApplicationIdsShrink = &v
	return s
}

func (s *DetachApplication2ConnectorShrinkRequest) SetConnectorId(v string) *DetachApplication2ConnectorShrinkRequest {
	s.ConnectorId = &v
	return s
}

func (s *DetachApplication2ConnectorShrinkRequest) Validate() error {
	return dara.Validate(s)
}
