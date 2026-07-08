// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeActionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionId(v string) *InvokeActionShrinkRequest
	GetActionId() *string
	SetActionVersion(v string) *InvokeActionShrinkRequest
	GetActionVersion() *string
	SetAuthConfigShrink(v string) *InvokeActionShrinkRequest
	GetAuthConfigShrink() *string
	SetBodyShrink(v string) *InvokeActionShrinkRequest
	GetBodyShrink() *string
	SetConnectorId(v string) *InvokeActionShrinkRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *InvokeActionShrinkRequest
	GetConnectorVersion() *string
	SetHeadersShrink(v string) *InvokeActionShrinkRequest
	GetHeadersShrink() *string
	SetPathShrink(v string) *InvokeActionShrinkRequest
	GetPathShrink() *string
	SetQueryShrink(v string) *InvokeActionShrinkRequest
	GetQueryShrink() *string
	SetStream(v bool) *InvokeActionShrinkRequest
	GetStream() *bool
}

type InvokeActionShrinkRequest struct {
	// The action ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// action-xxx
	ActionId *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	// The action version.
	//
	// example:
	//
	// 1
	ActionVersion *string `json:"ActionVersion,omitempty" xml:"ActionVersion,omitempty"`
	// The authentication information for the action.
	AuthConfigShrink *string `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty"`
	// The request body for the action.
	BodyShrink *string `json:"Body,omitempty" xml:"Body,omitempty"`
	// The connector ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-xxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The connector version.
	//
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	// The request header parameters for the action.
	HeadersShrink *string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	// The path parameters for the action.
	PathShrink *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The query parameters for the action.
	QueryShrink *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// Specifies whether to use streaming output.
	//
	// example:
	//
	// false
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s InvokeActionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionShrinkRequest) GoString() string {
	return s.String()
}

func (s *InvokeActionShrinkRequest) GetActionId() *string {
	return s.ActionId
}

func (s *InvokeActionShrinkRequest) GetActionVersion() *string {
	return s.ActionVersion
}

func (s *InvokeActionShrinkRequest) GetAuthConfigShrink() *string {
	return s.AuthConfigShrink
}

func (s *InvokeActionShrinkRequest) GetBodyShrink() *string {
	return s.BodyShrink
}

func (s *InvokeActionShrinkRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *InvokeActionShrinkRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *InvokeActionShrinkRequest) GetHeadersShrink() *string {
	return s.HeadersShrink
}

func (s *InvokeActionShrinkRequest) GetPathShrink() *string {
	return s.PathShrink
}

func (s *InvokeActionShrinkRequest) GetQueryShrink() *string {
	return s.QueryShrink
}

func (s *InvokeActionShrinkRequest) GetStream() *bool {
	return s.Stream
}

func (s *InvokeActionShrinkRequest) SetActionId(v string) *InvokeActionShrinkRequest {
	s.ActionId = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetActionVersion(v string) *InvokeActionShrinkRequest {
	s.ActionVersion = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetAuthConfigShrink(v string) *InvokeActionShrinkRequest {
	s.AuthConfigShrink = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetBodyShrink(v string) *InvokeActionShrinkRequest {
	s.BodyShrink = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetConnectorId(v string) *InvokeActionShrinkRequest {
	s.ConnectorId = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetConnectorVersion(v string) *InvokeActionShrinkRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetHeadersShrink(v string) *InvokeActionShrinkRequest {
	s.HeadersShrink = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetPathShrink(v string) *InvokeActionShrinkRequest {
	s.PathShrink = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetQueryShrink(v string) *InvokeActionShrinkRequest {
	s.QueryShrink = &v
	return s
}

func (s *InvokeActionShrinkRequest) SetStream(v bool) *InvokeActionShrinkRequest {
	s.Stream = &v
	return s
}

func (s *InvokeActionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
