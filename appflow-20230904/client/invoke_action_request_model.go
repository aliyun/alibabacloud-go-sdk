// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvokeActionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionId(v string) *InvokeActionRequest
	GetActionId() *string
	SetActionVersion(v string) *InvokeActionRequest
	GetActionVersion() *string
	SetAuthConfig(v *InvokeActionRequestAuthConfig) *InvokeActionRequest
	GetAuthConfig() *InvokeActionRequestAuthConfig
	SetBody(v map[string]interface{}) *InvokeActionRequest
	GetBody() map[string]interface{}
	SetConnectorId(v string) *InvokeActionRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *InvokeActionRequest
	GetConnectorVersion() *string
	SetHeaders(v map[string]*string) *InvokeActionRequest
	GetHeaders() map[string]*string
	SetPath(v map[string]*string) *InvokeActionRequest
	GetPath() map[string]*string
	SetQuery(v map[string]*string) *InvokeActionRequest
	GetQuery() map[string]*string
	SetStream(v bool) *InvokeActionRequest
	GetStream() *bool
}

type InvokeActionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// action-xxx
	ActionId *string `json:"ActionId,omitempty" xml:"ActionId,omitempty"`
	// example:
	//
	// 1
	ActionVersion *string                        `json:"ActionVersion,omitempty" xml:"ActionVersion,omitempty"`
	AuthConfig    *InvokeActionRequestAuthConfig `json:"AuthConfig,omitempty" xml:"AuthConfig,omitempty" type:"Struct"`
	Body          map[string]interface{}         `json:"Body,omitempty" xml:"Body,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-xxx
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1
	ConnectorVersion *string            `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
	Headers          map[string]*string `json:"Headers,omitempty" xml:"Headers,omitempty"`
	Path             map[string]*string `json:"Path,omitempty" xml:"Path,omitempty"`
	Query            map[string]*string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// false
	Stream *bool `json:"Stream,omitempty" xml:"Stream,omitempty"`
}

func (s InvokeActionRequest) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionRequest) GoString() string {
	return s.String()
}

func (s *InvokeActionRequest) GetActionId() *string {
	return s.ActionId
}

func (s *InvokeActionRequest) GetActionVersion() *string {
	return s.ActionVersion
}

func (s *InvokeActionRequest) GetAuthConfig() *InvokeActionRequestAuthConfig {
	return s.AuthConfig
}

func (s *InvokeActionRequest) GetBody() map[string]interface{} {
	return s.Body
}

func (s *InvokeActionRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *InvokeActionRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *InvokeActionRequest) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InvokeActionRequest) GetPath() map[string]*string {
	return s.Path
}

func (s *InvokeActionRequest) GetQuery() map[string]*string {
	return s.Query
}

func (s *InvokeActionRequest) GetStream() *bool {
	return s.Stream
}

func (s *InvokeActionRequest) SetActionId(v string) *InvokeActionRequest {
	s.ActionId = &v
	return s
}

func (s *InvokeActionRequest) SetActionVersion(v string) *InvokeActionRequest {
	s.ActionVersion = &v
	return s
}

func (s *InvokeActionRequest) SetAuthConfig(v *InvokeActionRequestAuthConfig) *InvokeActionRequest {
	s.AuthConfig = v
	return s
}

func (s *InvokeActionRequest) SetBody(v map[string]interface{}) *InvokeActionRequest {
	s.Body = v
	return s
}

func (s *InvokeActionRequest) SetConnectorId(v string) *InvokeActionRequest {
	s.ConnectorId = &v
	return s
}

func (s *InvokeActionRequest) SetConnectorVersion(v string) *InvokeActionRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *InvokeActionRequest) SetHeaders(v map[string]*string) *InvokeActionRequest {
	s.Headers = v
	return s
}

func (s *InvokeActionRequest) SetPath(v map[string]*string) *InvokeActionRequest {
	s.Path = v
	return s
}

func (s *InvokeActionRequest) SetQuery(v map[string]*string) *InvokeActionRequest {
	s.Query = v
	return s
}

func (s *InvokeActionRequest) SetStream(v bool) *InvokeActionRequest {
	s.Stream = &v
	return s
}

func (s *InvokeActionRequest) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InvokeActionRequestAuthConfig struct {
	// example:
	//
	// raw
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// {"AppId":"xxxx","AppSecret":"sk-xxx"}
	Value interface{} `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s InvokeActionRequestAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s InvokeActionRequestAuthConfig) GoString() string {
	return s.String()
}

func (s *InvokeActionRequestAuthConfig) GetType() *string {
	return s.Type
}

func (s *InvokeActionRequestAuthConfig) GetValue() interface{} {
	return s.Value
}

func (s *InvokeActionRequestAuthConfig) SetType(v string) *InvokeActionRequestAuthConfig {
	s.Type = &v
	return s
}

func (s *InvokeActionRequestAuthConfig) SetValue(v interface{}) *InvokeActionRequestAuthConfig {
	s.Value = v
	return s
}

func (s *InvokeActionRequestAuthConfig) Validate() error {
	return dara.Validate(s)
}
