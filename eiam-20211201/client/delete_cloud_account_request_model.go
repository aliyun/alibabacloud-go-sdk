// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *DeleteCloudAccountRequest
	GetCloudAccountId() *string
	SetInstanceId(v string) *DeleteCloudAccountRequest
	GetInstanceId() *string
}

type DeleteCloudAccountRequest struct {
	// 云账号ID
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

func (s DeleteCloudAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudAccountRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudAccountRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *DeleteCloudAccountRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudAccountRequest) SetCloudAccountId(v string) *DeleteCloudAccountRequest {
	s.CloudAccountId = &v
	return s
}

func (s *DeleteCloudAccountRequest) SetInstanceId(v string) *DeleteCloudAccountRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudAccountRequest) Validate() error {
	return dara.Validate(s)
}
