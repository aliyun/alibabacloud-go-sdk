// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *GetCloudAccountRequest
	GetCloudAccountId() *string
	SetInstanceId(v string) *GetCloudAccountRequest
	GetInstanceId() *string
}

type GetCloudAccountRequest struct {
	// 云账号ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ca_01kmegjc11qa1txxxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAccountRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAccountRequest) SetCloudAccountId(v string) *GetCloudAccountRequest {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountRequest) SetInstanceId(v string) *GetCloudAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
