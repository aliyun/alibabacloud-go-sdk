// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceControlConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceControlConfigurationRequest
	GetInstanceId() *string
}

type GetInstanceControlConfigurationRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceControlConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceControlConfigurationRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceControlConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceControlConfigurationRequest) SetInstanceId(v string) *GetInstanceControlConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceControlConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
