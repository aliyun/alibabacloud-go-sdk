// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateHttpApiOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfig(v *BatchUpdateHttpApiOperationRequestAuthConfig) *BatchUpdateHttpApiOperationRequest
	GetAuthConfig() *BatchUpdateHttpApiOperationRequestAuthConfig
	SetEnableAuth(v bool) *BatchUpdateHttpApiOperationRequest
	GetEnableAuth() *bool
	SetOperationIds(v []*string) *BatchUpdateHttpApiOperationRequest
	GetOperationIds() []*string
}

type BatchUpdateHttpApiOperationRequest struct {
	AuthConfig *BatchUpdateHttpApiOperationRequestAuthConfig `json:"authConfig,omitempty" xml:"authConfig,omitempty" type:"Struct"`
	// example:
	//
	// true
	EnableAuth   *bool     `json:"enableAuth,omitempty" xml:"enableAuth,omitempty"`
	OperationIds []*string `json:"operationIds,omitempty" xml:"operationIds,omitempty" type:"Repeated"`
}

func (s BatchUpdateHttpApiOperationRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateHttpApiOperationRequest) GoString() string {
	return s.String()
}

func (s *BatchUpdateHttpApiOperationRequest) GetAuthConfig() *BatchUpdateHttpApiOperationRequestAuthConfig {
	return s.AuthConfig
}

func (s *BatchUpdateHttpApiOperationRequest) GetEnableAuth() *bool {
	return s.EnableAuth
}

func (s *BatchUpdateHttpApiOperationRequest) GetOperationIds() []*string {
	return s.OperationIds
}

func (s *BatchUpdateHttpApiOperationRequest) SetAuthConfig(v *BatchUpdateHttpApiOperationRequestAuthConfig) *BatchUpdateHttpApiOperationRequest {
	s.AuthConfig = v
	return s
}

func (s *BatchUpdateHttpApiOperationRequest) SetEnableAuth(v bool) *BatchUpdateHttpApiOperationRequest {
	s.EnableAuth = &v
	return s
}

func (s *BatchUpdateHttpApiOperationRequest) SetOperationIds(v []*string) *BatchUpdateHttpApiOperationRequest {
	s.OperationIds = v
	return s
}

func (s *BatchUpdateHttpApiOperationRequest) Validate() error {
	if s.AuthConfig != nil {
		if err := s.AuthConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUpdateHttpApiOperationRequestAuthConfig struct {
	// example:
	//
	// Custom
	AuthMode *string `json:"authMode,omitempty" xml:"authMode,omitempty"`
	// example:
	//
	// Jwt
	AuthType *string `json:"authType,omitempty" xml:"authType,omitempty"`
}

func (s BatchUpdateHttpApiOperationRequestAuthConfig) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateHttpApiOperationRequestAuthConfig) GoString() string {
	return s.String()
}

func (s *BatchUpdateHttpApiOperationRequestAuthConfig) GetAuthMode() *string {
	return s.AuthMode
}

func (s *BatchUpdateHttpApiOperationRequestAuthConfig) GetAuthType() *string {
	return s.AuthType
}

func (s *BatchUpdateHttpApiOperationRequestAuthConfig) SetAuthMode(v string) *BatchUpdateHttpApiOperationRequestAuthConfig {
	s.AuthMode = &v
	return s
}

func (s *BatchUpdateHttpApiOperationRequestAuthConfig) SetAuthType(v string) *BatchUpdateHttpApiOperationRequestAuthConfig {
	s.AuthType = &v
	return s
}

func (s *BatchUpdateHttpApiOperationRequestAuthConfig) Validate() error {
	return dara.Validate(s)
}
