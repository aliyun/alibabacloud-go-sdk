// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateCloudAccountRequest
	GetClientToken() *string
	SetCloudAccountId(v string) *UpdateCloudAccountRequest
	GetCloudAccountId() *string
	SetCloudAccountName(v string) *UpdateCloudAccountRequest
	GetCloudAccountName() *string
	SetInstanceId(v string) *UpdateCloudAccountRequest
	GetInstanceId() *string
}

type UpdateCloudAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// 云账号ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cloud_account_xxx
	CloudAccountName *string `json:"CloudAccountName,omitempty" xml:"CloudAccountName,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudAccountRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *UpdateCloudAccountRequest) GetCloudAccountName() *string {
	return s.CloudAccountName
}

func (s *UpdateCloudAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudAccountRequest) SetClientToken(v string) *UpdateCloudAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudAccountRequest) SetCloudAccountId(v string) *UpdateCloudAccountRequest {
	s.CloudAccountId = &v
	return s
}

func (s *UpdateCloudAccountRequest) SetCloudAccountName(v string) *UpdateCloudAccountRequest {
	s.CloudAccountName = &v
	return s
}

func (s *UpdateCloudAccountRequest) SetInstanceId(v string) *UpdateCloudAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
