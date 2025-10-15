// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationSupportedProvisionProtocolTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationSupportedProvisionProtocolTypesRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListApplicationSupportedProvisionProtocolTypesRequest
	GetInstanceId() *string
}

type ListApplicationSupportedProvisionProtocolTypesRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// IDaaS EIAM的实例id
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationSupportedProvisionProtocolTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationSupportedProvisionProtocolTypesRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationSupportedProvisionProtocolTypesRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationSupportedProvisionProtocolTypesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationSupportedProvisionProtocolTypesRequest) SetApplicationId(v string) *ListApplicationSupportedProvisionProtocolTypesRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesRequest) SetInstanceId(v string) *ListApplicationSupportedProvisionProtocolTypesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationSupportedProvisionProtocolTypesRequest) Validate() error {
	return dara.Validate(s)
}
