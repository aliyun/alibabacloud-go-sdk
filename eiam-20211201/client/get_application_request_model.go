// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *GetApplicationRequest
	GetApplicationId() *string
	SetInstanceId(v string) *GetApplicationRequest
	GetInstanceId() *string
}

type GetApplicationRequest struct {
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

func (s GetApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *GetApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetApplicationRequest) SetApplicationId(v string) *GetApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *GetApplicationRequest) SetInstanceId(v string) *GetApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *GetApplicationRequest) Validate() error {
	return dara.Validate(s)
}
