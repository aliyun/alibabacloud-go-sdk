// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAccountResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAccountId(v string) *GetCloudAccountResourceRequest
	GetCloudAccountId() *string
	SetCloudAccountResourceId(v string) *GetCloudAccountResourceRequest
	GetCloudAccountResourceId() *string
	SetInstanceId(v string) *GetCloudAccountResourceRequest
	GetInstanceId() *string
}

type GetCloudAccountResourceRequest struct {
	// The ID of the cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789012xxxx
	CloudAccountId *string `json:"CloudAccountId,omitempty" xml:"CloudAccountId,omitempty"`
	// The ID of the cloud account resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// car_xxxxxxxxxxxxxxxxxxxxxxxxxxxx
	CloudAccountResourceId *string `json:"CloudAccountResourceId,omitempty" xml:"CloudAccountResourceId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCloudAccountResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAccountResourceRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAccountResourceRequest) GetCloudAccountId() *string {
	return s.CloudAccountId
}

func (s *GetCloudAccountResourceRequest) GetCloudAccountResourceId() *string {
	return s.CloudAccountResourceId
}

func (s *GetCloudAccountResourceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAccountResourceRequest) SetCloudAccountId(v string) *GetCloudAccountResourceRequest {
	s.CloudAccountId = &v
	return s
}

func (s *GetCloudAccountResourceRequest) SetCloudAccountResourceId(v string) *GetCloudAccountResourceRequest {
	s.CloudAccountResourceId = &v
	return s
}

func (s *GetCloudAccountResourceRequest) SetInstanceId(v string) *GetCloudAccountResourceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAccountResourceRequest) Validate() error {
	return dara.Validate(s)
}
