// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectNotifyReceiverRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEmail(v string) *CreateProjectNotifyReceiverRequest
	GetEmail() *string
	SetInstanceId(v string) *CreateProjectNotifyReceiverRequest
	GetInstanceId() *string
	SetName(v string) *CreateProjectNotifyReceiverRequest
	GetName() *string
	SetPhone(v string) *CreateProjectNotifyReceiverRequest
	GetPhone() *string
	SetProjectId(v int64) *CreateProjectNotifyReceiverRequest
	GetProjectId() *int64
	SetRegionId(v string) *CreateProjectNotifyReceiverRequest
	GetRegionId() *string
}

type CreateProjectNotifyReceiverRequest struct {
	// The email address of the recipient.
	//
	// > You must specify at least one of the mobile phone number and the email address.
	//
	// example:
	//
	// 1099**@qq.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can call the [DescribeInstances](https://help.aliyun.com/document_detail/462953.html) operation to retrieve this parameter.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the recipient. The name must be 1 to 128 characters in length. It cannot start with a special character and can contain only the following special characters: periods (.), underscores (_), hyphens (-), backslashes (\\), and spaces.
	//
	// This parameter is required.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The mobile phone number of the recipient.
	//
	// > You must specify at least one of the mobile phone number and the email address.
	//
	// example:
	//
	// 131****1421
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The project ID.
	//
	// > You can call the ListProjects operation to retrieve this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the Bastionhost instance.
	//
	// > For the mappings between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateProjectNotifyReceiverRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectNotifyReceiverRequest) GoString() string {
	return s.String()
}

func (s *CreateProjectNotifyReceiverRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateProjectNotifyReceiverRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateProjectNotifyReceiverRequest) GetName() *string {
	return s.Name
}

func (s *CreateProjectNotifyReceiverRequest) GetPhone() *string {
	return s.Phone
}

func (s *CreateProjectNotifyReceiverRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateProjectNotifyReceiverRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateProjectNotifyReceiverRequest) SetEmail(v string) *CreateProjectNotifyReceiverRequest {
	s.Email = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) SetInstanceId(v string) *CreateProjectNotifyReceiverRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) SetName(v string) *CreateProjectNotifyReceiverRequest {
	s.Name = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) SetPhone(v string) *CreateProjectNotifyReceiverRequest {
	s.Phone = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) SetProjectId(v int64) *CreateProjectNotifyReceiverRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) SetRegionId(v string) *CreateProjectNotifyReceiverRequest {
	s.RegionId = &v
	return s
}

func (s *CreateProjectNotifyReceiverRequest) Validate() error {
	return dara.Validate(s)
}
