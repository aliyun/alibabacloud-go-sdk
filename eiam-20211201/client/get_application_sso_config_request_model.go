// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationSsoConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationSsoConfigRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationSsoConfigRequest
	GetInstanceId() *string
}

type GetApplicationSsoConfigRequest struct {
	// The ID of the application.
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

func (s GetApplicationSsoConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationSsoConfigRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationSsoConfigRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationSsoConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationSsoConfigRequest) SetApplicationId(v string) *GetApplicationSsoConfigRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationSsoConfigRequest) SetInstanceId(v string) *GetApplicationSsoConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationSsoConfigRequest) Validate() error {
	return dara.Validate(s)
}
