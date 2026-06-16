// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCloudAccountDescriptionRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *UpdateCloudAccountDescriptionRequest
	GetCloudAccountId() *string
	SetDescription(v string) *UpdateCloudAccountDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCloudAccountDescriptionRequest
	GetInstanceId() *string
}

type UpdateCloudAccountDescriptionRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique value for this parameter from your client for each request. The ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://www.alibabacloud.com/help/zh/ecs/developer-reference/how-to-ensure-idempotence).
	//
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The description of the Alibaba Cloud account.
	//
	// example:
	//
	// cloud_account_description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudAccountDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountDescriptionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudAccountDescriptionRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *UpdateCloudAccountDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCloudAccountDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudAccountDescriptionRequest) SetClientToken(v string) *UpdateCloudAccountDescriptionRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudAccountDescriptionRequest) SetCloudAccountId(v string) *UpdateCloudAccountDescriptionRequest {
	s.CloudAccountId = &v
	return s
}

func (s *UpdateCloudAccountDescriptionRequest) SetDescription(v string) *UpdateCloudAccountDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateCloudAccountDescriptionRequest) SetInstanceId(v string) *UpdateCloudAccountDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudAccountDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
