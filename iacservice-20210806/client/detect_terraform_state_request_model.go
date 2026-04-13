// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectTerraformStateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetectTerraformStateRequest
	GetClientToken() *string
	SetIdentifier(v string) *DetectTerraformStateRequest
	GetIdentifier() *string
	SetType(v string) *DetectTerraformStateRequest
	GetType() *string
}

type DetectTerraformStateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// a65451293e64979ba7a4b573950217fe
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// stack-as11xxxxxxxxx:developmentA
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Stack
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DetectTerraformStateRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectTerraformStateRequest) GoString() string {
	return s.String()
}

func (s *DetectTerraformStateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetectTerraformStateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *DetectTerraformStateRequest) GetType() *string {
	return s.Type
}

func (s *DetectTerraformStateRequest) SetClientToken(v string) *DetectTerraformStateRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectTerraformStateRequest) SetIdentifier(v string) *DetectTerraformStateRequest {
	s.Identifier = &v
	return s
}

func (s *DetectTerraformStateRequest) SetType(v string) *DetectTerraformStateRequest {
	s.Type = &v
	return s
}

func (s *DetectTerraformStateRequest) Validate() error {
	return dara.Validate(s)
}
