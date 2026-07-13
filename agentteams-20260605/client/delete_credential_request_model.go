// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteCredentialRequest
	GetClientToken() *string
	SetInstanceId(v string) *DeleteCredentialRequest
	GetInstanceId() *string
	SetName(v string) *DeleteCredentialRequest
	GetName() *string
}

type DeleteCredentialRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCredentialRequest) GetName() *string {
	return s.Name
}

func (s *DeleteCredentialRequest) SetClientToken(v string) *DeleteCredentialRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCredentialRequest) SetInstanceId(v string) *DeleteCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCredentialRequest) SetName(v string) *DeleteCredentialRequest {
	s.Name = &v
	return s
}

func (s *DeleteCredentialRequest) Validate() error {
	return dara.Validate(s)
}
