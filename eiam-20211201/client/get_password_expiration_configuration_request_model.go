// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordExpirationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPasswordExpirationConfigurationRequest
	GetInstanceId() *string
}

type GetPasswordExpirationConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordExpirationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordExpirationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPasswordExpirationConfigurationRequest) SetInstanceId(v string) *GetPasswordExpirationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPasswordExpirationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
