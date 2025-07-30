// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordInitializationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPasswordInitializationConfigurationRequest
	GetInstanceId() *string
}

type GetPasswordInitializationConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordInitializationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordInitializationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPasswordInitializationConfigurationRequest) SetInstanceId(v string) *GetPasswordInitializationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPasswordInitializationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
