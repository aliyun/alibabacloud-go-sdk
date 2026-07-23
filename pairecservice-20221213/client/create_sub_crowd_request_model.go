// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateSubCrowdRequest
	GetInstanceId() *string
	SetSource(v string) *CreateSubCrowdRequest
	GetSource() *string
	SetUsers(v string) *CreateSubCrowdRequest
	GetUsers() *string
}

type CreateSubCrowdRequest struct {
	// The instance ID. You can obtain the instance ID by calling the ListInstances API.
	//
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source of users for the sub-crowd. Valid values: ● ManualInput: manual input ● UploadFile: file upload
	//
	// This parameter is required.
	//
	// example:
	//
	// ManualInput
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// A comma-separated list of users in the sub-crowd.
	//
	// This parameter is required.
	//
	// example:
	//
	// user1,user2,user3
	Users *string `json:"Users,omitempty" xml:"Users,omitempty"`
}

func (s CreateSubCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *CreateSubCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSubCrowdRequest) GetSource() *string {
	return s.Source
}

func (s *CreateSubCrowdRequest) GetUsers() *string {
	return s.Users
}

func (s *CreateSubCrowdRequest) SetInstanceId(v string) *CreateSubCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSubCrowdRequest) SetSource(v string) *CreateSubCrowdRequest {
	s.Source = &v
	return s
}

func (s *CreateSubCrowdRequest) SetUsers(v string) *CreateSubCrowdRequest {
	s.Users = &v
	return s
}

func (s *CreateSubCrowdRequest) Validate() error {
	return dara.Validate(s)
}
