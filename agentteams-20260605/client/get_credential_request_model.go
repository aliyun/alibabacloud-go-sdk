// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetCredentialRequest
	GetInstanceId() *string
	SetName(v string) *GetCredentialRequest
	GetName() *string
}

type GetCredentialRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCredentialRequest) GetName() *string {
	return s.Name
}

func (s *GetCredentialRequest) SetInstanceId(v string) *GetCredentialRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCredentialRequest) SetName(v string) *GetCredentialRequest {
	s.Name = &v
	return s
}

func (s *GetCredentialRequest) Validate() error {
	return dara.Validate(s)
}
