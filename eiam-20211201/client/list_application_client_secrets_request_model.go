// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationClientSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationClientSecretsRequest
	GetApplicationId() *string
	SetInstanceId(v string) *ListApplicationClientSecretsRequest
	GetInstanceId() *string
}

type ListApplicationClientSecretsRequest struct {
	// The ID of the application that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationClientSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationClientSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationClientSecretsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationClientSecretsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationClientSecretsRequest) SetApplicationId(v string) *ListApplicationClientSecretsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationClientSecretsRequest) SetInstanceId(v string) *ListApplicationClientSecretsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationClientSecretsRequest) Validate() error {
	return dara.Validate(s)
}
