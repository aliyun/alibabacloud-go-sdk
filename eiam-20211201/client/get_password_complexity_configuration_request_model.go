// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordComplexityConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetPasswordComplexityConfigurationRequest
	GetInstanceId() *string
}

type GetPasswordComplexityConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetPasswordComplexityConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordComplexityConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetPasswordComplexityConfigurationRequest) SetInstanceId(v string) *GetPasswordComplexityConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetPasswordComplexityConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
