// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListApplicationTokensRequest
	GetApplicationId() *string
	SetApplicationTokenType(v string) *ListApplicationTokensRequest
	GetApplicationTokenType() *string
	SetInstanceId(v string) *ListApplicationTokensRequest
	GetInstanceId() *string
}

type ListApplicationTokensRequest struct {
	// IDaaS的应用资源ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// ApplicationToken的类型
	//
	// This parameter is required.
	//
	// example:
	//
	// bearer_token
	ApplicationTokenType *string `json:"ApplicationTokenType,omitempty" xml:"ApplicationTokenType,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationTokensRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationTokensRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationTokensRequest) GetApplicationTokenType() *string {
	return s.ApplicationTokenType
}

func (s *ListApplicationTokensRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationTokensRequest) SetApplicationId(v string) *ListApplicationTokensRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationTokensRequest) SetApplicationTokenType(v string) *ListApplicationTokensRequest {
	s.ApplicationTokenType = &v
	return s
}

func (s *ListApplicationTokensRequest) SetInstanceId(v string) *ListApplicationTokensRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationTokensRequest) Validate() error {
	return dara.Validate(s)
}
