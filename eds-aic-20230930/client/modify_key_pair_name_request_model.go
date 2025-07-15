// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyKeyPairNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairId(v string) *ModifyKeyPairNameRequest
	GetKeyPairId() *string
	SetNewKeyPairName(v string) *ModifyKeyPairNameRequest
	GetNewKeyPairName() *string
}

type ModifyKeyPairNameRequest struct {
	// The ID of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The name of the ADB key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// newKeyPairName
	NewKeyPairName *string `json:"NewKeyPairName,omitempty" xml:"NewKeyPairName,omitempty"`
}

func (s ModifyKeyPairNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyKeyPairNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyKeyPairNameRequest) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *ModifyKeyPairNameRequest) GetNewKeyPairName() *string {
	return s.NewKeyPairName
}

func (s *ModifyKeyPairNameRequest) SetKeyPairId(v string) *ModifyKeyPairNameRequest {
	s.KeyPairId = &v
	return s
}

func (s *ModifyKeyPairNameRequest) SetNewKeyPairName(v string) *ModifyKeyPairNameRequest {
	s.NewKeyPairName = &v
	return s
}

func (s *ModifyKeyPairNameRequest) Validate() error {
	return dara.Validate(s)
}
