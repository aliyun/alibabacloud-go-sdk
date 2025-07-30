// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForgetPasswordConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetForgetPasswordConfigurationRequest
	GetInstanceId() *string
}

type GetForgetPasswordConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetForgetPasswordConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetForgetPasswordConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetForgetPasswordConfigurationRequest) SetInstanceId(v string) *GetForgetPasswordConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetForgetPasswordConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
