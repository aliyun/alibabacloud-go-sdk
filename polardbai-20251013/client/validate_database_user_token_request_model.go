// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDatabaseUserTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ValidateDatabaseUserTokenRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ValidateDatabaseUserTokenRequest
	GetAuthType() *string
	SetDBClusterId(v string) *ValidateDatabaseUserTokenRequest
	GetDBClusterId() *string
	SetDBName(v string) *ValidateDatabaseUserTokenRequest
	GetDBName() *string
}

type ValidateDatabaseUserTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// yEtNjzuM21NVLVJTuL9Trw****
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// token
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ai_test
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
}

func (s ValidateDatabaseUserTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateDatabaseUserTokenRequest) GoString() string {
	return s.String()
}

func (s *ValidateDatabaseUserTokenRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ValidateDatabaseUserTokenRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ValidateDatabaseUserTokenRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ValidateDatabaseUserTokenRequest) GetDBName() *string {
	return s.DBName
}

func (s *ValidateDatabaseUserTokenRequest) SetAuthMessage(v string) *ValidateDatabaseUserTokenRequest {
	s.AuthMessage = &v
	return s
}

func (s *ValidateDatabaseUserTokenRequest) SetAuthType(v string) *ValidateDatabaseUserTokenRequest {
	s.AuthType = &v
	return s
}

func (s *ValidateDatabaseUserTokenRequest) SetDBClusterId(v string) *ValidateDatabaseUserTokenRequest {
	s.DBClusterId = &v
	return s
}

func (s *ValidateDatabaseUserTokenRequest) SetDBName(v string) *ValidateDatabaseUserTokenRequest {
	s.DBName = &v
	return s
}

func (s *ValidateDatabaseUserTokenRequest) Validate() error {
	return dara.Validate(s)
}
