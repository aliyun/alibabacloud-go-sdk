// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationAdvancedConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationAdvancedConfigRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationAdvancedConfigRequest
	GetInstanceId() *string
}

type GetApplicationAdvancedConfigRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetApplicationAdvancedConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationAdvancedConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationAdvancedConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationAdvancedConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationAdvancedConfigRequest) SetApplicationId(v string) *GetApplicationAdvancedConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationAdvancedConfigRequest) SetInstanceId(v string) *GetApplicationAdvancedConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationAdvancedConfigRequest) Validate() error {
	return dara.Validate(s)
}
